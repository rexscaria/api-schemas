// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountCniSlotService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCniSlotService] method instead.
type AccountCniSlotService struct {
	Options []option.RequestOption
}

// NewAccountCniSlotService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountCniSlotService(opts ...option.RequestOption) (r *AccountCniSlotService) {
	r = &AccountCniSlotService{}
	r.Options = opts
	return
}

// Get information about the specified slot
func (r *AccountCniSlotService) Get(ctx context.Context, accountID string, slot string, opts ...option.RequestOption) (res *NscSlotInfo, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if slot == "" {
		err = errors.New("missing required slot parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/slots/%s", accountID, slot)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a list of all slots matching the specified parameters
func (r *AccountCniSlotService) List(ctx context.Context, accountID string, query AccountCniSlotListParams, opts ...option.RequestOption) (res *AccountCniSlotListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/slots", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type NscFacilityInfo struct {
	Address []string            `json:"address,required"`
	Name    string              `json:"name,required"`
	JSON    nscFacilityInfoJSON `json:"-"`
}

// nscFacilityInfoJSON contains the JSON metadata for the struct [NscFacilityInfo]
type nscFacilityInfoJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NscFacilityInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nscFacilityInfoJSON) RawJSON() string {
	return r.raw
}

type NscSlotInfo struct {
	// Slot ID
	ID       string          `json:"id,required" format:"uuid"`
	Facility NscFacilityInfo `json:"facility,required"`
	// Whether the slot is occupied or not
	Occupied bool   `json:"occupied,required"`
	Site     string `json:"site,required"`
	Speed    string `json:"speed,required"`
	// Customer account tag
	Account string          `json:"account"`
	JSON    nscSlotInfoJSON `json:"-"`
}

// nscSlotInfoJSON contains the JSON metadata for the struct [NscSlotInfo]
type nscSlotInfoJSON struct {
	ID          apijson.Field
	Facility    apijson.Field
	Occupied    apijson.Field
	Site        apijson.Field
	Speed       apijson.Field
	Account     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NscSlotInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nscSlotInfoJSON) RawJSON() string {
	return r.raw
}

type AccountCniSlotListResponse struct {
	Items []NscSlotInfo                  `json:"items,required"`
	Next  int64                          `json:"next,nullable"`
	JSON  accountCniSlotListResponseJSON `json:"-"`
}

// accountCniSlotListResponseJSON contains the JSON metadata for the struct
// [AccountCniSlotListResponse]
type accountCniSlotListResponseJSON struct {
	Items       apijson.Field
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCniSlotListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCniSlotListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCniSlotListParams struct {
	// If specified, only show slots with the given text in their address field
	AddressContains param.Field[string] `query:"address_contains"`
	Cursor          param.Field[int64]  `query:"cursor"`
	Limit           param.Field[int64]  `query:"limit"`
	// If specified, only show slots with a specific occupied/unoccupied state
	Occupied param.Field[bool] `query:"occupied"`
	// If specified, only show slots located at the given site
	Site param.Field[string] `query:"site"`
	// If specified, only show slots that support the given speed
	Speed param.Field[string] `query:"speed"`
}

// URLQuery serializes [AccountCniSlotListParams]'s query parameters as
// `url.Values`.
func (r AccountCniSlotListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
