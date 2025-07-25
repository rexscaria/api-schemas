// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountMagicCloudOnrampService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicCloudOnrampService] method instead.
type AccountMagicCloudOnrampService struct {
	Options              []option.RequestOption
	MagicWanAddressSpace *AccountMagicCloudOnrampMagicWanAddressSpaceService
}

// NewAccountMagicCloudOnrampService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMagicCloudOnrampService(opts ...option.RequestOption) (r *AccountMagicCloudOnrampService) {
	r = &AccountMagicCloudOnrampService{}
	r.Options = opts
	r.MagicWanAddressSpace = NewAccountMagicCloudOnrampMagicWanAddressSpaceService(opts...)
	return
}

// Create a new On-ramp (Closed Beta).
func (r *AccountMagicCloudOnrampService) New(ctx context.Context, accountID string, params AccountMagicCloudOnrampNewParams, opts ...option.RequestOption) (res *AccountMagicCloudOnrampNewResponse, err error) {
	if params.Forwarded.Present {
		opts = append(opts, option.WithHeader("forwarded", fmt.Sprintf("%s", params.Forwarded)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/onramps", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Read an On-ramp (Closed Beta).
func (r *AccountMagicCloudOnrampService) Get(ctx context.Context, accountID string, onrampID string, query AccountMagicCloudOnrampGetParams, opts ...option.RequestOption) (res *AccountMagicCloudOnrampGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if onrampID == "" {
		err = errors.New("missing required onramp_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/onramps/%s", accountID, onrampID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update an On-ramp (Closed Beta).
func (r *AccountMagicCloudOnrampService) Update(ctx context.Context, accountID string, onrampID string, body AccountMagicCloudOnrampUpdateParams, opts ...option.RequestOption) (res *McnUpdateOnrampResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if onrampID == "" {
		err = errors.New("missing required onramp_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/onramps/%s", accountID, onrampID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List On-ramps (Closed Beta).
func (r *AccountMagicCloudOnrampService) List(ctx context.Context, accountID string, query AccountMagicCloudOnrampListParams, opts ...option.RequestOption) (res *AccountMagicCloudOnrampListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/onramps", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an On-ramp (Closed Beta).
func (r *AccountMagicCloudOnrampService) Delete(ctx context.Context, accountID string, onrampID string, body AccountMagicCloudOnrampDeleteParams, opts ...option.RequestOption) (res *AccountMagicCloudOnrampDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if onrampID == "" {
		err = errors.New("missing required onramp_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/onramps/%s", accountID, onrampID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Apply an On-ramp (Closed Beta).
func (r *AccountMagicCloudOnrampService) Apply(ctx context.Context, accountID string, onrampID string, opts ...option.RequestOption) (res *McnGoodResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if onrampID == "" {
		err = errors.New("missing required onramp_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/onramps/%s/apply", accountID, onrampID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Export an On-ramp to terraform ready file(s) (Closed Beta).
func (r *AccountMagicCloudOnrampService) Export(ctx context.Context, accountID string, onrampID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/zip")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if onrampID == "" {
		err = errors.New("missing required onramp_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/onramps/%s/export", accountID, onrampID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Update an On-ramp (Closed Beta).
func (r *AccountMagicCloudOnrampService) Patch(ctx context.Context, accountID string, onrampID string, body AccountMagicCloudOnrampPatchParams, opts ...option.RequestOption) (res *McnUpdateOnrampResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if onrampID == "" {
		err = errors.New("missing required onramp_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/onramps/%s", accountID, onrampID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Plan an On-ramp (Closed Beta).
func (r *AccountMagicCloudOnrampService) Plan(ctx context.Context, accountID string, onrampID string, opts ...option.RequestOption) (res *McnGoodResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if onrampID == "" {
		err = errors.New("missing required onramp_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/onramps/%s/plan", accountID, onrampID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type McnCost struct {
	Currency    string      `json:"currency,required"`
	MonthlyCost float64     `json:"monthly_cost,required"`
	JSON        mcnCostJSON `json:"-"`
}

// mcnCostJSON contains the JSON metadata for the struct [McnCost]
type mcnCostJSON struct {
	Currency    apijson.Field
	MonthlyCost apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McnCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnCostJSON) RawJSON() string {
	return r.raw
}

type McnCostDiff struct {
	Currency            string          `json:"currency,required"`
	CurrentMonthlyCost  float64         `json:"current_monthly_cost,required"`
	Diff                float64         `json:"diff,required"`
	ProposedMonthlyCost float64         `json:"proposed_monthly_cost,required"`
	JSON                mcnCostDiffJSON `json:"-"`
}

// mcnCostDiffJSON contains the JSON metadata for the struct [McnCostDiff]
type mcnCostDiffJSON struct {
	Currency            apijson.Field
	CurrentMonthlyCost  apijson.Field
	Diff                apijson.Field
	ProposedMonthlyCost apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *McnCostDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnCostDiffJSON) RawJSON() string {
	return r.raw
}

type McnError struct {
	Code             McnErrorCode   `json:"code,required"`
	Message          string         `json:"message,required"`
	DocumentationURL string         `json:"documentation_url"`
	Meta             McnErrorMeta   `json:"meta"`
	Source           McnErrorSource `json:"source"`
	JSON             mcnErrorJSON   `json:"-"`
}

// mcnErrorJSON contains the JSON metadata for the struct [McnError]
type mcnErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *McnError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnErrorJSON) RawJSON() string {
	return r.raw
}

type McnErrorCode int64

const (
	McnErrorCode1001   McnErrorCode = 1001
	McnErrorCode1002   McnErrorCode = 1002
	McnErrorCode1003   McnErrorCode = 1003
	McnErrorCode1004   McnErrorCode = 1004
	McnErrorCode1005   McnErrorCode = 1005
	McnErrorCode1006   McnErrorCode = 1006
	McnErrorCode1007   McnErrorCode = 1007
	McnErrorCode1008   McnErrorCode = 1008
	McnErrorCode1009   McnErrorCode = 1009
	McnErrorCode1010   McnErrorCode = 1010
	McnErrorCode1011   McnErrorCode = 1011
	McnErrorCode1012   McnErrorCode = 1012
	McnErrorCode1013   McnErrorCode = 1013
	McnErrorCode1014   McnErrorCode = 1014
	McnErrorCode1015   McnErrorCode = 1015
	McnErrorCode1016   McnErrorCode = 1016
	McnErrorCode1017   McnErrorCode = 1017
	McnErrorCode2001   McnErrorCode = 2001
	McnErrorCode2002   McnErrorCode = 2002
	McnErrorCode2003   McnErrorCode = 2003
	McnErrorCode2004   McnErrorCode = 2004
	McnErrorCode2005   McnErrorCode = 2005
	McnErrorCode2006   McnErrorCode = 2006
	McnErrorCode2007   McnErrorCode = 2007
	McnErrorCode2008   McnErrorCode = 2008
	McnErrorCode2009   McnErrorCode = 2009
	McnErrorCode2010   McnErrorCode = 2010
	McnErrorCode2011   McnErrorCode = 2011
	McnErrorCode2012   McnErrorCode = 2012
	McnErrorCode2013   McnErrorCode = 2013
	McnErrorCode2014   McnErrorCode = 2014
	McnErrorCode2015   McnErrorCode = 2015
	McnErrorCode2016   McnErrorCode = 2016
	McnErrorCode2017   McnErrorCode = 2017
	McnErrorCode2018   McnErrorCode = 2018
	McnErrorCode2019   McnErrorCode = 2019
	McnErrorCode2020   McnErrorCode = 2020
	McnErrorCode2021   McnErrorCode = 2021
	McnErrorCode2022   McnErrorCode = 2022
	McnErrorCode3001   McnErrorCode = 3001
	McnErrorCode3002   McnErrorCode = 3002
	McnErrorCode3003   McnErrorCode = 3003
	McnErrorCode3004   McnErrorCode = 3004
	McnErrorCode3005   McnErrorCode = 3005
	McnErrorCode3006   McnErrorCode = 3006
	McnErrorCode3007   McnErrorCode = 3007
	McnErrorCode4001   McnErrorCode = 4001
	McnErrorCode4002   McnErrorCode = 4002
	McnErrorCode4003   McnErrorCode = 4003
	McnErrorCode4004   McnErrorCode = 4004
	McnErrorCode4005   McnErrorCode = 4005
	McnErrorCode4006   McnErrorCode = 4006
	McnErrorCode4007   McnErrorCode = 4007
	McnErrorCode4008   McnErrorCode = 4008
	McnErrorCode4009   McnErrorCode = 4009
	McnErrorCode4010   McnErrorCode = 4010
	McnErrorCode4011   McnErrorCode = 4011
	McnErrorCode4012   McnErrorCode = 4012
	McnErrorCode4013   McnErrorCode = 4013
	McnErrorCode4014   McnErrorCode = 4014
	McnErrorCode4015   McnErrorCode = 4015
	McnErrorCode4016   McnErrorCode = 4016
	McnErrorCode4017   McnErrorCode = 4017
	McnErrorCode4018   McnErrorCode = 4018
	McnErrorCode4019   McnErrorCode = 4019
	McnErrorCode4020   McnErrorCode = 4020
	McnErrorCode4021   McnErrorCode = 4021
	McnErrorCode4022   McnErrorCode = 4022
	McnErrorCode4023   McnErrorCode = 4023
	McnErrorCode5001   McnErrorCode = 5001
	McnErrorCode5002   McnErrorCode = 5002
	McnErrorCode5003   McnErrorCode = 5003
	McnErrorCode5004   McnErrorCode = 5004
	McnErrorCode102000 McnErrorCode = 102000
	McnErrorCode102001 McnErrorCode = 102001
	McnErrorCode102002 McnErrorCode = 102002
	McnErrorCode102003 McnErrorCode = 102003
	McnErrorCode102004 McnErrorCode = 102004
	McnErrorCode102005 McnErrorCode = 102005
	McnErrorCode102006 McnErrorCode = 102006
	McnErrorCode102007 McnErrorCode = 102007
	McnErrorCode102008 McnErrorCode = 102008
	McnErrorCode102009 McnErrorCode = 102009
	McnErrorCode102010 McnErrorCode = 102010
	McnErrorCode102011 McnErrorCode = 102011
	McnErrorCode102012 McnErrorCode = 102012
	McnErrorCode102013 McnErrorCode = 102013
	McnErrorCode102014 McnErrorCode = 102014
	McnErrorCode102015 McnErrorCode = 102015
	McnErrorCode102016 McnErrorCode = 102016
	McnErrorCode102017 McnErrorCode = 102017
	McnErrorCode102018 McnErrorCode = 102018
	McnErrorCode102019 McnErrorCode = 102019
	McnErrorCode102020 McnErrorCode = 102020
	McnErrorCode102021 McnErrorCode = 102021
	McnErrorCode102022 McnErrorCode = 102022
	McnErrorCode102023 McnErrorCode = 102023
	McnErrorCode102024 McnErrorCode = 102024
	McnErrorCode102025 McnErrorCode = 102025
	McnErrorCode102026 McnErrorCode = 102026
	McnErrorCode102027 McnErrorCode = 102027
	McnErrorCode102028 McnErrorCode = 102028
	McnErrorCode102029 McnErrorCode = 102029
	McnErrorCode102030 McnErrorCode = 102030
	McnErrorCode102031 McnErrorCode = 102031
	McnErrorCode102032 McnErrorCode = 102032
	McnErrorCode102033 McnErrorCode = 102033
	McnErrorCode102034 McnErrorCode = 102034
	McnErrorCode102035 McnErrorCode = 102035
	McnErrorCode102036 McnErrorCode = 102036
	McnErrorCode102037 McnErrorCode = 102037
	McnErrorCode102038 McnErrorCode = 102038
	McnErrorCode102039 McnErrorCode = 102039
	McnErrorCode102040 McnErrorCode = 102040
	McnErrorCode102041 McnErrorCode = 102041
	McnErrorCode102042 McnErrorCode = 102042
	McnErrorCode102043 McnErrorCode = 102043
	McnErrorCode102044 McnErrorCode = 102044
	McnErrorCode102045 McnErrorCode = 102045
	McnErrorCode102046 McnErrorCode = 102046
	McnErrorCode102047 McnErrorCode = 102047
	McnErrorCode102048 McnErrorCode = 102048
	McnErrorCode102049 McnErrorCode = 102049
	McnErrorCode102050 McnErrorCode = 102050
	McnErrorCode102051 McnErrorCode = 102051
	McnErrorCode102052 McnErrorCode = 102052
	McnErrorCode102053 McnErrorCode = 102053
	McnErrorCode102054 McnErrorCode = 102054
	McnErrorCode102055 McnErrorCode = 102055
	McnErrorCode102056 McnErrorCode = 102056
	McnErrorCode102057 McnErrorCode = 102057
	McnErrorCode102058 McnErrorCode = 102058
	McnErrorCode102059 McnErrorCode = 102059
	McnErrorCode102060 McnErrorCode = 102060
	McnErrorCode102061 McnErrorCode = 102061
	McnErrorCode102062 McnErrorCode = 102062
	McnErrorCode102063 McnErrorCode = 102063
	McnErrorCode102064 McnErrorCode = 102064
	McnErrorCode102065 McnErrorCode = 102065
	McnErrorCode102066 McnErrorCode = 102066
	McnErrorCode103001 McnErrorCode = 103001
	McnErrorCode103002 McnErrorCode = 103002
	McnErrorCode103003 McnErrorCode = 103003
	McnErrorCode103004 McnErrorCode = 103004
	McnErrorCode103005 McnErrorCode = 103005
	McnErrorCode103006 McnErrorCode = 103006
	McnErrorCode103007 McnErrorCode = 103007
	McnErrorCode103008 McnErrorCode = 103008
)

func (r McnErrorCode) IsKnown() bool {
	switch r {
	case McnErrorCode1001, McnErrorCode1002, McnErrorCode1003, McnErrorCode1004, McnErrorCode1005, McnErrorCode1006, McnErrorCode1007, McnErrorCode1008, McnErrorCode1009, McnErrorCode1010, McnErrorCode1011, McnErrorCode1012, McnErrorCode1013, McnErrorCode1014, McnErrorCode1015, McnErrorCode1016, McnErrorCode1017, McnErrorCode2001, McnErrorCode2002, McnErrorCode2003, McnErrorCode2004, McnErrorCode2005, McnErrorCode2006, McnErrorCode2007, McnErrorCode2008, McnErrorCode2009, McnErrorCode2010, McnErrorCode2011, McnErrorCode2012, McnErrorCode2013, McnErrorCode2014, McnErrorCode2015, McnErrorCode2016, McnErrorCode2017, McnErrorCode2018, McnErrorCode2019, McnErrorCode2020, McnErrorCode2021, McnErrorCode2022, McnErrorCode3001, McnErrorCode3002, McnErrorCode3003, McnErrorCode3004, McnErrorCode3005, McnErrorCode3006, McnErrorCode3007, McnErrorCode4001, McnErrorCode4002, McnErrorCode4003, McnErrorCode4004, McnErrorCode4005, McnErrorCode4006, McnErrorCode4007, McnErrorCode4008, McnErrorCode4009, McnErrorCode4010, McnErrorCode4011, McnErrorCode4012, McnErrorCode4013, McnErrorCode4014, McnErrorCode4015, McnErrorCode4016, McnErrorCode4017, McnErrorCode4018, McnErrorCode4019, McnErrorCode4020, McnErrorCode4021, McnErrorCode4022, McnErrorCode4023, McnErrorCode5001, McnErrorCode5002, McnErrorCode5003, McnErrorCode5004, McnErrorCode102000, McnErrorCode102001, McnErrorCode102002, McnErrorCode102003, McnErrorCode102004, McnErrorCode102005, McnErrorCode102006, McnErrorCode102007, McnErrorCode102008, McnErrorCode102009, McnErrorCode102010, McnErrorCode102011, McnErrorCode102012, McnErrorCode102013, McnErrorCode102014, McnErrorCode102015, McnErrorCode102016, McnErrorCode102017, McnErrorCode102018, McnErrorCode102019, McnErrorCode102020, McnErrorCode102021, McnErrorCode102022, McnErrorCode102023, McnErrorCode102024, McnErrorCode102025, McnErrorCode102026, McnErrorCode102027, McnErrorCode102028, McnErrorCode102029, McnErrorCode102030, McnErrorCode102031, McnErrorCode102032, McnErrorCode102033, McnErrorCode102034, McnErrorCode102035, McnErrorCode102036, McnErrorCode102037, McnErrorCode102038, McnErrorCode102039, McnErrorCode102040, McnErrorCode102041, McnErrorCode102042, McnErrorCode102043, McnErrorCode102044, McnErrorCode102045, McnErrorCode102046, McnErrorCode102047, McnErrorCode102048, McnErrorCode102049, McnErrorCode102050, McnErrorCode102051, McnErrorCode102052, McnErrorCode102053, McnErrorCode102054, McnErrorCode102055, McnErrorCode102056, McnErrorCode102057, McnErrorCode102058, McnErrorCode102059, McnErrorCode102060, McnErrorCode102061, McnErrorCode102062, McnErrorCode102063, McnErrorCode102064, McnErrorCode102065, McnErrorCode102066, McnErrorCode103001, McnErrorCode103002, McnErrorCode103003, McnErrorCode103004, McnErrorCode103005, McnErrorCode103006, McnErrorCode103007, McnErrorCode103008:
		return true
	}
	return false
}

type McnErrorMeta struct {
	L10nKey       string           `json:"l10n_key"`
	LoggableError string           `json:"loggable_error"`
	TemplateData  interface{}      `json:"template_data"`
	TraceID       string           `json:"trace_id"`
	JSON          mcnErrorMetaJSON `json:"-"`
}

// mcnErrorMetaJSON contains the JSON metadata for the struct [McnErrorMeta]
type mcnErrorMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *McnErrorMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnErrorMetaJSON) RawJSON() string {
	return r.raw
}

type McnErrorSource struct {
	Parameter           string             `json:"parameter"`
	ParameterValueIndex int64              `json:"parameter_value_index"`
	Pointer             string             `json:"pointer"`
	JSON                mcnErrorSourceJSON `json:"-"`
}

// mcnErrorSourceJSON contains the JSON metadata for the struct [McnErrorSource]
type mcnErrorSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *McnErrorSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnErrorSourceJSON) RawJSON() string {
	return r.raw
}

type McnGoodResponse struct {
	Errors   []McnError          `json:"errors,required"`
	Messages []McnError          `json:"messages,required"`
	Success  bool                `json:"success,required"`
	JSON     mcnGoodResponseJSON `json:"-"`
}

// mcnGoodResponseJSON contains the JSON metadata for the struct [McnGoodResponse]
type mcnGoodResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McnGoodResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnGoodResponseJSON) RawJSON() string {
	return r.raw
}

type McnOnramp struct {
	ID                            string                        `json:"id,required" format:"uuid"`
	CloudType                     McnOnrampCloudType            `json:"cloud_type,required"`
	InstallRoutesInCloud          bool                          `json:"install_routes_in_cloud,required"`
	InstallRoutesInMagicWan       bool                          `json:"install_routes_in_magic_wan,required"`
	Name                          string                        `json:"name,required"`
	Type                          McnOnrampType                 `json:"type,required"`
	UpdatedAt                     string                        `json:"updated_at,required"`
	AttachedHubs                  []string                      `json:"attached_hubs" format:"uuid"`
	AttachedVpcs                  []string                      `json:"attached_vpcs" format:"uuid"`
	Description                   string                        `json:"description"`
	Hub                           string                        `json:"hub" format:"uuid"`
	LastAppliedAt                 string                        `json:"last_applied_at"`
	LastExportedAt                string                        `json:"last_exported_at"`
	LastPlannedAt                 string                        `json:"last_planned_at"`
	ManageHubToHubAttachments     bool                          `json:"manage_hub_to_hub_attachments"`
	ManageVpcToHubAttachments     bool                          `json:"manage_vpc_to_hub_attachments"`
	PlannedMonthlyCostEstimate    McnCostDiff                   `json:"planned_monthly_cost_estimate"`
	PlannedResources              []McnOnrampPlannedResource    `json:"planned_resources"`
	PlannedResourcesUnavailable   bool                          `json:"planned_resources_unavailable"`
	PostApplyMonthlyCostEstimate  McnCost                       `json:"post_apply_monthly_cost_estimate"`
	PostApplyResources            map[string]McnResourceDetails `json:"post_apply_resources"`
	PostApplyResourcesUnavailable bool                          `json:"post_apply_resources_unavailable"`
	Region                        string                        `json:"region"`
	Status                        McnOnrampStatus               `json:"status"`
	Vpc                           string                        `json:"vpc" format:"uuid"`
	VpcsByID                      map[string]McnResourceDetails `json:"vpcs_by_id"`
	// The list of vpc IDs for which resource details failed to generate.
	VpcsByIDUnavailable []string      `json:"vpcs_by_id_unavailable" format:"uuid"`
	JSON                mcnOnrampJSON `json:"-"`
}

// mcnOnrampJSON contains the JSON metadata for the struct [McnOnramp]
type mcnOnrampJSON struct {
	ID                            apijson.Field
	CloudType                     apijson.Field
	InstallRoutesInCloud          apijson.Field
	InstallRoutesInMagicWan       apijson.Field
	Name                          apijson.Field
	Type                          apijson.Field
	UpdatedAt                     apijson.Field
	AttachedHubs                  apijson.Field
	AttachedVpcs                  apijson.Field
	Description                   apijson.Field
	Hub                           apijson.Field
	LastAppliedAt                 apijson.Field
	LastExportedAt                apijson.Field
	LastPlannedAt                 apijson.Field
	ManageHubToHubAttachments     apijson.Field
	ManageVpcToHubAttachments     apijson.Field
	PlannedMonthlyCostEstimate    apijson.Field
	PlannedResources              apijson.Field
	PlannedResourcesUnavailable   apijson.Field
	PostApplyMonthlyCostEstimate  apijson.Field
	PostApplyResources            apijson.Field
	PostApplyResourcesUnavailable apijson.Field
	Region                        apijson.Field
	Status                        apijson.Field
	Vpc                           apijson.Field
	VpcsByID                      apijson.Field
	VpcsByIDUnavailable           apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *McnOnramp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnOnrampJSON) RawJSON() string {
	return r.raw
}

type McnOnrampPlannedResource struct {
	Diff                    McnYamlDiff                            `json:"diff,required"`
	KeysRequireReplace      []string                               `json:"keys_require_replace,required"`
	MonthlyCostEstimateDiff McnCostDiff                            `json:"monthly_cost_estimate_diff,required"`
	PlannedAction           McnOnrampPlannedResourcesPlannedAction `json:"planned_action,required"`
	Resource                McnResourcePreview                     `json:"resource,required"`
	JSON                    mcnOnrampPlannedResourceJSON           `json:"-"`
}

// mcnOnrampPlannedResourceJSON contains the JSON metadata for the struct
// [McnOnrampPlannedResource]
type mcnOnrampPlannedResourceJSON struct {
	Diff                    apijson.Field
	KeysRequireReplace      apijson.Field
	MonthlyCostEstimateDiff apijson.Field
	PlannedAction           apijson.Field
	Resource                apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *McnOnrampPlannedResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnOnrampPlannedResourceJSON) RawJSON() string {
	return r.raw
}

type McnOnrampPlannedResourcesPlannedAction string

const (
	McnOnrampPlannedResourcesPlannedActionNoOp    McnOnrampPlannedResourcesPlannedAction = "no_op"
	McnOnrampPlannedResourcesPlannedActionCreate  McnOnrampPlannedResourcesPlannedAction = "create"
	McnOnrampPlannedResourcesPlannedActionUpdate  McnOnrampPlannedResourcesPlannedAction = "update"
	McnOnrampPlannedResourcesPlannedActionReplace McnOnrampPlannedResourcesPlannedAction = "replace"
	McnOnrampPlannedResourcesPlannedActionDestroy McnOnrampPlannedResourcesPlannedAction = "destroy"
)

func (r McnOnrampPlannedResourcesPlannedAction) IsKnown() bool {
	switch r {
	case McnOnrampPlannedResourcesPlannedActionNoOp, McnOnrampPlannedResourcesPlannedActionCreate, McnOnrampPlannedResourcesPlannedActionUpdate, McnOnrampPlannedResourcesPlannedActionReplace, McnOnrampPlannedResourcesPlannedActionDestroy:
		return true
	}
	return false
}

type McnOnrampStatus struct {
	ApplyProgress   McnOnrampStatusApplyProgress  `json:"apply_progress,required"`
	LifecycleState  McnOnrampStatusLifecycleState `json:"lifecycle_state,required"`
	PlanProgress    McnOnrampStatusPlanProgress   `json:"plan_progress,required"`
	Routes          []string                      `json:"routes,required" format:"uuid"`
	Tunnels         []string                      `json:"tunnels,required" format:"uuid"`
	LifecycleErrors map[string]McnError           `json:"lifecycle_errors"`
	JSON            mcnOnrampStatusJSON           `json:"-"`
}

// mcnOnrampStatusJSON contains the JSON metadata for the struct [McnOnrampStatus]
type mcnOnrampStatusJSON struct {
	ApplyProgress   apijson.Field
	LifecycleState  apijson.Field
	PlanProgress    apijson.Field
	Routes          apijson.Field
	Tunnels         apijson.Field
	LifecycleErrors apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *McnOnrampStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnOnrampStatusJSON) RawJSON() string {
	return r.raw
}

type McnOnrampStatusApplyProgress struct {
	Done  int64                            `json:"done,required"`
	Total int64                            `json:"total,required"`
	JSON  mcnOnrampStatusApplyProgressJSON `json:"-"`
}

// mcnOnrampStatusApplyProgressJSON contains the JSON metadata for the struct
// [McnOnrampStatusApplyProgress]
type mcnOnrampStatusApplyProgressJSON struct {
	Done        apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McnOnrampStatusApplyProgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnOnrampStatusApplyProgressJSON) RawJSON() string {
	return r.raw
}

type McnOnrampStatusLifecycleState string

const (
	McnOnrampStatusLifecycleStateOnrampNeedsApply      McnOnrampStatusLifecycleState = "OnrampNeedsApply"
	McnOnrampStatusLifecycleStateOnrampPendingPlan     McnOnrampStatusLifecycleState = "OnrampPendingPlan"
	McnOnrampStatusLifecycleStateOnrampPlanning        McnOnrampStatusLifecycleState = "OnrampPlanning"
	McnOnrampStatusLifecycleStateOnrampPlanFailed      McnOnrampStatusLifecycleState = "OnrampPlanFailed"
	McnOnrampStatusLifecycleStateOnrampPendingApproval McnOnrampStatusLifecycleState = "OnrampPendingApproval"
	McnOnrampStatusLifecycleStateOnrampPendingApply    McnOnrampStatusLifecycleState = "OnrampPendingApply"
	McnOnrampStatusLifecycleStateOnrampApplying        McnOnrampStatusLifecycleState = "OnrampApplying"
	McnOnrampStatusLifecycleStateOnrampApplyFailed     McnOnrampStatusLifecycleState = "OnrampApplyFailed"
	McnOnrampStatusLifecycleStateOnrampActive          McnOnrampStatusLifecycleState = "OnrampActive"
	McnOnrampStatusLifecycleStateOnrampPendingDestroy  McnOnrampStatusLifecycleState = "OnrampPendingDestroy"
	McnOnrampStatusLifecycleStateOnrampDestroying      McnOnrampStatusLifecycleState = "OnrampDestroying"
	McnOnrampStatusLifecycleStateOnrampDestroyFailed   McnOnrampStatusLifecycleState = "OnrampDestroyFailed"
)

func (r McnOnrampStatusLifecycleState) IsKnown() bool {
	switch r {
	case McnOnrampStatusLifecycleStateOnrampNeedsApply, McnOnrampStatusLifecycleStateOnrampPendingPlan, McnOnrampStatusLifecycleStateOnrampPlanning, McnOnrampStatusLifecycleStateOnrampPlanFailed, McnOnrampStatusLifecycleStateOnrampPendingApproval, McnOnrampStatusLifecycleStateOnrampPendingApply, McnOnrampStatusLifecycleStateOnrampApplying, McnOnrampStatusLifecycleStateOnrampApplyFailed, McnOnrampStatusLifecycleStateOnrampActive, McnOnrampStatusLifecycleStateOnrampPendingDestroy, McnOnrampStatusLifecycleStateOnrampDestroying, McnOnrampStatusLifecycleStateOnrampDestroyFailed:
		return true
	}
	return false
}

type McnOnrampStatusPlanProgress struct {
	Done  int64                           `json:"done,required"`
	Total int64                           `json:"total,required"`
	JSON  mcnOnrampStatusPlanProgressJSON `json:"-"`
}

// mcnOnrampStatusPlanProgressJSON contains the JSON metadata for the struct
// [McnOnrampStatusPlanProgress]
type mcnOnrampStatusPlanProgressJSON struct {
	Done        apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McnOnrampStatusPlanProgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnOnrampStatusPlanProgressJSON) RawJSON() string {
	return r.raw
}

type McnOnrampCloudType string

const (
	McnOnrampCloudTypeAws    McnOnrampCloudType = "AWS"
	McnOnrampCloudTypeAzure  McnOnrampCloudType = "AZURE"
	McnOnrampCloudTypeGoogle McnOnrampCloudType = "GOOGLE"
)

func (r McnOnrampCloudType) IsKnown() bool {
	switch r {
	case McnOnrampCloudTypeAws, McnOnrampCloudTypeAzure, McnOnrampCloudTypeGoogle:
		return true
	}
	return false
}

type McnOnrampType string

const (
	McnOnrampTypeOnrampTypeSingle McnOnrampType = "OnrampTypeSingle"
	McnOnrampTypeOnrampTypeHub    McnOnrampType = "OnrampTypeHub"
)

func (r McnOnrampType) IsKnown() bool {
	switch r {
	case McnOnrampTypeOnrampTypeSingle, McnOnrampTypeOnrampTypeHub:
		return true
	}
	return false
}

type McnResourcePreview struct {
	ID           string                 `json:"id,required" format:"uuid"`
	CloudType    McnCloudType           `json:"cloud_type,required"`
	Detail       string                 `json:"detail,required"`
	Name         string                 `json:"name,required"`
	ResourceType McnResourceType        `json:"resource_type,required"`
	Title        string                 `json:"title,required"`
	JSON         mcnResourcePreviewJSON `json:"-"`
}

// mcnResourcePreviewJSON contains the JSON metadata for the struct
// [McnResourcePreview]
type mcnResourcePreviewJSON struct {
	ID           apijson.Field
	CloudType    apijson.Field
	Detail       apijson.Field
	Name         apijson.Field
	ResourceType apijson.Field
	Title        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *McnResourcePreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnResourcePreviewJSON) RawJSON() string {
	return r.raw
}

type McnUpdateOnrampRequestParam struct {
	AttachedHubs              param.Field[[]string] `json:"attached_hubs" format:"uuid"`
	AttachedVpcs              param.Field[[]string] `json:"attached_vpcs" format:"uuid"`
	Description               param.Field[string]   `json:"description"`
	InstallRoutesInCloud      param.Field[bool]     `json:"install_routes_in_cloud"`
	InstallRoutesInMagicWan   param.Field[bool]     `json:"install_routes_in_magic_wan"`
	ManageHubToHubAttachments param.Field[bool]     `json:"manage_hub_to_hub_attachments"`
	ManageVpcToHubAttachments param.Field[bool]     `json:"manage_vpc_to_hub_attachments"`
	Name                      param.Field[string]   `json:"name"`
	Vpc                       param.Field[string]   `json:"vpc" format:"uuid"`
}

func (r McnUpdateOnrampRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type McnUpdateOnrampResponse struct {
	Errors   []McnError                  `json:"errors,required"`
	Messages []McnError                  `json:"messages,required"`
	Result   McnOnramp                   `json:"result,required"`
	Success  bool                        `json:"success,required"`
	JSON     mcnUpdateOnrampResponseJSON `json:"-"`
}

// mcnUpdateOnrampResponseJSON contains the JSON metadata for the struct
// [McnUpdateOnrampResponse]
type mcnUpdateOnrampResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McnUpdateOnrampResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnUpdateOnrampResponseJSON) RawJSON() string {
	return r.raw
}

type McnYamlDiff struct {
	Diff             string          `json:"diff,required"`
	LeftDescription  string          `json:"left_description,required"`
	LeftYaml         string          `json:"left_yaml,required"`
	RightDescription string          `json:"right_description,required"`
	RightYaml        string          `json:"right_yaml,required"`
	JSON             mcnYamlDiffJSON `json:"-"`
}

// mcnYamlDiffJSON contains the JSON metadata for the struct [McnYamlDiff]
type mcnYamlDiffJSON struct {
	Diff             apijson.Field
	LeftDescription  apijson.Field
	LeftYaml         apijson.Field
	RightDescription apijson.Field
	RightYaml        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *McnYamlDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnYamlDiffJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudOnrampNewResponse struct {
	Errors   []McnError                             `json:"errors,required"`
	Messages []McnError                             `json:"messages,required"`
	Result   McnOnramp                              `json:"result,required"`
	Success  bool                                   `json:"success,required"`
	JSON     accountMagicCloudOnrampNewResponseJSON `json:"-"`
}

// accountMagicCloudOnrampNewResponseJSON contains the JSON metadata for the struct
// [AccountMagicCloudOnrampNewResponse]
type accountMagicCloudOnrampNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudOnrampNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudOnrampNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudOnrampGetResponse struct {
	Errors   []McnError                             `json:"errors,required"`
	Messages []McnError                             `json:"messages,required"`
	Result   McnOnramp                              `json:"result,required"`
	Success  bool                                   `json:"success,required"`
	JSON     accountMagicCloudOnrampGetResponseJSON `json:"-"`
}

// accountMagicCloudOnrampGetResponseJSON contains the JSON metadata for the struct
// [AccountMagicCloudOnrampGetResponse]
type accountMagicCloudOnrampGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudOnrampGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudOnrampGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudOnrampListResponse struct {
	Errors   []McnError                              `json:"errors,required"`
	Messages []McnError                              `json:"messages,required"`
	Result   []McnOnramp                             `json:"result,required"`
	Success  bool                                    `json:"success,required"`
	JSON     accountMagicCloudOnrampListResponseJSON `json:"-"`
}

// accountMagicCloudOnrampListResponseJSON contains the JSON metadata for the
// struct [AccountMagicCloudOnrampListResponse]
type accountMagicCloudOnrampListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudOnrampListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudOnrampListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudOnrampDeleteResponse struct {
	Errors   []McnError                                  `json:"errors,required"`
	Messages []McnError                                  `json:"messages,required"`
	Result   AccountMagicCloudOnrampDeleteResponseResult `json:"result,required"`
	Success  bool                                        `json:"success,required"`
	JSON     accountMagicCloudOnrampDeleteResponseJSON   `json:"-"`
}

// accountMagicCloudOnrampDeleteResponseJSON contains the JSON metadata for the
// struct [AccountMagicCloudOnrampDeleteResponse]
type accountMagicCloudOnrampDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudOnrampDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudOnrampDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudOnrampDeleteResponseResult struct {
	ID   string                                          `json:"id,required" format:"uuid"`
	JSON accountMagicCloudOnrampDeleteResponseResultJSON `json:"-"`
}

// accountMagicCloudOnrampDeleteResponseResultJSON contains the JSON metadata for
// the struct [AccountMagicCloudOnrampDeleteResponseResult]
type accountMagicCloudOnrampDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudOnrampDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudOnrampDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudOnrampNewParams struct {
	CloudType                 param.Field[McnOnrampCloudType] `json:"cloud_type,required"`
	InstallRoutesInCloud      param.Field[bool]               `json:"install_routes_in_cloud,required"`
	InstallRoutesInMagicWan   param.Field[bool]               `json:"install_routes_in_magic_wan,required"`
	Name                      param.Field[string]             `json:"name,required"`
	Type                      param.Field[McnOnrampType]      `json:"type,required"`
	AdoptedHubID              param.Field[string]             `json:"adopted_hub_id" format:"uuid"`
	AttachedHubs              param.Field[[]string]           `json:"attached_hubs" format:"uuid"`
	AttachedVpcs              param.Field[[]string]           `json:"attached_vpcs" format:"uuid"`
	Description               param.Field[string]             `json:"description"`
	HubProviderID             param.Field[string]             `json:"hub_provider_id" format:"uuid"`
	ManageHubToHubAttachments param.Field[bool]               `json:"manage_hub_to_hub_attachments"`
	ManageVpcToHubAttachments param.Field[bool]               `json:"manage_vpc_to_hub_attachments"`
	Region                    param.Field[string]             `json:"region"`
	Vpc                       param.Field[string]             `json:"vpc" format:"uuid"`
	Forwarded                 param.Field[string]             `header:"forwarded"`
}

func (r AccountMagicCloudOnrampNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicCloudOnrampGetParams struct {
	PlannedResources   param.Field[bool] `query:"planned_resources"`
	PostApplyResources param.Field[bool] `query:"post_apply_resources"`
	Status             param.Field[bool] `query:"status"`
	Vpcs               param.Field[bool] `query:"vpcs"`
}

// URLQuery serializes [AccountMagicCloudOnrampGetParams]'s query parameters as
// `url.Values`.
func (r AccountMagicCloudOnrampGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountMagicCloudOnrampUpdateParams struct {
	McnUpdateOnrampRequest McnUpdateOnrampRequestParam `json:"mcn_update_onramp_request,required"`
}

func (r AccountMagicCloudOnrampUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.McnUpdateOnrampRequest)
}

type AccountMagicCloudOnrampListParams struct {
	Desc param.Field[bool] `query:"desc"`
	// One of ["updated_at", "id", "cloud_type", "name"].
	OrderBy param.Field[string] `query:"order_by"`
	Status  param.Field[bool]   `query:"status"`
	Vpcs    param.Field[bool]   `query:"vpcs"`
}

// URLQuery serializes [AccountMagicCloudOnrampListParams]'s query parameters as
// `url.Values`.
func (r AccountMagicCloudOnrampListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountMagicCloudOnrampDeleteParams struct {
	Destroy param.Field[bool] `query:"destroy"`
	Force   param.Field[bool] `query:"force"`
}

// URLQuery serializes [AccountMagicCloudOnrampDeleteParams]'s query parameters as
// `url.Values`.
func (r AccountMagicCloudOnrampDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountMagicCloudOnrampPatchParams struct {
	McnUpdateOnrampRequest McnUpdateOnrampRequestParam `json:"mcn_update_onramp_request,required"`
}

func (r AccountMagicCloudOnrampPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.McnUpdateOnrampRequest)
}
