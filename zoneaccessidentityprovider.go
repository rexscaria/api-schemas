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

// ZoneAccessIdentityProviderService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAccessIdentityProviderService] method instead.
type ZoneAccessIdentityProviderService struct {
	Options []option.RequestOption
}

// NewZoneAccessIdentityProviderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneAccessIdentityProviderService(opts ...option.RequestOption) (r *ZoneAccessIdentityProviderService) {
	r = &ZoneAccessIdentityProviderService{}
	r.Options = opts
	return
}

// Adds a new identity provider to Access.
func (r *ZoneAccessIdentityProviderService) New(ctx context.Context, zoneID string, body ZoneAccessIdentityProviderNewParams, opts ...option.RequestOption) (res *SingleResponseProviderZone, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/identity_providers", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a configured identity provider.
func (r *ZoneAccessIdentityProviderService) Get(ctx context.Context, zoneID string, identityProviderID string, opts ...option.RequestOption) (res *SingleResponseProviderZone, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identityProviderID == "" {
		err = errors.New("missing required identity_provider_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/identity_providers/%s", zoneID, identityProviderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured identity provider.
func (r *ZoneAccessIdentityProviderService) Update(ctx context.Context, zoneID string, identityProviderID string, body ZoneAccessIdentityProviderUpdateParams, opts ...option.RequestOption) (res *SingleResponseProviderZone, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identityProviderID == "" {
		err = errors.New("missing required identity_provider_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/identity_providers/%s", zoneID, identityProviderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all configured identity providers.
func (r *ZoneAccessIdentityProviderService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneAccessIdentityProviderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/identity_providers", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an identity provider from Access.
func (r *ZoneAccessIdentityProviderService) Delete(ctx context.Context, zoneID string, identityProviderID string, opts ...option.RequestOption) (res *IDResponseApps, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identityProviderID == "" {
		err = errors.New("missing required identity_provider_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/identity_providers/%s", zoneID, identityProviderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AzureAd struct {
	Config AzureAdConfig `json:"config"`
	JSON   azureAdJSON   `json:"-"`
	IdentityProviderZone
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

func (r AzureAd) implementsZoneIdentityProviders() {}

func (r AzureAd) implementsZoneAccessIdentityProviderListResponseResult() {}

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
	IdentityProviderZoneParam
}

func (r AzureAdParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AzureAdParam) implementsZoneIdentityProvidersUnionParam() {}

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
}

func (r AzureAdConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CentrifySchemaProvider struct {
	Config CentrifySchemaProviderConfig `json:"config"`
	JSON   centrifySchemaProviderJSON   `json:"-"`
	IdentityProviderZone
}

// centrifySchemaProviderJSON contains the JSON metadata for the struct
// [CentrifySchemaProvider]
type centrifySchemaProviderJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CentrifySchemaProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r centrifySchemaProviderJSON) RawJSON() string {
	return r.raw
}

func (r CentrifySchemaProvider) implementsZoneIdentityProviders() {}

func (r CentrifySchemaProvider) implementsZoneAccessIdentityProviderListResponseResult() {}

type CentrifySchemaProviderConfig struct {
	// Your centrify account url
	CentrifyAccount string `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID string                           `json:"centrify_app_id"`
	JSON          centrifySchemaProviderConfigJSON `json:"-"`
	GenericOAuthConfig
}

// centrifySchemaProviderConfigJSON contains the JSON metadata for the struct
// [CentrifySchemaProviderConfig]
type centrifySchemaProviderConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CentrifySchemaProviderConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r centrifySchemaProviderConfigJSON) RawJSON() string {
	return r.raw
}

type CentrifySchemaProviderParam struct {
	Config param.Field[CentrifySchemaProviderConfigParam] `json:"config"`
	IdentityProviderZoneParam
}

func (r CentrifySchemaProviderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CentrifySchemaProviderParam) implementsZoneIdentityProvidersUnionParam() {}

type CentrifySchemaProviderConfigParam struct {
	// Your centrify account url
	CentrifyAccount param.Field[string] `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID param.Field[string] `json:"centrify_app_id"`
	GenericOAuthConfigParam
}

func (r CentrifySchemaProviderConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FacebookZone struct {
	Config GenericOAuthConfig `json:"config"`
	JSON   facebookZoneJSON   `json:"-"`
	IdentityProviderZone
}

// facebookZoneJSON contains the JSON metadata for the struct [FacebookZone]
type facebookZoneJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FacebookZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r facebookZoneJSON) RawJSON() string {
	return r.raw
}

func (r FacebookZone) implementsZoneIdentityProviders() {}

func (r FacebookZone) implementsZoneAccessIdentityProviderListResponseResult() {}

type FacebookZoneParam struct {
	Config param.Field[GenericOAuthConfigParam] `json:"config"`
	IdentityProviderZoneParam
}

func (r FacebookZoneParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FacebookZoneParam) implementsZoneIdentityProvidersUnionParam() {}

type GitHubAccessSchema struct {
	Config GenericOAuthConfig     `json:"config"`
	JSON   githubAccessSchemaJSON `json:"-"`
	IdentityProviderZone
}

// githubAccessSchemaJSON contains the JSON metadata for the struct
// [GitHubAccessSchema]
type githubAccessSchemaJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GitHubAccessSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r githubAccessSchemaJSON) RawJSON() string {
	return r.raw
}

func (r GitHubAccessSchema) implementsZoneIdentityProviders() {}

func (r GitHubAccessSchema) implementsZoneAccessIdentityProviderListResponseResult() {}

type GitHubAccessSchemaParam struct {
	Config param.Field[GenericOAuthConfigParam] `json:"config"`
	IdentityProviderZoneParam
}

func (r GitHubAccessSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GitHubAccessSchemaParam) implementsZoneIdentityProvidersUnionParam() {}

type GoogleAppsSchemaIdentityProvider struct {
	Config GoogleAppsSchemaIdentityProviderConfig `json:"config"`
	JSON   googleAppsSchemaIdentityProviderJSON   `json:"-"`
	IdentityProviderZone
}

// googleAppsSchemaIdentityProviderJSON contains the JSON metadata for the struct
// [GoogleAppsSchemaIdentityProvider]
type googleAppsSchemaIdentityProviderJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GoogleAppsSchemaIdentityProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r googleAppsSchemaIdentityProviderJSON) RawJSON() string {
	return r.raw
}

func (r GoogleAppsSchemaIdentityProvider) implementsZoneIdentityProviders() {}

func (r GoogleAppsSchemaIdentityProvider) implementsZoneAccessIdentityProviderListResponseResult() {}

type GoogleAppsSchemaIdentityProviderConfig struct {
	// Your companies TLD
	AppsDomain string                                     `json:"apps_domain"`
	JSON       googleAppsSchemaIdentityProviderConfigJSON `json:"-"`
	GenericOAuthConfig
}

// googleAppsSchemaIdentityProviderConfigJSON contains the JSON metadata for the
// struct [GoogleAppsSchemaIdentityProviderConfig]
type googleAppsSchemaIdentityProviderConfigJSON struct {
	AppsDomain  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GoogleAppsSchemaIdentityProviderConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r googleAppsSchemaIdentityProviderConfigJSON) RawJSON() string {
	return r.raw
}

type GoogleAppsSchemaIdentityProviderParam struct {
	Config param.Field[GoogleAppsSchemaIdentityProviderConfigParam] `json:"config"`
	IdentityProviderZoneParam
}

func (r GoogleAppsSchemaIdentityProviderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GoogleAppsSchemaIdentityProviderParam) implementsZoneIdentityProvidersUnionParam() {}

type GoogleAppsSchemaIdentityProviderConfigParam struct {
	// Your companies TLD
	AppsDomain param.Field[string] `json:"apps_domain"`
	GenericOAuthConfigParam
}

func (r GoogleAppsSchemaIdentityProviderConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GoogleSchema struct {
	Config GenericOAuthConfig `json:"config"`
	JSON   googleSchemaJSON   `json:"-"`
	IdentityProviderZone
}

// googleSchemaJSON contains the JSON metadata for the struct [GoogleSchema]
type googleSchemaJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GoogleSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r googleSchemaJSON) RawJSON() string {
	return r.raw
}

func (r GoogleSchema) implementsZoneIdentityProviders() {}

func (r GoogleSchema) implementsZoneAccessIdentityProviderListResponseResult() {}

type GoogleSchemaParam struct {
	Config param.Field[GenericOAuthConfigParam] `json:"config"`
	IdentityProviderZoneParam
}

func (r GoogleSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GoogleSchemaParam) implementsZoneIdentityProvidersUnionParam() {}

type IdentityProviderZone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config interface{} `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderZoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderZoneScimConfig `json:"scim_config"`
	JSON       identityProviderZoneJSON       `json:"-"`
}

// identityProviderZoneJSON contains the JSON metadata for the struct
// [IdentityProviderZone]
type identityProviderZoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderZoneJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderZoneType string

const (
	IdentityProviderZoneTypeOnetimepin IdentityProviderZoneType = "onetimepin"
	IdentityProviderZoneTypeAzureAd    IdentityProviderZoneType = "azureAD"
	IdentityProviderZoneTypeSAML       IdentityProviderZoneType = "saml"
	IdentityProviderZoneTypeCentrify   IdentityProviderZoneType = "centrify"
	IdentityProviderZoneTypeFacebook   IdentityProviderZoneType = "facebook"
	IdentityProviderZoneTypeGitHub     IdentityProviderZoneType = "github"
	IdentityProviderZoneTypeGoogleApps IdentityProviderZoneType = "google-apps"
	IdentityProviderZoneTypeGoogle     IdentityProviderZoneType = "google"
	IdentityProviderZoneTypeLinkedin   IdentityProviderZoneType = "linkedin"
	IdentityProviderZoneTypeOidc       IdentityProviderZoneType = "oidc"
	IdentityProviderZoneTypeOkta       IdentityProviderZoneType = "okta"
	IdentityProviderZoneTypeOnelogin   IdentityProviderZoneType = "onelogin"
	IdentityProviderZoneTypePingone    IdentityProviderZoneType = "pingone"
	IdentityProviderZoneTypeYandex     IdentityProviderZoneType = "yandex"
)

func (r IdentityProviderZoneType) IsKnown() bool {
	switch r {
	case IdentityProviderZoneTypeOnetimepin, IdentityProviderZoneTypeAzureAd, IdentityProviderZoneTypeSAML, IdentityProviderZoneTypeCentrify, IdentityProviderZoneTypeFacebook, IdentityProviderZoneTypeGitHub, IdentityProviderZoneTypeGoogleApps, IdentityProviderZoneTypeGoogle, IdentityProviderZoneTypeLinkedin, IdentityProviderZoneTypeOidc, IdentityProviderZoneTypeOkta, IdentityProviderZoneTypeOnelogin, IdentityProviderZoneTypePingone, IdentityProviderZoneTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderZoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// Indicates how a SCIM event updates a user identity used for policy evaluation.
	// Use "automatic" to automatically update a user's identity and augment it with
	// fields from the SCIM user resource. Use "reauth" to force re-authentication on
	// group membership updates, user identity update will only occur after successful
	// re-authentication. With "reauth" identities will not contain fields from the
	// SCIM user resource. With "no_action" identities will not be changed by SCIM
	// updates in any way and users will not be prompted to reauthenticate.
	IdentityUpdateBehavior IdentityProviderZoneScimConfigIdentityUpdateBehavior `json:"identity_update_behavior"`
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
	UserDeprovision bool                               `json:"user_deprovision"`
	JSON            identityProviderZoneScimConfigJSON `json:"-"`
}

// identityProviderZoneScimConfigJSON contains the JSON metadata for the struct
// [IdentityProviderZoneScimConfig]
type identityProviderZoneScimConfigJSON struct {
	Enabled                apijson.Field
	IdentityUpdateBehavior apijson.Field
	ScimBaseURL            apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderZoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderZoneScimConfigJSON) RawJSON() string {
	return r.raw
}

// Indicates how a SCIM event updates a user identity used for policy evaluation.
// Use "automatic" to automatically update a user's identity and augment it with
// fields from the SCIM user resource. Use "reauth" to force re-authentication on
// group membership updates, user identity update will only occur after successful
// re-authentication. With "reauth" identities will not contain fields from the
// SCIM user resource. With "no_action" identities will not be changed by SCIM
// updates in any way and users will not be prompted to reauthenticate.
type IdentityProviderZoneScimConfigIdentityUpdateBehavior string

const (
	IdentityProviderZoneScimConfigIdentityUpdateBehaviorAutomatic IdentityProviderZoneScimConfigIdentityUpdateBehavior = "automatic"
	IdentityProviderZoneScimConfigIdentityUpdateBehaviorReauth    IdentityProviderZoneScimConfigIdentityUpdateBehavior = "reauth"
	IdentityProviderZoneScimConfigIdentityUpdateBehaviorNoAction  IdentityProviderZoneScimConfigIdentityUpdateBehavior = "no_action"
)

func (r IdentityProviderZoneScimConfigIdentityUpdateBehavior) IsKnown() bool {
	switch r {
	case IdentityProviderZoneScimConfigIdentityUpdateBehaviorAutomatic, IdentityProviderZoneScimConfigIdentityUpdateBehaviorReauth, IdentityProviderZoneScimConfigIdentityUpdateBehaviorNoAction:
		return true
	}
	return false
}

type IdentityProviderZoneParam struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[interface{}] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderZoneType] `json:"type,required"`
	// UUID
	ID param.Field[string] `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderZoneScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderZoneParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderZoneScimConfigParam struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled param.Field[bool] `json:"enabled"`
	// Indicates how a SCIM event updates a user identity used for policy evaluation.
	// Use "automatic" to automatically update a user's identity and augment it with
	// fields from the SCIM user resource. Use "reauth" to force re-authentication on
	// group membership updates, user identity update will only occur after successful
	// re-authentication. With "reauth" identities will not contain fields from the
	// SCIM user resource. With "no_action" identities will not be changed by SCIM
	// updates in any way and users will not be prompted to reauthenticate.
	IdentityUpdateBehavior param.Field[IdentityProviderZoneScimConfigIdentityUpdateBehavior] `json:"identity_update_behavior"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision param.Field[bool] `json:"seat_deprovision"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision param.Field[bool] `json:"user_deprovision"`
}

func (r IdentityProviderZoneScimConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LinkedinZone struct {
	Config GenericOAuthConfig `json:"config"`
	JSON   linkedinZoneJSON   `json:"-"`
	IdentityProviderZone
}

// linkedinZoneJSON contains the JSON metadata for the struct [LinkedinZone]
type linkedinZoneJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LinkedinZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkedinZoneJSON) RawJSON() string {
	return r.raw
}

func (r LinkedinZone) implementsZoneIdentityProviders() {}

func (r LinkedinZone) implementsZoneAccessIdentityProviderListResponseResult() {}

type LinkedinZoneParam struct {
	Config param.Field[GenericOAuthConfigParam] `json:"config"`
	IdentityProviderZoneParam
}

func (r LinkedinZoneParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r LinkedinZoneParam) implementsZoneIdentityProvidersUnionParam() {}

type OidcSchema struct {
	Config OidcSchemaConfig `json:"config"`
	JSON   oidcSchemaJSON   `json:"-"`
	IdentityProviderZone
}

// oidcSchemaJSON contains the JSON metadata for the struct [OidcSchema]
type oidcSchemaJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OidcSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oidcSchemaJSON) RawJSON() string {
	return r.raw
}

func (r OidcSchema) implementsZoneIdentityProviders() {}

func (r OidcSchema) implementsZoneAccessIdentityProviderListResponseResult() {}

type OidcSchemaConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL string `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL string `json:"certs_url"`
	// List of custom claims that will be pulled from your id_token and added to your
	// signed Access JWT token.
	Claims []string `json:"claims"`
	// OAuth scopes
	Scopes []string `json:"scopes"`
	// The token_endpoint URL of your IdP
	TokenURL string               `json:"token_url"`
	JSON     oidcSchemaConfigJSON `json:"-"`
	GenericOAuthConfig
}

// oidcSchemaConfigJSON contains the JSON metadata for the struct
// [OidcSchemaConfig]
type oidcSchemaConfigJSON struct {
	AuthURL     apijson.Field
	CertsURL    apijson.Field
	Claims      apijson.Field
	Scopes      apijson.Field
	TokenURL    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OidcSchemaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oidcSchemaConfigJSON) RawJSON() string {
	return r.raw
}

type OidcSchemaParam struct {
	Config param.Field[OidcSchemaConfigParam] `json:"config"`
	IdentityProviderZoneParam
}

func (r OidcSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OidcSchemaParam) implementsZoneIdentityProvidersUnionParam() {}

type OidcSchemaConfigParam struct {
	// The authorization_endpoint URL of your IdP
	AuthURL param.Field[string] `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL param.Field[string] `json:"certs_url"`
	// List of custom claims that will be pulled from your id_token and added to your
	// signed Access JWT token.
	Claims param.Field[[]string] `json:"claims"`
	// OAuth scopes
	Scopes param.Field[[]string] `json:"scopes"`
	// The token_endpoint URL of your IdP
	TokenURL param.Field[string] `json:"token_url"`
	GenericOAuthConfigParam
}

func (r OidcSchemaConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OktaSchemas struct {
	Config OktaSchemasConfig `json:"config"`
	JSON   oktaSchemasJSON   `json:"-"`
	IdentityProviderZone
}

// oktaSchemasJSON contains the JSON metadata for the struct [OktaSchemas]
type oktaSchemasJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OktaSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oktaSchemasJSON) RawJSON() string {
	return r.raw
}

func (r OktaSchemas) implementsZoneIdentityProviders() {}

func (r OktaSchemas) implementsZoneAccessIdentityProviderListResponseResult() {}

type OktaSchemasConfig struct {
	// Your okta account url
	OktaAccount string                `json:"okta_account"`
	JSON        oktaSchemasConfigJSON `json:"-"`
	GenericOAuthConfig
}

// oktaSchemasConfigJSON contains the JSON metadata for the struct
// [OktaSchemasConfig]
type oktaSchemasConfigJSON struct {
	OktaAccount apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OktaSchemasConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oktaSchemasConfigJSON) RawJSON() string {
	return r.raw
}

type OktaSchemasParam struct {
	Config param.Field[OktaSchemasConfigParam] `json:"config"`
	IdentityProviderZoneParam
}

func (r OktaSchemasParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OktaSchemasParam) implementsZoneIdentityProvidersUnionParam() {}

type OktaSchemasConfigParam struct {
	// Your okta account url
	OktaAccount param.Field[string] `json:"okta_account"`
	GenericOAuthConfigParam
}

func (r OktaSchemasConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OneloginZone struct {
	Config OneloginZoneConfig `json:"config"`
	JSON   oneloginZoneJSON   `json:"-"`
	IdentityProviderZone
}

// oneloginZoneJSON contains the JSON metadata for the struct [OneloginZone]
type oneloginZoneJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OneloginZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oneloginZoneJSON) RawJSON() string {
	return r.raw
}

func (r OneloginZone) implementsZoneIdentityProviders() {}

func (r OneloginZone) implementsZoneAccessIdentityProviderListResponseResult() {}

type OneloginZoneConfig struct {
	// Your OneLogin account url
	OneloginAccount string                 `json:"onelogin_account"`
	JSON            oneloginZoneConfigJSON `json:"-"`
	GenericOAuthConfig
}

// oneloginZoneConfigJSON contains the JSON metadata for the struct
// [OneloginZoneConfig]
type oneloginZoneConfigJSON struct {
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *OneloginZoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oneloginZoneConfigJSON) RawJSON() string {
	return r.raw
}

type OneloginZoneParam struct {
	Config param.Field[OneloginZoneConfigParam] `json:"config"`
	IdentityProviderZoneParam
}

func (r OneloginZoneParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OneloginZoneParam) implementsZoneIdentityProvidersUnionParam() {}

type OneloginZoneConfigParam struct {
	// Your OneLogin account url
	OneloginAccount param.Field[string] `json:"onelogin_account"`
	GenericOAuthConfigParam
}

func (r OneloginZoneConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PingoneZone struct {
	Config PingoneZoneConfig `json:"config"`
	JSON   pingoneZoneJSON   `json:"-"`
	IdentityProviderZone
}

// pingoneZoneJSON contains the JSON metadata for the struct [PingoneZone]
type pingoneZoneJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PingoneZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pingoneZoneJSON) RawJSON() string {
	return r.raw
}

func (r PingoneZone) implementsZoneIdentityProviders() {}

func (r PingoneZone) implementsZoneAccessIdentityProviderListResponseResult() {}

type PingoneZoneConfig struct {
	// Your PingOne environment identifier
	PingEnvID string                `json:"ping_env_id"`
	JSON      pingoneZoneConfigJSON `json:"-"`
	GenericOAuthConfig
}

// pingoneZoneConfigJSON contains the JSON metadata for the struct
// [PingoneZoneConfig]
type pingoneZoneConfigJSON struct {
	PingEnvID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PingoneZoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pingoneZoneConfigJSON) RawJSON() string {
	return r.raw
}

type PingoneZoneParam struct {
	Config param.Field[PingoneZoneConfigParam] `json:"config"`
	IdentityProviderZoneParam
}

func (r PingoneZoneParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PingoneZoneParam) implementsZoneIdentityProvidersUnionParam() {}

type PingoneZoneConfigParam struct {
	// Your PingOne environment identifier
	PingEnvID param.Field[string] `json:"ping_env_id"`
	GenericOAuthConfigParam
}

func (r PingoneZoneConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SAMLAuthentication struct {
	Config SAMLAuthenticationConfig `json:"config"`
	JSON   samlAuthenticationJSON   `json:"-"`
	IdentityProviderZone
}

// samlAuthenticationJSON contains the JSON metadata for the struct
// [SAMLAuthentication]
type samlAuthenticationJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SAMLAuthentication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r samlAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r SAMLAuthentication) implementsZoneIdentityProviders() {}

func (r SAMLAuthentication) implementsZoneAccessIdentityProviderListResponseResult() {}

type SAMLAuthenticationConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []SAMLAuthenticationConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                       `json:"sso_target_url"`
	JSON         samlAuthenticationConfigJSON `json:"-"`
}

// samlAuthenticationConfigJSON contains the JSON metadata for the struct
// [SAMLAuthenticationConfig]
type samlAuthenticationConfigJSON struct {
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

func (r *SAMLAuthenticationConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r samlAuthenticationConfigJSON) RawJSON() string {
	return r.raw
}

type SAMLAuthenticationConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                      `json:"header_name"`
	JSON       samlAuthenticationConfigHeaderAttributeJSON `json:"-"`
}

// samlAuthenticationConfigHeaderAttributeJSON contains the JSON metadata for the
// struct [SAMLAuthenticationConfigHeaderAttribute]
type samlAuthenticationConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SAMLAuthenticationConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r samlAuthenticationConfigHeaderAttributeJSON) RawJSON() string {
	return r.raw
}

type SAMLAuthenticationParam struct {
	Config param.Field[SAMLAuthenticationConfigParam] `json:"config"`
	IdentityProviderZoneParam
}

func (r SAMLAuthenticationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SAMLAuthenticationParam) implementsZoneIdentityProvidersUnionParam() {}

type SAMLAuthenticationConfigParam struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]SAMLAuthenticationConfigHeaderAttributeParam] `json:"header_attributes"`
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

func (r SAMLAuthenticationConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SAMLAuthenticationConfigHeaderAttributeParam struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r SAMLAuthenticationConfigHeaderAttributeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SingleResponseProviderZone struct {
	Result ZoneIdentityProviders          `json:"result"`
	JSON   singleResponseProviderZoneJSON `json:"-"`
	APIResponseSingleAccess
}

// singleResponseProviderZoneJSON contains the JSON metadata for the struct
// [SingleResponseProviderZone]
type singleResponseProviderZoneJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseProviderZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseProviderZoneJSON) RawJSON() string {
	return r.raw
}

type YandexSchema struct {
	Config GenericOAuthConfig `json:"config"`
	JSON   yandexSchemaJSON   `json:"-"`
	IdentityProviderZone
}

// yandexSchemaJSON contains the JSON metadata for the struct [YandexSchema]
type yandexSchemaJSON struct {
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *YandexSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r yandexSchemaJSON) RawJSON() string {
	return r.raw
}

func (r YandexSchema) implementsZoneIdentityProviders() {}

func (r YandexSchema) implementsZoneAccessIdentityProviderListResponseResult() {}

type YandexSchemaParam struct {
	Config param.Field[GenericOAuthConfigParam] `json:"config"`
	IdentityProviderZoneParam
}

func (r YandexSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r YandexSchemaParam) implementsZoneIdentityProvidersUnionParam() {}

type ZoneIdentityProviders struct {
	// UUID
	ID string `json:"id"`
	// This field can have the runtime type of [interface{}].
	Config interface{} `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// This field can have the runtime type of [IdentityProviderZoneScimConfig].
	ScimConfig interface{} `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type  ZoneIdentityProvidersType `json:"type"`
	JSON  zoneIdentityProvidersJSON `json:"-"`
	union ZoneIdentityProvidersUnion
}

// zoneIdentityProvidersJSON contains the JSON metadata for the struct
// [ZoneIdentityProviders]
type zoneIdentityProvidersJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r zoneIdentityProvidersJSON) RawJSON() string {
	return r.raw
}

func (r *ZoneIdentityProviders) UnmarshalJSON(data []byte) (err error) {
	*r = ZoneIdentityProviders{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZoneIdentityProvidersUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [AzureAd], [CentrifySchemaProvider],
// [FacebookZone], [GitHubAccessSchema], [GoogleSchema],
// [GoogleAppsSchemaIdentityProvider], [LinkedinZone], [OidcSchema], [OktaSchemas],
// [OneloginZone], [PingoneZone], [SAMLAuthentication], [YandexSchema].
func (r ZoneIdentityProviders) AsUnion() ZoneIdentityProvidersUnion {
	return r.union
}

// Union satisfied by [AzureAd], [CentrifySchemaProvider], [FacebookZone],
// [GitHubAccessSchema], [GoogleSchema], [GoogleAppsSchemaIdentityProvider],
// [LinkedinZone], [OidcSchema], [OktaSchemas], [OneloginZone], [PingoneZone],
// [SAMLAuthentication] or [YandexSchema].
type ZoneIdentityProvidersUnion interface {
	implementsZoneIdentityProviders()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneIdentityProvidersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AzureAd{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CentrifySchemaProvider{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FacebookZone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GitHubAccessSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GoogleSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GoogleAppsSchemaIdentityProvider{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LinkedinZone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OidcSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OktaSchemas{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OneloginZone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PingoneZone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SAMLAuthentication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(YandexSchema{}),
		},
	)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneIdentityProvidersType string

const (
	ZoneIdentityProvidersTypeOnetimepin ZoneIdentityProvidersType = "onetimepin"
	ZoneIdentityProvidersTypeAzureAd    ZoneIdentityProvidersType = "azureAD"
	ZoneIdentityProvidersTypeSAML       ZoneIdentityProvidersType = "saml"
	ZoneIdentityProvidersTypeCentrify   ZoneIdentityProvidersType = "centrify"
	ZoneIdentityProvidersTypeFacebook   ZoneIdentityProvidersType = "facebook"
	ZoneIdentityProvidersTypeGitHub     ZoneIdentityProvidersType = "github"
	ZoneIdentityProvidersTypeGoogleApps ZoneIdentityProvidersType = "google-apps"
	ZoneIdentityProvidersTypeGoogle     ZoneIdentityProvidersType = "google"
	ZoneIdentityProvidersTypeLinkedin   ZoneIdentityProvidersType = "linkedin"
	ZoneIdentityProvidersTypeOidc       ZoneIdentityProvidersType = "oidc"
	ZoneIdentityProvidersTypeOkta       ZoneIdentityProvidersType = "okta"
	ZoneIdentityProvidersTypeOnelogin   ZoneIdentityProvidersType = "onelogin"
	ZoneIdentityProvidersTypePingone    ZoneIdentityProvidersType = "pingone"
	ZoneIdentityProvidersTypeYandex     ZoneIdentityProvidersType = "yandex"
)

func (r ZoneIdentityProvidersType) IsKnown() bool {
	switch r {
	case ZoneIdentityProvidersTypeOnetimepin, ZoneIdentityProvidersTypeAzureAd, ZoneIdentityProvidersTypeSAML, ZoneIdentityProvidersTypeCentrify, ZoneIdentityProvidersTypeFacebook, ZoneIdentityProvidersTypeGitHub, ZoneIdentityProvidersTypeGoogleApps, ZoneIdentityProvidersTypeGoogle, ZoneIdentityProvidersTypeLinkedin, ZoneIdentityProvidersTypeOidc, ZoneIdentityProvidersTypeOkta, ZoneIdentityProvidersTypeOnelogin, ZoneIdentityProvidersTypePingone, ZoneIdentityProvidersTypeYandex:
		return true
	}
	return false
}

// Satisfied by [AzureAdParam], [CentrifySchemaProviderParam], [FacebookZoneParam],
// [GitHubAccessSchemaParam], [GoogleSchemaParam],
// [GoogleAppsSchemaIdentityProviderParam], [LinkedinZoneParam], [OidcSchemaParam],
// [OktaSchemasParam], [OneloginZoneParam], [PingoneZoneParam],
// [SAMLAuthenticationParam], [YandexSchemaParam].
type ZoneIdentityProvidersUnionParam interface {
	implementsZoneIdentityProvidersUnionParam()
}

type ZoneAccessIdentityProviderListResponse struct {
	Result []ZoneAccessIdentityProviderListResponseResult `json:"result"`
	JSON   zoneAccessIdentityProviderListResponseJSON     `json:"-"`
	APIResponseCollectionAccess
}

// zoneAccessIdentityProviderListResponseJSON contains the JSON metadata for the
// struct [ZoneAccessIdentityProviderListResponse]
type zoneAccessIdentityProviderListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAccessIdentityProviderListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAccessIdentityProviderListResponseResult struct {
	// UUID
	ID string `json:"id"`
	// This field can have the runtime type of [interface{}].
	Config interface{} `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// This field can have the runtime type of [IdentityProviderZoneScimConfig].
	ScimConfig interface{} `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type  ZoneAccessIdentityProviderListResponseResultType `json:"type"`
	JSON  zoneAccessIdentityProviderListResponseResultJSON `json:"-"`
	union ZoneAccessIdentityProviderListResponseResultUnion
}

// zoneAccessIdentityProviderListResponseResultJSON contains the JSON metadata for
// the struct [ZoneAccessIdentityProviderListResponseResult]
type zoneAccessIdentityProviderListResponseResultJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r zoneAccessIdentityProviderListResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *ZoneAccessIdentityProviderListResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = ZoneAccessIdentityProviderListResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZoneAccessIdentityProviderListResponseResultUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [AzureAd], [CentrifySchemaProvider],
// [FacebookZone], [GitHubAccessSchema], [GoogleSchema],
// [GoogleAppsSchemaIdentityProvider], [LinkedinZone], [OidcSchema], [OktaSchemas],
// [OneloginZone], [PingoneZone], [SAMLAuthentication], [YandexSchema],
// [ZoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepin].
func (r ZoneAccessIdentityProviderListResponseResult) AsUnion() ZoneAccessIdentityProviderListResponseResultUnion {
	return r.union
}

// Union satisfied by [AzureAd], [CentrifySchemaProvider], [FacebookZone],
// [GitHubAccessSchema], [GoogleSchema], [GoogleAppsSchemaIdentityProvider],
// [LinkedinZone], [OidcSchema], [OktaSchemas], [OneloginZone], [PingoneZone],
// [SAMLAuthentication], [YandexSchema] or
// [ZoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepin].
type ZoneAccessIdentityProviderListResponseResultUnion interface {
	implementsZoneAccessIdentityProviderListResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAccessIdentityProviderListResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AzureAd{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CentrifySchemaProvider{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FacebookZone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GitHubAccessSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GoogleSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GoogleAppsSchemaIdentityProvider{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LinkedinZone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OidcSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OktaSchemas{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OneloginZone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PingoneZone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SAMLAuthentication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(YandexSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepin{}),
		},
	)
}

type ZoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepin struct {
	Config ZoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepinConfig `json:"config"`
	Type   ZoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepinType   `json:"type"`
	JSON   zoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepinJSON   `json:"-"`
	IdentityProviderZone
}

// zoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepinJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepin]
type zoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepinJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepinJSON) RawJSON() string {
	return r.raw
}

func (r ZoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepin) implementsZoneAccessIdentityProviderListResponseResult() {
}

type ZoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepinConfig struct {
	RedirectURL string                                                                        `json:"redirect_url"`
	JSON        zoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepinConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepinConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepinConfig]
type zoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepinConfigJSON struct {
	RedirectURL apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepinConfigJSON) RawJSON() string {
	return r.raw
}

type ZoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepinType string

const (
	ZoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepinTypeOnetimepin ZoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepinType = "onetimepin"
)

func (r ZoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepinType) IsKnown() bool {
	switch r {
	case ZoneAccessIdentityProviderListResponseResultAccessSchemasOnetimepinTypeOnetimepin:
		return true
	}
	return false
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderListResponseResultType string

const (
	ZoneAccessIdentityProviderListResponseResultTypeOnetimepin ZoneAccessIdentityProviderListResponseResultType = "onetimepin"
	ZoneAccessIdentityProviderListResponseResultTypeAzureAd    ZoneAccessIdentityProviderListResponseResultType = "azureAD"
	ZoneAccessIdentityProviderListResponseResultTypeSAML       ZoneAccessIdentityProviderListResponseResultType = "saml"
	ZoneAccessIdentityProviderListResponseResultTypeCentrify   ZoneAccessIdentityProviderListResponseResultType = "centrify"
	ZoneAccessIdentityProviderListResponseResultTypeFacebook   ZoneAccessIdentityProviderListResponseResultType = "facebook"
	ZoneAccessIdentityProviderListResponseResultTypeGitHub     ZoneAccessIdentityProviderListResponseResultType = "github"
	ZoneAccessIdentityProviderListResponseResultTypeGoogleApps ZoneAccessIdentityProviderListResponseResultType = "google-apps"
	ZoneAccessIdentityProviderListResponseResultTypeGoogle     ZoneAccessIdentityProviderListResponseResultType = "google"
	ZoneAccessIdentityProviderListResponseResultTypeLinkedin   ZoneAccessIdentityProviderListResponseResultType = "linkedin"
	ZoneAccessIdentityProviderListResponseResultTypeOidc       ZoneAccessIdentityProviderListResponseResultType = "oidc"
	ZoneAccessIdentityProviderListResponseResultTypeOkta       ZoneAccessIdentityProviderListResponseResultType = "okta"
	ZoneAccessIdentityProviderListResponseResultTypeOnelogin   ZoneAccessIdentityProviderListResponseResultType = "onelogin"
	ZoneAccessIdentityProviderListResponseResultTypePingone    ZoneAccessIdentityProviderListResponseResultType = "pingone"
	ZoneAccessIdentityProviderListResponseResultTypeYandex     ZoneAccessIdentityProviderListResponseResultType = "yandex"
)

func (r ZoneAccessIdentityProviderListResponseResultType) IsKnown() bool {
	switch r {
	case ZoneAccessIdentityProviderListResponseResultTypeOnetimepin, ZoneAccessIdentityProviderListResponseResultTypeAzureAd, ZoneAccessIdentityProviderListResponseResultTypeSAML, ZoneAccessIdentityProviderListResponseResultTypeCentrify, ZoneAccessIdentityProviderListResponseResultTypeFacebook, ZoneAccessIdentityProviderListResponseResultTypeGitHub, ZoneAccessIdentityProviderListResponseResultTypeGoogleApps, ZoneAccessIdentityProviderListResponseResultTypeGoogle, ZoneAccessIdentityProviderListResponseResultTypeLinkedin, ZoneAccessIdentityProviderListResponseResultTypeOidc, ZoneAccessIdentityProviderListResponseResultTypeOkta, ZoneAccessIdentityProviderListResponseResultTypeOnelogin, ZoneAccessIdentityProviderListResponseResultTypePingone, ZoneAccessIdentityProviderListResponseResultTypeYandex:
		return true
	}
	return false
}

type ZoneAccessIdentityProviderNewParams struct {
	ZoneIdentityProviders ZoneIdentityProvidersUnionParam `json:"zone_identity_providers,required"`
}

func (r ZoneAccessIdentityProviderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ZoneIdentityProviders)
}

type ZoneAccessIdentityProviderUpdateParams struct {
	ZoneIdentityProviders ZoneIdentityProvidersUnionParam `json:"zone_identity_providers,required"`
}

func (r ZoneAccessIdentityProviderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ZoneIdentityProviders)
}
