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

// AccountCloudforceOneEventCategoryService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCloudforceOneEventCategoryService] method instead.
type AccountCloudforceOneEventCategoryService struct {
	Options []option.RequestOption
}

// NewAccountCloudforceOneEventCategoryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCloudforceOneEventCategoryService(opts ...option.RequestOption) (r *AccountCloudforceOneEventCategoryService) {
	r = &AccountCloudforceOneEventCategoryService{}
	r.Options = opts
	return
}

// Creates a new category
func (r *AccountCloudforceOneEventCategoryService) New(ctx context.Context, accountID float64, body AccountCloudforceOneEventCategoryNewParams, opts ...option.RequestOption) (res *AccountCloudforceOneEventCategoryNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/categories/create", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Reads a category
func (r *AccountCloudforceOneEventCategoryService) Get(ctx context.Context, accountID float64, categoryID string, opts ...option.RequestOption) (res *AccountCloudforceOneEventCategoryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if categoryID == "" {
		err = errors.New("missing required category_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/categories/%s", accountID, categoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a category
func (r *AccountCloudforceOneEventCategoryService) Update(ctx context.Context, accountID float64, categoryID string, body AccountCloudforceOneEventCategoryUpdateParams, opts ...option.RequestOption) (res *AccountCloudforceOneEventCategoryUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if categoryID == "" {
		err = errors.New("missing required category_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/categories/%s", accountID, categoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists categories
func (r *AccountCloudforceOneEventCategoryService) List(ctx context.Context, accountID float64, opts ...option.RequestOption) (res *[]AccountCloudforceOneEventCategoryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/categories", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a category
func (r *AccountCloudforceOneEventCategoryService) Delete(ctx context.Context, accountID float64, categoryID string, opts ...option.RequestOption) (res *AccountCloudforceOneEventCategoryDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if categoryID == "" {
		err = errors.New("missing required category_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/categories/%s", accountID, categoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccountCloudforceOneEventCategoryNewResponse struct {
	KillChain   float64                                          `json:"killChain,required"`
	Name        string                                           `json:"name,required"`
	Uuid        string                                           `json:"uuid,required"`
	MitreAttack []string                                         `json:"mitreAttack"`
	Shortname   string                                           `json:"shortname"`
	JSON        accountCloudforceOneEventCategoryNewResponseJSON `json:"-"`
}

// accountCloudforceOneEventCategoryNewResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneEventCategoryNewResponse]
type accountCloudforceOneEventCategoryNewResponseJSON struct {
	KillChain   apijson.Field
	Name        apijson.Field
	Uuid        apijson.Field
	MitreAttack apijson.Field
	Shortname   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventCategoryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventCategoryNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventCategoryGetResponse struct {
	KillChain   float64                                          `json:"killChain,required"`
	Name        string                                           `json:"name,required"`
	Uuid        string                                           `json:"uuid,required"`
	MitreAttack []string                                         `json:"mitreAttack"`
	Shortname   string                                           `json:"shortname"`
	JSON        accountCloudforceOneEventCategoryGetResponseJSON `json:"-"`
}

// accountCloudforceOneEventCategoryGetResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneEventCategoryGetResponse]
type accountCloudforceOneEventCategoryGetResponseJSON struct {
	KillChain   apijson.Field
	Name        apijson.Field
	Uuid        apijson.Field
	MitreAttack apijson.Field
	Shortname   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventCategoryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventCategoryGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventCategoryUpdateResponse struct {
	KillChain   float64                                             `json:"killChain,required"`
	Name        string                                              `json:"name,required"`
	Uuid        string                                              `json:"uuid,required"`
	MitreAttack []string                                            `json:"mitreAttack"`
	Shortname   string                                              `json:"shortname"`
	JSON        accountCloudforceOneEventCategoryUpdateResponseJSON `json:"-"`
}

// accountCloudforceOneEventCategoryUpdateResponseJSON contains the JSON metadata
// for the struct [AccountCloudforceOneEventCategoryUpdateResponse]
type accountCloudforceOneEventCategoryUpdateResponseJSON struct {
	KillChain   apijson.Field
	Name        apijson.Field
	Uuid        apijson.Field
	MitreAttack apijson.Field
	Shortname   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventCategoryUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventCategoryUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventCategoryListResponse struct {
	KillChain   float64                                           `json:"killChain,required"`
	Name        string                                            `json:"name,required"`
	Uuid        string                                            `json:"uuid,required"`
	MitreAttack []string                                          `json:"mitreAttack"`
	Shortname   string                                            `json:"shortname"`
	JSON        accountCloudforceOneEventCategoryListResponseJSON `json:"-"`
}

// accountCloudforceOneEventCategoryListResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneEventCategoryListResponse]
type accountCloudforceOneEventCategoryListResponseJSON struct {
	KillChain   apijson.Field
	Name        apijson.Field
	Uuid        apijson.Field
	MitreAttack apijson.Field
	Shortname   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventCategoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventCategoryListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventCategoryDeleteResponse struct {
	Uuid string                                              `json:"uuid,required"`
	JSON accountCloudforceOneEventCategoryDeleteResponseJSON `json:"-"`
}

// accountCloudforceOneEventCategoryDeleteResponseJSON contains the JSON metadata
// for the struct [AccountCloudforceOneEventCategoryDeleteResponse]
type accountCloudforceOneEventCategoryDeleteResponseJSON struct {
	Uuid        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventCategoryDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventCategoryDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventCategoryNewParams struct {
	KillChain   param.Field[float64]  `json:"killChain,required"`
	Name        param.Field[string]   `json:"name,required"`
	MitreAttack param.Field[[]string] `json:"mitreAttack"`
	Shortname   param.Field[string]   `json:"shortname"`
}

func (r AccountCloudforceOneEventCategoryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCloudforceOneEventCategoryUpdateParams struct {
	KillChain   param.Field[float64]  `json:"killChain"`
	MitreAttack param.Field[[]string] `json:"mitreAttack"`
	Name        param.Field[string]   `json:"name"`
	Shortname   param.Field[string]   `json:"shortname"`
}

func (r AccountCloudforceOneEventCategoryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
