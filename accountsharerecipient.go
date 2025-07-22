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
	Result ShareRecipientObject             `json:"result"`
	JSON   shareRecipientResponseSingleJSON `json:"-"`
	APIResponseResourceSharing
}

// shareRecipientResponseSingleJSON contains the JSON metadata for the struct
// [ShareRecipientResponseSingle]
type shareRecipientResponseSingleJSON struct {
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

type AccountShareRecipientListResponse struct {
	Result []ShareRecipientObject                `json:"result"`
	JSON   accountShareRecipientListResponseJSON `json:"-"`
	APIResponseCollectionResourceSharing
}

// accountShareRecipientListResponseJSON contains the JSON metadata for the struct
// [AccountShareRecipientListResponse]
type accountShareRecipientListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountShareRecipientListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountShareRecipientListResponseJSON) RawJSON() string {
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
