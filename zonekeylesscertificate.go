// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneKeylessCertificateService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneKeylessCertificateService] method instead.
type ZoneKeylessCertificateService struct {
	Options []option.RequestOption
}

// NewZoneKeylessCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneKeylessCertificateService(opts ...option.RequestOption) (r *ZoneKeylessCertificateService) {
	r = &ZoneKeylessCertificateService{}
	r.Options = opts
	return
}

// Create Keyless SSL Configuration
func (r *ZoneKeylessCertificateService) New(ctx context.Context, zoneID string, body ZoneKeylessCertificateNewParams, opts ...option.RequestOption) (res *KeylessResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/keyless_certificates", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details for one Keyless SSL configuration.
func (r *ZoneKeylessCertificateService) Get(ctx context.Context, zoneID string, keylessCertificateID string, opts ...option.RequestOption) (res *KeylessResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if keylessCertificateID == "" {
		err = errors.New("missing required keyless_certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/keyless_certificates/%s", zoneID, keylessCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This will update attributes of a Keyless SSL. Consists of one or more of the
// following: host,name,port.
func (r *ZoneKeylessCertificateService) Update(ctx context.Context, zoneID string, keylessCertificateID string, body ZoneKeylessCertificateUpdateParams, opts ...option.RequestOption) (res *KeylessResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if keylessCertificateID == "" {
		err = errors.New("missing required keyless_certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/keyless_certificates/%s", zoneID, keylessCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all Keyless SSL configurations for a given zone.
func (r *ZoneKeylessCertificateService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneKeylessCertificateListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/keyless_certificates", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Keyless SSL Configuration
func (r *ZoneKeylessCertificateService) Delete(ctx context.Context, zoneID string, keylessCertificateID string, body ZoneKeylessCertificateDeleteParams, opts ...option.RequestOption) (res *ZoneKeylessCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if keylessCertificateID == "" {
		err = errors.New("missing required keyless_certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/keyless_certificates/%s", zoneID, keylessCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type KeylessCertificate struct {
	// Keyless certificate identifier tag.
	ID string `json:"id,required"`
	// When the Keyless SSL was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Whether or not the Keyless SSL is on or off.
	Enabled bool `json:"enabled,required"`
	// The keyless SSL name.
	Host string `json:"host,required" format:"hostname"`
	// When the Keyless SSL was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The keyless SSL name.
	Name string `json:"name,required"`
	// Available permissions for the Keyless SSL for the current user requesting the
	// item.
	Permissions []string `json:"permissions,required"`
	// The keyless SSL port used to communicate between Cloudflare and the client's
	// Keyless SSL server.
	Port float64 `json:"port,required"`
	// Status of the Keyless SSL.
	Status KeylessCertificateStatus `json:"status,required"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel KeylessTunnel          `json:"tunnel"`
	JSON   keylessCertificateJSON `json:"-"`
}

// keylessCertificateJSON contains the JSON metadata for the struct
// [KeylessCertificate]
type keylessCertificateJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Enabled     apijson.Field
	Host        apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	Port        apijson.Field
	Status      apijson.Field
	Tunnel      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateJSON) RawJSON() string {
	return r.raw
}

// Status of the Keyless SSL.
type KeylessCertificateStatus string

const (
	KeylessCertificateStatusActive  KeylessCertificateStatus = "active"
	KeylessCertificateStatusDeleted KeylessCertificateStatus = "deleted"
)

func (r KeylessCertificateStatus) IsKnown() bool {
	switch r {
	case KeylessCertificateStatusActive, KeylessCertificateStatusDeleted:
		return true
	}
	return false
}

type KeylessResponseSingle struct {
	Result TlsCertificateBase        `json:"result"`
	JSON   keylessResponseSingleJSON `json:"-"`
	APIResponseSingleTlsCertificates
}

// keylessResponseSingleJSON contains the JSON metadata for the struct
// [KeylessResponseSingle]
type keylessResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessResponseSingleJSON) RawJSON() string {
	return r.raw
}

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type KeylessTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP string `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID string            `json:"vnet_id,required"`
	JSON   keylessTunnelJSON `json:"-"`
}

// keylessTunnelJSON contains the JSON metadata for the struct [KeylessTunnel]
type keylessTunnelJSON struct {
	PrivateIP   apijson.Field
	VnetID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessTunnelJSON) RawJSON() string {
	return r.raw
}

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type KeylessTunnelParam struct {
	// Private IP of the Key Server Host
	PrivateIP param.Field[string] `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID param.Field[string] `json:"vnet_id,required"`
}

func (r KeylessTunnelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TlsCertificateBase struct {
	// Keyless certificate identifier tag.
	ID string `json:"id,required"`
	// When the Keyless SSL was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Whether or not the Keyless SSL is on or off.
	Enabled bool `json:"enabled,required"`
	// The keyless SSL name.
	Host string `json:"host,required" format:"hostname"`
	// When the Keyless SSL was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The keyless SSL name.
	Name string `json:"name,required"`
	// Available permissions for the Keyless SSL for the current user requesting the
	// item.
	Permissions []string `json:"permissions,required"`
	// The keyless SSL port used to communicate between Cloudflare and the client's
	// Keyless SSL server.
	Port float64 `json:"port,required"`
	// Status of the Keyless SSL.
	Status TlsCertificateBaseStatus `json:"status,required"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel KeylessTunnel          `json:"tunnel"`
	JSON   tlsCertificateBaseJSON `json:"-"`
}

// tlsCertificateBaseJSON contains the JSON metadata for the struct
// [TlsCertificateBase]
type tlsCertificateBaseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Enabled     apijson.Field
	Host        apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	Port        apijson.Field
	Status      apijson.Field
	Tunnel      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TlsCertificateBase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsCertificateBaseJSON) RawJSON() string {
	return r.raw
}

// Status of the Keyless SSL.
type TlsCertificateBaseStatus string

const (
	TlsCertificateBaseStatusActive  TlsCertificateBaseStatus = "active"
	TlsCertificateBaseStatusDeleted TlsCertificateBaseStatus = "deleted"
)

func (r TlsCertificateBaseStatus) IsKnown() bool {
	switch r {
	case TlsCertificateBaseStatusActive, TlsCertificateBaseStatusDeleted:
		return true
	}
	return false
}

type ZoneKeylessCertificateListResponse struct {
	Result []KeylessCertificate                   `json:"result"`
	JSON   zoneKeylessCertificateListResponseJSON `json:"-"`
	APIResponseCollectionTlsCertificates
}

// zoneKeylessCertificateListResponseJSON contains the JSON metadata for the struct
// [ZoneKeylessCertificateListResponse]
type zoneKeylessCertificateListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneKeylessCertificateListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneKeylessCertificateDeleteResponse struct {
	Result ZoneKeylessCertificateDeleteResponseResult `json:"result"`
	JSON   zoneKeylessCertificateDeleteResponseJSON   `json:"-"`
	APIResponseSingleTlsCertificates
}

// zoneKeylessCertificateDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneKeylessCertificateDeleteResponse]
type zoneKeylessCertificateDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneKeylessCertificateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneKeylessCertificateDeleteResponseResult struct {
	// Identifier
	ID   string                                         `json:"id"`
	JSON zoneKeylessCertificateDeleteResponseResultJSON `json:"-"`
}

// zoneKeylessCertificateDeleteResponseResultJSON contains the JSON metadata for
// the struct [ZoneKeylessCertificateDeleteResponseResult]
type zoneKeylessCertificateDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneKeylessCertificateDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneKeylessCertificateNewParams struct {
	// The zone's SSL certificate or SSL certificate and intermediate(s).
	Certificate param.Field[string] `json:"certificate,required"`
	// The keyless SSL name.
	Host param.Field[string] `json:"host,required" format:"hostname"`
	// The keyless SSL port used to communicate between Cloudflare and the client's
	// Keyless SSL server.
	Port param.Field[float64] `json:"port,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[BundleMethod] `json:"bundle_method"`
	// The keyless SSL name.
	Name param.Field[string] `json:"name"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel param.Field[KeylessTunnelParam] `json:"tunnel"`
}

func (r ZoneKeylessCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneKeylessCertificateUpdateParams struct {
	// Whether or not the Keyless SSL is on or off.
	Enabled param.Field[bool] `json:"enabled"`
	// The keyless SSL name.
	Host param.Field[string] `json:"host" format:"hostname"`
	// The keyless SSL name.
	Name param.Field[string] `json:"name"`
	// The keyless SSL port used to communicate between Cloudflare and the client's
	// Keyless SSL server.
	Port param.Field[float64] `json:"port"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel param.Field[KeylessTunnelParam] `json:"tunnel"`
}

func (r ZoneKeylessCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneKeylessCertificateDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneKeylessCertificateDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
