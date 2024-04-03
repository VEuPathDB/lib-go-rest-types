package veupath

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/Foxcapades/Go-ChainRequest/simple"
	"github.com/vulpine-io/httpx/v1/pkg/httpx/header"
)

const (
	wdkServicePath = "/%s/service"

	wdkUsersPath = "/users"

	// "/users/{userId}"
	wdkUserByIDPath = wdkUsersPath + "/%s"

	// "/users/{userId}/strategies"
	wdkStrategiesPath = wdkUserByIDPath + "/strategies"

	// "/users/{userId}/strategies/{strategyId}"
	wdkStrategyByIDPath = wdkStrategiesPath + "/%s"

	wdkPublicStrategyListPath = "/strategy-lists/public"

	wdkRecordTypeListPath = "/record-types"

	wdkRecordTypeIDPath = wdkRecordTypeListPath + "/%s"

	// "/record-types/{record-type}/searches"
	wdkRecordTypeSearchPath = wdkRecordTypeIDPath + "/searches"

	// "/record-types/{record-type}/searches/{search}"
	wdkNamedSearchPath = wdkRecordTypeSearchPath + "/%s"

	// "/record-types/{record-type}/searches/{search}/reports/{report-name}"
	wdkAnswerPath = wdkNamedSearchPath + "/reports/%s"

	wdkStdReportName = "standard"

	authTktParam = "auth_tkt"

	currentUserID = "current"
)

// NewWDKURLBuilder attempts to create a WDKURLBuilder instance from the given
// URL string.
//
// This function will make one request at minimum to the target URL to validate
// it before returning a constructed WDKURLBuilder instance.
//
// If the given URL is prefixed with the protocol indicator 'http://' this
// function will test for a required https upgrade.
//
// If the given URL includes no protocol indicator, https will be assumed.
//
// If the given URL contains no path elements, a request will be made to the
// base URL to determine the name of the WDK app.
//
// If the URL is malformed or in an unexpected format, or if any of the HTTP
// requests to the target service fail, this function will return an error.
func NewWDKURLBuilder(url string) (WDKURLBuilder, error) {
	rgx := regexp.MustCompile(`^(https?://)?([^/]+)((?:/[^/]+)*)/?$`)
	res := rgx.FindStringSubmatch(url)

	if len(res) == 0 {
		return WDKURLBuilder{}, fmt.Errorf("malformed or unrecognized URL: %s", url)
	}

	prot := res[1]
	base := res[2]
	path := res[3]

	if len(prot) == 0 {
		prot = "https://"
	} else if prot == "http://" {
		// make request to test for upgrade to https.
		//
		// If an upgrade is required, site should respond with 301 with the correct
		// url in the Location header.
		response := simple.HeadRequest(prot + base).
			DisableRedirects().
			Submit()

		// If there was a request error.
		if response.GetError() != nil {
			return WDKURLBuilder{}, fmt.Errorf("failed to test for https upgrade: %s", response.GetError())
		}

		// There shouldn't be a body with a head request, but just to be safe.
		_ = response.Close()

		// If the server responded with a 301, get the correct location.
		if response.MustGetResponseCode() == 301 {
			if loc, ok, _ := response.LookupHeader(header.Location); !ok {
				return WDKURLBuilder{}, errors.New("invalid HTTP response while testing for https upgrade, server responded with 301 but no Location header")
			} else {
				res = rgx.FindStringSubmatch(loc)

				prot = res[1]
				base = res[2]
			}
		}

		// If the server responded with a 302, then the URL is correct, but the path
		// isn't, parse the correct path to get the app name.
		if response.MustGetResponseCode() == 302 {
			if loc, ok, _ := response.LookupHeader(header.Location); !ok {
				return WDKURLBuilder{}, errors.New("invalid HTTP response while testing for https upgrade, server responded with 302 but no Location header")
			} else {
				res = rgx.FindStringSubmatch(loc)

				prot = res[1]
				base = res[2]
				path = res[3]
			}
		}
	}

	if len(path) == 0 {
		// make request to get redirect for app name.  site should respond with a
		// 302 with the app location (e.g. https://plasmodb.org/plasmo/)
		response := simple.HeadRequest(prot + base).
			DisableRedirects().
			Submit()

		if response.GetError() != nil {
			return WDKURLBuilder{}, fmt.Errorf("failed to test for app name: %s", response.GetError())
		}

		// There shouldn't be a response body with a HEAD request, but just to be
		// safe.
		_ = response.Close()

		if response.MustGetResponseCode() == 302 {
			if loc, ok, _ := response.LookupHeader(header.Location); !ok {
				return WDKURLBuilder{}, errors.New("invalid HTTP response while fetching app name, server responded with 302 but no Location header")
			} else {
				res = rgx.FindStringSubmatch(loc)
				prot = res[1]
				base = res[2]
				path = res[3]
			}
		}
	}

	if len(path) == 0 {
		return WDKURLBuilder{}, fmt.Errorf("failed to resolve app name for given url: %s", url)
	}

	// split path and use the first segment as the app name
	idx := strings.IndexByte(path[1:], '/')
	if idx == -1 {
		path = path[1:]
	} else {
		path = path[1 : idx+1]
	}

	return WDKURLBuilder{url: prot + base + fmt.Sprintf(wdkServicePath, path)}, nil
}

type WDKURLBuilder struct {
	url string
	tkt string
}

func (w WDKURLBuilder) CurrentUserStrategiesURL() string {
	return w.appendAuthTkt(w.url + fmt.Sprintf(wdkStrategiesPath, currentUserID))
}

func (w WDKURLBuilder) CurrentUserURL() string {
	return w.appendAuthTkt(w.url + fmt.Sprintf(wdkUserByIDPath, currentUserID))
}

func (w WDKURLBuilder) PublicStrategiesURL() string {
	return w.appendAuthTkt(w.url + wdkPublicStrategyListPath)
}

func (w WDKURLBuilder) RecordSearchListURL(recordType string) string {
	return w.appendAuthTkt(w.url + fmt.Sprintf(wdkRecordTypeSearchPath, recordType))
}

func (w WDKURLBuilder) RecordSearchStdReportURL(recordType, search string) string {
	return w.appendAuthTkt(w.url + fmt.Sprintf(wdkAnswerPath, recordType, search, wdkStdReportName))
}

func (w WDKURLBuilder) RecordSearchURL(recordType, search string) string {
	return w.appendAuthTkt(w.url + fmt.Sprintf(wdkNamedSearchPath, recordType, search))
}

func (w WDKURLBuilder) RecordTypeListURL() string {
	return w.appendAuthTkt(w.url + wdkRecordTypeListPath)
}

func (w WDKURLBuilder) RecordTypeURL(recordType string) string {
	return w.appendAuthTkt(w.url + fmt.Sprintf(wdkRecordTypeIDPath, recordType))
}

func (w WDKURLBuilder) ServiceURL() string {
	return w.url
}

func (w *WDKURLBuilder) SetAuthTkt(auth string) {
	if len(auth) == 0 {
		w.tkt = ""
	} else {
		w.tkt = "?" + authTktParam + "=" + auth
	}
}

func (w WDKURLBuilder) appendAuthTkt(url string) string {
	if len(w.tkt) > 0 {
		return url + w.tkt
	}

	return url
}
