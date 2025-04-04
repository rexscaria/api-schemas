// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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
	Result DetailsResponseResult `json:"result"`
	JSON   detailsResponseJSON   `json:"-"`
	APIResponseSingleWeb3
}

// detailsResponseJSON contains the JSON metadata for the struct [DetailsResponse]
type detailsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detailsResponseJSON) RawJSON() string {
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

type ZoneWeb3HostnameIpfsUniversalPathContentListUpdateParams struct {
	// Behavior of the content list.
	Action param.Field[ActionBehavior] `json:"action,required"`
	// Content list entries.
	Entries param.Field[[]ContentListEntryParam] `json:"entries,required"`
}

func (r ZoneWeb3HostnameIpfsUniversalPathContentListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
