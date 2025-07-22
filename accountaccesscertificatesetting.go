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

// AccountAccessCertificateSettingService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessCertificateSettingService] method instead.
type AccountAccessCertificateSettingService struct {
	Options []option.RequestOption
}

// NewAccountAccessCertificateSettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessCertificateSettingService(opts ...option.RequestOption) (r *AccountAccessCertificateSettingService) {
	r = &AccountAccessCertificateSettingService{}
	r.Options = opts
	return
}

// Updates an mTLS certificate's hostname settings.
func (r *AccountAccessCertificateSettingService) Update(ctx context.Context, accountID string, body AccountAccessCertificateSettingUpdateParams, opts ...option.RequestOption) (res *ResponseCollectionHostnames, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/certificates/settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all mTLS hostname settings for this account.
func (r *AccountAccessCertificateSettingService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *ResponseCollectionHostnames, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/certificates/settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessSettings struct {
	// Request client certificates for this hostname in China. Can only be set to true
	// if this zone is china network enabled.
	ChinaNetwork bool `json:"china_network,required"`
	// Client Certificate Forwarding is a feature that takes the client cert provided
	// by the eyeball to the edge, and forwards it to the origin as a HTTP header to
	// allow logging on the origin.
	ClientCertificateForwarding bool `json:"client_certificate_forwarding,required"`
	// The hostname that these settings apply to.
	Hostname string             `json:"hostname,required"`
	JSON     accessSettingsJSON `json:"-"`
}

// accessSettingsJSON contains the JSON metadata for the struct [AccessSettings]
type accessSettingsJSON struct {
	ChinaNetwork                apijson.Field
	ClientCertificateForwarding apijson.Field
	Hostname                    apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *AccessSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessSettingsJSON) RawJSON() string {
	return r.raw
}

type AccessSettingsParam struct {
	// Request client certificates for this hostname in China. Can only be set to true
	// if this zone is china network enabled.
	ChinaNetwork param.Field[bool] `json:"china_network,required"`
	// Client Certificate Forwarding is a feature that takes the client cert provided
	// by the eyeball to the edge, and forwards it to the origin as a HTTP header to
	// allow logging on the origin.
	ClientCertificateForwarding param.Field[bool] `json:"client_certificate_forwarding,required"`
	// The hostname that these settings apply to.
	Hostname param.Field[string] `json:"hostname,required"`
}

func (r AccessSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResponseCollectionHostnames struct {
	Result []AccessSettings                `json:"result"`
	JSON   responseCollectionHostnamesJSON `json:"-"`
	APIResponseCollectionAccess
}

// responseCollectionHostnamesJSON contains the JSON metadata for the struct
// [ResponseCollectionHostnames]
type responseCollectionHostnamesJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionHostnames) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionHostnamesJSON) RawJSON() string {
	return r.raw
}

type AccountAccessCertificateSettingUpdateParams struct {
	Settings param.Field[[]AccessSettingsParam] `json:"settings,required"`
}

func (r AccountAccessCertificateSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
