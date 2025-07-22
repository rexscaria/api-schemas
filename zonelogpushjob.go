// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneLogpushJobService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneLogpushJobService] method instead.
type ZoneLogpushJobService struct {
	Options []option.RequestOption
}

// NewZoneLogpushJobService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneLogpushJobService(opts ...option.RequestOption) (r *ZoneLogpushJobService) {
	r = &ZoneLogpushJobService{}
	r.Options = opts
	return
}

// Creates a new Logpush job for a zone.
func (r *ZoneLogpushJobService) NewJob(ctx context.Context, zoneID string, body ZoneLogpushJobNewJobParams, opts ...option.RequestOption) (res *JobResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/logpush/jobs", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deletes a Logpush job.
func (r *ZoneLogpushJobService) DeleteJob(ctx context.Context, zoneID string, jobID int64, body ZoneLogpushJobDeleteJobParams, opts ...option.RequestOption) (res *ZoneLogpushJobDeleteJobResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/logpush/jobs/%v", zoneID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Lists Logpush jobs for a zone.
func (r *ZoneLogpushJobService) ListJobs(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *JobResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/logpush/jobs", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets the details of a Logpush job.
func (r *ZoneLogpushJobService) GetJob(ctx context.Context, zoneID string, jobID int64, opts ...option.RequestOption) (res *JobResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/logpush/jobs/%v", zoneID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a Logpush job.
func (r *ZoneLogpushJobService) UpdateJob(ctx context.Context, zoneID string, jobID int64, body ZoneLogpushJobUpdateJobParams, opts ...option.RequestOption) (res *JobResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/logpush/jobs/%v", zoneID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneLogpushJobDeleteJobResponse struct {
	Result ZoneLogpushJobDeleteJobResponseResult `json:"result"`
	JSON   zoneLogpushJobDeleteJobResponseJSON   `json:"-"`
	CommonResponseLogPush
}

// zoneLogpushJobDeleteJobResponseJSON contains the JSON metadata for the struct
// [ZoneLogpushJobDeleteJobResponse]
type zoneLogpushJobDeleteJobResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushJobDeleteJobResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneLogpushJobDeleteJobResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneLogpushJobDeleteJobResponseResult struct {
	// Unique id of the job.
	ID   int64                                     `json:"id"`
	JSON zoneLogpushJobDeleteJobResponseResultJSON `json:"-"`
}

// zoneLogpushJobDeleteJobResponseResultJSON contains the JSON metadata for the
// struct [ZoneLogpushJobDeleteJobResponseResult]
type zoneLogpushJobDeleteJobResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushJobDeleteJobResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneLogpushJobDeleteJobResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneLogpushJobNewJobParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
	// Name of the dataset. A list of supported datasets can be found on the
	// [Developer Docs](https://developers.cloudflare.com/logs/reference/log-fields/).
	Dataset param.Field[string] `json:"dataset"`
	// Flag that indicates if the job is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// This field is deprecated. Please use `max_upload_*` parameters instead. The
	// frequency at which Cloudflare sends batches of logs to your destination. Setting
	// frequency to high sends your logs in larger quantities of smaller files. Setting
	// frequency to low sends logs in smaller quantities of larger files.
	Frequency param.Field[Frequency] `json:"frequency"`
	// The kind parameter (optional) is used to differentiate between Logpush and Edge
	// Log Delivery jobs. Currently, Edge Log Delivery is only supported for the
	// `http_requests` dataset.
	Kind param.Field[LogpushKind] `json:"kind"`
	// This field is deprecated. Use `output_options` instead. Configuration string. It
	// specifies things like requested fields and timestamp formats. If migrating from
	// the logpull api, copy the url (full url or just the query string) of your call
	// here, and logpush will keep on making this call for you, setting start and end
	// times appropriately.
	LogpullOptions param.Field[string] `json:"logpull_options" format:"uri-reference"`
	// The maximum uncompressed file size of a batch of logs. This setting value must
	// be between `5 MB` and `1 GB`, or `0` to disable it. Note that you cannot set a
	// minimum file size; this means that log files may be much smaller than this batch
	// size. This parameter is not available for jobs with `edge` as its kind.
	MaxUploadBytes param.Field[int64] `json:"max_upload_bytes"`
	// The maximum interval in seconds for log batches. This setting must be between 30
	// and 300 seconds (5 minutes), or `0` to disable it. Note that you cannot specify
	// a minimum interval for log batches; this means that log files may be sent in
	// shorter intervals than this. This parameter is only used for jobs with `edge` as
	// its kind.
	MaxUploadIntervalSeconds param.Field[int64] `json:"max_upload_interval_seconds"`
	// The maximum number of log lines per batch. This setting must be between 1000 and
	// 1,000,000 lines, or `0` to disable it. Note that you cannot specify a minimum
	// number of log lines per batch; this means that log files may contain many fewer
	// lines than this. This parameter is not available for jobs with `edge` as its
	// kind.
	MaxUploadRecords param.Field[int64] `json:"max_upload_records"`
	// Optional human readable job name. Not unique. Cloudflare suggests that you set
	// this to a meaningful string, like the domain name, to make it easier to identify
	// your job.
	Name param.Field[string] `json:"name"`
	// The structured replacement for `logpull_options`. When including this field, the
	// `logpull_option` field will be ignored.
	OutputOptions param.Field[OutputOptionsParam] `json:"output_options"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge"`
}

func (r ZoneLogpushJobNewJobParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneLogpushJobDeleteJobParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneLogpushJobDeleteJobParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneLogpushJobUpdateJobParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf" format:"uri"`
	// Flag that indicates if the job is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// This field is deprecated. Please use `max_upload_*` parameters instead. The
	// frequency at which Cloudflare sends batches of logs to your destination. Setting
	// frequency to high sends your logs in larger quantities of smaller files. Setting
	// frequency to low sends logs in smaller quantities of larger files.
	Frequency param.Field[Frequency] `json:"frequency"`
	// The kind parameter (optional) is used to differentiate between Logpush and Edge
	// Log Delivery jobs. Currently, Edge Log Delivery is only supported for the
	// `http_requests` dataset.
	Kind param.Field[LogpushKind] `json:"kind"`
	// This field is deprecated. Use `output_options` instead. Configuration string. It
	// specifies things like requested fields and timestamp formats. If migrating from
	// the logpull api, copy the url (full url or just the query string) of your call
	// here, and logpush will keep on making this call for you, setting start and end
	// times appropriately.
	LogpullOptions param.Field[string] `json:"logpull_options" format:"uri-reference"`
	// The maximum uncompressed file size of a batch of logs. This setting value must
	// be between `5 MB` and `1 GB`, or `0` to disable it. Note that you cannot set a
	// minimum file size; this means that log files may be much smaller than this batch
	// size. This parameter is not available for jobs with `edge` as its kind.
	MaxUploadBytes param.Field[int64] `json:"max_upload_bytes"`
	// The maximum interval in seconds for log batches. This setting must be between 30
	// and 300 seconds (5 minutes), or `0` to disable it. Note that you cannot specify
	// a minimum interval for log batches; this means that log files may be sent in
	// shorter intervals than this. This parameter is only used for jobs with `edge` as
	// its kind.
	MaxUploadIntervalSeconds param.Field[int64] `json:"max_upload_interval_seconds"`
	// The maximum number of log lines per batch. This setting must be between 1000 and
	// 1,000,000 lines, or `0` to disable it. Note that you cannot specify a minimum
	// number of log lines per batch; this means that log files may contain many fewer
	// lines than this. This parameter is not available for jobs with `edge` as its
	// kind.
	MaxUploadRecords param.Field[int64] `json:"max_upload_records"`
	// Optional human readable job name. Not unique. Cloudflare suggests that you set
	// this to a meaningful string, like the domain name, to make it easier to identify
	// your job.
	Name param.Field[string] `json:"name"`
	// The structured replacement for `logpull_options`. When including this field, the
	// `logpull_option` field will be ignored.
	OutputOptions param.Field[OutputOptionsParam] `json:"output_options"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge"`
}

func (r ZoneLogpushJobUpdateJobParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
