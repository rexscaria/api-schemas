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

// AccountCloudforceOneEventCronService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCloudforceOneEventCronService] method instead.
type AccountCloudforceOneEventCronService struct {
	Options []option.RequestOption
}

// NewAccountCloudforceOneEventCronService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCloudforceOneEventCronService(opts ...option.RequestOption) (r *AccountCloudforceOneEventCronService) {
	r = &AccountCloudforceOneEventCronService{}
	r.Options = opts
	return
}

// Reads the last cron update time
func (r *AccountCloudforceOneEventCronService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountCloudforceOneEventCronGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/cron", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Reads the last cron update time
func (r *AccountCloudforceOneEventCronService) Update(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountCloudforceOneEventCronUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/cron", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AccountCloudforceOneEventCronGetResponse struct {
	Update string                                       `json:"update,required"`
	JSON   accountCloudforceOneEventCronGetResponseJSON `json:"-"`
}

// accountCloudforceOneEventCronGetResponseJSON contains the JSON metadata for the
// struct [AccountCloudforceOneEventCronGetResponse]
type accountCloudforceOneEventCronGetResponseJSON struct {
	Update      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventCronGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventCronGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventCronUpdateResponse struct {
	ID     float64                                         `json:"id,required"`
	Update string                                          `json:"update,required"`
	JSON   accountCloudforceOneEventCronUpdateResponseJSON `json:"-"`
}

// accountCloudforceOneEventCronUpdateResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneEventCronUpdateResponse]
type accountCloudforceOneEventCronUpdateResponseJSON struct {
	ID          apijson.Field
	Update      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventCronUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventCronUpdateResponseJSON) RawJSON() string {
	return r.raw
}
