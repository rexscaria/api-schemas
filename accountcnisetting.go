// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountCniSettingService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCniSettingService] method instead.
type AccountCniSettingService struct {
	Options []option.RequestOption
}

// NewAccountCniSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountCniSettingService(opts ...option.RequestOption) (r *AccountCniSettingService) {
	r = &AccountCniSettingService{}
	r.Options = opts
	return
}

// Get the current settings for the active account
func (r *AccountCniSettingService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *NscSettings, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the current settings for the active account
func (r *AccountCniSettingService) Update(ctx context.Context, accountID string, body AccountCniSettingUpdateParams, opts ...option.RequestOption) (res *NscSettings, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type NscSettings struct {
	DefaultAsn int64           `json:"default_asn,required"`
	JSON       nscSettingsJSON `json:"-"`
}

// nscSettingsJSON contains the JSON metadata for the struct [NscSettings]
type nscSettingsJSON struct {
	DefaultAsn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NscSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nscSettingsJSON) RawJSON() string {
	return r.raw
}

type AccountCniSettingUpdateParams struct {
	DefaultAsn param.Field[int64] `json:"default_asn"`
}

func (r AccountCniSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
