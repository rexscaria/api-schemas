// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneCustomHostnameFallbackOriginService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCustomHostnameFallbackOriginService] method instead.
type ZoneCustomHostnameFallbackOriginService struct {
	Options []option.RequestOption
}

// NewZoneCustomHostnameFallbackOriginService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneCustomHostnameFallbackOriginService(opts ...option.RequestOption) (r *ZoneCustomHostnameFallbackOriginService) {
	r = &ZoneCustomHostnameFallbackOriginService{}
	r.Options = opts
	return
}

// Get Fallback Origin for Custom Hostnames
func (r *ZoneCustomHostnameFallbackOriginService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *FallbackOriginResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Fallback Origin for Custom Hostnames
func (r *ZoneCustomHostnameFallbackOriginService) Update(ctx context.Context, zoneID string, body ZoneCustomHostnameFallbackOriginUpdateParams, opts ...option.RequestOption) (res *FallbackOriginResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete Fallback Origin for Custom Hostnames
func (r *ZoneCustomHostnameFallbackOriginService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *FallbackOriginResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type FallbackOriginResponse struct {
	Errors   []MessagesTlsCertificatesItem `json:"errors,required"`
	Messages []MessagesTlsCertificatesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success FallbackOriginResponseSuccess `json:"success,required"`
	Result  FallbackOriginResponseResult  `json:"result"`
	JSON    fallbackOriginResponseJSON    `json:"-"`
}

// fallbackOriginResponseJSON contains the JSON metadata for the struct
// [FallbackOriginResponse]
type fallbackOriginResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type FallbackOriginResponseSuccess bool

const (
	FallbackOriginResponseSuccessTrue FallbackOriginResponseSuccess = true
)

func (r FallbackOriginResponseSuccess) IsKnown() bool {
	switch r {
	case FallbackOriginResponseSuccessTrue:
		return true
	}
	return false
}

type FallbackOriginResponseResult struct {
	// This is the time the fallback origin was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// These are errors that were encountered while trying to activate a fallback
	// origin.
	Errors []string `json:"errors"`
	// Your origin hostname that requests to your custom hostnames will be sent to.
	Origin string `json:"origin"`
	// Status of the fallback origin's activation.
	Status FallbackOriginResponseResultStatus `json:"status"`
	// This is the time the fallback origin was updated.
	UpdatedAt time.Time                        `json:"updated_at" format:"date-time"`
	JSON      fallbackOriginResponseResultJSON `json:"-"`
}

// fallbackOriginResponseResultJSON contains the JSON metadata for the struct
// [FallbackOriginResponseResult]
type fallbackOriginResponseResultJSON struct {
	CreatedAt   apijson.Field
	Errors      apijson.Field
	Origin      apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginResponseResultJSON) RawJSON() string {
	return r.raw
}

// Status of the fallback origin's activation.
type FallbackOriginResponseResultStatus string

const (
	FallbackOriginResponseResultStatusInitializing       FallbackOriginResponseResultStatus = "initializing"
	FallbackOriginResponseResultStatusPendingDeployment  FallbackOriginResponseResultStatus = "pending_deployment"
	FallbackOriginResponseResultStatusPendingDeletion    FallbackOriginResponseResultStatus = "pending_deletion"
	FallbackOriginResponseResultStatusActive             FallbackOriginResponseResultStatus = "active"
	FallbackOriginResponseResultStatusDeploymentTimedOut FallbackOriginResponseResultStatus = "deployment_timed_out"
	FallbackOriginResponseResultStatusDeletionTimedOut   FallbackOriginResponseResultStatus = "deletion_timed_out"
)

func (r FallbackOriginResponseResultStatus) IsKnown() bool {
	switch r {
	case FallbackOriginResponseResultStatusInitializing, FallbackOriginResponseResultStatusPendingDeployment, FallbackOriginResponseResultStatusPendingDeletion, FallbackOriginResponseResultStatusActive, FallbackOriginResponseResultStatusDeploymentTimedOut, FallbackOriginResponseResultStatusDeletionTimedOut:
		return true
	}
	return false
}

type ZoneCustomHostnameFallbackOriginUpdateParams struct {
	// Your origin hostname that requests to your custom hostnames will be sent to.
	Origin param.Field[string] `json:"origin,required"`
}

func (r ZoneCustomHostnameFallbackOriginUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
