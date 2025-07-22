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

// AccountAccessIdentityProviderService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessIdentityProviderService] method instead.
type AccountAccessIdentityProviderService struct {
	Options []option.RequestOption
	Scim    *AccountAccessIdentityProviderScimService
}

// NewAccountAccessIdentityProviderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessIdentityProviderService(opts ...option.RequestOption) (r *AccountAccessIdentityProviderService) {
	r = &AccountAccessIdentityProviderService{}
	r.Options = opts
	r.Scim = NewAccountAccessIdentityProviderScimService(opts...)
	return
}

// Adds a new identity provider to Access.
func (r *AccountAccessIdentityProviderService) New(ctx context.Context, accountID string, body AccountAccessIdentityProviderNewParams, opts ...option.RequestOption) (res *SingleResponseProvider, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/identity_providers", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a configured identity provider.
func (r *AccountAccessIdentityProviderService) Get(ctx context.Context, accountID string, identityProviderID string, opts ...option.RequestOption) (res *SingleResponseProvider, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identityProviderID == "" {
		err = errors.New("missing required identity_provider_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/identity_providers/%s", accountID, identityProviderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured identity provider.
func (r *AccountAccessIdentityProviderService) Update(ctx context.Context, accountID string, identityProviderID string, body AccountAccessIdentityProviderUpdateParams, opts ...option.RequestOption) (res *SingleResponseProvider, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identityProviderID == "" {
		err = errors.New("missing required identity_provider_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/identity_providers/%s", accountID, identityProviderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all configured identity providers.
func (r *AccountAccessIdentityProviderService) List(ctx context.Context, accountID string, query AccountAccessIdentityProviderListParams, opts ...option.RequestOption) (res *AccountAccessIdentityProviderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/identity_providers", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an identity provider from Access.
func (r *AccountAccessIdentityProviderService) Delete(ctx context.Context, accountID string, identityProviderID string, opts ...option.RequestOption) (res *IDResponseApps, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identityProviderID == "" {
		err = errors.New("missing required identity_provider_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/identity_providers/%s", accountID, identityProviderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccountIdentityProviders struct {
	// UUID
	ID string `json:"id"`
	// This field can have the runtime type of [interface{}].
	Config interface{} `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// This field can have the runtime type of [IdentityProviderAccountScimConfig].
	ScimConfig interface{} `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type  AccountIdentityProvidersType `json:"type"`
	JSON  accountIdentityProvidersJSON `json:"-"`
	union AccountIdentityProvidersUnion
}

// accountIdentityProvidersJSON contains the JSON metadata for the struct
// [AccountIdentityProviders]
type accountIdentityProvidersJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r accountIdentityProvidersJSON) RawJSON() string {
	return r.raw
}

func (r *AccountIdentityProviders) UnmarshalJSON(data []byte) (err error) {
	*r = AccountIdentityProviders{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountIdentityProvidersUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [AzureAd], [CentrifyIdentityProvider],
// [FacebookAccount], [GitHubAccess], [GoogleAccount],
// [GoogleAppsIdentityProvider], [LinkedinAccount], [OidcAccess], [OktaAccess],
// [OneloginAccount], [PingoneAccount], [SAMLIdentityProvider], [YandexAccess],
// [AccountIdentityProvidersAccessOnetimepin].
func (r AccountIdentityProviders) AsUnion() AccountIdentityProvidersUnion {
	return r.union
}

// Union satisfied by [AzureAd], [CentrifyIdentityProvider], [FacebookAccount],
// [GitHubAccess], [GoogleAccount], [GoogleAppsIdentityProvider],
// [LinkedinAccount], [OidcAccess], [OktaAccess], [OneloginAccount],
// [PingoneAccount], [SAMLIdentityProvider], [YandexAccess] or
// [AccountIdentityProvidersAccessOnetimepin].
type AccountIdentityProvidersUnion interface {
	implementsAccountIdentityProviders()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountIdentityProvidersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AzureAd{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CentrifyIdentityProvider{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FacebookAccount{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GitHubAccess{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GoogleAccount{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GoogleAppsIdentityProvider{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LinkedinAccount{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OidcAccess{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OktaAccess{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OneloginAccount{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PingoneAccount{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SAMLIdentityProvider{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(YandexAccess{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountIdentityProvidersAccessOnetimepin{}),
		},
	)
}

type AccountIdentityProvidersAccessOnetimepin struct {
	Config AccountIdentityProvidersAccessOnetimepinConfig `json:"config"`
	Type   AccountIdentityProvidersAccessOnetimepinType   `json:"type"`
	JSON   accountIdentityProvidersAccessOnetimepinJSON   `json:"-"`
	IdentityProviderAccount
}

// accountIdentityProvidersAccessOnetimepinJSON contains the JSON metadata for the
// struct [AccountIdentityProvidersAccessOnetimepin]
type accountIdentityProvidersAccessOnetimepinJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIdentityProvidersAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIdentityProvidersAccessOnetimepinJSON) RawJSON() string {
	return r.raw
}

func (r AccountIdentityProvidersAccessOnetimepin) implementsAccountIdentityProviders() {}

type AccountIdentityProvidersAccessOnetimepinConfig struct {
	RedirectURL string                                             `json:"redirect_url"`
	JSON        accountIdentityProvidersAccessOnetimepinConfigJSON `json:"-"`
}

// accountIdentityProvidersAccessOnetimepinConfigJSON contains the JSON metadata
// for the struct [AccountIdentityProvidersAccessOnetimepinConfig]
type accountIdentityProvidersAccessOnetimepinConfigJSON struct {
	RedirectURL apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIdentityProvidersAccessOnetimepinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIdentityProvidersAccessOnetimepinConfigJSON) RawJSON() string {
	return r.raw
}

type AccountIdentityProvidersAccessOnetimepinType string

const (
	AccountIdentityProvidersAccessOnetimepinTypeOnetimepin AccountIdentityProvidersAccessOnetimepinType = "onetimepin"
)

func (r AccountIdentityProvidersAccessOnetimepinType) IsKnown() bool {
	switch r {
	case AccountIdentityProvidersAccessOnetimepinTypeOnetimepin:
		return true
	}
	return false
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountIdentityProvidersType string

const (
	AccountIdentityProvidersTypeOnetimepin AccountIdentityProvidersType = "onetimepin"
	AccountIdentityProvidersTypeAzureAd    AccountIdentityProvidersType = "azureAD"
	AccountIdentityProvidersTypeSAML       AccountIdentityProvidersType = "saml"
	AccountIdentityProvidersTypeCentrify   AccountIdentityProvidersType = "centrify"
	AccountIdentityProvidersTypeFacebook   AccountIdentityProvidersType = "facebook"
	AccountIdentityProvidersTypeGitHub     AccountIdentityProvidersType = "github"
	AccountIdentityProvidersTypeGoogleApps AccountIdentityProvidersType = "google-apps"
	AccountIdentityProvidersTypeGoogle     AccountIdentityProvidersType = "google"
	AccountIdentityProvidersTypeLinkedin   AccountIdentityProvidersType = "linkedin"
	AccountIdentityProvidersTypeOidc       AccountIdentityProvidersType = "oidc"
	AccountIdentityProvidersTypeOkta       AccountIdentityProvidersType = "okta"
	AccountIdentityProvidersTypeOnelogin   AccountIdentityProvidersType = "onelogin"
	AccountIdentityProvidersTypePingone    AccountIdentityProvidersType = "pingone"
	AccountIdentityProvidersTypeYandex     AccountIdentityProvidersType = "yandex"
)

func (r AccountIdentityProvidersType) IsKnown() bool {
	switch r {
	case AccountIdentityProvidersTypeOnetimepin, AccountIdentityProvidersTypeAzureAd, AccountIdentityProvidersTypeSAML, AccountIdentityProvidersTypeCentrify, AccountIdentityProvidersTypeFacebook, AccountIdentityProvidersTypeGitHub, AccountIdentityProvidersTypeGoogleApps, AccountIdentityProvidersTypeGoogle, AccountIdentityProvidersTypeLinkedin, AccountIdentityProvidersTypeOidc, AccountIdentityProvidersTypeOkta, AccountIdentityProvidersTypeOnelogin, AccountIdentityProvidersTypePingone, AccountIdentityProvidersTypeYandex:
		return true
	}
	return false
}

// Satisfied by [AzureAdParam], [CentrifyIdentityProviderParam],
// [FacebookAccountParam], [GitHubAccessParam], [GoogleAccountParam],
// [GoogleAppsIdentityProviderParam], [LinkedinAccountParam], [OidcAccessParam],
// [OktaAccessParam], [OneloginAccountParam], [PingoneAccountParam],
// [SAMLIdentityProviderParam], [YandexAccessParam],
// [AccountIdentityProvidersAccessOnetimepinParam].
type AccountIdentityProvidersUnionParam interface {
	implementsAccountIdentityProvidersUnionParam()
}

type AccountIdentityProvidersAccessOnetimepinParam struct {
	Config param.Field[AccountIdentityProvidersAccessOnetimepinConfigParam] `json:"config"`
	Type   param.Field[AccountIdentityProvidersAccessOnetimepinType]        `json:"type"`
	IdentityProviderAccountParam
}

func (r AccountIdentityProvidersAccessOnetimepinParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountIdentityProvidersAccessOnetimepinParam) implementsAccountIdentityProvidersUnionParam() {
}

type AccountIdentityProvidersAccessOnetimepinConfigParam struct {
}

func (r AccountIdentityProvidersAccessOnetimepinConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AzureAd struct {
	Config AzureAdConfig `json:"config"`
	JSON   azureAdJSON   `json:"-"`
	IdentityProviderAccount
}

// azureAdJSON contains the JSON metadata for the struct [AzureAd]
type azureAdJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r azureAdJSON) RawJSON() string {
	return r.raw
}

func (r AzureAd) implementsAccountIdentityProviders() {}

func (r AzureAd) implementsAccountAccessIdentityProviderListResponseResult() {}

type AzureAdConfig struct {
	// Should Cloudflare try to load authentication contexts from your account
	ConditionalAccessEnabled bool `json:"conditional_access_enabled"`
	// Your Azure directory uuid
	DirectoryID string `json:"directory_id"`
	// Indicates the type of user interaction that is required. prompt=login forces the
	// user to enter their credentials on that request, negating single-sign on.
	// prompt=none is the opposite. It ensures that the user isn't presented with any
	// interactive prompt. If the request can't be completed silently by using
	// single-sign on, the Microsoft identity platform returns an interaction_required
	// error. prompt=select_account interrupts single sign-on providing account
	// selection experience listing all the accounts either in session or any
	// remembered account or an option to choose to use a different account altogether.
	Prompt AzureAdConfigPrompt `json:"prompt"`
	// Should Cloudflare try to load groups from your account
	SupportGroups bool              `json:"support_groups"`
	JSON          azureAdConfigJSON `json:"-"`
	GenericOAuthConfig
	CustomClaimsSupport
}

// azureAdConfigJSON contains the JSON metadata for the struct [AzureAdConfig]
type azureAdConfigJSON struct {
	ConditionalAccessEnabled apijson.Field
	DirectoryID              apijson.Field
	Prompt                   apijson.Field
	SupportGroups            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r azureAdConfigJSON) RawJSON() string {
	return r.raw
}

// Indicates the type of user interaction that is required. prompt=login forces the
// user to enter their credentials on that request, negating single-sign on.
// prompt=none is the opposite. It ensures that the user isn't presented with any
// interactive prompt. If the request can't be completed silently by using
// single-sign on, the Microsoft identity platform returns an interaction_required
// error. prompt=select_account interrupts single sign-on providing account
// selection experience listing all the accounts either in session or any
// remembered account or an option to choose to use a different account altogether.
type AzureAdConfigPrompt string

const (
	AzureAdConfigPromptLogin         AzureAdConfigPrompt = "login"
	AzureAdConfigPromptSelectAccount AzureAdConfigPrompt = "select_account"
	AzureAdConfigPromptNone          AzureAdConfigPrompt = "none"
)

func (r AzureAdConfigPrompt) IsKnown() bool {
	switch r {
	case AzureAdConfigPromptLogin, AzureAdConfigPromptSelectAccount, AzureAdConfigPromptNone:
		return true
	}
	return false
}

type AzureAdParam struct {
	Config param.Field[AzureAdConfigParam] `json:"config"`
	IdentityProviderAccountParam
}

func (r AzureAdParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AzureAdParam) implementsAccountIdentityProvidersUnionParam() {}

type AzureAdConfigParam struct {
	// Should Cloudflare try to load authentication contexts from your account
	ConditionalAccessEnabled param.Field[bool] `json:"conditional_access_enabled"`
	// Your Azure directory uuid
	DirectoryID param.Field[string] `json:"directory_id"`
	// Indicates the type of user interaction that is required. prompt=login forces the
	// user to enter their credentials on that request, negating single-sign on.
	// prompt=none is the opposite. It ensures that the user isn't presented with any
	// interactive prompt. If the request can't be completed silently by using
	// single-sign on, the Microsoft identity platform returns an interaction_required
	// error. prompt=select_account interrupts single sign-on providing account
	// selection experience listing all the accounts either in session or any
	// remembered account or an option to choose to use a different account altogether.
	Prompt param.Field[AzureAdConfigPrompt] `json:"prompt"`
	// Should Cloudflare try to load groups from your account
	SupportGroups param.Field[bool] `json:"support_groups"`
	GenericOAuthConfigParam
	CustomClaimsSupportParam
}

func (r AzureAdConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CentrifyIdentityProvider struct {
	Config CentrifyIdentityProviderConfig `json:"config"`
	JSON   centrifyIdentityProviderJSON   `json:"-"`
	IdentityProviderAccount
}

// centrifyIdentityProviderJSON contains the JSON metadata for the struct
// [CentrifyIdentityProvider]
type centrifyIdentityProviderJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CentrifyIdentityProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r centrifyIdentityProviderJSON) RawJSON() string {
	return r.raw
}

func (r CentrifyIdentityProvider) implementsAccountIdentityProviders() {}

func (r CentrifyIdentityProvider) implementsAccountAccessIdentityProviderListResponseResult() {}

type CentrifyIdentityProviderConfig struct {
	// Your centrify account url
	CentrifyAccount string `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID string                             `json:"centrify_app_id"`
	JSON          centrifyIdentityProviderConfigJSON `json:"-"`
	GenericOAuthConfig
	CustomClaimsSupport
}

// centrifyIdentityProviderConfigJSON contains the JSON metadata for the struct
// [CentrifyIdentityProviderConfig]
type centrifyIdentityProviderConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CentrifyIdentityProviderConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r centrifyIdentityProviderConfigJSON) RawJSON() string {
	return r.raw
}

type CentrifyIdentityProviderParam struct {
	Config param.Field[CentrifyIdentityProviderConfigParam] `json:"config"`
	IdentityProviderAccountParam
}

func (r CentrifyIdentityProviderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CentrifyIdentityProviderParam) implementsAccountIdentityProvidersUnionParam() {}

type CentrifyIdentityProviderConfigParam struct {
	// Your centrify account url
	CentrifyAccount param.Field[string] `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID param.Field[string] `json:"centrify_app_id"`
	GenericOAuthConfigParam
	CustomClaimsSupportParam
}

func (r CentrifyIdentityProviderConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomClaimsSupport struct {
	// Custom claims
	Claims []string `json:"claims"`
	// The claim name for email in the id_token response.
	EmailClaimName string                  `json:"email_claim_name"`
	JSON           customClaimsSupportJSON `json:"-"`
}

// customClaimsSupportJSON contains the JSON metadata for the struct
// [CustomClaimsSupport]
type customClaimsSupportJSON struct {
	Claims         apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomClaimsSupport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customClaimsSupportJSON) RawJSON() string {
	return r.raw
}

type CustomClaimsSupportParam struct {
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
}

func (r CustomClaimsSupportParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FacebookAccount struct {
	Config GenericOAuthConfig  `json:"config"`
	JSON   facebookAccountJSON `json:"-"`
	IdentityProviderAccount
}

// facebookAccountJSON contains the JSON metadata for the struct [FacebookAccount]
type facebookAccountJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FacebookAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r facebookAccountJSON) RawJSON() string {
	return r.raw
}

func (r FacebookAccount) implementsAccountIdentityProviders() {}

func (r FacebookAccount) implementsAccountAccessIdentityProviderListResponseResult() {}

type FacebookAccountParam struct {
	Config param.Field[GenericOAuthConfigParam] `json:"config"`
	IdentityProviderAccountParam
}

func (r FacebookAccountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FacebookAccountParam) implementsAccountIdentityProvidersUnionParam() {}

type GenericOAuthConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                 `json:"client_secret"`
	JSON         genericOAuthConfigJSON `json:"-"`
}

// genericOAuthConfigJSON contains the JSON metadata for the struct
// [GenericOAuthConfig]
type genericOAuthConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GenericOAuthConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r genericOAuthConfigJSON) RawJSON() string {
	return r.raw
}

type GenericOAuthConfigParam struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r GenericOAuthConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GitHubAccess struct {
	Config GenericOAuthConfig `json:"config"`
	JSON   githubAccessJSON   `json:"-"`
	IdentityProviderAccount
}

// githubAccessJSON contains the JSON metadata for the struct [GitHubAccess]
type githubAccessJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GitHubAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r githubAccessJSON) RawJSON() string {
	return r.raw
}

func (r GitHubAccess) implementsAccountIdentityProviders() {}

func (r GitHubAccess) implementsAccountAccessIdentityProviderListResponseResult() {}

type GitHubAccessParam struct {
	Config param.Field[GenericOAuthConfigParam] `json:"config"`
	IdentityProviderAccountParam
}

func (r GitHubAccessParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GitHubAccessParam) implementsAccountIdentityProvidersUnionParam() {}

type GoogleAccount struct {
	Config GoogleAccountConfig `json:"config"`
	JSON   googleAccountJSON   `json:"-"`
	IdentityProviderAccount
}

// googleAccountJSON contains the JSON metadata for the struct [GoogleAccount]
type googleAccountJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GoogleAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r googleAccountJSON) RawJSON() string {
	return r.raw
}

func (r GoogleAccount) implementsAccountIdentityProviders() {}

func (r GoogleAccount) implementsAccountAccessIdentityProviderListResponseResult() {}

type GoogleAccountConfig struct {
	JSON googleAccountConfigJSON `json:"-"`
	GenericOAuthConfig
	CustomClaimsSupport
}

// googleAccountConfigJSON contains the JSON metadata for the struct
// [GoogleAccountConfig]
type googleAccountConfigJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GoogleAccountConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r googleAccountConfigJSON) RawJSON() string {
	return r.raw
}

type GoogleAccountParam struct {
	Config param.Field[GoogleAccountConfigParam] `json:"config"`
	IdentityProviderAccountParam
}

func (r GoogleAccountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GoogleAccountParam) implementsAccountIdentityProvidersUnionParam() {}

type GoogleAccountConfigParam struct {
	GenericOAuthConfigParam
	CustomClaimsSupportParam
}

func (r GoogleAccountConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GoogleAppsIdentityProvider struct {
	Config GoogleAppsIdentityProviderConfig `json:"config"`
	JSON   googleAppsIdentityProviderJSON   `json:"-"`
	IdentityProviderAccount
}

// googleAppsIdentityProviderJSON contains the JSON metadata for the struct
// [GoogleAppsIdentityProvider]
type googleAppsIdentityProviderJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GoogleAppsIdentityProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r googleAppsIdentityProviderJSON) RawJSON() string {
	return r.raw
}

func (r GoogleAppsIdentityProvider) implementsAccountIdentityProviders() {}

func (r GoogleAppsIdentityProvider) implementsAccountAccessIdentityProviderListResponseResult() {}

type GoogleAppsIdentityProviderConfig struct {
	// Your companies TLD
	AppsDomain string                               `json:"apps_domain"`
	JSON       googleAppsIdentityProviderConfigJSON `json:"-"`
	GenericOAuthConfig
	CustomClaimsSupport
}

// googleAppsIdentityProviderConfigJSON contains the JSON metadata for the struct
// [GoogleAppsIdentityProviderConfig]
type googleAppsIdentityProviderConfigJSON struct {
	AppsDomain  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GoogleAppsIdentityProviderConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r googleAppsIdentityProviderConfigJSON) RawJSON() string {
	return r.raw
}

type GoogleAppsIdentityProviderParam struct {
	Config param.Field[GoogleAppsIdentityProviderConfigParam] `json:"config"`
	IdentityProviderAccountParam
}

func (r GoogleAppsIdentityProviderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GoogleAppsIdentityProviderParam) implementsAccountIdentityProvidersUnionParam() {}

type GoogleAppsIdentityProviderConfigParam struct {
	// Your companies TLD
	AppsDomain param.Field[string] `json:"apps_domain"`
	GenericOAuthConfigParam
	CustomClaimsSupportParam
}

func (r GoogleAppsIdentityProviderConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderAccount struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config interface{} `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderAccountType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderAccountScimConfig `json:"scim_config"`
	JSON       identityProviderAccountJSON       `json:"-"`
}

// identityProviderAccountJSON contains the JSON metadata for the struct
// [IdentityProviderAccount]
type identityProviderAccountJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccountJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccountType string

const (
	IdentityProviderAccountTypeOnetimepin IdentityProviderAccountType = "onetimepin"
	IdentityProviderAccountTypeAzureAd    IdentityProviderAccountType = "azureAD"
	IdentityProviderAccountTypeSAML       IdentityProviderAccountType = "saml"
	IdentityProviderAccountTypeCentrify   IdentityProviderAccountType = "centrify"
	IdentityProviderAccountTypeFacebook   IdentityProviderAccountType = "facebook"
	IdentityProviderAccountTypeGitHub     IdentityProviderAccountType = "github"
	IdentityProviderAccountTypeGoogleApps IdentityProviderAccountType = "google-apps"
	IdentityProviderAccountTypeGoogle     IdentityProviderAccountType = "google"
	IdentityProviderAccountTypeLinkedin   IdentityProviderAccountType = "linkedin"
	IdentityProviderAccountTypeOidc       IdentityProviderAccountType = "oidc"
	IdentityProviderAccountTypeOkta       IdentityProviderAccountType = "okta"
	IdentityProviderAccountTypeOnelogin   IdentityProviderAccountType = "onelogin"
	IdentityProviderAccountTypePingone    IdentityProviderAccountType = "pingone"
	IdentityProviderAccountTypeYandex     IdentityProviderAccountType = "yandex"
)

func (r IdentityProviderAccountType) IsKnown() bool {
	switch r {
	case IdentityProviderAccountTypeOnetimepin, IdentityProviderAccountTypeAzureAd, IdentityProviderAccountTypeSAML, IdentityProviderAccountTypeCentrify, IdentityProviderAccountTypeFacebook, IdentityProviderAccountTypeGitHub, IdentityProviderAccountTypeGoogleApps, IdentityProviderAccountTypeGoogle, IdentityProviderAccountTypeLinkedin, IdentityProviderAccountTypeOidc, IdentityProviderAccountTypeOkta, IdentityProviderAccountTypeOnelogin, IdentityProviderAccountTypePingone, IdentityProviderAccountTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderAccountScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// Indicates how a SCIM event updates a user identity used for policy evaluation.
	// Use "automatic" to automatically update a user's identity and augment it with
	// fields from the SCIM user resource. Use "reauth" to force re-authentication on
	// group membership updates, user identity update will only occur after successful
	// re-authentication. With "reauth" identities will not contain fields from the
	// SCIM user resource. With "no_action" identities will not be changed by SCIM
	// updates in any way and users will not be prompted to reauthenticate.
	IdentityUpdateBehavior IdentityProviderAccountScimConfigIdentityUpdateBehavior `json:"identity_update_behavior"`
	// The base URL of Cloudflare's SCIM V2.0 API endpoint.
	ScimBaseURL string `json:"scim_base_url"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                  `json:"user_deprovision"`
	JSON            identityProviderAccountScimConfigJSON `json:"-"`
}

// identityProviderAccountScimConfigJSON contains the JSON metadata for the struct
// [IdentityProviderAccountScimConfig]
type identityProviderAccountScimConfigJSON struct {
	Enabled                apijson.Field
	IdentityUpdateBehavior apijson.Field
	ScimBaseURL            apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderAccountScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccountScimConfigJSON) RawJSON() string {
	return r.raw
}

// Indicates how a SCIM event updates a user identity used for policy evaluation.
// Use "automatic" to automatically update a user's identity and augment it with
// fields from the SCIM user resource. Use "reauth" to force re-authentication on
// group membership updates, user identity update will only occur after successful
// re-authentication. With "reauth" identities will not contain fields from the
// SCIM user resource. With "no_action" identities will not be changed by SCIM
// updates in any way and users will not be prompted to reauthenticate.
type IdentityProviderAccountScimConfigIdentityUpdateBehavior string

const (
	IdentityProviderAccountScimConfigIdentityUpdateBehaviorAutomatic IdentityProviderAccountScimConfigIdentityUpdateBehavior = "automatic"
	IdentityProviderAccountScimConfigIdentityUpdateBehaviorReauth    IdentityProviderAccountScimConfigIdentityUpdateBehavior = "reauth"
	IdentityProviderAccountScimConfigIdentityUpdateBehaviorNoAction  IdentityProviderAccountScimConfigIdentityUpdateBehavior = "no_action"
)

func (r IdentityProviderAccountScimConfigIdentityUpdateBehavior) IsKnown() bool {
	switch r {
	case IdentityProviderAccountScimConfigIdentityUpdateBehaviorAutomatic, IdentityProviderAccountScimConfigIdentityUpdateBehaviorReauth, IdentityProviderAccountScimConfigIdentityUpdateBehaviorNoAction:
		return true
	}
	return false
}

type IdentityProviderAccountParam struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[interface{}] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderAccountType] `json:"type,required"`
	// UUID
	ID param.Field[string] `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderAccountScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderAccountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderAccountScimConfigParam struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// Indicates how a SCIM event updates a user identity used for policy evaluation.
	// Use "automatic" to automatically update a user's identity and augment it with
	// fields from the SCIM user resource. Use "reauth" to force re-authentication on
	// group membership updates, user identity update will only occur after successful
	// re-authentication. With "reauth" identities will not contain fields from the
	// SCIM user resource. With "no_action" identities will not be changed by SCIM
	// updates in any way and users will not be prompted to reauthenticate.
	IdentityUpdateBehavior param.Field[IdentityProviderAccountScimConfigIdentityUpdateBehavior] `json:"identity_update_behavior"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r IdentityProviderAccountScimConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LinkedinAccount struct {
	Config GenericOAuthConfig  `json:"config"`
	JSON   linkedinAccountJSON `json:"-"`
	IdentityProviderAccount
}

// linkedinAccountJSON contains the JSON metadata for the struct [LinkedinAccount]
type linkedinAccountJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LinkedinAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkedinAccountJSON) RawJSON() string {
	return r.raw
}

func (r LinkedinAccount) implementsAccountIdentityProviders() {}

func (r LinkedinAccount) implementsAccountAccessIdentityProviderListResponseResult() {}

type LinkedinAccountParam struct {
	Config param.Field[GenericOAuthConfigParam] `json:"config"`
	IdentityProviderAccountParam
}

func (r LinkedinAccountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r LinkedinAccountParam) implementsAccountIdentityProvidersUnionParam() {}

type OidcAccess struct {
	Config OidcAccessConfig `json:"config"`
	JSON   oidcAccessJSON   `json:"-"`
	IdentityProviderAccount
}

// oidcAccessJSON contains the JSON metadata for the struct [OidcAccess]
type oidcAccessJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OidcAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oidcAccessJSON) RawJSON() string {
	return r.raw
}

func (r OidcAccess) implementsAccountIdentityProviders() {}

func (r OidcAccess) implementsAccountAccessIdentityProviderListResponseResult() {}

type OidcAccessConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL string `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL string `json:"certs_url"`
	// Enable Proof Key for Code Exchange (PKCE)
	PkceEnabled bool `json:"pkce_enabled"`
	// OAuth scopes
	Scopes []string `json:"scopes"`
	// The token_endpoint URL of your IdP
	TokenURL string               `json:"token_url"`
	JSON     oidcAccessConfigJSON `json:"-"`
	GenericOAuthConfig
	CustomClaimsSupport
}

// oidcAccessConfigJSON contains the JSON metadata for the struct
// [OidcAccessConfig]
type oidcAccessConfigJSON struct {
	AuthURL     apijson.Field
	CertsURL    apijson.Field
	PkceEnabled apijson.Field
	Scopes      apijson.Field
	TokenURL    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OidcAccessConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oidcAccessConfigJSON) RawJSON() string {
	return r.raw
}

type OidcAccessParam struct {
	Config param.Field[OidcAccessConfigParam] `json:"config"`
	IdentityProviderAccountParam
}

func (r OidcAccessParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OidcAccessParam) implementsAccountIdentityProvidersUnionParam() {}

type OidcAccessConfigParam struct {
	// The authorization_endpoint URL of your IdP
	AuthURL param.Field[string] `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL param.Field[string] `json:"certs_url"`
	// Enable Proof Key for Code Exchange (PKCE)
	PkceEnabled param.Field[bool] `json:"pkce_enabled"`
	// OAuth scopes
	Scopes param.Field[[]string] `json:"scopes"`
	// The token_endpoint URL of your IdP
	TokenURL param.Field[string] `json:"token_url"`
	GenericOAuthConfigParam
	CustomClaimsSupportParam
}

func (r OidcAccessConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OktaAccess struct {
	Config OktaAccessConfig `json:"config"`
	JSON   oktaAccessJSON   `json:"-"`
	IdentityProviderAccount
}

// oktaAccessJSON contains the JSON metadata for the struct [OktaAccess]
type oktaAccessJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OktaAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oktaAccessJSON) RawJSON() string {
	return r.raw
}

func (r OktaAccess) implementsAccountIdentityProviders() {}

func (r OktaAccess) implementsAccountAccessIdentityProviderListResponseResult() {}

type OktaAccessConfig struct {
	// Your okta authorization server id
	AuthorizationServerID string `json:"authorization_server_id"`
	// Your okta account url
	OktaAccount string               `json:"okta_account"`
	JSON        oktaAccessConfigJSON `json:"-"`
	GenericOAuthConfig
	CustomClaimsSupport
}

// oktaAccessConfigJSON contains the JSON metadata for the struct
// [OktaAccessConfig]
type oktaAccessConfigJSON struct {
	AuthorizationServerID apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *OktaAccessConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oktaAccessConfigJSON) RawJSON() string {
	return r.raw
}

type OktaAccessParam struct {
	Config param.Field[OktaAccessConfigParam] `json:"config"`
	IdentityProviderAccountParam
}

func (r OktaAccessParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OktaAccessParam) implementsAccountIdentityProvidersUnionParam() {}

type OktaAccessConfigParam struct {
	// Your okta authorization server id
	AuthorizationServerID param.Field[string] `json:"authorization_server_id"`
	// Your okta account url
	OktaAccount param.Field[string] `json:"okta_account"`
	GenericOAuthConfigParam
	CustomClaimsSupportParam
}

func (r OktaAccessConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OneloginAccount struct {
	Config OneloginAccountConfig `json:"config"`
	JSON   oneloginAccountJSON   `json:"-"`
	IdentityProviderAccount
}

// oneloginAccountJSON contains the JSON metadata for the struct [OneloginAccount]
type oneloginAccountJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OneloginAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oneloginAccountJSON) RawJSON() string {
	return r.raw
}

func (r OneloginAccount) implementsAccountIdentityProviders() {}

func (r OneloginAccount) implementsAccountAccessIdentityProviderListResponseResult() {}

type OneloginAccountConfig struct {
	// Your OneLogin account url
	OneloginAccount string                    `json:"onelogin_account"`
	JSON            oneloginAccountConfigJSON `json:"-"`
	GenericOAuthConfig
	CustomClaimsSupport
}

// oneloginAccountConfigJSON contains the JSON metadata for the struct
// [OneloginAccountConfig]
type oneloginAccountConfigJSON struct {
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *OneloginAccountConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oneloginAccountConfigJSON) RawJSON() string {
	return r.raw
}

type OneloginAccountParam struct {
	Config param.Field[OneloginAccountConfigParam] `json:"config"`
	IdentityProviderAccountParam
}

func (r OneloginAccountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OneloginAccountParam) implementsAccountIdentityProvidersUnionParam() {}

type OneloginAccountConfigParam struct {
	// Your OneLogin account url
	OneloginAccount param.Field[string] `json:"onelogin_account"`
	GenericOAuthConfigParam
	CustomClaimsSupportParam
}

func (r OneloginAccountConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PingoneAccount struct {
	Config PingoneAccountConfig `json:"config"`
	JSON   pingoneAccountJSON   `json:"-"`
	IdentityProviderAccount
}

// pingoneAccountJSON contains the JSON metadata for the struct [PingoneAccount]
type pingoneAccountJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PingoneAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pingoneAccountJSON) RawJSON() string {
	return r.raw
}

func (r PingoneAccount) implementsAccountIdentityProviders() {}

func (r PingoneAccount) implementsAccountAccessIdentityProviderListResponseResult() {}

type PingoneAccountConfig struct {
	// Your PingOne environment identifier
	PingEnvID string                   `json:"ping_env_id"`
	JSON      pingoneAccountConfigJSON `json:"-"`
	GenericOAuthConfig
	CustomClaimsSupport
}

// pingoneAccountConfigJSON contains the JSON metadata for the struct
// [PingoneAccountConfig]
type pingoneAccountConfigJSON struct {
	PingEnvID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PingoneAccountConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pingoneAccountConfigJSON) RawJSON() string {
	return r.raw
}

type PingoneAccountParam struct {
	Config param.Field[PingoneAccountConfigParam] `json:"config"`
	IdentityProviderAccountParam
}

func (r PingoneAccountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PingoneAccountParam) implementsAccountIdentityProvidersUnionParam() {}

type PingoneAccountConfigParam struct {
	// Your PingOne environment identifier
	PingEnvID param.Field[string] `json:"ping_env_id"`
	GenericOAuthConfigParam
	CustomClaimsSupportParam
}

func (r PingoneAccountConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SAMLIdentityProvider struct {
	Config SAMLIdentityProviderConfig `json:"config"`
	JSON   samlIdentityProviderJSON   `json:"-"`
	IdentityProviderAccount
}

// samlIdentityProviderJSON contains the JSON metadata for the struct
// [SAMLIdentityProvider]
type samlIdentityProviderJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SAMLIdentityProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r samlIdentityProviderJSON) RawJSON() string {
	return r.raw
}

func (r SAMLIdentityProvider) implementsAccountIdentityProviders() {}

func (r SAMLIdentityProvider) implementsAccountAccessIdentityProviderListResponseResult() {}

type SAMLIdentityProviderConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []SAMLIdentityProviderConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                         `json:"sso_target_url"`
	JSON         samlIdentityProviderConfigJSON `json:"-"`
}

// samlIdentityProviderConfigJSON contains the JSON metadata for the struct
// [SAMLIdentityProviderConfig]
type samlIdentityProviderConfigJSON struct {
	Attributes         apijson.Field
	EmailAttributeName apijson.Field
	HeaderAttributes   apijson.Field
	IdpPublicCerts     apijson.Field
	IssuerURL          apijson.Field
	SignRequest        apijson.Field
	SSOTargetURL       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SAMLIdentityProviderConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r samlIdentityProviderConfigJSON) RawJSON() string {
	return r.raw
}

type SAMLIdentityProviderConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                        `json:"header_name"`
	JSON       samlIdentityProviderConfigHeaderAttributeJSON `json:"-"`
}

// samlIdentityProviderConfigHeaderAttributeJSON contains the JSON metadata for the
// struct [SAMLIdentityProviderConfigHeaderAttribute]
type samlIdentityProviderConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SAMLIdentityProviderConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r samlIdentityProviderConfigHeaderAttributeJSON) RawJSON() string {
	return r.raw
}

type SAMLIdentityProviderParam struct {
	Config param.Field[SAMLIdentityProviderConfigParam] `json:"config"`
	IdentityProviderAccountParam
}

func (r SAMLIdentityProviderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SAMLIdentityProviderParam) implementsAccountIdentityProvidersUnionParam() {}

type SAMLIdentityProviderConfigParam struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]SAMLIdentityProviderConfigHeaderAttributeParam] `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts param.Field[[]string] `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL param.Field[string] `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest param.Field[bool] `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL param.Field[string] `json:"sso_target_url"`
}

func (r SAMLIdentityProviderConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SAMLIdentityProviderConfigHeaderAttributeParam struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r SAMLIdentityProviderConfigHeaderAttributeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SingleResponseProvider struct {
	Result AccountIdentityProviders   `json:"result"`
	JSON   singleResponseProviderJSON `json:"-"`
	APIResponseSingleAccess
}

// singleResponseProviderJSON contains the JSON metadata for the struct
// [SingleResponseProvider]
type singleResponseProviderJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseProviderJSON) RawJSON() string {
	return r.raw
}

type YandexAccess struct {
	Config GenericOAuthConfig `json:"config"`
	JSON   yandexAccessJSON   `json:"-"`
	IdentityProviderAccount
}

// yandexAccessJSON contains the JSON metadata for the struct [YandexAccess]
type yandexAccessJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *YandexAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r yandexAccessJSON) RawJSON() string {
	return r.raw
}

func (r YandexAccess) implementsAccountIdentityProviders() {}

func (r YandexAccess) implementsAccountAccessIdentityProviderListResponseResult() {}

type YandexAccessParam struct {
	Config param.Field[GenericOAuthConfigParam] `json:"config"`
	IdentityProviderAccountParam
}

func (r YandexAccessParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r YandexAccessParam) implementsAccountIdentityProvidersUnionParam() {}

type AccountAccessIdentityProviderListResponse struct {
	Result []AccountAccessIdentityProviderListResponseResult `json:"result"`
	JSON   accountAccessIdentityProviderListResponseJSON     `json:"-"`
	APIResponseCollectionAccess
}

// accountAccessIdentityProviderListResponseJSON contains the JSON metadata for the
// struct [AccountAccessIdentityProviderListResponse]
type accountAccessIdentityProviderListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessIdentityProviderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessIdentityProviderListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessIdentityProviderListResponseResult struct {
	// UUID
	ID string `json:"id"`
	// This field can have the runtime type of [interface{}].
	Config interface{} `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// This field can have the runtime type of [IdentityProviderAccountScimConfig].
	ScimConfig interface{} `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type  AccountAccessIdentityProviderListResponseResultType `json:"type"`
	JSON  accountAccessIdentityProviderListResponseResultJSON `json:"-"`
	union AccountAccessIdentityProviderListResponseResultUnion
}

// accountAccessIdentityProviderListResponseResultJSON contains the JSON metadata
// for the struct [AccountAccessIdentityProviderListResponseResult]
type accountAccessIdentityProviderListResponseResultJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r accountAccessIdentityProviderListResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *AccountAccessIdentityProviderListResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = AccountAccessIdentityProviderListResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountAccessIdentityProviderListResponseResultUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [AzureAd], [CentrifyIdentityProvider],
// [FacebookAccount], [GitHubAccess], [GoogleAccount],
// [GoogleAppsIdentityProvider], [LinkedinAccount], [OidcAccess], [OktaAccess],
// [OneloginAccount], [PingoneAccount], [SAMLIdentityProvider], [YandexAccess].
func (r AccountAccessIdentityProviderListResponseResult) AsUnion() AccountAccessIdentityProviderListResponseResultUnion {
	return r.union
}

// Union satisfied by [AzureAd], [CentrifyIdentityProvider], [FacebookAccount],
// [GitHubAccess], [GoogleAccount], [GoogleAppsIdentityProvider],
// [LinkedinAccount], [OidcAccess], [OktaAccess], [OneloginAccount],
// [PingoneAccount], [SAMLIdentityProvider] or [YandexAccess].
type AccountAccessIdentityProviderListResponseResultUnion interface {
	implementsAccountAccessIdentityProviderListResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAccessIdentityProviderListResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AzureAd{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CentrifyIdentityProvider{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FacebookAccount{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GitHubAccess{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GoogleAccount{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GoogleAppsIdentityProvider{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LinkedinAccount{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OidcAccess{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OktaAccess{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OneloginAccount{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PingoneAccount{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SAMLIdentityProvider{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(YandexAccess{}),
		},
	)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccountAccessIdentityProviderListResponseResultType string

const (
	AccountAccessIdentityProviderListResponseResultTypeOnetimepin AccountAccessIdentityProviderListResponseResultType = "onetimepin"
	AccountAccessIdentityProviderListResponseResultTypeAzureAd    AccountAccessIdentityProviderListResponseResultType = "azureAD"
	AccountAccessIdentityProviderListResponseResultTypeSAML       AccountAccessIdentityProviderListResponseResultType = "saml"
	AccountAccessIdentityProviderListResponseResultTypeCentrify   AccountAccessIdentityProviderListResponseResultType = "centrify"
	AccountAccessIdentityProviderListResponseResultTypeFacebook   AccountAccessIdentityProviderListResponseResultType = "facebook"
	AccountAccessIdentityProviderListResponseResultTypeGitHub     AccountAccessIdentityProviderListResponseResultType = "github"
	AccountAccessIdentityProviderListResponseResultTypeGoogleApps AccountAccessIdentityProviderListResponseResultType = "google-apps"
	AccountAccessIdentityProviderListResponseResultTypeGoogle     AccountAccessIdentityProviderListResponseResultType = "google"
	AccountAccessIdentityProviderListResponseResultTypeLinkedin   AccountAccessIdentityProviderListResponseResultType = "linkedin"
	AccountAccessIdentityProviderListResponseResultTypeOidc       AccountAccessIdentityProviderListResponseResultType = "oidc"
	AccountAccessIdentityProviderListResponseResultTypeOkta       AccountAccessIdentityProviderListResponseResultType = "okta"
	AccountAccessIdentityProviderListResponseResultTypeOnelogin   AccountAccessIdentityProviderListResponseResultType = "onelogin"
	AccountAccessIdentityProviderListResponseResultTypePingone    AccountAccessIdentityProviderListResponseResultType = "pingone"
	AccountAccessIdentityProviderListResponseResultTypeYandex     AccountAccessIdentityProviderListResponseResultType = "yandex"
)

func (r AccountAccessIdentityProviderListResponseResultType) IsKnown() bool {
	switch r {
	case AccountAccessIdentityProviderListResponseResultTypeOnetimepin, AccountAccessIdentityProviderListResponseResultTypeAzureAd, AccountAccessIdentityProviderListResponseResultTypeSAML, AccountAccessIdentityProviderListResponseResultTypeCentrify, AccountAccessIdentityProviderListResponseResultTypeFacebook, AccountAccessIdentityProviderListResponseResultTypeGitHub, AccountAccessIdentityProviderListResponseResultTypeGoogleApps, AccountAccessIdentityProviderListResponseResultTypeGoogle, AccountAccessIdentityProviderListResponseResultTypeLinkedin, AccountAccessIdentityProviderListResponseResultTypeOidc, AccountAccessIdentityProviderListResponseResultTypeOkta, AccountAccessIdentityProviderListResponseResultTypeOnelogin, AccountAccessIdentityProviderListResponseResultTypePingone, AccountAccessIdentityProviderListResponseResultTypeYandex:
		return true
	}
	return false
}

type AccountAccessIdentityProviderNewParams struct {
	AccountIdentityProviders AccountIdentityProvidersUnionParam `json:"account_identity_providers,required"`
}

func (r AccountAccessIdentityProviderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AccountIdentityProviders)
}

type AccountAccessIdentityProviderUpdateParams struct {
	AccountIdentityProviders AccountIdentityProvidersUnionParam `json:"account_identity_providers,required"`
}

func (r AccountAccessIdentityProviderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AccountIdentityProviders)
}

type AccountAccessIdentityProviderListParams struct {
	// Indicates to Access to only retrieve identity providers that have the System for
	// Cross-Domain Identity Management (SCIM) enabled.
	ScimEnabled param.Field[string] `query:"scim_enabled"`
}

// URLQuery serializes [AccountAccessIdentityProviderListParams]'s query parameters
// as `url.Values`.
func (r AccountAccessIdentityProviderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
