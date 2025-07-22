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

// AccountAIGatewayGatewayEvaluationService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIGatewayGatewayEvaluationService] method instead.
type AccountAIGatewayGatewayEvaluationService struct {
	Options []option.RequestOption
}

// NewAccountAIGatewayGatewayEvaluationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAIGatewayGatewayEvaluationService(opts ...option.RequestOption) (r *AccountAIGatewayGatewayEvaluationService) {
	r = &AccountAIGatewayGatewayEvaluationService{}
	r.Options = opts
	return
}

// Create a new Evaluation
func (r *AccountAIGatewayGatewayEvaluationService) NewEvaluation(ctx context.Context, accountID string, gatewayID string, body AccountAIGatewayGatewayEvaluationNewEvaluationParams, opts ...option.RequestOption) (res *AccountAIGatewayGatewayEvaluationNewEvaluationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/evaluations", accountID, gatewayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a Evaluation
func (r *AccountAIGatewayGatewayEvaluationService) DeleteEvaluation(ctx context.Context, accountID string, gatewayID string, id string, opts ...option.RequestOption) (res *AccountAIGatewayGatewayEvaluationDeleteEvaluationResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/evaluations/%s", accountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Fetch a Evaluation
func (r *AccountAIGatewayGatewayEvaluationService) FetchEvaluation(ctx context.Context, accountID string, gatewayID string, id string, opts ...option.RequestOption) (res *AccountAIGatewayGatewayEvaluationFetchEvaluationResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/evaluations/%s", accountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Evaluations
func (r *AccountAIGatewayGatewayEvaluationService) ListEvaluations(ctx context.Context, accountID string, gatewayID string, query AccountAIGatewayGatewayEvaluationListEvaluationsParams, opts ...option.RequestOption) (res *AccountAIGatewayGatewayEvaluationListEvaluationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/evaluations", accountID, gatewayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountAIGatewayGatewayEvaluationNewEvaluationResponse struct {
	Result  AccountAIGatewayGatewayEvaluationNewEvaluationResponseResult `json:"result,required"`
	Success bool                                                         `json:"success,required"`
	JSON    accountAIGatewayGatewayEvaluationNewEvaluationResponseJSON   `json:"-"`
}

// accountAIGatewayGatewayEvaluationNewEvaluationResponseJSON contains the JSON
// metadata for the struct [AccountAIGatewayGatewayEvaluationNewEvaluationResponse]
type accountAIGatewayGatewayEvaluationNewEvaluationResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayEvaluationNewEvaluationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationNewEvaluationResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationNewEvaluationResponseResult struct {
	ID         string                                                                `json:"id,required"`
	AccountID  string                                                                `json:"account_id,required"`
	AccountTag string                                                                `json:"account_tag,required"`
	CreatedAt  time.Time                                                             `json:"created_at,required" format:"date-time"`
	Datasets   []AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDataset `json:"datasets,required"`
	// gateway id
	GatewayID  string                                                               `json:"gateway_id,required"`
	ModifiedAt time.Time                                                            `json:"modified_at,required" format:"date-time"`
	Name       string                                                               `json:"name,required"`
	Processed  bool                                                                 `json:"processed,required"`
	Results    []AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultResult `json:"results,required"`
	TotalLogs  float64                                                              `json:"total_logs,required"`
	JSON       accountAIGatewayGatewayEvaluationNewEvaluationResponseResultJSON     `json:"-"`
}

// accountAIGatewayGatewayEvaluationNewEvaluationResponseResultJSON contains the
// JSON metadata for the struct
// [AccountAIGatewayGatewayEvaluationNewEvaluationResponseResult]
type accountAIGatewayGatewayEvaluationNewEvaluationResponseResultJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Datasets    apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Processed   apijson.Field
	Results     apijson.Field
	TotalLogs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayEvaluationNewEvaluationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationNewEvaluationResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDataset struct {
	ID         string                                                                       `json:"id,required"`
	AccountID  string                                                                       `json:"account_id,required"`
	AccountTag string                                                                       `json:"account_tag,required"`
	CreatedAt  time.Time                                                                    `json:"created_at,required" format:"date-time"`
	Enable     bool                                                                         `json:"enable,required"`
	Filters    []AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFilter `json:"filters,required"`
	// gateway id
	GatewayID  string                                                                  `json:"gateway_id,required"`
	ModifiedAt time.Time                                                               `json:"modified_at,required" format:"date-time"`
	Name       string                                                                  `json:"name,required"`
	JSON       accountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetJSON `json:"-"`
}

// accountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetJSON contains
// the JSON metadata for the struct
// [AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDataset]
type accountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetJSON struct {
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

func (r *AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFilter struct {
	Key      AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKey          `json:"key,required"`
	Operator AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersOperator     `json:"operator,required"`
	Value    []AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersValueUnion `json:"value,required"`
	JSON     accountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFilterJSON          `json:"-"`
}

// accountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFilterJSON
// contains the JSON metadata for the struct
// [AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFilter]
type accountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFilterJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFilterJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKey string

const (
	AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyCreatedAt           AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKey = "created_at"
	AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyRequestContentType  AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKey = "request_content_type"
	AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyResponseContentType AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKey = "response_content_type"
	AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeySuccess             AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKey = "success"
	AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyCached              AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKey = "cached"
	AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyProvider            AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKey = "provider"
	AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyModel               AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKey = "model"
	AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyCost                AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKey = "cost"
	AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyTokens              AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKey = "tokens"
	AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyTokensIn            AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKey = "tokens_in"
	AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyTokensOut           AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKey = "tokens_out"
	AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyDuration            AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKey = "duration"
	AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyFeedback            AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKey = "feedback"
)

func (r AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKey) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyCreatedAt, AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyRequestContentType, AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyResponseContentType, AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeySuccess, AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyCached, AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyProvider, AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyModel, AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyCost, AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyTokens, AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyTokensIn, AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyTokensOut, AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyDuration, AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersKeyFeedback:
		return true
	}
	return false
}

type AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersOperator string

const (
	AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersOperatorEq       AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersOperator = "eq"
	AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersOperatorContains AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersOperator = "contains"
	AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersOperatorLt       AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersOperator = "lt"
	AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersOperatorGt       AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersOperator = "gt"
)

func (r AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersOperator) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersOperatorEq, AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersOperatorContains, AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersOperatorLt, AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersOperatorGt:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersValueUnion interface {
	ImplementsAccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersValueUnion)(nil)).Elem(),
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

type AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultResult struct {
	ID                string                                                                 `json:"id,required"`
	CreatedAt         time.Time                                                              `json:"created_at,required" format:"date-time"`
	EvaluationID      string                                                                 `json:"evaluation_id,required"`
	EvaluationTypeID  string                                                                 `json:"evaluation_type_id,required"`
	ModifiedAt        time.Time                                                              `json:"modified_at,required" format:"date-time"`
	Result            string                                                                 `json:"result,required"`
	Status            float64                                                                `json:"status,required"`
	StatusDescription string                                                                 `json:"status_description,required"`
	TotalLogs         float64                                                                `json:"total_logs,required"`
	JSON              accountAIGatewayGatewayEvaluationNewEvaluationResponseResultResultJSON `json:"-"`
}

// accountAIGatewayGatewayEvaluationNewEvaluationResponseResultResultJSON contains
// the JSON metadata for the struct
// [AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultResult]
type accountAIGatewayGatewayEvaluationNewEvaluationResponseResultResultJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	EvaluationID      apijson.Field
	EvaluationTypeID  apijson.Field
	ModifiedAt        apijson.Field
	Result            apijson.Field
	Status            apijson.Field
	StatusDescription apijson.Field
	TotalLogs         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayEvaluationNewEvaluationResponseResultResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationNewEvaluationResponseResultResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationDeleteEvaluationResponse struct {
	Result  AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResult `json:"result,required"`
	Success bool                                                            `json:"success,required"`
	JSON    accountAIGatewayGatewayEvaluationDeleteEvaluationResponseJSON   `json:"-"`
}

// accountAIGatewayGatewayEvaluationDeleteEvaluationResponseJSON contains the JSON
// metadata for the struct
// [AccountAIGatewayGatewayEvaluationDeleteEvaluationResponse]
type accountAIGatewayGatewayEvaluationDeleteEvaluationResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayEvaluationDeleteEvaluationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationDeleteEvaluationResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResult struct {
	ID         string                                                                   `json:"id,required"`
	AccountID  string                                                                   `json:"account_id,required"`
	AccountTag string                                                                   `json:"account_tag,required"`
	CreatedAt  time.Time                                                                `json:"created_at,required" format:"date-time"`
	Datasets   []AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDataset `json:"datasets,required"`
	// gateway id
	GatewayID  string                                                                  `json:"gateway_id,required"`
	ModifiedAt time.Time                                                               `json:"modified_at,required" format:"date-time"`
	Name       string                                                                  `json:"name,required"`
	Processed  bool                                                                    `json:"processed,required"`
	Results    []AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultResult `json:"results,required"`
	TotalLogs  float64                                                                 `json:"total_logs,required"`
	JSON       accountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultJSON     `json:"-"`
}

// accountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultJSON contains the
// JSON metadata for the struct
// [AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResult]
type accountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Datasets    apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Processed   apijson.Field
	Results     apijson.Field
	TotalLogs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDataset struct {
	ID         string                                                                          `json:"id,required"`
	AccountID  string                                                                          `json:"account_id,required"`
	AccountTag string                                                                          `json:"account_tag,required"`
	CreatedAt  time.Time                                                                       `json:"created_at,required" format:"date-time"`
	Enable     bool                                                                            `json:"enable,required"`
	Filters    []AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFilter `json:"filters,required"`
	// gateway id
	GatewayID  string                                                                     `json:"gateway_id,required"`
	ModifiedAt time.Time                                                                  `json:"modified_at,required" format:"date-time"`
	Name       string                                                                     `json:"name,required"`
	JSON       accountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetJSON `json:"-"`
}

// accountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetJSON
// contains the JSON metadata for the struct
// [AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDataset]
type accountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetJSON struct {
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

func (r *AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFilter struct {
	Key      AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKey          `json:"key,required"`
	Operator AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersOperator     `json:"operator,required"`
	Value    []AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersValueUnion `json:"value,required"`
	JSON     accountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFilterJSON          `json:"-"`
}

// accountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFilterJSON
// contains the JSON metadata for the struct
// [AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFilter]
type accountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFilterJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFilterJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKey string

const (
	AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyCreatedAt           AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKey = "created_at"
	AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyRequestContentType  AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKey = "request_content_type"
	AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyResponseContentType AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKey = "response_content_type"
	AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeySuccess             AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKey = "success"
	AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyCached              AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKey = "cached"
	AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyProvider            AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKey = "provider"
	AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyModel               AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKey = "model"
	AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyCost                AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKey = "cost"
	AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyTokens              AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKey = "tokens"
	AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyTokensIn            AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKey = "tokens_in"
	AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyTokensOut           AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKey = "tokens_out"
	AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyDuration            AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKey = "duration"
	AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyFeedback            AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKey = "feedback"
)

func (r AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKey) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyCreatedAt, AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyRequestContentType, AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyResponseContentType, AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeySuccess, AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyCached, AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyProvider, AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyModel, AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyCost, AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyTokens, AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyTokensIn, AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyTokensOut, AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyDuration, AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersKeyFeedback:
		return true
	}
	return false
}

type AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersOperator string

const (
	AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersOperatorEq       AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersOperator = "eq"
	AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersOperatorContains AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersOperator = "contains"
	AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersOperatorLt       AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersOperator = "lt"
	AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersOperatorGt       AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersOperator = "gt"
)

func (r AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersOperator) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersOperatorEq, AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersOperatorContains, AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersOperatorLt, AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersOperatorGt:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersValueUnion interface {
	ImplementsAccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersValueUnion)(nil)).Elem(),
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

type AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultResult struct {
	ID                string                                                                    `json:"id,required"`
	CreatedAt         time.Time                                                                 `json:"created_at,required" format:"date-time"`
	EvaluationID      string                                                                    `json:"evaluation_id,required"`
	EvaluationTypeID  string                                                                    `json:"evaluation_type_id,required"`
	ModifiedAt        time.Time                                                                 `json:"modified_at,required" format:"date-time"`
	Result            string                                                                    `json:"result,required"`
	Status            float64                                                                   `json:"status,required"`
	StatusDescription string                                                                    `json:"status_description,required"`
	TotalLogs         float64                                                                   `json:"total_logs,required"`
	JSON              accountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultResultJSON `json:"-"`
}

// accountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultResultJSON
// contains the JSON metadata for the struct
// [AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultResult]
type accountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultResultJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	EvaluationID      apijson.Field
	EvaluationTypeID  apijson.Field
	ModifiedAt        apijson.Field
	Result            apijson.Field
	Status            apijson.Field
	StatusDescription apijson.Field
	TotalLogs         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationFetchEvaluationResponse struct {
	Result  AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResult `json:"result,required"`
	Success bool                                                           `json:"success,required"`
	JSON    accountAIGatewayGatewayEvaluationFetchEvaluationResponseJSON   `json:"-"`
}

// accountAIGatewayGatewayEvaluationFetchEvaluationResponseJSON contains the JSON
// metadata for the struct
// [AccountAIGatewayGatewayEvaluationFetchEvaluationResponse]
type accountAIGatewayGatewayEvaluationFetchEvaluationResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayEvaluationFetchEvaluationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationFetchEvaluationResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResult struct {
	ID         string                                                                  `json:"id,required"`
	AccountID  string                                                                  `json:"account_id,required"`
	AccountTag string                                                                  `json:"account_tag,required"`
	CreatedAt  time.Time                                                               `json:"created_at,required" format:"date-time"`
	Datasets   []AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDataset `json:"datasets,required"`
	// gateway id
	GatewayID  string                                                                 `json:"gateway_id,required"`
	ModifiedAt time.Time                                                              `json:"modified_at,required" format:"date-time"`
	Name       string                                                                 `json:"name,required"`
	Processed  bool                                                                   `json:"processed,required"`
	Results    []AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultResult `json:"results,required"`
	TotalLogs  float64                                                                `json:"total_logs,required"`
	JSON       accountAIGatewayGatewayEvaluationFetchEvaluationResponseResultJSON     `json:"-"`
}

// accountAIGatewayGatewayEvaluationFetchEvaluationResponseResultJSON contains the
// JSON metadata for the struct
// [AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResult]
type accountAIGatewayGatewayEvaluationFetchEvaluationResponseResultJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Datasets    apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Processed   apijson.Field
	Results     apijson.Field
	TotalLogs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationFetchEvaluationResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDataset struct {
	ID         string                                                                         `json:"id,required"`
	AccountID  string                                                                         `json:"account_id,required"`
	AccountTag string                                                                         `json:"account_tag,required"`
	CreatedAt  time.Time                                                                      `json:"created_at,required" format:"date-time"`
	Enable     bool                                                                           `json:"enable,required"`
	Filters    []AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFilter `json:"filters,required"`
	// gateway id
	GatewayID  string                                                                    `json:"gateway_id,required"`
	ModifiedAt time.Time                                                                 `json:"modified_at,required" format:"date-time"`
	Name       string                                                                    `json:"name,required"`
	JSON       accountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetJSON `json:"-"`
}

// accountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetJSON
// contains the JSON metadata for the struct
// [AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDataset]
type accountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetJSON struct {
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

func (r *AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFilter struct {
	Key      AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKey          `json:"key,required"`
	Operator AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersOperator     `json:"operator,required"`
	Value    []AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersValueUnion `json:"value,required"`
	JSON     accountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFilterJSON          `json:"-"`
}

// accountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFilterJSON
// contains the JSON metadata for the struct
// [AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFilter]
type accountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFilterJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFilterJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKey string

const (
	AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyCreatedAt           AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKey = "created_at"
	AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyRequestContentType  AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKey = "request_content_type"
	AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyResponseContentType AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKey = "response_content_type"
	AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeySuccess             AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKey = "success"
	AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyCached              AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKey = "cached"
	AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyProvider            AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKey = "provider"
	AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyModel               AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKey = "model"
	AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyCost                AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKey = "cost"
	AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyTokens              AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKey = "tokens"
	AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyTokensIn            AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKey = "tokens_in"
	AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyTokensOut           AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKey = "tokens_out"
	AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyDuration            AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKey = "duration"
	AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyFeedback            AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKey = "feedback"
)

func (r AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKey) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyCreatedAt, AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyRequestContentType, AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyResponseContentType, AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeySuccess, AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyCached, AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyProvider, AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyModel, AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyCost, AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyTokens, AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyTokensIn, AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyTokensOut, AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyDuration, AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersKeyFeedback:
		return true
	}
	return false
}

type AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersOperator string

const (
	AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersOperatorEq       AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersOperator = "eq"
	AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersOperatorContains AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersOperator = "contains"
	AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersOperatorLt       AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersOperator = "lt"
	AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersOperatorGt       AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersOperator = "gt"
)

func (r AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersOperator) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersOperatorEq, AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersOperatorContains, AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersOperatorLt, AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersOperatorGt:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersValueUnion interface {
	ImplementsAccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersValueUnion)(nil)).Elem(),
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

type AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultResult struct {
	ID                string                                                                   `json:"id,required"`
	CreatedAt         time.Time                                                                `json:"created_at,required" format:"date-time"`
	EvaluationID      string                                                                   `json:"evaluation_id,required"`
	EvaluationTypeID  string                                                                   `json:"evaluation_type_id,required"`
	ModifiedAt        time.Time                                                                `json:"modified_at,required" format:"date-time"`
	Result            string                                                                   `json:"result,required"`
	Status            float64                                                                  `json:"status,required"`
	StatusDescription string                                                                   `json:"status_description,required"`
	TotalLogs         float64                                                                  `json:"total_logs,required"`
	JSON              accountAIGatewayGatewayEvaluationFetchEvaluationResponseResultResultJSON `json:"-"`
}

// accountAIGatewayGatewayEvaluationFetchEvaluationResponseResultResultJSON
// contains the JSON metadata for the struct
// [AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultResult]
type accountAIGatewayGatewayEvaluationFetchEvaluationResponseResultResultJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	EvaluationID      apijson.Field
	EvaluationTypeID  apijson.Field
	ModifiedAt        apijson.Field
	Result            apijson.Field
	Status            apijson.Field
	StatusDescription apijson.Field
	TotalLogs         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationFetchEvaluationResponseResultResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationListEvaluationsResponse struct {
	Result  []AccountAIGatewayGatewayEvaluationListEvaluationsResponseResult `json:"result,required"`
	Success bool                                                             `json:"success,required"`
	JSON    accountAIGatewayGatewayEvaluationListEvaluationsResponseJSON     `json:"-"`
}

// accountAIGatewayGatewayEvaluationListEvaluationsResponseJSON contains the JSON
// metadata for the struct
// [AccountAIGatewayGatewayEvaluationListEvaluationsResponse]
type accountAIGatewayGatewayEvaluationListEvaluationsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayEvaluationListEvaluationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationListEvaluationsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationListEvaluationsResponseResult struct {
	ID         string                                                                  `json:"id,required"`
	AccountID  string                                                                  `json:"account_id,required"`
	AccountTag string                                                                  `json:"account_tag,required"`
	CreatedAt  time.Time                                                               `json:"created_at,required" format:"date-time"`
	Datasets   []AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDataset `json:"datasets,required"`
	// gateway id
	GatewayID  string                                                                 `json:"gateway_id,required"`
	ModifiedAt time.Time                                                              `json:"modified_at,required" format:"date-time"`
	Name       string                                                                 `json:"name,required"`
	Processed  bool                                                                   `json:"processed,required"`
	Results    []AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultResult `json:"results,required"`
	TotalLogs  float64                                                                `json:"total_logs,required"`
	JSON       accountAIGatewayGatewayEvaluationListEvaluationsResponseResultJSON     `json:"-"`
}

// accountAIGatewayGatewayEvaluationListEvaluationsResponseResultJSON contains the
// JSON metadata for the struct
// [AccountAIGatewayGatewayEvaluationListEvaluationsResponseResult]
type accountAIGatewayGatewayEvaluationListEvaluationsResponseResultJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Datasets    apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Processed   apijson.Field
	Results     apijson.Field
	TotalLogs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayEvaluationListEvaluationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationListEvaluationsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDataset struct {
	ID         string                                                                         `json:"id,required"`
	AccountID  string                                                                         `json:"account_id,required"`
	AccountTag string                                                                         `json:"account_tag,required"`
	CreatedAt  time.Time                                                                      `json:"created_at,required" format:"date-time"`
	Enable     bool                                                                           `json:"enable,required"`
	Filters    []AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFilter `json:"filters,required"`
	// gateway id
	GatewayID  string                                                                    `json:"gateway_id,required"`
	ModifiedAt time.Time                                                                 `json:"modified_at,required" format:"date-time"`
	Name       string                                                                    `json:"name,required"`
	JSON       accountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetJSON `json:"-"`
}

// accountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetJSON
// contains the JSON metadata for the struct
// [AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDataset]
type accountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetJSON struct {
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

func (r *AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFilter struct {
	Key      AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKey          `json:"key,required"`
	Operator AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersOperator     `json:"operator,required"`
	Value    []AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersValueUnion `json:"value,required"`
	JSON     accountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFilterJSON          `json:"-"`
}

// accountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFilterJSON
// contains the JSON metadata for the struct
// [AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFilter]
type accountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFilterJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFilterJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKey string

const (
	AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyCreatedAt           AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKey = "created_at"
	AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyRequestContentType  AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKey = "request_content_type"
	AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyResponseContentType AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKey = "response_content_type"
	AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeySuccess             AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKey = "success"
	AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyCached              AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKey = "cached"
	AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyProvider            AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKey = "provider"
	AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyModel               AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKey = "model"
	AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyCost                AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKey = "cost"
	AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyTokens              AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKey = "tokens"
	AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyTokensIn            AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKey = "tokens_in"
	AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyTokensOut           AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKey = "tokens_out"
	AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyDuration            AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKey = "duration"
	AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyFeedback            AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKey = "feedback"
)

func (r AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKey) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyCreatedAt, AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyRequestContentType, AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyResponseContentType, AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeySuccess, AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyCached, AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyProvider, AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyModel, AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyCost, AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyTokens, AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyTokensIn, AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyTokensOut, AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyDuration, AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersKeyFeedback:
		return true
	}
	return false
}

type AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersOperator string

const (
	AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersOperatorEq       AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersOperator = "eq"
	AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersOperatorContains AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersOperator = "contains"
	AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersOperatorLt       AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersOperator = "lt"
	AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersOperatorGt       AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersOperator = "gt"
)

func (r AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersOperator) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersOperatorEq, AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersOperatorContains, AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersOperatorLt, AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersOperatorGt:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersValueUnion interface {
	ImplementsAccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersValueUnion)(nil)).Elem(),
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

type AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultResult struct {
	ID                string                                                                   `json:"id,required"`
	CreatedAt         time.Time                                                                `json:"created_at,required" format:"date-time"`
	EvaluationID      string                                                                   `json:"evaluation_id,required"`
	EvaluationTypeID  string                                                                   `json:"evaluation_type_id,required"`
	ModifiedAt        time.Time                                                                `json:"modified_at,required" format:"date-time"`
	Result            string                                                                   `json:"result,required"`
	Status            float64                                                                  `json:"status,required"`
	StatusDescription string                                                                   `json:"status_description,required"`
	TotalLogs         float64                                                                  `json:"total_logs,required"`
	JSON              accountAIGatewayGatewayEvaluationListEvaluationsResponseResultResultJSON `json:"-"`
}

// accountAIGatewayGatewayEvaluationListEvaluationsResponseResultResultJSON
// contains the JSON metadata for the struct
// [AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultResult]
type accountAIGatewayGatewayEvaluationListEvaluationsResponseResultResultJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	EvaluationID      apijson.Field
	EvaluationTypeID  apijson.Field
	ModifiedAt        apijson.Field
	Result            apijson.Field
	Status            apijson.Field
	StatusDescription apijson.Field
	TotalLogs         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayEvaluationListEvaluationsResponseResultResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayEvaluationListEvaluationsResponseResultResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayEvaluationNewEvaluationParams struct {
	DatasetIDs        param.Field[[]string] `json:"dataset_ids,required"`
	EvaluationTypeIDs param.Field[[]string] `json:"evaluation_type_ids,required"`
	Name              param.Field[string]   `json:"name,required"`
}

func (r AccountAIGatewayGatewayEvaluationNewEvaluationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIGatewayGatewayEvaluationListEvaluationsParams struct {
	Name      param.Field[string] `query:"name"`
	Page      param.Field[int64]  `query:"page"`
	PerPage   param.Field[int64]  `query:"per_page"`
	Processed param.Field[bool]   `query:"processed"`
	// Search by id, name
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccountAIGatewayGatewayEvaluationListEvaluationsParams]'s
// query parameters as `url.Values`.
func (r AccountAIGatewayGatewayEvaluationListEvaluationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
