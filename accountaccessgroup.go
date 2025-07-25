// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountAccessGroupService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessGroupService] method instead.
type AccountAccessGroupService struct {
	Options []option.RequestOption
}

// NewAccountAccessGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessGroupService(opts ...option.RequestOption) (r *AccountAccessGroupService) {
	r = &AccountAccessGroupService{}
	r.Options = opts
	return
}

// Matches an Access group.
type AccessRule struct {
	// This field can have the runtime type of
	// [AccessRuleAccessAnyValidServiceTokenRuleAnyValidServiceToken].
	AnyValidServiceToken interface{} `json:"any_valid_service_token"`
	// This field can have the runtime type of
	// [AccessRuleAccessAuthContextRuleAuthContext].
	AuthContext interface{} `json:"auth_context"`
	// This field can have the runtime type of
	// [AccessRuleAccessAuthenticationMethodRuleAuthMethod].
	AuthMethod interface{} `json:"auth_method"`
	// This field can have the runtime type of [AccessRuleAccessAzureGroupRuleAzureAd].
	AzureAd interface{} `json:"azureAD"`
	// This field can have the runtime type of
	// [AccessRuleAccessCertificateRuleCertificate].
	Certificate interface{} `json:"certificate"`
	// This field can have the runtime type of
	// [AccessRuleAccessCommonNameRuleCommonName].
	CommonName interface{} `json:"common_name"`
	// This field can have the runtime type of
	// [AccessRuleAccessDevicePostureRuleDevicePosture].
	DevicePosture interface{} `json:"device_posture"`
	// This field can have the runtime type of [AccessRuleAccessEmailRuleEmail].
	Email interface{} `json:"email"`
	// This field can have the runtime type of [AccessRuleAccessDomainRuleEmailDomain].
	EmailDomain interface{} `json:"email_domain"`
	// This field can have the runtime type of
	// [AccessRuleAccessEmailListRuleEmailList].
	EmailList interface{} `json:"email_list"`
	// This field can have the runtime type of [AccessRuleAccessEveryoneRuleEveryone].
	Everyone interface{} `json:"everyone"`
	// This field can have the runtime type of
	// [AccessRuleAccessExternalEvaluationRuleExternalEvaluation].
	ExternalEvaluation interface{} `json:"external_evaluation"`
	// This field can have the runtime type of [AccessRuleAccessCountryRuleGeo].
	Geo interface{} `json:"geo"`
	// This field can have the runtime type of
	// [AccessRuleAccessGitHubOrganizationRuleGitHubOrganization].
	GitHubOrganization interface{} `json:"github-organization"`
	// This field can have the runtime type of [AccessRuleAccessAccessGroupRuleGroup].
	Group interface{} `json:"group"`
	// This field can have the runtime type of [AccessRuleAccessGsuiteGroupRuleGsuite].
	Gsuite interface{} `json:"gsuite"`
	// This field can have the runtime type of [AccessRuleAccessIPRuleIP].
	IP interface{} `json:"ip"`
	// This field can have the runtime type of [AccessRuleAccessIPListRuleIPList].
	IPList interface{} `json:"ip_list"`
	// This field can have the runtime type of
	// [AccessRuleAccessLinkedAppTokenRuleLinkedAppToken].
	LinkedAppToken interface{} `json:"linked_app_token"`
	// This field can have the runtime type of
	// [AccessRuleAccessLoginMethodRuleLoginMethod].
	LoginMethod interface{} `json:"login_method"`
	// This field can have the runtime type of [AccessRuleAccessOidcClaimRuleOidc].
	Oidc interface{} `json:"oidc"`
	// This field can have the runtime type of [AccessRuleAccessOktaGroupRuleOkta].
	Okta interface{} `json:"okta"`
	// This field can have the runtime type of [AccessRuleAccessSAMLGroupRuleSAML].
	SAML interface{} `json:"saml"`
	// This field can have the runtime type of
	// [AccessRuleAccessServiceTokenRuleServiceToken].
	ServiceToken interface{}    `json:"service_token"`
	JSON         accessRuleJSON `json:"-"`
	union        AccessRuleUnion
}

// accessRuleJSON contains the JSON metadata for the struct [AccessRule]
type accessRuleJSON struct {
	AnyValidServiceToken apijson.Field
	AuthContext          apijson.Field
	AuthMethod           apijson.Field
	AzureAd              apijson.Field
	Certificate          apijson.Field
	CommonName           apijson.Field
	DevicePosture        apijson.Field
	Email                apijson.Field
	EmailDomain          apijson.Field
	EmailList            apijson.Field
	Everyone             apijson.Field
	ExternalEvaluation   apijson.Field
	Geo                  apijson.Field
	GitHubOrganization   apijson.Field
	Group                apijson.Field
	Gsuite               apijson.Field
	IP                   apijson.Field
	IPList               apijson.Field
	LinkedAppToken       apijson.Field
	LoginMethod          apijson.Field
	Oidc                 apijson.Field
	Okta                 apijson.Field
	SAML                 apijson.Field
	ServiceToken         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r accessRuleJSON) RawJSON() string {
	return r.raw
}

func (r *AccessRule) UnmarshalJSON(data []byte) (err error) {
	*r = AccessRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessRuleUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [AccessRuleAccessAccessGroupRule],
// [AccessRuleAccessAnyValidServiceTokenRule], [AccessRuleAccessAuthContextRule],
// [AccessRuleAccessAuthenticationMethodRule], [AccessRuleAccessAzureGroupRule],
// [AccessRuleAccessCertificateRule], [AccessRuleAccessCommonNameRule],
// [AccessRuleAccessCountryRule], [AccessRuleAccessDevicePostureRule],
// [AccessRuleAccessDomainRule], [AccessRuleAccessEmailListRule],
// [AccessRuleAccessEmailRule], [AccessRuleAccessEveryoneRule],
// [AccessRuleAccessExternalEvaluationRule],
// [AccessRuleAccessGitHubOrganizationRule], [AccessRuleAccessGsuiteGroupRule],
// [AccessRuleAccessLoginMethodRule], [AccessRuleAccessIPListRule],
// [AccessRuleAccessIPRule], [AccessRuleAccessOktaGroupRule],
// [AccessRuleAccessSAMLGroupRule], [AccessRuleAccessOidcClaimRule],
// [AccessRuleAccessServiceTokenRule], [AccessRuleAccessLinkedAppTokenRule].
func (r AccessRule) AsUnion() AccessRuleUnion {
	return r.union
}

// Matches an Access group.
//
// Union satisfied by [AccessRuleAccessAccessGroupRule],
// [AccessRuleAccessAnyValidServiceTokenRule], [AccessRuleAccessAuthContextRule],
// [AccessRuleAccessAuthenticationMethodRule], [AccessRuleAccessAzureGroupRule],
// [AccessRuleAccessCertificateRule], [AccessRuleAccessCommonNameRule],
// [AccessRuleAccessCountryRule], [AccessRuleAccessDevicePostureRule],
// [AccessRuleAccessDomainRule], [AccessRuleAccessEmailListRule],
// [AccessRuleAccessEmailRule], [AccessRuleAccessEveryoneRule],
// [AccessRuleAccessExternalEvaluationRule],
// [AccessRuleAccessGitHubOrganizationRule], [AccessRuleAccessGsuiteGroupRule],
// [AccessRuleAccessLoginMethodRule], [AccessRuleAccessIPListRule],
// [AccessRuleAccessIPRule], [AccessRuleAccessOktaGroupRule],
// [AccessRuleAccessSAMLGroupRule], [AccessRuleAccessOidcClaimRule],
// [AccessRuleAccessServiceTokenRule] or [AccessRuleAccessLinkedAppTokenRule].
type AccessRuleUnion interface {
	implementsAccessRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessRuleUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessAuthContextRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessCommonNameRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessDevicePostureRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessLoginMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessSAMLGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessOidcClaimRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessLinkedAppTokenRule{}),
		},
	)
}

// Matches an Access group.
type AccessRuleAccessAccessGroupRule struct {
	Group AccessRuleAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessRuleAccessAccessGroupRuleJSON  `json:"-"`
}

// accessRuleAccessAccessGroupRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessAccessGroupRule]
type accessRuleAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessAccessGroupRule) implementsAccessRule() {}

type AccessRuleAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                   `json:"id,required"`
	JSON accessRuleAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessRuleAccessAccessGroupRuleGroupJSON contains the JSON metadata for the
// struct [AccessRuleAccessAccessGroupRuleGroup]
type accessRuleAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type AccessRuleAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken AccessRuleAccessAnyValidServiceTokenRuleAnyValidServiceToken `json:"any_valid_service_token,required"`
	JSON                 accessRuleAccessAnyValidServiceTokenRuleJSON                 `json:"-"`
}

// accessRuleAccessAnyValidServiceTokenRuleJSON contains the JSON metadata for the
// struct [AccessRuleAccessAnyValidServiceTokenRule]
type accessRuleAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessRuleAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessAnyValidServiceTokenRule) implementsAccessRule() {}

// An empty object which matches on all service tokens.
type AccessRuleAccessAnyValidServiceTokenRuleAnyValidServiceToken struct {
	JSON accessRuleAccessAnyValidServiceTokenRuleAnyValidServiceTokenJSON `json:"-"`
}

// accessRuleAccessAnyValidServiceTokenRuleAnyValidServiceTokenJSON contains the
// JSON metadata for the struct
// [AccessRuleAccessAnyValidServiceTokenRuleAnyValidServiceToken]
type accessRuleAccessAnyValidServiceTokenRuleAnyValidServiceTokenJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessAnyValidServiceTokenRuleAnyValidServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAnyValidServiceTokenRuleAnyValidServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure Authentication Context. Requires an Azure identity provider.
type AccessRuleAccessAuthContextRule struct {
	AuthContext AccessRuleAccessAuthContextRuleAuthContext `json:"auth_context,required"`
	JSON        accessRuleAccessAuthContextRuleJSON        `json:"-"`
}

// accessRuleAccessAuthContextRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessAuthContextRule]
type accessRuleAccessAuthContextRuleJSON struct {
	AuthContext apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessAuthContextRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAuthContextRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessAuthContextRule) implementsAccessRule() {}

type AccessRuleAccessAuthContextRuleAuthContext struct {
	// The ID of an Authentication context.
	ID string `json:"id,required"`
	// The ACID of an Authentication context.
	AcID string `json:"ac_id,required"`
	// The ID of your Azure identity provider.
	IdentityProviderID string                                         `json:"identity_provider_id,required"`
	JSON               accessRuleAccessAuthContextRuleAuthContextJSON `json:"-"`
}

// accessRuleAccessAuthContextRuleAuthContextJSON contains the JSON metadata for
// the struct [AccessRuleAccessAuthContextRuleAuthContext]
type accessRuleAccessAuthContextRuleAuthContextJSON struct {
	ID                 apijson.Field
	AcID               apijson.Field
	IdentityProviderID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessRuleAccessAuthContextRuleAuthContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAuthContextRuleAuthContextJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type AccessRuleAccessAuthenticationMethodRule struct {
	AuthMethod AccessRuleAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessRuleAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessRuleAccessAuthenticationMethodRuleJSON contains the JSON metadata for the
// struct [AccessRuleAccessAuthenticationMethodRule]
type accessRuleAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessAuthenticationMethodRule) implementsAccessRule() {}

type AccessRuleAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method
	// https://datatracker.ietf.org/doc/html/rfc8176#section-2.
	AuthMethod string                                                 `json:"auth_method,required"`
	JSON       accessRuleAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessRuleAccessAuthenticationMethodRuleAuthMethodJSON contains the JSON
// metadata for the struct [AccessRuleAccessAuthenticationMethodRuleAuthMethod]
type accessRuleAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessRuleAccessAzureGroupRule struct {
	AzureAd AccessRuleAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessRuleAccessAzureGroupRuleJSON    `json:"-"`
}

// accessRuleAccessAzureGroupRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessAzureGroupRule]
type accessRuleAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessAzureGroupRule) implementsAccessRule() {}

type AccessRuleAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	IdentityProviderID string                                    `json:"identity_provider_id,required"`
	JSON               accessRuleAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessRuleAccessAzureGroupRuleAzureAdJSON contains the JSON metadata for the
// struct [AccessRuleAccessAzureGroupRuleAzureAd]
type accessRuleAccessAzureGroupRuleAzureAdJSON struct {
	ID                 apijson.Field
	IdentityProviderID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessRuleAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type AccessRuleAccessCertificateRule struct {
	Certificate AccessRuleAccessCertificateRuleCertificate `json:"certificate,required"`
	JSON        accessRuleAccessCertificateRuleJSON        `json:"-"`
}

// accessRuleAccessCertificateRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessCertificateRule]
type accessRuleAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessCertificateRule) implementsAccessRule() {}

type AccessRuleAccessCertificateRuleCertificate struct {
	JSON accessRuleAccessCertificateRuleCertificateJSON `json:"-"`
}

// accessRuleAccessCertificateRuleCertificateJSON contains the JSON metadata for
// the struct [AccessRuleAccessCertificateRuleCertificate]
type accessRuleAccessCertificateRuleCertificateJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessCertificateRuleCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessCertificateRuleCertificateJSON) RawJSON() string {
	return r.raw
}

// Matches a specific common name.
type AccessRuleAccessCommonNameRule struct {
	CommonName AccessRuleAccessCommonNameRuleCommonName `json:"common_name,required"`
	JSON       accessRuleAccessCommonNameRuleJSON       `json:"-"`
}

// accessRuleAccessCommonNameRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessCommonNameRule]
type accessRuleAccessCommonNameRuleJSON struct {
	CommonName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessCommonNameRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessCommonNameRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessCommonNameRule) implementsAccessRule() {}

type AccessRuleAccessCommonNameRuleCommonName struct {
	// The common name to match.
	CommonName string                                       `json:"common_name,required"`
	JSON       accessRuleAccessCommonNameRuleCommonNameJSON `json:"-"`
}

// accessRuleAccessCommonNameRuleCommonNameJSON contains the JSON metadata for the
// struct [AccessRuleAccessCommonNameRuleCommonName]
type accessRuleAccessCommonNameRuleCommonNameJSON struct {
	CommonName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessCommonNameRuleCommonName) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessCommonNameRuleCommonNameJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type AccessRuleAccessCountryRule struct {
	Geo  AccessRuleAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessRuleAccessCountryRuleJSON `json:"-"`
}

// accessRuleAccessCountryRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessCountryRule]
type accessRuleAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessCountryRule) implementsAccessRule() {}

type AccessRuleAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                             `json:"country_code,required"`
	JSON        accessRuleAccessCountryRuleGeoJSON `json:"-"`
}

// accessRuleAccessCountryRuleGeoJSON contains the JSON metadata for the struct
// [AccessRuleAccessCountryRuleGeo]
type accessRuleAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type AccessRuleAccessDevicePostureRule struct {
	DevicePosture AccessRuleAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessRuleAccessDevicePostureRuleJSON          `json:"-"`
}

// accessRuleAccessDevicePostureRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessDevicePostureRule]
type accessRuleAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessRuleAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessDevicePostureRule) implementsAccessRule() {}

type AccessRuleAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                             `json:"integration_uid,required"`
	JSON           accessRuleAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessRuleAccessDevicePostureRuleDevicePostureJSON contains the JSON metadata
// for the struct [AccessRuleAccessDevicePostureRuleDevicePosture]
type accessRuleAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessRuleAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type AccessRuleAccessDomainRule struct {
	EmailDomain AccessRuleAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessRuleAccessDomainRuleJSON        `json:"-"`
}

// accessRuleAccessDomainRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessDomainRule]
type accessRuleAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessDomainRule) implementsAccessRule() {}

type AccessRuleAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                    `json:"domain,required"`
	JSON   accessRuleAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessRuleAccessDomainRuleEmailDomainJSON contains the JSON metadata for the
// struct [AccessRuleAccessDomainRuleEmailDomain]
type accessRuleAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type AccessRuleAccessEmailListRule struct {
	EmailList AccessRuleAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessRuleAccessEmailListRuleJSON      `json:"-"`
}

// accessRuleAccessEmailListRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessEmailListRule]
type accessRuleAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessEmailListRule) implementsAccessRule() {}

type AccessRuleAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                     `json:"id,required"`
	JSON accessRuleAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessRuleAccessEmailListRuleEmailListJSON contains the JSON metadata for the
// struct [AccessRuleAccessEmailListRuleEmailList]
type accessRuleAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
type AccessRuleAccessEmailRule struct {
	Email AccessRuleAccessEmailRuleEmail `json:"email,required"`
	JSON  accessRuleAccessEmailRuleJSON  `json:"-"`
}

// accessRuleAccessEmailRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessEmailRule]
type accessRuleAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessEmailRule) implementsAccessRule() {}

type AccessRuleAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                             `json:"email,required" format:"email"`
	JSON  accessRuleAccessEmailRuleEmailJSON `json:"-"`
}

// accessRuleAccessEmailRuleEmailJSON contains the JSON metadata for the struct
// [AccessRuleAccessEmailRuleEmail]
type accessRuleAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type AccessRuleAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone AccessRuleAccessEveryoneRuleEveryone `json:"everyone,required"`
	JSON     accessRuleAccessEveryoneRuleJSON     `json:"-"`
}

// accessRuleAccessEveryoneRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessEveryoneRule]
type accessRuleAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessEveryoneRule) implementsAccessRule() {}

// An empty object which matches on all users.
type AccessRuleAccessEveryoneRuleEveryone struct {
	JSON accessRuleAccessEveryoneRuleEveryoneJSON `json:"-"`
}

// accessRuleAccessEveryoneRuleEveryoneJSON contains the JSON metadata for the
// struct [AccessRuleAccessEveryoneRuleEveryone]
type accessRuleAccessEveryoneRuleEveryoneJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessEveryoneRuleEveryone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessEveryoneRuleEveryoneJSON) RawJSON() string {
	return r.raw
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessRuleAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessRuleAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessRuleAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessRuleAccessExternalEvaluationRuleJSON contains the JSON metadata for the
// struct [AccessRuleAccessExternalEvaluationRule]
type accessRuleAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessRuleAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessExternalEvaluationRule) implementsAccessRule() {}

type AccessRuleAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                       `json:"keys_url,required"`
	JSON    accessRuleAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessRuleAccessExternalEvaluationRuleExternalEvaluationJSON contains the JSON
// metadata for the struct
// [AccessRuleAccessExternalEvaluationRuleExternalEvaluation]
type accessRuleAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type AccessRuleAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessRuleAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessRuleAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessRuleAccessGitHubOrganizationRuleJSON contains the JSON metadata for the
// struct [AccessRuleAccessGitHubOrganizationRule]
type accessRuleAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessRuleAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessGitHubOrganizationRule) implementsAccessRule() {}

type AccessRuleAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	IdentityProviderID string `json:"identity_provider_id,required"`
	// The name of the organization.
	Name string `json:"name,required"`
	// The name of the team
	Team string                                                       `json:"team"`
	JSON accessRuleAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessRuleAccessGitHubOrganizationRuleGitHubOrganizationJSON contains the JSON
// metadata for the struct
// [AccessRuleAccessGitHubOrganizationRuleGitHubOrganization]
type accessRuleAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	IdentityProviderID apijson.Field
	Name               apijson.Field
	Team               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessRuleAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessRuleAccessGsuiteGroupRule struct {
	Gsuite AccessRuleAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessRuleAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessRuleAccessGsuiteGroupRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessGsuiteGroupRule]
type accessRuleAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessGsuiteGroupRule) implementsAccessRule() {}

type AccessRuleAccessGsuiteGroupRuleGsuite struct {
	// The email of the Google Workspace group.
	Email string `json:"email,required"`
	// The ID of your Google Workspace identity provider.
	IdentityProviderID string                                    `json:"identity_provider_id,required"`
	JSON               accessRuleAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessRuleAccessGsuiteGroupRuleGsuiteJSON contains the JSON metadata for the
// struct [AccessRuleAccessGsuiteGroupRuleGsuite]
type accessRuleAccessGsuiteGroupRuleGsuiteJSON struct {
	Email              apijson.Field
	IdentityProviderID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessRuleAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches a specific identity provider id.
type AccessRuleAccessLoginMethodRule struct {
	LoginMethod AccessRuleAccessLoginMethodRuleLoginMethod `json:"login_method,required"`
	JSON        accessRuleAccessLoginMethodRuleJSON        `json:"-"`
}

// accessRuleAccessLoginMethodRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessLoginMethodRule]
type accessRuleAccessLoginMethodRuleJSON struct {
	LoginMethod apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessLoginMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessLoginMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessLoginMethodRule) implementsAccessRule() {}

type AccessRuleAccessLoginMethodRuleLoginMethod struct {
	// The ID of an identity provider.
	ID   string                                         `json:"id,required"`
	JSON accessRuleAccessLoginMethodRuleLoginMethodJSON `json:"-"`
}

// accessRuleAccessLoginMethodRuleLoginMethodJSON contains the JSON metadata for
// the struct [AccessRuleAccessLoginMethodRuleLoginMethod]
type accessRuleAccessLoginMethodRuleLoginMethodJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessLoginMethodRuleLoginMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessLoginMethodRuleLoginMethodJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type AccessRuleAccessIPListRule struct {
	IPList AccessRuleAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessRuleAccessIPListRuleJSON   `json:"-"`
}

// accessRuleAccessIPListRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessIPListRule]
type accessRuleAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessIPListRule) implementsAccessRule() {}

type AccessRuleAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                               `json:"id,required"`
	JSON accessRuleAccessIPListRuleIPListJSON `json:"-"`
}

// accessRuleAccessIPListRuleIPListJSON contains the JSON metadata for the struct
// [AccessRuleAccessIPListRuleIPList]
type accessRuleAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address block.
type AccessRuleAccessIPRule struct {
	IP   AccessRuleAccessIPRuleIP   `json:"ip,required"`
	JSON accessRuleAccessIPRuleJSON `json:"-"`
}

// accessRuleAccessIPRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessIPRule]
type accessRuleAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessIPRule) implementsAccessRule() {}

type AccessRuleAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                       `json:"ip,required"`
	JSON accessRuleAccessIPRuleIPJSON `json:"-"`
}

// accessRuleAccessIPRuleIPJSON contains the JSON metadata for the struct
// [AccessRuleAccessIPRuleIP]
type accessRuleAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessRuleAccessOktaGroupRule struct {
	Okta AccessRuleAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessRuleAccessOktaGroupRuleJSON `json:"-"`
}

// accessRuleAccessOktaGroupRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessOktaGroupRule]
type accessRuleAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessOktaGroupRule) implementsAccessRule() {}

type AccessRuleAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	IdentityProviderID string `json:"identity_provider_id,required"`
	// The name of the Okta group.
	Name string                                `json:"name,required"`
	JSON accessRuleAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessRuleAccessOktaGroupRuleOktaJSON contains the JSON metadata for the struct
// [AccessRuleAccessOktaGroupRuleOkta]
type accessRuleAccessOktaGroupRuleOktaJSON struct {
	IdentityProviderID apijson.Field
	Name               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessRuleAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessRuleAccessSAMLGroupRule struct {
	SAML AccessRuleAccessSAMLGroupRuleSAML `json:"saml,required"`
	JSON accessRuleAccessSAMLGroupRuleJSON `json:"-"`
}

// accessRuleAccessSAMLGroupRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessSAMLGroupRule]
type accessRuleAccessSAMLGroupRuleJSON struct {
	SAML        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessSAMLGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessSAMLGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessSAMLGroupRule) implementsAccessRule() {}

type AccessRuleAccessSAMLGroupRuleSAML struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string `json:"attribute_value,required"`
	// The ID of your SAML identity provider.
	IdentityProviderID string                                `json:"identity_provider_id,required"`
	JSON               accessRuleAccessSAMLGroupRuleSAMLJSON `json:"-"`
}

// accessRuleAccessSAMLGroupRuleSAMLJSON contains the JSON metadata for the struct
// [AccessRuleAccessSAMLGroupRuleSAML]
type accessRuleAccessSAMLGroupRuleSAMLJSON struct {
	AttributeName      apijson.Field
	AttributeValue     apijson.Field
	IdentityProviderID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessRuleAccessSAMLGroupRuleSAML) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessSAMLGroupRuleSAMLJSON) RawJSON() string {
	return r.raw
}

// Matches an OIDC claim. Requires an OIDC identity provider.
type AccessRuleAccessOidcClaimRule struct {
	Oidc AccessRuleAccessOidcClaimRuleOidc `json:"oidc,required"`
	JSON accessRuleAccessOidcClaimRuleJSON `json:"-"`
}

// accessRuleAccessOidcClaimRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessOidcClaimRule]
type accessRuleAccessOidcClaimRuleJSON struct {
	Oidc        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessOidcClaimRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessOidcClaimRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessOidcClaimRule) implementsAccessRule() {}

type AccessRuleAccessOidcClaimRuleOidc struct {
	// The name of the OIDC claim.
	ClaimName string `json:"claim_name,required"`
	// The OIDC claim value to look for.
	ClaimValue string `json:"claim_value,required"`
	// The ID of your OIDC identity provider.
	IdentityProviderID string                                `json:"identity_provider_id,required"`
	JSON               accessRuleAccessOidcClaimRuleOidcJSON `json:"-"`
}

// accessRuleAccessOidcClaimRuleOidcJSON contains the JSON metadata for the struct
// [AccessRuleAccessOidcClaimRuleOidc]
type accessRuleAccessOidcClaimRuleOidcJSON struct {
	ClaimName          apijson.Field
	ClaimValue         apijson.Field
	IdentityProviderID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessRuleAccessOidcClaimRuleOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessOidcClaimRuleOidcJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type AccessRuleAccessServiceTokenRule struct {
	ServiceToken AccessRuleAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessRuleAccessServiceTokenRuleJSON         `json:"-"`
}

// accessRuleAccessServiceTokenRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessServiceTokenRule]
type accessRuleAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessRuleAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessServiceTokenRule) implementsAccessRule() {}

type AccessRuleAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                           `json:"token_id,required"`
	JSON    accessRuleAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessRuleAccessServiceTokenRuleServiceTokenJSON contains the JSON metadata for
// the struct [AccessRuleAccessServiceTokenRuleServiceToken]
type accessRuleAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches OAuth 2.0 access tokens issued by the specified Access OIDC SaaS
// application. Only compatible with non_identity and bypass decisions.
type AccessRuleAccessLinkedAppTokenRule struct {
	LinkedAppToken AccessRuleAccessLinkedAppTokenRuleLinkedAppToken `json:"linked_app_token,required"`
	JSON           accessRuleAccessLinkedAppTokenRuleJSON           `json:"-"`
}

// accessRuleAccessLinkedAppTokenRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessLinkedAppTokenRule]
type accessRuleAccessLinkedAppTokenRuleJSON struct {
	LinkedAppToken apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessRuleAccessLinkedAppTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessLinkedAppTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessLinkedAppTokenRule) implementsAccessRule() {}

type AccessRuleAccessLinkedAppTokenRuleLinkedAppToken struct {
	// The ID of an Access OIDC SaaS application
	AppUid string                                               `json:"app_uid,required"`
	JSON   accessRuleAccessLinkedAppTokenRuleLinkedAppTokenJSON `json:"-"`
}

// accessRuleAccessLinkedAppTokenRuleLinkedAppTokenJSON contains the JSON metadata
// for the struct [AccessRuleAccessLinkedAppTokenRuleLinkedAppToken]
type accessRuleAccessLinkedAppTokenRuleLinkedAppTokenJSON struct {
	AppUid      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessLinkedAppTokenRuleLinkedAppToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessLinkedAppTokenRuleLinkedAppTokenJSON) RawJSON() string {
	return r.raw
}

// Matches an Access group.
type AccessRuleParam struct {
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token"`
	AuthContext          param.Field[interface{}] `json:"auth_context"`
	AuthMethod           param.Field[interface{}] `json:"auth_method"`
	AzureAd              param.Field[interface{}] `json:"azureAD"`
	Certificate          param.Field[interface{}] `json:"certificate"`
	CommonName           param.Field[interface{}] `json:"common_name"`
	DevicePosture        param.Field[interface{}] `json:"device_posture"`
	Email                param.Field[interface{}] `json:"email"`
	EmailDomain          param.Field[interface{}] `json:"email_domain"`
	EmailList            param.Field[interface{}] `json:"email_list"`
	Everyone             param.Field[interface{}] `json:"everyone"`
	ExternalEvaluation   param.Field[interface{}] `json:"external_evaluation"`
	Geo                  param.Field[interface{}] `json:"geo"`
	GitHubOrganization   param.Field[interface{}] `json:"github-organization"`
	Group                param.Field[interface{}] `json:"group"`
	Gsuite               param.Field[interface{}] `json:"gsuite"`
	IP                   param.Field[interface{}] `json:"ip"`
	IPList               param.Field[interface{}] `json:"ip_list"`
	LinkedAppToken       param.Field[interface{}] `json:"linked_app_token"`
	LoginMethod          param.Field[interface{}] `json:"login_method"`
	Oidc                 param.Field[interface{}] `json:"oidc"`
	Okta                 param.Field[interface{}] `json:"okta"`
	SAML                 param.Field[interface{}] `json:"saml"`
	ServiceToken         param.Field[interface{}] `json:"service_token"`
}

func (r AccessRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleParam) implementsAccessRuleUnionParam() {}

// Matches an Access group.
//
// Satisfied by [AccessRuleAccessAccessGroupRuleParam],
// [AccessRuleAccessAnyValidServiceTokenRuleParam],
// [AccessRuleAccessAuthContextRuleParam],
// [AccessRuleAccessAuthenticationMethodRuleParam],
// [AccessRuleAccessAzureGroupRuleParam], [AccessRuleAccessCertificateRuleParam],
// [AccessRuleAccessCommonNameRuleParam], [AccessRuleAccessCountryRuleParam],
// [AccessRuleAccessDevicePostureRuleParam], [AccessRuleAccessDomainRuleParam],
// [AccessRuleAccessEmailListRuleParam], [AccessRuleAccessEmailRuleParam],
// [AccessRuleAccessEveryoneRuleParam],
// [AccessRuleAccessExternalEvaluationRuleParam],
// [AccessRuleAccessGitHubOrganizationRuleParam],
// [AccessRuleAccessGsuiteGroupRuleParam], [AccessRuleAccessLoginMethodRuleParam],
// [AccessRuleAccessIPListRuleParam], [AccessRuleAccessIPRuleParam],
// [AccessRuleAccessOktaGroupRuleParam], [AccessRuleAccessSAMLGroupRuleParam],
// [AccessRuleAccessOidcClaimRuleParam], [AccessRuleAccessServiceTokenRuleParam],
// [AccessRuleAccessLinkedAppTokenRuleParam], [AccessRuleParam].
type AccessRuleUnionParam interface {
	implementsAccessRuleUnionParam()
}

// Matches an Access group.
type AccessRuleAccessAccessGroupRuleParam struct {
	Group param.Field[AccessRuleAccessAccessGroupRuleGroupParam] `json:"group,required"`
}

func (r AccessRuleAccessAccessGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessAccessGroupRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessAccessGroupRuleGroupParam struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessRuleAccessAccessGroupRuleGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessRuleAccessAnyValidServiceTokenRuleParam struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[AccessRuleAccessAnyValidServiceTokenRuleAnyValidServiceTokenParam] `json:"any_valid_service_token,required"`
}

func (r AccessRuleAccessAnyValidServiceTokenRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessAnyValidServiceTokenRuleParam) implementsAccessRuleUnionParam() {}

// An empty object which matches on all service tokens.
type AccessRuleAccessAnyValidServiceTokenRuleAnyValidServiceTokenParam struct {
}

func (r AccessRuleAccessAnyValidServiceTokenRuleAnyValidServiceTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure Authentication Context. Requires an Azure identity provider.
type AccessRuleAccessAuthContextRuleParam struct {
	AuthContext param.Field[AccessRuleAccessAuthContextRuleAuthContextParam] `json:"auth_context,required"`
}

func (r AccessRuleAccessAuthContextRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessAuthContextRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessAuthContextRuleAuthContextParam struct {
	// The ID of an Authentication context.
	ID param.Field[string] `json:"id,required"`
	// The ACID of an Authentication context.
	AcID param.Field[string] `json:"ac_id,required"`
	// The ID of your Azure identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id,required"`
}

func (r AccessRuleAccessAuthContextRuleAuthContextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessRuleAccessAuthenticationMethodRuleParam struct {
	AuthMethod param.Field[AccessRuleAccessAuthenticationMethodRuleAuthMethodParam] `json:"auth_method,required"`
}

func (r AccessRuleAccessAuthenticationMethodRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessAuthenticationMethodRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessAuthenticationMethodRuleAuthMethodParam struct {
	// The type of authentication method
	// https://datatracker.ietf.org/doc/html/rfc8176#section-2.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessRuleAccessAuthenticationMethodRuleAuthMethodParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessRuleAccessAzureGroupRuleParam struct {
	AzureAd param.Field[AccessRuleAccessAzureGroupRuleAzureAdParam] `json:"azureAD,required"`
}

func (r AccessRuleAccessAzureGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessAzureGroupRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessAzureGroupRuleAzureAdParam struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id,required"`
}

func (r AccessRuleAccessAzureGroupRuleAzureAdParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessRuleAccessCertificateRuleParam struct {
	Certificate param.Field[AccessRuleAccessCertificateRuleCertificateParam] `json:"certificate,required"`
}

func (r AccessRuleAccessCertificateRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessCertificateRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessCertificateRuleCertificateParam struct {
}

func (r AccessRuleAccessCertificateRuleCertificateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific common name.
type AccessRuleAccessCommonNameRuleParam struct {
	CommonName param.Field[AccessRuleAccessCommonNameRuleCommonNameParam] `json:"common_name,required"`
}

func (r AccessRuleAccessCommonNameRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessCommonNameRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessCommonNameRuleCommonNameParam struct {
	// The common name to match.
	CommonName param.Field[string] `json:"common_name,required"`
}

func (r AccessRuleAccessCommonNameRuleCommonNameParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessRuleAccessCountryRuleParam struct {
	Geo param.Field[AccessRuleAccessCountryRuleGeoParam] `json:"geo,required"`
}

func (r AccessRuleAccessCountryRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessCountryRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessCountryRuleGeoParam struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessRuleAccessCountryRuleGeoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessRuleAccessDevicePostureRuleParam struct {
	DevicePosture param.Field[AccessRuleAccessDevicePostureRuleDevicePostureParam] `json:"device_posture,required"`
}

func (r AccessRuleAccessDevicePostureRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessDevicePostureRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessDevicePostureRuleDevicePostureParam struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessRuleAccessDevicePostureRuleDevicePostureParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessRuleAccessDomainRuleParam struct {
	EmailDomain param.Field[AccessRuleAccessDomainRuleEmailDomainParam] `json:"email_domain,required"`
}

func (r AccessRuleAccessDomainRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessDomainRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessDomainRuleEmailDomainParam struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessRuleAccessDomainRuleEmailDomainParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessRuleAccessEmailListRuleParam struct {
	EmailList param.Field[AccessRuleAccessEmailListRuleEmailListParam] `json:"email_list,required"`
}

func (r AccessRuleAccessEmailListRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessEmailListRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessEmailListRuleEmailListParam struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessRuleAccessEmailListRuleEmailListParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
type AccessRuleAccessEmailRuleParam struct {
	Email param.Field[AccessRuleAccessEmailRuleEmailParam] `json:"email,required"`
}

func (r AccessRuleAccessEmailRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessEmailRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessEmailRuleEmailParam struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessRuleAccessEmailRuleEmailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessRuleAccessEveryoneRuleParam struct {
	// An empty object which matches on all users.
	Everyone param.Field[AccessRuleAccessEveryoneRuleEveryoneParam] `json:"everyone,required"`
}

func (r AccessRuleAccessEveryoneRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessEveryoneRuleParam) implementsAccessRuleUnionParam() {}

// An empty object which matches on all users.
type AccessRuleAccessEveryoneRuleEveryoneParam struct {
}

func (r AccessRuleAccessEveryoneRuleEveryoneParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessRuleAccessExternalEvaluationRuleParam struct {
	ExternalEvaluation param.Field[AccessRuleAccessExternalEvaluationRuleExternalEvaluationParam] `json:"external_evaluation,required"`
}

func (r AccessRuleAccessExternalEvaluationRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessExternalEvaluationRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessExternalEvaluationRuleExternalEvaluationParam struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessRuleAccessExternalEvaluationRuleExternalEvaluationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessRuleAccessGitHubOrganizationRuleParam struct {
	GitHubOrganization param.Field[AccessRuleAccessGitHubOrganizationRuleGitHubOrganizationParam] `json:"github-organization,required"`
}

func (r AccessRuleAccessGitHubOrganizationRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessGitHubOrganizationRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessGitHubOrganizationRuleGitHubOrganizationParam struct {
	// The ID of your Github identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
	// The name of the team
	Team param.Field[string] `json:"team"`
}

func (r AccessRuleAccessGitHubOrganizationRuleGitHubOrganizationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessRuleAccessGsuiteGroupRuleParam struct {
	Gsuite param.Field[AccessRuleAccessGsuiteGroupRuleGsuiteParam] `json:"gsuite,required"`
}

func (r AccessRuleAccessGsuiteGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessGsuiteGroupRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessGsuiteGroupRuleGsuiteParam struct {
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
	// The ID of your Google Workspace identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id,required"`
}

func (r AccessRuleAccessGsuiteGroupRuleGsuiteParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific identity provider id.
type AccessRuleAccessLoginMethodRuleParam struct {
	LoginMethod param.Field[AccessRuleAccessLoginMethodRuleLoginMethodParam] `json:"login_method,required"`
}

func (r AccessRuleAccessLoginMethodRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessLoginMethodRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessLoginMethodRuleLoginMethodParam struct {
	// The ID of an identity provider.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessRuleAccessLoginMethodRuleLoginMethodParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessRuleAccessIPListRuleParam struct {
	IPList param.Field[AccessRuleAccessIPListRuleIPListParam] `json:"ip_list,required"`
}

func (r AccessRuleAccessIPListRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessIPListRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessIPListRuleIPListParam struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessRuleAccessIPListRuleIPListParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address block.
type AccessRuleAccessIPRuleParam struct {
	IP param.Field[AccessRuleAccessIPRuleIPParam] `json:"ip,required"`
}

func (r AccessRuleAccessIPRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessIPRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessIPRuleIPParam struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessRuleAccessIPRuleIPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessRuleAccessOktaGroupRuleParam struct {
	Okta param.Field[AccessRuleAccessOktaGroupRuleOktaParam] `json:"okta,required"`
}

func (r AccessRuleAccessOktaGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessOktaGroupRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessOktaGroupRuleOktaParam struct {
	// The ID of your Okta identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id,required"`
	// The name of the Okta group.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessRuleAccessOktaGroupRuleOktaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessRuleAccessSAMLGroupRuleParam struct {
	SAML param.Field[AccessRuleAccessSAMLGroupRuleSAMLParam] `json:"saml,required"`
}

func (r AccessRuleAccessSAMLGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessSAMLGroupRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessSAMLGroupRuleSAMLParam struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
	// The ID of your SAML identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id,required"`
}

func (r AccessRuleAccessSAMLGroupRuleSAMLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an OIDC claim. Requires an OIDC identity provider.
type AccessRuleAccessOidcClaimRuleParam struct {
	Oidc param.Field[AccessRuleAccessOidcClaimRuleOidcParam] `json:"oidc,required"`
}

func (r AccessRuleAccessOidcClaimRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessOidcClaimRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessOidcClaimRuleOidcParam struct {
	// The name of the OIDC claim.
	ClaimName param.Field[string] `json:"claim_name,required"`
	// The OIDC claim value to look for.
	ClaimValue param.Field[string] `json:"claim_value,required"`
	// The ID of your OIDC identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id,required"`
}

func (r AccessRuleAccessOidcClaimRuleOidcParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessRuleAccessServiceTokenRuleParam struct {
	ServiceToken param.Field[AccessRuleAccessServiceTokenRuleServiceTokenParam] `json:"service_token,required"`
}

func (r AccessRuleAccessServiceTokenRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessServiceTokenRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessServiceTokenRuleServiceTokenParam struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessRuleAccessServiceTokenRuleServiceTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches OAuth 2.0 access tokens issued by the specified Access OIDC SaaS
// application. Only compatible with non_identity and bypass decisions.
type AccessRuleAccessLinkedAppTokenRuleParam struct {
	LinkedAppToken param.Field[AccessRuleAccessLinkedAppTokenRuleLinkedAppTokenParam] `json:"linked_app_token,required"`
}

func (r AccessRuleAccessLinkedAppTokenRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessLinkedAppTokenRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessLinkedAppTokenRuleLinkedAppTokenParam struct {
	// The ID of an Access OIDC SaaS application
	AppUid param.Field[string] `json:"app_uid,required"`
}

func (r AccessRuleAccessLinkedAppTokenRuleLinkedAppTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
