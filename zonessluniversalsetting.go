// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneSslUniversalSettingService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSslUniversalSettingService] method instead.
type ZoneSslUniversalSettingService struct {
	Options []option.RequestOption
}

// NewZoneSslUniversalSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSslUniversalSettingService(opts ...option.RequestOption) (r *ZoneSslUniversalSettingService) {
	r = &ZoneSslUniversalSettingService{}
	r.Options = opts
	return
}

// Get Universal SSL Settings for a Zone.
func (r *ZoneSslUniversalSettingService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SslUniversalSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/universal/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patch Universal SSL Settings for a Zone.
func (r *ZoneSslUniversalSettingService) Update(ctx context.Context, zoneID string, body ZoneSslUniversalSettingUpdateParams, opts ...option.RequestOption) (res *SslUniversalSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/universal/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type SslUniversalSettingsResponse struct {
	Result Universal                        `json:"result"`
	JSON   sslUniversalSettingsResponseJSON `json:"-"`
	APIResponseSingleTlsCertificates
}

// sslUniversalSettingsResponseJSON contains the JSON metadata for the struct
// [SslUniversalSettingsResponse]
type sslUniversalSettingsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SslUniversalSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sslUniversalSettingsResponseJSON) RawJSON() string {
	return r.raw
}

type Universal struct {
	// Disabling Universal SSL removes any currently active Universal SSL certificates
	// for your zone from the edge and prevents any future Universal SSL certificates
	// from being ordered. If there are no advanced certificates or custom certificates
	// uploaded for the domain, visitors will be unable to access the domain over
	// HTTPS.
	//
	// By disabling Universal SSL, you understand that the following Cloudflare
	// settings and preferences will result in visitors being unable to visit your
	// domain unless you have uploaded a custom certificate or purchased an advanced
	// certificate.
	//
	// - HSTS
	// - Always Use HTTPS
	// - Opportunistic Encryption
	// - Onion Routing
	// - Any Page Rules redirecting traffic to HTTPS
	//
	// Similarly, any HTTP redirect to HTTPS at the origin while the Cloudflare proxy
	// is enabled will result in users being unable to visit your site without a valid
	// certificate at Cloudflare's edge.
	//
	// If you do not have a valid custom or advanced certificate at Cloudflare's edge
	// and are unsure if any of the above Cloudflare settings are enabled, or if any
	// HTTP redirects exist at your origin, we advise leaving Universal SSL enabled for
	// your domain.
	Enabled bool          `json:"enabled"`
	JSON    universalJSON `json:"-"`
}

// universalJSON contains the JSON metadata for the struct [Universal]
type universalJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Universal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r universalJSON) RawJSON() string {
	return r.raw
}

type UniversalParam struct {
	// Disabling Universal SSL removes any currently active Universal SSL certificates
	// for your zone from the edge and prevents any future Universal SSL certificates
	// from being ordered. If there are no advanced certificates or custom certificates
	// uploaded for the domain, visitors will be unable to access the domain over
	// HTTPS.
	//
	// By disabling Universal SSL, you understand that the following Cloudflare
	// settings and preferences will result in visitors being unable to visit your
	// domain unless you have uploaded a custom certificate or purchased an advanced
	// certificate.
	//
	// - HSTS
	// - Always Use HTTPS
	// - Opportunistic Encryption
	// - Onion Routing
	// - Any Page Rules redirecting traffic to HTTPS
	//
	// Similarly, any HTTP redirect to HTTPS at the origin while the Cloudflare proxy
	// is enabled will result in users being unable to visit your site without a valid
	// certificate at Cloudflare's edge.
	//
	// If you do not have a valid custom or advanced certificate at Cloudflare's edge
	// and are unsure if any of the above Cloudflare settings are enabled, or if any
	// HTTP redirects exist at your origin, we advise leaving Universal SSL enabled for
	// your domain.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r UniversalParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSslUniversalSettingUpdateParams struct {
	Universal UniversalParam `json:"universal,required"`
}

func (r ZoneSslUniversalSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Universal)
}
