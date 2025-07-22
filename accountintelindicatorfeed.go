// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apiform"
	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountIntelIndicatorFeedService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountIntelIndicatorFeedService] method instead.
type AccountIntelIndicatorFeedService struct {
	Options     []option.RequestOption
	Permissions *AccountIntelIndicatorFeedPermissionService
}

// NewAccountIntelIndicatorFeedService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountIntelIndicatorFeedService(opts ...option.RequestOption) (r *AccountIntelIndicatorFeedService) {
	r = &AccountIntelIndicatorFeedService{}
	r.Options = opts
	r.Permissions = NewAccountIntelIndicatorFeedPermissionService(opts...)
	return
}

// Create new indicator feed
func (r *AccountIntelIndicatorFeedService) NewFeed(ctx context.Context, accountID string, body AccountIntelIndicatorFeedNewFeedParams, opts ...option.RequestOption) (res *AccountIntelIndicatorFeedNewFeedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Download indicator feed data
func (r *AccountIntelIndicatorFeedService) DownloadData(ctx context.Context, accountID string, feedID int64, opts ...option.RequestOption) (res *UpdateFeedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator_feeds/%v/download", accountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get indicator feed data
func (r *AccountIntelIndicatorFeedService) GetData(ctx context.Context, accountID string, feedID int64, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/csv")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v/data", accountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get indicator feeds owned by this account
func (r *AccountIntelIndicatorFeedService) ListFeeds(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountIntelIndicatorFeedListFeedsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get indicator feed metadata
func (r *AccountIntelIndicatorFeedService) GetMetadata(ctx context.Context, accountID string, feedID int64, opts ...option.RequestOption) (res *AccountIntelIndicatorFeedGetMetadataResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v", accountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update indicator feed data
func (r *AccountIntelIndicatorFeedService) UpdateData(ctx context.Context, accountID string, feedID int64, body AccountIntelIndicatorFeedUpdateDataParams, opts ...option.RequestOption) (res *UpdateFeedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v/snapshot", accountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Update indicator feed metadata
func (r *AccountIntelIndicatorFeedService) UpdateMetadata(ctx context.Context, accountID string, feedID int64, body AccountIntelIndicatorFeedUpdateMetadataParams, opts ...option.RequestOption) (res *AccountIntelIndicatorFeedUpdateMetadataResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v", accountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CommonResponseIndicatorFeeds struct {
	Errors   []CustomIndicatorFeedMessage `json:"errors,required"`
	Messages []CustomIndicatorFeedMessage `json:"messages,required"`
	// Whether the API call was successful
	Success CommonResponseIndicatorFeedsSuccess `json:"success,required"`
	JSON    commonResponseIndicatorFeedsJSON    `json:"-"`
}

// commonResponseIndicatorFeedsJSON contains the JSON metadata for the struct
// [CommonResponseIndicatorFeeds]
type commonResponseIndicatorFeedsJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommonResponseIndicatorFeeds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commonResponseIndicatorFeedsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CommonResponseIndicatorFeedsSuccess bool

const (
	CommonResponseIndicatorFeedsSuccessTrue CommonResponseIndicatorFeedsSuccess = true
)

func (r CommonResponseIndicatorFeedsSuccess) IsKnown() bool {
	switch r {
	case CommonResponseIndicatorFeedsSuccessTrue:
		return true
	}
	return false
}

type CustomIndicatorFeedMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    customIndicatorFeedMessageJSON `json:"-"`
}

// customIndicatorFeedMessageJSON contains the JSON metadata for the struct
// [CustomIndicatorFeedMessage]
type customIndicatorFeedMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomIndicatorFeedMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customIndicatorFeedMessageJSON) RawJSON() string {
	return r.raw
}

type FeedItem struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The date and time when the data entry was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The description of the example test
	Description string `json:"description"`
	// Whether the indicator feed can be attributed to a provider
	IsAttributable bool `json:"is_attributable"`
	// Whether the indicator feed can be downloaded
	IsDownloadable bool `json:"is_downloadable"`
	// Whether the indicator feed is exposed to customers
	IsPublic bool `json:"is_public"`
	// The date and time when the data entry was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the indicator feed
	Name string       `json:"name"`
	JSON feedItemJSON `json:"-"`
}

// feedItemJSON contains the JSON metadata for the struct [FeedItem]
type feedItemJSON struct {
	ID             apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	IsAttributable apijson.Field
	IsDownloadable apijson.Field
	IsPublic       apijson.Field
	ModifiedOn     apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FeedItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r feedItemJSON) RawJSON() string {
	return r.raw
}

type SingleResponseFeed struct {
	Errors   []CustomIndicatorFeedMessage `json:"errors,required"`
	Messages []CustomIndicatorFeedMessage `json:"messages,required"`
	// Whether the API call was successful
	Success SingleResponseFeedSuccess `json:"success,required"`
	JSON    singleResponseFeedJSON    `json:"-"`
}

// singleResponseFeedJSON contains the JSON metadata for the struct
// [SingleResponseFeed]
type singleResponseFeedJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseFeed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseFeedJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SingleResponseFeedSuccess bool

const (
	SingleResponseFeedSuccessTrue SingleResponseFeedSuccess = true
)

func (r SingleResponseFeedSuccess) IsKnown() bool {
	switch r {
	case SingleResponseFeedSuccessTrue:
		return true
	}
	return false
}

type UpdateFeedResponse struct {
	Result UpdateFeedResponseResult `json:"result"`
	JSON   updateFeedResponseJSON   `json:"-"`
	SingleResponseFeed
}

// updateFeedResponseJSON contains the JSON metadata for the struct
// [UpdateFeedResponse]
type updateFeedResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UpdateFeedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r updateFeedResponseJSON) RawJSON() string {
	return r.raw
}

type UpdateFeedResponseResult struct {
	// Feed id
	FileID int64 `json:"file_id"`
	// Name of the file unified in our system
	Filename string `json:"filename"`
	// Current status of upload, should be unified
	Status string                       `json:"status"`
	JSON   updateFeedResponseResultJSON `json:"-"`
}

// updateFeedResponseResultJSON contains the JSON metadata for the struct
// [UpdateFeedResponseResult]
type updateFeedResponseResultJSON struct {
	FileID      apijson.Field
	Filename    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UpdateFeedResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r updateFeedResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedNewFeedResponse struct {
	Result FeedItem                                     `json:"result"`
	JSON   accountIntelIndicatorFeedNewFeedResponseJSON `json:"-"`
	SingleResponseFeed
}

// accountIntelIndicatorFeedNewFeedResponseJSON contains the JSON metadata for the
// struct [AccountIntelIndicatorFeedNewFeedResponse]
type accountIntelIndicatorFeedNewFeedResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedNewFeedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedNewFeedResponseJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedListFeedsResponse struct {
	Result []FeedItem                                     `json:"result"`
	JSON   accountIntelIndicatorFeedListFeedsResponseJSON `json:"-"`
	CommonResponseIndicatorFeeds
}

// accountIntelIndicatorFeedListFeedsResponseJSON contains the JSON metadata for
// the struct [AccountIntelIndicatorFeedListFeedsResponse]
type accountIntelIndicatorFeedListFeedsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedListFeedsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedListFeedsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedGetMetadataResponse struct {
	Result AccountIntelIndicatorFeedGetMetadataResponseResult `json:"result"`
	JSON   accountIntelIndicatorFeedGetMetadataResponseJSON   `json:"-"`
	SingleResponseFeed
}

// accountIntelIndicatorFeedGetMetadataResponseJSON contains the JSON metadata for
// the struct [AccountIntelIndicatorFeedGetMetadataResponse]
type accountIntelIndicatorFeedGetMetadataResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedGetMetadataResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedGetMetadataResponseJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedGetMetadataResponseResult struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The date and time when the data entry was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The description of the example test
	Description string `json:"description"`
	// Whether the indicator feed can be attributed to a provider
	IsAttributable bool `json:"is_attributable"`
	// Whether the indicator feed can be downloaded
	IsDownloadable bool `json:"is_downloadable"`
	// Whether the indicator feed is exposed to customers
	IsPublic bool `json:"is_public"`
	// Status of the latest snapshot uploaded
	LatestUploadStatus AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatus `json:"latest_upload_status"`
	// The date and time when the data entry was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the indicator feed
	Name string `json:"name"`
	// The unique identifier for the provider
	ProviderID string `json:"provider_id"`
	// The provider of the indicator feed
	ProviderName string                                                 `json:"provider_name"`
	JSON         accountIntelIndicatorFeedGetMetadataResponseResultJSON `json:"-"`
}

// accountIntelIndicatorFeedGetMetadataResponseResultJSON contains the JSON
// metadata for the struct [AccountIntelIndicatorFeedGetMetadataResponseResult]
type accountIntelIndicatorFeedGetMetadataResponseResultJSON struct {
	ID                 apijson.Field
	CreatedOn          apijson.Field
	Description        apijson.Field
	IsAttributable     apijson.Field
	IsDownloadable     apijson.Field
	IsPublic           apijson.Field
	LatestUploadStatus apijson.Field
	ModifiedOn         apijson.Field
	Name               apijson.Field
	ProviderID         apijson.Field
	ProviderName       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedGetMetadataResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedGetMetadataResponseResultJSON) RawJSON() string {
	return r.raw
}

// Status of the latest snapshot uploaded
type AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatus string

const (
	AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatusMirroring    AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatus = "Mirroring"
	AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatusUnifying     AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatus = "Unifying"
	AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatusLoading      AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatus = "Loading"
	AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatusProvisioning AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatus = "Provisioning"
	AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatusComplete     AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatus = "Complete"
	AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatusError        AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatus = "Error"
)

func (r AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatus) IsKnown() bool {
	switch r {
	case AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatusMirroring, AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatusUnifying, AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatusLoading, AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatusProvisioning, AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatusComplete, AccountIntelIndicatorFeedGetMetadataResponseResultLatestUploadStatusError:
		return true
	}
	return false
}

type AccountIntelIndicatorFeedUpdateMetadataResponse struct {
	Result FeedItem                                            `json:"result"`
	JSON   accountIntelIndicatorFeedUpdateMetadataResponseJSON `json:"-"`
	SingleResponseFeed
}

// accountIntelIndicatorFeedUpdateMetadataResponseJSON contains the JSON metadata
// for the struct [AccountIntelIndicatorFeedUpdateMetadataResponse]
type accountIntelIndicatorFeedUpdateMetadataResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedUpdateMetadataResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedUpdateMetadataResponseJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedNewFeedParams struct {
	// The description of the example test
	Description param.Field[string] `json:"description"`
	// The name of the indicator feed
	Name param.Field[string] `json:"name"`
}

func (r AccountIntelIndicatorFeedNewFeedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountIntelIndicatorFeedUpdateDataParams struct {
	// The file to upload
	Source param.Field[string] `json:"source"`
}

func (r AccountIntelIndicatorFeedUpdateDataParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type AccountIntelIndicatorFeedUpdateMetadataParams struct {
	// The new description of the feed
	Description param.Field[string] `json:"description"`
	// The new is_attributable value of the feed
	IsAttributable param.Field[bool] `json:"is_attributable"`
	// The new is_downloadable value of the feed
	IsDownloadable param.Field[bool] `json:"is_downloadable"`
	// The new is_public value of the feed
	IsPublic param.Field[bool] `json:"is_public"`
	// The new name of the feed
	Name param.Field[string] `json:"name"`
}

func (r AccountIntelIndicatorFeedUpdateMetadataParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
