// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneCacheService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCacheService] method instead.
type ZoneCacheService struct {
	Options                        []option.RequestOption
	CacheReserve                   *ZoneCacheCacheReserveService
	CacheReserveClear              *ZoneCacheCacheReserveClearService
	OriginPostQuantumEncryption    *ZoneCacheOriginPostQuantumEncryptionService
	RegionalTieredCache            *ZoneCacheRegionalTieredCacheService
	TieredCacheSmartTopologyEnable *ZoneCacheTieredCacheSmartTopologyEnableService
	Variants                       *ZoneCacheVariantService
}

// NewZoneCacheService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneCacheService(opts ...option.RequestOption) (r *ZoneCacheService) {
	r = &ZoneCacheService{}
	r.Options = opts
	r.CacheReserve = NewZoneCacheCacheReserveService(opts...)
	r.CacheReserveClear = NewZoneCacheCacheReserveClearService(opts...)
	r.OriginPostQuantumEncryption = NewZoneCacheOriginPostQuantumEncryptionService(opts...)
	r.RegionalTieredCache = NewZoneCacheRegionalTieredCacheService(opts...)
	r.TieredCacheSmartTopologyEnable = NewZoneCacheTieredCacheSmartTopologyEnableService(opts...)
	r.Variants = NewZoneCacheVariantService(opts...)
	return
}
