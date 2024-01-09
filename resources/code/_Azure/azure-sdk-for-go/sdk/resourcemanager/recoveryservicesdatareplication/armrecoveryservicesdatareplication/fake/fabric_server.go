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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
	"net/http"
	"net/url"
	"regexp"
)

// FabricServer is a fake server for instances of the armrecoveryservicesdatareplication.FabricClient type.
type FabricServer struct {
	// BeginCreate is the fake for method FabricClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, fabricName string, body armrecoveryservicesdatareplication.FabricModel, options *armrecoveryservicesdatareplication.FabricClientBeginCreateOptions) (resp azfake.PollerResponder[armrecoveryservicesdatareplication.FabricClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method FabricClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, fabricName string, options *armrecoveryservicesdatareplication.FabricClientBeginDeleteOptions) (resp azfake.PollerResponder[armrecoveryservicesdatareplication.FabricClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FabricClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, fabricName string, options *armrecoveryservicesdatareplication.FabricClientGetOptions) (resp azfake.Responder[armrecoveryservicesdatareplication.FabricClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method FabricClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armrecoveryservicesdatareplication.FabricClientListOptions) (resp azfake.PagerResponder[armrecoveryservicesdatareplication.FabricClientListResponse])

	// NewListBySubscriptionPager is the fake for method FabricClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armrecoveryservicesdatareplication.FabricClientListBySubscriptionOptions) (resp azfake.PagerResponder[armrecoveryservicesdatareplication.FabricClientListBySubscriptionResponse])

	// BeginUpdate is the fake for method FabricClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, fabricName string, body armrecoveryservicesdatareplication.FabricModelUpdate, options *armrecoveryservicesdatareplication.FabricClientBeginUpdateOptions) (resp azfake.PollerResponder[armrecoveryservicesdatareplication.FabricClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewFabricServerTransport creates a new instance of FabricServerTransport with the provided implementation.
// The returned FabricServerTransport instance is connected to an instance of armrecoveryservicesdatareplication.FabricClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFabricServerTransport(srv *FabricServer) *FabricServerTransport {
	return &FabricServerTransport{
		srv:                        srv,
		beginCreate:                newTracker[azfake.PollerResponder[armrecoveryservicesdatareplication.FabricClientCreateResponse]](),
		beginDelete:                newTracker[azfake.PollerResponder[armrecoveryservicesdatareplication.FabricClientDeleteResponse]](),
		newListPager:               newTracker[azfake.PagerResponder[armrecoveryservicesdatareplication.FabricClientListResponse]](),
		newListBySubscriptionPager: newTracker[azfake.PagerResponder[armrecoveryservicesdatareplication.FabricClientListBySubscriptionResponse]](),
		beginUpdate:                newTracker[azfake.PollerResponder[armrecoveryservicesdatareplication.FabricClientUpdateResponse]](),
	}
}

// FabricServerTransport connects instances of armrecoveryservicesdatareplication.FabricClient to instances of FabricServer.
// Don't use this type directly, use NewFabricServerTransport instead.
type FabricServerTransport struct {
	srv                        *FabricServer
	beginCreate                *tracker[azfake.PollerResponder[armrecoveryservicesdatareplication.FabricClientCreateResponse]]
	beginDelete                *tracker[azfake.PollerResponder[armrecoveryservicesdatareplication.FabricClientDeleteResponse]]
	newListPager               *tracker[azfake.PagerResponder[armrecoveryservicesdatareplication.FabricClientListResponse]]
	newListBySubscriptionPager *tracker[azfake.PagerResponder[armrecoveryservicesdatareplication.FabricClientListBySubscriptionResponse]]
	beginUpdate                *tracker[azfake.PollerResponder[armrecoveryservicesdatareplication.FabricClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for FabricServerTransport.
func (f *FabricServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FabricClient.BeginCreate":
		resp, err = f.dispatchBeginCreate(req)
	case "FabricClient.BeginDelete":
		resp, err = f.dispatchBeginDelete(req)
	case "FabricClient.Get":
		resp, err = f.dispatchGet(req)
	case "FabricClient.NewListPager":
		resp, err = f.dispatchNewListPager(req)
	case "FabricClient.NewListBySubscriptionPager":
		resp, err = f.dispatchNewListBySubscriptionPager(req)
	case "FabricClient.BeginUpdate":
		resp, err = f.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FabricServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if f.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := f.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicesdatareplication.FabricModel](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginCreate(req.Context(), resourceGroupNameParam, fabricNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		f.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		f.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		f.beginCreate.remove(req)
	}

	return resp, nil
}

func (f *FabricServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if f.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := f.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginDelete(req.Context(), resourceGroupNameParam, fabricNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		f.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		f.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		f.beginDelete.remove(req)
	}

	return resp, nil
}

func (f *FabricServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Get(req.Context(), resourceGroupNameParam, fabricNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FabricModel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FabricServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := f.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationFabrics`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		var options *armrecoveryservicesdatareplication.FabricClientListOptions
		if continuationTokenParam != nil {
			options = &armrecoveryservicesdatareplication.FabricClientListOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := f.srv.NewListPager(resourceGroupNameParam, options)
		newListPager = &resp
		f.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armrecoveryservicesdatareplication.FabricClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		f.newListPager.remove(req)
	}
	return resp, nil
}

func (f *FabricServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := f.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationFabrics`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		var options *armrecoveryservicesdatareplication.FabricClientListBySubscriptionOptions
		if continuationTokenParam != nil {
			options = &armrecoveryservicesdatareplication.FabricClientListBySubscriptionOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := f.srv.NewListBySubscriptionPager(options)
		newListBySubscriptionPager = &resp
		f.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armrecoveryservicesdatareplication.FabricClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		f.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (f *FabricServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if f.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := f.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicesdatareplication.FabricModelUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginUpdate(req.Context(), resourceGroupNameParam, fabricNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		f.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		f.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		f.beginUpdate.remove(req)
	}

	return resp, nil
}
