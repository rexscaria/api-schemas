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

// ZoneCustomNService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCustomNService] method instead.
type ZoneCustomNService struct {
	Options []option.RequestOption
}

// NewZoneCustomNService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneCustomNService(opts ...option.RequestOption) (r *ZoneCustomNService) {
	r = &ZoneCustomNService{}
	r.Options = opts
	return
}

// Get metadata for account-level custom nameservers on a zone.
//
// Deprecated in favor of
// [Show DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-a-zone-list-dns-settings).
//
// Deprecated: deprecated
func (r *ZoneCustomNService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneCustomNGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_ns", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set metadata for account-level custom nameservers on a zone.
//
// If you would like new zones in the account to use account custom nameservers by
// default, use PUT /accounts/:identifier to set the account setting
// use_account_custom_ns_by_default to true.
//
// Deprecated in favor of
// [Update DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-a-zone-update-dns-settings).
//
// Deprecated: deprecated
func (r *ZoneCustomNService) Update(ctx context.Context, zoneID string, body ZoneCustomNUpdateParams, opts ...option.RequestOption) (res *ZoneCustomNUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_ns", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneMetadataParam struct {
	// Whether zone uses account-level custom nameservers.
	Enabled param.Field[bool] `json:"enabled"`
	// The number of the name server set to assign to the zone.
	NsSet param.Field[float64] `json:"ns_set"`
}

func (r ZoneMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneCustomNGetResponse struct {
	Errors   []ZoneCustomNGetResponseError   `json:"errors,required"`
	Messages []ZoneCustomNGetResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneCustomNGetResponseSuccess `json:"success,required"`
	// Whether zone uses account-level custom nameservers.
	Enabled bool `json:"enabled"`
	// The number of the name server set to assign to the zone.
	NsSet      float64                          `json:"ns_set"`
	ResultInfo ZoneCustomNGetResponseResultInfo `json:"result_info"`
	JSON       zoneCustomNGetResponseJSON       `json:"-"`
}

// zoneCustomNGetResponseJSON contains the JSON metadata for the struct
// [ZoneCustomNGetResponse]
type zoneCustomNGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Enabled     apijson.Field
	NsSet       apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomNGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCustomNGetResponseError struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           ZoneCustomNGetResponseErrorsSource `json:"source"`
	JSON             zoneCustomNGetResponseErrorJSON    `json:"-"`
}

// zoneCustomNGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneCustomNGetResponseError]
type zoneCustomNGetResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneCustomNGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomNGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ZoneCustomNGetResponseErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    zoneCustomNGetResponseErrorsSourceJSON `json:"-"`
}

// zoneCustomNGetResponseErrorsSourceJSON contains the JSON metadata for the struct
// [ZoneCustomNGetResponseErrorsSource]
type zoneCustomNGetResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNGetResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomNGetResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneCustomNGetResponseMessage struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           ZoneCustomNGetResponseMessagesSource `json:"source"`
	JSON             zoneCustomNGetResponseMessageJSON    `json:"-"`
}

// zoneCustomNGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneCustomNGetResponseMessage]
type zoneCustomNGetResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneCustomNGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomNGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ZoneCustomNGetResponseMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    zoneCustomNGetResponseMessagesSourceJSON `json:"-"`
}

// zoneCustomNGetResponseMessagesSourceJSON contains the JSON metadata for the
// struct [ZoneCustomNGetResponseMessagesSource]
type zoneCustomNGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomNGetResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneCustomNGetResponseSuccess bool

const (
	ZoneCustomNGetResponseSuccessTrue ZoneCustomNGetResponseSuccess = true
)

func (r ZoneCustomNGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneCustomNGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneCustomNGetResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                              `json:"total_count"`
	JSON       zoneCustomNGetResponseResultInfoJSON `json:"-"`
}

// zoneCustomNGetResponseResultInfoJSON contains the JSON metadata for the struct
// [ZoneCustomNGetResponseResultInfo]
type zoneCustomNGetResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNGetResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomNGetResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneCustomNUpdateResponse struct {
	Errors   []CustomNsMessages `json:"errors,required"`
	Messages []CustomNsMessages `json:"messages,required"`
	// Whether the API call was successful
	Success    ZoneCustomNUpdateResponseSuccess    `json:"success,required"`
	Result     []string                            `json:"result"`
	ResultInfo ZoneCustomNUpdateResponseResultInfo `json:"result_info"`
	JSON       zoneCustomNUpdateResponseJSON       `json:"-"`
}

// zoneCustomNUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneCustomNUpdateResponse]
type zoneCustomNUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomNUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneCustomNUpdateResponseSuccess bool

const (
	ZoneCustomNUpdateResponseSuccessTrue ZoneCustomNUpdateResponseSuccess = true
)

func (r ZoneCustomNUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneCustomNUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneCustomNUpdateResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count"`
	JSON       zoneCustomNUpdateResponseResultInfoJSON `json:"-"`
}

// zoneCustomNUpdateResponseResultInfoJSON contains the JSON metadata for the
// struct [ZoneCustomNUpdateResponseResultInfo]
type zoneCustomNUpdateResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNUpdateResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomNUpdateResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneCustomNUpdateParams struct {
	ZoneMetadata ZoneMetadataParam `json:"zone_metadata,required"`
}

func (r ZoneCustomNUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ZoneMetadata)
}
