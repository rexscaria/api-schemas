// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneAccessAppCaService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAccessAppCaService] method instead.
type ZoneAccessAppCaService struct {
	Options []option.RequestOption
}

// NewZoneAccessAppCaService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneAccessAppCaService(opts ...option.RequestOption) (r *ZoneAccessAppCaService) {
	r = &ZoneAccessAppCaService{}
	r.Options = opts
	return
}

// Generates a new short-lived certificate CA and public key.
func (r *ZoneAccessAppCaService) New(ctx context.Context, zoneID string, appID string, opts ...option.RequestOption) (res *SingleResponseCa, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/apps/%s/ca", zoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Fetches a short-lived certificate CA and its public key.
func (r *ZoneAccessAppCaService) Get(ctx context.Context, zoneID string, appID string, opts ...option.RequestOption) (res *SingleResponseCa, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/apps/%s/ca", zoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists short-lived certificate CAs and their public keys.
func (r *ZoneAccessAppCaService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ResponseCollectionCa, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/apps/ca", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a short-lived certificate CA.
func (r *ZoneAccessAppCaService) Delete(ctx context.Context, zoneID string, appID string, opts ...option.RequestOption) (res *SchemasIDResponseAccess, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/access/apps/%s/ca", zoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}
