// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/cf-rex-go/internal/apiform"
	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountWorkerScriptVersionService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerScriptVersionService] method instead.
type AccountWorkerScriptVersionService struct {
	Options []option.RequestOption
}

// NewAccountWorkerScriptVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerScriptVersionService(opts ...option.RequestOption) (r *AccountWorkerScriptVersionService) {
	r = &AccountWorkerScriptVersionService{}
	r.Options = opts
	return
}

// List of Worker Versions. The first version in the list is the latest version.
func (r *AccountWorkerScriptVersionService) List(ctx context.Context, accountID string, scriptName string, query AccountWorkerScriptVersionListParams, opts ...option.RequestOption) (res *AccountWorkerScriptVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/versions", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Version Detail
func (r *AccountWorkerScriptVersionService) GetDetail(ctx context.Context, accountID string, scriptName string, versionID string, opts ...option.RequestOption) (res *AccountWorkerScriptVersionGetDetailResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/versions/%s", accountID, scriptName, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a Worker Version without deploying to Cloudflare's network. You can find
// more about the multipart metadata on our docs:
// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.
func (r *AccountWorkerScriptVersionService) Upload(ctx context.Context, accountID string, scriptName string, body AccountWorkerScriptVersionUploadParams, opts ...option.RequestOption) (res *AccountWorkerScriptVersionUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/versions", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type VersionItemFull struct {
	Resources interface{}         `json:"resources,required"`
	JSON      versionItemFullJSON `json:"-"`
	VersionItemShort
}

// versionItemFullJSON contains the JSON metadata for the struct [VersionItemFull]
type versionItemFullJSON struct {
	Resources   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionItemFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionItemFullJSON) RawJSON() string {
	return r.raw
}

type VersionItemShort struct {
	ID       string               `json:"id"`
	Metadata interface{}          `json:"metadata"`
	Number   float64              `json:"number"`
	JSON     versionItemShortJSON `json:"-"`
}

// versionItemShortJSON contains the JSON metadata for the struct
// [VersionItemShort]
type versionItemShortJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Number      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionItemShort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionItemShortJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptVersionListResponse struct {
	Result AccountWorkerScriptVersionListResponseResult `json:"result"`
	JSON   accountWorkerScriptVersionListResponseJSON   `json:"-"`
	CommonResponseWorkers
}

// accountWorkerScriptVersionListResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptVersionListResponse]
type accountWorkerScriptVersionListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptVersionListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptVersionListResponseResult struct {
	Items []VersionItemShort                               `json:"items"`
	JSON  accountWorkerScriptVersionListResponseResultJSON `json:"-"`
}

// accountWorkerScriptVersionListResponseResultJSON contains the JSON metadata for
// the struct [AccountWorkerScriptVersionListResponseResult]
type accountWorkerScriptVersionListResponseResultJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptVersionListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptVersionGetDetailResponse struct {
	Result VersionItemFull                                 `json:"result"`
	JSON   accountWorkerScriptVersionGetDetailResponseJSON `json:"-"`
	CommonResponseWorkers
}

// accountWorkerScriptVersionGetDetailResponseJSON contains the JSON metadata for
// the struct [AccountWorkerScriptVersionGetDetailResponse]
type accountWorkerScriptVersionGetDetailResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptVersionGetDetailResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptVersionGetDetailResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptVersionUploadResponse struct {
	Result AccountWorkerScriptVersionUploadResponseResult `json:"result"`
	JSON   accountWorkerScriptVersionUploadResponseJSON   `json:"-"`
	CommonResponseWorkers
}

// accountWorkerScriptVersionUploadResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptVersionUploadResponse]
type accountWorkerScriptVersionUploadResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptVersionUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptVersionUploadResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptVersionUploadResponseResult struct {
	StartupTimeMs int64                                              `json:"startup_time_ms"`
	JSON          accountWorkerScriptVersionUploadResponseResultJSON `json:"-"`
	VersionItemFull
}

// accountWorkerScriptVersionUploadResponseResultJSON contains the JSON metadata
// for the struct [AccountWorkerScriptVersionUploadResponseResult]
type accountWorkerScriptVersionUploadResponseResultJSON struct {
	StartupTimeMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountWorkerScriptVersionUploadResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptVersionUploadResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptVersionListParams struct {
	// Only return versions that can be used in a deployment. Ignores pagination.
	Deployable param.Field[bool] `query:"deployable"`
	// Current page.
	Page param.Field[int64] `query:"page"`
	// Items per-page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [AccountWorkerScriptVersionListParams]'s query parameters as
// `url.Values`.
func (r AccountWorkerScriptVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountWorkerScriptVersionUploadParams struct {
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[AccountWorkerScriptVersionUploadParamsMetadata] `json:"metadata,required"`
}

func (r AccountWorkerScriptVersionUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// JSON encoded metadata about the uploaded parts and Worker configuration.
type AccountWorkerScriptVersionUploadParamsMetadata struct {
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker, which
	// is required for Version Upload.
	MainModule  param.Field[string]                                                    `json:"main_module,required"`
	Annotations param.Field[AccountWorkerScriptVersionUploadParamsMetadataAnnotations] `json:"annotations"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]BindingItemUnionParam] `json:"bindings"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// List of binding types to keep from previous_upload.
	KeepBindings param.Field[[]string] `json:"keep_bindings"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[UsageModel] `json:"usage_model"`
}

func (r AccountWorkerScriptVersionUploadParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptVersionUploadParamsMetadataAnnotations struct {
	// Human-readable message about the version. Truncated to 100 bytes.
	WorkersMessage param.Field[string] `json:"workers/message"`
	// User-provided identifier for the version.
	WorkersTag param.Field[string] `json:"workers/tag"`
}

func (r AccountWorkerScriptVersionUploadParamsMetadataAnnotations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
