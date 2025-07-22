// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountCloudforceOneScanConfigService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCloudforceOneScanConfigService] method instead.
type AccountCloudforceOneScanConfigService struct {
	Options []option.RequestOption
}

// NewAccountCloudforceOneScanConfigService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCloudforceOneScanConfigService(opts ...option.RequestOption) (r *AccountCloudforceOneScanConfigService) {
	r = &AccountCloudforceOneScanConfigService{}
	r.Options = opts
	return
}

// Create a new Scan Config
func (r *AccountCloudforceOneScanConfigService) New(ctx context.Context, accountID string, body AccountCloudforceOneScanConfigNewParams, opts ...option.RequestOption) (res *AccountCloudforceOneScanConfigNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/scans/config", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an existing Scan Config
func (r *AccountCloudforceOneScanConfigService) Update(ctx context.Context, accountID string, configID string, body AccountCloudforceOneScanConfigUpdateParams, opts ...option.RequestOption) (res *AccountCloudforceOneScanConfigUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/scans/config/%s", accountID, configID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Scan Configs
func (r *AccountCloudforceOneScanConfigService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountCloudforceOneScanConfigListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/scans/config", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a Scan Config
func (r *AccountCloudforceOneScanConfigService) Delete(ctx context.Context, accountID string, configID string, opts ...option.RequestOption) (res *AccountCloudforceOneScanConfigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/scans/config/%s", accountID, configID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type APIResponsePortScan struct {
	Errors   []MessagesPortScanItems `json:"errors,required"`
	Messages []MessagesPortScanItems `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponsePortScanSuccess `json:"success,required"`
	JSON    apiResponsePortScanJSON    `json:"-"`
}

// apiResponsePortScanJSON contains the JSON metadata for the struct
// [APIResponsePortScan]
type apiResponsePortScanJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponsePortScan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponsePortScanJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponsePortScanSuccess bool

const (
	APIResponsePortScanSuccessTrue APIResponsePortScanSuccess = true
)

func (r APIResponsePortScanSuccess) IsKnown() bool {
	switch r {
	case APIResponsePortScanSuccessTrue:
		return true
	}
	return false
}

type MessagesPortScanItems struct {
	Code    int64                     `json:"code,required"`
	Message string                    `json:"message,required"`
	JSON    messagesPortScanItemsJSON `json:"-"`
}

// messagesPortScanItemsJSON contains the JSON metadata for the struct
// [MessagesPortScanItems]
type messagesPortScanItemsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesPortScanItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesPortScanItemsJSON) RawJSON() string {
	return r.raw
}

type ScanConfig struct {
	// Config ID
	ID        string `json:"id,required"`
	AccountID string `json:"account_id,required"`
	// The number of days between each scan (0 = no recurring scans).
	Frequency float64 `json:"frequency,required"`
	// A list of IP addresses or CIDR blocks to scan. The maximum number of total IP
	// addresses allowed is 5000.
	IPs []string `json:"ips,required"`
	// A list of ports to scan. Allowed values:"default", "all", or a comma-separated
	// list of ports or range of ports (e.g. ["1-80", "443"]). Default will scan the
	// 100 most commonly open ports.
	Ports []string       `json:"ports,required"`
	JSON  scanConfigJSON `json:"-"`
}

// scanConfigJSON contains the JSON metadata for the struct [ScanConfig]
type scanConfigJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Frequency   apijson.Field
	IPs         apijson.Field
	Ports       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneScanConfigNewResponse struct {
	Result ScanConfig                                    `json:"result"`
	JSON   accountCloudforceOneScanConfigNewResponseJSON `json:"-"`
	APIResponsePortScan
}

// accountCloudforceOneScanConfigNewResponseJSON contains the JSON metadata for the
// struct [AccountCloudforceOneScanConfigNewResponse]
type accountCloudforceOneScanConfigNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneScanConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneScanConfigNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneScanConfigUpdateResponse struct {
	Result ScanConfig                                       `json:"result"`
	JSON   accountCloudforceOneScanConfigUpdateResponseJSON `json:"-"`
	APIResponsePortScan
}

// accountCloudforceOneScanConfigUpdateResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneScanConfigUpdateResponse]
type accountCloudforceOneScanConfigUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneScanConfigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneScanConfigUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneScanConfigListResponse struct {
	Result []ScanConfig                                   `json:"result"`
	JSON   accountCloudforceOneScanConfigListResponseJSON `json:"-"`
	APIResponsePortScan
}

// accountCloudforceOneScanConfigListResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneScanConfigListResponse]
type accountCloudforceOneScanConfigListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneScanConfigListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneScanConfigListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneScanConfigDeleteResponse struct {
	Errors   []string                                         `json:"errors,required"`
	Messages []string                                         `json:"messages,required"`
	Result   interface{}                                      `json:"result,required"`
	Success  bool                                             `json:"success,required"`
	JSON     accountCloudforceOneScanConfigDeleteResponseJSON `json:"-"`
}

// accountCloudforceOneScanConfigDeleteResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneScanConfigDeleteResponse]
type accountCloudforceOneScanConfigDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneScanConfigDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneScanConfigDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneScanConfigNewParams struct {
	// A list of IP addresses or CIDR blocks to scan. The maximum number of total IP
	// addresses allowed is 5000.
	IPs param.Field[[]string] `json:"ips,required"`
	// The number of days between each scan (0 = no recurring scans).
	Frequency param.Field[float64] `json:"frequency"`
	// A list of ports to scan. Allowed values:"default", "all", or a comma-separated
	// list of ports or range of ports (e.g. ["1-80", "443"]). Default will scan the
	// 100 most commonly open ports.
	Ports param.Field[[]string] `json:"ports"`
}

func (r AccountCloudforceOneScanConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCloudforceOneScanConfigUpdateParams struct {
	// The number of days between each scan (0 = no recurring scans).
	Frequency param.Field[float64] `json:"frequency"`
	// A list of IP addresses or CIDR blocks to scan. The maximum number of total IP
	// addresses allowed is 5000.
	IPs param.Field[[]string] `json:"ips"`
	// A list of ports to scan. Allowed values:"default", "all", or a comma-separated
	// list of ports or range of ports (e.g. ["1-80", "443"]). Default will scan the
	// 100 most commonly open ports.
	Ports param.Field[[]string] `json:"ports"`
}

func (r AccountCloudforceOneScanConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
