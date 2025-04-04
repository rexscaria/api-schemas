// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/stainless-sdks/cf-rex-go/option"
)

// RadarService contains methods and other services that help with interacting with
// the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarService] method instead.
type RadarService struct {
	Options                []option.RequestOption
	AI                     *RadarAIService
	Annotations            *RadarAnnotationService
	As112                  *RadarAs112Service
	Attacks                *RadarAttackService
	Bgp                    *RadarBgpService
	Datasets               *RadarDatasetService
	DNS                    *RadarDNSService
	Email                  *RadarEmailService
	Entities               *RadarEntityService
	HTTP                   *RadarHTTPService
	LeakedCredentialChecks *RadarLeakedCredentialCheckService
	Netflows               *RadarNetflowService
	Quality                *RadarQualityService
	Ranking                *RadarRankingService
	RobotsTxt              *RadarRobotsTxtService
	Search                 *RadarSearchService
	TcpResetsTimeouts      *RadarTcpResetsTimeoutService
	TrafficAnomalies       *RadarTrafficAnomalyService
	VerifiedBots           *RadarVerifiedBotService
}

// NewRadarService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRadarService(opts ...option.RequestOption) (r *RadarService) {
	r = &RadarService{}
	r.Options = opts
	r.AI = NewRadarAIService(opts...)
	r.Annotations = NewRadarAnnotationService(opts...)
	r.As112 = NewRadarAs112Service(opts...)
	r.Attacks = NewRadarAttackService(opts...)
	r.Bgp = NewRadarBgpService(opts...)
	r.Datasets = NewRadarDatasetService(opts...)
	r.DNS = NewRadarDNSService(opts...)
	r.Email = NewRadarEmailService(opts...)
	r.Entities = NewRadarEntityService(opts...)
	r.HTTP = NewRadarHTTPService(opts...)
	r.LeakedCredentialChecks = NewRadarLeakedCredentialCheckService(opts...)
	r.Netflows = NewRadarNetflowService(opts...)
	r.Quality = NewRadarQualityService(opts...)
	r.Ranking = NewRadarRankingService(opts...)
	r.RobotsTxt = NewRadarRobotsTxtService(opts...)
	r.Search = NewRadarSearchService(opts...)
	r.TcpResetsTimeouts = NewRadarTcpResetsTimeoutService(opts...)
	r.TrafficAnomalies = NewRadarTrafficAnomalyService(opts...)
	r.VerifiedBots = NewRadarVerifiedBotService(opts...)
	return
}
