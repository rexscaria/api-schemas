// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/stainless-sdks/cf-rex-go/shared"
	"github.com/tidwall/gjson"
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

// Create user-defined detection pattern for Leaked Credential Checks
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

// Update user-defined detection pattern for Leaked Credential Checks
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

// List user-defined detection patterns for Leaked Credential Checks
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

// Remove user-defined detection pattern for Leaked Credential Checks
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
	Errors   []WafProductAPIBundleMessages          `json:"errors,required"`
	Messages []WafProductAPIBundleMessages          `json:"messages,required"`
	Result   APIResponseWafProductBundleResultUnion `json:"result,required"`
	// Whether the API call was successful
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

// Union satisfied by [APIResponseWafProductBundleResultArray] or
// [shared.UnionString].
type APIResponseWafProductBundleResultUnion interface {
	ImplementsAPIResponseWafProductBundleResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIResponseWafProductBundleResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIResponseWafProductBundleResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type APIResponseWafProductBundleResultArray []interface{}

func (r APIResponseWafProductBundleResultArray) ImplementsAPIResponseWafProductBundleResultUnion() {}

// Whether the API call was successful
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

// A custom set of username/password expressions to match Leaked Credential Checks
// on
type CustomDetection struct {
	// Identifier
	ID DetectionID `json:"id"`
	// The ruleset expression to use in matching the password in a request
	Password string `json:"password"`
	// The ruleset expression to use in matching the username in a request
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

// A custom set of username/password expressions to match Leaked Credential Checks
// on
type CustomDetectionParam struct {
	// The ruleset expression to use in matching the password in a request
	Password param.Field[string] `json:"password"`
	// The ruleset expression to use in matching the username in a request
	Username param.Field[string] `json:"username"`
}

func (r CustomDetectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DetectionID = string

type DetectionIDParam = string

type ResponseCustomDetection struct {
	// A custom set of username/password expressions to match Leaked Credential Checks
	// on
	Result CustomDetection             `json:"result"`
	JSON   responseCustomDetectionJSON `json:"-"`
	APIResponseSingleWafProductBundle
}

// responseCustomDetectionJSON contains the JSON metadata for the struct
// [ResponseCustomDetection]
type responseCustomDetectionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCustomDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCustomDetectionJSON) RawJSON() string {
	return r.raw
}

type ZoneLeakedCredentialCheckDetectionListResponse struct {
	Result []CustomDetection                                  `json:"result,nullable"`
	JSON   zoneLeakedCredentialCheckDetectionListResponseJSON `json:"-"`
	APIResponseWafProductBundle
}

// zoneLeakedCredentialCheckDetectionListResponseJSON contains the JSON metadata
// for the struct [ZoneLeakedCredentialCheckDetectionListResponse]
type zoneLeakedCredentialCheckDetectionListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLeakedCredentialCheckDetectionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneLeakedCredentialCheckDetectionListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneLeakedCredentialCheckDetectionNewParams struct {
	// A custom set of username/password expressions to match Leaked Credential Checks
	// on
	CustomDetection CustomDetectionParam `json:"custom_detection,required"`
}

func (r ZoneLeakedCredentialCheckDetectionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CustomDetection)
}

type ZoneLeakedCredentialCheckDetectionUpdateParams struct {
	// A custom set of username/password expressions to match Leaked Credential Checks
	// on
	CustomDetection CustomDetectionParam `json:"custom_detection,required"`
}

func (r ZoneLeakedCredentialCheckDetectionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CustomDetection)
}
