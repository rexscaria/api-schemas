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

// AccountMagicConnectorService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicConnectorService] method instead.
type AccountMagicConnectorService struct {
	Options   []option.RequestOption
	Telemetry *AccountMagicConnectorTelemetryService
}

// NewAccountMagicConnectorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMagicConnectorService(opts ...option.RequestOption) (r *AccountMagicConnectorService) {
	r = &AccountMagicConnectorService{}
	r.Options = opts
	r.Telemetry = NewAccountMagicConnectorTelemetryService(opts...)
	return
}

// Fetch Connector
func (r *AccountMagicConnectorService) Get(ctx context.Context, accountID string, connectorID string, opts ...option.RequestOption) (res *AccountMagicConnectorGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s", accountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Connectors
func (r *AccountMagicConnectorService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountMagicConnectorListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Connector
func (r *AccountMagicConnectorService) Patch(ctx context.Context, accountID string, connectorID string, body AccountMagicConnectorPatchParams, opts ...option.RequestOption) (res *MconnCustomerConnectorUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s", accountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Replace Connector
func (r *AccountMagicConnectorService) Replace(ctx context.Context, accountID string, connectorID string, body AccountMagicConnectorReplaceParams, opts ...option.RequestOption) (res *MconnCustomerConnectorUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s", accountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type MconnCustomerConnector struct {
	ID                           string                       `json:"id,required"`
	Activated                    bool                         `json:"activated,required"`
	InterruptWindowDurationHours float64                      `json:"interrupt_window_duration_hours,required"`
	InterruptWindowHourOfDay     float64                      `json:"interrupt_window_hour_of_day,required"`
	LastUpdated                  string                       `json:"last_updated,required"`
	Notes                        string                       `json:"notes,required"`
	Timezone                     string                       `json:"timezone,required"`
	Device                       MconnCustomerConnectorDevice `json:"device"`
	LastHeartbeat                string                       `json:"last_heartbeat"`
	LastSeenVersion              string                       `json:"last_seen_version"`
	JSON                         mconnCustomerConnectorJSON   `json:"-"`
}

// mconnCustomerConnectorJSON contains the JSON metadata for the struct
// [MconnCustomerConnector]
type mconnCustomerConnectorJSON struct {
	ID                           apijson.Field
	Activated                    apijson.Field
	InterruptWindowDurationHours apijson.Field
	InterruptWindowHourOfDay     apijson.Field
	LastUpdated                  apijson.Field
	Notes                        apijson.Field
	Timezone                     apijson.Field
	Device                       apijson.Field
	LastHeartbeat                apijson.Field
	LastSeenVersion              apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *MconnCustomerConnector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mconnCustomerConnectorJSON) RawJSON() string {
	return r.raw
}

type MconnCustomerConnectorDevice struct {
	ID           string                           `json:"id,required"`
	SerialNumber string                           `json:"serial_number"`
	JSON         mconnCustomerConnectorDeviceJSON `json:"-"`
}

// mconnCustomerConnectorDeviceJSON contains the JSON metadata for the struct
// [MconnCustomerConnectorDevice]
type mconnCustomerConnectorDeviceJSON struct {
	ID           apijson.Field
	SerialNumber apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MconnCustomerConnectorDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mconnCustomerConnectorDeviceJSON) RawJSON() string {
	return r.raw
}

type MconnCustomerConnectorFieldsParam struct {
	Activated                    param.Field[bool]    `json:"activated"`
	InterruptWindowDurationHours param.Field[float64] `json:"interrupt_window_duration_hours"`
	InterruptWindowHourOfDay     param.Field[float64] `json:"interrupt_window_hour_of_day"`
	Notes                        param.Field[string]  `json:"notes"`
	Timezone                     param.Field[string]  `json:"timezone"`
}

func (r MconnCustomerConnectorFieldsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MconnCustomerConnectorUpdateResponse struct {
	Result MconnCustomerConnector                   `json:"result"`
	JSON   mconnCustomerConnectorUpdateResponseJSON `json:"-"`
	MconnGoodResponse
}

// mconnCustomerConnectorUpdateResponseJSON contains the JSON metadata for the
// struct [MconnCustomerConnectorUpdateResponse]
type mconnCustomerConnectorUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MconnCustomerConnectorUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mconnCustomerConnectorUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type MconnGoodResponse struct {
	Errors   []MconnCodedMessage   `json:"errors,required"`
	Messages []MconnCodedMessage   `json:"messages,required"`
	Success  bool                  `json:"success,required"`
	JSON     mconnGoodResponseJSON `json:"-"`
}

// mconnGoodResponseJSON contains the JSON metadata for the struct
// [MconnGoodResponse]
type mconnGoodResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MconnGoodResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mconnGoodResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicConnectorGetResponse struct {
	Result MconnCustomerConnector               `json:"result"`
	JSON   accountMagicConnectorGetResponseJSON `json:"-"`
	MconnGoodResponse
}

// accountMagicConnectorGetResponseJSON contains the JSON metadata for the struct
// [AccountMagicConnectorGetResponse]
type accountMagicConnectorGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicConnectorListResponse struct {
	Result []MconnCustomerConnector              `json:"result"`
	JSON   accountMagicConnectorListResponseJSON `json:"-"`
	MconnGoodResponse
}

// accountMagicConnectorListResponseJSON contains the JSON metadata for the struct
// [AccountMagicConnectorListResponse]
type accountMagicConnectorListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicConnectorPatchParams struct {
	MconnCustomerConnectorFields MconnCustomerConnectorFieldsParam `json:"mconn_customer_connector_fields,required"`
}

func (r AccountMagicConnectorPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MconnCustomerConnectorFields)
}

type AccountMagicConnectorReplaceParams struct {
	MconnCustomerConnectorFields MconnCustomerConnectorFieldsParam `json:"mconn_customer_connector_fields,required"`
}

func (r AccountMagicConnectorReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MconnCustomerConnectorFields)
}
