// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneOriginTlsClientAuthSettingService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneOriginTlsClientAuthSettingService] method instead.
type ZoneOriginTlsClientAuthSettingService struct {
	Options []option.RequestOption
}

// NewZoneOriginTlsClientAuthSettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneOriginTlsClientAuthSettingService(opts ...option.RequestOption) (r *ZoneOriginTlsClientAuthSettingService) {
	r = &ZoneOriginTlsClientAuthSettingService{}
	r.Options = opts
	return
}

// Get whether zone-level authenticated origin pulls is enabled or not. It is false
// by default.
func (r *ZoneOriginTlsClientAuthSettingService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *EnabledResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable or disable zone-level authenticated origin pulls. 'enabled' should be set
// true either before/after the certificate is uploaded to see the certificate in
// use.
func (r *ZoneOriginTlsClientAuthSettingService) Update(ctx context.Context, zoneID string, body ZoneOriginTlsClientAuthSettingUpdateParams, opts ...option.RequestOption) (res *EnabledResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type EnabledResponse struct {
	Result EnabledResponseResult `json:"result"`
	JSON   enabledResponseJSON   `json:"-"`
	APIResponseSingleTlsCertificates
}

// enabledResponseJSON contains the JSON metadata for the struct [EnabledResponse]
type enabledResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnabledResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enabledResponseJSON) RawJSON() string {
	return r.raw
}

type EnabledResponseResult struct {
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool                      `json:"enabled"`
	JSON    enabledResponseResultJSON `json:"-"`
}

// enabledResponseResultJSON contains the JSON metadata for the struct
// [EnabledResponseResult]
type enabledResponseResultJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnabledResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enabledResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneOriginTlsClientAuthSettingUpdateParams struct {
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneOriginTlsClientAuthSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
