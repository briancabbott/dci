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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ArtifactManifestsServer is a fake server for instances of the armhybridnetwork.ArtifactManifestsClient type.
type ArtifactManifestsServer struct {
	// BeginCreateOrUpdate is the fake for method ArtifactManifestsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, artifactManifestName string, parameters armhybridnetwork.ArtifactManifest, options *armhybridnetwork.ArtifactManifestsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armhybridnetwork.ArtifactManifestsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ArtifactManifestsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, artifactManifestName string, options *armhybridnetwork.ArtifactManifestsClientBeginDeleteOptions) (resp azfake.PollerResponder[armhybridnetwork.ArtifactManifestsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ArtifactManifestsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, artifactManifestName string, options *armhybridnetwork.ArtifactManifestsClientGetOptions) (resp azfake.Responder[armhybridnetwork.ArtifactManifestsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByArtifactStorePager is the fake for method ArtifactManifestsClient.NewListByArtifactStorePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByArtifactStorePager func(resourceGroupName string, publisherName string, artifactStoreName string, options *armhybridnetwork.ArtifactManifestsClientListByArtifactStoreOptions) (resp azfake.PagerResponder[armhybridnetwork.ArtifactManifestsClientListByArtifactStoreResponse])

	// ListCredential is the fake for method ArtifactManifestsClient.ListCredential
	// HTTP status codes to indicate success: http.StatusOK
	ListCredential func(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, artifactManifestName string, options *armhybridnetwork.ArtifactManifestsClientListCredentialOptions) (resp azfake.Responder[armhybridnetwork.ArtifactManifestsClientListCredentialResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method ArtifactManifestsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, artifactManifestName string, parameters armhybridnetwork.TagsObject, options *armhybridnetwork.ArtifactManifestsClientUpdateOptions) (resp azfake.Responder[armhybridnetwork.ArtifactManifestsClientUpdateResponse], errResp azfake.ErrorResponder)

	// BeginUpdateState is the fake for method ArtifactManifestsClient.BeginUpdateState
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateState func(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, artifactManifestName string, parameters armhybridnetwork.ArtifactManifestUpdateState, options *armhybridnetwork.ArtifactManifestsClientBeginUpdateStateOptions) (resp azfake.PollerResponder[armhybridnetwork.ArtifactManifestsClientUpdateStateResponse], errResp azfake.ErrorResponder)
}

// NewArtifactManifestsServerTransport creates a new instance of ArtifactManifestsServerTransport with the provided implementation.
// The returned ArtifactManifestsServerTransport instance is connected to an instance of armhybridnetwork.ArtifactManifestsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewArtifactManifestsServerTransport(srv *ArtifactManifestsServer) *ArtifactManifestsServerTransport {
	return &ArtifactManifestsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armhybridnetwork.ArtifactManifestsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armhybridnetwork.ArtifactManifestsClientDeleteResponse]](),
		newListByArtifactStorePager: newTracker[azfake.PagerResponder[armhybridnetwork.ArtifactManifestsClientListByArtifactStoreResponse]](),
		beginUpdateState:            newTracker[azfake.PollerResponder[armhybridnetwork.ArtifactManifestsClientUpdateStateResponse]](),
	}
}

// ArtifactManifestsServerTransport connects instances of armhybridnetwork.ArtifactManifestsClient to instances of ArtifactManifestsServer.
// Don't use this type directly, use NewArtifactManifestsServerTransport instead.
type ArtifactManifestsServerTransport struct {
	srv                         *ArtifactManifestsServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armhybridnetwork.ArtifactManifestsClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armhybridnetwork.ArtifactManifestsClientDeleteResponse]]
	newListByArtifactStorePager *tracker[azfake.PagerResponder[armhybridnetwork.ArtifactManifestsClientListByArtifactStoreResponse]]
	beginUpdateState            *tracker[azfake.PollerResponder[armhybridnetwork.ArtifactManifestsClientUpdateStateResponse]]
}

// Do implements the policy.Transporter interface for ArtifactManifestsServerTransport.
func (a *ArtifactManifestsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ArtifactManifestsClient.BeginCreateOrUpdate":
		resp, err = a.dispatchBeginCreateOrUpdate(req)
	case "ArtifactManifestsClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "ArtifactManifestsClient.Get":
		resp, err = a.dispatchGet(req)
	case "ArtifactManifestsClient.NewListByArtifactStorePager":
		resp, err = a.dispatchNewListByArtifactStorePager(req)
	case "ArtifactManifestsClient.ListCredential":
		resp, err = a.dispatchListCredential(req)
	case "ArtifactManifestsClient.Update":
		resp, err = a.dispatchUpdate(req)
	case "ArtifactManifestsClient.BeginUpdateState":
		resp, err = a.dispatchBeginUpdateState(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *ArtifactManifestsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := a.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactStores/(?P<artifactStoreName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactManifests/(?P<artifactManifestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhybridnetwork.ArtifactManifest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
		if err != nil {
			return nil, err
		}
		artifactStoreNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactStoreName")])
		if err != nil {
			return nil, err
		}
		artifactManifestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactManifestName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, publisherNameParam, artifactStoreNameParam, artifactManifestNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		a.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		a.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		a.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (a *ArtifactManifestsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactStores/(?P<artifactStoreName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactManifests/(?P<artifactManifestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
		if err != nil {
			return nil, err
		}
		artifactStoreNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactStoreName")])
		if err != nil {
			return nil, err
		}
		artifactManifestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactManifestName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, publisherNameParam, artifactStoreNameParam, artifactManifestNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		a.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		a.beginDelete.remove(req)
	}

	return resp, nil
}

func (a *ArtifactManifestsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactStores/(?P<artifactStoreName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactManifests/(?P<artifactManifestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	artifactStoreNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactStoreName")])
	if err != nil {
		return nil, err
	}
	artifactManifestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactManifestName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, publisherNameParam, artifactStoreNameParam, artifactManifestNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ArtifactManifest, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ArtifactManifestsServerTransport) dispatchNewListByArtifactStorePager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByArtifactStorePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByArtifactStorePager not implemented")}
	}
	newListByArtifactStorePager := a.newListByArtifactStorePager.get(req)
	if newListByArtifactStorePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactStores/(?P<artifactStoreName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactManifests`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
		if err != nil {
			return nil, err
		}
		artifactStoreNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactStoreName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByArtifactStorePager(resourceGroupNameParam, publisherNameParam, artifactStoreNameParam, nil)
		newListByArtifactStorePager = &resp
		a.newListByArtifactStorePager.add(req, newListByArtifactStorePager)
		server.PagerResponderInjectNextLinks(newListByArtifactStorePager, req, func(page *armhybridnetwork.ArtifactManifestsClientListByArtifactStoreResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByArtifactStorePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByArtifactStorePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByArtifactStorePager) {
		a.newListByArtifactStorePager.remove(req)
	}
	return resp, nil
}

func (a *ArtifactManifestsServerTransport) dispatchListCredential(req *http.Request) (*http.Response, error) {
	if a.srv.ListCredential == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListCredential not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactStores/(?P<artifactStoreName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactManifests/(?P<artifactManifestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listCredential`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	artifactStoreNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactStoreName")])
	if err != nil {
		return nil, err
	}
	artifactManifestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactManifestName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.ListCredential(req.Context(), resourceGroupNameParam, publisherNameParam, artifactStoreNameParam, artifactManifestNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ArtifactAccessCredentialClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ArtifactManifestsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactStores/(?P<artifactStoreName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactManifests/(?P<artifactManifestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armhybridnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	artifactStoreNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactStoreName")])
	if err != nil {
		return nil, err
	}
	artifactManifestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactManifestName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Update(req.Context(), resourceGroupNameParam, publisherNameParam, artifactStoreNameParam, artifactManifestNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ArtifactManifest, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ArtifactManifestsServerTransport) dispatchBeginUpdateState(req *http.Request) (*http.Response, error) {
	if a.srv.BeginUpdateState == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateState not implemented")}
	}
	beginUpdateState := a.beginUpdateState.get(req)
	if beginUpdateState == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactStores/(?P<artifactStoreName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactManifests/(?P<artifactManifestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateState`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhybridnetwork.ArtifactManifestUpdateState](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
		if err != nil {
			return nil, err
		}
		artifactStoreNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactStoreName")])
		if err != nil {
			return nil, err
		}
		artifactManifestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactManifestName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginUpdateState(req.Context(), resourceGroupNameParam, publisherNameParam, artifactStoreNameParam, artifactManifestNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateState = &respr
		a.beginUpdateState.add(req, beginUpdateState)
	}

	resp, err := server.PollerResponderNext(beginUpdateState, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginUpdateState.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateState) {
		a.beginUpdateState.remove(req)
	}

	return resp, nil
}
