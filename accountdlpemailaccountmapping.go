// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/tidwall/gjson"
)

// AccountDlpEmailAccountMappingService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDlpEmailAccountMappingService] method instead.
type AccountDlpEmailAccountMappingService struct {
	Options []option.RequestOption
}

// NewAccountDlpEmailAccountMappingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDlpEmailAccountMappingService(opts ...option.RequestOption) (r *AccountDlpEmailAccountMappingService) {
	r = &AccountDlpEmailAccountMappingService{}
	r.Options = opts
	return
}

// Create mapping
func (r *AccountDlpEmailAccountMappingService) New(ctx context.Context, accountID string, body AccountDlpEmailAccountMappingNewParams, opts ...option.RequestOption) (res *AccountDlpEmailAccountMappingNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/email/account_mapping", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get mapping
func (r *AccountDlpEmailAccountMappingService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountDlpEmailAccountMappingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/email/account_mapping", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountMapping struct {
	AddinIdentifierToken string             `json:"addin_identifier_token,required" format:"uuid"`
	AuthRequirements     Auth               `json:"auth_requirements,required"`
	JSON                 accountMappingJSON `json:"-"`
}

// accountMappingJSON contains the JSON metadata for the struct [AccountMapping]
type accountMappingJSON struct {
	AddinIdentifierToken apijson.Field
	AuthRequirements     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountMapping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMappingJSON) RawJSON() string {
	return r.raw
}

type Auth struct {
	Type AuthType `json:"type,required"`
	// This field can have the runtime type of [[]string].
	AllowedMicrosoftOrganizations interface{} `json:"allowed_microsoft_organizations"`
	JSON                          authJSON    `json:"-"`
	union                         AuthUnion
}

// authJSON contains the JSON metadata for the struct [Auth]
type authJSON struct {
	Type                          apijson.Field
	AllowedMicrosoftOrganizations apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r authJSON) RawJSON() string {
	return r.raw
}

func (r *Auth) UnmarshalJSON(data []byte) (err error) {
	*r = Auth{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthUnion] interface which you can cast to the specific types
// for more type safety.
//
// Possible runtime types of the union are [AuthObject], [AuthType].
func (r Auth) AsUnion() AuthUnion {
	return r.union
}

// Union satisfied by [AuthObject] or [AuthType].
type AuthUnion interface {
	implementsAuth()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthType{}),
		},
	)
}

type AuthObject struct {
	AllowedMicrosoftOrganizations []string       `json:"allowed_microsoft_organizations,required"`
	Type                          AuthObjectType `json:"type,required"`
	JSON                          authObjectJSON `json:"-"`
}

// authObjectJSON contains the JSON metadata for the struct [AuthObject]
type authObjectJSON struct {
	AllowedMicrosoftOrganizations apijson.Field
	Type                          apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *AuthObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authObjectJSON) RawJSON() string {
	return r.raw
}

func (r AuthObject) implementsAuth() {}

type AuthObjectType string

const (
	AuthObjectTypeOrg AuthObjectType = "Org"
)

func (r AuthObjectType) IsKnown() bool {
	switch r {
	case AuthObjectTypeOrg:
		return true
	}
	return false
}

type AuthType struct {
	Type AuthTypeType `json:"type,required"`
	JSON authTypeJSON `json:"-"`
}

// authTypeJSON contains the JSON metadata for the struct [AuthType]
type authTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authTypeJSON) RawJSON() string {
	return r.raw
}

func (r AuthType) implementsAuth() {}

type AuthTypeType string

const (
	AuthTypeTypeNoAuth AuthTypeType = "NoAuth"
)

func (r AuthTypeType) IsKnown() bool {
	switch r {
	case AuthTypeTypeNoAuth:
		return true
	}
	return false
}

type AuthParam struct {
	Type                          param.Field[AuthType]    `json:"type,required"`
	AllowedMicrosoftOrganizations param.Field[interface{}] `json:"allowed_microsoft_organizations"`
}

func (r AuthParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthParam) implementsAuthUnionParam() {}

// Satisfied by [AuthObjectParam], [AuthTypeParam], [AuthParam].
type AuthUnionParam interface {
	implementsAuthUnionParam()
}

type AuthObjectParam struct {
	AllowedMicrosoftOrganizations param.Field[[]string]       `json:"allowed_microsoft_organizations,required"`
	Type                          param.Field[AuthObjectType] `json:"type,required"`
}

func (r AuthObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthObjectParam) implementsAuthUnionParam() {}

type AuthTypeParam struct {
	Type param.Field[AuthTypeType] `json:"type,required"`
}

func (r AuthTypeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthTypeParam) implementsAuthUnionParam() {}

type AccountDlpEmailAccountMappingNewResponse struct {
	Result AccountMapping                               `json:"result"`
	JSON   accountDlpEmailAccountMappingNewResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountDlpEmailAccountMappingNewResponseJSON contains the JSON metadata for the
// struct [AccountDlpEmailAccountMappingNewResponse]
type accountDlpEmailAccountMappingNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpEmailAccountMappingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpEmailAccountMappingNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpEmailAccountMappingGetResponse struct {
	Result AccountMapping                               `json:"result"`
	JSON   accountDlpEmailAccountMappingGetResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountDlpEmailAccountMappingGetResponseJSON contains the JSON metadata for the
// struct [AccountDlpEmailAccountMappingGetResponse]
type accountDlpEmailAccountMappingGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpEmailAccountMappingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpEmailAccountMappingGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpEmailAccountMappingNewParams struct {
	AuthRequirements param.Field[AuthUnionParam] `json:"auth_requirements,required"`
}

func (r AccountDlpEmailAccountMappingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
