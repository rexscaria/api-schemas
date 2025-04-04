// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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

// Fetches device certificate provisioning
func (r *ZoneDevicePolicyCertificateService) Get(ctx context.Context, zoneID interface{}, opts ...option.RequestOption) (res *DevicePolicyCertificate, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/devices/policy/certificates", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable Zero Trust Clients to provision a certificate, containing a x509 subject,
// and referenced by Access device posture policies when the client visits MTLS
// protected domains. This facilitates device posture without a WARP session.
func (r *ZoneDevicePolicyCertificateService) Update(ctx context.Context, zoneID interface{}, body ZoneDevicePolicyCertificateUpdateParams, opts ...option.RequestOption) (res *DevicePolicyCertificate, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/devices/policy/certificates", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type DevicePolicyCertificate struct {
	Result interface{}                 `json:"result,nullable"`
	JSON   devicePolicyCertificateJSON `json:"-"`
	APIResponseTeamsDevices
}

// devicePolicyCertificateJSON contains the JSON metadata for the struct
// [DevicePolicyCertificate]
type devicePolicyCertificateJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCertificateJSON) RawJSON() string {
	return r.raw
}

type ZoneDevicePolicyCertificateUpdateParams struct {
	// The current status of the device policy certificate provisioning feature for
	// WARP clients.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneDevicePolicyCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
