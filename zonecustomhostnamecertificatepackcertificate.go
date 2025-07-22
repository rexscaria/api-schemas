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

// ZoneCustomHostnameCertificatePackCertificateService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCustomHostnameCertificatePackCertificateService] method instead.
type ZoneCustomHostnameCertificatePackCertificateService struct {
	Options []option.RequestOption
}

// NewZoneCustomHostnameCertificatePackCertificateService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewZoneCustomHostnameCertificatePackCertificateService(opts ...option.RequestOption) (r *ZoneCustomHostnameCertificatePackCertificateService) {
	r = &ZoneCustomHostnameCertificatePackCertificateService{}
	r.Options = opts
	return
}

// Delete a single custom certificate from a certificate pack that contains two
// bundled certificates. Deletion is subject to the following constraints. You
// cannot delete a certificate if it is the only remaining certificate in the pack.
// At least one certificate must remain in the pack.
func (r *ZoneCustomHostnameCertificatePackCertificateService) Delete(ctx context.Context, zoneID string, customHostnameID string, certificatePackID string, certificateID string, body ZoneCustomHostnameCertificatePackCertificateDeleteParams, opts ...option.RequestOption) (res *ZoneCustomHostnameCertificatePackCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if customHostnameID == "" {
		err = errors.New("missing required custom_hostname_id parameter")
		return
	}
	if certificatePackID == "" {
		err = errors.New("missing required certificate_pack_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s/certificate_pack/%s/certificates/%s", zoneID, customHostnameID, certificatePackID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Replace a single custom certificate within a certificate pack that contains two
// bundled certificates. The replacement must adhere to the following constraints.
// You can only replace an RSA certificate with another RSA certificate or an ECDSA
// certificate with another ECDSA certificate.
func (r *ZoneCustomHostnameCertificatePackCertificateService) Replace(ctx context.Context, zoneID string, customHostnameID string, certificatePackID string, certificateID string, body ZoneCustomHostnameCertificatePackCertificateReplaceParams, opts ...option.RequestOption) (res *CustomHostnameResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if customHostnameID == "" {
		err = errors.New("missing required custom_hostname_id parameter")
		return
	}
	if certificatePackID == "" {
		err = errors.New("missing required certificate_pack_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s/certificate_pack/%s/certificates/%s", zoneID, customHostnameID, certificatePackID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CustomCertAndKeyParam struct {
	// If a custom uploaded certificate is used.
	CustomCertificate param.Field[string] `json:"custom_certificate,required"`
	// The key for a custom uploaded certificate.
	CustomKey param.Field[string] `json:"custom_key,required"`
}

func (r CustomCertAndKeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneCustomHostnameCertificatePackCertificateDeleteResponse struct {
	// Identifier
	ID   string                                                         `json:"id"`
	JSON zoneCustomHostnameCertificatePackCertificateDeleteResponseJSON `json:"-"`
}

// zoneCustomHostnameCertificatePackCertificateDeleteResponseJSON contains the JSON
// metadata for the struct
// [ZoneCustomHostnameCertificatePackCertificateDeleteResponse]
type zoneCustomHostnameCertificatePackCertificateDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameCertificatePackCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomHostnameCertificatePackCertificateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCustomHostnameCertificatePackCertificateDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneCustomHostnameCertificatePackCertificateDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneCustomHostnameCertificatePackCertificateReplaceParams struct {
	CustomCertAndKey CustomCertAndKeyParam `json:"custom_cert_and_key,required"`
}

func (r ZoneCustomHostnameCertificatePackCertificateReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CustomCertAndKey)
}
