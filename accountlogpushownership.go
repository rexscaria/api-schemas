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

// AccountLogpushOwnershipService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountLogpushOwnershipService] method instead.
type AccountLogpushOwnershipService struct {
	Options []option.RequestOption
}

// NewAccountLogpushOwnershipService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountLogpushOwnershipService(opts ...option.RequestOption) (r *AccountLogpushOwnershipService) {
	r = &AccountLogpushOwnershipService{}
	r.Options = opts
	return
}

// Gets a new ownership challenge sent to your destination.
func (r *AccountLogpushOwnershipService) NewChallenge(ctx context.Context, accountID string, body AccountLogpushOwnershipNewChallengeParams, opts ...option.RequestOption) (res *GetOwnershipResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/logpush/ownership", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Validates ownership challenge of the destination.
func (r *AccountLogpushOwnershipService) ValidateChallenge(ctx context.Context, accountID string, body AccountLogpushOwnershipValidateChallengeParams, opts ...option.RequestOption) (res *ValidateOwnershipResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/logpush/ownership/validate", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type GetOwnershipResponse struct {
	Result GetOwnershipResponseResult `json:"result,nullable"`
	JSON   getOwnershipResponseJSON   `json:"-"`
	CommonResponseLogPush
}

// getOwnershipResponseJSON contains the JSON metadata for the struct
// [GetOwnershipResponse]
type getOwnershipResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetOwnershipResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r getOwnershipResponseJSON) RawJSON() string {
	return r.raw
}

type GetOwnershipResponseResult struct {
	Filename string                         `json:"filename"`
	Message  string                         `json:"message"`
	Valid    bool                           `json:"valid"`
	JSON     getOwnershipResponseResultJSON `json:"-"`
}

// getOwnershipResponseResultJSON contains the JSON metadata for the struct
// [GetOwnershipResponseResult]
type getOwnershipResponseResultJSON struct {
	Filename    apijson.Field
	Message     apijson.Field
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetOwnershipResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r getOwnershipResponseResultJSON) RawJSON() string {
	return r.raw
}

type ValidateOwnershipResponse struct {
	Result ValidateOwnershipResponseResult `json:"result,nullable"`
	JSON   validateOwnershipResponseJSON   `json:"-"`
	CommonResponseLogPush
}

// validateOwnershipResponseJSON contains the JSON metadata for the struct
// [ValidateOwnershipResponse]
type validateOwnershipResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateOwnershipResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateOwnershipResponseJSON) RawJSON() string {
	return r.raw
}

type ValidateOwnershipResponseResult struct {
	Valid bool                                `json:"valid"`
	JSON  validateOwnershipResponseResultJSON `json:"-"`
}

// validateOwnershipResponseResultJSON contains the JSON metadata for the struct
// [ValidateOwnershipResponseResult]
type validateOwnershipResponseResultJSON struct {
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateOwnershipResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateOwnershipResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountLogpushOwnershipNewChallengeParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
}

func (r AccountLogpushOwnershipNewChallengeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountLogpushOwnershipValidateChallengeParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge,required"`
}

func (r AccountLogpushOwnershipValidateChallengeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
