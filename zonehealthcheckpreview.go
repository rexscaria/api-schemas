// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneHealthcheckPreviewService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneHealthcheckPreviewService] method instead.
type ZoneHealthcheckPreviewService struct {
	Options []option.RequestOption
}

// NewZoneHealthcheckPreviewService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneHealthcheckPreviewService(opts ...option.RequestOption) (r *ZoneHealthcheckPreviewService) {
	r = &ZoneHealthcheckPreviewService{}
	r.Options = opts
	return
}

// Create a new preview health check.
func (r *ZoneHealthcheckPreviewService) New(ctx context.Context, zoneID string, body ZoneHealthcheckPreviewNewParams, opts ...option.RequestOption) (res *HealthcheckSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/healthchecks/preview", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a single configured health check preview.
func (r *ZoneHealthcheckPreviewService) Get(ctx context.Context, zoneID string, healthcheckID string, opts ...option.RequestOption) (res *HealthcheckSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if healthcheckID == "" {
		err = errors.New("missing required healthcheck_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/healthchecks/preview/%s", zoneID, healthcheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a health check.
func (r *ZoneHealthcheckPreviewService) Delete(ctx context.Context, zoneID string, healthcheckID string, body ZoneHealthcheckPreviewDeleteParams, opts ...option.RequestOption) (res *HealthcheckIDResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if healthcheckID == "" {
		err = errors.New("missing required healthcheck_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/healthchecks/preview/%s", zoneID, healthcheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type ZoneHealthcheckPreviewNewParams struct {
	HealthcheckQuery HealthcheckQueryParam `json:"healthcheck_query,required"`
}

func (r ZoneHealthcheckPreviewNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.HealthcheckQuery)
}

type ZoneHealthcheckPreviewDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneHealthcheckPreviewDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
