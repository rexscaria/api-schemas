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

// AccountCloudforceOneRequestPriorityService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCloudforceOneRequestPriorityService] method instead.
type AccountCloudforceOneRequestPriorityService struct {
	Options []option.RequestOption
}

// NewAccountCloudforceOneRequestPriorityService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountCloudforceOneRequestPriorityService(opts ...option.RequestOption) (r *AccountCloudforceOneRequestPriorityService) {
	r = &AccountCloudforceOneRequestPriorityService{}
	r.Options = opts
	return
}

// Create a New Priority Intelligence Requirement
func (r *AccountCloudforceOneRequestPriorityService) New(ctx context.Context, accountIdentifier string, body AccountCloudforceOneRequestPriorityNewParams, opts ...option.RequestOption) (res *AccountCloudforceOneRequestPriorityNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/new", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a Priority Intelligence Requirement
func (r *AccountCloudforceOneRequestPriorityService) Get(ctx context.Context, accountIdentifier string, priorityIdentifer string, opts ...option.RequestOption) (res *AccountCloudforceOneRequestPriorityGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if priorityIdentifer == "" {
		err = errors.New("missing required priority_identifer parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", accountIdentifier, priorityIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Priority Intelligence Requirement
func (r *AccountCloudforceOneRequestPriorityService) Update(ctx context.Context, accountIdentifier string, priorityIdentifer string, body AccountCloudforceOneRequestPriorityUpdateParams, opts ...option.RequestOption) (res *AccountCloudforceOneRequestPriorityUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if priorityIdentifer == "" {
		err = errors.New("missing required priority_identifer parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", accountIdentifier, priorityIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List Priority Intelligence Requirements
func (r *AccountCloudforceOneRequestPriorityService) List(ctx context.Context, accountIdentifier string, body AccountCloudforceOneRequestPriorityListParams, opts ...option.RequestOption) (res *AccountCloudforceOneRequestPriorityListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a Priority Intelligence Requirement
func (r *AccountCloudforceOneRequestPriorityService) Delete(ctx context.Context, accountIdentifier string, priorityIdentifer string, opts ...option.RequestOption) (res *APIResponseRequests, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if priorityIdentifer == "" {
		err = errors.New("missing required priority_identifer parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", accountIdentifier, priorityIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type PriorityEditParam struct {
	// List of labels
	Labels param.Field[[]string] `json:"labels,required"`
	// Priority
	Priority param.Field[int64] `json:"priority,required"`
	// Requirement
	Requirement param.Field[string] `json:"requirement,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp param.Field[Tlp] `json:"tlp,required"`
}

func (r PriorityEditParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriorityItem struct {
	// UUID
	ID string `json:"id,required"`
	// Priority creation time
	Created time.Time `json:"created,required" format:"date-time"`
	// List of labels
	Labels []string `json:"labels,required"`
	// Priority
	Priority int64 `json:"priority,required"`
	// Requirement
	Requirement string `json:"requirement,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp Tlp `json:"tlp,required"`
	// Priority last updated time
	Updated time.Time        `json:"updated,required" format:"date-time"`
	JSON    priorityItemJSON `json:"-"`
}

// priorityItemJSON contains the JSON metadata for the struct [PriorityItem]
type priorityItemJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Labels      apijson.Field
	Priority    apijson.Field
	Requirement apijson.Field
	Tlp         apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriorityItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priorityItemJSON) RawJSON() string {
	return r.raw
}

type Quota struct {
	// Anniversary date is when annual quota limit is refresh
	AnniversaryDate QuotaAnniversaryDate `json:"anniversary_date"`
	// Quater anniversary date is when quota limit is refreshed each quarter
	QuarterAnniversaryDate QuotaQuarterAnniversaryDate `json:"quarter_anniversary_date"`
	// Tokens for the quarter
	Quota int64 `json:"quota"`
	// Tokens remaining for the quarter
	Remaining int64     `json:"remaining"`
	JSON      quotaJSON `json:"-"`
}

// quotaJSON contains the JSON metadata for the struct [Quota]
type quotaJSON struct {
	AnniversaryDate        apijson.Field
	QuarterAnniversaryDate apijson.Field
	Quota                  apijson.Field
	Remaining              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *Quota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotaJSON) RawJSON() string {
	return r.raw
}

// Anniversary date is when annual quota limit is refresh
type QuotaAnniversaryDate struct {
	JSON quotaAnniversaryDateJSON `json:"-"`
}

// quotaAnniversaryDateJSON contains the JSON metadata for the struct
// [QuotaAnniversaryDate]
type quotaAnniversaryDateJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuotaAnniversaryDate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotaAnniversaryDateJSON) RawJSON() string {
	return r.raw
}

// Quater anniversary date is when quota limit is refreshed each quarter
type QuotaQuarterAnniversaryDate struct {
	JSON quotaQuarterAnniversaryDateJSON `json:"-"`
}

// quotaQuarterAnniversaryDateJSON contains the JSON metadata for the struct
// [QuotaQuarterAnniversaryDate]
type quotaQuarterAnniversaryDateJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuotaQuarterAnniversaryDate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotaQuarterAnniversaryDateJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestPriorityNewResponse struct {
	Result PriorityItem                                       `json:"result"`
	JSON   accountCloudforceOneRequestPriorityNewResponseJSON `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestPriorityNewResponseJSON contains the JSON metadata
// for the struct [AccountCloudforceOneRequestPriorityNewResponse]
type accountCloudforceOneRequestPriorityNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestPriorityNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestPriorityNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestPriorityGetResponse struct {
	Result RequestItem                                        `json:"result"`
	JSON   accountCloudforceOneRequestPriorityGetResponseJSON `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestPriorityGetResponseJSON contains the JSON metadata
// for the struct [AccountCloudforceOneRequestPriorityGetResponse]
type accountCloudforceOneRequestPriorityGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestPriorityGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestPriorityGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestPriorityUpdateResponse struct {
	Result RequestItem                                           `json:"result"`
	JSON   accountCloudforceOneRequestPriorityUpdateResponseJSON `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestPriorityUpdateResponseJSON contains the JSON metadata
// for the struct [AccountCloudforceOneRequestPriorityUpdateResponse]
type accountCloudforceOneRequestPriorityUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestPriorityUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestPriorityUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestPriorityListResponse struct {
	Result []PriorityItem                                      `json:"result"`
	JSON   accountCloudforceOneRequestPriorityListResponseJSON `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestPriorityListResponseJSON contains the JSON metadata
// for the struct [AccountCloudforceOneRequestPriorityListResponse]
type accountCloudforceOneRequestPriorityListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestPriorityListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestPriorityListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestPriorityNewParams struct {
	PriorityEdit PriorityEditParam `json:"priority_edit,required"`
}

func (r AccountCloudforceOneRequestPriorityNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PriorityEdit)
}

type AccountCloudforceOneRequestPriorityUpdateParams struct {
	PriorityEdit PriorityEditParam `json:"priority_edit,required"`
}

func (r AccountCloudforceOneRequestPriorityUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PriorityEdit)
}

type AccountCloudforceOneRequestPriorityListParams struct {
	// Page number of results
	Page param.Field[int64] `json:"page,required"`
	// Number of results per page
	PerPage param.Field[int64] `json:"per_page,required"`
}

func (r AccountCloudforceOneRequestPriorityListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
