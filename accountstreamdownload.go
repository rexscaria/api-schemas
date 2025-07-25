// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountStreamDownloadService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountStreamDownloadService] method instead.
type AccountStreamDownloadService struct {
	Options []option.RequestOption
}

// NewAccountStreamDownloadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamDownloadService(opts ...option.RequestOption) (r *AccountStreamDownloadService) {
	r = &AccountStreamDownloadService{}
	r.Options = opts
	return
}

// Creates a download for a video when a video is ready to view.
func (r *AccountStreamDownloadService) New(ctx context.Context, accountID string, identifier string, body AccountStreamDownloadNewParams, opts ...option.RequestOption) (res *DownloadsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists the downloads created for a video.
func (r *AccountStreamDownloadService) List(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *DownloadsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete the downloads for a video.
func (r *AccountStreamDownloadService) Delete(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *DeletedStreamResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type DownloadsResponse struct {
	Errors   []StreamMessages `json:"errors,required"`
	Messages []StreamMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DownloadsResponseSuccess `json:"success,required"`
	Result  interface{}              `json:"result"`
	JSON    downloadsResponseJSON    `json:"-"`
}

// downloadsResponseJSON contains the JSON metadata for the struct
// [DownloadsResponse]
type downloadsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DownloadsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadsResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DownloadsResponseSuccess bool

const (
	DownloadsResponseSuccessTrue DownloadsResponseSuccess = true
)

func (r DownloadsResponseSuccess) IsKnown() bool {
	switch r {
	case DownloadsResponseSuccessTrue:
		return true
	}
	return false
}

type AccountStreamDownloadNewParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountStreamDownloadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
