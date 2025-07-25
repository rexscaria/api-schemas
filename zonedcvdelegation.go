// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneDcvDelegationService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneDcvDelegationService] method instead.
type ZoneDcvDelegationService struct {
	Options []option.RequestOption
}

// NewZoneDcvDelegationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneDcvDelegationService(opts ...option.RequestOption) (r *ZoneDcvDelegationService) {
	r = &ZoneDcvDelegationService{}
	r.Options = opts
	return
}

// Retrieve the account and zone specific unique identifier used as part of the
// CNAME target for DCV Delegation.
func (r *ZoneDcvDelegationService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneDcvDelegationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dcv_delegation/uuid", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneDcvDelegationGetResponse struct {
	Errors   []MessagesTlsCertificatesItem `json:"errors,required"`
	Messages []MessagesTlsCertificatesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneDcvDelegationGetResponseSuccess `json:"success,required"`
	Result  ZoneDcvDelegationGetResponseResult  `json:"result"`
	JSON    zoneDcvDelegationGetResponseJSON    `json:"-"`
}

// zoneDcvDelegationGetResponseJSON contains the JSON metadata for the struct
// [ZoneDcvDelegationGetResponse]
type zoneDcvDelegationGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDcvDelegationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDcvDelegationGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneDcvDelegationGetResponseSuccess bool

const (
	ZoneDcvDelegationGetResponseSuccessTrue ZoneDcvDelegationGetResponseSuccess = true
)

func (r ZoneDcvDelegationGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneDcvDelegationGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneDcvDelegationGetResponseResult struct {
	// The DCV Delegation unique identifier.
	Uuid string                                 `json:"uuid"`
	JSON zoneDcvDelegationGetResponseResultJSON `json:"-"`
}

// zoneDcvDelegationGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneDcvDelegationGetResponseResult]
type zoneDcvDelegationGetResponseResultJSON struct {
	Uuid        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDcvDelegationGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDcvDelegationGetResponseResultJSON) RawJSON() string {
	return r.raw
}
