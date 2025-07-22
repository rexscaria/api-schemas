// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneCustomCertificateService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCustomCertificateService] method instead.
type ZoneCustomCertificateService struct {
	Options []option.RequestOption
}

// NewZoneCustomCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneCustomCertificateService(opts ...option.RequestOption) (r *ZoneCustomCertificateService) {
	r = &ZoneCustomCertificateService{}
	r.Options = opts
	return
}

// Upload a new SSL certificate for a zone.
func (r *ZoneCustomCertificateService) New(ctx context.Context, zoneID string, body ZoneCustomCertificateNewParams, opts ...option.RequestOption) (res *CertificateResponseSingleCustom, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_certificates", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// SSL Configuration Details
func (r *ZoneCustomCertificateService) Get(ctx context.Context, zoneID string, customCertificateID string, opts ...option.RequestOption) (res *CertificateResponseSingleCustom, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if customCertificateID == "" {
		err = errors.New("missing required custom_certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_certificates/%s", zoneID, customCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a new private key and/or PEM/CRT for the SSL certificate. Note: PATCHing
// a configuration for sni_custom certificates will result in a new resource id
// being returned, and the previous one being deleted.
func (r *ZoneCustomCertificateService) Update(ctx context.Context, zoneID string, customCertificateID string, body ZoneCustomCertificateUpdateParams, opts ...option.RequestOption) (res *CertificateResponseSingleCustom, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if customCertificateID == "" {
		err = errors.New("missing required custom_certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_certificates/%s", zoneID, customCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List, search, and filter all of your custom SSL certificates. The higher
// priority will break ties across overlapping 'legacy_custom' certificates, but
// 'legacy_custom' certificates will always supercede 'sni_custom' certificates.
func (r *ZoneCustomCertificateService) List(ctx context.Context, zoneID string, query ZoneCustomCertificateListParams, opts ...option.RequestOption) (res *CertificateResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_certificates", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Remove a SSL certificate from a zone.
func (r *ZoneCustomCertificateService) Delete(ctx context.Context, zoneID string, customCertificateID string, body ZoneCustomCertificateDeleteParams, opts ...option.RequestOption) (res *ZoneCustomCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if customCertificateID == "" {
		err = errors.New("missing required custom_certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_certificates/%s", zoneID, customCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// If a zone has multiple SSL certificates, you can set the order in which they
// should be used during a request. The higher priority will break ties across
// overlapping 'legacy_custom' certificates.
func (r *ZoneCustomCertificateService) Prioritize(ctx context.Context, zoneID string, body ZoneCustomCertificatePrioritizeParams, opts ...option.RequestOption) (res *CertificateResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_certificates/prioritize", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type BundleMethod string

const (
	BundleMethodUbiquitous BundleMethod = "ubiquitous"
	BundleMethodOptimal    BundleMethod = "optimal"
	BundleMethodForce      BundleMethod = "force"
)

func (r BundleMethod) IsKnown() bool {
	switch r {
	case BundleMethodUbiquitous, BundleMethodOptimal, BundleMethodForce:
		return true
	}
	return false
}

type CertificateResponseCollection struct {
	Result []CustomCertificate               `json:"result"`
	JSON   certificateResponseCollectionJSON `json:"-"`
	APIResponseCollectionTlsCertificates
}

// certificateResponseCollectionJSON contains the JSON metadata for the struct
// [CertificateResponseCollection]
type certificateResponseCollectionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type CertificateResponseSingleCustom struct {
	Result CustomCertificate                   `json:"result"`
	JSON   certificateResponseSingleCustomJSON `json:"-"`
	APIResponseSingleTlsCertificates
}

// certificateResponseSingleCustomJSON contains the JSON metadata for the struct
// [CertificateResponseSingleCustom]
type certificateResponseSingleCustomJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseSingleCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateResponseSingleCustomJSON) RawJSON() string {
	return r.raw
}

type CustomCertificate struct {
	// Identifier
	ID string `json:"id,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod BundleMethod `json:"bundle_method,required"`
	// When the certificate from the authority expires.
	ExpiresOn time.Time `json:"expires_on,required" format:"date-time"`
	Hosts     []string  `json:"hosts,required"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer,required"`
	// When the certificate was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The order/priority in which the certificate will be used in a request. The
	// higher priority will break ties across overlapping 'legacy_custom' certificates,
	// but 'legacy_custom' certificates will always supercede 'sni_custom'
	// certificates.
	Priority float64 `json:"priority,required"`
	// The type of hash used for the certificate.
	Signature string `json:"signature,required"`
	// Status of the zone's custom SSL.
	Status CustomCertificateStatus `json:"status,required"`
	// When the certificate was uploaded to Cloudflare.
	UploadedOn time.Time `json:"uploaded_on,required" format:"date-time"`
	// Identifier
	ZoneID string `json:"zone_id,required"`
	// Specify the region where your private key can be held locally for optimal TLS
	// performance. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Options allow distribution to
	// only to U.S. data centers, only to E.U. data centers, or only to highest
	// security data centers. Default distribution is to all Cloudflare datacenters,
	// for optimal performance.
	GeoRestrictions GeoRestrictions    `json:"geo_restrictions"`
	KeylessServer   KeylessCertificate `json:"keyless_server"`
	// Specify the policy that determines the region where your private key will be
	// held locally. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Any combination of countries,
	// specified by their two letter country code
	// (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements)
	// can be chosen, such as 'country: IN', as well as 'region: EU' which refers to
	// the EU region. If there are too few data centers satisfying the policy, it will
	// be rejected.
	Policy string                `json:"policy"`
	JSON   customCertificateJSON `json:"-"`
}

// customCertificateJSON contains the JSON metadata for the struct
// [CustomCertificate]
type customCertificateJSON struct {
	ID              apijson.Field
	BundleMethod    apijson.Field
	ExpiresOn       apijson.Field
	Hosts           apijson.Field
	Issuer          apijson.Field
	ModifiedOn      apijson.Field
	Priority        apijson.Field
	Signature       apijson.Field
	Status          apijson.Field
	UploadedOn      apijson.Field
	ZoneID          apijson.Field
	GeoRestrictions apijson.Field
	KeylessServer   apijson.Field
	Policy          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCertificateJSON) RawJSON() string {
	return r.raw
}

// Status of the zone's custom SSL.
type CustomCertificateStatus string

const (
	CustomCertificateStatusActive       CustomCertificateStatus = "active"
	CustomCertificateStatusExpired      CustomCertificateStatus = "expired"
	CustomCertificateStatusDeleted      CustomCertificateStatus = "deleted"
	CustomCertificateStatusPending      CustomCertificateStatus = "pending"
	CustomCertificateStatusInitializing CustomCertificateStatus = "initializing"
)

func (r CustomCertificateStatus) IsKnown() bool {
	switch r {
	case CustomCertificateStatusActive, CustomCertificateStatusExpired, CustomCertificateStatusDeleted, CustomCertificateStatusPending, CustomCertificateStatusInitializing:
		return true
	}
	return false
}

// Specify the region where your private key can be held locally for optimal TLS
// performance. HTTPS connections to any excluded data center will still be fully
// encrypted, but will incur some latency while Keyless SSL is used to complete the
// handshake with the nearest allowed data center. Options allow distribution to
// only to U.S. data centers, only to E.U. data centers, or only to highest
// security data centers. Default distribution is to all Cloudflare datacenters,
// for optimal performance.
type GeoRestrictions struct {
	Label GeoRestrictionsLabel `json:"label"`
	JSON  geoRestrictionsJSON  `json:"-"`
}

// geoRestrictionsJSON contains the JSON metadata for the struct [GeoRestrictions]
type geoRestrictionsJSON struct {
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeoRestrictions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geoRestrictionsJSON) RawJSON() string {
	return r.raw
}

type GeoRestrictionsLabel string

const (
	GeoRestrictionsLabelUs              GeoRestrictionsLabel = "us"
	GeoRestrictionsLabelEu              GeoRestrictionsLabel = "eu"
	GeoRestrictionsLabelHighestSecurity GeoRestrictionsLabel = "highest_security"
)

func (r GeoRestrictionsLabel) IsKnown() bool {
	switch r {
	case GeoRestrictionsLabelUs, GeoRestrictionsLabelEu, GeoRestrictionsLabelHighestSecurity:
		return true
	}
	return false
}

// Specify the region where your private key can be held locally for optimal TLS
// performance. HTTPS connections to any excluded data center will still be fully
// encrypted, but will incur some latency while Keyless SSL is used to complete the
// handshake with the nearest allowed data center. Options allow distribution to
// only to U.S. data centers, only to E.U. data centers, or only to highest
// security data centers. Default distribution is to all Cloudflare datacenters,
// for optimal performance.
type GeoRestrictionsParam struct {
	Label param.Field[GeoRestrictionsLabel] `json:"label"`
}

func (r GeoRestrictionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneCustomCertificateDeleteResponse struct {
	Result ZoneCustomCertificateDeleteResponseResult `json:"result"`
	JSON   zoneCustomCertificateDeleteResponseJSON   `json:"-"`
	APIResponseSingleTlsCertificates
}

// zoneCustomCertificateDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneCustomCertificateDeleteResponse]
type zoneCustomCertificateDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomCertificateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCustomCertificateDeleteResponseResult struct {
	// Identifier
	ID   string                                        `json:"id"`
	JSON zoneCustomCertificateDeleteResponseResultJSON `json:"-"`
}

// zoneCustomCertificateDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneCustomCertificateDeleteResponseResult]
type zoneCustomCertificateDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomCertificateDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomCertificateDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneCustomCertificateNewParams struct {
	// The zone's SSL certificate or certificate and the intermediate(s).
	Certificate param.Field[string] `json:"certificate,required"`
	// The zone's private key.
	PrivateKey param.Field[string] `json:"private_key,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[BundleMethod] `json:"bundle_method"`
	// Specify the region where your private key can be held locally for optimal TLS
	// performance. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Options allow distribution to
	// only to U.S. data centers, only to E.U. data centers, or only to highest
	// security data centers. Default distribution is to all Cloudflare datacenters,
	// for optimal performance.
	GeoRestrictions param.Field[GeoRestrictionsParam] `json:"geo_restrictions"`
	// Specify the policy that determines the region where your private key will be
	// held locally. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Any combination of countries,
	// specified by their two letter country code
	// (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements)
	// can be chosen, such as 'country: IN', as well as 'region: EU' which refers to
	// the EU region. If there are too few data centers satisfying the policy, it will
	// be rejected.
	Policy param.Field[string] `json:"policy"`
	// The type 'legacy_custom' enables support for legacy clients which do not include
	// SNI in the TLS handshake.
	Type param.Field[ZoneCustomCertificateNewParamsType] `json:"type"`
}

func (r ZoneCustomCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type 'legacy_custom' enables support for legacy clients which do not include
// SNI in the TLS handshake.
type ZoneCustomCertificateNewParamsType string

const (
	ZoneCustomCertificateNewParamsTypeLegacyCustom ZoneCustomCertificateNewParamsType = "legacy_custom"
	ZoneCustomCertificateNewParamsTypeSniCustom    ZoneCustomCertificateNewParamsType = "sni_custom"
)

func (r ZoneCustomCertificateNewParamsType) IsKnown() bool {
	switch r {
	case ZoneCustomCertificateNewParamsTypeLegacyCustom, ZoneCustomCertificateNewParamsTypeSniCustom:
		return true
	}
	return false
}

type ZoneCustomCertificateUpdateParams struct {
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[BundleMethod] `json:"bundle_method"`
	// The zone's SSL certificate or certificate and the intermediate(s).
	Certificate param.Field[string] `json:"certificate"`
	// Specify the region where your private key can be held locally for optimal TLS
	// performance. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Options allow distribution to
	// only to U.S. data centers, only to E.U. data centers, or only to highest
	// security data centers. Default distribution is to all Cloudflare datacenters,
	// for optimal performance.
	GeoRestrictions param.Field[GeoRestrictionsParam] `json:"geo_restrictions"`
	// Specify the policy that determines the region where your private key will be
	// held locally. HTTPS connections to any excluded data center will still be fully
	// encrypted, but will incur some latency while Keyless SSL is used to complete the
	// handshake with the nearest allowed data center. Any combination of countries,
	// specified by their two letter country code
	// (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements)
	// can be chosen, such as 'country: IN', as well as 'region: EU' which refers to
	// the EU region. If there are too few data centers satisfying the policy, it will
	// be rejected.
	Policy param.Field[string] `json:"policy"`
	// The zone's private key.
	PrivateKey param.Field[string] `json:"private_key"`
}

func (r ZoneCustomCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneCustomCertificateListParams struct {
	// Whether to match all search requirements or at least one (any).
	Match param.Field[ZoneCustomCertificateListParamsMatch] `query:"match"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of zones per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Status of the zone's custom SSL.
	Status param.Field[ZoneCustomCertificateListParamsStatus] `query:"status"`
}

// URLQuery serializes [ZoneCustomCertificateListParams]'s query parameters as
// `url.Values`.
func (r ZoneCustomCertificateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether to match all search requirements or at least one (any).
type ZoneCustomCertificateListParamsMatch string

const (
	ZoneCustomCertificateListParamsMatchAny ZoneCustomCertificateListParamsMatch = "any"
	ZoneCustomCertificateListParamsMatchAll ZoneCustomCertificateListParamsMatch = "all"
)

func (r ZoneCustomCertificateListParamsMatch) IsKnown() bool {
	switch r {
	case ZoneCustomCertificateListParamsMatchAny, ZoneCustomCertificateListParamsMatchAll:
		return true
	}
	return false
}

// Status of the zone's custom SSL.
type ZoneCustomCertificateListParamsStatus string

const (
	ZoneCustomCertificateListParamsStatusActive       ZoneCustomCertificateListParamsStatus = "active"
	ZoneCustomCertificateListParamsStatusExpired      ZoneCustomCertificateListParamsStatus = "expired"
	ZoneCustomCertificateListParamsStatusDeleted      ZoneCustomCertificateListParamsStatus = "deleted"
	ZoneCustomCertificateListParamsStatusPending      ZoneCustomCertificateListParamsStatus = "pending"
	ZoneCustomCertificateListParamsStatusInitializing ZoneCustomCertificateListParamsStatus = "initializing"
)

func (r ZoneCustomCertificateListParamsStatus) IsKnown() bool {
	switch r {
	case ZoneCustomCertificateListParamsStatusActive, ZoneCustomCertificateListParamsStatusExpired, ZoneCustomCertificateListParamsStatusDeleted, ZoneCustomCertificateListParamsStatusPending, ZoneCustomCertificateListParamsStatusInitializing:
		return true
	}
	return false
}

type ZoneCustomCertificateDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneCustomCertificateDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneCustomCertificatePrioritizeParams struct {
	// Array of ordered certificates.
	Certificates param.Field[[]ZoneCustomCertificatePrioritizeParamsCertificate] `json:"certificates,required"`
}

func (r ZoneCustomCertificatePrioritizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneCustomCertificatePrioritizeParamsCertificate struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// The order/priority in which the certificate will be used in a request. The
	// higher priority will break ties across overlapping 'legacy_custom' certificates,
	// but 'legacy_custom' certificates will always supercede 'sni_custom'
	// certificates.
	Priority param.Field[float64] `json:"priority"`
}

func (r ZoneCustomCertificatePrioritizeParamsCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
