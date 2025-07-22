// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ScanService contains methods and other services that help with interacting with
// the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScanService] method instead.
type ScanService struct {
	Options []option.RequestOption
}

// NewScanService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScanService(opts ...option.RequestOption) (r *ScanService) {
	r = &ScanService{}
	r.Options = opts
	return
}

// Fetches All Due Scans
func (r *ScanService) List(ctx context.Context, query ScanListParams, opts ...option.RequestOption) (res *ScanListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "scans"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Mark a Scan for Retry
func (r *ScanService) Retry(ctx context.Context, body ScanRetryParams, opts ...option.RequestOption) (res *ScanRetryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "scans"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ScanListResponse struct {
	Errors   []string               `json:"errors,required"`
	Messages []string               `json:"messages,required"`
	Result   ScanListResponseResult `json:"result,required"`
	Success  bool                   `json:"success,required"`
	JSON     scanListResponseJSON   `json:"-"`
}

// scanListResponseJSON contains the JSON metadata for the struct
// [ScanListResponse]
type scanListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanListResponseJSON) RawJSON() string {
	return r.raw
}

type ScanListResponseResult struct {
	AccountID string                     `json:"account_id,required"`
	IPs       []string                   `json:"ips,required"`
	JSON      scanListResponseResultJSON `json:"-"`
}

// scanListResponseResultJSON contains the JSON metadata for the struct
// [ScanListResponseResult]
type scanListResponseResultJSON struct {
	AccountID   apijson.Field
	IPs         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanListResponseResultJSON) RawJSON() string {
	return r.raw
}

type ScanRetryResponse struct {
	Errors   []string              `json:"errors,required"`
	Messages []string              `json:"messages,required"`
	Result   interface{}           `json:"result,required"`
	Success  bool                  `json:"success,required"`
	JSON     scanRetryResponseJSON `json:"-"`
}

// scanRetryResponseJSON contains the JSON metadata for the struct
// [ScanRetryResponse]
type scanRetryResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanRetryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanRetryResponseJSON) RawJSON() string {
	return r.raw
}

type ScanListParams struct {
	// Toggle IP Authorization
	VerifyIPs param.Field[bool] `query:"verifyIPs"`
}

// URLQuery serializes [ScanListParams]'s query parameters as `url.Values`.
func (r ScanListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScanRetryParams struct {
	AccountID param.Field[string] `json:"account_id,required"`
}

func (r ScanRetryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
