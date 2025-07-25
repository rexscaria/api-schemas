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

// ZoneDevicePolicyCertificateService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneDevicePolicyCertificateService] method instead.
type ZoneDevicePolicyCertificateService struct {
	Options []option.RequestOption
}

// NewZoneDevicePolicyCertificateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneDevicePolicyCertificateService(opts ...option.RequestOption) (r *ZoneDevicePolicyCertificateService) {
	r = &ZoneDevicePolicyCertificateService{}
	r.Options = opts
	return
}

// Fetches device certificate provisioning.
func (r *ZoneDevicePolicyCertificateService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *DevicePolicyCertificate, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/devices/policy/certificates", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable Zero Trust Clients to provision a certificate, containing a x509 subject,
// and referenced by Access device posture policies when the client visits MTLS
// protected domains. This facilitates device posture without a WARP session.
func (r *ZoneDevicePolicyCertificateService) Update(ctx context.Context, zoneID string, body ZoneDevicePolicyCertificateUpdateParams, opts ...option.RequestOption) (res *DevicePolicyCertificate, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/devices/policy/certificates", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type DevicePolicyCertificate struct {
	Errors   []DevicePolicyCertificateError   `json:"errors,required"`
	Messages []DevicePolicyCertificateMessage `json:"messages,required"`
	Result   DevicePolicyCertificateResult    `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePolicyCertificateSuccess `json:"success,required"`
	JSON    devicePolicyCertificateJSON    `json:"-"`
}

// devicePolicyCertificateJSON contains the JSON metadata for the struct
// [DevicePolicyCertificate]
type devicePolicyCertificateJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCertificateJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyCertificateError struct {
	Code             int64                               `json:"code,required"`
	Message          string                              `json:"message,required"`
	DocumentationURL string                              `json:"documentation_url"`
	Source           DevicePolicyCertificateErrorsSource `json:"source"`
	JSON             devicePolicyCertificateErrorJSON    `json:"-"`
}

// devicePolicyCertificateErrorJSON contains the JSON metadata for the struct
// [DevicePolicyCertificateError]
type devicePolicyCertificateErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePolicyCertificateError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCertificateErrorJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyCertificateErrorsSource struct {
	Pointer string                                  `json:"pointer"`
	JSON    devicePolicyCertificateErrorsSourceJSON `json:"-"`
}

// devicePolicyCertificateErrorsSourceJSON contains the JSON metadata for the
// struct [DevicePolicyCertificateErrorsSource]
type devicePolicyCertificateErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCertificateErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCertificateErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyCertificateMessage struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           DevicePolicyCertificateMessagesSource `json:"source"`
	JSON             devicePolicyCertificateMessageJSON    `json:"-"`
}

// devicePolicyCertificateMessageJSON contains the JSON metadata for the struct
// [DevicePolicyCertificateMessage]
type devicePolicyCertificateMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DevicePolicyCertificateMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCertificateMessageJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyCertificateMessagesSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    devicePolicyCertificateMessagesSourceJSON `json:"-"`
}

// devicePolicyCertificateMessagesSourceJSON contains the JSON metadata for the
// struct [DevicePolicyCertificateMessagesSource]
type devicePolicyCertificateMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCertificateMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCertificateMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyCertificateResult struct {
	// The current status of the device policy certificate provisioning feature for
	// WARP clients.
	Enabled bool                              `json:"enabled,required"`
	JSON    devicePolicyCertificateResultJSON `json:"-"`
}

// devicePolicyCertificateResultJSON contains the JSON metadata for the struct
// [DevicePolicyCertificateResult]
type devicePolicyCertificateResultJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCertificateResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCertificateResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyCertificateSuccess bool

const (
	DevicePolicyCertificateSuccessTrue DevicePolicyCertificateSuccess = true
)

func (r DevicePolicyCertificateSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyCertificateSuccessTrue:
		return true
	}
	return false
}

type ZoneDevicePolicyCertificateUpdateParams struct {
	// The current status of the device policy certificate provisioning feature for
	// WARP clients.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneDevicePolicyCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
