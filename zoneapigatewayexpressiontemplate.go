// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneAPIGatewayExpressionTemplateService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAPIGatewayExpressionTemplateService] method instead.
type ZoneAPIGatewayExpressionTemplateService struct {
	Options []option.RequestOption
}

// NewZoneAPIGatewayExpressionTemplateService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneAPIGatewayExpressionTemplateService(opts ...option.RequestOption) (r *ZoneAPIGatewayExpressionTemplateService) {
	r = &ZoneAPIGatewayExpressionTemplateService{}
	r.Options = opts
	return
}

// Generate fallthrough WAF expression template from a set of API hosts
func (r *ZoneAPIGatewayExpressionTemplateService) GenerateFallthrough(ctx context.Context, zoneID string, body ZoneAPIGatewayExpressionTemplateGenerateFallthroughParams, opts ...option.RequestOption) (res *ZoneAPIGatewayExpressionTemplateGenerateFallthroughResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/expression-template/fallthrough", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneAPIGatewayExpressionTemplateGenerateFallthroughResponse struct {
	Result ZoneAPIGatewayExpressionTemplateGenerateFallthroughResponseResult `json:"result,required"`
	JSON   zoneAPIGatewayExpressionTemplateGenerateFallthroughResponseJSON   `json:"-"`
	APIResponseAPIShield
}

// zoneAPIGatewayExpressionTemplateGenerateFallthroughResponseJSON contains the
// JSON metadata for the struct
// [ZoneAPIGatewayExpressionTemplateGenerateFallthroughResponse]
type zoneAPIGatewayExpressionTemplateGenerateFallthroughResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayExpressionTemplateGenerateFallthroughResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayExpressionTemplateGenerateFallthroughResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayExpressionTemplateGenerateFallthroughResponseResult struct {
	// WAF Expression for fallthrough
	Expression string `json:"expression,required"`
	// Title for the expression
	Title string                                                                `json:"title,required"`
	JSON  zoneAPIGatewayExpressionTemplateGenerateFallthroughResponseResultJSON `json:"-"`
}

// zoneAPIGatewayExpressionTemplateGenerateFallthroughResponseResultJSON contains
// the JSON metadata for the struct
// [ZoneAPIGatewayExpressionTemplateGenerateFallthroughResponseResult]
type zoneAPIGatewayExpressionTemplateGenerateFallthroughResponseResultJSON struct {
	Expression  apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayExpressionTemplateGenerateFallthroughResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayExpressionTemplateGenerateFallthroughResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayExpressionTemplateGenerateFallthroughParams struct {
	// List of hosts to be targeted in the expression
	Hosts param.Field[[]string] `json:"hosts,required"`
}

func (r ZoneAPIGatewayExpressionTemplateGenerateFallthroughParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
