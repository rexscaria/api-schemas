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

// AccountAlertingV3PolicyService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAlertingV3PolicyService] method instead.
type AccountAlertingV3PolicyService struct {
	Options []option.RequestOption
}

// NewAccountAlertingV3PolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAlertingV3PolicyService(opts ...option.RequestOption) (r *AccountAlertingV3PolicyService) {
	r = &AccountAlertingV3PolicyService{}
	r.Options = opts
	return
}

// Creates a new Notification policy.
func (r *AccountAlertingV3PolicyService) New(ctx context.Context, accountID string, body AccountAlertingV3PolicyNewParams, opts ...option.RequestOption) (res *IDResponseAlerting, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details for a single policy.
func (r *AccountAlertingV3PolicyService) Get(ctx context.Context, accountID string, policyID string, opts ...option.RequestOption) (res *AccountAlertingV3PolicyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Notification policy.
func (r *AccountAlertingV3PolicyService) Update(ctx context.Context, accountID string, policyID string, body AccountAlertingV3PolicyUpdateParams, opts ...option.RequestOption) (res *IDResponseAlerting, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get a list of all Notification policies.
func (r *AccountAlertingV3PolicyService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAlertingV3PolicyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a Notification policy.
func (r *AccountAlertingV3PolicyService) Delete(ctx context.Context, accountID string, policyID string, opts ...option.RequestOption) (res *APIResponseCollectionAlerting, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type AlertFilters struct {
	// Usage depends on specific alert type
	Actions []string `json:"actions"`
	// Used for configuring radar_notification
	AffectedAsns []string `json:"affected_asns"`
	// Used for configuring incident_alert
	AffectedComponents []string `json:"affected_components"`
	// Used for configuring radar_notification
	AffectedLocations []string `json:"affected_locations"`
	// Used for configuring maintenance_event_notification
	AirportCode []string `json:"airport_code"`
	// Usage depends on specific alert type
	AlertTriggerPreferences []string `json:"alert_trigger_preferences"`
	// Usage depends on specific alert type
	AlertTriggerPreferencesValue []string `json:"alert_trigger_preferences_value"`
	// Used for configuring load_balancing_pool_enablement_alert
	Enabled []string `json:"enabled"`
	// Used for configuring pages_event_alert
	Environment []string `json:"environment"`
	// Used for configuring pages_event_alert
	Event []string `json:"event"`
	// Used for configuring load_balancing_health_alert
	EventSource []string `json:"event_source"`
	// Usage depends on specific alert type
	EventType []string `json:"event_type"`
	// Usage depends on specific alert type
	GroupBy []string `json:"group_by"`
	// Used for configuring health_check_status_notification
	HealthCheckID []string `json:"health_check_id"`
	// Used for configuring incident_alert
	IncidentImpact []AlertFiltersIncidentImpact `json:"incident_impact"`
	// Used for configuring stream_live_notifications
	InputID []string `json:"input_id"`
	// Used for configuring security_insights_alert
	InsightClass []string `json:"insight_class"`
	// Used for configuring billing_usage_alert
	Limit []string `json:"limit"`
	// Used for configuring logo_match_alert
	LogoTag []string `json:"logo_tag"`
	// Used for configuring advanced_ddos_attack_l4_alert
	MegabitsPerSecond []string `json:"megabits_per_second"`
	// Used for configuring load_balancing_health_alert
	NewHealth []string `json:"new_health"`
	// Used for configuring tunnel_health_event
	NewStatus []string `json:"new_status"`
	// Used for configuring advanced_ddos_attack_l4_alert
	PacketsPerSecond []string `json:"packets_per_second"`
	// Usage depends on specific alert type
	PoolID []string `json:"pool_id"`
	// Usage depends on specific alert type
	PopNames []string `json:"pop_names"`
	// Used for configuring billing_usage_alert
	Product []string `json:"product"`
	// Used for configuring pages_event_alert
	ProjectID []string `json:"project_id"`
	// Used for configuring advanced_ddos_attack_l4_alert
	Protocol []string `json:"protocol"`
	// Usage depends on specific alert type
	QueryTag []string `json:"query_tag"`
	// Used for configuring advanced_ddos_attack_l7_alert
	RequestsPerSecond []string `json:"requests_per_second"`
	// Usage depends on specific alert type
	Selectors []string `json:"selectors"`
	// Used for configuring clickhouse_alert_fw_ent_anomaly
	Services []string `json:"services"`
	// Usage depends on specific alert type
	Slo []string `json:"slo"`
	// Used for configuring health_check_status_notification
	Status []string `json:"status"`
	// Used for configuring advanced_ddos_attack_l7_alert
	TargetHostname []string `json:"target_hostname"`
	// Used for configuring advanced_ddos_attack_l4_alert
	TargetIP []string `json:"target_ip"`
	// Used for configuring advanced_ddos_attack_l7_alert
	TargetZoneName []string `json:"target_zone_name"`
	// Used for configuring traffic_anomalies_alert
	TrafficExclusions []AlertFiltersTrafficExclusion `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID []string `json:"tunnel_id"`
	// Usage depends on specific alert type
	TunnelName []string `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where []string `json:"where"`
	// Usage depends on specific alert type
	Zones []string         `json:"zones"`
	JSON  alertFiltersJSON `json:"-"`
}

// alertFiltersJSON contains the JSON metadata for the struct [AlertFilters]
type alertFiltersJSON struct {
	Actions                      apijson.Field
	AffectedAsns                 apijson.Field
	AffectedComponents           apijson.Field
	AffectedLocations            apijson.Field
	AirportCode                  apijson.Field
	AlertTriggerPreferences      apijson.Field
	AlertTriggerPreferencesValue apijson.Field
	Enabled                      apijson.Field
	Environment                  apijson.Field
	Event                        apijson.Field
	EventSource                  apijson.Field
	EventType                    apijson.Field
	GroupBy                      apijson.Field
	HealthCheckID                apijson.Field
	IncidentImpact               apijson.Field
	InputID                      apijson.Field
	InsightClass                 apijson.Field
	Limit                        apijson.Field
	LogoTag                      apijson.Field
	MegabitsPerSecond            apijson.Field
	NewHealth                    apijson.Field
	NewStatus                    apijson.Field
	PacketsPerSecond             apijson.Field
	PoolID                       apijson.Field
	PopNames                     apijson.Field
	Product                      apijson.Field
	ProjectID                    apijson.Field
	Protocol                     apijson.Field
	QueryTag                     apijson.Field
	RequestsPerSecond            apijson.Field
	Selectors                    apijson.Field
	Services                     apijson.Field
	Slo                          apijson.Field
	Status                       apijson.Field
	TargetHostname               apijson.Field
	TargetIP                     apijson.Field
	TargetZoneName               apijson.Field
	TrafficExclusions            apijson.Field
	TunnelID                     apijson.Field
	TunnelName                   apijson.Field
	Where                        apijson.Field
	Zones                        apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AlertFilters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertFiltersJSON) RawJSON() string {
	return r.raw
}

type AlertFiltersIncidentImpact string

const (
	AlertFiltersIncidentImpactIncidentImpactNone     AlertFiltersIncidentImpact = "INCIDENT_IMPACT_NONE"
	AlertFiltersIncidentImpactIncidentImpactMinor    AlertFiltersIncidentImpact = "INCIDENT_IMPACT_MINOR"
	AlertFiltersIncidentImpactIncidentImpactMajor    AlertFiltersIncidentImpact = "INCIDENT_IMPACT_MAJOR"
	AlertFiltersIncidentImpactIncidentImpactCritical AlertFiltersIncidentImpact = "INCIDENT_IMPACT_CRITICAL"
)

func (r AlertFiltersIncidentImpact) IsKnown() bool {
	switch r {
	case AlertFiltersIncidentImpactIncidentImpactNone, AlertFiltersIncidentImpactIncidentImpactMinor, AlertFiltersIncidentImpactIncidentImpactMajor, AlertFiltersIncidentImpactIncidentImpactCritical:
		return true
	}
	return false
}

type AlertFiltersTrafficExclusion string

const (
	AlertFiltersTrafficExclusionSecurityEvents AlertFiltersTrafficExclusion = "security_events"
)

func (r AlertFiltersTrafficExclusion) IsKnown() bool {
	switch r {
	case AlertFiltersTrafficExclusionSecurityEvents:
		return true
	}
	return false
}

// Optional filters that allow you to be alerted only on a subset of events for
// that alert type based on some criteria. This is only available for select alert
// types. See alert type documentation for more details.
type AlertFiltersParam struct {
	// Usage depends on specific alert type
	Actions param.Field[[]string] `json:"actions"`
	// Used for configuring radar_notification
	AffectedAsns param.Field[[]string] `json:"affected_asns"`
	// Used for configuring incident_alert
	AffectedComponents param.Field[[]string] `json:"affected_components"`
	// Used for configuring radar_notification
	AffectedLocations param.Field[[]string] `json:"affected_locations"`
	// Used for configuring maintenance_event_notification
	AirportCode param.Field[[]string] `json:"airport_code"`
	// Usage depends on specific alert type
	AlertTriggerPreferences param.Field[[]string] `json:"alert_trigger_preferences"`
	// Usage depends on specific alert type
	AlertTriggerPreferencesValue param.Field[[]string] `json:"alert_trigger_preferences_value"`
	// Used for configuring load_balancing_pool_enablement_alert
	Enabled param.Field[[]string] `json:"enabled"`
	// Used for configuring pages_event_alert
	Environment param.Field[[]string] `json:"environment"`
	// Used for configuring pages_event_alert
	Event param.Field[[]string] `json:"event"`
	// Used for configuring load_balancing_health_alert
	EventSource param.Field[[]string] `json:"event_source"`
	// Usage depends on specific alert type
	EventType param.Field[[]string] `json:"event_type"`
	// Usage depends on specific alert type
	GroupBy param.Field[[]string] `json:"group_by"`
	// Used for configuring health_check_status_notification
	HealthCheckID param.Field[[]string] `json:"health_check_id"`
	// Used for configuring incident_alert
	IncidentImpact param.Field[[]AlertFiltersIncidentImpact] `json:"incident_impact"`
	// Used for configuring stream_live_notifications
	InputID param.Field[[]string] `json:"input_id"`
	// Used for configuring security_insights_alert
	InsightClass param.Field[[]string] `json:"insight_class"`
	// Used for configuring billing_usage_alert
	Limit param.Field[[]string] `json:"limit"`
	// Used for configuring logo_match_alert
	LogoTag param.Field[[]string] `json:"logo_tag"`
	// Used for configuring advanced_ddos_attack_l4_alert
	MegabitsPerSecond param.Field[[]string] `json:"megabits_per_second"`
	// Used for configuring load_balancing_health_alert
	NewHealth param.Field[[]string] `json:"new_health"`
	// Used for configuring tunnel_health_event
	NewStatus param.Field[[]string] `json:"new_status"`
	// Used for configuring advanced_ddos_attack_l4_alert
	PacketsPerSecond param.Field[[]string] `json:"packets_per_second"`
	// Usage depends on specific alert type
	PoolID param.Field[[]string] `json:"pool_id"`
	// Usage depends on specific alert type
	PopNames param.Field[[]string] `json:"pop_names"`
	// Used for configuring billing_usage_alert
	Product param.Field[[]string] `json:"product"`
	// Used for configuring pages_event_alert
	ProjectID param.Field[[]string] `json:"project_id"`
	// Used for configuring advanced_ddos_attack_l4_alert
	Protocol param.Field[[]string] `json:"protocol"`
	// Usage depends on specific alert type
	QueryTag param.Field[[]string] `json:"query_tag"`
	// Used for configuring advanced_ddos_attack_l7_alert
	RequestsPerSecond param.Field[[]string] `json:"requests_per_second"`
	// Usage depends on specific alert type
	Selectors param.Field[[]string] `json:"selectors"`
	// Used for configuring clickhouse_alert_fw_ent_anomaly
	Services param.Field[[]string] `json:"services"`
	// Usage depends on specific alert type
	Slo param.Field[[]string] `json:"slo"`
	// Used for configuring health_check_status_notification
	Status param.Field[[]string] `json:"status"`
	// Used for configuring advanced_ddos_attack_l7_alert
	TargetHostname param.Field[[]string] `json:"target_hostname"`
	// Used for configuring advanced_ddos_attack_l4_alert
	TargetIP param.Field[[]string] `json:"target_ip"`
	// Used for configuring advanced_ddos_attack_l7_alert
	TargetZoneName param.Field[[]string] `json:"target_zone_name"`
	// Used for configuring traffic_anomalies_alert
	TrafficExclusions param.Field[[]AlertFiltersTrafficExclusion] `json:"traffic_exclusions"`
	// Used for configuring tunnel_health_event
	TunnelID param.Field[[]string] `json:"tunnel_id"`
	// Usage depends on specific alert type
	TunnelName param.Field[[]string] `json:"tunnel_name"`
	// Usage depends on specific alert type
	Where param.Field[[]string] `json:"where"`
	// Usage depends on specific alert type
	Zones param.Field[[]string] `json:"zones"`
}

func (r AlertFiltersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Refers to which event will trigger a Notification dispatch. You can use the
// endpoint to get available alert types which then will give you a list of
// possible values.
type AlertType string

const (
	AlertTypeAccessCustomCertificateExpirationType         AlertType = "access_custom_certificate_expiration_type"
	AlertTypeAdvancedDdosAttackL4Alert                     AlertType = "advanced_ddos_attack_l4_alert"
	AlertTypeAdvancedDdosAttackL7Alert                     AlertType = "advanced_ddos_attack_l7_alert"
	AlertTypeAdvancedHTTPAlertError                        AlertType = "advanced_http_alert_error"
	AlertTypeBgpHijackNotification                         AlertType = "bgp_hijack_notification"
	AlertTypeBillingUsageAlert                             AlertType = "billing_usage_alert"
	AlertTypeBlockNotificationBlockRemoved                 AlertType = "block_notification_block_removed"
	AlertTypeBlockNotificationNewBlock                     AlertType = "block_notification_new_block"
	AlertTypeBlockNotificationReviewRejected               AlertType = "block_notification_review_rejected"
	AlertTypeBotTrafficBasicAlert                          AlertType = "bot_traffic_basic_alert"
	AlertTypeBrandProtectionAlert                          AlertType = "brand_protection_alert"
	AlertTypeBrandProtectionDigest                         AlertType = "brand_protection_digest"
	AlertTypeClickhouseAlertFwAnomaly                      AlertType = "clickhouse_alert_fw_anomaly"
	AlertTypeClickhouseAlertFwEntAnomaly                   AlertType = "clickhouse_alert_fw_ent_anomaly"
	AlertTypeCloudforceOneRequestNotification              AlertType = "cloudforce_one_request_notification"
	AlertTypeCustomAnalytics                               AlertType = "custom_analytics"
	AlertTypeCustomBotDetectionAlert                       AlertType = "custom_bot_detection_alert"
	AlertTypeCustomSslCertificateEventType                 AlertType = "custom_ssl_certificate_event_type"
	AlertTypeDedicatedSslCertificateEventType              AlertType = "dedicated_ssl_certificate_event_type"
	AlertTypeDeviceConnectivityAnomalyAlert                AlertType = "device_connectivity_anomaly_alert"
	AlertTypeDosAttackL4                                   AlertType = "dos_attack_l4"
	AlertTypeDosAttackL7                                   AlertType = "dos_attack_l7"
	AlertTypeExpiringServiceTokenAlert                     AlertType = "expiring_service_token_alert"
	AlertTypeFailingLogpushJobDisabledAlert                AlertType = "failing_logpush_job_disabled_alert"
	AlertTypeFbmAutoAdvertisement                          AlertType = "fbm_auto_advertisement"
	AlertTypeFbmDosdAttack                                 AlertType = "fbm_dosd_attack"
	AlertTypeFbmVolumetricAttack                           AlertType = "fbm_volumetric_attack"
	AlertTypeHealthCheckStatusNotification                 AlertType = "health_check_status_notification"
	AlertTypeHostnameAopCustomCertificateExpirationType    AlertType = "hostname_aop_custom_certificate_expiration_type"
	AlertTypeHTTPAlertEdgeError                            AlertType = "http_alert_edge_error"
	AlertTypeHTTPAlertOriginError                          AlertType = "http_alert_origin_error"
	AlertTypeImageNotification                             AlertType = "image_notification"
	AlertTypeImageResizingNotification                     AlertType = "image_resizing_notification"
	AlertTypeIncidentAlert                                 AlertType = "incident_alert"
	AlertTypeLoadBalancingHealthAlert                      AlertType = "load_balancing_health_alert"
	AlertTypeLoadBalancingPoolEnablementAlert              AlertType = "load_balancing_pool_enablement_alert"
	AlertTypeLogoMatchAlert                                AlertType = "logo_match_alert"
	AlertTypeMagicTunnelHealthCheckEvent                   AlertType = "magic_tunnel_health_check_event"
	AlertTypeMagicWanTunnelHealth                          AlertType = "magic_wan_tunnel_health"
	AlertTypeMaintenanceEventNotification                  AlertType = "maintenance_event_notification"
	AlertTypeMtlsCertificateStoreCertificateExpirationType AlertType = "mtls_certificate_store_certificate_expiration_type"
	AlertTypePagesEventAlert                               AlertType = "pages_event_alert"
	AlertTypeRadarNotification                             AlertType = "radar_notification"
	AlertTypeRealOriginMonitoring                          AlertType = "real_origin_monitoring"
	AlertTypeScriptmonitorAlertNewCodeChangeDetections     AlertType = "scriptmonitor_alert_new_code_change_detections"
	AlertTypeScriptmonitorAlertNewHosts                    AlertType = "scriptmonitor_alert_new_hosts"
	AlertTypeScriptmonitorAlertNewMaliciousHosts           AlertType = "scriptmonitor_alert_new_malicious_hosts"
	AlertTypeScriptmonitorAlertNewMaliciousScripts         AlertType = "scriptmonitor_alert_new_malicious_scripts"
	AlertTypeScriptmonitorAlertNewMaliciousURL             AlertType = "scriptmonitor_alert_new_malicious_url"
	AlertTypeScriptmonitorAlertNewMaxLengthResourceURL     AlertType = "scriptmonitor_alert_new_max_length_resource_url"
	AlertTypeScriptmonitorAlertNewResources                AlertType = "scriptmonitor_alert_new_resources"
	AlertTypeSecondaryDNSAllPrimariesFailing               AlertType = "secondary_dns_all_primaries_failing"
	AlertTypeSecondaryDNSPrimariesFailing                  AlertType = "secondary_dns_primaries_failing"
	AlertTypeSecondaryDNSWarning                           AlertType = "secondary_dns_warning"
	AlertTypeSecondaryDNSZoneSuccessfullyUpdated           AlertType = "secondary_dns_zone_successfully_updated"
	AlertTypeSecondaryDNSZoneValidationWarning             AlertType = "secondary_dns_zone_validation_warning"
	AlertTypeSecurityInsightsAlert                         AlertType = "security_insights_alert"
	AlertTypeSentinelAlert                                 AlertType = "sentinel_alert"
	AlertTypeStreamLiveNotifications                       AlertType = "stream_live_notifications"
	AlertTypeSyntheticTestLatencyAlert                     AlertType = "synthetic_test_latency_alert"
	AlertTypeSyntheticTestLowAvailabilityAlert             AlertType = "synthetic_test_low_availability_alert"
	AlertTypeTrafficAnomaliesAlert                         AlertType = "traffic_anomalies_alert"
	AlertTypeTunnelHealthEvent                             AlertType = "tunnel_health_event"
	AlertTypeTunnelUpdateEvent                             AlertType = "tunnel_update_event"
	AlertTypeUniversalSslEventType                         AlertType = "universal_ssl_event_type"
	AlertTypeWebAnalyticsMetricsUpdate                     AlertType = "web_analytics_metrics_update"
	AlertTypeZoneAopCustomCertificateExpirationType        AlertType = "zone_aop_custom_certificate_expiration_type"
)

func (r AlertType) IsKnown() bool {
	switch r {
	case AlertTypeAccessCustomCertificateExpirationType, AlertTypeAdvancedDdosAttackL4Alert, AlertTypeAdvancedDdosAttackL7Alert, AlertTypeAdvancedHTTPAlertError, AlertTypeBgpHijackNotification, AlertTypeBillingUsageAlert, AlertTypeBlockNotificationBlockRemoved, AlertTypeBlockNotificationNewBlock, AlertTypeBlockNotificationReviewRejected, AlertTypeBotTrafficBasicAlert, AlertTypeBrandProtectionAlert, AlertTypeBrandProtectionDigest, AlertTypeClickhouseAlertFwAnomaly, AlertTypeClickhouseAlertFwEntAnomaly, AlertTypeCloudforceOneRequestNotification, AlertTypeCustomAnalytics, AlertTypeCustomBotDetectionAlert, AlertTypeCustomSslCertificateEventType, AlertTypeDedicatedSslCertificateEventType, AlertTypeDeviceConnectivityAnomalyAlert, AlertTypeDosAttackL4, AlertTypeDosAttackL7, AlertTypeExpiringServiceTokenAlert, AlertTypeFailingLogpushJobDisabledAlert, AlertTypeFbmAutoAdvertisement, AlertTypeFbmDosdAttack, AlertTypeFbmVolumetricAttack, AlertTypeHealthCheckStatusNotification, AlertTypeHostnameAopCustomCertificateExpirationType, AlertTypeHTTPAlertEdgeError, AlertTypeHTTPAlertOriginError, AlertTypeImageNotification, AlertTypeImageResizingNotification, AlertTypeIncidentAlert, AlertTypeLoadBalancingHealthAlert, AlertTypeLoadBalancingPoolEnablementAlert, AlertTypeLogoMatchAlert, AlertTypeMagicTunnelHealthCheckEvent, AlertTypeMagicWanTunnelHealth, AlertTypeMaintenanceEventNotification, AlertTypeMtlsCertificateStoreCertificateExpirationType, AlertTypePagesEventAlert, AlertTypeRadarNotification, AlertTypeRealOriginMonitoring, AlertTypeScriptmonitorAlertNewCodeChangeDetections, AlertTypeScriptmonitorAlertNewHosts, AlertTypeScriptmonitorAlertNewMaliciousHosts, AlertTypeScriptmonitorAlertNewMaliciousScripts, AlertTypeScriptmonitorAlertNewMaliciousURL, AlertTypeScriptmonitorAlertNewMaxLengthResourceURL, AlertTypeScriptmonitorAlertNewResources, AlertTypeSecondaryDNSAllPrimariesFailing, AlertTypeSecondaryDNSPrimariesFailing, AlertTypeSecondaryDNSWarning, AlertTypeSecondaryDNSZoneSuccessfullyUpdated, AlertTypeSecondaryDNSZoneValidationWarning, AlertTypeSecurityInsightsAlert, AlertTypeSentinelAlert, AlertTypeStreamLiveNotifications, AlertTypeSyntheticTestLatencyAlert, AlertTypeSyntheticTestLowAvailabilityAlert, AlertTypeTrafficAnomaliesAlert, AlertTypeTunnelHealthEvent, AlertTypeTunnelUpdateEvent, AlertTypeUniversalSslEventType, AlertTypeWebAnalyticsMetricsUpdate, AlertTypeZoneAopCustomCertificateExpirationType:
		return true
	}
	return false
}

// List of IDs that will be used when dispatching a notification. IDs for email
// type will be the email address.
type Mechanisms struct {
	Email     []MechanismsEmail     `json:"email"`
	Pagerduty []MechanismsPagerduty `json:"pagerduty"`
	Webhooks  []MechanismsWebhook   `json:"webhooks"`
	JSON      mechanismsJSON        `json:"-"`
}

// mechanismsJSON contains the JSON metadata for the struct [Mechanisms]
type mechanismsJSON struct {
	Email       apijson.Field
	Pagerduty   apijson.Field
	Webhooks    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Mechanisms) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mechanismsJSON) RawJSON() string {
	return r.raw
}

type MechanismsEmail struct {
	// The email address
	ID   string              `json:"id"`
	JSON mechanismsEmailJSON `json:"-"`
}

// mechanismsEmailJSON contains the JSON metadata for the struct [MechanismsEmail]
type mechanismsEmailJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MechanismsEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mechanismsEmailJSON) RawJSON() string {
	return r.raw
}

type MechanismsPagerduty struct {
	// UUID
	ID   string                  `json:"id"`
	JSON mechanismsPagerdutyJSON `json:"-"`
}

// mechanismsPagerdutyJSON contains the JSON metadata for the struct
// [MechanismsPagerduty]
type mechanismsPagerdutyJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MechanismsPagerduty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mechanismsPagerdutyJSON) RawJSON() string {
	return r.raw
}

type MechanismsWebhook struct {
	// UUID
	ID   string                `json:"id"`
	JSON mechanismsWebhookJSON `json:"-"`
}

// mechanismsWebhookJSON contains the JSON metadata for the struct
// [MechanismsWebhook]
type mechanismsWebhookJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MechanismsWebhook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mechanismsWebhookJSON) RawJSON() string {
	return r.raw
}

// List of IDs that will be used when dispatching a notification. IDs for email
// type will be the email address.
type MechanismsParam struct {
	Email     param.Field[[]MechanismsEmailParam]     `json:"email"`
	Pagerduty param.Field[[]MechanismsPagerdutyParam] `json:"pagerduty"`
	Webhooks  param.Field[[]MechanismsWebhookParam]   `json:"webhooks"`
}

func (r MechanismsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MechanismsEmailParam struct {
	// The email address
	ID param.Field[string] `json:"id"`
}

func (r MechanismsEmailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MechanismsPagerdutyParam struct {
	// UUID
	ID param.Field[string] `json:"id"`
}

func (r MechanismsPagerdutyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MechanismsWebhookParam struct {
	// UUID
	ID param.Field[string] `json:"id"`
}

func (r MechanismsWebhookParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Policies struct {
	// The unique identifier of a notification policy
	ID string `json:"id"`
	// Optional specification of how often to re-alert from the same incident, not
	// support on all alert types.
	AlertInterval string `json:"alert_interval"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType AlertType `json:"alert_type"`
	Created   time.Time `json:"created" format:"date-time"`
	// Optional description for the Notification policy.
	Description string `json:"description"`
	// Whether or not the Notification policy is enabled.
	Enabled bool `json:"enabled"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters AlertFilters `json:"filters"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms Mechanisms `json:"mechanisms"`
	Modified   time.Time  `json:"modified" format:"date-time"`
	// Name of the policy.
	Name string       `json:"name"`
	JSON policiesJSON `json:"-"`
}

// policiesJSON contains the JSON metadata for the struct [Policies]
type policiesJSON struct {
	ID            apijson.Field
	AlertInterval apijson.Field
	AlertType     apijson.Field
	Created       apijson.Field
	Description   apijson.Field
	Enabled       apijson.Field
	Filters       apijson.Field
	Mechanisms    apijson.Field
	Modified      apijson.Field
	Name          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Policies) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policiesJSON) RawJSON() string {
	return r.raw
}

type AccountAlertingV3PolicyGetResponse struct {
	Errors   []AaaMessage `json:"errors,required"`
	Messages []AaaMessage `json:"messages,required"`
	// Whether the API call was successful
	Success AccountAlertingV3PolicyGetResponseSuccess `json:"success,required"`
	Result  Policies                                  `json:"result"`
	JSON    accountAlertingV3PolicyGetResponseJSON    `json:"-"`
}

// accountAlertingV3PolicyGetResponseJSON contains the JSON metadata for the struct
// [AccountAlertingV3PolicyGetResponse]
type accountAlertingV3PolicyGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3PolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3PolicyGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountAlertingV3PolicyGetResponseSuccess bool

const (
	AccountAlertingV3PolicyGetResponseSuccessTrue AccountAlertingV3PolicyGetResponseSuccess = true
)

func (r AccountAlertingV3PolicyGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountAlertingV3PolicyGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountAlertingV3PolicyListResponse struct {
	Errors   []AaaMessage `json:"errors,required"`
	Messages []AaaMessage `json:"messages,required"`
	// Whether the API call was successful
	Success    AccountAlertingV3PolicyListResponseSuccess    `json:"success,required"`
	Result     []Policies                                    `json:"result"`
	ResultInfo AccountAlertingV3PolicyListResponseResultInfo `json:"result_info"`
	JSON       accountAlertingV3PolicyListResponseJSON       `json:"-"`
}

// accountAlertingV3PolicyListResponseJSON contains the JSON metadata for the
// struct [AccountAlertingV3PolicyListResponse]
type accountAlertingV3PolicyListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3PolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3PolicyListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountAlertingV3PolicyListResponseSuccess bool

const (
	AccountAlertingV3PolicyListResponseSuccessTrue AccountAlertingV3PolicyListResponseSuccess = true
)

func (r AccountAlertingV3PolicyListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountAlertingV3PolicyListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountAlertingV3PolicyListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       accountAlertingV3PolicyListResponseResultInfoJSON `json:"-"`
}

// accountAlertingV3PolicyListResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountAlertingV3PolicyListResponseResultInfo]
type accountAlertingV3PolicyListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3PolicyListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3PolicyListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountAlertingV3PolicyNewParams struct {
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType param.Field[AlertType] `json:"alert_type,required"`
	// Whether or not the Notification policy is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms param.Field[MechanismsParam] `json:"mechanisms,required"`
	// Name of the policy.
	Name param.Field[string] `json:"name,required"`
	// Optional specification of how often to re-alert from the same incident, not
	// support on all alert types.
	AlertInterval param.Field[string] `json:"alert_interval"`
	// Optional description for the Notification policy.
	Description param.Field[string] `json:"description"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters param.Field[AlertFiltersParam] `json:"filters"`
}

func (r AccountAlertingV3PolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAlertingV3PolicyUpdateParams struct {
	// Optional specification of how often to re-alert from the same incident, not
	// support on all alert types.
	AlertInterval param.Field[string] `json:"alert_interval"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType param.Field[AlertType] `json:"alert_type"`
	// Optional description for the Notification policy.
	Description param.Field[string] `json:"description"`
	// Whether or not the Notification policy is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Optional filters that allow you to be alerted only on a subset of events for
	// that alert type based on some criteria. This is only available for select alert
	// types. See alert type documentation for more details.
	Filters param.Field[AlertFiltersParam] `json:"filters"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms param.Field[MechanismsParam] `json:"mechanisms"`
	// Name of the policy.
	Name param.Field[string] `json:"name"`
}

func (r AccountAlertingV3PolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
