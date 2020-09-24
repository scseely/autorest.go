// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func newInheritanceClient() InheritanceOperations {
	return NewInheritanceClient(NewDefaultClient(nil))
}

func TestInheritanceGetValid(t *testing.T) {
	client := newInheritanceClient()
	result, err := client.GetValid(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.Siamese, &Siamese{
		Cat: Cat{
			Pet: Pet{
				ID:   to.Int32Ptr(2),
				Name: to.StringPtr("Siameeee"),
			},
			Color: to.StringPtr("green"),
			Hates: &[]Dog{
				{
					Pet: Pet{
						ID:   to.Int32Ptr(1),
						Name: to.StringPtr("Potato"),
					},
					Food: to.StringPtr("tomato"),
				},
				{
					Pet: Pet{
						ID:   to.Int32Ptr(-1),
						Name: to.StringPtr("Tomato"),
					},
					Food: to.StringPtr("french fries"),
				},
			},
		},
		Breed: to.StringPtr("persian"),
	})
}

func TestInheritancePutValid(t *testing.T) {
	client := newInheritanceClient()
	result, err := client.PutValid(context.Background(), Siamese{
		Cat: Cat{
			Pet: Pet{
				ID:   to.Int32Ptr(2),
				Name: to.StringPtr("Siameeee"),
			},
			Color: to.StringPtr("green"),
			Hates: &[]Dog{
				{
					Pet: Pet{
						ID:   to.Int32Ptr(1),
						Name: to.StringPtr("Potato"),
					},
					Food: to.StringPtr("tomato"),
				},
				{
					Pet: Pet{
						ID:   to.Int32Ptr(-1),
						Name: to.StringPtr("Tomato"),
					},
					Food: to.StringPtr("french fries"),
				},
			},
		},
		Breed: to.StringPtr("persian"),
	}, nil)
	if err != nil {
		t.Fatalf("PutValid: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}