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

// ZoneWeb3HostnameService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneWeb3HostnameService] method instead.
type ZoneWeb3HostnameService struct {
	Options           []option.RequestOption
	IpfsUniversalPath *ZoneWeb3HostnameIpfsUniversalPathService
}

// NewZoneWeb3HostnameService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneWeb3HostnameService(opts ...option.RequestOption) (r *ZoneWeb3HostnameService) {
	r = &ZoneWeb3HostnameService{}
	r.Options = opts
	r.IpfsUniversalPath = NewZoneWeb3HostnameIpfsUniversalPathService(opts...)
	return
}

// Create Web3 Hostname
func (r *ZoneWeb3HostnameService) New(ctx context.Context, zoneID string, body ZoneWeb3HostnameNewParams, opts ...option.RequestOption) (res *SingleResponseWeb3, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Web3 Hostname Details
func (r *ZoneWeb3HostnameService) Get(ctx context.Context, zoneID string, identifier string, opts ...option.RequestOption) (res *SingleResponseWeb3, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s", zoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Web3 Hostnames
func (r *ZoneWeb3HostnameService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneWeb3HostnameListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Web3 Hostname
func (r *ZoneWeb3HostnameService) Delete(ctx context.Context, zoneID string, identifier string, opts ...option.RequestOption) (res *APIResponseSingleID, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s", zoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Edit Web3 Hostname
func (r *ZoneWeb3HostnameService) Patch(ctx context.Context, zoneID string, identifier string, body ZoneWeb3HostnamePatchParams, opts ...option.RequestOption) (res *SingleResponseWeb3, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s", zoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type APIResponseSingleID struct {
	Errors   []APIResponseSingleIDError   `json:"errors,required"`
	Messages []APIResponseSingleIDMessage `json:"messages,required"`
	Result   APIResponseSingleIDResult    `json:"result,required,nullable"`
	// Specifies whether the API call was successful.
	Success APIResponseSingleIDSuccess `json:"success,required"`
	JSON    apiResponseSingleIDJSON    `json:"-"`
}

// apiResponseSingleIDJSON contains the JSON metadata for the struct
// [APIResponseSingleID]
type apiResponseSingleIDJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleIDJSON) RawJSON() string {
	return r.raw
}

type APIResponseSingleIDError struct {
	Code             int64                           `json:"code,required"`
	Message          string                          `json:"message,required"`
	DocumentationURL string                          `json:"documentation_url"`
	Source           APIResponseSingleIDErrorsSource `json:"source"`
	JSON             apiResponseSingleIDErrorJSON    `json:"-"`
}

// apiResponseSingleIDErrorJSON contains the JSON metadata for the struct
// [APIResponseSingleIDError]
type apiResponseSingleIDErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIResponseSingleIDError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleIDErrorJSON) RawJSON() string {
	return r.raw
}

type APIResponseSingleIDErrorsSource struct {
	Pointer string                              `json:"pointer"`
	JSON    apiResponseSingleIDErrorsSourceJSON `json:"-"`
}

// apiResponseSingleIDErrorsSourceJSON contains the JSON metadata for the struct
// [APIResponseSingleIDErrorsSource]
type apiResponseSingleIDErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleIDErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleIDErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type APIResponseSingleIDMessage struct {
	Code             int64                             `json:"code,required"`
	Message          string                            `json:"message,required"`
	DocumentationURL string                            `json:"documentation_url"`
	Source           APIResponseSingleIDMessagesSource `json:"source"`
	JSON             apiResponseSingleIDMessageJSON    `json:"-"`
}

// apiResponseSingleIDMessageJSON contains the JSON metadata for the struct
// [APIResponseSingleIDMessage]
type apiResponseSingleIDMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIResponseSingleIDMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleIDMessageJSON) RawJSON() string {
	return r.raw
}

type APIResponseSingleIDMessagesSource struct {
	Pointer string                                `json:"pointer"`
	JSON    apiResponseSingleIDMessagesSourceJSON `json:"-"`
}

// apiResponseSingleIDMessagesSourceJSON contains the JSON metadata for the struct
// [APIResponseSingleIDMessagesSource]
type apiResponseSingleIDMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleIDMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleIDMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type APIResponseSingleIDResult struct {
	// Specify the identifier of the hostname.
	ID   string                        `json:"id,required"`
	JSON apiResponseSingleIDResultJSON `json:"-"`
}

// apiResponseSingleIDResultJSON contains the JSON metadata for the struct
// [APIResponseSingleIDResult]
type apiResponseSingleIDResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleIDResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleIDResultJSON) RawJSON() string {
	return r.raw
}

// Specifies whether the API call was successful.
type APIResponseSingleIDSuccess bool

const (
	APIResponseSingleIDSuccessTrue APIResponseSingleIDSuccess = true
)

func (r APIResponseSingleIDSuccess) IsKnown() bool {
	switch r {
	case APIResponseSingleIDSuccessTrue:
		return true
	}
	return false
}

type MessageItems struct {
	Code             int64              `json:"code,required"`
	Message          string             `json:"message,required"`
	DocumentationURL string             `json:"documentation_url"`
	Source           MessageItemsSource `json:"source"`
	JSON             messageItemsJSON   `json:"-"`
}

// messageItemsJSON contains the JSON metadata for the struct [MessageItems]
type messageItemsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MessageItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageItemsJSON) RawJSON() string {
	return r.raw
}

type MessageItemsSource struct {
	Pointer string                 `json:"pointer"`
	JSON    messageItemsSourceJSON `json:"-"`
}

// messageItemsSourceJSON contains the JSON metadata for the struct
// [MessageItemsSource]
type messageItemsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageItemsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageItemsSourceJSON) RawJSON() string {
	return r.raw
}

type SingleResponseWeb3 struct {
	Errors   []MessageItems `json:"errors,required"`
	Messages []MessageItems `json:"messages,required"`
	Result   Web3Hostname   `json:"result,required"`
	// Specifies whether the API call was successful.
	Success SingleResponseWeb3Success `json:"success,required"`
	// Provides the API response.
	ResultInfo interface{}            `json:"result_info"`
	JSON       singleResponseWeb3JSON `json:"-"`
}

// singleResponseWeb3JSON contains the JSON metadata for the struct
// [SingleResponseWeb3]
type singleResponseWeb3JSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseWeb3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseWeb3JSON) RawJSON() string {
	return r.raw
}

// Specifies whether the API call was successful.
type SingleResponseWeb3Success bool

const (
	SingleResponseWeb3SuccessTrue SingleResponseWeb3Success = true
)

func (r SingleResponseWeb3Success) IsKnown() bool {
	switch r {
	case SingleResponseWeb3SuccessTrue:
		return true
	}
	return false
}

// Specify the target gateway of the hostname.
type TargetGateway string

const (
	TargetGatewayEthereum          TargetGateway = "ethereum"
	TargetGatewayIpfs              TargetGateway = "ipfs"
	TargetGatewayIpfsUniversalPath TargetGateway = "ipfs_universal_path"
)

func (r TargetGateway) IsKnown() bool {
	switch r {
	case TargetGatewayEthereum, TargetGatewayIpfs, TargetGatewayIpfsUniversalPath:
		return true
	}
	return false
}

type Web3Hostname struct {
	// Specify the identifier of the hostname.
	ID        string    `json:"id"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Specify an optional description of the hostname.
	Description string `json:"description"`
	// Specify the DNSLink value used if the target is ipfs.
	Dnslink    string    `json:"dnslink"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Specify the hostname that points to the target gateway via CNAME.
	Name string `json:"name"`
	// Specifies the status of the hostname's activation.
	Status Web3HostnameStatus `json:"status"`
	// Specify the target gateway of the hostname.
	Target TargetGateway    `json:"target"`
	JSON   web3HostnameJSON `json:"-"`
}

// web3HostnameJSON contains the JSON metadata for the struct [Web3Hostname]
type web3HostnameJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	Dnslink     apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3Hostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameJSON) RawJSON() string {
	return r.raw
}

// Specifies the status of the hostname's activation.
type Web3HostnameStatus string

const (
	Web3HostnameStatusActive   Web3HostnameStatus = "active"
	Web3HostnameStatusPending  Web3HostnameStatus = "pending"
	Web3HostnameStatusDeleting Web3HostnameStatus = "deleting"
	Web3HostnameStatusError    Web3HostnameStatus = "error"
)

func (r Web3HostnameStatus) IsKnown() bool {
	switch r {
	case Web3HostnameStatusActive, Web3HostnameStatusPending, Web3HostnameStatusDeleting, Web3HostnameStatusError:
		return true
	}
	return false
}

type ZoneWeb3HostnameListResponse struct {
	Errors   []ZoneWeb3HostnameListResponseError   `json:"errors,required"`
	Messages []ZoneWeb3HostnameListResponseMessage `json:"messages,required"`
	Result   []Web3Hostname                        `json:"result,required,nullable"`
	// Specifies whether the API call was successful.
	Success    ZoneWeb3HostnameListResponseSuccess    `json:"success,required"`
	ResultInfo ZoneWeb3HostnameListResponseResultInfo `json:"result_info"`
	JSON       zoneWeb3HostnameListResponseJSON       `json:"-"`
}

// zoneWeb3HostnameListResponseJSON contains the JSON metadata for the struct
// [ZoneWeb3HostnameListResponse]
type zoneWeb3HostnameListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWeb3HostnameListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneWeb3HostnameListResponseError struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           ZoneWeb3HostnameListResponseErrorsSource `json:"source"`
	JSON             zoneWeb3HostnameListResponseErrorJSON    `json:"-"`
}

// zoneWeb3HostnameListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneWeb3HostnameListResponseError]
type zoneWeb3HostnameListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneWeb3HostnameListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWeb3HostnameListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ZoneWeb3HostnameListResponseErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    zoneWeb3HostnameListResponseErrorsSourceJSON `json:"-"`
}

// zoneWeb3HostnameListResponseErrorsSourceJSON contains the JSON metadata for the
// struct [ZoneWeb3HostnameListResponseErrorsSource]
type zoneWeb3HostnameListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWeb3HostnameListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneWeb3HostnameListResponseMessage struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           ZoneWeb3HostnameListResponseMessagesSource `json:"source"`
	JSON             zoneWeb3HostnameListResponseMessageJSON    `json:"-"`
}

// zoneWeb3HostnameListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneWeb3HostnameListResponseMessage]
type zoneWeb3HostnameListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneWeb3HostnameListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWeb3HostnameListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ZoneWeb3HostnameListResponseMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    zoneWeb3HostnameListResponseMessagesSourceJSON `json:"-"`
}

// zoneWeb3HostnameListResponseMessagesSourceJSON contains the JSON metadata for
// the struct [ZoneWeb3HostnameListResponseMessagesSource]
type zoneWeb3HostnameListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWeb3HostnameListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Specifies whether the API call was successful.
type ZoneWeb3HostnameListResponseSuccess bool

const (
	ZoneWeb3HostnameListResponseSuccessTrue ZoneWeb3HostnameListResponseSuccess = true
)

func (r ZoneWeb3HostnameListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneWeb3HostnameListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneWeb3HostnameListResponseResultInfo struct {
	// Specifies the total number of results for the requested service.
	Count float64 `json:"count"`
	// Specifies the current page within paginated list of results.
	Page float64 `json:"page"`
	// Specifies the number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Specifies the total results available without any search parameters.
	TotalCount float64                                    `json:"total_count"`
	JSON       zoneWeb3HostnameListResponseResultInfoJSON `json:"-"`
}

// zoneWeb3HostnameListResponseResultInfoJSON contains the JSON metadata for the
// struct [ZoneWeb3HostnameListResponseResultInfo]
type zoneWeb3HostnameListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWeb3HostnameListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneWeb3HostnameNewParams struct {
	// Specify the hostname that points to the target gateway via CNAME.
	Name param.Field[string] `json:"name,required"`
	// Specify the target gateway of the hostname.
	Target param.Field[TargetGateway] `json:"target,required"`
	// Specify an optional description of the hostname.
	Description param.Field[string] `json:"description"`
	// Specify the DNSLink value used if the target is ipfs.
	Dnslink param.Field[string] `json:"dnslink"`
}

func (r ZoneWeb3HostnameNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneWeb3HostnamePatchParams struct {
	// Specify an optional description of the hostname.
	Description param.Field[string] `json:"description"`
	// Specify the DNSLink value used if the target is ipfs.
	Dnslink param.Field[string] `json:"dnslink"`
}

func (r ZoneWeb3HostnamePatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
