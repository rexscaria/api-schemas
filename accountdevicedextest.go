// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// AccountDeviceDexTestService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDeviceDexTestService] method instead.
type AccountDeviceDexTestService struct {
	Options []option.RequestOption
}

// NewAccountDeviceDexTestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDeviceDexTestService(opts ...option.RequestOption) (r *AccountDeviceDexTestService) {
	r = &AccountDeviceDexTestService{}
	r.Options = opts
	return
}

// Create a DEX test.
func (r *AccountDeviceDexTestService) New(ctx context.Context, accountID interface{}, body AccountDeviceDexTestNewParams, opts ...option.RequestOption) (res *DexSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/dex_tests", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a single DEX test.
func (r *AccountDeviceDexTestService) Get(ctx context.Context, accountID interface{}, dexTestID string, opts ...option.RequestOption) (res *DexSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if dexTestID == "" {
		err = errors.New("missing required dex_test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/devices/dex_tests/%s", accountID, dexTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a DEX test.
func (r *AccountDeviceDexTestService) Update(ctx context.Context, accountID interface{}, dexTestID string, body AccountDeviceDexTestUpdateParams, opts ...option.RequestOption) (res *DexSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if dexTestID == "" {
		err = errors.New("missing required dex_test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/devices/dex_tests/%s", accountID, dexTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetch all DEX tests.
func (r *AccountDeviceDexTestService) List(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *AccountDeviceDexTestListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/dex_tests", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a Device DEX test. Returns the remaining device dex tests for the
// account.
func (r *AccountDeviceDexTestService) Delete(ctx context.Context, accountID interface{}, dexTestID string, opts ...option.RequestOption) (res *AccountDeviceDexTestDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if dexTestID == "" {
		err = errors.New("missing required dex_test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/devices/dex_tests/%s", accountID, dexTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type APIResponseTeamsDevices struct {
	Errors   []MessagesDeviceTestsItems         `json:"errors,required"`
	Messages []MessagesDeviceTestsItems         `json:"messages,required"`
	Result   APIResponseTeamsDevicesResultUnion `json:"result,required"`
	// Whether the API call was successful.
	Success APIResponseTeamsDevicesSuccess `json:"success,required"`
	JSON    apiResponseTeamsDevicesJSON    `json:"-"`
}

// apiResponseTeamsDevicesJSON contains the JSON metadata for the struct
// [APIResponseTeamsDevices]
type apiResponseTeamsDevicesJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseTeamsDevices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseTeamsDevicesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [APIResponseTeamsDevicesResultArray] or [shared.UnionString].
type APIResponseTeamsDevicesResultUnion interface {
	ImplementsAPIResponseTeamsDevicesResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIResponseTeamsDevicesResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIResponseTeamsDevicesResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type APIResponseTeamsDevicesResultArray []interface{}

func (r APIResponseTeamsDevicesResultArray) ImplementsAPIResponseTeamsDevicesResultUnion() {}

// Whether the API call was successful.
type APIResponseTeamsDevicesSuccess bool

const (
	APIResponseTeamsDevicesSuccessTrue APIResponseTeamsDevicesSuccess = true
)

func (r APIResponseTeamsDevicesSuccess) IsKnown() bool {
	switch r {
	case APIResponseTeamsDevicesSuccessTrue:
		return true
	}
	return false
}

type DeviceDexTestHTTP struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data DeviceDexTestHTTPData `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled bool `json:"enabled,required"`
	// How often the test will run.
	Interval string `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name string `json:"name,required"`
	// Additional details about the test.
	Description string `json:"description"`
	// Device settings profiles targeted by this test
	TargetPolicies []DeviceDexTestHTTPTargetPolicy `json:"target_policies"`
	Targeted       bool                            `json:"targeted"`
	// The unique identifier for the test.
	TestID string                `json:"test_id"`
	JSON   deviceDexTestHTTPJSON `json:"-"`
}

// deviceDexTestHTTPJSON contains the JSON metadata for the struct
// [DeviceDexTestHTTP]
type deviceDexTestHTTPJSON struct {
	Data           apijson.Field
	Enabled        apijson.Field
	Interval       apijson.Field
	Name           apijson.Field
	Description    apijson.Field
	TargetPolicies apijson.Field
	Targeted       apijson.Field
	TestID         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DeviceDexTestHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDexTestHTTPJSON) RawJSON() string {
	return r.raw
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDexTestHTTPData struct {
	// The desired endpoint to test.
	Host string `json:"host"`
	// The type of test.
	Kind string `json:"kind"`
	// The HTTP request method type.
	Method string                    `json:"method"`
	JSON   deviceDexTestHTTPDataJSON `json:"-"`
}

// deviceDexTestHTTPDataJSON contains the JSON metadata for the struct
// [DeviceDexTestHTTPData]
type deviceDexTestHTTPDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDexTestHTTPData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDexTestHTTPDataJSON) RawJSON() string {
	return r.raw
}

type DeviceDexTestHTTPTargetPolicy struct {
	// The id of the device settings profile
	ID string `json:"id"`
	// Whether the profile is the account default
	Default bool `json:"default"`
	// The name of the device settings profile
	Name string                            `json:"name"`
	JSON deviceDexTestHTTPTargetPolicyJSON `json:"-"`
}

// deviceDexTestHTTPTargetPolicyJSON contains the JSON metadata for the struct
// [DeviceDexTestHTTPTargetPolicy]
type deviceDexTestHTTPTargetPolicyJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDexTestHTTPTargetPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDexTestHTTPTargetPolicyJSON) RawJSON() string {
	return r.raw
}

type DeviceDexTestHTTPParam struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data param.Field[DeviceDexTestHTTPDataParam] `json:"data,required"`
	// Determines whether or not the test is active.
	Enabled param.Field[bool] `json:"enabled,required"`
	// How often the test will run.
	Interval param.Field[string] `json:"interval,required"`
	// The name of the DEX test. Must be unique.
	Name param.Field[string] `json:"name,required"`
	// Additional details about the test.
	Description param.Field[string] `json:"description"`
	// Device settings profiles targeted by this test
	TargetPolicies param.Field[[]DeviceDexTestHTTPTargetPolicyParam] `json:"target_policies"`
	Targeted       param.Field[bool]                                 `json:"targeted"`
}

func (r DeviceDexTestHTTPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DeviceDexTestHTTPDataParam struct {
	// The desired endpoint to test.
	Host param.Field[string] `json:"host"`
	// The type of test.
	Kind param.Field[string] `json:"kind"`
	// The HTTP request method type.
	Method param.Field[string] `json:"method"`
}

func (r DeviceDexTestHTTPDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceDexTestHTTPTargetPolicyParam struct {
	// The id of the device settings profile
	ID param.Field[string] `json:"id"`
	// Whether the profile is the account default
	Default param.Field[bool] `json:"default"`
	// The name of the device settings profile
	Name param.Field[string] `json:"name"`
}

func (r DeviceDexTestHTTPTargetPolicyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DexSingleResponse struct {
	Result DeviceDexTestHTTP     `json:"result"`
	JSON   dexSingleResponseJSON `json:"-"`
	APIResponseSingleTeamsDevices
}

// dexSingleResponseJSON contains the JSON metadata for the struct
// [DexSingleResponse]
type dexSingleResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexSingleResponseJSON) RawJSON() string {
	return r.raw
}

type MessagesDeviceTestsItems struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    messagesDeviceTestsItemsJSON `json:"-"`
}

// messagesDeviceTestsItemsJSON contains the JSON metadata for the struct
// [MessagesDeviceTestsItems]
type messagesDeviceTestsItemsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesDeviceTestsItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesDeviceTestsItemsJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceDexTestListResponse struct {
	Result []DeviceDexTestHTTP                  `json:"result,nullable"`
	JSON   accountDeviceDexTestListResponseJSON `json:"-"`
	APIResponseTeamsDevices
}

// accountDeviceDexTestListResponseJSON contains the JSON metadata for the struct
// [AccountDeviceDexTestListResponse]
type accountDeviceDexTestListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceDexTestListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceDexTestDeleteResponse struct {
	Result AccountDeviceDexTestDeleteResponseResult `json:"result"`
	JSON   accountDeviceDexTestDeleteResponseJSON   `json:"-"`
	APIResponseTeamsDevices
}

// accountDeviceDexTestDeleteResponseJSON contains the JSON metadata for the struct
// [AccountDeviceDexTestDeleteResponse]
type accountDeviceDexTestDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceDexTestDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceDexTestDeleteResponseResult struct {
	DexTests []DeviceDexTestHTTP                          `json:"dex_tests"`
	JSON     accountDeviceDexTestDeleteResponseResultJSON `json:"-"`
}

// accountDeviceDexTestDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountDeviceDexTestDeleteResponseResult]
type accountDeviceDexTestDeleteResponseResultJSON struct {
	DexTests    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceDexTestDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeviceDexTestDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDeviceDexTestNewParams struct {
	DeviceDexTestHTTP DeviceDexTestHTTPParam `json:"device_dex_test_http,required"`
}

func (r AccountDeviceDexTestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DeviceDexTestHTTP)
}

type AccountDeviceDexTestUpdateParams struct {
	DeviceDexTestHTTP DeviceDexTestHTTPParam `json:"device_dex_test_http,required"`
}

func (r AccountDeviceDexTestUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DeviceDexTestHTTP)
}
