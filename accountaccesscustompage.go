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

// AccountAccessCustomPageService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessCustomPageService] method instead.
type AccountAccessCustomPageService struct {
	Options []option.RequestOption
}

// NewAccountAccessCustomPageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessCustomPageService(opts ...option.RequestOption) (r *AccountAccessCustomPageService) {
	r = &AccountAccessCustomPageService{}
	r.Options = opts
	return
}

// Create a custom page
func (r *AccountAccessCustomPageService) New(ctx context.Context, accountID string, body AccountAccessCustomPageNewParams, opts ...option.RequestOption) (res *SingleResponseWithoutHTML, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/custom_pages", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a custom page and also returns its HTML.
func (r *AccountAccessCustomPageService) Get(ctx context.Context, accountID string, customPageID string, opts ...option.RequestOption) (res *AccountAccessCustomPageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if customPageID == "" {
		err = errors.New("missing required custom_page_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", accountID, customPageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a custom page
func (r *AccountAccessCustomPageService) Update(ctx context.Context, accountID string, customPageID string, body AccountAccessCustomPageUpdateParams, opts ...option.RequestOption) (res *SingleResponseWithoutHTML, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if customPageID == "" {
		err = errors.New("missing required custom_page_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", accountID, customPageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List custom pages
func (r *AccountAccessCustomPageService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAccessCustomPageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/custom_pages", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a custom page
func (r *AccountAccessCustomPageService) Delete(ctx context.Context, accountID string, customPageID string, opts ...option.RequestOption) (res *IDResponseCertificates, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if customPageID == "" {
		err = errors.New("missing required custom_page_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", accountID, customPageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CustomPage struct {
	// Custom page HTML.
	CustomHTML string `json:"custom_html,required"`
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type SchemasCustomPageType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string         `json:"uid"`
	UpdatedAt time.Time      `json:"updated_at" format:"date-time"`
	JSON      customPageJSON `json:"-"`
}

// customPageJSON contains the JSON metadata for the struct [CustomPage]
type customPageJSON struct {
	CustomHTML  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	Uid         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageJSON) RawJSON() string {
	return r.raw
}

type CustomPageParam struct {
	// Custom page HTML.
	CustomHTML param.Field[string] `json:"custom_html,required"`
	// Custom page name.
	Name param.Field[string] `json:"name,required"`
	// Custom page type.
	Type param.Field[SchemasCustomPageType] `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  param.Field[int64]     `json:"app_count"`
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// UUID
	Uid       param.Field[string]    `json:"uid"`
	UpdatedAt param.Field[time.Time] `json:"updated_at" format:"date-time"`
}

func (r CustomPageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomPageWithoutHTML struct {
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type SchemasCustomPageType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                    `json:"uid"`
	UpdatedAt time.Time                 `json:"updated_at" format:"date-time"`
	JSON      customPageWithoutHTMLJSON `json:"-"`
}

// customPageWithoutHTMLJSON contains the JSON metadata for the struct
// [CustomPageWithoutHTML]
type customPageWithoutHTMLJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	Uid         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPageWithoutHTML) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageWithoutHTMLJSON) RawJSON() string {
	return r.raw
}

// Custom page type.
type SchemasCustomPageType string

const (
	SchemasCustomPageTypeIdentityDenied SchemasCustomPageType = "identity_denied"
	SchemasCustomPageTypeForbidden      SchemasCustomPageType = "forbidden"
)

func (r SchemasCustomPageType) IsKnown() bool {
	switch r {
	case SchemasCustomPageTypeIdentityDenied, SchemasCustomPageTypeForbidden:
		return true
	}
	return false
}

type SingleResponseWithoutHTML struct {
	Result CustomPageWithoutHTML         `json:"result"`
	JSON   singleResponseWithoutHTMLJSON `json:"-"`
	APIResponseSingleAccess
}

// singleResponseWithoutHTMLJSON contains the JSON metadata for the struct
// [SingleResponseWithoutHTML]
type singleResponseWithoutHTMLJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseWithoutHTML) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseWithoutHTMLJSON) RawJSON() string {
	return r.raw
}

type AccountAccessCustomPageGetResponse struct {
	Result CustomPage                             `json:"result"`
	JSON   accountAccessCustomPageGetResponseJSON `json:"-"`
	APIResponseSingleAccess
}

// accountAccessCustomPageGetResponseJSON contains the JSON metadata for the struct
// [AccountAccessCustomPageGetResponse]
type accountAccessCustomPageGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessCustomPageGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessCustomPageListResponse struct {
	Result []CustomPageWithoutHTML                 `json:"result"`
	JSON   accountAccessCustomPageListResponseJSON `json:"-"`
	APIResponseCollectionAccess
}

// accountAccessCustomPageListResponseJSON contains the JSON metadata for the
// struct [AccountAccessCustomPageListResponse]
type accountAccessCustomPageListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCustomPageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessCustomPageListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessCustomPageNewParams struct {
	CustomPage CustomPageParam `json:"custom_page,required"`
}

func (r AccountAccessCustomPageNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CustomPage)
}

type AccountAccessCustomPageUpdateParams struct {
	CustomPage CustomPageParam `json:"custom_page,required"`
}

func (r AccountAccessCustomPageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CustomPage)
}
