// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// RadarLeakedCredentialCheckService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarLeakedCredentialCheckService] method instead.
type RadarLeakedCredentialCheckService struct {
	Options          []option.RequestOption
	Summary          *RadarLeakedCredentialCheckSummaryService
	TimeseriesGroups *RadarLeakedCredentialCheckTimeseriesGroupService
}

// NewRadarLeakedCredentialCheckService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarLeakedCredentialCheckService(opts ...option.RequestOption) (r *RadarLeakedCredentialCheckService) {
	r = &RadarLeakedCredentialCheckService{}
	r.Options = opts
	r.Summary = NewRadarLeakedCredentialCheckSummaryService(opts...)
	r.TimeseriesGroups = NewRadarLeakedCredentialCheckTimeseriesGroupService(opts...)
	return
}
