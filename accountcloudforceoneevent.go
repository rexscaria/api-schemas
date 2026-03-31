// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountCloudforceOneEventService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCloudforceOneEventService] method instead.
type AccountCloudforceOneEventService struct {
	Options    []option.RequestOption
	Insight    *AccountCloudforceOneEventInsightService
	Raw        *AccountCloudforceOneEventRawService
	Categories *AccountCloudforceOneEventCategoryService
	Cron       *AccountCloudforceOneEventCronService
	Dataset    *AccountCloudforceOneEventDatasetService
	EventTag   *AccountCloudforceOneEventEventTagService
	Relate     *AccountCloudforceOneEventRelateService
	Tags       *AccountCloudforceOneEventTagService
}

// NewAccountCloudforceOneEventService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCloudforceOneEventService(opts ...option.RequestOption) (r *AccountCloudforceOneEventService) {
	r = &AccountCloudforceOneEventService{}
	r.Options = opts
	r.Insight = NewAccountCloudforceOneEventInsightService(opts...)
	r.Raw = NewAccountCloudforceOneEventRawService(opts...)
	r.Categories = NewAccountCloudforceOneEventCategoryService(opts...)
	r.Cron = NewAccountCloudforceOneEventCronService(opts...)
	r.Dataset = NewAccountCloudforceOneEventDatasetService(opts...)
	r.EventTag = NewAccountCloudforceOneEventEventTagService(opts...)
	r.Relate = NewAccountCloudforceOneEventRelateService(opts...)
	r.Tags = NewAccountCloudforceOneEventTagService(opts...)
	return
}

// To create a dataset, see the
// [`Create Dataset`](https://developers.cloudflare.com/api/resources/cloudforce_one/subresources/threat_events/subresources/datasets/methods/create/)
// endpoint. When `datasetId` parameter is unspecified, it will be created in a
// default dataset named `Cloudforce One Threat Events`.
func (r *AccountCloudforceOneEventService) New(ctx context.Context, accountID string, body AccountCloudforceOneEventNewParams, opts ...option.RequestOption) (res *AccountCloudforceOneEventNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/create", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Reads an event
func (r *AccountCloudforceOneEventService) Get(ctx context.Context, accountID string, eventID string, opts ...option.RequestOption) (res *AccountCloudforceOneEventGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/%s", accountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an event
func (r *AccountCloudforceOneEventService) Update(ctx context.Context, accountID string, eventID string, body AccountCloudforceOneEventUpdateParams, opts ...option.RequestOption) (res *AccountCloudforceOneEventUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/%s", accountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The `datasetId` parameter must be defined. To list existing datasets (and their
// IDs) in your account, use the
// [`List Datasets`](https://developers.cloudflare.com/api/resources/cloudforce_one/subresources/threat_events/subresources/datasets/methods/list/)
// endpoint.
func (r *AccountCloudforceOneEventService) Delete(ctx context.Context, accountID string, eventID string, opts ...option.RequestOption) (res *AccountCloudforceOneEventDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/%s", accountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// The `datasetId` parameter must be defined. To list existing datasets (and their
// IDs) in your account, use the
// [`List Datasets`](https://developers.cloudflare.com/api/resources/cloudforce_one/subresources/threat_events/subresources/datasets/methods/list/)
// endpoint.
func (r *AccountCloudforceOneEventService) NewBulk(ctx context.Context, accountID string, body AccountCloudforceOneEventNewBulkParams, opts ...option.RequestOption) (res *float64, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/create/bulk", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists attackers
func (r *AccountCloudforceOneEventService) ListAttackers(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountCloudforceOneEventListAttackersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/attackers", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves countries information for all countries
func (r *AccountCloudforceOneEventService) ListCountries(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]AccountCloudforceOneEventListCountriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/countries", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists all indicator types
func (r *AccountCloudforceOneEventService) ListIndicatorTypes(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountCloudforceOneEventListIndicatorTypesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/indicatorTypes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists all target industries
func (r *AccountCloudforceOneEventService) ListTargetIndustries(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountCloudforceOneEventListTargetIndustriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/targetIndustries", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AccountCloudforceOneEventNewResponse struct {
	ID              float64                                  `json:"id" api:"required"`
	AccountID       float64                                  `json:"accountId" api:"required"`
	Attacker        string                                   `json:"attacker" api:"required"`
	AttackerCountry string                                   `json:"attackerCountry" api:"required"`
	Category        string                                   `json:"category" api:"required"`
	Date            string                                   `json:"date" api:"required"`
	Event           string                                   `json:"event" api:"required"`
	Indicator       string                                   `json:"indicator" api:"required"`
	IndicatorType   string                                   `json:"indicatorType" api:"required"`
	IndicatorTypeID float64                                  `json:"indicatorTypeId" api:"required"`
	KillChain       float64                                  `json:"killChain" api:"required"`
	MitreAttack     []string                                 `json:"mitreAttack" api:"required"`
	NumReferenced   float64                                  `json:"numReferenced" api:"required"`
	NumReferences   float64                                  `json:"numReferences" api:"required"`
	RawID           string                                   `json:"rawId" api:"required"`
	Referenced      []string                                 `json:"referenced" api:"required"`
	ReferencedIDs   []float64                                `json:"referencedIds" api:"required"`
	References      []string                                 `json:"references" api:"required"`
	ReferencesIDs   []float64                                `json:"referencesIds" api:"required"`
	Tags            []string                                 `json:"tags" api:"required"`
	TargetCountry   string                                   `json:"targetCountry" api:"required"`
	TargetIndustry  string                                   `json:"targetIndustry" api:"required"`
	Tlp             string                                   `json:"tlp" api:"required"`
	Uuid            string                                   `json:"uuid" api:"required"`
	Insight         string                                   `json:"insight"`
	ReleasabilityID string                                   `json:"releasabilityId"`
	JSON            accountCloudforceOneEventNewResponseJSON `json:"-"`
}

// accountCloudforceOneEventNewResponseJSON contains the JSON metadata for the
// struct [AccountCloudforceOneEventNewResponse]
type accountCloudforceOneEventNewResponseJSON struct {
	ID              apijson.Field
	AccountID       apijson.Field
	Attacker        apijson.Field
	AttackerCountry apijson.Field
	Category        apijson.Field
	Date            apijson.Field
	Event           apijson.Field
	Indicator       apijson.Field
	IndicatorType   apijson.Field
	IndicatorTypeID apijson.Field
	KillChain       apijson.Field
	MitreAttack     apijson.Field
	NumReferenced   apijson.Field
	NumReferences   apijson.Field
	RawID           apijson.Field
	Referenced      apijson.Field
	ReferencedIDs   apijson.Field
	References      apijson.Field
	ReferencesIDs   apijson.Field
	Tags            apijson.Field
	TargetCountry   apijson.Field
	TargetIndustry  apijson.Field
	Tlp             apijson.Field
	Uuid            apijson.Field
	Insight         apijson.Field
	ReleasabilityID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountCloudforceOneEventNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventGetResponse struct {
	ID              float64                                  `json:"id" api:"required"`
	AccountID       float64                                  `json:"accountId" api:"required"`
	Attacker        string                                   `json:"attacker" api:"required"`
	AttackerCountry string                                   `json:"attackerCountry" api:"required"`
	Category        string                                   `json:"category" api:"required"`
	Date            string                                   `json:"date" api:"required"`
	Event           string                                   `json:"event" api:"required"`
	Indicator       string                                   `json:"indicator" api:"required"`
	IndicatorType   string                                   `json:"indicatorType" api:"required"`
	IndicatorTypeID float64                                  `json:"indicatorTypeId" api:"required"`
	KillChain       float64                                  `json:"killChain" api:"required"`
	MitreAttack     []string                                 `json:"mitreAttack" api:"required"`
	NumReferenced   float64                                  `json:"numReferenced" api:"required"`
	NumReferences   float64                                  `json:"numReferences" api:"required"`
	RawID           string                                   `json:"rawId" api:"required"`
	Referenced      []string                                 `json:"referenced" api:"required"`
	ReferencedIDs   []float64                                `json:"referencedIds" api:"required"`
	References      []string                                 `json:"references" api:"required"`
	ReferencesIDs   []float64                                `json:"referencesIds" api:"required"`
	Tags            []string                                 `json:"tags" api:"required"`
	TargetCountry   string                                   `json:"targetCountry" api:"required"`
	TargetIndustry  string                                   `json:"targetIndustry" api:"required"`
	Tlp             string                                   `json:"tlp" api:"required"`
	Uuid            string                                   `json:"uuid" api:"required"`
	Insight         string                                   `json:"insight"`
	ReleasabilityID string                                   `json:"releasabilityId"`
	JSON            accountCloudforceOneEventGetResponseJSON `json:"-"`
}

// accountCloudforceOneEventGetResponseJSON contains the JSON metadata for the
// struct [AccountCloudforceOneEventGetResponse]
type accountCloudforceOneEventGetResponseJSON struct {
	ID              apijson.Field
	AccountID       apijson.Field
	Attacker        apijson.Field
	AttackerCountry apijson.Field
	Category        apijson.Field
	Date            apijson.Field
	Event           apijson.Field
	Indicator       apijson.Field
	IndicatorType   apijson.Field
	IndicatorTypeID apijson.Field
	KillChain       apijson.Field
	MitreAttack     apijson.Field
	NumReferenced   apijson.Field
	NumReferences   apijson.Field
	RawID           apijson.Field
	Referenced      apijson.Field
	ReferencedIDs   apijson.Field
	References      apijson.Field
	ReferencesIDs   apijson.Field
	Tags            apijson.Field
	TargetCountry   apijson.Field
	TargetIndustry  apijson.Field
	Tlp             apijson.Field
	Uuid            apijson.Field
	Insight         apijson.Field
	ReleasabilityID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountCloudforceOneEventGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventUpdateResponse struct {
	ID              float64                                     `json:"id" api:"required"`
	AccountID       float64                                     `json:"accountId" api:"required"`
	Attacker        string                                      `json:"attacker" api:"required"`
	AttackerCountry string                                      `json:"attackerCountry" api:"required"`
	Category        string                                      `json:"category" api:"required"`
	Date            string                                      `json:"date" api:"required"`
	Event           string                                      `json:"event" api:"required"`
	Indicator       string                                      `json:"indicator" api:"required"`
	IndicatorType   string                                      `json:"indicatorType" api:"required"`
	IndicatorTypeID float64                                     `json:"indicatorTypeId" api:"required"`
	KillChain       float64                                     `json:"killChain" api:"required"`
	MitreAttack     []string                                    `json:"mitreAttack" api:"required"`
	NumReferenced   float64                                     `json:"numReferenced" api:"required"`
	NumReferences   float64                                     `json:"numReferences" api:"required"`
	RawID           string                                      `json:"rawId" api:"required"`
	Referenced      []string                                    `json:"referenced" api:"required"`
	ReferencedIDs   []float64                                   `json:"referencedIds" api:"required"`
	References      []string                                    `json:"references" api:"required"`
	ReferencesIDs   []float64                                   `json:"referencesIds" api:"required"`
	Tags            []string                                    `json:"tags" api:"required"`
	TargetCountry   string                                      `json:"targetCountry" api:"required"`
	TargetIndustry  string                                      `json:"targetIndustry" api:"required"`
	Tlp             string                                      `json:"tlp" api:"required"`
	Uuid            string                                      `json:"uuid" api:"required"`
	Insight         string                                      `json:"insight"`
	ReleasabilityID string                                      `json:"releasabilityId"`
	JSON            accountCloudforceOneEventUpdateResponseJSON `json:"-"`
}

// accountCloudforceOneEventUpdateResponseJSON contains the JSON metadata for the
// struct [AccountCloudforceOneEventUpdateResponse]
type accountCloudforceOneEventUpdateResponseJSON struct {
	ID              apijson.Field
	AccountID       apijson.Field
	Attacker        apijson.Field
	AttackerCountry apijson.Field
	Category        apijson.Field
	Date            apijson.Field
	Event           apijson.Field
	Indicator       apijson.Field
	IndicatorType   apijson.Field
	IndicatorTypeID apijson.Field
	KillChain       apijson.Field
	MitreAttack     apijson.Field
	NumReferenced   apijson.Field
	NumReferences   apijson.Field
	RawID           apijson.Field
	Referenced      apijson.Field
	ReferencedIDs   apijson.Field
	References      apijson.Field
	ReferencesIDs   apijson.Field
	Tags            apijson.Field
	TargetCountry   apijson.Field
	TargetIndustry  apijson.Field
	Tlp             apijson.Field
	Uuid            apijson.Field
	Insight         apijson.Field
	ReleasabilityID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountCloudforceOneEventUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventDeleteResponse struct {
	Uuid string                                      `json:"uuid" api:"required"`
	JSON accountCloudforceOneEventDeleteResponseJSON `json:"-"`
}

// accountCloudforceOneEventDeleteResponseJSON contains the JSON metadata for the
// struct [AccountCloudforceOneEventDeleteResponse]
type accountCloudforceOneEventDeleteResponseJSON struct {
	Uuid        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventListAttackersResponse struct {
	Items AccountCloudforceOneEventListAttackersResponseItems `json:"items" api:"required"`
	Type  string                                              `json:"type" api:"required"`
	JSON  accountCloudforceOneEventListAttackersResponseJSON  `json:"-"`
}

// accountCloudforceOneEventListAttackersResponseJSON contains the JSON metadata
// for the struct [AccountCloudforceOneEventListAttackersResponse]
type accountCloudforceOneEventListAttackersResponseJSON struct {
	Items       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventListAttackersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventListAttackersResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventListAttackersResponseItems struct {
	Type string                                                  `json:"type" api:"required"`
	JSON accountCloudforceOneEventListAttackersResponseItemsJSON `json:"-"`
}

// accountCloudforceOneEventListAttackersResponseItemsJSON contains the JSON
// metadata for the struct [AccountCloudforceOneEventListAttackersResponseItems]
type accountCloudforceOneEventListAttackersResponseItemsJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventListAttackersResponseItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventListAttackersResponseItemsJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventListCountriesResponse struct {
	Result  []AccountCloudforceOneEventListCountriesResponseResult `json:"result" api:"required"`
	Success string                                                 `json:"success" api:"required"`
	JSON    accountCloudforceOneEventListCountriesResponseJSON     `json:"-"`
}

// accountCloudforceOneEventListCountriesResponseJSON contains the JSON metadata
// for the struct [AccountCloudforceOneEventListCountriesResponse]
type accountCloudforceOneEventListCountriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventListCountriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventListCountriesResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventListCountriesResponseResult struct {
	Alpha3 string                                                   `json:"alpha3" api:"required"`
	Name   string                                                   `json:"name" api:"required"`
	JSON   accountCloudforceOneEventListCountriesResponseResultJSON `json:"-"`
}

// accountCloudforceOneEventListCountriesResponseResultJSON contains the JSON
// metadata for the struct [AccountCloudforceOneEventListCountriesResponseResult]
type accountCloudforceOneEventListCountriesResponseResultJSON struct {
	Alpha3      apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventListCountriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventListCountriesResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventListIndicatorTypesResponse struct {
	Items AccountCloudforceOneEventListIndicatorTypesResponseItems `json:"items" api:"required"`
	Type  string                                                   `json:"type" api:"required"`
	JSON  accountCloudforceOneEventListIndicatorTypesResponseJSON  `json:"-"`
}

// accountCloudforceOneEventListIndicatorTypesResponseJSON contains the JSON
// metadata for the struct [AccountCloudforceOneEventListIndicatorTypesResponse]
type accountCloudforceOneEventListIndicatorTypesResponseJSON struct {
	Items       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventListIndicatorTypesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventListIndicatorTypesResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventListIndicatorTypesResponseItems struct {
	Type string                                                       `json:"type" api:"required"`
	JSON accountCloudforceOneEventListIndicatorTypesResponseItemsJSON `json:"-"`
}

// accountCloudforceOneEventListIndicatorTypesResponseItemsJSON contains the JSON
// metadata for the struct
// [AccountCloudforceOneEventListIndicatorTypesResponseItems]
type accountCloudforceOneEventListIndicatorTypesResponseItemsJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventListIndicatorTypesResponseItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventListIndicatorTypesResponseItemsJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventListTargetIndustriesResponse struct {
	Items AccountCloudforceOneEventListTargetIndustriesResponseItems `json:"items" api:"required"`
	Type  string                                                     `json:"type" api:"required"`
	JSON  accountCloudforceOneEventListTargetIndustriesResponseJSON  `json:"-"`
}

// accountCloudforceOneEventListTargetIndustriesResponseJSON contains the JSON
// metadata for the struct [AccountCloudforceOneEventListTargetIndustriesResponse]
type accountCloudforceOneEventListTargetIndustriesResponseJSON struct {
	Items       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventListTargetIndustriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventListTargetIndustriesResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventListTargetIndustriesResponseItems struct {
	Type string                                                         `json:"type" api:"required"`
	JSON accountCloudforceOneEventListTargetIndustriesResponseItemsJSON `json:"-"`
}

// accountCloudforceOneEventListTargetIndustriesResponseItemsJSON contains the JSON
// metadata for the struct
// [AccountCloudforceOneEventListTargetIndustriesResponseItems]
type accountCloudforceOneEventListTargetIndustriesResponseItemsJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventListTargetIndustriesResponseItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventListTargetIndustriesResponseItemsJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventNewParams struct {
	Category        param.Field[string]                                `json:"category" api:"required"`
	Date            param.Field[time.Time]                             `json:"date" api:"required" format:"date-time"`
	Event           param.Field[string]                                `json:"event" api:"required"`
	IndicatorType   param.Field[string]                                `json:"indicatorType" api:"required"`
	Raw             param.Field[AccountCloudforceOneEventNewParamsRaw] `json:"raw" api:"required"`
	Tlp             param.Field[string]                                `json:"tlp" api:"required"`
	AccountID       param.Field[float64]                               `json:"accountId"`
	Attacker        param.Field[string]                                `json:"attacker"`
	AttackerCountry param.Field[string]                                `json:"attackerCountry"`
	DatasetID       param.Field[string]                                `json:"datasetId"`
	Indicator       param.Field[string]                                `json:"indicator"`
	Tags            param.Field[[]string]                              `json:"tags"`
	TargetCountry   param.Field[string]                                `json:"targetCountry"`
	TargetIndustry  param.Field[string]                                `json:"targetIndustry"`
}

func (r AccountCloudforceOneEventNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCloudforceOneEventNewParamsRaw struct {
	Data   param.Field[map[string]interface{}] `json:"data" api:"required"`
	Source param.Field[string]                 `json:"source"`
	Tlp    param.Field[string]                 `json:"tlp"`
}

func (r AccountCloudforceOneEventNewParamsRaw) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCloudforceOneEventUpdateParams struct {
	Attacker        param.Field[string]                                   `json:"attacker"`
	AttackerCountry param.Field[string]                                   `json:"attackerCountry"`
	Category        param.Field[string]                                   `json:"category"`
	Date            param.Field[time.Time]                                `json:"date" format:"date-time"`
	Event           param.Field[string]                                   `json:"event"`
	Indicator       param.Field[string]                                   `json:"indicator"`
	IndicatorType   param.Field[string]                                   `json:"indicatorType"`
	Insight         param.Field[string]                                   `json:"insight"`
	Raw             param.Field[AccountCloudforceOneEventUpdateParamsRaw] `json:"raw"`
	TargetCountry   param.Field[string]                                   `json:"targetCountry"`
	TargetIndustry  param.Field[string]                                   `json:"targetIndustry"`
	Tlp             param.Field[string]                                   `json:"tlp"`
}

func (r AccountCloudforceOneEventUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCloudforceOneEventUpdateParamsRaw struct {
	Data   param.Field[map[string]interface{}] `json:"data"`
	Source param.Field[string]                 `json:"source"`
	Tlp    param.Field[string]                 `json:"tlp"`
}

func (r AccountCloudforceOneEventUpdateParamsRaw) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCloudforceOneEventNewBulkParams struct {
	Data      param.Field[[]AccountCloudforceOneEventNewBulkParamsData] `json:"data" api:"required"`
	DatasetID param.Field[string]                                       `json:"datasetId" api:"required"`
}

func (r AccountCloudforceOneEventNewBulkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCloudforceOneEventNewBulkParamsData struct {
	Category        param.Field[string]                                        `json:"category" api:"required"`
	Date            param.Field[time.Time]                                     `json:"date" api:"required" format:"date-time"`
	Event           param.Field[string]                                        `json:"event" api:"required"`
	IndicatorType   param.Field[string]                                        `json:"indicatorType" api:"required"`
	Raw             param.Field[AccountCloudforceOneEventNewBulkParamsDataRaw] `json:"raw" api:"required"`
	Tlp             param.Field[string]                                        `json:"tlp" api:"required"`
	AccountID       param.Field[float64]                                       `json:"accountId"`
	Attacker        param.Field[string]                                        `json:"attacker"`
	AttackerCountry param.Field[string]                                        `json:"attackerCountry"`
	DatasetID       param.Field[string]                                        `json:"datasetId"`
	Indicator       param.Field[string]                                        `json:"indicator"`
	Tags            param.Field[[]string]                                      `json:"tags"`
	TargetCountry   param.Field[string]                                        `json:"targetCountry"`
	TargetIndustry  param.Field[string]                                        `json:"targetIndustry"`
}

func (r AccountCloudforceOneEventNewBulkParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCloudforceOneEventNewBulkParamsDataRaw struct {
	Data   param.Field[map[string]interface{}] `json:"data" api:"required"`
	Source param.Field[string]                 `json:"source"`
	Tlp    param.Field[string]                 `json:"tlp"`
}

func (r AccountCloudforceOneEventNewBulkParamsDataRaw) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
