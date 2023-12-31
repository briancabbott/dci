//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
	"net/http"
	"net/url"
	"regexp"
)

// GenerateReservationDetailsReportServer is a fake server for instances of the armcostmanagement.GenerateReservationDetailsReportClient type.
type GenerateReservationDetailsReportServer struct {
	// BeginByBillingAccountID is the fake for method GenerateReservationDetailsReportClient.BeginByBillingAccountID
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginByBillingAccountID func(ctx context.Context, billingAccountID string, startDate string, endDate string, options *armcostmanagement.GenerateReservationDetailsReportClientBeginByBillingAccountIDOptions) (resp azfake.PollerResponder[armcostmanagement.GenerateReservationDetailsReportClientByBillingAccountIDResponse], errResp azfake.ErrorResponder)

	// BeginByBillingProfileID is the fake for method GenerateReservationDetailsReportClient.BeginByBillingProfileID
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginByBillingProfileID func(ctx context.Context, billingAccountID string, billingProfileID string, startDate string, endDate string, options *armcostmanagement.GenerateReservationDetailsReportClientBeginByBillingProfileIDOptions) (resp azfake.PollerResponder[armcostmanagement.GenerateReservationDetailsReportClientByBillingProfileIDResponse], errResp azfake.ErrorResponder)
}

// NewGenerateReservationDetailsReportServerTransport creates a new instance of GenerateReservationDetailsReportServerTransport with the provided implementation.
// The returned GenerateReservationDetailsReportServerTransport instance is connected to an instance of armcostmanagement.GenerateReservationDetailsReportClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGenerateReservationDetailsReportServerTransport(srv *GenerateReservationDetailsReportServer) *GenerateReservationDetailsReportServerTransport {
	return &GenerateReservationDetailsReportServerTransport{
		srv:                     srv,
		beginByBillingAccountID: newTracker[azfake.PollerResponder[armcostmanagement.GenerateReservationDetailsReportClientByBillingAccountIDResponse]](),
		beginByBillingProfileID: newTracker[azfake.PollerResponder[armcostmanagement.GenerateReservationDetailsReportClientByBillingProfileIDResponse]](),
	}
}

// GenerateReservationDetailsReportServerTransport connects instances of armcostmanagement.GenerateReservationDetailsReportClient to instances of GenerateReservationDetailsReportServer.
// Don't use this type directly, use NewGenerateReservationDetailsReportServerTransport instead.
type GenerateReservationDetailsReportServerTransport struct {
	srv                     *GenerateReservationDetailsReportServer
	beginByBillingAccountID *tracker[azfake.PollerResponder[armcostmanagement.GenerateReservationDetailsReportClientByBillingAccountIDResponse]]
	beginByBillingProfileID *tracker[azfake.PollerResponder[armcostmanagement.GenerateReservationDetailsReportClientByBillingProfileIDResponse]]
}

// Do implements the policy.Transporter interface for GenerateReservationDetailsReportServerTransport.
func (g *GenerateReservationDetailsReportServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GenerateReservationDetailsReportClient.BeginByBillingAccountID":
		resp, err = g.dispatchBeginByBillingAccountID(req)
	case "GenerateReservationDetailsReportClient.BeginByBillingProfileID":
		resp, err = g.dispatchBeginByBillingProfileID(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GenerateReservationDetailsReportServerTransport) dispatchBeginByBillingAccountID(req *http.Request) (*http.Response, error) {
	if g.srv.BeginByBillingAccountID == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginByBillingAccountID not implemented")}
	}
	beginByBillingAccountID := g.beginByBillingAccountID.get(req)
	if beginByBillingAccountID == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/generateReservationDetailsReport`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		billingAccountIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountId")])
		if err != nil {
			return nil, err
		}
		startDateParam, err := url.QueryUnescape(qp.Get("startDate"))
		if err != nil {
			return nil, err
		}
		endDateParam, err := url.QueryUnescape(qp.Get("endDate"))
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginByBillingAccountID(req.Context(), billingAccountIDParam, startDateParam, endDateParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginByBillingAccountID = &respr
		g.beginByBillingAccountID.add(req, beginByBillingAccountID)
	}

	resp, err := server.PollerResponderNext(beginByBillingAccountID, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		g.beginByBillingAccountID.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginByBillingAccountID) {
		g.beginByBillingAccountID.remove(req)
	}

	return resp, nil
}

func (g *GenerateReservationDetailsReportServerTransport) dispatchBeginByBillingProfileID(req *http.Request) (*http.Response, error) {
	if g.srv.BeginByBillingProfileID == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginByBillingProfileID not implemented")}
	}
	beginByBillingProfileID := g.beginByBillingProfileID.get(req)
	if beginByBillingProfileID == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/generateReservationDetailsReport`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		billingAccountIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountId")])
		if err != nil {
			return nil, err
		}
		billingProfileIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileId")])
		if err != nil {
			return nil, err
		}
		startDateParam, err := url.QueryUnescape(qp.Get("startDate"))
		if err != nil {
			return nil, err
		}
		endDateParam, err := url.QueryUnescape(qp.Get("endDate"))
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginByBillingProfileID(req.Context(), billingAccountIDParam, billingProfileIDParam, startDateParam, endDateParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginByBillingProfileID = &respr
		g.beginByBillingProfileID.add(req, beginByBillingProfileID)
	}

	resp, err := server.PollerResponderNext(beginByBillingProfileID, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		g.beginByBillingProfileID.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginByBillingProfileID) {
		g.beginByBillingProfileID.remove(req)
	}

	return resp, nil
}
