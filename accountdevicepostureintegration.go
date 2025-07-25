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

// AccountDevicePostureIntegrationService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDevicePostureIntegrationService] method instead.
type AccountDevicePostureIntegrationService struct {
	Options []option.RequestOption
}

// NewAccountDevicePostureIntegrationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDevicePostureIntegrationService(opts ...option.RequestOption) (r *AccountDevicePostureIntegrationService) {
	r = &AccountDevicePostureIntegrationService{}
	r.Options = opts
	return
}

// Create a new device posture integration.
func (r *AccountDevicePostureIntegrationService) New(ctx context.Context, accountID string, body AccountDevicePostureIntegrationNewParams, opts ...option.RequestOption) (res *SingleResponseIntegration, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/posture/integration", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches details for a single device posture integration.
func (r *AccountDevicePostureIntegrationService) Get(ctx context.Context, accountID string, integrationID string, opts ...option.RequestOption) (res *SingleResponseIntegration, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if integrationID == "" {
		err = errors.New("missing required integration_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/posture/integration/%s", accountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured device posture integration.
func (r *AccountDevicePostureIntegrationService) Update(ctx context.Context, accountID string, integrationID string, body AccountDevicePostureIntegrationUpdateParams, opts ...option.RequestOption) (res *SingleResponseIntegration, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if integrationID == "" {
		err = errors.New("missing required integration_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/posture/integration/%s", accountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Fetches the list of device posture integrations for an account.
func (r *AccountDevicePostureIntegrationService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountDevicePostureIntegrationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/posture/integration", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a configured device posture integration.
func (r *AccountDevicePostureIntegrationService) Delete(ctx context.Context, accountID string, integrationID string, opts ...option.RequestOption) (res *AccountDevicePostureIntegrationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if integrationID == "" {
		err = errors.New("missing required integration_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/posture/integration/%s", accountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// The configuration object containing third-party integration information.
type ConfigRequestIntegrationParam struct {
	// If present, this id will be passed in the `CF-Access-Client-ID` header when
	// hitting the `api_url`.
	AccessClientID param.Field[string] `json:"access_client_id"`
	// If present, this secret will be passed in the `CF-Access-Client-Secret` header
	// when hitting the `api_url`.
	AccessClientSecret param.Field[string] `json:"access_client_secret"`
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL param.Field[string] `json:"api_url"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL param.Field[string] `json:"auth_url"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID param.Field[string] `json:"client_id"`
	// The Uptycs client secret.
	ClientKey param.Field[string] `json:"client_key"`
	// The Workspace One client secret provided in the Workspace One Admin Dashboard.
	ClientSecret param.Field[string] `json:"client_secret"`
	// The Crowdstrike customer ID.
	CustomerID param.Field[string] `json:"customer_id"`
}

func (r ConfigRequestIntegrationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigRequestIntegrationParam) implementsConfigRequestIntegrationUnionParam() {}

// The configuration object containing third-party integration information.
//
// Satisfied by
// [ConfigRequestIntegrationTeamsDevicesWorkspaceOneConfigRequestParam],
// [ConfigRequestIntegrationTeamsDevicesCrowdstrikeConfigRequestParam],
// [ConfigRequestIntegrationTeamsDevicesUptycsConfigRequestParam],
// [ConfigRequestIntegrationTeamsDevicesIntuneConfigRequestParam],
// [ConfigRequestIntegrationTeamsDevicesKolideConfigRequestParam],
// [ConfigRequestIntegrationTeamsDevicesTaniumConfigRequestParam],
// [ConfigRequestIntegrationTeamsDevicesSentineloneS2sConfigRequestParam],
// [ConfigRequestIntegrationTeamsDevicesCustomS2sConfigRequestParam],
// [ConfigRequestIntegrationParam].
type ConfigRequestIntegrationUnionParam interface {
	implementsConfigRequestIntegrationUnionParam()
}

type ConfigRequestIntegrationTeamsDevicesWorkspaceOneConfigRequestParam struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL param.Field[string] `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Workspace One client secret provided in the Workspace One Admin Dashboard.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r ConfigRequestIntegrationTeamsDevicesWorkspaceOneConfigRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigRequestIntegrationTeamsDevicesWorkspaceOneConfigRequestParam) implementsConfigRequestIntegrationUnionParam() {
}

type ConfigRequestIntegrationTeamsDevicesCrowdstrikeConfigRequestParam struct {
	// The Crowdstrike API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Crowdstrike client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Crowdstrike client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Crowdstrike customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r ConfigRequestIntegrationTeamsDevicesCrowdstrikeConfigRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigRequestIntegrationTeamsDevicesCrowdstrikeConfigRequestParam) implementsConfigRequestIntegrationUnionParam() {
}

type ConfigRequestIntegrationTeamsDevicesUptycsConfigRequestParam struct {
	// The Uptycs API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Uptycs client secret.
	ClientKey param.Field[string] `json:"client_key,required"`
	// The Uptycs client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Uptycs customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r ConfigRequestIntegrationTeamsDevicesUptycsConfigRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigRequestIntegrationTeamsDevicesUptycsConfigRequestParam) implementsConfigRequestIntegrationUnionParam() {
}

type ConfigRequestIntegrationTeamsDevicesIntuneConfigRequestParam struct {
	// The Intune client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Intune client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The Intune customer ID.
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r ConfigRequestIntegrationTeamsDevicesIntuneConfigRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigRequestIntegrationTeamsDevicesIntuneConfigRequestParam) implementsConfigRequestIntegrationUnionParam() {
}

type ConfigRequestIntegrationTeamsDevicesKolideConfigRequestParam struct {
	// The Kolide client ID.
	ClientID param.Field[string] `json:"client_id,required"`
	// The Kolide client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r ConfigRequestIntegrationTeamsDevicesKolideConfigRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigRequestIntegrationTeamsDevicesKolideConfigRequestParam) implementsConfigRequestIntegrationUnionParam() {
}

type ConfigRequestIntegrationTeamsDevicesTaniumConfigRequestParam struct {
	// The Tanium API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The Tanium client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// If present, this id will be passed in the `CF-Access-Client-ID` header when
	// hitting the `api_url`.
	AccessClientID param.Field[string] `json:"access_client_id"`
	// If present, this secret will be passed in the `CF-Access-Client-Secret` header
	// when hitting the `api_url`.
	AccessClientSecret param.Field[string] `json:"access_client_secret"`
}

func (r ConfigRequestIntegrationTeamsDevicesTaniumConfigRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigRequestIntegrationTeamsDevicesTaniumConfigRequestParam) implementsConfigRequestIntegrationUnionParam() {
}

type ConfigRequestIntegrationTeamsDevicesSentineloneS2sConfigRequestParam struct {
	// The SentinelOne S2S API URL.
	APIURL param.Field[string] `json:"api_url,required"`
	// The SentinelOne S2S client secret.
	ClientSecret param.Field[string] `json:"client_secret,required"`
}

func (r ConfigRequestIntegrationTeamsDevicesSentineloneS2sConfigRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigRequestIntegrationTeamsDevicesSentineloneS2sConfigRequestParam) implementsConfigRequestIntegrationUnionParam() {
}

type ConfigRequestIntegrationTeamsDevicesCustomS2sConfigRequestParam struct {
	// This id will be passed in the `CF-Access-Client-ID` header when hitting the
	// `api_url`.
	AccessClientID param.Field[string] `json:"access_client_id,required"`
	// This secret will be passed in the `CF-Access-Client-Secret` header when hitting
	// the `api_url`.
	AccessClientSecret param.Field[string] `json:"access_client_secret,required"`
	// The Custom Device Posture Integration API URL.
	APIURL param.Field[string] `json:"api_url,required"`
}

func (r ConfigRequestIntegrationTeamsDevicesCustomS2sConfigRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigRequestIntegrationTeamsDevicesCustomS2sConfigRequestParam) implementsConfigRequestIntegrationUnionParam() {
}

type DevicePostureIntegrations struct {
	// API UUID.
	ID string `json:"id"`
	// The configuration object containing third-party integration information.
	Config DevicePostureIntegrationsConfig `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval string `json:"interval"`
	// The name of the device posture integration.
	Name string `json:"name"`
	// The type of device posture integration.
	Type DeviceTypePostureIntegration  `json:"type"`
	JSON devicePostureIntegrationsJSON `json:"-"`
}

// devicePostureIntegrationsJSON contains the JSON metadata for the struct
// [DevicePostureIntegrations]
type devicePostureIntegrationsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Interval    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureIntegrationsJSON) RawJSON() string {
	return r.raw
}

// The configuration object containing third-party integration information.
type DevicePostureIntegrationsConfig struct {
	// The Workspace One API URL provided in the Workspace One Admin Dashboard.
	APIURL string `json:"api_url,required"`
	// The Workspace One Authorization URL depending on your region.
	AuthURL string `json:"auth_url,required"`
	// The Workspace One client ID provided in the Workspace One Admin Dashboard.
	ClientID string                              `json:"client_id,required"`
	JSON     devicePostureIntegrationsConfigJSON `json:"-"`
}

// devicePostureIntegrationsConfigJSON contains the JSON metadata for the struct
// [DevicePostureIntegrationsConfig]
type devicePostureIntegrationsConfigJSON struct {
	APIURL      apijson.Field
	AuthURL     apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureIntegrationsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureIntegrationsConfigJSON) RawJSON() string {
	return r.raw
}

// The type of device posture integration.
type DeviceTypePostureIntegration string

const (
	DeviceTypePostureIntegrationWorkspaceOne   DeviceTypePostureIntegration = "workspace_one"
	DeviceTypePostureIntegrationCrowdstrikeS2s DeviceTypePostureIntegration = "crowdstrike_s2s"
	DeviceTypePostureIntegrationUptycs         DeviceTypePostureIntegration = "uptycs"
	DeviceTypePostureIntegrationIntune         DeviceTypePostureIntegration = "intune"
	DeviceTypePostureIntegrationKolide         DeviceTypePostureIntegration = "kolide"
	DeviceTypePostureIntegrationTaniumS2s      DeviceTypePostureIntegration = "tanium_s2s"
	DeviceTypePostureIntegrationSentineloneS2s DeviceTypePostureIntegration = "sentinelone_s2s"
	DeviceTypePostureIntegrationCustomS2s      DeviceTypePostureIntegration = "custom_s2s"
)

func (r DeviceTypePostureIntegration) IsKnown() bool {
	switch r {
	case DeviceTypePostureIntegrationWorkspaceOne, DeviceTypePostureIntegrationCrowdstrikeS2s, DeviceTypePostureIntegrationUptycs, DeviceTypePostureIntegrationIntune, DeviceTypePostureIntegrationKolide, DeviceTypePostureIntegrationTaniumS2s, DeviceTypePostureIntegrationSentineloneS2s, DeviceTypePostureIntegrationCustomS2s:
		return true
	}
	return false
}

type SingleResponseIntegration struct {
	Errors   []MessagesDeviceTestsItems `json:"errors,required"`
	Messages []MessagesDeviceTestsItems `json:"messages,required"`
	Result   DevicePostureIntegrations  `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success SingleResponseIntegrationSuccess `json:"success,required"`
	JSON    singleResponseIntegrationJSON    `json:"-"`
}

// singleResponseIntegrationJSON contains the JSON metadata for the struct
// [SingleResponseIntegration]
type singleResponseIntegrationJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SingleResponseIntegrationSuccess bool

const (
	SingleResponseIntegrationSuccessTrue SingleResponseIntegrationSuccess = true
)

func (r SingleResponseIntegrationSuccess) IsKnown() bool {
	switch r {
	case SingleResponseIntegrationSuccessTrue:
		return true
	}
	return false
}

type AccountDevicePostureIntegrationListResponse struct {
	Errors   []MessagesDeviceTestsItems  `json:"errors,required"`
	Messages []MessagesDeviceTestsItems  `json:"messages,required"`
	Result   []DevicePostureIntegrations `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    AccountDevicePostureIntegrationListResponseSuccess    `json:"success,required"`
	ResultInfo AccountDevicePostureIntegrationListResponseResultInfo `json:"result_info"`
	JSON       accountDevicePostureIntegrationListResponseJSON       `json:"-"`
}

// accountDevicePostureIntegrationListResponseJSON contains the JSON metadata for
// the struct [AccountDevicePostureIntegrationListResponse]
type accountDevicePostureIntegrationListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDevicePostureIntegrationListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDevicePostureIntegrationListResponseSuccess bool

const (
	AccountDevicePostureIntegrationListResponseSuccessTrue AccountDevicePostureIntegrationListResponseSuccess = true
)

func (r AccountDevicePostureIntegrationListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDevicePostureIntegrationListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDevicePostureIntegrationListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                                   `json:"total_count"`
	JSON       accountDevicePostureIntegrationListResponseResultInfoJSON `json:"-"`
}

// accountDevicePostureIntegrationListResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountDevicePostureIntegrationListResponseResultInfo]
type accountDevicePostureIntegrationListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDevicePostureIntegrationListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountDevicePostureIntegrationDeleteResponse struct {
	Errors   []MessagesDeviceTestsItems `json:"errors,required"`
	Messages []MessagesDeviceTestsItems `json:"messages,required"`
	Result   interface{}                `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success AccountDevicePostureIntegrationDeleteResponseSuccess `json:"success,required"`
	JSON    accountDevicePostureIntegrationDeleteResponseJSON    `json:"-"`
}

// accountDevicePostureIntegrationDeleteResponseJSON contains the JSON metadata for
// the struct [AccountDevicePostureIntegrationDeleteResponse]
type accountDevicePostureIntegrationDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureIntegrationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDevicePostureIntegrationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDevicePostureIntegrationDeleteResponseSuccess bool

const (
	AccountDevicePostureIntegrationDeleteResponseSuccessTrue AccountDevicePostureIntegrationDeleteResponseSuccess = true
)

func (r AccountDevicePostureIntegrationDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDevicePostureIntegrationDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDevicePostureIntegrationNewParams struct {
	// The configuration object containing third-party integration information.
	Config param.Field[ConfigRequestIntegrationUnionParam] `json:"config,required"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval param.Field[string] `json:"interval,required"`
	// The name of the device posture integration.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture integration.
	Type param.Field[DeviceTypePostureIntegration] `json:"type,required"`
}

func (r AccountDevicePostureIntegrationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePostureIntegrationUpdateParams struct {
	// The configuration object containing third-party integration information.
	Config param.Field[ConfigRequestIntegrationUnionParam] `json:"config"`
	// The interval between each posture check with the third-party API. Use `m` for
	// minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
	Interval param.Field[string] `json:"interval"`
	// The name of the device posture integration.
	Name param.Field[string] `json:"name"`
	// The type of device posture integration.
	Type param.Field[DeviceTypePostureIntegration] `json:"type"`
}

func (r AccountDevicePostureIntegrationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
