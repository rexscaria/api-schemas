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

// AccountCloudforceOneEventRawService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCloudforceOneEventRawService] method instead.
type AccountCloudforceOneEventRawService struct {
	Options []option.RequestOption
}

// NewAccountCloudforceOneEventRawService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCloudforceOneEventRawService(opts ...option.RequestOption) (r *AccountCloudforceOneEventRawService) {
	r = &AccountCloudforceOneEventRawService{}
	r.Options = opts
	return
}

// Reads data for a raw event
func (r *AccountCloudforceOneEventRawService) Get(ctx context.Context, accountID float64, datasetID string, eventID string, opts ...option.RequestOption) (res *AccountCloudforceOneEventRawGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/raw/%s/%s", accountID, datasetID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a raw event
func (r *AccountCloudforceOneEventRawService) Update(ctx context.Context, accountID float64, eventID string, rawID string, body AccountCloudforceOneEventRawUpdateParams, opts ...option.RequestOption) (res *AccountCloudforceOneEventRawUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	if rawID == "" {
		err = errors.New("missing required raw_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/%s/raw/%s", accountID, eventID, rawID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountCloudforceOneEventRawGetResponse struct {
	ID        string                                      `json:"id,required"`
	AccountID float64                                     `json:"accountId,required"`
	Created   string                                      `json:"created,required"`
	Data      interface{}                                 `json:"data,required"`
	Source    string                                      `json:"source,required"`
	Tlp       string                                      `json:"tlp,required"`
	JSON      accountCloudforceOneEventRawGetResponseJSON `json:"-"`
}

// accountCloudforceOneEventRawGetResponseJSON contains the JSON metadata for the
// struct [AccountCloudforceOneEventRawGetResponse]
type accountCloudforceOneEventRawGetResponseJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Created     apijson.Field
	Data        apijson.Field
	Source      apijson.Field
	Tlp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventRawGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventRawGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventRawUpdateResponse struct {
	ID   string                                         `json:"id,required"`
	Data interface{}                                    `json:"data,required"`
	JSON accountCloudforceOneEventRawUpdateResponseJSON `json:"-"`
}

// accountCloudforceOneEventRawUpdateResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneEventRawUpdateResponse]
type accountCloudforceOneEventRawUpdateResponseJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventRawUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventRawUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventRawUpdateParams struct {
	Data   param.Field[interface{}] `json:"data"`
	Source param.Field[string]      `json:"source"`
	Tlp    param.Field[string]      `json:"tlp"`
}

func (r AccountCloudforceOneEventRawUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
