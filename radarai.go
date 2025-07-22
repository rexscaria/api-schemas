// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// RadarAIService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAIService] method instead.
type RadarAIService struct {
	Options   []option.RequestOption
	Bots      *RadarAIBotService
	Inference *RadarAIInferenceService
}

// NewRadarAIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRadarAIService(opts ...option.RequestOption) (r *RadarAIService) {
	r = &RadarAIService{}
	r.Options = opts
	r.Bots = NewRadarAIBotService(opts...)
	r.Inference = NewRadarAIInferenceService(opts...)
	return
}
