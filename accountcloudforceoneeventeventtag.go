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

// AccountCloudforceOneEventEventTagService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCloudforceOneEventEventTagService] method instead.
type AccountCloudforceOneEventEventTagService struct {
	Options []option.RequestOption
}

// NewAccountCloudforceOneEventEventTagService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCloudforceOneEventEventTagService(opts ...option.RequestOption) (r *AccountCloudforceOneEventEventTagService) {
	r = &AccountCloudforceOneEventEventTagService{}
	r.Options = opts
	return
}

// Adds a tag to an event
func (r *AccountCloudforceOneEventEventTagService) Add(ctx context.Context, accountID float64, eventID string, body AccountCloudforceOneEventEventTagAddParams, opts ...option.RequestOption) (res *AccountCloudforceOneEventEventTagAddResponse, err error) {
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/event_tag/%s/create", accountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Removes a tag from an event
func (r *AccountCloudforceOneEventEventTagService) Remove(ctx context.Context, accountID float64, eventID string, body AccountCloudforceOneEventEventTagRemoveParams, opts ...option.RequestOption) (res *AccountCloudforceOneEventEventTagRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/event_tag/%s", accountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type AccountCloudforceOneEventEventTagAddResponse struct {
	Result  AccountCloudforceOneEventEventTagAddResponseResult `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    accountCloudforceOneEventEventTagAddResponseJSON   `json:"-"`
}

// accountCloudforceOneEventEventTagAddResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneEventEventTagAddResponse]
type accountCloudforceOneEventEventTagAddResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventEventTagAddResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventEventTagAddResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventEventTagAddResponseResult struct {
	Success bool                                                   `json:"success,required"`
	JSON    accountCloudforceOneEventEventTagAddResponseResultJSON `json:"-"`
}

// accountCloudforceOneEventEventTagAddResponseResultJSON contains the JSON
// metadata for the struct [AccountCloudforceOneEventEventTagAddResponseResult]
type accountCloudforceOneEventEventTagAddResponseResultJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventEventTagAddResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventEventTagAddResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventEventTagRemoveResponse struct {
	Result  AccountCloudforceOneEventEventTagRemoveResponseResult `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    accountCloudforceOneEventEventTagRemoveResponseJSON   `json:"-"`
}

// accountCloudforceOneEventEventTagRemoveResponseJSON contains the JSON metadata
// for the struct [AccountCloudforceOneEventEventTagRemoveResponse]
type accountCloudforceOneEventEventTagRemoveResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventEventTagRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventEventTagRemoveResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventEventTagRemoveResponseResult struct {
	Success bool                                                      `json:"success,required"`
	JSON    accountCloudforceOneEventEventTagRemoveResponseResultJSON `json:"-"`
}

// accountCloudforceOneEventEventTagRemoveResponseResultJSON contains the JSON
// metadata for the struct [AccountCloudforceOneEventEventTagRemoveResponseResult]
type accountCloudforceOneEventEventTagRemoveResponseResultJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventEventTagRemoveResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventEventTagRemoveResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventEventTagAddParams struct {
	Tags param.Field[[]string] `json:"tags,required"`
}

func (r AccountCloudforceOneEventEventTagAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCloudforceOneEventEventTagRemoveParams struct {
	Tags param.Field[[]string] `json:"tags,required"`
}

func (r AccountCloudforceOneEventEventTagRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
