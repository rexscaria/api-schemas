// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// ZoneLogpushService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneLogpushService] method instead.
type ZoneLogpushService struct {
	Options   []option.RequestOption
	Datasets  *ZoneLogpushDatasetService
	Edge      *ZoneLogpushEdgeService
	Jobs      *ZoneLogpushJobService
	Ownership *ZoneLogpushOwnershipService
	Validate  *ZoneLogpushValidateService
}

// NewZoneLogpushService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneLogpushService(opts ...option.RequestOption) (r *ZoneLogpushService) {
	r = &ZoneLogpushService{}
	r.Options = opts
	r.Datasets = NewZoneLogpushDatasetService(opts...)
	r.Edge = NewZoneLogpushEdgeService(opts...)
	r.Jobs = NewZoneLogpushJobService(opts...)
	r.Ownership = NewZoneLogpushOwnershipService(opts...)
	r.Validate = NewZoneLogpushValidateService(opts...)
	return
}
