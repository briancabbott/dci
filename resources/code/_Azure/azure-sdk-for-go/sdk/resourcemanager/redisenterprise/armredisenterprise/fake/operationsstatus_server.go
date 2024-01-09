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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise"
	"net/http"
	"net/url"
	"regexp"
)

// OperationsStatusServer is a fake server for instances of the armredisenterprise.OperationsStatusClient type.
type OperationsStatusServer struct {
	// Get is the fake for method OperationsStatusClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, operationID string, options *armredisenterprise.OperationsStatusClientGetOptions) (resp azfake.Responder[armredisenterprise.OperationsStatusClientGetResponse], errResp azfake.ErrorResponder)
}

// NewOperationsStatusServerTransport creates a new instance of OperationsStatusServerTransport with the provided implementation.
// The returned OperationsStatusServerTransport instance is connected to an instance of armredisenterprise.OperationsStatusClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOperationsStatusServerTransport(srv *OperationsStatusServer) *OperationsStatusServerTransport {
	return &OperationsStatusServerTransport{srv: srv}
}

// OperationsStatusServerTransport connects instances of armredisenterprise.OperationsStatusClient to instances of OperationsStatusServer.
// Don't use this type directly, use NewOperationsStatusServerTransport instead.
type OperationsStatusServerTransport struct {
	srv *OperationsStatusServer
}

// Do implements the policy.Transporter interface for OperationsStatusServerTransport.
func (o *OperationsStatusServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "OperationsStatusClient.Get":
		resp, err = o.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *OperationsStatusServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if o.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cache/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operationsStatus/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Get(req.Context(), locationParam, operationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OperationStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
