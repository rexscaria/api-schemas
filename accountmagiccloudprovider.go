// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountMagicCloudProviderService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicCloudProviderService] method instead.
type AccountMagicCloudProviderService struct {
	Options []option.RequestOption
}

// NewAccountMagicCloudProviderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountMagicCloudProviderService(opts ...option.RequestOption) (r *AccountMagicCloudProviderService) {
	r = &AccountMagicCloudProviderService{}
	r.Options = opts
	return
}

// Create a new Cloud Integration (Closed Beta)
func (r *AccountMagicCloudProviderService) New(ctx context.Context, accountID string, params AccountMagicCloudProviderNewParams, opts ...option.RequestOption) (res *AccountMagicCloudProviderNewResponse, err error) {
	if params.Forwarded.Present {
		opts = append(opts, option.WithHeader("forwarded", fmt.Sprintf("%s", params.Forwarded)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/providers", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Read a Cloud Integration (Closed Beta)
func (r *AccountMagicCloudProviderService) Get(ctx context.Context, accountID string, providerID string, query AccountMagicCloudProviderGetParams, opts ...option.RequestOption) (res *AccountMagicCloudProviderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if providerID == "" {
		err = errors.New("missing required provider_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/providers/%s", accountID, providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update a Cloud Integration (Closed Beta)
func (r *AccountMagicCloudProviderService) Update(ctx context.Context, accountID string, providerID string, body AccountMagicCloudProviderUpdateParams, opts ...option.RequestOption) (res *McnUpdateProviderResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if providerID == "" {
		err = errors.New("missing required provider_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/providers/%s", accountID, providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List Cloud Integrations (Closed Beta)
func (r *AccountMagicCloudProviderService) List(ctx context.Context, accountID string, query AccountMagicCloudProviderListParams, opts ...option.RequestOption) (res *AccountMagicCloudProviderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/providers", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a Cloud Integration (Closed Beta)
func (r *AccountMagicCloudProviderService) Delete(ctx context.Context, accountID string, providerID string, opts ...option.RequestOption) (res *AccountMagicCloudProviderDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if providerID == "" {
		err = errors.New("missing required provider_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/providers/%s", accountID, providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Run discovery for a Cloud Integration (Closed Beta)
func (r *AccountMagicCloudProviderService) Discover(ctx context.Context, accountID string, providerID string, body AccountMagicCloudProviderDiscoverParams, opts ...option.RequestOption) (res *McnGoodResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if providerID == "" {
		err = errors.New("missing required provider_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/providers/%s/discover", accountID, providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Run discovery for all Cloud Integrations in an account (Closed Beta)
func (r *AccountMagicCloudProviderService) DiscoverAll(ctx context.Context, accountID string, opts ...option.RequestOption) (res *McnGoodResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/providers/discover", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get initial configuration to complete Cloud Integration setup (Closed Beta)
func (r *AccountMagicCloudProviderService) GetSetupConfig(ctx context.Context, accountID string, providerID string, opts ...option.RequestOption) (res *AccountMagicCloudProviderGetSetupConfigResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if providerID == "" {
		err = errors.New("missing required provider_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/providers/%s/initial_setup", accountID, providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Cloud Integration (Closed Beta)
func (r *AccountMagicCloudProviderService) Patch(ctx context.Context, accountID string, providerID string, body AccountMagicCloudProviderPatchParams, opts ...option.RequestOption) (res *McnUpdateProviderResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if providerID == "" {
		err = errors.New("missing required provider_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/providers/%s", accountID, providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type McnCloudType string

const (
	McnCloudTypeAws        McnCloudType = "AWS"
	McnCloudTypeAzure      McnCloudType = "AZURE"
	McnCloudTypeGoogle     McnCloudType = "GOOGLE"
	McnCloudTypeCloudflare McnCloudType = "CLOUDFLARE"
)

func (r McnCloudType) IsKnown() bool {
	switch r {
	case McnCloudTypeAws, McnCloudTypeAzure, McnCloudTypeGoogle, McnCloudTypeCloudflare:
		return true
	}
	return false
}

type McnProvider struct {
	ID                     string                     `json:"id,required" format:"uuid"`
	CloudType              McnCloudType               `json:"cloud_type,required"`
	FriendlyName           string                     `json:"friendly_name,required"`
	LastUpdated            string                     `json:"last_updated,required"`
	LifecycleState         McnProviderLifecycleState  `json:"lifecycle_state,required"`
	State                  McnProviderDiscoveryStatus `json:"state,required"`
	StateV2                McnProviderDiscoveryStatus `json:"state_v2,required"`
	AwsArn                 string                     `json:"aws_arn"`
	AzureSubscriptionID    string                     `json:"azure_subscription_id"`
	AzureTenantID          string                     `json:"azure_tenant_id"`
	Description            string                     `json:"description"`
	GcpProjectID           string                     `json:"gcp_project_id"`
	GcpServiceAccountEmail string                     `json:"gcp_service_account_email"`
	Status                 McnProviderStatus          `json:"status"`
	JSON                   mcnProviderJSON            `json:"-"`
}

// mcnProviderJSON contains the JSON metadata for the struct [McnProvider]
type mcnProviderJSON struct {
	ID                     apijson.Field
	CloudType              apijson.Field
	FriendlyName           apijson.Field
	LastUpdated            apijson.Field
	LifecycleState         apijson.Field
	State                  apijson.Field
	StateV2                apijson.Field
	AwsArn                 apijson.Field
	AzureSubscriptionID    apijson.Field
	AzureTenantID          apijson.Field
	Description            apijson.Field
	GcpProjectID           apijson.Field
	GcpServiceAccountEmail apijson.Field
	Status                 apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *McnProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnProviderJSON) RawJSON() string {
	return r.raw
}

type McnProviderLifecycleState string

const (
	McnProviderLifecycleStateActive       McnProviderLifecycleState = "ACTIVE"
	McnProviderLifecycleStatePendingSetup McnProviderLifecycleState = "PENDING_SETUP"
	McnProviderLifecycleStateRetired      McnProviderLifecycleState = "RETIRED"
)

func (r McnProviderLifecycleState) IsKnown() bool {
	switch r {
	case McnProviderLifecycleStateActive, McnProviderLifecycleStatePendingSetup, McnProviderLifecycleStateRetired:
		return true
	}
	return false
}

type McnProviderStatus struct {
	DiscoveryProgress          McnProviderDiscoveryProgress `json:"discovery_progress,required"`
	DiscoveryProgressV2        McnProviderDiscoveryProgress `json:"discovery_progress_v2,required"`
	LastDiscoveryStatus        McnProviderDiscoveryStatus   `json:"last_discovery_status,required"`
	LastDiscoveryStatusV2      McnProviderDiscoveryStatus   `json:"last_discovery_status_v2,required"`
	Regions                    []string                     `json:"regions,required"`
	CredentialsGoodSince       string                       `json:"credentials_good_since"`
	CredentialsMissingSince    string                       `json:"credentials_missing_since"`
	CredentialsRejectedSince   string                       `json:"credentials_rejected_since"`
	DiscoveryMessage           string                       `json:"discovery_message"`
	DiscoveryMessageV2         string                       `json:"discovery_message_v2"`
	InUseBy                    []McnCloudPlatformClient     `json:"in_use_by"`
	LastDiscoveryCompletedAt   string                       `json:"last_discovery_completed_at"`
	LastDiscoveryCompletedAtV2 string                       `json:"last_discovery_completed_at_v2"`
	LastDiscoveryStartedAt     string                       `json:"last_discovery_started_at"`
	LastDiscoveryStartedAtV2   string                       `json:"last_discovery_started_at_v2"`
	LastUpdated                string                       `json:"last_updated"`
	JSON                       mcnProviderStatusJSON        `json:"-"`
}

// mcnProviderStatusJSON contains the JSON metadata for the struct
// [McnProviderStatus]
type mcnProviderStatusJSON struct {
	DiscoveryProgress          apijson.Field
	DiscoveryProgressV2        apijson.Field
	LastDiscoveryStatus        apijson.Field
	LastDiscoveryStatusV2      apijson.Field
	Regions                    apijson.Field
	CredentialsGoodSince       apijson.Field
	CredentialsMissingSince    apijson.Field
	CredentialsRejectedSince   apijson.Field
	DiscoveryMessage           apijson.Field
	DiscoveryMessageV2         apijson.Field
	InUseBy                    apijson.Field
	LastDiscoveryCompletedAt   apijson.Field
	LastDiscoveryCompletedAtV2 apijson.Field
	LastDiscoveryStartedAt     apijson.Field
	LastDiscoveryStartedAtV2   apijson.Field
	LastUpdated                apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *McnProviderStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnProviderStatusJSON) RawJSON() string {
	return r.raw
}

type McnProviderDiscoveryProgress struct {
	Done  int64                            `json:"done,required"`
	Total int64                            `json:"total,required"`
	Unit  string                           `json:"unit,required"`
	JSON  mcnProviderDiscoveryProgressJSON `json:"-"`
}

// mcnProviderDiscoveryProgressJSON contains the JSON metadata for the struct
// [McnProviderDiscoveryProgress]
type mcnProviderDiscoveryProgressJSON struct {
	Done        apijson.Field
	Total       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McnProviderDiscoveryProgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnProviderDiscoveryProgressJSON) RawJSON() string {
	return r.raw
}

type McnProviderDiscoveryStatus string

const (
	McnProviderDiscoveryStatusUnspecified McnProviderDiscoveryStatus = "UNSPECIFIED"
	McnProviderDiscoveryStatusPending     McnProviderDiscoveryStatus = "PENDING"
	McnProviderDiscoveryStatusDiscovering McnProviderDiscoveryStatus = "DISCOVERING"
	McnProviderDiscoveryStatusFailed      McnProviderDiscoveryStatus = "FAILED"
	McnProviderDiscoveryStatusSucceeded   McnProviderDiscoveryStatus = "SUCCEEDED"
)

func (r McnProviderDiscoveryStatus) IsKnown() bool {
	switch r {
	case McnProviderDiscoveryStatusUnspecified, McnProviderDiscoveryStatusPending, McnProviderDiscoveryStatusDiscovering, McnProviderDiscoveryStatusFailed, McnProviderDiscoveryStatusSucceeded:
		return true
	}
	return false
}

type McnUpdateProviderRequestParam struct {
	AwsArn                 param.Field[string] `json:"aws_arn"`
	AzureSubscriptionID    param.Field[string] `json:"azure_subscription_id"`
	AzureTenantID          param.Field[string] `json:"azure_tenant_id"`
	Description            param.Field[string] `json:"description"`
	FriendlyName           param.Field[string] `json:"friendly_name"`
	GcpProjectID           param.Field[string] `json:"gcp_project_id"`
	GcpServiceAccountEmail param.Field[string] `json:"gcp_service_account_email"`
}

func (r McnUpdateProviderRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type McnUpdateProviderResponse struct {
	Result McnProvider                   `json:"result"`
	JSON   mcnUpdateProviderResponseJSON `json:"-"`
	McnGoodResponse
}

// mcnUpdateProviderResponseJSON contains the JSON metadata for the struct
// [McnUpdateProviderResponse]
type mcnUpdateProviderResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McnUpdateProviderResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnUpdateProviderResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudProviderNewResponse struct {
	Result McnProvider                              `json:"result"`
	JSON   accountMagicCloudProviderNewResponseJSON `json:"-"`
	McnGoodResponse
}

// accountMagicCloudProviderNewResponseJSON contains the JSON metadata for the
// struct [AccountMagicCloudProviderNewResponse]
type accountMagicCloudProviderNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudProviderNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudProviderNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudProviderGetResponse struct {
	Result McnProvider                              `json:"result"`
	JSON   accountMagicCloudProviderGetResponseJSON `json:"-"`
	McnGoodResponse
}

// accountMagicCloudProviderGetResponseJSON contains the JSON metadata for the
// struct [AccountMagicCloudProviderGetResponse]
type accountMagicCloudProviderGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudProviderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudProviderGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudProviderListResponse struct {
	Result []McnProvider                             `json:"result"`
	JSON   accountMagicCloudProviderListResponseJSON `json:"-"`
	McnGoodResponse
}

// accountMagicCloudProviderListResponseJSON contains the JSON metadata for the
// struct [AccountMagicCloudProviderListResponse]
type accountMagicCloudProviderListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudProviderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudProviderListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudProviderDeleteResponse struct {
	Result AccountMagicCloudProviderDeleteResponseResult `json:"result"`
	JSON   accountMagicCloudProviderDeleteResponseJSON   `json:"-"`
	McnGoodResponse
}

// accountMagicCloudProviderDeleteResponseJSON contains the JSON metadata for the
// struct [AccountMagicCloudProviderDeleteResponse]
type accountMagicCloudProviderDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudProviderDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudProviderDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudProviderDeleteResponseResult struct {
	ID   string                                            `json:"id,required" format:"uuid"`
	JSON accountMagicCloudProviderDeleteResponseResultJSON `json:"-"`
}

// accountMagicCloudProviderDeleteResponseResultJSON contains the JSON metadata for
// the struct [AccountMagicCloudProviderDeleteResponseResult]
type accountMagicCloudProviderDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudProviderDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudProviderDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudProviderGetSetupConfigResponse struct {
	Result AccountMagicCloudProviderGetSetupConfigResponseResult `json:"result"`
	JSON   accountMagicCloudProviderGetSetupConfigResponseJSON   `json:"-"`
	McnGoodResponse
}

// accountMagicCloudProviderGetSetupConfigResponseJSON contains the JSON metadata
// for the struct [AccountMagicCloudProviderGetSetupConfigResponse]
type accountMagicCloudProviderGetSetupConfigResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudProviderGetSetupConfigResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudProviderGetSetupConfigResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudProviderGetSetupConfigResponseResult struct {
	ItemType               string                                                    `json:"item_type,required"`
	AwsTrustPolicy         string                                                    `json:"aws_trust_policy"`
	AzureConsentURL        string                                                    `json:"azure_consent_url"`
	IntegrationIdentityTag string                                                    `json:"integration_identity_tag"`
	TagCliCommand          string                                                    `json:"tag_cli_command"`
	JSON                   accountMagicCloudProviderGetSetupConfigResponseResultJSON `json:"-"`
	union                  AccountMagicCloudProviderGetSetupConfigResponseResultUnion
}

// accountMagicCloudProviderGetSetupConfigResponseResultJSON contains the JSON
// metadata for the struct [AccountMagicCloudProviderGetSetupConfigResponseResult]
type accountMagicCloudProviderGetSetupConfigResponseResultJSON struct {
	ItemType               apijson.Field
	AwsTrustPolicy         apijson.Field
	AzureConsentURL        apijson.Field
	IntegrationIdentityTag apijson.Field
	TagCliCommand          apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r accountMagicCloudProviderGetSetupConfigResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *AccountMagicCloudProviderGetSetupConfigResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = AccountMagicCloudProviderGetSetupConfigResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountMagicCloudProviderGetSetupConfigResponseResultUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccountMagicCloudProviderGetSetupConfigResponseResultMcnAwsTrustPolicy],
// [AccountMagicCloudProviderGetSetupConfigResponseResultMcnAzureSetup],
// [AccountMagicCloudProviderGetSetupConfigResponseResultMcnGcpSetup].
func (r AccountMagicCloudProviderGetSetupConfigResponseResult) AsUnion() AccountMagicCloudProviderGetSetupConfigResponseResultUnion {
	return r.union
}

// Union satisfied by
// [AccountMagicCloudProviderGetSetupConfigResponseResultMcnAwsTrustPolicy],
// [AccountMagicCloudProviderGetSetupConfigResponseResultMcnAzureSetup] or
// [AccountMagicCloudProviderGetSetupConfigResponseResultMcnGcpSetup].
type AccountMagicCloudProviderGetSetupConfigResponseResultUnion interface {
	implementsAccountMagicCloudProviderGetSetupConfigResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountMagicCloudProviderGetSetupConfigResponseResultUnion)(nil)).Elem(),
		"item_type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountMagicCloudProviderGetSetupConfigResponseResultMcnAwsTrustPolicy{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountMagicCloudProviderGetSetupConfigResponseResultMcnAzureSetup{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountMagicCloudProviderGetSetupConfigResponseResultMcnGcpSetup{}),
		},
	)
}

type AccountMagicCloudProviderGetSetupConfigResponseResultMcnAwsTrustPolicy struct {
	AwsTrustPolicy string                                                                     `json:"aws_trust_policy,required"`
	ItemType       string                                                                     `json:"item_type,required"`
	JSON           accountMagicCloudProviderGetSetupConfigResponseResultMcnAwsTrustPolicyJSON `json:"-"`
}

// accountMagicCloudProviderGetSetupConfigResponseResultMcnAwsTrustPolicyJSON
// contains the JSON metadata for the struct
// [AccountMagicCloudProviderGetSetupConfigResponseResultMcnAwsTrustPolicy]
type accountMagicCloudProviderGetSetupConfigResponseResultMcnAwsTrustPolicyJSON struct {
	AwsTrustPolicy apijson.Field
	ItemType       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountMagicCloudProviderGetSetupConfigResponseResultMcnAwsTrustPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudProviderGetSetupConfigResponseResultMcnAwsTrustPolicyJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicCloudProviderGetSetupConfigResponseResultMcnAwsTrustPolicy) implementsAccountMagicCloudProviderGetSetupConfigResponseResult() {
}

type AccountMagicCloudProviderGetSetupConfigResponseResultMcnAzureSetup struct {
	AzureConsentURL        string                                                                 `json:"azure_consent_url,required"`
	IntegrationIdentityTag string                                                                 `json:"integration_identity_tag,required"`
	ItemType               string                                                                 `json:"item_type,required"`
	TagCliCommand          string                                                                 `json:"tag_cli_command,required"`
	JSON                   accountMagicCloudProviderGetSetupConfigResponseResultMcnAzureSetupJSON `json:"-"`
}

// accountMagicCloudProviderGetSetupConfigResponseResultMcnAzureSetupJSON contains
// the JSON metadata for the struct
// [AccountMagicCloudProviderGetSetupConfigResponseResultMcnAzureSetup]
type accountMagicCloudProviderGetSetupConfigResponseResultMcnAzureSetupJSON struct {
	AzureConsentURL        apijson.Field
	IntegrationIdentityTag apijson.Field
	ItemType               apijson.Field
	TagCliCommand          apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountMagicCloudProviderGetSetupConfigResponseResultMcnAzureSetup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudProviderGetSetupConfigResponseResultMcnAzureSetupJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicCloudProviderGetSetupConfigResponseResultMcnAzureSetup) implementsAccountMagicCloudProviderGetSetupConfigResponseResult() {
}

type AccountMagicCloudProviderGetSetupConfigResponseResultMcnGcpSetup struct {
	IntegrationIdentityTag string                                                               `json:"integration_identity_tag,required"`
	ItemType               string                                                               `json:"item_type,required"`
	TagCliCommand          string                                                               `json:"tag_cli_command,required"`
	JSON                   accountMagicCloudProviderGetSetupConfigResponseResultMcnGcpSetupJSON `json:"-"`
}

// accountMagicCloudProviderGetSetupConfigResponseResultMcnGcpSetupJSON contains
// the JSON metadata for the struct
// [AccountMagicCloudProviderGetSetupConfigResponseResultMcnGcpSetup]
type accountMagicCloudProviderGetSetupConfigResponseResultMcnGcpSetupJSON struct {
	IntegrationIdentityTag apijson.Field
	ItemType               apijson.Field
	TagCliCommand          apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountMagicCloudProviderGetSetupConfigResponseResultMcnGcpSetup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudProviderGetSetupConfigResponseResultMcnGcpSetupJSON) RawJSON() string {
	return r.raw
}

func (r AccountMagicCloudProviderGetSetupConfigResponseResultMcnGcpSetup) implementsAccountMagicCloudProviderGetSetupConfigResponseResult() {
}

type AccountMagicCloudProviderNewParams struct {
	CloudType    param.Field[McnCloudType] `json:"cloud_type,required"`
	FriendlyName param.Field[string]       `json:"friendly_name,required"`
	Description  param.Field[string]       `json:"description"`
	Forwarded    param.Field[string]       `header:"forwarded"`
}

func (r AccountMagicCloudProviderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicCloudProviderGetParams struct {
	Status param.Field[bool] `query:"status"`
}

// URLQuery serializes [AccountMagicCloudProviderGetParams]'s query parameters as
// `url.Values`.
func (r AccountMagicCloudProviderGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountMagicCloudProviderUpdateParams struct {
	McnUpdateProviderRequest McnUpdateProviderRequestParam `json:"mcn_update_provider_request,required"`
}

func (r AccountMagicCloudProviderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.McnUpdateProviderRequest)
}

type AccountMagicCloudProviderListParams struct {
	Cloudflare param.Field[bool] `query:"cloudflare"`
	Desc       param.Field[bool] `query:"desc"`
	// one of ["updated_at", "id", "cloud_type", "name"]
	OrderBy param.Field[string] `query:"order_by"`
	Status  param.Field[bool]   `query:"status"`
}

// URLQuery serializes [AccountMagicCloudProviderListParams]'s query parameters as
// `url.Values`.
func (r AccountMagicCloudProviderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountMagicCloudProviderDiscoverParams struct {
	V2 param.Field[bool] `query:"v2"`
}

// URLQuery serializes [AccountMagicCloudProviderDiscoverParams]'s query parameters
// as `url.Values`.
func (r AccountMagicCloudProviderDiscoverParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountMagicCloudProviderPatchParams struct {
	McnUpdateProviderRequest McnUpdateProviderRequestParam `json:"mcn_update_provider_request,required"`
}

func (r AccountMagicCloudProviderPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.McnUpdateProviderRequest)
}
