// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountSlurperJobService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountSlurperJobService] method instead.
type AccountSlurperJobService struct {
	Options []option.RequestOption
}

// NewAccountSlurperJobService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountSlurperJobService(opts ...option.RequestOption) (r *AccountSlurperJobService) {
	r = &AccountSlurperJobService{}
	r.Options = opts
	return
}

// Create a job
func (r *AccountSlurperJobService) New(ctx context.Context, accountID string, body AccountSlurperJobNewParams, opts ...option.RequestOption) (res *AccountSlurperJobNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/slurper/jobs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get job details
func (r *AccountSlurperJobService) Get(ctx context.Context, accountID string, jobID string, opts ...option.RequestOption) (res *AccountSlurperJobGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/slurper/jobs/%s", accountID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List jobs
func (r *AccountSlurperJobService) List(ctx context.Context, accountID string, query AccountSlurperJobListParams, opts ...option.RequestOption) (res *AccountSlurperJobListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/slurper/jobs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Abort a job
func (r *AccountSlurperJobService) Abort(ctx context.Context, accountID string, jobID string, opts ...option.RequestOption) (res *AccountSlurperJobAbortResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/slurper/jobs/%s/abort", accountID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Abort all jobs
func (r *AccountSlurperJobService) AbortAll(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountSlurperJobAbortAllResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/slurper/jobs/abortAll", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Get job logs
func (r *AccountSlurperJobService) GetLogs(ctx context.Context, accountID string, jobID string, query AccountSlurperJobGetLogsParams, opts ...option.RequestOption) (res *AccountSlurperJobGetLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/slurper/jobs/%s/logs", accountID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get job progress
func (r *AccountSlurperJobService) GetProgress(ctx context.Context, accountID string, jobID string, opts ...option.RequestOption) (res *AccountSlurperJobGetProgressResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/slurper/jobs/%s/progress", accountID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Pause a job
func (r *AccountSlurperJobService) Pause(ctx context.Context, accountID string, jobID string, opts ...option.RequestOption) (res *AccountSlurperJobPauseResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/slurper/jobs/%s/pause", accountID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Resume a job
func (r *AccountSlurperJobService) Resume(ctx context.Context, accountID string, jobID string, opts ...option.RequestOption) (res *AccountSlurperJobResumeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/slurper/jobs/%s/resume", accountID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

type JobResponse struct {
	ID         string            `json:"id"`
	CreatedAt  string            `json:"createdAt"`
	FinishedAt string            `json:"finishedAt,nullable"`
	Overwrite  bool              `json:"overwrite"`
	Source     JobResponseSource `json:"source"`
	Status     JobStatus         `json:"status"`
	Target     JobResponseTarget `json:"target"`
	JSON       jobResponseJSON   `json:"-"`
}

// jobResponseJSON contains the JSON metadata for the struct [JobResponse]
type jobResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	FinishedAt  apijson.Field
	Overwrite   apijson.Field
	Source      apijson.Field
	Status      apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobResponseJSON) RawJSON() string {
	return r.raw
}

type JobResponseSource struct {
	Bucket       string                  `json:"bucket"`
	Endpoint     string                  `json:"endpoint,nullable"`
	Jurisdiction Jurisdiction            `json:"jurisdiction"`
	PathPrefix   string                  `json:"pathPrefix,nullable"`
	Vendor       JobResponseSourceVendor `json:"vendor"`
	JSON         jobResponseSourceJSON   `json:"-"`
	union        JobResponseSourceUnion
}

// jobResponseSourceJSON contains the JSON metadata for the struct
// [JobResponseSource]
type jobResponseSourceJSON struct {
	Bucket       apijson.Field
	Endpoint     apijson.Field
	Jurisdiction apijson.Field
	PathPrefix   apijson.Field
	Vendor       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r jobResponseSourceJSON) RawJSON() string {
	return r.raw
}

func (r *JobResponseSource) UnmarshalJSON(data []byte) (err error) {
	*r = JobResponseSource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [JobResponseSourceUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [JobResponseSourceS3SourceResponseSchema],
// [JobResponseSourceGcsSourceResponseSchema],
// [JobResponseSourceR2SourceResponseSchema].
func (r JobResponseSource) AsUnion() JobResponseSourceUnion {
	return r.union
}

// Union satisfied by [JobResponseSourceS3SourceResponseSchema],
// [JobResponseSourceGcsSourceResponseSchema] or
// [JobResponseSourceR2SourceResponseSchema].
type JobResponseSourceUnion interface {
	implementsJobResponseSource()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*JobResponseSourceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(JobResponseSourceS3SourceResponseSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(JobResponseSourceGcsSourceResponseSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(JobResponseSourceR2SourceResponseSchema{}),
		},
	)
}

type JobResponseSourceS3SourceResponseSchema struct {
	Bucket     string                                        `json:"bucket"`
	Endpoint   string                                        `json:"endpoint,nullable"`
	PathPrefix string                                        `json:"pathPrefix,nullable"`
	Vendor     JobResponseSourceS3SourceResponseSchemaVendor `json:"vendor"`
	JSON       jobResponseSourceS3SourceResponseSchemaJSON   `json:"-"`
}

// jobResponseSourceS3SourceResponseSchemaJSON contains the JSON metadata for the
// struct [JobResponseSourceS3SourceResponseSchema]
type jobResponseSourceS3SourceResponseSchemaJSON struct {
	Bucket      apijson.Field
	Endpoint    apijson.Field
	PathPrefix  apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobResponseSourceS3SourceResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobResponseSourceS3SourceResponseSchemaJSON) RawJSON() string {
	return r.raw
}

func (r JobResponseSourceS3SourceResponseSchema) implementsJobResponseSource() {}

type JobResponseSourceS3SourceResponseSchemaVendor string

const (
	JobResponseSourceS3SourceResponseSchemaVendorS3 JobResponseSourceS3SourceResponseSchemaVendor = "s3"
)

func (r JobResponseSourceS3SourceResponseSchemaVendor) IsKnown() bool {
	switch r {
	case JobResponseSourceS3SourceResponseSchemaVendorS3:
		return true
	}
	return false
}

type JobResponseSourceGcsSourceResponseSchema struct {
	Bucket     string                                         `json:"bucket"`
	PathPrefix string                                         `json:"pathPrefix,nullable"`
	Vendor     JobResponseSourceGcsSourceResponseSchemaVendor `json:"vendor"`
	JSON       jobResponseSourceGcsSourceResponseSchemaJSON   `json:"-"`
}

// jobResponseSourceGcsSourceResponseSchemaJSON contains the JSON metadata for the
// struct [JobResponseSourceGcsSourceResponseSchema]
type jobResponseSourceGcsSourceResponseSchemaJSON struct {
	Bucket      apijson.Field
	PathPrefix  apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobResponseSourceGcsSourceResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobResponseSourceGcsSourceResponseSchemaJSON) RawJSON() string {
	return r.raw
}

func (r JobResponseSourceGcsSourceResponseSchema) implementsJobResponseSource() {}

type JobResponseSourceGcsSourceResponseSchemaVendor string

const (
	JobResponseSourceGcsSourceResponseSchemaVendorGcs JobResponseSourceGcsSourceResponseSchemaVendor = "gcs"
)

func (r JobResponseSourceGcsSourceResponseSchemaVendor) IsKnown() bool {
	switch r {
	case JobResponseSourceGcsSourceResponseSchemaVendorGcs:
		return true
	}
	return false
}

type JobResponseSourceR2SourceResponseSchema struct {
	Bucket       string                                        `json:"bucket"`
	Jurisdiction Jurisdiction                                  `json:"jurisdiction"`
	PathPrefix   string                                        `json:"pathPrefix,nullable"`
	Vendor       JobResponseSourceR2SourceResponseSchemaVendor `json:"vendor"`
	JSON         jobResponseSourceR2SourceResponseSchemaJSON   `json:"-"`
}

// jobResponseSourceR2SourceResponseSchemaJSON contains the JSON metadata for the
// struct [JobResponseSourceR2SourceResponseSchema]
type jobResponseSourceR2SourceResponseSchemaJSON struct {
	Bucket       apijson.Field
	Jurisdiction apijson.Field
	PathPrefix   apijson.Field
	Vendor       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *JobResponseSourceR2SourceResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobResponseSourceR2SourceResponseSchemaJSON) RawJSON() string {
	return r.raw
}

func (r JobResponseSourceR2SourceResponseSchema) implementsJobResponseSource() {}

type JobResponseSourceR2SourceResponseSchemaVendor string

const (
	JobResponseSourceR2SourceResponseSchemaVendorR2 JobResponseSourceR2SourceResponseSchemaVendor = "r2"
)

func (r JobResponseSourceR2SourceResponseSchemaVendor) IsKnown() bool {
	switch r {
	case JobResponseSourceR2SourceResponseSchemaVendorR2:
		return true
	}
	return false
}

type JobResponseSourceVendor string

const (
	JobResponseSourceVendorS3  JobResponseSourceVendor = "s3"
	JobResponseSourceVendorGcs JobResponseSourceVendor = "gcs"
	JobResponseSourceVendorR2  JobResponseSourceVendor = "r2"
)

func (r JobResponseSourceVendor) IsKnown() bool {
	switch r {
	case JobResponseSourceVendorS3, JobResponseSourceVendorGcs, JobResponseSourceVendorR2:
		return true
	}
	return false
}

type JobResponseTarget struct {
	Bucket       string                  `json:"bucket"`
	Jurisdiction Jurisdiction            `json:"jurisdiction"`
	Vendor       JobResponseTargetVendor `json:"vendor"`
	JSON         jobResponseTargetJSON   `json:"-"`
}

// jobResponseTargetJSON contains the JSON metadata for the struct
// [JobResponseTarget]
type jobResponseTargetJSON struct {
	Bucket       apijson.Field
	Jurisdiction apijson.Field
	Vendor       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *JobResponseTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobResponseTargetJSON) RawJSON() string {
	return r.raw
}

type JobResponseTargetVendor string

const (
	JobResponseTargetVendorR2 JobResponseTargetVendor = "r2"
)

func (r JobResponseTargetVendor) IsKnown() bool {
	switch r {
	case JobResponseTargetVendorR2:
		return true
	}
	return false
}

type JobStatus string

const (
	JobStatusRunning   JobStatus = "running"
	JobStatusPaused    JobStatus = "paused"
	JobStatusAborted   JobStatus = "aborted"
	JobStatusCompleted JobStatus = "completed"
)

func (r JobStatus) IsKnown() bool {
	switch r {
	case JobStatusRunning, JobStatusPaused, JobStatusAborted, JobStatusCompleted:
		return true
	}
	return false
}

type Jurisdiction string

const (
	JurisdictionDefault Jurisdiction = "default"
	JurisdictionEu      Jurisdiction = "eu"
	JurisdictionFedramp Jurisdiction = "fedramp"
)

func (r Jurisdiction) IsKnown() bool {
	switch r {
	case JurisdictionDefault, JurisdictionEu, JurisdictionFedramp:
		return true
	}
	return false
}

type AccountSlurperJobNewResponse struct {
	Errors   []AccountSlurperJobNewResponseError `json:"errors"`
	Messages []string                            `json:"messages"`
	Result   AccountSlurperJobNewResponseResult  `json:"result"`
	// Indicates if the API call was successful or not.
	Success AccountSlurperJobNewResponseSuccess `json:"success"`
	JSON    accountSlurperJobNewResponseJSON    `json:"-"`
}

// accountSlurperJobNewResponseJSON contains the JSON metadata for the struct
// [AccountSlurperJobNewResponse]
type accountSlurperJobNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobNewResponseError struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           AccountSlurperJobNewResponseErrorsSource `json:"source"`
	JSON             accountSlurperJobNewResponseErrorJSON    `json:"-"`
}

// accountSlurperJobNewResponseErrorJSON contains the JSON metadata for the struct
// [AccountSlurperJobNewResponseError]
type accountSlurperJobNewResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountSlurperJobNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobNewResponseErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    accountSlurperJobNewResponseErrorsSourceJSON `json:"-"`
}

// accountSlurperJobNewResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountSlurperJobNewResponseErrorsSource]
type accountSlurperJobNewResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobNewResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobNewResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobNewResponseResult struct {
	ID   string                                 `json:"id"`
	JSON accountSlurperJobNewResponseResultJSON `json:"-"`
}

// accountSlurperJobNewResponseResultJSON contains the JSON metadata for the struct
// [AccountSlurperJobNewResponseResult]
type accountSlurperJobNewResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobNewResponseResultJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type AccountSlurperJobNewResponseSuccess bool

const (
	AccountSlurperJobNewResponseSuccessTrue AccountSlurperJobNewResponseSuccess = true
)

func (r AccountSlurperJobNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountSlurperJobNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountSlurperJobGetResponse struct {
	Errors   []AccountSlurperJobGetResponseError `json:"errors"`
	Messages []string                            `json:"messages"`
	Result   JobResponse                         `json:"result"`
	// Indicates if the API call was successful or not.
	Success AccountSlurperJobGetResponseSuccess `json:"success"`
	JSON    accountSlurperJobGetResponseJSON    `json:"-"`
}

// accountSlurperJobGetResponseJSON contains the JSON metadata for the struct
// [AccountSlurperJobGetResponse]
type accountSlurperJobGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobGetResponseError struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           AccountSlurperJobGetResponseErrorsSource `json:"source"`
	JSON             accountSlurperJobGetResponseErrorJSON    `json:"-"`
}

// accountSlurperJobGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountSlurperJobGetResponseError]
type accountSlurperJobGetResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountSlurperJobGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobGetResponseErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    accountSlurperJobGetResponseErrorsSourceJSON `json:"-"`
}

// accountSlurperJobGetResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountSlurperJobGetResponseErrorsSource]
type accountSlurperJobGetResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobGetResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobGetResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type AccountSlurperJobGetResponseSuccess bool

const (
	AccountSlurperJobGetResponseSuccessTrue AccountSlurperJobGetResponseSuccess = true
)

func (r AccountSlurperJobGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountSlurperJobGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountSlurperJobListResponse struct {
	Errors   []AccountSlurperJobListResponseError `json:"errors"`
	Messages []string                             `json:"messages"`
	Result   []JobResponse                        `json:"result"`
	// Indicates if the API call was successful or not.
	Success AccountSlurperJobListResponseSuccess `json:"success"`
	JSON    accountSlurperJobListResponseJSON    `json:"-"`
}

// accountSlurperJobListResponseJSON contains the JSON metadata for the struct
// [AccountSlurperJobListResponse]
type accountSlurperJobListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobListResponseError struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           AccountSlurperJobListResponseErrorsSource `json:"source"`
	JSON             accountSlurperJobListResponseErrorJSON    `json:"-"`
}

// accountSlurperJobListResponseErrorJSON contains the JSON metadata for the struct
// [AccountSlurperJobListResponseError]
type accountSlurperJobListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountSlurperJobListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobListResponseErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    accountSlurperJobListResponseErrorsSourceJSON `json:"-"`
}

// accountSlurperJobListResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountSlurperJobListResponseErrorsSource]
type accountSlurperJobListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type AccountSlurperJobListResponseSuccess bool

const (
	AccountSlurperJobListResponseSuccessTrue AccountSlurperJobListResponseSuccess = true
)

func (r AccountSlurperJobListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountSlurperJobListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountSlurperJobAbortResponse struct {
	Errors   []AccountSlurperJobAbortResponseError `json:"errors"`
	Messages []string                              `json:"messages"`
	Result   string                                `json:"result"`
	// Indicates if the API call was successful or not.
	Success AccountSlurperJobAbortResponseSuccess `json:"success"`
	JSON    accountSlurperJobAbortResponseJSON    `json:"-"`
}

// accountSlurperJobAbortResponseJSON contains the JSON metadata for the struct
// [AccountSlurperJobAbortResponse]
type accountSlurperJobAbortResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobAbortResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobAbortResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobAbortResponseError struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           AccountSlurperJobAbortResponseErrorsSource `json:"source"`
	JSON             accountSlurperJobAbortResponseErrorJSON    `json:"-"`
}

// accountSlurperJobAbortResponseErrorJSON contains the JSON metadata for the
// struct [AccountSlurperJobAbortResponseError]
type accountSlurperJobAbortResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountSlurperJobAbortResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobAbortResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobAbortResponseErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    accountSlurperJobAbortResponseErrorsSourceJSON `json:"-"`
}

// accountSlurperJobAbortResponseErrorsSourceJSON contains the JSON metadata for
// the struct [AccountSlurperJobAbortResponseErrorsSource]
type accountSlurperJobAbortResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobAbortResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobAbortResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type AccountSlurperJobAbortResponseSuccess bool

const (
	AccountSlurperJobAbortResponseSuccessTrue AccountSlurperJobAbortResponseSuccess = true
)

func (r AccountSlurperJobAbortResponseSuccess) IsKnown() bool {
	switch r {
	case AccountSlurperJobAbortResponseSuccessTrue:
		return true
	}
	return false
}

type AccountSlurperJobAbortAllResponse struct {
	Errors   []AccountSlurperJobAbortAllResponseError `json:"errors"`
	Messages []string                                 `json:"messages"`
	Result   string                                   `json:"result"`
	// Indicates if the API call was successful or not.
	Success AccountSlurperJobAbortAllResponseSuccess `json:"success"`
	JSON    accountSlurperJobAbortAllResponseJSON    `json:"-"`
}

// accountSlurperJobAbortAllResponseJSON contains the JSON metadata for the struct
// [AccountSlurperJobAbortAllResponse]
type accountSlurperJobAbortAllResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobAbortAllResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobAbortAllResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobAbortAllResponseError struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           AccountSlurperJobAbortAllResponseErrorsSource `json:"source"`
	JSON             accountSlurperJobAbortAllResponseErrorJSON    `json:"-"`
}

// accountSlurperJobAbortAllResponseErrorJSON contains the JSON metadata for the
// struct [AccountSlurperJobAbortAllResponseError]
type accountSlurperJobAbortAllResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountSlurperJobAbortAllResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobAbortAllResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobAbortAllResponseErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    accountSlurperJobAbortAllResponseErrorsSourceJSON `json:"-"`
}

// accountSlurperJobAbortAllResponseErrorsSourceJSON contains the JSON metadata for
// the struct [AccountSlurperJobAbortAllResponseErrorsSource]
type accountSlurperJobAbortAllResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobAbortAllResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobAbortAllResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type AccountSlurperJobAbortAllResponseSuccess bool

const (
	AccountSlurperJobAbortAllResponseSuccessTrue AccountSlurperJobAbortAllResponseSuccess = true
)

func (r AccountSlurperJobAbortAllResponseSuccess) IsKnown() bool {
	switch r {
	case AccountSlurperJobAbortAllResponseSuccessTrue:
		return true
	}
	return false
}

type AccountSlurperJobGetLogsResponse struct {
	Errors   []AccountSlurperJobGetLogsResponseError  `json:"errors"`
	Messages []string                                 `json:"messages"`
	Result   []AccountSlurperJobGetLogsResponseResult `json:"result"`
	// Indicates if the API call was successful or not.
	Success AccountSlurperJobGetLogsResponseSuccess `json:"success"`
	JSON    accountSlurperJobGetLogsResponseJSON    `json:"-"`
}

// accountSlurperJobGetLogsResponseJSON contains the JSON metadata for the struct
// [AccountSlurperJobGetLogsResponse]
type accountSlurperJobGetLogsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobGetLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobGetLogsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobGetLogsResponseError struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           AccountSlurperJobGetLogsResponseErrorsSource `json:"source"`
	JSON             accountSlurperJobGetLogsResponseErrorJSON    `json:"-"`
}

// accountSlurperJobGetLogsResponseErrorJSON contains the JSON metadata for the
// struct [AccountSlurperJobGetLogsResponseError]
type accountSlurperJobGetLogsResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountSlurperJobGetLogsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobGetLogsResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobGetLogsResponseErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    accountSlurperJobGetLogsResponseErrorsSourceJSON `json:"-"`
}

// accountSlurperJobGetLogsResponseErrorsSourceJSON contains the JSON metadata for
// the struct [AccountSlurperJobGetLogsResponseErrorsSource]
type accountSlurperJobGetLogsResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobGetLogsResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobGetLogsResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobGetLogsResponseResult struct {
	CreatedAt string                                        `json:"createdAt"`
	Job       string                                        `json:"job"`
	LogType   AccountSlurperJobGetLogsResponseResultLogType `json:"logType"`
	Message   string                                        `json:"message,nullable"`
	ObjectKey string                                        `json:"objectKey,nullable"`
	JSON      accountSlurperJobGetLogsResponseResultJSON    `json:"-"`
}

// accountSlurperJobGetLogsResponseResultJSON contains the JSON metadata for the
// struct [AccountSlurperJobGetLogsResponseResult]
type accountSlurperJobGetLogsResponseResultJSON struct {
	CreatedAt   apijson.Field
	Job         apijson.Field
	LogType     apijson.Field
	Message     apijson.Field
	ObjectKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobGetLogsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobGetLogsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobGetLogsResponseResultLogType string

const (
	AccountSlurperJobGetLogsResponseResultLogTypeMigrationStart                      AccountSlurperJobGetLogsResponseResultLogType = "migrationStart"
	AccountSlurperJobGetLogsResponseResultLogTypeMigrationComplete                   AccountSlurperJobGetLogsResponseResultLogType = "migrationComplete"
	AccountSlurperJobGetLogsResponseResultLogTypeMigrationAbort                      AccountSlurperJobGetLogsResponseResultLogType = "migrationAbort"
	AccountSlurperJobGetLogsResponseResultLogTypeMigrationError                      AccountSlurperJobGetLogsResponseResultLogType = "migrationError"
	AccountSlurperJobGetLogsResponseResultLogTypeMigrationPause                      AccountSlurperJobGetLogsResponseResultLogType = "migrationPause"
	AccountSlurperJobGetLogsResponseResultLogTypeMigrationResume                     AccountSlurperJobGetLogsResponseResultLogType = "migrationResume"
	AccountSlurperJobGetLogsResponseResultLogTypeMigrationErrorFailedContinuation    AccountSlurperJobGetLogsResponseResultLogType = "migrationErrorFailedContinuation"
	AccountSlurperJobGetLogsResponseResultLogTypeImportErrorRetryExhaustion          AccountSlurperJobGetLogsResponseResultLogType = "importErrorRetryExhaustion"
	AccountSlurperJobGetLogsResponseResultLogTypeImportSkippedStorageClass           AccountSlurperJobGetLogsResponseResultLogType = "importSkippedStorageClass"
	AccountSlurperJobGetLogsResponseResultLogTypeImportSkippedOversized              AccountSlurperJobGetLogsResponseResultLogType = "importSkippedOversized"
	AccountSlurperJobGetLogsResponseResultLogTypeImportSkippedEmptyObject            AccountSlurperJobGetLogsResponseResultLogType = "importSkippedEmptyObject"
	AccountSlurperJobGetLogsResponseResultLogTypeImportSkippedUnsupportedContentType AccountSlurperJobGetLogsResponseResultLogType = "importSkippedUnsupportedContentType"
	AccountSlurperJobGetLogsResponseResultLogTypeImportSkippedExcludedContentType    AccountSlurperJobGetLogsResponseResultLogType = "importSkippedExcludedContentType"
	AccountSlurperJobGetLogsResponseResultLogTypeImportSkippedInvalidMedia           AccountSlurperJobGetLogsResponseResultLogType = "importSkippedInvalidMedia"
	AccountSlurperJobGetLogsResponseResultLogTypeImportSkippedRequiresRetrieval      AccountSlurperJobGetLogsResponseResultLogType = "importSkippedRequiresRetrieval"
)

func (r AccountSlurperJobGetLogsResponseResultLogType) IsKnown() bool {
	switch r {
	case AccountSlurperJobGetLogsResponseResultLogTypeMigrationStart, AccountSlurperJobGetLogsResponseResultLogTypeMigrationComplete, AccountSlurperJobGetLogsResponseResultLogTypeMigrationAbort, AccountSlurperJobGetLogsResponseResultLogTypeMigrationError, AccountSlurperJobGetLogsResponseResultLogTypeMigrationPause, AccountSlurperJobGetLogsResponseResultLogTypeMigrationResume, AccountSlurperJobGetLogsResponseResultLogTypeMigrationErrorFailedContinuation, AccountSlurperJobGetLogsResponseResultLogTypeImportErrorRetryExhaustion, AccountSlurperJobGetLogsResponseResultLogTypeImportSkippedStorageClass, AccountSlurperJobGetLogsResponseResultLogTypeImportSkippedOversized, AccountSlurperJobGetLogsResponseResultLogTypeImportSkippedEmptyObject, AccountSlurperJobGetLogsResponseResultLogTypeImportSkippedUnsupportedContentType, AccountSlurperJobGetLogsResponseResultLogTypeImportSkippedExcludedContentType, AccountSlurperJobGetLogsResponseResultLogTypeImportSkippedInvalidMedia, AccountSlurperJobGetLogsResponseResultLogTypeImportSkippedRequiresRetrieval:
		return true
	}
	return false
}

// Indicates if the API call was successful or not.
type AccountSlurperJobGetLogsResponseSuccess bool

const (
	AccountSlurperJobGetLogsResponseSuccessTrue AccountSlurperJobGetLogsResponseSuccess = true
)

func (r AccountSlurperJobGetLogsResponseSuccess) IsKnown() bool {
	switch r {
	case AccountSlurperJobGetLogsResponseSuccessTrue:
		return true
	}
	return false
}

type AccountSlurperJobGetProgressResponse struct {
	Errors   []AccountSlurperJobGetProgressResponseError `json:"errors"`
	Messages []string                                    `json:"messages"`
	Result   AccountSlurperJobGetProgressResponseResult  `json:"result"`
	// Indicates if the API call was successful or not.
	Success AccountSlurperJobGetProgressResponseSuccess `json:"success"`
	JSON    accountSlurperJobGetProgressResponseJSON    `json:"-"`
}

// accountSlurperJobGetProgressResponseJSON contains the JSON metadata for the
// struct [AccountSlurperJobGetProgressResponse]
type accountSlurperJobGetProgressResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobGetProgressResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobGetProgressResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobGetProgressResponseError struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           AccountSlurperJobGetProgressResponseErrorsSource `json:"source"`
	JSON             accountSlurperJobGetProgressResponseErrorJSON    `json:"-"`
}

// accountSlurperJobGetProgressResponseErrorJSON contains the JSON metadata for the
// struct [AccountSlurperJobGetProgressResponseError]
type accountSlurperJobGetProgressResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountSlurperJobGetProgressResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobGetProgressResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobGetProgressResponseErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    accountSlurperJobGetProgressResponseErrorsSourceJSON `json:"-"`
}

// accountSlurperJobGetProgressResponseErrorsSourceJSON contains the JSON metadata
// for the struct [AccountSlurperJobGetProgressResponseErrorsSource]
type accountSlurperJobGetProgressResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobGetProgressResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobGetProgressResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobGetProgressResponseResult struct {
	ID                 string                                         `json:"id"`
	CreatedAt          string                                         `json:"createdAt"`
	FailedObjects      int64                                          `json:"failedObjects"`
	Objects            int64                                          `json:"objects"`
	SkippedObjects     int64                                          `json:"skippedObjects"`
	Status             JobStatus                                      `json:"status"`
	TransferredObjects int64                                          `json:"transferredObjects"`
	JSON               accountSlurperJobGetProgressResponseResultJSON `json:"-"`
}

// accountSlurperJobGetProgressResponseResultJSON contains the JSON metadata for
// the struct [AccountSlurperJobGetProgressResponseResult]
type accountSlurperJobGetProgressResponseResultJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	FailedObjects      apijson.Field
	Objects            apijson.Field
	SkippedObjects     apijson.Field
	Status             apijson.Field
	TransferredObjects apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountSlurperJobGetProgressResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobGetProgressResponseResultJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type AccountSlurperJobGetProgressResponseSuccess bool

const (
	AccountSlurperJobGetProgressResponseSuccessTrue AccountSlurperJobGetProgressResponseSuccess = true
)

func (r AccountSlurperJobGetProgressResponseSuccess) IsKnown() bool {
	switch r {
	case AccountSlurperJobGetProgressResponseSuccessTrue:
		return true
	}
	return false
}

type AccountSlurperJobPauseResponse struct {
	Errors   []AccountSlurperJobPauseResponseError `json:"errors"`
	Messages []string                              `json:"messages"`
	Result   string                                `json:"result"`
	// Indicates if the API call was successful or not.
	Success AccountSlurperJobPauseResponseSuccess `json:"success"`
	JSON    accountSlurperJobPauseResponseJSON    `json:"-"`
}

// accountSlurperJobPauseResponseJSON contains the JSON metadata for the struct
// [AccountSlurperJobPauseResponse]
type accountSlurperJobPauseResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobPauseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobPauseResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobPauseResponseError struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           AccountSlurperJobPauseResponseErrorsSource `json:"source"`
	JSON             accountSlurperJobPauseResponseErrorJSON    `json:"-"`
}

// accountSlurperJobPauseResponseErrorJSON contains the JSON metadata for the
// struct [AccountSlurperJobPauseResponseError]
type accountSlurperJobPauseResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountSlurperJobPauseResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobPauseResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobPauseResponseErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    accountSlurperJobPauseResponseErrorsSourceJSON `json:"-"`
}

// accountSlurperJobPauseResponseErrorsSourceJSON contains the JSON metadata for
// the struct [AccountSlurperJobPauseResponseErrorsSource]
type accountSlurperJobPauseResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobPauseResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobPauseResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type AccountSlurperJobPauseResponseSuccess bool

const (
	AccountSlurperJobPauseResponseSuccessTrue AccountSlurperJobPauseResponseSuccess = true
)

func (r AccountSlurperJobPauseResponseSuccess) IsKnown() bool {
	switch r {
	case AccountSlurperJobPauseResponseSuccessTrue:
		return true
	}
	return false
}

type AccountSlurperJobResumeResponse struct {
	Errors   []AccountSlurperJobResumeResponseError `json:"errors"`
	Messages []string                               `json:"messages"`
	Result   string                                 `json:"result"`
	// Indicates if the API call was successful or not.
	Success AccountSlurperJobResumeResponseSuccess `json:"success"`
	JSON    accountSlurperJobResumeResponseJSON    `json:"-"`
}

// accountSlurperJobResumeResponseJSON contains the JSON metadata for the struct
// [AccountSlurperJobResumeResponse]
type accountSlurperJobResumeResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobResumeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobResumeResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobResumeResponseError struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           AccountSlurperJobResumeResponseErrorsSource `json:"source"`
	JSON             accountSlurperJobResumeResponseErrorJSON    `json:"-"`
}

// accountSlurperJobResumeResponseErrorJSON contains the JSON metadata for the
// struct [AccountSlurperJobResumeResponseError]
type accountSlurperJobResumeResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountSlurperJobResumeResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobResumeResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobResumeResponseErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    accountSlurperJobResumeResponseErrorsSourceJSON `json:"-"`
}

// accountSlurperJobResumeResponseErrorsSourceJSON contains the JSON metadata for
// the struct [AccountSlurperJobResumeResponseErrorsSource]
type accountSlurperJobResumeResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobResumeResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobResumeResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type AccountSlurperJobResumeResponseSuccess bool

const (
	AccountSlurperJobResumeResponseSuccessTrue AccountSlurperJobResumeResponseSuccess = true
)

func (r AccountSlurperJobResumeResponseSuccess) IsKnown() bool {
	switch r {
	case AccountSlurperJobResumeResponseSuccessTrue:
		return true
	}
	return false
}

type AccountSlurperJobNewParams struct {
	Overwrite param.Field[bool]                      `json:"overwrite"`
	Source    param.Field[SourceJobSchemaUnionParam] `json:"source"`
	Target    param.Field[R2TargetSchemaParam]       `json:"target"`
}

func (r AccountSlurperJobNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountSlurperJobListParams struct {
	Limit  param.Field[int64] `query:"limit"`
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [AccountSlurperJobListParams]'s query parameters as
// `url.Values`.
func (r AccountSlurperJobListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountSlurperJobGetLogsParams struct {
	Limit  param.Field[int64] `query:"limit"`
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [AccountSlurperJobGetLogsParams]'s query parameters as
// `url.Values`.
func (r AccountSlurperJobGetLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
