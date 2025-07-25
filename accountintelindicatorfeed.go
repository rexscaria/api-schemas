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

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
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

type UpdateFeedResponse struct {
	Errors   []UpdateFeedResponseError   `json:"errors,required"`
	Messages []UpdateFeedResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success UpdateFeedResponseSuccess `json:"success,required"`
	Result  UpdateFeedResponseResult  `json:"result"`
	JSON    updateFeedResponseJSON    `json:"-"`
}

// updateFeedResponseJSON contains the JSON metadata for the struct
// [UpdateFeedResponse]
type updateFeedResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

type UpdateFeedResponseError struct {
	Code             int64                          `json:"code,required"`
	Message          string                         `json:"message,required"`
	DocumentationURL string                         `json:"documentation_url"`
	Source           UpdateFeedResponseErrorsSource `json:"source"`
	JSON             updateFeedResponseErrorJSON    `json:"-"`
}

// updateFeedResponseErrorJSON contains the JSON metadata for the struct
// [UpdateFeedResponseError]
type updateFeedResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UpdateFeedResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r updateFeedResponseErrorJSON) RawJSON() string {
	return r.raw
}

type UpdateFeedResponseErrorsSource struct {
	Pointer string                             `json:"pointer"`
	JSON    updateFeedResponseErrorsSourceJSON `json:"-"`
}

// updateFeedResponseErrorsSourceJSON contains the JSON metadata for the struct
// [UpdateFeedResponseErrorsSource]
type updateFeedResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UpdateFeedResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r updateFeedResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UpdateFeedResponseMessage struct {
	Code             int64                            `json:"code,required"`
	Message          string                           `json:"message,required"`
	DocumentationURL string                           `json:"documentation_url"`
	Source           UpdateFeedResponseMessagesSource `json:"source"`
	JSON             updateFeedResponseMessageJSON    `json:"-"`
}

// updateFeedResponseMessageJSON contains the JSON metadata for the struct
// [UpdateFeedResponseMessage]
type updateFeedResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UpdateFeedResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r updateFeedResponseMessageJSON) RawJSON() string {
	return r.raw
}

type UpdateFeedResponseMessagesSource struct {
	Pointer string                               `json:"pointer"`
	JSON    updateFeedResponseMessagesSourceJSON `json:"-"`
}

// updateFeedResponseMessagesSourceJSON contains the JSON metadata for the struct
// [UpdateFeedResponseMessagesSource]
type updateFeedResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UpdateFeedResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r updateFeedResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UpdateFeedResponseSuccess bool

const (
	UpdateFeedResponseSuccessTrue UpdateFeedResponseSuccess = true
)

func (r UpdateFeedResponseSuccess) IsKnown() bool {
	switch r {
	case UpdateFeedResponseSuccessTrue:
		return true
	}
	return false
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
	Errors   []AccountIntelIndicatorFeedNewFeedResponseError   `json:"errors,required"`
	Messages []AccountIntelIndicatorFeedNewFeedResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountIntelIndicatorFeedNewFeedResponseSuccess `json:"success,required"`
	Result  FeedItem                                        `json:"result"`
	JSON    accountIntelIndicatorFeedNewFeedResponseJSON    `json:"-"`
}

// accountIntelIndicatorFeedNewFeedResponseJSON contains the JSON metadata for the
// struct [AccountIntelIndicatorFeedNewFeedResponse]
type accountIntelIndicatorFeedNewFeedResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

type AccountIntelIndicatorFeedNewFeedResponseError struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           AccountIntelIndicatorFeedNewFeedResponseErrorsSource `json:"source"`
	JSON             accountIntelIndicatorFeedNewFeedResponseErrorJSON    `json:"-"`
}

// accountIntelIndicatorFeedNewFeedResponseErrorJSON contains the JSON metadata for
// the struct [AccountIntelIndicatorFeedNewFeedResponseError]
type accountIntelIndicatorFeedNewFeedResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedNewFeedResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedNewFeedResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedNewFeedResponseErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    accountIntelIndicatorFeedNewFeedResponseErrorsSourceJSON `json:"-"`
}

// accountIntelIndicatorFeedNewFeedResponseErrorsSourceJSON contains the JSON
// metadata for the struct [AccountIntelIndicatorFeedNewFeedResponseErrorsSource]
type accountIntelIndicatorFeedNewFeedResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedNewFeedResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedNewFeedResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedNewFeedResponseMessage struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           AccountIntelIndicatorFeedNewFeedResponseMessagesSource `json:"source"`
	JSON             accountIntelIndicatorFeedNewFeedResponseMessageJSON    `json:"-"`
}

// accountIntelIndicatorFeedNewFeedResponseMessageJSON contains the JSON metadata
// for the struct [AccountIntelIndicatorFeedNewFeedResponseMessage]
type accountIntelIndicatorFeedNewFeedResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedNewFeedResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedNewFeedResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedNewFeedResponseMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    accountIntelIndicatorFeedNewFeedResponseMessagesSourceJSON `json:"-"`
}

// accountIntelIndicatorFeedNewFeedResponseMessagesSourceJSON contains the JSON
// metadata for the struct [AccountIntelIndicatorFeedNewFeedResponseMessagesSource]
type accountIntelIndicatorFeedNewFeedResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedNewFeedResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedNewFeedResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountIntelIndicatorFeedNewFeedResponseSuccess bool

const (
	AccountIntelIndicatorFeedNewFeedResponseSuccessTrue AccountIntelIndicatorFeedNewFeedResponseSuccess = true
)

func (r AccountIntelIndicatorFeedNewFeedResponseSuccess) IsKnown() bool {
	switch r {
	case AccountIntelIndicatorFeedNewFeedResponseSuccessTrue:
		return true
	}
	return false
}

type AccountIntelIndicatorFeedListFeedsResponse struct {
	Errors   []AccountIntelIndicatorFeedListFeedsResponseError   `json:"errors,required"`
	Messages []AccountIntelIndicatorFeedListFeedsResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountIntelIndicatorFeedListFeedsResponseSuccess `json:"success,required"`
	Result  []FeedItem                                        `json:"result"`
	JSON    accountIntelIndicatorFeedListFeedsResponseJSON    `json:"-"`
}

// accountIntelIndicatorFeedListFeedsResponseJSON contains the JSON metadata for
// the struct [AccountIntelIndicatorFeedListFeedsResponse]
type accountIntelIndicatorFeedListFeedsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

type AccountIntelIndicatorFeedListFeedsResponseError struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           AccountIntelIndicatorFeedListFeedsResponseErrorsSource `json:"source"`
	JSON             accountIntelIndicatorFeedListFeedsResponseErrorJSON    `json:"-"`
}

// accountIntelIndicatorFeedListFeedsResponseErrorJSON contains the JSON metadata
// for the struct [AccountIntelIndicatorFeedListFeedsResponseError]
type accountIntelIndicatorFeedListFeedsResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedListFeedsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedListFeedsResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedListFeedsResponseErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    accountIntelIndicatorFeedListFeedsResponseErrorsSourceJSON `json:"-"`
}

// accountIntelIndicatorFeedListFeedsResponseErrorsSourceJSON contains the JSON
// metadata for the struct [AccountIntelIndicatorFeedListFeedsResponseErrorsSource]
type accountIntelIndicatorFeedListFeedsResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedListFeedsResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedListFeedsResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedListFeedsResponseMessage struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           AccountIntelIndicatorFeedListFeedsResponseMessagesSource `json:"source"`
	JSON             accountIntelIndicatorFeedListFeedsResponseMessageJSON    `json:"-"`
}

// accountIntelIndicatorFeedListFeedsResponseMessageJSON contains the JSON metadata
// for the struct [AccountIntelIndicatorFeedListFeedsResponseMessage]
type accountIntelIndicatorFeedListFeedsResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedListFeedsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedListFeedsResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedListFeedsResponseMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    accountIntelIndicatorFeedListFeedsResponseMessagesSourceJSON `json:"-"`
}

// accountIntelIndicatorFeedListFeedsResponseMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccountIntelIndicatorFeedListFeedsResponseMessagesSource]
type accountIntelIndicatorFeedListFeedsResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedListFeedsResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedListFeedsResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountIntelIndicatorFeedListFeedsResponseSuccess bool

const (
	AccountIntelIndicatorFeedListFeedsResponseSuccessTrue AccountIntelIndicatorFeedListFeedsResponseSuccess = true
)

func (r AccountIntelIndicatorFeedListFeedsResponseSuccess) IsKnown() bool {
	switch r {
	case AccountIntelIndicatorFeedListFeedsResponseSuccessTrue:
		return true
	}
	return false
}

type AccountIntelIndicatorFeedGetMetadataResponse struct {
	Errors   []AccountIntelIndicatorFeedGetMetadataResponseError   `json:"errors,required"`
	Messages []AccountIntelIndicatorFeedGetMetadataResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountIntelIndicatorFeedGetMetadataResponseSuccess `json:"success,required"`
	Result  AccountIntelIndicatorFeedGetMetadataResponseResult  `json:"result"`
	JSON    accountIntelIndicatorFeedGetMetadataResponseJSON    `json:"-"`
}

// accountIntelIndicatorFeedGetMetadataResponseJSON contains the JSON metadata for
// the struct [AccountIntelIndicatorFeedGetMetadataResponse]
type accountIntelIndicatorFeedGetMetadataResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

type AccountIntelIndicatorFeedGetMetadataResponseError struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           AccountIntelIndicatorFeedGetMetadataResponseErrorsSource `json:"source"`
	JSON             accountIntelIndicatorFeedGetMetadataResponseErrorJSON    `json:"-"`
}

// accountIntelIndicatorFeedGetMetadataResponseErrorJSON contains the JSON metadata
// for the struct [AccountIntelIndicatorFeedGetMetadataResponseError]
type accountIntelIndicatorFeedGetMetadataResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedGetMetadataResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedGetMetadataResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedGetMetadataResponseErrorsSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    accountIntelIndicatorFeedGetMetadataResponseErrorsSourceJSON `json:"-"`
}

// accountIntelIndicatorFeedGetMetadataResponseErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccountIntelIndicatorFeedGetMetadataResponseErrorsSource]
type accountIntelIndicatorFeedGetMetadataResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedGetMetadataResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedGetMetadataResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedGetMetadataResponseMessage struct {
	Code             int64                                                      `json:"code,required"`
	Message          string                                                     `json:"message,required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           AccountIntelIndicatorFeedGetMetadataResponseMessagesSource `json:"source"`
	JSON             accountIntelIndicatorFeedGetMetadataResponseMessageJSON    `json:"-"`
}

// accountIntelIndicatorFeedGetMetadataResponseMessageJSON contains the JSON
// metadata for the struct [AccountIntelIndicatorFeedGetMetadataResponseMessage]
type accountIntelIndicatorFeedGetMetadataResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedGetMetadataResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedGetMetadataResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedGetMetadataResponseMessagesSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    accountIntelIndicatorFeedGetMetadataResponseMessagesSourceJSON `json:"-"`
}

// accountIntelIndicatorFeedGetMetadataResponseMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccountIntelIndicatorFeedGetMetadataResponseMessagesSource]
type accountIntelIndicatorFeedGetMetadataResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedGetMetadataResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedGetMetadataResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountIntelIndicatorFeedGetMetadataResponseSuccess bool

const (
	AccountIntelIndicatorFeedGetMetadataResponseSuccessTrue AccountIntelIndicatorFeedGetMetadataResponseSuccess = true
)

func (r AccountIntelIndicatorFeedGetMetadataResponseSuccess) IsKnown() bool {
	switch r {
	case AccountIntelIndicatorFeedGetMetadataResponseSuccessTrue:
		return true
	}
	return false
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
	Errors   []AccountIntelIndicatorFeedUpdateMetadataResponseError   `json:"errors,required"`
	Messages []AccountIntelIndicatorFeedUpdateMetadataResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountIntelIndicatorFeedUpdateMetadataResponseSuccess `json:"success,required"`
	Result  FeedItem                                               `json:"result"`
	JSON    accountIntelIndicatorFeedUpdateMetadataResponseJSON    `json:"-"`
}

// accountIntelIndicatorFeedUpdateMetadataResponseJSON contains the JSON metadata
// for the struct [AccountIntelIndicatorFeedUpdateMetadataResponse]
type accountIntelIndicatorFeedUpdateMetadataResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

type AccountIntelIndicatorFeedUpdateMetadataResponseError struct {
	Code             int64                                                       `json:"code,required"`
	Message          string                                                      `json:"message,required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           AccountIntelIndicatorFeedUpdateMetadataResponseErrorsSource `json:"source"`
	JSON             accountIntelIndicatorFeedUpdateMetadataResponseErrorJSON    `json:"-"`
}

// accountIntelIndicatorFeedUpdateMetadataResponseErrorJSON contains the JSON
// metadata for the struct [AccountIntelIndicatorFeedUpdateMetadataResponseError]
type accountIntelIndicatorFeedUpdateMetadataResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedUpdateMetadataResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedUpdateMetadataResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedUpdateMetadataResponseErrorsSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    accountIntelIndicatorFeedUpdateMetadataResponseErrorsSourceJSON `json:"-"`
}

// accountIntelIndicatorFeedUpdateMetadataResponseErrorsSourceJSON contains the
// JSON metadata for the struct
// [AccountIntelIndicatorFeedUpdateMetadataResponseErrorsSource]
type accountIntelIndicatorFeedUpdateMetadataResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedUpdateMetadataResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedUpdateMetadataResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedUpdateMetadataResponseMessage struct {
	Code             int64                                                         `json:"code,required"`
	Message          string                                                        `json:"message,required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           AccountIntelIndicatorFeedUpdateMetadataResponseMessagesSource `json:"source"`
	JSON             accountIntelIndicatorFeedUpdateMetadataResponseMessageJSON    `json:"-"`
}

// accountIntelIndicatorFeedUpdateMetadataResponseMessageJSON contains the JSON
// metadata for the struct [AccountIntelIndicatorFeedUpdateMetadataResponseMessage]
type accountIntelIndicatorFeedUpdateMetadataResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedUpdateMetadataResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedUpdateMetadataResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountIntelIndicatorFeedUpdateMetadataResponseMessagesSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    accountIntelIndicatorFeedUpdateMetadataResponseMessagesSourceJSON `json:"-"`
}

// accountIntelIndicatorFeedUpdateMetadataResponseMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccountIntelIndicatorFeedUpdateMetadataResponseMessagesSource]
type accountIntelIndicatorFeedUpdateMetadataResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIndicatorFeedUpdateMetadataResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelIndicatorFeedUpdateMetadataResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountIntelIndicatorFeedUpdateMetadataResponseSuccess bool

const (
	AccountIntelIndicatorFeedUpdateMetadataResponseSuccessTrue AccountIntelIndicatorFeedUpdateMetadataResponseSuccess = true
)

func (r AccountIntelIndicatorFeedUpdateMetadataResponseSuccess) IsKnown() bool {
	switch r {
	case AccountIntelIndicatorFeedUpdateMetadataResponseSuccessTrue:
		return true
	}
	return false
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
