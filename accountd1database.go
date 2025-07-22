// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountD1DatabaseService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountD1DatabaseService] method instead.
type AccountD1DatabaseService struct {
	Options []option.RequestOption
}

// NewAccountD1DatabaseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountD1DatabaseService(opts ...option.RequestOption) (r *AccountD1DatabaseService) {
	r = &AccountD1DatabaseService{}
	r.Options = opts
	return
}

// Returns the created D1 database.
func (r *AccountD1DatabaseService) New(ctx context.Context, accountID string, body AccountD1DatabaseNewParams, opts ...option.RequestOption) (res *AccountD1DatabaseNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the specified D1 database.
func (r *AccountD1DatabaseService) Get(ctx context.Context, accountID string, databaseID string, opts ...option.RequestOption) (res *AccountD1DatabaseGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s", accountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of D1 databases.
func (r *AccountD1DatabaseService) List(ctx context.Context, accountID string, query AccountD1DatabaseListParams, opts ...option.RequestOption) (res *AccountD1DatabaseListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes the specified D1 database.
func (r *AccountD1DatabaseService) Delete(ctx context.Context, accountID string, databaseID string, opts ...option.RequestOption) (res *AccountD1DatabaseDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s", accountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Returns a URL where the SQL contents of your D1 can be downloaded. Note: this
// process may take some time for larger DBs, during which your D1 will be
// unavailable to serve queries. To avoid blocking your DB unnecessarily, an
// in-progress export must be continually polled or will automatically cancel.
func (r *AccountD1DatabaseService) Export(ctx context.Context, accountID string, databaseID string, body AccountD1DatabaseExportParams, opts ...option.RequestOption) (res *AccountD1DatabaseExportResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s/export", accountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generates a temporary URL for uploading an SQL file to, then instructing the D1
// to import it and polling it for status updates. Imports block the D1 for their
// duration.
func (r *AccountD1DatabaseService) Import(ctx context.Context, accountID string, databaseID string, body AccountD1DatabaseImportParams, opts ...option.RequestOption) (res *AccountD1DatabaseImportResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s/import", accountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the query result as an object.
func (r *AccountD1DatabaseService) Query(ctx context.Context, accountID string, databaseID string, body AccountD1DatabaseQueryParams, opts ...option.RequestOption) (res *AccountD1DatabaseQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s/query", accountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the query result rows as arrays rather than objects. This is a
// performance-optimized version of the /query endpoint.
func (r *AccountD1DatabaseService) RawQuery(ctx context.Context, accountID string, databaseID string, body AccountD1DatabaseRawQueryParams, opts ...option.RequestOption) (res *AccountD1DatabaseRawQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s/raw", accountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type APIResponseD1 struct {
	Errors   []MessagesD1Item `json:"errors,required"`
	Messages []MessagesD1Item `json:"messages,required"`
	Result   interface{}      `json:"result,required"`
	// Whether the API call was successful
	Success APIResponseD1Success `json:"success,required"`
	JSON    apiResponseD1JSON    `json:"-"`
}

// apiResponseD1JSON contains the JSON metadata for the struct [APIResponseD1]
type apiResponseD1JSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseD1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseD1JSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseD1Success bool

const (
	APIResponseD1SuccessTrue APIResponseD1Success = true
)

func (r APIResponseD1Success) IsKnown() bool {
	switch r {
	case APIResponseD1SuccessTrue:
		return true
	}
	return false
}

type DatabaseDetailsResponse struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The D1 database's size, in bytes.
	FileSize float64 `json:"file_size"`
	// D1 database name.
	Name      string  `json:"name"`
	NumTables float64 `json:"num_tables"`
	// D1 database identifier (UUID).
	Uuid    string                      `json:"uuid"`
	Version string                      `json:"version"`
	JSON    databaseDetailsResponseJSON `json:"-"`
}

// databaseDetailsResponseJSON contains the JSON metadata for the struct
// [DatabaseDetailsResponse]
type databaseDetailsResponseJSON struct {
	CreatedAt   apijson.Field
	FileSize    apijson.Field
	Name        apijson.Field
	NumTables   apijson.Field
	Uuid        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseDetailsResponseJSON) RawJSON() string {
	return r.raw
}

type MessagesD1Item struct {
	Code    int64              `json:"code,required"`
	Message string             `json:"message,required"`
	JSON    messagesD1ItemJSON `json:"-"`
}

// messagesD1ItemJSON contains the JSON metadata for the struct [MessagesD1Item]
type messagesD1ItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesD1Item) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesD1ItemJSON) RawJSON() string {
	return r.raw
}

type QueryMeta struct {
	ChangedDB   bool          `json:"changed_db"`
	Changes     float64       `json:"changes"`
	Duration    float64       `json:"duration"`
	LastRowID   float64       `json:"last_row_id"`
	RowsRead    float64       `json:"rows_read"`
	RowsWritten float64       `json:"rows_written"`
	SizeAfter   float64       `json:"size_after"`
	JSON        queryMetaJSON `json:"-"`
}

// queryMetaJSON contains the JSON metadata for the struct [QueryMeta]
type queryMetaJSON struct {
	ChangedDB   apijson.Field
	Changes     apijson.Field
	Duration    apijson.Field
	LastRowID   apijson.Field
	RowsRead    apijson.Field
	RowsWritten apijson.Field
	SizeAfter   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueryMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queryMetaJSON) RawJSON() string {
	return r.raw
}

type AccountD1DatabaseNewResponse struct {
	Result DatabaseDetailsResponse          `json:"result"`
	JSON   accountD1DatabaseNewResponseJSON `json:"-"`
	APIResponseD1
}

// accountD1DatabaseNewResponseJSON contains the JSON metadata for the struct
// [AccountD1DatabaseNewResponse]
type accountD1DatabaseNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountD1DatabaseNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountD1DatabaseGetResponse struct {
	Result DatabaseDetailsResponse          `json:"result"`
	JSON   accountD1DatabaseGetResponseJSON `json:"-"`
	APIResponseD1
}

// accountD1DatabaseGetResponseJSON contains the JSON metadata for the struct
// [AccountD1DatabaseGetResponse]
type accountD1DatabaseGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountD1DatabaseGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountD1DatabaseListResponse struct {
	Result     []AccountD1DatabaseListResponseResult   `json:"result"`
	ResultInfo AccountD1DatabaseListResponseResultInfo `json:"result_info"`
	JSON       accountD1DatabaseListResponseJSON       `json:"-"`
	APIResponseD1
}

// accountD1DatabaseListResponseJSON contains the JSON metadata for the struct
// [AccountD1DatabaseListResponse]
type accountD1DatabaseListResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountD1DatabaseListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountD1DatabaseListResponseResult struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// D1 database name.
	Name string `json:"name"`
	// D1 database identifier (UUID).
	Uuid    string                                  `json:"uuid"`
	Version string                                  `json:"version"`
	JSON    accountD1DatabaseListResponseResultJSON `json:"-"`
}

// accountD1DatabaseListResponseResultJSON contains the JSON metadata for the
// struct [AccountD1DatabaseListResponseResult]
type accountD1DatabaseListResponseResultJSON struct {
	CreatedAt   apijson.Field
	Name        apijson.Field
	Uuid        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountD1DatabaseListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountD1DatabaseListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       accountD1DatabaseListResponseResultInfoJSON `json:"-"`
}

// accountD1DatabaseListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountD1DatabaseListResponseResultInfo]
type accountD1DatabaseListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountD1DatabaseListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountD1DatabaseDeleteResponse struct {
	Result interface{}                         `json:"result,nullable"`
	JSON   accountD1DatabaseDeleteResponseJSON `json:"-"`
	APIResponseD1
}

// accountD1DatabaseDeleteResponseJSON contains the JSON metadata for the struct
// [AccountD1DatabaseDeleteResponse]
type accountD1DatabaseDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountD1DatabaseDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountD1DatabaseExportResponse struct {
	Result AccountD1DatabaseExportResponseResult `json:"result"`
	JSON   accountD1DatabaseExportResponseJSON   `json:"-"`
	APIResponseD1
}

// accountD1DatabaseExportResponseJSON contains the JSON metadata for the struct
// [AccountD1DatabaseExportResponse]
type accountD1DatabaseExportResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseExportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountD1DatabaseExportResponseJSON) RawJSON() string {
	return r.raw
}

type AccountD1DatabaseExportResponseResult struct {
	// The current time-travel bookmark for your D1, used to poll for updates. Will not
	// change for the duration of the export task.
	AtBookmark string `json:"at_bookmark"`
	// Only present when status = 'error'. Contains the error message.
	Error string `json:"error"`
	// Logs since the last time you polled
	Messages []string `json:"messages"`
	// Only present when status = 'complete'
	Result  AccountD1DatabaseExportResponseResultResult `json:"result"`
	Status  AccountD1DatabaseExportResponseResultStatus `json:"status"`
	Success bool                                        `json:"success"`
	Type    AccountD1DatabaseExportResponseResultType   `json:"type"`
	JSON    accountD1DatabaseExportResponseResultJSON   `json:"-"`
}

// accountD1DatabaseExportResponseResultJSON contains the JSON metadata for the
// struct [AccountD1DatabaseExportResponseResult]
type accountD1DatabaseExportResponseResultJSON struct {
	AtBookmark  apijson.Field
	Error       apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Status      apijson.Field
	Success     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseExportResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountD1DatabaseExportResponseResultJSON) RawJSON() string {
	return r.raw
}

// Only present when status = 'complete'
type AccountD1DatabaseExportResponseResultResult struct {
	// The generated SQL filename.
	Filename string `json:"filename"`
	// The URL to download the exported SQL. Available for one hour.
	SignedURL string                                          `json:"signed_url"`
	JSON      accountD1DatabaseExportResponseResultResultJSON `json:"-"`
}

// accountD1DatabaseExportResponseResultResultJSON contains the JSON metadata for
// the struct [AccountD1DatabaseExportResponseResultResult]
type accountD1DatabaseExportResponseResultResultJSON struct {
	Filename    apijson.Field
	SignedURL   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseExportResponseResultResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountD1DatabaseExportResponseResultResultJSON) RawJSON() string {
	return r.raw
}

type AccountD1DatabaseExportResponseResultStatus string

const (
	AccountD1DatabaseExportResponseResultStatusComplete AccountD1DatabaseExportResponseResultStatus = "complete"
	AccountD1DatabaseExportResponseResultStatusError    AccountD1DatabaseExportResponseResultStatus = "error"
)

func (r AccountD1DatabaseExportResponseResultStatus) IsKnown() bool {
	switch r {
	case AccountD1DatabaseExportResponseResultStatusComplete, AccountD1DatabaseExportResponseResultStatusError:
		return true
	}
	return false
}

type AccountD1DatabaseExportResponseResultType string

const (
	AccountD1DatabaseExportResponseResultTypeExport AccountD1DatabaseExportResponseResultType = "export"
)

func (r AccountD1DatabaseExportResponseResultType) IsKnown() bool {
	switch r {
	case AccountD1DatabaseExportResponseResultTypeExport:
		return true
	}
	return false
}

type AccountD1DatabaseImportResponse struct {
	Result AccountD1DatabaseImportResponseResult `json:"result"`
	JSON   accountD1DatabaseImportResponseJSON   `json:"-"`
	APIResponseD1
}

// accountD1DatabaseImportResponseJSON contains the JSON metadata for the struct
// [AccountD1DatabaseImportResponse]
type accountD1DatabaseImportResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseImportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountD1DatabaseImportResponseJSON) RawJSON() string {
	return r.raw
}

type AccountD1DatabaseImportResponseResult struct {
	// The current time-travel bookmark for your D1, used to poll for updates. Will not
	// change for the duration of the import. Only returned if an import process is
	// currently running or recently finished.
	AtBookmark string `json:"at_bookmark"`
	// Only present when status = 'error'. Contains the error message that prevented
	// the import from succeeding.
	Error string `json:"error"`
	// Derived from the database ID and etag, to use in avoiding repeated uploads. Only
	// returned when for the 'init' action.
	Filename string `json:"filename"`
	// Logs since the last time you polled
	Messages []string `json:"messages"`
	// Only present when status = 'complete'
	Result  AccountD1DatabaseImportResponseResultResult `json:"result"`
	Status  AccountD1DatabaseImportResponseResultStatus `json:"status"`
	Success bool                                        `json:"success"`
	Type    AccountD1DatabaseImportResponseResultType   `json:"type"`
	// The R2 presigned URL to use for uploading. Only returned when for the 'init'
	// action.
	UploadURL string                                    `json:"upload_url"`
	JSON      accountD1DatabaseImportResponseResultJSON `json:"-"`
}

// accountD1DatabaseImportResponseResultJSON contains the JSON metadata for the
// struct [AccountD1DatabaseImportResponseResult]
type accountD1DatabaseImportResponseResultJSON struct {
	AtBookmark  apijson.Field
	Error       apijson.Field
	Filename    apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Status      apijson.Field
	Success     apijson.Field
	Type        apijson.Field
	UploadURL   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseImportResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountD1DatabaseImportResponseResultJSON) RawJSON() string {
	return r.raw
}

// Only present when status = 'complete'
type AccountD1DatabaseImportResponseResultResult struct {
	// The time-travel bookmark if you need restore your D1 to directly after the
	// import succeeded.
	FinalBookmark string    `json:"final_bookmark"`
	Meta          QueryMeta `json:"meta"`
	// The total number of queries that were executed during the import.
	NumQueries float64                                         `json:"num_queries"`
	JSON       accountD1DatabaseImportResponseResultResultJSON `json:"-"`
}

// accountD1DatabaseImportResponseResultResultJSON contains the JSON metadata for
// the struct [AccountD1DatabaseImportResponseResultResult]
type accountD1DatabaseImportResponseResultResultJSON struct {
	FinalBookmark apijson.Field
	Meta          apijson.Field
	NumQueries    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountD1DatabaseImportResponseResultResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountD1DatabaseImportResponseResultResultJSON) RawJSON() string {
	return r.raw
}

type AccountD1DatabaseImportResponseResultStatus string

const (
	AccountD1DatabaseImportResponseResultStatusComplete AccountD1DatabaseImportResponseResultStatus = "complete"
	AccountD1DatabaseImportResponseResultStatusError    AccountD1DatabaseImportResponseResultStatus = "error"
)

func (r AccountD1DatabaseImportResponseResultStatus) IsKnown() bool {
	switch r {
	case AccountD1DatabaseImportResponseResultStatusComplete, AccountD1DatabaseImportResponseResultStatusError:
		return true
	}
	return false
}

type AccountD1DatabaseImportResponseResultType string

const (
	AccountD1DatabaseImportResponseResultTypeImport AccountD1DatabaseImportResponseResultType = "import"
)

func (r AccountD1DatabaseImportResponseResultType) IsKnown() bool {
	switch r {
	case AccountD1DatabaseImportResponseResultTypeImport:
		return true
	}
	return false
}

type AccountD1DatabaseQueryResponse struct {
	Result []AccountD1DatabaseQueryResponseResult `json:"result"`
	JSON   accountD1DatabaseQueryResponseJSON     `json:"-"`
	APIResponseD1
}

// accountD1DatabaseQueryResponseJSON contains the JSON metadata for the struct
// [AccountD1DatabaseQueryResponse]
type accountD1DatabaseQueryResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountD1DatabaseQueryResponseJSON) RawJSON() string {
	return r.raw
}

type AccountD1DatabaseQueryResponseResult struct {
	Meta    QueryMeta                                `json:"meta"`
	Results []interface{}                            `json:"results"`
	Success bool                                     `json:"success"`
	JSON    accountD1DatabaseQueryResponseResultJSON `json:"-"`
}

// accountD1DatabaseQueryResponseResultJSON contains the JSON metadata for the
// struct [AccountD1DatabaseQueryResponseResult]
type accountD1DatabaseQueryResponseResultJSON struct {
	Meta        apijson.Field
	Results     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseQueryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountD1DatabaseQueryResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountD1DatabaseRawQueryResponse struct {
	Result []AccountD1DatabaseRawQueryResponseResult `json:"result"`
	JSON   accountD1DatabaseRawQueryResponseJSON     `json:"-"`
	APIResponseD1
}

// accountD1DatabaseRawQueryResponseJSON contains the JSON metadata for the struct
// [AccountD1DatabaseRawQueryResponse]
type accountD1DatabaseRawQueryResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseRawQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountD1DatabaseRawQueryResponseJSON) RawJSON() string {
	return r.raw
}

type AccountD1DatabaseRawQueryResponseResult struct {
	Meta    QueryMeta                                      `json:"meta"`
	Results AccountD1DatabaseRawQueryResponseResultResults `json:"results"`
	Success bool                                           `json:"success"`
	JSON    accountD1DatabaseRawQueryResponseResultJSON    `json:"-"`
}

// accountD1DatabaseRawQueryResponseResultJSON contains the JSON metadata for the
// struct [AccountD1DatabaseRawQueryResponseResult]
type accountD1DatabaseRawQueryResponseResultJSON struct {
	Meta        apijson.Field
	Results     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseRawQueryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountD1DatabaseRawQueryResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountD1DatabaseRawQueryResponseResultResults struct {
	Columns []string                                           `json:"columns"`
	Rows    [][]interface{}                                    `json:"rows"`
	JSON    accountD1DatabaseRawQueryResponseResultResultsJSON `json:"-"`
}

// accountD1DatabaseRawQueryResponseResultResultsJSON contains the JSON metadata
// for the struct [AccountD1DatabaseRawQueryResponseResultResults]
type accountD1DatabaseRawQueryResponseResultResultsJSON struct {
	Columns     apijson.Field
	Rows        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountD1DatabaseRawQueryResponseResultResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountD1DatabaseRawQueryResponseResultResultsJSON) RawJSON() string {
	return r.raw
}

type AccountD1DatabaseNewParams struct {
	// D1 database name.
	Name param.Field[string] `json:"name,required"`
	// Specify the region to create the D1 primary, if available. If this option is
	// omitted, the D1 will be created as close as possible to the current user.
	PrimaryLocationHint param.Field[AccountD1DatabaseNewParamsPrimaryLocationHint] `json:"primary_location_hint"`
}

func (r AccountD1DatabaseNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify the region to create the D1 primary, if available. If this option is
// omitted, the D1 will be created as close as possible to the current user.
type AccountD1DatabaseNewParamsPrimaryLocationHint string

const (
	AccountD1DatabaseNewParamsPrimaryLocationHintWnam AccountD1DatabaseNewParamsPrimaryLocationHint = "wnam"
	AccountD1DatabaseNewParamsPrimaryLocationHintEnam AccountD1DatabaseNewParamsPrimaryLocationHint = "enam"
	AccountD1DatabaseNewParamsPrimaryLocationHintWeur AccountD1DatabaseNewParamsPrimaryLocationHint = "weur"
	AccountD1DatabaseNewParamsPrimaryLocationHintEeur AccountD1DatabaseNewParamsPrimaryLocationHint = "eeur"
	AccountD1DatabaseNewParamsPrimaryLocationHintApac AccountD1DatabaseNewParamsPrimaryLocationHint = "apac"
	AccountD1DatabaseNewParamsPrimaryLocationHintOc   AccountD1DatabaseNewParamsPrimaryLocationHint = "oc"
)

func (r AccountD1DatabaseNewParamsPrimaryLocationHint) IsKnown() bool {
	switch r {
	case AccountD1DatabaseNewParamsPrimaryLocationHintWnam, AccountD1DatabaseNewParamsPrimaryLocationHintEnam, AccountD1DatabaseNewParamsPrimaryLocationHintWeur, AccountD1DatabaseNewParamsPrimaryLocationHintEeur, AccountD1DatabaseNewParamsPrimaryLocationHintApac, AccountD1DatabaseNewParamsPrimaryLocationHintOc:
		return true
	}
	return false
}

type AccountD1DatabaseListParams struct {
	// a database name to search for.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountD1DatabaseListParams]'s query parameters as
// `url.Values`.
func (r AccountD1DatabaseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountD1DatabaseExportParams struct {
	// Specifies that you will poll this endpoint until the export completes
	OutputFormat param.Field[AccountD1DatabaseExportParamsOutputFormat] `json:"output_format,required"`
	// To poll an in-progress export, provide the current bookmark (returned by your
	// first polling response)
	CurrentBookmark param.Field[string]                                   `json:"current_bookmark"`
	DumpOptions     param.Field[AccountD1DatabaseExportParamsDumpOptions] `json:"dump_options"`
}

func (r AccountD1DatabaseExportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies that you will poll this endpoint until the export completes
type AccountD1DatabaseExportParamsOutputFormat string

const (
	AccountD1DatabaseExportParamsOutputFormatPolling AccountD1DatabaseExportParamsOutputFormat = "polling"
)

func (r AccountD1DatabaseExportParamsOutputFormat) IsKnown() bool {
	switch r {
	case AccountD1DatabaseExportParamsOutputFormatPolling:
		return true
	}
	return false
}

type AccountD1DatabaseExportParamsDumpOptions struct {
	// Export only the table definitions, not their contents
	NoData param.Field[bool] `json:"no_data"`
	// Export only each table's contents, not its definition
	NoSchema param.Field[bool] `json:"no_schema"`
	// Filter the export to just one or more tables. Passing an empty array is the same
	// as not passing anything and means: export all tables.
	Tables param.Field[[]string] `json:"tables"`
}

func (r AccountD1DatabaseExportParamsDumpOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountD1DatabaseImportParams struct {
	Body AccountD1DatabaseImportParamsBodyUnion `json:"body,required"`
}

func (r AccountD1DatabaseImportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountD1DatabaseImportParamsBody struct {
	// Indicates you have a new SQL file to upload.
	Action param.Field[AccountD1DatabaseImportParamsBodyAction] `json:"action,required"`
	// This identifies the currently-running import, checking its status.
	CurrentBookmark param.Field[string] `json:"current_bookmark"`
	// Required when action is 'init' or 'ingest'. An md5 hash of the file you're
	// uploading. Used to check if it already exists, and validate its contents before
	// ingesting.
	Etag param.Field[string] `json:"etag"`
	// The filename you have successfully uploaded.
	Filename param.Field[string] `json:"filename"`
}

func (r AccountD1DatabaseImportParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountD1DatabaseImportParamsBody) implementsAccountD1DatabaseImportParamsBodyUnion() {}

// Satisfied by [AccountD1DatabaseImportParamsBodyObject],
// [AccountD1DatabaseImportParamsBodyObject],
// [AccountD1DatabaseImportParamsBodyObject], [AccountD1DatabaseImportParamsBody].
type AccountD1DatabaseImportParamsBodyUnion interface {
	implementsAccountD1DatabaseImportParamsBodyUnion()
}

type AccountD1DatabaseImportParamsBodyObject struct {
	// Indicates you have a new SQL file to upload.
	Action param.Field[AccountD1DatabaseImportParamsBodyObjectAction] `json:"action,required"`
	// Required when action is 'init' or 'ingest'. An md5 hash of the file you're
	// uploading. Used to check if it already exists, and validate its contents before
	// ingesting.
	Etag param.Field[string] `json:"etag,required"`
}

func (r AccountD1DatabaseImportParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountD1DatabaseImportParamsBodyObject) implementsAccountD1DatabaseImportParamsBodyUnion() {}

// Indicates you have a new SQL file to upload.
type AccountD1DatabaseImportParamsBodyObjectAction string

const (
	AccountD1DatabaseImportParamsBodyObjectActionInit AccountD1DatabaseImportParamsBodyObjectAction = "init"
)

func (r AccountD1DatabaseImportParamsBodyObjectAction) IsKnown() bool {
	switch r {
	case AccountD1DatabaseImportParamsBodyObjectActionInit:
		return true
	}
	return false
}

// Indicates you have a new SQL file to upload.
type AccountD1DatabaseImportParamsBodyAction string

const (
	AccountD1DatabaseImportParamsBodyActionInit   AccountD1DatabaseImportParamsBodyAction = "init"
	AccountD1DatabaseImportParamsBodyActionIngest AccountD1DatabaseImportParamsBodyAction = "ingest"
	AccountD1DatabaseImportParamsBodyActionPoll   AccountD1DatabaseImportParamsBodyAction = "poll"
)

func (r AccountD1DatabaseImportParamsBodyAction) IsKnown() bool {
	switch r {
	case AccountD1DatabaseImportParamsBodyActionInit, AccountD1DatabaseImportParamsBodyActionIngest, AccountD1DatabaseImportParamsBodyActionPoll:
		return true
	}
	return false
}

type AccountD1DatabaseQueryParams struct {
	// Your SQL query. Supports multiple statements, joined by semicolons, which will
	// be executed as a batch.
	Sql    param.Field[string]   `json:"sql,required"`
	Params param.Field[[]string] `json:"params"`
}

func (r AccountD1DatabaseQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountD1DatabaseRawQueryParams struct {
	// Your SQL query. Supports multiple statements, joined by semicolons, which will
	// be executed as a batch.
	Sql    param.Field[string]   `json:"sql,required"`
	Params param.Field[[]string] `json:"params"`
}

func (r AccountD1DatabaseRawQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
