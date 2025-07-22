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

// ZoneAvailablePlanService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAvailablePlanService] method instead.
type ZoneAvailablePlanService struct {
	Options []option.RequestOption
}

// NewZoneAvailablePlanService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAvailablePlanService(opts ...option.RequestOption) (r *ZoneAvailablePlanService) {
	r = &ZoneAvailablePlanService{}
	r.Options = opts
	return
}

// Details of the available plan that the zone can subscribe to.
func (r *ZoneAvailablePlanService) Get(ctx context.Context, zoneID string, planIdentifier string, opts ...option.RequestOption) (res *ZoneAvailablePlanGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if planIdentifier == "" {
		err = errors.New("missing required plan_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/available_plans/%s", zoneID, planIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists available plans the zone can subscribe to.
func (r *ZoneAvailablePlanService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneAvailablePlanListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/available_plans", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BillSubsAPIAvailableRatePlan struct {
	// Identifier
	ID string `json:"id"`
	// Indicates whether you can subscribe to this plan.
	CanSubscribe bool `json:"can_subscribe"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// Indicates whether this plan is managed externally.
	ExternallyManaged bool `json:"externally_managed"`
	// The frequency at which you will be billed for this plan.
	Frequency BillSubsAPIFrequency `json:"frequency"`
	// Indicates whether you are currently subscribed to this plan.
	IsSubscribed bool `json:"is_subscribed"`
	// Indicates whether this plan has a legacy discount applied.
	LegacyDiscount bool `json:"legacy_discount"`
	// The legacy identifier for this rate plan, if any.
	LegacyID string `json:"legacy_id"`
	// The plan name.
	Name string `json:"name"`
	// The amount you will be billed for this plan.
	Price float64                          `json:"price"`
	JSON  billSubsAPIAvailableRatePlanJSON `json:"-"`
}

// billSubsAPIAvailableRatePlanJSON contains the JSON metadata for the struct
// [BillSubsAPIAvailableRatePlan]
type billSubsAPIAvailableRatePlanJSON struct {
	ID                apijson.Field
	CanSubscribe      apijson.Field
	Currency          apijson.Field
	ExternallyManaged apijson.Field
	Frequency         apijson.Field
	IsSubscribed      apijson.Field
	LegacyDiscount    apijson.Field
	LegacyID          apijson.Field
	Name              apijson.Field
	Price             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *BillSubsAPIAvailableRatePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billSubsAPIAvailableRatePlanJSON) RawJSON() string {
	return r.raw
}

// The frequency at which you will be billed for this plan.
type BillSubsAPIFrequency string

const (
	BillSubsAPIFrequencyWeekly    BillSubsAPIFrequency = "weekly"
	BillSubsAPIFrequencyMonthly   BillSubsAPIFrequency = "monthly"
	BillSubsAPIFrequencyQuarterly BillSubsAPIFrequency = "quarterly"
	BillSubsAPIFrequencyYearly    BillSubsAPIFrequency = "yearly"
)

func (r BillSubsAPIFrequency) IsKnown() bool {
	switch r {
	case BillSubsAPIFrequencyWeekly, BillSubsAPIFrequencyMonthly, BillSubsAPIFrequencyQuarterly, BillSubsAPIFrequencyYearly:
		return true
	}
	return false
}

type BillSubsAPIResponseCollection struct {
	Result     []interface{}                           `json:"result,nullable"`
	ResultInfo BillSubsAPIResponseCollectionResultInfo `json:"result_info"`
	JSON       billSubsAPIResponseCollectionJSON       `json:"-"`
	APIResponseBilling
}

// billSubsAPIResponseCollectionJSON contains the JSON metadata for the struct
// [BillSubsAPIResponseCollection]
type billSubsAPIResponseCollectionJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillSubsAPIResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billSubsAPIResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type BillSubsAPIResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       billSubsAPIResponseCollectionResultInfoJSON `json:"-"`
}

// billSubsAPIResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [BillSubsAPIResponseCollectionResultInfo]
type billSubsAPIResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillSubsAPIResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billSubsAPIResponseCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneAvailablePlanGetResponse struct {
	Result BillSubsAPIAvailableRatePlan     `json:"result"`
	JSON   zoneAvailablePlanGetResponseJSON `json:"-"`
	APIResponseSingleBilling
}

// zoneAvailablePlanGetResponseJSON contains the JSON metadata for the struct
// [ZoneAvailablePlanGetResponse]
type zoneAvailablePlanGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAvailablePlanGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAvailablePlanGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAvailablePlanListResponse struct {
	Result []BillSubsAPIAvailableRatePlan    `json:"result"`
	JSON   zoneAvailablePlanListResponseJSON `json:"-"`
	BillSubsAPIResponseCollection
}

// zoneAvailablePlanListResponseJSON contains the JSON metadata for the struct
// [ZoneAvailablePlanListResponse]
type zoneAvailablePlanListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAvailablePlanListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAvailablePlanListResponseJSON) RawJSON() string {
	return r.raw
}
