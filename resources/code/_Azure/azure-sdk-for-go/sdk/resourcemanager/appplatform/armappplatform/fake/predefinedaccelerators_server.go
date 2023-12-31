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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
	"net/http"
	"net/url"
	"regexp"
)

// PredefinedAcceleratorsServer is a fake server for instances of the armappplatform.PredefinedAcceleratorsClient type.
type PredefinedAcceleratorsServer struct {
	// BeginDisable is the fake for method PredefinedAcceleratorsClient.BeginDisable
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDisable func(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, predefinedAcceleratorName string, options *armappplatform.PredefinedAcceleratorsClientBeginDisableOptions) (resp azfake.PollerResponder[armappplatform.PredefinedAcceleratorsClientDisableResponse], errResp azfake.ErrorResponder)

	// BeginEnable is the fake for method PredefinedAcceleratorsClient.BeginEnable
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginEnable func(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, predefinedAcceleratorName string, options *armappplatform.PredefinedAcceleratorsClientBeginEnableOptions) (resp azfake.PollerResponder[armappplatform.PredefinedAcceleratorsClientEnableResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PredefinedAcceleratorsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, predefinedAcceleratorName string, options *armappplatform.PredefinedAcceleratorsClientGetOptions) (resp azfake.Responder[armappplatform.PredefinedAcceleratorsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PredefinedAcceleratorsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, serviceName string, applicationAcceleratorName string, options *armappplatform.PredefinedAcceleratorsClientListOptions) (resp azfake.PagerResponder[armappplatform.PredefinedAcceleratorsClientListResponse])
}

// NewPredefinedAcceleratorsServerTransport creates a new instance of PredefinedAcceleratorsServerTransport with the provided implementation.
// The returned PredefinedAcceleratorsServerTransport instance is connected to an instance of armappplatform.PredefinedAcceleratorsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPredefinedAcceleratorsServerTransport(srv *PredefinedAcceleratorsServer) *PredefinedAcceleratorsServerTransport {
	return &PredefinedAcceleratorsServerTransport{
		srv:          srv,
		beginDisable: newTracker[azfake.PollerResponder[armappplatform.PredefinedAcceleratorsClientDisableResponse]](),
		beginEnable:  newTracker[azfake.PollerResponder[armappplatform.PredefinedAcceleratorsClientEnableResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armappplatform.PredefinedAcceleratorsClientListResponse]](),
	}
}

// PredefinedAcceleratorsServerTransport connects instances of armappplatform.PredefinedAcceleratorsClient to instances of PredefinedAcceleratorsServer.
// Don't use this type directly, use NewPredefinedAcceleratorsServerTransport instead.
type PredefinedAcceleratorsServerTransport struct {
	srv          *PredefinedAcceleratorsServer
	beginDisable *tracker[azfake.PollerResponder[armappplatform.PredefinedAcceleratorsClientDisableResponse]]
	beginEnable  *tracker[azfake.PollerResponder[armappplatform.PredefinedAcceleratorsClientEnableResponse]]
	newListPager *tracker[azfake.PagerResponder[armappplatform.PredefinedAcceleratorsClientListResponse]]
}

// Do implements the policy.Transporter interface for PredefinedAcceleratorsServerTransport.
func (p *PredefinedAcceleratorsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PredefinedAcceleratorsClient.BeginDisable":
		resp, err = p.dispatchBeginDisable(req)
	case "PredefinedAcceleratorsClient.BeginEnable":
		resp, err = p.dispatchBeginEnable(req)
	case "PredefinedAcceleratorsClient.Get":
		resp, err = p.dispatchGet(req)
	case "PredefinedAcceleratorsClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PredefinedAcceleratorsServerTransport) dispatchBeginDisable(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDisable == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDisable not implemented")}
	}
	beginDisable := p.beginDisable.get(req)
	if beginDisable == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationAccelerators/(?P<applicationAcceleratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/predefinedAccelerators/(?P<predefinedAcceleratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disable`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		applicationAcceleratorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationAcceleratorName")])
		if err != nil {
			return nil, err
		}
		predefinedAcceleratorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("predefinedAcceleratorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDisable(req.Context(), resourceGroupNameParam, serviceNameParam, applicationAcceleratorNameParam, predefinedAcceleratorNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDisable = &respr
		p.beginDisable.add(req, beginDisable)
	}

	resp, err := server.PollerResponderNext(beginDisable, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		p.beginDisable.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDisable) {
		p.beginDisable.remove(req)
	}

	return resp, nil
}

func (p *PredefinedAcceleratorsServerTransport) dispatchBeginEnable(req *http.Request) (*http.Response, error) {
	if p.srv.BeginEnable == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginEnable not implemented")}
	}
	beginEnable := p.beginEnable.get(req)
	if beginEnable == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationAccelerators/(?P<applicationAcceleratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/predefinedAccelerators/(?P<predefinedAcceleratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/enable`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		applicationAcceleratorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationAcceleratorName")])
		if err != nil {
			return nil, err
		}
		predefinedAcceleratorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("predefinedAcceleratorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginEnable(req.Context(), resourceGroupNameParam, serviceNameParam, applicationAcceleratorNameParam, predefinedAcceleratorNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginEnable = &respr
		p.beginEnable.add(req, beginEnable)
	}

	resp, err := server.PollerResponderNext(beginEnable, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		p.beginEnable.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginEnable) {
		p.beginEnable.remove(req)
	}

	return resp, nil
}

func (p *PredefinedAcceleratorsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationAccelerators/(?P<applicationAcceleratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/predefinedAccelerators/(?P<predefinedAcceleratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	applicationAcceleratorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationAcceleratorName")])
	if err != nil {
		return nil, err
	}
	predefinedAcceleratorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("predefinedAcceleratorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, applicationAcceleratorNameParam, predefinedAcceleratorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PredefinedAcceleratorResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PredefinedAcceleratorsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationAccelerators/(?P<applicationAcceleratorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/predefinedAccelerators`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		applicationAcceleratorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationAcceleratorName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListPager(resourceGroupNameParam, serviceNameParam, applicationAcceleratorNameParam, nil)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappplatform.PredefinedAcceleratorsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		p.newListPager.remove(req)
	}
	return resp, nil
}
