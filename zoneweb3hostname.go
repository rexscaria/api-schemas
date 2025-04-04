// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/stainless-sdks/cf-rex-go/shared"
	"github.com/tidwall/gjson"
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
func (r *ZoneWeb3HostnameService) Delete(ctx context.Context, zoneID string, identifier string, body ZoneWeb3HostnameDeleteParams, opts ...option.RequestOption) (res *APIResponseSingleID, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
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

type APIResponseCollectionWeb3 struct {
	Result     []interface{}                       `json:"result,nullable"`
	ResultInfo APIResponseCollectionWeb3ResultInfo `json:"result_info"`
	JSON       apiResponseCollectionWeb3JSON       `json:"-"`
	APIResponseWeb3
}

// apiResponseCollectionWeb3JSON contains the JSON metadata for the struct
// [APIResponseCollectionWeb3]
type apiResponseCollectionWeb3JSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionWeb3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionWeb3JSON) RawJSON() string {
	return r.raw
}

type APIResponseCollectionWeb3ResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count"`
	JSON       apiResponseCollectionWeb3ResultInfoJSON `json:"-"`
}

// apiResponseCollectionWeb3ResultInfoJSON contains the JSON metadata for the
// struct [APIResponseCollectionWeb3ResultInfo]
type apiResponseCollectionWeb3ResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionWeb3ResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionWeb3ResultInfoJSON) RawJSON() string {
	return r.raw
}

type APIResponseSingleID struct {
	Result APIResponseSingleIDResult `json:"result,nullable"`
	JSON   apiResponseSingleIDJSON   `json:"-"`
	APIResponseWeb3
}

// apiResponseSingleIDJSON contains the JSON metadata for the struct
// [APIResponseSingleID]
type apiResponseSingleIDJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleIDJSON) RawJSON() string {
	return r.raw
}

type APIResponseSingleIDResult struct {
	// Identifier
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

type APIResponseSingleWeb3 struct {
	Result interface{}               `json:"result"`
	JSON   apiResponseSingleWeb3JSON `json:"-"`
	APIResponseWeb3
}

// apiResponseSingleWeb3JSON contains the JSON metadata for the struct
// [APIResponseSingleWeb3]
type apiResponseSingleWeb3JSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleWeb3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleWeb3JSON) RawJSON() string {
	return r.raw
}

type APIResponseWeb3 struct {
	Errors   []MessageItems             `json:"errors,required"`
	Messages []MessageItems             `json:"messages,required"`
	Result   APIResponseWeb3ResultUnion `json:"result,required"`
	// Whether the API call was successful
	Success APIResponseWeb3Success `json:"success,required"`
	JSON    apiResponseWeb3JSON    `json:"-"`
}

// apiResponseWeb3JSON contains the JSON metadata for the struct [APIResponseWeb3]
type apiResponseWeb3JSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseWeb3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseWeb3JSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [APIResponseWeb3ResultArray] or [shared.UnionString].
type APIResponseWeb3ResultUnion interface {
	ImplementsAPIResponseWeb3ResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIResponseWeb3ResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIResponseWeb3ResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type APIResponseWeb3ResultArray []interface{}

func (r APIResponseWeb3ResultArray) ImplementsAPIResponseWeb3ResultUnion() {}

// Whether the API call was successful
type APIResponseWeb3Success bool

const (
	APIResponseWeb3SuccessTrue APIResponseWeb3Success = true
)

func (r APIResponseWeb3Success) IsKnown() bool {
	switch r {
	case APIResponseWeb3SuccessTrue:
		return true
	}
	return false
}

type MessageItems struct {
	Code    int64            `json:"code,required"`
	Message string           `json:"message,required"`
	JSON    messageItemsJSON `json:"-"`
}

// messageItemsJSON contains the JSON metadata for the struct [MessageItems]
type messageItemsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageItemsJSON) RawJSON() string {
	return r.raw
}

type SingleResponseWeb3 struct {
	Result Web3Hostname           `json:"result"`
	JSON   singleResponseWeb3JSON `json:"-"`
	APIResponseSingleWeb3
}

// singleResponseWeb3JSON contains the JSON metadata for the struct
// [SingleResponseWeb3]
type singleResponseWeb3JSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseWeb3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseWeb3JSON) RawJSON() string {
	return r.raw
}

// Target gateway of the hostname.
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
	// Identifier
	ID        string    `json:"id"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the hostname.
	Description string `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink    string    `json:"dnslink"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The hostname that will point to the target gateway via CNAME.
	Name string `json:"name"`
	// Status of the hostname's activation.
	Status Web3HostnameStatus `json:"status"`
	// Target gateway of the hostname.
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

// Status of the hostname's activation.
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
	Result []Web3Hostname                   `json:"result"`
	JSON   zoneWeb3HostnameListResponseJSON `json:"-"`
	APIResponseCollectionWeb3
}

// zoneWeb3HostnameListResponseJSON contains the JSON metadata for the struct
// [ZoneWeb3HostnameListResponse]
type zoneWeb3HostnameListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWeb3HostnameListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneWeb3HostnameNewParams struct {
	// The hostname that will point to the target gateway via CNAME.
	Name param.Field[string] `json:"name,required"`
	// Target gateway of the hostname.
	Target param.Field[TargetGateway] `json:"target,required"`
	// An optional description of the hostname.
	Description param.Field[string] `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink param.Field[string] `json:"dnslink"`
}

func (r ZoneWeb3HostnameNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneWeb3HostnameDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneWeb3HostnameDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneWeb3HostnamePatchParams struct {
	// An optional description of the hostname.
	Description param.Field[string] `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink param.Field[string] `json:"dnslink"`
}

func (r ZoneWeb3HostnamePatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
