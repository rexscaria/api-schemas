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

// ZoneLeakedCredentialCheckDetectionService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneLeakedCredentialCheckDetectionService] method instead.
type ZoneLeakedCredentialCheckDetectionService struct {
	Options []option.RequestOption
}

// NewZoneLeakedCredentialCheckDetectionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZoneLeakedCredentialCheckDetectionService(opts ...option.RequestOption) (r *ZoneLeakedCredentialCheckDetectionService) {
	r = &ZoneLeakedCredentialCheckDetectionService{}
	r.Options = opts
	return
}

// Create user-defined detection pattern for Leaked Credential Checks.
func (r *ZoneLeakedCredentialCheckDetectionService) New(ctx context.Context, zoneID string, body ZoneLeakedCredentialCheckDetectionNewParams, opts ...option.RequestOption) (res *ResponseCustomDetection, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/leaked-credential-checks/detections", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update user-defined detection pattern for Leaked Credential Checks.
func (r *ZoneLeakedCredentialCheckDetectionService) Update(ctx context.Context, zoneID string, detectionID DetectionIDParam, body ZoneLeakedCredentialCheckDetectionUpdateParams, opts ...option.RequestOption) (res *ResponseCustomDetection, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if detectionID == "" {
		err = errors.New("missing required detection_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/leaked-credential-checks/detections/%s", zoneID, detectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List user-defined detection patterns for Leaked Credential Checks.
func (r *ZoneLeakedCredentialCheckDetectionService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneLeakedCredentialCheckDetectionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/leaked-credential-checks/detections", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove user-defined detection pattern for Leaked Credential Checks.
func (r *ZoneLeakedCredentialCheckDetectionService) Delete(ctx context.Context, zoneID string, detectionID DetectionIDParam, opts ...option.RequestOption) (res *APIResponseWafProductBundle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if detectionID == "" {
		err = errors.New("missing required detection_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/leaked-credential-checks/detections/%s", zoneID, detectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type APIResponseWafProductBundle struct {
	Errors   []APIResponseWafProductBundleError   `json:"errors,required"`
	Messages []APIResponseWafProductBundleMessage `json:"messages,required"`
	Result   interface{}                          `json:"result,required"`
	// Defines whether the API call was successful.
	Success APIResponseWafProductBundleSuccess `json:"success,required"`
	JSON    apiResponseWafProductBundleJSON    `json:"-"`
}

// apiResponseWafProductBundleJSON contains the JSON metadata for the struct
// [APIResponseWafProductBundle]
type apiResponseWafProductBundleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseWafProductBundle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseWafProductBundleJSON) RawJSON() string {
	return r.raw
}

type APIResponseWafProductBundleError struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           APIResponseWafProductBundleErrorsSource `json:"source"`
	JSON             apiResponseWafProductBundleErrorJSON    `json:"-"`
}

// apiResponseWafProductBundleErrorJSON contains the JSON metadata for the struct
// [APIResponseWafProductBundleError]
type apiResponseWafProductBundleErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIResponseWafProductBundleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseWafProductBundleErrorJSON) RawJSON() string {
	return r.raw
}

type APIResponseWafProductBundleErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    apiResponseWafProductBundleErrorsSourceJSON `json:"-"`
}

// apiResponseWafProductBundleErrorsSourceJSON contains the JSON metadata for the
// struct [APIResponseWafProductBundleErrorsSource]
type apiResponseWafProductBundleErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseWafProductBundleErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseWafProductBundleErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type APIResponseWafProductBundleMessage struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           APIResponseWafProductBundleMessagesSource `json:"source"`
	JSON             apiResponseWafProductBundleMessageJSON    `json:"-"`
}

// apiResponseWafProductBundleMessageJSON contains the JSON metadata for the struct
// [APIResponseWafProductBundleMessage]
type apiResponseWafProductBundleMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIResponseWafProductBundleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseWafProductBundleMessageJSON) RawJSON() string {
	return r.raw
}

type APIResponseWafProductBundleMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    apiResponseWafProductBundleMessagesSourceJSON `json:"-"`
}

// apiResponseWafProductBundleMessagesSourceJSON contains the JSON metadata for the
// struct [APIResponseWafProductBundleMessagesSource]
type apiResponseWafProductBundleMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseWafProductBundleMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseWafProductBundleMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type APIResponseWafProductBundleSuccess bool

const (
	APIResponseWafProductBundleSuccessTrue APIResponseWafProductBundleSuccess = true
)

func (r APIResponseWafProductBundleSuccess) IsKnown() bool {
	switch r {
	case APIResponseWafProductBundleSuccessTrue:
		return true
	}
	return false
}

// Defines a custom set of username/password expressions to match Leaked Credential
// Checks on.
type CustomDetection struct {
	// Defines the unique ID for this custom detection.
	ID DetectionID `json:"id"`
	// Defines ehe ruleset expression to use in matching the password in a request.
	Password string `json:"password"`
	// Defines the ruleset expression to use in matching the username in a request.
	Username string              `json:"username"`
	JSON     customDetectionJSON `json:"-"`
}

// customDetectionJSON contains the JSON metadata for the struct [CustomDetection]
type customDetectionJSON struct {
	ID          apijson.Field
	Password    apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customDetectionJSON) RawJSON() string {
	return r.raw
}

// Defines a custom set of username/password expressions to match Leaked Credential
// Checks on.
type CustomDetectionParam struct {
	// Defines ehe ruleset expression to use in matching the password in a request.
	Password param.Field[string] `json:"password"`
	// Defines the ruleset expression to use in matching the username in a request.
	Username param.Field[string] `json:"username"`
}

func (r CustomDetectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DetectionID = string

type DetectionIDParam = string

type ResponseCustomDetection struct {
	Errors   []WafProductAPIBundleMessages `json:"errors,required"`
	Messages []WafProductAPIBundleMessages `json:"messages,required"`
	// Defines a custom set of username/password expressions to match Leaked Credential
	// Checks on.
	Result CustomDetection `json:"result,required"`
	// Defines whether the API call was successful.
	Success ResponseCustomDetectionSuccess `json:"success,required"`
	JSON    responseCustomDetectionJSON    `json:"-"`
}

// responseCustomDetectionJSON contains the JSON metadata for the struct
// [ResponseCustomDetection]
type responseCustomDetectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCustomDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCustomDetectionJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type ResponseCustomDetectionSuccess bool

const (
	ResponseCustomDetectionSuccessTrue ResponseCustomDetectionSuccess = true
)

func (r ResponseCustomDetectionSuccess) IsKnown() bool {
	switch r {
	case ResponseCustomDetectionSuccessTrue:
		return true
	}
	return false
}

type ZoneLeakedCredentialCheckDetectionListResponse struct {
	Errors   []WafProductAPIBundleMessages `json:"errors,required"`
	Messages []WafProductAPIBundleMessages `json:"messages,required"`
	Result   []CustomDetection             `json:"result,required,nullable"`
	// Defines whether the API call was successful.
	Success ZoneLeakedCredentialCheckDetectionListResponseSuccess `json:"success,required"`
	JSON    zoneLeakedCredentialCheckDetectionListResponseJSON    `json:"-"`
}

// zoneLeakedCredentialCheckDetectionListResponseJSON contains the JSON metadata
// for the struct [ZoneLeakedCredentialCheckDetectionListResponse]
type zoneLeakedCredentialCheckDetectionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLeakedCredentialCheckDetectionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneLeakedCredentialCheckDetectionListResponseJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type ZoneLeakedCredentialCheckDetectionListResponseSuccess bool

const (
	ZoneLeakedCredentialCheckDetectionListResponseSuccessTrue ZoneLeakedCredentialCheckDetectionListResponseSuccess = true
)

func (r ZoneLeakedCredentialCheckDetectionListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneLeakedCredentialCheckDetectionListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneLeakedCredentialCheckDetectionNewParams struct {
	// Defines a custom set of username/password expressions to match Leaked Credential
	// Checks on.
	CustomDetection CustomDetectionParam `json:"custom_detection,required"`
}

func (r ZoneLeakedCredentialCheckDetectionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CustomDetection)
}

type ZoneLeakedCredentialCheckDetectionUpdateParams struct {
	// Defines a custom set of username/password expressions to match Leaked Credential
	// Checks on.
	CustomDetection CustomDetectionParam `json:"custom_detection,required"`
}

func (r ZoneLeakedCredentialCheckDetectionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CustomDetection)
}
