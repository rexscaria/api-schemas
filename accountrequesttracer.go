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

// AccountRequestTracerService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountRequestTracerService] method instead.
type AccountRequestTracerService struct {
	Options []option.RequestOption
}

// NewAccountRequestTracerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRequestTracerService(opts ...option.RequestOption) (r *AccountRequestTracerService) {
	r = &AccountRequestTracerService{}
	r.Options = opts
	return
}

// Request Trace
func (r *AccountRequestTracerService) Trace(ctx context.Context, accountID string, body AccountRequestTracerTraceParams, opts ...option.RequestOption) (res *AccountRequestTracerTraceResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/request-tracer/trace", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RequestTracerMessagesItems struct {
	Code             int64                            `json:"code,required"`
	Message          string                           `json:"message,required"`
	DocumentationURL string                           `json:"documentation_url"`
	Source           RequestTracerMessagesItemsSource `json:"source"`
	JSON             requestTracerMessagesItemsJSON   `json:"-"`
}

// requestTracerMessagesItemsJSON contains the JSON metadata for the struct
// [RequestTracerMessagesItems]
type requestTracerMessagesItemsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RequestTracerMessagesItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestTracerMessagesItemsJSON) RawJSON() string {
	return r.raw
}

type RequestTracerMessagesItemsSource struct {
	Pointer string                               `json:"pointer"`
	JSON    requestTracerMessagesItemsSourceJSON `json:"-"`
}

// requestTracerMessagesItemsSourceJSON contains the JSON metadata for the struct
// [RequestTracerMessagesItemsSource]
type requestTracerMessagesItemsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestTracerMessagesItemsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestTracerMessagesItemsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountRequestTracerTraceResponse struct {
	Errors   []RequestTracerMessagesItems `json:"errors,required"`
	Messages []RequestTracerMessagesItems `json:"messages,required"`
	// Whether the API call was successful
	Success AccountRequestTracerTraceResponseSuccess `json:"success,required"`
	// Trace result with an origin status code
	Result AccountRequestTracerTraceResponseResult `json:"result"`
	JSON   accountRequestTracerTraceResponseJSON   `json:"-"`
}

// accountRequestTracerTraceResponseJSON contains the JSON metadata for the struct
// [AccountRequestTracerTraceResponse]
type accountRequestTracerTraceResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRequestTracerTraceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRequestTracerTraceResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountRequestTracerTraceResponseSuccess bool

const (
	AccountRequestTracerTraceResponseSuccessTrue AccountRequestTracerTraceResponseSuccess = true
)

func (r AccountRequestTracerTraceResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRequestTracerTraceResponseSuccessTrue:
		return true
	}
	return false
}

// Trace result with an origin status code
type AccountRequestTracerTraceResponseResult struct {
	// HTTP Status code of zone response
	StatusCode int64                                          `json:"status_code"`
	Trace      []AccountRequestTracerTraceResponseResultTrace `json:"trace"`
	JSON       accountRequestTracerTraceResponseResultJSON    `json:"-"`
}

// accountRequestTracerTraceResponseResultJSON contains the JSON metadata for the
// struct [AccountRequestTracerTraceResponseResult]
type accountRequestTracerTraceResponseResultJSON struct {
	StatusCode  apijson.Field
	Trace       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRequestTracerTraceResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRequestTracerTraceResponseResultJSON) RawJSON() string {
	return r.raw
}

// List of steps acting on request/response
type AccountRequestTracerTraceResponseResultTrace struct {
	// If step type is rule, then action performed by this rule
	Action string `json:"action"`
	// If step type is rule, then action parameters of this rule as JSON
	ActionParameters interface{} `json:"action_parameters"`
	// If step type is rule or ruleset, the description of this entity
	Description string `json:"description"`
	// If step type is rule, then expression used to match for this rule
	Expression string `json:"expression"`
	// If step type is ruleset, then kind of this ruleset
	Kind string `json:"kind"`
	// Whether tracing step affected tracing request/response
	Matched bool `json:"matched"`
	// If step type is ruleset, then name of this ruleset
	Name string `json:"name"`
	// Tracing step identifying name
	StepName string      `json:"step_name"`
	Trace    interface{} `json:"trace"`
	// Tracing step type
	Type string                                           `json:"type"`
	JSON accountRequestTracerTraceResponseResultTraceJSON `json:"-"`
}

// accountRequestTracerTraceResponseResultTraceJSON contains the JSON metadata for
// the struct [AccountRequestTracerTraceResponseResultTrace]
type accountRequestTracerTraceResponseResultTraceJSON struct {
	Action           apijson.Field
	ActionParameters apijson.Field
	Description      apijson.Field
	Expression       apijson.Field
	Kind             apijson.Field
	Matched          apijson.Field
	Name             apijson.Field
	StepName         apijson.Field
	Trace            apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountRequestTracerTraceResponseResultTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRequestTracerTraceResponseResultTraceJSON) RawJSON() string {
	return r.raw
}

type AccountRequestTracerTraceParams struct {
	// HTTP Method of tracing request
	Method param.Field[string] `json:"method,required"`
	// URL to which perform tracing request
	URL  param.Field[string]                              `json:"url,required"`
	Body param.Field[AccountRequestTracerTraceParamsBody] `json:"body"`
	// Additional request parameters
	Context param.Field[AccountRequestTracerTraceParamsContext] `json:"context"`
	// Cookies added to tracing request
	Cookies param.Field[map[string]string] `json:"cookies"`
	// Headers added to tracing request
	Headers param.Field[map[string]string] `json:"headers"`
	// HTTP Protocol of tracing request
	Protocol param.Field[string] `json:"protocol"`
	// Skip sending the request to the Origin server after all rules evaluation
	SkipResponse param.Field[bool] `json:"skip_response"`
}

func (r AccountRequestTracerTraceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRequestTracerTraceParamsBody struct {
	// Base64 encoded request body
	Base64 param.Field[string] `json:"base64"`
	// Arbitrary json as request body
	Json param.Field[interface{}] `json:"json"`
	// Request body as plain text
	PlainText param.Field[string] `json:"plain_text"`
}

func (r AccountRequestTracerTraceParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional request parameters
type AccountRequestTracerTraceParamsContext struct {
	// Bot score used for evaluating tracing request processing
	BotScore param.Field[int64] `json:"bot_score"`
	// Geodata for tracing request
	Geoloc param.Field[AccountRequestTracerTraceParamsContextGeoloc] `json:"geoloc"`
	// Whether to skip any challenges for tracing request (e.g.: captcha)
	SkipChallenge param.Field[bool] `json:"skip_challenge"`
	// Threat score used for evaluating tracing request processing
	ThreatScore param.Field[int64] `json:"threat_score"`
}

func (r AccountRequestTracerTraceParamsContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Geodata for tracing request
type AccountRequestTracerTraceParamsContextGeoloc struct {
	City                param.Field[string]  `json:"city"`
	Continent           param.Field[string]  `json:"continent"`
	IsEuCountry         param.Field[bool]    `json:"is_eu_country"`
	ISOCode             param.Field[string]  `json:"iso_code"`
	Latitude            param.Field[float64] `json:"latitude"`
	Longitude           param.Field[float64] `json:"longitude"`
	PostalCode          param.Field[string]  `json:"postal_code"`
	RegionCode          param.Field[string]  `json:"region_code"`
	Subdivision2ISOCode param.Field[string]  `json:"subdivision_2_iso_code"`
	Timezone            param.Field[string]  `json:"timezone"`
}

func (r AccountRequestTracerTraceParamsContextGeoloc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
