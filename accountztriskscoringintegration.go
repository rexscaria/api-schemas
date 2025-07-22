// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountZtRiskScoringIntegrationService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountZtRiskScoringIntegrationService] method instead.
type AccountZtRiskScoringIntegrationService struct {
	Options []option.RequestOption
}

// NewAccountZtRiskScoringIntegrationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountZtRiskScoringIntegrationService(opts ...option.RequestOption) (r *AccountZtRiskScoringIntegrationService) {
	r = &AccountZtRiskScoringIntegrationService{}
	r.Options = opts
	return
}

// Create new risk score integration.
func (r *AccountZtRiskScoringIntegrationService) NewIntegration(ctx context.Context, accountID string, body AccountZtRiskScoringIntegrationNewIntegrationParams, opts ...option.RequestOption) (res *AccountZtRiskScoringIntegrationNewIntegrationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/integrations", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a risk score integration.
func (r *AccountZtRiskScoringIntegrationService) DeleteIntegration(ctx context.Context, accountID string, integrationID string, opts ...option.RequestOption) (res *AccountZtRiskScoringIntegrationDeleteIntegrationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if integrationID == "" {
		err = errors.New("missing required integration_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/integrations/%s", accountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List all risk score integrations for the account.
func (r *AccountZtRiskScoringIntegrationService) ListIntegrations(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountZtRiskScoringIntegrationListIntegrationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/integrations", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get risk score integration by id.
func (r *AccountZtRiskScoringIntegrationService) GetIntegration(ctx context.Context, accountID string, integrationID string, opts ...option.RequestOption) (res *AccountZtRiskScoringIntegrationGetIntegrationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if integrationID == "" {
		err = errors.New("missing required integration_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/integrations/%s", accountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get risk score integration by reference id.
func (r *AccountZtRiskScoringIntegrationService) GetIntegrationByReferenceID(ctx context.Context, accountID string, referenceID string, opts ...option.RequestOption) (res *AccountZtRiskScoringIntegrationGetIntegrationByReferenceIDResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if referenceID == "" {
		err = errors.New("missing required reference_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/integrations/reference_id/%s", accountID, referenceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Overwrite the reference_id, tenant_url, and active values with the ones provided
func (r *AccountZtRiskScoringIntegrationService) UpdateIntegration(ctx context.Context, accountID string, integrationID string, body AccountZtRiskScoringIntegrationUpdateIntegrationParams, opts ...option.RequestOption) (res *AccountZtRiskScoringIntegrationUpdateIntegrationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if integrationID == "" {
		err = errors.New("missing required integration_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/integrations/%s", accountID, integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type RiskScoreIntegration struct {
	// The id of the integration, a UUIDv4.
	ID string `json:"id,required" format:"uuid"`
	// The Cloudflare account tag.
	AccountTag string `json:"account_tag,required"`
	// Whether this integration is enabled and should export changes in risk score.
	Active bool `json:"active,required"`
	// When the integration was created in RFC3339 format.
	CreatedAt       time.Time                `json:"created_at,required" format:"date-time"`
	IntegrationType RiskScoreIntegrationType `json:"integration_type,required"`
	// A reference ID defined by the client. Should be set to the Access-Okta IDP
	// integration ID. Useful when the risk-score integration needs to be associated
	// with a secondary asset and recalled using that ID.
	ReferenceID string `json:"reference_id,required"`
	// The base URL for the tenant. E.g. "https://tenant.okta.com"
	TenantURL string `json:"tenant_url,required"`
	// The URL for the Shared Signals Framework configuration, e.g.
	// "/.well-known/sse-configuration/{integration_uuid}/".
	// https://openid.net/specs/openid-sse-framework-1_0.html#rfc.section.6.2.1
	WellKnownURL string                   `json:"well_known_url,required"`
	JSON         riskScoreIntegrationJSON `json:"-"`
}

// riskScoreIntegrationJSON contains the JSON metadata for the struct
// [RiskScoreIntegration]
type riskScoreIntegrationJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Active          apijson.Field
	CreatedAt       apijson.Field
	IntegrationType apijson.Field
	ReferenceID     apijson.Field
	TenantURL       apijson.Field
	WellKnownURL    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RiskScoreIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoreIntegrationJSON) RawJSON() string {
	return r.raw
}

type RiskScoreIntegrationType string

const (
	RiskScoreIntegrationTypeOkta RiskScoreIntegrationType = "Okta"
)

func (r RiskScoreIntegrationType) IsKnown() bool {
	switch r {
	case RiskScoreIntegrationTypeOkta:
		return true
	}
	return false
}

type AccountZtRiskScoringIntegrationNewIntegrationResponse struct {
	Result RiskScoreIntegration                                      `json:"result"`
	JSON   accountZtRiskScoringIntegrationNewIntegrationResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountZtRiskScoringIntegrationNewIntegrationResponseJSON contains the JSON
// metadata for the struct [AccountZtRiskScoringIntegrationNewIntegrationResponse]
type accountZtRiskScoringIntegrationNewIntegrationResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZtRiskScoringIntegrationNewIntegrationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringIntegrationNewIntegrationResponseJSON) RawJSON() string {
	return r.raw
}

type AccountZtRiskScoringIntegrationDeleteIntegrationResponse struct {
	Result interface{}                                                  `json:"result,nullable"`
	JSON   accountZtRiskScoringIntegrationDeleteIntegrationResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountZtRiskScoringIntegrationDeleteIntegrationResponseJSON contains the JSON
// metadata for the struct
// [AccountZtRiskScoringIntegrationDeleteIntegrationResponse]
type accountZtRiskScoringIntegrationDeleteIntegrationResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZtRiskScoringIntegrationDeleteIntegrationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringIntegrationDeleteIntegrationResponseJSON) RawJSON() string {
	return r.raw
}

type AccountZtRiskScoringIntegrationListIntegrationsResponse struct {
	Result []RiskScoreIntegration                                      `json:"result"`
	JSON   accountZtRiskScoringIntegrationListIntegrationsResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountZtRiskScoringIntegrationListIntegrationsResponseJSON contains the JSON
// metadata for the struct
// [AccountZtRiskScoringIntegrationListIntegrationsResponse]
type accountZtRiskScoringIntegrationListIntegrationsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZtRiskScoringIntegrationListIntegrationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringIntegrationListIntegrationsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountZtRiskScoringIntegrationGetIntegrationResponse struct {
	Result RiskScoreIntegration                                      `json:"result"`
	JSON   accountZtRiskScoringIntegrationGetIntegrationResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountZtRiskScoringIntegrationGetIntegrationResponseJSON contains the JSON
// metadata for the struct [AccountZtRiskScoringIntegrationGetIntegrationResponse]
type accountZtRiskScoringIntegrationGetIntegrationResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZtRiskScoringIntegrationGetIntegrationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringIntegrationGetIntegrationResponseJSON) RawJSON() string {
	return r.raw
}

type AccountZtRiskScoringIntegrationGetIntegrationByReferenceIDResponse struct {
	Result RiskScoreIntegration                                                   `json:"result"`
	JSON   accountZtRiskScoringIntegrationGetIntegrationByReferenceIDResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountZtRiskScoringIntegrationGetIntegrationByReferenceIDResponseJSON contains
// the JSON metadata for the struct
// [AccountZtRiskScoringIntegrationGetIntegrationByReferenceIDResponse]
type accountZtRiskScoringIntegrationGetIntegrationByReferenceIDResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZtRiskScoringIntegrationGetIntegrationByReferenceIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringIntegrationGetIntegrationByReferenceIDResponseJSON) RawJSON() string {
	return r.raw
}

type AccountZtRiskScoringIntegrationUpdateIntegrationResponse struct {
	Result RiskScoreIntegration                                         `json:"result"`
	JSON   accountZtRiskScoringIntegrationUpdateIntegrationResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountZtRiskScoringIntegrationUpdateIntegrationResponseJSON contains the JSON
// metadata for the struct
// [AccountZtRiskScoringIntegrationUpdateIntegrationResponse]
type accountZtRiskScoringIntegrationUpdateIntegrationResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZtRiskScoringIntegrationUpdateIntegrationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringIntegrationUpdateIntegrationResponseJSON) RawJSON() string {
	return r.raw
}

type AccountZtRiskScoringIntegrationNewIntegrationParams struct {
	IntegrationType param.Field[RiskScoreIntegrationType] `json:"integration_type,required"`
	// The base url of the tenant, e.g. "https://tenant.okta.com"
	TenantURL param.Field[string] `json:"tenant_url,required" format:"uri"`
	// A reference id that can be supplied by the client. Currently this should be set
	// to the Access-Okta IDP ID (a UUIDv4).
	// https://developers.cloudflare.com/api/operations/access-identity-providers-get-an-access-identity-provider
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AccountZtRiskScoringIntegrationNewIntegrationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountZtRiskScoringIntegrationUpdateIntegrationParams struct {
	// Whether this integration is enabled. If disabled, no risk changes will be
	// exported to the third-party.
	Active param.Field[bool] `json:"active,required"`
	// The base url of the tenant, e.g. "https://tenant.okta.com"
	TenantURL param.Field[string] `json:"tenant_url,required" format:"uri"`
	// A reference id that can be supplied by the client. Currently this should be set
	// to the Access-Okta IDP ID (a UUIDv4).
	// https://developers.cloudflare.com/api/operations/access-identity-providers-get-an-access-identity-provider
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r AccountZtRiskScoringIntegrationUpdateIntegrationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
