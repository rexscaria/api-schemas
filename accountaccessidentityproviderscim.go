// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountAccessIdentityProviderScimService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessIdentityProviderScimService] method instead.
type AccountAccessIdentityProviderScimService struct {
	Options []option.RequestOption
}

// NewAccountAccessIdentityProviderScimService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessIdentityProviderScimService(opts ...option.RequestOption) (r *AccountAccessIdentityProviderScimService) {
	r = &AccountAccessIdentityProviderScimService{}
	r.Options = opts
	return
}

// Lists SCIM Group resources synced to Cloudflare via the System for Cross-domain
// Identity Management (SCIM).
func (r *AccountAccessIdentityProviderScimService) ListGroups(ctx context.Context, accountID string, identityProviderID string, query AccountAccessIdentityProviderScimListGroupsParams, opts ...option.RequestOption) (res *AccountAccessIdentityProviderScimListGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identityProviderID == "" {
		err = errors.New("missing required identity_provider_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/identity_providers/%s/scim/groups", accountID, identityProviderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Lists SCIM User resources synced to Cloudflare via the System for Cross-domain
// Identity Management (SCIM).
func (r *AccountAccessIdentityProviderScimService) ListUsers(ctx context.Context, accountID string, identityProviderID string, query AccountAccessIdentityProviderScimListUsersParams, opts ...option.RequestOption) (res *AccountAccessIdentityProviderScimListUsersResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identityProviderID == "" {
		err = errors.New("missing required identity_provider_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/identity_providers/%s/scim/users", accountID, identityProviderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The metadata of the SCIM resource.
type Meta struct {
	// The timestamp of when the SCIM resource was created.
	Created time.Time `json:"created" format:"date-time"`
	// The timestamp of when the SCIM resource was last modified.
	LastModified time.Time `json:"lastModified" format:"date-time"`
	JSON         metaJSON  `json:"-"`
}

// metaJSON contains the JSON metadata for the struct [Meta]
type metaJSON struct {
	Created      apijson.Field
	LastModified apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Meta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metaJSON) RawJSON() string {
	return r.raw
}

type AccountAccessIdentityProviderScimListGroupsResponse struct {
	Result []AccountAccessIdentityProviderScimListGroupsResponseResult `json:"result"`
	JSON   accountAccessIdentityProviderScimListGroupsResponseJSON     `json:"-"`
	APIResponseCollectionAccess
}

// accountAccessIdentityProviderScimListGroupsResponseJSON contains the JSON
// metadata for the struct [AccountAccessIdentityProviderScimListGroupsResponse]
type accountAccessIdentityProviderScimListGroupsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderScimListGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessIdentityProviderScimListGroupsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessIdentityProviderScimListGroupsResponseResult struct {
	// The unique Cloudflare-generated Id of the SCIM resource.
	ID string `json:"id"`
	// The display name of the SCIM Group resource.
	DisplayName string `json:"displayName"`
	// The IdP-generated Id of the SCIM resource.
	ExternalID string `json:"externalId"`
	// The metadata of the SCIM resource.
	Meta Meta `json:"meta"`
	// The list of URIs which indicate the attributes contained within a SCIM resource.
	Schemas []string                                                      `json:"schemas"`
	JSON    accountAccessIdentityProviderScimListGroupsResponseResultJSON `json:"-"`
}

// accountAccessIdentityProviderScimListGroupsResponseResultJSON contains the JSON
// metadata for the struct
// [AccountAccessIdentityProviderScimListGroupsResponseResult]
type accountAccessIdentityProviderScimListGroupsResponseResultJSON struct {
	ID          apijson.Field
	DisplayName apijson.Field
	ExternalID  apijson.Field
	Meta        apijson.Field
	Schemas     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderScimListGroupsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessIdentityProviderScimListGroupsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAccessIdentityProviderScimListUsersResponse struct {
	Result []AccountAccessIdentityProviderScimListUsersResponseResult `json:"result"`
	JSON   accountAccessIdentityProviderScimListUsersResponseJSON     `json:"-"`
	APIResponseCollectionAccess
}

// accountAccessIdentityProviderScimListUsersResponseJSON contains the JSON
// metadata for the struct [AccountAccessIdentityProviderScimListUsersResponse]
type accountAccessIdentityProviderScimListUsersResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderScimListUsersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessIdentityProviderScimListUsersResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessIdentityProviderScimListUsersResponseResult struct {
	// The unique Cloudflare-generated Id of the SCIM resource.
	ID string `json:"id"`
	// Determines the status of the SCIM User resource.
	Active bool `json:"active"`
	// The name of the SCIM User resource.
	DisplayName string                                                          `json:"displayName"`
	Emails      []AccountAccessIdentityProviderScimListUsersResponseResultEmail `json:"emails"`
	// The IdP-generated Id of the SCIM resource.
	ExternalID string `json:"externalId"`
	// The metadata of the SCIM resource.
	Meta Meta `json:"meta"`
	// The list of URIs which indicate the attributes contained within a SCIM resource.
	Schemas []string                                                     `json:"schemas"`
	JSON    accountAccessIdentityProviderScimListUsersResponseResultJSON `json:"-"`
}

// accountAccessIdentityProviderScimListUsersResponseResultJSON contains the JSON
// metadata for the struct
// [AccountAccessIdentityProviderScimListUsersResponseResult]
type accountAccessIdentityProviderScimListUsersResponseResultJSON struct {
	ID          apijson.Field
	Active      apijson.Field
	DisplayName apijson.Field
	Emails      apijson.Field
	ExternalID  apijson.Field
	Meta        apijson.Field
	Schemas     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderScimListUsersResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessIdentityProviderScimListUsersResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAccessIdentityProviderScimListUsersResponseResultEmail struct {
	// Indicates if the email address is the primary email belonging to the SCIM User
	// resource.
	Primary bool `json:"primary"`
	// Indicates the type of the email address.
	Type string `json:"type"`
	// The email address of the SCIM User resource.
	Value string                                                            `json:"value" format:"email"`
	JSON  accountAccessIdentityProviderScimListUsersResponseResultEmailJSON `json:"-"`
}

// accountAccessIdentityProviderScimListUsersResponseResultEmailJSON contains the
// JSON metadata for the struct
// [AccountAccessIdentityProviderScimListUsersResponseResultEmail]
type accountAccessIdentityProviderScimListUsersResponseResultEmailJSON struct {
	Primary     apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderScimListUsersResponseResultEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessIdentityProviderScimListUsersResponseResultEmailJSON) RawJSON() string {
	return r.raw
}

type AccountAccessIdentityProviderScimListGroupsParams struct {
	// The unique Cloudflare-generated Id of the SCIM Group resource; also known as the
	// "Id".
	CfResourceID param.Field[string] `query:"cf_resource_id"`
	// The IdP-generated Id of the SCIM Group resource; also known as the "external
	// Id".
	IdpResourceID param.Field[string] `query:"idp_resource_id"`
	// The display name of the SCIM Group resource.
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [AccountAccessIdentityProviderScimListGroupsParams]'s query
// parameters as `url.Values`.
func (r AccountAccessIdentityProviderScimListGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAccessIdentityProviderScimListUsersParams struct {
	// The unique Cloudflare-generated Id of the SCIM User resource; also known as the
	// "Id".
	CfResourceID param.Field[string] `query:"cf_resource_id"`
	// The email address of the SCIM User resource.
	Email param.Field[string] `query:"email"`
	// The IdP-generated Id of the SCIM User resource; also known as the "external Id".
	IdpResourceID param.Field[string] `query:"idp_resource_id"`
	// The name of the SCIM User resource.
	Name param.Field[string] `query:"name"`
	// The username of the SCIM User resource.
	Username param.Field[string] `query:"username"`
}

// URLQuery serializes [AccountAccessIdentityProviderScimListUsersParams]'s query
// parameters as `url.Values`.
func (r AccountAccessIdentityProviderScimListUsersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
