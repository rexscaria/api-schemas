// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountCallAppService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCallAppService] method instead.
type AccountCallAppService struct {
	Options []option.RequestOption
}

// NewAccountCallAppService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountCallAppService(opts ...option.RequestOption) (r *AccountCallAppService) {
	r = &AccountCallAppService{}
	r.Options = opts
	return
}

// Creates a new Cloudflare calls app. An app is an unique enviroment where each
// Session can access all Tracks within the app.
func (r *AccountCallAppService) New(ctx context.Context, accountID string, body AccountCallAppNewParams, opts ...option.RequestOption) (res *AccountCallAppNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/apps", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches details for a single Calls app.
func (r *AccountCallAppService) Get(ctx context.Context, accountID string, appID string, opts ...option.RequestOption) (res *CallsAppResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/apps/%s", accountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Edit details for a single app.
func (r *AccountCallAppService) Update(ctx context.Context, accountID string, appID string, body AccountCallAppUpdateParams, opts ...option.RequestOption) (res *CallsAppResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/apps/%s", accountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all apps in the Cloudflare account
func (r *AccountCallAppService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountCallAppListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/apps", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an app from Cloudflare Calls
func (r *AccountCallAppService) Delete(ctx context.Context, accountID string, appID string, opts ...option.RequestOption) (res *CallsAppResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/apps/%s", accountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CallsAPIResponseCommon struct {
	Errors   []CallsMessageItem `json:"errors,required"`
	Messages []CallsMessageItem `json:"messages,required"`
	// Whether the API call was successful
	Success CallsAPIResponseCommonSuccess `json:"success,required"`
	JSON    callsAPIResponseCommonJSON    `json:"-"`
}

// callsAPIResponseCommonJSON contains the JSON metadata for the struct
// [CallsAPIResponseCommon]
type callsAPIResponseCommonJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallsAPIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callsAPIResponseCommonJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CallsAPIResponseCommonSuccess bool

const (
	CallsAPIResponseCommonSuccessTrue CallsAPIResponseCommonSuccess = true
)

func (r CallsAPIResponseCommonSuccess) IsKnown() bool {
	switch r {
	case CallsAPIResponseCommonSuccessTrue:
		return true
	}
	return false
}

type CallsAPIResponseSingle struct {
	Errors   []CallsMessageItem `json:"errors,required"`
	Messages []CallsMessageItem `json:"messages,required"`
	// Whether the API call was successful
	Success CallsAPIResponseSingleSuccess `json:"success,required"`
	JSON    callsAPIResponseSingleJSON    `json:"-"`
}

// callsAPIResponseSingleJSON contains the JSON metadata for the struct
// [CallsAPIResponseSingle]
type callsAPIResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallsAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callsAPIResponseSingleJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CallsAPIResponseSingleSuccess bool

const (
	CallsAPIResponseSingleSuccessTrue CallsAPIResponseSingleSuccess = true
)

func (r CallsAPIResponseSingleSuccess) IsKnown() bool {
	switch r {
	case CallsAPIResponseSingleSuccessTrue:
		return true
	}
	return false
}

type CallsApp struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	Uid  string       `json:"uid"`
	JSON callsAppJSON `json:"-"`
}

// callsAppJSON contains the JSON metadata for the struct [CallsApp]
type callsAppJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallsApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callsAppJSON) RawJSON() string {
	return r.raw
}

type CallsAppEditableFieldsParam struct {
	// A short description of Calls app, not shown to end users.
	Name param.Field[string] `json:"name"`
}

func (r CallsAppEditableFieldsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CallsAppResponseSingle struct {
	Result CallsApp                   `json:"result"`
	JSON   callsAppResponseSingleJSON `json:"-"`
	CallsAPIResponseSingle
}

// callsAppResponseSingleJSON contains the JSON metadata for the struct
// [CallsAppResponseSingle]
type callsAppResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallsAppResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callsAppResponseSingleJSON) RawJSON() string {
	return r.raw
}

type CallsMessageItem struct {
	Code    int64                `json:"code,required"`
	Message string               `json:"message,required"`
	JSON    callsMessageItemJSON `json:"-"`
}

// callsMessageItemJSON contains the JSON metadata for the struct
// [CallsMessageItem]
type callsMessageItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallsMessageItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callsMessageItemJSON) RawJSON() string {
	return r.raw
}

type AccountCallAppNewResponse struct {
	Result AccountCallAppNewResponseResult `json:"result"`
	JSON   accountCallAppNewResponseJSON   `json:"-"`
	CallsAPIResponseSingle
}

// accountCallAppNewResponseJSON contains the JSON metadata for the struct
// [AccountCallAppNewResponse]
type accountCallAppNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCallAppNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCallAppNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCallAppNewResponseResult struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// Bearer token
	Secret string `json:"secret"`
	// A Cloudflare-generated unique identifier for a item.
	Uid  string                              `json:"uid"`
	JSON accountCallAppNewResponseResultJSON `json:"-"`
}

// accountCallAppNewResponseResultJSON contains the JSON metadata for the struct
// [AccountCallAppNewResponseResult]
type accountCallAppNewResponseResultJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCallAppNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCallAppNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountCallAppListResponse struct {
	Result []CallsApp                     `json:"result"`
	JSON   accountCallAppListResponseJSON `json:"-"`
	CallsAPIResponseCommon
}

// accountCallAppListResponseJSON contains the JSON metadata for the struct
// [AccountCallAppListResponse]
type accountCallAppListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCallAppListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCallAppListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCallAppNewParams struct {
	CallsAppEditableFields CallsAppEditableFieldsParam `json:"calls_app_editable_fields,required"`
}

func (r AccountCallAppNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CallsAppEditableFields)
}

type AccountCallAppUpdateParams struct {
	CallsAppEditableFields CallsAppEditableFieldsParam `json:"calls_app_editable_fields,required"`
}

func (r AccountCallAppUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CallsAppEditableFields)
}
