// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneSslVerificationService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSslVerificationService] method instead.
type ZoneSslVerificationService struct {
	Options []option.RequestOption
}

// NewZoneSslVerificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSslVerificationService(opts ...option.RequestOption) (r *ZoneSslVerificationService) {
	r = &ZoneSslVerificationService{}
	r.Options = opts
	return
}

// Get SSL Verification Info for a Zone.
func (r *ZoneSslVerificationService) Get(ctx context.Context, zoneID string, query ZoneSslVerificationGetParams, opts ...option.RequestOption) (res *ZoneSslVerificationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/verification", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Edit SSL validation method for a certificate pack. A PATCH request will request
// an immediate validation check on any certificate, and return the updated status.
// If a validation method is provided, the validation will be immediately attempted
// using that method.
func (r *ZoneSslVerificationService) Update(ctx context.Context, zoneID string, certificatePackID string, body ZoneSslVerificationUpdateParams, opts ...option.RequestOption) (res *ZoneSslVerificationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificatePackID == "" {
		err = errors.New("missing required certificate_pack_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/verification/%s", zoneID, certificatePackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Desired validation method.
type ValidationMethodDefinition string

const (
	ValidationMethodDefinitionHTTP  ValidationMethodDefinition = "http"
	ValidationMethodDefinitionCname ValidationMethodDefinition = "cname"
	ValidationMethodDefinitionTxt   ValidationMethodDefinition = "txt"
	ValidationMethodDefinitionEmail ValidationMethodDefinition = "email"
)

func (r ValidationMethodDefinition) IsKnown() bool {
	switch r {
	case ValidationMethodDefinitionHTTP, ValidationMethodDefinitionCname, ValidationMethodDefinitionTxt, ValidationMethodDefinitionEmail:
		return true
	}
	return false
}

type ZoneSslVerificationGetResponse struct {
	Result []ZoneSslVerificationGetResponseResult `json:"result"`
	JSON   zoneSslVerificationGetResponseJSON     `json:"-"`
}

// zoneSslVerificationGetResponseJSON contains the JSON metadata for the struct
// [ZoneSslVerificationGetResponse]
type zoneSslVerificationGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslVerificationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslVerificationGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSslVerificationGetResponseResult struct {
	// Current status of certificate.
	CertificateStatus ZoneSslVerificationGetResponseResultCertificateStatus `json:"certificate_status,required"`
	// Certificate Authority is manually reviewing the order.
	BrandCheck bool `json:"brand_check"`
	// Certificate Pack UUID.
	CertPackUuid string `json:"cert_pack_uuid"`
	// Certificate's signature algorithm.
	Signature ZoneSslVerificationGetResponseResultSignature `json:"signature"`
	// Validation method in use for a certificate pack order.
	ValidationMethod ZoneSslVerificationGetResponseResultValidationMethod `json:"validation_method"`
	// Certificate's required verification information.
	VerificationInfo ZoneSslVerificationGetResponseResultVerificationInfo `json:"verification_info"`
	// Status of the required verification information, omitted if verification status
	// is unknown.
	VerificationStatus bool `json:"verification_status"`
	// Method of verification.
	VerificationType ZoneSslVerificationGetResponseResultVerificationType `json:"verification_type"`
	JSON             zoneSslVerificationGetResponseResultJSON             `json:"-"`
}

// zoneSslVerificationGetResponseResultJSON contains the JSON metadata for the
// struct [ZoneSslVerificationGetResponseResult]
type zoneSslVerificationGetResponseResultJSON struct {
	CertificateStatus  apijson.Field
	BrandCheck         apijson.Field
	CertPackUuid       apijson.Field
	Signature          apijson.Field
	ValidationMethod   apijson.Field
	VerificationInfo   apijson.Field
	VerificationStatus apijson.Field
	VerificationType   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneSslVerificationGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslVerificationGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// Current status of certificate.
type ZoneSslVerificationGetResponseResultCertificateStatus string

const (
	ZoneSslVerificationGetResponseResultCertificateStatusInitializing      ZoneSslVerificationGetResponseResultCertificateStatus = "initializing"
	ZoneSslVerificationGetResponseResultCertificateStatusAuthorizing       ZoneSslVerificationGetResponseResultCertificateStatus = "authorizing"
	ZoneSslVerificationGetResponseResultCertificateStatusActive            ZoneSslVerificationGetResponseResultCertificateStatus = "active"
	ZoneSslVerificationGetResponseResultCertificateStatusExpired           ZoneSslVerificationGetResponseResultCertificateStatus = "expired"
	ZoneSslVerificationGetResponseResultCertificateStatusIssuing           ZoneSslVerificationGetResponseResultCertificateStatus = "issuing"
	ZoneSslVerificationGetResponseResultCertificateStatusTimingOut         ZoneSslVerificationGetResponseResultCertificateStatus = "timing_out"
	ZoneSslVerificationGetResponseResultCertificateStatusPendingDeployment ZoneSslVerificationGetResponseResultCertificateStatus = "pending_deployment"
)

func (r ZoneSslVerificationGetResponseResultCertificateStatus) IsKnown() bool {
	switch r {
	case ZoneSslVerificationGetResponseResultCertificateStatusInitializing, ZoneSslVerificationGetResponseResultCertificateStatusAuthorizing, ZoneSslVerificationGetResponseResultCertificateStatusActive, ZoneSslVerificationGetResponseResultCertificateStatusExpired, ZoneSslVerificationGetResponseResultCertificateStatusIssuing, ZoneSslVerificationGetResponseResultCertificateStatusTimingOut, ZoneSslVerificationGetResponseResultCertificateStatusPendingDeployment:
		return true
	}
	return false
}

// Certificate's signature algorithm.
type ZoneSslVerificationGetResponseResultSignature string

const (
	ZoneSslVerificationGetResponseResultSignatureEcdsaWithSha256 ZoneSslVerificationGetResponseResultSignature = "ECDSAWithSHA256"
	ZoneSslVerificationGetResponseResultSignatureSha1WithRsa     ZoneSslVerificationGetResponseResultSignature = "SHA1WithRSA"
	ZoneSslVerificationGetResponseResultSignatureSha256WithRsa   ZoneSslVerificationGetResponseResultSignature = "SHA256WithRSA"
)

func (r ZoneSslVerificationGetResponseResultSignature) IsKnown() bool {
	switch r {
	case ZoneSslVerificationGetResponseResultSignatureEcdsaWithSha256, ZoneSslVerificationGetResponseResultSignatureSha1WithRsa, ZoneSslVerificationGetResponseResultSignatureSha256WithRsa:
		return true
	}
	return false
}

// Validation method in use for a certificate pack order.
type ZoneSslVerificationGetResponseResultValidationMethod string

const (
	ZoneSslVerificationGetResponseResultValidationMethodHTTP  ZoneSslVerificationGetResponseResultValidationMethod = "http"
	ZoneSslVerificationGetResponseResultValidationMethodCname ZoneSslVerificationGetResponseResultValidationMethod = "cname"
	ZoneSslVerificationGetResponseResultValidationMethodTxt   ZoneSslVerificationGetResponseResultValidationMethod = "txt"
)

func (r ZoneSslVerificationGetResponseResultValidationMethod) IsKnown() bool {
	switch r {
	case ZoneSslVerificationGetResponseResultValidationMethodHTTP, ZoneSslVerificationGetResponseResultValidationMethodCname, ZoneSslVerificationGetResponseResultValidationMethodTxt:
		return true
	}
	return false
}

// Certificate's required verification information.
type ZoneSslVerificationGetResponseResultVerificationInfo struct {
	// Name of CNAME record.
	RecordName ZoneSslVerificationGetResponseResultVerificationInfoRecordName `json:"record_name"`
	// Target of CNAME record.
	RecordTarget ZoneSslVerificationGetResponseResultVerificationInfoRecordTarget `json:"record_target"`
	JSON         zoneSslVerificationGetResponseResultVerificationInfoJSON         `json:"-"`
}

// zoneSslVerificationGetResponseResultVerificationInfoJSON contains the JSON
// metadata for the struct [ZoneSslVerificationGetResponseResultVerificationInfo]
type zoneSslVerificationGetResponseResultVerificationInfoJSON struct {
	RecordName   apijson.Field
	RecordTarget apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneSslVerificationGetResponseResultVerificationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslVerificationGetResponseResultVerificationInfoJSON) RawJSON() string {
	return r.raw
}

// Name of CNAME record.
type ZoneSslVerificationGetResponseResultVerificationInfoRecordName string

const (
	ZoneSslVerificationGetResponseResultVerificationInfoRecordNameRecordName ZoneSslVerificationGetResponseResultVerificationInfoRecordName = "record_name"
	ZoneSslVerificationGetResponseResultVerificationInfoRecordNameHTTPURL    ZoneSslVerificationGetResponseResultVerificationInfoRecordName = "http_url"
	ZoneSslVerificationGetResponseResultVerificationInfoRecordNameCname      ZoneSslVerificationGetResponseResultVerificationInfoRecordName = "cname"
	ZoneSslVerificationGetResponseResultVerificationInfoRecordNameTxtName    ZoneSslVerificationGetResponseResultVerificationInfoRecordName = "txt_name"
)

func (r ZoneSslVerificationGetResponseResultVerificationInfoRecordName) IsKnown() bool {
	switch r {
	case ZoneSslVerificationGetResponseResultVerificationInfoRecordNameRecordName, ZoneSslVerificationGetResponseResultVerificationInfoRecordNameHTTPURL, ZoneSslVerificationGetResponseResultVerificationInfoRecordNameCname, ZoneSslVerificationGetResponseResultVerificationInfoRecordNameTxtName:
		return true
	}
	return false
}

// Target of CNAME record.
type ZoneSslVerificationGetResponseResultVerificationInfoRecordTarget string

const (
	ZoneSslVerificationGetResponseResultVerificationInfoRecordTargetRecordValue ZoneSslVerificationGetResponseResultVerificationInfoRecordTarget = "record_value"
	ZoneSslVerificationGetResponseResultVerificationInfoRecordTargetHTTPBody    ZoneSslVerificationGetResponseResultVerificationInfoRecordTarget = "http_body"
	ZoneSslVerificationGetResponseResultVerificationInfoRecordTargetCnameTarget ZoneSslVerificationGetResponseResultVerificationInfoRecordTarget = "cname_target"
	ZoneSslVerificationGetResponseResultVerificationInfoRecordTargetTxtValue    ZoneSslVerificationGetResponseResultVerificationInfoRecordTarget = "txt_value"
)

func (r ZoneSslVerificationGetResponseResultVerificationInfoRecordTarget) IsKnown() bool {
	switch r {
	case ZoneSslVerificationGetResponseResultVerificationInfoRecordTargetRecordValue, ZoneSslVerificationGetResponseResultVerificationInfoRecordTargetHTTPBody, ZoneSslVerificationGetResponseResultVerificationInfoRecordTargetCnameTarget, ZoneSslVerificationGetResponseResultVerificationInfoRecordTargetTxtValue:
		return true
	}
	return false
}

// Method of verification.
type ZoneSslVerificationGetResponseResultVerificationType string

const (
	ZoneSslVerificationGetResponseResultVerificationTypeCname   ZoneSslVerificationGetResponseResultVerificationType = "cname"
	ZoneSslVerificationGetResponseResultVerificationTypeMetaTag ZoneSslVerificationGetResponseResultVerificationType = "meta tag"
)

func (r ZoneSslVerificationGetResponseResultVerificationType) IsKnown() bool {
	switch r {
	case ZoneSslVerificationGetResponseResultVerificationTypeCname, ZoneSslVerificationGetResponseResultVerificationTypeMetaTag:
		return true
	}
	return false
}

type ZoneSslVerificationUpdateResponse struct {
	Result ZoneSslVerificationUpdateResponseResult `json:"result"`
	JSON   zoneSslVerificationUpdateResponseJSON   `json:"-"`
	APIResponseSingleTlsCertificates
}

// zoneSslVerificationUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSslVerificationUpdateResponse]
type zoneSslVerificationUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslVerificationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslVerificationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSslVerificationUpdateResponseResult struct {
	// Result status.
	Status string `json:"status"`
	// Desired validation method.
	ValidationMethod ValidationMethodDefinition                  `json:"validation_method"`
	JSON             zoneSslVerificationUpdateResponseResultJSON `json:"-"`
}

// zoneSslVerificationUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSslVerificationUpdateResponseResult]
type zoneSslVerificationUpdateResponseResultJSON struct {
	Status           apijson.Field
	ValidationMethod apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneSslVerificationUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslVerificationUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneSslVerificationGetParams struct {
	// Immediately retry SSL Verification.
	Retry param.Field[ZoneSslVerificationGetParamsRetry] `query:"retry"`
}

// URLQuery serializes [ZoneSslVerificationGetParams]'s query parameters as
// `url.Values`.
func (r ZoneSslVerificationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Immediately retry SSL Verification.
type ZoneSslVerificationGetParamsRetry bool

const (
	ZoneSslVerificationGetParamsRetryTrue ZoneSslVerificationGetParamsRetry = true
)

func (r ZoneSslVerificationGetParamsRetry) IsKnown() bool {
	switch r {
	case ZoneSslVerificationGetParamsRetryTrue:
		return true
	}
	return false
}

type ZoneSslVerificationUpdateParams struct {
	// Desired validation method.
	ValidationMethod param.Field[ValidationMethodDefinition] `json:"validation_method,required"`
}

func (r ZoneSslVerificationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
