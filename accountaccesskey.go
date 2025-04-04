// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountAccessKeyService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessKeyService] method instead.
type AccountAccessKeyService struct {
	Options []option.RequestOption
}

// NewAccountAccessKeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessKeyService(opts ...option.RequestOption) (r *AccountAccessKeyService) {
	r = &AccountAccessKeyService{}
	r.Options = opts
	return
}

// Gets the Access key rotation settings for an account.
func (r *AccountAccessKeyService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *SingleResponseKey, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/keys", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Access key rotation settings for an account.
func (r *AccountAccessKeyService) Update(ctx context.Context, accountID string, body AccountAccessKeyUpdateParams, opts ...option.RequestOption) (res *SingleResponseKey, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/keys", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Perfoms a key rotation for an account.
func (r *AccountAccessKeyService) Rotate(ctx context.Context, accountID string, opts ...option.RequestOption) (res *SingleResponseKey, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/keys/rotate", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type SingleResponseKey struct {
	Result SingleResponseKeyResult `json:"result"`
	JSON   singleResponseKeyJSON   `json:"-"`
	APIResponseSingleAccess
}

// singleResponseKeyJSON contains the JSON metadata for the struct
// [SingleResponseKey]
type singleResponseKeyJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseKeyJSON) RawJSON() string {
	return r.raw
}

type SingleResponseKeyResult struct {
	// The number of days until the next key rotation.
	DaysUntilNextRotation float64 `json:"days_until_next_rotation"`
	// The number of days between key rotations.
	KeyRotationIntervalDays float64 `json:"key_rotation_interval_days"`
	// The timestamp of the previous key rotation.
	LastKeyRotationAt time.Time                   `json:"last_key_rotation_at" format:"date-time"`
	JSON              singleResponseKeyResultJSON `json:"-"`
}

// singleResponseKeyResultJSON contains the JSON metadata for the struct
// [SingleResponseKeyResult]
type singleResponseKeyResultJSON struct {
	DaysUntilNextRotation   apijson.Field
	KeyRotationIntervalDays apijson.Field
	LastKeyRotationAt       apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SingleResponseKeyResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseKeyResultJSON) RawJSON() string {
	return r.raw
}

type AccountAccessKeyUpdateParams struct {
	// The number of days between key rotations.
	KeyRotationIntervalDays param.Field[float64] `json:"key_rotation_interval_days,required"`
}

func (r AccountAccessKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
