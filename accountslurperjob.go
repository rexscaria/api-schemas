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

type APIV4Success struct {
	Errors   []Apiv4SuccessError `json:"errors"`
	Messages []string            `json:"messages"`
	// Indicates if the API call was successful or not.
	Success APIV4SuccessSuccess `json:"success"`
	JSON    apiv4SuccessJSON    `json:"-"`
}

// apiv4SuccessJSON contains the JSON metadata for the struct [APIV4Success]
type apiv4SuccessJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIV4Success) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiv4SuccessJSON) RawJSON() string {
	return r.raw
}

type Apiv4SuccessError struct {
	Code    int64                 `json:"code,required"`
	Message string                `json:"message,required"`
	JSON    apiv4SuccessErrorJSON `json:"-"`
}

// apiv4SuccessErrorJSON contains the JSON metadata for the struct
// [Apiv4SuccessError]
type apiv4SuccessErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Apiv4SuccessError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiv4SuccessErrorJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type APIV4SuccessSuccess bool

const (
	APIV4SuccessSuccessTrue APIV4SuccessSuccess = true
)

func (r APIV4SuccessSuccess) IsKnown() bool {
	switch r {
	case APIV4SuccessSuccessTrue:
		return true
	}
	return false
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
	Result AccountSlurperJobNewResponseResult `json:"result"`
	JSON   accountSlurperJobNewResponseJSON   `json:"-"`
	APIV4Success
}

// accountSlurperJobNewResponseJSON contains the JSON metadata for the struct
// [AccountSlurperJobNewResponse]
type accountSlurperJobNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobNewResponseJSON) RawJSON() string {
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

type AccountSlurperJobGetResponse struct {
	Result JobResponse                      `json:"result"`
	JSON   accountSlurperJobGetResponseJSON `json:"-"`
	APIV4Success
}

// accountSlurperJobGetResponseJSON contains the JSON metadata for the struct
// [AccountSlurperJobGetResponse]
type accountSlurperJobGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobListResponse struct {
	Result []JobResponse                     `json:"result"`
	JSON   accountSlurperJobListResponseJSON `json:"-"`
	APIV4Success
}

// accountSlurperJobListResponseJSON contains the JSON metadata for the struct
// [AccountSlurperJobListResponse]
type accountSlurperJobListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobAbortResponse struct {
	Result string                             `json:"result"`
	JSON   accountSlurperJobAbortResponseJSON `json:"-"`
	APIV4Success
}

// accountSlurperJobAbortResponseJSON contains the JSON metadata for the struct
// [AccountSlurperJobAbortResponse]
type accountSlurperJobAbortResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobAbortResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobAbortResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobAbortAllResponse struct {
	Result string                                `json:"result"`
	JSON   accountSlurperJobAbortAllResponseJSON `json:"-"`
	APIV4Success
}

// accountSlurperJobAbortAllResponseJSON contains the JSON metadata for the struct
// [AccountSlurperJobAbortAllResponse]
type accountSlurperJobAbortAllResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobAbortAllResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobAbortAllResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobGetLogsResponse struct {
	Result []AccountSlurperJobGetLogsResponseResult `json:"result"`
	JSON   accountSlurperJobGetLogsResponseJSON     `json:"-"`
	APIV4Success
}

// accountSlurperJobGetLogsResponseJSON contains the JSON metadata for the struct
// [AccountSlurperJobGetLogsResponse]
type accountSlurperJobGetLogsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobGetLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobGetLogsResponseJSON) RawJSON() string {
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

type AccountSlurperJobGetProgressResponse struct {
	Result AccountSlurperJobGetProgressResponseResult `json:"result"`
	JSON   accountSlurperJobGetProgressResponseJSON   `json:"-"`
	APIV4Success
}

// accountSlurperJobGetProgressResponseJSON contains the JSON metadata for the
// struct [AccountSlurperJobGetProgressResponse]
type accountSlurperJobGetProgressResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobGetProgressResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobGetProgressResponseJSON) RawJSON() string {
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

type AccountSlurperJobPauseResponse struct {
	Result string                             `json:"result"`
	JSON   accountSlurperJobPauseResponseJSON `json:"-"`
	APIV4Success
}

// accountSlurperJobPauseResponseJSON contains the JSON metadata for the struct
// [AccountSlurperJobPauseResponse]
type accountSlurperJobPauseResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobPauseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobPauseResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSlurperJobResumeResponse struct {
	Result string                              `json:"result"`
	JSON   accountSlurperJobResumeResponseJSON `json:"-"`
	APIV4Success
}

// accountSlurperJobResumeResponseJSON contains the JSON metadata for the struct
// [AccountSlurperJobResumeResponse]
type accountSlurperJobResumeResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSlurperJobResumeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSlurperJobResumeResponseJSON) RawJSON() string {
	return r.raw
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
