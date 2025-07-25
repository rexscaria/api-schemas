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

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
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
func (r *ZoneDNSRecordService) Delete(ctx context.Context, zoneID string, dnsRecordID string, opts ...option.RequestOption) (res *ZoneDNSRecordDeleteResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
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

type DNSRecordMessageItem struct {
	Code             int64                      `json:"code,required"`
	Message          string                     `json:"message,required"`
	DocumentationURL string                     `json:"documentation_url"`
	Source           DNSRecordMessageItemSource `json:"source"`
	JSON             dnsRecordMessageItemJSON   `json:"-"`
}

// dnsRecordMessageItemJSON contains the JSON metadata for the struct
// [DNSRecordMessageItem]
type dnsRecordMessageItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DNSRecordMessageItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordMessageItemJSON) RawJSON() string {
	return r.raw
}

type DNSRecordMessageItemSource struct {
	Pointer string                         `json:"pointer"`
	JSON    dnsRecordMessageItemSourceJSON `json:"-"`
}

// dnsRecordMessageItemSourceJSON contains the JSON metadata for the struct
// [DNSRecordMessageItemSource]
type dnsRecordMessageItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordMessageItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordMessageItemSourceJSON) RawJSON() string {
	return r.raw
}

type DNSRecordPatchParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string]      `json:"name,required"`
	Ttl  param.Field[interface{}] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid IPv4 address.
	Content param.Field[string]      `json:"content" format:"ipv4"`
	Data    param.Field[interface{}] `json:"data"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied  param.Field[bool]        `json:"proxied"`
	Settings param.Field[interface{}] `json:"settings"`
	Tags     param.Field[interface{}] `json:"tags"`
}

func (r DNSRecordPatchParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchParam) implementsDNSRecordPatchUnionParam() {}

// Satisfied by [DNSRecordPatchDNSRecordsARecordParam],
// [DNSRecordPatchDNSRecordsAaaaRecordParam],
// [DNSRecordPatchDNSRecordsCnameRecordParam],
// [DNSRecordPatchDNSRecordsMxRecordParam],
// [DNSRecordPatchDNSRecordsNsRecordParam],
// [DNSRecordPatchDNSRecordsOpenpgpkeyRecordParam],
// [DNSRecordPatchDNSRecordsPtrRecordParam],
// [DNSRecordPatchDNSRecordsTxtRecordParam],
// [DNSRecordPatchDNSRecordsCaaRecordParam],
// [DNSRecordPatchDNSRecordsCertRecordParam],
// [DNSRecordPatchDNSRecordsDnskeyRecordParam],
// [DNSRecordPatchDNSRecordsDsRecordParam],
// [DNSRecordPatchDNSRecordsHTTPSRecordParam],
// [DNSRecordPatchDNSRecordsLocRecordParam],
// [DNSRecordPatchDNSRecordsNaptrRecordParam],
// [DNSRecordPatchDNSRecordsSmimeaRecordParam],
// [DNSRecordPatchDNSRecordsSrvRecordParam],
// [DNSRecordPatchDNSRecordsSshfpRecordParam],
// [DNSRecordPatchDNSRecordsSvcbRecordParam],
// [DNSRecordPatchDNSRecordsTlsaRecordParam],
// [DNSRecordPatchDNSRecordsUriRecordParam], [DNSRecordPatchParam].
type DNSRecordPatchUnionParam interface {
	implementsDNSRecordPatchUnionParam()
}

type DNSRecordPatchDNSRecordsARecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsARecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid IPv4 address.
	Content param.Field[string] `json:"content" format:"ipv4"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsARecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsARecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsARecordTtl float64

const (
	DNSRecordPatchDNSRecordsARecordTtl1 DNSRecordPatchDNSRecordsARecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsARecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsARecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsARecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsARecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsAaaaRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsAaaaRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsAaaaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid IPv6 address.
	Content param.Field[string] `json:"content" format:"ipv6"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsAaaaRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsAaaaRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsAaaaRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsAaaaRecordTtl float64

const (
	DNSRecordPatchDNSRecordsAaaaRecordTtl1 DNSRecordPatchDNSRecordsAaaaRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsAaaaRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsAaaaRecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsAaaaRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsAaaaRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsCnameRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsCnameRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsCnameRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid hostname. Must not match the record's name.
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsCnameRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsCnameRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsCnameRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsCnameRecordTtl float64

const (
	DNSRecordPatchDNSRecordsCnameRecordTtl1 DNSRecordPatchDNSRecordsCnameRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsCnameRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsCnameRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsCnameRecordSettingsParam struct {
	// If enabled, causes the CNAME record to be resolved externally and the resulting
	// address records (e.g., A and AAAA) to be returned instead of the CNAME record
	// itself. This setting is unavailable for proxied records, since they are always
	// flattened.
	FlattenCname param.Field[bool] `json:"flatten_cname"`
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

func (r DNSRecordPatchDNSRecordsCnameRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsMxRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsMxRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsMxRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid mail server hostname.
	Content param.Field[string] `json:"content" format:"hostname"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsMxRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsMxRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsMxRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsMxRecordTtl float64

const (
	DNSRecordPatchDNSRecordsMxRecordTtl1 DNSRecordPatchDNSRecordsMxRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsMxRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsMxRecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsMxRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsMxRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsNsRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsNsRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsNsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid name server host name.
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsNsRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsNsRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsNsRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsNsRecordTtl float64

const (
	DNSRecordPatchDNSRecordsNsRecordTtl1 DNSRecordPatchDNSRecordsNsRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsNsRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsNsRecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsNsRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsNsRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsOpenpgpkeyRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsOpenpgpkeyRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsOpenpgpkeyRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsOpenpgpkeyRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsOpenpgpkeyRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsOpenpgpkeyRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsOpenpgpkeyRecordTtl float64

const (
	DNSRecordPatchDNSRecordsOpenpgpkeyRecordTtl1 DNSRecordPatchDNSRecordsOpenpgpkeyRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsOpenpgpkeyRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsOpenpgpkeyRecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsOpenpgpkeyRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsOpenpgpkeyRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsPtrRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsPtrRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsPtrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsPtrRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsPtrRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsPtrRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsPtrRecordTtl float64

const (
	DNSRecordPatchDNSRecordsPtrRecordTtl1 DNSRecordPatchDNSRecordsPtrRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsPtrRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsPtrRecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsPtrRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsPtrRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsTxtRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsTxtRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsTxtRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Text content for the record. The content must consist of quoted "character
	// strings" (RFC 1035), each with a length of up to 255 bytes. Strings exceeding
	// this allowed maximum length are automatically split.
	//
	// Learn more at
	// <https://www.cloudflare.com/learning/dns/dns-records/dns-txt-record/>.
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsTxtRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsTxtRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsTxtRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsTxtRecordTtl float64

const (
	DNSRecordPatchDNSRecordsTxtRecordTtl1 DNSRecordPatchDNSRecordsTxtRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsTxtRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsTxtRecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsTxtRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsTxtRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsCaaRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsCaaRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsCaaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a CAA record.
	Data param.Field[DNSRecordPatchDNSRecordsCaaRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsCaaRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsCaaRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsCaaRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsCaaRecordTtl float64

const (
	DNSRecordPatchDNSRecordsCaaRecordTtl1 DNSRecordPatchDNSRecordsCaaRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsCaaRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsCaaRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsCaaRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsCaaRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsCertRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsCertRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsCertRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a CERT record.
	Data param.Field[DNSRecordPatchDNSRecordsCertRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsCertRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsCertRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsCertRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsCertRecordTtl float64

const (
	DNSRecordPatchDNSRecordsCertRecordTtl1 DNSRecordPatchDNSRecordsCertRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsCertRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsCertRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsCertRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsCertRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsDnskeyRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsDnskeyRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsDnskeyRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a DNSKEY record.
	Data param.Field[DNSRecordPatchDNSRecordsDnskeyRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsDnskeyRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsDnskeyRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsDnskeyRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsDnskeyRecordTtl float64

const (
	DNSRecordPatchDNSRecordsDnskeyRecordTtl1 DNSRecordPatchDNSRecordsDnskeyRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsDnskeyRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsDnskeyRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsDnskeyRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsDnskeyRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsDsRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsDsRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsDsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a DS record.
	Data param.Field[DNSRecordPatchDNSRecordsDsRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsDsRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsDsRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsDsRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsDsRecordTtl float64

const (
	DNSRecordPatchDNSRecordsDsRecordTtl1 DNSRecordPatchDNSRecordsDsRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsDsRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsDsRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsDsRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsDsRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsHTTPSRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsHTTPSRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsHTTPSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a HTTPS record.
	Data param.Field[DNSRecordPatchDNSRecordsHTTPSRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsHTTPSRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsHTTPSRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsHTTPSRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsHTTPSRecordTtl float64

const (
	DNSRecordPatchDNSRecordsHTTPSRecordTtl1 DNSRecordPatchDNSRecordsHTTPSRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsHTTPSRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsHTTPSRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsHTTPSRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsHTTPSRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsLocRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsLocRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsLocRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a LOC record.
	Data param.Field[DNSRecordPatchDNSRecordsLocRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsLocRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsLocRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsLocRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsLocRecordTtl float64

const (
	DNSRecordPatchDNSRecordsLocRecordTtl1 DNSRecordPatchDNSRecordsLocRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsLocRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsLocRecordTtl1:
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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsLocRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsLocRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsNaptrRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsNaptrRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsNaptrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a NAPTR record.
	Data param.Field[DNSRecordPatchDNSRecordsNaptrRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsNaptrRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsNaptrRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsNaptrRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsNaptrRecordTtl float64

const (
	DNSRecordPatchDNSRecordsNaptrRecordTtl1 DNSRecordPatchDNSRecordsNaptrRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsNaptrRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsNaptrRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsNaptrRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsNaptrRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsSmimeaRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsSmimeaRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsSmimeaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SMIMEA record.
	Data param.Field[DNSRecordPatchDNSRecordsSmimeaRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsSmimeaRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsSmimeaRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsSmimeaRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsSmimeaRecordTtl float64

const (
	DNSRecordPatchDNSRecordsSmimeaRecordTtl1 DNSRecordPatchDNSRecordsSmimeaRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsSmimeaRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsSmimeaRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsSmimeaRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsSmimeaRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsSrvRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsSrvRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsSrvRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SRV record.
	Data param.Field[DNSRecordPatchDNSRecordsSrvRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsSrvRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsSrvRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsSrvRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsSrvRecordTtl float64

const (
	DNSRecordPatchDNSRecordsSrvRecordTtl1 DNSRecordPatchDNSRecordsSrvRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsSrvRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsSrvRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsSrvRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsSrvRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsSshfpRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsSshfpRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsSshfpRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SSHFP record.
	Data param.Field[DNSRecordPatchDNSRecordsSshfpRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsSshfpRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsSshfpRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsSshfpRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsSshfpRecordTtl float64

const (
	DNSRecordPatchDNSRecordsSshfpRecordTtl1 DNSRecordPatchDNSRecordsSshfpRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsSshfpRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsSshfpRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsSshfpRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsSshfpRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsSvcbRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsSvcbRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsSvcbRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SVCB record.
	Data param.Field[DNSRecordPatchDNSRecordsSvcbRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsSvcbRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsSvcbRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsSvcbRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsSvcbRecordTtl float64

const (
	DNSRecordPatchDNSRecordsSvcbRecordTtl1 DNSRecordPatchDNSRecordsSvcbRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsSvcbRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsSvcbRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsSvcbRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsSvcbRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsTlsaRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsTlsaRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsTlsaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a TLSA record.
	Data param.Field[DNSRecordPatchDNSRecordsTlsaRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsTlsaRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsTlsaRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsTlsaRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsTlsaRecordTtl float64

const (
	DNSRecordPatchDNSRecordsTlsaRecordTtl1 DNSRecordPatchDNSRecordsTlsaRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsTlsaRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsTlsaRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsTlsaRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsTlsaRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPatchDNSRecordsUriRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPatchDNSRecordsUriRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPatchDNSRecordsUriRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a URI record.
	Data param.Field[DNSRecordPatchDNSRecordsUriRecordDataParam] `json:"data"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPatchDNSRecordsUriRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPatchDNSRecordsUriRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPatchDNSRecordsUriRecordParam) implementsDNSRecordPatchUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPatchDNSRecordsUriRecordTtl float64

const (
	DNSRecordPatchDNSRecordsUriRecordTtl1 DNSRecordPatchDNSRecordsUriRecordTtl = 1
)

func (r DNSRecordPatchDNSRecordsUriRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPatchDNSRecordsUriRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPatchDNSRecordsUriRecordSettingsParam struct {
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

func (r DNSRecordPatchDNSRecordsUriRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPatchType string

const (
	DNSRecordPatchTypeA          DNSRecordPatchType = "A"
	DNSRecordPatchTypeAaaa       DNSRecordPatchType = "AAAA"
	DNSRecordPatchTypeCname      DNSRecordPatchType = "CNAME"
	DNSRecordPatchTypeMx         DNSRecordPatchType = "MX"
	DNSRecordPatchTypeNs         DNSRecordPatchType = "NS"
	DNSRecordPatchTypeOpenpgpkey DNSRecordPatchType = "OPENPGPKEY"
	DNSRecordPatchTypePtr        DNSRecordPatchType = "PTR"
	DNSRecordPatchTypeTxt        DNSRecordPatchType = "TXT"
	DNSRecordPatchTypeCaa        DNSRecordPatchType = "CAA"
	DNSRecordPatchTypeCert       DNSRecordPatchType = "CERT"
	DNSRecordPatchTypeDnskey     DNSRecordPatchType = "DNSKEY"
	DNSRecordPatchTypeDs         DNSRecordPatchType = "DS"
	DNSRecordPatchTypeHTTPS      DNSRecordPatchType = "HTTPS"
	DNSRecordPatchTypeLoc        DNSRecordPatchType = "LOC"
	DNSRecordPatchTypeNaptr      DNSRecordPatchType = "NAPTR"
	DNSRecordPatchTypeSmimea     DNSRecordPatchType = "SMIMEA"
	DNSRecordPatchTypeSrv        DNSRecordPatchType = "SRV"
	DNSRecordPatchTypeSshfp      DNSRecordPatchType = "SSHFP"
	DNSRecordPatchTypeSvcb       DNSRecordPatchType = "SVCB"
	DNSRecordPatchTypeTlsa       DNSRecordPatchType = "TLSA"
	DNSRecordPatchTypeUri        DNSRecordPatchType = "URI"
)

func (r DNSRecordPatchType) IsKnown() bool {
	switch r {
	case DNSRecordPatchTypeA, DNSRecordPatchTypeAaaa, DNSRecordPatchTypeCname, DNSRecordPatchTypeMx, DNSRecordPatchTypeNs, DNSRecordPatchTypeOpenpgpkey, DNSRecordPatchTypePtr, DNSRecordPatchTypeTxt, DNSRecordPatchTypeCaa, DNSRecordPatchTypeCert, DNSRecordPatchTypeDnskey, DNSRecordPatchTypeDs, DNSRecordPatchTypeHTTPS, DNSRecordPatchTypeLoc, DNSRecordPatchTypeNaptr, DNSRecordPatchTypeSmimea, DNSRecordPatchTypeSrv, DNSRecordPatchTypeSshfp, DNSRecordPatchTypeSvcb, DNSRecordPatchTypeTlsa, DNSRecordPatchTypeUri:
		return true
	}
	return false
}

type DNSRecordPostParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string]      `json:"name,required"`
	Ttl  param.Field[interface{}] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid IPv4 address.
	Content param.Field[string]      `json:"content" format:"ipv4"`
	Data    param.Field[interface{}] `json:"data"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied  param.Field[bool]        `json:"proxied"`
	Settings param.Field[interface{}] `json:"settings"`
	Tags     param.Field[interface{}] `json:"tags"`
}

func (r DNSRecordPostParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostParam) implementsDNSRecordPostUnionParam() {}

// Satisfied by [DNSRecordPostDNSRecordsARecordParam],
// [DNSRecordPostDNSRecordsAaaaRecordParam],
// [DNSRecordPostDNSRecordsCnameRecordParam],
// [DNSRecordPostDNSRecordsMxRecordParam], [DNSRecordPostDNSRecordsNsRecordParam],
// [DNSRecordPostDNSRecordsOpenpgpkeyRecordParam],
// [DNSRecordPostDNSRecordsPtrRecordParam],
// [DNSRecordPostDNSRecordsTxtRecordParam],
// [DNSRecordPostDNSRecordsCaaRecordParam],
// [DNSRecordPostDNSRecordsCertRecordParam],
// [DNSRecordPostDNSRecordsDnskeyRecordParam],
// [DNSRecordPostDNSRecordsDsRecordParam],
// [DNSRecordPostDNSRecordsHTTPSRecordParam],
// [DNSRecordPostDNSRecordsLocRecordParam],
// [DNSRecordPostDNSRecordsNaptrRecordParam],
// [DNSRecordPostDNSRecordsSmimeaRecordParam],
// [DNSRecordPostDNSRecordsSrvRecordParam],
// [DNSRecordPostDNSRecordsSshfpRecordParam],
// [DNSRecordPostDNSRecordsSvcbRecordParam],
// [DNSRecordPostDNSRecordsTlsaRecordParam],
// [DNSRecordPostDNSRecordsUriRecordParam], [DNSRecordPostParam].
type DNSRecordPostUnionParam interface {
	implementsDNSRecordPostUnionParam()
}

type DNSRecordPostDNSRecordsARecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsARecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid IPv4 address.
	Content param.Field[string] `json:"content" format:"ipv4"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsARecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsARecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsARecordTtl float64

const (
	DNSRecordPostDNSRecordsARecordTtl1 DNSRecordPostDNSRecordsARecordTtl = 1
)

func (r DNSRecordPostDNSRecordsARecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsARecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsARecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsARecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsAaaaRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsAaaaRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsAaaaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid IPv6 address.
	Content param.Field[string] `json:"content" format:"ipv6"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsAaaaRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsAaaaRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsAaaaRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsAaaaRecordTtl float64

const (
	DNSRecordPostDNSRecordsAaaaRecordTtl1 DNSRecordPostDNSRecordsAaaaRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsAaaaRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsAaaaRecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsAaaaRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsAaaaRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsCnameRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsCnameRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsCnameRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid hostname. Must not match the record's name.
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsCnameRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsCnameRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsCnameRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsCnameRecordTtl float64

const (
	DNSRecordPostDNSRecordsCnameRecordTtl1 DNSRecordPostDNSRecordsCnameRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsCnameRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsCnameRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsCnameRecordSettingsParam struct {
	// If enabled, causes the CNAME record to be resolved externally and the resulting
	// address records (e.g., A and AAAA) to be returned instead of the CNAME record
	// itself. This setting is unavailable for proxied records, since they are always
	// flattened.
	FlattenCname param.Field[bool] `json:"flatten_cname"`
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

func (r DNSRecordPostDNSRecordsCnameRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsMxRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsMxRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsMxRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid mail server hostname.
	Content param.Field[string] `json:"content" format:"hostname"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsMxRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsMxRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsMxRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsMxRecordTtl float64

const (
	DNSRecordPostDNSRecordsMxRecordTtl1 DNSRecordPostDNSRecordsMxRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsMxRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsMxRecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsMxRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsMxRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsNsRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsNsRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsNsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid name server host name.
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsNsRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsNsRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsNsRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsNsRecordTtl float64

const (
	DNSRecordPostDNSRecordsNsRecordTtl1 DNSRecordPostDNSRecordsNsRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsNsRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsNsRecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsNsRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsNsRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsOpenpgpkeyRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsOpenpgpkeyRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsOpenpgpkeyRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsOpenpgpkeyRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsOpenpgpkeyRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsOpenpgpkeyRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsOpenpgpkeyRecordTtl float64

const (
	DNSRecordPostDNSRecordsOpenpgpkeyRecordTtl1 DNSRecordPostDNSRecordsOpenpgpkeyRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsOpenpgpkeyRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsOpenpgpkeyRecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsOpenpgpkeyRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsOpenpgpkeyRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsPtrRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsPtrRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsPtrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsPtrRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsPtrRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsPtrRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsPtrRecordTtl float64

const (
	DNSRecordPostDNSRecordsPtrRecordTtl1 DNSRecordPostDNSRecordsPtrRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsPtrRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsPtrRecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsPtrRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsPtrRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsTxtRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsTxtRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsTxtRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Text content for the record. The content must consist of quoted "character
	// strings" (RFC 1035), each with a length of up to 255 bytes. Strings exceeding
	// this allowed maximum length are automatically split.
	//
	// Learn more at
	// <https://www.cloudflare.com/learning/dns/dns-records/dns-txt-record/>.
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsTxtRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsTxtRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsTxtRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsTxtRecordTtl float64

const (
	DNSRecordPostDNSRecordsTxtRecordTtl1 DNSRecordPostDNSRecordsTxtRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsTxtRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsTxtRecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsTxtRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsTxtRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsCaaRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsCaaRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsCaaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a CAA record.
	Data param.Field[DNSRecordPostDNSRecordsCaaRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsCaaRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsCaaRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsCaaRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsCaaRecordTtl float64

const (
	DNSRecordPostDNSRecordsCaaRecordTtl1 DNSRecordPostDNSRecordsCaaRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsCaaRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsCaaRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsCaaRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsCaaRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsCertRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsCertRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsCertRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a CERT record.
	Data param.Field[DNSRecordPostDNSRecordsCertRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsCertRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsCertRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsCertRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsCertRecordTtl float64

const (
	DNSRecordPostDNSRecordsCertRecordTtl1 DNSRecordPostDNSRecordsCertRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsCertRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsCertRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsCertRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsCertRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsDnskeyRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsDnskeyRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsDnskeyRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a DNSKEY record.
	Data param.Field[DNSRecordPostDNSRecordsDnskeyRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsDnskeyRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsDnskeyRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsDnskeyRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsDnskeyRecordTtl float64

const (
	DNSRecordPostDNSRecordsDnskeyRecordTtl1 DNSRecordPostDNSRecordsDnskeyRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsDnskeyRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsDnskeyRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsDnskeyRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsDnskeyRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsDsRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsDsRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsDsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a DS record.
	Data param.Field[DNSRecordPostDNSRecordsDsRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsDsRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsDsRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsDsRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsDsRecordTtl float64

const (
	DNSRecordPostDNSRecordsDsRecordTtl1 DNSRecordPostDNSRecordsDsRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsDsRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsDsRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsDsRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsDsRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsHTTPSRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsHTTPSRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsHTTPSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a HTTPS record.
	Data param.Field[DNSRecordPostDNSRecordsHTTPSRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsHTTPSRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsHTTPSRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsHTTPSRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsHTTPSRecordTtl float64

const (
	DNSRecordPostDNSRecordsHTTPSRecordTtl1 DNSRecordPostDNSRecordsHTTPSRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsHTTPSRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsHTTPSRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsHTTPSRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsHTTPSRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsLocRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsLocRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsLocRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a LOC record.
	Data param.Field[DNSRecordPostDNSRecordsLocRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsLocRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsLocRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsLocRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsLocRecordTtl float64

const (
	DNSRecordPostDNSRecordsLocRecordTtl1 DNSRecordPostDNSRecordsLocRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsLocRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsLocRecordTtl1:
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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsLocRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsLocRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsNaptrRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsNaptrRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsNaptrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a NAPTR record.
	Data param.Field[DNSRecordPostDNSRecordsNaptrRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsNaptrRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsNaptrRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsNaptrRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsNaptrRecordTtl float64

const (
	DNSRecordPostDNSRecordsNaptrRecordTtl1 DNSRecordPostDNSRecordsNaptrRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsNaptrRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsNaptrRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsNaptrRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsNaptrRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsSmimeaRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsSmimeaRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsSmimeaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SMIMEA record.
	Data param.Field[DNSRecordPostDNSRecordsSmimeaRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsSmimeaRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsSmimeaRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsSmimeaRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsSmimeaRecordTtl float64

const (
	DNSRecordPostDNSRecordsSmimeaRecordTtl1 DNSRecordPostDNSRecordsSmimeaRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsSmimeaRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsSmimeaRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsSmimeaRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsSmimeaRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsSrvRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsSrvRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsSrvRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SRV record.
	Data param.Field[DNSRecordPostDNSRecordsSrvRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsSrvRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsSrvRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsSrvRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsSrvRecordTtl float64

const (
	DNSRecordPostDNSRecordsSrvRecordTtl1 DNSRecordPostDNSRecordsSrvRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsSrvRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsSrvRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsSrvRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsSrvRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsSshfpRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsSshfpRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsSshfpRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SSHFP record.
	Data param.Field[DNSRecordPostDNSRecordsSshfpRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsSshfpRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsSshfpRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsSshfpRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsSshfpRecordTtl float64

const (
	DNSRecordPostDNSRecordsSshfpRecordTtl1 DNSRecordPostDNSRecordsSshfpRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsSshfpRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsSshfpRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsSshfpRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsSshfpRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsSvcbRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsSvcbRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsSvcbRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SVCB record.
	Data param.Field[DNSRecordPostDNSRecordsSvcbRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsSvcbRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsSvcbRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsSvcbRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsSvcbRecordTtl float64

const (
	DNSRecordPostDNSRecordsSvcbRecordTtl1 DNSRecordPostDNSRecordsSvcbRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsSvcbRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsSvcbRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsSvcbRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsSvcbRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsTlsaRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsTlsaRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsTlsaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a TLSA record.
	Data param.Field[DNSRecordPostDNSRecordsTlsaRecordDataParam] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsTlsaRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsTlsaRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsTlsaRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsTlsaRecordTtl float64

const (
	DNSRecordPostDNSRecordsTlsaRecordTtl1 DNSRecordPostDNSRecordsTlsaRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsTlsaRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsTlsaRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsTlsaRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsTlsaRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordPostDNSRecordsUriRecordParam struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[DNSRecordPostDNSRecordsUriRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[DNSRecordPostDNSRecordsUriRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a URI record.
	Data param.Field[DNSRecordPostDNSRecordsUriRecordDataParam] `json:"data"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[DNSRecordPostDNSRecordsUriRecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r DNSRecordPostDNSRecordsUriRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSRecordPostDNSRecordsUriRecordParam) implementsDNSRecordPostUnionParam() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordPostDNSRecordsUriRecordTtl float64

const (
	DNSRecordPostDNSRecordsUriRecordTtl1 DNSRecordPostDNSRecordsUriRecordTtl = 1
)

func (r DNSRecordPostDNSRecordsUriRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordPostDNSRecordsUriRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordPostDNSRecordsUriRecordSettingsParam struct {
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

func (r DNSRecordPostDNSRecordsUriRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordPostType string

const (
	DNSRecordPostTypeA          DNSRecordPostType = "A"
	DNSRecordPostTypeAaaa       DNSRecordPostType = "AAAA"
	DNSRecordPostTypeCname      DNSRecordPostType = "CNAME"
	DNSRecordPostTypeMx         DNSRecordPostType = "MX"
	DNSRecordPostTypeNs         DNSRecordPostType = "NS"
	DNSRecordPostTypeOpenpgpkey DNSRecordPostType = "OPENPGPKEY"
	DNSRecordPostTypePtr        DNSRecordPostType = "PTR"
	DNSRecordPostTypeTxt        DNSRecordPostType = "TXT"
	DNSRecordPostTypeCaa        DNSRecordPostType = "CAA"
	DNSRecordPostTypeCert       DNSRecordPostType = "CERT"
	DNSRecordPostTypeDnskey     DNSRecordPostType = "DNSKEY"
	DNSRecordPostTypeDs         DNSRecordPostType = "DS"
	DNSRecordPostTypeHTTPS      DNSRecordPostType = "HTTPS"
	DNSRecordPostTypeLoc        DNSRecordPostType = "LOC"
	DNSRecordPostTypeNaptr      DNSRecordPostType = "NAPTR"
	DNSRecordPostTypeSmimea     DNSRecordPostType = "SMIMEA"
	DNSRecordPostTypeSrv        DNSRecordPostType = "SRV"
	DNSRecordPostTypeSshfp      DNSRecordPostType = "SSHFP"
	DNSRecordPostTypeSvcb       DNSRecordPostType = "SVCB"
	DNSRecordPostTypeTlsa       DNSRecordPostType = "TLSA"
	DNSRecordPostTypeUri        DNSRecordPostType = "URI"
)

func (r DNSRecordPostType) IsKnown() bool {
	switch r {
	case DNSRecordPostTypeA, DNSRecordPostTypeAaaa, DNSRecordPostTypeCname, DNSRecordPostTypeMx, DNSRecordPostTypeNs, DNSRecordPostTypeOpenpgpkey, DNSRecordPostTypePtr, DNSRecordPostTypeTxt, DNSRecordPostTypeCaa, DNSRecordPostTypeCert, DNSRecordPostTypeDnskey, DNSRecordPostTypeDs, DNSRecordPostTypeHTTPS, DNSRecordPostTypeLoc, DNSRecordPostTypeNaptr, DNSRecordPostTypeSmimea, DNSRecordPostTypeSrv, DNSRecordPostTypeSshfp, DNSRecordPostTypeSvcb, DNSRecordPostTypeTlsa, DNSRecordPostTypeUri:
		return true
	}
	return false
}

type DNSRecordResponse struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// This field can have the runtime type of [DNSRecordResponseDNSRecordsARecordTtl],
	// [DNSRecordResponseDNSRecordsAaaaRecordTtl],
	// [DNSRecordResponseDNSRecordsCnameRecordTtl],
	// [DNSRecordResponseDNSRecordsMxRecordTtl],
	// [DNSRecordResponseDNSRecordsNsRecordTtl],
	// [DNSRecordResponseDNSRecordsOpenpgpkeyRecordTtl],
	// [DNSRecordResponseDNSRecordsPtrRecordTtl],
	// [DNSRecordResponseDNSRecordsTxtRecordTtl],
	// [DNSRecordResponseDNSRecordsCaaRecordTtl],
	// [DNSRecordResponseDNSRecordsCertRecordTtl],
	// [DNSRecordResponseDNSRecordsDnskeyRecordTtl],
	// [DNSRecordResponseDNSRecordsDsRecordTtl],
	// [DNSRecordResponseDNSRecordsHTTPSRecordTtl],
	// [DNSRecordResponseDNSRecordsLocRecordTtl],
	// [DNSRecordResponseDNSRecordsNaptrRecordTtl],
	// [DNSRecordResponseDNSRecordsSmimeaRecordTtl],
	// [DNSRecordResponseDNSRecordsSrvRecordTtl],
	// [DNSRecordResponseDNSRecordsSshfpRecordTtl],
	// [DNSRecordResponseDNSRecordsSvcbRecordTtl],
	// [DNSRecordResponseDNSRecordsTlsaRecordTtl],
	// [DNSRecordResponseDNSRecordsUriRecordTtl].
	Ttl interface{} `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseType `json:"type,required"`
	// Identifier.
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
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// This field can have the runtime type of
	// [DNSRecordResponseDNSRecordsARecordSettings],
	// [DNSRecordResponseDNSRecordsAaaaRecordSettings],
	// [DNSRecordResponseDNSRecordsCnameRecordSettings],
	// [DNSRecordResponseDNSRecordsMxRecordSettings],
	// [DNSRecordResponseDNSRecordsNsRecordSettings],
	// [DNSRecordResponseDNSRecordsOpenpgpkeyRecordSettings],
	// [DNSRecordResponseDNSRecordsPtrRecordSettings],
	// [DNSRecordResponseDNSRecordsTxtRecordSettings],
	// [DNSRecordResponseDNSRecordsCaaRecordSettings],
	// [DNSRecordResponseDNSRecordsCertRecordSettings],
	// [DNSRecordResponseDNSRecordsDnskeyRecordSettings],
	// [DNSRecordResponseDNSRecordsDsRecordSettings],
	// [DNSRecordResponseDNSRecordsHTTPSRecordSettings],
	// [DNSRecordResponseDNSRecordsLocRecordSettings],
	// [DNSRecordResponseDNSRecordsNaptrRecordSettings],
	// [DNSRecordResponseDNSRecordsSmimeaRecordSettings],
	// [DNSRecordResponseDNSRecordsSrvRecordSettings],
	// [DNSRecordResponseDNSRecordsSshfpRecordSettings],
	// [DNSRecordResponseDNSRecordsSvcbRecordSettings],
	// [DNSRecordResponseDNSRecordsTlsaRecordSettings],
	// [DNSRecordResponseDNSRecordsUriRecordSettings].
	Settings interface{} `json:"settings"`
	// This field can have the runtime type of [[]string].
	Tags interface{} `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time             `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseJSON `json:"-"`
	union          DNSRecordResponseUnion
}

// dnsRecordResponseJSON contains the JSON metadata for the struct
// [DNSRecordResponse]
type dnsRecordResponseJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Priority          apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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
// [DNSRecordResponseDNSRecordsAaaaRecord],
// [DNSRecordResponseDNSRecordsCnameRecord], [DNSRecordResponseDNSRecordsMxRecord],
// [DNSRecordResponseDNSRecordsNsRecord],
// [DNSRecordResponseDNSRecordsOpenpgpkeyRecord],
// [DNSRecordResponseDNSRecordsPtrRecord], [DNSRecordResponseDNSRecordsTxtRecord],
// [DNSRecordResponseDNSRecordsCaaRecord], [DNSRecordResponseDNSRecordsCertRecord],
// [DNSRecordResponseDNSRecordsDnskeyRecord],
// [DNSRecordResponseDNSRecordsDsRecord], [DNSRecordResponseDNSRecordsHTTPSRecord],
// [DNSRecordResponseDNSRecordsLocRecord],
// [DNSRecordResponseDNSRecordsNaptrRecord],
// [DNSRecordResponseDNSRecordsSmimeaRecord],
// [DNSRecordResponseDNSRecordsSrvRecord],
// [DNSRecordResponseDNSRecordsSshfpRecord],
// [DNSRecordResponseDNSRecordsSvcbRecord],
// [DNSRecordResponseDNSRecordsTlsaRecord], [DNSRecordResponseDNSRecordsUriRecord].
func (r DNSRecordResponse) AsUnion() DNSRecordResponseUnion {
	return r.union
}

// Union satisfied by [DNSRecordResponseDNSRecordsARecord],
// [DNSRecordResponseDNSRecordsAaaaRecord],
// [DNSRecordResponseDNSRecordsCnameRecord], [DNSRecordResponseDNSRecordsMxRecord],
// [DNSRecordResponseDNSRecordsNsRecord],
// [DNSRecordResponseDNSRecordsOpenpgpkeyRecord],
// [DNSRecordResponseDNSRecordsPtrRecord], [DNSRecordResponseDNSRecordsTxtRecord],
// [DNSRecordResponseDNSRecordsCaaRecord], [DNSRecordResponseDNSRecordsCertRecord],
// [DNSRecordResponseDNSRecordsDnskeyRecord],
// [DNSRecordResponseDNSRecordsDsRecord], [DNSRecordResponseDNSRecordsHTTPSRecord],
// [DNSRecordResponseDNSRecordsLocRecord],
// [DNSRecordResponseDNSRecordsNaptrRecord],
// [DNSRecordResponseDNSRecordsSmimeaRecord],
// [DNSRecordResponseDNSRecordsSrvRecord],
// [DNSRecordResponseDNSRecordsSshfpRecord],
// [DNSRecordResponseDNSRecordsSvcbRecord], [DNSRecordResponseDNSRecordsTlsaRecord]
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
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsCnameRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsMxRecord{}),
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
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsTxtRecord{}),
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
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsNaptrRecord{}),
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
			Type:       reflect.TypeOf(DNSRecordResponseDNSRecordsUriRecord{}),
		},
	)
}

type DNSRecordResponseDNSRecordsARecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsARecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsARecordType `json:"type,required"`
	// Identifier.
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
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsARecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                              `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsARecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsARecordJSON contains the JSON metadata for the struct
// [DNSRecordResponseDNSRecordsARecord]
type dnsRecordResponseDNSRecordsARecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsARecordTtl float64

const (
	DNSRecordResponseDNSRecordsARecordTtl1 DNSRecordResponseDNSRecordsARecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsARecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsARecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsARecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                           `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsARecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsARecordSettingsJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsARecordSettings]
type dnsRecordResponseDNSRecordsARecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsARecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsARecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsAaaaRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsAaaaRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsAaaaRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsAaaaRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                 `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsAaaaRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsAaaaRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsAaaaRecord]
type dnsRecordResponseDNSRecordsAaaaRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsAaaaRecordTtl float64

const (
	DNSRecordResponseDNSRecordsAaaaRecordTtl1 DNSRecordResponseDNSRecordsAaaaRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsAaaaRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsAaaaRecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsAaaaRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                              `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsAaaaRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsAaaaRecordSettingsJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsAaaaRecordSettings]
type dnsRecordResponseDNSRecordsAaaaRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsAaaaRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsAaaaRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsCnameRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsCnameRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsCnameRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	Proxiable bool `json:"proxiable"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsCnameRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                  `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsCnameRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsCnameRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsCnameRecord]
type dnsRecordResponseDNSRecordsCnameRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsCnameRecordTtl float64

const (
	DNSRecordResponseDNSRecordsCnameRecordTtl1 DNSRecordResponseDNSRecordsCnameRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsCnameRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsCnameRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsCnameRecordSettings struct {
	// If enabled, causes the CNAME record to be resolved externally and the resulting
	// address records (e.g., A and AAAA) to be returned instead of the CNAME record
	// itself. This setting is unavailable for proxied records, since they are always
	// flattened.
	FlattenCname bool `json:"flatten_cname"`
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                               `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsCnameRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsCnameRecordSettingsJSON contains the JSON metadata
// for the struct [DNSRecordResponseDNSRecordsCnameRecordSettings]
type dnsRecordResponseDNSRecordsCnameRecordSettingsJSON struct {
	FlattenCname apijson.Field
	Ipv4Only     apijson.Field
	Ipv6Only     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsCnameRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsCnameRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsMxRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsMxRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsMxRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsMxRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsMxRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsMxRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsMxRecord]
type dnsRecordResponseDNSRecordsMxRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Priority          apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsMxRecordTtl float64

const (
	DNSRecordResponseDNSRecordsMxRecordTtl1 DNSRecordResponseDNSRecordsMxRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsMxRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsMxRecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsMxRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                            `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsMxRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsMxRecordSettingsJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsMxRecordSettings]
type dnsRecordResponseDNSRecordsMxRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsMxRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsMxRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsNsRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsNsRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsNsRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsNsRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsNsRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsNsRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsNsRecord]
type dnsRecordResponseDNSRecordsNsRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsNsRecordTtl float64

const (
	DNSRecordResponseDNSRecordsNsRecordTtl1 DNSRecordResponseDNSRecordsNsRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsNsRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsNsRecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsNsRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                            `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsNsRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsNsRecordSettingsJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsNsRecordSettings]
type dnsRecordResponseDNSRecordsNsRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsNsRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsNsRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsOpenpgpkeyRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsOpenpgpkeyRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsOpenpgpkeyRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsOpenpgpkeyRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                       `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsOpenpgpkeyRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsOpenpgpkeyRecordJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsOpenpgpkeyRecord]
type dnsRecordResponseDNSRecordsOpenpgpkeyRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsOpenpgpkeyRecordTtl float64

const (
	DNSRecordResponseDNSRecordsOpenpgpkeyRecordTtl1 DNSRecordResponseDNSRecordsOpenpgpkeyRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsOpenpgpkeyRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsOpenpgpkeyRecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsOpenpgpkeyRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                                    `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsOpenpgpkeyRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsOpenpgpkeyRecordSettingsJSON contains the JSON
// metadata for the struct [DNSRecordResponseDNSRecordsOpenpgpkeyRecordSettings]
type dnsRecordResponseDNSRecordsOpenpgpkeyRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsOpenpgpkeyRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsOpenpgpkeyRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsPtrRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsPtrRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsPtrRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsPtrRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsPtrRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsPtrRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsPtrRecord]
type dnsRecordResponseDNSRecordsPtrRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsPtrRecordTtl float64

const (
	DNSRecordResponseDNSRecordsPtrRecordTtl1 DNSRecordResponseDNSRecordsPtrRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsPtrRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsPtrRecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsPtrRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                             `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsPtrRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsPtrRecordSettingsJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsPtrRecordSettings]
type dnsRecordResponseDNSRecordsPtrRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsPtrRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsPtrRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsTxtRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsTxtRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsTxtRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsTxtRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsTxtRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsTxtRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsTxtRecord]
type dnsRecordResponseDNSRecordsTxtRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsTxtRecordTtl float64

const (
	DNSRecordResponseDNSRecordsTxtRecordTtl1 DNSRecordResponseDNSRecordsTxtRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsTxtRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsTxtRecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsTxtRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                             `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsTxtRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsTxtRecordSettingsJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsTxtRecordSettings]
type dnsRecordResponseDNSRecordsTxtRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsTxtRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsTxtRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsCaaRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsCaaRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsCaaRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsCaaRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsCaaRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsCaaRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsCaaRecord]
type dnsRecordResponseDNSRecordsCaaRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsCaaRecordTtl float64

const (
	DNSRecordResponseDNSRecordsCaaRecordTtl1 DNSRecordResponseDNSRecordsCaaRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsCaaRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsCaaRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsCaaRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                             `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsCaaRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsCaaRecordSettingsJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsCaaRecordSettings]
type dnsRecordResponseDNSRecordsCaaRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsCaaRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsCaaRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsCertRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsCertRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsCertRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsCertRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                 `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsCertRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsCertRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsCertRecord]
type dnsRecordResponseDNSRecordsCertRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsCertRecordTtl float64

const (
	DNSRecordResponseDNSRecordsCertRecordTtl1 DNSRecordResponseDNSRecordsCertRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsCertRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsCertRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsCertRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                              `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsCertRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsCertRecordSettingsJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsCertRecordSettings]
type dnsRecordResponseDNSRecordsCertRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsCertRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsCertRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsDnskeyRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsDnskeyRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsDnskeyRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsDnskeyRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                   `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsDnskeyRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsDnskeyRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsDnskeyRecord]
type dnsRecordResponseDNSRecordsDnskeyRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsDnskeyRecordTtl float64

const (
	DNSRecordResponseDNSRecordsDnskeyRecordTtl1 DNSRecordResponseDNSRecordsDnskeyRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsDnskeyRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsDnskeyRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsDnskeyRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                                `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsDnskeyRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsDnskeyRecordSettingsJSON contains the JSON metadata
// for the struct [DNSRecordResponseDNSRecordsDnskeyRecordSettings]
type dnsRecordResponseDNSRecordsDnskeyRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsDnskeyRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsDnskeyRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsDsRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsDsRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsDsRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsDsRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsDsRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsDsRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsDsRecord]
type dnsRecordResponseDNSRecordsDsRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsDsRecordTtl float64

const (
	DNSRecordResponseDNSRecordsDsRecordTtl1 DNSRecordResponseDNSRecordsDsRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsDsRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsDsRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsDsRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                            `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsDsRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsDsRecordSettingsJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsDsRecordSettings]
type dnsRecordResponseDNSRecordsDsRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsDsRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsDsRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsHTTPSRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsHTTPSRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsHTTPSRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsHTTPSRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                  `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsHTTPSRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsHTTPSRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsHTTPSRecord]
type dnsRecordResponseDNSRecordsHTTPSRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsHTTPSRecordTtl float64

const (
	DNSRecordResponseDNSRecordsHTTPSRecordTtl1 DNSRecordResponseDNSRecordsHTTPSRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsHTTPSRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsHTTPSRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsHTTPSRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                               `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsHTTPSRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsHTTPSRecordSettingsJSON contains the JSON metadata
// for the struct [DNSRecordResponseDNSRecordsHTTPSRecordSettings]
type dnsRecordResponseDNSRecordsHTTPSRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsHTTPSRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsHTTPSRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsLocRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsLocRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsLocRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsLocRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsLocRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsLocRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsLocRecord]
type dnsRecordResponseDNSRecordsLocRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsLocRecordTtl float64

const (
	DNSRecordResponseDNSRecordsLocRecordTtl1 DNSRecordResponseDNSRecordsLocRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsLocRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsLocRecordTtl1:
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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsLocRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                             `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsLocRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsLocRecordSettingsJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsLocRecordSettings]
type dnsRecordResponseDNSRecordsLocRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsLocRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsLocRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsNaptrRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsNaptrRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsNaptrRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsNaptrRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                  `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsNaptrRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsNaptrRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsNaptrRecord]
type dnsRecordResponseDNSRecordsNaptrRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsNaptrRecordTtl float64

const (
	DNSRecordResponseDNSRecordsNaptrRecordTtl1 DNSRecordResponseDNSRecordsNaptrRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsNaptrRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsNaptrRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsNaptrRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                               `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsNaptrRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsNaptrRecordSettingsJSON contains the JSON metadata
// for the struct [DNSRecordResponseDNSRecordsNaptrRecordSettings]
type dnsRecordResponseDNSRecordsNaptrRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsNaptrRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsNaptrRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsSmimeaRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsSmimeaRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsSmimeaRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsSmimeaRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                   `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsSmimeaRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsSmimeaRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsSmimeaRecord]
type dnsRecordResponseDNSRecordsSmimeaRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsSmimeaRecordTtl float64

const (
	DNSRecordResponseDNSRecordsSmimeaRecordTtl1 DNSRecordResponseDNSRecordsSmimeaRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsSmimeaRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsSmimeaRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsSmimeaRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                                `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsSmimeaRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsSmimeaRecordSettingsJSON contains the JSON metadata
// for the struct [DNSRecordResponseDNSRecordsSmimeaRecordSettings]
type dnsRecordResponseDNSRecordsSmimeaRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsSmimeaRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsSmimeaRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsSrvRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsSrvRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsSrvRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsSrvRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsSrvRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsSrvRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsSrvRecord]
type dnsRecordResponseDNSRecordsSrvRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsSrvRecordTtl float64

const (
	DNSRecordResponseDNSRecordsSrvRecordTtl1 DNSRecordResponseDNSRecordsSrvRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsSrvRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsSrvRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsSrvRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                             `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsSrvRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsSrvRecordSettingsJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsSrvRecordSettings]
type dnsRecordResponseDNSRecordsSrvRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsSrvRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsSrvRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsSshfpRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsSshfpRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsSshfpRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsSshfpRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                  `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsSshfpRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsSshfpRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsSshfpRecord]
type dnsRecordResponseDNSRecordsSshfpRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsSshfpRecordTtl float64

const (
	DNSRecordResponseDNSRecordsSshfpRecordTtl1 DNSRecordResponseDNSRecordsSshfpRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsSshfpRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsSshfpRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsSshfpRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                               `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsSshfpRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsSshfpRecordSettingsJSON contains the JSON metadata
// for the struct [DNSRecordResponseDNSRecordsSshfpRecordSettings]
type dnsRecordResponseDNSRecordsSshfpRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsSshfpRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsSshfpRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsSvcbRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsSvcbRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsSvcbRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsSvcbRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                 `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsSvcbRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsSvcbRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsSvcbRecord]
type dnsRecordResponseDNSRecordsSvcbRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsSvcbRecordTtl float64

const (
	DNSRecordResponseDNSRecordsSvcbRecordTtl1 DNSRecordResponseDNSRecordsSvcbRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsSvcbRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsSvcbRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsSvcbRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                              `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsSvcbRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsSvcbRecordSettingsJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsSvcbRecordSettings]
type dnsRecordResponseDNSRecordsSvcbRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsSvcbRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsSvcbRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsTlsaRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsTlsaRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsTlsaRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsTlsaRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                 `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsTlsaRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsTlsaRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsTlsaRecord]
type dnsRecordResponseDNSRecordsTlsaRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsTlsaRecordTtl float64

const (
	DNSRecordResponseDNSRecordsTlsaRecordTtl1 DNSRecordResponseDNSRecordsTlsaRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsTlsaRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsTlsaRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsTlsaRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                              `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsTlsaRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsTlsaRecordSettingsJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsTlsaRecordSettings]
type dnsRecordResponseDNSRecordsTlsaRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsTlsaRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsTlsaRecordSettingsJSON) RawJSON() string {
	return r.raw
}

type DNSRecordResponseDNSRecordsUriRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name string `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResponseDNSRecordsUriRecordTtl `json:"ttl,required"`
	// Record type.
	Type DNSRecordResponseDNSRecordsUriRecordType `json:"type,required"`
	// Identifier.
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
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
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
	Settings DNSRecordResponseDNSRecordsUriRecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                `json:"tags_modified_on" format:"date-time"`
	JSON           dnsRecordResponseDNSRecordsUriRecordJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsUriRecordJSON contains the JSON metadata for the
// struct [DNSRecordResponseDNSRecordsUriRecord]
type dnsRecordResponseDNSRecordsUriRecordJSON struct {
	Name              apijson.Field
	Ttl               apijson.Field
	Type              apijson.Field
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Priority          apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
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

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type DNSRecordResponseDNSRecordsUriRecordTtl float64

const (
	DNSRecordResponseDNSRecordsUriRecordTtl1 DNSRecordResponseDNSRecordsUriRecordTtl = 1
)

func (r DNSRecordResponseDNSRecordsUriRecordTtl) IsKnown() bool {
	switch r {
	case DNSRecordResponseDNSRecordsUriRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type DNSRecordResponseDNSRecordsUriRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	Ipv6Only bool                                             `json:"ipv6_only"`
	JSON     dnsRecordResponseDNSRecordsUriRecordSettingsJSON `json:"-"`
}

// dnsRecordResponseDNSRecordsUriRecordSettingsJSON contains the JSON metadata for
// the struct [DNSRecordResponseDNSRecordsUriRecordSettings]
type dnsRecordResponseDNSRecordsUriRecordSettingsJSON struct {
	Ipv4Only    apijson.Field
	Ipv6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResponseDNSRecordsUriRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordResponseDNSRecordsUriRecordSettingsJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordResponseType string

const (
	DNSRecordResponseTypeA          DNSRecordResponseType = "A"
	DNSRecordResponseTypeAaaa       DNSRecordResponseType = "AAAA"
	DNSRecordResponseTypeCname      DNSRecordResponseType = "CNAME"
	DNSRecordResponseTypeMx         DNSRecordResponseType = "MX"
	DNSRecordResponseTypeNs         DNSRecordResponseType = "NS"
	DNSRecordResponseTypeOpenpgpkey DNSRecordResponseType = "OPENPGPKEY"
	DNSRecordResponseTypePtr        DNSRecordResponseType = "PTR"
	DNSRecordResponseTypeTxt        DNSRecordResponseType = "TXT"
	DNSRecordResponseTypeCaa        DNSRecordResponseType = "CAA"
	DNSRecordResponseTypeCert       DNSRecordResponseType = "CERT"
	DNSRecordResponseTypeDnskey     DNSRecordResponseType = "DNSKEY"
	DNSRecordResponseTypeDs         DNSRecordResponseType = "DS"
	DNSRecordResponseTypeHTTPS      DNSRecordResponseType = "HTTPS"
	DNSRecordResponseTypeLoc        DNSRecordResponseType = "LOC"
	DNSRecordResponseTypeNaptr      DNSRecordResponseType = "NAPTR"
	DNSRecordResponseTypeSmimea     DNSRecordResponseType = "SMIMEA"
	DNSRecordResponseTypeSrv        DNSRecordResponseType = "SRV"
	DNSRecordResponseTypeSshfp      DNSRecordResponseType = "SSHFP"
	DNSRecordResponseTypeSvcb       DNSRecordResponseType = "SVCB"
	DNSRecordResponseTypeTlsa       DNSRecordResponseType = "TLSA"
	DNSRecordResponseTypeUri        DNSRecordResponseType = "URI"
)

func (r DNSRecordResponseType) IsKnown() bool {
	switch r {
	case DNSRecordResponseTypeA, DNSRecordResponseTypeAaaa, DNSRecordResponseTypeCname, DNSRecordResponseTypeMx, DNSRecordResponseTypeNs, DNSRecordResponseTypeOpenpgpkey, DNSRecordResponseTypePtr, DNSRecordResponseTypeTxt, DNSRecordResponseTypeCaa, DNSRecordResponseTypeCert, DNSRecordResponseTypeDnskey, DNSRecordResponseTypeDs, DNSRecordResponseTypeHTTPS, DNSRecordResponseTypeLoc, DNSRecordResponseTypeNaptr, DNSRecordResponseTypeSmimea, DNSRecordResponseTypeSrv, DNSRecordResponseTypeSshfp, DNSRecordResponseTypeSvcb, DNSRecordResponseTypeTlsa, DNSRecordResponseTypeUri:
		return true
	}
	return false
}

type ImportScanResponse struct {
	Errors   []DNSRecordMessageItem `json:"errors,required"`
	Messages []DNSRecordMessageItem `json:"messages,required"`
	// Whether the API call was successful.
	Success ImportScanResponseSuccess `json:"success,required"`
	Result  ImportScanResponseResult  `json:"result"`
	JSON    importScanResponseJSON    `json:"-"`
}

// importScanResponseJSON contains the JSON metadata for the struct
// [ImportScanResponse]
type importScanResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful.
type ImportScanResponseSuccess bool

const (
	ImportScanResponseSuccessTrue ImportScanResponseSuccess = true
)

func (r ImportScanResponseSuccess) IsKnown() bool {
	switch r {
	case ImportScanResponseSuccessTrue:
		return true
	}
	return false
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

type SingleResponseDNSResponse struct {
	Errors   []DNSRecordMessageItem `json:"errors,required"`
	Messages []DNSRecordMessageItem `json:"messages,required"`
	// Whether the API call was successful.
	Success SingleResponseDNSResponseSuccess `json:"success,required"`
	Result  DNSRecordResponse                `json:"result"`
	JSON    singleResponseDNSResponseJSON    `json:"-"`
}

// singleResponseDNSResponseJSON contains the JSON metadata for the struct
// [SingleResponseDNSResponse]
type singleResponseDNSResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful.
type SingleResponseDNSResponseSuccess bool

const (
	SingleResponseDNSResponseSuccessTrue SingleResponseDNSResponseSuccess = true
)

func (r SingleResponseDNSResponseSuccess) IsKnown() bool {
	switch r {
	case SingleResponseDNSResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneDNSRecordListResponse struct {
	Errors   []DNSRecordMessageItem `json:"errors,required"`
	Messages []DNSRecordMessageItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    ZoneDNSRecordListResponseSuccess    `json:"success,required"`
	Result     []DNSRecordResponse                 `json:"result"`
	ResultInfo ZoneDNSRecordListResponseResultInfo `json:"result_info"`
	JSON       zoneDNSRecordListResponseJSON       `json:"-"`
}

// zoneDNSRecordListResponseJSON contains the JSON metadata for the struct
// [ZoneDNSRecordListResponse]
type zoneDNSRecordListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful.
type ZoneDNSRecordListResponseSuccess bool

const (
	ZoneDNSRecordListResponseSuccessTrue ZoneDNSRecordListResponseSuccess = true
)

func (r ZoneDNSRecordListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneDNSRecordListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneDNSRecordListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
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
	// Identifier.
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
	Errors   []DNSRecordMessageItem `json:"errors,required"`
	Messages []DNSRecordMessageItem `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneDNSRecordBatchResponseSuccess `json:"success,required"`
	Result  ZoneDNSRecordBatchResponseResult  `json:"result"`
	JSON    zoneDNSRecordBatchResponseJSON    `json:"-"`
}

// zoneDNSRecordBatchResponseJSON contains the JSON metadata for the struct
// [ZoneDNSRecordBatchResponse]
type zoneDNSRecordBatchResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful.
type ZoneDNSRecordBatchResponseSuccess bool

const (
	ZoneDNSRecordBatchResponseSuccessTrue ZoneDNSRecordBatchResponseSuccess = true
)

func (r ZoneDNSRecordBatchResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchResponseSuccessTrue:
		return true
	}
	return false
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
	// Identifier.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneDNSRecordBatchParamsDelete) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatch struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string]      `json:"name,required"`
	Ttl  param.Field[interface{}] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid IPv4 address.
	Content param.Field[string]      `json:"content" format:"ipv4"`
	Data    param.Field[interface{}] `json:"data"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied  param.Field[bool]        `json:"proxied"`
	Settings param.Field[interface{}] `json:"settings"`
	Tags     param.Field[interface{}] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatch) implementsZoneDNSRecordBatchParamsPatchUnion() {}

// Satisfied by [ZoneDNSRecordBatchParamsPatchesDNSRecordsARecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecord],
// [ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecord],
// [ZoneDNSRecordBatchParamsPatch].
type ZoneDNSRecordBatchParamsPatchUnion interface {
	implementsZoneDNSRecordBatchParamsPatchUnion()
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsARecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsARecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsARecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid IPv4 address.
	Content param.Field[string] `json:"content" format:"ipv4"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsARecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsARecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsARecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsARecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsARecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsARecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsARecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsARecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsARecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid IPv6 address.
	Content param.Field[string] `json:"content" format:"ipv6"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsAaaaRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid hostname. Must not match the record's name.
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordSettings struct {
	// If enabled, causes the CNAME record to be resolved externally and the resulting
	// address records (e.g., A and AAAA) to be returned instead of the CNAME record
	// itself. This setting is unavailable for proxied records, since they are always
	// flattened.
	FlattenCname param.Field[bool] `json:"flatten_cname"`
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCnameRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid mail server hostname.
	Content param.Field[string] `json:"content" format:"hostname"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsMxRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid name server host name.
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsNsRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsOpenpgpkeyRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsPtrRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Text content for the record. The content must consist of quoted "character
	// strings" (RFC 1035), each with a length of up to 255 bytes. Strings exceeding
	// this allowed maximum length are automatically split.
	//
	// Learn more at
	// <https://www.cloudflare.com/learning/dns/dns-records/dns-txt-record/>.
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsTxtRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a CAA record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCaaRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a CERT record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsCertRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a DNSKEY record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsDnskeyRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a DS record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsDsRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a HTTPS record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsHTTPSRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a LOC record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordTtl1:
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsLocRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a NAPTR record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsNaptrRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SMIMEA record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSmimeaRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SRV record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSrvRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SSHFP record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSshfpRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SVCB record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsSvcbRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a TLSA record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsTlsaRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a URI record.
	Data param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordData] `json:"data"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecord) implementsZoneDNSRecordBatchParamsPatchUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordTtl1 ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPatchesDNSRecordsUriRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPatchesType string

const (
	ZoneDNSRecordBatchParamsPatchesTypeA          ZoneDNSRecordBatchParamsPatchesType = "A"
	ZoneDNSRecordBatchParamsPatchesTypeAaaa       ZoneDNSRecordBatchParamsPatchesType = "AAAA"
	ZoneDNSRecordBatchParamsPatchesTypeCname      ZoneDNSRecordBatchParamsPatchesType = "CNAME"
	ZoneDNSRecordBatchParamsPatchesTypeMx         ZoneDNSRecordBatchParamsPatchesType = "MX"
	ZoneDNSRecordBatchParamsPatchesTypeNs         ZoneDNSRecordBatchParamsPatchesType = "NS"
	ZoneDNSRecordBatchParamsPatchesTypeOpenpgpkey ZoneDNSRecordBatchParamsPatchesType = "OPENPGPKEY"
	ZoneDNSRecordBatchParamsPatchesTypePtr        ZoneDNSRecordBatchParamsPatchesType = "PTR"
	ZoneDNSRecordBatchParamsPatchesTypeTxt        ZoneDNSRecordBatchParamsPatchesType = "TXT"
	ZoneDNSRecordBatchParamsPatchesTypeCaa        ZoneDNSRecordBatchParamsPatchesType = "CAA"
	ZoneDNSRecordBatchParamsPatchesTypeCert       ZoneDNSRecordBatchParamsPatchesType = "CERT"
	ZoneDNSRecordBatchParamsPatchesTypeDnskey     ZoneDNSRecordBatchParamsPatchesType = "DNSKEY"
	ZoneDNSRecordBatchParamsPatchesTypeDs         ZoneDNSRecordBatchParamsPatchesType = "DS"
	ZoneDNSRecordBatchParamsPatchesTypeHTTPS      ZoneDNSRecordBatchParamsPatchesType = "HTTPS"
	ZoneDNSRecordBatchParamsPatchesTypeLoc        ZoneDNSRecordBatchParamsPatchesType = "LOC"
	ZoneDNSRecordBatchParamsPatchesTypeNaptr      ZoneDNSRecordBatchParamsPatchesType = "NAPTR"
	ZoneDNSRecordBatchParamsPatchesTypeSmimea     ZoneDNSRecordBatchParamsPatchesType = "SMIMEA"
	ZoneDNSRecordBatchParamsPatchesTypeSrv        ZoneDNSRecordBatchParamsPatchesType = "SRV"
	ZoneDNSRecordBatchParamsPatchesTypeSshfp      ZoneDNSRecordBatchParamsPatchesType = "SSHFP"
	ZoneDNSRecordBatchParamsPatchesTypeSvcb       ZoneDNSRecordBatchParamsPatchesType = "SVCB"
	ZoneDNSRecordBatchParamsPatchesTypeTlsa       ZoneDNSRecordBatchParamsPatchesType = "TLSA"
	ZoneDNSRecordBatchParamsPatchesTypeUri        ZoneDNSRecordBatchParamsPatchesType = "URI"
)

func (r ZoneDNSRecordBatchParamsPatchesType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPatchesTypeA, ZoneDNSRecordBatchParamsPatchesTypeAaaa, ZoneDNSRecordBatchParamsPatchesTypeCname, ZoneDNSRecordBatchParamsPatchesTypeMx, ZoneDNSRecordBatchParamsPatchesTypeNs, ZoneDNSRecordBatchParamsPatchesTypeOpenpgpkey, ZoneDNSRecordBatchParamsPatchesTypePtr, ZoneDNSRecordBatchParamsPatchesTypeTxt, ZoneDNSRecordBatchParamsPatchesTypeCaa, ZoneDNSRecordBatchParamsPatchesTypeCert, ZoneDNSRecordBatchParamsPatchesTypeDnskey, ZoneDNSRecordBatchParamsPatchesTypeDs, ZoneDNSRecordBatchParamsPatchesTypeHTTPS, ZoneDNSRecordBatchParamsPatchesTypeLoc, ZoneDNSRecordBatchParamsPatchesTypeNaptr, ZoneDNSRecordBatchParamsPatchesTypeSmimea, ZoneDNSRecordBatchParamsPatchesTypeSrv, ZoneDNSRecordBatchParamsPatchesTypeSshfp, ZoneDNSRecordBatchParamsPatchesTypeSvcb, ZoneDNSRecordBatchParamsPatchesTypeTlsa, ZoneDNSRecordBatchParamsPatchesTypeUri:
		return true
	}
	return false
}

type ZoneDNSRecordBatchParamsPut struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string]      `json:"name,required"`
	Ttl  param.Field[interface{}] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid IPv4 address.
	Content param.Field[string]      `json:"content" format:"ipv4"`
	Data    param.Field[interface{}] `json:"data"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied  param.Field[bool]        `json:"proxied"`
	Settings param.Field[interface{}] `json:"settings"`
	Tags     param.Field[interface{}] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPut) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPut) implementsZoneDNSRecordBatchParamsPutUnion() {}

// Satisfied by [ZoneDNSRecordBatchParamsPutsDNSRecordsARecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecord],
// [ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecord],
// [ZoneDNSRecordBatchParamsPut].
type ZoneDNSRecordBatchParamsPutUnion interface {
	implementsZoneDNSRecordBatchParamsPutUnion()
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsARecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsARecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsARecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid IPv4 address.
	Content param.Field[string] `json:"content" format:"ipv4"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsARecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsARecord) implementsZoneDNSRecordBatchParamsPutUnion() {}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsARecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsARecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsARecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsARecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsARecordTtl1:
		return true
	}
	return false
}

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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsARecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsARecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid IPv6 address.
	Content param.Field[string] `json:"content" format:"ipv6"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsAaaaRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid hostname. Must not match the record's name.
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordSettings struct {
	// If enabled, causes the CNAME record to be resolved externally and the resulting
	// address records (e.g., A and AAAA) to be returned instead of the CNAME record
	// itself. This setting is unavailable for proxied records, since they are always
	// flattened.
	FlattenCname param.Field[bool] `json:"flatten_cname"`
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCnameRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid mail server hostname.
	Content param.Field[string] `json:"content" format:"hostname"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsMxRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid name server host name.
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsNsRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsOpenpgpkeyRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsPtrRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Text content for the record. The content must consist of quoted "character
	// strings" (RFC 1035), each with a length of up to 255 bytes. Strings exceeding
	// this allowed maximum length are automatically split.
	//
	// Learn more at
	// <https://www.cloudflare.com/learning/dns/dns-records/dns-txt-record/>.
	Content param.Field[string] `json:"content"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsTxtRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a CAA record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCaaRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a CERT record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsCertRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a DNSKEY record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsDnskeyRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a DS record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsDsRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a HTTPS record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsHTTPSRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a LOC record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordTtl1:
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsLocRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a NAPTR record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsNaptrRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SMIMEA record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSmimeaRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SRV record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSrvRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SSHFP record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSshfpRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SVCB record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsSvcbRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a TLSA record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordData] `json:"data"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsTlsaRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecord struct {
	// Complete DNS record name, including the zone name, in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordTtl] `json:"ttl,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordType] `json:"type,required"`
	// Identifier.
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a URI record.
	Data param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordData] `json:"data"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordSettings] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecord) implementsZoneDNSRecordBatchParamsPutUnion() {
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordTtl float64

const (
	ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordTtl1 ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordTtl = 1
)

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordTtl) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordTtl1:
		return true
	}
	return false
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

// Settings for the DNS record.
type ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordSettings struct {
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

func (r ZoneDNSRecordBatchParamsPutsDNSRecordsUriRecordSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordBatchParamsPutsType string

const (
	ZoneDNSRecordBatchParamsPutsTypeA          ZoneDNSRecordBatchParamsPutsType = "A"
	ZoneDNSRecordBatchParamsPutsTypeAaaa       ZoneDNSRecordBatchParamsPutsType = "AAAA"
	ZoneDNSRecordBatchParamsPutsTypeCname      ZoneDNSRecordBatchParamsPutsType = "CNAME"
	ZoneDNSRecordBatchParamsPutsTypeMx         ZoneDNSRecordBatchParamsPutsType = "MX"
	ZoneDNSRecordBatchParamsPutsTypeNs         ZoneDNSRecordBatchParamsPutsType = "NS"
	ZoneDNSRecordBatchParamsPutsTypeOpenpgpkey ZoneDNSRecordBatchParamsPutsType = "OPENPGPKEY"
	ZoneDNSRecordBatchParamsPutsTypePtr        ZoneDNSRecordBatchParamsPutsType = "PTR"
	ZoneDNSRecordBatchParamsPutsTypeTxt        ZoneDNSRecordBatchParamsPutsType = "TXT"
	ZoneDNSRecordBatchParamsPutsTypeCaa        ZoneDNSRecordBatchParamsPutsType = "CAA"
	ZoneDNSRecordBatchParamsPutsTypeCert       ZoneDNSRecordBatchParamsPutsType = "CERT"
	ZoneDNSRecordBatchParamsPutsTypeDnskey     ZoneDNSRecordBatchParamsPutsType = "DNSKEY"
	ZoneDNSRecordBatchParamsPutsTypeDs         ZoneDNSRecordBatchParamsPutsType = "DS"
	ZoneDNSRecordBatchParamsPutsTypeHTTPS      ZoneDNSRecordBatchParamsPutsType = "HTTPS"
	ZoneDNSRecordBatchParamsPutsTypeLoc        ZoneDNSRecordBatchParamsPutsType = "LOC"
	ZoneDNSRecordBatchParamsPutsTypeNaptr      ZoneDNSRecordBatchParamsPutsType = "NAPTR"
	ZoneDNSRecordBatchParamsPutsTypeSmimea     ZoneDNSRecordBatchParamsPutsType = "SMIMEA"
	ZoneDNSRecordBatchParamsPutsTypeSrv        ZoneDNSRecordBatchParamsPutsType = "SRV"
	ZoneDNSRecordBatchParamsPutsTypeSshfp      ZoneDNSRecordBatchParamsPutsType = "SSHFP"
	ZoneDNSRecordBatchParamsPutsTypeSvcb       ZoneDNSRecordBatchParamsPutsType = "SVCB"
	ZoneDNSRecordBatchParamsPutsTypeTlsa       ZoneDNSRecordBatchParamsPutsType = "TLSA"
	ZoneDNSRecordBatchParamsPutsTypeUri        ZoneDNSRecordBatchParamsPutsType = "URI"
)

func (r ZoneDNSRecordBatchParamsPutsType) IsKnown() bool {
	switch r {
	case ZoneDNSRecordBatchParamsPutsTypeA, ZoneDNSRecordBatchParamsPutsTypeAaaa, ZoneDNSRecordBatchParamsPutsTypeCname, ZoneDNSRecordBatchParamsPutsTypeMx, ZoneDNSRecordBatchParamsPutsTypeNs, ZoneDNSRecordBatchParamsPutsTypeOpenpgpkey, ZoneDNSRecordBatchParamsPutsTypePtr, ZoneDNSRecordBatchParamsPutsTypeTxt, ZoneDNSRecordBatchParamsPutsTypeCaa, ZoneDNSRecordBatchParamsPutsTypeCert, ZoneDNSRecordBatchParamsPutsTypeDnskey, ZoneDNSRecordBatchParamsPutsTypeDs, ZoneDNSRecordBatchParamsPutsTypeHTTPS, ZoneDNSRecordBatchParamsPutsTypeLoc, ZoneDNSRecordBatchParamsPutsTypeNaptr, ZoneDNSRecordBatchParamsPutsTypeSmimea, ZoneDNSRecordBatchParamsPutsTypeSrv, ZoneDNSRecordBatchParamsPutsTypeSshfp, ZoneDNSRecordBatchParamsPutsTypeSvcb, ZoneDNSRecordBatchParamsPutsTypeTlsa, ZoneDNSRecordBatchParamsPutsTypeUri:
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
