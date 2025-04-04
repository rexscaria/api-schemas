// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneLogpushEdgeService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneLogpushEdgeService] method instead.
type ZoneLogpushEdgeService struct {
	Options []option.RequestOption
}

// NewZoneLogpushEdgeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneLogpushEdgeService(opts ...option.RequestOption) (r *ZoneLogpushEdgeService) {
	r = &ZoneLogpushEdgeService{}
	r.Options = opts
	return
}

// Creates a new Instant Logs job for a zone.
func (r *ZoneLogpushEdgeService) NewJob(ctx context.Context, zoneID string, body ZoneLogpushEdgeNewJobParams, opts ...option.RequestOption) (res *ZoneLogpushEdgeNewJobResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/logpush/edge", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists Instant Logs jobs for a zone.
func (r *ZoneLogpushEdgeService) ListJobs(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneLogpushEdgeListJobsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/logpush/edge", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type InstantLogsJob struct {
	// Unique WebSocket address that will receive messages from Cloudflare’s edge.
	DestinationConf string `json:"destination_conf" format:"uri"`
	// Comma-separated list of fields.
	Fields string `json:"fields"`
	// Filters to drill down into specific events.
	Filter string `json:"filter"`
	// The sample parameter is the sample rate of the records set by the client:
	// "sample": 1 is 100% of records "sample": 10 is 10% and so on.
	Sample int64 `json:"sample"`
	// Unique session id of the job.
	SessionID string             `json:"session_id"`
	JSON      instantLogsJobJSON `json:"-"`
}

// instantLogsJobJSON contains the JSON metadata for the struct [InstantLogsJob]
type instantLogsJobJSON struct {
	DestinationConf apijson.Field
	Fields          apijson.Field
	Filter          apijson.Field
	Sample          apijson.Field
	SessionID       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InstantLogsJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instantLogsJobJSON) RawJSON() string {
	return r.raw
}

type ZoneLogpushEdgeNewJobResponse struct {
	Result InstantLogsJob                    `json:"result,nullable"`
	JSON   zoneLogpushEdgeNewJobResponseJSON `json:"-"`
	SingleResponseJob
}

// zoneLogpushEdgeNewJobResponseJSON contains the JSON metadata for the struct
// [ZoneLogpushEdgeNewJobResponse]
type zoneLogpushEdgeNewJobResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushEdgeNewJobResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneLogpushEdgeNewJobResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneLogpushEdgeListJobsResponse struct {
	Result []InstantLogsJob                    `json:"result"`
	JSON   zoneLogpushEdgeListJobsResponseJSON `json:"-"`
	CommonResponseLogPush
}

// zoneLogpushEdgeListJobsResponseJSON contains the JSON metadata for the struct
// [ZoneLogpushEdgeListJobsResponse]
type zoneLogpushEdgeListJobsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushEdgeListJobsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneLogpushEdgeListJobsResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneLogpushEdgeNewJobParams struct {
	// Comma-separated list of fields.
	Fields param.Field[string] `json:"fields"`
	// Filters to drill down into specific events.
	Filter param.Field[string] `json:"filter"`
	// The sample parameter is the sample rate of the records set by the client:
	// "sample": 1 is 100% of records "sample": 10 is 10% and so on.
	Sample param.Field[int64] `json:"sample"`
}

func (r ZoneLogpushEdgeNewJobParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
