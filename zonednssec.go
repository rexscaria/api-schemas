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

// ZoneDnssecService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneDnssecService] method instead.
type ZoneDnssecService struct {
	Options []option.RequestOption
}

// NewZoneDnssecService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneDnssecService(opts ...option.RequestOption) (r *ZoneDnssecService) {
	r = &ZoneDnssecService{}
	r.Options = opts
	return
}

// Details about DNSSEC status and configuration.
func (r *ZoneDnssecService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *DnssecResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dnssec", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable or disable DNSSEC.
func (r *ZoneDnssecService) Update(ctx context.Context, zoneID string, body ZoneDnssecUpdateParams, opts ...option.RequestOption) (res *DnssecResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dnssec", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete DNSSEC.
func (r *ZoneDnssecService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneDnssecDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dnssec", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type DnssecMessageItem struct {
	Code             int64                   `json:"code,required"`
	Message          string                  `json:"message,required"`
	DocumentationURL string                  `json:"documentation_url"`
	Source           DnssecMessageItemSource `json:"source"`
	JSON             dnssecMessageItemJSON   `json:"-"`
}

// dnssecMessageItemJSON contains the JSON metadata for the struct
// [DnssecMessageItem]
type dnssecMessageItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DnssecMessageItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnssecMessageItemJSON) RawJSON() string {
	return r.raw
}

type DnssecMessageItemSource struct {
	Pointer string                      `json:"pointer"`
	JSON    dnssecMessageItemSourceJSON `json:"-"`
}

// dnssecMessageItemSourceJSON contains the JSON metadata for the struct
// [DnssecMessageItemSource]
type dnssecMessageItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DnssecMessageItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnssecMessageItemSourceJSON) RawJSON() string {
	return r.raw
}

type DnssecResponse struct {
	Errors   []DnssecMessageItem `json:"errors,required"`
	Messages []DnssecMessageItem `json:"messages,required"`
	// Whether the API call was successful.
	Success DnssecResponseSuccess `json:"success,required"`
	Result  DnssecResponseResult  `json:"result"`
	JSON    dnssecResponseJSON    `json:"-"`
}

// dnssecResponseJSON contains the JSON metadata for the struct [DnssecResponse]
type dnssecResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DnssecResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnssecResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DnssecResponseSuccess bool

const (
	DnssecResponseSuccessTrue DnssecResponseSuccess = true
)

func (r DnssecResponseSuccess) IsKnown() bool {
	switch r {
	case DnssecResponseSuccessTrue:
		return true
	}
	return false
}

type DnssecResponseResult struct {
	// Algorithm key code.
	Algorithm string `json:"algorithm,nullable"`
	// Digest hash.
	Digest string `json:"digest,nullable"`
	// Type of digest algorithm.
	DigestAlgorithm string `json:"digest_algorithm,nullable"`
	// Coded type for digest algorithm.
	DigestType string `json:"digest_type,nullable"`
	// If true, multi-signer DNSSEC is enabled on the zone, allowing multiple providers
	// to serve a DNSSEC-signed zone at the same time. This is required for DNSKEY
	// records (except those automatically generated by Cloudflare) to be added to the
	// zone.
	//
	// See
	// [Multi-signer DNSSEC](https://developers.cloudflare.com/dns/dnssec/multi-signer-dnssec/)
	// for details.
	DnssecMultiSigner bool `json:"dnssec_multi_signer"`
	// If true, allows Cloudflare to transfer in a DNSSEC-signed zone including
	// signatures from an external provider, without requiring Cloudflare to sign any
	// records on the fly.
	//
	// Note that this feature has some limitations. See
	// [Cloudflare as Secondary](https://developers.cloudflare.com/dns/zone-setups/zone-transfers/cloudflare-as-secondary/setup/#dnssec)
	// for details.
	DnssecPresigned bool `json:"dnssec_presigned"`
	// If true, enables the use of NSEC3 together with DNSSEC on the zone. Combined
	// with setting dnssec_presigned to true, this enables the use of NSEC3 records
	// when transferring in from an external provider. If dnssec_presigned is instead
	// set to false (default), NSEC3 records will be generated and signed at request
	// time.
	//
	// See
	// [DNSSEC with NSEC3](https://developers.cloudflare.com/dns/dnssec/enable-nsec3/)
	// for details.
	DnssecUseNsec3 bool `json:"dnssec_use_nsec3"`
	// Full DS record.
	Ds string `json:"ds,nullable"`
	// Flag for DNSSEC record.
	Flags float64 `json:"flags,nullable"`
	// Code for key tag.
	KeyTag float64 `json:"key_tag,nullable"`
	// Algorithm key type.
	KeyType string `json:"key_type,nullable"`
	// When DNSSEC was last modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Public key for DS record.
	PublicKey string `json:"public_key,nullable"`
	// Status of DNSSEC, based on user-desired state and presence of necessary records.
	Status DnssecResponseResultStatus `json:"status"`
	JSON   dnssecResponseResultJSON   `json:"-"`
}

// dnssecResponseResultJSON contains the JSON metadata for the struct
// [DnssecResponseResult]
type dnssecResponseResultJSON struct {
	Algorithm         apijson.Field
	Digest            apijson.Field
	DigestAlgorithm   apijson.Field
	DigestType        apijson.Field
	DnssecMultiSigner apijson.Field
	DnssecPresigned   apijson.Field
	DnssecUseNsec3    apijson.Field
	Ds                apijson.Field
	Flags             apijson.Field
	KeyTag            apijson.Field
	KeyType           apijson.Field
	ModifiedOn        apijson.Field
	PublicKey         apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DnssecResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnssecResponseResultJSON) RawJSON() string {
	return r.raw
}

// Status of DNSSEC, based on user-desired state and presence of necessary records.
type DnssecResponseResultStatus string

const (
	DnssecResponseResultStatusActive          DnssecResponseResultStatus = "active"
	DnssecResponseResultStatusPending         DnssecResponseResultStatus = "pending"
	DnssecResponseResultStatusDisabled        DnssecResponseResultStatus = "disabled"
	DnssecResponseResultStatusPendingDisabled DnssecResponseResultStatus = "pending-disabled"
	DnssecResponseResultStatusError           DnssecResponseResultStatus = "error"
)

func (r DnssecResponseResultStatus) IsKnown() bool {
	switch r {
	case DnssecResponseResultStatusActive, DnssecResponseResultStatusPending, DnssecResponseResultStatusDisabled, DnssecResponseResultStatusPendingDisabled, DnssecResponseResultStatusError:
		return true
	}
	return false
}

type ZoneDnssecDeleteResponse struct {
	Errors   []DnssecMessageItem `json:"errors,required"`
	Messages []DnssecMessageItem `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneDnssecDeleteResponseSuccess `json:"success,required"`
	Result  string                          `json:"result"`
	JSON    zoneDnssecDeleteResponseJSON    `json:"-"`
}

// zoneDnssecDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneDnssecDeleteResponse]
type zoneDnssecDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDnssecDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDnssecDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneDnssecDeleteResponseSuccess bool

const (
	ZoneDnssecDeleteResponseSuccessTrue ZoneDnssecDeleteResponseSuccess = true
)

func (r ZoneDnssecDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneDnssecDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneDnssecUpdateParams struct {
	// If true, multi-signer DNSSEC is enabled on the zone, allowing multiple providers
	// to serve a DNSSEC-signed zone at the same time. This is required for DNSKEY
	// records (except those automatically generated by Cloudflare) to be added to the
	// zone.
	//
	// See
	// [Multi-signer DNSSEC](https://developers.cloudflare.com/dns/dnssec/multi-signer-dnssec/)
	// for details.
	DnssecMultiSigner param.Field[bool] `json:"dnssec_multi_signer"`
	// If true, allows Cloudflare to transfer in a DNSSEC-signed zone including
	// signatures from an external provider, without requiring Cloudflare to sign any
	// records on the fly.
	//
	// Note that this feature has some limitations. See
	// [Cloudflare as Secondary](https://developers.cloudflare.com/dns/zone-setups/zone-transfers/cloudflare-as-secondary/setup/#dnssec)
	// for details.
	DnssecPresigned param.Field[bool] `json:"dnssec_presigned"`
	// If true, enables the use of NSEC3 together with DNSSEC on the zone. Combined
	// with setting dnssec_presigned to true, this enables the use of NSEC3 records
	// when transferring in from an external provider. If dnssec_presigned is instead
	// set to false (default), NSEC3 records will be generated and signed at request
	// time.
	//
	// See
	// [DNSSEC with NSEC3](https://developers.cloudflare.com/dns/dnssec/enable-nsec3/)
	// for details.
	DnssecUseNsec3 param.Field[bool] `json:"dnssec_use_nsec3"`
	// Status of DNSSEC, based on user-desired state and presence of necessary records.
	Status param.Field[ZoneDnssecUpdateParamsStatus] `json:"status"`
}

func (r ZoneDnssecUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status of DNSSEC, based on user-desired state and presence of necessary records.
type ZoneDnssecUpdateParamsStatus string

const (
	ZoneDnssecUpdateParamsStatusActive   ZoneDnssecUpdateParamsStatus = "active"
	ZoneDnssecUpdateParamsStatusDisabled ZoneDnssecUpdateParamsStatus = "disabled"
)

func (r ZoneDnssecUpdateParamsStatus) IsKnown() bool {
	switch r {
	case ZoneDnssecUpdateParamsStatusActive, ZoneDnssecUpdateParamsStatusDisabled:
		return true
	}
	return false
}
