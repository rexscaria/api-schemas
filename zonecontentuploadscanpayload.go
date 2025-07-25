// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneContentUploadScanPayloadService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneContentUploadScanPayloadService] method instead.
type ZoneContentUploadScanPayloadService struct {
	Options []option.RequestOption
}

// NewZoneContentUploadScanPayloadService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneContentUploadScanPayloadService(opts ...option.RequestOption) (r *ZoneContentUploadScanPayloadService) {
	r = &ZoneContentUploadScanPayloadService{}
	r.Options = opts
	return
}

// Add custom scan expressions for Content Scanning.
func (r *ZoneContentUploadScanPayloadService) New(ctx context.Context, zoneID string, body ZoneContentUploadScanPayloadNewParams, opts ...option.RequestOption) (res *CustomScanCollection, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/payloads", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a list of existing custom scan expressions for Content Scanning.
func (r *ZoneContentUploadScanPayloadService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *CustomScanCollection, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/payloads", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a Content Scan Custom Expression.
func (r *ZoneContentUploadScanPayloadService) Delete(ctx context.Context, zoneID string, expressionID CustomScanIDParam, opts ...option.RequestOption) (res *CustomScanCollection, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if expressionID == "" {
		err = errors.New("missing required expression_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/payloads/%s", zoneID, expressionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CustomScanCollection struct {
	Errors   []WafProductAPIBundleMessages `json:"errors,required"`
	Messages []WafProductAPIBundleMessages `json:"messages,required"`
	Result   []CustomScanCollectionResult  `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success CustomScanCollectionSuccess `json:"success,required"`
	JSON    customScanCollectionJSON    `json:"-"`
}

// customScanCollectionJSON contains the JSON metadata for the struct
// [CustomScanCollection]
type customScanCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomScanCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customScanCollectionJSON) RawJSON() string {
	return r.raw
}

// Defines a custom scan expression to match Content Scanning on.
type CustomScanCollectionResult struct {
	// defines the unique ID for this custom scan expression.
	ID CustomScanID `json:"id"`
	// Defines the ruleset expression to use in matching content objects.
	Payload string                         `json:"payload"`
	JSON    customScanCollectionResultJSON `json:"-"`
}

// customScanCollectionResultJSON contains the JSON metadata for the struct
// [CustomScanCollectionResult]
type customScanCollectionResultJSON struct {
	ID          apijson.Field
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomScanCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customScanCollectionResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CustomScanCollectionSuccess bool

const (
	CustomScanCollectionSuccessTrue CustomScanCollectionSuccess = true
)

func (r CustomScanCollectionSuccess) IsKnown() bool {
	switch r {
	case CustomScanCollectionSuccessTrue:
		return true
	}
	return false
}

type CustomScanID = string

type CustomScanIDParam = string

type ZoneContentUploadScanPayloadNewParams struct {
	Body []ZoneContentUploadScanPayloadNewParamsBody `json:"body,required"`
}

func (r ZoneContentUploadScanPayloadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneContentUploadScanPayloadNewParamsBody struct {
	// Defines the ruleset expression to use in matching content objects.
	Payload param.Field[string] `json:"payload,required"`
}

func (r ZoneContentUploadScanPayloadNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
