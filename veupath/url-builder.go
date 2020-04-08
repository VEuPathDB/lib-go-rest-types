package veupath

import (
	"fmt"
	"strings"
)

const (
	ubServiceUrl = "service/"

	// Record-type API
	ubRecordTypeListUrl   = ubServiceUrl + "record-types/"
	ubRecordTypeUrl       = ubRecordTypeListUrl + "%s/"
	ubRecordSearchListUrl = ubRecordTypeUrl + "searches/"
	ubRecordSearchUrl     = ubRecordSearchListUrl + "%s/"
	ubRecordReportListUrl = ubRecordSearchUrl + "reports/"
	ubRecordReportUrl     = ubRecordReportListUrl + "%s/"

	ubStratLists   = ubServiceUrl + "strategy-lists/"
	ubPublicStrats = ubStratLists + "public/"
)

type ApiUrlBuilder struct {
	baseUrl string
	auth    string
}

func NewApiUrlBuilder(baseUrl string) ApiUrlBuilder {
	if strings.HasSuffix(baseUrl, "/") {
		return ApiUrlBuilder{baseUrl: baseUrl}
	} else {
		return ApiUrlBuilder{baseUrl: baseUrl + "/"}
	}
}

func (a *ApiUrlBuilder) SetAuthTkt(auth string) {
	if len(auth) > 0 {
		a.auth = "?auth_tkt=" + auth
	} else {
		a.auth = ""
	}
}

func (a *ApiUrlBuilder) ServiceUrl() string {
	return a.baseUrl + ubServiceUrl + a.auth
}

func (a *ApiUrlBuilder) RecordTypeListUrl() string {
	return a.baseUrl + ubRecordTypeListUrl + a.auth
}

func (a *ApiUrlBuilder) RecordTypeUrl(recType string) string {
	return a.baseUrl + fmt.Sprintf(ubRecordTypeUrl, recType) + a.auth
}

func (a *ApiUrlBuilder) RecordSearchListUrl(recType string) string {
	return a.baseUrl + fmt.Sprintf(ubRecordSearchListUrl, recType) + a.auth
}

func (a *ApiUrlBuilder) RecordSearchUrl(recType, search string) string {
	return a.baseUrl + fmt.Sprintf(ubRecordSearchUrl, recType, search) + a.auth
}

func (a *ApiUrlBuilder) RecordSearchStdReportUrl(recType, search string) string {
	return a.baseUrl + fmt.Sprintf(ubRecordReportUrl, recType, search, "standard") + a.auth
}

func (a *ApiUrlBuilder) PublicStrategiesUrl() string {
	return a.baseUrl + ubPublicStrats
}
