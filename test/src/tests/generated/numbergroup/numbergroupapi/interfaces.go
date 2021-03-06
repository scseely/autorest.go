package numbergroupapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/shopspring/decimal"
	"tests/generated/numbergroup"
)

// NumberClientAPI contains the set of methods on the NumberClient type.
type NumberClientAPI interface {
	GetBigDecimal(ctx context.Context) (result numbergroup.Decimal, err error)
	GetBigDecimalNegativeDecimal(ctx context.Context) (result numbergroup.Decimal, err error)
	GetBigDecimalPositiveDecimal(ctx context.Context) (result numbergroup.Decimal, err error)
	GetBigDouble(ctx context.Context) (result numbergroup.Float64, err error)
	GetBigDoubleNegativeDecimal(ctx context.Context) (result numbergroup.Float64, err error)
	GetBigDoublePositiveDecimal(ctx context.Context) (result numbergroup.Float64, err error)
	GetBigFloat(ctx context.Context) (result numbergroup.Float64, err error)
	GetInvalidDecimal(ctx context.Context) (result numbergroup.Decimal, err error)
	GetInvalidDouble(ctx context.Context) (result numbergroup.Float64, err error)
	GetInvalidFloat(ctx context.Context) (result numbergroup.Float64, err error)
	GetNull(ctx context.Context) (result numbergroup.Float64, err error)
	GetSmallDecimal(ctx context.Context) (result numbergroup.Decimal, err error)
	GetSmallDouble(ctx context.Context) (result numbergroup.Float64, err error)
	GetSmallFloat(ctx context.Context) (result numbergroup.Float64, err error)
	PutBigDecimal(ctx context.Context, numberBody decimal.Decimal) (result autorest.Response, err error)
	PutBigDecimalNegativeDecimal(ctx context.Context) (result autorest.Response, err error)
	PutBigDecimalPositiveDecimal(ctx context.Context) (result autorest.Response, err error)
	PutBigDouble(ctx context.Context, numberBody float64) (result autorest.Response, err error)
	PutBigDoubleNegativeDecimal(ctx context.Context) (result autorest.Response, err error)
	PutBigDoublePositiveDecimal(ctx context.Context) (result autorest.Response, err error)
	PutBigFloat(ctx context.Context, numberBody float64) (result autorest.Response, err error)
	PutSmallDecimal(ctx context.Context, numberBody decimal.Decimal) (result autorest.Response, err error)
	PutSmallDouble(ctx context.Context, numberBody float64) (result autorest.Response, err error)
	PutSmallFloat(ctx context.Context, numberBody float64) (result autorest.Response, err error)
}

var _ NumberClientAPI = (*numbergroup.NumberClient)(nil)
