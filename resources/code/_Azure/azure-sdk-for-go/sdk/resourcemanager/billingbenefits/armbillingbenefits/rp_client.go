//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbillingbenefits

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// RPClient contains the methods for the BillingBenefitsRP group.
// Don't use this type directly, use NewRPClient() instead.
type RPClient struct {
	internal *arm.Client
}

// NewRPClient creates a new instance of RPClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRPClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*RPClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RPClient{
		internal: cl,
	}
	return client, nil
}

// ValidatePurchase - Validate savings plan purchase.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - body - Request body for validating the purchase of a savings plan
//   - options - RPClientValidatePurchaseOptions contains the optional parameters for the RPClient.ValidatePurchase method.
func (client *RPClient) ValidatePurchase(ctx context.Context, body SavingsPlanPurchaseValidateRequest, options *RPClientValidatePurchaseOptions) (RPClientValidatePurchaseResponse, error) {
	var err error
	const operationName = "RPClient.ValidatePurchase"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.validatePurchaseCreateRequest(ctx, body, options)
	if err != nil {
		return RPClientValidatePurchaseResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RPClientValidatePurchaseResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RPClientValidatePurchaseResponse{}, err
	}
	resp, err := client.validatePurchaseHandleResponse(httpResp)
	return resp, err
}

// validatePurchaseCreateRequest creates the ValidatePurchase request.
func (client *RPClient) validatePurchaseCreateRequest(ctx context.Context, body SavingsPlanPurchaseValidateRequest, options *RPClientValidatePurchaseOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BillingBenefits/validate"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// validatePurchaseHandleResponse handles the ValidatePurchase response.
func (client *RPClient) validatePurchaseHandleResponse(resp *http.Response) (RPClientValidatePurchaseResponse, error) {
	result := RPClientValidatePurchaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SavingsPlanValidateResponse); err != nil {
		return RPClientValidatePurchaseResponse{}, err
	}
	return result, nil
}
