// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"reflect"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/tidwall/gjson"
)

// AccountService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options            []option.RequestOption
	Access             *AccountAccessService
	Addressing         *AccountAddressingService
	AIGateway          *AccountAIGatewayService
	AI                 *AccountAIService
	Alerting           *AccountAlertingService
	Billing            *AccountBillingService
	BotnetFeed         *AccountBotnetFeedService
	BrandProtection    *AccountBrandProtectionService
	BrowserRendering   *AccountBrowserRenderingService
	Calls              *AccountCallService
	CfdTunnel          *AccountCfdTunnelService
	Challenges         *AccountChallengeService
	CloudforceOne      *AccountCloudforceOneService
	Cni                *AccountCniService
	CustomNs           *AccountCustomNService
	D1                 *AccountD1Service
	Devices            *AccountDeviceService
	Dex                *AccountDexService
	Diagnostics        *AccountDiagnosticService
	Dlp                *AccountDlpService
	DNSFirewall        *AccountDNSFirewallService
	DNSSettings        *AccountDNSSettingService
	EmailSecurity      *AccountEmailSecurityService
	Email              *AccountEmailService
	EventNotifications *AccountEventNotificationService
	Firewall           *AccountFirewallService
	Gateway            *AccountGatewayService
	Hyperdrive         *AccountHyperdriveService
	Iam                *AccountIamService
	Images             *AccountImageService
	Infrastructure     *AccountInfrastructureService
	Intel              *AccountIntelService
	LoadBalancers      *AccountLoadBalancerService
	Logpush            *AccountLogpushService
	Logs               *AccountLogService
	Magic              *AccountMagicService
	Members            *AccountMemberService
	Mnm                *AccountMnmService
	MtlsCertificates   *AccountMtlsCertificateService
	Pages              *AccountPageService
	Pcaps              *AccountPcapService
	Queues             *AccountQueueService
	R2                 *AccountR2Service
	Registrar          *AccountRegistrarService
	RequestTracer      *AccountRequestTracerService
	Roles              *AccountRoleService
	Rules              *AccountRuleService
	Rulesets           *AccountRulesetService
	Rum                *AccountRumService
	SecondaryDNS       *AccountSecondaryDNSService
	SecurityCenter     *AccountSecurityCenterService
	Shares             *AccountShareService
	Slurper            *AccountSlurperService
	Storage            *AccountStorageService
	Stream             *AccountStreamService
	Subscriptions      *AccountSubscriptionService
	Teamnet            *AccountTeamnetService
	Tokens             *AccountTokenService
	Urlscanner         *AccountUrlscannerService
	Vectorize          *AccountVectorizeService
	WarpConnector      *AccountWarpConnectorService
	Workers            *AccountWorkerService
	Workflows          *AccountWorkflowService
	Zerotrust          *AccountZerotrustService
	ZtRiskScoring      *AccountZtRiskScoringService
	CustomPages        *AccountCustomPageService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	r.Access = NewAccountAccessService(opts...)
	r.Addressing = NewAccountAddressingService(opts...)
	r.AIGateway = NewAccountAIGatewayService(opts...)
	r.AI = NewAccountAIService(opts...)
	r.Alerting = NewAccountAlertingService(opts...)
	r.Billing = NewAccountBillingService(opts...)
	r.BotnetFeed = NewAccountBotnetFeedService(opts...)
	r.BrandProtection = NewAccountBrandProtectionService(opts...)
	r.BrowserRendering = NewAccountBrowserRenderingService(opts...)
	r.Calls = NewAccountCallService(opts...)
	r.CfdTunnel = NewAccountCfdTunnelService(opts...)
	r.Challenges = NewAccountChallengeService(opts...)
	r.CloudforceOne = NewAccountCloudforceOneService(opts...)
	r.Cni = NewAccountCniService(opts...)
	r.CustomNs = NewAccountCustomNService(opts...)
	r.D1 = NewAccountD1Service(opts...)
	r.Devices = NewAccountDeviceService(opts...)
	r.Dex = NewAccountDexService(opts...)
	r.Diagnostics = NewAccountDiagnosticService(opts...)
	r.Dlp = NewAccountDlpService(opts...)
	r.DNSFirewall = NewAccountDNSFirewallService(opts...)
	r.DNSSettings = NewAccountDNSSettingService(opts...)
	r.EmailSecurity = NewAccountEmailSecurityService(opts...)
	r.Email = NewAccountEmailService(opts...)
	r.EventNotifications = NewAccountEventNotificationService(opts...)
	r.Firewall = NewAccountFirewallService(opts...)
	r.Gateway = NewAccountGatewayService(opts...)
	r.Hyperdrive = NewAccountHyperdriveService(opts...)
	r.Iam = NewAccountIamService(opts...)
	r.Images = NewAccountImageService(opts...)
	r.Infrastructure = NewAccountInfrastructureService(opts...)
	r.Intel = NewAccountIntelService(opts...)
	r.LoadBalancers = NewAccountLoadBalancerService(opts...)
	r.Logpush = NewAccountLogpushService(opts...)
	r.Logs = NewAccountLogService(opts...)
	r.Magic = NewAccountMagicService(opts...)
	r.Members = NewAccountMemberService(opts...)
	r.Mnm = NewAccountMnmService(opts...)
	r.MtlsCertificates = NewAccountMtlsCertificateService(opts...)
	r.Pages = NewAccountPageService(opts...)
	r.Pcaps = NewAccountPcapService(opts...)
	r.Queues = NewAccountQueueService(opts...)
	r.R2 = NewAccountR2Service(opts...)
	r.Registrar = NewAccountRegistrarService(opts...)
	r.RequestTracer = NewAccountRequestTracerService(opts...)
	r.Roles = NewAccountRoleService(opts...)
	r.Rules = NewAccountRuleService(opts...)
	r.Rulesets = NewAccountRulesetService(opts...)
	r.Rum = NewAccountRumService(opts...)
	r.SecondaryDNS = NewAccountSecondaryDNSService(opts...)
	r.SecurityCenter = NewAccountSecurityCenterService(opts...)
	r.Shares = NewAccountShareService(opts...)
	r.Slurper = NewAccountSlurperService(opts...)
	r.Storage = NewAccountStorageService(opts...)
	r.Stream = NewAccountStreamService(opts...)
	r.Subscriptions = NewAccountSubscriptionService(opts...)
	r.Teamnet = NewAccountTeamnetService(opts...)
	r.Tokens = NewAccountTokenService(opts...)
	r.Urlscanner = NewAccountUrlscannerService(opts...)
	r.Vectorize = NewAccountVectorizeService(opts...)
	r.WarpConnector = NewAccountWarpConnectorService(opts...)
	r.Workers = NewAccountWorkerService(opts...)
	r.Workflows = NewAccountWorkflowService(opts...)
	r.Zerotrust = NewAccountZerotrustService(opts...)
	r.ZtRiskScoring = NewAccountZtRiskScoringService(opts...)
	r.CustomPages = NewAccountCustomPageService(opts...)
	return
}

type AaaAuditLogs struct {
	// This field can have the runtime type of [[]AaaMessage].
	Errors interface{} `json:"errors"`
	// This field can have the runtime type of [[]AaaMessage].
	Messages interface{} `json:"messages"`
	// This field can have the runtime type of [[]AaaAuditLogsObjectResult].
	Result  interface{}      `json:"result"`
	Success bool             `json:"success"`
	JSON    aaaAuditLogsJSON `json:"-"`
	union   AaaAuditLogsUnion
}

// aaaAuditLogsJSON contains the JSON metadata for the struct [AaaAuditLogs]
type aaaAuditLogsJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r aaaAuditLogsJSON) RawJSON() string {
	return r.raw
}

func (r *AaaAuditLogs) UnmarshalJSON(data []byte) (err error) {
	*r = AaaAuditLogs{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AaaAuditLogsUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [AaaAuditLogsObject],
// [APIResponseAlerting].
func (r AaaAuditLogs) AsUnion() AaaAuditLogsUnion {
	return r.union
}

// Union satisfied by [AaaAuditLogsObject] or [APIResponseAlerting].
type AaaAuditLogsUnion interface {
	implementsAaaAuditLogs()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AaaAuditLogsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AaaAuditLogsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIResponseAlerting{}),
		},
	)
}

type AaaAuditLogsObject struct {
	Errors   []AaaMessage               `json:"errors"`
	Messages []AaaMessage               `json:"messages"`
	Result   []AaaAuditLogsObjectResult `json:"result"`
	Success  bool                       `json:"success"`
	JSON     aaaAuditLogsObjectJSON     `json:"-"`
}

// aaaAuditLogsObjectJSON contains the JSON metadata for the struct
// [AaaAuditLogsObject]
type aaaAuditLogsObjectJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AaaAuditLogsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aaaAuditLogsObjectJSON) RawJSON() string {
	return r.raw
}

func (r AaaAuditLogsObject) implementsAaaAuditLogs() {}

type AaaAuditLogsObjectResult struct {
	// A string that uniquely identifies the audit log.
	ID     string                         `json:"id"`
	Action AaaAuditLogsObjectResultAction `json:"action"`
	Actor  AaaAuditLogsObjectResultActor  `json:"actor"`
	// The source of the event.
	Interface string `json:"interface"`
	// An object which can lend more context to the action being logged. This is a
	// flexible value and varies between different actions.
	Metadata interface{} `json:"metadata"`
	// The new value of the resource that was modified.
	NewValue string `json:"newValue"`
	// The value of the resource before it was modified.
	OldValue string                           `json:"oldValue"`
	Owner    AaaAuditLogsObjectResultOwner    `json:"owner"`
	Resource AaaAuditLogsObjectResultResource `json:"resource"`
	// A UTC RFC3339 timestamp that specifies when the action being logged occured.
	When time.Time                    `json:"when" format:"date-time"`
	JSON aaaAuditLogsObjectResultJSON `json:"-"`
}

// aaaAuditLogsObjectResultJSON contains the JSON metadata for the struct
// [AaaAuditLogsObjectResult]
type aaaAuditLogsObjectResultJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Actor       apijson.Field
	Interface   apijson.Field
	Metadata    apijson.Field
	NewValue    apijson.Field
	OldValue    apijson.Field
	Owner       apijson.Field
	Resource    apijson.Field
	When        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AaaAuditLogsObjectResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aaaAuditLogsObjectResultJSON) RawJSON() string {
	return r.raw
}

type AaaAuditLogsObjectResultAction struct {
	// A boolean that indicates if the action attempted was successful.
	Result bool `json:"result"`
	// A short string that describes the action that was performed.
	Type string                             `json:"type"`
	JSON aaaAuditLogsObjectResultActionJSON `json:"-"`
}

// aaaAuditLogsObjectResultActionJSON contains the JSON metadata for the struct
// [AaaAuditLogsObjectResultAction]
type aaaAuditLogsObjectResultActionJSON struct {
	Result      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AaaAuditLogsObjectResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aaaAuditLogsObjectResultActionJSON) RawJSON() string {
	return r.raw
}

type AaaAuditLogsObjectResultActor struct {
	// The ID of the actor that performed the action. If a user performed the action,
	// this will be their User ID.
	ID string `json:"id"`
	// The email of the user that performed the action.
	Email string `json:"email" format:"email"`
	// The IP address of the request that performed the action.
	IP string `json:"ip"`
	// The type of actor, whether a User, Cloudflare Admin, or an Automated System.
	Type AaaAuditLogsObjectResultActorType `json:"type"`
	JSON aaaAuditLogsObjectResultActorJSON `json:"-"`
}

// aaaAuditLogsObjectResultActorJSON contains the JSON metadata for the struct
// [AaaAuditLogsObjectResultActor]
type aaaAuditLogsObjectResultActorJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	IP          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AaaAuditLogsObjectResultActor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aaaAuditLogsObjectResultActorJSON) RawJSON() string {
	return r.raw
}

// The type of actor, whether a User, Cloudflare Admin, or an Automated System.
type AaaAuditLogsObjectResultActorType string

const (
	AaaAuditLogsObjectResultActorTypeUser       AaaAuditLogsObjectResultActorType = "user"
	AaaAuditLogsObjectResultActorTypeAdmin      AaaAuditLogsObjectResultActorType = "admin"
	AaaAuditLogsObjectResultActorTypeCloudflare AaaAuditLogsObjectResultActorType = "Cloudflare"
)

func (r AaaAuditLogsObjectResultActorType) IsKnown() bool {
	switch r {
	case AaaAuditLogsObjectResultActorTypeUser, AaaAuditLogsObjectResultActorTypeAdmin, AaaAuditLogsObjectResultActorTypeCloudflare:
		return true
	}
	return false
}

type AaaAuditLogsObjectResultOwner struct {
	// Identifier
	ID   string                            `json:"id"`
	JSON aaaAuditLogsObjectResultOwnerJSON `json:"-"`
}

// aaaAuditLogsObjectResultOwnerJSON contains the JSON metadata for the struct
// [AaaAuditLogsObjectResultOwner]
type aaaAuditLogsObjectResultOwnerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AaaAuditLogsObjectResultOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aaaAuditLogsObjectResultOwnerJSON) RawJSON() string {
	return r.raw
}

type AaaAuditLogsObjectResultResource struct {
	// An identifier for the resource that was affected by the action.
	ID string `json:"id"`
	// A short string that describes the resource that was affected by the action.
	Type string                               `json:"type"`
	JSON aaaAuditLogsObjectResultResourceJSON `json:"-"`
}

// aaaAuditLogsObjectResultResourceJSON contains the JSON metadata for the struct
// [AaaAuditLogsObjectResultResource]
type aaaAuditLogsObjectResultResourceJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AaaAuditLogsObjectResultResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aaaAuditLogsObjectResultResourceJSON) RawJSON() string {
	return r.raw
}

type AaaMessage struct {
	Code    int64          `json:"code,required"`
	Message string         `json:"message,required"`
	JSON    aaaMessageJSON `json:"-"`
}

// aaaMessageJSON contains the JSON metadata for the struct [AaaMessage]
type aaaMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AaaMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aaaMessageJSON) RawJSON() string {
	return r.raw
}

type IamAPIResponseCollection struct {
	ResultInfo IamAPIResponseCollectionResultInfo `json:"result_info"`
	JSON       iamAPIResponseCollectionJSON       `json:"-"`
	IamAPIResponseCommon
}

// iamAPIResponseCollectionJSON contains the JSON metadata for the struct
// [IamAPIResponseCollection]
type iamAPIResponseCollectionJSON struct {
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamAPIResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamAPIResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type IamAPIResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                `json:"total_count"`
	JSON       iamAPIResponseCollectionResultInfoJSON `json:"-"`
}

// iamAPIResponseCollectionResultInfoJSON contains the JSON metadata for the struct
// [IamAPIResponseCollectionResultInfo]
type iamAPIResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamAPIResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamAPIResponseCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type IamAPIResponseCommon struct {
	Errors   []IamMessage `json:"errors,required"`
	Messages []IamMessage `json:"messages,required"`
	// Whether the API call was successful
	Success IamAPIResponseCommonSuccess `json:"success,required"`
	JSON    iamAPIResponseCommonJSON    `json:"-"`
}

// iamAPIResponseCommonJSON contains the JSON metadata for the struct
// [IamAPIResponseCommon]
type iamAPIResponseCommonJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamAPIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamAPIResponseCommonJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IamAPIResponseCommonSuccess bool

const (
	IamAPIResponseCommonSuccessTrue IamAPIResponseCommonSuccess = true
)

func (r IamAPIResponseCommonSuccess) IsKnown() bool {
	switch r {
	case IamAPIResponseCommonSuccessTrue:
		return true
	}
	return false
}

type IamAPIResponseSingleID struct {
	Result IamAPIResponseSingleIDResult `json:"result,nullable"`
	JSON   iamAPIResponseSingleIDJSON   `json:"-"`
	IamAPIResponseCommon
}

// iamAPIResponseSingleIDJSON contains the JSON metadata for the struct
// [IamAPIResponseSingleID]
type iamAPIResponseSingleIDJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamAPIResponseSingleID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamAPIResponseSingleIDJSON) RawJSON() string {
	return r.raw
}

type IamAPIResponseSingleIDResult struct {
	// Identifier
	ID   string                           `json:"id,required"`
	JSON iamAPIResponseSingleIDResultJSON `json:"-"`
}

// iamAPIResponseSingleIDResultJSON contains the JSON metadata for the struct
// [IamAPIResponseSingleIDResult]
type iamAPIResponseSingleIDResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamAPIResponseSingleIDResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamAPIResponseSingleIDResultJSON) RawJSON() string {
	return r.raw
}

type IamMessage struct {
	Code    int64          `json:"code,required"`
	Message string         `json:"message,required"`
	JSON    iamMessageJSON `json:"-"`
}

// iamMessageJSON contains the JSON metadata for the struct [IamMessage]
type iamMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamMessageJSON) RawJSON() string {
	return r.raw
}

type WaitingroomAPIResponseCollection struct {
	Errors   []WaitingroomMessage `json:"errors,required"`
	Messages []WaitingroomMessage `json:"messages,required"`
	// Whether the API call was successful
	Success    WaitingroomAPIResponseCollectionSuccess    `json:"success,required"`
	ResultInfo WaitingroomAPIResponseCollectionResultInfo `json:"result_info"`
	JSON       waitingroomAPIResponseCollectionJSON       `json:"-"`
}

// waitingroomAPIResponseCollectionJSON contains the JSON metadata for the struct
// [WaitingroomAPIResponseCollection]
type waitingroomAPIResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingroomAPIResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingroomAPIResponseCollectionJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WaitingroomAPIResponseCollectionSuccess bool

const (
	WaitingroomAPIResponseCollectionSuccessTrue WaitingroomAPIResponseCollectionSuccess = true
)

func (r WaitingroomAPIResponseCollectionSuccess) IsKnown() bool {
	switch r {
	case WaitingroomAPIResponseCollectionSuccessTrue:
		return true
	}
	return false
}

type WaitingroomAPIResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       waitingroomAPIResponseCollectionResultInfoJSON `json:"-"`
}

// waitingroomAPIResponseCollectionResultInfoJSON contains the JSON metadata for
// the struct [WaitingroomAPIResponseCollectionResultInfo]
type waitingroomAPIResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingroomAPIResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingroomAPIResponseCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type WaitingroomMessage struct {
	Code    int64                  `json:"code,required"`
	Message string                 `json:"message,required"`
	JSON    waitingroomMessageJSON `json:"-"`
}

// waitingroomMessageJSON contains the JSON metadata for the struct
// [WaitingroomMessage]
type waitingroomMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingroomMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingroomMessageJSON) RawJSON() string {
	return r.raw
}

type WaitingroomResponseCollection struct {
	Result []Waitingroom                     `json:"result"`
	JSON   waitingroomResponseCollectionJSON `json:"-"`
	WaitingroomAPIResponseCollection
}

// waitingroomResponseCollectionJSON contains the JSON metadata for the struct
// [WaitingroomResponseCollection]
type waitingroomResponseCollectionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingroomResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingroomResponseCollectionJSON) RawJSON() string {
	return r.raw
}
