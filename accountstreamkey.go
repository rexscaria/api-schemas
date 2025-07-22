// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountStreamKeyService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountStreamKeyService] method instead.
type AccountStreamKeyService struct {
	Options []option.RequestOption
}

// NewAccountStreamKeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamKeyService(opts ...option.RequestOption) (r *AccountStreamKeyService) {
	r = &AccountStreamKeyService{}
	r.Options = opts
	return
}

// Creates an RSA private key in PEM and JWK formats. Key files are only displayed
// once after creation. Keys are created, used, and deleted independently of
// videos, and every key can sign any video.
func (r *AccountStreamKeyService) New(ctx context.Context, accountID string, body AccountStreamKeyNewParams, opts ...option.RequestOption) (res *AccountStreamKeyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/keys", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists the video ID and creation date and time when a signing key was created.
func (r *AccountStreamKeyService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountStreamKeyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/keys", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes signing keys and revokes all signed URLs generated with the key.
func (r *AccountStreamKeyService) Delete(ctx context.Context, accountID string, identifier string, body AccountStreamKeyDeleteParams, opts ...option.RequestOption) (res *DeletedStreamResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/keys/%s", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type AccountStreamKeyNewResponse struct {
	Result AccountStreamKeyNewResponseResult `json:"result"`
	JSON   accountStreamKeyNewResponseJSON   `json:"-"`
	APIResponseStream
}

// accountStreamKeyNewResponseJSON contains the JSON metadata for the struct
// [AccountStreamKeyNewResponse]
type accountStreamKeyNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamKeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamKeyNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountStreamKeyNewResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// The date and time a signing key was created.
	Created time.Time `json:"created" format:"date-time"`
	// The signing key in JWK format.
	Jwk string `json:"jwk"`
	// The signing key in PEM format.
	Pem  string                                `json:"pem"`
	JSON accountStreamKeyNewResponseResultJSON `json:"-"`
}

// accountStreamKeyNewResponseResultJSON contains the JSON metadata for the struct
// [AccountStreamKeyNewResponseResult]
type accountStreamKeyNewResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Jwk         apijson.Field
	Pem         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamKeyNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamKeyNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountStreamKeyListResponse struct {
	Result []AccountStreamKeyListResponseResult `json:"result"`
	JSON   accountStreamKeyListResponseJSON     `json:"-"`
	APIResponseStream
}

// accountStreamKeyListResponseJSON contains the JSON metadata for the struct
// [AccountStreamKeyListResponse]
type accountStreamKeyListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamKeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamKeyListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountStreamKeyListResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// The date and time a signing key was created.
	Created time.Time                              `json:"created" format:"date-time"`
	JSON    accountStreamKeyListResponseResultJSON `json:"-"`
}

// accountStreamKeyListResponseResultJSON contains the JSON metadata for the struct
// [AccountStreamKeyListResponseResult]
type accountStreamKeyListResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamKeyListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStreamKeyListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountStreamKeyNewParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountStreamKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountStreamKeyDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountStreamKeyDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
