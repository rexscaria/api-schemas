// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountMagicConnectorTelemetryEventService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicConnectorTelemetryEventService] method instead.
type AccountMagicConnectorTelemetryEventService struct {
	Options []option.RequestOption
}

// NewAccountMagicConnectorTelemetryEventService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountMagicConnectorTelemetryEventService(opts ...option.RequestOption) (r *AccountMagicConnectorTelemetryEventService) {
	r = &AccountMagicConnectorTelemetryEventService{}
	r.Options = opts
	return
}

// List Events
func (r *AccountMagicConnectorTelemetryEventService) List(ctx context.Context, accountID float64, connectorID string, query AccountMagicConnectorTelemetryEventListParams, opts ...option.RequestOption) (res *AccountMagicConnectorTelemetryEventListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/magic/connectors/%s/telemetry/events", accountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Event
func (r *AccountMagicConnectorTelemetryEventService) Get(ctx context.Context, accountID float64, connectorID string, eventT float64, eventN float64, opts ...option.RequestOption) (res *AccountMagicConnectorTelemetryEventGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/magic/connectors/%s/telemetry/events/%v.%v", accountID, connectorID, eventT, eventN)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type MconnCodedMessage struct {
	Code    float64               `json:"code,required"`
	Message string                `json:"message,required"`
	JSON    mconnCodedMessageJSON `json:"-"`
}

// mconnCodedMessageJSON contains the JSON metadata for the struct
// [MconnCodedMessage]
type mconnCodedMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MconnCodedMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mconnCodedMessageJSON) RawJSON() string {
	return r.raw
}

type MconnEnvelope struct {
	Success  bool                `json:"success,required"`
	Errors   []MconnCodedMessage `json:"errors"`
	Messages []MconnCodedMessage `json:"messages"`
	JSON     mconnEnvelopeJSON   `json:"-"`
}

// mconnEnvelopeJSON contains the JSON metadata for the struct [MconnEnvelope]
type mconnEnvelopeJSON struct {
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MconnEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mconnEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccountMagicConnectorTelemetryEventListResponse struct {
	Result AccountMagicConnectorTelemetryEventListResponseResult `json:"result,required"`
	JSON   accountMagicConnectorTelemetryEventListResponseJSON   `json:"-"`
	MconnEnvelope
}

// accountMagicConnectorTelemetryEventListResponseJSON contains the JSON metadata
// for the struct [AccountMagicConnectorTelemetryEventListResponse]
type accountMagicConnectorTelemetryEventListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicConnectorTelemetryEventListResponseResult struct {
	Count  float64                                                     `json:"count,required"`
	Items  []AccountMagicConnectorTelemetryEventListResponseResultItem `json:"items,required"`
	Cursor string                                                      `json:"cursor"`
	JSON   accountMagicConnectorTelemetryEventListResponseResultJSON   `json:"-"`
}

// accountMagicConnectorTelemetryEventListResponseResultJSON contains the JSON
// metadata for the struct [AccountMagicConnectorTelemetryEventListResponseResult]
type accountMagicConnectorTelemetryEventListResponseResultJSON struct {
	Count       apijson.Field
	Items       apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicConnectorTelemetryEventListResponseResultItem struct {
	// Time the Event was collected (seconds since the Unix epoch)
	A float64 `json:"a,required"`
	// Kind
	K string `json:"k,required"`
	// Sequence number, used to order events with the same timestamp
	N float64 `json:"n,required"`
	// Time the Event was recorded (seconds since the Unix epoch)
	T    float64                                                       `json:"t,required"`
	JSON accountMagicConnectorTelemetryEventListResponseResultItemJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventListResponseResultItemJSON contains the JSON
// metadata for the struct
// [AccountMagicConnectorTelemetryEventListResponseResultItem]
type accountMagicConnectorTelemetryEventListResponseResultItemJSON struct {
	A           apijson.Field
	K           apijson.Field
	N           apijson.Field
	T           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventListResponseResultItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventListResponseResultItemJSON) RawJSON() string {
	return r.raw
}

type AccountMagicConnectorTelemetryEventGetResponse struct {
	// Recorded Event
	Result AccountMagicConnectorTelemetryEventGetResponseResult `json:"result,required"`
	JSON   accountMagicConnectorTelemetryEventGetResponseJSON   `json:"-"`
	MconnEnvelope
}

// accountMagicConnectorTelemetryEventGetResponseJSON contains the JSON metadata
// for the struct [AccountMagicConnectorTelemetryEventGetResponse]
type accountMagicConnectorTelemetryEventGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseJSON) RawJSON() string {
	return r.raw
}

// Recorded Event
type AccountMagicConnectorTelemetryEventGetResponseResult struct {
	E AccountMagicConnectorTelemetryEventGetResponseResultE `json:"e,required"`
	// Sequence number, used to order events with the same timestamp
	N float64 `json:"n,required"`
	// Time the Event was recorded (seconds since the Unix epoch)
	T    float64                                                  `json:"t,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultJSON contains the JSON
// metadata for the struct [AccountMagicConnectorTelemetryEventGetResponseResult]
type accountMagicConnectorTelemetryEventGetResponseResultJSON struct {
	E           apijson.Field
	N           apijson.Field
	T           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicConnectorTelemetryEventGetResponseResultE struct {
	// Initialized process
	K AccountMagicConnectorTelemetryEventGetResponseResultEK `json:"k,required"`
	// Location of upgrade bundle
	URL   string                                                    `json:"url"`
	JSON  accountMagicConnectorTelemetryEventGetResponseResultEJSON `json:"-"`
	union AccountMagicConnectorTelemetryEventGetResponseResultEUnion
}

// accountMagicConnectorTelemetryEventGetResponseResultEJSON contains the JSON
// metadata for the struct [AccountMagicConnectorTelemetryEventGetResponseResultE]
type accountMagicConnectorTelemetryEventGetResponseResultEJSON struct {
	K           apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r accountMagicConnectorTelemetryEventGetResponseResultEJSON) RawJSON() string {
	return r.raw
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultE) UnmarshalJSON(data []byte) (err error) {
	*r = AccountMagicConnectorTelemetryEventGetResponseResultE{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountMagicConnectorTelemetryEventGetResponseResultEUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEObject],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK].
func (r AccountMagicConnectorTelemetryEventGetResponseResultE) AsUnion() AccountMagicConnectorTelemetryEventGetResponseResultEUnion {
	return r.union
}

// Union satisfied by [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEObject],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK],
// [AccountMagicConnectorTelemetryEventGetResponseResultEK] or
// [AccountMagicConnectorTelemetryEventGetResponseResultEK].
type AccountMagicConnectorTelemetryEventGetResponseResultEUnion interface {
	implementsAccountMagicConnectorTelemetryEventGetResponseResultE()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountMagicConnectorTelemetryEventGetResponseResultEUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEK{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEK{}),
		},
	)
}

type AccountMagicConnectorTelemetryEventGetResponseResultEK struct {
	// Initialized process
	K    AccountMagicConnectorTelemetryEventGetResponseResultEKK    `json:"k,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultEkJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultEkJSON contains the JSON
// metadata for the struct [AccountMagicConnectorTelemetryEventGetResponseResultEK]
type accountMagicConnectorTelemetryEventGetResponseResultEkJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultEK) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultEkJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicConnectorTelemetryEventGetResponseResultEK) implementsAccountMagicConnectorTelemetryEventGetResponseResultE() {
}

// Initialized process
type AccountMagicConnectorTelemetryEventGetResponseResultEKK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultEKKInit AccountMagicConnectorTelemetryEventGetResponseResultEKK = "Init"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultEKK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultEKKInit:
		return true
	}
	return false
}

type AccountMagicConnectorTelemetryEventGetResponseResultEObject struct {
	// Started upgrade
	K AccountMagicConnectorTelemetryEventGetResponseResultEObjectK `json:"k,required"`
	// Location of upgrade bundle
	URL  string                                                          `json:"url,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultEObjectJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultEObjectJSON contains the
// JSON metadata for the struct
// [AccountMagicConnectorTelemetryEventGetResponseResultEObject]
type accountMagicConnectorTelemetryEventGetResponseResultEObjectJSON struct {
	K           apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultEObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultEObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicConnectorTelemetryEventGetResponseResultEObject) implementsAccountMagicConnectorTelemetryEventGetResponseResultE() {
}

// Started upgrade
type AccountMagicConnectorTelemetryEventGetResponseResultEObjectK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultEObjectKStartUpgrade AccountMagicConnectorTelemetryEventGetResponseResultEObjectK = "StartUpgrade"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultEObjectK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultEObjectKStartUpgrade:
		return true
	}
	return false
}

type AccountMagicConnectorTelemetryEventListParams struct {
	From   param.Field[float64] `query:"from,required"`
	To     param.Field[float64] `query:"to,required"`
	Cursor param.Field[string]  `query:"cursor"`
	Limit  param.Field[float64] `query:"limit"`
}

// URLQuery serializes [AccountMagicConnectorTelemetryEventListParams]'s query
// parameters as `url.Values`.
func (r AccountMagicConnectorTelemetryEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
