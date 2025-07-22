// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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
func (r *AccountAccessBookmarkService) Delete(ctx context.Context, accountID string, bookmarkID string, body AccountAccessBookmarkDeleteParams, opts ...option.RequestOption) (res *IDResponseApps, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
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
	Result Bookmarks                  `json:"result"`
	JSON   singleResponseBookmarkJSON `json:"-"`
	APIResponseSingleAccess
}

// singleResponseBookmarkJSON contains the JSON metadata for the struct
// [SingleResponseBookmark]
type singleResponseBookmarkJSON struct {
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

type AccountAccessBookmarkListResponse struct {
	Result []Bookmarks                           `json:"result"`
	JSON   accountAccessBookmarkListResponseJSON `json:"-"`
	APIResponseCollectionAccess
}

// accountAccessBookmarkListResponseJSON contains the JSON metadata for the struct
// [AccountAccessBookmarkListResponse]
type accountAccessBookmarkListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessBookmarkListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessBookmarkListResponseJSON) RawJSON() string {
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

type AccountAccessBookmarkDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountAccessBookmarkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
