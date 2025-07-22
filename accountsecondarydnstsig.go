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

// AccountSecondaryDNSTsigService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountSecondaryDNSTsigService] method instead.
type AccountSecondaryDNSTsigService struct {
	Options []option.RequestOption
}

// NewAccountSecondaryDNSTsigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountSecondaryDNSTsigService(opts ...option.RequestOption) (r *AccountSecondaryDNSTsigService) {
	r = &AccountSecondaryDNSTsigService{}
	r.Options = opts
	return
}

// Create TSIG.
func (r *AccountSecondaryDNSTsigService) New(ctx context.Context, accountID string, body AccountSecondaryDNSTsigNewParams, opts ...option.RequestOption) (res *SingleResponseTsigs, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/tsigs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get TSIG.
func (r *AccountSecondaryDNSTsigService) Get(ctx context.Context, accountID string, tsigID string, opts ...option.RequestOption) (res *SingleResponseTsigs, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tsigID == "" {
		err = errors.New("missing required tsig_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/tsigs/%s", accountID, tsigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify TSIG.
func (r *AccountSecondaryDNSTsigService) Update(ctx context.Context, accountID string, tsigID string, body AccountSecondaryDNSTsigUpdateParams, opts ...option.RequestOption) (res *SingleResponseTsigs, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tsigID == "" {
		err = errors.New("missing required tsig_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/tsigs/%s", accountID, tsigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List TSIGs.
func (r *AccountSecondaryDNSTsigService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountSecondaryDNSTsigListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/tsigs", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete TSIG.
func (r *AccountSecondaryDNSTsigService) Delete(ctx context.Context, accountID string, tsigID string, body AccountSecondaryDNSTsigDeleteParams, opts ...option.RequestOption) (res *AccountSecondaryDNSTsigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tsigID == "" {
		err = errors.New("missing required tsig_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/tsigs/%s", accountID, tsigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type SingleResponseTsigs struct {
	Result Tsig                    `json:"result"`
	JSON   singleResponseTsigsJSON `json:"-"`
	ResponseSingleACL
}

// singleResponseTsigsJSON contains the JSON metadata for the struct
// [SingleResponseTsigs]
type singleResponseTsigsJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseTsigs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseTsigsJSON) RawJSON() string {
	return r.raw
}

type Tsig struct {
	ID string `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string   `json:"secret,required"`
	JSON   tsigJSON `json:"-"`
}

// tsigJSON contains the JSON metadata for the struct [Tsig]
type tsigJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Tsig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigJSON) RawJSON() string {
	return r.raw
}

type TsigParam struct {
	// TSIG algorithm.
	Algo param.Field[string] `json:"algo,required"`
	// TSIG key name.
	Name param.Field[string] `json:"name,required"`
	// TSIG secret.
	Secret param.Field[string] `json:"secret,required"`
}

func (r TsigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountSecondaryDNSTsigListResponse struct {
	Result []Tsig                                  `json:"result"`
	JSON   accountSecondaryDNSTsigListResponseJSON `json:"-"`
	ResponseCollectionDNSACLs
}

// accountSecondaryDNSTsigListResponseJSON contains the JSON metadata for the
// struct [AccountSecondaryDNSTsigListResponse]
type accountSecondaryDNSTsigListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSecondaryDNSTsigListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSecondaryDNSTsigDeleteResponse struct {
	Result AccountSecondaryDNSTsigDeleteResponseResult `json:"result"`
	JSON   accountSecondaryDNSTsigDeleteResponseJSON   `json:"-"`
	ResponseSingleACL
}

// accountSecondaryDNSTsigDeleteResponseJSON contains the JSON metadata for the
// struct [AccountSecondaryDNSTsigDeleteResponse]
type accountSecondaryDNSTsigDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSecondaryDNSTsigDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSecondaryDNSTsigDeleteResponseResult struct {
	ID   string                                          `json:"id"`
	JSON accountSecondaryDNSTsigDeleteResponseResultJSON `json:"-"`
}

// accountSecondaryDNSTsigDeleteResponseResultJSON contains the JSON metadata for
// the struct [AccountSecondaryDNSTsigDeleteResponseResult]
type accountSecondaryDNSTsigDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSecondaryDNSTsigDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountSecondaryDNSTsigNewParams struct {
	Tsig TsigParam `json:"tsig,required"`
}

func (r AccountSecondaryDNSTsigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Tsig)
}

type AccountSecondaryDNSTsigUpdateParams struct {
	Tsig TsigParam `json:"tsig,required"`
}

func (r AccountSecondaryDNSTsigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Tsig)
}

type AccountSecondaryDNSTsigDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountSecondaryDNSTsigDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
