// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package datetimegroup

import (
	"context"
	"generatortests/autorest/generated/datetimerfc1123group"
	"generatortests/helpers"
	"net/http"
	"testing"
	"time"
)

func getDatetimerfc1123Operations(t *testing.T) datetimerfc1123group.Datetimerfc1123Operations {
	client, err := datetimerfc1123group.NewClient("http://localhost:3000", nil)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	return client.Datetimerfc1123Operations()
}

func TestGetInvalid(t *testing.T) {
	client := getDatetimerfc1123Operations(t)
	_, err := client.GetInvalid(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}

func TestGetNull(t *testing.T) {
	client := getDatetimerfc1123Operations(t)
	result, err := client.GetNull(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	var expected *time.Time
	helpers.DeepEqualOrFatal(t, result.Value, expected)
}

func TestGetOverflow(t *testing.T) {
	client := getDatetimerfc1123Operations(t)
	_, err := client.GetOverflow(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}

// GetUTCLowercaseMaxDateTime - Get max datetime value fri, 31 dec 9999 23:59:59 gmt
func TestGetUTCLowercaseMaxDateTime(t *testing.T) {
	client := getDatetimerfc1123Operations(t)
	result, err := client.GetUTCLowercaseMaxDateTime(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	expected, err := time.Parse(time.RFC1123, "Fri, 31 Dec 9999 23:59:59 GMT")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, &expected)
}

// GetUTCMinDateTime - Get min datetime value Mon, 1 Jan 0001 00:00:00 GMT
func TestGetUTCMinDateTime(t *testing.T) {
	client := getDatetimerfc1123Operations(t)
	result, err := client.GetUTCMinDateTime(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	expected, err := time.Parse(time.RFC1123, "Mon, 01 Jan 0001 00:00:00 GMT")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, &expected)
}

// GetUTCUppercaseMaxDateTime - Get max datetime value FRI, 31 DEC 9999 23:59:59 GMT
func TestGetUTCUppercaseMaxDateTime(t *testing.T) {
	client := getDatetimerfc1123Operations(t)
	result, err := client.GetUTCUppercaseMaxDateTime(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	expected, err := time.Parse(time.RFC1123, "FRI, 31 DEC 9999 23:59:59 GMT")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, &expected)
}

func TestGetUnderflow(t *testing.T) {
	client := getDatetimerfc1123Operations(t)
	_, err := client.GetUnderflow(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}

// PutUTCMaxDateTime - Put max datetime value Fri, 31 Dec 9999 23:59:59 GMT
func TestPutUTCMaxDateTime(t *testing.T) {
	client := getDatetimerfc1123Operations(t)
	body, err := time.Parse(time.RFC1123, "Fri, 31 Dec 9999 23:59:59 GMT")
	if err != nil {
		t.Fatal(err)
	}
	result, err := client.PutUTCMaxDateTime(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// PutUTCMinDateTime - Put min datetime value Mon, 1 Jan 0001 00:00:00 GMT
func TestPutUTCMinDateTime(t *testing.T) {
	client := getDatetimerfc1123Operations(t)
	body, err := time.Parse(time.RFC1123, "Mon, 01 Jan 0001 00:00:00 GMT")
	if err != nil {
		t.Fatal(err)
	}
	result, err := client.PutUTCMinDateTime(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}