// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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

// Generates a new service token. **Note:** This is the only time you can get the
// Client Secret. If you lose the Client Secret, you will have to rotate the Client
// Secret or create a new service token.
func (r *AccountAccessServiceTokenService) New(ctx context.Context, accountID string, body AccountAccessServiceTokenNewParams, opts ...option.RequestOption) (res *CreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/service_tokens", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single service token.
func (r *AccountAccessServiceTokenService) Get(ctx context.Context, accountID string, serviceTokenID string, opts ...option.RequestOption) (res *SchemasAccessSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if serviceTokenID == "" {
		err = errors.New("missing required service_token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s", accountID, serviceTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured service token.
func (r *AccountAccessServiceTokenService) Update(ctx context.Context, accountID string, serviceTokenID string, body AccountAccessServiceTokenUpdateParams, opts ...option.RequestOption) (res *SchemasAccessSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if serviceTokenID == "" {
		err = errors.New("missing required service_token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s", accountID, serviceTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all service tokens.
func (r *AccountAccessServiceTokenService) List(ctx context.Context, accountID string, query AccountAccessServiceTokenListParams, opts ...option.RequestOption) (res *ResponseCollectionServiceTokens, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/service_tokens", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a service token.
func (r *AccountAccessServiceTokenService) Delete(ctx context.Context, accountID string, serviceTokenID string, opts ...option.RequestOption) (res *SchemasAccessSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if serviceTokenID == "" {
		err = errors.New("missing required service_token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s", accountID, serviceTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
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
	Result CreateResponseResult `json:"result"`
	JSON   createResponseJSON   `json:"-"`
	APIResponseSingleAccess
}

// createResponseJSON contains the JSON metadata for the struct [CreateResponse]
type createResponseJSON struct {
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

type ResponseCollectionServiceTokens struct {
	Result []ServiceTokens                     `json:"result"`
	JSON   responseCollectionServiceTokensJSON `json:"-"`
	APIResponseCollectionAccess
}

// responseCollectionServiceTokensJSON contains the JSON metadata for the struct
// [ResponseCollectionServiceTokens]
type responseCollectionServiceTokensJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionServiceTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionServiceTokensJSON) RawJSON() string {
	return r.raw
}

type SchemasAccessSingleResponse struct {
	Result ServiceTokens                   `json:"result"`
	JSON   schemasAccessSingleResponseJSON `json:"-"`
	APIResponseSingleAccess
}

// schemasAccessSingleResponseJSON contains the JSON metadata for the struct
// [SchemasAccessSingleResponse]
type schemasAccessSingleResponseJSON struct {
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

type ServiceTokens struct {
	// UUID
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

type AccountAccessServiceTokenNewParams struct {
	// The name of the service token.
	Name param.Field[string] `json:"name,required"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration param.Field[string] `json:"duration"`
}

func (r AccountAccessServiceTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessServiceTokenUpdateParams struct {
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration param.Field[string] `json:"duration"`
	// The name of the service token.
	Name param.Field[string] `json:"name"`
}

func (r AccountAccessServiceTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessServiceTokenListParams struct {
	// The name of the service token.
	Name param.Field[string] `query:"name"`
	// Search for service tokens by other listed query parameters.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccountAccessServiceTokenListParams]'s query parameters as
// `url.Values`.
func (r AccountAccessServiceTokenListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
