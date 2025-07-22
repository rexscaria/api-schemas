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

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/tidwall/gjson"
)

// AccountAccessAppService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessAppService] method instead.
type AccountAccessAppService struct {
	Options  []option.RequestOption
	Ca       *AccountAccessAppCaService
	Policies *AccountAccessAppPolicyService
}

// NewAccountAccessAppService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessAppService(opts ...option.RequestOption) (r *AccountAccessAppService) {
	r = &AccountAccessAppService{}
	r.Options = opts
	r.Ca = NewAccountAccessAppCaService(opts...)
	r.Policies = NewAccountAccessAppPolicyService(opts...)
	return
}

// Adds a new application to Access.
func (r *AccountAccessAppService) New(ctx context.Context, accountID string, body AccountAccessAppNewParams, opts ...option.RequestOption) (res *AccountAccessAppNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/apps", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches information about an Access application.
func (r *AccountAccessAppService) Get(ctx context.Context, accountID string, appID AppIDParam, opts ...option.RequestOption) (res *SingleResponseApp, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/apps/%s", accountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an Access application.
func (r *AccountAccessAppService) Update(ctx context.Context, accountID string, appID AppIDParam, body AccountAccessAppUpdateParams, opts ...option.RequestOption) (res *AccountAccessAppUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/apps/%s", accountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all Access applications in an account.
func (r *AccountAccessAppService) List(ctx context.Context, accountID string, query AccountAccessAppListParams, opts ...option.RequestOption) (res *AccountAccessAppListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/apps", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an application from Access.
func (r *AccountAccessAppService) Delete(ctx context.Context, accountID string, appID AppIDParam, opts ...option.RequestOption) (res *IDResponseApps, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/apps/%s", accountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Revokes all tokens issued for an application.
func (r *AccountAccessAppService) RevokeTokens(ctx context.Context, accountID string, appID AppIDParam, opts ...option.RequestOption) (res *SchemasEmptyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/apps/%s/revoke_tokens", accountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Tests if a specific user has permission to access an application.
func (r *AccountAccessAppService) UserPolicyChecks(ctx context.Context, accountID string, appID AppIDParam, opts ...option.RequestOption) (res *AccountAccessAppUserPolicyChecksResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/apps/%s/user_policy_checks", accountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AllowedMethodsItem string

const (
	AllowedMethodsItemGet     AllowedMethodsItem = "GET"
	AllowedMethodsItemPost    AllowedMethodsItem = "POST"
	AllowedMethodsItemHead    AllowedMethodsItem = "HEAD"
	AllowedMethodsItemPut     AllowedMethodsItem = "PUT"
	AllowedMethodsItemDelete  AllowedMethodsItem = "DELETE"
	AllowedMethodsItemConnect AllowedMethodsItem = "CONNECT"
	AllowedMethodsItemOptions AllowedMethodsItem = "OPTIONS"
	AllowedMethodsItemTrace   AllowedMethodsItem = "TRACE"
	AllowedMethodsItemPatch   AllowedMethodsItem = "PATCH"
)

func (r AllowedMethodsItem) IsKnown() bool {
	switch r {
	case AllowedMethodsItemGet, AllowedMethodsItemPost, AllowedMethodsItemHead, AllowedMethodsItemPut, AllowedMethodsItemDelete, AllowedMethodsItemConnect, AllowedMethodsItemOptions, AllowedMethodsItemTrace, AllowedMethodsItemPatch:
		return true
	}
	return false
}

type APIResponseCollectionAccess struct {
	ResultInfo APIResponseCollectionAccessResultInfo `json:"result_info"`
	JSON       apiResponseCollectionAccessJSON       `json:"-"`
	APIResponseAccess
}

// apiResponseCollectionAccessJSON contains the JSON metadata for the struct
// [APIResponseCollectionAccess]
type apiResponseCollectionAccessJSON struct {
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionAccessJSON) RawJSON() string {
	return r.raw
}

type APIResponseCollectionAccessResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       apiResponseCollectionAccessResultInfoJSON `json:"-"`
}

// apiResponseCollectionAccessResultInfoJSON contains the JSON metadata for the
// struct [APIResponseCollectionAccessResultInfo]
type apiResponseCollectionAccessResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionAccessResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionAccessResultInfoJSON) RawJSON() string {
	return r.raw
}

type APIResponseSingleAccess struct {
	Errors   []MessagesAccessItem `json:"errors,required"`
	Messages []MessagesAccessItem `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseSingleAccessSuccess `json:"success,required"`
	JSON    apiResponseSingleAccessJSON    `json:"-"`
}

// apiResponseSingleAccessJSON contains the JSON metadata for the struct
// [APIResponseSingleAccess]
type apiResponseSingleAccessJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleAccessJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseSingleAccessSuccess bool

const (
	APIResponseSingleAccessSuccessTrue APIResponseSingleAccessSuccess = true
)

func (r APIResponseSingleAccessSuccess) IsKnown() bool {
	switch r {
	case APIResponseSingleAccessSuccessTrue:
		return true
	}
	return false
}

type AppIDParam = string

// Contains the targets secured by the application.
//
// Satisfied by [AppRequestSelfHostedApplicationParam],
// [AppRequestSaaSApplicationParam], [AppRequestBrowserSSHApplicationParam],
// [AppRequestBrowserVncApplicationParam], [AppRequestAppLauncherApplicationParam],
// [AppRequestDeviceEnrollmentPermissionsApplicationParam],
// [AppRequestBrowserIsolationPermissionsApplicationParam],
// [AppRequestBookmarkApplicationParam],
// [AppRequestInfrastructureApplicationParam],
// [AppRequestBrowserRdpApplicationParam].
type AppRequestUnionParam interface {
	implementsAppRequestUnionParam()
}

type AppRequestSelfHostedApplicationParam struct {
	SelfHostedPropsAccountParam
	EmbeddedPoliciesParam
	EmbeddedScimConfigParam
}

func (r AppRequestSelfHostedApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppRequestSelfHostedApplicationParam) implementsAppRequestUnionParam() {}

type AppRequestSaaSApplicationParam struct {
	SaasPropsParam
	EmbeddedPoliciesParam
	EmbeddedScimConfigParam
}

func (r AppRequestSaaSApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppRequestSaaSApplicationParam) implementsAppRequestUnionParam() {}

type AppRequestBrowserSSHApplicationParam struct {
	SSHPropsParam
	EmbeddedPoliciesParam
	EmbeddedScimConfigParam
}

func (r AppRequestBrowserSSHApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppRequestBrowserSSHApplicationParam) implementsAppRequestUnionParam() {}

type AppRequestBrowserVncApplicationParam struct {
	VncPropsParam
	EmbeddedPoliciesParam
	EmbeddedScimConfigParam
}

func (r AppRequestBrowserVncApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppRequestBrowserVncApplicationParam) implementsAppRequestUnionParam() {}

type AppRequestAppLauncherApplicationParam struct {
	LauncherPropsParam
	EmbeddedPoliciesParam
	EmbeddedScimConfigParam
}

func (r AppRequestAppLauncherApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppRequestAppLauncherApplicationParam) implementsAppRequestUnionParam() {}

type AppRequestDeviceEnrollmentPermissionsApplicationParam struct {
	WarpPropsParam
	EmbeddedPoliciesParam
	EmbeddedScimConfigParam
}

func (r AppRequestDeviceEnrollmentPermissionsApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppRequestDeviceEnrollmentPermissionsApplicationParam) implementsAppRequestUnionParam() {}

type AppRequestBrowserIsolationPermissionsApplicationParam struct {
	BisoPropsParam
	EmbeddedPoliciesParam
	EmbeddedScimConfigParam
}

func (r AppRequestBrowserIsolationPermissionsApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppRequestBrowserIsolationPermissionsApplicationParam) implementsAppRequestUnionParam() {}

type AppRequestBookmarkApplicationParam struct {
	BookmarkPropsParam
	EmbeddedScimConfigParam
}

func (r AppRequestBookmarkApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppRequestBookmarkApplicationParam) implementsAppRequestUnionParam() {}

// Contains the targets secured by the application.
type AppRequestInfrastructureApplicationParam struct {
	// The policies that Access applies to the application.
	Policies param.Field[[]AppRequestInfrastructureApplicationPolicyParam] `json:"policies"`
	InfraPropsParam
}

func (r AppRequestInfrastructureApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppRequestInfrastructureApplicationParam) implementsAppRequestUnionParam() {}

type AppRequestInfrastructureApplicationPolicyParam struct {
	// The rules that define how users may connect to the targets secured by your
	// application.
	ConnectionRules param.Field[ConnectionRulesParam] `json:"connection_rules"`
	BasePolicyRequestParam
}

func (r AppRequestInfrastructureApplicationPolicyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Contains the targets secured by the application.
type AppRequestBrowserRdpApplicationParam struct {
	RdpPropsParam
	EmbeddedPoliciesParam
	EmbeddedScimConfigParam
}

func (r AppRequestBrowserRdpApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppRequestBrowserRdpApplicationParam) implementsAppRequestUnionParam() {}

// The policies that Access applies to the application.
type AppResponse struct {
	// UUID
	ID string `json:"id"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWarp bool `json:"allow_authenticate_via_warp"`
	// This field can have the runtime type of [[]string].
	AllowedIdps interface{} `json:"allowed_idps"`
	// The image URL of the logo shown in the App Launcher header.
	AppLauncherLogoURL string `json:"app_launcher_logo_url"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// Audience tag.
	Aud string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool `json:"auto_redirect_to_identity"`
	// The background color of the App Launcher page.
	BgColor string `json:"bg_color"`
	// This field can have the runtime type of [SelfHostedPropsAccountCorsHeaders].
	CorsHeaders interface{} `json:"cors_headers"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage string `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing identity-based rules.
	CustomDenyURL string `json:"custom_deny_url"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing non-identity rules.
	CustomNonIdentityDenyURL string `json:"custom_non_identity_deny_url"`
	// This field can have the runtime type of [[]string].
	CustomPages interface{} `json:"custom_pages"`
	// This field can have the runtime type of [[]SelfHostedPropsAccountDestination].
	Destinations interface{} `json:"destinations"`
	// This field can have the runtime type of [string].
	Domain interface{} `json:"domain"`
	// Enables the binding cookie, which increases security against compromised
	// authorization tokens and CSRF attacks.
	EnableBindingCookie bool `json:"enable_binding_cookie"`
	// This field can have the runtime type of [[]FeatureAppPropsAccountFooterLink].
	FooterLinks interface{} `json:"footer_links"`
	// The background color of the App Launcher header.
	HeaderBgColor string `json:"header_bg_color"`
	// Enables the HttpOnly cookie attribute, which increases security against XSS
	// attacks.
	HTTPOnlyCookieAttribute bool `json:"http_only_cookie_attribute"`
	// This field can have the runtime type of
	// [FeatureAppPropsAccountLandingPageDesign].
	LandingPageDesign interface{} `json:"landing_page_design"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// This field can have the runtime type of [string].
	Name interface{} `json:"name"`
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool `json:"path_cookie_attribute"`
	// This field can have the runtime type of [[]PolicyResponseApp],
	// [[]AppResponseInfrastructureApplicationPolicy].
	Policies interface{} `json:"policies"`
	// This field can have the runtime type of [SaasPropsSaasApp].
	SaasApp interface{} `json:"saas_app"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	ScimConfig ScimConfig `json:"scim_config"`
	// This field can have the runtime type of [[]string].
	SelfHostedDomains interface{} `json:"self_hosted_domains"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect bool `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration string `json:"session_duration"`
	// Determines when to skip the App Launcher landing page.
	SkipAppLauncherLoginPage bool `json:"skip_app_launcher_login_page"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial bool `json:"skip_interstitial"`
	// This field can have the runtime type of [[]string].
	Tags interface{} `json:"tags"`
	// This field can have the runtime type of
	// [[]EmbeddedTargetCriteriaTargetCriterion].
	TargetCriteria interface{} `json:"target_criteria"`
	// The application type.
	Type      string          `json:"type"`
	UpdatedAt time.Time       `json:"updated_at" format:"date-time"`
	JSON      appResponseJSON `json:"-"`
	union     AppResponseUnion
}

// appResponseJSON contains the JSON metadata for the struct [AppResponse]
type appResponseJSON struct {
	ID                       apijson.Field
	AllowAuthenticateViaWarp apijson.Field
	AllowedIdps              apijson.Field
	AppLauncherLogoURL       apijson.Field
	AppLauncherVisible       apijson.Field
	Aud                      apijson.Field
	AutoRedirectToIdentity   apijson.Field
	BgColor                  apijson.Field
	CorsHeaders              apijson.Field
	CreatedAt                apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	Destinations             apijson.Field
	Domain                   apijson.Field
	EnableBindingCookie      apijson.Field
	FooterLinks              apijson.Field
	HeaderBgColor            apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LandingPageDesign        apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	Policies                 apijson.Field
	SaasApp                  apijson.Field
	SameSiteCookieAttribute  apijson.Field
	ScimConfig               apijson.Field
	SelfHostedDomains        apijson.Field
	ServiceAuth401Redirect   apijson.Field
	SessionDuration          apijson.Field
	SkipAppLauncherLoginPage apijson.Field
	SkipInterstitial         apijson.Field
	Tags                     apijson.Field
	TargetCriteria           apijson.Field
	Type                     apijson.Field
	UpdatedAt                apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r appResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AppResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AppResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AppResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [AppResponseSelfHostedApplication],
// [AppResponseSaaSApplication], [AppResponseBrowserSSHApplication],
// [AppResponseBrowserVncApplication], [AppResponseAppLauncherApplication],
// [AppResponseDeviceEnrollmentPermissionsApplication],
// [AppResponseBrowserIsolationPermissionsApplication],
// [AppResponseBookmarkApplication], [AppResponseInfrastructureApplication],
// [AppResponseBrowserRdpApplication].
func (r AppResponse) AsUnion() AppResponseUnion {
	return r.union
}

// The policies that Access applies to the application.
//
// Union satisfied by [AppResponseSelfHostedApplication],
// [AppResponseSaaSApplication], [AppResponseBrowserSSHApplication],
// [AppResponseBrowserVncApplication], [AppResponseAppLauncherApplication],
// [AppResponseDeviceEnrollmentPermissionsApplication],
// [AppResponseBrowserIsolationPermissionsApplication],
// [AppResponseBookmarkApplication], [AppResponseInfrastructureApplication] or
// [AppResponseBrowserRdpApplication].
type AppResponseUnion interface {
	implementsAppResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppResponseSelfHostedApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppResponseSaaSApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppResponseBrowserSSHApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppResponseBrowserVncApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppResponseAppLauncherApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppResponseDeviceEnrollmentPermissionsApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppResponseBrowserIsolationPermissionsApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppResponseBookmarkApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppResponseInfrastructureApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppResponseBrowserRdpApplication{}),
		},
	)
}

// The policies that Access applies to the application.
type AppResponseSelfHostedApplication struct {
	JSON appResponseSelfHostedApplicationJSON `json:"-"`
	BasicResponseProps
	SelfHostedPropsAccount
	EmbeddedResponsePolicies
}

// appResponseSelfHostedApplicationJSON contains the JSON metadata for the struct
// [AppResponseSelfHostedApplication]
type appResponseSelfHostedApplicationJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppResponseSelfHostedApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appResponseSelfHostedApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AppResponseSelfHostedApplication) implementsAppResponse() {}

// The policies that Access applies to the application.
type AppResponseSaaSApplication struct {
	JSON appResponseSaaSApplicationJSON `json:"-"`
	BasicResponseProps
	SaasProps
	EmbeddedResponsePolicies
}

// appResponseSaaSApplicationJSON contains the JSON metadata for the struct
// [AppResponseSaaSApplication]
type appResponseSaaSApplicationJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppResponseSaaSApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appResponseSaaSApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AppResponseSaaSApplication) implementsAppResponse() {}

// The policies that Access applies to the application.
type AppResponseBrowserSSHApplication struct {
	JSON appResponseBrowserSSHApplicationJSON `json:"-"`
	BasicResponseProps
	SSHProps
	EmbeddedResponsePolicies
}

// appResponseBrowserSSHApplicationJSON contains the JSON metadata for the struct
// [AppResponseBrowserSSHApplication]
type appResponseBrowserSSHApplicationJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppResponseBrowserSSHApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appResponseBrowserSSHApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AppResponseBrowserSSHApplication) implementsAppResponse() {}

// The policies that Access applies to the application.
type AppResponseBrowserVncApplication struct {
	JSON appResponseBrowserVncApplicationJSON `json:"-"`
	BasicResponseProps
	VncProps
	EmbeddedResponsePolicies
}

// appResponseBrowserVncApplicationJSON contains the JSON metadata for the struct
// [AppResponseBrowserVncApplication]
type appResponseBrowserVncApplicationJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppResponseBrowserVncApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appResponseBrowserVncApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AppResponseBrowserVncApplication) implementsAppResponse() {}

// The policies that Access applies to the application.
type AppResponseAppLauncherApplication struct {
	JSON appResponseAppLauncherApplicationJSON `json:"-"`
	BasicResponseProps
	LauncherProps
	EmbeddedResponsePolicies
}

// appResponseAppLauncherApplicationJSON contains the JSON metadata for the struct
// [AppResponseAppLauncherApplication]
type appResponseAppLauncherApplicationJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppResponseAppLauncherApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appResponseAppLauncherApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AppResponseAppLauncherApplication) implementsAppResponse() {}

// The policies that Access applies to the application.
type AppResponseDeviceEnrollmentPermissionsApplication struct {
	JSON appResponseDeviceEnrollmentPermissionsApplicationJSON `json:"-"`
	BasicResponseProps
	WarpProps
	EmbeddedResponsePolicies
}

// appResponseDeviceEnrollmentPermissionsApplicationJSON contains the JSON metadata
// for the struct [AppResponseDeviceEnrollmentPermissionsApplication]
type appResponseDeviceEnrollmentPermissionsApplicationJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppResponseDeviceEnrollmentPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appResponseDeviceEnrollmentPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AppResponseDeviceEnrollmentPermissionsApplication) implementsAppResponse() {}

// The policies that Access applies to the application.
type AppResponseBrowserIsolationPermissionsApplication struct {
	JSON appResponseBrowserIsolationPermissionsApplicationJSON `json:"-"`
	BasicResponseProps
	BisoProps
	EmbeddedResponsePolicies
}

// appResponseBrowserIsolationPermissionsApplicationJSON contains the JSON metadata
// for the struct [AppResponseBrowserIsolationPermissionsApplication]
type appResponseBrowserIsolationPermissionsApplicationJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppResponseBrowserIsolationPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appResponseBrowserIsolationPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AppResponseBrowserIsolationPermissionsApplication) implementsAppResponse() {}

type AppResponseBookmarkApplication struct {
	JSON appResponseBookmarkApplicationJSON `json:"-"`
	BasicResponseProps
	BookmarkProps
}

// appResponseBookmarkApplicationJSON contains the JSON metadata for the struct
// [AppResponseBookmarkApplication]
type appResponseBookmarkApplicationJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppResponseBookmarkApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appResponseBookmarkApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AppResponseBookmarkApplication) implementsAppResponse() {}

// Contains the targets secured by the application.
type AppResponseInfrastructureApplication struct {
	Policies []AppResponseInfrastructureApplicationPolicy `json:"policies"`
	JSON     appResponseInfrastructureApplicationJSON     `json:"-"`
	BasicResponseProps
	InfraProps
}

// appResponseInfrastructureApplicationJSON contains the JSON metadata for the
// struct [AppResponseInfrastructureApplication]
type appResponseInfrastructureApplicationJSON struct {
	Policies    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppResponseInfrastructureApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appResponseInfrastructureApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AppResponseInfrastructureApplication) implementsAppResponse() {}

type AppResponseInfrastructureApplicationPolicy struct {
	// The rules that define how users may connect to the targets secured by your
	// application.
	ConnectionRules ConnectionRules                                `json:"connection_rules"`
	JSON            appResponseInfrastructureApplicationPolicyJSON `json:"-"`
	BasePolicyResponse
}

// appResponseInfrastructureApplicationPolicyJSON contains the JSON metadata for
// the struct [AppResponseInfrastructureApplicationPolicy]
type appResponseInfrastructureApplicationPolicyJSON struct {
	ConnectionRules apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppResponseInfrastructureApplicationPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appResponseInfrastructureApplicationPolicyJSON) RawJSON() string {
	return r.raw
}

// Contains the targets secured by the application.
type AppResponseBrowserRdpApplication struct {
	JSON appResponseBrowserRdpApplicationJSON `json:"-"`
	BasicResponseProps
	RdpProps
	EmbeddedResponsePolicies
}

// appResponseBrowserRdpApplicationJSON contains the JSON metadata for the struct
// [AppResponseBrowserRdpApplication]
type appResponseBrowserRdpApplicationJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppResponseBrowserRdpApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appResponseBrowserRdpApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AppResponseBrowserRdpApplication) implementsAppResponse() {}

// The application type.
type ApplicationType string

const (
	ApplicationTypeSelfHosted     ApplicationType = "self_hosted"
	ApplicationTypeSaas           ApplicationType = "saas"
	ApplicationTypeSSH            ApplicationType = "ssh"
	ApplicationTypeVnc            ApplicationType = "vnc"
	ApplicationTypeAppLauncher    ApplicationType = "app_launcher"
	ApplicationTypeWarp           ApplicationType = "warp"
	ApplicationTypeBiso           ApplicationType = "biso"
	ApplicationTypeBookmark       ApplicationType = "bookmark"
	ApplicationTypeDashSSO        ApplicationType = "dash_sso"
	ApplicationTypeInfrastructure ApplicationType = "infrastructure"
	ApplicationTypeRdp            ApplicationType = "rdp"
)

func (r ApplicationType) IsKnown() bool {
	switch r {
	case ApplicationTypeSelfHosted, ApplicationTypeSaas, ApplicationTypeSSH, ApplicationTypeVnc, ApplicationTypeAppLauncher, ApplicationTypeWarp, ApplicationTypeBiso, ApplicationTypeBookmark, ApplicationTypeDashSSO, ApplicationTypeInfrastructure, ApplicationTypeRdp:
		return true
	}
	return false
}

type BasicResponseProps struct {
	// UUID
	ID string `json:"id"`
	// Audience tag.
	Aud       string    `json:"aud"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	ScimConfig ScimConfig             `json:"scim_config"`
	UpdatedAt  time.Time              `json:"updated_at" format:"date-time"`
	JSON       basicResponsePropsJSON `json:"-"`
}

// basicResponsePropsJSON contains the JSON metadata for the struct
// [BasicResponseProps]
type basicResponsePropsJSON struct {
	ID          apijson.Field
	Aud         apijson.Field
	CreatedAt   apijson.Field
	ScimConfig  apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BasicResponseProps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basicResponsePropsJSON) RawJSON() string {
	return r.raw
}

type BisoProps struct {
	Domain interface{} `json:"domain"`
	Name   interface{} `json:"name"`
	// The application type.
	Type string        `json:"type"`
	JSON bisoPropsJSON `json:"-"`
	FeatureAppPropsAccount
}

// bisoPropsJSON contains the JSON metadata for the struct [BisoProps]
type bisoPropsJSON struct {
	Domain      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BisoProps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bisoPropsJSON) RawJSON() string {
	return r.raw
}

type BisoPropsParam struct {
	// The application type.
	Type param.Field[string] `json:"type"`
	FeatureAppPropsAccountParam
}

func (r BisoPropsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BookmarkProps struct {
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// The URL or domain of the bookmark.
	Domain string `json:"domain"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name string `json:"name"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags []string `json:"tags"`
	// The application type.
	Type string            `json:"type"`
	JSON bookmarkPropsJSON `json:"-"`
}

// bookmarkPropsJSON contains the JSON metadata for the struct [BookmarkProps]
type bookmarkPropsJSON struct {
	AppLauncherVisible apijson.Field
	Domain             apijson.Field
	LogoURL            apijson.Field
	Name               apijson.Field
	Tags               apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BookmarkProps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bookmarkPropsJSON) RawJSON() string {
	return r.raw
}

type BookmarkPropsParam struct {
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// The URL or domain of the bookmark.
	Domain param.Field[string] `json:"domain"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL param.Field[string] `json:"logo_url"`
	// The name of the application.
	Name param.Field[string] `json:"name"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r BookmarkPropsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rules that define how users may connect to the targets secured by your
// application.
type ConnectionRules struct {
	// The SSH-specific rules that define how users may connect to the targets secured
	// by your application.
	SSH  ConnectionRulesSSH  `json:"ssh"`
	JSON connectionRulesJSON `json:"-"`
}

// connectionRulesJSON contains the JSON metadata for the struct [ConnectionRules]
type connectionRulesJSON struct {
	SSH         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionRules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionRulesJSON) RawJSON() string {
	return r.raw
}

// The SSH-specific rules that define how users may connect to the targets secured
// by your application.
type ConnectionRulesSSH struct {
	// Contains the Unix usernames that may be used when connecting over SSH.
	Usernames []string `json:"usernames,required"`
	// Enables using Identity Provider email alias as SSH username.
	AllowEmailAlias bool                   `json:"allow_email_alias"`
	JSON            connectionRulesSSHJSON `json:"-"`
}

// connectionRulesSSHJSON contains the JSON metadata for the struct
// [ConnectionRulesSSH]
type connectionRulesSSHJSON struct {
	Usernames       apijson.Field
	AllowEmailAlias apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConnectionRulesSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionRulesSSHJSON) RawJSON() string {
	return r.raw
}

// The rules that define how users may connect to the targets secured by your
// application.
type ConnectionRulesParam struct {
	// The SSH-specific rules that define how users may connect to the targets secured
	// by your application.
	SSH param.Field[ConnectionRulesSSHParam] `json:"ssh"`
}

func (r ConnectionRulesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The SSH-specific rules that define how users may connect to the targets secured
// by your application.
type ConnectionRulesSSHParam struct {
	// Contains the Unix usernames that may be used when connecting over SSH.
	Usernames param.Field[[]string] `json:"usernames,required"`
	// Enables using Identity Provider email alias as SSH username.
	AllowEmailAlias param.Field[bool] `json:"allow_email_alias"`
}

func (r ConnectionRulesSSHParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmbeddedPoliciesParam struct {
	// The policies that Access applies to the application, in ascending order of
	// precedence. Items can reference existing policies or create new policies
	// exclusive to the application.
	Policies param.Field[[]EmbeddedPoliciesPoliciesUnionParam] `json:"policies"`
}

func (r EmbeddedPoliciesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A JSON that links a reusable policy to an application.
//
// Satisfied by [EmbeddedPoliciesPoliciesAccessAppPolicyLinkParam],
// [shared.UnionString], [EmbeddedPoliciesPoliciesIntersectionParam].
type EmbeddedPoliciesPoliciesUnionParam interface {
	ImplementsEmbeddedPoliciesPoliciesUnionParam()
}

// A JSON that links a reusable policy to an application.
type EmbeddedPoliciesPoliciesAccessAppPolicyLinkParam struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
}

func (r EmbeddedPoliciesPoliciesAccessAppPolicyLinkParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmbeddedPoliciesPoliciesAccessAppPolicyLinkParam) ImplementsEmbeddedPoliciesPoliciesUnionParam() {
}

type EmbeddedPoliciesPoliciesIntersectionParam struct {
	// The UUID of the policy
	ID param.Field[string] `json:"id"`
	PolicyRequestForAppParam
}

func (r EmbeddedPoliciesPoliciesIntersectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmbeddedPoliciesPoliciesIntersectionParam) ImplementsEmbeddedPoliciesPoliciesUnionParam() {}

// The policies that Access applies to the application.
type EmbeddedResponsePolicies struct {
	Policies []PolicyResponseApp          `json:"policies"`
	JSON     embeddedResponsePoliciesJSON `json:"-"`
}

// embeddedResponsePoliciesJSON contains the JSON metadata for the struct
// [EmbeddedResponsePolicies]
type embeddedResponsePoliciesJSON struct {
	Policies    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmbeddedResponsePolicies) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r embeddedResponsePoliciesJSON) RawJSON() string {
	return r.raw
}

type EmbeddedScimConfigParam struct {
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r EmbeddedScimConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Contains the targets secured by the application.
type EmbeddedTargetCriteria struct {
	TargetCriteria []EmbeddedTargetCriteriaTargetCriterion `json:"target_criteria"`
	JSON           embeddedTargetCriteriaJSON              `json:"-"`
}

// embeddedTargetCriteriaJSON contains the JSON metadata for the struct
// [EmbeddedTargetCriteria]
type embeddedTargetCriteriaJSON struct {
	TargetCriteria apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmbeddedTargetCriteria) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r embeddedTargetCriteriaJSON) RawJSON() string {
	return r.raw
}

type EmbeddedTargetCriteriaTargetCriterion struct {
	// The port that the targets use for the chosen communication protocol. A port
	// cannot be assigned to multiple protocols.
	Port int64 `json:"port,required"`
	// The communication protocol your application secures.
	Protocol EmbeddedTargetCriteriaTargetCriteriaProtocol `json:"protocol,required"`
	// Contains a map of target attribute keys to target attribute values.
	TargetAttributes map[string][]string                       `json:"target_attributes,required"`
	JSON             embeddedTargetCriteriaTargetCriterionJSON `json:"-"`
}

// embeddedTargetCriteriaTargetCriterionJSON contains the JSON metadata for the
// struct [EmbeddedTargetCriteriaTargetCriterion]
type embeddedTargetCriteriaTargetCriterionJSON struct {
	Port             apijson.Field
	Protocol         apijson.Field
	TargetAttributes apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EmbeddedTargetCriteriaTargetCriterion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r embeddedTargetCriteriaTargetCriterionJSON) RawJSON() string {
	return r.raw
}

// The communication protocol your application secures.
type EmbeddedTargetCriteriaTargetCriteriaProtocol string

const (
	EmbeddedTargetCriteriaTargetCriteriaProtocolSSH EmbeddedTargetCriteriaTargetCriteriaProtocol = "ssh"
)

func (r EmbeddedTargetCriteriaTargetCriteriaProtocol) IsKnown() bool {
	switch r {
	case EmbeddedTargetCriteriaTargetCriteriaProtocolSSH:
		return true
	}
	return false
}

// Contains the targets secured by the application.
type EmbeddedTargetCriteriaParam struct {
	TargetCriteria param.Field[[]EmbeddedTargetCriteriaTargetCriterionParam] `json:"target_criteria"`
}

func (r EmbeddedTargetCriteriaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmbeddedTargetCriteriaTargetCriterionParam struct {
	// The port that the targets use for the chosen communication protocol. A port
	// cannot be assigned to multiple protocols.
	Port param.Field[int64] `json:"port,required"`
	// The communication protocol your application secures.
	Protocol param.Field[EmbeddedTargetCriteriaTargetCriteriaProtocol] `json:"protocol,required"`
	// Contains a map of target attribute keys to target attribute values.
	TargetAttributes param.Field[map[string][]string] `json:"target_attributes,required"`
}

func (r EmbeddedTargetCriteriaTargetCriterionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FeatureAppPropsAccount struct {
	// The application type.
	Type ApplicationType `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps []string `json:"allowed_idps"`
	// The image URL of the logo shown in the App Launcher header.
	AppLauncherLogoURL string `json:"app_launcher_logo_url"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool `json:"auto_redirect_to_identity"`
	// The background color of the App Launcher page.
	BgColor string `json:"bg_color"`
	// The primary hostname and path secured by Access. This domain will be displayed
	// if the app is visible in the App Launcher.
	Domain string `json:"domain"`
	// The links in the App Launcher footer.
	FooterLinks []FeatureAppPropsAccountFooterLink `json:"footer_links"`
	// The background color of the App Launcher header.
	HeaderBgColor string `json:"header_bg_color"`
	// The design of the App Launcher landing page shown to users when they log in.
	LandingPageDesign FeatureAppPropsAccountLandingPageDesign `json:"landing_page_design"`
	// The name of the application.
	Name string `json:"name"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration string `json:"session_duration"`
	// Determines when to skip the App Launcher landing page.
	SkipAppLauncherLoginPage bool                       `json:"skip_app_launcher_login_page"`
	JSON                     featureAppPropsAccountJSON `json:"-"`
}

// featureAppPropsAccountJSON contains the JSON metadata for the struct
// [FeatureAppPropsAccount]
type featureAppPropsAccountJSON struct {
	Type                     apijson.Field
	AllowedIdps              apijson.Field
	AppLauncherLogoURL       apijson.Field
	AutoRedirectToIdentity   apijson.Field
	BgColor                  apijson.Field
	Domain                   apijson.Field
	FooterLinks              apijson.Field
	HeaderBgColor            apijson.Field
	LandingPageDesign        apijson.Field
	Name                     apijson.Field
	SessionDuration          apijson.Field
	SkipAppLauncherLoginPage apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *FeatureAppPropsAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r featureAppPropsAccountJSON) RawJSON() string {
	return r.raw
}

type FeatureAppPropsAccountFooterLink struct {
	// The hypertext in the footer link.
	Name string `json:"name,required"`
	// the hyperlink in the footer link.
	URL  string                               `json:"url,required"`
	JSON featureAppPropsAccountFooterLinkJSON `json:"-"`
}

// featureAppPropsAccountFooterLinkJSON contains the JSON metadata for the struct
// [FeatureAppPropsAccountFooterLink]
type featureAppPropsAccountFooterLinkJSON struct {
	Name        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FeatureAppPropsAccountFooterLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r featureAppPropsAccountFooterLinkJSON) RawJSON() string {
	return r.raw
}

// The design of the App Launcher landing page shown to users when they log in.
type FeatureAppPropsAccountLandingPageDesign struct {
	// The background color of the log in button on the landing page.
	ButtonColor string `json:"button_color"`
	// The color of the text in the log in button on the landing page.
	ButtonTextColor string `json:"button_text_color"`
	// The URL of the image shown on the landing page.
	ImageURL string `json:"image_url"`
	// The message shown on the landing page.
	Message string `json:"message"`
	// The title shown on the landing page.
	Title string                                      `json:"title"`
	JSON  featureAppPropsAccountLandingPageDesignJSON `json:"-"`
}

// featureAppPropsAccountLandingPageDesignJSON contains the JSON metadata for the
// struct [FeatureAppPropsAccountLandingPageDesign]
type featureAppPropsAccountLandingPageDesignJSON struct {
	ButtonColor     apijson.Field
	ButtonTextColor apijson.Field
	ImageURL        apijson.Field
	Message         apijson.Field
	Title           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *FeatureAppPropsAccountLandingPageDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r featureAppPropsAccountLandingPageDesignJSON) RawJSON() string {
	return r.raw
}

type FeatureAppPropsAccountParam struct {
	// The application type.
	Type param.Field[ApplicationType] `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
	// The image URL of the logo shown in the App Launcher header.
	AppLauncherLogoURL param.Field[string] `json:"app_launcher_logo_url"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The background color of the App Launcher page.
	BgColor param.Field[string] `json:"bg_color"`
	// The primary hostname and path secured by Access. This domain will be displayed
	// if the app is visible in the App Launcher.
	Domain param.Field[string] `json:"domain"`
	// The links in the App Launcher footer.
	FooterLinks param.Field[[]FeatureAppPropsAccountFooterLinkParam] `json:"footer_links"`
	// The background color of the App Launcher header.
	HeaderBgColor param.Field[string] `json:"header_bg_color"`
	// The design of the App Launcher landing page shown to users when they log in.
	LandingPageDesign param.Field[FeatureAppPropsAccountLandingPageDesignParam] `json:"landing_page_design"`
	// The name of the application.
	Name param.Field[string] `json:"name"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// Determines when to skip the App Launcher landing page.
	SkipAppLauncherLoginPage param.Field[bool] `json:"skip_app_launcher_login_page"`
}

func (r FeatureAppPropsAccountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FeatureAppPropsAccountFooterLinkParam struct {
	// The hypertext in the footer link.
	Name param.Field[string] `json:"name,required"`
	// the hyperlink in the footer link.
	URL param.Field[string] `json:"url,required"`
}

func (r FeatureAppPropsAccountFooterLinkParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The design of the App Launcher landing page shown to users when they log in.
type FeatureAppPropsAccountLandingPageDesignParam struct {
	// The background color of the log in button on the landing page.
	ButtonColor param.Field[string] `json:"button_color"`
	// The color of the text in the log in button on the landing page.
	ButtonTextColor param.Field[string] `json:"button_text_color"`
	// The URL of the image shown on the landing page.
	ImageURL param.Field[string] `json:"image_url"`
	// The message shown on the landing page.
	Message param.Field[string] `json:"message"`
	// The title shown on the landing page.
	Title param.Field[string] `json:"title"`
}

func (r FeatureAppPropsAccountLandingPageDesignParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IDResponseApps struct {
	Result IDResponseAppsResult `json:"result"`
	JSON   idResponseAppsJSON   `json:"-"`
	APIResponseSingleAccess
}

// idResponseAppsJSON contains the JSON metadata for the struct [IDResponseApps]
type idResponseAppsJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseAppsJSON) RawJSON() string {
	return r.raw
}

type IDResponseAppsResult struct {
	// UUID
	ID   string                   `json:"id"`
	JSON idResponseAppsResultJSON `json:"-"`
}

// idResponseAppsResultJSON contains the JSON metadata for the struct
// [IDResponseAppsResult]
type idResponseAppsResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseAppsResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseAppsResultJSON) RawJSON() string {
	return r.raw
}

// Contains the targets secured by the application.
type InfraProps struct {
	// The name of the application.
	Name string `json:"name"`
	// The application type.
	Type ApplicationType `json:"type"`
	JSON infraPropsJSON  `json:"-"`
	EmbeddedTargetCriteria
}

// infraPropsJSON contains the JSON metadata for the struct [InfraProps]
type infraPropsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InfraProps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r infraPropsJSON) RawJSON() string {
	return r.raw
}

// Contains the targets secured by the application.
type InfraPropsParam struct {
	// The name of the application.
	Name param.Field[string] `json:"name"`
	// The application type.
	Type param.Field[ApplicationType] `json:"type"`
	EmbeddedTargetCriteriaParam
}

func (r InfraPropsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LauncherProps struct {
	Domain interface{} `json:"domain"`
	Name   interface{} `json:"name"`
	// The application type.
	Type string            `json:"type"`
	JSON launcherPropsJSON `json:"-"`
	FeatureAppPropsAccount
}

// launcherPropsJSON contains the JSON metadata for the struct [LauncherProps]
type launcherPropsJSON struct {
	Domain      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LauncherProps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r launcherPropsJSON) RawJSON() string {
	return r.raw
}

type LauncherPropsParam struct {
	// The application type.
	Type param.Field[string] `json:"type"`
	FeatureAppPropsAccountParam
}

func (r LauncherPropsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Contains the targets secured by the application.
type RdpProps struct {
	// The application type.
	Type string       `json:"type"`
	JSON rdpPropsJSON `json:"-"`
	EmbeddedTargetCriteria
	SelfHostedPropsAccount
}

// rdpPropsJSON contains the JSON metadata for the struct [RdpProps]
type rdpPropsJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RdpProps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rdpPropsJSON) RawJSON() string {
	return r.raw
}

// Contains the targets secured by the application.
type RdpPropsParam struct {
	// The application type.
	Type param.Field[string] `json:"type"`
	EmbeddedTargetCriteriaParam
	SelfHostedPropsAccountParam
}

func (r RdpPropsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SaasProps struct {
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps []string `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool `json:"auto_redirect_to_identity"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages []string `json:"custom_pages"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name    string           `json:"name"`
	SaasApp SaasPropsSaasApp `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags []string `json:"tags"`
	// The application type.
	Type string        `json:"type"`
	JSON saasPropsJSON `json:"-"`
}

// saasPropsJSON contains the JSON metadata for the struct [SaasProps]
type saasPropsJSON struct {
	AllowedIdps            apijson.Field
	AppLauncherVisible     apijson.Field
	AutoRedirectToIdentity apijson.Field
	CustomPages            apijson.Field
	LogoURL                apijson.Field
	Name                   apijson.Field
	SaasApp                apijson.Field
	Tags                   apijson.Field
	Type                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SaasProps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r saasPropsJSON) RawJSON() string {
	return r.raw
}

type SaasPropsSaasApp struct {
	// The lifetime of the OIDC Access Token after creation. Valid units are m,h. Must
	// be greater than or equal to 1m and less than or equal to 24h.
	AccessTokenLifetime string `json:"access_token_lifetime"`
	// If client secret should be required on the token endpoint when
	// authorization_code_with_pkce grant is used.
	AllowPkceWithoutClientSecret bool `json:"allow_pkce_without_client_secret"`
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType SaasPropsSaasAppAuthType `json:"auth_type"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string `json:"client_secret"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string    `json:"consumer_service_url"`
	CreatedAt          time.Time `json:"created_at" format:"date-time"`
	// This field can have the runtime type of
	// [[]SaasPropsSaasAppAccessSAMLSaasAppCustomAttribute].
	CustomAttributes interface{} `json:"custom_attributes"`
	// This field can have the runtime type of
	// [[]SaasPropsSaasAppAccessOidcSaasAppCustomClaim].
	CustomClaims interface{} `json:"custom_claims"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// This field can have the runtime type of
	// [[]SaasPropsSaasAppAccessOidcSaasAppGrantType].
	GrantTypes interface{} `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex string `json:"group_filter_regex"`
	// This field can have the runtime type of
	// [SaasPropsSaasAppAccessOidcSaasAppHybridAndImplicitOptions].
	HybridAndImplicitOptions interface{} `json:"hybrid_and_implicit_options"`
	// The unique identifier for your SaaS application.
	IdpEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat SaasPropsSaasAppNameIDFormat `json:"name_id_format"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's
	// user identities into a NameID value for its SAML assertion. This expression
	// should evaluate to a singular string. The output of this expression can override
	// the `name_id_format` setting.
	NameIDTransformJsonata string `json:"name_id_transform_jsonata"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// This field can have the runtime type of [[]string].
	RedirectUris interface{} `json:"redirect_uris"`
	// This field can have the runtime type of
	// [SaasPropsSaasAppAccessOidcSaasAppRefreshTokenOptions].
	RefreshTokenOptions interface{} `json:"refresh_token_options"`
	// A [JSONata] (https://jsonata.org/) expression that transforms an application's
	// user identities into attribute assertions in the SAML response. The expression
	// can transform id, email, name, and groups values. It can also transform fields
	// listed in the saml_attributes or oidc_fields of the identity provider used to
	// authenticate. The output of this expression must be a JSON object.
	SAMLAttributeTransformJsonata string `json:"saml_attribute_transform_jsonata"`
	// This field can have the runtime type of
	// [[]SaasPropsSaasAppAccessOidcSaasAppScope].
	Scopes interface{} `json:"scopes"`
	// A globally unique name for an identity or service provider.
	SpEntityID string `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint string               `json:"sso_endpoint"`
	UpdatedAt   time.Time            `json:"updated_at" format:"date-time"`
	JSON        saasPropsSaasAppJSON `json:"-"`
	union       SaasPropsSaasAppUnion
}

// saasPropsSaasAppJSON contains the JSON metadata for the struct
// [SaasPropsSaasApp]
type saasPropsSaasAppJSON struct {
	AccessTokenLifetime           apijson.Field
	AllowPkceWithoutClientSecret  apijson.Field
	AppLauncherURL                apijson.Field
	AuthType                      apijson.Field
	ClientID                      apijson.Field
	ClientSecret                  apijson.Field
	ConsumerServiceURL            apijson.Field
	CreatedAt                     apijson.Field
	CustomAttributes              apijson.Field
	CustomClaims                  apijson.Field
	DefaultRelayState             apijson.Field
	GrantTypes                    apijson.Field
	GroupFilterRegex              apijson.Field
	HybridAndImplicitOptions      apijson.Field
	IdpEntityID                   apijson.Field
	NameIDFormat                  apijson.Field
	NameIDTransformJsonata        apijson.Field
	PublicKey                     apijson.Field
	RedirectUris                  apijson.Field
	RefreshTokenOptions           apijson.Field
	SAMLAttributeTransformJsonata apijson.Field
	Scopes                        apijson.Field
	SpEntityID                    apijson.Field
	SSOEndpoint                   apijson.Field
	UpdatedAt                     apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r saasPropsSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r *SaasPropsSaasApp) UnmarshalJSON(data []byte) (err error) {
	*r = SaasPropsSaasApp{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SaasPropsSaasAppUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [SaasPropsSaasAppAccessSAMLSaasApp],
// [SaasPropsSaasAppAccessOidcSaasApp].
func (r SaasPropsSaasApp) AsUnion() SaasPropsSaasAppUnion {
	return r.union
}

// Union satisfied by [SaasPropsSaasAppAccessSAMLSaasApp] or
// [SaasPropsSaasAppAccessOidcSaasApp].
type SaasPropsSaasAppUnion interface {
	implementsSaasPropsSaasApp()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SaasPropsSaasAppUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SaasPropsSaasAppAccessSAMLSaasApp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SaasPropsSaasAppAccessOidcSaasApp{}),
		},
	)
}

type SaasPropsSaasAppAccessSAMLSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType SaasPropsSaasAppAccessSAMLSaasAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string                                             `json:"consumer_service_url"`
	CreatedAt          time.Time                                          `json:"created_at" format:"date-time"`
	CustomAttributes   []SaasPropsSaasAppAccessSAMLSaasAppCustomAttribute `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState string `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat SaasPropsSaasAppAccessSAMLSaasAppNameIDFormat `json:"name_id_format"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's
	// user identities into a NameID value for its SAML assertion. This expression
	// should evaluate to a singular string. The output of this expression can override
	// the `name_id_format` setting.
	NameIDTransformJsonata string `json:"name_id_transform_jsonata"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// A [JSONata] (https://jsonata.org/) expression that transforms an application's
	// user identities into attribute assertions in the SAML response. The expression
	// can transform id, email, name, and groups values. It can also transform fields
	// listed in the saml_attributes or oidc_fields of the identity provider used to
	// authenticate. The output of this expression must be a JSON object.
	SAMLAttributeTransformJsonata string `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SpEntityID string `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint string                                `json:"sso_endpoint"`
	UpdatedAt   time.Time                             `json:"updated_at" format:"date-time"`
	JSON        saasPropsSaasAppAccessSAMLSaasAppJSON `json:"-"`
}

// saasPropsSaasAppAccessSAMLSaasAppJSON contains the JSON metadata for the struct
// [SaasPropsSaasAppAccessSAMLSaasApp]
type saasPropsSaasAppAccessSAMLSaasAppJSON struct {
	AuthType                      apijson.Field
	ConsumerServiceURL            apijson.Field
	CreatedAt                     apijson.Field
	CustomAttributes              apijson.Field
	DefaultRelayState             apijson.Field
	IdpEntityID                   apijson.Field
	NameIDFormat                  apijson.Field
	NameIDTransformJsonata        apijson.Field
	PublicKey                     apijson.Field
	SAMLAttributeTransformJsonata apijson.Field
	SpEntityID                    apijson.Field
	SSOEndpoint                   apijson.Field
	UpdatedAt                     apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *SaasPropsSaasAppAccessSAMLSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r saasPropsSaasAppAccessSAMLSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r SaasPropsSaasAppAccessSAMLSaasApp) implementsSaasPropsSaasApp() {}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type SaasPropsSaasAppAccessSAMLSaasAppAuthType string

const (
	SaasPropsSaasAppAccessSAMLSaasAppAuthTypeSAML SaasPropsSaasAppAccessSAMLSaasAppAuthType = "saml"
	SaasPropsSaasAppAccessSAMLSaasAppAuthTypeOidc SaasPropsSaasAppAccessSAMLSaasAppAuthType = "oidc"
)

func (r SaasPropsSaasAppAccessSAMLSaasAppAuthType) IsKnown() bool {
	switch r {
	case SaasPropsSaasAppAccessSAMLSaasAppAuthTypeSAML, SaasPropsSaasAppAccessSAMLSaasAppAuthTypeOidc:
		return true
	}
	return false
}

type SaasPropsSaasAppAccessSAMLSaasAppCustomAttribute struct {
	// The SAML FriendlyName of the attribute.
	FriendlyName string `json:"friendly_name"`
	// The name of the attribute.
	Name string `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesNameFormat `json:"name_format"`
	// If the attribute is required when building a SAML assertion.
	Required bool                                                    `json:"required"`
	Source   SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesSource `json:"source"`
	JSON     saasPropsSaasAppAccessSAMLSaasAppCustomAttributeJSON    `json:"-"`
}

// saasPropsSaasAppAccessSAMLSaasAppCustomAttributeJSON contains the JSON metadata
// for the struct [SaasPropsSaasAppAccessSAMLSaasAppCustomAttribute]
type saasPropsSaasAppAccessSAMLSaasAppCustomAttributeJSON struct {
	FriendlyName apijson.Field
	Name         apijson.Field
	NameFormat   apijson.Field
	Required     apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SaasPropsSaasAppAccessSAMLSaasAppCustomAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r saasPropsSaasAppAccessSAMLSaasAppCustomAttributeJSON) RawJSON() string {
	return r.raw
}

// A globally unique name for an identity or service provider.
type SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesNameFormat string

const (
	SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatUnspecified SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatBasic       SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatUri         SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

func (r SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesNameFormat) IsKnown() bool {
	switch r {
	case SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatUnspecified, SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatBasic, SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatUri:
		return true
	}
	return false
}

type SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name string `json:"name"`
	// A mapping from IdP ID to attribute name.
	NameByIdp []SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesSourceNameByIdp `json:"name_by_idp"`
	JSON      saasPropsSaasAppAccessSAMLSaasAppCustomAttributesSourceJSON        `json:"-"`
}

// saasPropsSaasAppAccessSAMLSaasAppCustomAttributesSourceJSON contains the JSON
// metadata for the struct
// [SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesSource]
type saasPropsSaasAppAccessSAMLSaasAppCustomAttributesSourceJSON struct {
	Name        apijson.Field
	NameByIdp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r saasPropsSaasAppAccessSAMLSaasAppCustomAttributesSourceJSON) RawJSON() string {
	return r.raw
}

type SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesSourceNameByIdp struct {
	// The UID of the IdP.
	IdpID string `json:"idp_id"`
	// The name of the IdP provided attribute.
	SourceName string                                                               `json:"source_name"`
	JSON       saasPropsSaasAppAccessSAMLSaasAppCustomAttributesSourceNameByIdpJSON `json:"-"`
}

// saasPropsSaasAppAccessSAMLSaasAppCustomAttributesSourceNameByIdpJSON contains
// the JSON metadata for the struct
// [SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesSourceNameByIdp]
type saasPropsSaasAppAccessSAMLSaasAppCustomAttributesSourceNameByIdpJSON struct {
	IdpID       apijson.Field
	SourceName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesSourceNameByIdp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r saasPropsSaasAppAccessSAMLSaasAppCustomAttributesSourceNameByIdpJSON) RawJSON() string {
	return r.raw
}

// The format of the name identifier sent to the SaaS application.
type SaasPropsSaasAppAccessSAMLSaasAppNameIDFormat string

const (
	SaasPropsSaasAppAccessSAMLSaasAppNameIDFormatID    SaasPropsSaasAppAccessSAMLSaasAppNameIDFormat = "id"
	SaasPropsSaasAppAccessSAMLSaasAppNameIDFormatEmail SaasPropsSaasAppAccessSAMLSaasAppNameIDFormat = "email"
)

func (r SaasPropsSaasAppAccessSAMLSaasAppNameIDFormat) IsKnown() bool {
	switch r {
	case SaasPropsSaasAppAccessSAMLSaasAppNameIDFormatID, SaasPropsSaasAppAccessSAMLSaasAppNameIDFormatEmail:
		return true
	}
	return false
}

type SaasPropsSaasAppAccessOidcSaasApp struct {
	// The lifetime of the OIDC Access Token after creation. Valid units are m,h. Must
	// be greater than or equal to 1m and less than or equal to 24h.
	AccessTokenLifetime string `json:"access_token_lifetime"`
	// If client secret should be required on the token endpoint when
	// authorization_code_with_pkce grant is used.
	AllowPkceWithoutClientSecret bool `json:"allow_pkce_without_client_secret"`
	// The URL where this applications tile redirects users
	AppLauncherURL string `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType SaasPropsSaasAppAccessOidcSaasAppAuthType `json:"auth_type"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string                                         `json:"client_secret"`
	CreatedAt    time.Time                                      `json:"created_at" format:"date-time"`
	CustomClaims []SaasPropsSaasAppAccessOidcSaasAppCustomClaim `json:"custom_claims"`
	// The OIDC flows supported by this application
	GrantTypes []SaasPropsSaasAppAccessOidcSaasAppGrantType `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex         string                                                    `json:"group_filter_regex"`
	HybridAndImplicitOptions SaasPropsSaasAppAccessOidcSaasAppHybridAndImplicitOptions `json:"hybrid_and_implicit_options"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectUris        []string                                             `json:"redirect_uris"`
	RefreshTokenOptions SaasPropsSaasAppAccessOidcSaasAppRefreshTokenOptions `json:"refresh_token_options"`
	// Define the user information shared with access, "offline_access" scope will be
	// automatically enabled if refresh tokens are enabled
	Scopes    []SaasPropsSaasAppAccessOidcSaasAppScope `json:"scopes"`
	UpdatedAt time.Time                                `json:"updated_at" format:"date-time"`
	JSON      saasPropsSaasAppAccessOidcSaasAppJSON    `json:"-"`
}

// saasPropsSaasAppAccessOidcSaasAppJSON contains the JSON metadata for the struct
// [SaasPropsSaasAppAccessOidcSaasApp]
type saasPropsSaasAppAccessOidcSaasAppJSON struct {
	AccessTokenLifetime          apijson.Field
	AllowPkceWithoutClientSecret apijson.Field
	AppLauncherURL               apijson.Field
	AuthType                     apijson.Field
	ClientID                     apijson.Field
	ClientSecret                 apijson.Field
	CreatedAt                    apijson.Field
	CustomClaims                 apijson.Field
	GrantTypes                   apijson.Field
	GroupFilterRegex             apijson.Field
	HybridAndImplicitOptions     apijson.Field
	PublicKey                    apijson.Field
	RedirectUris                 apijson.Field
	RefreshTokenOptions          apijson.Field
	Scopes                       apijson.Field
	UpdatedAt                    apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SaasPropsSaasAppAccessOidcSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r saasPropsSaasAppAccessOidcSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r SaasPropsSaasAppAccessOidcSaasApp) implementsSaasPropsSaasApp() {}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type SaasPropsSaasAppAccessOidcSaasAppAuthType string

const (
	SaasPropsSaasAppAccessOidcSaasAppAuthTypeSAML SaasPropsSaasAppAccessOidcSaasAppAuthType = "saml"
	SaasPropsSaasAppAccessOidcSaasAppAuthTypeOidc SaasPropsSaasAppAccessOidcSaasAppAuthType = "oidc"
)

func (r SaasPropsSaasAppAccessOidcSaasAppAuthType) IsKnown() bool {
	switch r {
	case SaasPropsSaasAppAccessOidcSaasAppAuthTypeSAML, SaasPropsSaasAppAccessOidcSaasAppAuthTypeOidc:
		return true
	}
	return false
}

type SaasPropsSaasAppAccessOidcSaasAppCustomClaim struct {
	// The name of the claim.
	Name string `json:"name"`
	// If the claim is required when building an OIDC token.
	Required bool `json:"required"`
	// The scope of the claim.
	Scope  SaasPropsSaasAppAccessOidcSaasAppCustomClaimsScope  `json:"scope"`
	Source SaasPropsSaasAppAccessOidcSaasAppCustomClaimsSource `json:"source"`
	JSON   saasPropsSaasAppAccessOidcSaasAppCustomClaimJSON    `json:"-"`
}

// saasPropsSaasAppAccessOidcSaasAppCustomClaimJSON contains the JSON metadata for
// the struct [SaasPropsSaasAppAccessOidcSaasAppCustomClaim]
type saasPropsSaasAppAccessOidcSaasAppCustomClaimJSON struct {
	Name        apijson.Field
	Required    apijson.Field
	Scope       apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SaasPropsSaasAppAccessOidcSaasAppCustomClaim) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r saasPropsSaasAppAccessOidcSaasAppCustomClaimJSON) RawJSON() string {
	return r.raw
}

// The scope of the claim.
type SaasPropsSaasAppAccessOidcSaasAppCustomClaimsScope string

const (
	SaasPropsSaasAppAccessOidcSaasAppCustomClaimsScopeGroups  SaasPropsSaasAppAccessOidcSaasAppCustomClaimsScope = "groups"
	SaasPropsSaasAppAccessOidcSaasAppCustomClaimsScopeProfile SaasPropsSaasAppAccessOidcSaasAppCustomClaimsScope = "profile"
	SaasPropsSaasAppAccessOidcSaasAppCustomClaimsScopeEmail   SaasPropsSaasAppAccessOidcSaasAppCustomClaimsScope = "email"
	SaasPropsSaasAppAccessOidcSaasAppCustomClaimsScopeOpenid  SaasPropsSaasAppAccessOidcSaasAppCustomClaimsScope = "openid"
)

func (r SaasPropsSaasAppAccessOidcSaasAppCustomClaimsScope) IsKnown() bool {
	switch r {
	case SaasPropsSaasAppAccessOidcSaasAppCustomClaimsScopeGroups, SaasPropsSaasAppAccessOidcSaasAppCustomClaimsScopeProfile, SaasPropsSaasAppAccessOidcSaasAppCustomClaimsScopeEmail, SaasPropsSaasAppAccessOidcSaasAppCustomClaimsScopeOpenid:
		return true
	}
	return false
}

type SaasPropsSaasAppAccessOidcSaasAppCustomClaimsSource struct {
	// The name of the IdP claim.
	Name string `json:"name"`
	// A mapping from IdP ID to claim name.
	NameByIdp map[string]string                                       `json:"name_by_idp"`
	JSON      saasPropsSaasAppAccessOidcSaasAppCustomClaimsSourceJSON `json:"-"`
}

// saasPropsSaasAppAccessOidcSaasAppCustomClaimsSourceJSON contains the JSON
// metadata for the struct [SaasPropsSaasAppAccessOidcSaasAppCustomClaimsSource]
type saasPropsSaasAppAccessOidcSaasAppCustomClaimsSourceJSON struct {
	Name        apijson.Field
	NameByIdp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SaasPropsSaasAppAccessOidcSaasAppCustomClaimsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r saasPropsSaasAppAccessOidcSaasAppCustomClaimsSourceJSON) RawJSON() string {
	return r.raw
}

type SaasPropsSaasAppAccessOidcSaasAppGrantType string

const (
	SaasPropsSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode         SaasPropsSaasAppAccessOidcSaasAppGrantType = "authorization_code"
	SaasPropsSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce SaasPropsSaasAppAccessOidcSaasAppGrantType = "authorization_code_with_pkce"
	SaasPropsSaasAppAccessOidcSaasAppGrantTypeRefreshTokens             SaasPropsSaasAppAccessOidcSaasAppGrantType = "refresh_tokens"
	SaasPropsSaasAppAccessOidcSaasAppGrantTypeHybrid                    SaasPropsSaasAppAccessOidcSaasAppGrantType = "hybrid"
	SaasPropsSaasAppAccessOidcSaasAppGrantTypeImplicit                  SaasPropsSaasAppAccessOidcSaasAppGrantType = "implicit"
)

func (r SaasPropsSaasAppAccessOidcSaasAppGrantType) IsKnown() bool {
	switch r {
	case SaasPropsSaasAppAccessOidcSaasAppGrantTypeAuthorizationCode, SaasPropsSaasAppAccessOidcSaasAppGrantTypeAuthorizationCodeWithPkce, SaasPropsSaasAppAccessOidcSaasAppGrantTypeRefreshTokens, SaasPropsSaasAppAccessOidcSaasAppGrantTypeHybrid, SaasPropsSaasAppAccessOidcSaasAppGrantTypeImplicit:
		return true
	}
	return false
}

type SaasPropsSaasAppAccessOidcSaasAppHybridAndImplicitOptions struct {
	// If an Access Token should be returned from the OIDC Authorization endpoint
	ReturnAccessTokenFromAuthorizationEndpoint bool `json:"return_access_token_from_authorization_endpoint"`
	// If an ID Token should be returned from the OIDC Authorization endpoint
	ReturnIDTokenFromAuthorizationEndpoint bool                                                          `json:"return_id_token_from_authorization_endpoint"`
	JSON                                   saasPropsSaasAppAccessOidcSaasAppHybridAndImplicitOptionsJSON `json:"-"`
}

// saasPropsSaasAppAccessOidcSaasAppHybridAndImplicitOptionsJSON contains the JSON
// metadata for the struct
// [SaasPropsSaasAppAccessOidcSaasAppHybridAndImplicitOptions]
type saasPropsSaasAppAccessOidcSaasAppHybridAndImplicitOptionsJSON struct {
	ReturnAccessTokenFromAuthorizationEndpoint apijson.Field
	ReturnIDTokenFromAuthorizationEndpoint     apijson.Field
	raw                                        string
	ExtraFields                                map[string]apijson.Field
}

func (r *SaasPropsSaasAppAccessOidcSaasAppHybridAndImplicitOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r saasPropsSaasAppAccessOidcSaasAppHybridAndImplicitOptionsJSON) RawJSON() string {
	return r.raw
}

type SaasPropsSaasAppAccessOidcSaasAppRefreshTokenOptions struct {
	// How long a refresh token will be valid for after creation. Valid units are
	// m,h,d. Must be longer than 1m.
	Lifetime string                                                   `json:"lifetime"`
	JSON     saasPropsSaasAppAccessOidcSaasAppRefreshTokenOptionsJSON `json:"-"`
}

// saasPropsSaasAppAccessOidcSaasAppRefreshTokenOptionsJSON contains the JSON
// metadata for the struct [SaasPropsSaasAppAccessOidcSaasAppRefreshTokenOptions]
type saasPropsSaasAppAccessOidcSaasAppRefreshTokenOptionsJSON struct {
	Lifetime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SaasPropsSaasAppAccessOidcSaasAppRefreshTokenOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r saasPropsSaasAppAccessOidcSaasAppRefreshTokenOptionsJSON) RawJSON() string {
	return r.raw
}

type SaasPropsSaasAppAccessOidcSaasAppScope string

const (
	SaasPropsSaasAppAccessOidcSaasAppScopeOpenid  SaasPropsSaasAppAccessOidcSaasAppScope = "openid"
	SaasPropsSaasAppAccessOidcSaasAppScopeGroups  SaasPropsSaasAppAccessOidcSaasAppScope = "groups"
	SaasPropsSaasAppAccessOidcSaasAppScopeEmail   SaasPropsSaasAppAccessOidcSaasAppScope = "email"
	SaasPropsSaasAppAccessOidcSaasAppScopeProfile SaasPropsSaasAppAccessOidcSaasAppScope = "profile"
)

func (r SaasPropsSaasAppAccessOidcSaasAppScope) IsKnown() bool {
	switch r {
	case SaasPropsSaasAppAccessOidcSaasAppScopeOpenid, SaasPropsSaasAppAccessOidcSaasAppScopeGroups, SaasPropsSaasAppAccessOidcSaasAppScopeEmail, SaasPropsSaasAppAccessOidcSaasAppScopeProfile:
		return true
	}
	return false
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type SaasPropsSaasAppAuthType string

const (
	SaasPropsSaasAppAuthTypeSAML SaasPropsSaasAppAuthType = "saml"
	SaasPropsSaasAppAuthTypeOidc SaasPropsSaasAppAuthType = "oidc"
)

func (r SaasPropsSaasAppAuthType) IsKnown() bool {
	switch r {
	case SaasPropsSaasAppAuthTypeSAML, SaasPropsSaasAppAuthTypeOidc:
		return true
	}
	return false
}

// The format of the name identifier sent to the SaaS application.
type SaasPropsSaasAppNameIDFormat string

const (
	SaasPropsSaasAppNameIDFormatID    SaasPropsSaasAppNameIDFormat = "id"
	SaasPropsSaasAppNameIDFormatEmail SaasPropsSaasAppNameIDFormat = "email"
)

func (r SaasPropsSaasAppNameIDFormat) IsKnown() bool {
	switch r {
	case SaasPropsSaasAppNameIDFormatID, SaasPropsSaasAppNameIDFormatEmail:
		return true
	}
	return false
}

type SaasPropsParam struct {
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages param.Field[[]string] `json:"custom_pages"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL param.Field[string] `json:"logo_url"`
	// The name of the application.
	Name    param.Field[string]                     `json:"name"`
	SaasApp param.Field[SaasPropsSaasAppUnionParam] `json:"saas_app"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
	// The application type.
	Type param.Field[string] `json:"type"`
}

func (r SaasPropsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SaasPropsSaasAppParam struct {
	// The lifetime of the OIDC Access Token after creation. Valid units are m,h. Must
	// be greater than or equal to 1m and less than or equal to 24h.
	AccessTokenLifetime param.Field[string] `json:"access_token_lifetime"`
	// If client secret should be required on the token endpoint when
	// authorization_code_with_pkce grant is used.
	AllowPkceWithoutClientSecret param.Field[bool] `json:"allow_pkce_without_client_secret"`
	// The URL where this applications tile redirects users
	AppLauncherURL param.Field[string] `json:"app_launcher_url"`
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[SaasPropsSaasAppAuthType] `json:"auth_type"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string] `json:"client_secret"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]      `json:"consumer_service_url"`
	CreatedAt          param.Field[time.Time]   `json:"created_at" format:"date-time"`
	CustomAttributes   param.Field[interface{}] `json:"custom_attributes"`
	CustomClaims       param.Field[interface{}] `json:"custom_claims"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string]      `json:"default_relay_state"`
	GrantTypes        param.Field[interface{}] `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex         param.Field[string]      `json:"group_filter_regex"`
	HybridAndImplicitOptions param.Field[interface{}] `json:"hybrid_and_implicit_options"`
	// The unique identifier for your SaaS application.
	IdpEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[SaasPropsSaasAppNameIDFormat] `json:"name_id_format"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's
	// user identities into a NameID value for its SAML assertion. This expression
	// should evaluate to a singular string. The output of this expression can override
	// the `name_id_format` setting.
	NameIDTransformJsonata param.Field[string] `json:"name_id_transform_jsonata"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey           param.Field[string]      `json:"public_key"`
	RedirectUris        param.Field[interface{}] `json:"redirect_uris"`
	RefreshTokenOptions param.Field[interface{}] `json:"refresh_token_options"`
	// A [JSONata] (https://jsonata.org/) expression that transforms an application's
	// user identities into attribute assertions in the SAML response. The expression
	// can transform id, email, name, and groups values. It can also transform fields
	// listed in the saml_attributes or oidc_fields of the identity provider used to
	// authenticate. The output of this expression must be a JSON object.
	SAMLAttributeTransformJsonata param.Field[string]      `json:"saml_attribute_transform_jsonata"`
	Scopes                        param.Field[interface{}] `json:"scopes"`
	// A globally unique name for an identity or service provider.
	SpEntityID param.Field[string] `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint param.Field[string]    `json:"sso_endpoint"`
	UpdatedAt   param.Field[time.Time] `json:"updated_at" format:"date-time"`
}

func (r SaasPropsSaasAppParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SaasPropsSaasAppParam) implementsSaasPropsSaasAppUnionParam() {}

// Satisfied by [SaasPropsSaasAppAccessSAMLSaasAppParam],
// [SaasPropsSaasAppAccessOidcSaasAppParam], [SaasPropsSaasAppParam].
type SaasPropsSaasAppUnionParam interface {
	implementsSaasPropsSaasAppUnionParam()
}

type SaasPropsSaasAppAccessSAMLSaasAppParam struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[SaasPropsSaasAppAccessSAMLSaasAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]                                                  `json:"consumer_service_url"`
	CreatedAt          param.Field[time.Time]                                               `json:"created_at" format:"date-time"`
	CustomAttributes   param.Field[[]SaasPropsSaasAppAccessSAMLSaasAppCustomAttributeParam] `json:"custom_attributes"`
	// The URL that the user will be redirected to after a successful login for IDP
	// initiated logins.
	DefaultRelayState param.Field[string] `json:"default_relay_state"`
	// The unique identifier for your SaaS application.
	IdpEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[SaasPropsSaasAppAccessSAMLSaasAppNameIDFormat] `json:"name_id_format"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's
	// user identities into a NameID value for its SAML assertion. This expression
	// should evaluate to a singular string. The output of this expression can override
	// the `name_id_format` setting.
	NameIDTransformJsonata param.Field[string] `json:"name_id_transform_jsonata"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// A [JSONata] (https://jsonata.org/) expression that transforms an application's
	// user identities into attribute assertions in the SAML response. The expression
	// can transform id, email, name, and groups values. It can also transform fields
	// listed in the saml_attributes or oidc_fields of the identity provider used to
	// authenticate. The output of this expression must be a JSON object.
	SAMLAttributeTransformJsonata param.Field[string] `json:"saml_attribute_transform_jsonata"`
	// A globally unique name for an identity or service provider.
	SpEntityID param.Field[string] `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint param.Field[string]    `json:"sso_endpoint"`
	UpdatedAt   param.Field[time.Time] `json:"updated_at" format:"date-time"`
}

func (r SaasPropsSaasAppAccessSAMLSaasAppParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SaasPropsSaasAppAccessSAMLSaasAppParam) implementsSaasPropsSaasAppUnionParam() {}

type SaasPropsSaasAppAccessSAMLSaasAppCustomAttributeParam struct {
	// The SAML FriendlyName of the attribute.
	FriendlyName param.Field[string] `json:"friendly_name"`
	// The name of the attribute.
	Name param.Field[string] `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat param.Field[SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesNameFormat] `json:"name_format"`
	// If the attribute is required when building a SAML assertion.
	Required param.Field[bool]                                                         `json:"required"`
	Source   param.Field[SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesSourceParam] `json:"source"`
}

func (r SaasPropsSaasAppAccessSAMLSaasAppCustomAttributeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesSourceParam struct {
	// The name of the IdP attribute.
	Name param.Field[string] `json:"name"`
	// A mapping from IdP ID to attribute name.
	NameByIdp param.Field[[]SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesSourceNameByIdpParam] `json:"name_by_idp"`
}

func (r SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesSourceNameByIdpParam struct {
	// The UID of the IdP.
	IdpID param.Field[string] `json:"idp_id"`
	// The name of the IdP provided attribute.
	SourceName param.Field[string] `json:"source_name"`
}

func (r SaasPropsSaasAppAccessSAMLSaasAppCustomAttributesSourceNameByIdpParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SaasPropsSaasAppAccessOidcSaasAppParam struct {
	// The lifetime of the OIDC Access Token after creation. Valid units are m,h. Must
	// be greater than or equal to 1m and less than or equal to 24h.
	AccessTokenLifetime param.Field[string] `json:"access_token_lifetime"`
	// If client secret should be required on the token endpoint when
	// authorization_code_with_pkce grant is used.
	AllowPkceWithoutClientSecret param.Field[bool] `json:"allow_pkce_without_client_secret"`
	// The URL where this applications tile redirects users
	AppLauncherURL param.Field[string] `json:"app_launcher_url"`
	// Identifier of the authentication protocol used for the saas app. Required for
	// OIDC.
	AuthType param.Field[SaasPropsSaasAppAccessOidcSaasAppAuthType] `json:"auth_type"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string]                                              `json:"client_secret"`
	CreatedAt    param.Field[time.Time]                                           `json:"created_at" format:"date-time"`
	CustomClaims param.Field[[]SaasPropsSaasAppAccessOidcSaasAppCustomClaimParam] `json:"custom_claims"`
	// The OIDC flows supported by this application
	GrantTypes param.Field[[]SaasPropsSaasAppAccessOidcSaasAppGrantType] `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint
	GroupFilterRegex         param.Field[string]                                                         `json:"group_filter_regex"`
	HybridAndImplicitOptions param.Field[SaasPropsSaasAppAccessOidcSaasAppHybridAndImplicitOptionsParam] `json:"hybrid_and_implicit_options"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectUris        param.Field[[]string]                                                  `json:"redirect_uris"`
	RefreshTokenOptions param.Field[SaasPropsSaasAppAccessOidcSaasAppRefreshTokenOptionsParam] `json:"refresh_token_options"`
	// Define the user information shared with access, "offline_access" scope will be
	// automatically enabled if refresh tokens are enabled
	Scopes    param.Field[[]SaasPropsSaasAppAccessOidcSaasAppScope] `json:"scopes"`
	UpdatedAt param.Field[time.Time]                                `json:"updated_at" format:"date-time"`
}

func (r SaasPropsSaasAppAccessOidcSaasAppParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SaasPropsSaasAppAccessOidcSaasAppParam) implementsSaasPropsSaasAppUnionParam() {}

type SaasPropsSaasAppAccessOidcSaasAppCustomClaimParam struct {
	// The name of the claim.
	Name param.Field[string] `json:"name"`
	// If the claim is required when building an OIDC token.
	Required param.Field[bool] `json:"required"`
	// The scope of the claim.
	Scope  param.Field[SaasPropsSaasAppAccessOidcSaasAppCustomClaimsScope]       `json:"scope"`
	Source param.Field[SaasPropsSaasAppAccessOidcSaasAppCustomClaimsSourceParam] `json:"source"`
}

func (r SaasPropsSaasAppAccessOidcSaasAppCustomClaimParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SaasPropsSaasAppAccessOidcSaasAppCustomClaimsSourceParam struct {
	// The name of the IdP claim.
	Name param.Field[string] `json:"name"`
	// A mapping from IdP ID to claim name.
	NameByIdp param.Field[map[string]string] `json:"name_by_idp"`
}

func (r SaasPropsSaasAppAccessOidcSaasAppCustomClaimsSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SaasPropsSaasAppAccessOidcSaasAppHybridAndImplicitOptionsParam struct {
	// If an Access Token should be returned from the OIDC Authorization endpoint
	ReturnAccessTokenFromAuthorizationEndpoint param.Field[bool] `json:"return_access_token_from_authorization_endpoint"`
	// If an ID Token should be returned from the OIDC Authorization endpoint
	ReturnIDTokenFromAuthorizationEndpoint param.Field[bool] `json:"return_id_token_from_authorization_endpoint"`
}

func (r SaasPropsSaasAppAccessOidcSaasAppHybridAndImplicitOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SaasPropsSaasAppAccessOidcSaasAppRefreshTokenOptionsParam struct {
	// How long a refresh token will be valid for after creation. Valid units are
	// m,h,d. Must be longer than 1m.
	Lifetime param.Field[string] `json:"lifetime"`
}

func (r SaasPropsSaasAppAccessOidcSaasAppRefreshTokenOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SchemasEmptyResponse struct {
	Result  interface{}                 `json:"result,nullable"`
	Success SchemasEmptyResponseSuccess `json:"success"`
	JSON    schemasEmptyResponseJSON    `json:"-"`
}

// schemasEmptyResponseJSON contains the JSON metadata for the struct
// [SchemasEmptyResponse]
type schemasEmptyResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasEmptyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemasEmptyResponseJSON) RawJSON() string {
	return r.raw
}

type SchemasEmptyResponseSuccess bool

const (
	SchemasEmptyResponseSuccessTrue  SchemasEmptyResponseSuccess = true
	SchemasEmptyResponseSuccessFalse SchemasEmptyResponseSuccess = false
)

func (r SchemasEmptyResponseSuccess) IsKnown() bool {
	switch r {
	case SchemasEmptyResponseSuccessTrue, SchemasEmptyResponseSuccessFalse:
		return true
	}
	return false
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type ScimConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdpUid string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteUri string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication ScimConfigAuthenticationUnion `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []ScimConfigMapping `json:"mappings"`
	JSON     scimConfigJSON      `json:"-"`
}

// scimConfigJSON contains the JSON metadata for the struct [ScimConfig]
type scimConfigJSON struct {
	IdpUid             apijson.Field
	RemoteUri          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scimConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [ScimConfigAuthHTTPBasic],
// [ScimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerToken],
// [ScimConfigAuthOauth2], [ScimConfigAuthServiceToken] or
// [ScimConfigAuthenticationAccessScimConfigMultiAuthentication].
type ScimConfigAuthenticationUnion interface {
	implementsScimConfigAuthenticationUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScimConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScimConfigAuthHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScimConfigAuthOauth2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScimConfigAuthServiceToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScimConfigAuthenticationAccessScimConfigMultiAuthentication{}),
		},
	)
}

// Attributes for configuring OAuth Bearer Token authentication scheme for SCIM
// provisioning to an application.
type ScimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerToken struct {
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme ScimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerTokenScheme `json:"scheme,required"`
	JSON   scimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerTokenJSON   `json:"-"`
}

// scimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerTokenJSON
// contains the JSON metadata for the struct
// [ScimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerToken]
type scimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerTokenJSON struct {
	Token       apijson.Field
	Scheme      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerTokenJSON) RawJSON() string {
	return r.raw
}

func (r ScimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerToken) implementsScimConfigAuthenticationUnion() {
}

// The authentication scheme to use when making SCIM requests to this application.
type ScimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerTokenScheme string

const (
	ScimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerTokenSchemeOauthbearertoken ScimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerTokenScheme = "oauthbearertoken"
)

func (r ScimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerTokenScheme) IsKnown() bool {
	switch r {
	case ScimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerTokenSchemeOauthbearertoken:
		return true
	}
	return false
}

type ScimConfigAuthenticationAccessScimConfigMultiAuthentication []ScimConfigSingleAuth

func (r ScimConfigAuthenticationAccessScimConfigMultiAuthentication) implementsScimConfigAuthenticationUnion() {
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type ScimConfigParam struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdpUid param.Field[string] `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteUri param.Field[string] `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication param.Field[ScimConfigAuthenticationUnionParam] `json:"authentication"`
	// If false, propagates DELETE requests to the target application for SCIM
	// resources. If true, sets 'active' to false on the SCIM resource. Note: Some
	// targets do not support DELETE operations.
	DeactivateOnDelete param.Field[bool] `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings param.Field[[]ScimConfigMappingParam] `json:"mappings"`
}

func (r ScimConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type ScimConfigAuthenticationParam struct {
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[ScimConfigAuthenticationScheme] `json:"scheme,required"`
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string] `json:"client_secret"`
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string]      `json:"password"`
	Scopes   param.Field[interface{}] `json:"scopes"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user"`
}

func (r ScimConfigAuthenticationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScimConfigAuthenticationParam) implementsScimConfigAuthenticationUnionParam() {}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [ScimConfigAuthHTTPBasicParam],
// [ScimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerTokenParam],
// [ScimConfigAuthOauth2Param], [ScimConfigAuthServiceTokenParam],
// [ScimConfigAuthenticationAccessScimConfigMultiAuthenticationParam],
// [ScimConfigAuthenticationParam].
type ScimConfigAuthenticationUnionParam interface {
	implementsScimConfigAuthenticationUnionParam()
}

// Attributes for configuring OAuth Bearer Token authentication scheme for SCIM
// provisioning to an application.
type ScimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerTokenParam struct {
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[ScimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerTokenScheme] `json:"scheme,required"`
}

func (r ScimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScimConfigAuthenticationAccessScimConfigAuthenticationOAuthBearerTokenParam) implementsScimConfigAuthenticationUnionParam() {
}

type ScimConfigAuthenticationAccessScimConfigMultiAuthenticationParam []ScimConfigSingleAuthUnionParam

func (r ScimConfigAuthenticationAccessScimConfigMultiAuthenticationParam) implementsScimConfigAuthenticationUnionParam() {
}

// The authentication scheme to use when making SCIM requests to this application.
type ScimConfigAuthenticationScheme string

const (
	ScimConfigAuthenticationSchemeHttpbasic          ScimConfigAuthenticationScheme = "httpbasic"
	ScimConfigAuthenticationSchemeOauthbearertoken   ScimConfigAuthenticationScheme = "oauthbearertoken"
	ScimConfigAuthenticationSchemeOauth2             ScimConfigAuthenticationScheme = "oauth2"
	ScimConfigAuthenticationSchemeAccessServiceToken ScimConfigAuthenticationScheme = "access_service_token"
)

func (r ScimConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case ScimConfigAuthenticationSchemeHttpbasic, ScimConfigAuthenticationSchemeOauthbearertoken, ScimConfigAuthenticationSchemeOauth2, ScimConfigAuthenticationSchemeAccessServiceToken:
		return true
	}
	return false
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type ScimConfigAuthHTTPBasic struct {
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme ScimConfigAuthHTTPBasicScheme `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User string                      `json:"user,required"`
	JSON scimConfigAuthHTTPBasicJSON `json:"-"`
}

// scimConfigAuthHTTPBasicJSON contains the JSON metadata for the struct
// [ScimConfigAuthHTTPBasic]
type scimConfigAuthHTTPBasicJSON struct {
	Password    apijson.Field
	Scheme      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScimConfigAuthHTTPBasic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scimConfigAuthHTTPBasicJSON) RawJSON() string {
	return r.raw
}

func (r ScimConfigAuthHTTPBasic) implementsScimConfigAuthenticationUnion() {}

func (r ScimConfigAuthHTTPBasic) implementsScimConfigSingleAuth() {}

func (r ScimConfigAuthHTTPBasic) implementsBasicAppResponsePropsScimConfigAuthenticationUnion() {}

func (r ScimConfigAuthHTTPBasic) implementsScimConfigSingleAuthentication() {}

// The authentication scheme to use when making SCIM requests to this application.
type ScimConfigAuthHTTPBasicScheme string

const (
	ScimConfigAuthHTTPBasicSchemeHttpbasic ScimConfigAuthHTTPBasicScheme = "httpbasic"
)

func (r ScimConfigAuthHTTPBasicScheme) IsKnown() bool {
	switch r {
	case ScimConfigAuthHTTPBasicSchemeHttpbasic:
		return true
	}
	return false
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type ScimConfigAuthHTTPBasicParam struct {
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string] `json:"password,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[ScimConfigAuthHTTPBasicScheme] `json:"scheme,required"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user,required"`
}

func (r ScimConfigAuthHTTPBasicParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScimConfigAuthHTTPBasicParam) implementsScimConfigAuthenticationUnionParam() {}

func (r ScimConfigAuthHTTPBasicParam) implementsScimConfigSingleAuthUnionParam() {}

func (r ScimConfigAuthHTTPBasicParam) implementsBasicAppResponsePropsScimConfigAuthenticationUnionParam() {
}

func (r ScimConfigAuthHTTPBasicParam) implementsScimConfigSingleAuthenticationUnionParam() {}

// Attributes for configuring OAuth 2 authentication scheme for SCIM provisioning
// to an application.
type ScimConfigAuthOauth2 struct {
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url,required"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id,required"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme ScimConfigAuthOauth2Scheme `json:"scheme,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string `json:"token_url,required"`
	// The authorization scopes to request when generating the token used to
	// authenticate with the remove SCIM service.
	Scopes []string                 `json:"scopes"`
	JSON   scimConfigAuthOauth2JSON `json:"-"`
}

// scimConfigAuthOauth2JSON contains the JSON metadata for the struct
// [ScimConfigAuthOauth2]
type scimConfigAuthOauth2JSON struct {
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Scheme           apijson.Field
	TokenURL         apijson.Field
	Scopes           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScimConfigAuthOauth2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scimConfigAuthOauth2JSON) RawJSON() string {
	return r.raw
}

func (r ScimConfigAuthOauth2) implementsScimConfigAuthenticationUnion() {}

func (r ScimConfigAuthOauth2) implementsScimConfigSingleAuth() {}

func (r ScimConfigAuthOauth2) implementsBasicAppResponsePropsScimConfigAuthenticationUnion() {}

func (r ScimConfigAuthOauth2) implementsScimConfigSingleAuthentication() {}

// The authentication scheme to use when making SCIM requests to this application.
type ScimConfigAuthOauth2Scheme string

const (
	ScimConfigAuthOauth2SchemeOauth2 ScimConfigAuthOauth2Scheme = "oauth2"
)

func (r ScimConfigAuthOauth2Scheme) IsKnown() bool {
	switch r {
	case ScimConfigAuthOauth2SchemeOauth2:
		return true
	}
	return false
}

// Attributes for configuring OAuth 2 authentication scheme for SCIM provisioning
// to an application.
type ScimConfigAuthOauth2Param struct {
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url,required"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id,required"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[ScimConfigAuthOauth2Scheme] `json:"scheme,required"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url,required"`
	// The authorization scopes to request when generating the token used to
	// authenticate with the remove SCIM service.
	Scopes param.Field[[]string] `json:"scopes"`
}

func (r ScimConfigAuthOauth2Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScimConfigAuthOauth2Param) implementsScimConfigAuthenticationUnionParam() {}

func (r ScimConfigAuthOauth2Param) implementsScimConfigSingleAuthUnionParam() {}

func (r ScimConfigAuthOauth2Param) implementsBasicAppResponsePropsScimConfigAuthenticationUnionParam() {
}

func (r ScimConfigAuthOauth2Param) implementsScimConfigSingleAuthenticationUnionParam() {}

// Attributes for configuring Access Service Token authentication scheme for SCIM
// provisioning to an application.
type ScimConfigAuthServiceToken struct {
	// Client ID of the Access service token used to authenticate with the remote
	// service.
	ClientID string `json:"client_id,required"`
	// Client secret of the Access service token used to authenticate with the remote
	// service.
	ClientSecret string `json:"client_secret,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme ScimConfigAuthServiceTokenScheme `json:"scheme,required"`
	JSON   scimConfigAuthServiceTokenJSON   `json:"-"`
}

// scimConfigAuthServiceTokenJSON contains the JSON metadata for the struct
// [ScimConfigAuthServiceToken]
type scimConfigAuthServiceTokenJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	Scheme       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScimConfigAuthServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scimConfigAuthServiceTokenJSON) RawJSON() string {
	return r.raw
}

func (r ScimConfigAuthServiceToken) implementsScimConfigAuthenticationUnion() {}

func (r ScimConfigAuthServiceToken) implementsScimConfigSingleAuth() {}

func (r ScimConfigAuthServiceToken) implementsBasicAppResponsePropsScimConfigAuthenticationUnion() {}

func (r ScimConfigAuthServiceToken) implementsScimConfigSingleAuthentication() {}

// The authentication scheme to use when making SCIM requests to this application.
type ScimConfigAuthServiceTokenScheme string

const (
	ScimConfigAuthServiceTokenSchemeAccessServiceToken ScimConfigAuthServiceTokenScheme = "access_service_token"
)

func (r ScimConfigAuthServiceTokenScheme) IsKnown() bool {
	switch r {
	case ScimConfigAuthServiceTokenSchemeAccessServiceToken:
		return true
	}
	return false
}

// Attributes for configuring Access Service Token authentication scheme for SCIM
// provisioning to an application.
type ScimConfigAuthServiceTokenParam struct {
	// Client ID of the Access service token used to authenticate with the remote
	// service.
	ClientID param.Field[string] `json:"client_id,required"`
	// Client secret of the Access service token used to authenticate with the remote
	// service.
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[ScimConfigAuthServiceTokenScheme] `json:"scheme,required"`
}

func (r ScimConfigAuthServiceTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScimConfigAuthServiceTokenParam) implementsScimConfigAuthenticationUnionParam() {}

func (r ScimConfigAuthServiceTokenParam) implementsScimConfigSingleAuthUnionParam() {}

func (r ScimConfigAuthServiceTokenParam) implementsBasicAppResponsePropsScimConfigAuthenticationUnionParam() {
}

func (r ScimConfigAuthServiceTokenParam) implementsScimConfigSingleAuthenticationUnionParam() {}

// Transformations and filters applied to resources before they are provisioned in
// the remote SCIM service.
type ScimConfigMapping struct {
	// Which SCIM resource type this mapping applies to.
	Schema string `json:"schema,required"`
	// Whether or not this mapping is enabled.
	Enabled bool `json:"enabled"`
	// A
	// [SCIM filter expression](https://datatracker.ietf.org/doc/html/rfc7644#section-3.4.2.2)
	// that matches resources that should be provisioned to this application.
	Filter string `json:"filter"`
	// Whether or not this mapping applies to creates, updates, or deletes.
	Operations ScimConfigMappingOperations `json:"operations"`
	// The level of adherence to outbound resource schemas when provisioning to this
	// mapping. ‘Strict’ removes unknown values, while ‘passthrough’ passes unknown
	// values to the target.
	Strictness ScimConfigMappingStrictness `json:"strictness"`
	// A [JSONata](https://jsonata.org/) expression that transforms the resource before
	// provisioning it in the application.
	TransformJsonata string                `json:"transform_jsonata"`
	JSON             scimConfigMappingJSON `json:"-"`
}

// scimConfigMappingJSON contains the JSON metadata for the struct
// [ScimConfigMapping]
type scimConfigMappingJSON struct {
	Schema           apijson.Field
	Enabled          apijson.Field
	Filter           apijson.Field
	Operations       apijson.Field
	Strictness       apijson.Field
	TransformJsonata apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScimConfigMapping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scimConfigMappingJSON) RawJSON() string {
	return r.raw
}

// Whether or not this mapping applies to creates, updates, or deletes.
type ScimConfigMappingOperations struct {
	// Whether or not this mapping applies to create (POST) operations.
	Create bool `json:"create"`
	// Whether or not this mapping applies to DELETE operations.
	Delete bool `json:"delete"`
	// Whether or not this mapping applies to update (PATCH/PUT) operations.
	Update bool                            `json:"update"`
	JSON   scimConfigMappingOperationsJSON `json:"-"`
}

// scimConfigMappingOperationsJSON contains the JSON metadata for the struct
// [ScimConfigMappingOperations]
type scimConfigMappingOperationsJSON struct {
	Create      apijson.Field
	Delete      apijson.Field
	Update      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScimConfigMappingOperations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scimConfigMappingOperationsJSON) RawJSON() string {
	return r.raw
}

// The level of adherence to outbound resource schemas when provisioning to this
// mapping. ‘Strict’ removes unknown values, while ‘passthrough’ passes unknown
// values to the target.
type ScimConfigMappingStrictness string

const (
	ScimConfigMappingStrictnessStrict      ScimConfigMappingStrictness = "strict"
	ScimConfigMappingStrictnessPassthrough ScimConfigMappingStrictness = "passthrough"
)

func (r ScimConfigMappingStrictness) IsKnown() bool {
	switch r {
	case ScimConfigMappingStrictnessStrict, ScimConfigMappingStrictnessPassthrough:
		return true
	}
	return false
}

// Transformations and filters applied to resources before they are provisioned in
// the remote SCIM service.
type ScimConfigMappingParam struct {
	// Which SCIM resource type this mapping applies to.
	Schema param.Field[string] `json:"schema,required"`
	// Whether or not this mapping is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// A
	// [SCIM filter expression](https://datatracker.ietf.org/doc/html/rfc7644#section-3.4.2.2)
	// that matches resources that should be provisioned to this application.
	Filter param.Field[string] `json:"filter"`
	// Whether or not this mapping applies to creates, updates, or deletes.
	Operations param.Field[ScimConfigMappingOperationsParam] `json:"operations"`
	// The level of adherence to outbound resource schemas when provisioning to this
	// mapping. ‘Strict’ removes unknown values, while ‘passthrough’ passes unknown
	// values to the target.
	Strictness param.Field[ScimConfigMappingStrictness] `json:"strictness"`
	// A [JSONata](https://jsonata.org/) expression that transforms the resource before
	// provisioning it in the application.
	TransformJsonata param.Field[string] `json:"transform_jsonata"`
}

func (r ScimConfigMappingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not this mapping applies to creates, updates, or deletes.
type ScimConfigMappingOperationsParam struct {
	// Whether or not this mapping applies to create (POST) operations.
	Create param.Field[bool] `json:"create"`
	// Whether or not this mapping applies to DELETE operations.
	Delete param.Field[bool] `json:"delete"`
	// Whether or not this mapping applies to update (PATCH/PUT) operations.
	Update param.Field[bool] `json:"update"`
}

func (r ScimConfigMappingOperationsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type ScimConfigSingleAuth struct {
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme ScimConfigSingleAuthScheme `json:"scheme,required"`
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL string `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID string `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret string `json:"client_secret"`
	// Password used to authenticate with the remote SCIM service.
	Password string `json:"password"`
	// This field can have the runtime type of [[]string].
	Scopes interface{} `json:"scopes"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL string `json:"token_url"`
	// User name used to authenticate with the remote SCIM service.
	User  string                   `json:"user"`
	JSON  scimConfigSingleAuthJSON `json:"-"`
	union ScimConfigSingleAuthUnion
}

// scimConfigSingleAuthJSON contains the JSON metadata for the struct
// [ScimConfigSingleAuth]
type scimConfigSingleAuthJSON struct {
	Scheme           apijson.Field
	Token            apijson.Field
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	Password         apijson.Field
	Scopes           apijson.Field
	TokenURL         apijson.Field
	User             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r scimConfigSingleAuthJSON) RawJSON() string {
	return r.raw
}

func (r *ScimConfigSingleAuth) UnmarshalJSON(data []byte) (err error) {
	*r = ScimConfigSingleAuth{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScimConfigSingleAuthUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [ScimConfigAuthHTTPBasic],
// [ScimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerToken],
// [ScimConfigAuthOauth2], [ScimConfigAuthServiceToken].
func (r ScimConfigSingleAuth) AsUnion() ScimConfigSingleAuthUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [ScimConfigAuthHTTPBasic],
// [ScimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerToken],
// [ScimConfigAuthOauth2] or [ScimConfigAuthServiceToken].
type ScimConfigSingleAuthUnion interface {
	implementsScimConfigSingleAuth()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScimConfigSingleAuthUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScimConfigAuthHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerToken{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScimConfigAuthOauth2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScimConfigAuthServiceToken{}),
		},
	)
}

// Attributes for configuring OAuth Bearer Token authentication scheme for SCIM
// provisioning to an application.
type ScimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerToken struct {
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme ScimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerTokenScheme `json:"scheme,required"`
	JSON   scimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerTokenJSON   `json:"-"`
}

// scimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerTokenJSON contains
// the JSON metadata for the struct
// [ScimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerToken]
type scimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerTokenJSON struct {
	Token       apijson.Field
	Scheme      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerTokenJSON) RawJSON() string {
	return r.raw
}

func (r ScimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerToken) implementsScimConfigSingleAuth() {
}

// The authentication scheme to use when making SCIM requests to this application.
type ScimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerTokenScheme string

const (
	ScimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerTokenSchemeOauthbearertoken ScimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerTokenScheme = "oauthbearertoken"
)

func (r ScimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerTokenScheme) IsKnown() bool {
	switch r {
	case ScimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerTokenSchemeOauthbearertoken:
		return true
	}
	return false
}

// The authentication scheme to use when making SCIM requests to this application.
type ScimConfigSingleAuthScheme string

const (
	ScimConfigSingleAuthSchemeHttpbasic          ScimConfigSingleAuthScheme = "httpbasic"
	ScimConfigSingleAuthSchemeOauthbearertoken   ScimConfigSingleAuthScheme = "oauthbearertoken"
	ScimConfigSingleAuthSchemeOauth2             ScimConfigSingleAuthScheme = "oauth2"
	ScimConfigSingleAuthSchemeAccessServiceToken ScimConfigSingleAuthScheme = "access_service_token"
)

func (r ScimConfigSingleAuthScheme) IsKnown() bool {
	switch r {
	case ScimConfigSingleAuthSchemeHttpbasic, ScimConfigSingleAuthSchemeOauthbearertoken, ScimConfigSingleAuthSchemeOauth2, ScimConfigSingleAuthSchemeAccessServiceToken:
		return true
	}
	return false
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type ScimConfigSingleAuthParam struct {
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[ScimConfigSingleAuthScheme] `json:"scheme,required"`
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token"`
	// URL used to generate the auth code used during token generation.
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	// Client ID used to authenticate when generating a token for authenticating with
	// the remote SCIM service.
	ClientID param.Field[string] `json:"client_id"`
	// Secret used to authenticate when generating a token for authenticating with the
	// remove SCIM service.
	ClientSecret param.Field[string] `json:"client_secret"`
	// Password used to authenticate with the remote SCIM service.
	Password param.Field[string]      `json:"password"`
	Scopes   param.Field[interface{}] `json:"scopes"`
	// URL used to generate the token used to authenticate with the remote SCIM
	// service.
	TokenURL param.Field[string] `json:"token_url"`
	// User name used to authenticate with the remote SCIM service.
	User param.Field[string] `json:"user"`
}

func (r ScimConfigSingleAuthParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScimConfigSingleAuthParam) implementsScimConfigSingleAuthUnionParam() {}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [ScimConfigAuthHTTPBasicParam],
// [ScimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerTokenParam],
// [ScimConfigAuthOauth2Param], [ScimConfigAuthServiceTokenParam],
// [ScimConfigSingleAuthParam].
type ScimConfigSingleAuthUnionParam interface {
	implementsScimConfigSingleAuthUnionParam()
}

// Attributes for configuring OAuth Bearer Token authentication scheme for SCIM
// provisioning to an application.
type ScimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerTokenParam struct {
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[ScimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerTokenScheme] `json:"scheme,required"`
}

func (r ScimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScimConfigSingleAuthAccessScimConfigAuthenticationOAuthBearerTokenParam) implementsScimConfigSingleAuthUnionParam() {
}

type SelfHostedPropsAccount struct {
	// The primary hostname and path secured by Access. This domain will be displayed
	// if the app is visible in the App Launcher.
	Domain string `json:"domain,required"`
	// The application type.
	Type string `json:"type,required"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWarp bool `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps []string `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool                              `json:"auto_redirect_to_identity"`
	CorsHeaders            SelfHostedPropsAccountCorsHeaders `json:"cors_headers"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage string `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing identity-based rules.
	CustomDenyURL string `json:"custom_deny_url"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing non-identity rules.
	CustomNonIdentityDenyURL string `json:"custom_non_identity_deny_url"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages []string `json:"custom_pages"`
	// List of destinations secured by Access. This supersedes `self_hosted_domains` to
	// allow for more flexibility in defining different types of domains. If
	// `destinations` are provided, then `self_hosted_domains` will be ignored.
	Destinations []SelfHostedPropsAccountDestination `json:"destinations"`
	// Enables the binding cookie, which increases security against compromised
	// authorization tokens and CSRF attacks.
	EnableBindingCookie bool `json:"enable_binding_cookie"`
	// Enables the HttpOnly cookie attribute, which increases security against XSS
	// attacks.
	HTTPOnlyCookieAttribute bool `json:"http_only_cookie_attribute"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name string `json:"name"`
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute bool `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// List of public domains that Access will secure. This field is deprecated in
	// favor of `destinations` and will be supported until **November 21, 2025.** If
	// `destinations` are provided, then `self_hosted_domains` will be ignored.
	//
	// Deprecated: deprecated
	SelfHostedDomains []string `json:"self_hosted_domains"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect bool `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration string `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial bool `json:"skip_interstitial"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags []string                   `json:"tags"`
	JSON selfHostedPropsAccountJSON `json:"-"`
}

// selfHostedPropsAccountJSON contains the JSON metadata for the struct
// [SelfHostedPropsAccount]
type selfHostedPropsAccountJSON struct {
	Domain                   apijson.Field
	Type                     apijson.Field
	AllowAuthenticateViaWarp apijson.Field
	AllowedIdps              apijson.Field
	AppLauncherVisible       apijson.Field
	AutoRedirectToIdentity   apijson.Field
	CorsHeaders              apijson.Field
	CustomDenyMessage        apijson.Field
	CustomDenyURL            apijson.Field
	CustomNonIdentityDenyURL apijson.Field
	CustomPages              apijson.Field
	Destinations             apijson.Field
	EnableBindingCookie      apijson.Field
	HTTPOnlyCookieAttribute  apijson.Field
	LogoURL                  apijson.Field
	Name                     apijson.Field
	OptionsPreflightBypass   apijson.Field
	PathCookieAttribute      apijson.Field
	SameSiteCookieAttribute  apijson.Field
	SelfHostedDomains        apijson.Field
	ServiceAuth401Redirect   apijson.Field
	SessionDuration          apijson.Field
	SkipInterstitial         apijson.Field
	Tags                     apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *SelfHostedPropsAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r selfHostedPropsAccountJSON) RawJSON() string {
	return r.raw
}

type SelfHostedPropsAccountCorsHeaders struct {
	// Allows all HTTP request headers.
	AllowAllHeaders bool `json:"allow_all_headers"`
	// Allows all HTTP request methods.
	AllowAllMethods bool `json:"allow_all_methods"`
	// Allows all origins.
	AllowAllOrigins bool `json:"allow_all_origins"`
	// When set to `true`, includes credentials (cookies, authorization headers, or TLS
	// client certificates) with requests.
	AllowCredentials bool `json:"allow_credentials"`
	// Allowed HTTP request headers.
	AllowedHeaders []string `json:"allowed_headers"`
	// Allowed HTTP request methods.
	AllowedMethods []AllowedMethodsItem `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []string `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                               `json:"max_age"`
	JSON   selfHostedPropsAccountCorsHeadersJSON `json:"-"`
}

// selfHostedPropsAccountCorsHeadersJSON contains the JSON metadata for the struct
// [SelfHostedPropsAccountCorsHeaders]
type selfHostedPropsAccountCorsHeadersJSON struct {
	AllowAllHeaders  apijson.Field
	AllowAllMethods  apijson.Field
	AllowAllOrigins  apijson.Field
	AllowCredentials apijson.Field
	AllowedHeaders   apijson.Field
	AllowedMethods   apijson.Field
	AllowedOrigins   apijson.Field
	MaxAge           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SelfHostedPropsAccountCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r selfHostedPropsAccountCorsHeadersJSON) RawJSON() string {
	return r.raw
}

// A public hostname that Access will secure. Public destinations support
// sub-domain and path. Wildcard '\*' can be used in the definition.
type SelfHostedPropsAccountDestination struct {
	// The CIDR range of the destination. Single IPs will be computed as /32.
	Cidr string `json:"cidr"`
	// The hostname of the destination. Matches a valid SNI served by an HTTPS origin.
	Hostname string `json:"hostname"`
	// The L4 protocol of the destination. When omitted, both UDP and TCP traffic will
	// match.
	L4Protocol SelfHostedPropsAccountDestinationsL4Protocol `json:"l4_protocol"`
	// The port range of the destination. Can be a single port or a range of ports.
	// When omitted, all ports will match.
	PortRange string                                 `json:"port_range"`
	Type      SelfHostedPropsAccountDestinationsType `json:"type"`
	// The URI of the destination. Public destinations' URIs can include a domain and
	// path with
	// [wildcards](https://developers.cloudflare.com/cloudflare-one/policies/access/app-paths/).
	Uri string `json:"uri"`
	// The VNET ID to match the destination. When omitted, all VNETs will match.
	VnetID string                                `json:"vnet_id"`
	JSON   selfHostedPropsAccountDestinationJSON `json:"-"`
	union  SelfHostedPropsAccountDestinationsUnion
}

// selfHostedPropsAccountDestinationJSON contains the JSON metadata for the struct
// [SelfHostedPropsAccountDestination]
type selfHostedPropsAccountDestinationJSON struct {
	Cidr        apijson.Field
	Hostname    apijson.Field
	L4Protocol  apijson.Field
	PortRange   apijson.Field
	Type        apijson.Field
	Uri         apijson.Field
	VnetID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r selfHostedPropsAccountDestinationJSON) RawJSON() string {
	return r.raw
}

func (r *SelfHostedPropsAccountDestination) UnmarshalJSON(data []byte) (err error) {
	*r = SelfHostedPropsAccountDestination{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SelfHostedPropsAccountDestinationsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SelfHostedPropsAccountDestinationsPublicDestination],
// [SelfHostedPropsAccountDestinationsPrivateDestination].
func (r SelfHostedPropsAccountDestination) AsUnion() SelfHostedPropsAccountDestinationsUnion {
	return r.union
}

// A public hostname that Access will secure. Public destinations support
// sub-domain and path. Wildcard '\*' can be used in the definition.
//
// Union satisfied by [SelfHostedPropsAccountDestinationsPublicDestination] or
// [SelfHostedPropsAccountDestinationsPrivateDestination].
type SelfHostedPropsAccountDestinationsUnion interface {
	implementsSelfHostedPropsAccountDestination()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SelfHostedPropsAccountDestinationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SelfHostedPropsAccountDestinationsPublicDestination{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SelfHostedPropsAccountDestinationsPrivateDestination{}),
		},
	)
}

// A public hostname that Access will secure. Public destinations support
// sub-domain and path. Wildcard '\*' can be used in the definition.
type SelfHostedPropsAccountDestinationsPublicDestination struct {
	Type SelfHostedPropsAccountDestinationsPublicDestinationType `json:"type"`
	// The URI of the destination. Public destinations' URIs can include a domain and
	// path with
	// [wildcards](https://developers.cloudflare.com/cloudflare-one/policies/access/app-paths/).
	Uri  string                                                  `json:"uri"`
	JSON selfHostedPropsAccountDestinationsPublicDestinationJSON `json:"-"`
}

// selfHostedPropsAccountDestinationsPublicDestinationJSON contains the JSON
// metadata for the struct [SelfHostedPropsAccountDestinationsPublicDestination]
type selfHostedPropsAccountDestinationsPublicDestinationJSON struct {
	Type        apijson.Field
	Uri         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SelfHostedPropsAccountDestinationsPublicDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r selfHostedPropsAccountDestinationsPublicDestinationJSON) RawJSON() string {
	return r.raw
}

func (r SelfHostedPropsAccountDestinationsPublicDestination) implementsSelfHostedPropsAccountDestination() {
}

type SelfHostedPropsAccountDestinationsPublicDestinationType string

const (
	SelfHostedPropsAccountDestinationsPublicDestinationTypePublic SelfHostedPropsAccountDestinationsPublicDestinationType = "public"
)

func (r SelfHostedPropsAccountDestinationsPublicDestinationType) IsKnown() bool {
	switch r {
	case SelfHostedPropsAccountDestinationsPublicDestinationTypePublic:
		return true
	}
	return false
}

// Private destinations are an early access feature and gated behind a feature
// flag.
type SelfHostedPropsAccountDestinationsPrivateDestination struct {
	// The CIDR range of the destination. Single IPs will be computed as /32.
	Cidr string `json:"cidr"`
	// The hostname of the destination. Matches a valid SNI served by an HTTPS origin.
	Hostname string `json:"hostname"`
	// The L4 protocol of the destination. When omitted, both UDP and TCP traffic will
	// match.
	L4Protocol SelfHostedPropsAccountDestinationsPrivateDestinationL4Protocol `json:"l4_protocol"`
	// The port range of the destination. Can be a single port or a range of ports.
	// When omitted, all ports will match.
	PortRange string                                                   `json:"port_range"`
	Type      SelfHostedPropsAccountDestinationsPrivateDestinationType `json:"type"`
	// The VNET ID to match the destination. When omitted, all VNETs will match.
	VnetID string                                                   `json:"vnet_id"`
	JSON   selfHostedPropsAccountDestinationsPrivateDestinationJSON `json:"-"`
}

// selfHostedPropsAccountDestinationsPrivateDestinationJSON contains the JSON
// metadata for the struct [SelfHostedPropsAccountDestinationsPrivateDestination]
type selfHostedPropsAccountDestinationsPrivateDestinationJSON struct {
	Cidr        apijson.Field
	Hostname    apijson.Field
	L4Protocol  apijson.Field
	PortRange   apijson.Field
	Type        apijson.Field
	VnetID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SelfHostedPropsAccountDestinationsPrivateDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r selfHostedPropsAccountDestinationsPrivateDestinationJSON) RawJSON() string {
	return r.raw
}

func (r SelfHostedPropsAccountDestinationsPrivateDestination) implementsSelfHostedPropsAccountDestination() {
}

// The L4 protocol of the destination. When omitted, both UDP and TCP traffic will
// match.
type SelfHostedPropsAccountDestinationsPrivateDestinationL4Protocol string

const (
	SelfHostedPropsAccountDestinationsPrivateDestinationL4ProtocolTcp SelfHostedPropsAccountDestinationsPrivateDestinationL4Protocol = "tcp"
	SelfHostedPropsAccountDestinationsPrivateDestinationL4ProtocolUdp SelfHostedPropsAccountDestinationsPrivateDestinationL4Protocol = "udp"
)

func (r SelfHostedPropsAccountDestinationsPrivateDestinationL4Protocol) IsKnown() bool {
	switch r {
	case SelfHostedPropsAccountDestinationsPrivateDestinationL4ProtocolTcp, SelfHostedPropsAccountDestinationsPrivateDestinationL4ProtocolUdp:
		return true
	}
	return false
}

type SelfHostedPropsAccountDestinationsPrivateDestinationType string

const (
	SelfHostedPropsAccountDestinationsPrivateDestinationTypePrivate SelfHostedPropsAccountDestinationsPrivateDestinationType = "private"
)

func (r SelfHostedPropsAccountDestinationsPrivateDestinationType) IsKnown() bool {
	switch r {
	case SelfHostedPropsAccountDestinationsPrivateDestinationTypePrivate:
		return true
	}
	return false
}

// The L4 protocol of the destination. When omitted, both UDP and TCP traffic will
// match.
type SelfHostedPropsAccountDestinationsL4Protocol string

const (
	SelfHostedPropsAccountDestinationsL4ProtocolTcp SelfHostedPropsAccountDestinationsL4Protocol = "tcp"
	SelfHostedPropsAccountDestinationsL4ProtocolUdp SelfHostedPropsAccountDestinationsL4Protocol = "udp"
)

func (r SelfHostedPropsAccountDestinationsL4Protocol) IsKnown() bool {
	switch r {
	case SelfHostedPropsAccountDestinationsL4ProtocolTcp, SelfHostedPropsAccountDestinationsL4ProtocolUdp:
		return true
	}
	return false
}

type SelfHostedPropsAccountDestinationsType string

const (
	SelfHostedPropsAccountDestinationsTypePublic  SelfHostedPropsAccountDestinationsType = "public"
	SelfHostedPropsAccountDestinationsTypePrivate SelfHostedPropsAccountDestinationsType = "private"
)

func (r SelfHostedPropsAccountDestinationsType) IsKnown() bool {
	switch r {
	case SelfHostedPropsAccountDestinationsTypePublic, SelfHostedPropsAccountDestinationsTypePrivate:
		return true
	}
	return false
}

type SelfHostedPropsAccountParam struct {
	// The primary hostname and path secured by Access. This domain will be displayed
	// if the app is visible in the App Launcher.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// When set to true, users can authenticate to this application using their WARP
	// session. When set to false this application will always require direct IdP
	// authentication. This setting always overrides the organization setting for WARP
	// authentication.
	AllowAuthenticateViaWarp param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]                                   `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[SelfHostedPropsAccountCorsHeadersParam] `json:"cors_headers"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage param.Field[string] `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing identity-based rules.
	CustomDenyURL param.Field[string] `json:"custom_deny_url"`
	// The custom URL a user is redirected to when they are denied access to the
	// application when failing non-identity rules.
	CustomNonIdentityDenyURL param.Field[string] `json:"custom_non_identity_deny_url"`
	// The custom pages that will be displayed when applicable for this application
	CustomPages param.Field[[]string] `json:"custom_pages"`
	// List of destinations secured by Access. This supersedes `self_hosted_domains` to
	// allow for more flexibility in defining different types of domains. If
	// `destinations` are provided, then `self_hosted_domains` will be ignored.
	Destinations param.Field[[]SelfHostedPropsAccountDestinationsUnionParam] `json:"destinations"`
	// Enables the binding cookie, which increases security against compromised
	// authorization tokens and CSRF attacks.
	EnableBindingCookie param.Field[bool] `json:"enable_binding_cookie"`
	// Enables the HttpOnly cookie attribute, which increases security against XSS
	// attacks.
	HTTPOnlyCookieAttribute param.Field[bool] `json:"http_only_cookie_attribute"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL param.Field[string] `json:"logo_url"`
	// The name of the application.
	Name param.Field[string] `json:"name"`
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass param.Field[bool] `json:"options_preflight_bypass"`
	// Enables cookie paths to scope an application's JWT to the application path. If
	// disabled, the JWT will scope to the hostname by default
	PathCookieAttribute param.Field[bool] `json:"path_cookie_attribute"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// List of public domains that Access will secure. This field is deprecated in
	// favor of `destinations` and will be supported until **November 21, 2025.** If
	// `destinations` are provided, then `self_hosted_domains` will be ignored.
	//
	// Deprecated: deprecated
	SelfHostedDomains param.Field[[]string] `json:"self_hosted_domains"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect param.Field[bool] `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial param.Field[bool] `json:"skip_interstitial"`
	// The tags you want assigned to an application. Tags are used to filter
	// applications in the App Launcher dashboard.
	Tags param.Field[[]string] `json:"tags"`
}

func (r SelfHostedPropsAccountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SelfHostedPropsAccountCorsHeadersParam struct {
	// Allows all HTTP request headers.
	AllowAllHeaders param.Field[bool] `json:"allow_all_headers"`
	// Allows all HTTP request methods.
	AllowAllMethods param.Field[bool] `json:"allow_all_methods"`
	// Allows all origins.
	AllowAllOrigins param.Field[bool] `json:"allow_all_origins"`
	// When set to `true`, includes credentials (cookies, authorization headers, or TLS
	// client certificates) with requests.
	AllowCredentials param.Field[bool] `json:"allow_credentials"`
	// Allowed HTTP request headers.
	AllowedHeaders param.Field[[]string] `json:"allowed_headers"`
	// Allowed HTTP request methods.
	AllowedMethods param.Field[[]AllowedMethodsItem] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]string] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r SelfHostedPropsAccountCorsHeadersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A public hostname that Access will secure. Public destinations support
// sub-domain and path. Wildcard '\*' can be used in the definition.
type SelfHostedPropsAccountDestinationParam struct {
	// The CIDR range of the destination. Single IPs will be computed as /32.
	Cidr param.Field[string] `json:"cidr"`
	// The hostname of the destination. Matches a valid SNI served by an HTTPS origin.
	Hostname param.Field[string] `json:"hostname"`
	// The L4 protocol of the destination. When omitted, both UDP and TCP traffic will
	// match.
	L4Protocol param.Field[SelfHostedPropsAccountDestinationsL4Protocol] `json:"l4_protocol"`
	// The port range of the destination. Can be a single port or a range of ports.
	// When omitted, all ports will match.
	PortRange param.Field[string]                                 `json:"port_range"`
	Type      param.Field[SelfHostedPropsAccountDestinationsType] `json:"type"`
	// The URI of the destination. Public destinations' URIs can include a domain and
	// path with
	// [wildcards](https://developers.cloudflare.com/cloudflare-one/policies/access/app-paths/).
	Uri param.Field[string] `json:"uri"`
	// The VNET ID to match the destination. When omitted, all VNETs will match.
	VnetID param.Field[string] `json:"vnet_id"`
}

func (r SelfHostedPropsAccountDestinationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SelfHostedPropsAccountDestinationParam) implementsSelfHostedPropsAccountDestinationsUnionParam() {
}

// A public hostname that Access will secure. Public destinations support
// sub-domain and path. Wildcard '\*' can be used in the definition.
//
// Satisfied by [SelfHostedPropsAccountDestinationsPublicDestinationParam],
// [SelfHostedPropsAccountDestinationsPrivateDestinationParam],
// [SelfHostedPropsAccountDestinationParam].
type SelfHostedPropsAccountDestinationsUnionParam interface {
	implementsSelfHostedPropsAccountDestinationsUnionParam()
}

// A public hostname that Access will secure. Public destinations support
// sub-domain and path. Wildcard '\*' can be used in the definition.
type SelfHostedPropsAccountDestinationsPublicDestinationParam struct {
	Type param.Field[SelfHostedPropsAccountDestinationsPublicDestinationType] `json:"type"`
	// The URI of the destination. Public destinations' URIs can include a domain and
	// path with
	// [wildcards](https://developers.cloudflare.com/cloudflare-one/policies/access/app-paths/).
	Uri param.Field[string] `json:"uri"`
}

func (r SelfHostedPropsAccountDestinationsPublicDestinationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SelfHostedPropsAccountDestinationsPublicDestinationParam) implementsSelfHostedPropsAccountDestinationsUnionParam() {
}

// Private destinations are an early access feature and gated behind a feature
// flag.
type SelfHostedPropsAccountDestinationsPrivateDestinationParam struct {
	// The CIDR range of the destination. Single IPs will be computed as /32.
	Cidr param.Field[string] `json:"cidr"`
	// The hostname of the destination. Matches a valid SNI served by an HTTPS origin.
	Hostname param.Field[string] `json:"hostname"`
	// The L4 protocol of the destination. When omitted, both UDP and TCP traffic will
	// match.
	L4Protocol param.Field[SelfHostedPropsAccountDestinationsPrivateDestinationL4Protocol] `json:"l4_protocol"`
	// The port range of the destination. Can be a single port or a range of ports.
	// When omitted, all ports will match.
	PortRange param.Field[string]                                                   `json:"port_range"`
	Type      param.Field[SelfHostedPropsAccountDestinationsPrivateDestinationType] `json:"type"`
	// The VNET ID to match the destination. When omitted, all VNETs will match.
	VnetID param.Field[string] `json:"vnet_id"`
}

func (r SelfHostedPropsAccountDestinationsPrivateDestinationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SelfHostedPropsAccountDestinationsPrivateDestinationParam) implementsSelfHostedPropsAccountDestinationsUnionParam() {
}

type SingleResponseApp struct {
	// The policies that Access applies to the application.
	Result AppResponse           `json:"result"`
	JSON   singleResponseAppJSON `json:"-"`
	APIResponseSingleAccess
}

// singleResponseAppJSON contains the JSON metadata for the struct
// [SingleResponseApp]
type singleResponseAppJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseAppJSON) RawJSON() string {
	return r.raw
}

type SSHProps struct {
	// The application type.
	Type string       `json:"type"`
	JSON sshPropsJSON `json:"-"`
	SelfHostedPropsAccount
}

// sshPropsJSON contains the JSON metadata for the struct [SSHProps]
type sshPropsJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSHProps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sshPropsJSON) RawJSON() string {
	return r.raw
}

type SSHPropsParam struct {
	// The application type.
	Type param.Field[string] `json:"type"`
	SelfHostedPropsAccountParam
}

func (r SSHPropsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VncProps struct {
	// The application type.
	Type string       `json:"type"`
	JSON vncPropsJSON `json:"-"`
	SelfHostedPropsAccount
}

// vncPropsJSON contains the JSON metadata for the struct [VncProps]
type vncPropsJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VncProps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vncPropsJSON) RawJSON() string {
	return r.raw
}

type VncPropsParam struct {
	// The application type.
	Type param.Field[string] `json:"type"`
	SelfHostedPropsAccountParam
}

func (r VncPropsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WarpProps struct {
	Domain interface{} `json:"domain"`
	Name   interface{} `json:"name"`
	// The application type.
	Type string        `json:"type"`
	JSON warpPropsJSON `json:"-"`
	FeatureAppPropsAccount
}

// warpPropsJSON contains the JSON metadata for the struct [WarpProps]
type warpPropsJSON struct {
	Domain      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WarpProps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpPropsJSON) RawJSON() string {
	return r.raw
}

type WarpPropsParam struct {
	// The application type.
	Type param.Field[string] `json:"type"`
	FeatureAppPropsAccountParam
}

func (r WarpPropsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessAppNewResponse struct {
	// The policies that Access applies to the application.
	Result AppResponse                     `json:"result"`
	JSON   accountAccessAppNewResponseJSON `json:"-"`
	SingleResponseApp
}

// accountAccessAppNewResponseJSON contains the JSON metadata for the struct
// [AccountAccessAppNewResponse]
type accountAccessAppNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessAppNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessAppNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessAppUpdateResponse struct {
	// The policies that Access applies to the application.
	Result AppResponse                        `json:"result"`
	JSON   accountAccessAppUpdateResponseJSON `json:"-"`
	SingleResponseApp
}

// accountAccessAppUpdateResponseJSON contains the JSON metadata for the struct
// [AccountAccessAppUpdateResponse]
type accountAccessAppUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessAppUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessAppUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessAppListResponse struct {
	Result []AppResponse                    `json:"result"`
	JSON   accountAccessAppListResponseJSON `json:"-"`
	APIResponseCollectionAccess
}

// accountAccessAppListResponseJSON contains the JSON metadata for the struct
// [AccountAccessAppListResponse]
type accountAccessAppListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessAppListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessAppListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessAppUserPolicyChecksResponse struct {
	Result AccountAccessAppUserPolicyChecksResponseResult `json:"result"`
	JSON   accountAccessAppUserPolicyChecksResponseJSON   `json:"-"`
	APIResponseSingleAccess
}

// accountAccessAppUserPolicyChecksResponseJSON contains the JSON metadata for the
// struct [AccountAccessAppUserPolicyChecksResponse]
type accountAccessAppUserPolicyChecksResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessAppUserPolicyChecksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessAppUserPolicyChecksResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessAppUserPolicyChecksResponseResult struct {
	AppState     AccountAccessAppUserPolicyChecksResponseResultAppState     `json:"app_state"`
	UserIdentity AccountAccessAppUserPolicyChecksResponseResultUserIdentity `json:"user_identity"`
	JSON         accountAccessAppUserPolicyChecksResponseResultJSON         `json:"-"`
}

// accountAccessAppUserPolicyChecksResponseResultJSON contains the JSON metadata
// for the struct [AccountAccessAppUserPolicyChecksResponseResult]
type accountAccessAppUserPolicyChecksResponseResultJSON struct {
	AppState     apijson.Field
	UserIdentity apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessAppUserPolicyChecksResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessAppUserPolicyChecksResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAccessAppUserPolicyChecksResponseResultAppState struct {
	// UUID
	AppUid   string                                                     `json:"app_uid"`
	Aud      string                                                     `json:"aud"`
	Hostname string                                                     `json:"hostname"`
	Name     string                                                     `json:"name"`
	Policies []interface{}                                              `json:"policies"`
	Status   string                                                     `json:"status"`
	JSON     accountAccessAppUserPolicyChecksResponseResultAppStateJSON `json:"-"`
}

// accountAccessAppUserPolicyChecksResponseResultAppStateJSON contains the JSON
// metadata for the struct [AccountAccessAppUserPolicyChecksResponseResultAppState]
type accountAccessAppUserPolicyChecksResponseResultAppStateJSON struct {
	AppUid      apijson.Field
	Aud         apijson.Field
	Hostname    apijson.Field
	Name        apijson.Field
	Policies    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessAppUserPolicyChecksResponseResultAppState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessAppUserPolicyChecksResponseResultAppStateJSON) RawJSON() string {
	return r.raw
}

type AccountAccessAppUserPolicyChecksResponseResultUserIdentity struct {
	ID             string                                                        `json:"id"`
	AccountID      string                                                        `json:"account_id"`
	DeviceSessions interface{}                                                   `json:"device_sessions"`
	Email          string                                                        `json:"email"`
	Geo            AccountAccessAppUserPolicyChecksResponseResultUserIdentityGeo `json:"geo"`
	Iat            int64                                                         `json:"iat"`
	IsGateway      bool                                                          `json:"is_gateway"`
	IsWarp         bool                                                          `json:"is_warp"`
	Name           string                                                        `json:"name"`
	// UUID
	UserUuid string                                                         `json:"user_uuid"`
	Version  int64                                                          `json:"version"`
	JSON     accountAccessAppUserPolicyChecksResponseResultUserIdentityJSON `json:"-"`
}

// accountAccessAppUserPolicyChecksResponseResultUserIdentityJSON contains the JSON
// metadata for the struct
// [AccountAccessAppUserPolicyChecksResponseResultUserIdentity]
type accountAccessAppUserPolicyChecksResponseResultUserIdentityJSON struct {
	ID             apijson.Field
	AccountID      apijson.Field
	DeviceSessions apijson.Field
	Email          apijson.Field
	Geo            apijson.Field
	Iat            apijson.Field
	IsGateway      apijson.Field
	IsWarp         apijson.Field
	Name           apijson.Field
	UserUuid       apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessAppUserPolicyChecksResponseResultUserIdentity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessAppUserPolicyChecksResponseResultUserIdentityJSON) RawJSON() string {
	return r.raw
}

type AccountAccessAppUserPolicyChecksResponseResultUserIdentityGeo struct {
	Country string                                                            `json:"country"`
	JSON    accountAccessAppUserPolicyChecksResponseResultUserIdentityGeoJSON `json:"-"`
}

// accountAccessAppUserPolicyChecksResponseResultUserIdentityGeoJSON contains the
// JSON metadata for the struct
// [AccountAccessAppUserPolicyChecksResponseResultUserIdentityGeo]
type accountAccessAppUserPolicyChecksResponseResultUserIdentityGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessAppUserPolicyChecksResponseResultUserIdentityGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessAppUserPolicyChecksResponseResultUserIdentityGeoJSON) RawJSON() string {
	return r.raw
}

type AccountAccessAppNewParams struct {
	// Contains the targets secured by the application.
	AppRequest AppRequestUnionParam `json:"app_request,required"`
}

func (r AccountAccessAppNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AppRequest)
}

type AccountAccessAppUpdateParams struct {
	// Contains the targets secured by the application.
	AppRequest AppRequestUnionParam `json:"app_request,required"`
}

func (r AccountAccessAppUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AppRequest)
}

type AccountAccessAppListParams struct {
	// The aud of the app.
	Aud param.Field[string] `query:"aud"`
	// The domain of the app.
	Domain param.Field[string] `query:"domain"`
	// The name of the app.
	Name param.Field[string] `query:"name"`
	// Search for apps by other listed query parameters.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccountAccessAppListParams]'s query parameters as
// `url.Values`.
func (r AccountAccessAppListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
