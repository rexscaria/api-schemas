// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountShareRecipientService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountShareRecipientService] method instead.
type AccountShareRecipientService struct {
	Options []option.RequestOption
}

// NewAccountShareRecipientService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountShareRecipientService(opts ...option.RequestOption) (r *AccountShareRecipientService) {
	r = &AccountShareRecipientService{}
	r.Options = opts
	return
}

// Create a new share recipient
func (r *AccountShareRecipientService) New(ctx context.Context, accountID string, shareID string, body AccountShareRecipientNewParams, opts ...option.RequestOption) (res *ShareRecipientResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/recipients", accountID, shareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get share recipient by ID.
func (r *AccountShareRecipientService) Get(ctx context.Context, accountID string, shareID string, recipientID string, opts ...option.RequestOption) (res *ShareRecipientResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	if recipientID == "" {
		err = errors.New("missing required recipient_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/recipients/%s", accountID, shareID, recipientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List share recipients by share ID.
func (r *AccountShareRecipientService) List(ctx context.Context, accountID string, shareID string, query AccountShareRecipientListParams, opts ...option.RequestOption) (res *AccountShareRecipientListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/recipients", accountID, shareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletion is not immediate, an updated share recipient object with a new status
// will be returned.
func (r *AccountShareRecipientService) Delete(ctx context.Context, accountID string, shareID string, recipientID string, opts ...option.RequestOption) (res *ShareRecipientResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	if recipientID == "" {
		err = errors.New("missing required recipient_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/recipients/%s", accountID, shareID, recipientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Account or organization ID must be provided.
type CreateShareRecipientRequestParam struct {
	// Account identifier.
	AccountID param.Field[string] `json:"account_id"`
	// Organization identifier.
	OrganizationID param.Field[string] `json:"organization_id"`
}

func (r CreateShareRecipientRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ShareRecipientObject struct {
	// Share Recipient identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// Share Recipient association status.
	AssociationStatus ShareRecipientObjectAssociationStatus `json:"association_status,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Share Recipient status message.
	StatusMessage string                   `json:"status_message,required"`
	JSON          shareRecipientObjectJSON `json:"-"`
}

// shareRecipientObjectJSON contains the JSON metadata for the struct
// [ShareRecipientObject]
type shareRecipientObjectJSON struct {
	ID                apijson.Field
	AccountID         apijson.Field
	AssociationStatus apijson.Field
	Created           apijson.Field
	Modified          apijson.Field
	StatusMessage     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ShareRecipientObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r shareRecipientObjectJSON) RawJSON() string {
	return r.raw
}

// Share Recipient association status.
type ShareRecipientObjectAssociationStatus string

const (
	ShareRecipientObjectAssociationStatusAssociating    ShareRecipientObjectAssociationStatus = "associating"
	ShareRecipientObjectAssociationStatusAssociated     ShareRecipientObjectAssociationStatus = "associated"
	ShareRecipientObjectAssociationStatusDisassociating ShareRecipientObjectAssociationStatus = "disassociating"
	ShareRecipientObjectAssociationStatusDisassociated  ShareRecipientObjectAssociationStatus = "disassociated"
)

func (r ShareRecipientObjectAssociationStatus) IsKnown() bool {
	switch r {
	case ShareRecipientObjectAssociationStatusAssociating, ShareRecipientObjectAssociationStatusAssociated, ShareRecipientObjectAssociationStatusDisassociating, ShareRecipientObjectAssociationStatusDisassociated:
		return true
	}
	return false
}

type ShareRecipientResponseSingle struct {
	Errors []ShareRecipientResponseSingleError `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                             `json:"success,required"`
	Result  ShareRecipientObject             `json:"result"`
	JSON    shareRecipientResponseSingleJSON `json:"-"`
}

// shareRecipientResponseSingleJSON contains the JSON metadata for the struct
// [ShareRecipientResponseSingle]
type shareRecipientResponseSingleJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ShareRecipientResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r shareRecipientResponseSingleJSON) RawJSON() string {
	return r.raw
}

type ShareRecipientResponseSingleError struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           ShareRecipientResponseSingleErrorsSource `json:"source"`
	JSON             shareRecipientResponseSingleErrorJSON    `json:"-"`
}

// shareRecipientResponseSingleErrorJSON contains the JSON metadata for the struct
// [ShareRecipientResponseSingleError]
type shareRecipientResponseSingleErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ShareRecipientResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r shareRecipientResponseSingleErrorJSON) RawJSON() string {
	return r.raw
}

type ShareRecipientResponseSingleErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    shareRecipientResponseSingleErrorsSourceJSON `json:"-"`
}

// shareRecipientResponseSingleErrorsSourceJSON contains the JSON metadata for the
// struct [ShareRecipientResponseSingleErrorsSource]
type shareRecipientResponseSingleErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ShareRecipientResponseSingleErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r shareRecipientResponseSingleErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountShareRecipientListResponse struct {
	Errors []AccountShareRecipientListResponseError `json:"errors,required"`
	// Whether the API call was successful.
	Success    bool                                        `json:"success,required"`
	Result     []ShareRecipientObject                      `json:"result,nullable"`
	ResultInfo AccountShareRecipientListResponseResultInfo `json:"result_info"`
	JSON       accountShareRecipientListResponseJSON       `json:"-"`
}

// accountShareRecipientListResponseJSON contains the JSON metadata for the struct
// [AccountShareRecipientListResponse]
type accountShareRecipientListResponseJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountShareRecipientListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountShareRecipientListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountShareRecipientListResponseError struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           AccountShareRecipientListResponseErrorsSource `json:"source"`
	JSON             accountShareRecipientListResponseErrorJSON    `json:"-"`
}

// accountShareRecipientListResponseErrorJSON contains the JSON metadata for the
// struct [AccountShareRecipientListResponseError]
type accountShareRecipientListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountShareRecipientListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountShareRecipientListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountShareRecipientListResponseErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    accountShareRecipientListResponseErrorsSourceJSON `json:"-"`
}

// accountShareRecipientListResponseErrorsSourceJSON contains the JSON metadata for
// the struct [AccountShareRecipientListResponseErrorsSource]
type accountShareRecipientListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountShareRecipientListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountShareRecipientListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountShareRecipientListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64 `json:"total_count"`
	// Total number of pages using the given per page.
	TotalPages float64                                         `json:"total_pages"`
	JSON       accountShareRecipientListResponseResultInfoJSON `json:"-"`
}

// accountShareRecipientListResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountShareRecipientListResponseResultInfo]
type accountShareRecipientListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountShareRecipientListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountShareRecipientListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountShareRecipientNewParams struct {
	// Account or organization ID must be provided.
	CreateShareRecipientRequest CreateShareRecipientRequestParam `json:"create_share_recipient_request,required"`
}

func (r AccountShareRecipientNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateShareRecipientRequest)
}

type AccountShareRecipientListParams struct {
	// Page number.
	Page param.Field[int64] `query:"page"`
	// Number of objects to return per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [AccountShareRecipientListParams]'s query parameters as
// `url.Values`.
func (r AccountShareRecipientListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
