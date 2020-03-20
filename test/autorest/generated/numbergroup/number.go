// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package numbergroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// NumberOperations contains the methods for the Number group.
type NumberOperations interface {
	// GetBigDecimal - Get big decimal value 2.5976931e+101
	GetBigDecimal(ctx context.Context) (*Float64Response, error)
	// GetBigDecimalNegativeDecimal - Get big decimal value -99999999.99
	GetBigDecimalNegativeDecimal(ctx context.Context) (*Float64Response, error)
	// GetBigDecimalPositiveDecimal - Get big decimal value 99999999.99
	GetBigDecimalPositiveDecimal(ctx context.Context) (*Float64Response, error)
	// GetBigDouble - Get big double value 2.5976931e+101
	GetBigDouble(ctx context.Context) (*Float64Response, error)
	// GetBigDoubleNegativeDecimal - Get big double value -99999999.99
	GetBigDoubleNegativeDecimal(ctx context.Context) (*Float64Response, error)
	// GetBigDoublePositiveDecimal - Get big double value 99999999.99
	GetBigDoublePositiveDecimal(ctx context.Context) (*Float64Response, error)
	// GetBigFloat - Get big float value 3.402823e+20
	GetBigFloat(ctx context.Context) (*Float32Response, error)
	// GetInvalidDecimal - Get invalid decimal Number value
	GetInvalidDecimal(ctx context.Context) (*Float64Response, error)
	// GetInvalidDouble - Get invalid double Number value
	GetInvalidDouble(ctx context.Context) (*Float64Response, error)
	// GetInvalidFloat - Get invalid float Number value
	GetInvalidFloat(ctx context.Context) (*Float32Response, error)
	// GetNull - Get null Number value
	GetNull(ctx context.Context) (*Float32Response, error)
	// GetSmallDecimal - Get small decimal value 2.5976931e-101
	GetSmallDecimal(ctx context.Context) (*Float64Response, error)
	// GetSmallDouble - Get big double value 2.5976931e-101
	GetSmallDouble(ctx context.Context) (*Float64Response, error)
	// GetSmallFloat - Get big double value 3.402823e-20
	GetSmallFloat(ctx context.Context) (*Float64Response, error)
	// PutBigDecimal - Put big decimal value 2.5976931e+101
	PutBigDecimal(ctx context.Context, numberBody float64) (*http.Response, error)
	// PutBigDecimalNegativeDecimal - Put big decimal value -99999999.99
	PutBigDecimalNegativeDecimal(ctx context.Context) (*http.Response, error)
	// PutBigDecimalPositiveDecimal - Put big decimal value 99999999.99
	PutBigDecimalPositiveDecimal(ctx context.Context) (*http.Response, error)
	// PutBigDouble - Put big double value 2.5976931e+101
	PutBigDouble(ctx context.Context, numberBody float64) (*http.Response, error)
	// PutBigDoubleNegativeDecimal - Put big double value -99999999.99
	PutBigDoubleNegativeDecimal(ctx context.Context) (*http.Response, error)
	// PutBigDoublePositiveDecimal - Put big double value 99999999.99
	PutBigDoublePositiveDecimal(ctx context.Context) (*http.Response, error)
	// PutBigFloat - Put big float value 3.402823e+20
	PutBigFloat(ctx context.Context, numberBody float32) (*http.Response, error)
	// PutSmallDecimal - Put small decimal value 2.5976931e-101
	PutSmallDecimal(ctx context.Context, numberBody float64) (*http.Response, error)
	// PutSmallDouble - Put small double value 2.5976931e-101
	PutSmallDouble(ctx context.Context, numberBody float64) (*http.Response, error)
	// PutSmallFloat - Put small float value 3.402823e-20
	PutSmallFloat(ctx context.Context, numberBody float32) (*http.Response, error)
}

// numberOperations implements the NumberOperations interface.
type numberOperations struct {
	*Client
}

// GetBigDecimal - Get big decimal value 2.5976931e+101
func (client *numberOperations) GetBigDecimal(ctx context.Context) (*Float64Response, error) {
	req, err := client.getBigDecimalCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getBigDecimalHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getBigDecimalCreateRequest creates the GetBigDecimal request.
func (client *numberOperations) getBigDecimalCreateRequest() (*azcore.Request, error) {
	urlPath := "/number/big/decimal/2.5976931e+101"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getBigDecimalHandleResponse handles the GetBigDecimal response.
func (client *numberOperations) getBigDecimalHandleResponse(resp *azcore.Response) (*Float64Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := Float64Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetBigDecimalNegativeDecimal - Get big decimal value -99999999.99
func (client *numberOperations) GetBigDecimalNegativeDecimal(ctx context.Context) (*Float64Response, error) {
	req, err := client.getBigDecimalNegativeDecimalCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getBigDecimalNegativeDecimalHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getBigDecimalNegativeDecimalCreateRequest creates the GetBigDecimalNegativeDecimal request.
func (client *numberOperations) getBigDecimalNegativeDecimalCreateRequest() (*azcore.Request, error) {
	urlPath := "/number/big/decimal/-99999999.99"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getBigDecimalNegativeDecimalHandleResponse handles the GetBigDecimalNegativeDecimal response.
func (client *numberOperations) getBigDecimalNegativeDecimalHandleResponse(resp *azcore.Response) (*Float64Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := Float64Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetBigDecimalPositiveDecimal - Get big decimal value 99999999.99
func (client *numberOperations) GetBigDecimalPositiveDecimal(ctx context.Context) (*Float64Response, error) {
	req, err := client.getBigDecimalPositiveDecimalCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getBigDecimalPositiveDecimalHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getBigDecimalPositiveDecimalCreateRequest creates the GetBigDecimalPositiveDecimal request.
func (client *numberOperations) getBigDecimalPositiveDecimalCreateRequest() (*azcore.Request, error) {
	urlPath := "/number/big/decimal/99999999.99"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getBigDecimalPositiveDecimalHandleResponse handles the GetBigDecimalPositiveDecimal response.
func (client *numberOperations) getBigDecimalPositiveDecimalHandleResponse(resp *azcore.Response) (*Float64Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := Float64Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetBigDouble - Get big double value 2.5976931e+101
func (client *numberOperations) GetBigDouble(ctx context.Context) (*Float64Response, error) {
	req, err := client.getBigDoubleCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getBigDoubleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getBigDoubleCreateRequest creates the GetBigDouble request.
func (client *numberOperations) getBigDoubleCreateRequest() (*azcore.Request, error) {
	urlPath := "/number/big/double/2.5976931e+101"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getBigDoubleHandleResponse handles the GetBigDouble response.
func (client *numberOperations) getBigDoubleHandleResponse(resp *azcore.Response) (*Float64Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := Float64Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetBigDoubleNegativeDecimal - Get big double value -99999999.99
func (client *numberOperations) GetBigDoubleNegativeDecimal(ctx context.Context) (*Float64Response, error) {
	req, err := client.getBigDoubleNegativeDecimalCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getBigDoubleNegativeDecimalHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getBigDoubleNegativeDecimalCreateRequest creates the GetBigDoubleNegativeDecimal request.
func (client *numberOperations) getBigDoubleNegativeDecimalCreateRequest() (*azcore.Request, error) {
	urlPath := "/number/big/double/-99999999.99"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getBigDoubleNegativeDecimalHandleResponse handles the GetBigDoubleNegativeDecimal response.
func (client *numberOperations) getBigDoubleNegativeDecimalHandleResponse(resp *azcore.Response) (*Float64Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := Float64Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetBigDoublePositiveDecimal - Get big double value 99999999.99
func (client *numberOperations) GetBigDoublePositiveDecimal(ctx context.Context) (*Float64Response, error) {
	req, err := client.getBigDoublePositiveDecimalCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getBigDoublePositiveDecimalHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getBigDoublePositiveDecimalCreateRequest creates the GetBigDoublePositiveDecimal request.
func (client *numberOperations) getBigDoublePositiveDecimalCreateRequest() (*azcore.Request, error) {
	urlPath := "/number/big/double/99999999.99"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getBigDoublePositiveDecimalHandleResponse handles the GetBigDoublePositiveDecimal response.
func (client *numberOperations) getBigDoublePositiveDecimalHandleResponse(resp *azcore.Response) (*Float64Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := Float64Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetBigFloat - Get big float value 3.402823e+20
func (client *numberOperations) GetBigFloat(ctx context.Context) (*Float32Response, error) {
	req, err := client.getBigFloatCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getBigFloatHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getBigFloatCreateRequest creates the GetBigFloat request.
func (client *numberOperations) getBigFloatCreateRequest() (*azcore.Request, error) {
	urlPath := "/number/big/float/3.402823e+20"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getBigFloatHandleResponse handles the GetBigFloat response.
func (client *numberOperations) getBigFloatHandleResponse(resp *azcore.Response) (*Float32Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := Float32Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetInvalidDecimal - Get invalid decimal Number value
func (client *numberOperations) GetInvalidDecimal(ctx context.Context) (*Float64Response, error) {
	req, err := client.getInvalidDecimalCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getInvalidDecimalHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getInvalidDecimalCreateRequest creates the GetInvalidDecimal request.
func (client *numberOperations) getInvalidDecimalCreateRequest() (*azcore.Request, error) {
	urlPath := "/number/invaliddecimal"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getInvalidDecimalHandleResponse handles the GetInvalidDecimal response.
func (client *numberOperations) getInvalidDecimalHandleResponse(resp *azcore.Response) (*Float64Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := Float64Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetInvalidDouble - Get invalid double Number value
func (client *numberOperations) GetInvalidDouble(ctx context.Context) (*Float64Response, error) {
	req, err := client.getInvalidDoubleCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getInvalidDoubleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getInvalidDoubleCreateRequest creates the GetInvalidDouble request.
func (client *numberOperations) getInvalidDoubleCreateRequest() (*azcore.Request, error) {
	urlPath := "/number/invaliddouble"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getInvalidDoubleHandleResponse handles the GetInvalidDouble response.
func (client *numberOperations) getInvalidDoubleHandleResponse(resp *azcore.Response) (*Float64Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := Float64Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetInvalidFloat - Get invalid float Number value
func (client *numberOperations) GetInvalidFloat(ctx context.Context) (*Float32Response, error) {
	req, err := client.getInvalidFloatCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getInvalidFloatHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getInvalidFloatCreateRequest creates the GetInvalidFloat request.
func (client *numberOperations) getInvalidFloatCreateRequest() (*azcore.Request, error) {
	urlPath := "/number/invalidfloat"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getInvalidFloatHandleResponse handles the GetInvalidFloat response.
func (client *numberOperations) getInvalidFloatHandleResponse(resp *azcore.Response) (*Float32Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := Float32Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetNull - Get null Number value
func (client *numberOperations) GetNull(ctx context.Context) (*Float32Response, error) {
	req, err := client.getNullCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getNullCreateRequest creates the GetNull request.
func (client *numberOperations) getNullCreateRequest() (*azcore.Request, error) {
	urlPath := "/number/null"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *numberOperations) getNullHandleResponse(resp *azcore.Response) (*Float32Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := Float32Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetSmallDecimal - Get small decimal value 2.5976931e-101
func (client *numberOperations) GetSmallDecimal(ctx context.Context) (*Float64Response, error) {
	req, err := client.getSmallDecimalCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getSmallDecimalHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getSmallDecimalCreateRequest creates the GetSmallDecimal request.
func (client *numberOperations) getSmallDecimalCreateRequest() (*azcore.Request, error) {
	urlPath := "/number/small/decimal/2.5976931e-101"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getSmallDecimalHandleResponse handles the GetSmallDecimal response.
func (client *numberOperations) getSmallDecimalHandleResponse(resp *azcore.Response) (*Float64Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := Float64Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetSmallDouble - Get big double value 2.5976931e-101
func (client *numberOperations) GetSmallDouble(ctx context.Context) (*Float64Response, error) {
	req, err := client.getSmallDoubleCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getSmallDoubleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getSmallDoubleCreateRequest creates the GetSmallDouble request.
func (client *numberOperations) getSmallDoubleCreateRequest() (*azcore.Request, error) {
	urlPath := "/number/small/double/2.5976931e-101"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getSmallDoubleHandleResponse handles the GetSmallDouble response.
func (client *numberOperations) getSmallDoubleHandleResponse(resp *azcore.Response) (*Float64Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := Float64Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetSmallFloat - Get big double value 3.402823e-20
func (client *numberOperations) GetSmallFloat(ctx context.Context) (*Float64Response, error) {
	req, err := client.getSmallFloatCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getSmallFloatHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getSmallFloatCreateRequest creates the GetSmallFloat request.
func (client *numberOperations) getSmallFloatCreateRequest() (*azcore.Request, error) {
	urlPath := "/number/small/float/3.402823e-20"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getSmallFloatHandleResponse handles the GetSmallFloat response.
func (client *numberOperations) getSmallFloatHandleResponse(resp *azcore.Response) (*Float64Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := Float64Response{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// PutBigDecimal - Put big decimal value 2.5976931e+101
func (client *numberOperations) PutBigDecimal(ctx context.Context, numberBody float64) (*http.Response, error) {
	req, err := client.putBigDecimalCreateRequest(numberBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putBigDecimalHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putBigDecimalCreateRequest creates the PutBigDecimal request.
func (client *numberOperations) putBigDecimalCreateRequest(numberBody float64) (*azcore.Request, error) {
	urlPath := "/number/big/decimal/2.5976931e+101"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	err = req.MarshalAsJSON(numberBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putBigDecimalHandleResponse handles the PutBigDecimal response.
func (client *numberOperations) putBigDecimalHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return resp.Response, nil
}

// PutBigDecimalNegativeDecimal - Put big decimal value -99999999.99
func (client *numberOperations) PutBigDecimalNegativeDecimal(ctx context.Context) (*http.Response, error) {
	req, err := client.putBigDecimalNegativeDecimalCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putBigDecimalNegativeDecimalHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putBigDecimalNegativeDecimalCreateRequest creates the PutBigDecimalNegativeDecimal request.
func (client *numberOperations) putBigDecimalNegativeDecimalCreateRequest() (*azcore.Request, error) {
	urlPath := "/number/big/decimal/-99999999.99"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	err = req.MarshalAsJSON(-99999999.99)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putBigDecimalNegativeDecimalHandleResponse handles the PutBigDecimalNegativeDecimal response.
func (client *numberOperations) putBigDecimalNegativeDecimalHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return resp.Response, nil
}

// PutBigDecimalPositiveDecimal - Put big decimal value 99999999.99
func (client *numberOperations) PutBigDecimalPositiveDecimal(ctx context.Context) (*http.Response, error) {
	req, err := client.putBigDecimalPositiveDecimalCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putBigDecimalPositiveDecimalHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putBigDecimalPositiveDecimalCreateRequest creates the PutBigDecimalPositiveDecimal request.
func (client *numberOperations) putBigDecimalPositiveDecimalCreateRequest() (*azcore.Request, error) {
	urlPath := "/number/big/decimal/99999999.99"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	err = req.MarshalAsJSON(99999999.99)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putBigDecimalPositiveDecimalHandleResponse handles the PutBigDecimalPositiveDecimal response.
func (client *numberOperations) putBigDecimalPositiveDecimalHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return resp.Response, nil
}

// PutBigDouble - Put big double value 2.5976931e+101
func (client *numberOperations) PutBigDouble(ctx context.Context, numberBody float64) (*http.Response, error) {
	req, err := client.putBigDoubleCreateRequest(numberBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putBigDoubleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putBigDoubleCreateRequest creates the PutBigDouble request.
func (client *numberOperations) putBigDoubleCreateRequest(numberBody float64) (*azcore.Request, error) {
	urlPath := "/number/big/double/2.5976931e+101"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	err = req.MarshalAsJSON(numberBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putBigDoubleHandleResponse handles the PutBigDouble response.
func (client *numberOperations) putBigDoubleHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return resp.Response, nil
}

// PutBigDoubleNegativeDecimal - Put big double value -99999999.99
func (client *numberOperations) PutBigDoubleNegativeDecimal(ctx context.Context) (*http.Response, error) {
	req, err := client.putBigDoubleNegativeDecimalCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putBigDoubleNegativeDecimalHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putBigDoubleNegativeDecimalCreateRequest creates the PutBigDoubleNegativeDecimal request.
func (client *numberOperations) putBigDoubleNegativeDecimalCreateRequest() (*azcore.Request, error) {
	urlPath := "/number/big/double/-99999999.99"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	err = req.MarshalAsJSON(-99999999.99)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putBigDoubleNegativeDecimalHandleResponse handles the PutBigDoubleNegativeDecimal response.
func (client *numberOperations) putBigDoubleNegativeDecimalHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return resp.Response, nil
}

// PutBigDoublePositiveDecimal - Put big double value 99999999.99
func (client *numberOperations) PutBigDoublePositiveDecimal(ctx context.Context) (*http.Response, error) {
	req, err := client.putBigDoublePositiveDecimalCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putBigDoublePositiveDecimalHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putBigDoublePositiveDecimalCreateRequest creates the PutBigDoublePositiveDecimal request.
func (client *numberOperations) putBigDoublePositiveDecimalCreateRequest() (*azcore.Request, error) {
	urlPath := "/number/big/double/99999999.99"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	err = req.MarshalAsJSON(99999999.99)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putBigDoublePositiveDecimalHandleResponse handles the PutBigDoublePositiveDecimal response.
func (client *numberOperations) putBigDoublePositiveDecimalHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return resp.Response, nil
}

// PutBigFloat - Put big float value 3.402823e+20
func (client *numberOperations) PutBigFloat(ctx context.Context, numberBody float32) (*http.Response, error) {
	req, err := client.putBigFloatCreateRequest(numberBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putBigFloatHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putBigFloatCreateRequest creates the PutBigFloat request.
func (client *numberOperations) putBigFloatCreateRequest(numberBody float32) (*azcore.Request, error) {
	urlPath := "/number/big/float/3.402823e+20"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	err = req.MarshalAsJSON(numberBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putBigFloatHandleResponse handles the PutBigFloat response.
func (client *numberOperations) putBigFloatHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return resp.Response, nil
}

// PutSmallDecimal - Put small decimal value 2.5976931e-101
func (client *numberOperations) PutSmallDecimal(ctx context.Context, numberBody float64) (*http.Response, error) {
	req, err := client.putSmallDecimalCreateRequest(numberBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putSmallDecimalHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putSmallDecimalCreateRequest creates the PutSmallDecimal request.
func (client *numberOperations) putSmallDecimalCreateRequest(numberBody float64) (*azcore.Request, error) {
	urlPath := "/number/small/decimal/2.5976931e-101"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	err = req.MarshalAsJSON(numberBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putSmallDecimalHandleResponse handles the PutSmallDecimal response.
func (client *numberOperations) putSmallDecimalHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return resp.Response, nil
}

// PutSmallDouble - Put small double value 2.5976931e-101
func (client *numberOperations) PutSmallDouble(ctx context.Context, numberBody float64) (*http.Response, error) {
	req, err := client.putSmallDoubleCreateRequest(numberBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putSmallDoubleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putSmallDoubleCreateRequest creates the PutSmallDouble request.
func (client *numberOperations) putSmallDoubleCreateRequest(numberBody float64) (*azcore.Request, error) {
	urlPath := "/number/small/double/2.5976931e-101"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	err = req.MarshalAsJSON(numberBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putSmallDoubleHandleResponse handles the PutSmallDouble response.
func (client *numberOperations) putSmallDoubleHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return resp.Response, nil
}

// PutSmallFloat - Put small float value 3.402823e-20
func (client *numberOperations) PutSmallFloat(ctx context.Context, numberBody float32) (*http.Response, error) {
	req, err := client.putSmallFloatCreateRequest(numberBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putSmallFloatHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putSmallFloatCreateRequest creates the PutSmallFloat request.
func (client *numberOperations) putSmallFloatCreateRequest(numberBody float32) (*azcore.Request, error) {
	urlPath := "/number/small/float/3.402823e-20"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	err = req.MarshalAsJSON(numberBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putSmallFloatHandleResponse handles the PutSmallFloat response.
func (client *numberOperations) putSmallFloatHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return resp.Response, nil
}
