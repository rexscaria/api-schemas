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

// AccountLogpushValidateDestinationService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountLogpushValidateDestinationService] method instead.
type AccountLogpushValidateDestinationService struct {
	Options []option.RequestOption
}

// NewAccountLogpushValidateDestinationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLogpushValidateDestinationService(opts ...option.RequestOption) (r *AccountLogpushValidateDestinationService) {
	r = &AccountLogpushValidateDestinationService{}
	r.Options = opts
	return
}

// Checks if there is an existing job with a destination.
func (r *AccountLogpushValidateDestinationService) CheckExists(ctx context.Context, accountID string, body AccountLogpushValidateDestinationCheckExistsParams, opts ...option.RequestOption) (res *DestinationExistsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/logpush/validate/destination/exists", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Validates destination.
func (r *AccountLogpushValidateDestinationService) Validate(ctx context.Context, accountID string, body AccountLogpushValidateDestinationValidateParams, opts ...option.RequestOption) (res *ValidateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/logpush/validate/destination", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DestinationExistsResponse struct {
	Result DestinationExistsResponseResult `json:"result,nullable"`
	JSON   destinationExistsResponseJSON   `json:"-"`
	CommonResponseLogPush
}

// destinationExistsResponseJSON contains the JSON metadata for the struct
// [DestinationExistsResponse]
type destinationExistsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationExistsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationExistsResponseJSON) RawJSON() string {
	return r.raw
}

type DestinationExistsResponseResult struct {
	Exists bool                                `json:"exists"`
	JSON   destinationExistsResponseResultJSON `json:"-"`
}

// destinationExistsResponseResultJSON contains the JSON metadata for the struct
// [DestinationExistsResponseResult]
type destinationExistsResponseResultJSON struct {
	Exists      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationExistsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationExistsResponseResultJSON) RawJSON() string {
	return r.raw
}

type ValidateResponse struct {
	Result ValidateResponseResult `json:"result,nullable"`
	JSON   validateResponseJSON   `json:"-"`
	CommonResponseLogPush
}

// validateResponseJSON contains the JSON metadata for the struct
// [ValidateResponse]
type validateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateResponseJSON) RawJSON() string {
	return r.raw
}

type ValidateResponseResult struct {
	Message string                     `json:"message"`
	Valid   bool                       `json:"valid"`
	JSON    validateResponseResultJSON `json:"-"`
}

// validateResponseResultJSON contains the JSON metadata for the struct
// [ValidateResponseResult]
type validateResponseResultJSON struct {
	Message     apijson.Field
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountLogpushValidateDestinationCheckExistsParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
}

func (r AccountLogpushValidateDestinationCheckExistsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountLogpushValidateDestinationValidateParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
}

func (r AccountLogpushValidateDestinationValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
