// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// AccountAIGatewayGatewayDatasetService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIGatewayGatewayDatasetService] method instead.
type AccountAIGatewayGatewayDatasetService struct {
	Options []option.RequestOption
}

// NewAccountAIGatewayGatewayDatasetService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAIGatewayGatewayDatasetService(opts ...option.RequestOption) (r *AccountAIGatewayGatewayDatasetService) {
	r = &AccountAIGatewayGatewayDatasetService{}
	r.Options = opts
	return
}

// Create a new Dataset
func (r *AccountAIGatewayGatewayDatasetService) NewDataset(ctx context.Context, accountID string, gatewayID string, body AccountAIGatewayGatewayDatasetNewDatasetParams, opts ...option.RequestOption) (res *AccountAIGatewayGatewayDatasetNewDatasetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/datasets", accountID, gatewayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a Dataset
func (r *AccountAIGatewayGatewayDatasetService) DeleteDataset(ctx context.Context, accountID string, gatewayID string, id string, opts ...option.RequestOption) (res *AccountAIGatewayGatewayDatasetDeleteDatasetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/datasets/%s", accountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Fetch a Dataset
func (r *AccountAIGatewayGatewayDatasetService) FetchDataset(ctx context.Context, accountID string, gatewayID string, id string, opts ...option.RequestOption) (res *AccountAIGatewayGatewayDatasetFetchDatasetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/datasets/%s", accountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Datasets
func (r *AccountAIGatewayGatewayDatasetService) ListDatasets(ctx context.Context, accountID string, gatewayID string, query AccountAIGatewayGatewayDatasetListDatasetsParams, opts ...option.RequestOption) (res *AccountAIGatewayGatewayDatasetListDatasetsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/datasets", accountID, gatewayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update a Dataset
func (r *AccountAIGatewayGatewayDatasetService) UpdateDataset(ctx context.Context, accountID string, gatewayID string, id string, body AccountAIGatewayGatewayDatasetUpdateDatasetParams, opts ...option.RequestOption) (res *AccountAIGatewayGatewayDatasetUpdateDatasetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/datasets/%s", accountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountAIGatewayGatewayDatasetNewDatasetResponse struct {
	Result  AccountAIGatewayGatewayDatasetNewDatasetResponseResult `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    accountAIGatewayGatewayDatasetNewDatasetResponseJSON   `json:"-"`
}

// accountAIGatewayGatewayDatasetNewDatasetResponseJSON contains the JSON metadata
// for the struct [AccountAIGatewayGatewayDatasetNewDatasetResponse]
type accountAIGatewayGatewayDatasetNewDatasetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayDatasetNewDatasetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayDatasetNewDatasetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayDatasetNewDatasetResponseResult struct {
	ID         string                                                         `json:"id,required"`
	AccountID  string                                                         `json:"account_id,required"`
	AccountTag string                                                         `json:"account_tag,required"`
	CreatedAt  time.Time                                                      `json:"created_at,required" format:"date-time"`
	Enable     bool                                                           `json:"enable,required"`
	Filters    []AccountAIGatewayGatewayDatasetNewDatasetResponseResultFilter `json:"filters,required"`
	// gateway id
	GatewayID  string                                                     `json:"gateway_id,required"`
	ModifiedAt time.Time                                                  `json:"modified_at,required" format:"date-time"`
	Name       string                                                     `json:"name,required"`
	JSON       accountAIGatewayGatewayDatasetNewDatasetResponseResultJSON `json:"-"`
}

// accountAIGatewayGatewayDatasetNewDatasetResponseResultJSON contains the JSON
// metadata for the struct [AccountAIGatewayGatewayDatasetNewDatasetResponseResult]
type accountAIGatewayGatewayDatasetNewDatasetResponseResultJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Enable      apijson.Field
	Filters     apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayDatasetNewDatasetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayDatasetNewDatasetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayDatasetNewDatasetResponseResultFilter struct {
	Key      AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKey          `json:"key,required"`
	Operator AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersOperator     `json:"operator,required"`
	Value    []AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersValueUnion `json:"value,required"`
	JSON     accountAIGatewayGatewayDatasetNewDatasetResponseResultFilterJSON          `json:"-"`
}

// accountAIGatewayGatewayDatasetNewDatasetResponseResultFilterJSON contains the
// JSON metadata for the struct
// [AccountAIGatewayGatewayDatasetNewDatasetResponseResultFilter]
type accountAIGatewayGatewayDatasetNewDatasetResponseResultFilterJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayDatasetNewDatasetResponseResultFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayDatasetNewDatasetResponseResultFilterJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKey string

const (
	AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyCreatedAt           AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKey = "created_at"
	AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyRequestContentType  AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKey = "request_content_type"
	AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyResponseContentType AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKey = "response_content_type"
	AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeySuccess             AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKey = "success"
	AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyCached              AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKey = "cached"
	AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyProvider            AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKey = "provider"
	AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyModel               AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKey = "model"
	AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyCost                AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKey = "cost"
	AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyTokens              AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKey = "tokens"
	AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyTokensIn            AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKey = "tokens_in"
	AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyTokensOut           AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKey = "tokens_out"
	AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyDuration            AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKey = "duration"
	AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyFeedback            AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKey = "feedback"
)

func (r AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKey) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyCreatedAt, AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyRequestContentType, AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyResponseContentType, AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeySuccess, AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyCached, AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyProvider, AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyModel, AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyCost, AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyTokens, AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyTokensIn, AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyTokensOut, AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyDuration, AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersKeyFeedback:
		return true
	}
	return false
}

type AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersOperator string

const (
	AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersOperatorEq       AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersOperator = "eq"
	AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersOperatorContains AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersOperator = "contains"
	AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersOperatorLt       AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersOperator = "lt"
	AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersOperatorGt       AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersOperator = "gt"
)

func (r AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersOperator) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersOperatorEq, AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersOperatorContains, AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersOperatorLt, AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersOperatorGt:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersValueUnion interface {
	ImplementsAccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type AccountAIGatewayGatewayDatasetDeleteDatasetResponse struct {
	Result  AccountAIGatewayGatewayDatasetDeleteDatasetResponseResult `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    accountAIGatewayGatewayDatasetDeleteDatasetResponseJSON   `json:"-"`
}

// accountAIGatewayGatewayDatasetDeleteDatasetResponseJSON contains the JSON
// metadata for the struct [AccountAIGatewayGatewayDatasetDeleteDatasetResponse]
type accountAIGatewayGatewayDatasetDeleteDatasetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayDatasetDeleteDatasetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayDatasetDeleteDatasetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayDatasetDeleteDatasetResponseResult struct {
	ID         string                                                            `json:"id,required"`
	AccountID  string                                                            `json:"account_id,required"`
	AccountTag string                                                            `json:"account_tag,required"`
	CreatedAt  time.Time                                                         `json:"created_at,required" format:"date-time"`
	Enable     bool                                                              `json:"enable,required"`
	Filters    []AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFilter `json:"filters,required"`
	// gateway id
	GatewayID  string                                                        `json:"gateway_id,required"`
	ModifiedAt time.Time                                                     `json:"modified_at,required" format:"date-time"`
	Name       string                                                        `json:"name,required"`
	JSON       accountAIGatewayGatewayDatasetDeleteDatasetResponseResultJSON `json:"-"`
}

// accountAIGatewayGatewayDatasetDeleteDatasetResponseResultJSON contains the JSON
// metadata for the struct
// [AccountAIGatewayGatewayDatasetDeleteDatasetResponseResult]
type accountAIGatewayGatewayDatasetDeleteDatasetResponseResultJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Enable      apijson.Field
	Filters     apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayDatasetDeleteDatasetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayDatasetDeleteDatasetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFilter struct {
	Key      AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKey          `json:"key,required"`
	Operator AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersOperator     `json:"operator,required"`
	Value    []AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersValueUnion `json:"value,required"`
	JSON     accountAIGatewayGatewayDatasetDeleteDatasetResponseResultFilterJSON          `json:"-"`
}

// accountAIGatewayGatewayDatasetDeleteDatasetResponseResultFilterJSON contains the
// JSON metadata for the struct
// [AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFilter]
type accountAIGatewayGatewayDatasetDeleteDatasetResponseResultFilterJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayDatasetDeleteDatasetResponseResultFilterJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKey string

const (
	AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyCreatedAt           AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKey = "created_at"
	AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyRequestContentType  AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKey = "request_content_type"
	AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyResponseContentType AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKey = "response_content_type"
	AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeySuccess             AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKey = "success"
	AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyCached              AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKey = "cached"
	AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyProvider            AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKey = "provider"
	AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyModel               AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKey = "model"
	AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyCost                AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKey = "cost"
	AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyTokens              AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKey = "tokens"
	AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyTokensIn            AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKey = "tokens_in"
	AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyTokensOut           AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKey = "tokens_out"
	AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyDuration            AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKey = "duration"
	AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyFeedback            AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKey = "feedback"
)

func (r AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKey) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyCreatedAt, AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyRequestContentType, AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyResponseContentType, AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeySuccess, AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyCached, AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyProvider, AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyModel, AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyCost, AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyTokens, AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyTokensIn, AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyTokensOut, AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyDuration, AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersKeyFeedback:
		return true
	}
	return false
}

type AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersOperator string

const (
	AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersOperatorEq       AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersOperator = "eq"
	AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersOperatorContains AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersOperator = "contains"
	AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersOperatorLt       AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersOperator = "lt"
	AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersOperatorGt       AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersOperator = "gt"
)

func (r AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersOperator) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersOperatorEq, AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersOperatorContains, AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersOperatorLt, AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersOperatorGt:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersValueUnion interface {
	ImplementsAccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type AccountAIGatewayGatewayDatasetFetchDatasetResponse struct {
	Result  AccountAIGatewayGatewayDatasetFetchDatasetResponseResult `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    accountAIGatewayGatewayDatasetFetchDatasetResponseJSON   `json:"-"`
}

// accountAIGatewayGatewayDatasetFetchDatasetResponseJSON contains the JSON
// metadata for the struct [AccountAIGatewayGatewayDatasetFetchDatasetResponse]
type accountAIGatewayGatewayDatasetFetchDatasetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayDatasetFetchDatasetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayDatasetFetchDatasetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayDatasetFetchDatasetResponseResult struct {
	ID         string                                                           `json:"id,required"`
	AccountID  string                                                           `json:"account_id,required"`
	AccountTag string                                                           `json:"account_tag,required"`
	CreatedAt  time.Time                                                        `json:"created_at,required" format:"date-time"`
	Enable     bool                                                             `json:"enable,required"`
	Filters    []AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFilter `json:"filters,required"`
	// gateway id
	GatewayID  string                                                       `json:"gateway_id,required"`
	ModifiedAt time.Time                                                    `json:"modified_at,required" format:"date-time"`
	Name       string                                                       `json:"name,required"`
	JSON       accountAIGatewayGatewayDatasetFetchDatasetResponseResultJSON `json:"-"`
}

// accountAIGatewayGatewayDatasetFetchDatasetResponseResultJSON contains the JSON
// metadata for the struct
// [AccountAIGatewayGatewayDatasetFetchDatasetResponseResult]
type accountAIGatewayGatewayDatasetFetchDatasetResponseResultJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Enable      apijson.Field
	Filters     apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayDatasetFetchDatasetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayDatasetFetchDatasetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFilter struct {
	Key      AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKey          `json:"key,required"`
	Operator AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersOperator     `json:"operator,required"`
	Value    []AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersValueUnion `json:"value,required"`
	JSON     accountAIGatewayGatewayDatasetFetchDatasetResponseResultFilterJSON          `json:"-"`
}

// accountAIGatewayGatewayDatasetFetchDatasetResponseResultFilterJSON contains the
// JSON metadata for the struct
// [AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFilter]
type accountAIGatewayGatewayDatasetFetchDatasetResponseResultFilterJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayDatasetFetchDatasetResponseResultFilterJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKey string

const (
	AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyCreatedAt           AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKey = "created_at"
	AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyRequestContentType  AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKey = "request_content_type"
	AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyResponseContentType AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKey = "response_content_type"
	AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeySuccess             AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKey = "success"
	AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyCached              AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKey = "cached"
	AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyProvider            AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKey = "provider"
	AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyModel               AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKey = "model"
	AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyCost                AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKey = "cost"
	AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyTokens              AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKey = "tokens"
	AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyTokensIn            AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKey = "tokens_in"
	AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyTokensOut           AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKey = "tokens_out"
	AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyDuration            AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKey = "duration"
	AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyFeedback            AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKey = "feedback"
)

func (r AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKey) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyCreatedAt, AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyRequestContentType, AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyResponseContentType, AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeySuccess, AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyCached, AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyProvider, AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyModel, AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyCost, AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyTokens, AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyTokensIn, AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyTokensOut, AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyDuration, AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersKeyFeedback:
		return true
	}
	return false
}

type AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersOperator string

const (
	AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersOperatorEq       AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersOperator = "eq"
	AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersOperatorContains AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersOperator = "contains"
	AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersOperatorLt       AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersOperator = "lt"
	AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersOperatorGt       AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersOperator = "gt"
)

func (r AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersOperator) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersOperatorEq, AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersOperatorContains, AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersOperatorLt, AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersOperatorGt:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersValueUnion interface {
	ImplementsAccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type AccountAIGatewayGatewayDatasetListDatasetsResponse struct {
	Result  []AccountAIGatewayGatewayDatasetListDatasetsResponseResult `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    accountAIGatewayGatewayDatasetListDatasetsResponseJSON     `json:"-"`
}

// accountAIGatewayGatewayDatasetListDatasetsResponseJSON contains the JSON
// metadata for the struct [AccountAIGatewayGatewayDatasetListDatasetsResponse]
type accountAIGatewayGatewayDatasetListDatasetsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayDatasetListDatasetsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayDatasetListDatasetsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayDatasetListDatasetsResponseResult struct {
	ID         string                                                           `json:"id,required"`
	AccountID  string                                                           `json:"account_id,required"`
	AccountTag string                                                           `json:"account_tag,required"`
	CreatedAt  time.Time                                                        `json:"created_at,required" format:"date-time"`
	Enable     bool                                                             `json:"enable,required"`
	Filters    []AccountAIGatewayGatewayDatasetListDatasetsResponseResultFilter `json:"filters,required"`
	// gateway id
	GatewayID  string                                                       `json:"gateway_id,required"`
	ModifiedAt time.Time                                                    `json:"modified_at,required" format:"date-time"`
	Name       string                                                       `json:"name,required"`
	JSON       accountAIGatewayGatewayDatasetListDatasetsResponseResultJSON `json:"-"`
}

// accountAIGatewayGatewayDatasetListDatasetsResponseResultJSON contains the JSON
// metadata for the struct
// [AccountAIGatewayGatewayDatasetListDatasetsResponseResult]
type accountAIGatewayGatewayDatasetListDatasetsResponseResultJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Enable      apijson.Field
	Filters     apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayDatasetListDatasetsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayDatasetListDatasetsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayDatasetListDatasetsResponseResultFilter struct {
	Key      AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKey          `json:"key,required"`
	Operator AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersOperator     `json:"operator,required"`
	Value    []AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersValueUnion `json:"value,required"`
	JSON     accountAIGatewayGatewayDatasetListDatasetsResponseResultFilterJSON          `json:"-"`
}

// accountAIGatewayGatewayDatasetListDatasetsResponseResultFilterJSON contains the
// JSON metadata for the struct
// [AccountAIGatewayGatewayDatasetListDatasetsResponseResultFilter]
type accountAIGatewayGatewayDatasetListDatasetsResponseResultFilterJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayDatasetListDatasetsResponseResultFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayDatasetListDatasetsResponseResultFilterJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKey string

const (
	AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyCreatedAt           AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKey = "created_at"
	AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyRequestContentType  AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKey = "request_content_type"
	AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyResponseContentType AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKey = "response_content_type"
	AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeySuccess             AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKey = "success"
	AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyCached              AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKey = "cached"
	AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyProvider            AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKey = "provider"
	AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyModel               AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKey = "model"
	AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyCost                AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKey = "cost"
	AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyTokens              AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKey = "tokens"
	AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyTokensIn            AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKey = "tokens_in"
	AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyTokensOut           AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKey = "tokens_out"
	AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyDuration            AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKey = "duration"
	AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyFeedback            AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKey = "feedback"
)

func (r AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKey) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyCreatedAt, AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyRequestContentType, AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyResponseContentType, AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeySuccess, AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyCached, AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyProvider, AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyModel, AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyCost, AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyTokens, AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyTokensIn, AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyTokensOut, AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyDuration, AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersKeyFeedback:
		return true
	}
	return false
}

type AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersOperator string

const (
	AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersOperatorEq       AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersOperator = "eq"
	AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersOperatorContains AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersOperator = "contains"
	AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersOperatorLt       AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersOperator = "lt"
	AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersOperatorGt       AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersOperator = "gt"
)

func (r AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersOperator) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersOperatorEq, AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersOperatorContains, AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersOperatorLt, AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersOperatorGt:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersValueUnion interface {
	ImplementsAccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type AccountAIGatewayGatewayDatasetUpdateDatasetResponse struct {
	Result  AccountAIGatewayGatewayDatasetUpdateDatasetResponseResult `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    accountAIGatewayGatewayDatasetUpdateDatasetResponseJSON   `json:"-"`
}

// accountAIGatewayGatewayDatasetUpdateDatasetResponseJSON contains the JSON
// metadata for the struct [AccountAIGatewayGatewayDatasetUpdateDatasetResponse]
type accountAIGatewayGatewayDatasetUpdateDatasetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayDatasetUpdateDatasetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayDatasetUpdateDatasetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayDatasetUpdateDatasetResponseResult struct {
	ID         string                                                            `json:"id,required"`
	AccountID  string                                                            `json:"account_id,required"`
	AccountTag string                                                            `json:"account_tag,required"`
	CreatedAt  time.Time                                                         `json:"created_at,required" format:"date-time"`
	Enable     bool                                                              `json:"enable,required"`
	Filters    []AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFilter `json:"filters,required"`
	// gateway id
	GatewayID  string                                                        `json:"gateway_id,required"`
	ModifiedAt time.Time                                                     `json:"modified_at,required" format:"date-time"`
	Name       string                                                        `json:"name,required"`
	JSON       accountAIGatewayGatewayDatasetUpdateDatasetResponseResultJSON `json:"-"`
}

// accountAIGatewayGatewayDatasetUpdateDatasetResponseResultJSON contains the JSON
// metadata for the struct
// [AccountAIGatewayGatewayDatasetUpdateDatasetResponseResult]
type accountAIGatewayGatewayDatasetUpdateDatasetResponseResultJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Enable      apijson.Field
	Filters     apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayDatasetUpdateDatasetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayDatasetUpdateDatasetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFilter struct {
	Key      AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKey          `json:"key,required"`
	Operator AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersOperator     `json:"operator,required"`
	Value    []AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersValueUnion `json:"value,required"`
	JSON     accountAIGatewayGatewayDatasetUpdateDatasetResponseResultFilterJSON          `json:"-"`
}

// accountAIGatewayGatewayDatasetUpdateDatasetResponseResultFilterJSON contains the
// JSON metadata for the struct
// [AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFilter]
type accountAIGatewayGatewayDatasetUpdateDatasetResponseResultFilterJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayDatasetUpdateDatasetResponseResultFilterJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKey string

const (
	AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyCreatedAt           AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKey = "created_at"
	AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyRequestContentType  AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKey = "request_content_type"
	AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyResponseContentType AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKey = "response_content_type"
	AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeySuccess             AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKey = "success"
	AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyCached              AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKey = "cached"
	AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyProvider            AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKey = "provider"
	AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyModel               AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKey = "model"
	AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyCost                AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKey = "cost"
	AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyTokens              AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKey = "tokens"
	AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyTokensIn            AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKey = "tokens_in"
	AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyTokensOut           AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKey = "tokens_out"
	AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyDuration            AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKey = "duration"
	AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyFeedback            AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKey = "feedback"
)

func (r AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKey) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyCreatedAt, AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyRequestContentType, AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyResponseContentType, AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeySuccess, AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyCached, AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyProvider, AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyModel, AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyCost, AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyTokens, AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyTokensIn, AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyTokensOut, AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyDuration, AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersKeyFeedback:
		return true
	}
	return false
}

type AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersOperator string

const (
	AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersOperatorEq       AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersOperator = "eq"
	AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersOperatorContains AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersOperator = "contains"
	AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersOperatorLt       AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersOperator = "lt"
	AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersOperatorGt       AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersOperator = "gt"
)

func (r AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersOperator) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersOperatorEq, AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersOperatorContains, AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersOperatorLt, AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersOperatorGt:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersValueUnion interface {
	ImplementsAccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type AccountAIGatewayGatewayDatasetNewDatasetParams struct {
	Enable  param.Field[bool]                                                   `json:"enable,required"`
	Filters param.Field[[]AccountAIGatewayGatewayDatasetNewDatasetParamsFilter] `json:"filters,required"`
	Name    param.Field[string]                                                 `json:"name,required"`
}

func (r AccountAIGatewayGatewayDatasetNewDatasetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIGatewayGatewayDatasetNewDatasetParamsFilter struct {
	Key      param.Field[AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKey]          `json:"key,required"`
	Operator param.Field[AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersOperator]     `json:"operator,required"`
	Value    param.Field[[]AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersValueUnion] `json:"value,required"`
}

func (r AccountAIGatewayGatewayDatasetNewDatasetParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKey string

const (
	AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyCreatedAt           AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKey = "created_at"
	AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyRequestContentType  AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKey = "request_content_type"
	AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyResponseContentType AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKey = "response_content_type"
	AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeySuccess             AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKey = "success"
	AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyCached              AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKey = "cached"
	AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyProvider            AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKey = "provider"
	AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyModel               AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKey = "model"
	AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyCost                AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKey = "cost"
	AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyTokens              AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKey = "tokens"
	AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyTokensIn            AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKey = "tokens_in"
	AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyTokensOut           AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKey = "tokens_out"
	AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyDuration            AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKey = "duration"
	AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyFeedback            AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKey = "feedback"
)

func (r AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKey) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyCreatedAt, AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyRequestContentType, AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyResponseContentType, AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeySuccess, AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyCached, AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyProvider, AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyModel, AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyCost, AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyTokens, AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyTokensIn, AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyTokensOut, AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyDuration, AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersKeyFeedback:
		return true
	}
	return false
}

type AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersOperator string

const (
	AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersOperatorEq       AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersOperator = "eq"
	AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersOperatorContains AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersOperator = "contains"
	AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersOperatorLt       AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersOperator = "lt"
	AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersOperatorGt       AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersOperator = "gt"
)

func (r AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersOperator) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersOperatorEq, AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersOperatorContains, AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersOperatorLt, AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersOperatorGt:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type AccountAIGatewayGatewayDatasetNewDatasetParamsFiltersValueUnion interface {
	ImplementsAccountAIGatewayGatewayDatasetNewDatasetParamsFiltersValueUnion()
}

type AccountAIGatewayGatewayDatasetListDatasetsParams struct {
	Enable  param.Field[bool]   `query:"enable"`
	Name    param.Field[string] `query:"name"`
	Page    param.Field[int64]  `query:"page"`
	PerPage param.Field[int64]  `query:"per_page"`
	// Search by id, name, filters
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccountAIGatewayGatewayDatasetListDatasetsParams]'s query
// parameters as `url.Values`.
func (r AccountAIGatewayGatewayDatasetListDatasetsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIGatewayGatewayDatasetUpdateDatasetParams struct {
	Enable  param.Field[bool]                                                      `json:"enable,required"`
	Filters param.Field[[]AccountAIGatewayGatewayDatasetUpdateDatasetParamsFilter] `json:"filters,required"`
	Name    param.Field[string]                                                    `json:"name,required"`
}

func (r AccountAIGatewayGatewayDatasetUpdateDatasetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIGatewayGatewayDatasetUpdateDatasetParamsFilter struct {
	Key      param.Field[AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKey]          `json:"key,required"`
	Operator param.Field[AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersOperator]     `json:"operator,required"`
	Value    param.Field[[]AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersValueUnion] `json:"value,required"`
}

func (r AccountAIGatewayGatewayDatasetUpdateDatasetParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKey string

const (
	AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyCreatedAt           AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKey = "created_at"
	AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyRequestContentType  AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKey = "request_content_type"
	AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyResponseContentType AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKey = "response_content_type"
	AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeySuccess             AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKey = "success"
	AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyCached              AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKey = "cached"
	AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyProvider            AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKey = "provider"
	AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyModel               AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKey = "model"
	AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyCost                AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKey = "cost"
	AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyTokens              AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKey = "tokens"
	AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyTokensIn            AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKey = "tokens_in"
	AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyTokensOut           AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKey = "tokens_out"
	AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyDuration            AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKey = "duration"
	AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyFeedback            AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKey = "feedback"
)

func (r AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKey) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyCreatedAt, AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyRequestContentType, AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyResponseContentType, AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeySuccess, AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyCached, AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyProvider, AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyModel, AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyCost, AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyTokens, AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyTokensIn, AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyTokensOut, AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyDuration, AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersKeyFeedback:
		return true
	}
	return false
}

type AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersOperator string

const (
	AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersOperatorEq       AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersOperator = "eq"
	AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersOperatorContains AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersOperator = "contains"
	AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersOperatorLt       AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersOperator = "lt"
	AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersOperatorGt       AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersOperator = "gt"
)

func (r AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersOperator) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersOperatorEq, AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersOperatorContains, AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersOperatorLt, AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersOperatorGt:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type AccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersValueUnion interface {
	ImplementsAccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersValueUnion()
}
