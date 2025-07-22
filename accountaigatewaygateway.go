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

// AccountAIGatewayGatewayService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIGatewayGatewayService] method instead.
type AccountAIGatewayGatewayService struct {
	Options     []option.RequestOption
	Datasets    *AccountAIGatewayGatewayDatasetService
	Evaluations *AccountAIGatewayGatewayEvaluationService
	Logs        *AccountAIGatewayGatewayLogService
}

// NewAccountAIGatewayGatewayService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIGatewayGatewayService(opts ...option.RequestOption) (r *AccountAIGatewayGatewayService) {
	r = &AccountAIGatewayGatewayService{}
	r.Options = opts
	r.Datasets = NewAccountAIGatewayGatewayDatasetService(opts...)
	r.Evaluations = NewAccountAIGatewayGatewayEvaluationService(opts...)
	r.Logs = NewAccountAIGatewayGatewayLogService(opts...)
	return
}

// Create a new Gateway
func (r *AccountAIGatewayGatewayService) NewGateway(ctx context.Context, accountID string, body AccountAIGatewayGatewayNewGatewayParams, opts ...option.RequestOption) (res *AccountAIGatewayGatewayNewGatewayResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a Gateway
func (r *AccountAIGatewayGatewayService) DeleteGateway(ctx context.Context, accountID string, id string, opts ...option.RequestOption) (res *AccountAIGatewayGatewayDeleteGatewayResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s", accountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Fetch a Gateway
func (r *AccountAIGatewayGatewayService) FetchGateway(ctx context.Context, accountID string, id string, opts ...option.RequestOption) (res *AccountAIGatewayGatewayFetchGatewayResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s", accountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get Gateway URL
func (r *AccountAIGatewayGatewayService) GetGatewayURL(ctx context.Context, accountID string, gatewayID string, provider string, opts ...option.RequestOption) (res *AccountAIGatewayGatewayGetGatewayURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if provider == "" {
		err = errors.New("missing required provider parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/url/%s", accountID, gatewayID, provider)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Gateways
func (r *AccountAIGatewayGatewayService) ListGateways(ctx context.Context, accountID string, query AccountAIGatewayGatewayListGatewaysParams, opts ...option.RequestOption) (res *AccountAIGatewayGatewayListGatewaysResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update a Gateway
func (r *AccountAIGatewayGatewayService) UpdateGateway(ctx context.Context, accountID string, id string, body AccountAIGatewayGatewayUpdateGatewayParams, opts ...option.RequestOption) (res *AccountAIGatewayGatewayUpdateGatewayResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s", accountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountAIGatewayGatewayNewGatewayResponse struct {
	Result  AccountAIGatewayGatewayNewGatewayResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    accountAIGatewayGatewayNewGatewayResponseJSON   `json:"-"`
}

// accountAIGatewayGatewayNewGatewayResponseJSON contains the JSON metadata for the
// struct [AccountAIGatewayGatewayNewGatewayResponse]
type accountAIGatewayGatewayNewGatewayResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayNewGatewayResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayNewGatewayResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayNewGatewayResponseResult struct {
	// gateway id
	ID                      string                                                               `json:"id,required"`
	AccountID               string                                                               `json:"account_id,required"`
	AccountTag              string                                                               `json:"account_tag,required"`
	CacheInvalidateOnUpdate bool                                                                 `json:"cache_invalidate_on_update,required"`
	CacheTtl                int64                                                                `json:"cache_ttl,required,nullable"`
	CollectLogs             bool                                                                 `json:"collect_logs,required"`
	CreatedAt               time.Time                                                            `json:"created_at,required" format:"date-time"`
	InternalID              string                                                               `json:"internal_id,required" format:"uuid"`
	ModifiedAt              time.Time                                                            `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                                                                `json:"rate_limiting_interval,required,nullable"`
	RateLimitingLimit       int64                                                                `json:"rate_limiting_limit,required,nullable"`
	RateLimitingTechnique   AccountAIGatewayGatewayNewGatewayResponseResultRateLimitingTechnique `json:"rate_limiting_technique,required"`
	Authentication          bool                                                                 `json:"authentication"`
	LogManagement           int64                                                                `json:"log_management,nullable"`
	LogManagementStrategy   AccountAIGatewayGatewayNewGatewayResponseResultLogManagementStrategy `json:"log_management_strategy,nullable"`
	Logpush                 bool                                                                 `json:"logpush"`
	LogpushPublicKey        string                                                               `json:"logpush_public_key,nullable"`
	JSON                    accountAIGatewayGatewayNewGatewayResponseResultJSON                  `json:"-"`
}

// accountAIGatewayGatewayNewGatewayResponseResultJSON contains the JSON metadata
// for the struct [AccountAIGatewayGatewayNewGatewayResponseResult]
type accountAIGatewayGatewayNewGatewayResponseResultJSON struct {
	ID                      apijson.Field
	AccountID               apijson.Field
	AccountTag              apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTtl                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	InternalID              apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	Authentication          apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayNewGatewayResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayNewGatewayResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayNewGatewayResponseResultRateLimitingTechnique string

const (
	AccountAIGatewayGatewayNewGatewayResponseResultRateLimitingTechniqueFixed   AccountAIGatewayGatewayNewGatewayResponseResultRateLimitingTechnique = "fixed"
	AccountAIGatewayGatewayNewGatewayResponseResultRateLimitingTechniqueSliding AccountAIGatewayGatewayNewGatewayResponseResultRateLimitingTechnique = "sliding"
)

func (r AccountAIGatewayGatewayNewGatewayResponseResultRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayNewGatewayResponseResultRateLimitingTechniqueFixed, AccountAIGatewayGatewayNewGatewayResponseResultRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AccountAIGatewayGatewayNewGatewayResponseResultLogManagementStrategy string

const (
	AccountAIGatewayGatewayNewGatewayResponseResultLogManagementStrategyStopInserting AccountAIGatewayGatewayNewGatewayResponseResultLogManagementStrategy = "STOP_INSERTING"
	AccountAIGatewayGatewayNewGatewayResponseResultLogManagementStrategyDeleteOldest  AccountAIGatewayGatewayNewGatewayResponseResultLogManagementStrategy = "DELETE_OLDEST"
)

func (r AccountAIGatewayGatewayNewGatewayResponseResultLogManagementStrategy) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayNewGatewayResponseResultLogManagementStrategyStopInserting, AccountAIGatewayGatewayNewGatewayResponseResultLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AccountAIGatewayGatewayDeleteGatewayResponse struct {
	Result  AccountAIGatewayGatewayDeleteGatewayResponseResult `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    accountAIGatewayGatewayDeleteGatewayResponseJSON   `json:"-"`
}

// accountAIGatewayGatewayDeleteGatewayResponseJSON contains the JSON metadata for
// the struct [AccountAIGatewayGatewayDeleteGatewayResponse]
type accountAIGatewayGatewayDeleteGatewayResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayDeleteGatewayResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayDeleteGatewayResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayDeleteGatewayResponseResult struct {
	// gateway id
	ID                      string                                                                  `json:"id,required"`
	AccountID               string                                                                  `json:"account_id,required"`
	AccountTag              string                                                                  `json:"account_tag,required"`
	CacheInvalidateOnUpdate bool                                                                    `json:"cache_invalidate_on_update,required"`
	CacheTtl                int64                                                                   `json:"cache_ttl,required,nullable"`
	CollectLogs             bool                                                                    `json:"collect_logs,required"`
	CreatedAt               time.Time                                                               `json:"created_at,required" format:"date-time"`
	InternalID              string                                                                  `json:"internal_id,required" format:"uuid"`
	ModifiedAt              time.Time                                                               `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                                                                   `json:"rate_limiting_interval,required,nullable"`
	RateLimitingLimit       int64                                                                   `json:"rate_limiting_limit,required,nullable"`
	RateLimitingTechnique   AccountAIGatewayGatewayDeleteGatewayResponseResultRateLimitingTechnique `json:"rate_limiting_technique,required"`
	Authentication          bool                                                                    `json:"authentication"`
	LogManagement           int64                                                                   `json:"log_management,nullable"`
	LogManagementStrategy   AccountAIGatewayGatewayDeleteGatewayResponseResultLogManagementStrategy `json:"log_management_strategy,nullable"`
	Logpush                 bool                                                                    `json:"logpush"`
	LogpushPublicKey        string                                                                  `json:"logpush_public_key,nullable"`
	JSON                    accountAIGatewayGatewayDeleteGatewayResponseResultJSON                  `json:"-"`
}

// accountAIGatewayGatewayDeleteGatewayResponseResultJSON contains the JSON
// metadata for the struct [AccountAIGatewayGatewayDeleteGatewayResponseResult]
type accountAIGatewayGatewayDeleteGatewayResponseResultJSON struct {
	ID                      apijson.Field
	AccountID               apijson.Field
	AccountTag              apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTtl                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	InternalID              apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	Authentication          apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayDeleteGatewayResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayDeleteGatewayResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayDeleteGatewayResponseResultRateLimitingTechnique string

const (
	AccountAIGatewayGatewayDeleteGatewayResponseResultRateLimitingTechniqueFixed   AccountAIGatewayGatewayDeleteGatewayResponseResultRateLimitingTechnique = "fixed"
	AccountAIGatewayGatewayDeleteGatewayResponseResultRateLimitingTechniqueSliding AccountAIGatewayGatewayDeleteGatewayResponseResultRateLimitingTechnique = "sliding"
)

func (r AccountAIGatewayGatewayDeleteGatewayResponseResultRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayDeleteGatewayResponseResultRateLimitingTechniqueFixed, AccountAIGatewayGatewayDeleteGatewayResponseResultRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AccountAIGatewayGatewayDeleteGatewayResponseResultLogManagementStrategy string

const (
	AccountAIGatewayGatewayDeleteGatewayResponseResultLogManagementStrategyStopInserting AccountAIGatewayGatewayDeleteGatewayResponseResultLogManagementStrategy = "STOP_INSERTING"
	AccountAIGatewayGatewayDeleteGatewayResponseResultLogManagementStrategyDeleteOldest  AccountAIGatewayGatewayDeleteGatewayResponseResultLogManagementStrategy = "DELETE_OLDEST"
)

func (r AccountAIGatewayGatewayDeleteGatewayResponseResultLogManagementStrategy) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayDeleteGatewayResponseResultLogManagementStrategyStopInserting, AccountAIGatewayGatewayDeleteGatewayResponseResultLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AccountAIGatewayGatewayFetchGatewayResponse struct {
	Result  AccountAIGatewayGatewayFetchGatewayResponseResult `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    accountAIGatewayGatewayFetchGatewayResponseJSON   `json:"-"`
}

// accountAIGatewayGatewayFetchGatewayResponseJSON contains the JSON metadata for
// the struct [AccountAIGatewayGatewayFetchGatewayResponse]
type accountAIGatewayGatewayFetchGatewayResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayFetchGatewayResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayFetchGatewayResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayFetchGatewayResponseResult struct {
	// gateway id
	ID                      string                                                                 `json:"id,required"`
	AccountID               string                                                                 `json:"account_id,required"`
	AccountTag              string                                                                 `json:"account_tag,required"`
	CacheInvalidateOnUpdate bool                                                                   `json:"cache_invalidate_on_update,required"`
	CacheTtl                int64                                                                  `json:"cache_ttl,required,nullable"`
	CollectLogs             bool                                                                   `json:"collect_logs,required"`
	CreatedAt               time.Time                                                              `json:"created_at,required" format:"date-time"`
	InternalID              string                                                                 `json:"internal_id,required" format:"uuid"`
	ModifiedAt              time.Time                                                              `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                                                                  `json:"rate_limiting_interval,required,nullable"`
	RateLimitingLimit       int64                                                                  `json:"rate_limiting_limit,required,nullable"`
	RateLimitingTechnique   AccountAIGatewayGatewayFetchGatewayResponseResultRateLimitingTechnique `json:"rate_limiting_technique,required"`
	Authentication          bool                                                                   `json:"authentication"`
	LogManagement           int64                                                                  `json:"log_management,nullable"`
	LogManagementStrategy   AccountAIGatewayGatewayFetchGatewayResponseResultLogManagementStrategy `json:"log_management_strategy,nullable"`
	Logpush                 bool                                                                   `json:"logpush"`
	LogpushPublicKey        string                                                                 `json:"logpush_public_key,nullable"`
	JSON                    accountAIGatewayGatewayFetchGatewayResponseResultJSON                  `json:"-"`
}

// accountAIGatewayGatewayFetchGatewayResponseResultJSON contains the JSON metadata
// for the struct [AccountAIGatewayGatewayFetchGatewayResponseResult]
type accountAIGatewayGatewayFetchGatewayResponseResultJSON struct {
	ID                      apijson.Field
	AccountID               apijson.Field
	AccountTag              apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTtl                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	InternalID              apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	Authentication          apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayFetchGatewayResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayFetchGatewayResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayFetchGatewayResponseResultRateLimitingTechnique string

const (
	AccountAIGatewayGatewayFetchGatewayResponseResultRateLimitingTechniqueFixed   AccountAIGatewayGatewayFetchGatewayResponseResultRateLimitingTechnique = "fixed"
	AccountAIGatewayGatewayFetchGatewayResponseResultRateLimitingTechniqueSliding AccountAIGatewayGatewayFetchGatewayResponseResultRateLimitingTechnique = "sliding"
)

func (r AccountAIGatewayGatewayFetchGatewayResponseResultRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayFetchGatewayResponseResultRateLimitingTechniqueFixed, AccountAIGatewayGatewayFetchGatewayResponseResultRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AccountAIGatewayGatewayFetchGatewayResponseResultLogManagementStrategy string

const (
	AccountAIGatewayGatewayFetchGatewayResponseResultLogManagementStrategyStopInserting AccountAIGatewayGatewayFetchGatewayResponseResultLogManagementStrategy = "STOP_INSERTING"
	AccountAIGatewayGatewayFetchGatewayResponseResultLogManagementStrategyDeleteOldest  AccountAIGatewayGatewayFetchGatewayResponseResultLogManagementStrategy = "DELETE_OLDEST"
)

func (r AccountAIGatewayGatewayFetchGatewayResponseResultLogManagementStrategy) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayFetchGatewayResponseResultLogManagementStrategyStopInserting, AccountAIGatewayGatewayFetchGatewayResponseResultLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AccountAIGatewayGatewayGetGatewayURLResponse struct {
	Result  string                                           `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    accountAIGatewayGatewayGetGatewayURLResponseJSON `json:"-"`
}

// accountAIGatewayGatewayGetGatewayURLResponseJSON contains the JSON metadata for
// the struct [AccountAIGatewayGatewayGetGatewayURLResponse]
type accountAIGatewayGatewayGetGatewayURLResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayGetGatewayURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayGetGatewayURLResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayListGatewaysResponse struct {
	Result  []AccountAIGatewayGatewayListGatewaysResponseResult `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    accountAIGatewayGatewayListGatewaysResponseJSON     `json:"-"`
}

// accountAIGatewayGatewayListGatewaysResponseJSON contains the JSON metadata for
// the struct [AccountAIGatewayGatewayListGatewaysResponse]
type accountAIGatewayGatewayListGatewaysResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayListGatewaysResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayListGatewaysResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayListGatewaysResponseResult struct {
	// gateway id
	ID                      string                                                                 `json:"id,required"`
	AccountID               string                                                                 `json:"account_id,required"`
	AccountTag              string                                                                 `json:"account_tag,required"`
	CacheInvalidateOnUpdate bool                                                                   `json:"cache_invalidate_on_update,required"`
	CacheTtl                int64                                                                  `json:"cache_ttl,required,nullable"`
	CollectLogs             bool                                                                   `json:"collect_logs,required"`
	CreatedAt               time.Time                                                              `json:"created_at,required" format:"date-time"`
	InternalID              string                                                                 `json:"internal_id,required" format:"uuid"`
	ModifiedAt              time.Time                                                              `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                                                                  `json:"rate_limiting_interval,required,nullable"`
	RateLimitingLimit       int64                                                                  `json:"rate_limiting_limit,required,nullable"`
	RateLimitingTechnique   AccountAIGatewayGatewayListGatewaysResponseResultRateLimitingTechnique `json:"rate_limiting_technique,required"`
	Authentication          bool                                                                   `json:"authentication"`
	LogManagement           int64                                                                  `json:"log_management,nullable"`
	LogManagementStrategy   AccountAIGatewayGatewayListGatewaysResponseResultLogManagementStrategy `json:"log_management_strategy,nullable"`
	Logpush                 bool                                                                   `json:"logpush"`
	LogpushPublicKey        string                                                                 `json:"logpush_public_key,nullable"`
	JSON                    accountAIGatewayGatewayListGatewaysResponseResultJSON                  `json:"-"`
}

// accountAIGatewayGatewayListGatewaysResponseResultJSON contains the JSON metadata
// for the struct [AccountAIGatewayGatewayListGatewaysResponseResult]
type accountAIGatewayGatewayListGatewaysResponseResultJSON struct {
	ID                      apijson.Field
	AccountID               apijson.Field
	AccountTag              apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTtl                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	InternalID              apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	Authentication          apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayListGatewaysResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayListGatewaysResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayListGatewaysResponseResultRateLimitingTechnique string

const (
	AccountAIGatewayGatewayListGatewaysResponseResultRateLimitingTechniqueFixed   AccountAIGatewayGatewayListGatewaysResponseResultRateLimitingTechnique = "fixed"
	AccountAIGatewayGatewayListGatewaysResponseResultRateLimitingTechniqueSliding AccountAIGatewayGatewayListGatewaysResponseResultRateLimitingTechnique = "sliding"
)

func (r AccountAIGatewayGatewayListGatewaysResponseResultRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayListGatewaysResponseResultRateLimitingTechniqueFixed, AccountAIGatewayGatewayListGatewaysResponseResultRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AccountAIGatewayGatewayListGatewaysResponseResultLogManagementStrategy string

const (
	AccountAIGatewayGatewayListGatewaysResponseResultLogManagementStrategyStopInserting AccountAIGatewayGatewayListGatewaysResponseResultLogManagementStrategy = "STOP_INSERTING"
	AccountAIGatewayGatewayListGatewaysResponseResultLogManagementStrategyDeleteOldest  AccountAIGatewayGatewayListGatewaysResponseResultLogManagementStrategy = "DELETE_OLDEST"
)

func (r AccountAIGatewayGatewayListGatewaysResponseResultLogManagementStrategy) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayListGatewaysResponseResultLogManagementStrategyStopInserting, AccountAIGatewayGatewayListGatewaysResponseResultLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AccountAIGatewayGatewayUpdateGatewayResponse struct {
	Result  AccountAIGatewayGatewayUpdateGatewayResponseResult `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    accountAIGatewayGatewayUpdateGatewayResponseJSON   `json:"-"`
}

// accountAIGatewayGatewayUpdateGatewayResponseJSON contains the JSON metadata for
// the struct [AccountAIGatewayGatewayUpdateGatewayResponse]
type accountAIGatewayGatewayUpdateGatewayResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayUpdateGatewayResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayUpdateGatewayResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayUpdateGatewayResponseResult struct {
	// gateway id
	ID                      string                                                                  `json:"id,required"`
	AccountID               string                                                                  `json:"account_id,required"`
	AccountTag              string                                                                  `json:"account_tag,required"`
	CacheInvalidateOnUpdate bool                                                                    `json:"cache_invalidate_on_update,required"`
	CacheTtl                int64                                                                   `json:"cache_ttl,required,nullable"`
	CollectLogs             bool                                                                    `json:"collect_logs,required"`
	CreatedAt               time.Time                                                               `json:"created_at,required" format:"date-time"`
	InternalID              string                                                                  `json:"internal_id,required" format:"uuid"`
	ModifiedAt              time.Time                                                               `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                                                                   `json:"rate_limiting_interval,required,nullable"`
	RateLimitingLimit       int64                                                                   `json:"rate_limiting_limit,required,nullable"`
	RateLimitingTechnique   AccountAIGatewayGatewayUpdateGatewayResponseResultRateLimitingTechnique `json:"rate_limiting_technique,required"`
	Authentication          bool                                                                    `json:"authentication"`
	LogManagement           int64                                                                   `json:"log_management,nullable"`
	LogManagementStrategy   AccountAIGatewayGatewayUpdateGatewayResponseResultLogManagementStrategy `json:"log_management_strategy,nullable"`
	Logpush                 bool                                                                    `json:"logpush"`
	LogpushPublicKey        string                                                                  `json:"logpush_public_key,nullable"`
	JSON                    accountAIGatewayGatewayUpdateGatewayResponseResultJSON                  `json:"-"`
}

// accountAIGatewayGatewayUpdateGatewayResponseResultJSON contains the JSON
// metadata for the struct [AccountAIGatewayGatewayUpdateGatewayResponseResult]
type accountAIGatewayGatewayUpdateGatewayResponseResultJSON struct {
	ID                      apijson.Field
	AccountID               apijson.Field
	AccountTag              apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTtl                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	InternalID              apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	Authentication          apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayUpdateGatewayResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayUpdateGatewayResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayUpdateGatewayResponseResultRateLimitingTechnique string

const (
	AccountAIGatewayGatewayUpdateGatewayResponseResultRateLimitingTechniqueFixed   AccountAIGatewayGatewayUpdateGatewayResponseResultRateLimitingTechnique = "fixed"
	AccountAIGatewayGatewayUpdateGatewayResponseResultRateLimitingTechniqueSliding AccountAIGatewayGatewayUpdateGatewayResponseResultRateLimitingTechnique = "sliding"
)

func (r AccountAIGatewayGatewayUpdateGatewayResponseResultRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayUpdateGatewayResponseResultRateLimitingTechniqueFixed, AccountAIGatewayGatewayUpdateGatewayResponseResultRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AccountAIGatewayGatewayUpdateGatewayResponseResultLogManagementStrategy string

const (
	AccountAIGatewayGatewayUpdateGatewayResponseResultLogManagementStrategyStopInserting AccountAIGatewayGatewayUpdateGatewayResponseResultLogManagementStrategy = "STOP_INSERTING"
	AccountAIGatewayGatewayUpdateGatewayResponseResultLogManagementStrategyDeleteOldest  AccountAIGatewayGatewayUpdateGatewayResponseResultLogManagementStrategy = "DELETE_OLDEST"
)

func (r AccountAIGatewayGatewayUpdateGatewayResponseResultLogManagementStrategy) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayUpdateGatewayResponseResultLogManagementStrategyStopInserting, AccountAIGatewayGatewayUpdateGatewayResponseResultLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AccountAIGatewayGatewayNewGatewayParams struct {
	// gateway id
	ID                      param.Field[string]                                                       `json:"id,required"`
	CacheInvalidateOnUpdate param.Field[bool]                                                         `json:"cache_invalidate_on_update,required"`
	CacheTtl                param.Field[int64]                                                        `json:"cache_ttl,required"`
	CollectLogs             param.Field[bool]                                                         `json:"collect_logs,required"`
	RateLimitingInterval    param.Field[int64]                                                        `json:"rate_limiting_interval,required"`
	RateLimitingLimit       param.Field[int64]                                                        `json:"rate_limiting_limit,required"`
	RateLimitingTechnique   param.Field[AccountAIGatewayGatewayNewGatewayParamsRateLimitingTechnique] `json:"rate_limiting_technique,required"`
	Authentication          param.Field[bool]                                                         `json:"authentication"`
	LogManagement           param.Field[int64]                                                        `json:"log_management"`
	LogManagementStrategy   param.Field[AccountAIGatewayGatewayNewGatewayParamsLogManagementStrategy] `json:"log_management_strategy"`
	Logpush                 param.Field[bool]                                                         `json:"logpush"`
	LogpushPublicKey        param.Field[string]                                                       `json:"logpush_public_key"`
}

func (r AccountAIGatewayGatewayNewGatewayParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIGatewayGatewayNewGatewayParamsRateLimitingTechnique string

const (
	AccountAIGatewayGatewayNewGatewayParamsRateLimitingTechniqueFixed   AccountAIGatewayGatewayNewGatewayParamsRateLimitingTechnique = "fixed"
	AccountAIGatewayGatewayNewGatewayParamsRateLimitingTechniqueSliding AccountAIGatewayGatewayNewGatewayParamsRateLimitingTechnique = "sliding"
)

func (r AccountAIGatewayGatewayNewGatewayParamsRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayNewGatewayParamsRateLimitingTechniqueFixed, AccountAIGatewayGatewayNewGatewayParamsRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AccountAIGatewayGatewayNewGatewayParamsLogManagementStrategy string

const (
	AccountAIGatewayGatewayNewGatewayParamsLogManagementStrategyStopInserting AccountAIGatewayGatewayNewGatewayParamsLogManagementStrategy = "STOP_INSERTING"
	AccountAIGatewayGatewayNewGatewayParamsLogManagementStrategyDeleteOldest  AccountAIGatewayGatewayNewGatewayParamsLogManagementStrategy = "DELETE_OLDEST"
)

func (r AccountAIGatewayGatewayNewGatewayParamsLogManagementStrategy) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayNewGatewayParamsLogManagementStrategyStopInserting, AccountAIGatewayGatewayNewGatewayParamsLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AccountAIGatewayGatewayListGatewaysParams struct {
	Page    param.Field[int64] `query:"page"`
	PerPage param.Field[int64] `query:"per_page"`
	// Search by id
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccountAIGatewayGatewayListGatewaysParams]'s query
// parameters as `url.Values`.
func (r AccountAIGatewayGatewayListGatewaysParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIGatewayGatewayUpdateGatewayParams struct {
	CacheInvalidateOnUpdate param.Field[bool]                                                            `json:"cache_invalidate_on_update,required"`
	CacheTtl                param.Field[int64]                                                           `json:"cache_ttl,required"`
	CollectLogs             param.Field[bool]                                                            `json:"collect_logs,required"`
	RateLimitingInterval    param.Field[int64]                                                           `json:"rate_limiting_interval,required"`
	RateLimitingLimit       param.Field[int64]                                                           `json:"rate_limiting_limit,required"`
	RateLimitingTechnique   param.Field[AccountAIGatewayGatewayUpdateGatewayParamsRateLimitingTechnique] `json:"rate_limiting_technique,required"`
	Authentication          param.Field[bool]                                                            `json:"authentication"`
	LogManagement           param.Field[int64]                                                           `json:"log_management"`
	LogManagementStrategy   param.Field[AccountAIGatewayGatewayUpdateGatewayParamsLogManagementStrategy] `json:"log_management_strategy"`
	Logpush                 param.Field[bool]                                                            `json:"logpush"`
	LogpushPublicKey        param.Field[string]                                                          `json:"logpush_public_key"`
}

func (r AccountAIGatewayGatewayUpdateGatewayParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIGatewayGatewayUpdateGatewayParamsRateLimitingTechnique string

const (
	AccountAIGatewayGatewayUpdateGatewayParamsRateLimitingTechniqueFixed   AccountAIGatewayGatewayUpdateGatewayParamsRateLimitingTechnique = "fixed"
	AccountAIGatewayGatewayUpdateGatewayParamsRateLimitingTechniqueSliding AccountAIGatewayGatewayUpdateGatewayParamsRateLimitingTechnique = "sliding"
)

func (r AccountAIGatewayGatewayUpdateGatewayParamsRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayUpdateGatewayParamsRateLimitingTechniqueFixed, AccountAIGatewayGatewayUpdateGatewayParamsRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AccountAIGatewayGatewayUpdateGatewayParamsLogManagementStrategy string

const (
	AccountAIGatewayGatewayUpdateGatewayParamsLogManagementStrategyStopInserting AccountAIGatewayGatewayUpdateGatewayParamsLogManagementStrategy = "STOP_INSERTING"
	AccountAIGatewayGatewayUpdateGatewayParamsLogManagementStrategyDeleteOldest  AccountAIGatewayGatewayUpdateGatewayParamsLogManagementStrategy = "DELETE_OLDEST"
)

func (r AccountAIGatewayGatewayUpdateGatewayParamsLogManagementStrategy) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayUpdateGatewayParamsLogManagementStrategyStopInserting, AccountAIGatewayGatewayUpdateGatewayParamsLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}
