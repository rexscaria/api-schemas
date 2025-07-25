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
	"github.com/tidwall/gjson"
)

// MembershipService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMembershipService] method instead.
type MembershipService struct {
	Options []option.RequestOption
}

// NewMembershipService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMembershipService(opts ...option.RequestOption) (r *MembershipService) {
	r = &MembershipService{}
	r.Options = opts
	return
}

// Get a specific membership.
func (r *MembershipService) Get(ctx context.Context, membershipID string, opts ...option.RequestOption) (res *SingleMembershipResponseWithPolicies, err error) {
	opts = append(r.Options[:], opts...)
	if membershipID == "" {
		err = errors.New("missing required membership_id parameter")
		return
	}
	path := fmt.Sprintf("memberships/%s", membershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Accept or reject this account invitation.
func (r *MembershipService) Update(ctx context.Context, membershipID string, body MembershipUpdateParams, opts ...option.RequestOption) (res *SingleMembershipResponseWithPolicies, err error) {
	opts = append(r.Options[:], opts...)
	if membershipID == "" {
		err = errors.New("missing required membership_id parameter")
		return
	}
	path := fmt.Sprintf("memberships/%s", membershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List memberships of accounts the user can access.
func (r *MembershipService) List(ctx context.Context, query MembershipListParams, opts ...option.RequestOption) (res *MembershipListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "memberships"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Remove the associated member from an account.
func (r *MembershipService) Remove(ctx context.Context, membershipID string, opts ...option.RequestOption) (res *MembershipRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	if membershipID == "" {
		err = errors.New("missing required membership_id parameter")
		return
	}
	path := fmt.Sprintf("memberships/%s", membershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type MembershipWithPolicies struct {
	// Membership identifier tag.
	ID      string         `json:"id"`
	Account SchemasAccount `json:"account"`
	// Enterprise only. Indicates whether or not API access is enabled specifically for
	// this user on a given account.
	APIAccessEnabled bool `json:"api_access_enabled,nullable"`
	// All access permissions for the user at the account.
	Permissions IamPermissions `json:"permissions"`
	// Access policy for the membership
	Policies []IamListMemberPolicy `json:"policies"`
	// List of role names the membership has for this account.
	Roles []string `json:"roles"`
	// Status of this membership.
	Status SchemasStatus              `json:"status"`
	JSON   membershipWithPoliciesJSON `json:"-"`
}

// membershipWithPoliciesJSON contains the JSON metadata for the struct
// [MembershipWithPolicies]
type membershipWithPoliciesJSON struct {
	ID               apijson.Field
	Account          apijson.Field
	APIAccessEnabled apijson.Field
	Permissions      apijson.Field
	Policies         apijson.Field
	Roles            apijson.Field
	Status           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipWithPolicies) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipWithPoliciesJSON) RawJSON() string {
	return r.raw
}

type SchemasAccount struct {
	// Identifier
	ID string `json:"id,required"`
	// Account name
	Name string `json:"name,required"`
	// Timestamp for the creation of the account
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Account settings
	Settings SchemasAccountSettings `json:"settings"`
	JSON     schemasAccountJSON     `json:"-"`
}

// schemasAccountJSON contains the JSON metadata for the struct [SchemasAccount]
type schemasAccountJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	CreatedOn   apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemasAccountJSON) RawJSON() string {
	return r.raw
}

// Account settings
type SchemasAccountSettings struct {
	// Sets an abuse contact email to notify for abuse reports.
	AbuseContactEmail string `json:"abuse_contact_email"`
	// Indicates whether membership in this account requires that Two-Factor
	// Authentication is enabled
	EnforceTwofactor bool                       `json:"enforce_twofactor"`
	JSON             schemasAccountSettingsJSON `json:"-"`
}

// schemasAccountSettingsJSON contains the JSON metadata for the struct
// [SchemasAccountSettings]
type schemasAccountSettingsJSON struct {
	AbuseContactEmail apijson.Field
	EnforceTwofactor  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SchemasAccountSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemasAccountSettingsJSON) RawJSON() string {
	return r.raw
}

// Status of this membership.
type SchemasStatus string

const (
	SchemasStatusAccepted SchemasStatus = "accepted"
	SchemasStatusPending  SchemasStatus = "pending"
	SchemasStatusRejected SchemasStatus = "rejected"
)

func (r SchemasStatus) IsKnown() bool {
	switch r {
	case SchemasStatusAccepted, SchemasStatusPending, SchemasStatusRejected:
		return true
	}
	return false
}

type SingleMembershipResponseWithPolicies struct {
	Errors   []SingleMembershipResponseWithPoliciesError   `json:"errors,required"`
	Messages []SingleMembershipResponseWithPoliciesMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success SingleMembershipResponseWithPoliciesSuccess `json:"success,required"`
	Result  MembershipWithPolicies                      `json:"result"`
	JSON    singleMembershipResponseWithPoliciesJSON    `json:"-"`
}

// singleMembershipResponseWithPoliciesJSON contains the JSON metadata for the
// struct [SingleMembershipResponseWithPolicies]
type singleMembershipResponseWithPoliciesJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleMembershipResponseWithPolicies) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleMembershipResponseWithPoliciesJSON) RawJSON() string {
	return r.raw
}

type SingleMembershipResponseWithPoliciesError struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           SingleMembershipResponseWithPoliciesErrorsSource `json:"source"`
	JSON             singleMembershipResponseWithPoliciesErrorJSON    `json:"-"`
}

// singleMembershipResponseWithPoliciesErrorJSON contains the JSON metadata for the
// struct [SingleMembershipResponseWithPoliciesError]
type singleMembershipResponseWithPoliciesErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SingleMembershipResponseWithPoliciesError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleMembershipResponseWithPoliciesErrorJSON) RawJSON() string {
	return r.raw
}

type SingleMembershipResponseWithPoliciesErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    singleMembershipResponseWithPoliciesErrorsSourceJSON `json:"-"`
}

// singleMembershipResponseWithPoliciesErrorsSourceJSON contains the JSON metadata
// for the struct [SingleMembershipResponseWithPoliciesErrorsSource]
type singleMembershipResponseWithPoliciesErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleMembershipResponseWithPoliciesErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleMembershipResponseWithPoliciesErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SingleMembershipResponseWithPoliciesMessage struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           SingleMembershipResponseWithPoliciesMessagesSource `json:"source"`
	JSON             singleMembershipResponseWithPoliciesMessageJSON    `json:"-"`
}

// singleMembershipResponseWithPoliciesMessageJSON contains the JSON metadata for
// the struct [SingleMembershipResponseWithPoliciesMessage]
type singleMembershipResponseWithPoliciesMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SingleMembershipResponseWithPoliciesMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleMembershipResponseWithPoliciesMessageJSON) RawJSON() string {
	return r.raw
}

type SingleMembershipResponseWithPoliciesMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    singleMembershipResponseWithPoliciesMessagesSourceJSON `json:"-"`
}

// singleMembershipResponseWithPoliciesMessagesSourceJSON contains the JSON
// metadata for the struct [SingleMembershipResponseWithPoliciesMessagesSource]
type singleMembershipResponseWithPoliciesMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleMembershipResponseWithPoliciesMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleMembershipResponseWithPoliciesMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SingleMembershipResponseWithPoliciesSuccess bool

const (
	SingleMembershipResponseWithPoliciesSuccessTrue SingleMembershipResponseWithPoliciesSuccess = true
)

func (r SingleMembershipResponseWithPoliciesSuccess) IsKnown() bool {
	switch r {
	case SingleMembershipResponseWithPoliciesSuccessTrue:
		return true
	}
	return false
}

type MembershipListResponse struct {
	// This field can have the runtime type of
	// [[]MembershipListResponseIamCollectionMembershipResponseError],
	// [[]MembershipListResponseIamCollectionMembershipResponseWithPoliciesError].
	Errors interface{} `json:"errors,required"`
	// This field can have the runtime type of
	// [[]MembershipListResponseIamCollectionMembershipResponseMessage],
	// [[]MembershipListResponseIamCollectionMembershipResponseWithPoliciesMessage].
	Messages interface{} `json:"messages,required"`
	// Whether the API call was successful.
	Success MembershipListResponseSuccess `json:"success,required"`
	// This field can have the runtime type of
	// [[]MembershipListResponseIamCollectionMembershipResponseResult],
	// [[]MembershipWithPolicies].
	Result interface{} `json:"result"`
	// This field can have the runtime type of
	// [MembershipListResponseIamCollectionMembershipResponseResultInfo],
	// [MembershipListResponseIamCollectionMembershipResponseWithPoliciesResultInfo].
	ResultInfo interface{}                `json:"result_info"`
	JSON       membershipListResponseJSON `json:"-"`
	union      MembershipListResponseUnion
}

// membershipListResponseJSON contains the JSON metadata for the struct
// [MembershipListResponse]
type membershipListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r membershipListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *MembershipListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = MembershipListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [MembershipListResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [MembershipListResponseIamCollectionMembershipResponse],
// [MembershipListResponseIamCollectionMembershipResponseWithPolicies].
func (r MembershipListResponse) AsUnion() MembershipListResponseUnion {
	return r.union
}

// Union satisfied by [MembershipListResponseIamCollectionMembershipResponse] or
// [MembershipListResponseIamCollectionMembershipResponseWithPolicies].
type MembershipListResponseUnion interface {
	implementsMembershipListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MembershipListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MembershipListResponseIamCollectionMembershipResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MembershipListResponseIamCollectionMembershipResponseWithPolicies{}),
		},
	)
}

type MembershipListResponseIamCollectionMembershipResponse struct {
	Errors   []MembershipListResponseIamCollectionMembershipResponseError   `json:"errors,required"`
	Messages []MembershipListResponseIamCollectionMembershipResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success    MembershipListResponseIamCollectionMembershipResponseSuccess    `json:"success,required"`
	Result     []MembershipListResponseIamCollectionMembershipResponseResult   `json:"result"`
	ResultInfo MembershipListResponseIamCollectionMembershipResponseResultInfo `json:"result_info"`
	JSON       membershipListResponseIamCollectionMembershipResponseJSON       `json:"-"`
}

// membershipListResponseIamCollectionMembershipResponseJSON contains the JSON
// metadata for the struct [MembershipListResponseIamCollectionMembershipResponse]
type membershipListResponseIamCollectionMembershipResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponseIamCollectionMembershipResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipListResponseIamCollectionMembershipResponseJSON) RawJSON() string {
	return r.raw
}

func (r MembershipListResponseIamCollectionMembershipResponse) implementsMembershipListResponse() {}

type MembershipListResponseIamCollectionMembershipResponseError struct {
	Code             int64                                                             `json:"code,required"`
	Message          string                                                            `json:"message,required"`
	DocumentationURL string                                                            `json:"documentation_url"`
	Source           MembershipListResponseIamCollectionMembershipResponseErrorsSource `json:"source"`
	JSON             membershipListResponseIamCollectionMembershipResponseErrorJSON    `json:"-"`
}

// membershipListResponseIamCollectionMembershipResponseErrorJSON contains the JSON
// metadata for the struct
// [MembershipListResponseIamCollectionMembershipResponseError]
type membershipListResponseIamCollectionMembershipResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipListResponseIamCollectionMembershipResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipListResponseIamCollectionMembershipResponseErrorJSON) RawJSON() string {
	return r.raw
}

type MembershipListResponseIamCollectionMembershipResponseErrorsSource struct {
	Pointer string                                                                `json:"pointer"`
	JSON    membershipListResponseIamCollectionMembershipResponseErrorsSourceJSON `json:"-"`
}

// membershipListResponseIamCollectionMembershipResponseErrorsSourceJSON contains
// the JSON metadata for the struct
// [MembershipListResponseIamCollectionMembershipResponseErrorsSource]
type membershipListResponseIamCollectionMembershipResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponseIamCollectionMembershipResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipListResponseIamCollectionMembershipResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type MembershipListResponseIamCollectionMembershipResponseMessage struct {
	Code             int64                                                               `json:"code,required"`
	Message          string                                                              `json:"message,required"`
	DocumentationURL string                                                              `json:"documentation_url"`
	Source           MembershipListResponseIamCollectionMembershipResponseMessagesSource `json:"source"`
	JSON             membershipListResponseIamCollectionMembershipResponseMessageJSON    `json:"-"`
}

// membershipListResponseIamCollectionMembershipResponseMessageJSON contains the
// JSON metadata for the struct
// [MembershipListResponseIamCollectionMembershipResponseMessage]
type membershipListResponseIamCollectionMembershipResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipListResponseIamCollectionMembershipResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipListResponseIamCollectionMembershipResponseMessageJSON) RawJSON() string {
	return r.raw
}

type MembershipListResponseIamCollectionMembershipResponseMessagesSource struct {
	Pointer string                                                                  `json:"pointer"`
	JSON    membershipListResponseIamCollectionMembershipResponseMessagesSourceJSON `json:"-"`
}

// membershipListResponseIamCollectionMembershipResponseMessagesSourceJSON contains
// the JSON metadata for the struct
// [MembershipListResponseIamCollectionMembershipResponseMessagesSource]
type membershipListResponseIamCollectionMembershipResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponseIamCollectionMembershipResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipListResponseIamCollectionMembershipResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type MembershipListResponseIamCollectionMembershipResponseSuccess bool

const (
	MembershipListResponseIamCollectionMembershipResponseSuccessTrue MembershipListResponseIamCollectionMembershipResponseSuccess = true
)

func (r MembershipListResponseIamCollectionMembershipResponseSuccess) IsKnown() bool {
	switch r {
	case MembershipListResponseIamCollectionMembershipResponseSuccessTrue:
		return true
	}
	return false
}

type MembershipListResponseIamCollectionMembershipResponseResult struct {
	// Membership identifier tag.
	ID      string         `json:"id"`
	Account SchemasAccount `json:"account"`
	// Enterprise only. Indicates whether or not API access is enabled specifically for
	// this user on a given account.
	APIAccessEnabled bool `json:"api_access_enabled,nullable"`
	// All access permissions for the user at the account.
	Permissions IamPermissions `json:"permissions"`
	// List of role names the membership has for this account.
	Roles []string `json:"roles"`
	// Status of this membership.
	Status SchemasStatus                                                   `json:"status"`
	JSON   membershipListResponseIamCollectionMembershipResponseResultJSON `json:"-"`
}

// membershipListResponseIamCollectionMembershipResponseResultJSON contains the
// JSON metadata for the struct
// [MembershipListResponseIamCollectionMembershipResponseResult]
type membershipListResponseIamCollectionMembershipResponseResultJSON struct {
	ID               apijson.Field
	Account          apijson.Field
	APIAccessEnabled apijson.Field
	Permissions      apijson.Field
	Roles            apijson.Field
	Status           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipListResponseIamCollectionMembershipResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipListResponseIamCollectionMembershipResponseResultJSON) RawJSON() string {
	return r.raw
}

type MembershipListResponseIamCollectionMembershipResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                             `json:"total_count"`
	JSON       membershipListResponseIamCollectionMembershipResponseResultInfoJSON `json:"-"`
}

// membershipListResponseIamCollectionMembershipResponseResultInfoJSON contains the
// JSON metadata for the struct
// [MembershipListResponseIamCollectionMembershipResponseResultInfo]
type membershipListResponseIamCollectionMembershipResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponseIamCollectionMembershipResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipListResponseIamCollectionMembershipResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type MembershipListResponseIamCollectionMembershipResponseWithPolicies struct {
	Errors   []MembershipListResponseIamCollectionMembershipResponseWithPoliciesError   `json:"errors,required"`
	Messages []MembershipListResponseIamCollectionMembershipResponseWithPoliciesMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success    MembershipListResponseIamCollectionMembershipResponseWithPoliciesSuccess    `json:"success,required"`
	Result     []MembershipWithPolicies                                                    `json:"result"`
	ResultInfo MembershipListResponseIamCollectionMembershipResponseWithPoliciesResultInfo `json:"result_info"`
	JSON       membershipListResponseIamCollectionMembershipResponseWithPoliciesJSON       `json:"-"`
}

// membershipListResponseIamCollectionMembershipResponseWithPoliciesJSON contains
// the JSON metadata for the struct
// [MembershipListResponseIamCollectionMembershipResponseWithPolicies]
type membershipListResponseIamCollectionMembershipResponseWithPoliciesJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponseIamCollectionMembershipResponseWithPolicies) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipListResponseIamCollectionMembershipResponseWithPoliciesJSON) RawJSON() string {
	return r.raw
}

func (r MembershipListResponseIamCollectionMembershipResponseWithPolicies) implementsMembershipListResponse() {
}

type MembershipListResponseIamCollectionMembershipResponseWithPoliciesError struct {
	Code             int64                                                                         `json:"code,required"`
	Message          string                                                                        `json:"message,required"`
	DocumentationURL string                                                                        `json:"documentation_url"`
	Source           MembershipListResponseIamCollectionMembershipResponseWithPoliciesErrorsSource `json:"source"`
	JSON             membershipListResponseIamCollectionMembershipResponseWithPoliciesErrorJSON    `json:"-"`
}

// membershipListResponseIamCollectionMembershipResponseWithPoliciesErrorJSON
// contains the JSON metadata for the struct
// [MembershipListResponseIamCollectionMembershipResponseWithPoliciesError]
type membershipListResponseIamCollectionMembershipResponseWithPoliciesErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipListResponseIamCollectionMembershipResponseWithPoliciesError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipListResponseIamCollectionMembershipResponseWithPoliciesErrorJSON) RawJSON() string {
	return r.raw
}

type MembershipListResponseIamCollectionMembershipResponseWithPoliciesErrorsSource struct {
	Pointer string                                                                            `json:"pointer"`
	JSON    membershipListResponseIamCollectionMembershipResponseWithPoliciesErrorsSourceJSON `json:"-"`
}

// membershipListResponseIamCollectionMembershipResponseWithPoliciesErrorsSourceJSON
// contains the JSON metadata for the struct
// [MembershipListResponseIamCollectionMembershipResponseWithPoliciesErrorsSource]
type membershipListResponseIamCollectionMembershipResponseWithPoliciesErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponseIamCollectionMembershipResponseWithPoliciesErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipListResponseIamCollectionMembershipResponseWithPoliciesErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type MembershipListResponseIamCollectionMembershipResponseWithPoliciesMessage struct {
	Code             int64                                                                           `json:"code,required"`
	Message          string                                                                          `json:"message,required"`
	DocumentationURL string                                                                          `json:"documentation_url"`
	Source           MembershipListResponseIamCollectionMembershipResponseWithPoliciesMessagesSource `json:"source"`
	JSON             membershipListResponseIamCollectionMembershipResponseWithPoliciesMessageJSON    `json:"-"`
}

// membershipListResponseIamCollectionMembershipResponseWithPoliciesMessageJSON
// contains the JSON metadata for the struct
// [MembershipListResponseIamCollectionMembershipResponseWithPoliciesMessage]
type membershipListResponseIamCollectionMembershipResponseWithPoliciesMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipListResponseIamCollectionMembershipResponseWithPoliciesMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipListResponseIamCollectionMembershipResponseWithPoliciesMessageJSON) RawJSON() string {
	return r.raw
}

type MembershipListResponseIamCollectionMembershipResponseWithPoliciesMessagesSource struct {
	Pointer string                                                                              `json:"pointer"`
	JSON    membershipListResponseIamCollectionMembershipResponseWithPoliciesMessagesSourceJSON `json:"-"`
}

// membershipListResponseIamCollectionMembershipResponseWithPoliciesMessagesSourceJSON
// contains the JSON metadata for the struct
// [MembershipListResponseIamCollectionMembershipResponseWithPoliciesMessagesSource]
type membershipListResponseIamCollectionMembershipResponseWithPoliciesMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponseIamCollectionMembershipResponseWithPoliciesMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipListResponseIamCollectionMembershipResponseWithPoliciesMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type MembershipListResponseIamCollectionMembershipResponseWithPoliciesSuccess bool

const (
	MembershipListResponseIamCollectionMembershipResponseWithPoliciesSuccessTrue MembershipListResponseIamCollectionMembershipResponseWithPoliciesSuccess = true
)

func (r MembershipListResponseIamCollectionMembershipResponseWithPoliciesSuccess) IsKnown() bool {
	switch r {
	case MembershipListResponseIamCollectionMembershipResponseWithPoliciesSuccessTrue:
		return true
	}
	return false
}

type MembershipListResponseIamCollectionMembershipResponseWithPoliciesResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                         `json:"total_count"`
	JSON       membershipListResponseIamCollectionMembershipResponseWithPoliciesResultInfoJSON `json:"-"`
}

// membershipListResponseIamCollectionMembershipResponseWithPoliciesResultInfoJSON
// contains the JSON metadata for the struct
// [MembershipListResponseIamCollectionMembershipResponseWithPoliciesResultInfo]
type membershipListResponseIamCollectionMembershipResponseWithPoliciesResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponseIamCollectionMembershipResponseWithPoliciesResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipListResponseIamCollectionMembershipResponseWithPoliciesResultInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type MembershipListResponseSuccess bool

const (
	MembershipListResponseSuccessTrue MembershipListResponseSuccess = true
)

func (r MembershipListResponseSuccess) IsKnown() bool {
	switch r {
	case MembershipListResponseSuccessTrue:
		return true
	}
	return false
}

type MembershipRemoveResponse struct {
	Errors   []MembershipRemoveResponseError   `json:"errors,required"`
	Messages []MembershipRemoveResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success MembershipRemoveResponseSuccess `json:"success,required"`
	Result  MembershipRemoveResponseResult  `json:"result"`
	JSON    membershipRemoveResponseJSON    `json:"-"`
}

// membershipRemoveResponseJSON contains the JSON metadata for the struct
// [MembershipRemoveResponse]
type membershipRemoveResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipRemoveResponseJSON) RawJSON() string {
	return r.raw
}

type MembershipRemoveResponseError struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           MembershipRemoveResponseErrorsSource `json:"source"`
	JSON             membershipRemoveResponseErrorJSON    `json:"-"`
}

// membershipRemoveResponseErrorJSON contains the JSON metadata for the struct
// [MembershipRemoveResponseError]
type membershipRemoveResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipRemoveResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipRemoveResponseErrorJSON) RawJSON() string {
	return r.raw
}

type MembershipRemoveResponseErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    membershipRemoveResponseErrorsSourceJSON `json:"-"`
}

// membershipRemoveResponseErrorsSourceJSON contains the JSON metadata for the
// struct [MembershipRemoveResponseErrorsSource]
type membershipRemoveResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipRemoveResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipRemoveResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type MembershipRemoveResponseMessage struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           MembershipRemoveResponseMessagesSource `json:"source"`
	JSON             membershipRemoveResponseMessageJSON    `json:"-"`
}

// membershipRemoveResponseMessageJSON contains the JSON metadata for the struct
// [MembershipRemoveResponseMessage]
type membershipRemoveResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipRemoveResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipRemoveResponseMessageJSON) RawJSON() string {
	return r.raw
}

type MembershipRemoveResponseMessagesSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    membershipRemoveResponseMessagesSourceJSON `json:"-"`
}

// membershipRemoveResponseMessagesSourceJSON contains the JSON metadata for the
// struct [MembershipRemoveResponseMessagesSource]
type membershipRemoveResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipRemoveResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipRemoveResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type MembershipRemoveResponseSuccess bool

const (
	MembershipRemoveResponseSuccessTrue MembershipRemoveResponseSuccess = true
)

func (r MembershipRemoveResponseSuccess) IsKnown() bool {
	switch r {
	case MembershipRemoveResponseSuccessTrue:
		return true
	}
	return false
}

type MembershipRemoveResponseResult struct {
	// Membership identifier tag.
	ID   string                             `json:"id"`
	JSON membershipRemoveResponseResultJSON `json:"-"`
}

// membershipRemoveResponseResultJSON contains the JSON metadata for the struct
// [MembershipRemoveResponseResult]
type membershipRemoveResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipRemoveResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipRemoveResponseResultJSON) RawJSON() string {
	return r.raw
}

type MembershipUpdateParams struct {
	// Whether to accept or reject this account invitation.
	Status param.Field[MembershipUpdateParamsStatus] `json:"status,required"`
}

func (r MembershipUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether to accept or reject this account invitation.
type MembershipUpdateParamsStatus string

const (
	MembershipUpdateParamsStatusAccepted MembershipUpdateParamsStatus = "accepted"
	MembershipUpdateParamsStatusRejected MembershipUpdateParamsStatus = "rejected"
)

func (r MembershipUpdateParamsStatus) IsKnown() bool {
	switch r {
	case MembershipUpdateParamsStatusAccepted, MembershipUpdateParamsStatusRejected:
		return true
	}
	return false
}

type MembershipListParams struct {
	Account param.Field[MembershipListParamsAccount] `query:"account"`
	// Direction to order memberships.
	Direction param.Field[MembershipListParamsDirection] `query:"direction"`
	// Account name
	Name param.Field[string] `query:"name"`
	// Field to order memberships by.
	Order param.Field[MembershipListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of memberships per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Status of this membership.
	Status param.Field[MembershipListParamsStatus] `query:"status"`
}

// URLQuery serializes [MembershipListParams]'s query parameters as `url.Values`.
func (r MembershipListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MembershipListParamsAccount struct {
	// Account name
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [MembershipListParamsAccount]'s query parameters as
// `url.Values`.
func (r MembershipListParamsAccount) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order memberships.
type MembershipListParamsDirection string

const (
	MembershipListParamsDirectionAsc  MembershipListParamsDirection = "asc"
	MembershipListParamsDirectionDesc MembershipListParamsDirection = "desc"
)

func (r MembershipListParamsDirection) IsKnown() bool {
	switch r {
	case MembershipListParamsDirectionAsc, MembershipListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order memberships by.
type MembershipListParamsOrder string

const (
	MembershipListParamsOrderID          MembershipListParamsOrder = "id"
	MembershipListParamsOrderAccountName MembershipListParamsOrder = "account.name"
	MembershipListParamsOrderStatus      MembershipListParamsOrder = "status"
)

func (r MembershipListParamsOrder) IsKnown() bool {
	switch r {
	case MembershipListParamsOrderID, MembershipListParamsOrderAccountName, MembershipListParamsOrderStatus:
		return true
	}
	return false
}

// Status of this membership.
type MembershipListParamsStatus string

const (
	MembershipListParamsStatusAccepted MembershipListParamsStatus = "accepted"
	MembershipListParamsStatusPending  MembershipListParamsStatus = "pending"
	MembershipListParamsStatusRejected MembershipListParamsStatus = "rejected"
)

func (r MembershipListParamsStatus) IsKnown() bool {
	switch r {
	case MembershipListParamsStatusAccepted, MembershipListParamsStatusPending, MembershipListParamsStatusRejected:
		return true
	}
	return false
}
