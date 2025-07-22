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

// AccountAccessTagService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessTagService] method instead.
type AccountAccessTagService struct {
	Options []option.RequestOption
}

// NewAccountAccessTagService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessTagService(opts ...option.RequestOption) (r *AccountAccessTagService) {
	r = &AccountAccessTagService{}
	r.Options = opts
	return
}

// Create a tag
func (r *AccountAccessTagService) New(ctx context.Context, accountID string, body AccountAccessTagNewParams, opts ...option.RequestOption) (res *SingleResponseTag, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/tags", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a tag
func (r *AccountAccessTagService) Get(ctx context.Context, accountID string, tagName string, opts ...option.RequestOption) (res *SingleResponseTag, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tagName == "" {
		err = errors.New("missing required tag_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/tags/%s", accountID, tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a tag
func (r *AccountAccessTagService) Update(ctx context.Context, accountID string, tagName string, body AccountAccessTagUpdateParams, opts ...option.RequestOption) (res *SingleResponseTag, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tagName == "" {
		err = errors.New("missing required tag_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/tags/%s", accountID, tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List tags
func (r *AccountAccessTagService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAccessTagListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/tags", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a tag
func (r *AccountAccessTagService) Delete(ctx context.Context, accountID string, tagName string, opts ...option.RequestOption) (res *AccountAccessTagDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tagName == "" {
		err = errors.New("missing required tag_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/tags/%s", accountID, tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type SingleResponseTag struct {
	// A tag
	Result Tag                   `json:"result"`
	JSON   singleResponseTagJSON `json:"-"`
	APIResponseSingleAccess
}

// singleResponseTagJSON contains the JSON metadata for the struct
// [SingleResponseTag]
type singleResponseTagJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseTagJSON) RawJSON() string {
	return r.raw
}

// A tag
type Tag struct {
	// The name of the tag
	Name string `json:"name,required"`
	// The number of applications that have this tag
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	JSON      tagJSON   `json:"-"`
}

// tagJSON contains the JSON metadata for the struct [Tag]
type tagJSON struct {
	Name        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Tag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tagJSON) RawJSON() string {
	return r.raw
}

type AccountAccessTagListResponse struct {
	Result []Tag                            `json:"result"`
	JSON   accountAccessTagListResponseJSON `json:"-"`
	APIResponseCollectionAccess
}

// accountAccessTagListResponseJSON contains the JSON metadata for the struct
// [AccountAccessTagListResponse]
type accountAccessTagListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessTagListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessTagDeleteResponse struct {
	Result AccountAccessTagDeleteResponseResult `json:"result"`
	JSON   accountAccessTagDeleteResponseJSON   `json:"-"`
	APIResponseSingleAccess
}

// accountAccessTagDeleteResponseJSON contains the JSON metadata for the struct
// [AccountAccessTagDeleteResponse]
type accountAccessTagDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessTagDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessTagDeleteResponseResult struct {
	// The name of the tag
	Name string                                   `json:"name"`
	JSON accountAccessTagDeleteResponseResultJSON `json:"-"`
}

// accountAccessTagDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessTagDeleteResponseResult]
type accountAccessTagDeleteResponseResultJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessTagDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessTagDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAccessTagNewParams struct {
	// The name of the tag
	Name param.Field[string] `json:"name"`
}

func (r AccountAccessTagNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessTagUpdateParams struct {
	// The name of the tag
	Name      param.Field[string]    `json:"name,required"`
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	UpdatedAt param.Field[time.Time] `json:"updated_at" format:"date-time"`
}

func (r AccountAccessTagUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
