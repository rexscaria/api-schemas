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

// ZoneAcmTotalTlService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAcmTotalTlService] method instead.
type ZoneAcmTotalTlService struct {
	Options []option.RequestOption
}

// NewZoneAcmTotalTlService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneAcmTotalTlService(opts ...option.RequestOption) (r *ZoneAcmTotalTlService) {
	r = &ZoneAcmTotalTlService{}
	r.Options = opts
	return
}

// Get Total TLS Settings for a Zone.
func (r *ZoneAcmTotalTlService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingsResponseTotalTls, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/acm/total_tls", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set Total TLS Settings or disable the feature for a Zone.
func (r *ZoneAcmTotalTlService) Update(ctx context.Context, zoneID string, body ZoneAcmTotalTlUpdateParams, opts ...option.RequestOption) (res *SettingsResponseTotalTls, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/acm/total_tls", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The Certificate Authority that Total TLS certificates will be issued through.
type CertificateAuthorityTotalTls string

const (
	CertificateAuthorityTotalTlsGoogle      CertificateAuthorityTotalTls = "google"
	CertificateAuthorityTotalTlsLetsEncrypt CertificateAuthorityTotalTls = "lets_encrypt"
	CertificateAuthorityTotalTlsSslCom      CertificateAuthorityTotalTls = "ssl_com"
)

func (r CertificateAuthorityTotalTls) IsKnown() bool {
	switch r {
	case CertificateAuthorityTotalTlsGoogle, CertificateAuthorityTotalTlsLetsEncrypt, CertificateAuthorityTotalTlsSslCom:
		return true
	}
	return false
}

type SettingsResponseTotalTls struct {
	Result SettingsResponseTotalTlsResult `json:"result"`
	JSON   settingsResponseTotalTlsJSON   `json:"-"`
	APIResponseSingleTlsCertificates
}

// settingsResponseTotalTlsJSON contains the JSON metadata for the struct
// [SettingsResponseTotalTls]
type settingsResponseTotalTlsJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingsResponseTotalTls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingsResponseTotalTlsJSON) RawJSON() string {
	return r.raw
}

type SettingsResponseTotalTlsResult struct {
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority CertificateAuthorityTotalTls `json:"certificate_authority"`
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled bool `json:"enabled"`
	// The validity period in days for the certificates ordered via Total TLS.
	ValidityPeriod SettingsResponseTotalTlsResultValidityPeriod `json:"validity_period"`
	JSON           settingsResponseTotalTlsResultJSON           `json:"-"`
}

// settingsResponseTotalTlsResultJSON contains the JSON metadata for the struct
// [SettingsResponseTotalTlsResult]
type settingsResponseTotalTlsResultJSON struct {
	CertificateAuthority apijson.Field
	Enabled              apijson.Field
	ValidityPeriod       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SettingsResponseTotalTlsResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingsResponseTotalTlsResultJSON) RawJSON() string {
	return r.raw
}

// The validity period in days for the certificates ordered via Total TLS.
type SettingsResponseTotalTlsResultValidityPeriod int64

const (
	SettingsResponseTotalTlsResultValidityPeriod90 SettingsResponseTotalTlsResultValidityPeriod = 90
)

func (r SettingsResponseTotalTlsResultValidityPeriod) IsKnown() bool {
	switch r {
	case SettingsResponseTotalTlsResultValidityPeriod90:
		return true
	}
	return false
}

type ZoneAcmTotalTlUpdateParams struct {
	// If enabled, Total TLS will order a hostname specific TLS certificate for any
	// proxied A, AAAA, or CNAME record in your zone.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The Certificate Authority that Total TLS certificates will be issued through.
	CertificateAuthority param.Field[CertificateAuthorityTotalTls] `json:"certificate_authority"`
}

func (r ZoneAcmTotalTlUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
