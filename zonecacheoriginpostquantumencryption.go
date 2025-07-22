// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneCacheOriginPostQuantumEncryptionService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCacheOriginPostQuantumEncryptionService] method instead.
type ZoneCacheOriginPostQuantumEncryptionService struct {
	Options []option.RequestOption
}

// NewZoneCacheOriginPostQuantumEncryptionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZoneCacheOriginPostQuantumEncryptionService(opts ...option.RequestOption) (r *ZoneCacheOriginPostQuantumEncryptionService) {
	r = &ZoneCacheOriginPostQuantumEncryptionService{}
	r.Options = opts
	return
}

// Instructs Cloudflare to use Post-Quantum (PQ) key agreement algorithms when
// connecting to your origin. Preferred instructs Cloudflare to opportunistically
// send a Post-Quantum keyshare in the first message to the origin (for fastest
// connections when the origin supports and prefers PQ), supported means that PQ
// algorithms are advertised but only used when requested by the origin, and off
// means that PQ algorithms are not advertised
func (r *ZoneCacheOriginPostQuantumEncryptionService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneCacheOriginPostQuantumEncryptionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/origin_post_quantum_encryption", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Instructs Cloudflare to use Post-Quantum (PQ) key agreement algorithms when
// connecting to your origin. Preferred instructs Cloudflare to opportunistically
// send a Post-Quantum keyshare in the first message to the origin (for fastest
// connections when the origin supports and prefers PQ), supported means that PQ
// algorithms are advertised but only used when requested by the origin, and off
// means that PQ algorithms are not advertised
func (r *ZoneCacheOriginPostQuantumEncryptionService) Update(ctx context.Context, zoneID string, body ZoneCacheOriginPostQuantumEncryptionUpdateParams, opts ...option.RequestOption) (res *ZoneCacheOriginPostQuantumEncryptionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/origin_post_quantum_encryption", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Value of the Origin Post Quantum Encryption Setting.
type OriginPostQuantumEncryptionValue string

const (
	OriginPostQuantumEncryptionValuePreferred OriginPostQuantumEncryptionValue = "preferred"
	OriginPostQuantumEncryptionValueSupported OriginPostQuantumEncryptionValue = "supported"
	OriginPostQuantumEncryptionValueOff       OriginPostQuantumEncryptionValue = "off"
)

func (r OriginPostQuantumEncryptionValue) IsKnown() bool {
	switch r {
	case OriginPostQuantumEncryptionValuePreferred, OriginPostQuantumEncryptionValueSupported, OriginPostQuantumEncryptionValueOff:
		return true
	}
	return false
}

type ResponseValueEncryption struct {
	Result ResponseValueEncryptionResult `json:"result"`
	JSON   responseValueEncryptionJSON   `json:"-"`
}

// responseValueEncryptionJSON contains the JSON metadata for the struct
// [ResponseValueEncryption]
type responseValueEncryptionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseValueEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseValueEncryptionJSON) RawJSON() string {
	return r.raw
}

type ResponseValueEncryptionResult struct {
	// Value of the Origin Post Quantum Encryption Setting.
	Value OriginPostQuantumEncryptionValue `json:"value,required"`
	// Value of the zone setting.
	ID   ResponseValueEncryptionResultID   `json:"id"`
	JSON responseValueEncryptionResultJSON `json:"-"`
	BaseCacheRule
}

// responseValueEncryptionResultJSON contains the JSON metadata for the struct
// [ResponseValueEncryptionResult]
type responseValueEncryptionResultJSON struct {
	Value       apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseValueEncryptionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseValueEncryptionResultJSON) RawJSON() string {
	return r.raw
}

// Value of the zone setting.
type ResponseValueEncryptionResultID string

const (
	ResponseValueEncryptionResultIDOriginPqe ResponseValueEncryptionResultID = "origin_pqe"
)

func (r ResponseValueEncryptionResultID) IsKnown() bool {
	switch r {
	case ResponseValueEncryptionResultIDOriginPqe:
		return true
	}
	return false
}

type ZoneCacheOriginPostQuantumEncryptionGetResponse struct {
	JSON zoneCacheOriginPostQuantumEncryptionGetResponseJSON `json:"-"`
	CacheRulesZoneCacheSettingsResponse
	ResponseValueEncryption
}

// zoneCacheOriginPostQuantumEncryptionGetResponseJSON contains the JSON metadata
// for the struct [ZoneCacheOriginPostQuantumEncryptionGetResponse]
type zoneCacheOriginPostQuantumEncryptionGetResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheOriginPostQuantumEncryptionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheOriginPostQuantumEncryptionGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCacheOriginPostQuantumEncryptionUpdateResponse struct {
	JSON zoneCacheOriginPostQuantumEncryptionUpdateResponseJSON `json:"-"`
	CacheRulesZoneCacheSettingsResponse
	ResponseValueEncryption
}

// zoneCacheOriginPostQuantumEncryptionUpdateResponseJSON contains the JSON
// metadata for the struct [ZoneCacheOriginPostQuantumEncryptionUpdateResponse]
type zoneCacheOriginPostQuantumEncryptionUpdateResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheOriginPostQuantumEncryptionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheOriginPostQuantumEncryptionUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCacheOriginPostQuantumEncryptionUpdateParams struct {
	// Value of the Origin Post Quantum Encryption Setting.
	Value param.Field[OriginPostQuantumEncryptionValue] `json:"value,required"`
}

func (r ZoneCacheOriginPostQuantumEncryptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
