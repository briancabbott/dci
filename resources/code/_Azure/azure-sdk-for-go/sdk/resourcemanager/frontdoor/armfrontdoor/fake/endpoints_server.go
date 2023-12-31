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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
	"net/http"
	"net/url"
	"regexp"
)

// EndpointsServer is a fake server for instances of the armfrontdoor.EndpointsClient type.
type EndpointsServer struct {
	// BeginPurgeContent is the fake for method EndpointsClient.BeginPurgeContent
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginPurgeContent func(ctx context.Context, resourceGroupName string, frontDoorName string, contentFilePaths armfrontdoor.PurgeParameters, options *armfrontdoor.EndpointsClientBeginPurgeContentOptions) (resp azfake.PollerResponder[armfrontdoor.EndpointsClientPurgeContentResponse], errResp azfake.ErrorResponder)
}

// NewEndpointsServerTransport creates a new instance of EndpointsServerTransport with the provided implementation.
// The returned EndpointsServerTransport instance is connected to an instance of armfrontdoor.EndpointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEndpointsServerTransport(srv *EndpointsServer) *EndpointsServerTransport {
	return &EndpointsServerTransport{
		srv:               srv,
		beginPurgeContent: newTracker[azfake.PollerResponder[armfrontdoor.EndpointsClientPurgeContentResponse]](),
	}
}

// EndpointsServerTransport connects instances of armfrontdoor.EndpointsClient to instances of EndpointsServer.
// Don't use this type directly, use NewEndpointsServerTransport instead.
type EndpointsServerTransport struct {
	srv               *EndpointsServer
	beginPurgeContent *tracker[azfake.PollerResponder[armfrontdoor.EndpointsClientPurgeContentResponse]]
}

// Do implements the policy.Transporter interface for EndpointsServerTransport.
func (e *EndpointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "EndpointsClient.BeginPurgeContent":
		resp, err = e.dispatchBeginPurgeContent(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *EndpointsServerTransport) dispatchBeginPurgeContent(req *http.Request) (*http.Response, error) {
	if e.srv.BeginPurgeContent == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPurgeContent not implemented")}
	}
	beginPurgeContent := e.beginPurgeContent.get(req)
	if beginPurgeContent == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/frontDoors/(?P<frontDoorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/purge`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armfrontdoor.PurgeParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		frontDoorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("frontDoorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginPurgeContent(req.Context(), resourceGroupNameParam, frontDoorNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPurgeContent = &respr
		e.beginPurgeContent.add(req, beginPurgeContent)
	}

	resp, err := server.PollerResponderNext(beginPurgeContent, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		e.beginPurgeContent.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPurgeContent) {
		e.beginPurgeContent.remove(req)
	}

	return resp, nil
}
