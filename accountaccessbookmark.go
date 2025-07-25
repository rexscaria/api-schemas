// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountAccessBookmarkService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessBookmarkService] method instead.
type AccountAccessBookmarkService struct {
	Options []option.RequestOption
}

// NewAccountAccessBookmarkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessBookmarkService(opts ...option.RequestOption) (r *AccountAccessBookmarkService) {
	r = &AccountAccessBookmarkService{}
	r.Options = opts
	return
}

// Create a new Bookmark application.
//
// Deprecated: deprecated
func (r *AccountAccessBookmarkService) New(ctx context.Context, accountID string, bookmarkID string, body AccountAccessBookmarkNewParams, opts ...option.RequestOption) (res *SingleResponseBookmark, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bookmarkID == "" {
		err = errors.New("missing required bookmark_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/bookmarks/%s", accountID, bookmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single Bookmark application.
//
// Deprecated: deprecated
func (r *AccountAccessBookmarkService) Get(ctx context.Context, accountID string, bookmarkID string, opts ...option.RequestOption) (res *SingleResponseBookmark, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bookmarkID == "" {
		err = errors.New("missing required bookmark_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/bookmarks/%s", accountID, bookmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured Bookmark application.
//
// Deprecated: deprecated
func (r *AccountAccessBookmarkService) Update(ctx context.Context, accountID string, bookmarkID string, body AccountAccessBookmarkUpdateParams, opts ...option.RequestOption) (res *SingleResponseBookmark, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bookmarkID == "" {
		err = errors.New("missing required bookmark_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/bookmarks/%s", accountID, bookmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists Bookmark applications.
//
// Deprecated: deprecated
func (r *AccountAccessBookmarkService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAccessBookmarkListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/bookmarks", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a Bookmark application.
//
// Deprecated: deprecated
func (r *AccountAccessBookmarkService) Delete(ctx context.Context, accountID string, bookmarkID string, opts ...option.RequestOption) (res *IDResponseApps, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bookmarkID == "" {
		err = errors.New("missing required bookmark_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/bookmarks/%s", accountID, bookmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Bookmarks struct {
	// The unique identifier for the Bookmark application.
	ID string `json:"id"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool      `json:"app_launcher_visible"`
	CreatedAt          time.Time `json:"created_at" format:"date-time"`
	// The domain of the Bookmark application.
	Domain string `json:"domain"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the Bookmark application.
	Name      string        `json:"name"`
	UpdatedAt time.Time     `json:"updated_at" format:"date-time"`
	JSON      bookmarksJSON `json:"-"`
}

// bookmarksJSON contains the JSON metadata for the struct [Bookmarks]
type bookmarksJSON struct {
	ID                 apijson.Field
	AppLauncherVisible apijson.Field
	CreatedAt          apijson.Field
	Domain             apijson.Field
	LogoURL            apijson.Field
	Name               apijson.Field
	UpdatedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Bookmarks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bookmarksJSON) RawJSON() string {
	return r.raw
}

type SingleResponseBookmark struct {
	Errors   []MessagesAccessItem `json:"errors,required"`
	Messages []MessagesAccessItem `json:"messages,required"`
	// Whether the API call was successful.
	Success SingleResponseBookmarkSuccess `json:"success,required"`
	Result  Bookmarks                     `json:"result"`
	JSON    singleResponseBookmarkJSON    `json:"-"`
}

// singleResponseBookmarkJSON contains the JSON metadata for the struct
// [SingleResponseBookmark]
type singleResponseBookmarkJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseBookmark) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseBookmarkJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SingleResponseBookmarkSuccess bool

const (
	SingleResponseBookmarkSuccessTrue SingleResponseBookmarkSuccess = true
)

func (r SingleResponseBookmarkSuccess) IsKnown() bool {
	switch r {
	case SingleResponseBookmarkSuccessTrue:
		return true
	}
	return false
}

type AccountAccessBookmarkListResponse struct {
	Errors   []MessagesAccessItem `json:"errors,required"`
	Messages []MessagesAccessItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    AccountAccessBookmarkListResponseSuccess    `json:"success,required"`
	Result     []Bookmarks                                 `json:"result"`
	ResultInfo AccountAccessBookmarkListResponseResultInfo `json:"result_info"`
	JSON       accountAccessBookmarkListResponseJSON       `json:"-"`
}

// accountAccessBookmarkListResponseJSON contains the JSON metadata for the struct
// [AccountAccessBookmarkListResponse]
type accountAccessBookmarkListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessBookmarkListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountAccessBookmarkListResponseSuccess bool

const (
	AccountAccessBookmarkListResponseSuccessTrue AccountAccessBookmarkListResponseSuccess = true
)

func (r AccountAccessBookmarkListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountAccessBookmarkListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountAccessBookmarkListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                         `json:"total_count"`
	JSON       accountAccessBookmarkListResponseResultInfoJSON `json:"-"`
}

// accountAccessBookmarkListResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountAccessBookmarkListResponseResultInfo]
type accountAccessBookmarkListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessBookmarkListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountAccessBookmarkNewParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountAccessBookmarkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountAccessBookmarkUpdateParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountAccessBookmarkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
