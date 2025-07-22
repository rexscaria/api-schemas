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

// Events must be created in a client-specific dataset, which means the `datasetId`
// parameter must be defined. To create a dataset, see the
// [`Create Dataset`](https://developers.cloudflare.com/api/resources/cloudforce_one/subresources/threat_events/subresources/datasets/methods/create/)
// endpoint.
func (r *AccountCloudforceOneEventService) New(ctx context.Context, accountID float64, body AccountCloudforceOneEventNewParams, opts ...option.RequestOption) (res *AccountCloudforceOneEventNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/create", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Reads an event
func (r *AccountCloudforceOneEventService) Get(ctx context.Context, accountID float64, eventID string, opts ...option.RequestOption) (res *AccountCloudforceOneEventGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/%s", accountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an event
func (r *AccountCloudforceOneEventService) Update(ctx context.Context, accountID float64, eventID string, body AccountCloudforceOneEventUpdateParams, opts ...option.RequestOption) (res *AccountCloudforceOneEventUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/%s", accountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The `datasetId` parameter must be defined. To list existing datasets (and their
// IDs) in your account, use the
// [`List Datasets`](https://developers.cloudflare.com/api/resources/cloudforce_one/subresources/threat_events/subresources/datasets/methods/list/)
// endpoint.
func (r *AccountCloudforceOneEventService) Delete(ctx context.Context, accountID float64, eventID string, opts ...option.RequestOption) (res *AccountCloudforceOneEventDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/%s", accountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// The `datasetId` parameter must be defined. To list existing datasets (and their
// IDs) in your account, use the
// [`List Datasets`](https://developers.cloudflare.com/api/resources/cloudforce_one/subresources/threat_events/subresources/datasets/methods/list/)
// endpoint.
func (r *AccountCloudforceOneEventService) NewBulk(ctx context.Context, accountID float64, body AccountCloudforceOneEventNewBulkParams, opts ...option.RequestOption) (res *[]AccountCloudforceOneEventNewBulkResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/create/bulk", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists attackers
func (r *AccountCloudforceOneEventService) ListAttackers(ctx context.Context, accountID float64, opts ...option.RequestOption) (res *AccountCloudforceOneEventListAttackersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/attackers", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves countries information for all countries
func (r *AccountCloudforceOneEventService) ListCountries(ctx context.Context, accountID float64, opts ...option.RequestOption) (res *[]AccountCloudforceOneEventListCountriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/countries", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all indicator types
func (r *AccountCloudforceOneEventService) ListIndicatorTypes(ctx context.Context, accountID float64, opts ...option.RequestOption) (res *AccountCloudforceOneEventListIndicatorTypesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/indicatorTypes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all target industries
func (r *AccountCloudforceOneEventService) ListTargetIndustries(ctx context.Context, accountID float64, opts ...option.RequestOption) (res *AccountCloudforceOneEventListTargetIndustriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/targetIndustries", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountCloudforceOneEventNewResponse struct {
	ID              float64                                  `json:"id,required"`
	AccountID       float64                                  `json:"accountId,required"`
	Attacker        string                                   `json:"attacker,required"`
	AttackerCountry string                                   `json:"attackerCountry,required"`
	Category        string                                   `json:"category,required"`
	CategoryID      float64                                  `json:"categoryId,required"`
	Date            string                                   `json:"date,required"`
	Event           string                                   `json:"event,required"`
	Indicator       string                                   `json:"indicator,required"`
	IndicatorType   string                                   `json:"indicatorType,required"`
	IndicatorTypeID float64                                  `json:"indicatorTypeId,required"`
	KillChain       float64                                  `json:"killChain,required"`
	MitreAttack     []string                                 `json:"mitreAttack,required"`
	NumReferenced   float64                                  `json:"numReferenced,required"`
	NumReferences   float64                                  `json:"numReferences,required"`
	RawID           string                                   `json:"rawId,required"`
	Referenced      []string                                 `json:"referenced,required"`
	ReferencedIDs   []float64                                `json:"referencedIds,required"`
	References      []string                                 `json:"references,required"`
	ReferencesIDs   []float64                                `json:"referencesIds,required"`
	Tags            []string                                 `json:"tags,required"`
	TargetCountry   string                                   `json:"targetCountry,required"`
	TargetIndustry  string                                   `json:"targetIndustry,required"`
	Tlp             string                                   `json:"tlp,required"`
	Uuid            string                                   `json:"uuid,required"`
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
	CategoryID      apijson.Field
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
	ID              float64                                  `json:"id,required"`
	AccountID       float64                                  `json:"accountId,required"`
	Attacker        string                                   `json:"attacker,required"`
	AttackerCountry string                                   `json:"attackerCountry,required"`
	Category        string                                   `json:"category,required"`
	CategoryID      float64                                  `json:"categoryId,required"`
	Date            string                                   `json:"date,required"`
	Event           string                                   `json:"event,required"`
	Indicator       string                                   `json:"indicator,required"`
	IndicatorType   string                                   `json:"indicatorType,required"`
	IndicatorTypeID float64                                  `json:"indicatorTypeId,required"`
	KillChain       float64                                  `json:"killChain,required"`
	MitreAttack     []string                                 `json:"mitreAttack,required"`
	NumReferenced   float64                                  `json:"numReferenced,required"`
	NumReferences   float64                                  `json:"numReferences,required"`
	RawID           string                                   `json:"rawId,required"`
	Referenced      []string                                 `json:"referenced,required"`
	ReferencedIDs   []float64                                `json:"referencedIds,required"`
	References      []string                                 `json:"references,required"`
	ReferencesIDs   []float64                                `json:"referencesIds,required"`
	Tags            []string                                 `json:"tags,required"`
	TargetCountry   string                                   `json:"targetCountry,required"`
	TargetIndustry  string                                   `json:"targetIndustry,required"`
	Tlp             string                                   `json:"tlp,required"`
	Uuid            string                                   `json:"uuid,required"`
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
	CategoryID      apijson.Field
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
	ID              float64                                     `json:"id,required"`
	AccountID       float64                                     `json:"accountId,required"`
	Attacker        string                                      `json:"attacker,required"`
	AttackerCountry string                                      `json:"attackerCountry,required"`
	Category        string                                      `json:"category,required"`
	CategoryID      float64                                     `json:"categoryId,required"`
	Date            string                                      `json:"date,required"`
	Event           string                                      `json:"event,required"`
	Indicator       string                                      `json:"indicator,required"`
	IndicatorType   string                                      `json:"indicatorType,required"`
	IndicatorTypeID float64                                     `json:"indicatorTypeId,required"`
	KillChain       float64                                     `json:"killChain,required"`
	MitreAttack     []string                                    `json:"mitreAttack,required"`
	NumReferenced   float64                                     `json:"numReferenced,required"`
	NumReferences   float64                                     `json:"numReferences,required"`
	RawID           string                                      `json:"rawId,required"`
	Referenced      []string                                    `json:"referenced,required"`
	ReferencedIDs   []float64                                   `json:"referencedIds,required"`
	References      []string                                    `json:"references,required"`
	ReferencesIDs   []float64                                   `json:"referencesIds,required"`
	Tags            []string                                    `json:"tags,required"`
	TargetCountry   string                                      `json:"targetCountry,required"`
	TargetIndustry  string                                      `json:"targetIndustry,required"`
	Tlp             string                                      `json:"tlp,required"`
	Uuid            string                                      `json:"uuid,required"`
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
	CategoryID      apijson.Field
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
	Uuid string                                      `json:"uuid,required"`
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

type AccountCloudforceOneEventNewBulkResponse struct {
	ID              float64                                      `json:"id,required"`
	AccountID       float64                                      `json:"accountId,required"`
	Attacker        string                                       `json:"attacker,required"`
	AttackerCountry string                                       `json:"attackerCountry,required"`
	Category        string                                       `json:"category,required"`
	CategoryID      float64                                      `json:"categoryId,required"`
	Date            string                                       `json:"date,required"`
	Event           string                                       `json:"event,required"`
	Indicator       string                                       `json:"indicator,required"`
	IndicatorType   string                                       `json:"indicatorType,required"`
	IndicatorTypeID float64                                      `json:"indicatorTypeId,required"`
	KillChain       float64                                      `json:"killChain,required"`
	MitreAttack     []string                                     `json:"mitreAttack,required"`
	NumReferenced   float64                                      `json:"numReferenced,required"`
	NumReferences   float64                                      `json:"numReferences,required"`
	RawID           string                                       `json:"rawId,required"`
	Referenced      []string                                     `json:"referenced,required"`
	ReferencedIDs   []float64                                    `json:"referencedIds,required"`
	References      []string                                     `json:"references,required"`
	ReferencesIDs   []float64                                    `json:"referencesIds,required"`
	Tags            []string                                     `json:"tags,required"`
	TargetCountry   string                                       `json:"targetCountry,required"`
	TargetIndustry  string                                       `json:"targetIndustry,required"`
	Tlp             string                                       `json:"tlp,required"`
	Uuid            string                                       `json:"uuid,required"`
	Insight         string                                       `json:"insight"`
	ReleasabilityID string                                       `json:"releasabilityId"`
	JSON            accountCloudforceOneEventNewBulkResponseJSON `json:"-"`
}

// accountCloudforceOneEventNewBulkResponseJSON contains the JSON metadata for the
// struct [AccountCloudforceOneEventNewBulkResponse]
type accountCloudforceOneEventNewBulkResponseJSON struct {
	ID              apijson.Field
	AccountID       apijson.Field
	Attacker        apijson.Field
	AttackerCountry apijson.Field
	Category        apijson.Field
	CategoryID      apijson.Field
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

func (r *AccountCloudforceOneEventNewBulkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventNewBulkResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventListAttackersResponse struct {
	Items AccountCloudforceOneEventListAttackersResponseItems `json:"items,required"`
	Type  string                                              `json:"type,required"`
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
	Type string                                                  `json:"type,required"`
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
	Result  []AccountCloudforceOneEventListCountriesResponseResult `json:"result,required"`
	Success string                                                 `json:"success,required"`
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
	Alpha3 string                                                   `json:"alpha3,required"`
	Name   string                                                   `json:"name,required"`
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
	Items AccountCloudforceOneEventListIndicatorTypesResponseItems `json:"items,required"`
	Type  string                                                   `json:"type,required"`
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
	Type string                                                       `json:"type,required"`
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
	Items AccountCloudforceOneEventListTargetIndustriesResponseItems `json:"items,required"`
	Type  string                                                     `json:"type,required"`
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
	Type string                                                         `json:"type,required"`
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
	Attacker        param.Field[string]                                `json:"attacker,required"`
	AttackerCountry param.Field[string]                                `json:"attackerCountry,required"`
	Category        param.Field[string]                                `json:"category,required"`
	Date            param.Field[time.Time]                             `json:"date,required" format:"date-time"`
	Event           param.Field[string]                                `json:"event,required"`
	IndicatorType   param.Field[string]                                `json:"indicatorType,required"`
	Raw             param.Field[AccountCloudforceOneEventNewParamsRaw] `json:"raw,required"`
	Tlp             param.Field[string]                                `json:"tlp,required"`
	AccountID       param.Field[float64]                               `json:"accountId"`
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
	Data   param.Field[interface{}] `json:"data"`
	Source param.Field[string]      `json:"source"`
	Tlp    param.Field[string]      `json:"tlp"`
}

func (r AccountCloudforceOneEventNewParamsRaw) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCloudforceOneEventUpdateParams struct {
	Attacker        param.Field[string]    `json:"attacker"`
	AttackerCountry param.Field[string]    `json:"attackerCountry"`
	Category        param.Field[string]    `json:"category"`
	Date            param.Field[time.Time] `json:"date" format:"date-time"`
	Event           param.Field[string]    `json:"event"`
	Indicator       param.Field[string]    `json:"indicator"`
	IndicatorType   param.Field[string]    `json:"indicatorType"`
	TargetCountry   param.Field[string]    `json:"targetCountry"`
	TargetIndustry  param.Field[string]    `json:"targetIndustry"`
	Tlp             param.Field[string]    `json:"tlp"`
}

func (r AccountCloudforceOneEventUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCloudforceOneEventNewBulkParams struct {
	Data      param.Field[[]AccountCloudforceOneEventNewBulkParamsData] `json:"data,required"`
	DatasetID param.Field[string]                                       `json:"datasetId,required"`
}

func (r AccountCloudforceOneEventNewBulkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCloudforceOneEventNewBulkParamsData struct {
	Attacker        param.Field[string]                                        `json:"attacker,required"`
	AttackerCountry param.Field[string]                                        `json:"attackerCountry,required"`
	Category        param.Field[string]                                        `json:"category,required"`
	Date            param.Field[time.Time]                                     `json:"date,required" format:"date-time"`
	Event           param.Field[string]                                        `json:"event,required"`
	IndicatorType   param.Field[string]                                        `json:"indicatorType,required"`
	Raw             param.Field[AccountCloudforceOneEventNewBulkParamsDataRaw] `json:"raw,required"`
	Tlp             param.Field[string]                                        `json:"tlp,required"`
	AccountID       param.Field[float64]                                       `json:"accountId"`
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
	Data   param.Field[interface{}] `json:"data"`
	Source param.Field[string]      `json:"source"`
	Tlp    param.Field[string]      `json:"tlp"`
}

func (r AccountCloudforceOneEventNewBulkParamsDataRaw) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
