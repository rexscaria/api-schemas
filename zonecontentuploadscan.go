// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneContentUploadScanService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneContentUploadScanService] method instead.
type ZoneContentUploadScanService struct {
	Options  []option.RequestOption
	Payloads *ZoneContentUploadScanPayloadService
}

// NewZoneContentUploadScanService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneContentUploadScanService(opts ...option.RequestOption) (r *ZoneContentUploadScanService) {
	r = &ZoneContentUploadScanService{}
	r.Options = opts
	r.Payloads = NewZoneContentUploadScanPayloadService(opts...)
	return
}

// Disable Content Scanning.
func (r *ZoneContentUploadScanService) Disable(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *APIResponseCommon, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/disable", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Enable Content Scanning.
func (r *ZoneContentUploadScanService) Enable(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *APIResponseCommon, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/enable", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Retrieve the current status of Content Scanning.
func (r *ZoneContentUploadScanService) GetStatus(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneContentUploadScanGetStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type APIResponseCommon struct {
	Errors   []WafProductAPIBundleMessages `json:"errors,required"`
	Messages []WafProductAPIBundleMessages `json:"messages,required"`
	Result   interface{}                   `json:"result,required"`
	// Whether the API call was successful.
	Success APIResponseCommonSuccess `json:"success,required"`
	JSON    apiResponseCommonJSON    `json:"-"`
}

// apiResponseCommonJSON contains the JSON metadata for the struct
// [APIResponseCommon]
type apiResponseCommonJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCommonJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type APIResponseCommonSuccess bool

const (
	APIResponseCommonSuccessTrue APIResponseCommonSuccess = true
)

func (r APIResponseCommonSuccess) IsKnown() bool {
	switch r {
	case APIResponseCommonSuccessTrue:
		return true
	}
	return false
}

type WafProductAPIBundleMessages struct {
	Code             int64                             `json:"code,required"`
	Message          string                            `json:"message,required"`
	DocumentationURL string                            `json:"documentation_url"`
	Source           WafProductAPIBundleMessagesSource `json:"source"`
	JSON             wafProductAPIBundleMessagesJSON   `json:"-"`
}

// wafProductAPIBundleMessagesJSON contains the JSON metadata for the struct
// [WafProductAPIBundleMessages]
type wafProductAPIBundleMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WafProductAPIBundleMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafProductAPIBundleMessagesJSON) RawJSON() string {
	return r.raw
}

type WafProductAPIBundleMessagesSource struct {
	Pointer string                                `json:"pointer"`
	JSON    wafProductAPIBundleMessagesSourceJSON `json:"-"`
}

// wafProductAPIBundleMessagesSourceJSON contains the JSON metadata for the struct
// [WafProductAPIBundleMessagesSource]
type wafProductAPIBundleMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WafProductAPIBundleMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafProductAPIBundleMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneContentUploadScanGetStatusResponse struct {
	Errors   []WafProductAPIBundleMessages `json:"errors,required"`
	Messages []WafProductAPIBundleMessages `json:"messages,required"`
	// Defines the status for Content Scanning.
	Result ZoneContentUploadScanGetStatusResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneContentUploadScanGetStatusResponseSuccess `json:"success,required"`
	JSON    zoneContentUploadScanGetStatusResponseJSON    `json:"-"`
}

// zoneContentUploadScanGetStatusResponseJSON contains the JSON metadata for the
// struct [ZoneContentUploadScanGetStatusResponse]
type zoneContentUploadScanGetStatusResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneContentUploadScanGetStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneContentUploadScanGetStatusResponseJSON) RawJSON() string {
	return r.raw
}

// Defines the status for Content Scanning.
type ZoneContentUploadScanGetStatusResponseResult struct {
	// Defines the last modification date (ISO 8601) of the Content Scanning status.
	Modified string `json:"modified"`
	// Defines the status of Content Scanning.
	Value string                                           `json:"value"`
	JSON  zoneContentUploadScanGetStatusResponseResultJSON `json:"-"`
}

// zoneContentUploadScanGetStatusResponseResultJSON contains the JSON metadata for
// the struct [ZoneContentUploadScanGetStatusResponseResult]
type zoneContentUploadScanGetStatusResponseResultJSON struct {
	Modified    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneContentUploadScanGetStatusResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneContentUploadScanGetStatusResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneContentUploadScanGetStatusResponseSuccess bool

const (
	ZoneContentUploadScanGetStatusResponseSuccessTrue ZoneContentUploadScanGetStatusResponseSuccess = true
)

func (r ZoneContentUploadScanGetStatusResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneContentUploadScanGetStatusResponseSuccessTrue:
		return true
	}
	return false
}
