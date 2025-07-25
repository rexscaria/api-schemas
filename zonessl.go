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

// ZoneSslService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSslService] method instead.
type ZoneSslService struct {
	Options          []option.RequestOption
	CertificatePacks *ZoneSslCertificatePackService
	Universal        *ZoneSslUniversalService
	Verification     *ZoneSslVerificationService
}

// NewZoneSslService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewZoneSslService(opts ...option.RequestOption) (r *ZoneSslService) {
	r = &ZoneSslService{}
	r.Options = opts
	r.CertificatePacks = NewZoneSslCertificatePackService(opts...)
	r.Universal = NewZoneSslUniversalService(opts...)
	r.Verification = NewZoneSslVerificationService(opts...)
	return
}

// Returns the set of hostnames, the signature algorithm, and the expiration date
// of the certificate.
func (r *ZoneSslService) AnalyzeCertificate(ctx context.Context, zoneID string, body ZoneSslAnalyzeCertificateParams, opts ...option.RequestOption) (res *ZoneSslAnalyzeCertificateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/analyze", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve the SSL/TLS Recommender's recommendation for a zone.
//
// Deprecated: SSL/TLS Recommender has been decommissioned in favor of Automatic
// SSL/TLS
func (r *ZoneSslService) GetRecommendation(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneSslGetRecommendationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/recommendation", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSslAnalyzeCertificateResponse struct {
	Errors   []MessagesTlsCertificatesItem `json:"errors,required"`
	Messages []MessagesTlsCertificatesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneSslAnalyzeCertificateResponseSuccess `json:"success,required"`
	Result  interface{}                              `json:"result"`
	JSON    zoneSslAnalyzeCertificateResponseJSON    `json:"-"`
}

// zoneSslAnalyzeCertificateResponseJSON contains the JSON metadata for the struct
// [ZoneSslAnalyzeCertificateResponse]
type zoneSslAnalyzeCertificateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslAnalyzeCertificateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslAnalyzeCertificateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneSslAnalyzeCertificateResponseSuccess bool

const (
	ZoneSslAnalyzeCertificateResponseSuccessTrue ZoneSslAnalyzeCertificateResponseSuccess = true
)

func (r ZoneSslAnalyzeCertificateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSslAnalyzeCertificateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneSslGetRecommendationResponse struct {
	Errors   []ZoneSslGetRecommendationResponseError   `json:"errors,required"`
	Messages []ZoneSslGetRecommendationResponseMessage `json:"messages,required"`
	Result   ZoneSslGetRecommendationResponseResult    `json:"result,required"`
	// Indicates the API call's success or failure.
	Success bool                                 `json:"success,required"`
	JSON    zoneSslGetRecommendationResponseJSON `json:"-"`
}

// zoneSslGetRecommendationResponseJSON contains the JSON metadata for the struct
// [ZoneSslGetRecommendationResponse]
type zoneSslGetRecommendationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslGetRecommendationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslGetRecommendationResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSslGetRecommendationResponseError struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ZoneSslGetRecommendationResponseErrorsSource `json:"source"`
	JSON             zoneSslGetRecommendationResponseErrorJSON    `json:"-"`
}

// zoneSslGetRecommendationResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSslGetRecommendationResponseError]
type zoneSslGetRecommendationResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneSslGetRecommendationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslGetRecommendationResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ZoneSslGetRecommendationResponseErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    zoneSslGetRecommendationResponseErrorsSourceJSON `json:"-"`
}

// zoneSslGetRecommendationResponseErrorsSourceJSON contains the JSON metadata for
// the struct [ZoneSslGetRecommendationResponseErrorsSource]
type zoneSslGetRecommendationResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslGetRecommendationResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslGetRecommendationResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneSslGetRecommendationResponseMessage struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ZoneSslGetRecommendationResponseMessagesSource `json:"source"`
	JSON             zoneSslGetRecommendationResponseMessageJSON    `json:"-"`
}

// zoneSslGetRecommendationResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSslGetRecommendationResponseMessage]
type zoneSslGetRecommendationResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneSslGetRecommendationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslGetRecommendationResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ZoneSslGetRecommendationResponseMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    zoneSslGetRecommendationResponseMessagesSourceJSON `json:"-"`
}

// zoneSslGetRecommendationResponseMessagesSourceJSON contains the JSON metadata
// for the struct [ZoneSslGetRecommendationResponseMessagesSource]
type zoneSslGetRecommendationResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslGetRecommendationResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslGetRecommendationResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneSslGetRecommendationResponseResult struct {
	ID string `json:"id,required"`
	// Whether this setting can be updated or not.
	Editable bool `json:"editable,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Current setting of the automatic SSL/TLS.
	Value ZoneSslGetRecommendationResponseResultValue `json:"value,required"`
	// Next time this zone will be scanned by the Automatic SSL/TLS.
	NextScheduledScan time.Time                                  `json:"next_scheduled_scan,nullable" format:"date-time"`
	JSON              zoneSslGetRecommendationResponseResultJSON `json:"-"`
}

// zoneSslGetRecommendationResponseResultJSON contains the JSON metadata for the
// struct [ZoneSslGetRecommendationResponseResult]
type zoneSslGetRecommendationResponseResultJSON struct {
	ID                apijson.Field
	Editable          apijson.Field
	ModifiedOn        apijson.Field
	Value             apijson.Field
	NextScheduledScan apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSslGetRecommendationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslGetRecommendationResponseResultJSON) RawJSON() string {
	return r.raw
}

// Current setting of the automatic SSL/TLS.
type ZoneSslGetRecommendationResponseResultValue string

const (
	ZoneSslGetRecommendationResponseResultValueAuto   ZoneSslGetRecommendationResponseResultValue = "auto"
	ZoneSslGetRecommendationResponseResultValueCustom ZoneSslGetRecommendationResponseResultValue = "custom"
)

func (r ZoneSslGetRecommendationResponseResultValue) IsKnown() bool {
	switch r {
	case ZoneSslGetRecommendationResponseResultValueAuto, ZoneSslGetRecommendationResponseResultValueCustom:
		return true
	}
	return false
}

type ZoneSslAnalyzeCertificateParams struct {
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[BundleMethod] `json:"bundle_method"`
	// The zone's SSL certificate or certificate and the intermediate(s).
	Certificate param.Field[string] `json:"certificate"`
}

func (r ZoneSslAnalyzeCertificateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
