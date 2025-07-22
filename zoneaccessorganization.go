// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneAccessOrganizationService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAccessOrganizationService] method instead.
type ZoneAccessOrganizationService struct {
	Options []option.RequestOption
}

// NewZoneAccessOrganizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAccessOrganizationService(opts ...option.RequestOption) (r *ZoneAccessOrganizationService) {
	r = &ZoneAccessOrganizationService{}
	r.Options = opts
	return
}

// Sets up a Zero Trust organization for your account.
func (r *ZoneAccessOrganizationService) New(ctx context.Context, zoneID interface{}, body ZoneAccessOrganizationNewParams, opts ...option.RequestOption) (res *SingleResponseOrganizationZone, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/access/organizations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the configuration for your Zero Trust organization.
func (r *ZoneAccessOrganizationService) Get(ctx context.Context, zoneID interface{}, opts ...option.RequestOption) (res *SingleResponseOrganizationZone, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/access/organizations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the configuration for your Zero Trust organization.
func (r *ZoneAccessOrganizationService) Update(ctx context.Context, zoneID interface{}, body ZoneAccessOrganizationUpdateParams, opts ...option.RequestOption) (res *SingleResponseOrganizationZone, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/access/organizations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Revokes a user's access across all applications.
func (r *ZoneAccessOrganizationService) RevokeUser(ctx context.Context, zoneID interface{}, body ZoneAccessOrganizationRevokeUserParams, opts ...option.RequestOption) (res *AccessEmptyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/access/organizations/revoke_user", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SingleResponseOrganizationZone struct {
	Result SingleResponseOrganizationZoneResult `json:"result"`
	JSON   singleResponseOrganizationZoneJSON   `json:"-"`
	APIResponseSingleAccess
}

// singleResponseOrganizationZoneJSON contains the JSON metadata for the struct
// [SingleResponseOrganizationZone]
type singleResponseOrganizationZoneJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseOrganizationZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseOrganizationZoneJSON) RawJSON() string {
	return r.raw
}

type SingleResponseOrganizationZoneResult struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string    `json:"auth_domain"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool        `json:"is_ui_read_only"`
	LoginDesign  LoginDesign `json:"login_design"`
	// The name of your Zero Trust organization.
	Name string `json:"name"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason string    `json:"ui_read_only_toggle_reason"`
	UpdatedAt              time.Time `json:"updated_at" format:"date-time"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime string                                   `json:"user_seat_expiration_inactive_time"`
	JSON                           singleResponseOrganizationZoneResultJSON `json:"-"`
}

// singleResponseOrganizationZoneResultJSON contains the JSON metadata for the
// struct [SingleResponseOrganizationZoneResult]
type singleResponseOrganizationZoneResultJSON struct {
	AuthDomain                     apijson.Field
	CreatedAt                      apijson.Field
	IsUiReadOnly                   apijson.Field
	LoginDesign                    apijson.Field
	Name                           apijson.Field
	UiReadOnlyToggleReason         apijson.Field
	UpdatedAt                      apijson.Field
	UserSeatExpirationInactiveTime apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *SingleResponseOrganizationZoneResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseOrganizationZoneResultJSON) RawJSON() string {
	return r.raw
}

type ZoneAccessOrganizationNewParams struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain,required"`
	// The name of your Zero Trust organization.
	Name param.Field[string] `json:"name,required"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly param.Field[bool]             `json:"is_ui_read_only"`
	LoginDesign  param.Field[LoginDesignParam] `json:"login_design"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason param.Field[string] `json:"ui_read_only_toggle_reason"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime param.Field[string] `json:"user_seat_expiration_inactive_time"`
}

func (r ZoneAccessOrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessOrganizationUpdateParams struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly param.Field[bool]             `json:"is_ui_read_only"`
	LoginDesign  param.Field[LoginDesignParam] `json:"login_design"`
	// The name of your Zero Trust organization.
	Name param.Field[string] `json:"name"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason param.Field[string] `json:"ui_read_only_toggle_reason"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime param.Field[string] `json:"user_seat_expiration_inactive_time"`
}

func (r ZoneAccessOrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessOrganizationRevokeUserParams struct {
	// The email of the user to revoke.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessOrganizationRevokeUserParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
