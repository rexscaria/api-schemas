// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apiform"
	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountCloudforceOneRequestAssetService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCloudforceOneRequestAssetService] method instead.
type AccountCloudforceOneRequestAssetService struct {
	Options []option.RequestOption
}

// NewAccountCloudforceOneRequestAssetService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCloudforceOneRequestAssetService(opts ...option.RequestOption) (r *AccountCloudforceOneRequestAssetService) {
	r = &AccountCloudforceOneRequestAssetService{}
	r.Options = opts
	return
}

// Create a New Request Asset
func (r *AccountCloudforceOneRequestAssetService) New(ctx context.Context, accountIdentifier string, requestIdentifier string, body AccountCloudforceOneRequestAssetNewParams, opts ...option.RequestOption) (res *AccountCloudforceOneRequestAssetNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/asset/new", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a Request Asset
func (r *AccountCloudforceOneRequestAssetService) Get(ctx context.Context, accountIdentifier string, requestIdentifier string, assetIdentifer string, opts ...option.RequestOption) (res *AccountCloudforceOneRequestAssetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	if assetIdentifer == "" {
		err = errors.New("missing required asset_identifer parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/asset/%s", accountIdentifier, requestIdentifier, assetIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Request Asset
func (r *AccountCloudforceOneRequestAssetService) Update(ctx context.Context, accountIdentifier string, requestIdentifier string, assetIdentifer string, body AccountCloudforceOneRequestAssetUpdateParams, opts ...option.RequestOption) (res *AccountCloudforceOneRequestAssetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	if assetIdentifer == "" {
		err = errors.New("missing required asset_identifer parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/asset/%s", accountIdentifier, requestIdentifier, assetIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List Request Assets
func (r *AccountCloudforceOneRequestAssetService) List(ctx context.Context, accountIdentifier string, requestIdentifier string, body AccountCloudforceOneRequestAssetListParams, opts ...option.RequestOption) (res *AccountCloudforceOneRequestAssetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/asset", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a Request Asset
func (r *AccountCloudforceOneRequestAssetService) Delete(ctx context.Context, accountIdentifier string, requestIdentifier string, assetIdentifer string, opts ...option.RequestOption) (res *APIResponseRequests, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	if assetIdentifer == "" {
		err = errors.New("missing required asset_identifer parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/asset/%s", accountIdentifier, requestIdentifier, assetIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AssetEditParam struct {
	// Asset file to upload
	Source param.Field[string] `json:"source"`
}

func (r AssetEditParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AssetItem struct {
	// Asset ID
	ID int64 `json:"id,required"`
	// Asset name
	Name string `json:"name,required"`
	// Asset creation time
	Created AssetItemCreated `json:"created"`
	// Asset description
	Description string `json:"description"`
	// Asset file type
	FileType string        `json:"file_type"`
	JSON     assetItemJSON `json:"-"`
}

// assetItemJSON contains the JSON metadata for the struct [AssetItem]
type assetItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Created     apijson.Field
	Description apijson.Field
	FileType    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetItemJSON) RawJSON() string {
	return r.raw
}

// Asset creation time
type AssetItemCreated struct {
	JSON assetItemCreatedJSON `json:"-"`
}

// assetItemCreatedJSON contains the JSON metadata for the struct
// [AssetItemCreated]
type assetItemCreatedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetItemCreated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetItemCreatedJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestAssetNewResponse struct {
	Result AssetItem                                       `json:"result"`
	JSON   accountCloudforceOneRequestAssetNewResponseJSON `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestAssetNewResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneRequestAssetNewResponse]
type accountCloudforceOneRequestAssetNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestAssetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestAssetNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestAssetGetResponse struct {
	Result []AssetItem                                     `json:"result"`
	JSON   accountCloudforceOneRequestAssetGetResponseJSON `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestAssetGetResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneRequestAssetGetResponse]
type accountCloudforceOneRequestAssetGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestAssetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestAssetGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestAssetUpdateResponse struct {
	Result AssetItem                                          `json:"result"`
	JSON   accountCloudforceOneRequestAssetUpdateResponseJSON `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestAssetUpdateResponseJSON contains the JSON metadata
// for the struct [AccountCloudforceOneRequestAssetUpdateResponse]
type accountCloudforceOneRequestAssetUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestAssetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestAssetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestAssetListResponse struct {
	Result []AssetItem                                      `json:"result"`
	JSON   accountCloudforceOneRequestAssetListResponseJSON `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestAssetListResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneRequestAssetListResponse]
type accountCloudforceOneRequestAssetListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestAssetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestAssetListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestAssetNewParams struct {
	AssetEdit AssetEditParam `json:"asset_edit,required"`
}

func (r AccountCloudforceOneRequestAssetNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type AccountCloudforceOneRequestAssetUpdateParams struct {
	AssetEdit AssetEditParam `json:"asset_edit,required"`
}

func (r AccountCloudforceOneRequestAssetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AssetEdit)
}

type AccountCloudforceOneRequestAssetListParams struct {
	// Page number of results
	Page param.Field[int64] `json:"page,required"`
	// Number of results per page
	PerPage param.Field[int64] `json:"per_page,required"`
}

func (r AccountCloudforceOneRequestAssetListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
