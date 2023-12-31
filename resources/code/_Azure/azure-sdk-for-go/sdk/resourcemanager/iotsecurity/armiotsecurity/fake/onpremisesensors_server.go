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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity"
	"net/http"
	"net/url"
	"regexp"
)

// OnPremiseSensorsServer is a fake server for instances of the armiotsecurity.OnPremiseSensorsClient type.
type OnPremiseSensorsServer struct {
	// CreateOrUpdate is the fake for method OnPremiseSensorsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, onPremiseSensorName string, options *armiotsecurity.OnPremiseSensorsClientCreateOrUpdateOptions) (resp azfake.Responder[armiotsecurity.OnPremiseSensorsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method OnPremiseSensorsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, onPremiseSensorName string, options *armiotsecurity.OnPremiseSensorsClientDeleteOptions) (resp azfake.Responder[armiotsecurity.OnPremiseSensorsClientDeleteResponse], errResp azfake.ErrorResponder)

	// DownloadActivation is the fake for method OnPremiseSensorsClient.DownloadActivation
	// HTTP status codes to indicate success: http.StatusOK
	DownloadActivation func(ctx context.Context, onPremiseSensorName string, options *armiotsecurity.OnPremiseSensorsClientDownloadActivationOptions) (resp azfake.Responder[armiotsecurity.OnPremiseSensorsClientDownloadActivationResponse], errResp azfake.ErrorResponder)

	// DownloadResetPassword is the fake for method OnPremiseSensorsClient.DownloadResetPassword
	// HTTP status codes to indicate success: http.StatusOK
	DownloadResetPassword func(ctx context.Context, onPremiseSensorName string, body armiotsecurity.ResetPasswordInput, options *armiotsecurity.OnPremiseSensorsClientDownloadResetPasswordOptions) (resp azfake.Responder[armiotsecurity.OnPremiseSensorsClientDownloadResetPasswordResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method OnPremiseSensorsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, onPremiseSensorName string, options *armiotsecurity.OnPremiseSensorsClientGetOptions) (resp azfake.Responder[armiotsecurity.OnPremiseSensorsClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method OnPremiseSensorsClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, options *armiotsecurity.OnPremiseSensorsClientListOptions) (resp azfake.Responder[armiotsecurity.OnPremiseSensorsClientListResponse], errResp azfake.ErrorResponder)
}

// NewOnPremiseSensorsServerTransport creates a new instance of OnPremiseSensorsServerTransport with the provided implementation.
// The returned OnPremiseSensorsServerTransport instance is connected to an instance of armiotsecurity.OnPremiseSensorsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOnPremiseSensorsServerTransport(srv *OnPremiseSensorsServer) *OnPremiseSensorsServerTransport {
	return &OnPremiseSensorsServerTransport{srv: srv}
}

// OnPremiseSensorsServerTransport connects instances of armiotsecurity.OnPremiseSensorsClient to instances of OnPremiseSensorsServer.
// Don't use this type directly, use NewOnPremiseSensorsServerTransport instead.
type OnPremiseSensorsServerTransport struct {
	srv *OnPremiseSensorsServer
}

// Do implements the policy.Transporter interface for OnPremiseSensorsServerTransport.
func (o *OnPremiseSensorsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "OnPremiseSensorsClient.CreateOrUpdate":
		resp, err = o.dispatchCreateOrUpdate(req)
	case "OnPremiseSensorsClient.Delete":
		resp, err = o.dispatchDelete(req)
	case "OnPremiseSensorsClient.DownloadActivation":
		resp, err = o.dispatchDownloadActivation(req)
	case "OnPremiseSensorsClient.DownloadResetPassword":
		resp, err = o.dispatchDownloadResetPassword(req)
	case "OnPremiseSensorsClient.Get":
		resp, err = o.dispatchGet(req)
	case "OnPremiseSensorsClient.List":
		resp, err = o.dispatchList(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *OnPremiseSensorsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if o.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IoTSecurity/onPremiseSensors/(?P<onPremiseSensorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	onPremiseSensorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("onPremiseSensorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.CreateOrUpdate(req.Context(), onPremiseSensorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OnPremiseSensor, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OnPremiseSensorsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if o.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IoTSecurity/onPremiseSensors/(?P<onPremiseSensorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	onPremiseSensorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("onPremiseSensorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Delete(req.Context(), onPremiseSensorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OnPremiseSensorsServerTransport) dispatchDownloadActivation(req *http.Request) (*http.Response, error) {
	if o.srv.DownloadActivation == nil {
		return nil, &nonRetriableError{errors.New("fake for method DownloadActivation not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IoTSecurity/onPremiseSensors/(?P<onPremiseSensorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/downloadActivation`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	onPremiseSensorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("onPremiseSensorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.DownloadActivation(req.Context(), onPremiseSensorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{
		Body:        server.GetResponse(respr).Body,
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OnPremiseSensorsServerTransport) dispatchDownloadResetPassword(req *http.Request) (*http.Response, error) {
	if o.srv.DownloadResetPassword == nil {
		return nil, &nonRetriableError{errors.New("fake for method DownloadResetPassword not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IoTSecurity/onPremiseSensors/(?P<onPremiseSensorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/downloadResetPassword`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armiotsecurity.ResetPasswordInput](req)
	if err != nil {
		return nil, err
	}
	onPremiseSensorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("onPremiseSensorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.DownloadResetPassword(req.Context(), onPremiseSensorNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, &server.ResponseOptions{
		Body:        server.GetResponse(respr).Body,
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OnPremiseSensorsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if o.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IoTSecurity/onPremiseSensors/(?P<onPremiseSensorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	onPremiseSensorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("onPremiseSensorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Get(req.Context(), onPremiseSensorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OnPremiseSensor, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OnPremiseSensorsServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if o.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IoTSecurity/onPremiseSensors`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := o.srv.List(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OnPremiseSensorsList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
