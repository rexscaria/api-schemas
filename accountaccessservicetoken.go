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

// AccountAccessServiceTokenService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessServiceTokenService] method instead.
type AccountAccessServiceTokenService struct {
	Options []option.RequestOption
}

// NewAccountAccessServiceTokenService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessServiceTokenService(opts ...option.RequestOption) (r *AccountAccessServiceTokenService) {
	r = &AccountAccessServiceTokenService{}
	r.Options = opts
	return
}

// Refreshes the expiration of a service token.
func (r *AccountAccessServiceTokenService) Refresh(ctx context.Context, accountID string, serviceTokenID string, opts ...option.RequestOption) (res *SchemasAccessSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if serviceTokenID == "" {
		err = errors.New("missing required service_token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s/refresh", accountID, serviceTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Generates a new Client Secret for a service token and revokes the old one.
func (r *AccountAccessServiceTokenService) Rotate(ctx context.Context, accountID string, serviceTokenID string, opts ...option.RequestOption) (res *CreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if serviceTokenID == "" {
		err = errors.New("missing required service_token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s/rotate", accountID, serviceTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type CreateResponse struct {
	Errors   []MessagesAccessItem `json:"errors,required"`
	Messages []MessagesAccessItem `json:"messages,required"`
	// Whether the API call was successful.
	Success CreateResponseSuccess `json:"success,required"`
	Result  CreateResponseResult  `json:"result"`
	JSON    createResponseJSON    `json:"-"`
}

// createResponseJSON contains the JSON metadata for the struct [CreateResponse]
type createResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CreateResponseSuccess bool

const (
	CreateResponseSuccessTrue CreateResponseSuccess = true
)

func (r CreateResponseSuccess) IsKnown() bool {
	switch r {
	case CreateResponseSuccessTrue:
		return true
	}
	return false
}

type CreateResponseResult struct {
	// The ID of the service token.
	ID string `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID string `json:"client_id"`
	// The Client Secret for the service token. Access will check for this value in the
	// `CF-Access-Client-Secret` request header.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                   `json:"name"`
	UpdatedAt time.Time                `json:"updated_at" format:"date-time"`
	JSON      createResponseResultJSON `json:"-"`
}

// createResponseResultJSON contains the JSON metadata for the struct
// [CreateResponseResult]
type createResponseResultJSON struct {
	ID           apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	CreatedAt    apijson.Field
	Duration     apijson.Field
	Name         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CreateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createResponseResultJSON) RawJSON() string {
	return r.raw
}

type SchemasAccessSingleResponse struct {
	Errors   []MessagesAccessItem `json:"errors,required"`
	Messages []MessagesAccessItem `json:"messages,required"`
	// Whether the API call was successful.
	Success SchemasAccessSingleResponseSuccess `json:"success,required"`
	Result  ServiceTokens                      `json:"result"`
	JSON    schemasAccessSingleResponseJSON    `json:"-"`
}

// schemasAccessSingleResponseJSON contains the JSON metadata for the struct
// [SchemasAccessSingleResponse]
type schemasAccessSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasAccessSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemasAccessSingleResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SchemasAccessSingleResponseSuccess bool

const (
	SchemasAccessSingleResponseSuccessTrue SchemasAccessSingleResponseSuccess = true
)

func (r SchemasAccessSingleResponseSuccess) IsKnown() bool {
	switch r {
	case SchemasAccessSingleResponseSuccessTrue:
		return true
	}
	return false
}

type ServiceTokens struct {
	// The ID of the service token.
	ID string `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration   string    `json:"duration"`
	ExpiresAt  time.Time `json:"expires_at" format:"date-time"`
	LastSeenAt time.Time `json:"last_seen_at" format:"date-time"`
	// The name of the service token.
	Name      string            `json:"name"`
	UpdatedAt time.Time         `json:"updated_at" format:"date-time"`
	JSON      serviceTokensJSON `json:"-"`
}

// serviceTokensJSON contains the JSON metadata for the struct [ServiceTokens]
type serviceTokensJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	ExpiresAt   apijson.Field
	LastSeenAt  apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceTokensJSON) RawJSON() string {
	return r.raw
}
