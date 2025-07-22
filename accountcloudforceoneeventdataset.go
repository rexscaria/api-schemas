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

// AccountCloudforceOneEventDatasetService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCloudforceOneEventDatasetService] method instead.
type AccountCloudforceOneEventDatasetService struct {
	Options []option.RequestOption
}

// NewAccountCloudforceOneEventDatasetService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCloudforceOneEventDatasetService(opts ...option.RequestOption) (r *AccountCloudforceOneEventDatasetService) {
	r = &AccountCloudforceOneEventDatasetService{}
	r.Options = opts
	return
}

// Creates a dataset
func (r *AccountCloudforceOneEventDatasetService) New(ctx context.Context, accountID float64, body AccountCloudforceOneEventDatasetNewParams, opts ...option.RequestOption) (res *AccountCloudforceOneEventDatasetNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/dataset/create", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Reads a dataset
func (r *AccountCloudforceOneEventDatasetService) Get(ctx context.Context, accountID float64, datasetID string, opts ...option.RequestOption) (res *AccountCloudforceOneEventDatasetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/dataset/%s", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing dataset
func (r *AccountCloudforceOneEventDatasetService) Update(ctx context.Context, accountID float64, datasetID string, body AccountCloudforceOneEventDatasetUpdateParams, opts ...option.RequestOption) (res *AccountCloudforceOneEventDatasetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/dataset/%s", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all datasets in an account
func (r *AccountCloudforceOneEventDatasetService) List(ctx context.Context, accountID float64, opts ...option.RequestOption) (res *[]AccountCloudforceOneEventDatasetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/dataset", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountCloudforceOneEventDatasetNewResponse struct {
	IsPublic bool                                            `json:"isPublic,required"`
	Name     string                                          `json:"name,required"`
	Uuid     string                                          `json:"uuid,required"`
	JSON     accountCloudforceOneEventDatasetNewResponseJSON `json:"-"`
}

// accountCloudforceOneEventDatasetNewResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneEventDatasetNewResponse]
type accountCloudforceOneEventDatasetNewResponseJSON struct {
	IsPublic    apijson.Field
	Name        apijson.Field
	Uuid        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventDatasetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventDatasetNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventDatasetGetResponse struct {
	IsPublic bool                                            `json:"isPublic,required"`
	Name     string                                          `json:"name,required"`
	Uuid     string                                          `json:"uuid,required"`
	JSON     accountCloudforceOneEventDatasetGetResponseJSON `json:"-"`
}

// accountCloudforceOneEventDatasetGetResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneEventDatasetGetResponse]
type accountCloudforceOneEventDatasetGetResponseJSON struct {
	IsPublic    apijson.Field
	Name        apijson.Field
	Uuid        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventDatasetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventDatasetGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventDatasetUpdateResponse struct {
	IsPublic bool                                               `json:"isPublic,required"`
	Name     string                                             `json:"name,required"`
	Uuid     string                                             `json:"uuid,required"`
	JSON     accountCloudforceOneEventDatasetUpdateResponseJSON `json:"-"`
}

// accountCloudforceOneEventDatasetUpdateResponseJSON contains the JSON metadata
// for the struct [AccountCloudforceOneEventDatasetUpdateResponse]
type accountCloudforceOneEventDatasetUpdateResponseJSON struct {
	IsPublic    apijson.Field
	Name        apijson.Field
	Uuid        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventDatasetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventDatasetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventDatasetListResponse struct {
	IsPublic bool                                             `json:"isPublic,required"`
	Name     string                                           `json:"name,required"`
	Uuid     string                                           `json:"uuid,required"`
	JSON     accountCloudforceOneEventDatasetListResponseJSON `json:"-"`
}

// accountCloudforceOneEventDatasetListResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneEventDatasetListResponse]
type accountCloudforceOneEventDatasetListResponseJSON struct {
	IsPublic    apijson.Field
	Name        apijson.Field
	Uuid        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventDatasetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventDatasetListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventDatasetNewParams struct {
	// If true, then anyone can search the dataset. If false, then its limited to the
	// account.
	IsPublic param.Field[bool] `json:"isPublic,required"`
	// Used to describe the dataset within the account context
	Name param.Field[string] `json:"name,required"`
}

func (r AccountCloudforceOneEventDatasetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCloudforceOneEventDatasetUpdateParams struct {
	// If true, then anyone can search the dataset. If false, then its limited to the
	// account.
	IsPublic param.Field[bool] `json:"isPublic,required"`
	// Used to describe the dataset within the account context
	Name param.Field[string] `json:"name,required"`
}

func (r AccountCloudforceOneEventDatasetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
