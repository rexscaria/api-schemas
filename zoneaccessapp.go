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

// ZoneAccessAppService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAccessAppService] method instead.
type ZoneAccessAppService struct {
	Options  []option.RequestOption
	Ca       *ZoneAccessAppCaService
	Policies *ZoneAccessAppPolicyService
}

// NewZoneAccessAppService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneAccessAppService(opts ...option.RequestOption) (r *ZoneAccessAppService) {
	r = &ZoneAccessAppService{}
	r.Options = opts
	r.Ca = NewZoneAccessAppCaService(opts...)
	r.Policies = NewZoneAccessAppPolicyService(opts...)
	return
}

// Adds a new application to Access.
func (r *ZoneAccessAppService) New(ctx context.Context, zoneID string, body ZoneAccessAppNewParams, opts ...option.RequestOption) (res *ZoneAccessAppNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/apps", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches information about an Access application.
func (r *ZoneAccessAppService) Get(ctx context.Context, zoneID string, appID AppIDParam, opts ...option.RequestOption) (res *SingleResponseAppZone, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/apps/%s", zoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an Access application.
func (r *ZoneAccessAppService) Update(ctx context.Context, zoneID string, appID AppIDParam, body ZoneAccessAppUpdateParams, opts ...option.RequestOption) (res *ZoneAccessAppUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/apps/%s", zoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all Access Applications in a zone.
func (r *ZoneAccessAppService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneAccessAppListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/apps", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an application from Access.
func (r *ZoneAccessAppService) Delete(ctx context.Context, zoneID string, appID AppIDParam, opts ...option.RequestOption) (res *IDResponseApps, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/apps/%s", zoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Revokes all tokens issued for an application.
func (r *ZoneAccessAppService) RevokeTokens(ctx context.Context, zoneID string, appID AppIDParam, opts ...option.RequestOption) (res *SchemasEmptyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/apps/%s/revoke_tokens", zoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Tests if a specific user has permission to access an application.
func (r *ZoneAccessAppService) UserPolicyChecks(ctx context.Context, zoneID string, appID AppIDParam, opts ...option.RequestOption) (res *ZoneAccessAppUserPolicyChecksResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/apps/%s/user_policy_checks", zoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessApps struct {
	// UUID
	ID string `json:"id"`
	// This field can have the runtime type of [[]string].
	AllowedIdps interface{} `json:"allowed_idps"`
	// This field can have the runtime type of [bool], [interface{}].
	AppLauncherVisible interface{} `json:"app_launcher_visible"`
	// Audience tag.
	Aud string `json:"aud"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool `json:"auto_redirect_to_identity"`
	// This field can have the runtime type of [SelfHostedPropsZoneCorsHeaders].
	CorsHeaders interface{} `json:"cors_headers"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage string `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application.
	CustomDenyURL string `json:"custom_deny_url"`
	// This field can have the runtime type of [string], [interface{}].
	Domain interface{} `json:"domain"`
	// Enables the binding cookie, which increases security against compromised
	// authorization tokens and CSRF attacks.
	EnableBindingCookie bool `json:"enable_binding_cookie"`
	// Enables the HttpOnly cookie attribute, which increases security against XSS
	// attacks.
	HTTPOnlyCookieAttribute bool `json:"http_only_cookie_attribute"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// This field can have the runtime type of [string].
	Name interface{} `json:"name"`
	// Allows options preflight requests to bypass Access authentication and go
	// directly to the origin. Cannot turn on if cors_headers is set.
	OptionsPreflightBypass bool `json:"options_preflight_bypass"`
	// This field can have the runtime type of [AccessAppsSaaSApplicationSaasApp].
	SaasApp interface{} `json:"saas_app"`
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// This field can have the runtime type of [BasicAppResponsePropsScimConfig].
	ScimConfig interface{} `json:"scim_config"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect bool `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration string `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial bool `json:"skip_interstitial"`
	// The application type.
	Type      string         `json:"type"`
	UpdatedAt time.Time      `json:"updated_at" format:"date-time"`
	JSON      accessAppsJSON `json:"-"`
	union     AccessAppsUnion
}

// accessAppsJSON contains the JSON metadata for the struct [AccessApps]
type accessAppsJSON struct {
	ID                      apijson.Field
	AllowedIdps             apijson.Field
	AppLauncherVisible      apijson.Field
	Aud                     apijson.Field
	AutoRedirectToIdentity  apijson.Field
	CorsHeaders             apijson.Field
	CreatedAt               apijson.Field
	CustomDenyMessage       apijson.Field
	CustomDenyURL           apijson.Field
	Domain                  apijson.Field
	EnableBindingCookie     apijson.Field
	HTTPOnlyCookieAttribute apijson.Field
	LogoURL                 apijson.Field
	Name                    apijson.Field
	OptionsPreflightBypass  apijson.Field
	SaasApp                 apijson.Field
	SameSiteCookieAttribute apijson.Field
	ScimConfig              apijson.Field
	ServiceAuth401Redirect  apijson.Field
	SessionDuration         apijson.Field
	SkipInterstitial        apijson.Field
	Type                    apijson.Field
	UpdatedAt               apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r accessAppsJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApps) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApps{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessAppsUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [AccessAppsSelfHostedApplication],
// [AccessAppsSaaSApplication], [AccessAppsBrowserSSHApplication],
// [AccessAppsBrowserVncApplication], [AccessAppsAppLauncherApplication],
// [AccessAppsDeviceEnrollmentPermissionsApplication],
// [AccessAppsBrowserIsolationPermissionsApplication],
// [AccessAppsBookmarkApplication].
func (r AccessApps) AsUnion() AccessAppsUnion {
	return r.union
}

// Union satisfied by [AccessAppsSelfHostedApplication],
// [AccessAppsSaaSApplication], [AccessAppsBrowserSSHApplication],
// [AccessAppsBrowserVncApplication], [AccessAppsAppLauncherApplication],
// [AccessAppsDeviceEnrollmentPermissionsApplication],
// [AccessAppsBrowserIsolationPermissionsApplication] or
// [AccessAppsBookmarkApplication].
type AccessAppsUnion interface {
	implementsAccessApps()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessAppsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsSelfHostedApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsSaaSApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsBrowserSSHApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsBrowserVncApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsAppLauncherApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsDeviceEnrollmentPermissionsApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsBrowserIsolationPermissionsApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsBookmarkApplication{}),
		},
	)
}

type AccessAppsSelfHostedApplication struct {
	JSON accessAppsSelfHostedApplicationJSON `json:"-"`
	BasicAppResponseProps
	SelfHostedPropsZone
}

// accessAppsSelfHostedApplicationJSON contains the JSON metadata for the struct
// [AccessAppsSelfHostedApplication]
type accessAppsSelfHostedApplicationJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppsSelfHostedApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsSelfHostedApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsSelfHostedApplication) implementsAccessApps() {}

type AccessAppsSaaSApplication struct {
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps []string `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool `json:"auto_redirect_to_identity"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name    string                           `json:"name"`
	SaasApp AccessAppsSaaSApplicationSaasApp `json:"saas_app"`
	// The application type.
	Type string                        `json:"type"`
	JSON accessAppsSaaSApplicationJSON `json:"-"`
	BasicAppResponseProps
}

// accessAppsSaaSApplicationJSON contains the JSON metadata for the struct
// [AccessAppsSaaSApplication]
type accessAppsSaaSApplicationJSON struct {
	AllowedIdps            apijson.Field
	AppLauncherVisible     apijson.Field
	AutoRedirectToIdentity apijson.Field
	LogoURL                apijson.Field
	Name                   apijson.Field
	SaasApp                apijson.Field
	Type                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessAppsSaaSApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsSaaSApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsSaaSApplication) implementsAccessApps() {}

type AccessAppsSaaSApplicationSaasApp struct {
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
	AuthType AccessAppsSaaSApplicationSaasAppAuthType `json:"auth_type"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string `json:"client_secret"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string    `json:"consumer_service_url"`
	CreatedAt          time.Time `json:"created_at" format:"date-time"`
	// This field can have the runtime type of
	// [[]AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttribute].
	CustomAttributes interface{} `json:"custom_attributes"`
	// This field can have the runtime type of
	// [[]AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaim].
	CustomClaims interface{} `json:"custom_claims"`
	// This field can have the runtime type of
	// [[]AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantType].
	GrantTypes interface{} `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint.
	GroupFilterRegex string `json:"group_filter_regex"`
	// This field can have the runtime type of
	// [AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppHybridAndImplicitOptions].
	HybridAndImplicitOptions interface{} `json:"hybrid_and_implicit_options"`
	// The unique identifier for your SaaS application.
	IdpEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat AccessAppsSaaSApplicationSaasAppNameIDFormat `json:"name_id_format"`
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
	// [AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppRefreshTokenOptions].
	RefreshTokenOptions interface{} `json:"refresh_token_options"`
	// This field can have the runtime type of
	// [[]AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppScope].
	Scopes interface{} `json:"scopes"`
	// A globally unique name for an identity or service provider.
	SpEntityID string `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint string                               `json:"sso_endpoint"`
	UpdatedAt   time.Time                            `json:"updated_at" format:"date-time"`
	JSON        accessAppsSaaSApplicationSaasAppJSON `json:"-"`
	union       AccessAppsSaaSApplicationSaasAppUnion
}

// accessAppsSaaSApplicationSaasAppJSON contains the JSON metadata for the struct
// [AccessAppsSaaSApplicationSaasApp]
type accessAppsSaaSApplicationSaasAppJSON struct {
	AccessTokenLifetime          apijson.Field
	AllowPkceWithoutClientSecret apijson.Field
	AppLauncherURL               apijson.Field
	AuthType                     apijson.Field
	ClientID                     apijson.Field
	ClientSecret                 apijson.Field
	ConsumerServiceURL           apijson.Field
	CreatedAt                    apijson.Field
	CustomAttributes             apijson.Field
	CustomClaims                 apijson.Field
	GrantTypes                   apijson.Field
	GroupFilterRegex             apijson.Field
	HybridAndImplicitOptions     apijson.Field
	IdpEntityID                  apijson.Field
	NameIDFormat                 apijson.Field
	NameIDTransformJsonata       apijson.Field
	PublicKey                    apijson.Field
	RedirectUris                 apijson.Field
	RefreshTokenOptions          apijson.Field
	Scopes                       apijson.Field
	SpEntityID                   apijson.Field
	SSOEndpoint                  apijson.Field
	UpdatedAt                    apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r accessAppsSaaSApplicationSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r *AccessAppsSaaSApplicationSaasApp) UnmarshalJSON(data []byte) (err error) {
	*r = AccessAppsSaaSApplicationSaasApp{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessAppsSaaSApplicationSaasAppUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasApp],
// [AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasApp].
func (r AccessAppsSaaSApplicationSaasApp) AsUnion() AccessAppsSaaSApplicationSaasAppUnion {
	return r.union
}

// Union satisfied by [AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasApp] or
// [AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasApp].
type AccessAppsSaaSApplicationSaasAppUnion interface {
	implementsAccessAppsSaaSApplicationSaasApp()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessAppsSaaSApplicationSaasAppUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasApp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasApp{}),
		},
	)
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasApp struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppAuthType `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL string                                                                    `json:"consumer_service_url"`
	CreatedAt          time.Time                                                                 `json:"created_at" format:"date-time"`
	CustomAttributes   []AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttribute `json:"custom_attributes"`
	// The unique identifier for your SaaS application.
	IdpEntityID string `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppNameIDFormat `json:"name_id_format"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's
	// user identities into a NameID value for its SAML assertion. This expression
	// should evaluate to a singular string. The output of this expression can override
	// the `name_id_format` setting.
	NameIDTransformJsonata string `json:"name_id_transform_jsonata"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// A globally unique name for an identity or service provider.
	SpEntityID string `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint string                                                       `json:"sso_endpoint"`
	UpdatedAt   time.Time                                                    `json:"updated_at" format:"date-time"`
	JSON        accessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppJSON `json:"-"`
}

// accessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppJSON contains the JSON
// metadata for the struct
// [AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasApp]
type accessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppJSON struct {
	AuthType               apijson.Field
	ConsumerServiceURL     apijson.Field
	CreatedAt              apijson.Field
	CustomAttributes       apijson.Field
	IdpEntityID            apijson.Field
	NameIDFormat           apijson.Field
	NameIDTransformJsonata apijson.Field
	PublicKey              apijson.Field
	SpEntityID             apijson.Field
	SSOEndpoint            apijson.Field
	UpdatedAt              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasApp) implementsAccessAppsSaaSApplicationSaasApp() {
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppAuthType string

const (
	AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppAuthTypeSAML AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppAuthType = "saml"
	AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppAuthTypeOidc AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppAuthType = "oidc"
)

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppAuthType) IsKnown() bool {
	switch r {
	case AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppAuthTypeSAML, AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppAuthTypeOidc:
		return true
	}
	return false
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttribute struct {
	// The SAML FriendlyName of the attribute.
	FriendlyName string `json:"friendly_name"`
	// The name of the attribute.
	Name string `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesNameFormat `json:"name_format"`
	// If the attribute is required when building a SAML assertion.
	Required bool                                                                           `json:"required"`
	Source   AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesSource `json:"source"`
	JSON     accessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributeJSON    `json:"-"`
}

// accessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributeJSON
// contains the JSON metadata for the struct
// [AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttribute]
type accessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributeJSON struct {
	FriendlyName apijson.Field
	Name         apijson.Field
	NameFormat   apijson.Field
	Required     apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributeJSON) RawJSON() string {
	return r.raw
}

// A globally unique name for an identity or service provider.
type AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesNameFormat string

const (
	AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatUnspecified AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"
	AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatBasic       AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
	AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatUri         AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesNameFormat = "urn:oasis:names:tc:SAML:2.0:attrname-format:uri"
)

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesNameFormat) IsKnown() bool {
	switch r {
	case AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatUnspecified, AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatBasic, AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesNameFormatUrnOasisNamesTcSAML2_0AttrnameFormatUri:
		return true
	}
	return false
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesSource struct {
	// The name of the IdP attribute.
	Name string `json:"name"`
	// A mapping from IdP ID to attribute name.
	NameByIdp map[string]string                                                                  `json:"name_by_idp"`
	JSON      accessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesSourceJSON `json:"-"`
}

// accessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesSourceJSON
// contains the JSON metadata for the struct
// [AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesSource]
type accessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesSourceJSON struct {
	Name        apijson.Field
	NameByIdp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesSourceJSON) RawJSON() string {
	return r.raw
}

// The format of the name identifier sent to the SaaS application.
type AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppNameIDFormat string

const (
	AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppNameIDFormatID    AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppNameIDFormat = "id"
	AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppNameIDFormatEmail AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppNameIDFormat = "email"
)

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppNameIDFormat) IsKnown() bool {
	switch r {
	case AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppNameIDFormatID, AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppNameIDFormatEmail:
		return true
	}
	return false
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasApp struct {
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
	AuthType AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppAuthType `json:"auth_type"`
	// The application client id
	ClientID string `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret string                                                                `json:"client_secret"`
	CreatedAt    time.Time                                                             `json:"created_at" format:"date-time"`
	CustomClaims []AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaim `json:"custom_claims"`
	// The OIDC flows supported by this application
	GrantTypes []AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantType `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint.
	GroupFilterRegex         string                                                                           `json:"group_filter_regex"`
	HybridAndImplicitOptions AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppHybridAndImplicitOptions `json:"hybrid_and_implicit_options"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey string `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectUris        []string                                                                    `json:"redirect_uris"`
	RefreshTokenOptions AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppRefreshTokenOptions `json:"refresh_token_options"`
	// Define the user information shared with access, "offline_access" scope will be
	// automatically enabled if refresh tokens are enabled
	Scopes    []AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppScope `json:"scopes"`
	UpdatedAt time.Time                                                       `json:"updated_at" format:"date-time"`
	JSON      accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppJSON    `json:"-"`
}

// accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppJSON contains the JSON
// metadata for the struct
// [AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasApp]
type accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppJSON struct {
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

func (r *AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasApp) implementsAccessAppsSaaSApplicationSaasApp() {
}

// Identifier of the authentication protocol used for the saas app. Required for
// OIDC.
type AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppAuthType string

const (
	AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppAuthTypeSAML AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppAuthType = "saml"
	AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppAuthTypeOidc AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppAuthType = "oidc"
)

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppAuthType) IsKnown() bool {
	switch r {
	case AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppAuthTypeSAML, AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppAuthTypeOidc:
		return true
	}
	return false
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaim struct {
	// The name of the claim.
	Name string `json:"name"`
	// If the claim is required when building an OIDC token.
	Required bool `json:"required"`
	// The scope of the claim.
	Scope  AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsScope  `json:"scope"`
	Source AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSource `json:"source"`
	JSON   accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimJSON    `json:"-"`
}

// accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimJSON contains
// the JSON metadata for the struct
// [AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaim]
type accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimJSON struct {
	Name        apijson.Field
	Required    apijson.Field
	Scope       apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaim) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimJSON) RawJSON() string {
	return r.raw
}

// The scope of the claim.
type AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsScope string

const (
	AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsScopeGroups  AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsScope = "groups"
	AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsScopeProfile AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsScope = "profile"
	AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsScopeEmail   AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsScope = "email"
	AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsScopeOpenid  AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsScope = "openid"
)

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsScope) IsKnown() bool {
	switch r {
	case AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsScopeGroups, AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsScopeProfile, AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsScopeEmail, AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsScopeOpenid:
		return true
	}
	return false
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSource struct {
	// The name of the IdP claim.
	Name string `json:"name"`
	// A mapping from IdP ID to attribute name.
	NameByIdp []AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSourceNameByIdp `json:"name_by_idp"`
	JSON      accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSourceJSON        `json:"-"`
}

// accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSourceJSON
// contains the JSON metadata for the struct
// [AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSource]
type accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSourceJSON struct {
	Name        apijson.Field
	NameByIdp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSourceNameByIdp struct {
	// The UID of the IdP.
	IdpID string `json:"idp_id"`
	// The name of the IdP provided attribute.
	SourceName string                                                                                  `json:"source_name"`
	JSON       accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSourceNameByIdpJSON `json:"-"`
}

// accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSourceNameByIdpJSON
// contains the JSON metadata for the struct
// [AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSourceNameByIdp]
type accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSourceNameByIdpJSON struct {
	IdpID       apijson.Field
	SourceName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSourceNameByIdp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSourceNameByIdpJSON) RawJSON() string {
	return r.raw
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantType string

const (
	AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantTypeAuthorizationCode         AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantType = "authorization_code"
	AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantTypeAuthorizationCodeWithPkce AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantType = "authorization_code_with_pkce"
	AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantTypeRefreshTokens             AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantType = "refresh_tokens"
	AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantTypeHybrid                    AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantType = "hybrid"
	AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantTypeImplicit                  AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantType = "implicit"
)

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantType) IsKnown() bool {
	switch r {
	case AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantTypeAuthorizationCode, AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantTypeAuthorizationCodeWithPkce, AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantTypeRefreshTokens, AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantTypeHybrid, AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantTypeImplicit:
		return true
	}
	return false
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppHybridAndImplicitOptions struct {
	// If an Access Token should be returned from the OIDC Authorization endpoint
	ReturnAccessTokenFromAuthorizationEndpoint bool `json:"return_access_token_from_authorization_endpoint"`
	// If an ID Token should be returned from the OIDC Authorization endpoint
	ReturnIDTokenFromAuthorizationEndpoint bool                                                                                 `json:"return_id_token_from_authorization_endpoint"`
	JSON                                   accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppHybridAndImplicitOptionsJSON `json:"-"`
}

// accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppHybridAndImplicitOptionsJSON
// contains the JSON metadata for the struct
// [AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppHybridAndImplicitOptions]
type accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppHybridAndImplicitOptionsJSON struct {
	ReturnAccessTokenFromAuthorizationEndpoint apijson.Field
	ReturnIDTokenFromAuthorizationEndpoint     apijson.Field
	raw                                        string
	ExtraFields                                map[string]apijson.Field
}

func (r *AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppHybridAndImplicitOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppHybridAndImplicitOptionsJSON) RawJSON() string {
	return r.raw
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppRefreshTokenOptions struct {
	// How long a refresh token will be valid for after creation. Valid units are
	// m,h,d. Must be longer than 1m.
	Lifetime string                                                                          `json:"lifetime"`
	JSON     accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppRefreshTokenOptionsJSON `json:"-"`
}

// accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppRefreshTokenOptionsJSON
// contains the JSON metadata for the struct
// [AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppRefreshTokenOptions]
type accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppRefreshTokenOptionsJSON struct {
	Lifetime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppRefreshTokenOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppRefreshTokenOptionsJSON) RawJSON() string {
	return r.raw
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppScope string

const (
	AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppScopeOpenid  AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppScope = "openid"
	AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppScopeGroups  AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppScope = "groups"
	AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppScopeEmail   AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppScope = "email"
	AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppScopeProfile AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppScope = "profile"
)

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppScope) IsKnown() bool {
	switch r {
	case AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppScopeOpenid, AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppScopeGroups, AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppScopeEmail, AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppScopeProfile:
		return true
	}
	return false
}

// Optional identifier indicating the authentication protocol used for the saas
// app. Required for OIDC. Default if unset is "saml"
type AccessAppsSaaSApplicationSaasAppAuthType string

const (
	AccessAppsSaaSApplicationSaasAppAuthTypeSAML AccessAppsSaaSApplicationSaasAppAuthType = "saml"
	AccessAppsSaaSApplicationSaasAppAuthTypeOidc AccessAppsSaaSApplicationSaasAppAuthType = "oidc"
)

func (r AccessAppsSaaSApplicationSaasAppAuthType) IsKnown() bool {
	switch r {
	case AccessAppsSaaSApplicationSaasAppAuthTypeSAML, AccessAppsSaaSApplicationSaasAppAuthTypeOidc:
		return true
	}
	return false
}

// The format of the name identifier sent to the SaaS application.
type AccessAppsSaaSApplicationSaasAppNameIDFormat string

const (
	AccessAppsSaaSApplicationSaasAppNameIDFormatID    AccessAppsSaaSApplicationSaasAppNameIDFormat = "id"
	AccessAppsSaaSApplicationSaasAppNameIDFormatEmail AccessAppsSaaSApplicationSaasAppNameIDFormat = "email"
)

func (r AccessAppsSaaSApplicationSaasAppNameIDFormat) IsKnown() bool {
	switch r {
	case AccessAppsSaaSApplicationSaasAppNameIDFormatID, AccessAppsSaaSApplicationSaasAppNameIDFormatEmail:
		return true
	}
	return false
}

type AccessAppsBrowserSSHApplication struct {
	// The application type.
	Type string                              `json:"type"`
	JSON accessAppsBrowserSSHApplicationJSON `json:"-"`
	BasicAppResponseProps
	SelfHostedPropsZone
}

// accessAppsBrowserSSHApplicationJSON contains the JSON metadata for the struct
// [AccessAppsBrowserSSHApplication]
type accessAppsBrowserSSHApplicationJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppsBrowserSSHApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsBrowserSSHApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsBrowserSSHApplication) implementsAccessApps() {}

type AccessAppsBrowserVncApplication struct {
	// The application type.
	Type string                              `json:"type"`
	JSON accessAppsBrowserVncApplicationJSON `json:"-"`
	BasicAppResponseProps
	SelfHostedPropsZone
}

// accessAppsBrowserVncApplicationJSON contains the JSON metadata for the struct
// [AccessAppsBrowserVncApplication]
type accessAppsBrowserVncApplicationJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppsBrowserVncApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsBrowserVncApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsBrowserVncApplication) implementsAccessApps() {}

type AccessAppsAppLauncherApplication struct {
	Domain interface{} `json:"domain"`
	Name   interface{} `json:"name"`
	// The application type.
	Type string                               `json:"type"`
	JSON accessAppsAppLauncherApplicationJSON `json:"-"`
	BasicAppResponseProps
	FeatureAppPropsZone
}

// accessAppsAppLauncherApplicationJSON contains the JSON metadata for the struct
// [AccessAppsAppLauncherApplication]
type accessAppsAppLauncherApplicationJSON struct {
	Domain      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppsAppLauncherApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsAppLauncherApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsAppLauncherApplication) implementsAccessApps() {}

type AccessAppsDeviceEnrollmentPermissionsApplication struct {
	Domain interface{} `json:"domain"`
	Name   interface{} `json:"name"`
	// The application type.
	Type string                                               `json:"type"`
	JSON accessAppsDeviceEnrollmentPermissionsApplicationJSON `json:"-"`
	BasicAppResponseProps
	FeatureAppPropsZone
}

// accessAppsDeviceEnrollmentPermissionsApplicationJSON contains the JSON metadata
// for the struct [AccessAppsDeviceEnrollmentPermissionsApplication]
type accessAppsDeviceEnrollmentPermissionsApplicationJSON struct {
	Domain      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppsDeviceEnrollmentPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsDeviceEnrollmentPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsDeviceEnrollmentPermissionsApplication) implementsAccessApps() {}

type AccessAppsBrowserIsolationPermissionsApplication struct {
	Domain interface{} `json:"domain"`
	Name   interface{} `json:"name"`
	// The application type.
	Type string                                               `json:"type"`
	JSON accessAppsBrowserIsolationPermissionsApplicationJSON `json:"-"`
	BasicAppResponseProps
	FeatureAppPropsZone
}

// accessAppsBrowserIsolationPermissionsApplicationJSON contains the JSON metadata
// for the struct [AccessAppsBrowserIsolationPermissionsApplication]
type accessAppsBrowserIsolationPermissionsApplicationJSON struct {
	Domain      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppsBrowserIsolationPermissionsApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsBrowserIsolationPermissionsApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsBrowserIsolationPermissionsApplication) implementsAccessApps() {}

type AccessAppsBookmarkApplication struct {
	// The URL or domain of the bookmark.
	Domain interface{} `json:"domain,required"`
	// The application type.
	Type               string      `json:"type,required"`
	AppLauncherVisible interface{} `json:"app_launcher_visible"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the application.
	Name string                            `json:"name"`
	JSON accessAppsBookmarkApplicationJSON `json:"-"`
	BasicAppResponseProps
}

// accessAppsBookmarkApplicationJSON contains the JSON metadata for the struct
// [AccessAppsBookmarkApplication]
type accessAppsBookmarkApplicationJSON struct {
	Domain             apijson.Field
	Type               apijson.Field
	AppLauncherVisible apijson.Field
	LogoURL            apijson.Field
	Name               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppsBookmarkApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessAppsBookmarkApplicationJSON) RawJSON() string {
	return r.raw
}

func (r AccessAppsBookmarkApplication) implementsAccessApps() {}

// Satisfied by [AccessAppsSelfHostedApplicationParam],
// [AccessAppsSaaSApplicationParam], [AccessAppsBrowserSSHApplicationParam],
// [AccessAppsBrowserVncApplicationParam], [AccessAppsAppLauncherApplicationParam],
// [AccessAppsDeviceEnrollmentPermissionsApplicationParam],
// [AccessAppsBrowserIsolationPermissionsApplicationParam],
// [AccessAppsBookmarkApplicationParam].
type AccessAppsUnionParam interface {
	implementsAccessAppsUnionParam()
}

type AccessAppsSelfHostedApplicationParam struct {
	BasicAppResponsePropsParam
	SelfHostedPropsZoneParam
}

func (r AccessAppsSelfHostedApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppsSelfHostedApplicationParam) implementsAccessAppsUnionParam() {}

type AccessAppsSaaSApplicationParam struct {
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL param.Field[string] `json:"logo_url"`
	// The name of the application.
	Name    param.Field[string]                                     `json:"name"`
	SaasApp param.Field[AccessAppsSaaSApplicationSaasAppUnionParam] `json:"saas_app"`
	// The application type.
	Type param.Field[string] `json:"type"`
	BasicAppResponsePropsParam
}

func (r AccessAppsSaaSApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppsSaaSApplicationParam) implementsAccessAppsUnionParam() {}

type AccessAppsSaaSApplicationSaasAppParam struct {
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
	AuthType param.Field[AccessAppsSaaSApplicationSaasAppAuthType] `json:"auth_type"`
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
	GrantTypes         param.Field[interface{}] `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint.
	GroupFilterRegex         param.Field[string]      `json:"group_filter_regex"`
	HybridAndImplicitOptions param.Field[interface{}] `json:"hybrid_and_implicit_options"`
	// The unique identifier for your SaaS application.
	IdpEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[AccessAppsSaaSApplicationSaasAppNameIDFormat] `json:"name_id_format"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's
	// user identities into a NameID value for its SAML assertion. This expression
	// should evaluate to a singular string. The output of this expression can override
	// the `name_id_format` setting.
	NameIDTransformJsonata param.Field[string] `json:"name_id_transform_jsonata"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey           param.Field[string]      `json:"public_key"`
	RedirectUris        param.Field[interface{}] `json:"redirect_uris"`
	RefreshTokenOptions param.Field[interface{}] `json:"refresh_token_options"`
	Scopes              param.Field[interface{}] `json:"scopes"`
	// A globally unique name for an identity or service provider.
	SpEntityID param.Field[string] `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint param.Field[string]    `json:"sso_endpoint"`
	UpdatedAt   param.Field[time.Time] `json:"updated_at" format:"date-time"`
}

func (r AccessAppsSaaSApplicationSaasAppParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppsSaaSApplicationSaasAppParam) implementsAccessAppsSaaSApplicationSaasAppUnionParam() {
}

// Satisfied by [AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppParam],
// [AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppParam],
// [AccessAppsSaaSApplicationSaasAppParam].
type AccessAppsSaaSApplicationSaasAppUnionParam interface {
	implementsAccessAppsSaaSApplicationSaasAppUnionParam()
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppParam struct {
	// Optional identifier indicating the authentication protocol used for the saas
	// app. Required for OIDC. Default if unset is "saml"
	AuthType param.Field[AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppAuthType] `json:"auth_type"`
	// The service provider's endpoint that is responsible for receiving and parsing a
	// SAML assertion.
	ConsumerServiceURL param.Field[string]                                                                         `json:"consumer_service_url"`
	CreatedAt          param.Field[time.Time]                                                                      `json:"created_at" format:"date-time"`
	CustomAttributes   param.Field[[]AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributeParam] `json:"custom_attributes"`
	// The unique identifier for your SaaS application.
	IdpEntityID param.Field[string] `json:"idp_entity_id"`
	// The format of the name identifier sent to the SaaS application.
	NameIDFormat param.Field[AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppNameIDFormat] `json:"name_id_format"`
	// A [JSONata](https://jsonata.org/) expression that transforms an application's
	// user identities into a NameID value for its SAML assertion. This expression
	// should evaluate to a singular string. The output of this expression can override
	// the `name_id_format` setting.
	NameIDTransformJsonata param.Field[string] `json:"name_id_transform_jsonata"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// A globally unique name for an identity or service provider.
	SpEntityID param.Field[string] `json:"sp_entity_id"`
	// The endpoint where your SaaS application will send login requests.
	SSOEndpoint param.Field[string]    `json:"sso_endpoint"`
	UpdatedAt   param.Field[time.Time] `json:"updated_at" format:"date-time"`
}

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppParam) implementsAccessAppsSaaSApplicationSaasAppUnionParam() {
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributeParam struct {
	// The SAML FriendlyName of the attribute.
	FriendlyName param.Field[string] `json:"friendly_name"`
	// The name of the attribute.
	Name param.Field[string] `json:"name"`
	// A globally unique name for an identity or service provider.
	NameFormat param.Field[AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesNameFormat] `json:"name_format"`
	// If the attribute is required when building a SAML assertion.
	Required param.Field[bool]                                                                                `json:"required"`
	Source   param.Field[AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesSourceParam] `json:"source"`
}

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesSourceParam struct {
	// The name of the IdP attribute.
	Name param.Field[string] `json:"name"`
	// A mapping from IdP ID to attribute name.
	NameByIdp param.Field[map[string]string] `json:"name_by_idp"`
}

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasSAMLSaasAppCustomAttributesSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppParam struct {
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
	AuthType param.Field[AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppAuthType] `json:"auth_type"`
	// The application client id
	ClientID param.Field[string] `json:"client_id"`
	// The application client secret, only returned on POST request.
	ClientSecret param.Field[string]                                                                     `json:"client_secret"`
	CreatedAt    param.Field[time.Time]                                                                  `json:"created_at" format:"date-time"`
	CustomClaims param.Field[[]AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimParam] `json:"custom_claims"`
	// The OIDC flows supported by this application
	GrantTypes param.Field[[]AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppGrantType] `json:"grant_types"`
	// A regex to filter Cloudflare groups returned in ID token and userinfo endpoint.
	GroupFilterRegex         param.Field[string]                                                                                `json:"group_filter_regex"`
	HybridAndImplicitOptions param.Field[AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppHybridAndImplicitOptionsParam] `json:"hybrid_and_implicit_options"`
	// The Access public certificate that will be used to verify your identity.
	PublicKey param.Field[string] `json:"public_key"`
	// The permitted URL's for Cloudflare to return Authorization codes and Access/ID
	// tokens
	RedirectUris        param.Field[[]string]                                                                         `json:"redirect_uris"`
	RefreshTokenOptions param.Field[AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppRefreshTokenOptionsParam] `json:"refresh_token_options"`
	// Define the user information shared with access, "offline_access" scope will be
	// automatically enabled if refresh tokens are enabled
	Scopes    param.Field[[]AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppScope] `json:"scopes"`
	UpdatedAt param.Field[time.Time]                                                       `json:"updated_at" format:"date-time"`
}

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppParam) implementsAccessAppsSaaSApplicationSaasAppUnionParam() {
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimParam struct {
	// The name of the claim.
	Name param.Field[string] `json:"name"`
	// If the claim is required when building an OIDC token.
	Required param.Field[bool] `json:"required"`
	// The scope of the claim.
	Scope  param.Field[AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsScope]       `json:"scope"`
	Source param.Field[AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSourceParam] `json:"source"`
}

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSourceParam struct {
	// The name of the IdP claim.
	Name param.Field[string] `json:"name"`
	// A mapping from IdP ID to attribute name.
	NameByIdp param.Field[[]AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSourceNameByIdpParam] `json:"name_by_idp"`
}

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSourceNameByIdpParam struct {
	// The UID of the IdP.
	IdpID param.Field[string] `json:"idp_id"`
	// The name of the IdP provided attribute.
	SourceName param.Field[string] `json:"source_name"`
}

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppCustomClaimsSourceNameByIdpParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppHybridAndImplicitOptionsParam struct {
	// If an Access Token should be returned from the OIDC Authorization endpoint
	ReturnAccessTokenFromAuthorizationEndpoint param.Field[bool] `json:"return_access_token_from_authorization_endpoint"`
	// If an ID Token should be returned from the OIDC Authorization endpoint
	ReturnIDTokenFromAuthorizationEndpoint param.Field[bool] `json:"return_id_token_from_authorization_endpoint"`
}

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppHybridAndImplicitOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppRefreshTokenOptionsParam struct {
	// How long a refresh token will be valid for after creation. Valid units are
	// m,h,d. Must be longer than 1m.
	Lifetime param.Field[string] `json:"lifetime"`
}

func (r AccessAppsSaaSApplicationSaasAppAccessSchemasOidcSaasAppRefreshTokenOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppsBrowserSSHApplicationParam struct {
	// The application type.
	Type param.Field[string] `json:"type"`
	BasicAppResponsePropsParam
	SelfHostedPropsZoneParam
}

func (r AccessAppsBrowserSSHApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppsBrowserSSHApplicationParam) implementsAccessAppsUnionParam() {}

type AccessAppsBrowserVncApplicationParam struct {
	// The application type.
	Type param.Field[string] `json:"type"`
	BasicAppResponsePropsParam
	SelfHostedPropsZoneParam
}

func (r AccessAppsBrowserVncApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppsBrowserVncApplicationParam) implementsAccessAppsUnionParam() {}

type AccessAppsAppLauncherApplicationParam struct {
	// The application type.
	Type param.Field[string] `json:"type"`
	BasicAppResponsePropsParam
	FeatureAppPropsZoneParam
}

func (r AccessAppsAppLauncherApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppsAppLauncherApplicationParam) implementsAccessAppsUnionParam() {}

type AccessAppsDeviceEnrollmentPermissionsApplicationParam struct {
	// The application type.
	Type param.Field[string] `json:"type"`
	BasicAppResponsePropsParam
	FeatureAppPropsZoneParam
}

func (r AccessAppsDeviceEnrollmentPermissionsApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppsDeviceEnrollmentPermissionsApplicationParam) implementsAccessAppsUnionParam() {}

type AccessAppsBrowserIsolationPermissionsApplicationParam struct {
	// The application type.
	Type param.Field[string] `json:"type"`
	BasicAppResponsePropsParam
	FeatureAppPropsZoneParam
}

func (r AccessAppsBrowserIsolationPermissionsApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppsBrowserIsolationPermissionsApplicationParam) implementsAccessAppsUnionParam() {}

type AccessAppsBookmarkApplicationParam struct {
	// The URL or domain of the bookmark.
	Domain param.Field[interface{}] `json:"domain,required"`
	// The application type.
	Type               param.Field[string]      `json:"type,required"`
	AppLauncherVisible param.Field[interface{}] `json:"app_launcher_visible"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL param.Field[string] `json:"logo_url"`
	// The name of the application.
	Name param.Field[string] `json:"name"`
	BasicAppResponsePropsParam
}

func (r AccessAppsBookmarkApplicationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppsBookmarkApplicationParam) implementsAccessAppsUnionParam() {}

type BasicAppResponseProps struct {
	// UUID
	ID string `json:"id"`
	// Audience tag.
	Aud       string    `json:"aud"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	ScimConfig BasicAppResponsePropsScimConfig `json:"scim_config"`
	UpdatedAt  time.Time                       `json:"updated_at" format:"date-time"`
	JSON       basicAppResponsePropsJSON       `json:"-"`
}

// basicAppResponsePropsJSON contains the JSON metadata for the struct
// [BasicAppResponseProps]
type basicAppResponsePropsJSON struct {
	ID          apijson.Field
	Aud         apijson.Field
	CreatedAt   apijson.Field
	ScimConfig  apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BasicAppResponseProps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basicAppResponsePropsJSON) RawJSON() string {
	return r.raw
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type BasicAppResponsePropsScimConfig struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdpUid string `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteUri string `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication BasicAppResponsePropsScimConfigAuthenticationUnion `json:"authentication"`
	// If false, we propagate DELETE requests to the target application for SCIM
	// resources. If true, we only set `active` to false on the SCIM resource. This is
	// useful because some targets do not support DELETE operations.
	DeactivateOnDelete bool `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled bool `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings []ScimConfigMapping                 `json:"mappings"`
	JSON     basicAppResponsePropsScimConfigJSON `json:"-"`
}

// basicAppResponsePropsScimConfigJSON contains the JSON metadata for the struct
// [BasicAppResponsePropsScimConfig]
type basicAppResponsePropsScimConfigJSON struct {
	IdpUid             apijson.Field
	RemoteUri          apijson.Field
	Authentication     apijson.Field
	DeactivateOnDelete apijson.Field
	Enabled            apijson.Field
	Mappings           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BasicAppResponsePropsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basicAppResponsePropsScimConfigJSON) RawJSON() string {
	return r.raw
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [ScimConfigAuthHTTPBasic],
// [BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerToken],
// [ScimConfigAuthOauth2], [ScimConfigAuthServiceToken] or
// [BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigMultiAuthentication].
type BasicAppResponsePropsScimConfigAuthenticationUnion interface {
	implementsBasicAppResponsePropsScimConfigAuthenticationUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BasicAppResponsePropsScimConfigAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScimConfigAuthHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerToken{}),
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
			Type:       reflect.TypeOf(BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigMultiAuthentication{}),
		},
	)
}

// Attributes for configuring OAuth Bearer Token authentication scheme for SCIM
// provisioning to an application.
type BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerToken struct {
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenScheme `json:"scheme,required"`
	JSON   basicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenJSON   `json:"-"`
}

// basicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenJSON
// contains the JSON metadata for the struct
// [BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerToken]
type basicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenJSON struct {
	Token       apijson.Field
	Scheme      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenJSON) RawJSON() string {
	return r.raw
}

func (r BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerToken) implementsBasicAppResponsePropsScimConfigAuthenticationUnion() {
}

// The authentication scheme to use when making SCIM requests to this application.
type BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenScheme string

const (
	BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenSchemeOauthbearertoken BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenScheme = "oauthbearertoken"
)

func (r BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenScheme) IsKnown() bool {
	switch r {
	case BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenSchemeOauthbearertoken:
		return true
	}
	return false
}

type BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigMultiAuthentication []ScimConfigSingleAuthentication

func (r BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigMultiAuthentication) implementsBasicAppResponsePropsScimConfigAuthenticationUnion() {
}

type BasicAppResponsePropsParam struct {
	// UUID
	ID        param.Field[string]    `json:"id"`
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// Configuration for provisioning to this application via SCIM. This is currently
	// in closed beta.
	ScimConfig param.Field[BasicAppResponsePropsScimConfigParam] `json:"scim_config"`
	UpdatedAt  param.Field[time.Time]                            `json:"updated_at" format:"date-time"`
}

func (r BasicAppResponsePropsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for provisioning to this application via SCIM. This is currently
// in closed beta.
type BasicAppResponsePropsScimConfigParam struct {
	// The UID of the IdP to use as the source for SCIM resources to provision to this
	// application.
	IdpUid param.Field[string] `json:"idp_uid,required"`
	// The base URI for the application's SCIM-compatible API.
	RemoteUri param.Field[string] `json:"remote_uri,required"`
	// Attributes for configuring HTTP Basic authentication scheme for SCIM
	// provisioning to an application.
	Authentication param.Field[BasicAppResponsePropsScimConfigAuthenticationUnionParam] `json:"authentication"`
	// If false, we propagate DELETE requests to the target application for SCIM
	// resources. If true, we only set `active` to false on the SCIM resource. This is
	// useful because some targets do not support DELETE operations.
	DeactivateOnDelete param.Field[bool] `json:"deactivate_on_delete"`
	// Whether SCIM provisioning is turned on for this application.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of mappings to apply to SCIM resources before provisioning them in this
	// application. These can transform or filter the resources to be provisioned.
	Mappings param.Field[[]ScimConfigMappingParam] `json:"mappings"`
}

func (r BasicAppResponsePropsScimConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type BasicAppResponsePropsScimConfigAuthenticationParam struct {
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[BasicAppResponsePropsScimConfigAuthenticationScheme] `json:"scheme,required"`
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

func (r BasicAppResponsePropsScimConfigAuthenticationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BasicAppResponsePropsScimConfigAuthenticationParam) implementsBasicAppResponsePropsScimConfigAuthenticationUnionParam() {
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [ScimConfigAuthHTTPBasicParam],
// [BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenParam],
// [ScimConfigAuthOauth2Param], [ScimConfigAuthServiceTokenParam],
// [BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigMultiAuthenticationParam],
// [BasicAppResponsePropsScimConfigAuthenticationParam].
type BasicAppResponsePropsScimConfigAuthenticationUnionParam interface {
	implementsBasicAppResponsePropsScimConfigAuthenticationUnionParam()
}

// Attributes for configuring OAuth Bearer Token authentication scheme for SCIM
// provisioning to an application.
type BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenParam struct {
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenScheme] `json:"scheme,required"`
}

func (r BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenParam) implementsBasicAppResponsePropsScimConfigAuthenticationUnionParam() {
}

type BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigMultiAuthenticationParam []ScimConfigSingleAuthenticationUnionParam

func (r BasicAppResponsePropsScimConfigAuthenticationAccessSchemasScimConfigMultiAuthenticationParam) implementsBasicAppResponsePropsScimConfigAuthenticationUnionParam() {
}

// The authentication scheme to use when making SCIM requests to this application.
type BasicAppResponsePropsScimConfigAuthenticationScheme string

const (
	BasicAppResponsePropsScimConfigAuthenticationSchemeHttpbasic          BasicAppResponsePropsScimConfigAuthenticationScheme = "httpbasic"
	BasicAppResponsePropsScimConfigAuthenticationSchemeOauthbearertoken   BasicAppResponsePropsScimConfigAuthenticationScheme = "oauthbearertoken"
	BasicAppResponsePropsScimConfigAuthenticationSchemeOauth2             BasicAppResponsePropsScimConfigAuthenticationScheme = "oauth2"
	BasicAppResponsePropsScimConfigAuthenticationSchemeAccessServiceToken BasicAppResponsePropsScimConfigAuthenticationScheme = "access_service_token"
)

func (r BasicAppResponsePropsScimConfigAuthenticationScheme) IsKnown() bool {
	switch r {
	case BasicAppResponsePropsScimConfigAuthenticationSchemeHttpbasic, BasicAppResponsePropsScimConfigAuthenticationSchemeOauthbearertoken, BasicAppResponsePropsScimConfigAuthenticationSchemeOauth2, BasicAppResponsePropsScimConfigAuthenticationSchemeAccessServiceToken:
		return true
	}
	return false
}

type FeatureAppPropsZone struct {
	// The application type.
	Type FeatureAppPropsZoneType `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps []string `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool `json:"auto_redirect_to_identity"`
	// The domain and path that Access will secure.
	Domain string `json:"domain"`
	// The name of the application.
	Name string `json:"name"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration string                  `json:"session_duration"`
	JSON            featureAppPropsZoneJSON `json:"-"`
}

// featureAppPropsZoneJSON contains the JSON metadata for the struct
// [FeatureAppPropsZone]
type featureAppPropsZoneJSON struct {
	Type                   apijson.Field
	AllowedIdps            apijson.Field
	AutoRedirectToIdentity apijson.Field
	Domain                 apijson.Field
	Name                   apijson.Field
	SessionDuration        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *FeatureAppPropsZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r featureAppPropsZoneJSON) RawJSON() string {
	return r.raw
}

// The application type.
type FeatureAppPropsZoneType string

const (
	FeatureAppPropsZoneTypeSelfHosted  FeatureAppPropsZoneType = "self_hosted"
	FeatureAppPropsZoneTypeSaas        FeatureAppPropsZoneType = "saas"
	FeatureAppPropsZoneTypeSSH         FeatureAppPropsZoneType = "ssh"
	FeatureAppPropsZoneTypeVnc         FeatureAppPropsZoneType = "vnc"
	FeatureAppPropsZoneTypeAppLauncher FeatureAppPropsZoneType = "app_launcher"
	FeatureAppPropsZoneTypeWarp        FeatureAppPropsZoneType = "warp"
	FeatureAppPropsZoneTypeBiso        FeatureAppPropsZoneType = "biso"
	FeatureAppPropsZoneTypeBookmark    FeatureAppPropsZoneType = "bookmark"
	FeatureAppPropsZoneTypeDashSSO     FeatureAppPropsZoneType = "dash_sso"
)

func (r FeatureAppPropsZoneType) IsKnown() bool {
	switch r {
	case FeatureAppPropsZoneTypeSelfHosted, FeatureAppPropsZoneTypeSaas, FeatureAppPropsZoneTypeSSH, FeatureAppPropsZoneTypeVnc, FeatureAppPropsZoneTypeAppLauncher, FeatureAppPropsZoneTypeWarp, FeatureAppPropsZoneTypeBiso, FeatureAppPropsZoneTypeBookmark, FeatureAppPropsZoneTypeDashSSO:
		return true
	}
	return false
}

type FeatureAppPropsZoneParam struct {
	// The application type.
	Type param.Field[FeatureAppPropsZoneType] `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// The domain and path that Access will secure.
	Domain param.Field[string] `json:"domain"`
	// The name of the application.
	Name param.Field[string] `json:"name"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r FeatureAppPropsZoneParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type ScimConfigSingleAuthentication struct {
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme ScimConfigSingleAuthenticationScheme `json:"scheme,required"`
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
	User  string                             `json:"user"`
	JSON  scimConfigSingleAuthenticationJSON `json:"-"`
	union ScimConfigSingleAuthenticationUnion
}

// scimConfigSingleAuthenticationJSON contains the JSON metadata for the struct
// [ScimConfigSingleAuthentication]
type scimConfigSingleAuthenticationJSON struct {
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

func (r scimConfigSingleAuthenticationJSON) RawJSON() string {
	return r.raw
}

func (r *ScimConfigSingleAuthentication) UnmarshalJSON(data []byte) (err error) {
	*r = ScimConfigSingleAuthentication{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScimConfigSingleAuthenticationUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [ScimConfigAuthHTTPBasic],
// [ScimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerToken],
// [ScimConfigAuthOauth2], [ScimConfigAuthServiceToken].
func (r ScimConfigSingleAuthentication) AsUnion() ScimConfigSingleAuthenticationUnion {
	return r.union
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Union satisfied by [ScimConfigAuthHTTPBasic],
// [ScimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerToken],
// [ScimConfigAuthOauth2] or [ScimConfigAuthServiceToken].
type ScimConfigSingleAuthenticationUnion interface {
	implementsScimConfigSingleAuthentication()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScimConfigSingleAuthenticationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScimConfigAuthHTTPBasic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerToken{}),
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
type ScimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerToken struct {
	// Token used to authenticate with the remote SCIM service.
	Token string `json:"token,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme ScimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenScheme `json:"scheme,required"`
	JSON   scimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenJSON   `json:"-"`
}

// scimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenJSON
// contains the JSON metadata for the struct
// [ScimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerToken]
type scimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenJSON struct {
	Token       apijson.Field
	Scheme      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenJSON) RawJSON() string {
	return r.raw
}

func (r ScimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerToken) implementsScimConfigSingleAuthentication() {
}

// The authentication scheme to use when making SCIM requests to this application.
type ScimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenScheme string

const (
	ScimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenSchemeOauthbearertoken ScimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenScheme = "oauthbearertoken"
)

func (r ScimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenScheme) IsKnown() bool {
	switch r {
	case ScimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenSchemeOauthbearertoken:
		return true
	}
	return false
}

// The authentication scheme to use when making SCIM requests to this application.
type ScimConfigSingleAuthenticationScheme string

const (
	ScimConfigSingleAuthenticationSchemeHttpbasic          ScimConfigSingleAuthenticationScheme = "httpbasic"
	ScimConfigSingleAuthenticationSchemeOauthbearertoken   ScimConfigSingleAuthenticationScheme = "oauthbearertoken"
	ScimConfigSingleAuthenticationSchemeOauth2             ScimConfigSingleAuthenticationScheme = "oauth2"
	ScimConfigSingleAuthenticationSchemeAccessServiceToken ScimConfigSingleAuthenticationScheme = "access_service_token"
)

func (r ScimConfigSingleAuthenticationScheme) IsKnown() bool {
	switch r {
	case ScimConfigSingleAuthenticationSchemeHttpbasic, ScimConfigSingleAuthenticationSchemeOauthbearertoken, ScimConfigSingleAuthenticationSchemeOauth2, ScimConfigSingleAuthenticationSchemeAccessServiceToken:
		return true
	}
	return false
}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
type ScimConfigSingleAuthenticationParam struct {
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[ScimConfigSingleAuthenticationScheme] `json:"scheme,required"`
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

func (r ScimConfigSingleAuthenticationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScimConfigSingleAuthenticationParam) implementsScimConfigSingleAuthenticationUnionParam() {}

// Attributes for configuring HTTP Basic authentication scheme for SCIM
// provisioning to an application.
//
// Satisfied by [ScimConfigAuthHTTPBasicParam],
// [ScimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenParam],
// [ScimConfigAuthOauth2Param], [ScimConfigAuthServiceTokenParam],
// [ScimConfigSingleAuthenticationParam].
type ScimConfigSingleAuthenticationUnionParam interface {
	implementsScimConfigSingleAuthenticationUnionParam()
}

// Attributes for configuring OAuth Bearer Token authentication scheme for SCIM
// provisioning to an application.
type ScimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenParam struct {
	// Token used to authenticate with the remote SCIM service.
	Token param.Field[string] `json:"token,required"`
	// The authentication scheme to use when making SCIM requests to this application.
	Scheme param.Field[ScimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenScheme] `json:"scheme,required"`
}

func (r ScimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScimConfigSingleAuthenticationAccessSchemasScimConfigAuthenticationOAuthBearerTokenParam) implementsScimConfigSingleAuthenticationUnionParam() {
}

type SelfHostedPropsZone struct {
	// The domain and path that Access will secure.
	Domain string `json:"domain,required"`
	// The application type.
	Type string `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps []string `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity bool                           `json:"auto_redirect_to_identity"`
	CorsHeaders            SelfHostedPropsZoneCorsHeaders `json:"cors_headers"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage string `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application.
	CustomDenyURL string `json:"custom_deny_url"`
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
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute string `json:"same_site_cookie_attribute"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect bool `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration string `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial bool                    `json:"skip_interstitial"`
	JSON             selfHostedPropsZoneJSON `json:"-"`
}

// selfHostedPropsZoneJSON contains the JSON metadata for the struct
// [SelfHostedPropsZone]
type selfHostedPropsZoneJSON struct {
	Domain                  apijson.Field
	Type                    apijson.Field
	AllowedIdps             apijson.Field
	AppLauncherVisible      apijson.Field
	AutoRedirectToIdentity  apijson.Field
	CorsHeaders             apijson.Field
	CustomDenyMessage       apijson.Field
	CustomDenyURL           apijson.Field
	EnableBindingCookie     apijson.Field
	HTTPOnlyCookieAttribute apijson.Field
	LogoURL                 apijson.Field
	Name                    apijson.Field
	OptionsPreflightBypass  apijson.Field
	SameSiteCookieAttribute apijson.Field
	ServiceAuth401Redirect  apijson.Field
	SessionDuration         apijson.Field
	SkipInterstitial        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SelfHostedPropsZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r selfHostedPropsZoneJSON) RawJSON() string {
	return r.raw
}

type SelfHostedPropsZoneCorsHeaders struct {
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
	AllowedHeaders []interface{} `json:"allowed_headers"`
	// Allowed HTTP request methods.
	AllowedMethods []AllowedMethodsItem `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins []interface{} `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge float64                            `json:"max_age"`
	JSON   selfHostedPropsZoneCorsHeadersJSON `json:"-"`
}

// selfHostedPropsZoneCorsHeadersJSON contains the JSON metadata for the struct
// [SelfHostedPropsZoneCorsHeaders]
type selfHostedPropsZoneCorsHeadersJSON struct {
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

func (r *SelfHostedPropsZoneCorsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r selfHostedPropsZoneCorsHeadersJSON) RawJSON() string {
	return r.raw
}

type SelfHostedPropsZoneParam struct {
	// The domain and path that Access will secure.
	Domain param.Field[string] `json:"domain,required"`
	// The application type.
	Type param.Field[string] `json:"type,required"`
	// The identity providers your users can select when connecting to this
	// application. Defaults to all IdPs configured in your account.
	AllowedIdps param.Field[[]string] `json:"allowed_idps"`
	// Displays the application in the App Launcher.
	AppLauncherVisible param.Field[bool] `json:"app_launcher_visible"`
	// When set to `true`, users skip the identity provider selection step during
	// login. You must specify only one identity provider in allowed_idps.
	AutoRedirectToIdentity param.Field[bool]                                `json:"auto_redirect_to_identity"`
	CorsHeaders            param.Field[SelfHostedPropsZoneCorsHeadersParam] `json:"cors_headers"`
	// The custom error message shown to a user when they are denied access to the
	// application.
	CustomDenyMessage param.Field[string] `json:"custom_deny_message"`
	// The custom URL a user is redirected to when they are denied access to the
	// application.
	CustomDenyURL param.Field[string] `json:"custom_deny_url"`
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
	// Sets the SameSite cookie setting, which provides increased security against CSRF
	// attacks.
	SameSiteCookieAttribute param.Field[string] `json:"same_site_cookie_attribute"`
	// Returns a 401 status code when the request is blocked by a Service Auth policy.
	ServiceAuth401Redirect param.Field[bool] `json:"service_auth_401_redirect"`
	// The amount of time that tokens issued for this application will be valid. Must
	// be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms,
	// s, m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial param.Field[bool] `json:"skip_interstitial"`
}

func (r SelfHostedPropsZoneParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SelfHostedPropsZoneCorsHeadersParam struct {
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
	AllowedHeaders param.Field[[]interface{}] `json:"allowed_headers"`
	// Allowed HTTP request methods.
	AllowedMethods param.Field[[]AllowedMethodsItem] `json:"allowed_methods"`
	// Allowed origins.
	AllowedOrigins param.Field[[]interface{}] `json:"allowed_origins"`
	// The maximum number of seconds the results of a preflight request can be cached.
	MaxAge param.Field[float64] `json:"max_age"`
}

func (r SelfHostedPropsZoneCorsHeadersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SingleResponseAppZone struct {
	Result AccessApps                `json:"result"`
	JSON   singleResponseAppZoneJSON `json:"-"`
	APIResponseSingleAccess
}

// singleResponseAppZoneJSON contains the JSON metadata for the struct
// [SingleResponseAppZone]
type singleResponseAppZoneJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseAppZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseAppZoneJSON) RawJSON() string {
	return r.raw
}

type ZoneAccessAppNewResponse struct {
	Result AccessApps                   `json:"result"`
	JSON   zoneAccessAppNewResponseJSON `json:"-"`
	SingleResponseAppZone
}

// zoneAccessAppNewResponseJSON contains the JSON metadata for the struct
// [ZoneAccessAppNewResponse]
type zoneAccessAppNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessAppNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAccessAppNewResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAccessAppUpdateResponse struct {
	Result AccessApps                      `json:"result"`
	JSON   zoneAccessAppUpdateResponseJSON `json:"-"`
	SingleResponseAppZone
}

// zoneAccessAppUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneAccessAppUpdateResponse]
type zoneAccessAppUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessAppUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAccessAppUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAccessAppListResponse struct {
	Result []AccessApps                  `json:"result"`
	JSON   zoneAccessAppListResponseJSON `json:"-"`
	APIResponseCollectionAccess
}

// zoneAccessAppListResponseJSON contains the JSON metadata for the struct
// [ZoneAccessAppListResponse]
type zoneAccessAppListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessAppListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAccessAppListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAccessAppUserPolicyChecksResponse struct {
	Result ZoneAccessAppUserPolicyChecksResponseResult `json:"result"`
	JSON   zoneAccessAppUserPolicyChecksResponseJSON   `json:"-"`
	APIResponseSingleAccess
}

// zoneAccessAppUserPolicyChecksResponseJSON contains the JSON metadata for the
// struct [ZoneAccessAppUserPolicyChecksResponse]
type zoneAccessAppUserPolicyChecksResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessAppUserPolicyChecksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAccessAppUserPolicyChecksResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAccessAppUserPolicyChecksResponseResult struct {
	AppState     ZoneAccessAppUserPolicyChecksResponseResultAppState     `json:"app_state"`
	UserIdentity ZoneAccessAppUserPolicyChecksResponseResultUserIdentity `json:"user_identity"`
	JSON         zoneAccessAppUserPolicyChecksResponseResultJSON         `json:"-"`
}

// zoneAccessAppUserPolicyChecksResponseResultJSON contains the JSON metadata for
// the struct [ZoneAccessAppUserPolicyChecksResponseResult]
type zoneAccessAppUserPolicyChecksResponseResultJSON struct {
	AppState     apijson.Field
	UserIdentity apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessAppUserPolicyChecksResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAccessAppUserPolicyChecksResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneAccessAppUserPolicyChecksResponseResultAppState struct {
	// UUID
	AppUid   string                                                  `json:"app_uid"`
	Aud      string                                                  `json:"aud"`
	Hostname string                                                  `json:"hostname"`
	Name     string                                                  `json:"name"`
	Policies []interface{}                                           `json:"policies"`
	Status   string                                                  `json:"status"`
	JSON     zoneAccessAppUserPolicyChecksResponseResultAppStateJSON `json:"-"`
}

// zoneAccessAppUserPolicyChecksResponseResultAppStateJSON contains the JSON
// metadata for the struct [ZoneAccessAppUserPolicyChecksResponseResultAppState]
type zoneAccessAppUserPolicyChecksResponseResultAppStateJSON struct {
	AppUid      apijson.Field
	Aud         apijson.Field
	Hostname    apijson.Field
	Name        apijson.Field
	Policies    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessAppUserPolicyChecksResponseResultAppState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAccessAppUserPolicyChecksResponseResultAppStateJSON) RawJSON() string {
	return r.raw
}

type ZoneAccessAppUserPolicyChecksResponseResultUserIdentity struct {
	ID             string                                                     `json:"id"`
	AccountID      string                                                     `json:"account_id"`
	DeviceSessions interface{}                                                `json:"device_sessions"`
	Email          string                                                     `json:"email"`
	Geo            ZoneAccessAppUserPolicyChecksResponseResultUserIdentityGeo `json:"geo"`
	Iat            int64                                                      `json:"iat"`
	IsGateway      bool                                                       `json:"is_gateway"`
	IsWarp         bool                                                       `json:"is_warp"`
	Name           string                                                     `json:"name"`
	// UUID
	UserUuid string                                                      `json:"user_uuid"`
	Version  int64                                                       `json:"version"`
	JSON     zoneAccessAppUserPolicyChecksResponseResultUserIdentityJSON `json:"-"`
}

// zoneAccessAppUserPolicyChecksResponseResultUserIdentityJSON contains the JSON
// metadata for the struct
// [ZoneAccessAppUserPolicyChecksResponseResultUserIdentity]
type zoneAccessAppUserPolicyChecksResponseResultUserIdentityJSON struct {
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

func (r *ZoneAccessAppUserPolicyChecksResponseResultUserIdentity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAccessAppUserPolicyChecksResponseResultUserIdentityJSON) RawJSON() string {
	return r.raw
}

type ZoneAccessAppUserPolicyChecksResponseResultUserIdentityGeo struct {
	Country string                                                         `json:"country"`
	JSON    zoneAccessAppUserPolicyChecksResponseResultUserIdentityGeoJSON `json:"-"`
}

// zoneAccessAppUserPolicyChecksResponseResultUserIdentityGeoJSON contains the JSON
// metadata for the struct
// [ZoneAccessAppUserPolicyChecksResponseResultUserIdentityGeo]
type zoneAccessAppUserPolicyChecksResponseResultUserIdentityGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessAppUserPolicyChecksResponseResultUserIdentityGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAccessAppUserPolicyChecksResponseResultUserIdentityGeoJSON) RawJSON() string {
	return r.raw
}

type ZoneAccessAppNewParams struct {
	AccessApps AccessAppsUnionParam `json:"access_apps,required"`
}

func (r ZoneAccessAppNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AccessApps)
}

type ZoneAccessAppUpdateParams struct {
	AccessApps AccessAppsUnionParam `json:"access_apps,required"`
}

func (r ZoneAccessAppUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AccessApps)
}
