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

// ZoneSslCertificatePackService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSslCertificatePackService] method instead.
type ZoneSslCertificatePackService struct {
	Options []option.RequestOption
}

// NewZoneSslCertificatePackService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSslCertificatePackService(opts ...option.RequestOption) (r *ZoneSslCertificatePackService) {
	r = &ZoneSslCertificatePackService{}
	r.Options = opts
	return
}

// For a given zone, get a certificate pack.
func (r *ZoneSslCertificatePackService) Get(ctx context.Context, zoneID string, certificatePackID string, opts ...option.RequestOption) (res *ZoneSslCertificatePackGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificatePackID == "" {
		err = errors.New("missing required certificate_pack_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", zoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// For a given zone, restart validation or add cloudflare branding for an advanced
// certificate pack. The former is only a validation operation for a Certificate
// Pack in a validation_timed_out status.
func (r *ZoneSslCertificatePackService) Update(ctx context.Context, zoneID string, certificatePackID string, body ZoneSslCertificatePackUpdateParams, opts ...option.RequestOption) (res *AdvancedCertificatePackResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificatePackID == "" {
		err = errors.New("missing required certificate_pack_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", zoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// For a given zone, list all active certificate packs.
func (r *ZoneSslCertificatePackService) List(ctx context.Context, zoneID string, query ZoneSslCertificatePackListParams, opts ...option.RequestOption) (res *ZoneSslCertificatePackListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// For a given zone, delete an advanced certificate pack.
func (r *ZoneSslCertificatePackService) Delete(ctx context.Context, zoneID string, certificatePackID string, body ZoneSslCertificatePackDeleteParams, opts ...option.RequestOption) (res *ZoneSslCertificatePackDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificatePackID == "" {
		err = errors.New("missing required certificate_pack_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/%s", zoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// For a given zone, list certificate pack quotas.
func (r *ZoneSslCertificatePackService) GetQuota(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneSslCertificatePackGetQuotaResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/quota", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// For a given zone, order an advanced certificate pack.
func (r *ZoneSslCertificatePackService) Order(ctx context.Context, zoneID string, body ZoneSslCertificatePackOrderParams, opts ...option.RequestOption) (res *AdvancedCertificatePackResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/order", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AdvancedCertificatePackResponseSingle struct {
	Result AdvancedCertificatePackResponseSingleResult `json:"result"`
	JSON   advancedCertificatePackResponseSingleJSON   `json:"-"`
	APIResponseSingleTlsCertificates
}

// advancedCertificatePackResponseSingleJSON contains the JSON metadata for the
// struct [AdvancedCertificatePackResponseSingle]
type advancedCertificatePackResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedCertificatePackResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedCertificatePackResponseSingleJSON) RawJSON() string {
	return r.raw
}

type AdvancedCertificatePackResponseSingleResult struct {
	// Identifier
	ID string `json:"id"`
	// Certificate Authority selected for the order. For information on any certificate
	// authority specific details or restrictions
	// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	CertificateAuthority CertificateAuthoritySslPack `json:"certificate_authority"`
	// Whether or not to add Cloudflare Branding for the order. This will add a
	// subdomain of sni.cloudflaressl.com as the Common Name if set to true.
	CloudflareBranding bool `json:"cloudflare_branding"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts []string `json:"hosts"`
	// Status of certificate pack.
	Status AdvancedCertificatePackResponseSingleResultStatus `json:"status"`
	// Type of certificate pack.
	Type AdvancedType `json:"type"`
	// Validation Method selected for the order.
	ValidationMethod ValidationMethod `json:"validation_method"`
	// Validity Days selected for the order.
	ValidityDays ValidityDays                                    `json:"validity_days"`
	JSON         advancedCertificatePackResponseSingleResultJSON `json:"-"`
}

// advancedCertificatePackResponseSingleResultJSON contains the JSON metadata for
// the struct [AdvancedCertificatePackResponseSingleResult]
type advancedCertificatePackResponseSingleResultJSON struct {
	ID                   apijson.Field
	CertificateAuthority apijson.Field
	CloudflareBranding   apijson.Field
	Hosts                apijson.Field
	Status               apijson.Field
	Type                 apijson.Field
	ValidationMethod     apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AdvancedCertificatePackResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedCertificatePackResponseSingleResultJSON) RawJSON() string {
	return r.raw
}

// Status of certificate pack.
type AdvancedCertificatePackResponseSingleResultStatus string

const (
	AdvancedCertificatePackResponseSingleResultStatusInitializing         AdvancedCertificatePackResponseSingleResultStatus = "initializing"
	AdvancedCertificatePackResponseSingleResultStatusPendingValidation    AdvancedCertificatePackResponseSingleResultStatus = "pending_validation"
	AdvancedCertificatePackResponseSingleResultStatusDeleted              AdvancedCertificatePackResponseSingleResultStatus = "deleted"
	AdvancedCertificatePackResponseSingleResultStatusPendingIssuance      AdvancedCertificatePackResponseSingleResultStatus = "pending_issuance"
	AdvancedCertificatePackResponseSingleResultStatusPendingDeployment    AdvancedCertificatePackResponseSingleResultStatus = "pending_deployment"
	AdvancedCertificatePackResponseSingleResultStatusPendingDeletion      AdvancedCertificatePackResponseSingleResultStatus = "pending_deletion"
	AdvancedCertificatePackResponseSingleResultStatusPendingExpiration    AdvancedCertificatePackResponseSingleResultStatus = "pending_expiration"
	AdvancedCertificatePackResponseSingleResultStatusExpired              AdvancedCertificatePackResponseSingleResultStatus = "expired"
	AdvancedCertificatePackResponseSingleResultStatusActive               AdvancedCertificatePackResponseSingleResultStatus = "active"
	AdvancedCertificatePackResponseSingleResultStatusInitializingTimedOut AdvancedCertificatePackResponseSingleResultStatus = "initializing_timed_out"
	AdvancedCertificatePackResponseSingleResultStatusValidationTimedOut   AdvancedCertificatePackResponseSingleResultStatus = "validation_timed_out"
	AdvancedCertificatePackResponseSingleResultStatusIssuanceTimedOut     AdvancedCertificatePackResponseSingleResultStatus = "issuance_timed_out"
	AdvancedCertificatePackResponseSingleResultStatusDeploymentTimedOut   AdvancedCertificatePackResponseSingleResultStatus = "deployment_timed_out"
	AdvancedCertificatePackResponseSingleResultStatusDeletionTimedOut     AdvancedCertificatePackResponseSingleResultStatus = "deletion_timed_out"
	AdvancedCertificatePackResponseSingleResultStatusPendingCleanup       AdvancedCertificatePackResponseSingleResultStatus = "pending_cleanup"
	AdvancedCertificatePackResponseSingleResultStatusStagingDeployment    AdvancedCertificatePackResponseSingleResultStatus = "staging_deployment"
	AdvancedCertificatePackResponseSingleResultStatusStagingActive        AdvancedCertificatePackResponseSingleResultStatus = "staging_active"
	AdvancedCertificatePackResponseSingleResultStatusDeactivating         AdvancedCertificatePackResponseSingleResultStatus = "deactivating"
	AdvancedCertificatePackResponseSingleResultStatusInactive             AdvancedCertificatePackResponseSingleResultStatus = "inactive"
	AdvancedCertificatePackResponseSingleResultStatusBackupIssued         AdvancedCertificatePackResponseSingleResultStatus = "backup_issued"
	AdvancedCertificatePackResponseSingleResultStatusHoldingDeployment    AdvancedCertificatePackResponseSingleResultStatus = "holding_deployment"
)

func (r AdvancedCertificatePackResponseSingleResultStatus) IsKnown() bool {
	switch r {
	case AdvancedCertificatePackResponseSingleResultStatusInitializing, AdvancedCertificatePackResponseSingleResultStatusPendingValidation, AdvancedCertificatePackResponseSingleResultStatusDeleted, AdvancedCertificatePackResponseSingleResultStatusPendingIssuance, AdvancedCertificatePackResponseSingleResultStatusPendingDeployment, AdvancedCertificatePackResponseSingleResultStatusPendingDeletion, AdvancedCertificatePackResponseSingleResultStatusPendingExpiration, AdvancedCertificatePackResponseSingleResultStatusExpired, AdvancedCertificatePackResponseSingleResultStatusActive, AdvancedCertificatePackResponseSingleResultStatusInitializingTimedOut, AdvancedCertificatePackResponseSingleResultStatusValidationTimedOut, AdvancedCertificatePackResponseSingleResultStatusIssuanceTimedOut, AdvancedCertificatePackResponseSingleResultStatusDeploymentTimedOut, AdvancedCertificatePackResponseSingleResultStatusDeletionTimedOut, AdvancedCertificatePackResponseSingleResultStatusPendingCleanup, AdvancedCertificatePackResponseSingleResultStatusStagingDeployment, AdvancedCertificatePackResponseSingleResultStatusStagingActive, AdvancedCertificatePackResponseSingleResultStatusDeactivating, AdvancedCertificatePackResponseSingleResultStatusInactive, AdvancedCertificatePackResponseSingleResultStatusBackupIssued, AdvancedCertificatePackResponseSingleResultStatusHoldingDeployment:
		return true
	}
	return false
}

// Type of certificate pack.
type AdvancedType string

const (
	AdvancedTypeAdvanced AdvancedType = "advanced"
)

func (r AdvancedType) IsKnown() bool {
	switch r {
	case AdvancedTypeAdvanced:
		return true
	}
	return false
}

// Certificate Authority selected for the order. For information on any certificate
// authority specific details or restrictions
// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
type CertificateAuthoritySslPack string

const (
	CertificateAuthoritySslPackGoogle      CertificateAuthoritySslPack = "google"
	CertificateAuthoritySslPackLetsEncrypt CertificateAuthoritySslPack = "lets_encrypt"
	CertificateAuthoritySslPackSslCom      CertificateAuthoritySslPack = "ssl_com"
)

func (r CertificateAuthoritySslPack) IsKnown() bool {
	switch r {
	case CertificateAuthoritySslPackGoogle, CertificateAuthoritySslPackLetsEncrypt, CertificateAuthoritySslPackSslCom:
		return true
	}
	return false
}

// Validation Method selected for the order.
type ValidationMethod string

const (
	ValidationMethodTxt   ValidationMethod = "txt"
	ValidationMethodHTTP  ValidationMethod = "http"
	ValidationMethodEmail ValidationMethod = "email"
)

func (r ValidationMethod) IsKnown() bool {
	switch r {
	case ValidationMethodTxt, ValidationMethodHTTP, ValidationMethodEmail:
		return true
	}
	return false
}

// Validity Days selected for the order.
type ValidityDays int64

const (
	ValidityDays14  ValidityDays = 14
	ValidityDays30  ValidityDays = 30
	ValidityDays90  ValidityDays = 90
	ValidityDays365 ValidityDays = 365
)

func (r ValidityDays) IsKnown() bool {
	switch r {
	case ValidityDays14, ValidityDays30, ValidityDays90, ValidityDays365:
		return true
	}
	return false
}

type ZoneSslCertificatePackGetResponse struct {
	Result interface{}                           `json:"result"`
	JSON   zoneSslCertificatePackGetResponseJSON `json:"-"`
	APIResponseSingleTlsCertificates
}

// zoneSslCertificatePackGetResponseJSON contains the JSON metadata for the struct
// [ZoneSslCertificatePackGetResponse]
type zoneSslCertificatePackGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslCertificatePackGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSslCertificatePackListResponse struct {
	Result []interface{}                          `json:"result"`
	JSON   zoneSslCertificatePackListResponseJSON `json:"-"`
	APIResponseCollectionTlsCertificates
}

// zoneSslCertificatePackListResponseJSON contains the JSON metadata for the struct
// [ZoneSslCertificatePackListResponse]
type zoneSslCertificatePackListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslCertificatePackListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSslCertificatePackDeleteResponse struct {
	Result ZoneSslCertificatePackDeleteResponseResult `json:"result"`
	JSON   zoneSslCertificatePackDeleteResponseJSON   `json:"-"`
	APIResponseSingleTlsCertificates
}

// zoneSslCertificatePackDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneSslCertificatePackDeleteResponse]
type zoneSslCertificatePackDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslCertificatePackDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSslCertificatePackDeleteResponseResult struct {
	// Identifier
	ID   string                                         `json:"id"`
	JSON zoneSslCertificatePackDeleteResponseResultJSON `json:"-"`
}

// zoneSslCertificatePackDeleteResponseResultJSON contains the JSON metadata for
// the struct [ZoneSslCertificatePackDeleteResponseResult]
type zoneSslCertificatePackDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslCertificatePackDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneSslCertificatePackGetQuotaResponse struct {
	Result ZoneSslCertificatePackGetQuotaResponseResult `json:"result"`
	JSON   zoneSslCertificatePackGetQuotaResponseJSON   `json:"-"`
	APIResponseSingleTlsCertificates
}

// zoneSslCertificatePackGetQuotaResponseJSON contains the JSON metadata for the
// struct [ZoneSslCertificatePackGetQuotaResponse]
type zoneSslCertificatePackGetQuotaResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackGetQuotaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslCertificatePackGetQuotaResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSslCertificatePackGetQuotaResponseResult struct {
	Advanced ZoneSslCertificatePackGetQuotaResponseResultAdvanced `json:"advanced"`
	JSON     zoneSslCertificatePackGetQuotaResponseResultJSON     `json:"-"`
}

// zoneSslCertificatePackGetQuotaResponseResultJSON contains the JSON metadata for
// the struct [ZoneSslCertificatePackGetQuotaResponseResult]
type zoneSslCertificatePackGetQuotaResponseResultJSON struct {
	Advanced    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackGetQuotaResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslCertificatePackGetQuotaResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneSslCertificatePackGetQuotaResponseResultAdvanced struct {
	// Quantity Allocated.
	Allocated int64 `json:"allocated"`
	// Quantity Used.
	Used int64                                                    `json:"used"`
	JSON zoneSslCertificatePackGetQuotaResponseResultAdvancedJSON `json:"-"`
}

// zoneSslCertificatePackGetQuotaResponseResultAdvancedJSON contains the JSON
// metadata for the struct [ZoneSslCertificatePackGetQuotaResponseResultAdvanced]
type zoneSslCertificatePackGetQuotaResponseResultAdvancedJSON struct {
	Allocated   apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackGetQuotaResponseResultAdvanced) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslCertificatePackGetQuotaResponseResultAdvancedJSON) RawJSON() string {
	return r.raw
}

type ZoneSslCertificatePackUpdateParams struct {
	// Whether or not to add Cloudflare Branding for the order. This will add a
	// subdomain of sni.cloudflaressl.com as the Common Name if set to true.
	CloudflareBranding param.Field[bool] `json:"cloudflare_branding"`
}

func (r ZoneSslCertificatePackUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSslCertificatePackListParams struct {
	// Include Certificate Packs of all statuses, not just active ones.
	Status param.Field[ZoneSslCertificatePackListParamsStatus] `query:"status"`
}

// URLQuery serializes [ZoneSslCertificatePackListParams]'s query parameters as
// `url.Values`.
func (r ZoneSslCertificatePackListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Include Certificate Packs of all statuses, not just active ones.
type ZoneSslCertificatePackListParamsStatus string

const (
	ZoneSslCertificatePackListParamsStatusAll ZoneSslCertificatePackListParamsStatus = "all"
)

func (r ZoneSslCertificatePackListParamsStatus) IsKnown() bool {
	switch r {
	case ZoneSslCertificatePackListParamsStatusAll:
		return true
	}
	return false
}

type ZoneSslCertificatePackDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneSslCertificatePackDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneSslCertificatePackOrderParams struct {
	// Certificate Authority selected for the order. For information on any certificate
	// authority specific details or restrictions
	// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	CertificateAuthority param.Field[CertificateAuthoritySslPack] `json:"certificate_authority,required"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts param.Field[[]string] `json:"hosts,required"`
	// Type of certificate pack.
	Type param.Field[AdvancedType] `json:"type,required"`
	// Validation Method selected for the order.
	ValidationMethod param.Field[ValidationMethod] `json:"validation_method,required"`
	// Validity Days selected for the order.
	ValidityDays param.Field[ValidityDays] `json:"validity_days,required"`
	// Whether or not to add Cloudflare Branding for the order. This will add a
	// subdomain of sni.cloudflaressl.com as the Common Name if set to true.
	CloudflareBranding param.Field[bool] `json:"cloudflare_branding"`
}

func (r ZoneSslCertificatePackOrderParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
