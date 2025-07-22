// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountPcapService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountPcapService] method instead.
type AccountPcapService struct {
	Options   []option.RequestOption
	Ownership *AccountPcapOwnershipService
}

// NewAccountPcapService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountPcapService(opts ...option.RequestOption) (r *AccountPcapService) {
	r = &AccountPcapService{}
	r.Options = opts
	r.Ownership = NewAccountPcapOwnershipService(opts...)
	return
}

// Create new PCAP request for account.
func (r *AccountPcapService) New(ctx context.Context, accountID string, body AccountPcapNewParams, opts ...option.RequestOption) (res *SingleResponsePcaps, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get information for a PCAP request by id.
func (r *AccountPcapService) Get(ctx context.Context, accountID string, pcapID string, opts ...option.RequestOption) (res *SingleResponsePcaps, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if pcapID == "" {
		err = errors.New("missing required pcap_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps/%s", accountID, pcapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all packet capture requests for an account.
func (r *AccountPcapService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountPcapListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Download PCAP information into a file. Response is a binary PCAP file.
func (r *AccountPcapService) Download(ctx context.Context, accountID string, pcapID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.tcpdump.pcap")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if pcapID == "" {
		err = errors.New("missing required pcap_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps/%s/download", accountID, pcapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type APIResponseCollectionMagicVisibilityPcaps struct {
	Result     []interface{}                                       `json:"result,nullable"`
	ResultInfo APIResponseCollectionMagicVisibilityPcapsResultInfo `json:"result_info"`
	JSON       apiResponseCollectionMagicVisibilityPcapsJSON       `json:"-"`
	APIResponseMagicVisibilityPcaps
}

// apiResponseCollectionMagicVisibilityPcapsJSON contains the JSON metadata for the
// struct [APIResponseCollectionMagicVisibilityPcaps]
type apiResponseCollectionMagicVisibilityPcapsJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionMagicVisibilityPcaps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionMagicVisibilityPcapsJSON) RawJSON() string {
	return r.raw
}

type APIResponseCollectionMagicVisibilityPcapsResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                 `json:"total_count"`
	JSON       apiResponseCollectionMagicVisibilityPcapsResultInfoJSON `json:"-"`
}

// apiResponseCollectionMagicVisibilityPcapsResultInfoJSON contains the JSON
// metadata for the struct [APIResponseCollectionMagicVisibilityPcapsResultInfo]
type apiResponseCollectionMagicVisibilityPcapsResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionMagicVisibilityPcapsResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionMagicVisibilityPcapsResultInfoJSON) RawJSON() string {
	return r.raw
}

// The packet capture filter. When this field is empty, all packets are captured.
type FilterV1 struct {
	// The destination IP address of the packet.
	DestinationAddress string `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort float64 `json:"destination_port"`
	// The protocol number of the packet.
	Protocol float64 `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress string `json:"source_address"`
	// The source port of the packet.
	SourcePort float64      `json:"source_port"`
	JSON       filterV1JSON `json:"-"`
}

// filterV1JSON contains the JSON metadata for the struct [FilterV1]
type filterV1JSON struct {
	DestinationAddress apijson.Field
	DestinationPort    apijson.Field
	Protocol           apijson.Field
	SourceAddress      apijson.Field
	SourcePort         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *FilterV1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterV1JSON) RawJSON() string {
	return r.raw
}

// The packet capture filter. When this field is empty, all packets are captured.
type FilterV1Param struct {
	// The destination IP address of the packet.
	DestinationAddress param.Field[string] `json:"destination_address"`
	// The destination port of the packet.
	DestinationPort param.Field[float64] `json:"destination_port"`
	// The protocol number of the packet.
	Protocol param.Field[float64] `json:"protocol"`
	// The source IP address of the packet.
	SourceAddress param.Field[string] `json:"source_address"`
	// The source port of the packet.
	SourcePort param.Field[float64] `json:"source_port"`
}

func (r FilterV1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of packet capture. `Simple` captures sampled packets, and `full`
// captures entire payloads and non-sampled packets.
type PacketCaptureType string

const (
	PacketCaptureTypeSimple PacketCaptureType = "simple"
	PacketCaptureTypeFull   PacketCaptureType = "full"
)

func (r PacketCaptureType) IsKnown() bool {
	switch r {
	case PacketCaptureTypeSimple, PacketCaptureTypeFull:
		return true
	}
	return false
}

type ResponseFull struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit float64 `json:"byte_limit"`
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName string `json:"colo_name"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf"`
	// An error message that describes why the packet capture failed. This field only
	// applies to `full` packet captures.
	ErrorMessage string `json:"error_message"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 FilterV1 `json:"filter_v1"`
	// The status of the packet capture request.
	Status StatusPacketCapture `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System System `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PacketCaptureType `json:"type"`
	JSON responseFullJSON  `json:"-"`
}

// responseFullJSON contains the JSON metadata for the struct [ResponseFull]
type responseFullJSON struct {
	ID              apijson.Field
	ByteLimit       apijson.Field
	ColoName        apijson.Field
	DestinationConf apijson.Field
	ErrorMessage    apijson.Field
	FilterV1        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	System          apijson.Field
	TimeLimit       apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResponseFull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseFullJSON) RawJSON() string {
	return r.raw
}

func (r ResponseFull) implementsSingleResponsePcapsResult() {}

func (r ResponseFull) implementsAccountPcapListResponseResult() {}

type ResponseSimple struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 FilterV1 `json:"filter_v1"`
	// The RFC 3339 offset timestamp from which to query backwards for packets. Must be
	// within the last 24h. When this field is empty, defaults to time of request.
	OffsetTime time.Time `json:"offset_time" format:"date-time"`
	// The status of the packet capture request.
	Status StatusPacketCapture `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System System `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type PacketCaptureType  `json:"type"`
	JSON responseSimpleJSON `json:"-"`
}

// responseSimpleJSON contains the JSON metadata for the struct [ResponseSimple]
type responseSimpleJSON struct {
	ID          apijson.Field
	FilterV1    apijson.Field
	OffsetTime  apijson.Field
	Status      apijson.Field
	Submitted   apijson.Field
	System      apijson.Field
	TimeLimit   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseSimpleJSON) RawJSON() string {
	return r.raw
}

func (r ResponseSimple) implementsSingleResponsePcapsResult() {}

func (r ResponseSimple) implementsAccountPcapListResponseResult() {}

type SingleResponsePcaps struct {
	Result SingleResponsePcapsResult `json:"result"`
	JSON   singleResponsePcapsJSON   `json:"-"`
	APIResponseMagicVisibilityPcaps
}

// singleResponsePcapsJSON contains the JSON metadata for the struct
// [SingleResponsePcaps]
type singleResponsePcapsJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponsePcaps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponsePcapsJSON) RawJSON() string {
	return r.raw
}

type SingleResponsePcapsResult struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit float64 `json:"byte_limit"`
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName string `json:"colo_name"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf"`
	// An error message that describes why the packet capture failed. This field only
	// applies to `full` packet captures.
	ErrorMessage string `json:"error_message"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 FilterV1 `json:"filter_v1"`
	// The RFC 3339 offset timestamp from which to query backwards for packets. Must be
	// within the last 24h. When this field is empty, defaults to time of request.
	OffsetTime time.Time `json:"offset_time" format:"date-time"`
	// The status of the packet capture request.
	Status StatusPacketCapture `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System System `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type  PacketCaptureType             `json:"type"`
	JSON  singleResponsePcapsResultJSON `json:"-"`
	union SingleResponsePcapsResultUnion
}

// singleResponsePcapsResultJSON contains the JSON metadata for the struct
// [SingleResponsePcapsResult]
type singleResponsePcapsResultJSON struct {
	ID              apijson.Field
	ByteLimit       apijson.Field
	ColoName        apijson.Field
	DestinationConf apijson.Field
	ErrorMessage    apijson.Field
	FilterV1        apijson.Field
	OffsetTime      apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	System          apijson.Field
	TimeLimit       apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r singleResponsePcapsResultJSON) RawJSON() string {
	return r.raw
}

func (r *SingleResponsePcapsResult) UnmarshalJSON(data []byte) (err error) {
	*r = SingleResponsePcapsResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SingleResponsePcapsResultUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [ResponseSimple], [ResponseFull].
func (r SingleResponsePcapsResult) AsUnion() SingleResponsePcapsResultUnion {
	return r.union
}

// Union satisfied by [ResponseSimple] or [ResponseFull].
type SingleResponsePcapsResultUnion interface {
	implementsSingleResponsePcapsResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SingleResponsePcapsResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResponseSimple{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResponseFull{}),
		},
	)
}

// The status of the packet capture request.
type StatusPacketCapture string

const (
	StatusPacketCaptureUnknown           StatusPacketCapture = "unknown"
	StatusPacketCaptureSuccess           StatusPacketCapture = "success"
	StatusPacketCapturePending           StatusPacketCapture = "pending"
	StatusPacketCaptureRunning           StatusPacketCapture = "running"
	StatusPacketCaptureConversionPending StatusPacketCapture = "conversion_pending"
	StatusPacketCaptureConversionRunning StatusPacketCapture = "conversion_running"
	StatusPacketCaptureComplete          StatusPacketCapture = "complete"
	StatusPacketCaptureFailed            StatusPacketCapture = "failed"
)

func (r StatusPacketCapture) IsKnown() bool {
	switch r {
	case StatusPacketCaptureUnknown, StatusPacketCaptureSuccess, StatusPacketCapturePending, StatusPacketCaptureRunning, StatusPacketCaptureConversionPending, StatusPacketCaptureConversionRunning, StatusPacketCaptureComplete, StatusPacketCaptureFailed:
		return true
	}
	return false
}

// The system used to collect packet captures.
type System string

const (
	SystemMagicTransit System = "magic-transit"
)

func (r System) IsKnown() bool {
	switch r {
	case SystemMagicTransit:
		return true
	}
	return false
}

type AccountPcapListResponse struct {
	Result []AccountPcapListResponseResult `json:"result"`
	JSON   accountPcapListResponseJSON     `json:"-"`
	APIResponseCollectionMagicVisibilityPcaps
}

// accountPcapListResponseJSON contains the JSON metadata for the struct
// [AccountPcapListResponse]
type accountPcapListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPcapListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountPcapListResponseResult struct {
	// The ID for the packet capture.
	ID string `json:"id"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit float64 `json:"byte_limit"`
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName string `json:"colo_name"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf"`
	// An error message that describes why the packet capture failed. This field only
	// applies to `full` packet captures.
	ErrorMessage string `json:"error_message"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 FilterV1 `json:"filter_v1"`
	// The RFC 3339 offset timestamp from which to query backwards for packets. Must be
	// within the last 24h. When this field is empty, defaults to time of request.
	OffsetTime time.Time `json:"offset_time" format:"date-time"`
	// The status of the packet capture request.
	Status StatusPacketCapture `json:"status"`
	// The RFC 3339 timestamp when the packet capture was created.
	Submitted string `json:"submitted"`
	// The system used to collect packet captures.
	System System `json:"system"`
	// The packet capture duration in seconds.
	TimeLimit float64 `json:"time_limit"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type  PacketCaptureType                 `json:"type"`
	JSON  accountPcapListResponseResultJSON `json:"-"`
	union AccountPcapListResponseResultUnion
}

// accountPcapListResponseResultJSON contains the JSON metadata for the struct
// [AccountPcapListResponseResult]
type accountPcapListResponseResultJSON struct {
	ID              apijson.Field
	ByteLimit       apijson.Field
	ColoName        apijson.Field
	DestinationConf apijson.Field
	ErrorMessage    apijson.Field
	FilterV1        apijson.Field
	OffsetTime      apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	System          apijson.Field
	TimeLimit       apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r accountPcapListResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *AccountPcapListResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = AccountPcapListResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountPcapListResponseResultUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [ResponseSimple], [ResponseFull].
func (r AccountPcapListResponseResult) AsUnion() AccountPcapListResponseResultUnion {
	return r.union
}

// Union satisfied by [ResponseSimple] or [ResponseFull].
type AccountPcapListResponseResultUnion interface {
	implementsAccountPcapListResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountPcapListResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResponseSimple{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResponseFull{}),
		},
	)
}

type AccountPcapNewParams struct {
	Body AccountPcapNewParamsBodyUnion `json:"body,required"`
}

func (r AccountPcapNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountPcapNewParamsBody struct {
	// The system used to collect packet captures.
	System param.Field[System] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PacketCaptureType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[FilterV1Param] `json:"filter_v1"`
	// The RFC 3339 offset timestamp from which to query backwards for packets. Must be
	// within the last 24h. When this field is empty, defaults to time of request.
	OffsetTime param.Field[time.Time] `json:"offset_time" format:"date-time"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r AccountPcapNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountPcapNewParamsBody) implementsAccountPcapNewParamsBodyUnion() {}

// Satisfied by [AccountPcapNewParamsBodyMagicVisibilityPcapsPcapsRequestSimple],
// [AccountPcapNewParamsBodyMagicVisibilityPcapsPcapsRequestFull],
// [AccountPcapNewParamsBody].
type AccountPcapNewParamsBodyUnion interface {
	implementsAccountPcapNewParamsBodyUnion()
}

type AccountPcapNewParamsBodyMagicVisibilityPcapsPcapsRequestSimple struct {
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit,required"`
	// The system used to collect packet captures.
	System param.Field[System] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PacketCaptureType] `json:"type,required"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[FilterV1Param] `json:"filter_v1"`
	// The RFC 3339 offset timestamp from which to query backwards for packets. Must be
	// within the last 24h. When this field is empty, defaults to time of request.
	OffsetTime param.Field[time.Time] `json:"offset_time" format:"date-time"`
}

func (r AccountPcapNewParamsBodyMagicVisibilityPcapsPcapsRequestSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountPcapNewParamsBodyMagicVisibilityPcapsPcapsRequestSimple) implementsAccountPcapNewParamsBodyUnion() {
}

type AccountPcapNewParamsBodyMagicVisibilityPcapsPcapsRequestFull struct {
	// The name of the data center used for the packet capture. This can be a specific
	// colo (ord02) or a multi-colo name (ORD). This field only applies to `full`
	// packet captures.
	ColoName param.Field[string] `json:"colo_name,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The system used to collect packet captures.
	System param.Field[System] `json:"system,required"`
	// The packet capture duration in seconds.
	TimeLimit param.Field[float64] `json:"time_limit,required"`
	// The type of packet capture. `Simple` captures sampled packets, and `full`
	// captures entire payloads and non-sampled packets.
	Type param.Field[PacketCaptureType] `json:"type,required"`
	// The maximum number of bytes to capture. This field only applies to `full` packet
	// captures.
	ByteLimit param.Field[float64] `json:"byte_limit"`
	// The packet capture filter. When this field is empty, all packets are captured.
	FilterV1 param.Field[FilterV1Param] `json:"filter_v1"`
	// The limit of packets contained in a packet capture.
	PacketLimit param.Field[float64] `json:"packet_limit"`
}

func (r AccountPcapNewParamsBodyMagicVisibilityPcapsPcapsRequestFull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountPcapNewParamsBodyMagicVisibilityPcapsPcapsRequestFull) implementsAccountPcapNewParamsBodyUnion() {
}
