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

// ZoneLogpushOwnershipService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneLogpushOwnershipService] method instead.
type ZoneLogpushOwnershipService struct {
	Options []option.RequestOption
}

// NewZoneLogpushOwnershipService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneLogpushOwnershipService(opts ...option.RequestOption) (r *ZoneLogpushOwnershipService) {
	r = &ZoneLogpushOwnershipService{}
	r.Options = opts
	return
}

// Gets a new ownership challenge sent to your destination.
func (r *ZoneLogpushOwnershipService) GetChallenge(ctx context.Context, zoneID string, body ZoneLogpushOwnershipGetChallengeParams, opts ...option.RequestOption) (res *GetOwnershipResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/logpush/ownership", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Validates ownership challenge of the destination.
func (r *ZoneLogpushOwnershipService) ValidateChallenge(ctx context.Context, zoneID string, body ZoneLogpushOwnershipValidateChallengeParams, opts ...option.RequestOption) (res *ValidateOwnershipResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/logpush/ownership/validate", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneLogpushOwnershipGetChallengeParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
}

func (r ZoneLogpushOwnershipGetChallengeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneLogpushOwnershipValidateChallengeParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge,required"`
}

func (r ZoneLogpushOwnershipValidateChallengeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
