// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountShareResourceService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountShareResourceService] method instead.
type AccountShareResourceService struct {
	Options []option.RequestOption
}

// NewAccountShareResourceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountShareResourceService(opts ...option.RequestOption) (r *AccountShareResourceService) {
	r = &AccountShareResourceService{}
	r.Options = opts
	return
}

// Create a new share resource
func (r *AccountShareResourceService) New(ctx context.Context, accountID string, shareID string, body AccountShareResourceNewParams, opts ...option.RequestOption) (res *ShareResourceResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/resources", accountID, shareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get share resource by ID.
func (r *AccountShareResourceService) Get(ctx context.Context, accountID string, shareID string, resourceID string, opts ...option.RequestOption) (res *ShareResourceResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	if resourceID == "" {
		err = errors.New("missing required resource_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/resources/%s", accountID, shareID, resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update is not immediate, an updated share resource object with a new status will
// be returned.
func (r *AccountShareResourceService) Update(ctx context.Context, accountID string, shareID string, resourceID string, body AccountShareResourceUpdateParams, opts ...option.RequestOption) (res *ShareResourceResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	if resourceID == "" {
		err = errors.New("missing required resource_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/resources/%s", accountID, shareID, resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List share resources by share ID.
func (r *AccountShareResourceService) List(ctx context.Context, accountID string, shareID string, query AccountShareResourceListParams, opts ...option.RequestOption) (res *AccountShareResourceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/resources", accountID, shareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletion is not immediate, an updated share resource object with a new status
// will be returned.
func (r *AccountShareResourceService) Delete(ctx context.Context, accountID string, shareID string, resourceID string, opts ...option.RequestOption) (res *ShareResourceResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	if resourceID == "" {
		err = errors.New("missing required resource_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/resources/%s", accountID, shareID, resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CreateShareResourceRequestParam struct {
	// Resource Metadata.
	Meta param.Field[interface{}] `json:"meta,required"`
	// Account identifier.
	ResourceAccountID param.Field[string] `json:"resource_account_id,required"`
	// Share Resource identifier.
	ResourceID param.Field[string] `json:"resource_id,required"`
	// Resource Type.
	ResourceType param.Field[ResourceType] `json:"resource_type,required"`
}

func (r CreateShareResourceRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Resource Type.
type ResourceType string

const (
	ResourceTypeCustomRuleset ResourceType = "custom-ruleset"
	ResourceTypeWidget        ResourceType = "widget"
)

func (r ResourceType) IsKnown() bool {
	switch r {
	case ResourceTypeCustomRuleset, ResourceTypeWidget:
		return true
	}
	return false
}

type ShareResourceObject struct {
	// Share Resource identifier.
	ID string `json:"id,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Resource Metadata.
	Meta interface{} `json:"meta,required"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Account identifier.
	ResourceAccountID string `json:"resource_account_id,required"`
	// Share Resource identifier.
	ResourceID string `json:"resource_id,required"`
	// Resource Type.
	ResourceType ResourceType `json:"resource_type,required"`
	// Resource Version.
	ResourceVersion int64 `json:"resource_version,required"`
	// Resource Status.
	Status ShareResourceObjectStatus `json:"status,required"`
	JSON   shareResourceObjectJSON   `json:"-"`
}

// shareResourceObjectJSON contains the JSON metadata for the struct
// [ShareResourceObject]
type shareResourceObjectJSON struct {
	ID                apijson.Field
	Created           apijson.Field
	Meta              apijson.Field
	Modified          apijson.Field
	ResourceAccountID apijson.Field
	ResourceID        apijson.Field
	ResourceType      apijson.Field
	ResourceVersion   apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ShareResourceObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r shareResourceObjectJSON) RawJSON() string {
	return r.raw
}

// Resource Status.
type ShareResourceObjectStatus string

const (
	ShareResourceObjectStatusActive   ShareResourceObjectStatus = "active"
	ShareResourceObjectStatusDeleting ShareResourceObjectStatus = "deleting"
	ShareResourceObjectStatusDeleted  ShareResourceObjectStatus = "deleted"
)

func (r ShareResourceObjectStatus) IsKnown() bool {
	switch r {
	case ShareResourceObjectStatusActive, ShareResourceObjectStatusDeleting, ShareResourceObjectStatusDeleted:
		return true
	}
	return false
}

type ShareResourceResponseSingle struct {
	Result ShareResourceObject             `json:"result"`
	JSON   shareResourceResponseSingleJSON `json:"-"`
	APIResponseResourceSharing
}

// shareResourceResponseSingleJSON contains the JSON metadata for the struct
// [ShareResourceResponseSingle]
type shareResourceResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ShareResourceResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r shareResourceResponseSingleJSON) RawJSON() string {
	return r.raw
}

type AccountShareResourceListResponse struct {
	Result []ShareResourceObject                `json:"result"`
	JSON   accountShareResourceListResponseJSON `json:"-"`
	APIResponseCollectionResourceSharing
}

// accountShareResourceListResponseJSON contains the JSON metadata for the struct
// [AccountShareResourceListResponse]
type accountShareResourceListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountShareResourceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountShareResourceListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountShareResourceNewParams struct {
	CreateShareResourceRequest CreateShareResourceRequestParam `json:"create_share_resource_request,required"`
}

func (r AccountShareResourceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateShareResourceRequest)
}

type AccountShareResourceUpdateParams struct {
	// Resource Metadata.
	Meta param.Field[interface{}] `json:"meta,required"`
}

func (r AccountShareResourceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountShareResourceListParams struct {
	// Page number.
	Page param.Field[int64] `query:"page"`
	// Number of objects to return per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter share resources by resource_type.
	ResourceType param.Field[ResourceType] `query:"resource_type"`
	// Filter share resources by status.
	Status param.Field[AccountShareResourceListParamsStatus] `query:"status"`
}

// URLQuery serializes [AccountShareResourceListParams]'s query parameters as
// `url.Values`.
func (r AccountShareResourceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter share resources by status.
type AccountShareResourceListParamsStatus string

const (
	AccountShareResourceListParamsStatusActive   AccountShareResourceListParamsStatus = "active"
	AccountShareResourceListParamsStatusDeleting AccountShareResourceListParamsStatus = "deleting"
	AccountShareResourceListParamsStatusDeleted  AccountShareResourceListParamsStatus = "deleted"
)

func (r AccountShareResourceListParamsStatus) IsKnown() bool {
	switch r {
	case AccountShareResourceListParamsStatusActive, AccountShareResourceListParamsStatusDeleting, AccountShareResourceListParamsStatusDeleted:
		return true
	}
	return false
}
