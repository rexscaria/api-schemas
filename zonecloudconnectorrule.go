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

// ZoneCloudConnectorRuleService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCloudConnectorRuleService] method instead.
type ZoneCloudConnectorRuleService struct {
	Options []option.RequestOption
}

// NewZoneCloudConnectorRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneCloudConnectorRuleService(opts ...option.RequestOption) (r *ZoneCloudConnectorRuleService) {
	r = &ZoneCloudConnectorRuleService{}
	r.Options = opts
	return
}

// Put Rules
func (r *ZoneCloudConnectorRuleService) Update(ctx context.Context, zoneID string, body ZoneCloudConnectorRuleUpdateParams, opts ...option.RequestOption) (res *ZoneCloudConnectorRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cloud_connector/rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Rules
func (r *ZoneCloudConnectorRuleService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneCloudConnectorRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cloud_connector/rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type APIResponseCloudConnector struct {
	Errors   []MessagesCloudConnectorItem `json:"errors,required"`
	Messages []MessagesCloudConnectorItem `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseCloudConnectorSuccess `json:"success,required"`
	JSON    apiResponseCloudConnectorJSON    `json:"-"`
}

// apiResponseCloudConnectorJSON contains the JSON metadata for the struct
// [APIResponseCloudConnector]
type apiResponseCloudConnectorJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCloudConnector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCloudConnectorJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseCloudConnectorSuccess bool

const (
	APIResponseCloudConnectorSuccessTrue APIResponseCloudConnectorSuccess = true
)

func (r APIResponseCloudConnectorSuccess) IsKnown() bool {
	switch r {
	case APIResponseCloudConnectorSuccessTrue:
		return true
	}
	return false
}

type MessagesCloudConnectorItem struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    messagesCloudConnectorItemJSON `json:"-"`
}

// messagesCloudConnectorItemJSON contains the JSON metadata for the struct
// [MessagesCloudConnectorItem]
type messagesCloudConnectorItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesCloudConnectorItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesCloudConnectorItemJSON) RawJSON() string {
	return r.raw
}

type RuleItem struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Expression  string `json:"expression"`
	// Parameters of Cloud Connector Rule
	Parameters RuleItemParameters `json:"parameters"`
	// Cloud Provider type
	Provider RuleItemProvider `json:"provider"`
	JSON     ruleItemJSON     `json:"-"`
}

// ruleItemJSON contains the JSON metadata for the struct [RuleItem]
type ruleItemJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Parameters  apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleItemJSON) RawJSON() string {
	return r.raw
}

// Parameters of Cloud Connector Rule
type RuleItemParameters struct {
	// Host to perform Cloud Connection to
	Host string                 `json:"host"`
	JSON ruleItemParametersJSON `json:"-"`
}

// ruleItemParametersJSON contains the JSON metadata for the struct
// [RuleItemParameters]
type ruleItemParametersJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleItemParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleItemParametersJSON) RawJSON() string {
	return r.raw
}

// Cloud Provider type
type RuleItemProvider string

const (
	RuleItemProviderAwsS3        RuleItemProvider = "aws_s3"
	RuleItemProviderR2           RuleItemProvider = "r2"
	RuleItemProviderGcpStorage   RuleItemProvider = "gcp_storage"
	RuleItemProviderAzureStorage RuleItemProvider = "azure_storage"
)

func (r RuleItemProvider) IsKnown() bool {
	switch r {
	case RuleItemProviderAwsS3, RuleItemProviderR2, RuleItemProviderGcpStorage, RuleItemProviderAzureStorage:
		return true
	}
	return false
}

type RuleItemParam struct {
	ID          param.Field[string] `json:"id"`
	Description param.Field[string] `json:"description"`
	Enabled     param.Field[bool]   `json:"enabled"`
	Expression  param.Field[string] `json:"expression"`
	// Parameters of Cloud Connector Rule
	Parameters param.Field[RuleItemParametersParam] `json:"parameters"`
	// Cloud Provider type
	Provider param.Field[RuleItemProvider] `json:"provider"`
}

func (r RuleItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Parameters of Cloud Connector Rule
type RuleItemParametersParam struct {
	// Host to perform Cloud Connection to
	Host param.Field[string] `json:"host"`
}

func (r RuleItemParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneCloudConnectorRuleUpdateResponse struct {
	// List of Cloud Connector rules
	Result []RuleItem                               `json:"result"`
	JSON   zoneCloudConnectorRuleUpdateResponseJSON `json:"-"`
	APIResponseCloudConnector
}

// zoneCloudConnectorRuleUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneCloudConnectorRuleUpdateResponse]
type zoneCloudConnectorRuleUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCloudConnectorRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCloudConnectorRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCloudConnectorRuleListResponse struct {
	// List of Cloud Connector rules
	Result []RuleItem                             `json:"result"`
	JSON   zoneCloudConnectorRuleListResponseJSON `json:"-"`
	APIResponseCloudConnector
}

// zoneCloudConnectorRuleListResponseJSON contains the JSON metadata for the struct
// [ZoneCloudConnectorRuleListResponse]
type zoneCloudConnectorRuleListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCloudConnectorRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCloudConnectorRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCloudConnectorRuleUpdateParams struct {
	// List of Cloud Connector rules
	Body []RuleItemParam `json:"body"`
}

func (r ZoneCloudConnectorRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
