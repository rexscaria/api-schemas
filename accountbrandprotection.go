// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountBrandProtectionService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountBrandProtectionService] method instead.
type AccountBrandProtectionService struct {
	Options []option.RequestOption
}

// NewAccountBrandProtectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountBrandProtectionService(opts ...option.RequestOption) (r *AccountBrandProtectionService) {
	r = &AccountBrandProtectionService{}
	r.Options = opts
	return
}

// Return submitted URLs based on ID
func (r *AccountBrandProtectionService) GetURLInfo(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountBrandProtectionGetURLInfoResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/brand-protection/url-info", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Return new URL submissions
func (r *AccountBrandProtectionService) SubmitURL(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountBrandProtectionSubmitURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/brand-protection/submit", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AccountBrandProtectionGetURLInfoResponse struct {
	Result []map[string]interface{}                     `json:"result"`
	JSON   accountBrandProtectionGetURLInfoResponseJSON `json:"-"`
}

// accountBrandProtectionGetURLInfoResponseJSON contains the JSON metadata for the
// struct [AccountBrandProtectionGetURLInfoResponse]
type accountBrandProtectionGetURLInfoResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionGetURLInfoResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrandProtectionGetURLInfoResponseJSON) RawJSON() string {
	return r.raw
}

type AccountBrandProtectionSubmitURLResponse struct {
	SkippedURLs   []map[string]interface{}                    `json:"skipped_urls"`
	SubmittedURLs []map[string]interface{}                    `json:"submitted_urls"`
	JSON          accountBrandProtectionSubmitURLResponseJSON `json:"-"`
}

// accountBrandProtectionSubmitURLResponseJSON contains the JSON metadata for the
// struct [AccountBrandProtectionSubmitURLResponse]
type accountBrandProtectionSubmitURLResponseJSON struct {
	SkippedURLs   apijson.Field
	SubmittedURLs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountBrandProtectionSubmitURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrandProtectionSubmitURLResponseJSON) RawJSON() string {
	return r.raw
}
