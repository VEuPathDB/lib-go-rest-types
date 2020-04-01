package service

import (
	"fmt"
	"strings"
)

const (
	ubServiceUrl          = "service/"
	ubRecordTypeListUrl   = ubServiceUrl + "record-types/"
	ubRecordTypeUrl       = ubRecordTypeListUrl + "%s/"
	ubRecordSearchListUrl = ubRecordTypeUrl + "searches/"
	ubRecordSearchUrl     = ubRecordSearchListUrl + "%s/"
)

type ApiUrlBuilder struct {
	baseUrl string
	auth string
}

func NewApiUrlBuilder(baseUrl string) ApiUrlBuilder {
	if strings.HasSuffix(baseUrl, "/") {
		return ApiUrlBuilder{baseUrl: baseUrl}
	} else {
		return ApiUrlBuilder{baseUrl: baseUrl + "/"}
	}
}

func (a *ApiUrlBuilder) SetAuthTkt(auth string) {
	a.auth = "?auth_tkt=" + auth
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
