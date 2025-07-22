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

// ZoneAccessCertificateSettingService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAccessCertificateSettingService] method instead.
type ZoneAccessCertificateSettingService struct {
	Options []option.RequestOption
}

// NewZoneAccessCertificateSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneAccessCertificateSettingService(opts ...option.RequestOption) (r *ZoneAccessCertificateSettingService) {
	r = &ZoneAccessCertificateSettingService{}
	r.Options = opts
	return
}

// Updates an mTLS certificate's hostname settings.
func (r *ZoneAccessCertificateSettingService) Update(ctx context.Context, zoneID string, body ZoneAccessCertificateSettingUpdateParams, opts ...option.RequestOption) (res *ResponseCollectionHostnames, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/certificates/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all mTLS hostname settings for this zone.
func (r *ZoneAccessCertificateSettingService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ResponseCollectionHostnames, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/certificates/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneAccessCertificateSettingUpdateParams struct {
	Settings param.Field[[]AccessSettingsParam] `json:"settings,required"`
}

func (r ZoneAccessCertificateSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
