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

// ZoneAccessCertificateService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAccessCertificateService] method instead.
type ZoneAccessCertificateService struct {
	Options  []option.RequestOption
	Settings *ZoneAccessCertificateSettingService
}

// NewZoneAccessCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAccessCertificateService(opts ...option.RequestOption) (r *ZoneAccessCertificateService) {
	r = &ZoneAccessCertificateService{}
	r.Options = opts
	r.Settings = NewZoneAccessCertificateSettingService(opts...)
	return
}

// Adds a new mTLS root certificate to Access.
func (r *ZoneAccessCertificateService) New(ctx context.Context, zoneID string, body ZoneAccessCertificateNewParams, opts ...option.RequestOption) (res *SingleResponseCertificateZone, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/certificates", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single mTLS certificate.
func (r *ZoneAccessCertificateService) Get(ctx context.Context, zoneID string, certificateID string, opts ...option.RequestOption) (res *SingleResponseCertificateZone, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/certificates/%s", zoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured mTLS certificate.
func (r *ZoneAccessCertificateService) Update(ctx context.Context, zoneID string, certificateID string, body ZoneAccessCertificateUpdateParams, opts ...option.RequestOption) (res *SingleResponseCertificateZone, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/certificates/%s", zoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all mTLS certificates.
func (r *ZoneAccessCertificateService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneAccessCertificateListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/certificates", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an mTLS certificate.
func (r *ZoneAccessCertificateService) Delete(ctx context.Context, zoneID string, certificateID string, opts ...option.RequestOption) (res *IDResponseCertificates, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/certificates/%s", zoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccessComponentsCertificates struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                           `json:"name"`
	UpdatedAt time.Time                        `json:"updated_at" format:"date-time"`
	JSON      accessComponentsCertificatesJSON `json:"-"`
}

// accessComponentsCertificatesJSON contains the JSON metadata for the struct
// [AccessComponentsCertificates]
type accessComponentsCertificatesJSON struct {
	ID                  apijson.Field
	AssociatedHostnames apijson.Field
	CreatedAt           apijson.Field
	ExpiresOn           apijson.Field
	Fingerprint         apijson.Field
	Name                apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessComponentsCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessComponentsCertificatesJSON) RawJSON() string {
	return r.raw
}

type SingleResponseCertificateZone struct {
	Result AccessComponentsCertificates      `json:"result"`
	JSON   singleResponseCertificateZoneJSON `json:"-"`
	APIResponseSingleAccess
}

// singleResponseCertificateZoneJSON contains the JSON metadata for the struct
// [SingleResponseCertificateZone]
type singleResponseCertificateZoneJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseCertificateZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseCertificateZoneJSON) RawJSON() string {
	return r.raw
}

type ZoneAccessCertificateListResponse struct {
	Result []AccessComponentsCertificates        `json:"result"`
	JSON   zoneAccessCertificateListResponseJSON `json:"-"`
	APIResponseCollectionAccess
}

// zoneAccessCertificateListResponseJSON contains the JSON metadata for the struct
// [ZoneAccessCertificateListResponse]
type zoneAccessCertificateListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAccessCertificateListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAccessCertificateNewParams struct {
	// The certificate content.
	Certificate param.Field[string] `json:"certificate,required"`
	// The name of the certificate.
	Name param.Field[string] `json:"name,required"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]string] `json:"associated_hostnames"`
}

func (r ZoneAccessCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessCertificateUpdateParams struct {
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]string] `json:"associated_hostnames,required"`
	// The name of the certificate.
	Name param.Field[string] `json:"name"`
}

func (r ZoneAccessCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
