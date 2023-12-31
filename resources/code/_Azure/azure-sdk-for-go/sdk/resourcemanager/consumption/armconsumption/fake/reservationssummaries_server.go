//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"net/http"
	"net/url"
	"regexp"
)

// ReservationsSummariesServer is a fake server for instances of the armconsumption.ReservationsSummariesClient type.
type ReservationsSummariesServer struct {
	// NewListPager is the fake for method ReservationsSummariesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceScope string, grain armconsumption.Datagrain, options *armconsumption.ReservationsSummariesClientListOptions) (resp azfake.PagerResponder[armconsumption.ReservationsSummariesClientListResponse])

	// NewListByReservationOrderPager is the fake for method ReservationsSummariesClient.NewListByReservationOrderPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByReservationOrderPager func(reservationOrderID string, grain armconsumption.Datagrain, options *armconsumption.ReservationsSummariesClientListByReservationOrderOptions) (resp azfake.PagerResponder[armconsumption.ReservationsSummariesClientListByReservationOrderResponse])

	// NewListByReservationOrderAndReservationPager is the fake for method ReservationsSummariesClient.NewListByReservationOrderAndReservationPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByReservationOrderAndReservationPager func(reservationOrderID string, reservationID string, grain armconsumption.Datagrain, options *armconsumption.ReservationsSummariesClientListByReservationOrderAndReservationOptions) (resp azfake.PagerResponder[armconsumption.ReservationsSummariesClientListByReservationOrderAndReservationResponse])
}

// NewReservationsSummariesServerTransport creates a new instance of ReservationsSummariesServerTransport with the provided implementation.
// The returned ReservationsSummariesServerTransport instance is connected to an instance of armconsumption.ReservationsSummariesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReservationsSummariesServerTransport(srv *ReservationsSummariesServer) *ReservationsSummariesServerTransport {
	return &ReservationsSummariesServerTransport{
		srv:                            srv,
		newListPager:                   newTracker[azfake.PagerResponder[armconsumption.ReservationsSummariesClientListResponse]](),
		newListByReservationOrderPager: newTracker[azfake.PagerResponder[armconsumption.ReservationsSummariesClientListByReservationOrderResponse]](),
		newListByReservationOrderAndReservationPager: newTracker[azfake.PagerResponder[armconsumption.ReservationsSummariesClientListByReservationOrderAndReservationResponse]](),
	}
}

// ReservationsSummariesServerTransport connects instances of armconsumption.ReservationsSummariesClient to instances of ReservationsSummariesServer.
// Don't use this type directly, use NewReservationsSummariesServerTransport instead.
type ReservationsSummariesServerTransport struct {
	srv                                          *ReservationsSummariesServer
	newListPager                                 *tracker[azfake.PagerResponder[armconsumption.ReservationsSummariesClientListResponse]]
	newListByReservationOrderPager               *tracker[azfake.PagerResponder[armconsumption.ReservationsSummariesClientListByReservationOrderResponse]]
	newListByReservationOrderAndReservationPager *tracker[azfake.PagerResponder[armconsumption.ReservationsSummariesClientListByReservationOrderAndReservationResponse]]
}

// Do implements the policy.Transporter interface for ReservationsSummariesServerTransport.
func (r *ReservationsSummariesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ReservationsSummariesClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	case "ReservationsSummariesClient.NewListByReservationOrderPager":
		resp, err = r.dispatchNewListByReservationOrderPager(req)
	case "ReservationsSummariesClient.NewListByReservationOrderAndReservationPager":
		resp, err = r.dispatchNewListByReservationOrderAndReservationPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ReservationsSummariesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/(?P<resourceScope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Consumption/reservationSummaries`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceScopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceScope")])
		if err != nil {
			return nil, err
		}
		grainParam, err := parseWithCast(qp.Get("grain"), func(v string) (armconsumption.Datagrain, error) {
			p, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armconsumption.Datagrain(p), nil
		})
		if err != nil {
			return nil, err
		}
		startDateUnescaped, err := url.QueryUnescape(qp.Get("startDate"))
		if err != nil {
			return nil, err
		}
		startDateParam := getOptional(startDateUnescaped)
		endDateUnescaped, err := url.QueryUnescape(qp.Get("endDate"))
		if err != nil {
			return nil, err
		}
		endDateParam := getOptional(endDateUnescaped)
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		reservationIDUnescaped, err := url.QueryUnescape(qp.Get("reservationId"))
		if err != nil {
			return nil, err
		}
		reservationIDParam := getOptional(reservationIDUnescaped)
		reservationOrderIDUnescaped, err := url.QueryUnescape(qp.Get("reservationOrderId"))
		if err != nil {
			return nil, err
		}
		reservationOrderIDParam := getOptional(reservationOrderIDUnescaped)
		var options *armconsumption.ReservationsSummariesClientListOptions
		if startDateParam != nil || endDateParam != nil || filterParam != nil || reservationIDParam != nil || reservationOrderIDParam != nil {
			options = &armconsumption.ReservationsSummariesClientListOptions{
				StartDate:          startDateParam,
				EndDate:            endDateParam,
				Filter:             filterParam,
				ReservationID:      reservationIDParam,
				ReservationOrderID: reservationOrderIDParam,
			}
		}
		resp := r.srv.NewListPager(resourceScopeParam, grainParam, options)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armconsumption.ReservationsSummariesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}

func (r *ReservationsSummariesServerTransport) dispatchNewListByReservationOrderPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByReservationOrderPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByReservationOrderPager not implemented")}
	}
	newListByReservationOrderPager := r.newListByReservationOrderPager.get(req)
	if newListByReservationOrderPager == nil {
		const regexStr = `/providers/Microsoft\.Capacity/reservationorders/(?P<reservationOrderId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Consumption/reservationSummaries`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		reservationOrderIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reservationOrderId")])
		if err != nil {
			return nil, err
		}
		grainParam, err := parseWithCast(qp.Get("grain"), func(v string) (armconsumption.Datagrain, error) {
			p, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armconsumption.Datagrain(p), nil
		})
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armconsumption.ReservationsSummariesClientListByReservationOrderOptions
		if filterParam != nil {
			options = &armconsumption.ReservationsSummariesClientListByReservationOrderOptions{
				Filter: filterParam,
			}
		}
		resp := r.srv.NewListByReservationOrderPager(reservationOrderIDParam, grainParam, options)
		newListByReservationOrderPager = &resp
		r.newListByReservationOrderPager.add(req, newListByReservationOrderPager)
		server.PagerResponderInjectNextLinks(newListByReservationOrderPager, req, func(page *armconsumption.ReservationsSummariesClientListByReservationOrderResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByReservationOrderPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByReservationOrderPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByReservationOrderPager) {
		r.newListByReservationOrderPager.remove(req)
	}
	return resp, nil
}

func (r *ReservationsSummariesServerTransport) dispatchNewListByReservationOrderAndReservationPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByReservationOrderAndReservationPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByReservationOrderAndReservationPager not implemented")}
	}
	newListByReservationOrderAndReservationPager := r.newListByReservationOrderAndReservationPager.get(req)
	if newListByReservationOrderAndReservationPager == nil {
		const regexStr = `/providers/Microsoft\.Capacity/reservationorders/(?P<reservationOrderId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reservations/(?P<reservationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Consumption/reservationSummaries`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		reservationOrderIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reservationOrderId")])
		if err != nil {
			return nil, err
		}
		reservationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reservationId")])
		if err != nil {
			return nil, err
		}
		grainParam, err := parseWithCast(qp.Get("grain"), func(v string) (armconsumption.Datagrain, error) {
			p, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armconsumption.Datagrain(p), nil
		})
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armconsumption.ReservationsSummariesClientListByReservationOrderAndReservationOptions
		if filterParam != nil {
			options = &armconsumption.ReservationsSummariesClientListByReservationOrderAndReservationOptions{
				Filter: filterParam,
			}
		}
		resp := r.srv.NewListByReservationOrderAndReservationPager(reservationOrderIDParam, reservationIDParam, grainParam, options)
		newListByReservationOrderAndReservationPager = &resp
		r.newListByReservationOrderAndReservationPager.add(req, newListByReservationOrderAndReservationPager)
		server.PagerResponderInjectNextLinks(newListByReservationOrderAndReservationPager, req, func(page *armconsumption.ReservationsSummariesClientListByReservationOrderAndReservationResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByReservationOrderAndReservationPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByReservationOrderAndReservationPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByReservationOrderAndReservationPager) {
		r.newListByReservationOrderAndReservationPager.remove(req)
	}
	return resp, nil
}
