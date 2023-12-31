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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// PrivateStoreCollectionOfferServer is a fake server for instances of the armmarketplace.PrivateStoreCollectionOfferClient type.
type PrivateStoreCollectionOfferServer struct {
	// CreateOrUpdate is the fake for method PrivateStoreCollectionOfferClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, privateStoreID string, offerID string, collectionID string, options *armmarketplace.PrivateStoreCollectionOfferClientCreateOrUpdateOptions) (resp azfake.Responder[armmarketplace.PrivateStoreCollectionOfferClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method PrivateStoreCollectionOfferClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, privateStoreID string, offerID string, collectionID string, options *armmarketplace.PrivateStoreCollectionOfferClientDeleteOptions) (resp azfake.Responder[armmarketplace.PrivateStoreCollectionOfferClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PrivateStoreCollectionOfferClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, privateStoreID string, offerID string, collectionID string, options *armmarketplace.PrivateStoreCollectionOfferClientGetOptions) (resp azfake.Responder[armmarketplace.PrivateStoreCollectionOfferClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PrivateStoreCollectionOfferClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(privateStoreID string, collectionID string, options *armmarketplace.PrivateStoreCollectionOfferClientListOptions) (resp azfake.PagerResponder[armmarketplace.PrivateStoreCollectionOfferClientListResponse])

	// Post is the fake for method PrivateStoreCollectionOfferClient.Post
	// HTTP status codes to indicate success: http.StatusOK
	Post func(ctx context.Context, privateStoreID string, offerID string, collectionID string, options *armmarketplace.PrivateStoreCollectionOfferClientPostOptions) (resp azfake.Responder[armmarketplace.PrivateStoreCollectionOfferClientPostResponse], errResp azfake.ErrorResponder)
}

// NewPrivateStoreCollectionOfferServerTransport creates a new instance of PrivateStoreCollectionOfferServerTransport with the provided implementation.
// The returned PrivateStoreCollectionOfferServerTransport instance is connected to an instance of armmarketplace.PrivateStoreCollectionOfferClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateStoreCollectionOfferServerTransport(srv *PrivateStoreCollectionOfferServer) *PrivateStoreCollectionOfferServerTransport {
	return &PrivateStoreCollectionOfferServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armmarketplace.PrivateStoreCollectionOfferClientListResponse]](),
	}
}

// PrivateStoreCollectionOfferServerTransport connects instances of armmarketplace.PrivateStoreCollectionOfferClient to instances of PrivateStoreCollectionOfferServer.
// Don't use this type directly, use NewPrivateStoreCollectionOfferServerTransport instead.
type PrivateStoreCollectionOfferServerTransport struct {
	srv          *PrivateStoreCollectionOfferServer
	newListPager *tracker[azfake.PagerResponder[armmarketplace.PrivateStoreCollectionOfferClientListResponse]]
}

// Do implements the policy.Transporter interface for PrivateStoreCollectionOfferServerTransport.
func (p *PrivateStoreCollectionOfferServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateStoreCollectionOfferClient.CreateOrUpdate":
		resp, err = p.dispatchCreateOrUpdate(req)
	case "PrivateStoreCollectionOfferClient.Delete":
		resp, err = p.dispatchDelete(req)
	case "PrivateStoreCollectionOfferClient.Get":
		resp, err = p.dispatchGet(req)
	case "PrivateStoreCollectionOfferClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	case "PrivateStoreCollectionOfferClient.Post":
		resp, err = p.dispatchPost(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateStoreCollectionOfferServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Marketplace/privateStores/(?P<privateStoreId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/collections/(?P<collectionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/offers/(?P<offerId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmarketplace.Offer](req)
	if err != nil {
		return nil, err
	}
	privateStoreIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateStoreId")])
	if err != nil {
		return nil, err
	}
	offerIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("offerId")])
	if err != nil {
		return nil, err
	}
	collectionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("collectionId")])
	if err != nil {
		return nil, err
	}
	var options *armmarketplace.PrivateStoreCollectionOfferClientCreateOrUpdateOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armmarketplace.PrivateStoreCollectionOfferClientCreateOrUpdateOptions{
			Payload: &body,
		}
	}
	respr, errRespr := p.srv.CreateOrUpdate(req.Context(), privateStoreIDParam, offerIDParam, collectionIDParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Offer, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateStoreCollectionOfferServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if p.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Marketplace/privateStores/(?P<privateStoreId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/collections/(?P<collectionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/offers/(?P<offerId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	privateStoreIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateStoreId")])
	if err != nil {
		return nil, err
	}
	offerIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("offerId")])
	if err != nil {
		return nil, err
	}
	collectionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("collectionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Delete(req.Context(), privateStoreIDParam, offerIDParam, collectionIDParam, nil)
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

func (p *PrivateStoreCollectionOfferServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Marketplace/privateStores/(?P<privateStoreId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/collections/(?P<collectionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/offers/(?P<offerId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	privateStoreIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateStoreId")])
	if err != nil {
		return nil, err
	}
	offerIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("offerId")])
	if err != nil {
		return nil, err
	}
	collectionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("collectionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), privateStoreIDParam, offerIDParam, collectionIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Offer, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateStoreCollectionOfferServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/providers/Microsoft\.Marketplace/privateStores/(?P<privateStoreId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/collections/(?P<collectionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/offers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		privateStoreIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateStoreId")])
		if err != nil {
			return nil, err
		}
		collectionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("collectionId")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListPager(privateStoreIDParam, collectionIDParam, nil)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmarketplace.PrivateStoreCollectionOfferClientListResponse, createLink func() string) {
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

func (p *PrivateStoreCollectionOfferServerTransport) dispatchPost(req *http.Request) (*http.Response, error) {
	if p.srv.Post == nil {
		return nil, &nonRetriableError{errors.New("fake for method Post not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Marketplace/privateStores/(?P<privateStoreId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/collections/(?P<collectionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/offers/(?P<offerId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmarketplace.Operation](req)
	if err != nil {
		return nil, err
	}
	privateStoreIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateStoreId")])
	if err != nil {
		return nil, err
	}
	offerIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("offerId")])
	if err != nil {
		return nil, err
	}
	collectionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("collectionId")])
	if err != nil {
		return nil, err
	}
	var options *armmarketplace.PrivateStoreCollectionOfferClientPostOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armmarketplace.PrivateStoreCollectionOfferClientPostOptions{
			Payload: &body,
		}
	}
	respr, errRespr := p.srv.Post(req.Context(), privateStoreIDParam, offerIDParam, collectionIDParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
