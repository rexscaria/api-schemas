// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountCloudforceOneEventTagService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCloudforceOneEventTagService] method instead.
type AccountCloudforceOneEventTagService struct {
	Options []option.RequestOption
}

// NewAccountCloudforceOneEventTagService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCloudforceOneEventTagService(opts ...option.RequestOption) (r *AccountCloudforceOneEventTagService) {
	r = &AccountCloudforceOneEventTagService{}
	r.Options = opts
	return
}

// Creates a new tag
func (r *AccountCloudforceOneEventTagService) New(ctx context.Context, accountID float64, body AccountCloudforceOneEventTagNewParams, opts ...option.RequestOption) (res *AccountCloudforceOneEventTagNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/tags/create", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountCloudforceOneEventTagNewResponse struct {
	Name string                                      `json:"name,required"`
	Uuid string                                      `json:"uuid,required"`
	JSON accountCloudforceOneEventTagNewResponseJSON `json:"-"`
}

// accountCloudforceOneEventTagNewResponseJSON contains the JSON metadata for the
// struct [AccountCloudforceOneEventTagNewResponse]
type accountCloudforceOneEventTagNewResponseJSON struct {
	Name        apijson.Field
	Uuid        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneEventTagNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneEventTagNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneEventTagNewParams struct {
	Name param.Field[string] `json:"name,required"`
}

func (r AccountCloudforceOneEventTagNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
