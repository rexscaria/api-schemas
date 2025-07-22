// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneAccessGroupService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAccessGroupService] method instead.
type ZoneAccessGroupService struct {
	Options []option.RequestOption
}

// NewZoneAccessGroupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneAccessGroupService(opts ...option.RequestOption) (r *ZoneAccessGroupService) {
	r = &ZoneAccessGroupService{}
	r.Options = opts
	return
}

// Creates a new Access group.
func (r *ZoneAccessGroupService) New(ctx context.Context, zoneID string, body ZoneAccessGroupNewParams, opts ...option.RequestOption) (res *SingleResponseGroupZone, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/groups", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single Access group.
func (r *ZoneAccessGroupService) Get(ctx context.Context, zoneID string, groupID string, opts ...option.RequestOption) (res *SingleResponseGroupZone, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/groups/%s", zoneID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured Access group.
func (r *ZoneAccessGroupService) Update(ctx context.Context, zoneID string, groupID string, body ZoneAccessGroupUpdateParams, opts ...option.RequestOption) (res *SingleResponseGroupZone, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/groups/%s", zoneID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all Access groups.
func (r *ZoneAccessGroupService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneAccessGroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/groups", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an Access group.
func (r *ZoneAccessGroupService) Delete(ctx context.Context, zoneID string, groupID string, opts ...option.RequestOption) (res *IDResponseApps, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/groups/%s", zoneID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccessComponentsGroups struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessRule `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessRule `json:"include"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccessRule               `json:"require"`
	UpdatedAt time.Time                  `json:"updated_at" format:"date-time"`
	JSON      accessComponentsGroupsJSON `json:"-"`
}

// accessComponentsGroupsJSON contains the JSON metadata for the struct
// [AccessComponentsGroups]
type accessComponentsGroupsJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Exclude     apijson.Field
	Include     apijson.Field
	Name        apijson.Field
	Require     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessComponentsGroups) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessComponentsGroupsJSON) RawJSON() string {
	return r.raw
}

type SingleResponseGroupZone struct {
	Result AccessComponentsGroups      `json:"result"`
	JSON   singleResponseGroupZoneJSON `json:"-"`
	APIResponseSingleAccess
}

// singleResponseGroupZoneJSON contains the JSON metadata for the struct
// [SingleResponseGroupZone]
type singleResponseGroupZoneJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseGroupZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseGroupZoneJSON) RawJSON() string {
	return r.raw
}

type ZoneAccessGroupListResponse struct {
	Result []AccessComponentsGroups        `json:"result"`
	JSON   zoneAccessGroupListResponseJSON `json:"-"`
	APIResponseCollectionAccess
}

// zoneAccessGroupListResponseJSON contains the JSON metadata for the struct
// [ZoneAccessGroupListResponse]
type zoneAccessGroupListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAccessGroupListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAccessGroupNewParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
}

func (r ZoneAccessGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessGroupUpdateParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
}

func (r ZoneAccessGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
