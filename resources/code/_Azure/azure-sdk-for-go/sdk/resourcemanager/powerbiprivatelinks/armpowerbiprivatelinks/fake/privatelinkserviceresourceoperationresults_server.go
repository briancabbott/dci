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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks/v2"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateLinkServiceResourceOperationResultsServer is a fake server for instances of the armpowerbiprivatelinks.PrivateLinkServiceResourceOperationResultsClient type.
type PrivateLinkServiceResourceOperationResultsServer struct {
	// BeginGet is the fake for method PrivateLinkServiceResourceOperationResultsClient.BeginGet
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGet func(ctx context.Context, operationID string, options *armpowerbiprivatelinks.PrivateLinkServiceResourceOperationResultsClientBeginGetOptions) (resp azfake.PollerResponder[armpowerbiprivatelinks.PrivateLinkServiceResourceOperationResultsClientGetResponse], errResp azfake.ErrorResponder)
}

// NewPrivateLinkServiceResourceOperationResultsServerTransport creates a new instance of PrivateLinkServiceResourceOperationResultsServerTransport with the provided implementation.
// The returned PrivateLinkServiceResourceOperationResultsServerTransport instance is connected to an instance of armpowerbiprivatelinks.PrivateLinkServiceResourceOperationResultsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateLinkServiceResourceOperationResultsServerTransport(srv *PrivateLinkServiceResourceOperationResultsServer) *PrivateLinkServiceResourceOperationResultsServerTransport {
	return &PrivateLinkServiceResourceOperationResultsServerTransport{
		srv:      srv,
		beginGet: newTracker[azfake.PollerResponder[armpowerbiprivatelinks.PrivateLinkServiceResourceOperationResultsClientGetResponse]](),
	}
}

// PrivateLinkServiceResourceOperationResultsServerTransport connects instances of armpowerbiprivatelinks.PrivateLinkServiceResourceOperationResultsClient to instances of PrivateLinkServiceResourceOperationResultsServer.
// Don't use this type directly, use NewPrivateLinkServiceResourceOperationResultsServerTransport instead.
type PrivateLinkServiceResourceOperationResultsServerTransport struct {
	srv      *PrivateLinkServiceResourceOperationResultsServer
	beginGet *tracker[azfake.PollerResponder[armpowerbiprivatelinks.PrivateLinkServiceResourceOperationResultsClientGetResponse]]
}

// Do implements the policy.Transporter interface for PrivateLinkServiceResourceOperationResultsServerTransport.
func (p *PrivateLinkServiceResourceOperationResultsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateLinkServiceResourceOperationResultsClient.BeginGet":
		resp, err = p.dispatchBeginGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateLinkServiceResourceOperationResultsServerTransport) dispatchBeginGet(req *http.Request) (*http.Response, error) {
	if p.srv.BeginGet == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGet not implemented")}
	}
	beginGet := p.beginGet.get(req)
	if beginGet == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.PowerBI/operationResults/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginGet(req.Context(), operationIDParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGet = &respr
		p.beginGet.add(req, beginGet)
	}

	resp, err := server.PollerResponderNext(beginGet, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		p.beginGet.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGet) {
		p.beginGet.remove(req)
	}

	return resp, nil
}
