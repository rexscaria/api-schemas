// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ConfigService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigService] method instead.
type ConfigService struct {
	Options []option.RequestOption
}

// NewConfigService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConfigService(opts ...option.RequestOption) (r *ConfigService) {
	r = &ConfigService{}
	r.Options = opts
	return
}

// List Configs
func (r *ConfigService) List(ctx context.Context, query ConfigListParams, opts ...option.RequestOption) (res *ConfigListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "configs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ConfigListResponse struct {
	Result  []ConfigListResponseResult `json:"result,required"`
	Success bool                       `json:"success,required"`
	JSON    configListResponseJSON     `json:"-"`
}

// configListResponseJSON contains the JSON metadata for the struct
// [ConfigListResponse]
type configListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configListResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigListResponseResult struct {
	ID        float64                      `json:"id,required"`
	AccountID string                       `json:"account_id,required"`
	Frequency float64                      `json:"frequency,required"`
	IPs       []string                     `json:"ips,required"`
	Name      string                       `json:"name,required"`
	JSON      configListResponseResultJSON `json:"-"`
}

// configListResponseResultJSON contains the JSON metadata for the struct
// [ConfigListResponseResult]
type configListResponseResultJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Frequency   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configListResponseResultJSON) RawJSON() string {
	return r.raw
}

type ConfigListParams struct {
	// Account ID
	AccountID param.Field[string] `query:"account_id"`
	// Page number
	Page param.Field[float64] `query:"page"`
	// Page Size
	PageSize param.Field[float64] `query:"pageSize"`
}

// URLQuery serializes [ConfigListParams]'s query parameters as `url.Values`.
func (r ConfigListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
