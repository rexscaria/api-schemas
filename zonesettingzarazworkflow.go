// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneSettingZarazWorkflowService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSettingZarazWorkflowService] method instead.
type ZoneSettingZarazWorkflowService struct {
	Options []option.RequestOption
}

// NewZoneSettingZarazWorkflowService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingZarazWorkflowService(opts ...option.RequestOption) (r *ZoneSettingZarazWorkflowService) {
	r = &ZoneSettingZarazWorkflowService{}
	r.Options = opts
	return
}

// Gets Zaraz workflow for a zone.
func (r *ZoneSettingZarazWorkflowService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZarazWorkflowResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/workflow", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates Zaraz workflow for a zone.
func (r *ZoneSettingZarazWorkflowService) Update(ctx context.Context, zoneID string, body ZoneSettingZarazWorkflowUpdateParams, opts ...option.RequestOption) (res *ZarazWorkflowResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/zaraz/workflow", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Zaraz workflow
type ZarazWorkflow string

const (
	ZarazWorkflowRealtime ZarazWorkflow = "realtime"
	ZarazWorkflowPreview  ZarazWorkflow = "preview"
)

func (r ZarazWorkflow) IsKnown() bool {
	switch r {
	case ZarazWorkflowRealtime, ZarazWorkflowPreview:
		return true
	}
	return false
}

type ZarazWorkflowResponse struct {
	Errors   []ZarazMessagesItems `json:"errors,required"`
	Messages []ZarazMessagesItems `json:"messages,required"`
	// Zaraz workflow
	Result ZarazWorkflow `json:"result,required"`
	// Whether the API call was successful
	Success bool                      `json:"success,required"`
	JSON    zarazWorkflowResponseJSON `json:"-"`
}

// zarazWorkflowResponseJSON contains the JSON metadata for the struct
// [ZarazWorkflowResponse]
type zarazWorkflowResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazWorkflowResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zarazWorkflowResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingZarazWorkflowUpdateParams struct {
	// Zaraz workflow
	ZarazWorkflow ZarazWorkflow `json:"zaraz_workflow,required"`
}

func (r ZoneSettingZarazWorkflowUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ZarazWorkflow)
}
