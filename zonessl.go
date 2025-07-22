// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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

type CacheAPIResponseSingle struct {
	Errors   []CacheMessagesItem `json:"errors,required"`
	Messages []CacheMessagesItem `json:"messages,required"`
	// Whether the API call was successful
	Success CacheAPIResponseSingleSuccess `json:"success,required"`
	JSON    cacheAPIResponseSingleJSON    `json:"-"`
}

// cacheAPIResponseSingleJSON contains the JSON metadata for the struct
// [CacheAPIResponseSingle]
type cacheAPIResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheAPIResponseSingleJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CacheAPIResponseSingleSuccess bool

const (
	CacheAPIResponseSingleSuccessTrue CacheAPIResponseSingleSuccess = true
)

func (r CacheAPIResponseSingleSuccess) IsKnown() bool {
	switch r {
	case CacheAPIResponseSingleSuccessTrue:
		return true
	}
	return false
}

type CacheMessagesItem struct {
	Code    int64                 `json:"code,required"`
	Message string                `json:"message,required"`
	JSON    cacheMessagesItemJSON `json:"-"`
}

// cacheMessagesItemJSON contains the JSON metadata for the struct
// [CacheMessagesItem]
type cacheMessagesItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheMessagesItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheMessagesItemJSON) RawJSON() string {
	return r.raw
}

type ZoneSslAnalyzeCertificateResponse struct {
	Result interface{}                           `json:"result"`
	JSON   zoneSslAnalyzeCertificateResponseJSON `json:"-"`
	APIResponseSingleTlsCertificates
}

// zoneSslAnalyzeCertificateResponseJSON contains the JSON metadata for the struct
// [ZoneSslAnalyzeCertificateResponse]
type zoneSslAnalyzeCertificateResponseJSON struct {
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

type ZoneSslGetRecommendationResponse struct {
	Result ZoneSslGetRecommendationResponseResult `json:"result"`
	JSON   zoneSslGetRecommendationResponseJSON   `json:"-"`
	CacheAPIResponseSingle
}

// zoneSslGetRecommendationResponseJSON contains the JSON metadata for the struct
// [ZoneSslGetRecommendationResponse]
type zoneSslGetRecommendationResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslGetRecommendationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslGetRecommendationResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSslGetRecommendationResponseResult struct {
	// Identifier of a recommendation result.
	ID         string                                      `json:"id"`
	ModifiedOn time.Time                                   `json:"modified_on" format:"date-time"`
	Value      ZoneSslGetRecommendationResponseResultValue `json:"value"`
	JSON       zoneSslGetRecommendationResponseResultJSON  `json:"-"`
}

// zoneSslGetRecommendationResponseResultJSON contains the JSON metadata for the
// struct [ZoneSslGetRecommendationResponseResult]
type zoneSslGetRecommendationResponseResultJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslGetRecommendationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSslGetRecommendationResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneSslGetRecommendationResponseResultValue string

const (
	ZoneSslGetRecommendationResponseResultValueFlexible ZoneSslGetRecommendationResponseResultValue = "flexible"
	ZoneSslGetRecommendationResponseResultValueFull     ZoneSslGetRecommendationResponseResultValue = "full"
	ZoneSslGetRecommendationResponseResultValueStrict   ZoneSslGetRecommendationResponseResultValue = "strict"
)

func (r ZoneSslGetRecommendationResponseResultValue) IsKnown() bool {
	switch r {
	case ZoneSslGetRecommendationResponseResultValueFlexible, ZoneSslGetRecommendationResponseResultValueFull, ZoneSslGetRecommendationResponseResultValueStrict:
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
