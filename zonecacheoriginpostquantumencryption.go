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

type ZoneCacheOriginPostQuantumEncryptionGetResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneCacheOriginPostQuantumEncryptionGetResponseSuccess `json:"success,required"`
	Result  ZoneCacheOriginPostQuantumEncryptionGetResponseResult  `json:"result"`
	JSON    zoneCacheOriginPostQuantumEncryptionGetResponseJSON    `json:"-"`
}

// zoneCacheOriginPostQuantumEncryptionGetResponseJSON contains the JSON metadata
// for the struct [ZoneCacheOriginPostQuantumEncryptionGetResponse]
type zoneCacheOriginPostQuantumEncryptionGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheOriginPostQuantumEncryptionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheOriginPostQuantumEncryptionGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneCacheOriginPostQuantumEncryptionGetResponseSuccess bool

const (
	ZoneCacheOriginPostQuantumEncryptionGetResponseSuccessTrue ZoneCacheOriginPostQuantumEncryptionGetResponseSuccess = true
)

func (r ZoneCacheOriginPostQuantumEncryptionGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneCacheOriginPostQuantumEncryptionGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneCacheOriginPostQuantumEncryptionGetResponseResult struct {
	// Value of the zone setting.
	ID ZoneCacheOriginPostQuantumEncryptionGetResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value OriginPostQuantumEncryptionValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneCacheOriginPostQuantumEncryptionGetResponseResultJSON `json:"-"`
}

// zoneCacheOriginPostQuantumEncryptionGetResponseResultJSON contains the JSON
// metadata for the struct [ZoneCacheOriginPostQuantumEncryptionGetResponseResult]
type zoneCacheOriginPostQuantumEncryptionGetResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheOriginPostQuantumEncryptionGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheOriginPostQuantumEncryptionGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// Value of the zone setting.
type ZoneCacheOriginPostQuantumEncryptionGetResponseResultID string

const (
	ZoneCacheOriginPostQuantumEncryptionGetResponseResultIDOriginPqe ZoneCacheOriginPostQuantumEncryptionGetResponseResultID = "origin_pqe"
)

func (r ZoneCacheOriginPostQuantumEncryptionGetResponseResultID) IsKnown() bool {
	switch r {
	case ZoneCacheOriginPostQuantumEncryptionGetResponseResultIDOriginPqe:
		return true
	}
	return false
}

type ZoneCacheOriginPostQuantumEncryptionUpdateResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneCacheOriginPostQuantumEncryptionUpdateResponseSuccess `json:"success,required"`
	Result  ZoneCacheOriginPostQuantumEncryptionUpdateResponseResult  `json:"result"`
	JSON    zoneCacheOriginPostQuantumEncryptionUpdateResponseJSON    `json:"-"`
}

// zoneCacheOriginPostQuantumEncryptionUpdateResponseJSON contains the JSON
// metadata for the struct [ZoneCacheOriginPostQuantumEncryptionUpdateResponse]
type zoneCacheOriginPostQuantumEncryptionUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheOriginPostQuantumEncryptionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheOriginPostQuantumEncryptionUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneCacheOriginPostQuantumEncryptionUpdateResponseSuccess bool

const (
	ZoneCacheOriginPostQuantumEncryptionUpdateResponseSuccessTrue ZoneCacheOriginPostQuantumEncryptionUpdateResponseSuccess = true
)

func (r ZoneCacheOriginPostQuantumEncryptionUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneCacheOriginPostQuantumEncryptionUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneCacheOriginPostQuantumEncryptionUpdateResponseResult struct {
	// Value of the zone setting.
	ID ZoneCacheOriginPostQuantumEncryptionUpdateResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value OriginPostQuantumEncryptionValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneCacheOriginPostQuantumEncryptionUpdateResponseResultJSON `json:"-"`
}

// zoneCacheOriginPostQuantumEncryptionUpdateResponseResultJSON contains the JSON
// metadata for the struct
// [ZoneCacheOriginPostQuantumEncryptionUpdateResponseResult]
type zoneCacheOriginPostQuantumEncryptionUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheOriginPostQuantumEncryptionUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheOriginPostQuantumEncryptionUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

// Value of the zone setting.
type ZoneCacheOriginPostQuantumEncryptionUpdateResponseResultID string

const (
	ZoneCacheOriginPostQuantumEncryptionUpdateResponseResultIDOriginPqe ZoneCacheOriginPostQuantumEncryptionUpdateResponseResultID = "origin_pqe"
)

func (r ZoneCacheOriginPostQuantumEncryptionUpdateResponseResultID) IsKnown() bool {
	switch r {
	case ZoneCacheOriginPostQuantumEncryptionUpdateResponseResultIDOriginPqe:
		return true
	}
	return false
}

type ZoneCacheOriginPostQuantumEncryptionUpdateParams struct {
	// Value of the Origin Post Quantum Encryption Setting.
	Value param.Field[OriginPostQuantumEncryptionValue] `json:"value,required"`
}

func (r ZoneCacheOriginPostQuantumEncryptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
