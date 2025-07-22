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

// UserLoadBalancerMonitorService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserLoadBalancerMonitorService] method instead.
type UserLoadBalancerMonitorService struct {
	Options []option.RequestOption
}

// NewUserLoadBalancerMonitorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserLoadBalancerMonitorService(opts ...option.RequestOption) (r *UserLoadBalancerMonitorService) {
	r = &UserLoadBalancerMonitorService{}
	r.Options = opts
	return
}

// Create a configured monitor.
func (r *UserLoadBalancerMonitorService) New(ctx context.Context, body UserLoadBalancerMonitorNewParams, opts ...option.RequestOption) (res *ResponseSingleMonitor, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/load_balancers/monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List a single configured monitor for a user.
func (r *UserLoadBalancerMonitorService) Get(ctx context.Context, monitorID string, opts ...option.RequestOption) (res *ResponseSingleMonitor, err error) {
	opts = append(r.Options[:], opts...)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	path := fmt.Sprintf("user/load_balancers/monitors/%s", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify a configured monitor.
func (r *UserLoadBalancerMonitorService) Update(ctx context.Context, monitorID string, body UserLoadBalancerMonitorUpdateParams, opts ...option.RequestOption) (res *ResponseSingleMonitor, err error) {
	opts = append(r.Options[:], opts...)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	path := fmt.Sprintf("user/load_balancers/monitors/%s", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List configured monitors for a user.
func (r *UserLoadBalancerMonitorService) List(ctx context.Context, opts ...option.RequestOption) (res *ResponseCollectionMonitor, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/load_balancers/monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a configured monitor.
func (r *UserLoadBalancerMonitorService) Delete(ctx context.Context, monitorID string, body UserLoadBalancerMonitorDeleteParams, opts ...option.RequestOption) (res *IDResponseLoadBalancing, err error) {
	opts = append(r.Options[:], opts...)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	path := fmt.Sprintf("user/load_balancers/monitors/%s", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Get the list of resources that reference the provided monitor.
func (r *UserLoadBalancerMonitorService) ListReferences(ctx context.Context, monitorID string, opts ...option.RequestOption) (res *ReferencesMonitorResponse, err error) {
	opts = append(r.Options[:], opts...)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	path := fmt.Sprintf("user/load_balancers/monitors/%s/references", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Apply changes to an existing monitor, overwriting the supplied properties.
func (r *UserLoadBalancerMonitorService) Patch(ctx context.Context, monitorID string, body UserLoadBalancerMonitorPatchParams, opts ...option.RequestOption) (res *ResponseSingleMonitor, err error) {
	opts = append(r.Options[:], opts...)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	path := fmt.Sprintf("user/load_balancers/monitors/%s", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Preview pools using the specified monitor with provided monitor details. The
// returned preview_id can be used in the preview endpoint to retrieve the results.
func (r *UserLoadBalancerMonitorService) Preview(ctx context.Context, monitorID string, body UserLoadBalancerMonitorPreviewParams, opts ...option.RequestOption) (res *PreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	path := fmt.Sprintf("user/load_balancers/monitors/%s/preview", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type UserLoadBalancerMonitorNewParams struct {
	EditableMonitor EditableMonitorParam `json:"editable_monitor,required"`
}

func (r UserLoadBalancerMonitorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.EditableMonitor)
}

type UserLoadBalancerMonitorUpdateParams struct {
	EditableMonitor EditableMonitorParam `json:"editable_monitor,required"`
}

func (r UserLoadBalancerMonitorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.EditableMonitor)
}

type UserLoadBalancerMonitorDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r UserLoadBalancerMonitorDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type UserLoadBalancerMonitorPatchParams struct {
	EditableMonitor EditableMonitorParam `json:"editable_monitor,required"`
}

func (r UserLoadBalancerMonitorPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.EditableMonitor)
}

type UserLoadBalancerMonitorPreviewParams struct {
	EditableMonitor EditableMonitorParam `json:"editable_monitor,required"`
}

func (r UserLoadBalancerMonitorPreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.EditableMonitor)
}
