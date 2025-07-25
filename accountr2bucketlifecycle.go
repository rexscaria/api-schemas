// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountR2BucketLifecycleService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountR2BucketLifecycleService] method instead.
type AccountR2BucketLifecycleService struct {
	Options []option.RequestOption
}

// NewAccountR2BucketLifecycleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountR2BucketLifecycleService(opts ...option.RequestOption) (r *AccountR2BucketLifecycleService) {
	r = &AccountR2BucketLifecycleService{}
	r.Options = opts
	return
}

// Get object lifecycle rules for a bucket.
func (r *AccountR2BucketLifecycleService) Get(ctx context.Context, accountID string, bucketName string, query AccountR2BucketLifecycleGetParams, opts ...option.RequestOption) (res *AccountR2BucketLifecycleGetResponse, err error) {
	if query.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", query.Jurisdiction)))
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/lifecycle", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set the object lifecycle rules for a bucket.
func (r *AccountR2BucketLifecycleService) Update(ctx context.Context, accountID string, bucketName string, params AccountR2BucketLifecycleUpdateParams, opts ...option.RequestOption) (res *R2V4Response, err error) {
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/lifecycle", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type R2LifecycleAgeCondition struct {
	MaxAge int64                       `json:"maxAge,required"`
	Type   R2LifecycleAgeConditionType `json:"type,required"`
	JSON   r2LifecycleAgeConditionJSON `json:"-"`
}

// r2LifecycleAgeConditionJSON contains the JSON metadata for the struct
// [R2LifecycleAgeCondition]
type r2LifecycleAgeConditionJSON struct {
	MaxAge      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2LifecycleAgeCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2LifecycleAgeConditionJSON) RawJSON() string {
	return r.raw
}

func (r R2LifecycleAgeCondition) implementsR2LifecycleRuleDeleteObjectsTransitionCondition() {}

func (r R2LifecycleAgeCondition) implementsR2LifecycleRuleStorageClassTransitionsCondition() {}

type R2LifecycleAgeConditionType string

const (
	R2LifecycleAgeConditionTypeAge R2LifecycleAgeConditionType = "Age"
)

func (r R2LifecycleAgeConditionType) IsKnown() bool {
	switch r {
	case R2LifecycleAgeConditionTypeAge:
		return true
	}
	return false
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type R2LifecycleAgeConditionParam struct {
	MaxAge param.Field[int64]                       `json:"maxAge,required"`
	Type   param.Field[R2LifecycleAgeConditionType] `json:"type,required"`
}

func (r R2LifecycleAgeConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r R2LifecycleAgeConditionParam) implementsR2LifecycleRuleDeleteObjectsTransitionConditionUnionParam() {
}

func (r R2LifecycleAgeConditionParam) implementsR2LifecycleRuleStorageClassTransitionsConditionUnionParam() {
}

// Condition for lifecycle transitions to apply on a specific date.
type R2LifecycleDateCondition struct {
	Date time.Time                    `json:"date,required" format:"date"`
	Type R2LifecycleDateConditionType `json:"type,required"`
	JSON r2LifecycleDateConditionJSON `json:"-"`
}

// r2LifecycleDateConditionJSON contains the JSON metadata for the struct
// [R2LifecycleDateCondition]
type r2LifecycleDateConditionJSON struct {
	Date        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2LifecycleDateCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2LifecycleDateConditionJSON) RawJSON() string {
	return r.raw
}

func (r R2LifecycleDateCondition) implementsR2LifecycleRuleDeleteObjectsTransitionCondition() {}

func (r R2LifecycleDateCondition) implementsR2LifecycleRuleStorageClassTransitionsCondition() {}

type R2LifecycleDateConditionType string

const (
	R2LifecycleDateConditionTypeDate R2LifecycleDateConditionType = "Date"
)

func (r R2LifecycleDateConditionType) IsKnown() bool {
	switch r {
	case R2LifecycleDateConditionTypeDate:
		return true
	}
	return false
}

// Condition for lifecycle transitions to apply on a specific date.
type R2LifecycleDateConditionParam struct {
	Date param.Field[time.Time]                    `json:"date,required" format:"date"`
	Type param.Field[R2LifecycleDateConditionType] `json:"type,required"`
}

func (r R2LifecycleDateConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r R2LifecycleDateConditionParam) implementsR2LifecycleRuleDeleteObjectsTransitionConditionUnionParam() {
}

func (r R2LifecycleDateConditionParam) implementsR2LifecycleRuleStorageClassTransitionsConditionUnionParam() {
}

type R2LifecycleRule struct {
	// Unique identifier for this rule.
	ID string `json:"id,required"`
	// Conditions that apply to all transitions of this rule.
	Conditions R2LifecycleRuleConditions `json:"conditions,required"`
	// Whether or not this rule is in effect.
	Enabled bool `json:"enabled,required"`
	// Transition to abort ongoing multipart uploads.
	AbortMultipartUploadsTransition R2LifecycleRuleAbortMultipartUploadsTransition `json:"abortMultipartUploadsTransition"`
	// Transition to delete objects.
	DeleteObjectsTransition R2LifecycleRuleDeleteObjectsTransition `json:"deleteObjectsTransition"`
	// Transitions to change the storage class of objects.
	StorageClassTransitions []R2LifecycleRuleStorageClassTransition `json:"storageClassTransitions"`
	JSON                    r2LifecycleRuleJSON                     `json:"-"`
}

// r2LifecycleRuleJSON contains the JSON metadata for the struct [R2LifecycleRule]
type r2LifecycleRuleJSON struct {
	ID                              apijson.Field
	Conditions                      apijson.Field
	Enabled                         apijson.Field
	AbortMultipartUploadsTransition apijson.Field
	DeleteObjectsTransition         apijson.Field
	StorageClassTransitions         apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *R2LifecycleRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2LifecycleRuleJSON) RawJSON() string {
	return r.raw
}

// Conditions that apply to all transitions of this rule.
type R2LifecycleRuleConditions struct {
	// Transitions will only apply to objects/uploads in the bucket that start with the
	// given prefix, an empty prefix can be provided to scope rule to all
	// objects/uploads.
	Prefix string                        `json:"prefix,required"`
	JSON   r2LifecycleRuleConditionsJSON `json:"-"`
}

// r2LifecycleRuleConditionsJSON contains the JSON metadata for the struct
// [R2LifecycleRuleConditions]
type r2LifecycleRuleConditionsJSON struct {
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2LifecycleRuleConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2LifecycleRuleConditionsJSON) RawJSON() string {
	return r.raw
}

// Transition to abort ongoing multipart uploads.
type R2LifecycleRuleAbortMultipartUploadsTransition struct {
	// Condition for lifecycle transitions to apply after an object reaches an age in
	// seconds.
	Condition R2LifecycleAgeCondition                            `json:"condition"`
	JSON      r2LifecycleRuleAbortMultipartUploadsTransitionJSON `json:"-"`
}

// r2LifecycleRuleAbortMultipartUploadsTransitionJSON contains the JSON metadata
// for the struct [R2LifecycleRuleAbortMultipartUploadsTransition]
type r2LifecycleRuleAbortMultipartUploadsTransitionJSON struct {
	Condition   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2LifecycleRuleAbortMultipartUploadsTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2LifecycleRuleAbortMultipartUploadsTransitionJSON) RawJSON() string {
	return r.raw
}

// Transition to delete objects.
type R2LifecycleRuleDeleteObjectsTransition struct {
	// Condition for lifecycle transitions to apply after an object reaches an age in
	// seconds.
	Condition R2LifecycleRuleDeleteObjectsTransitionCondition `json:"condition"`
	JSON      r2LifecycleRuleDeleteObjectsTransitionJSON      `json:"-"`
}

// r2LifecycleRuleDeleteObjectsTransitionJSON contains the JSON metadata for the
// struct [R2LifecycleRuleDeleteObjectsTransition]
type r2LifecycleRuleDeleteObjectsTransitionJSON struct {
	Condition   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2LifecycleRuleDeleteObjectsTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2LifecycleRuleDeleteObjectsTransitionJSON) RawJSON() string {
	return r.raw
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type R2LifecycleRuleDeleteObjectsTransitionCondition struct {
	Type   R2LifecycleRuleDeleteObjectsTransitionConditionType `json:"type,required"`
	Date   time.Time                                           `json:"date" format:"date"`
	MaxAge int64                                               `json:"maxAge"`
	JSON   r2LifecycleRuleDeleteObjectsTransitionConditionJSON `json:"-"`
	union  R2LifecycleRuleDeleteObjectsTransitionConditionUnion
}

// r2LifecycleRuleDeleteObjectsTransitionConditionJSON contains the JSON metadata
// for the struct [R2LifecycleRuleDeleteObjectsTransitionCondition]
type r2LifecycleRuleDeleteObjectsTransitionConditionJSON struct {
	Type        apijson.Field
	Date        apijson.Field
	MaxAge      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r r2LifecycleRuleDeleteObjectsTransitionConditionJSON) RawJSON() string {
	return r.raw
}

func (r *R2LifecycleRuleDeleteObjectsTransitionCondition) UnmarshalJSON(data []byte) (err error) {
	*r = R2LifecycleRuleDeleteObjectsTransitionCondition{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [R2LifecycleRuleDeleteObjectsTransitionConditionUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [R2LifecycleAgeCondition],
// [R2LifecycleDateCondition].
func (r R2LifecycleRuleDeleteObjectsTransitionCondition) AsUnion() R2LifecycleRuleDeleteObjectsTransitionConditionUnion {
	return r.union
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
//
// Union satisfied by [R2LifecycleAgeCondition] or [R2LifecycleDateCondition].
type R2LifecycleRuleDeleteObjectsTransitionConditionUnion interface {
	implementsR2LifecycleRuleDeleteObjectsTransitionCondition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*R2LifecycleRuleDeleteObjectsTransitionConditionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(R2LifecycleAgeCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(R2LifecycleDateCondition{}),
		},
	)
}

type R2LifecycleRuleDeleteObjectsTransitionConditionType string

const (
	R2LifecycleRuleDeleteObjectsTransitionConditionTypeAge  R2LifecycleRuleDeleteObjectsTransitionConditionType = "Age"
	R2LifecycleRuleDeleteObjectsTransitionConditionTypeDate R2LifecycleRuleDeleteObjectsTransitionConditionType = "Date"
)

func (r R2LifecycleRuleDeleteObjectsTransitionConditionType) IsKnown() bool {
	switch r {
	case R2LifecycleRuleDeleteObjectsTransitionConditionTypeAge, R2LifecycleRuleDeleteObjectsTransitionConditionTypeDate:
		return true
	}
	return false
}

type R2LifecycleRuleStorageClassTransition struct {
	// Condition for lifecycle transitions to apply after an object reaches an age in
	// seconds.
	Condition    R2LifecycleRuleStorageClassTransitionsCondition    `json:"condition,required"`
	StorageClass R2LifecycleRuleStorageClassTransitionsStorageClass `json:"storageClass,required"`
	JSON         r2LifecycleRuleStorageClassTransitionJSON          `json:"-"`
}

// r2LifecycleRuleStorageClassTransitionJSON contains the JSON metadata for the
// struct [R2LifecycleRuleStorageClassTransition]
type r2LifecycleRuleStorageClassTransitionJSON struct {
	Condition    apijson.Field
	StorageClass apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *R2LifecycleRuleStorageClassTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2LifecycleRuleStorageClassTransitionJSON) RawJSON() string {
	return r.raw
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type R2LifecycleRuleStorageClassTransitionsCondition struct {
	Type   R2LifecycleRuleStorageClassTransitionsConditionType `json:"type,required"`
	Date   time.Time                                           `json:"date" format:"date"`
	MaxAge int64                                               `json:"maxAge"`
	JSON   r2LifecycleRuleStorageClassTransitionsConditionJSON `json:"-"`
	union  R2LifecycleRuleStorageClassTransitionsConditionUnion
}

// r2LifecycleRuleStorageClassTransitionsConditionJSON contains the JSON metadata
// for the struct [R2LifecycleRuleStorageClassTransitionsCondition]
type r2LifecycleRuleStorageClassTransitionsConditionJSON struct {
	Type        apijson.Field
	Date        apijson.Field
	MaxAge      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r r2LifecycleRuleStorageClassTransitionsConditionJSON) RawJSON() string {
	return r.raw
}

func (r *R2LifecycleRuleStorageClassTransitionsCondition) UnmarshalJSON(data []byte) (err error) {
	*r = R2LifecycleRuleStorageClassTransitionsCondition{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [R2LifecycleRuleStorageClassTransitionsConditionUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [R2LifecycleAgeCondition],
// [R2LifecycleDateCondition].
func (r R2LifecycleRuleStorageClassTransitionsCondition) AsUnion() R2LifecycleRuleStorageClassTransitionsConditionUnion {
	return r.union
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
//
// Union satisfied by [R2LifecycleAgeCondition] or [R2LifecycleDateCondition].
type R2LifecycleRuleStorageClassTransitionsConditionUnion interface {
	implementsR2LifecycleRuleStorageClassTransitionsCondition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*R2LifecycleRuleStorageClassTransitionsConditionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(R2LifecycleAgeCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(R2LifecycleDateCondition{}),
		},
	)
}

type R2LifecycleRuleStorageClassTransitionsConditionType string

const (
	R2LifecycleRuleStorageClassTransitionsConditionTypeAge  R2LifecycleRuleStorageClassTransitionsConditionType = "Age"
	R2LifecycleRuleStorageClassTransitionsConditionTypeDate R2LifecycleRuleStorageClassTransitionsConditionType = "Date"
)

func (r R2LifecycleRuleStorageClassTransitionsConditionType) IsKnown() bool {
	switch r {
	case R2LifecycleRuleStorageClassTransitionsConditionTypeAge, R2LifecycleRuleStorageClassTransitionsConditionTypeDate:
		return true
	}
	return false
}

type R2LifecycleRuleStorageClassTransitionsStorageClass string

const (
	R2LifecycleRuleStorageClassTransitionsStorageClassInfrequentAccess R2LifecycleRuleStorageClassTransitionsStorageClass = "InfrequentAccess"
)

func (r R2LifecycleRuleStorageClassTransitionsStorageClass) IsKnown() bool {
	switch r {
	case R2LifecycleRuleStorageClassTransitionsStorageClassInfrequentAccess:
		return true
	}
	return false
}

type R2LifecycleRuleParam struct {
	// Unique identifier for this rule.
	ID param.Field[string] `json:"id,required"`
	// Conditions that apply to all transitions of this rule.
	Conditions param.Field[R2LifecycleRuleConditionsParam] `json:"conditions,required"`
	// Whether or not this rule is in effect.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Transition to abort ongoing multipart uploads.
	AbortMultipartUploadsTransition param.Field[R2LifecycleRuleAbortMultipartUploadsTransitionParam] `json:"abortMultipartUploadsTransition"`
	// Transition to delete objects.
	DeleteObjectsTransition param.Field[R2LifecycleRuleDeleteObjectsTransitionParam] `json:"deleteObjectsTransition"`
	// Transitions to change the storage class of objects.
	StorageClassTransitions param.Field[[]R2LifecycleRuleStorageClassTransitionParam] `json:"storageClassTransitions"`
}

func (r R2LifecycleRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Conditions that apply to all transitions of this rule.
type R2LifecycleRuleConditionsParam struct {
	// Transitions will only apply to objects/uploads in the bucket that start with the
	// given prefix, an empty prefix can be provided to scope rule to all
	// objects/uploads.
	Prefix param.Field[string] `json:"prefix,required"`
}

func (r R2LifecycleRuleConditionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Transition to abort ongoing multipart uploads.
type R2LifecycleRuleAbortMultipartUploadsTransitionParam struct {
	// Condition for lifecycle transitions to apply after an object reaches an age in
	// seconds.
	Condition param.Field[R2LifecycleAgeConditionParam] `json:"condition"`
}

func (r R2LifecycleRuleAbortMultipartUploadsTransitionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Transition to delete objects.
type R2LifecycleRuleDeleteObjectsTransitionParam struct {
	// Condition for lifecycle transitions to apply after an object reaches an age in
	// seconds.
	Condition param.Field[R2LifecycleRuleDeleteObjectsTransitionConditionUnionParam] `json:"condition"`
}

func (r R2LifecycleRuleDeleteObjectsTransitionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type R2LifecycleRuleDeleteObjectsTransitionConditionParam struct {
	Type   param.Field[R2LifecycleRuleDeleteObjectsTransitionConditionType] `json:"type,required"`
	Date   param.Field[time.Time]                                           `json:"date" format:"date"`
	MaxAge param.Field[int64]                                               `json:"maxAge"`
}

func (r R2LifecycleRuleDeleteObjectsTransitionConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r R2LifecycleRuleDeleteObjectsTransitionConditionParam) implementsR2LifecycleRuleDeleteObjectsTransitionConditionUnionParam() {
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
//
// Satisfied by [R2LifecycleAgeConditionParam], [R2LifecycleDateConditionParam],
// [R2LifecycleRuleDeleteObjectsTransitionConditionParam].
type R2LifecycleRuleDeleteObjectsTransitionConditionUnionParam interface {
	implementsR2LifecycleRuleDeleteObjectsTransitionConditionUnionParam()
}

type R2LifecycleRuleStorageClassTransitionParam struct {
	// Condition for lifecycle transitions to apply after an object reaches an age in
	// seconds.
	Condition    param.Field[R2LifecycleRuleStorageClassTransitionsConditionUnionParam] `json:"condition,required"`
	StorageClass param.Field[R2LifecycleRuleStorageClassTransitionsStorageClass]        `json:"storageClass,required"`
}

func (r R2LifecycleRuleStorageClassTransitionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type R2LifecycleRuleStorageClassTransitionsConditionParam struct {
	Type   param.Field[R2LifecycleRuleStorageClassTransitionsConditionType] `json:"type,required"`
	Date   param.Field[time.Time]                                           `json:"date" format:"date"`
	MaxAge param.Field[int64]                                               `json:"maxAge"`
}

func (r R2LifecycleRuleStorageClassTransitionsConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r R2LifecycleRuleStorageClassTransitionsConditionParam) implementsR2LifecycleRuleStorageClassTransitionsConditionUnionParam() {
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
//
// Satisfied by [R2LifecycleAgeConditionParam], [R2LifecycleDateConditionParam],
// [R2LifecycleRuleStorageClassTransitionsConditionParam].
type R2LifecycleRuleStorageClassTransitionsConditionUnionParam interface {
	implementsR2LifecycleRuleStorageClassTransitionsConditionUnionParam()
}

type AccountR2BucketLifecycleGetResponse struct {
	Errors   []AccountR2BucketLifecycleGetResponseError `json:"errors,required"`
	Messages []string                                   `json:"messages,required"`
	Result   AccountR2BucketLifecycleGetResponseResult  `json:"result,required"`
	// Whether the API call was successful.
	Success AccountR2BucketLifecycleGetResponseSuccess `json:"success,required"`
	JSON    accountR2BucketLifecycleGetResponseJSON    `json:"-"`
}

// accountR2BucketLifecycleGetResponseJSON contains the JSON metadata for the
// struct [AccountR2BucketLifecycleGetResponse]
type accountR2BucketLifecycleGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketLifecycleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketLifecycleGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketLifecycleGetResponseError struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           AccountR2BucketLifecycleGetResponseErrorsSource `json:"source"`
	JSON             accountR2BucketLifecycleGetResponseErrorJSON    `json:"-"`
}

// accountR2BucketLifecycleGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountR2BucketLifecycleGetResponseError]
type accountR2BucketLifecycleGetResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountR2BucketLifecycleGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketLifecycleGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketLifecycleGetResponseErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    accountR2BucketLifecycleGetResponseErrorsSourceJSON `json:"-"`
}

// accountR2BucketLifecycleGetResponseErrorsSourceJSON contains the JSON metadata
// for the struct [AccountR2BucketLifecycleGetResponseErrorsSource]
type accountR2BucketLifecycleGetResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketLifecycleGetResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketLifecycleGetResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountR2BucketLifecycleGetResponseResult struct {
	Rules []R2LifecycleRule                             `json:"rules"`
	JSON  accountR2BucketLifecycleGetResponseResultJSON `json:"-"`
}

// accountR2BucketLifecycleGetResponseResultJSON contains the JSON metadata for the
// struct [AccountR2BucketLifecycleGetResponseResult]
type accountR2BucketLifecycleGetResponseResultJSON struct {
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketLifecycleGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountR2BucketLifecycleGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountR2BucketLifecycleGetResponseSuccess bool

const (
	AccountR2BucketLifecycleGetResponseSuccessTrue AccountR2BucketLifecycleGetResponseSuccess = true
)

func (r AccountR2BucketLifecycleGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountR2BucketLifecycleGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountR2BucketLifecycleGetParams struct {
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[AccountR2BucketLifecycleGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type AccountR2BucketLifecycleGetParamsCfR2Jurisdiction string

const (
	AccountR2BucketLifecycleGetParamsCfR2JurisdictionDefault AccountR2BucketLifecycleGetParamsCfR2Jurisdiction = "default"
	AccountR2BucketLifecycleGetParamsCfR2JurisdictionEu      AccountR2BucketLifecycleGetParamsCfR2Jurisdiction = "eu"
	AccountR2BucketLifecycleGetParamsCfR2JurisdictionFedramp AccountR2BucketLifecycleGetParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketLifecycleGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketLifecycleGetParamsCfR2JurisdictionDefault, AccountR2BucketLifecycleGetParamsCfR2JurisdictionEu, AccountR2BucketLifecycleGetParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type AccountR2BucketLifecycleUpdateParams struct {
	Rules param.Field[[]R2LifecycleRuleParam] `json:"rules"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[AccountR2BucketLifecycleUpdateParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r AccountR2BucketLifecycleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type AccountR2BucketLifecycleUpdateParamsCfR2Jurisdiction string

const (
	AccountR2BucketLifecycleUpdateParamsCfR2JurisdictionDefault AccountR2BucketLifecycleUpdateParamsCfR2Jurisdiction = "default"
	AccountR2BucketLifecycleUpdateParamsCfR2JurisdictionEu      AccountR2BucketLifecycleUpdateParamsCfR2Jurisdiction = "eu"
	AccountR2BucketLifecycleUpdateParamsCfR2JurisdictionFedramp AccountR2BucketLifecycleUpdateParamsCfR2Jurisdiction = "fedramp"
)

func (r AccountR2BucketLifecycleUpdateParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case AccountR2BucketLifecycleUpdateParamsCfR2JurisdictionDefault, AccountR2BucketLifecycleUpdateParamsCfR2JurisdictionEu, AccountR2BucketLifecycleUpdateParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}
