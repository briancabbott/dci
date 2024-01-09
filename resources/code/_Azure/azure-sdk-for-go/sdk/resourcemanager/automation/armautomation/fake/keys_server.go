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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
	"net/http"
	"net/url"
	"regexp"
)

// KeysServer is a fake server for instances of the armautomation.KeysClient type.
type KeysServer struct {
	// ListByAutomationAccount is the fake for method KeysClient.ListByAutomationAccount
	// HTTP status codes to indicate success: http.StatusOK
	ListByAutomationAccount func(ctx context.Context, resourceGroupName string, automationAccountName string, options *armautomation.KeysClientListByAutomationAccountOptions) (resp azfake.Responder[armautomation.KeysClientListByAutomationAccountResponse], errResp azfake.ErrorResponder)
}

// NewKeysServerTransport creates a new instance of KeysServerTransport with the provided implementation.
// The returned KeysServerTransport instance is connected to an instance of armautomation.KeysClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewKeysServerTransport(srv *KeysServer) *KeysServerTransport {
	return &KeysServerTransport{srv: srv}
}

// KeysServerTransport connects instances of armautomation.KeysClient to instances of KeysServer.
// Don't use this type directly, use NewKeysServerTransport instead.
type KeysServerTransport struct {
	srv *KeysServer
}

// Do implements the policy.Transporter interface for KeysServerTransport.
func (k *KeysServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "KeysClient.ListByAutomationAccount":
		resp, err = k.dispatchListByAutomationAccount(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (k *KeysServerTransport) dispatchListByAutomationAccount(req *http.Request) (*http.Response, error) {
	if k.srv.ListByAutomationAccount == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListByAutomationAccount not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listKeys`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := k.srv.ListByAutomationAccount(req.Context(), resourceGroupNameParam, automationAccountNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).KeyListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
