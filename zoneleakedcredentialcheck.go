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

// ZoneLeakedCredentialCheckService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneLeakedCredentialCheckService] method instead.
type ZoneLeakedCredentialCheckService struct {
	Options    []option.RequestOption
	Detections *ZoneLeakedCredentialCheckDetectionService
}

// NewZoneLeakedCredentialCheckService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneLeakedCredentialCheckService(opts ...option.RequestOption) (r *ZoneLeakedCredentialCheckService) {
	r = &ZoneLeakedCredentialCheckService{}
	r.Options = opts
	r.Detections = NewZoneLeakedCredentialCheckDetectionService(opts...)
	return
}

// Retrieves the current status of Leaked Credential Checks.
func (r *ZoneLeakedCredentialCheckService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ResponseStatus, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/leaked-credential-checks", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the current status of Leaked Credential Checks.
func (r *ZoneLeakedCredentialCheckService) Update(ctx context.Context, zoneID string, body ZoneLeakedCredentialCheckUpdateParams, opts ...option.RequestOption) (res *ResponseStatus, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/leaked-credential-checks", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ResponseStatus struct {
	Errors   []WafProductAPIBundleMessages `json:"errors,required"`
	Messages []WafProductAPIBundleMessages `json:"messages,required"`
	// Defines the overall status for Leaked Credential Checks.
	Result StatusLeakedCredentialChecks `json:"result,required"`
	// Defines whether the API call was successful.
	Success ResponseStatusSuccess `json:"success,required"`
	JSON    responseStatusJSON    `json:"-"`
}

// responseStatusJSON contains the JSON metadata for the struct [ResponseStatus]
type responseStatusJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseStatusJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type ResponseStatusSuccess bool

const (
	ResponseStatusSuccessTrue ResponseStatusSuccess = true
)

func (r ResponseStatusSuccess) IsKnown() bool {
	switch r {
	case ResponseStatusSuccessTrue:
		return true
	}
	return false
}

// Defines the overall status for Leaked Credential Checks.
type StatusLeakedCredentialChecks struct {
	// Determines whether or not Leaked Credential Checks are enabled.
	Enabled bool                             `json:"enabled"`
	JSON    statusLeakedCredentialChecksJSON `json:"-"`
}

// statusLeakedCredentialChecksJSON contains the JSON metadata for the struct
// [StatusLeakedCredentialChecks]
type statusLeakedCredentialChecksJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StatusLeakedCredentialChecks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statusLeakedCredentialChecksJSON) RawJSON() string {
	return r.raw
}

// Defines the overall status for Leaked Credential Checks.
type StatusLeakedCredentialChecksParam struct {
	// Determines whether or not Leaked Credential Checks are enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r StatusLeakedCredentialChecksParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneLeakedCredentialCheckUpdateParams struct {
	// Defines the overall status for Leaked Credential Checks.
	StatusLeakedCredentialChecks StatusLeakedCredentialChecksParam `json:"status_leaked_credential_checks,required"`
}

func (r ZoneLeakedCredentialCheckUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.StatusLeakedCredentialChecks)
}
