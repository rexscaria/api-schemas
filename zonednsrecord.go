// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apiform"
	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/tidwall/gjson"
)

// ZoneDNSRecordService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneDNSRecordService] method instead.
type ZoneDNSRecordService struct {
	Options []option.RequestOption
}

// NewZoneDNSRecordService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneDNSRecordService(opts ...option.RequestOption) (r *ZoneDNSRecordService) {
	r = &ZoneDNSRecordService{}
	r.Options = opts
	return
}

// Create a new DNS record for a zone.
//
// Notes:
//
//   - A/AAAA records cannot exist on the same name as CNAME records.
//   - NS records cannot exist on the same name as any other record type.
//   - Domain names are always represented in Punycode, even if Unicode characters
//     were used when creating the record.
func (r *ZoneDNSRecordService) New(ctx context.Context, zoneID string, body ZoneDNSRecordNewParams, opts ...option.RequestOption) (res *SingleResponseDNSResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_records", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// DNS Record Details
func (r *ZoneDNSRecordService) Get(ctx context.Context, zoneID string, dnsRecordID string, opts ...option.RequestOption) (res *SingleResponseDNSResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if dnsRecordID == "" {
		err = errors.New("missing required dns_record_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_records/%s", zoneID, dnsRecordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing DNS record.
//
// Notes:
//
//   - A/AAAA records cannot exist on the same name as CNAME records.
//   - NS records cannot exist on the same name as any other record type.
//   - Domain names are always represented in Punycode, even if Unicode characters
//     were used when creating the record.
func (r *ZoneDNSRecordService) Update(ctx context.Context, zoneID string, dnsRecordID string, body ZoneDNSRecordUpdateParams, opts ...option.RequestOption) (res *SingleResponseDNSResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if dnsRecordID == "" {
		err = errors.New("missing required dns_record_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_records/%s", zoneID, dnsRecordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List, search, sort, and filter a zones' DNS records.
func (r *ZoneDNSRecordService) List(ctx context.Context, zoneID string, query ZoneDNSRecordListParams, opts ...option.RequestOption) (res *ZoneDNSRecordListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_records", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete DNS Record
func (r *ZoneDNSRecordService) Delete(ctx context.Context, zoneID string, dnsRecordID string, body ZoneDNSRecordDeleteParams, opts ...option.RequestOption) (res *ZoneDNSRecordDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if dnsRecordID == "" {
		err = errors.New("missing required dns_record_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_records/%s", zoneID, dnsRecordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Send a Batch of DNS Record API calls to be executed together.
//
// Notes:
//
//   - Although Cloudflare will execute the batched operations in a single database
//     transaction, Cloudflare's distributed KV store must treat each record change
//     as a single key-value pair. This means that the propagation of changes is not
//     atomic. See
//     [the documentation](https://developers.cloudflare.com/dns/manage-dns-records/how-to/batch-record-changes/ "Batch DNS records")
//     for more information.
//
//   - The operations you specify within the /batch request body are always executed
//     in the following order:
//
//   - Deletes
//
//   - Patches
//
//   - Puts
//
//   - Posts
func (r *ZoneDNSRecordService) Batch(ctx context.Context, zoneID string, body ZoneDNSRecordBatchParams, opts ...option.RequestOption) (res *ZoneDNSRecordBatchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_records/batch", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// You can export your
// [BIND config](https://en.wikipedia.org/wiki/Zone_file "Zone file") through this
// endpoint.
//
// See
// [the documentation](https://developers.cloudflare.com/dns/manage-dns-records/how-to/import-and-export/ "Import and export records")
// for more information.
func (r *ZoneDNSRecordService) Export(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_records/export", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// You can upload your
// [BIND config](https://en.wikipedia.org/wiki/Zone_file "Zone file") through this
// endpoint. It assumes that cURL is called from a location with bind_config.txt
// (valid BIND config) present.
//
// See
// [the documentation](https://developers.cloudflare.com/dns/manage-dns-records/how-to/import-and-export/ "Import and export records")
// for more information.
func (r *ZoneDNSRecordService) Import(ctx context.Context, zoneID string, body ZoneDNSRecordImportParams, opts ...option.RequestOption) (res *ImportScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_records/import", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Overwrite an existing DNS record.
//
// Notes:
//
//   - A/AAAA records cannot exist on the same name as CNAME records.
//   - NS records cannot exist on the same name as any other record type.
//   - Domain names are always represented in Punycode, even if Unicode characters
//     were used when creating the record.
func (r *ZoneDNSRecordService) Overwrite(ctx context.Context, zoneID string, dnsRecordID string, body ZoneDNSRecordOverwriteParams, opts ...option.RequestOption) (res *SingleResponseDNSResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if dnsRecordID == "" {
		err = errors.New("missing required dns_record_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_records/%s", zoneID, dnsRecordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Scan for common DNS records on your domain and automatically add them to your
// zone. Useful if you haven't updated your nameservers yet.
func (r *ZoneDNSRecordService) Scan(ctx context.Context, zoneID string, body ZoneDNSRecordScanParams, opts ...option.RequestOption) (res *ImportScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_records/scan", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CommonResponseDNSRecords struct {
	Errors   []DNSRecordMessageItem `json:"errors,required"`
	Messages []DNSRecordMessageItem `json:"messages,required"`
	// Whether the API call was successful
	Success CommonResponseDNSRecordsSuccess `json:"success,required"`
	JSON    commonResponseDNSRecordsJSON    `json:"-"`
}

// commonResponseDNSRecordsJSON contains the JSON metadata for the struct
// [CommonResponseDNSRecords]
type commonResponseDNSRecordsJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommonResponseDNSRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commonResponseDNSRecordsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CommonResponseDNSRecordsSuccess bool

const (
	CommonResponseDNSRecordsSuccessTrue CommonResponseDNSRecordsSuccess = true
)

func (r CommonResponseDNSRecordsSuccess) IsKnown() bool {
	switch r {
	case CommonResponseDNSRecordsSuccessTrue:
		return true
	}
	return false
}

type DNSRecordMessageItem struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    dnsRecordMessageItemJSON `json:"-"`
}

// dnsRecordMessageItemJSON contains the JSON metadata for the struct
// [DNSRecordMessageItem]
type dnsRecordMessageItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordMessageItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordMessageItemJSON) RawJSON() string {
	return r.raw
}

// Satisfied by [DNSRecordPatchDNSRecordsARecordParam],
// [DNSRecordPatchDNSRecordsAaaaRecordParam],
// [DNSRecordPatchDNSRecordsCaaRecordParam],
// [DNSRecordPatchDNSRecordsCertRecordParam],
// [DNSRecordPatchDNSRecordsCnameRecordParam],
// [DNSRecordPatchDNSRecordsDnskeyRecordParam],
// [DNSRecordPatchDNSRecordsDsRecordParam],
// [DNSRecordPatchDNSRecordsHTTPSRecordParam],
// [DNSRecordPatchDNSRecordsLocRecordParam],
// [DNSRecordPatchDNSRecordsMxRecordParam],
// [DNSRecordPatchDNSRecordsNaptrRecordParam],
// [DNSRecordPatchDNSRecordsNsRecordParam],
// [DNSRecordPatchDNSRecordsOpenpgpkeyRecordParam],
// [DNSRecordPatchDNSRecordsPtrRecordParam],
// [DNSRecordPatchDNSRecordsSmimeaRecordParam],
// [DNSRecordPatchDNSRecordsSrvRecordParam],
// [DNSRecordPatchDNSRecordsSshfpRecordParam],
// [DNSRecordPatchDNSRecordsSvcbRecordParam],
// [DNSRecordPatchDNSRecordsTlsaRecordParam],
// [DNSRecordPatchDNSRecordsTxtRecordParam],
// [DNSRecordPatchDNSRecordsUriRecordParam].
type DNSRecordPatchUnionParam interface {
	implementsDNSRecordPatchUnionParam()
}

type DNSRecordPatchDNSRecordsARecordParam struct {
	// A valid IPv4 address.
	Content param.Field[string] `json:"content" format:"ipv4"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsARecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsARecordParam) implementsDNSRecordPatchUnionParam() {}

// Record type.
type DNSRecordPatchDNSRecordsARecordType string

const (
	DNSRecordPatchDNSRecordsARecordTypeA DNSRecordPatchDNSRecordsARecordType = "A"
)

func (r DNSRecordPatchDNSRecordsARecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsARecordTypeA:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsAaaaRecordParam struct {
	// A valid IPv6 address.
	Content param.Field[string] `json:"content" format:"ipv6"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsAaaaRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsAaaaRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsAaaaRecordParam) implementsDNSRecordPatchUnionParam() {}

// Record type.
type DNSRecordPatchDNSRecordsAaaaRecordType string

const (
	DNSRecordPatchDNSRecordsAaaaRecordTypeAaaa DNSRecordPatchDNSRecordsAaaaRecordType = "AAAA"
)

func (r DNSRecordPatchDNSRecordsAaaaRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsAaaaRecordTypeAaaa:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsCaaRecordParam struct {
	// Components of a CAA record.
	Data param.Field[DNSRecordPatchDNSRecordsCaaRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsCaaRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsCaaRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsCaaRecordParam) implementsDNSRecordPatchUnionParam() {}

// Components of a CAA record.
type DNSRecordPatchDNSRecordsCaaRecordDataParam struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r DNSRecordPatchDNSRecordsCaaRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPatchDNSRecordsCaaRecordType string

const (
	DNSRecordPatchDNSRecordsCaaRecordTypeCaa DNSRecordPatchDNSRecordsCaaRecordType = "CAA"
)

func (r DNSRecordPatchDNSRecordsCaaRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsCaaRecordTypeCaa:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsCertRecordParam struct {
	// Components of a CERT record.
	Data param.Field[DNSRecordPatchDNSRecordsCertRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsCertRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsCertRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsCertRecordParam) implementsDNSRecordPatchUnionParam() {}

// Components of a CERT record.
type DNSRecordPatchDNSRecordsCertRecordDataParam struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r DNSRecordPatchDNSRecordsCertRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPatchDNSRecordsCertRecordType string

const (
	DNSRecordPatchDNSRecordsCertRecordTypeCert DNSRecordPatchDNSRecordsCertRecordType = "CERT"
)

func (r DNSRecordPatchDNSRecordsCertRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsCertRecordTypeCert:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsCnameRecordParam struct {
	// A valid hostname. Must not match the record's name.
	Content  param.Field[string]                                           `json:"content"`
	Settings param.Field[DNSRecordPatchDNSRecordsCnameRecordSettingsParam] `json:"settings"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsCnameRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsCnameRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsCnameRecordParam) implementsDNSRecordPatchUnionParam() {}

type DNSRecordPatchDNSRecordsCnameRecordSettingsParam struct {
	// If enabled, causes the CNAME record to be resolved externally and the resulting
	// address records (e.g., A and AAAA) to be returned instead of the CNAME record
	// itself. This setting is unavailable for proxied records, since they are always
	// flattened.
	FlattenCname param.Field[bool] `json:"flatten_cname"`
}

func (r DNSRecordPatchDNSRecordsCnameRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPatchDNSRecordsCnameRecordType string

const (
	DNSRecordPatchDNSRecordsCnameRecordTypeCname DNSRecordPatchDNSRecordsCnameRecordType = "CNAME"
)

func (r DNSRecordPatchDNSRecordsCnameRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsCnameRecordTypeCname:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsDnskeyRecordParam struct {
	// Components of a DNSKEY record.
	Data param.Field[DNSRecordPatchDNSRecordsDnskeyRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsDnskeyRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsDnskeyRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsDnskeyRecordParam) implementsDNSRecordPatchUnionParam() {}

// Components of a DNSKEY record.
type DNSRecordPatchDNSRecordsDnskeyRecordDataParam struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r DNSRecordPatchDNSRecordsDnskeyRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPatchDNSRecordsDnskeyRecordType string

const (
	DNSRecordPatchDNSRecordsDnskeyRecordTypeDnskey DNSRecordPatchDNSRecordsDnskeyRecordType = "DNSKEY"
)

func (r DNSRecordPatchDNSRecordsDnskeyRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsDnskeyRecordTypeDnskey:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsDsRecordParam struct {
	// Components of a DS record.
	Data param.Field[DNSRecordPatchDNSRecordsDsRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsDsRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsDsRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsDsRecordParam) implementsDNSRecordPatchUnionParam() {}

// Components of a DS record.
type DNSRecordPatchDNSRecordsDsRecordDataParam struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r DNSRecordPatchDNSRecordsDsRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPatchDNSRecordsDsRecordType string

const (
	DNSRecordPatchDNSRecordsDsRecordTypeDs DNSRecordPatchDNSRecordsDsRecordType = "DS"
)

func (r DNSRecordPatchDNSRecordsDsRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsDsRecordTypeDs:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsHTTPSRecordParam struct {
	// Components of a HTTPS record.
	Data param.Field[DNSRecordPatchDNSRecordsHTTPSRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsHTTPSRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsHTTPSRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsHTTPSRecordParam) implementsDNSRecordPatchUnionParam() {}

// Components of a HTTPS record.
type DNSRecordPatchDNSRecordsHTTPSRecordDataParam struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r DNSRecordPatchDNSRecordsHTTPSRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPatchDNSRecordsHTTPSRecordType string

const (
	DNSRecordPatchDNSRecordsHTTPSRecordTypeHTTPS DNSRecordPatchDNSRecordsHTTPSRecordType = "HTTPS"
)

func (r DNSRecordPatchDNSRecordsHTTPSRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsHTTPSRecordTypeHTTPS:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsLocRecordParam struct {
	// Components of a LOC record.
	Data param.Field[DNSRecordPatchDNSRecordsLocRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsLocRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsLocRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsLocRecordParam) implementsDNSRecordPatchUnionParam() {}

// Components of a LOC record.
type DNSRecordPatchDNSRecordsLocRecordDataParam struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[DNSRecordPatchDNSRecordsLocRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[DNSRecordPatchDNSRecordsLocRecordDataLongDirection] `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes param.Field[float64] `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds param.Field[float64] `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz param.Field[float64] `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert param.Field[float64] `json:"precision_vert"`
	// Size of location in meters.
	Size param.Field[float64] `json:"size"`
}

func (r DNSRecordPatchDNSRecordsLocRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type DNSRecordPatchDNSRecordsLocRecordDataLatDirection string

const (
	DNSRecordPatchDNSRecordsLocRecordDataLatDirectionN DNSRecordPatchDNSRecordsLocRecordDataLatDirection = "N"
	DNSRecordPatchDNSRecordsLocRecordDataLatDirectionS DNSRecordPatchDNSRecordsLocRecordDataLatDirection = "S"
)

func (r DNSRecordPatchDNSRecordsLocRecordDataLatDirection) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsLocRecordDataLatDirectionN, DNSRecordPatchDNSRecordsLocRecordDataLatDirectionS:
		return true
	}
	return false
}

// Longitude direction.
type DNSRecordPatchDNSRecordsLocRecordDataLongDirection string

const (
	DNSRecordPatchDNSRecordsLocRecordDataLongDirectionE DNSRecordPatchDNSRecordsLocRecordDataLongDirection = "E"
	DNSRecordPatchDNSRecordsLocRecordDataLongDirectionW DNSRecordPatchDNSRecordsLocRecordDataLongDirection = "W"
)

func (r DNSRecordPatchDNSRecordsLocRecordDataLongDirection) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsLocRecordDataLongDirectionE, DNSRecordPatchDNSRecordsLocRecordDataLongDirectionW:
		return true
	}
	return false
}

// Record type.
type DNSRecordPatchDNSRecordsLocRecordType string

const (
	DNSRecordPatchDNSRecordsLocRecordTypeLoc DNSRecordPatchDNSRecordsLocRecordType = "LOC"
)

func (r DNSRecordPatchDNSRecordsLocRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsLocRecordTypeLoc:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsMxRecordParam struct {
	// A valid mail server hostname.
	Content param.Field[string] `json:"content" format:"hostname"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsMxRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsMxRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsMxRecordParam) implementsDNSRecordPatchUnionParam() {}

// Record type.
type DNSRecordPatchDNSRecordsMxRecordType string

const (
	DNSRecordPatchDNSRecordsMxRecordTypeMx DNSRecordPatchDNSRecordsMxRecordType = "MX"
)

func (r DNSRecordPatchDNSRecordsMxRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsMxRecordTypeMx:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsNaptrRecordParam struct {
	// Components of a NAPTR record.
	Data param.Field[DNSRecordPatchDNSRecordsNaptrRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsNaptrRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsNaptrRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsNaptrRecordParam) implementsDNSRecordPatchUnionParam() {}

// Components of a NAPTR record.
type DNSRecordPatchDNSRecordsNaptrRecordDataParam struct {
	// Flags.
	Flags param.Field[string] `json:"flags"`
	// Order.
	Order param.Field[float64] `json:"order"`
	// Preference.
	Preference param.Field[float64] `json:"preference"`
	// Regex.
	Regex param.Field[string] `json:"regex"`
	// Replacement.
	Replacement param.Field[string] `json:"replacement"`
	// Service.
	Service param.Field[string] `json:"service"`
}

func (r DNSRecordPatchDNSRecordsNaptrRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPatchDNSRecordsNaptrRecordType string

const (
	DNSRecordPatchDNSRecordsNaptrRecordTypeNaptr DNSRecordPatchDNSRecordsNaptrRecordType = "NAPTR"
)

func (r DNSRecordPatchDNSRecordsNaptrRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsNaptrRecordTypeNaptr:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsNsRecordParam struct {
	// A valid name server host name.
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsNsRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsNsRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsNsRecordParam) implementsDNSRecordPatchUnionParam() {}

// Record type.
type DNSRecordPatchDNSRecordsNsRecordType string

const (
	DNSRecordPatchDNSRecordsNsRecordTypeNs DNSRecordPatchDNSRecordsNsRecordType = "NS"
)

func (r DNSRecordPatchDNSRecordsNsRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsNsRecordTypeNs:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsOpenpgpkeyRecordParam struct {
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsOpenpgpkeyRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsOpenpgpkeyRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsOpenpgpkeyRecordParam) implementsDNSRecordPatchUnionParam() {}

// Record type.
type DNSRecordPatchDNSRecordsOpenpgpkeyRecordType string

const (
	DNSRecordPatchDNSRecordsOpenpgpkeyRecordTypeOpenpgpkey DNSRecordPatchDNSRecordsOpenpgpkeyRecordType = "OPENPGPKEY"
)

func (r DNSRecordPatchDNSRecordsOpenpgpkeyRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsOpenpgpkeyRecordTypeOpenpgpkey:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsPtrRecordParam struct {
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsPtrRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsPtrRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsPtrRecordParam) implementsDNSRecordPatchUnionParam() {}

// Record type.
type DNSRecordPatchDNSRecordsPtrRecordType string

const (
	DNSRecordPatchDNSRecordsPtrRecordTypePtr DNSRecordPatchDNSRecordsPtrRecordType = "PTR"
)

func (r DNSRecordPatchDNSRecordsPtrRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsPtrRecordTypePtr:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsSmimeaRecordParam struct {
	// Components of a SMIMEA record.
	Data param.Field[DNSRecordPatchDNSRecordsSmimeaRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsSmimeaRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsSmimeaRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsSmimeaRecordParam) implementsDNSRecordPatchUnionParam() {}

// Components of a SMIMEA record.
type DNSRecordPatchDNSRecordsSmimeaRecordDataParam struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r DNSRecordPatchDNSRecordsSmimeaRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPatchDNSRecordsSmimeaRecordType string

const (
	DNSRecordPatchDNSRecordsSmimeaRecordTypeSmimea DNSRecordPatchDNSRecordsSmimeaRecordType = "SMIMEA"
)

func (r DNSRecordPatchDNSRecordsSmimeaRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsSmimeaRecordTypeSmimea:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsSrvRecordParam struct {
	// Components of a SRV record.
	Data param.Field[DNSRecordPatchDNSRecordsSrvRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsSrvRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsSrvRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsSrvRecordParam) implementsDNSRecordPatchUnionParam() {}

// Components of a SRV record.
type DNSRecordPatchDNSRecordsSrvRecordDataParam struct {
	// The port of the service.
	Port param.Field[float64] `json:"port"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// A valid hostname.
	Target param.Field[string] `json:"target" format:"hostname"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r DNSRecordPatchDNSRecordsSrvRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPatchDNSRecordsSrvRecordType string

const (
	DNSRecordPatchDNSRecordsSrvRecordTypeSrv DNSRecordPatchDNSRecordsSrvRecordType = "SRV"
)

func (r DNSRecordPatchDNSRecordsSrvRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsSrvRecordTypeSrv:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsSshfpRecordParam struct {
	// Components of a SSHFP record.
	Data param.Field[DNSRecordPatchDNSRecordsSshfpRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsSshfpRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsSshfpRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsSshfpRecordParam) implementsDNSRecordPatchUnionParam() {}

// Components of a SSHFP record.
type DNSRecordPatchDNSRecordsSshfpRecordDataParam struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r DNSRecordPatchDNSRecordsSshfpRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPatchDNSRecordsSshfpRecordType string

const (
	DNSRecordPatchDNSRecordsSshfpRecordTypeSshfp DNSRecordPatchDNSRecordsSshfpRecordType = "SSHFP"
)

func (r DNSRecordPatchDNSRecordsSshfpRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsSshfpRecordTypeSshfp:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsSvcbRecordParam struct {
	// Components of a SVCB record.
	Data param.Field[DNSRecordPatchDNSRecordsSvcbRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsSvcbRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsSvcbRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsSvcbRecordParam) implementsDNSRecordPatchUnionParam() {}

// Components of a SVCB record.
type DNSRecordPatchDNSRecordsSvcbRecordDataParam struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r DNSRecordPatchDNSRecordsSvcbRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPatchDNSRecordsSvcbRecordType string

const (
	DNSRecordPatchDNSRecordsSvcbRecordTypeSvcb DNSRecordPatchDNSRecordsSvcbRecordType = "SVCB"
)

func (r DNSRecordPatchDNSRecordsSvcbRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsSvcbRecordTypeSvcb:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsTlsaRecordParam struct {
	// Components of a TLSA record.
	Data param.Field[DNSRecordPatchDNSRecordsTlsaRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsTlsaRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsTlsaRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsTlsaRecordParam) implementsDNSRecordPatchUnionParam() {}

// Components of a TLSA record.
type DNSRecordPatchDNSRecordsTlsaRecordDataParam struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r DNSRecordPatchDNSRecordsTlsaRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPatchDNSRecordsTlsaRecordType string

const (
	DNSRecordPatchDNSRecordsTlsaRecordTypeTlsa DNSRecordPatchDNSRecordsTlsaRecordType = "TLSA"
)

func (r DNSRecordPatchDNSRecordsTlsaRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsTlsaRecordTypeTlsa:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsTxtRecordParam struct {
	// Text content for the record. The content must consist of quoted "character
	// strings" (RFC 1035), each with a length of up to 255 bytes. Strings exceeding
	// this allowed maximum length are automatically split.
	//
	// Learn more at
	// <https://www.cloudflare.com/learning/dns/dns-records/dns-txt-record/>.
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsTxtRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsTxtRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsTxtRecordParam) implementsDNSRecordPatchUnionParam() {}

// Record type.
type DNSRecordPatchDNSRecordsTxtRecordType string

const (
	DNSRecordPatchDNSRecordsTxtRecordTypeTxt DNSRecordPatchDNSRecordsTxtRecordType = "TXT"
)

func (r DNSRecordPatchDNSRecordsTxtRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsTxtRecordTypeTxt:
		return true
	}
	return false
}

type DNSRecordPatchDNSRecordsUriRecordParam struct {
	// Components of a URI record.
	Data param.Field[DNSRecordPatchDNSRecordsUriRecordDataParam] `json:"data"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsUriRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPatchDNSRecordsUriRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsUriRecordParam) implementsDNSRecordPatchUnionParam() {}

// Components of a URI record.
type DNSRecordPatchDNSRecordsUriRecordDataParam struct {
	// The record content.
	Target param.Field[string] `json:"target"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r DNSRecordPatchDNSRecordsUriRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPatchDNSRecordsUriRecordType string

const (
	DNSRecordPatchDNSRecordsUriRecordTypeUri DNSRecordPatchDNSRecordsUriRecordType = "URI"
)

func (r DNSRecordPatchDNSRecordsUriRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsUriRecordTypeUri:
		return true
	}
	return false
}

// Satisfied by [DNSRecordPostDNSRecordsARecordParam],
// [DNSRecordPostDNSRecordsAaaaRecordParam],
// [DNSRecordPostDNSRecordsCaaRecordParam],
// [DNSRecordPostDNSRecordsCertRecordParam],
// [DNSRecordPostDNSRecordsCnameRecordParam],
// [DNSRecordPostDNSRecordsDnskeyRecordParam],
// [DNSRecordPostDNSRecordsDsRecordParam],
// [DNSRecordPostDNSRecordsHTTPSRecordParam],
// [DNSRecordPostDNSRecordsLocRecordParam], [DNSRecordPostDNSRecordsMxRecordParam],
// [DNSRecordPostDNSRecordsNaptrRecordParam],
// [DNSRecordPostDNSRecordsNsRecordParam],
// [DNSRecordPostDNSRecordsOpenpgpkeyRecordParam],
// [DNSRecordPostDNSRecordsPtrRecordParam],
// [DNSRecordPostDNSRecordsSmimeaRecordParam],
// [DNSRecordPostDNSRecordsSrvRecordParam],
// [DNSRecordPostDNSRecordsSshfpRecordParam],
// [DNSRecordPostDNSRecordsSvcbRecordParam],
// [DNSRecordPostDNSRecordsTlsaRecordParam],
// [DNSRecordPostDNSRecordsTxtRecordParam],
// [DNSRecordPostDNSRecordsUriRecordParam].
type DNSRecordPostUnionParam interface {
	implementsDNSRecordPostUnionParam()
}

type DNSRecordPostDNSRecordsARecordParam struct {
	// A valid IPv4 address.
	Content param.Field[string] `json:"content" format:"ipv4"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsARecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsARecordParam) implementsDNSRecordPostUnionParam() {}

// Record type.
type DNSRecordPostDNSRecordsARecordType string

const (
	DNSRecordPostDNSRecordsARecordTypeA DNSRecordPostDNSRecordsARecordType = "A"
)

func (r DNSRecordPostDNSRecordsARecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsARecordTypeA:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsAaaaRecordParam struct {
	// A valid IPv6 address.
	Content param.Field[string] `json:"content" format:"ipv6"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsAaaaRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsAaaaRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsAaaaRecordParam) implementsDNSRecordPostUnionParam() {}

// Record type.
type DNSRecordPostDNSRecordsAaaaRecordType string

const (
	DNSRecordPostDNSRecordsAaaaRecordTypeAaaa DNSRecordPostDNSRecordsAaaaRecordType = "AAAA"
)

func (r DNSRecordPostDNSRecordsAaaaRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsAaaaRecordTypeAaaa:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsCaaRecordParam struct {
	// Components of a CAA record.
	Data param.Field[DNSRecordPostDNSRecordsCaaRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsCaaRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsCaaRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsCaaRecordParam) implementsDNSRecordPostUnionParam() {}

// Components of a CAA record.
type DNSRecordPostDNSRecordsCaaRecordDataParam struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r DNSRecordPostDNSRecordsCaaRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPostDNSRecordsCaaRecordType string

const (
	DNSRecordPostDNSRecordsCaaRecordTypeCaa DNSRecordPostDNSRecordsCaaRecordType = "CAA"
)

func (r DNSRecordPostDNSRecordsCaaRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsCaaRecordTypeCaa:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsCertRecordParam struct {
	// Components of a CERT record.
	Data param.Field[DNSRecordPostDNSRecordsCertRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsCertRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsCertRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsCertRecordParam) implementsDNSRecordPostUnionParam() {}

// Components of a CERT record.
type DNSRecordPostDNSRecordsCertRecordDataParam struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r DNSRecordPostDNSRecordsCertRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPostDNSRecordsCertRecordType string

const (
	DNSRecordPostDNSRecordsCertRecordTypeCert DNSRecordPostDNSRecordsCertRecordType = "CERT"
)

func (r DNSRecordPostDNSRecordsCertRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsCertRecordTypeCert:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsCnameRecordParam struct {
	// A valid hostname. Must not match the record's name.
	Content  param.Field[string]                                          `json:"content"`
	Settings param.Field[DNSRecordPostDNSRecordsCnameRecordSettingsParam] `json:"settings"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsCnameRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsCnameRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsCnameRecordParam) implementsDNSRecordPostUnionParam() {}

type DNSRecordPostDNSRecordsCnameRecordSettingsParam struct {
	// If enabled, causes the CNAME record to be resolved externally and the resulting
	// address records (e.g., A and AAAA) to be returned instead of the CNAME record
	// itself. This setting is unavailable for proxied records, since they are always
	// flattened.
	FlattenCname param.Field[bool] `json:"flatten_cname"`
}

func (r DNSRecordPostDNSRecordsCnameRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPostDNSRecordsCnameRecordType string

const (
	DNSRecordPostDNSRecordsCnameRecordTypeCname DNSRecordPostDNSRecordsCnameRecordType = "CNAME"
)

func (r DNSRecordPostDNSRecordsCnameRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsCnameRecordTypeCname:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsDnskeyRecordParam struct {
	// Components of a DNSKEY record.
	Data param.Field[DNSRecordPostDNSRecordsDnskeyRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsDnskeyRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsDnskeyRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsDnskeyRecordParam) implementsDNSRecordPostUnionParam() {}

// Components of a DNSKEY record.
type DNSRecordPostDNSRecordsDnskeyRecordDataParam struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r DNSRecordPostDNSRecordsDnskeyRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPostDNSRecordsDnskeyRecordType string

const (
	DNSRecordPostDNSRecordsDnskeyRecordTypeDnskey DNSRecordPostDNSRecordsDnskeyRecordType = "DNSKEY"
)

func (r DNSRecordPostDNSRecordsDnskeyRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsDnskeyRecordTypeDnskey:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsDsRecordParam struct {
	// Components of a DS record.
	Data param.Field[DNSRecordPostDNSRecordsDsRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsDsRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsDsRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsDsRecordParam) implementsDNSRecordPostUnionParam() {}

// Components of a DS record.
type DNSRecordPostDNSRecordsDsRecordDataParam struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r DNSRecordPostDNSRecordsDsRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPostDNSRecordsDsRecordType string

const (
	DNSRecordPostDNSRecordsDsRecordTypeDs DNSRecordPostDNSRecordsDsRecordType = "DS"
)

func (r DNSRecordPostDNSRecordsDsRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsDsRecordTypeDs:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsHTTPSRecordParam struct {
	// Components of a HTTPS record.
	Data param.Field[DNSRecordPostDNSRecordsHTTPSRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsHTTPSRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsHTTPSRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsHTTPSRecordParam) implementsDNSRecordPostUnionParam() {}

// Components of a HTTPS record.
type DNSRecordPostDNSRecordsHTTPSRecordDataParam struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r DNSRecordPostDNSRecordsHTTPSRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPostDNSRecordsHTTPSRecordType string

const (
	DNSRecordPostDNSRecordsHTTPSRecordTypeHTTPS DNSRecordPostDNSRecordsHTTPSRecordType = "HTTPS"
)

func (r DNSRecordPostDNSRecordsHTTPSRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsHTTPSRecordTypeHTTPS:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsLocRecordParam struct {
	// Components of a LOC record.
	Data param.Field[DNSRecordPostDNSRecordsLocRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsLocRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsLocRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsLocRecordParam) implementsDNSRecordPostUnionParam() {}

// Components of a LOC record.
type DNSRecordPostDNSRecordsLocRecordDataParam struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[DNSRecordPostDNSRecordsLocRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[DNSRecordPostDNSRecordsLocRecordDataLongDirection] `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes param.Field[float64] `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds param.Field[float64] `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz param.Field[float64] `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert param.Field[float64] `json:"precision_vert"`
	// Size of location in meters.
	Size param.Field[float64] `json:"size"`
}

func (r DNSRecordPostDNSRecordsLocRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type DNSRecordPostDNSRecordsLocRecordDataLatDirection string

const (
	DNSRecordPostDNSRecordsLocRecordDataLatDirectionN DNSRecordPostDNSRecordsLocRecordDataLatDirection = "N"
	DNSRecordPostDNSRecordsLocRecordDataLatDirectionS DNSRecordPostDNSRecordsLocRecordDataLatDirection = "S"
)

func (r DNSRecordPostDNSRecordsLocRecordDataLatDirection) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsLocRecordDataLatDirectionN, DNSRecordPostDNSRecordsLocRecordDataLatDirectionS:
		return true
	}
	return false
}

// Longitude direction.
type DNSRecordPostDNSRecordsLocRecordDataLongDirection string

const (
	DNSRecordPostDNSRecordsLocRecordDataLongDirectionE DNSRecordPostDNSRecordsLocRecordDataLongDirection = "E"
	DNSRecordPostDNSRecordsLocRecordDataLongDirectionW DNSRecordPostDNSRecordsLocRecordDataLongDirection = "W"
)

func (r DNSRecordPostDNSRecordsLocRecordDataLongDirection) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsLocRecordDataLongDirectionE, DNSRecordPostDNSRecordsLocRecordDataLongDirectionW:
		return true
	}
	return false
}

// Record type.
type DNSRecordPostDNSRecordsLocRecordType string

const (
	DNSRecordPostDNSRecordsLocRecordTypeLoc DNSRecordPostDNSRecordsLocRecordType = "LOC"
)

func (r DNSRecordPostDNSRecordsLocRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsLocRecordTypeLoc:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsMxRecordParam struct {
	// A valid mail server hostname.
	Content param.Field[string] `json:"content" format:"hostname"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsMxRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsMxRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsMxRecordParam) implementsDNSRecordPostUnionParam() {}

// Record type.
type DNSRecordPostDNSRecordsMxRecordType string

const (
	DNSRecordPostDNSRecordsMxRecordTypeMx DNSRecordPostDNSRecordsMxRecordType = "MX"
)

func (r DNSRecordPostDNSRecordsMxRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsMxRecordTypeMx:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsNaptrRecordParam struct {
	// Components of a NAPTR record.
	Data param.Field[DNSRecordPostDNSRecordsNaptrRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsNaptrRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsNaptrRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsNaptrRecordParam) implementsDNSRecordPostUnionParam() {}

// Components of a NAPTR record.
type DNSRecordPostDNSRecordsNaptrRecordDataParam struct {
	// Flags.
	Flags param.Field[string] `json:"flags"`
	// Order.
	Order param.Field[float64] `json:"order"`
	// Preference.
	Preference param.Field[float64] `json:"preference"`
	// Regex.
	Regex param.Field[string] `json:"regex"`
	// Replacement.
	Replacement param.Field[string] `json:"replacement"`
	// Service.
	Service param.Field[string] `json:"service"`
}

func (r DNSRecordPostDNSRecordsNaptrRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPostDNSRecordsNaptrRecordType string

const (
	DNSRecordPostDNSRecordsNaptrRecordTypeNaptr DNSRecordPostDNSRecordsNaptrRecordType = "NAPTR"
)

func (r DNSRecordPostDNSRecordsNaptrRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsNaptrRecordTypeNaptr:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsNsRecordParam struct {
	// A valid name server host name.
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsNsRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsNsRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsNsRecordParam) implementsDNSRecordPostUnionParam() {}

// Record type.
type DNSRecordPostDNSRecordsNsRecordType string

const (
	DNSRecordPostDNSRecordsNsRecordTypeNs DNSRecordPostDNSRecordsNsRecordType = "NS"
)

func (r DNSRecordPostDNSRecordsNsRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsNsRecordTypeNs:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsOpenpgpkeyRecordParam struct {
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsOpenpgpkeyRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsOpenpgpkeyRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsOpenpgpkeyRecordParam) implementsDNSRecordPostUnionParam() {}

// Record type.
type DNSRecordPostDNSRecordsOpenpgpkeyRecordType string

const (
	DNSRecordPostDNSRecordsOpenpgpkeyRecordTypeOpenpgpkey DNSRecordPostDNSRecordsOpenpgpkeyRecordType = "OPENPGPKEY"
)

func (r DNSRecordPostDNSRecordsOpenpgpkeyRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsOpenpgpkeyRecordTypeOpenpgpkey:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsPtrRecordParam struct {
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsPtrRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsPtrRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsPtrRecordParam) implementsDNSRecordPostUnionParam() {}

// Record type.
type DNSRecordPostDNSRecordsPtrRecordType string

const (
	DNSRecordPostDNSRecordsPtrRecordTypePtr DNSRecordPostDNSRecordsPtrRecordType = "PTR"
)

func (r DNSRecordPostDNSRecordsPtrRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsPtrRecordTypePtr:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsSmimeaRecordParam struct {
	// Components of a SMIMEA record.
	Data param.Field[DNSRecordPostDNSRecordsSmimeaRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsSmimeaRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsSmimeaRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsSmimeaRecordParam) implementsDNSRecordPostUnionParam() {}

// Components of a SMIMEA record.
type DNSRecordPostDNSRecordsSmimeaRecordDataParam struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r DNSRecordPostDNSRecordsSmimeaRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPostDNSRecordsSmimeaRecordType string

const (
	DNSRecordPostDNSRecordsSmimeaRecordTypeSmimea DNSRecordPostDNSRecordsSmimeaRecordType = "SMIMEA"
)

func (r DNSRecordPostDNSRecordsSmimeaRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsSmimeaRecordTypeSmimea:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsSrvRecordParam struct {
	// Components of a SRV record.
	Data param.Field[DNSRecordPostDNSRecordsSrvRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsSrvRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsSrvRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsSrvRecordParam) implementsDNSRecordPostUnionParam() {}

// Components of a SRV record.
type DNSRecordPostDNSRecordsSrvRecordDataParam struct {
	// The port of the service.
	Port param.Field[float64] `json:"port"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// A valid hostname.
	Target param.Field[string] `json:"target" format:"hostname"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r DNSRecordPostDNSRecordsSrvRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPostDNSRecordsSrvRecordType string

const (
	DNSRecordPostDNSRecordsSrvRecordTypeSrv DNSRecordPostDNSRecordsSrvRecordType = "SRV"
)

func (r DNSRecordPostDNSRecordsSrvRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsSrvRecordTypeSrv:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsSshfpRecordParam struct {
	// Components of a SSHFP record.
	Data param.Field[DNSRecordPostDNSRecordsSshfpRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsSshfpRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsSshfpRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsSshfpRecordParam) implementsDNSRecordPostUnionParam() {}

// Components of a SSHFP record.
type DNSRecordPostDNSRecordsSshfpRecordDataParam struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r DNSRecordPostDNSRecordsSshfpRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPostDNSRecordsSshfpRecordType string

const (
	DNSRecordPostDNSRecordsSshfpRecordTypeSshfp DNSRecordPostDNSRecordsSshfpRecordType = "SSHFP"
)

func (r DNSRecordPostDNSRecordsSshfpRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsSshfpRecordTypeSshfp:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsSvcbRecordParam struct {
	// Components of a SVCB record.
	Data param.Field[DNSRecordPostDNSRecordsSvcbRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsSvcbRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsSvcbRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsSvcbRecordParam) implementsDNSRecordPostUnionParam() {}

// Components of a SVCB record.
type DNSRecordPostDNSRecordsSvcbRecordDataParam struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r DNSRecordPostDNSRecordsSvcbRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPostDNSRecordsSvcbRecordType string

const (
	DNSRecordPostDNSRecordsSvcbRecordTypeSvcb DNSRecordPostDNSRecordsSvcbRecordType = "SVCB"
)

func (r DNSRecordPostDNSRecordsSvcbRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsSvcbRecordTypeSvcb:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsTlsaRecordParam struct {
	// Components of a TLSA record.
	Data param.Field[DNSRecordPostDNSRecordsTlsaRecordDataParam] `json:"data"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsTlsaRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsTlsaRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsTlsaRecordParam) implementsDNSRecordPostUnionParam() {}

// Components of a TLSA record.
type DNSRecordPostDNSRecordsTlsaRecordDataParam struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r DNSRecordPostDNSRecordsTlsaRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPostDNSRecordsTlsaRecordType string

const (
	DNSRecordPostDNSRecordsTlsaRecordTypeTlsa DNSRecordPostDNSRecordsTlsaRecordType = "TLSA"
)

func (r DNSRecordPostDNSRecordsTlsaRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsTlsaRecordTypeTlsa:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsTxtRecordParam struct {
	// Text content for the record. The content must consist of quoted "character
	// strings" (RFC 1035), each with a length of up to 255 bytes. Strings exceeding
	// this allowed maximum length are automatically split.
	//
	// Learn more at
	// <https://www.cloudflare.com/learning/dns/dns-records/dns-txt-record/>.
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsTxtRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsTxtRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsTxtRecordParam) implementsDNSRecordPostUnionParam() {}

// Record type.
type DNSRecordPostDNSRecordsTxtRecordType string

const (
	DNSRecordPostDNSRecordsTxtRecordTypeTxt DNSRecordPostDNSRecordsTxtRecordType = "TXT"
)

func (r DNSRecordPostDNSRecordsTxtRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsTxtRecordTypeTxt:
		return true
	}
	return false
}

type DNSRecordPostDNSRecordsUriRecordParam struct {
	// Components of a URI record.
	Data param.Field[DNSRecordPostDNSRecordsUriRecordDataParam] `json:"data"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsUriRecordType] `json:"type"`
	SharedFieldsParam
}

func (r DNSRecordPostDNSRecordsUriRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsUriRecordParam) implementsDNSRecordPostUnionParam() {}

// Components of a URI record.
type DNSRecordPostDNSRecordsUriRecordDataParam struct {
	// The record content.
	Target param.Field[string] `json:"target"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r DNSRecordPostDNSRecordsUriRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPostDNSRecordsUriRecordType string

const (
	DNSRecordPostDNSRecordsUriRecordTypeUri DNSRecordPostDNSRecordsUriRecordType = "URI"
)

func (r DNSRecordPostDNSRecordsUriRecordType) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsUriRecordTypeUri:
		return true
	}
	return false
}

type DNSRecordResponse struct {
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// This field can have the runtime type of
	// [DNSRecordResponseDNSRecordsCaaRecordData],
	// [DNSRecordResponseDNSRecordsCertRecordData],
	// [DNSRecordResponseDNSRecordsDnskeyRecordData],
	// [DNSRecordResponseDNSRecordsDsRecordData],
	// [DNSRecordResponseDNSRecordsHTTPSRecordData],
	// [DNSRecordResponseDNSRecordsLocRecordData],
	// [DNSRecordResponseDNSRecordsNaptrRecordData],
	// [DNSRecordResponseDNSRecordsSmimeaRecordData],
	// [DNSRecordResponseDNSRecordsSrvRecordData],
	// [DNSRecordResponseDNSRecordsSshfpRecordData],
	// [DNSRecordResponseDNSRecordsSvcbRecordData],
	// [DNSRecordResponseDNSRecordsTlsaRecordData],
	// [DNSRecordResponseDNSRecordsUriRecordData].
	Data interface{} `json:"data"`
	// This field can have the runtime type of [interface{}].
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// This field can have the runtime type of [SharedFieldsSettings].
	Settings interface{} `json:"settings"`
	// This field can have the runtime type of [[]string].
	Tags interface{} `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// This field can have the runtime type of [SharedFieldsTtl].
	Ttl interface{} `json:"ttl"`
	// Record type.
	Type  DNSRecordResponseType `json:"type"`
	JSON  dnsRecordResponseJSON `json:"-"`
	union DNSRecordResponseUnion
}

// dnsRecordResponseJSON contains the JSON metadata for the struct
// [DNSRecordResponse]
type dnsRecordResponseJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Priority          apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r dnsRecordResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DNSRecordResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DNSRecordResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DNSRecordResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [DNSRecordResponseDNSRecordsARecord],
// [DNSRecordResponseDNSRecordsAaaaRecord], [DNSRecordResponseDNSRecordsCaaRecord],
// [DNSRecordResponseDNSRecordsCertRecord],
// [DNSRecordResponseDNSRecordsCnameRecord],
// [DNSRecordResponseDNSRecordsDnskeyRecord],
// [DNSRecordResponseDNSRecordsDsRecord], [DNSRecordResponseDNSRecordsHTTPSRecord],
// [DNSRecordResponseDNSRecordsLocRecord], [DNSRecordResponseDNSRecordsMxRecord],
// [DNSRecordResponseDNSRecordsNaptrRecord], [DNSRecordResponseDNSRecordsNsRecord],
// [DNSRecordResponseDNSRecordsOpenpgpkeyRecord],
// [DNSRecordResponseDNSRecordsPtrRecord],
// [DNSRecordResponseDNSRecordsSmimeaRecord],
// [DNSRecordResponseDNSRecordsSrvRecord],
// [DNSRecordResponseDNSRecordsSshfpRecord],
// [DNSRecordResponseDNSRecordsSvcbRecord],
// [DNSRecordResponseDNSRecordsTlsaRecord], [DNSRecordResponseDNSRecordsTxtRecord],
// [DNSRecordResponseDNSRecordsUriRecord].
func (r DNSRecordResponse) AsUnion() DNSRecordResponseUnion {
	return r.union
}

// Union satisfied by [DNSRecordResponseDNSRecordsARecord],
// [DNSRecordResponseDNSRecordsAaaaRecord], [DNSRecordResponseDNSRecordsCaaRecord],
// [DNSRecordResponseDNSRecordsCertRecord],
// [DNSRecordResponseDNSRecordsCnameRecord],
// [DNSRecordResponseDNSRecordsDnskeyRecord],
// [DNSRecordResponseDNSRecordsDsRecord], [DNSRecordResponseDNSRecordsHTTPSRecord],
// [DNSRecordResponseDNSRecordsLocRecord], [DNSRecordResponseDNSRecordsMxRecord],
// [DNSRecordResponseDNSRecordsNaptrRecord], [DNSRecordResponseDNSRecordsNsRecord],
// [DNSRecordResponseDNSRecordsOpenpgpkeyRecord],
// [DNSRecordResponseDNSRecordsPtrRecord],
// [DNSRecordResponseDNSRecordsSmimeaRecord],
// [DNSRecordResponseDNSRecordsSrvRecord],
// [DNSRecordResponseDNSRecordsSshfpRecord],
// [DNSRecordResponseDNSRecordsSvcbRecord],
// [DNSRecordResponseDNSRecordsTlsaRecord], [DNSRecordResponseDNSRecordsTxtRecord]
// or [DNSRecordResponseDNSRecordsUriRecord].
type DNSRecordResponseUnion interface {
	implementsDNSRecordResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsAaaaRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsCaaRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsCertRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsCnameRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsDnskeyRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsDsRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsHTTPSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsLocRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsMxRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsNaptrRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsNsRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsOpenpgpkeyRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsPtrRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsSmimeaRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsSrvRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsSshfpRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsSvcbRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsTlsaRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsTxtRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsUriRecord{}),
		},
	)
}

type DNSRecordResponseDNSRecordsARecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsARecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsARecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsARecordJSON contains the JSON metadata for the struct
// [DNSRecordResponseDNSRecordsARecord]
type dnsRecordResponseDNSRecordsARecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsARecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsARecord) implementsDNSRecordResponse() {}

// Record type.
type DNSRecordResponseDNSRecordsARecordType string

const (
	DNSRecordResponseDNSRecordsARecordTypeA DNSRecordResponseDNSRecordsARecordType = "A"
)

func (r DNSRecordResponseDNSRecordsARecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsARecordTypeA:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsAaaaRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// A valid IPv6 address.
	Content string `json:"content" format:"ipv6"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsAaaaRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsAaaaRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsAaaaRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsAaaaRecord]
type dnsRecordResponseDNSRecordsAaaaRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsAaaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsAaaaRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsAaaaRecord) implementsDNSRecordResponse() {}

// Record type.
type DNSRecordResponseDNSRecordsAaaaRecordType string

const (
	DNSRecordResponseDNSRecordsAaaaRecordTypeAaaa DNSRecordResponseDNSRecordsAaaaRecordType = "AAAA"
)

func (r DNSRecordResponseDNSRecordsAaaaRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsAaaaRecordTypeAaaa:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsCaaRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// Formatted CAA content. See 'data' to set CAA properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Components of a CAA record.
	Data DNSRecordResponseDNSRecordsCaaRecordData `json:"data"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsCaaRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsCaaRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsCaaRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsCaaRecord]
type dnsRecordResponseDNSRecordsCaaRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsCaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsCaaRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsCaaRecord) implementsDNSRecordResponse() {}

// Components of a CAA record.
type DNSRecordResponseDNSRecordsCaaRecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                       `json:"value"`
	JSON  dnsRecordResponseDNSRecordsCaaRecordDataJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsCaaRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsCaaRecordData]
type dnsRecordResponseDNSRecordsCaaRecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsCaaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsCaaRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordResponseDNSRecordsCaaRecordType string

const (
	DNSRecordResponseDNSRecordsCaaRecordTypeCaa DNSRecordResponseDNSRecordsCaaRecordType = "CAA"
)

func (r DNSRecordResponseDNSRecordsCaaRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsCaaRecordTypeCaa:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsCertRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// Formatted CERT content. See 'data' to set CERT properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Components of a CERT record.
	Data DNSRecordResponseDNSRecordsCertRecordData `json:"data"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsCertRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsCertRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsCertRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsCertRecord]
type dnsRecordResponseDNSRecordsCertRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsCertRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsCertRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsCertRecord) implementsDNSRecordResponse() {}

// Components of a CERT record.
type DNSRecordResponseDNSRecordsCertRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                                       `json:"type"`
	JSON dnsRecordResponseDNSRecordsCertRecordDataJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsCertRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsCertRecordData]
type dnsRecordResponseDNSRecordsCertRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsCertRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsCertRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordResponseDNSRecordsCertRecordType string

const (
	DNSRecordResponseDNSRecordsCertRecordTypeCert DNSRecordResponseDNSRecordsCertRecordType = "CERT"
)

func (r DNSRecordResponseDNSRecordsCertRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsCertRecordTypeCert:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsCnameRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// A valid hostname. Must not match the record's name.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool                                           `json:"proxiable"`
	Settings  DNSRecordResponseDNSRecordsCnameRecordSettings `json:"settings"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsCnameRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsCnameRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsCnameRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsCnameRecord]
type dnsRecordResponseDNSRecordsCnameRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Settings          apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsCnameRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsCnameRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsCnameRecord) implementsDNSRecordResponse() {}

type DNSRecordResponseDNSRecordsCnameRecordSettings struct {
	// If enabled, causes the CNAME record to be resolved externally and the resulting
	// address records (e.g., A and AAAA) to be returned instead of the CNAME record
	// itself. This setting is unavailable for proxied records, since they are always
	// flattened.
	FlattenCname bool                                               `json:"flatten_cname"`
	JSON         dnsRecordResponseDNSRecordsCnameRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsCnameRecordSettingsJSON contains the JSON metadata
// for the struct [DNSRecordResponseDNSRecordsCnameRecordSettings]
type dnsRecordResponseDNSRecordsCnameRecordSettingsJSON struct {
	FlattenCname apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsCnameRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsCnameRecordSettingsJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordResponseDNSRecordsCnameRecordType string

const (
	DNSRecordResponseDNSRecordsCnameRecordTypeCname DNSRecordResponseDNSRecordsCnameRecordType = "CNAME"
)

func (r DNSRecordResponseDNSRecordsCnameRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsCnameRecordTypeCname:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsDnskeyRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// Formatted DNSKEY content. See 'data' to set DNSKEY properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Components of a DNSKEY record.
	Data DNSRecordResponseDNSRecordsDnskeyRecordData `json:"data"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsDnskeyRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsDnskeyRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsDnskeyRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsDnskeyRecord]
type dnsRecordResponseDNSRecordsDnskeyRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsDnskeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsDnskeyRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsDnskeyRecord) implementsDNSRecordResponse() {}

// Components of a DNSKEY record.
type DNSRecordResponseDNSRecordsDnskeyRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                          `json:"public_key"`
	JSON      dnsRecordResponseDNSRecordsDnskeyRecordDataJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsDnskeyRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsDnskeyRecordData]
type dnsRecordResponseDNSRecordsDnskeyRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsDnskeyRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsDnskeyRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordResponseDNSRecordsDnskeyRecordType string

const (
	DNSRecordResponseDNSRecordsDnskeyRecordTypeDnskey DNSRecordResponseDNSRecordsDnskeyRecordType = "DNSKEY"
)

func (r DNSRecordResponseDNSRecordsDnskeyRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsDnskeyRecordTypeDnskey:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsDsRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// Formatted DS content. See 'data' to set DS properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Components of a DS record.
	Data DNSRecordResponseDNSRecordsDsRecordData `json:"data"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsDsRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsDsRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsDsRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsDsRecord]
type dnsRecordResponseDNSRecordsDsRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsDsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsDsRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsDsRecord) implementsDNSRecordResponse() {}

// Components of a DS record.
type DNSRecordResponseDNSRecordsDsRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                     `json:"key_tag"`
	JSON   dnsRecordResponseDNSRecordsDsRecordDataJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsDsRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsDsRecordData]
type dnsRecordResponseDNSRecordsDsRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsDsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsDsRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordResponseDNSRecordsDsRecordType string

const (
	DNSRecordResponseDNSRecordsDsRecordTypeDs DNSRecordResponseDNSRecordsDsRecordType = "DS"
)

func (r DNSRecordResponseDNSRecordsDsRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsDsRecordTypeDs:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsHTTPSRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// Formatted HTTPS content. See 'data' to set HTTPS properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Components of a HTTPS record.
	Data DNSRecordResponseDNSRecordsHTTPSRecordData `json:"data"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsHTTPSRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsHTTPSRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsHTTPSRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsHTTPSRecord]
type dnsRecordResponseDNSRecordsHTTPSRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsHTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsHTTPSRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsHTTPSRecord) implementsDNSRecordResponse() {}

// Components of a HTTPS record.
type DNSRecordResponseDNSRecordsHTTPSRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                         `json:"value"`
	JSON  dnsRecordResponseDNSRecordsHTTPSRecordDataJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsHTTPSRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsHTTPSRecordData]
type dnsRecordResponseDNSRecordsHTTPSRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsHTTPSRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsHTTPSRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordResponseDNSRecordsHTTPSRecordType string

const (
	DNSRecordResponseDNSRecordsHTTPSRecordTypeHTTPS DNSRecordResponseDNSRecordsHTTPSRecordType = "HTTPS"
)

func (r DNSRecordResponseDNSRecordsHTTPSRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsHTTPSRecordTypeHTTPS:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsLocRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// Formatted LOC content. See 'data' to set LOC properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Components of a LOC record.
	Data DNSRecordResponseDNSRecordsLocRecordData `json:"data"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsLocRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsLocRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsLocRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsLocRecord]
type dnsRecordResponseDNSRecordsLocRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsLocRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsLocRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsLocRecord) implementsDNSRecordResponse() {}

// Components of a LOC record.
type DNSRecordResponseDNSRecordsLocRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordResponseDNSRecordsLocRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordResponseDNSRecordsLocRecordDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                                      `json:"size"`
	JSON dnsRecordResponseDNSRecordsLocRecordDataJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsLocRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsLocRecordData]
type dnsRecordResponseDNSRecordsLocRecordDataJSON struct {
	Altitude      apijson.Field
	LatDegrees    apijson.Field
	LatDirection  apijson.Field
	LatMinutes    apijson.Field
	LatSeconds    apijson.Field
	LongDegrees   apijson.Field
	LongDirection apijson.Field
	LongMinutes   apijson.Field
	LongSeconds   apijson.Field
	PrecisionHorz apijson.Field
	PrecisionVert apijson.Field
	Size          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsLocRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsLocRecordDataJSON) RawJSON() string {
	return r.raw
}

// Latitude direction.
type DNSRecordResponseDNSRecordsLocRecordDataLatDirection string

const (
	DNSRecordResponseDNSRecordsLocRecordDataLatDirectionN DNSRecordResponseDNSRecordsLocRecordDataLatDirection = "N"
	DNSRecordResponseDNSRecordsLocRecordDataLatDirectionS DNSRecordResponseDNSRecordsLocRecordDataLatDirection = "S"
)

func (r DNSRecordResponseDNSRecordsLocRecordDataLatDirection) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsLocRecordDataLatDirectionN, DNSRecordResponseDNSRecordsLocRecordDataLatDirectionS:
		return true
	}
	return false
}

// Longitude direction.
type DNSRecordResponseDNSRecordsLocRecordDataLongDirection string

const (
	DNSRecordResponseDNSRecordsLocRecordDataLongDirectionE DNSRecordResponseDNSRecordsLocRecordDataLongDirection = "E"
	DNSRecordResponseDNSRecordsLocRecordDataLongDirectionW DNSRecordResponseDNSRecordsLocRecordDataLongDirection = "W"
)

func (r DNSRecordResponseDNSRecordsLocRecordDataLongDirection) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsLocRecordDataLongDirectionE, DNSRecordResponseDNSRecordsLocRecordDataLongDirectionW:
		return true
	}
	return false
}

// Record type.
type DNSRecordResponseDNSRecordsLocRecordType string

const (
	DNSRecordResponseDNSRecordsLocRecordTypeLoc DNSRecordResponseDNSRecordsLocRecordType = "LOC"
)

func (r DNSRecordResponseDNSRecordsLocRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsLocRecordTypeLoc:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsMxRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// A valid mail server hostname.
	Content string `json:"content" format:"hostname"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsMxRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsMxRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsMxRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsMxRecord]
type dnsRecordResponseDNSRecordsMxRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Priority          apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsMxRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsMxRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsMxRecord) implementsDNSRecordResponse() {}

// Record type.
type DNSRecordResponseDNSRecordsMxRecordType string

const (
	DNSRecordResponseDNSRecordsMxRecordTypeMx DNSRecordResponseDNSRecordsMxRecordType = "MX"
)

func (r DNSRecordResponseDNSRecordsMxRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsMxRecordTypeMx:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsNaptrRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// Formatted NAPTR content. See 'data' to set NAPTR properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Components of a NAPTR record.
	Data DNSRecordResponseDNSRecordsNaptrRecordData `json:"data"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsNaptrRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsNaptrRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsNaptrRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsNaptrRecord]
type dnsRecordResponseDNSRecordsNaptrRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsNaptrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsNaptrRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsNaptrRecord) implementsDNSRecordResponse() {}

// Components of a NAPTR record.
type DNSRecordResponseDNSRecordsNaptrRecordData struct {
	// Flags.
	Flags string `json:"flags"`
	// Order.
	Order float64 `json:"order"`
	// Preference.
	Preference float64 `json:"preference"`
	// Regex.
	Regex string `json:"regex"`
	// Replacement.
	Replacement string `json:"replacement"`
	// Service.
	Service string                                         `json:"service"`
	JSON    dnsRecordResponseDNSRecordsNaptrRecordDataJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsNaptrRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsNaptrRecordData]
type dnsRecordResponseDNSRecordsNaptrRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsNaptrRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsNaptrRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordResponseDNSRecordsNaptrRecordType string

const (
	DNSRecordResponseDNSRecordsNaptrRecordTypeNaptr DNSRecordResponseDNSRecordsNaptrRecordType = "NAPTR"
)

func (r DNSRecordResponseDNSRecordsNaptrRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsNaptrRecordTypeNaptr:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsNsRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// A valid name server host name.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsNsRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsNsRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsNsRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsNsRecord]
type dnsRecordResponseDNSRecordsNsRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsNsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsNsRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsNsRecord) implementsDNSRecordResponse() {}

// Record type.
type DNSRecordResponseDNSRecordsNsRecordType string

const (
	DNSRecordResponseDNSRecordsNsRecordTypeNs DNSRecordResponseDNSRecordsNsRecordType = "NS"
)

func (r DNSRecordResponseDNSRecordsNsRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsNsRecordTypeNs:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsOpenpgpkeyRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsOpenpgpkeyRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsOpenpgpkeyRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsOpenpgpkeyRecordJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsOpenpgpkeyRecord]
type dnsRecordResponseDNSRecordsOpenpgpkeyRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsOpenpgpkeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsOpenpgpkeyRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsOpenpgpkeyRecord) implementsDNSRecordResponse() {}

// Record type.
type DNSRecordResponseDNSRecordsOpenpgpkeyRecordType string

const (
	DNSRecordResponseDNSRecordsOpenpgpkeyRecordTypeOpenpgpkey DNSRecordResponseDNSRecordsOpenpgpkeyRecordType = "OPENPGPKEY"
)

func (r DNSRecordResponseDNSRecordsOpenpgpkeyRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsOpenpgpkeyRecordTypeOpenpgpkey:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsPtrRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// Domain name pointing to the address.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsPtrRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsPtrRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsPtrRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsPtrRecord]
type dnsRecordResponseDNSRecordsPtrRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsPtrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsPtrRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsPtrRecord) implementsDNSRecordResponse() {}

// Record type.
type DNSRecordResponseDNSRecordsPtrRecordType string

const (
	DNSRecordResponseDNSRecordsPtrRecordTypePtr DNSRecordResponseDNSRecordsPtrRecordType = "PTR"
)

func (r DNSRecordResponseDNSRecordsPtrRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsPtrRecordTypePtr:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsSmimeaRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// Formatted SMIMEA content. See 'data' to set SMIMEA properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Components of a SMIMEA record.
	Data DNSRecordResponseDNSRecordsSmimeaRecordData `json:"data"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsSmimeaRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsSmimeaRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsSmimeaRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsSmimeaRecord]
type dnsRecordResponseDNSRecordsSmimeaRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsSmimeaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsSmimeaRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsSmimeaRecord) implementsDNSRecordResponse() {}

// Components of a SMIMEA record.
type DNSRecordResponseDNSRecordsSmimeaRecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                         `json:"usage"`
	JSON  dnsRecordResponseDNSRecordsSmimeaRecordDataJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsSmimeaRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsSmimeaRecordData]
type dnsRecordResponseDNSRecordsSmimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsSmimeaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsSmimeaRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordResponseDNSRecordsSmimeaRecordType string

const (
	DNSRecordResponseDNSRecordsSmimeaRecordTypeSmimea DNSRecordResponseDNSRecordsSmimeaRecordType = "SMIMEA"
)

func (r DNSRecordResponseDNSRecordsSmimeaRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsSmimeaRecordTypeSmimea:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsSrvRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// Priority, weight, port, and SRV target. See 'data' for setting the individual
	// component values.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Components of a SRV record.
	Data DNSRecordResponseDNSRecordsSrvRecordData `json:"data"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsSrvRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsSrvRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsSrvRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsSrvRecord]
type dnsRecordResponseDNSRecordsSrvRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsSrvRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsSrvRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsSrvRecord) implementsDNSRecordResponse() {}

// Components of a SRV record.
type DNSRecordResponseDNSRecordsSrvRecordData struct {
	// The port of the service.
	Port float64 `json:"port"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// A valid hostname.
	Target string `json:"target" format:"hostname"`
	// The record weight.
	Weight float64                                      `json:"weight"`
	JSON   dnsRecordResponseDNSRecordsSrvRecordDataJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsSrvRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsSrvRecordData]
type dnsRecordResponseDNSRecordsSrvRecordDataJSON struct {
	Port        apijson.Field
	Priority    apijson.Field
	Target      apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsSrvRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsSrvRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordResponseDNSRecordsSrvRecordType string

const (
	DNSRecordResponseDNSRecordsSrvRecordTypeSrv DNSRecordResponseDNSRecordsSrvRecordType = "SRV"
)

func (r DNSRecordResponseDNSRecordsSrvRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsSrvRecordTypeSrv:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsSshfpRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// Formatted SSHFP content. See 'data' to set SSHFP properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Components of a SSHFP record.
	Data DNSRecordResponseDNSRecordsSshfpRecordData `json:"data"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsSshfpRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsSshfpRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsSshfpRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsSshfpRecord]
type dnsRecordResponseDNSRecordsSshfpRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsSshfpRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsSshfpRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsSshfpRecord) implementsDNSRecordResponse() {}

// Components of a SSHFP record.
type DNSRecordResponseDNSRecordsSshfpRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                        `json:"type"`
	JSON dnsRecordResponseDNSRecordsSshfpRecordDataJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsSshfpRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsSshfpRecordData]
type dnsRecordResponseDNSRecordsSshfpRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsSshfpRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsSshfpRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordResponseDNSRecordsSshfpRecordType string

const (
	DNSRecordResponseDNSRecordsSshfpRecordTypeSshfp DNSRecordResponseDNSRecordsSshfpRecordType = "SSHFP"
)

func (r DNSRecordResponseDNSRecordsSshfpRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsSshfpRecordTypeSshfp:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsSvcbRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// Formatted SVCB content. See 'data' to set SVCB properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Components of a SVCB record.
	Data DNSRecordResponseDNSRecordsSvcbRecordData `json:"data"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsSvcbRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsSvcbRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsSvcbRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsSvcbRecord]
type dnsRecordResponseDNSRecordsSvcbRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsSvcbRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsSvcbRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsSvcbRecord) implementsDNSRecordResponse() {}

// Components of a SVCB record.
type DNSRecordResponseDNSRecordsSvcbRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                        `json:"value"`
	JSON  dnsRecordResponseDNSRecordsSvcbRecordDataJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsSvcbRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsSvcbRecordData]
type dnsRecordResponseDNSRecordsSvcbRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsSvcbRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsSvcbRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordResponseDNSRecordsSvcbRecordType string

const (
	DNSRecordResponseDNSRecordsSvcbRecordTypeSvcb DNSRecordResponseDNSRecordsSvcbRecordType = "SVCB"
)

func (r DNSRecordResponseDNSRecordsSvcbRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsSvcbRecordTypeSvcb:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsTlsaRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// Formatted TLSA content. See 'data' to set TLSA properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Components of a TLSA record.
	Data DNSRecordResponseDNSRecordsTlsaRecordData `json:"data"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsTlsaRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsTlsaRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsTlsaRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsTlsaRecord]
type dnsRecordResponseDNSRecordsTlsaRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsTlsaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsTlsaRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsTlsaRecord) implementsDNSRecordResponse() {}

// Components of a TLSA record.
type DNSRecordResponseDNSRecordsTlsaRecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                       `json:"usage"`
	JSON  dnsRecordResponseDNSRecordsTlsaRecordDataJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsTlsaRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsTlsaRecordData]
type dnsRecordResponseDNSRecordsTlsaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsTlsaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsTlsaRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordResponseDNSRecordsTlsaRecordType string

const (
	DNSRecordResponseDNSRecordsTlsaRecordTypeTlsa DNSRecordResponseDNSRecordsTlsaRecordType = "TLSA"
)

func (r DNSRecordResponseDNSRecordsTlsaRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsTlsaRecordTypeTlsa:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsTxtRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// Text content for the record. The content must consist of quoted "character
	// strings" (RFC 1035), each with a length of up to 255 bytes. Strings exceeding
	// this allowed maximum length are automatically split.
	//
	// Learn more at
	// <https://www.cloudflare.com/learning/dns/dns-records/dns-txt-record/>.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsTxtRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsTxtRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsTxtRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsTxtRecord]
type dnsRecordResponseDNSRecordsTxtRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsTxtRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsTxtRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsTxtRecord) implementsDNSRecordResponse() {}

// Record type.
type DNSRecordResponseDNSRecordsTxtRecordType string

const (
	DNSRecordResponseDNSRecordsTxtRecordTypeTxt DNSRecordResponseDNSRecordsTxtRecordType = "TXT"
)

func (r DNSRecordResponseDNSRecordsTxtRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsTxtRecordTypeTxt:
		return true
	}
	return false
}

type DNSRecordResponseDNSRecordsUriRecord struct {
	// Identifier
	ID string `json:"id"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// Formatted URI content. See 'data' to set URI properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Components of a URI record.
	Data DNSRecordResponseDNSRecordsUriRecordData `json:"data"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Record type.
	Type DNSRecordResponseDNSRecordsUriRecordType `json:"type"`
	JSON dnsRecordResponseDNSRecordsUriRecordJSON `json:"-"`
	SharedFields
}

// dnsRecordResponseDNSRecordsUriRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsUriRecord]
type dnsRecordResponseDNSRecordsUriRecordJSON struct {
	ID                apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Priority          apijson.Field
	Proxiable         apijson.Field
	TagsModifiedOn    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsUriRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsUriRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordResponseDNSRecordsUriRecord) implementsDNSRecordResponse() {}

// Components of a URI record.
type DNSRecordResponseDNSRecordsUriRecordData struct {
	// The record content.
	Target string `json:"target"`
	// The record weight.
	Weight float64                                      `json:"weight"`
	JSON   dnsRecordResponseDNSRecordsUriRecordDataJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsUriRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsUriRecordData]
type dnsRecordResponseDNSRecordsUriRecordDataJSON struct {
	Target      apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsUriRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsUriRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordResponseDNSRecordsUriRecordType string

const (
	DNSRecordResponseDNSRecordsUriRecordTypeUri DNSRecordResponseDNSRecordsUriRecordType = "URI"
)

func (r DNSRecordResponseDNSRecordsUriRecordType) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsUriRecordTypeUri:
		return true
	}
	return false
}

// Record type.
type DNSRecordResponseType string

const (
	DNSRecordResponseTypeA          DNSRecordResponseType = "A"
	DNSRecordResponseTypeAaaa       DNSRecordResponseType = "AAAA"
	DNSRecordResponseTypeCaa        DNSRecordResponseType = "CAA"
	DNSRecordResponseTypeCert       DNSRecordResponseType = "CERT"
	DNSRecordResponseTypeCname      DNSRecordResponseType = "CNAME"
	DNSRecordResponseTypeDnskey     DNSRecordResponseType = "DNSKEY"
	DNSRecordResponseTypeDs         DNSRecordResponseType = "DS"
	DNSRecordResponseTypeHTTPS      DNSRecordResponseType = "HTTPS"
	DNSRecordResponseTypeLoc        DNSRecordResponseType = "LOC"
	DNSRecordResponseTypeMx         DNSRecordResponseType = "MX"
	DNSRecordResponseTypeNaptr      DNSRecordResponseType = "NAPTR"
	DNSRecordResponseTypeNs         DNSRecordResponseType = "NS"
	DNSRecordResponseTypeOpenpgpkey DNSRecordResponseType = "OPENPGPKEY"
	DNSRecordResponseTypePtr        DNSRecordResponseType = "PTR"
	DNSRecordResponseTypeSmimea     DNSRecordResponseType = "SMIMEA"
	DNSRecordResponseTypeSrv        DNSRecordResponseType = "SRV"
	DNSRecordResponseTypeSshfp      DNSRecordResponseType = "SSHFP"
	DNSRecordResponseTypeSvcb       DNSRecordResponseType = "SVCB"
	DNSRecordResponseTypeTlsa       DNSRecordResponseType = "TLSA"
	DNSRecordResponseTypeTxt        DNSRecordResponseType = "TXT"
	DNSRecordResponseTypeUri        DNSRecordResponseType = "URI"
)

func (r DNSRecordResponseType) IsKnown() bool {
	switch r {
	case DNSRecordResponseTypeA, DNSRecordResponseTypeAaaa, DNSRecordResponseTypeCaa, DNSRecordResponseTypeCert, DNSRecordResponseTypeCname, DNSRecordResponseTypeDnskey, DNSRecordResponseTypeDs, DNSRecordResponseTypeHTTPS, DNSRecordResponseTypeLoc, DNSRecordResponseTypeMx, DNSRecordResponseTypeNaptr, DNSRecordResponseTypeNs, DNSRecordResponseTypeOpenpgpkey, DNSRecordResponseTypePtr, DNSRecordResponseTypeSmimea, DNSRecordResponseTypeSrv, DNSRecordResponseTypeSshfp, DNSRecordResponseTypeSvcb, DNSRecordResponseTypeTlsa, DNSRecordResponseTypeTxt, DNSRecordResponseTypeUri:
		return true
	}
	return false
}

type ImportScanResponse struct {
	Result ImportScanResponseResult `json:"result"`
	JSON   importScanResponseJSON   `json:"-"`
	SingleResponseDNSRecord
}

// importScanResponseJSON contains the JSON metadata for the struct
// [ImportScanResponse]
type importScanResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImportScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r importScanResponseJSON) RawJSON() string {
	return r.raw
}

type ImportScanResponseResult struct {
	// Number of DNS records added.
	RecsAdded float64 `json:"recs_added"`
	// Total number of DNS records parsed.
	TotalRecordsParsed float64                      `json:"total_records_parsed"`
	JSON               importScanResponseResultJSON `json:"-"`
}

// importScanResponseResultJSON contains the JSON metadata for the struct
// [ImportScanResponseResult]
type importScanResponseResultJSON struct {
	RecsAdded          apijson.Field
	TotalRecordsParsed apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ImportScanResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r importScanResponseResultJSON) RawJSON() string {
	return r.raw
}

type SharedFields struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings SharedFieldsSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl  SharedFieldsTtl  `json:"ttl"`
	JSON sharedFieldsJSON `json:"-"`
}

// sharedFieldsJSON contains the JSON metadata for the struct [SharedFields]
type sharedFieldsJSON struct {
	Comment     apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Settings    apijson.Field
	Tags        apijson.Field
	Ttl         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SharedFields) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sharedFieldsJSON) RawJSON() string {
	return r.raw
}

// Settings for the DNS record.
type SharedFieldsSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                     `json:"ipv6_only"`
	JSON     sharedFieldsSettingsJSON `json:"-"`
}

// sharedFieldsSettingsJSON contains the JSON metadata for the struct
// [SharedFieldsSettings]
type sharedFieldsSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SharedFieldsSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sharedFieldsSettingsJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type SharedFieldsTtl float64

const (
	SharedFieldsTtl1 SharedFieldsTtl = 1
)

func (r SharedFieldsTtl) IsKnown() bool {
	switch r {
	case SharedFieldsTtl1:
		return true
	}
	return false
}

type SharedFieldsParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[SharedFieldsSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[SharedFieldsTtl] `json:"ttl"`
}

func (r SharedFieldsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for the DNS record.
type SharedFieldsSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only param.Field[bool] `json:"ipv6_only"`
}

func (r SharedFieldsSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SingleResponseDNSRecord struct {
	Errors   []DNSRecordMessageItem `json:"errors,required"`
	Messages []DNSRecordMessageItem `json:"messages,required"`
	// Whether the API call was successful
	Success SingleResponseDNSRecordSuccess `json:"success,required"`
	JSON    singleResponseDNSRecordJSON    `json:"-"`
}

// singleResponseDNSRecordJSON contains the JSON metadata for the struct
// [SingleResponseDNSRecord]
type singleResponseDNSRecordJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseDNSRecordJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SingleResponseDNSRecordSuccess bool

const (
	SingleResponseDNSRecordSuccessTrue SingleResponseDNSRecordSuccess = true
)

func (r SingleResponseDNSRecordSuccess) IsKnown() bool {
	switch r {
	case SingleResponseDNSRecordSuccessTrue:
		return true
	}
	return false
}

type SingleResponseDNSResponse struct {
	Result DNSRecordResponse             `json:"result"`
	JSON   singleResponseDNSResponseJSON `json:"-"`
	SingleResponseDNSRecord
}

// singleResponseDNSResponseJSON contains the JSON metadata for the struct
// [SingleResponseDNSResponse]
type singleResponseDNSResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseDNSResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseDNSResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneDNSRecordListResponse struct {
	Result     []DNSRecordResponse                 `json:"result"`
	ResultInfo ZoneDNSRecordListResponseResultInfo `json:"result_info"`
	JSON       zoneDNSRecordListResponseJSON       `json:"-"`
	CommonResponseDNSRecords
}

// zoneDNSRecordListResponseJSON contains the JSON metadata for the struct
// [ZoneDNSRecordListResponse]
type zoneDNSRecordListResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDNSRecordListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneDNSRecordListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count"`
	JSON       zoneDNSRecordListResponseResultInfoJSON `json:"-"`
}

// zoneDNSRecordListResponseResultInfoJSON contains the JSON metadata for the
// struct [ZoneDNSRecordListResponseResultInfo]
type zoneDNSRecordListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDNSRecordListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneDNSRecordDeleteResponse struct {
	Result ZoneDNSRecordDeleteResponseResult `json:"result"`
	JSON   zoneDNSRecordDeleteResponseJSON   `json:"-"`
}

// zoneDNSRecordDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneDNSRecordDeleteResponse]
type zoneDNSRecordDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDNSRecordDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneDNSRecordDeleteResponseResult struct {
	// Identifier
	ID   string                                `json:"id"`
	JSON zoneDNSRecordDeleteResponseResultJSON `json:"-"`
}

// zoneDNSRecordDeleteResponseResultJSON contains the JSON metadata for the struct
// [ZoneDNSRecordDeleteResponseResult]
type zoneDNSRecordDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDNSRecordDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneDNSRecordBatchResponse struct {
	Result ZoneDNSRecordBatchResponseResult `json:"result"`
	JSON   zoneDNSRecordBatchResponseJSON   `json:"-"`
	SingleResponseDNSRecord
}

// zoneDNSRecordBatchResponseJSON contains the JSON metadata for the struct
// [ZoneDNSRecordBatchResponse]
type zoneDNSRecordBatchResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordBatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDNSRecordBatchResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneDNSRecordBatchResponseResult struct {
	Deletes []DNSRecordResponse                  `json:"deletes"`
	Patches []DNSRecordResponse                  `json:"patches"`
	Posts   []DNSRecordResponse                  `json:"posts"`
	Puts    []DNSRecordResponse                  `json:"puts"`
	JSON    zoneDNSRecordBatchResponseResultJSON `json:"-"`
}

// zoneDNSRecordBatchResponseResultJSON contains the JSON metadata for the struct
// [ZoneDNSRecordBatchResponseResult]
type zoneDNSRecordBatchResponseResultJSON struct {
	Deletes     apijson.Field
	Patches     apijson.Field
	Posts       apijson.Field
	Puts        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordBatchResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDNSRecordBatchResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneDNSRecordNewParams struct {
	DNSRecordPost DNSRecordPostUnionParam `json:"dns_record_post,required"`
}

func (r ZoneDNSRecordNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DNSRecordPost)
}

type ZoneDNSRecordUpdateParams struct {
	DNSRecordPatch DNSRecordPatchUnionParam `json:"dns_record_patch,required"`
}

func (r ZoneDNSRecordUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DNSRecordPatch)
}

type ZoneDNSRecordListParams struct {
	Comment param.Field[ZoneDNSRecordListParamsComment] `query:"comment"`
	Content param.Field[ZoneDNSRecordListParamsContent] `query:"content"`
	// Direction to order DNS records in.
	Direction param.Field[ZoneDNSRecordListParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any). If set to `all`,
	// acts like a logical AND between filters. If set to `any`, acts like a logical OR
	// instead. Note that the interaction between tag filters is controlled by the
	// `tag-match` parameter instead.
	Match param.Field[ZoneDNSRecordListParamsMatch] `query:"match"`
	Name  param.Field[ZoneDNSRecordListParamsName]  `query:"name"`
	// Field to order DNS records by.
	Order param.Field[ZoneDNSRecordListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of DNS records per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `query:"proxied"`
	// Allows searching in multiple properties of a DNS record simultaneously. This
	// parameter is intended for human users, not automation. Its exact behavior is
	// intentionally left unspecified and is subject to change in the future. This
	// parameter works independently of the `match` setting. For automated searches,
	// please use the other available parameters.
	Search param.Field[string]                     `query:"search"`
	Tag    param.Field[ZoneDNSRecordListParamsTag] `query:"tag"`
	// Whether to match all tag search requirements or at least one (any). If set to
	// `all`, acts like a logical AND between tag filters. If set to `any`, acts like a
	// logical OR instead. Note that the regular `match` parameter is still used to
	// combine the resulting condition with other filters that aren't related to tags.
	TagMatch param.Field[ZoneDNSRecordListParamsTagMatch] `query:"tag_match"`
	// Record type.
	Type param.Field[ZoneDNSRecordListParamsType] `query:"type"`
}

// URLQuery serializes [ZoneDNSRecordListParams]'s query parameters as
// `url.Values`.
func (r ZoneDNSRecordListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneDNSRecordListParamsComment struct {
	// If this parameter is present, only records _without_ a comment are returned.
	Absent param.Field[string] `query:"absent"`
	// Substring of the DNS record comment. Comment filters are case-insensitive.
	Contains param.Field[string] `query:"contains"`
	// Suffix of the DNS record comment. Comment filters are case-insensitive.
	Endswith param.Field[string] `query:"endswith"`
	// Exact value of the DNS record comment. Comment filters are case-insensitive.
	Exact param.Field[string] `query:"exact"`
	// If this parameter is present, only records _with_ a comment are returned.
	Present param.Field[string] `query:"present"`
	// Prefix of the DNS record comment. Comment filters are case-insensitive.
	Startswith param.Field[string] `query:"startswith"`
}

// URLQuery serializes [ZoneDNSRecordListParamsComment]'s query parameters as
// `url.Values`.
func (r ZoneDNSRecordListParamsComment) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneDNSRecordListParamsContent struct {
	// Substring of the DNS record content. Content filters are case-insensitive.
	Contains param.Field[string] `query:"contains"`
	// Suffix of the DNS record content. Content filters are case-insensitive.
	Endswith param.Field[string] `query:"endswith"`
	// Exact value of the DNS record content. Content filters are case-insensitive.
	Exact param.Field[string] `query:"exact"`
	// Prefix of the DNS record content. Content filters are case-insensitive.
	Startswith param.Field[string] `query:"startswith"`
}

// URLQuery serializes [ZoneDNSRecordListParamsContent]'s query parameters as
// `url.Values`.
func (r ZoneDNSRecordListParamsContent) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order DNS records in.
type ZoneDNSRecordListParamsDirection string

const (
	ZoneDNSRecordListParamsDirectionAsc  ZoneDNSRecordListParamsDirection = "asc"
	ZoneDNSRecordListParamsDirectionDesc ZoneDNSRecordListParamsDirection = "desc"
)

func (r ZoneDNSRecordListParamsDirection) IsKnown() bool {
	switch r {
	case ZoneDNSRecordListParamsDirectionAsc, ZoneDNSRecordListParamsDirectionDesc:
		return true
	}
	return false
}

// Whether to match all search requirements or at least one (any). If set to `all`,
// acts like a logical AND between filters. If set to `any`, acts like a logical OR
// instead. Note that the interaction between tag filters is controlled by the
// `tag-match` parameter instead.
type ZoneDNSRecordListParamsMatch string

const (
	ZoneDNSRecordListParamsMatchAny ZoneDNSRecordListParamsMatch = "any"
	ZoneDNSRecordListParamsMatchAll ZoneDNSRecordListParamsMatch = "all"
)

func (r ZoneDNSRecordListParamsMatch) IsKnown() bool {
	switch r {
	case ZoneDNSRecordListParamsMatchAny, ZoneDNSRecordListParamsMatchAll:
		return true
	}
	return false
}

type ZoneDNSRecordListParamsName struct {
	// Substring of the DNS record name. Name filters are case-insensitive.
	Contains param.Field[string] `query:"contains"`
	// Suffix of the DNS record name. Name filters are case-insensitive.
	Endswith param.Field[string] `query:"endswith"`
	// Exact value of the DNS record name. Name filters are case-insensitive.
	Exact param.Field[string] `query:"exact"`
	// Prefix of the DNS record name. Name filters are case-insensitive.
	Startswith param.Field[string] `query:"startswith"`
}

// URLQuery serializes [ZoneDNSRecordListParamsName]'s query parameters as
// `url.Values`.
func (r ZoneDNSRecordListParamsName) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Field to order DNS records by.
type ZoneDNSRecordListParamsOrder string

const (
	ZoneDNSRecordListParamsOrderType    ZoneDNSRecordListParamsOrder = "type"
	ZoneDNSRecordListParamsOrderName    ZoneDNSRecordListParamsOrder = "name"
	ZoneDNSRecordListParamsOrderContent ZoneDNSRecordListParamsOrder = "content"
	ZoneDNSRecordListParamsOrderTtl     ZoneDNSRecordListParamsOrder = "ttl"
	ZoneDNSRecordListParamsOrderProxied ZoneDNSRecordListParamsOrder = "proxied"
)

func (r ZoneDNSRecordListParamsOrder) IsKnown() bool {
	switch r {
	case ZoneDNSRecordListParamsOrderType, ZoneDNSRecordListParamsOrderName, ZoneDNSRecordListParamsOrderContent, ZoneDNSRecordListParamsOrderTtl, ZoneDNSRecordListParamsOrderProxied:
		return true
	}
	return false
}

type ZoneDNSRecordListParamsTag struct {
	// Name of a tag which must _not_ be present on the DNS record. Tag filters are
	// case-insensitive.
	Absent param.Field[string] `query:"absent"`
	// A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return
	// DNS records that have a tag named `<tag-name>` whose value contains
	// `<tag-value>`. Tag filters are case-insensitive.
	Contains param.Field[string] `query:"contains"`
	// A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return
	// DNS records that have a tag named `<tag-name>` whose value ends with
	// `<tag-value>`. Tag filters are case-insensitive.
	Endswith param.Field[string] `query:"endswith"`
	// A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return
	// DNS records that have a tag named `<tag-name>` whose value is `<tag-value>`. Tag
	// filters are case-insensitive.
	Exact param.Field[string] `query:"exact"`
	// Name of a tag which must be present on the DNS record. Tag filters are
	// case-insensitive.
	Present param.Field[string] `query:"present"`
	// A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return
	// DNS records that have a tag named `<tag-name>` whose value starts with
	// `<tag-value>`. Tag filters are case-insensitive.
	Startswith param.Field[string] `query:"startswith"`
}

// URLQuery serializes [ZoneDNSRecordListParamsTag]'s query parameters as
// `url.Values`.
func (r ZoneDNSRecordListParamsTag) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether to match all tag search requirements or at least one (any). If set to
// `all`, acts like a logical AND between tag filters. If set to `any`, acts like a
// logical OR instead. Note that the regular `match` parameter is still used to
// combine the resulting condition with other filters that aren't related to tags.
type ZoneDNSRecordListParamsTagMatch string

const (
	ZoneDNSRecordListParamsTagMatchAny ZoneDNSRecordListParamsTagMatch = "any"
	ZoneDNSRecordListParamsTagMatchAll ZoneDNSRecordListParamsTagMatch = "all"
)

func (r ZoneDNSRecordListParamsTagMatch) IsKnown() bool {
	switch r {
	case ZoneDNSRecordListParamsTagMatchAny, ZoneDNSRecordListParamsTagMatchAll:
		return true
	}
	return false
}

// Record type.
type ZoneDNSRecordListParamsType string

const (
	ZoneDNSRecordListParamsTypeA          ZoneDNSRecordListParamsType = "A"
	ZoneDNSRecordListParamsTypeAaaa       ZoneDNSRecordListParamsType = "AAAA"
	ZoneDNSRecordListParamsTypeCaa        ZoneDNSRecordListParamsType = "CAA"
	ZoneDNSRecordListParamsTypeCert       ZoneDNSRecordListParamsType = "CERT"
	ZoneDNSRecordListParamsTypeCname      ZoneDNSRecordListParamsType = "CNAME"
	ZoneDNSRecordListParamsTypeDnskey     ZoneDNSRecordListParamsType = "DNSKEY"
	ZoneDNSRecordListParamsTypeDs         ZoneDNSRecordListParamsType = "DS"
	ZoneDNSRecordListParamsTypeHTTPS      ZoneDNSRecordListParamsType = "HTTPS"
	ZoneDNSRecordListParamsTypeLoc        ZoneDNSRecordListParamsType = "LOC"
	ZoneDNSRecordListParamsTypeMx         ZoneDNSRecordListParamsType = "MX"
	ZoneDNSRecordListParamsTypeNaptr      ZoneDNSRecordListParamsType = "NAPTR"
	ZoneDNSRecordListParamsTypeNs         ZoneDNSRecordListParamsType = "NS"
	ZoneDNSRecordListParamsTypeOpenpgpkey ZoneDNSRecordListParamsType = "OPENPGPKEY"
	ZoneDNSRecordListParamsTypePtr        ZoneDNSRecordListParamsType = "PTR"
	ZoneDNSRecordListParamsTypeSmimea     ZoneDNSRecordListParamsType = "SMIMEA"
	ZoneDNSRecordListParamsTypeSrv        ZoneDNSRecordListParamsType = "SRV"
	ZoneDNSRecordListParamsTypeSshfp      ZoneDNSRecordListParamsType = "SSHFP"
	ZoneDNSRecordListParamsTypeSvcb       ZoneDNSRecordListParamsType = "SVCB"
	ZoneDNSRecordListParamsTypeTlsa       ZoneDNSRecordListParamsType = "TLSA"
	ZoneDNSRecordListParamsTypeTxt        ZoneDNSRecordListParamsType = "TXT"
	ZoneDNSRecordListParamsTypeUri        ZoneDNSRecordListParamsType = "URI"
)

func (r ZoneDNSRecordListParamsType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordListParamsTypeA, ZoneDNSRecordListParamsTypeAaaa, ZoneDNSRecordListParamsTypeCaa, ZoneDNSRecordListParamsTypeCert, ZoneDNSRecordListParamsTypeCname, ZoneDNSRecordListParamsTypeDnskey, ZoneDNSRecordListParamsTypeDs, ZoneDNSRecordListParamsTypeHTTPS, ZoneDNSRecordListParamsTypeLoc, ZoneDNSRecordListParamsTypeMx, ZoneDNSRecordListParamsTypeNaptr, ZoneDNSRecordListParamsTypeNs, ZoneDNSRecordListParamsTypeOpenpgpkey, ZoneDNSRecordListParamsTypePtr, ZoneDNSRecordListParamsTypeSmimea, ZoneDNSRecordListParamsTypeSrv, ZoneDNSRecordListParamsTypeSshfp, ZoneDNSRecordListParamsTypeSvcb, ZoneDNSRecordListParamsTypeTlsa, ZoneDNSRecordListParamsTypeTxt, ZoneDNSRecordListParamsTypeUri:
		return true
	}
	return false
}

type ZoneDNSRecordDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneDNSRecordDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneDNSRecordBatchParams struct {
	Deletes param.Field[[]ZoneDNSRecordBatchParamsDelete]     `json:"deletes"`
	Patches param.Field[[]ZoneDNSRecordBatchParamsPatchUnion] `json:"patches"`
	Posts   param.Field[[]DNSRecordPostUnionParam]            `json:"posts"`
	Puts    param.Field[[]ZoneDNSRecordBatchParamsPutUnion]   `json:"puts"`
}

func (r ZoneDNSRecordBatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsDelete struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneDNSRecordBatchParamsDelete) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [ZoneDNSRecordBatchParamsPatchesDNSRecordsARecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecord].
type ZoneDNSRecordBatchParamsPatchUnion interface {
	implementsZoneDNSRecordBatchParamsPatchUnion()
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsARecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// A valid IPv4 address.
	Content param.Field[string] `json:"content" format:"ipv4"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsARecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsARecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsARecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsARecordTypeA ZoneDNSRecordBatchParamsPatchesDNSRecordsARecordType = "A"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsARecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsARecordTypeA:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// A valid IPv6 address.
	Content param.Field[string] `json:"content" format:"ipv6"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecordTypeAaaa ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecordType = "AAAA"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecordTypeAaaa:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a CAA record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Components of a CAA record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordData struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordTypeCaa ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordType = "CAA"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordTypeCaa:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a CERT record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Components of a CERT record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordTypeCert ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordType = "CERT"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordTypeCert:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// A valid hostname. Must not match the record's name.
	Content  param.Field[string]                                                       `json:"content"`
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordSettings] `json:"settings"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordSettings struct {
	// If enabled, causes the CNAME record to be resolved externally and the resulting
	// address records (e.g., A and AAAA) to be returned instead of the CNAME record
	// itself. This setting is unavailable for proxied records, since they are always
	// flattened.
	FlattenCname param.Field[bool] `json:"flatten_cname"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordTypeCname ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordType = "CNAME"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordTypeCname:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a DNSKEY record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Components of a DNSKEY record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordTypeDnskey ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordType = "DNSKEY"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordTypeDnskey:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a DS record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Components of a DS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordTypeDs ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordType = "DS"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordTypeDs:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a HTTPS record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Components of a HTTPS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordTypeHTTPS ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordType = "HTTPS"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordTypeHTTPS:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a LOC record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Components of a LOC record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordData struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordDataLongDirection] `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes param.Field[float64] `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds param.Field[float64] `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz param.Field[float64] `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert param.Field[float64] `json:"precision_vert"`
	// Size of location in meters.
	Size param.Field[float64] `json:"size"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordDataLatDirection string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordDataLatDirectionN ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordDataLatDirection = "N"
	ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordDataLatDirectionS ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordDataLatDirection = "S"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordDataLatDirection) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordDataLatDirectionN, ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordDataLatDirectionS:
		return true
	}
	return false
}

// Longitude direction.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordDataLongDirection string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordDataLongDirectionE ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordDataLongDirection = "E"
	ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordDataLongDirectionW ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordDataLongDirection = "W"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordDataLongDirection) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordDataLongDirectionE, ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordDataLongDirectionW:
		return true
	}
	return false
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordTypeLoc ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordType = "LOC"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordTypeLoc:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// A valid mail server hostname.
	Content param.Field[string] `json:"content" format:"hostname"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecordTypeMx ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecordType = "MX"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecordTypeMx:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a NAPTR record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Components of a NAPTR record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordData struct {
	// Flags.
	Flags param.Field[string] `json:"flags"`
	// Order.
	Order param.Field[float64] `json:"order"`
	// Preference.
	Preference param.Field[float64] `json:"preference"`
	// Regex.
	Regex param.Field[string] `json:"regex"`
	// Replacement.
	Replacement param.Field[string] `json:"replacement"`
	// Service.
	Service param.Field[string] `json:"service"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordTypeNaptr ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordType = "NAPTR"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordTypeNaptr:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// A valid name server host name.
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecordTypeNs ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecordType = "NS"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecordTypeNs:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecordTypeOpenpgpkey ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecordType = "OPENPGPKEY"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecordTypeOpenpgpkey:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecordTypePtr ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecordType = "PTR"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecordTypePtr:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a SMIMEA record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Components of a SMIMEA record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordData struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordTypeSmimea ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordType = "SMIMEA"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordTypeSmimea:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a SRV record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Components of a SRV record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordData struct {
	// The port of the service.
	Port param.Field[float64] `json:"port"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// A valid hostname.
	Target param.Field[string] `json:"target" format:"hostname"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordTypeSrv ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordType = "SRV"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordTypeSrv:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a SSHFP record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Components of a SSHFP record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordTypeSshfp ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordType = "SSHFP"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordTypeSshfp:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a SVCB record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Components of a SVCB record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordTypeSvcb ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordType = "SVCB"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordTypeSvcb:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a TLSA record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Components of a TLSA record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordData struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordTypeTlsa ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordType = "TLSA"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordTypeTlsa:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Text content for the record. The content must consist of quoted "character
	// strings" (RFC 1035), each with a length of up to 255 bytes. Strings exceeding
	// this allowed maximum length are automatically split.
	//
	// Learn more at
	// <https://www.cloudflare.com/learning/dns/dns-records/dns-txt-record/>.
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecordTypeTxt ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecordType = "TXT"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecordTypeTxt:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a URI record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordData] `json:"data"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Components of a URI record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordData struct {
	// The record content.
	Target param.Field[string] `json:"target"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordType string

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordTypeUri ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordType = "URI"
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordTypeUri:
		return true
	}
	return false
}

// Satisfied by [ZoneDNSRecordBatchParamsPutsDNSRecordsARecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecord].
type ZoneDNSRecordBatchParamsPutUnion interface {
	implementsZoneDNSRecordBatchParamsPutUnion()
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsARecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// A valid IPv4 address.
	Content param.Field[string] `json:"content" format:"ipv4"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsARecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsARecord) implementsZoneDNSRecordBatchParamsPutUnion() {}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsARecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsARecordTypeA ZoneDNSRecordBatchParamsPutsDNSRecordsARecordType = "A"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsARecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsARecordTypeA:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// A valid IPv6 address.
	Content param.Field[string] `json:"content" format:"ipv6"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecordTypeAaaa ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecordType = "AAAA"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecordTypeAaaa:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a CAA record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Components of a CAA record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordData struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordTypeCaa ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordType = "CAA"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordTypeCaa:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a CERT record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Components of a CERT record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordTypeCert ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordType = "CERT"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordTypeCert:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// A valid hostname. Must not match the record's name.
	Content  param.Field[string]                                                    `json:"content"`
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordSettings] `json:"settings"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordSettings struct {
	// If enabled, causes the CNAME record to be resolved externally and the resulting
	// address records (e.g., A and AAAA) to be returned instead of the CNAME record
	// itself. This setting is unavailable for proxied records, since they are always
	// flattened.
	FlattenCname param.Field[bool] `json:"flatten_cname"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordTypeCname ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordType = "CNAME"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordTypeCname:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a DNSKEY record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Components of a DNSKEY record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordTypeDnskey ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordType = "DNSKEY"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordTypeDnskey:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a DS record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Components of a DS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordTypeDs ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordType = "DS"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordTypeDs:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a HTTPS record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Components of a HTTPS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordTypeHTTPS ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordType = "HTTPS"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordTypeHTTPS:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a LOC record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Components of a LOC record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordData struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordDataLongDirection] `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes param.Field[float64] `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds param.Field[float64] `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz param.Field[float64] `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert param.Field[float64] `json:"precision_vert"`
	// Size of location in meters.
	Size param.Field[float64] `json:"size"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordDataLatDirection string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordDataLatDirectionN ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordDataLatDirection = "N"
	ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordDataLatDirectionS ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordDataLatDirection = "S"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordDataLatDirection) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordDataLatDirectionN, ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordDataLatDirectionS:
		return true
	}
	return false
}

// Longitude direction.
type ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordDataLongDirection string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordDataLongDirectionE ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordDataLongDirection = "E"
	ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordDataLongDirectionW ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordDataLongDirection = "W"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordDataLongDirection) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordDataLongDirectionE, ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordDataLongDirectionW:
		return true
	}
	return false
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordTypeLoc ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordType = "LOC"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordTypeLoc:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// A valid mail server hostname.
	Content param.Field[string] `json:"content" format:"hostname"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecordTypeMx ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecordType = "MX"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecordTypeMx:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a NAPTR record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Components of a NAPTR record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordData struct {
	// Flags.
	Flags param.Field[string] `json:"flags"`
	// Order.
	Order param.Field[float64] `json:"order"`
	// Preference.
	Preference param.Field[float64] `json:"preference"`
	// Regex.
	Regex param.Field[string] `json:"regex"`
	// Replacement.
	Replacement param.Field[string] `json:"replacement"`
	// Service.
	Service param.Field[string] `json:"service"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordTypeNaptr ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordType = "NAPTR"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordTypeNaptr:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// A valid name server host name.
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecordTypeNs ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecordType = "NS"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecordTypeNs:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecordTypeOpenpgpkey ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecordType = "OPENPGPKEY"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecordTypeOpenpgpkey:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecordTypePtr ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecordType = "PTR"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecordTypePtr:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a SMIMEA record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Components of a SMIMEA record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordData struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordTypeSmimea ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordType = "SMIMEA"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordTypeSmimea:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a SRV record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Components of a SRV record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordData struct {
	// The port of the service.
	Port param.Field[float64] `json:"port"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// A valid hostname.
	Target param.Field[string] `json:"target" format:"hostname"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordTypeSrv ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordType = "SRV"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordTypeSrv:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a SSHFP record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Components of a SSHFP record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordTypeSshfp ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordType = "SSHFP"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordTypeSshfp:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a SVCB record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Components of a SVCB record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordTypeSvcb ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordType = "SVCB"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordTypeSvcb:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a TLSA record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordData] `json:"data"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Components of a TLSA record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordData struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordTypeTlsa ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordType = "TLSA"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordTypeTlsa:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Text content for the record. The content must consist of quoted "character
	// strings" (RFC 1035), each with a length of up to 255 bytes. Strings exceeding
	// this allowed maximum length are automatically split.
	//
	// Learn more at
	// <https://www.cloudflare.com/learning/dns/dns-records/dns-txt-record/>.
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecordTypeTxt ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecordType = "TXT"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecordTypeTxt:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecord struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	// Components of a URI record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordData] `json:"data"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordType] `json:"type"`
	SharedFieldsParam
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Components of a URI record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordData struct {
	// The record content.
	Target param.Field[string] `json:"target"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordType string

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordTypeUri ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordType = "URI"
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordTypeUri:
		return true
	}
	return false
}

type ZoneDNSRecordImportParams struct {
	// BIND config to import.
	//
	// **Tip:** When using cURL, a file can be uploaded using
	// `--form 'file=@bind_config.txt'`.
	File param.Field[string] `json:"file,required"`
	// Whether or not proxiable records should receive the performance and security
	// benefits of Cloudflare.
	//
	// The value should be either `true` or `false`.
	Proxied param.Field[string] `json:"proxied"`
}

func (r ZoneDNSRecordImportParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type ZoneDNSRecordOverwriteParams struct {
	DNSRecordPost DNSRecordPostUnionParam `json:"dns_record_post,required"`
}

func (r ZoneDNSRecordOverwriteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DNSRecordPost)
}

type ZoneDNSRecordScanParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneDNSRecordScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
