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

// ZoneWeb3HostnameIpfsUniversalPathContentListService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneWeb3HostnameIpfsUniversalPathContentListService] method instead.
type ZoneWeb3HostnameIpfsUniversalPathContentListService struct {
	Options []option.RequestOption
	Entries *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService
}

// NewZoneWeb3HostnameIpfsUniversalPathContentListService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewZoneWeb3HostnameIpfsUniversalPathContentListService(opts ...option.RequestOption) (r *ZoneWeb3HostnameIpfsUniversalPathContentListService) {
	r = &ZoneWeb3HostnameIpfsUniversalPathContentListService{}
	r.Options = opts
	r.Entries = NewZoneWeb3HostnameIpfsUniversalPathContentListEntryService(opts...)
	return
}

// IPFS Universal Path Gateway Content List Details
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListService) Get(ctx context.Context, zoneID string, identifier string, opts ...option.RequestOption) (res *DetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list", zoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update IPFS Universal Path Gateway Content List
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListService) Update(ctx context.Context, zoneID string, identifier string, body ZoneWeb3HostnameIpfsUniversalPathContentListUpdateParams, opts ...option.RequestOption) (res *DetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list", zoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Behavior of the content list.
type ActionBehavior string

const (
	ActionBehaviorBlock ActionBehavior = "block"
)

func (r ActionBehavior) IsKnown() bool {
	switch r {
	case ActionBehaviorBlock:
		return true
	}
	return false
}

type DetailsResponse struct {
	Errors   []DetailsResponseError   `json:"errors,required"`
	Messages []DetailsResponseMessage `json:"messages,required"`
	Result   DetailsResponseResult    `json:"result,required"`
	// Specifies whether the API call was successful.
	Success DetailsResponseSuccess `json:"success,required"`
	// Provides the API response.
	ResultInfo interface{}         `json:"result_info"`
	JSON       detailsResponseJSON `json:"-"`
}

// detailsResponseJSON contains the JSON metadata for the struct [DetailsResponse]
type detailsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detailsResponseJSON) RawJSON() string {
	return r.raw
}

type DetailsResponseError struct {
	Code             int64                       `json:"code,required"`
	Message          string                      `json:"message,required"`
	DocumentationURL string                      `json:"documentation_url"`
	Source           DetailsResponseErrorsSource `json:"source"`
	JSON             detailsResponseErrorJSON    `json:"-"`
}

// detailsResponseErrorJSON contains the JSON metadata for the struct
// [DetailsResponseError]
type detailsResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DetailsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detailsResponseErrorJSON) RawJSON() string {
	return r.raw
}

type DetailsResponseErrorsSource struct {
	Pointer string                          `json:"pointer"`
	JSON    detailsResponseErrorsSourceJSON `json:"-"`
}

// detailsResponseErrorsSourceJSON contains the JSON metadata for the struct
// [DetailsResponseErrorsSource]
type detailsResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DetailsResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detailsResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DetailsResponseMessage struct {
	Code             int64                         `json:"code,required"`
	Message          string                        `json:"message,required"`
	DocumentationURL string                        `json:"documentation_url"`
	Source           DetailsResponseMessagesSource `json:"source"`
	JSON             detailsResponseMessageJSON    `json:"-"`
}

// detailsResponseMessageJSON contains the JSON metadata for the struct
// [DetailsResponseMessage]
type detailsResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DetailsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detailsResponseMessageJSON) RawJSON() string {
	return r.raw
}

type DetailsResponseMessagesSource struct {
	Pointer string                            `json:"pointer"`
	JSON    detailsResponseMessagesSourceJSON `json:"-"`
}

// detailsResponseMessagesSourceJSON contains the JSON metadata for the struct
// [DetailsResponseMessagesSource]
type detailsResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DetailsResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detailsResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type DetailsResponseResult struct {
	// Behavior of the content list.
	Action ActionBehavior            `json:"action"`
	JSON   detailsResponseResultJSON `json:"-"`
}

// detailsResponseResultJSON contains the JSON metadata for the struct
// [DetailsResponseResult]
type detailsResponseResultJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DetailsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detailsResponseResultJSON) RawJSON() string {
	return r.raw
}

// Specifies whether the API call was successful.
type DetailsResponseSuccess bool

const (
	DetailsResponseSuccessTrue DetailsResponseSuccess = true
)

func (r DetailsResponseSuccess) IsKnown() bool {
	switch r {
	case DetailsResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneWeb3HostnameIpfsUniversalPathContentListUpdateParams struct {
	// Behavior of the content list.
	Action param.Field[ActionBehavior] `json:"action,required"`
	// Provides content list entries.
	Entries param.Field[[]ContentListEntryParam] `json:"entries,required"`
}

func (r ZoneWeb3HostnameIpfsUniversalPathContentListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
