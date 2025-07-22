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

// ZoneHostnameSettingService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneHostnameSettingService] method instead.
type ZoneHostnameSettingService struct {
	Options []option.RequestOption
}

// NewZoneHostnameSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneHostnameSettingService(opts ...option.RequestOption) (r *ZoneHostnameSettingService) {
	r = &ZoneHostnameSettingService{}
	r.Options = opts
	return
}

// List the requested TLS setting for the hostnames under this zone.
func (r *ZoneHostnameSettingService) Get(ctx context.Context, zoneID string, settingID SettingID, opts ...option.RequestOption) (res *ZoneHostnameSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v", zoneID, settingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the tls setting value for the hostname.
func (r *ZoneHostnameSettingService) Update(ctx context.Context, zoneID string, settingID SettingID, hostname string, body ZoneHostnameSettingUpdateParams, opts ...option.RequestOption) (res *ZoneHostnameSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if hostname == "" {
		err = errors.New("missing required hostname parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v/%s", zoneID, settingID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete the tls setting value for the hostname.
func (r *ZoneHostnameSettingService) Delete(ctx context.Context, zoneID string, settingID SettingID, hostname string, opts ...option.RequestOption) (res *ZoneHostnameSettingDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if hostname == "" {
		err = errors.New("missing required hostname parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v/%s", zoneID, settingID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// The TLS Setting name.
type SettingID string

const (
	SettingIDCiphers       SettingID = "ciphers"
	SettingIDMinTlsVersion SettingID = "min_tls_version"
	SettingIDHttp2         SettingID = "http2"
)

func (r SettingID) IsKnown() bool {
	switch r {
	case SettingIDCiphers, SettingIDMinTlsVersion, SettingIDHttp2:
		return true
	}
	return false
}

// The tls setting value.
//
// Union satisfied by [shared.UnionFloat], [shared.UnionString] or
// [TlsSettingValueArray].
type TlsSettingValueUnion interface {
	ImplementsTlsSettingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TlsSettingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TlsSettingValueArray{}),
		},
	)
}

type TlsSettingValueArray []string

func (r TlsSettingValueArray) ImplementsTlsSettingValueUnion() {}

// The tls setting value.
//
// Satisfied by [shared.UnionFloat], [shared.UnionString],
// [TlsSettingValueArrayParam].
type TlsSettingValueUnionParam interface {
	ImplementsTlsSettingValueUnionParam()
}

type TlsSettingValueArrayParam []string

func (r TlsSettingValueArrayParam) ImplementsTlsSettingValueUnionParam() {}

type ZoneHostnameSettingGetResponse struct {
	Result     []ZoneHostnameSettingGetResponseResult   `json:"result"`
	ResultInfo ZoneHostnameSettingGetResponseResultInfo `json:"result_info"`
	JSON       zoneHostnameSettingGetResponseJSON       `json:"-"`
	APIResponseCollectionTlsCertificates
}

// zoneHostnameSettingGetResponseJSON contains the JSON metadata for the struct
// [ZoneHostnameSettingGetResponse]
type zoneHostnameSettingGetResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHostnameSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHostnameSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneHostnameSettingGetResponseResult struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string `json:"hostname"`
	// Deployment status for the given tls setting.
	Status string `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The tls setting value.
	Value TlsSettingValueUnion                     `json:"value"`
	JSON  zoneHostnameSettingGetResponseResultJSON `json:"-"`
}

// zoneHostnameSettingGetResponseResultJSON contains the JSON metadata for the
// struct [ZoneHostnameSettingGetResponseResult]
type zoneHostnameSettingGetResponseResultJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHostnameSettingGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHostnameSettingGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneHostnameSettingGetResponseResultInfo struct {
	Count      interface{} `json:"count"`
	Page       interface{} `json:"page"`
	PerPage    interface{} `json:"per_page"`
	TotalCount interface{} `json:"total_count"`
	// Total pages available of results
	TotalPages float64                                      `json:"total_pages"`
	JSON       zoneHostnameSettingGetResponseResultInfoJSON `json:"-"`
}

// zoneHostnameSettingGetResponseResultInfoJSON contains the JSON metadata for the
// struct [ZoneHostnameSettingGetResponseResultInfo]
type zoneHostnameSettingGetResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHostnameSettingGetResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHostnameSettingGetResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneHostnameSettingUpdateResponse struct {
	Result ZoneHostnameSettingUpdateResponseResult `json:"result"`
	JSON   zoneHostnameSettingUpdateResponseJSON   `json:"-"`
	APIResponseSingleTlsCertificates
}

// zoneHostnameSettingUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneHostnameSettingUpdateResponse]
type zoneHostnameSettingUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHostnameSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHostnameSettingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneHostnameSettingUpdateResponseResult struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string `json:"hostname"`
	// Deployment status for the given tls setting.
	Status string `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The tls setting value.
	Value TlsSettingValueUnion                        `json:"value"`
	JSON  zoneHostnameSettingUpdateResponseResultJSON `json:"-"`
}

// zoneHostnameSettingUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneHostnameSettingUpdateResponseResult]
type zoneHostnameSettingUpdateResponseResultJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHostnameSettingUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHostnameSettingUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneHostnameSettingDeleteResponse struct {
	Result ZoneHostnameSettingDeleteResponseResult `json:"result"`
	JSON   zoneHostnameSettingDeleteResponseJSON   `json:"-"`
	APIResponseSingleTlsCertificates
}

// zoneHostnameSettingDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneHostnameSettingDeleteResponse]
type zoneHostnameSettingDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHostnameSettingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHostnameSettingDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneHostnameSettingDeleteResponseResult struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string `json:"hostname"`
	// Deployment status for the given tls setting.
	Status string `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The tls setting value.
	Value TlsSettingValueUnion                        `json:"value"`
	JSON  zoneHostnameSettingDeleteResponseResultJSON `json:"-"`
}

// zoneHostnameSettingDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneHostnameSettingDeleteResponseResult]
type zoneHostnameSettingDeleteResponseResultJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHostnameSettingDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHostnameSettingDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneHostnameSettingUpdateParams struct {
	// The tls setting value.
	Value param.Field[TlsSettingValueUnionParam] `json:"value,required"`
}

func (r ZoneHostnameSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
