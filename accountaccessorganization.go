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

// AccountAccessOrganizationService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessOrganizationService] method instead.
type AccountAccessOrganizationService struct {
	Options []option.RequestOption
	Doh     *AccountAccessOrganizationDohService
}

// NewAccountAccessOrganizationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessOrganizationService(opts ...option.RequestOption) (r *AccountAccessOrganizationService) {
	r = &AccountAccessOrganizationService{}
	r.Options = opts
	r.Doh = NewAccountAccessOrganizationDohService(opts...)
	return
}

// Sets up a Zero Trust organization for your account.
func (r *AccountAccessOrganizationService) New(ctx context.Context, accountID string, body AccountAccessOrganizationNewParams, opts ...option.RequestOption) (res *SingleResponseOrganization, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/organizations", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the configuration for your Zero Trust organization.
func (r *AccountAccessOrganizationService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *SingleResponseOrganization, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/organizations", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the configuration for your Zero Trust organization.
func (r *AccountAccessOrganizationService) Update(ctx context.Context, accountID string, body AccountAccessOrganizationUpdateParams, opts ...option.RequestOption) (res *SingleResponseOrganization, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/organizations", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Revokes a user's access across all applications.
func (r *AccountAccessOrganizationService) RevokeUser(ctx context.Context, accountID string, params AccountAccessOrganizationRevokeUserParams, opts ...option.RequestOption) (res *AccessEmptyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/organizations/revoke_user", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccessEmptyResponse struct {
	Result  AccessEmptyResponseResult  `json:"result"`
	Success AccessEmptyResponseSuccess `json:"success"`
	JSON    accessEmptyResponseJSON    `json:"-"`
}

// accessEmptyResponseJSON contains the JSON metadata for the struct
// [AccessEmptyResponse]
type accessEmptyResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessEmptyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessEmptyResponseJSON) RawJSON() string {
	return r.raw
}

type AccessEmptyResponseResult bool

const (
	AccessEmptyResponseResultTrue  AccessEmptyResponseResult = true
	AccessEmptyResponseResultFalse AccessEmptyResponseResult = false
)

func (r AccessEmptyResponseResult) IsKnown() bool {
	switch r {
	case AccessEmptyResponseResultTrue, AccessEmptyResponseResultFalse:
		return true
	}
	return false
}

type AccessEmptyResponseSuccess bool

const (
	AccessEmptyResponseSuccessTrue  AccessEmptyResponseSuccess = true
	AccessEmptyResponseSuccessFalse AccessEmptyResponseSuccess = false
)

func (r AccessEmptyResponseSuccess) IsKnown() bool {
	switch r {
	case AccessEmptyResponseSuccessTrue, AccessEmptyResponseSuccessFalse:
		return true
	}
	return false
}

type CustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string          `json:"identity_denied"`
	JSON           customPagesJSON `json:"-"`
}

// customPagesJSON contains the JSON metadata for the struct [CustomPages]
type customPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPagesJSON) RawJSON() string {
	return r.raw
}

type CustomPagesParam struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden param.Field[string] `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied param.Field[string] `json:"identity_denied"`
}

func (r CustomPagesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string          `json:"text_color"`
	JSON      loginDesignJSON `json:"-"`
}

// loginDesignJSON contains the JSON metadata for the struct [LoginDesign]
type loginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loginDesignJSON) RawJSON() string {
	return r.raw
}

type LoginDesignParam struct {
	// The background color on your login page.
	BackgroundColor param.Field[string] `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText param.Field[string] `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText param.Field[string] `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath param.Field[string] `json:"logo_path"`
	// The text color on your login page.
	TextColor param.Field[string] `json:"text_color"`
}

func (r LoginDesignParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SingleResponseOrganization struct {
	Result SingleResponseOrganizationResult `json:"result"`
	JSON   singleResponseOrganizationJSON   `json:"-"`
	APIResponseSingleAccess
}

// singleResponseOrganizationJSON contains the JSON metadata for the struct
// [SingleResponseOrganization]
type singleResponseOrganizationJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseOrganizationJSON) RawJSON() string {
	return r.raw
}

type SingleResponseOrganizationResult struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp bool `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool        `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time   `json:"created_at" format:"date-time"`
	CustomPages            CustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool        `json:"is_ui_read_only"`
	LoginDesign  LoginDesign `json:"login_design"`
	// The name of your Zero Trust organization.
	Name string `json:"name"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration string `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason string    `json:"ui_read_only_toggle_reason"`
	UpdatedAt              time.Time `json:"updated_at" format:"date-time"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Minimum value for this setting is 1
	// month (730h). Must be in the format `300ms` or `2h45m`. Valid time units are:
	// `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime string `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WarpAuthSessionDuration string                               `json:"warp_auth_session_duration"`
	JSON                    singleResponseOrganizationResultJSON `json:"-"`
}

// singleResponseOrganizationResultJSON contains the JSON metadata for the struct
// [SingleResponseOrganizationResult]
type singleResponseOrganizationResultJSON struct {
	AllowAuthenticateViaWarp       apijson.Field
	AuthDomain                     apijson.Field
	AutoRedirectToIdentity         apijson.Field
	CreatedAt                      apijson.Field
	CustomPages                    apijson.Field
	IsUiReadOnly                   apijson.Field
	LoginDesign                    apijson.Field
	Name                           apijson.Field
	SessionDuration                apijson.Field
	UiReadOnlyToggleReason         apijson.Field
	UpdatedAt                      apijson.Field
	UserSeatExpirationInactiveTime apijson.Field
	WarpAuthSessionDuration        apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *SingleResponseOrganizationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseOrganizationResultJSON) RawJSON() string {
	return r.raw
}

type AccountAccessOrganizationNewParams struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain,required"`
	// The name of your Zero Trust organization.
	Name param.Field[string] `json:"name,required"`
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp param.Field[bool] `json:"allow_authenticate_via_warp"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly param.Field[bool]             `json:"is_ui_read_only"`
	LoginDesign  param.Field[LoginDesignParam] `json:"login_design"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason param.Field[string] `json:"ui_read_only_toggle_reason"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Minimum value for this setting is 1
	// month (730h). Must be in the format `300ms` or `2h45m`. Valid time units are:
	// `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime param.Field[string] `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WarpAuthSessionDuration param.Field[string] `json:"warp_auth_session_duration"`
}

func (r AccountAccessOrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessOrganizationUpdateParams struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity param.Field[bool]             `json:"auto_redirect_to_identity"`
	CustomPages            param.Field[CustomPagesParam] `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly param.Field[bool]             `json:"is_ui_read_only"`
	LoginDesign  param.Field[LoginDesignParam] `json:"login_design"`
	// The name of your Zero Trust organization.
	Name param.Field[string] `json:"name"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason param.Field[string] `json:"ui_read_only_toggle_reason"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Minimum value for this setting is 1
	// month (730h). Must be in the format `300ms` or `2h45m`. Valid time units are:
	// `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime param.Field[string] `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WarpAuthSessionDuration param.Field[string] `json:"warp_auth_session_duration"`
}

func (r AccountAccessOrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessOrganizationRevokeUserParams struct {
	// The email of the user to revoke.
	Email param.Field[string] `json:"email,required"`
	// When set to `true`, all devices associated with the user will be revoked.
	QueryDevices param.Field[bool] `query:"devices"`
	// When set to `true`, all devices associated with the user will be revoked.
	BodyDevices param.Field[bool] `json:"devices"`
	// The uuid of the user to revoke.
	UserUid param.Field[string] `json:"user_uid"`
	// When set to `true`, the user will be required to re-authenticate to WARP for all
	// Gateway policies that enforce a WARP client session duration. When `false`, the
	// user’s WARP session will remain active
	WarpSessionReauth param.Field[bool] `json:"warp_session_reauth"`
}

func (r AccountAccessOrganizationRevokeUserParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountAccessOrganizationRevokeUserParams]'s query
// parameters as `url.Values`.
func (r AccountAccessOrganizationRevokeUserParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
