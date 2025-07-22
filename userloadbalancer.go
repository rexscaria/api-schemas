// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// UserLoadBalancerService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserLoadBalancerService] method instead.
type UserLoadBalancerService struct {
	Options  []option.RequestOption
	Monitors *UserLoadBalancerMonitorService
	Pools    *UserLoadBalancerPoolService
}

// NewUserLoadBalancerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserLoadBalancerService(opts ...option.RequestOption) (r *UserLoadBalancerService) {
	r = &UserLoadBalancerService{}
	r.Options = opts
	r.Monitors = NewUserLoadBalancerMonitorService(opts...)
	r.Pools = NewUserLoadBalancerPoolService(opts...)
	return
}

// Get the result of a previous preview operation using the provided preview_id.
func (r *UserLoadBalancerService) PreviewResult(ctx context.Context, previewID interface{}, opts ...option.RequestOption) (res *PreviewResultResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/preview/%v", previewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
