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
func (r *AccountMagicConnectorTelemetryEventService) List(ctx context.Context, accountID string, connectorID string, query AccountMagicConnectorTelemetryEventListParams, opts ...option.RequestOption) (res *AccountMagicConnectorTelemetryEventListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s/telemetry/events", accountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Event
func (r *AccountMagicConnectorTelemetryEventService) Get(ctx context.Context, accountID string, connectorID string, eventT float64, eventN float64, opts ...option.RequestOption) (res *AccountMagicConnectorTelemetryEventGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s/telemetry/events/%v.%v", accountID, connectorID, eventT, eventN)
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

type AccountMagicConnectorTelemetryEventListResponse struct {
	Result   AccountMagicConnectorTelemetryEventListResponseResult `json:"result,required"`
	Success  bool                                                  `json:"success,required"`
	Errors   []MconnCodedMessage                                   `json:"errors"`
	Messages []MconnCodedMessage                                   `json:"messages"`
	JSON     accountMagicConnectorTelemetryEventListResponseJSON   `json:"-"`
}

// accountMagicConnectorTelemetryEventListResponseJSON contains the JSON metadata
// for the struct [AccountMagicConnectorTelemetryEventListResponse]
type accountMagicConnectorTelemetryEventListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
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
	Result   AccountMagicConnectorTelemetryEventGetResponseResult `json:"result,required"`
	Success  bool                                                 `json:"success,required"`
	Errors   []MconnCodedMessage                                  `json:"errors"`
	Messages []MconnCodedMessage                                  `json:"messages"`
	JSON     accountMagicConnectorTelemetryEventGetResponseJSON   `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseJSON contains the JSON metadata
// for the struct [AccountMagicConnectorTelemetryEventGetResponse]
type accountMagicConnectorTelemetryEventGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
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
// [AccountMagicConnectorTelemetryEventGetResponseResultEInit],
// [AccountMagicConnectorTelemetryEventGetResponseResultELeave],
// [AccountMagicConnectorTelemetryEventGetResponseResultEStartAttestation],
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationSuccess],
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationFailure],
// [AccountMagicConnectorTelemetryEventGetResponseResultEStartRotateCryptKey],
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeySuccess],
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeyFailure],
// [AccountMagicConnectorTelemetryEventGetResponseResultEStartRotatePki],
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiSuccess],
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiFailure],
// [AccountMagicConnectorTelemetryEventGetResponseResultEStartUpgrade],
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeSuccess],
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeFailure],
// [AccountMagicConnectorTelemetryEventGetResponseResultEReconcile],
// [AccountMagicConnectorTelemetryEventGetResponseResultEConfigureCloudflaredTunnel].
func (r AccountMagicConnectorTelemetryEventGetResponseResultE) AsUnion() AccountMagicConnectorTelemetryEventGetResponseResultEUnion {
	return r.union
}

// Union satisfied by [AccountMagicConnectorTelemetryEventGetResponseResultEInit],
// [AccountMagicConnectorTelemetryEventGetResponseResultELeave],
// [AccountMagicConnectorTelemetryEventGetResponseResultEStartAttestation],
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationSuccess],
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationFailure],
// [AccountMagicConnectorTelemetryEventGetResponseResultEStartRotateCryptKey],
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeySuccess],
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeyFailure],
// [AccountMagicConnectorTelemetryEventGetResponseResultEStartRotatePki],
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiSuccess],
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiFailure],
// [AccountMagicConnectorTelemetryEventGetResponseResultEStartUpgrade],
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeSuccess],
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeFailure],
// [AccountMagicConnectorTelemetryEventGetResponseResultEReconcile] or
// [AccountMagicConnectorTelemetryEventGetResponseResultEConfigureCloudflaredTunnel].
type AccountMagicConnectorTelemetryEventGetResponseResultEUnion interface {
	implementsAccountMagicConnectorTelemetryEventGetResponseResultE()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountMagicConnectorTelemetryEventGetResponseResultEUnion)(nil)).Elem(),
		"k",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEInit{}),
			DiscriminatorValue: "Init",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultELeave{}),
			DiscriminatorValue: "Leave",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEStartAttestation{}),
			DiscriminatorValue: "StartAttestation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationSuccess{}),
			DiscriminatorValue: "FinishAttestationSuccess",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationFailure{}),
			DiscriminatorValue: "FinishAttestationFailure",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEStartRotateCryptKey{}),
			DiscriminatorValue: "StartRotateCryptKey",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeySuccess{}),
			DiscriminatorValue: "FinishRotateCryptKeySuccess",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeyFailure{}),
			DiscriminatorValue: "FinishRotateCryptKeyFailure",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEStartRotatePki{}),
			DiscriminatorValue: "StartRotatePki",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiSuccess{}),
			DiscriminatorValue: "FinishRotatePkiSuccess",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiFailure{}),
			DiscriminatorValue: "FinishRotatePkiFailure",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEStartUpgrade{}),
			DiscriminatorValue: "StartUpgrade",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeSuccess{}),
			DiscriminatorValue: "FinishUpgradeSuccess",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeFailure{}),
			DiscriminatorValue: "FinishUpgradeFailure",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEReconcile{}),
			DiscriminatorValue: "Reconcile",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AccountMagicConnectorTelemetryEventGetResponseResultEConfigureCloudflaredTunnel{}),
			DiscriminatorValue: "ConfigureCloudflaredTunnel",
		},
	)
}

type AccountMagicConnectorTelemetryEventGetResponseResultEInit struct {
	// Initialized process
	K    AccountMagicConnectorTelemetryEventGetResponseResultEInitK    `json:"k,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultEInitJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultEInitJSON contains the JSON
// metadata for the struct
// [AccountMagicConnectorTelemetryEventGetResponseResultEInit]
type accountMagicConnectorTelemetryEventGetResponseResultEInitJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultEInit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultEInitJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicConnectorTelemetryEventGetResponseResultEInit) implementsAccountMagicConnectorTelemetryEventGetResponseResultE() {
}

// Initialized process
type AccountMagicConnectorTelemetryEventGetResponseResultEInitK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultEInitKInit AccountMagicConnectorTelemetryEventGetResponseResultEInitK = "Init"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultEInitK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultEInitKInit:
		return true
	}
	return false
}

type AccountMagicConnectorTelemetryEventGetResponseResultELeave struct {
	// Stopped process
	K    AccountMagicConnectorTelemetryEventGetResponseResultELeaveK    `json:"k,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultELeaveJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultELeaveJSON contains the JSON
// metadata for the struct
// [AccountMagicConnectorTelemetryEventGetResponseResultELeave]
type accountMagicConnectorTelemetryEventGetResponseResultELeaveJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultELeave) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultELeaveJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicConnectorTelemetryEventGetResponseResultELeave) implementsAccountMagicConnectorTelemetryEventGetResponseResultE() {
}

// Stopped process
type AccountMagicConnectorTelemetryEventGetResponseResultELeaveK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultELeaveKLeave AccountMagicConnectorTelemetryEventGetResponseResultELeaveK = "Leave"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultELeaveK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultELeaveKLeave:
		return true
	}
	return false
}

type AccountMagicConnectorTelemetryEventGetResponseResultEStartAttestation struct {
	// Started attestation
	K    AccountMagicConnectorTelemetryEventGetResponseResultEStartAttestationK    `json:"k,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultEStartAttestationJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultEStartAttestationJSON
// contains the JSON metadata for the struct
// [AccountMagicConnectorTelemetryEventGetResponseResultEStartAttestation]
type accountMagicConnectorTelemetryEventGetResponseResultEStartAttestationJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultEStartAttestation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultEStartAttestationJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicConnectorTelemetryEventGetResponseResultEStartAttestation) implementsAccountMagicConnectorTelemetryEventGetResponseResultE() {
}

// Started attestation
type AccountMagicConnectorTelemetryEventGetResponseResultEStartAttestationK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultEStartAttestationKStartAttestation AccountMagicConnectorTelemetryEventGetResponseResultEStartAttestationK = "StartAttestation"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultEStartAttestationK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultEStartAttestationKStartAttestation:
		return true
	}
	return false
}

type AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationSuccess struct {
	// Finished attestation
	K    AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationSuccessK    `json:"k,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationSuccessJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationSuccessJSON
// contains the JSON metadata for the struct
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationSuccess]
type accountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationSuccessJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationSuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationSuccessJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationSuccess) implementsAccountMagicConnectorTelemetryEventGetResponseResultE() {
}

// Finished attestation
type AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationSuccessK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationSuccessKFinishAttestationSuccess AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationSuccessK = "FinishAttestationSuccess"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationSuccessK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationSuccessKFinishAttestationSuccess:
		return true
	}
	return false
}

type AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationFailure struct {
	// Failed attestation
	K    AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationFailureK    `json:"k,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationFailureJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationFailureJSON
// contains the JSON metadata for the struct
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationFailure]
type accountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationFailureJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationFailure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationFailureJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationFailure) implementsAccountMagicConnectorTelemetryEventGetResponseResultE() {
}

// Failed attestation
type AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationFailureK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationFailureKFinishAttestationFailure AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationFailureK = "FinishAttestationFailure"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationFailureK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultEFinishAttestationFailureKFinishAttestationFailure:
		return true
	}
	return false
}

type AccountMagicConnectorTelemetryEventGetResponseResultEStartRotateCryptKey struct {
	// Started crypt key rotation
	K    AccountMagicConnectorTelemetryEventGetResponseResultEStartRotateCryptKeyK    `json:"k,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultEStartRotateCryptKeyJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultEStartRotateCryptKeyJSON
// contains the JSON metadata for the struct
// [AccountMagicConnectorTelemetryEventGetResponseResultEStartRotateCryptKey]
type accountMagicConnectorTelemetryEventGetResponseResultEStartRotateCryptKeyJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultEStartRotateCryptKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultEStartRotateCryptKeyJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicConnectorTelemetryEventGetResponseResultEStartRotateCryptKey) implementsAccountMagicConnectorTelemetryEventGetResponseResultE() {
}

// Started crypt key rotation
type AccountMagicConnectorTelemetryEventGetResponseResultEStartRotateCryptKeyK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultEStartRotateCryptKeyKStartRotateCryptKey AccountMagicConnectorTelemetryEventGetResponseResultEStartRotateCryptKeyK = "StartRotateCryptKey"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultEStartRotateCryptKeyK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultEStartRotateCryptKeyKStartRotateCryptKey:
		return true
	}
	return false
}

type AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeySuccess struct {
	// Finished crypt key rotation
	K    AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeySuccessK    `json:"k,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeySuccessJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeySuccessJSON
// contains the JSON metadata for the struct
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeySuccess]
type accountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeySuccessJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeySuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeySuccessJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeySuccess) implementsAccountMagicConnectorTelemetryEventGetResponseResultE() {
}

// Finished crypt key rotation
type AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeySuccessK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeySuccessKFinishRotateCryptKeySuccess AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeySuccessK = "FinishRotateCryptKeySuccess"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeySuccessK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeySuccessKFinishRotateCryptKeySuccess:
		return true
	}
	return false
}

type AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeyFailure struct {
	// Failed crypt key rotation
	K    AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeyFailureK    `json:"k,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeyFailureJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeyFailureJSON
// contains the JSON metadata for the struct
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeyFailure]
type accountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeyFailureJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeyFailure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeyFailureJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeyFailure) implementsAccountMagicConnectorTelemetryEventGetResponseResultE() {
}

// Failed crypt key rotation
type AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeyFailureK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeyFailureKFinishRotateCryptKeyFailure AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeyFailureK = "FinishRotateCryptKeyFailure"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeyFailureK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotateCryptKeyFailureKFinishRotateCryptKeyFailure:
		return true
	}
	return false
}

type AccountMagicConnectorTelemetryEventGetResponseResultEStartRotatePki struct {
	// Started PKI rotation
	K    AccountMagicConnectorTelemetryEventGetResponseResultEStartRotatePkiK    `json:"k,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultEStartRotatePkiJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultEStartRotatePkiJSON contains
// the JSON metadata for the struct
// [AccountMagicConnectorTelemetryEventGetResponseResultEStartRotatePki]
type accountMagicConnectorTelemetryEventGetResponseResultEStartRotatePkiJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultEStartRotatePki) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultEStartRotatePkiJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicConnectorTelemetryEventGetResponseResultEStartRotatePki) implementsAccountMagicConnectorTelemetryEventGetResponseResultE() {
}

// Started PKI rotation
type AccountMagicConnectorTelemetryEventGetResponseResultEStartRotatePkiK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultEStartRotatePkiKStartRotatePki AccountMagicConnectorTelemetryEventGetResponseResultEStartRotatePkiK = "StartRotatePki"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultEStartRotatePkiK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultEStartRotatePkiKStartRotatePki:
		return true
	}
	return false
}

type AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiSuccess struct {
	// Finished PKI rotation
	K    AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiSuccessK    `json:"k,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiSuccessJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiSuccessJSON
// contains the JSON metadata for the struct
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiSuccess]
type accountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiSuccessJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiSuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiSuccessJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiSuccess) implementsAccountMagicConnectorTelemetryEventGetResponseResultE() {
}

// Finished PKI rotation
type AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiSuccessK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiSuccessKFinishRotatePkiSuccess AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiSuccessK = "FinishRotatePkiSuccess"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiSuccessK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiSuccessKFinishRotatePkiSuccess:
		return true
	}
	return false
}

type AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiFailure struct {
	// Failed PKI rotation
	K    AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiFailureK    `json:"k,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiFailureJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiFailureJSON
// contains the JSON metadata for the struct
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiFailure]
type accountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiFailureJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiFailure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiFailureJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiFailure) implementsAccountMagicConnectorTelemetryEventGetResponseResultE() {
}

// Failed PKI rotation
type AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiFailureK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiFailureKFinishRotatePkiFailure AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiFailureK = "FinishRotatePkiFailure"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiFailureK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultEFinishRotatePkiFailureKFinishRotatePkiFailure:
		return true
	}
	return false
}

type AccountMagicConnectorTelemetryEventGetResponseResultEStartUpgrade struct {
	// Started upgrade
	K AccountMagicConnectorTelemetryEventGetResponseResultEStartUpgradeK `json:"k,required"`
	// Location of upgrade bundle
	URL  string                                                                `json:"url,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultEStartUpgradeJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultEStartUpgradeJSON contains
// the JSON metadata for the struct
// [AccountMagicConnectorTelemetryEventGetResponseResultEStartUpgrade]
type accountMagicConnectorTelemetryEventGetResponseResultEStartUpgradeJSON struct {
	K           apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultEStartUpgrade) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultEStartUpgradeJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicConnectorTelemetryEventGetResponseResultEStartUpgrade) implementsAccountMagicConnectorTelemetryEventGetResponseResultE() {
}

// Started upgrade
type AccountMagicConnectorTelemetryEventGetResponseResultEStartUpgradeK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultEStartUpgradeKStartUpgrade AccountMagicConnectorTelemetryEventGetResponseResultEStartUpgradeK = "StartUpgrade"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultEStartUpgradeK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultEStartUpgradeKStartUpgrade:
		return true
	}
	return false
}

type AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeSuccess struct {
	// Finished upgrade
	K    AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeSuccessK    `json:"k,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeSuccessJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeSuccessJSON
// contains the JSON metadata for the struct
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeSuccess]
type accountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeSuccessJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeSuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeSuccessJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeSuccess) implementsAccountMagicConnectorTelemetryEventGetResponseResultE() {
}

// Finished upgrade
type AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeSuccessK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeSuccessKFinishUpgradeSuccess AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeSuccessK = "FinishUpgradeSuccess"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeSuccessK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeSuccessKFinishUpgradeSuccess:
		return true
	}
	return false
}

type AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeFailure struct {
	// Failed upgrade
	K    AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeFailureK    `json:"k,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeFailureJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeFailureJSON
// contains the JSON metadata for the struct
// [AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeFailure]
type accountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeFailureJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeFailure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeFailureJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeFailure) implementsAccountMagicConnectorTelemetryEventGetResponseResultE() {
}

// Failed upgrade
type AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeFailureK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeFailureKFinishUpgradeFailure AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeFailureK = "FinishUpgradeFailure"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeFailureK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultEFinishUpgradeFailureKFinishUpgradeFailure:
		return true
	}
	return false
}

type AccountMagicConnectorTelemetryEventGetResponseResultEReconcile struct {
	// Reconciled
	K    AccountMagicConnectorTelemetryEventGetResponseResultEReconcileK    `json:"k,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultEReconcileJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultEReconcileJSON contains the
// JSON metadata for the struct
// [AccountMagicConnectorTelemetryEventGetResponseResultEReconcile]
type accountMagicConnectorTelemetryEventGetResponseResultEReconcileJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultEReconcile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultEReconcileJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicConnectorTelemetryEventGetResponseResultEReconcile) implementsAccountMagicConnectorTelemetryEventGetResponseResultE() {
}

// Reconciled
type AccountMagicConnectorTelemetryEventGetResponseResultEReconcileK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultEReconcileKReconcile AccountMagicConnectorTelemetryEventGetResponseResultEReconcileK = "Reconcile"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultEReconcileK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultEReconcileKReconcile:
		return true
	}
	return false
}

type AccountMagicConnectorTelemetryEventGetResponseResultEConfigureCloudflaredTunnel struct {
	// Configured Cloudflared tunnel
	K    AccountMagicConnectorTelemetryEventGetResponseResultEConfigureCloudflaredTunnelK    `json:"k,required"`
	JSON accountMagicConnectorTelemetryEventGetResponseResultEConfigureCloudflaredTunnelJSON `json:"-"`
}

// accountMagicConnectorTelemetryEventGetResponseResultEConfigureCloudflaredTunnelJSON
// contains the JSON metadata for the struct
// [AccountMagicConnectorTelemetryEventGetResponseResultEConfigureCloudflaredTunnel]
type accountMagicConnectorTelemetryEventGetResponseResultEConfigureCloudflaredTunnelJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetryEventGetResponseResultEConfigureCloudflaredTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetryEventGetResponseResultEConfigureCloudflaredTunnelJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicConnectorTelemetryEventGetResponseResultEConfigureCloudflaredTunnel) implementsAccountMagicConnectorTelemetryEventGetResponseResultE() {
}

// Configured Cloudflared tunnel
type AccountMagicConnectorTelemetryEventGetResponseResultEConfigureCloudflaredTunnelK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultEConfigureCloudflaredTunnelKConfigureCloudflaredTunnel AccountMagicConnectorTelemetryEventGetResponseResultEConfigureCloudflaredTunnelK = "ConfigureCloudflaredTunnel"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultEConfigureCloudflaredTunnelK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultEConfigureCloudflaredTunnelKConfigureCloudflaredTunnel:
		return true
	}
	return false
}

// Initialized process
type AccountMagicConnectorTelemetryEventGetResponseResultEK string

const (
	AccountMagicConnectorTelemetryEventGetResponseResultEKInit                        AccountMagicConnectorTelemetryEventGetResponseResultEK = "Init"
	AccountMagicConnectorTelemetryEventGetResponseResultEKLeave                       AccountMagicConnectorTelemetryEventGetResponseResultEK = "Leave"
	AccountMagicConnectorTelemetryEventGetResponseResultEKStartAttestation            AccountMagicConnectorTelemetryEventGetResponseResultEK = "StartAttestation"
	AccountMagicConnectorTelemetryEventGetResponseResultEKFinishAttestationSuccess    AccountMagicConnectorTelemetryEventGetResponseResultEK = "FinishAttestationSuccess"
	AccountMagicConnectorTelemetryEventGetResponseResultEKFinishAttestationFailure    AccountMagicConnectorTelemetryEventGetResponseResultEK = "FinishAttestationFailure"
	AccountMagicConnectorTelemetryEventGetResponseResultEKStartRotateCryptKey         AccountMagicConnectorTelemetryEventGetResponseResultEK = "StartRotateCryptKey"
	AccountMagicConnectorTelemetryEventGetResponseResultEKFinishRotateCryptKeySuccess AccountMagicConnectorTelemetryEventGetResponseResultEK = "FinishRotateCryptKeySuccess"
	AccountMagicConnectorTelemetryEventGetResponseResultEKFinishRotateCryptKeyFailure AccountMagicConnectorTelemetryEventGetResponseResultEK = "FinishRotateCryptKeyFailure"
	AccountMagicConnectorTelemetryEventGetResponseResultEKStartRotatePki              AccountMagicConnectorTelemetryEventGetResponseResultEK = "StartRotatePki"
	AccountMagicConnectorTelemetryEventGetResponseResultEKFinishRotatePkiSuccess      AccountMagicConnectorTelemetryEventGetResponseResultEK = "FinishRotatePkiSuccess"
	AccountMagicConnectorTelemetryEventGetResponseResultEKFinishRotatePkiFailure      AccountMagicConnectorTelemetryEventGetResponseResultEK = "FinishRotatePkiFailure"
	AccountMagicConnectorTelemetryEventGetResponseResultEKStartUpgrade                AccountMagicConnectorTelemetryEventGetResponseResultEK = "StartUpgrade"
	AccountMagicConnectorTelemetryEventGetResponseResultEKFinishUpgradeSuccess        AccountMagicConnectorTelemetryEventGetResponseResultEK = "FinishUpgradeSuccess"
	AccountMagicConnectorTelemetryEventGetResponseResultEKFinishUpgradeFailure        AccountMagicConnectorTelemetryEventGetResponseResultEK = "FinishUpgradeFailure"
	AccountMagicConnectorTelemetryEventGetResponseResultEKReconcile                   AccountMagicConnectorTelemetryEventGetResponseResultEK = "Reconcile"
	AccountMagicConnectorTelemetryEventGetResponseResultEKConfigureCloudflaredTunnel  AccountMagicConnectorTelemetryEventGetResponseResultEK = "ConfigureCloudflaredTunnel"
)

func (r AccountMagicConnectorTelemetryEventGetResponseResultEK) IsKnown() bool {
	switch r {
	case AccountMagicConnectorTelemetryEventGetResponseResultEKInit, AccountMagicConnectorTelemetryEventGetResponseResultEKLeave, AccountMagicConnectorTelemetryEventGetResponseResultEKStartAttestation, AccountMagicConnectorTelemetryEventGetResponseResultEKFinishAttestationSuccess, AccountMagicConnectorTelemetryEventGetResponseResultEKFinishAttestationFailure, AccountMagicConnectorTelemetryEventGetResponseResultEKStartRotateCryptKey, AccountMagicConnectorTelemetryEventGetResponseResultEKFinishRotateCryptKeySuccess, AccountMagicConnectorTelemetryEventGetResponseResultEKFinishRotateCryptKeyFailure, AccountMagicConnectorTelemetryEventGetResponseResultEKStartRotatePki, AccountMagicConnectorTelemetryEventGetResponseResultEKFinishRotatePkiSuccess, AccountMagicConnectorTelemetryEventGetResponseResultEKFinishRotatePkiFailure, AccountMagicConnectorTelemetryEventGetResponseResultEKStartUpgrade, AccountMagicConnectorTelemetryEventGetResponseResultEKFinishUpgradeSuccess, AccountMagicConnectorTelemetryEventGetResponseResultEKFinishUpgradeFailure, AccountMagicConnectorTelemetryEventGetResponseResultEKReconcile, AccountMagicConnectorTelemetryEventGetResponseResultEKConfigureCloudflaredTunnel:
		return true
	}
	return false
}

type AccountMagicConnectorTelemetryEventListParams struct {
	From   param.Field[float64] `query:"from,required"`
	To     param.Field[float64] `query:"to,required"`
	Cursor param.Field[string]  `query:"cursor"`
	// Filter by event kind
	K     param.Field[string]  `query:"k"`
	Limit param.Field[float64] `query:"limit"`
}

// URLQuery serializes [AccountMagicConnectorTelemetryEventListParams]'s query
// parameters as `url.Values`.
func (r AccountMagicConnectorTelemetryEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
