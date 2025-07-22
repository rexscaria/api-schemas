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

// AccountAccessAppCaService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessAppCaService] method instead.
type AccountAccessAppCaService struct {
	Options []option.RequestOption
}

// NewAccountAccessAppCaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessAppCaService(opts ...option.RequestOption) (r *AccountAccessAppCaService) {
	r = &AccountAccessAppCaService{}
	r.Options = opts
	return
}

// Generates a new short-lived certificate CA and public key.
func (r *AccountAccessAppCaService) New(ctx context.Context, accountID string, appID string, opts ...option.RequestOption) (res *SingleResponseCa, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/apps/%s/ca", accountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Fetches a short-lived certificate CA and its public key.
func (r *AccountAccessAppCaService) Get(ctx context.Context, accountID string, appID string, opts ...option.RequestOption) (res *SingleResponseCa, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/apps/%s/ca", accountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists short-lived certificate CAs and their public keys.
func (r *AccountAccessAppCaService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *ResponseCollectionCa, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/apps/ca", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a short-lived certificate CA.
func (r *AccountAccessAppCaService) Delete(ctx context.Context, accountID string, appID string, opts ...option.RequestOption) (res *SchemasIDResponseAccess, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/apps/%s/ca", accountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Ca struct {
	// The ID of the CA.
	ID string `json:"id"`
	// The Application Audience (AUD) tag. Identifies the application associated with
	// the CA.
	Aud string `json:"aud"`
	// The public key to add to your SSH server configuration.
	PublicKey string `json:"public_key"`
	JSON      caJSON `json:"-"`
}

// caJSON contains the JSON metadata for the struct [Ca]
type caJSON struct {
	ID          apijson.Field
	Aud         apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Ca) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caJSON) RawJSON() string {
	return r.raw
}

type ResponseCollectionCa struct {
	Result []Ca                     `json:"result"`
	JSON   responseCollectionCaJSON `json:"-"`
	APIResponseCollectionAccess
}

// responseCollectionCaJSON contains the JSON metadata for the struct
// [ResponseCollectionCa]
type responseCollectionCaJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionCa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionCaJSON) RawJSON() string {
	return r.raw
}

type SchemasIDResponseAccess struct {
	Result SchemasIDResponseAccessResult `json:"result"`
	JSON   schemasIDResponseAccessJSON   `json:"-"`
	APIResponseSingleAccess
}

// schemasIDResponseAccessJSON contains the JSON metadata for the struct
// [SchemasIDResponseAccess]
type schemasIDResponseAccessJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasIDResponseAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemasIDResponseAccessJSON) RawJSON() string {
	return r.raw
}

type SchemasIDResponseAccessResult struct {
	// The ID of the CA.
	ID   string                            `json:"id"`
	JSON schemasIDResponseAccessResultJSON `json:"-"`
}

// schemasIDResponseAccessResultJSON contains the JSON metadata for the struct
// [SchemasIDResponseAccessResult]
type schemasIDResponseAccessResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasIDResponseAccessResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemasIDResponseAccessResultJSON) RawJSON() string {
	return r.raw
}

type SingleResponseCa struct {
	Result Ca                   `json:"result"`
	JSON   singleResponseCaJSON `json:"-"`
	APIResponseSingleAccess
}

// singleResponseCaJSON contains the JSON metadata for the struct
// [SingleResponseCa]
type singleResponseCaJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseCa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseCaJSON) RawJSON() string {
	return r.raw
}
