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

// AccountAccessGatewayCaService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessGatewayCaService] method instead.
type AccountAccessGatewayCaService struct {
	Options []option.RequestOption
}

// NewAccountAccessGatewayCaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessGatewayCaService(opts ...option.RequestOption) (r *AccountAccessGatewayCaService) {
	r = &AccountAccessGatewayCaService{}
	r.Options = opts
	return
}

// Adds a new SSH Certificate Authority (CA).
func (r *AccountAccessGatewayCaService) New(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAccessGatewayCaNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/gateway_ca", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Lists SSH Certificate Authorities (CA).
func (r *AccountAccessGatewayCaService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAccessGatewayCaListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/gateway_ca", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an SSH Certificate Authority.
func (r *AccountAccessGatewayCaService) Delete(ctx context.Context, accountID string, certificateID string, opts ...option.RequestOption) (res *IDResponseApps, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/gateway_ca/%s", accountID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type SchemasCertificates struct {
	// The key ID of this certificate.
	ID string `json:"id"`
	// The public key of this certificate.
	PublicKey string                  `json:"public_key"`
	JSON      schemasCertificatesJSON `json:"-"`
}

// schemasCertificatesJSON contains the JSON metadata for the struct
// [SchemasCertificates]
type schemasCertificatesJSON struct {
	ID          apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemasCertificatesJSON) RawJSON() string {
	return r.raw
}

type AccountAccessGatewayCaNewResponse struct {
	Result SchemasCertificates                   `json:"result"`
	JSON   accountAccessGatewayCaNewResponseJSON `json:"-"`
	APIResponseSingleAccess
}

// accountAccessGatewayCaNewResponseJSON contains the JSON metadata for the struct
// [AccountAccessGatewayCaNewResponse]
type accountAccessGatewayCaNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGatewayCaNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessGatewayCaNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessGatewayCaListResponse struct {
	Result []SchemasCertificates                  `json:"result"`
	JSON   accountAccessGatewayCaListResponseJSON `json:"-"`
	APIResponseCollectionAccess
}

// accountAccessGatewayCaListResponseJSON contains the JSON metadata for the struct
// [AccountAccessGatewayCaListResponse]
type accountAccessGatewayCaListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGatewayCaListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessGatewayCaListResponseJSON) RawJSON() string {
	return r.raw
}
