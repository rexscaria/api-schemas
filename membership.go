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
func (r *MembershipService) Remove(ctx context.Context, membershipID string, body MembershipRemoveParams, opts ...option.RequestOption) (res *MembershipRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	if membershipID == "" {
		err = errors.New("missing required membership_id parameter")
		return
	}
	path := fmt.Sprintf("memberships/%s", membershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type APIResponseSingleIam struct {
	Errors   []IamMessage `json:"errors,required"`
	Messages []IamMessage `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseSingleIamSuccess `json:"success,required"`
	JSON    apiResponseSingleIamJSON    `json:"-"`
}

// apiResponseSingleIamJSON contains the JSON metadata for the struct
// [APIResponseSingleIam]
type apiResponseSingleIamJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleIam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleIamJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseSingleIamSuccess bool

const (
	APIResponseSingleIamSuccessTrue APIResponseSingleIamSuccess = true
)

func (r APIResponseSingleIamSuccess) IsKnown() bool {
	switch r {
	case APIResponseSingleIamSuccessTrue:
		return true
	}
	return false
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
	// Specifies the default nameservers to be used for new zones added to this
	// account.
	//
	// - `cloudflare.standard` for Cloudflare-branded nameservers
	// - `custom.account` for account custom nameservers
	// - `custom.tenant` for tenant custom nameservers
	//
	// See
	// [Custom Nameservers](https://developers.cloudflare.com/dns/additional-options/custom-nameservers/)
	// for more information.
	//
	// Deprecated in favor of
	// [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-update-dns-settings).
	//
	// Deprecated: deprecated
	DefaultNameservers SchemasAccountSettingsDefaultNameservers `json:"default_nameservers"`
	// Indicates whether membership in this account requires that Two-Factor
	// Authentication is enabled
	EnforceTwofactor bool `json:"enforce_twofactor"`
	// Indicates whether new zones should use the account-level custom nameservers by
	// default.
	//
	// Deprecated in favor of
	// [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-update-dns-settings).
	//
	// Deprecated: deprecated
	UseAccountCustomNsByDefault bool                       `json:"use_account_custom_ns_by_default"`
	JSON                        schemasAccountSettingsJSON `json:"-"`
}

// schemasAccountSettingsJSON contains the JSON metadata for the struct
// [SchemasAccountSettings]
type schemasAccountSettingsJSON struct {
	AbuseContactEmail           apijson.Field
	DefaultNameservers          apijson.Field
	EnforceTwofactor            apijson.Field
	UseAccountCustomNsByDefault apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *SchemasAccountSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemasAccountSettingsJSON) RawJSON() string {
	return r.raw
}

// Specifies the default nameservers to be used for new zones added to this
// account.
//
// - `cloudflare.standard` for Cloudflare-branded nameservers
// - `custom.account` for account custom nameservers
// - `custom.tenant` for tenant custom nameservers
//
// See
// [Custom Nameservers](https://developers.cloudflare.com/dns/additional-options/custom-nameservers/)
// for more information.
//
// Deprecated in favor of
// [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-update-dns-settings).
type SchemasAccountSettingsDefaultNameservers string

const (
	SchemasAccountSettingsDefaultNameserversCloudflareStandard SchemasAccountSettingsDefaultNameservers = "cloudflare.standard"
	SchemasAccountSettingsDefaultNameserversCustomAccount      SchemasAccountSettingsDefaultNameservers = "custom.account"
	SchemasAccountSettingsDefaultNameserversCustomTenant       SchemasAccountSettingsDefaultNameservers = "custom.tenant"
)

func (r SchemasAccountSettingsDefaultNameservers) IsKnown() bool {
	switch r {
	case SchemasAccountSettingsDefaultNameserversCloudflareStandard, SchemasAccountSettingsDefaultNameserversCustomAccount, SchemasAccountSettingsDefaultNameserversCustomTenant:
		return true
	}
	return false
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
	Result MembershipWithPolicies                   `json:"result"`
	JSON   singleMembershipResponseWithPoliciesJSON `json:"-"`
	APIResponseSingleIam
}

// singleMembershipResponseWithPoliciesJSON contains the JSON metadata for the
// struct [SingleMembershipResponseWithPolicies]
type singleMembershipResponseWithPoliciesJSON struct {
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

type MembershipListResponse struct {
	// This field can have the runtime type of [[]IamMessage].
	Errors interface{} `json:"errors"`
	// This field can have the runtime type of [[]IamMessage].
	Messages interface{} `json:"messages"`
	// This field can have the runtime type of
	// [[]MembershipListResponseIamCollectionMembershipResponseResult],
	// [[]MembershipWithPolicies].
	Result interface{} `json:"result"`
	// This field can have the runtime type of [IamAPIResponseCollectionResultInfo].
	ResultInfo interface{} `json:"result_info"`
	// Whether the API call was successful
	Success MembershipListResponseSuccess `json:"success"`
	JSON    membershipListResponseJSON    `json:"-"`
	union   MembershipListResponseUnion
}

// membershipListResponseJSON contains the JSON metadata for the struct
// [MembershipListResponse]
type membershipListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
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
	Result []MembershipListResponseIamCollectionMembershipResponseResult `json:"result"`
	JSON   membershipListResponseIamCollectionMembershipResponseJSON     `json:"-"`
	IamAPIResponseCollection
}

// membershipListResponseIamCollectionMembershipResponseJSON contains the JSON
// metadata for the struct [MembershipListResponseIamCollectionMembershipResponse]
type membershipListResponseIamCollectionMembershipResponseJSON struct {
	Result      apijson.Field
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

type MembershipListResponseIamCollectionMembershipResponseWithPolicies struct {
	Result []MembershipWithPolicies                                              `json:"result"`
	JSON   membershipListResponseIamCollectionMembershipResponseWithPoliciesJSON `json:"-"`
	IamAPIResponseCollection
}

// membershipListResponseIamCollectionMembershipResponseWithPoliciesJSON contains
// the JSON metadata for the struct
// [MembershipListResponseIamCollectionMembershipResponseWithPolicies]
type membershipListResponseIamCollectionMembershipResponseWithPoliciesJSON struct {
	Result      apijson.Field
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

// Whether the API call was successful
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
	Result MembershipRemoveResponseResult `json:"result"`
	JSON   membershipRemoveResponseJSON   `json:"-"`
	APIResponseSingleIam
}

// membershipRemoveResponseJSON contains the JSON metadata for the struct
// [MembershipRemoveResponse]
type membershipRemoveResponseJSON struct {
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

type MembershipRemoveParams struct {
	Body interface{} `json:"body,required"`
}

func (r MembershipRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
