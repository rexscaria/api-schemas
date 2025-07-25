// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneCertificateAuthorityHostnameAssociationService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCertificateAuthorityHostnameAssociationService] method instead.
type ZoneCertificateAuthorityHostnameAssociationService struct {
	Options []option.RequestOption
}

// NewZoneCertificateAuthorityHostnameAssociationService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewZoneCertificateAuthorityHostnameAssociationService(opts ...option.RequestOption) (r *ZoneCertificateAuthorityHostnameAssociationService) {
	r = &ZoneCertificateAuthorityHostnameAssociationService{}
	r.Options = opts
	return
}

// List Hostname Associations
func (r *ZoneCertificateAuthorityHostnameAssociationService) List(ctx context.Context, zoneID string, query ZoneCertificateAuthorityHostnameAssociationListParams, opts ...option.RequestOption) (res *HostnameAssociationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/certificate_authorities/hostname_associations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Replace Hostname Associations
func (r *ZoneCertificateAuthorityHostnameAssociationService) Replace(ctx context.Context, zoneID string, body ZoneCertificateAuthorityHostnameAssociationReplaceParams, opts ...option.RequestOption) (res *HostnameAssociationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/certificate_authorities/hostname_associations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type HostnameAssociationsResponse struct {
	Errors   []MessagesTlsCertificatesItem `json:"errors,required"`
	Messages []MessagesTlsCertificatesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success HostnameAssociationsResponseSuccess `json:"success,required"`
	Result  HostnameAssociationsResponseResult  `json:"result"`
	JSON    hostnameAssociationsResponseJSON    `json:"-"`
}

// hostnameAssociationsResponseJSON contains the JSON metadata for the struct
// [HostnameAssociationsResponse]
type hostnameAssociationsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAssociationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameAssociationsResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type HostnameAssociationsResponseSuccess bool

const (
	HostnameAssociationsResponseSuccessTrue HostnameAssociationsResponseSuccess = true
)

func (r HostnameAssociationsResponseSuccess) IsKnown() bool {
	switch r {
	case HostnameAssociationsResponseSuccessTrue:
		return true
	}
	return false
}

type HostnameAssociationsResponseResult struct {
	Hostnames []string                               `json:"hostnames"`
	JSON      hostnameAssociationsResponseResultJSON `json:"-"`
}

// hostnameAssociationsResponseResultJSON contains the JSON metadata for the struct
// [HostnameAssociationsResponseResult]
type hostnameAssociationsResponseResultJSON struct {
	Hostnames   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAssociationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameAssociationsResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneCertificateAuthorityHostnameAssociationListParams struct {
	// The UUID to match against for a certificate that was uploaded to the mTLS
	// Certificate Management endpoint. If no mtls_certificate_id is given, the results
	// will be the hostnames associated to your active Cloudflare Managed CA.
	MtlsCertificateID param.Field[string] `query:"mtls_certificate_id"`
}

// URLQuery serializes [ZoneCertificateAuthorityHostnameAssociationListParams]'s
// query parameters as `url.Values`.
func (r ZoneCertificateAuthorityHostnameAssociationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneCertificateAuthorityHostnameAssociationReplaceParams struct {
	Hostnames param.Field[[]string] `json:"hostnames"`
	// The UUID for a certificate that was uploaded to the mTLS Certificate Management
	// endpoint. If no mtls_certificate_id is given, the hostnames will be associated
	// to your active Cloudflare Managed CA.
	MtlsCertificateID param.Field[string] `json:"mtls_certificate_id"`
}

func (r ZoneCertificateAuthorityHostnameAssociationReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
