// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/tidwall/gjson"
)

// AccountR2BucketLockService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountR2BucketLockService] method instead.
type AccountR2BucketLockService struct {
	Options []option.RequestOption
}

// NewAccountR2BucketLockService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountR2BucketLockService(opts ...option.RequestOption) (r *AccountR2BucketLockService) {
	r = &AccountR2BucketLockService{}
	r.Options = opts
	return
}

// Get lock rules for a bucket
func (r *AccountR2BucketLockService) Get(ctx context.Context, accountID string, bucketName string, query AccountR2BucketLockGetParams, opts ...option.RequestOption) (res *AccountR2BucketLockGetResponse, err error) {
	if query.CfR2Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", query.CfR2Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/lock", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set lock rules for a bucket
func (r *AccountR2BucketLockService) Update(ctx context.Context, accountID string, bucketName string, params AccountR2BucketLockUpdateParams, opts ...option.RequestOption) (res *AccountR2BucketLockUpdateResponse, err error) {
	if params.CfR2Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.CfR2Jurisdiction)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/lock", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type R2BucketLockRule struct {
	// Unique identifier for this rule
	ID string `json:"id,required"`
	// Condition to apply a lock rule to an object for how long in seconds
	Condition R2BucketLockRuleCondition `json:"condition,required"`
	// Whether or not this rule is in effect
	Enabled bool `json:"enabled,required"`
	// Rule will only apply to objects/uploads in the bucket that start with the given
	// prefix, an empty prefix can be provided to scope rule to all objects/uploads
	Prefix string               `json:"prefix"`
	JSON   r2BucketLockRuleJSON `json:"-"`
}

// r2BucketLockRuleJSON contains the JSON metadata for the struct
// [R2BucketLockRule]
type r2BucketLockRuleJSON struct {
	ID          apijson.Field
	Condition   apijson.Field
	Enabled     apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketLockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2BucketLockRuleJSON) RawJSON() string {
	return r.raw
}

// Condition to apply a lock rule to an object for how long in seconds
type R2BucketLockRuleCondition struct {
	Type          R2BucketLockRuleConditionType `json:"type,required"`
	Date          time.Time                     `json:"date" format:"date"`
	MaxAgeSeconds int64                         `json:"maxAgeSeconds"`
	JSON          r2BucketLockRuleConditionJSON `json:"-"`
	union         R2BucketLockRuleConditionUnion
}

// r2BucketLockRuleConditionJSON contains the JSON metadata for the struct
// [R2BucketLockRuleCondition]
type r2BucketLockRuleConditionJSON struct {
	Type          apijson.Field
	Date          apijson.Field
	MaxAgeSeconds apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r r2BucketLockRuleConditionJSON) RawJSON() string {
	return r.raw
}

func (r *R2BucketLockRuleCondition) UnmarshalJSON(data []byte) (err error) {
	*r = R2BucketLockRuleCondition{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [R2BucketLockRuleConditionUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [R2BucketLockRuleConditionR2LockRuleAgeCondition],
// [R2BucketLockRuleConditionR2LockRuleDateCondition],
// [R2BucketLockRuleConditionR2LockRuleIndefiniteCondition].
func (r R2BucketLockRuleCondition) AsUnion() R2BucketLockRuleConditionUnion {
	return r.union
}

// Condition to apply a lock rule to an object for how long in seconds
//
// Union satisfied by [R2BucketLockRuleConditionR2LockRuleAgeCondition],
// [R2BucketLockRuleConditionR2LockRuleDateCondition] or
// [R2BucketLockRuleConditionR2LockRuleIndefiniteCondition].
type R2BucketLockRuleConditionUnion interface {
	implementsR2BucketLockRuleCondition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*R2BucketLockRuleConditionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(R2BucketLockRuleConditionR2LockRuleAgeCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(R2BucketLockRuleConditionR2LockRuleDateCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(R2BucketLockRuleConditionR2LockRuleIndefiniteCondition{}),
		},
	)
}

// Condition to apply a lock rule to an object for how long in seconds
type R2BucketLockRuleConditionR2LockRuleAgeCondition struct {
	MaxAgeSeconds int64                                               `json:"maxAgeSeconds,required"`
	Type          R2BucketLockRuleConditionR2LockRuleAgeConditionType `json:"type,required"`
	JSON          r2BucketLockRuleConditionR2LockRuleAgeConditionJSON `json:"-"`
}

// r2BucketLockRuleConditionR2LockRuleAgeConditionJSON contains the JSON metadata
// for the struct [R2BucketLockRuleConditionR2LockRuleAgeCondition]
type r2BucketLockRuleConditionR2LockRuleAgeConditionJSON struct {
	MaxAgeSeconds apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *R2BucketLockRuleConditionR2LockRuleAgeCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2BucketLockRuleConditionR2LockRuleAgeConditionJSON) RawJSON() string {
	return r.raw
}

func (r R2BucketLockRuleConditionR2LockRuleAgeCondition) implementsR2BucketLockRuleCondition() {}

type R2BucketLockRuleConditionR2LockRuleAgeConditionType string

const (
	R2BucketLockRuleConditionR2LockRuleAgeConditionTypeAge R2BucketLockRuleConditionR2LockRuleAgeConditionType = "Age"
)

func (r R2BucketLockRuleConditionR2LockRuleAgeConditionType) IsKnown() bool {
	switch r {
	case R2BucketLockRuleConditionR2LockRuleAgeConditionTypeAge:
		return true
	}
	return false
}

// Condition to apply a lock rule to an object until a specific date
type R2BucketLockRuleConditionR2LockRuleDateCondition struct {
	Date time.Time                                            `json:"date,required" format:"date"`
	Type R2BucketLockRuleConditionR2LockRuleDateConditionType `json:"type,required"`
	JSON r2BucketLockRuleConditionR2LockRuleDateConditionJSON `json:"-"`
}

// r2BucketLockRuleConditionR2LockRuleDateConditionJSON contains the JSON metadata
// for the struct [R2BucketLockRuleConditionR2LockRuleDateCondition]
type r2BucketLockRuleConditionR2LockRuleDateConditionJSON struct {
	Date        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketLockRuleConditionR2LockRuleDateCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2BucketLockRuleConditionR2LockRuleDateConditionJSON) RawJSON() string {
	return r.raw
}

func (r R2BucketLockRuleConditionR2LockRuleDateCondition) implementsR2BucketLockRuleCondition() {}

type R2BucketLockRuleConditionR2LockRuleDateConditionType string

const (
	R2BucketLockRuleConditionR2LockRuleDateConditionTypeDate R2BucketLockRuleConditionR2LockRuleDateConditionType = "Date"
)

func (r R2BucketLockRuleConditionR2LockRuleDateConditionType) IsKnown() bool {
	switch r {
	case R2BucketLockRuleConditionR2LockRuleDateConditionTypeDate:
		return true
	}
	return false
}

// Condition to apply a lock rule indefinitely
type R2BucketLockRuleConditionR2LockRuleIndefiniteCondition struct {
	Type R2BucketLockRuleConditionR2LockRuleIndefiniteConditionType `json:"type,required"`
	JSON r2BucketLockRuleConditionR2LockRuleIndefiniteConditionJSON `json:"-"`
}

// r2BucketLockRuleConditionR2LockRuleIndefiniteConditionJSON contains the JSON
// metadata for the struct [R2BucketLockRuleConditionR2LockRuleIndefiniteCondition]
type r2BucketLockRuleConditionR2LockRuleIndefiniteConditionJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketLockRuleConditionR2LockRuleIndefiniteCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2BucketLockRuleConditionR2LockRuleIndefiniteConditionJSON) RawJSON() string {
	return r.raw
}

func (r R2BucketLockRuleConditionR2LockRuleIndefiniteCondition) implementsR2BucketLockRuleCondition() {
}

type R2BucketLockRuleConditionR2LockRuleIndefiniteConditionType string

const (
	R2BucketLockRuleConditionR2LockRuleIndefiniteConditionTypeIndefinite R2BucketLockRuleConditionR2LockRuleIndefiniteConditionType = "Indefinite"
)

func (r R2BucketLockRuleConditionR2LockRuleIndefiniteConditionType) IsKnown() bool {
	switch r {
	case R2BucketLockRuleConditionR2LockRuleIndefiniteConditionTypeIndefinite:
		return true
	}
	return false
}

type R2BucketLockRuleConditionType string

const (
	R2BucketLockRuleConditionTypeAge        R2BucketLockRuleConditionType = "Age"
	R2BucketLockRuleConditionTypeDate       R2BucketLockRuleConditionType = "Date"
	R2BucketLockRuleConditionTypeIndefinite R2BucketLockRuleConditionType = "Indefinite"
)

func (r R2BucketLockRuleConditionType) IsKnown() bool {
	switch r {
	case R2BucketLockRuleConditionTypeAge, R2BucketLockRuleConditionTypeDate, R2BucketLockRuleConditionTypeIndefinite:
		return true
	}
	return false
}

type R2BucketLockRuleParam struct {
	// Unique identifier for this rule
	ID param.Field[string] `json:"id,required"`
	// Condition to apply a lock rule to an object for how long in seconds
	Condition param.Field[R2BucketLockRuleConditionUnionParam] `json:"condition,required"`
	// Whether or not this rule is in effect
	Enabled param.Field[bool] `json:"enabled,required"`
	// Rule will only apply to objects/uploads in the bucket that start with the given
	// prefix, an empty prefix can be provided to scope rule to all objects/uploads
	Prefix param.Field[string] `json:"prefix"`
}

func (r R2BucketLockRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Condition to apply a lock rule to an object for how long in seconds
type R2BucketLockRuleConditionParam struct {
	Type          param.Field[R2BucketLockRuleConditionType] `json:"type,required"`
	Date          param.Field[time.Time]                     `json:"date" format:"date"`
	MaxAgeSeconds param.Field[int64]                         `json:"maxAgeSeconds"`
}

func (r R2BucketLockRuleConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r R2BucketLockRuleConditionParam) implementsR2BucketLockRuleConditionUnionParam() {}

// Condition to apply a lock rule to an object for how long in seconds
//
// Satisfied by [R2BucketLockRuleConditionR2LockRuleAgeConditionParam],
// [R2BucketLockRuleConditionR2LockRuleDateConditionParam],
// [R2BucketLockRuleConditionR2LockRuleIndefiniteConditionParam],
// [R2BucketLockRuleConditionParam].
type R2BucketLockRuleConditionUnionParam interface {
	implementsR2BucketLockRuleConditionUnionParam()
}

// Condition to apply a lock rule to an object for how long in seconds
type R2BucketLockRuleConditionR2LockRuleAgeConditionParam struct {
	MaxAgeSeconds param.Field[int64]                                               `json:"maxAgeSeconds,required"`
	Type          param.Field[R2BucketLockRuleConditionR2LockRuleAgeConditionType] `json:"type,required"`
}

func (r R2BucketLockRuleConditionR2LockRuleAgeConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r R2BucketLockRuleConditionR2LockRuleAgeConditionParam) implementsR2BucketLockRuleConditionUnionParam() {
}

// Condition to apply a lock rule to an object until a specific date
type R2BucketLockRuleConditionR2LockRuleDateConditionParam struct {
	Date param.Field[time.Time]                                            `json:"date,required" format:"date"`
	Type param.Field[R2BucketLockRuleConditionR2LockRuleDateConditionType] `json:"type,required"`
}

func (r R2BucketLockRuleConditionR2LockRuleDateConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r R2BucketLockRuleConditionR2LockRuleDateConditionParam) implementsR2BucketLockRuleConditionUnionParam() {
}

// Condition to apply a lock rule indefinitely
type R2BucketLockRuleConditionR2LockRuleIndefiniteConditionParam struct {
	Type param.Field[R2BucketLockRuleConditionR2LockRuleIndefiniteConditionType] `json:"type,required"`
}

func (r R2BucketLockRuleConditionR2LockRuleIndefiniteConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r R2BucketLockRuleConditionR2LockRuleIndefiniteConditionParam) implementsR2BucketLockRuleConditionUnionParam() {
}

type AccountR2BucketLockGetResponse struct {
	Result AccountR2BucketLockGetResponseResult `json:"result"`
	JSON   accountR2BucketLockGetResponseJSON   `json:"-"`
	R2V4Response
}

// accountR2BucketLockGetResponseJSON contains the JSON metadata for the struct
// [AccountR2BucketLockGetResponse]
type accountR2BucketLockGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketLockGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketLockGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketLockGetResponseResult struct {
	Rules []R2BucketLockRule                       `json:"rules"`
	JSON  accountR2BucketLockGetResponseResultJSON `json:"-"`
}

// accountR2BucketLockGetResponseResultJSON contains the JSON metadata for the
// struct [AccountR2BucketLockGetResponseResult]
type accountR2BucketLockGetResponseResultJSON struct {
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketLockGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketLockGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketLockUpdateResponse struct {
	JSON accountR2BucketLockUpdateResponseJSON `json:"-"`
	R2V4Response
}

// accountR2BucketLockUpdateResponseJSON contains the JSON metadata for the struct
// [AccountR2BucketLockUpdateResponse]
type accountR2BucketLockUpdateResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketLockUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketLockUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketLockGetParams struct {
	// The bucket jurisdiction
	CfR2Jurisdiction param.Field[AccountR2BucketLockGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// The bucket jurisdiction
type AccountR2BucketLockGetParamsCfR2Jurisdiction string

const (
	AccountR2BucketLockGetParamsCfR2JurisdictionDefault AccountR2BucketLockGetParamsCfR2Jurisdiction = "default"
	AccountR2BucketLockGetParamsCfR2JurisdictionEu      AccountR2BucketLockGetParamsCfR2Jurisdiction = "eu"
	AccountR2BucketLockGetParamsCfR2JurisdictionFedramp AccountR2BucketLockGetParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketLockGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketLockGetParamsCfR2JurisdictionDefault, AccountR2BucketLockGetParamsCfR2JurisdictionEu, AccountR2BucketLockGetParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type AccountR2BucketLockUpdateParams struct {
	Rules param.Field[[]R2BucketLockRuleParam] `json:"rules"`
	// The bucket jurisdiction
	CfR2Jurisdiction param.Field[AccountR2BucketLockUpdateParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r AccountR2BucketLockUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The bucket jurisdiction
type AccountR2BucketLockUpdateParamsCfR2Jurisdiction string

const (
	AccountR2BucketLockUpdateParamsCfR2JurisdictionDefault AccountR2BucketLockUpdateParamsCfR2Jurisdiction = "default"
	AccountR2BucketLockUpdateParamsCfR2JurisdictionEu      AccountR2BucketLockUpdateParamsCfR2Jurisdiction = "eu"
	AccountR2BucketLockUpdateParamsCfR2JurisdictionFedramp AccountR2BucketLockUpdateParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketLockUpdateParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketLockUpdateParamsCfR2JurisdictionDefault, AccountR2BucketLockUpdateParamsCfR2JurisdictionEu, AccountR2BucketLockUpdateParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}
