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

// AccountCallTurnKeyService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCallTurnKeyService] method instead.
type AccountCallTurnKeyService struct {
	Options []option.RequestOption
}

// NewAccountCallTurnKeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountCallTurnKeyService(opts ...option.RequestOption) (r *AccountCallTurnKeyService) {
	r = &AccountCallTurnKeyService{}
	r.Options = opts
	return
}

// Creates a new Cloudflare Calls TURN key.
func (r *AccountCallTurnKeyService) New(ctx context.Context, accountID string, body AccountCallTurnKeyNewParams, opts ...option.RequestOption) (res *AccountCallTurnKeyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/turn_keys", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches details for a single TURN key.
func (r *AccountCallTurnKeyService) Get(ctx context.Context, accountID string, keyID string, opts ...option.RequestOption) (res *CallsTurnKeyResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/turn_keys/%s", accountID, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Edit details for a single TURN key.
func (r *AccountCallTurnKeyService) Update(ctx context.Context, accountID string, keyID string, body AccountCallTurnKeyUpdateParams, opts ...option.RequestOption) (res *CallsTurnKeyResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/turn_keys/%s", accountID, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all TURN keys in the Cloudflare account
func (r *AccountCallTurnKeyService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountCallTurnKeyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/turn_keys", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a TURN key from Cloudflare Calls
func (r *AccountCallTurnKeyService) Delete(ctx context.Context, accountID string, keyID string, opts ...option.RequestOption) (res *CallsTurnKeyResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/turn_keys/%s", accountID, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CallsTurnKeyEditableFieldsParam struct {
	// A short description of a TURN key, not shown to end users.
	Name param.Field[string] `json:"name"`
}

func (r CallsTurnKeyEditableFieldsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CallsTurnKeyObject struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	Uid  string                 `json:"uid"`
	JSON callsTurnKeyObjectJSON `json:"-"`
}

// callsTurnKeyObjectJSON contains the JSON metadata for the struct
// [CallsTurnKeyObject]
type callsTurnKeyObjectJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallsTurnKeyObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callsTurnKeyObjectJSON) RawJSON() string {
	return r.raw
}

type CallsTurnKeyResponseSingle struct {
	Result CallsTurnKeyObject             `json:"result"`
	JSON   callsTurnKeyResponseSingleJSON `json:"-"`
	CallsAPIResponseSingle
}

// callsTurnKeyResponseSingleJSON contains the JSON metadata for the struct
// [CallsTurnKeyResponseSingle]
type callsTurnKeyResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallsTurnKeyResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callsTurnKeyResponseSingleJSON) RawJSON() string {
	return r.raw
}

type AccountCallTurnKeyNewResponse struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// Bearer token
	Key string `json:"key"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of a TURN key, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	Uid  string                            `json:"uid"`
	JSON accountCallTurnKeyNewResponseJSON `json:"-"`
}

// accountCallTurnKeyNewResponseJSON contains the JSON metadata for the struct
// [AccountCallTurnKeyNewResponse]
type accountCallTurnKeyNewResponseJSON struct {
	Created     apijson.Field
	Key         apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCallTurnKeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCallTurnKeyNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCallTurnKeyListResponse struct {
	Result []CallsTurnKeyObject               `json:"result"`
	JSON   accountCallTurnKeyListResponseJSON `json:"-"`
	CallsAPIResponseCommon
}

// accountCallTurnKeyListResponseJSON contains the JSON metadata for the struct
// [AccountCallTurnKeyListResponse]
type accountCallTurnKeyListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCallTurnKeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCallTurnKeyListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCallTurnKeyNewParams struct {
	CallsTurnKeyEditableFields CallsTurnKeyEditableFieldsParam `json:"calls_turn_key_editable_fields,required"`
}

func (r AccountCallTurnKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CallsTurnKeyEditableFields)
}

type AccountCallTurnKeyUpdateParams struct {
	CallsTurnKeyEditableFields CallsTurnKeyEditableFieldsParam `json:"calls_turn_key_editable_fields,required"`
}

func (r AccountCallTurnKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CallsTurnKeyEditableFields)
}
