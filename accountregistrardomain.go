// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// AccountRegistrarDomainService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountRegistrarDomainService] method instead.
type AccountRegistrarDomainService struct {
	Options []option.RequestOption
}

// NewAccountRegistrarDomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRegistrarDomainService(opts ...option.RequestOption) (r *AccountRegistrarDomainService) {
	r = &AccountRegistrarDomainService{}
	r.Options = opts
	return
}

// Show individual domain.
func (r *AccountRegistrarDomainService) Get(ctx context.Context, accountID string, domainName string, opts ...option.RequestOption) (res *RegistrarAPIDomainResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/registrar/domains/%s", accountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update individual domain.
func (r *AccountRegistrarDomainService) Update(ctx context.Context, accountID string, domainName string, body AccountRegistrarDomainUpdateParams, opts ...option.RequestOption) (res *RegistrarAPIDomainResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/registrar/domains/%s", accountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List domains handled by Registrar.
func (r *AccountRegistrarDomainService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountRegistrarDomainListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/registrar/domains", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RegistrarAPIAPIResponseCommon struct {
	Errors   []RegistrarAPIMessagesItems              `json:"errors,required"`
	Messages []RegistrarAPIMessagesItems              `json:"messages,required"`
	Result   RegistrarAPIAPIResponseCommonResultUnion `json:"result,required"`
	// Whether the API call was successful
	Success RegistrarAPIAPIResponseCommonSuccess `json:"success,required"`
	JSON    registrarApiapiResponseCommonJSON    `json:"-"`
}

// registrarApiapiResponseCommonJSON contains the JSON metadata for the struct
// [RegistrarAPIAPIResponseCommon]
type registrarApiapiResponseCommonJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarAPIAPIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarApiapiResponseCommonJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [RegistrarAPIAPIResponseCommonResultArray] or
// [shared.UnionString].
type RegistrarAPIAPIResponseCommonResultUnion interface {
	ImplementsRegistrarApiapiResponseCommonResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RegistrarAPIAPIResponseCommonResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RegistrarAPIAPIResponseCommonResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RegistrarAPIAPIResponseCommonResultArray []interface{}

func (r RegistrarAPIAPIResponseCommonResultArray) ImplementsRegistrarApiapiResponseCommonResultUnion() {
}

// Whether the API call was successful
type RegistrarAPIAPIResponseCommonSuccess bool

const (
	RegistrarAPIAPIResponseCommonSuccessTrue RegistrarAPIAPIResponseCommonSuccess = true
)

func (r RegistrarAPIAPIResponseCommonSuccess) IsKnown() bool {
	switch r {
	case RegistrarAPIAPIResponseCommonSuccessTrue:
		return true
	}
	return false
}

type RegistrarAPIDomainResponseSingle struct {
	Result interface{}                          `json:"result,nullable"`
	JSON   registrarAPIDomainResponseSingleJSON `json:"-"`
	RegistrarAPIAPIResponseCommon
}

// registrarAPIDomainResponseSingleJSON contains the JSON metadata for the struct
// [RegistrarAPIDomainResponseSingle]
type registrarAPIDomainResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarAPIDomainResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarAPIDomainResponseSingleJSON) RawJSON() string {
	return r.raw
}

type RegistrarAPIMessagesItems struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    registrarAPIMessagesItemsJSON `json:"-"`
}

// registrarAPIMessagesItemsJSON contains the JSON metadata for the struct
// [RegistrarAPIMessagesItems]
type registrarAPIMessagesItemsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarAPIMessagesItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarAPIMessagesItemsJSON) RawJSON() string {
	return r.raw
}

type AccountRegistrarDomainListResponse struct {
	Result     []interface{}                                `json:"result,nullable"`
	ResultInfo AccountRegistrarDomainListResponseResultInfo `json:"result_info"`
	JSON       accountRegistrarDomainListResponseJSON       `json:"-"`
	RegistrarAPIAPIResponseCommon
}

// accountRegistrarDomainListResponseJSON contains the JSON metadata for the struct
// [AccountRegistrarDomainListResponse]
type accountRegistrarDomainListResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRegistrarDomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRegistrarDomainListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountRegistrarDomainListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       accountRegistrarDomainListResponseResultInfoJSON `json:"-"`
}

// accountRegistrarDomainListResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountRegistrarDomainListResponseResultInfo]
type accountRegistrarDomainListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRegistrarDomainListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRegistrarDomainListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountRegistrarDomainUpdateParams struct {
	// Auto-renew controls whether subscription is automatically renewed upon domain
	// expiration.
	AutoRenew param.Field[bool] `json:"auto_renew"`
	// Shows whether a registrar lock is in place for a domain.
	Locked param.Field[bool] `json:"locked"`
	// Privacy option controls redacting WHOIS information.
	Privacy param.Field[bool] `json:"privacy"`
}

func (r AccountRegistrarDomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
