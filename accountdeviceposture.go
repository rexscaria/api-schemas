// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountDevicePostureService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDevicePostureService] method instead.
type AccountDevicePostureService struct {
	Options     []option.RequestOption
	Integration *AccountDevicePostureIntegrationService
}

// NewAccountDevicePostureService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDevicePostureService(opts ...option.RequestOption) (r *AccountDevicePostureService) {
	r = &AccountDevicePostureService{}
	r.Options = opts
	r.Integration = NewAccountDevicePostureIntegrationService(opts...)
	return
}

// Creates a new device posture rule.
func (r *AccountDevicePostureService) New(ctx context.Context, accountID interface{}, body AccountDevicePostureNewParams, opts ...option.RequestOption) (res *SingleResponsePosture, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single device posture rule.
func (r *AccountDevicePostureService) Get(ctx context.Context, accountID interface{}, ruleID string, opts ...option.RequestOption) (res *SingleResponsePosture, err error) {
	opts = append(r.Options[:], opts...)
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a device posture rule.
func (r *AccountDevicePostureService) Update(ctx context.Context, accountID interface{}, ruleID string, body AccountDevicePostureUpdateParams, opts ...option.RequestOption) (res *SingleResponsePosture, err error) {
	opts = append(r.Options[:], opts...)
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches device posture rules for a Zero Trust account.
func (r *AccountDevicePostureService) List(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *AccountDevicePostureListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/posture", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a device posture rule.
func (r *AccountDevicePostureService) Delete(ctx context.Context, accountID interface{}, ruleID string, body AccountDevicePostureDeleteParams, opts ...option.RequestOption) (res *AccountDevicePostureDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/devices/posture/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type DevicePostureRules struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input InputValue `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []MatchItem `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type DeviceTypePostureRule  `json:"type"`
	JSON devicePostureRulesJSON `json:"-"`
}

// devicePostureRulesJSON contains the JSON metadata for the struct
// [DevicePostureRules]
type devicePostureRulesJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expiration  apijson.Field
	Input       apijson.Field
	Match       apijson.Field
	Name        apijson.Field
	Schedule    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureRules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRulesJSON) RawJSON() string {
	return r.raw
}

// The type of device posture rule.
type DeviceTypePostureRule string

const (
	DeviceTypePostureRuleFile                DeviceTypePostureRule = "file"
	DeviceTypePostureRuleApplication         DeviceTypePostureRule = "application"
	DeviceTypePostureRuleTanium              DeviceTypePostureRule = "tanium"
	DeviceTypePostureRuleGateway             DeviceTypePostureRule = "gateway"
	DeviceTypePostureRuleWarp                DeviceTypePostureRule = "warp"
	DeviceTypePostureRuleDiskEncryption      DeviceTypePostureRule = "disk_encryption"
	DeviceTypePostureRuleSentinelone         DeviceTypePostureRule = "sentinelone"
	DeviceTypePostureRuleCarbonblack         DeviceTypePostureRule = "carbonblack"
	DeviceTypePostureRuleFirewall            DeviceTypePostureRule = "firewall"
	DeviceTypePostureRuleOsVersion           DeviceTypePostureRule = "os_version"
	DeviceTypePostureRuleDomainJoined        DeviceTypePostureRule = "domain_joined"
	DeviceTypePostureRuleClientCertificate   DeviceTypePostureRule = "client_certificate"
	DeviceTypePostureRuleClientCertificateV2 DeviceTypePostureRule = "client_certificate_v2"
	DeviceTypePostureRuleUniqueClientID      DeviceTypePostureRule = "unique_client_id"
	DeviceTypePostureRuleKolide              DeviceTypePostureRule = "kolide"
	DeviceTypePostureRuleTaniumS2s           DeviceTypePostureRule = "tanium_s2s"
	DeviceTypePostureRuleCrowdstrikeS2s      DeviceTypePostureRule = "crowdstrike_s2s"
	DeviceTypePostureRuleIntune              DeviceTypePostureRule = "intune"
	DeviceTypePostureRuleWorkspaceOne        DeviceTypePostureRule = "workspace_one"
	DeviceTypePostureRuleSentineloneS2s      DeviceTypePostureRule = "sentinelone_s2s"
	DeviceTypePostureRuleCustomS2s           DeviceTypePostureRule = "custom_s2s"
)

func (r DeviceTypePostureRule) IsKnown() bool {
	switch r {
	case DeviceTypePostureRuleFile, DeviceTypePostureRuleApplication, DeviceTypePostureRuleTanium, DeviceTypePostureRuleGateway, DeviceTypePostureRuleWarp, DeviceTypePostureRuleDiskEncryption, DeviceTypePostureRuleSentinelone, DeviceTypePostureRuleCarbonblack, DeviceTypePostureRuleFirewall, DeviceTypePostureRuleOsVersion, DeviceTypePostureRuleDomainJoined, DeviceTypePostureRuleClientCertificate, DeviceTypePostureRuleClientCertificateV2, DeviceTypePostureRuleUniqueClientID, DeviceTypePostureRuleKolide, DeviceTypePostureRuleTaniumS2s, DeviceTypePostureRuleCrowdstrikeS2s, DeviceTypePostureRuleIntune, DeviceTypePostureRuleWorkspaceOne, DeviceTypePostureRuleSentineloneS2s, DeviceTypePostureRuleCustomS2s:
		return true
	}
	return false
}

// The value to be checked against.
type InputValue struct {
	// List ID.
	ID string `json:"id"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id"`
	// Confirm the certificate was not imported from another device. We recommend
	// keeping this enabled unless the certificate was deployed without a private key.
	CheckPrivateKey bool `json:"check_private_key"`
	// This field can have the runtime type of [[]string].
	CheckDisks interface{} `json:"checkDisks"`
	// Common Name that is protected by the certificate
	Cn string `json:"cn"`
	// Compliance Status
	ComplianceStatus InputValueComplianceStatus `json:"compliance_status"`
	// Posture Integration ID.
	ConnectionID string `json:"connection_id"`
	// Count Operator
	CountOperator InputValueCountOperator `json:"countOperator"`
	// Domain
	Domain string `json:"domain"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Enabled
	Enabled bool `json:"enabled"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// This field can have the runtime type of
	// [[]InputValueTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsage].
	ExtendedKeyUsage interface{} `json:"extended_key_usage"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// The Number of Issues.
	IssueCount string `json:"issue_count"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// This field can have the runtime type of
	// [InputValueTeamsDevicesClientCertificateV2InputRequestLocations].
	Locations interface{} `json:"locations"`
	// Network status of device.
	NetworkStatus InputValueNetworkStatus `json:"network_status"`
	// Operating system
	OperatingSystem InputValueOperatingSystem `json:"operating_system"`
	// Agent operational state.
	OperationalState InputValueOperationalState `json:"operational_state"`
	// Operator
	Operator InputValueOperator `json:"operator"`
	// Os Version
	Os string `json:"os"`
	// Operating System Distribution Name (linux only)
	OsDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Version Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OsVersionExtra string `json:"os_version_extra"`
	// overall
	Overall string `json:"overall"`
	// File path.
	Path string `json:"path"`
	// Whether to check all disks for encryption.
	RequireAll bool `json:"requireAll"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel InputValueRiskLevel `json:"risk_level"`
	// A value between 0-100 assigned to devices set by the 3rd party posture provider.
	Score float64 `json:"score"`
	// Score Operator
	ScoreOperator InputValueScoreOperator `json:"scoreOperator"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State InputValueState `json:"state"`
	// Signing certificate thumbprint.
	Thumbprint string `json:"thumbprint"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64 `json:"total_score"`
	// Version of OS
	Version string `json:"version"`
	// Version Operator
	VersionOperator InputValueVersionOperator `json:"versionOperator"`
	JSON            inputValueJSON            `json:"-"`
	union           InputValueUnion
}

// inputValueJSON contains the JSON metadata for the struct [InputValue]
type inputValueJSON struct {
	ID               apijson.Field
	ActiveThreats    apijson.Field
	CertificateID    apijson.Field
	CheckPrivateKey  apijson.Field
	CheckDisks       apijson.Field
	Cn               apijson.Field
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	CountOperator    apijson.Field
	Domain           apijson.Field
	EidLastSeen      apijson.Field
	Enabled          apijson.Field
	Exists           apijson.Field
	ExtendedKeyUsage apijson.Field
	Infected         apijson.Field
	IsActive         apijson.Field
	IssueCount       apijson.Field
	LastSeen         apijson.Field
	Locations        apijson.Field
	NetworkStatus    apijson.Field
	OperatingSystem  apijson.Field
	OperationalState apijson.Field
	Operator         apijson.Field
	Os               apijson.Field
	OsDistroName     apijson.Field
	OsDistroRevision apijson.Field
	OsVersionExtra   apijson.Field
	Overall          apijson.Field
	Path             apijson.Field
	RequireAll       apijson.Field
	RiskLevel        apijson.Field
	Score            apijson.Field
	ScoreOperator    apijson.Field
	SensorConfig     apijson.Field
	Sha256           apijson.Field
	State            apijson.Field
	Thumbprint       apijson.Field
	TotalScore       apijson.Field
	Version          apijson.Field
	VersionOperator  apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r inputValueJSON) RawJSON() string {
	return r.raw
}

func (r *InputValue) UnmarshalJSON(data []byte) (err error) {
	*r = InputValue{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InputValueUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are
// [InputValueTeamsDevicesFileInputRequest],
// [InputValueTeamsDevicesUniqueClientIDInputRequest],
// [InputValueTeamsDevicesDomainJoinedInputRequest],
// [InputValueTeamsDevicesOsVersionInputRequest],
// [InputValueTeamsDevicesFirewallInputRequest],
// [InputValueTeamsDevicesSentineloneInputRequest],
// [InputValueTeamsDevicesCarbonblackInputRequest],
// [InputValueTeamsDevicesDiskEncryptionInputRequest],
// [InputValueTeamsDevicesApplicationInputRequest],
// [InputValueTeamsDevicesClientCertificateInputRequest],
// [InputValueTeamsDevicesClientCertificateV2InputRequest],
// [InputValueTeamsDevicesWorkspaceOneInputRequest],
// [InputValueTeamsDevicesCrowdstrikeInputRequest],
// [InputValueTeamsDevicesIntuneInputRequest],
// [InputValueTeamsDevicesKolideInputRequest],
// [InputValueTeamsDevicesTaniumInputRequest],
// [InputValueTeamsDevicesSentineloneS2sInputRequest],
// [InputValueTeamsDevicesCustomS2sInputRequest].
func (r InputValue) AsUnion() InputValueUnion {
	return r.union
}

// The value to be checked against.
//
// Union satisfied by [InputValueTeamsDevicesFileInputRequest],
// [InputValueTeamsDevicesUniqueClientIDInputRequest],
// [InputValueTeamsDevicesDomainJoinedInputRequest],
// [InputValueTeamsDevicesOsVersionInputRequest],
// [InputValueTeamsDevicesFirewallInputRequest],
// [InputValueTeamsDevicesSentineloneInputRequest],
// [InputValueTeamsDevicesCarbonblackInputRequest],
// [InputValueTeamsDevicesDiskEncryptionInputRequest],
// [InputValueTeamsDevicesApplicationInputRequest],
// [InputValueTeamsDevicesClientCertificateInputRequest],
// [InputValueTeamsDevicesClientCertificateV2InputRequest],
// [InputValueTeamsDevicesWorkspaceOneInputRequest],
// [InputValueTeamsDevicesCrowdstrikeInputRequest],
// [InputValueTeamsDevicesIntuneInputRequest],
// [InputValueTeamsDevicesKolideInputRequest],
// [InputValueTeamsDevicesTaniumInputRequest],
// [InputValueTeamsDevicesSentineloneS2sInputRequest] or
// [InputValueTeamsDevicesCustomS2sInputRequest].
type InputValueUnion interface {
	implementsInputValue()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InputValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputValueTeamsDevicesFileInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputValueTeamsDevicesUniqueClientIDInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputValueTeamsDevicesDomainJoinedInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputValueTeamsDevicesOsVersionInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputValueTeamsDevicesFirewallInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputValueTeamsDevicesSentineloneInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputValueTeamsDevicesCarbonblackInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputValueTeamsDevicesDiskEncryptionInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputValueTeamsDevicesApplicationInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputValueTeamsDevicesClientCertificateInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputValueTeamsDevicesClientCertificateV2InputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputValueTeamsDevicesWorkspaceOneInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputValueTeamsDevicesCrowdstrikeInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputValueTeamsDevicesIntuneInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputValueTeamsDevicesKolideInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputValueTeamsDevicesTaniumInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputValueTeamsDevicesSentineloneS2sInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InputValueTeamsDevicesCustomS2sInputRequest{}),
		},
	)
}

type InputValueTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem InputValueTeamsDevicesFileInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                     `json:"thumbprint"`
	JSON       inputValueTeamsDevicesFileInputRequestJSON `json:"-"`
}

// inputValueTeamsDevicesFileInputRequestJSON contains the JSON metadata for the
// struct [InputValueTeamsDevicesFileInputRequest]
type inputValueTeamsDevicesFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InputValueTeamsDevicesFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesFileInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputValueTeamsDevicesFileInputRequest) implementsInputValue() {}

// Operating system
type InputValueTeamsDevicesFileInputRequestOperatingSystem string

const (
	InputValueTeamsDevicesFileInputRequestOperatingSystemWindows InputValueTeamsDevicesFileInputRequestOperatingSystem = "windows"
	InputValueTeamsDevicesFileInputRequestOperatingSystemLinux   InputValueTeamsDevicesFileInputRequestOperatingSystem = "linux"
	InputValueTeamsDevicesFileInputRequestOperatingSystemMac     InputValueTeamsDevicesFileInputRequestOperatingSystem = "mac"
)

func (r InputValueTeamsDevicesFileInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesFileInputRequestOperatingSystemWindows, InputValueTeamsDevicesFileInputRequestOperatingSystemLinux, InputValueTeamsDevicesFileInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type InputValueTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem InputValueTeamsDevicesUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            inputValueTeamsDevicesUniqueClientIDInputRequestJSON            `json:"-"`
}

// inputValueTeamsDevicesUniqueClientIDInputRequestJSON contains the JSON metadata
// for the struct [InputValueTeamsDevicesUniqueClientIDInputRequest]
type inputValueTeamsDevicesUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InputValueTeamsDevicesUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesUniqueClientIDInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputValueTeamsDevicesUniqueClientIDInputRequest) implementsInputValue() {}

// Operating System
type InputValueTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	InputValueTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  InputValueTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	InputValueTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      InputValueTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	InputValueTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos InputValueTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

func (r InputValueTeamsDevicesUniqueClientIDInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid, InputValueTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos, InputValueTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos:
		return true
	}
	return false
}

type InputValueTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem InputValueTeamsDevicesDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                             `json:"domain"`
	JSON   inputValueTeamsDevicesDomainJoinedInputRequestJSON `json:"-"`
}

// inputValueTeamsDevicesDomainJoinedInputRequestJSON contains the JSON metadata
// for the struct [InputValueTeamsDevicesDomainJoinedInputRequest]
type inputValueTeamsDevicesDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InputValueTeamsDevicesDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesDomainJoinedInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputValueTeamsDevicesDomainJoinedInputRequest) implementsInputValue() {}

// Operating System
type InputValueTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	InputValueTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows InputValueTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

func (r InputValueTeamsDevicesDomainJoinedInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows:
		return true
	}
	return false
}

type InputValueTeamsDevicesOsVersionInputRequest struct {
	// Operating System
	OperatingSystem InputValueTeamsDevicesOsVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// Operator
	Operator InputValueTeamsDevicesOsVersionInputRequestOperator `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OsDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Version Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OsVersionExtra string                                          `json:"os_version_extra"`
	JSON           inputValueTeamsDevicesOsVersionInputRequestJSON `json:"-"`
}

// inputValueTeamsDevicesOsVersionInputRequestJSON contains the JSON metadata for
// the struct [InputValueTeamsDevicesOsVersionInputRequest]
type inputValueTeamsDevicesOsVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OsDistroName     apijson.Field
	OsDistroRevision apijson.Field
	OsVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InputValueTeamsDevicesOsVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesOsVersionInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputValueTeamsDevicesOsVersionInputRequest) implementsInputValue() {}

// Operating System
type InputValueTeamsDevicesOsVersionInputRequestOperatingSystem string

const (
	InputValueTeamsDevicesOsVersionInputRequestOperatingSystemWindows InputValueTeamsDevicesOsVersionInputRequestOperatingSystem = "windows"
)

func (r InputValueTeamsDevicesOsVersionInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesOsVersionInputRequestOperatingSystemWindows:
		return true
	}
	return false
}

// Operator
type InputValueTeamsDevicesOsVersionInputRequestOperator string

const (
	InputValueTeamsDevicesOsVersionInputRequestOperatorLess            InputValueTeamsDevicesOsVersionInputRequestOperator = "<"
	InputValueTeamsDevicesOsVersionInputRequestOperatorLessOrEquals    InputValueTeamsDevicesOsVersionInputRequestOperator = "<="
	InputValueTeamsDevicesOsVersionInputRequestOperatorGreater         InputValueTeamsDevicesOsVersionInputRequestOperator = ">"
	InputValueTeamsDevicesOsVersionInputRequestOperatorGreaterOrEquals InputValueTeamsDevicesOsVersionInputRequestOperator = ">="
	InputValueTeamsDevicesOsVersionInputRequestOperatorEquals          InputValueTeamsDevicesOsVersionInputRequestOperator = "=="
)

func (r InputValueTeamsDevicesOsVersionInputRequestOperator) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesOsVersionInputRequestOperatorLess, InputValueTeamsDevicesOsVersionInputRequestOperatorLessOrEquals, InputValueTeamsDevicesOsVersionInputRequestOperatorGreater, InputValueTeamsDevicesOsVersionInputRequestOperatorGreaterOrEquals, InputValueTeamsDevicesOsVersionInputRequestOperatorEquals:
		return true
	}
	return false
}

type InputValueTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem InputValueTeamsDevicesFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            inputValueTeamsDevicesFirewallInputRequestJSON            `json:"-"`
}

// inputValueTeamsDevicesFirewallInputRequestJSON contains the JSON metadata for
// the struct [InputValueTeamsDevicesFirewallInputRequest]
type inputValueTeamsDevicesFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InputValueTeamsDevicesFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesFirewallInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputValueTeamsDevicesFirewallInputRequest) implementsInputValue() {}

// Operating System
type InputValueTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	InputValueTeamsDevicesFirewallInputRequestOperatingSystemWindows InputValueTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	InputValueTeamsDevicesFirewallInputRequestOperatingSystemMac     InputValueTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

func (r InputValueTeamsDevicesFirewallInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesFirewallInputRequestOperatingSystemWindows, InputValueTeamsDevicesFirewallInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type InputValueTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem InputValueTeamsDevicesSentineloneInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                            `json:"thumbprint"`
	JSON       inputValueTeamsDevicesSentineloneInputRequestJSON `json:"-"`
}

// inputValueTeamsDevicesSentineloneInputRequestJSON contains the JSON metadata for
// the struct [InputValueTeamsDevicesSentineloneInputRequest]
type inputValueTeamsDevicesSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InputValueTeamsDevicesSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesSentineloneInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputValueTeamsDevicesSentineloneInputRequest) implementsInputValue() {}

// Operating system
type InputValueTeamsDevicesSentineloneInputRequestOperatingSystem string

const (
	InputValueTeamsDevicesSentineloneInputRequestOperatingSystemWindows InputValueTeamsDevicesSentineloneInputRequestOperatingSystem = "windows"
	InputValueTeamsDevicesSentineloneInputRequestOperatingSystemLinux   InputValueTeamsDevicesSentineloneInputRequestOperatingSystem = "linux"
	InputValueTeamsDevicesSentineloneInputRequestOperatingSystemMac     InputValueTeamsDevicesSentineloneInputRequestOperatingSystem = "mac"
)

func (r InputValueTeamsDevicesSentineloneInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesSentineloneInputRequestOperatingSystemWindows, InputValueTeamsDevicesSentineloneInputRequestOperatingSystemLinux, InputValueTeamsDevicesSentineloneInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type InputValueTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem InputValueTeamsDevicesCarbonblackInputRequestOperatingSystem `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                            `json:"thumbprint"`
	JSON       inputValueTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// inputValueTeamsDevicesCarbonblackInputRequestJSON contains the JSON metadata for
// the struct [InputValueTeamsDevicesCarbonblackInputRequest]
type inputValueTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InputValueTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesCarbonblackInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputValueTeamsDevicesCarbonblackInputRequest) implementsInputValue() {}

// Operating system
type InputValueTeamsDevicesCarbonblackInputRequestOperatingSystem string

const (
	InputValueTeamsDevicesCarbonblackInputRequestOperatingSystemWindows InputValueTeamsDevicesCarbonblackInputRequestOperatingSystem = "windows"
	InputValueTeamsDevicesCarbonblackInputRequestOperatingSystemLinux   InputValueTeamsDevicesCarbonblackInputRequestOperatingSystem = "linux"
	InputValueTeamsDevicesCarbonblackInputRequestOperatingSystemMac     InputValueTeamsDevicesCarbonblackInputRequestOperatingSystem = "mac"
)

func (r InputValueTeamsDevicesCarbonblackInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesCarbonblackInputRequestOperatingSystemWindows, InputValueTeamsDevicesCarbonblackInputRequestOperatingSystemLinux, InputValueTeamsDevicesCarbonblackInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type InputValueTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                 `json:"requireAll"`
	JSON       inputValueTeamsDevicesDiskEncryptionInputRequestJSON `json:"-"`
}

// inputValueTeamsDevicesDiskEncryptionInputRequestJSON contains the JSON metadata
// for the struct [InputValueTeamsDevicesDiskEncryptionInputRequest]
type inputValueTeamsDevicesDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InputValueTeamsDevicesDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesDiskEncryptionInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputValueTeamsDevicesDiskEncryptionInputRequest) implementsInputValue() {}

type InputValueTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem InputValueTeamsDevicesApplicationInputRequestOperatingSystem `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                            `json:"thumbprint"`
	JSON       inputValueTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// inputValueTeamsDevicesApplicationInputRequestJSON contains the JSON metadata for
// the struct [InputValueTeamsDevicesApplicationInputRequest]
type inputValueTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InputValueTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesApplicationInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputValueTeamsDevicesApplicationInputRequest) implementsInputValue() {}

// Operating system
type InputValueTeamsDevicesApplicationInputRequestOperatingSystem string

const (
	InputValueTeamsDevicesApplicationInputRequestOperatingSystemWindows InputValueTeamsDevicesApplicationInputRequestOperatingSystem = "windows"
	InputValueTeamsDevicesApplicationInputRequestOperatingSystemLinux   InputValueTeamsDevicesApplicationInputRequestOperatingSystem = "linux"
	InputValueTeamsDevicesApplicationInputRequestOperatingSystemMac     InputValueTeamsDevicesApplicationInputRequestOperatingSystem = "mac"
)

func (r InputValueTeamsDevicesApplicationInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesApplicationInputRequestOperatingSystemWindows, InputValueTeamsDevicesApplicationInputRequestOperatingSystemLinux, InputValueTeamsDevicesApplicationInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type InputValueTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                  `json:"cn,required"`
	JSON inputValueTeamsDevicesClientCertificateInputRequestJSON `json:"-"`
}

// inputValueTeamsDevicesClientCertificateInputRequestJSON contains the JSON
// metadata for the struct [InputValueTeamsDevicesClientCertificateInputRequest]
type inputValueTeamsDevicesClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InputValueTeamsDevicesClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesClientCertificateInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputValueTeamsDevicesClientCertificateInputRequest) implementsInputValue() {}

type InputValueTeamsDevicesClientCertificateV2InputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Confirm the certificate was not imported from another device. We recommend
	// keeping this enabled unless the certificate was deployed without a private key.
	CheckPrivateKey bool `json:"check_private_key,required"`
	// Operating System
	OperatingSystem InputValueTeamsDevicesClientCertificateV2InputRequestOperatingSystem `json:"operating_system,required"`
	// Common Name that is protected by the client certificate. This may include one or
	// more variables in the ${ } notation. Only ${serial_number} and ${hostname} are
	// valid variables.
	Cn string `json:"cn"`
	// List of values indicating purposes for which the certificate public key can be
	// used
	ExtendedKeyUsage []InputValueTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsage `json:"extended_key_usage"`
	Locations        InputValueTeamsDevicesClientCertificateV2InputRequestLocations          `json:"locations"`
	JSON             inputValueTeamsDevicesClientCertificateV2InputRequestJSON               `json:"-"`
}

// inputValueTeamsDevicesClientCertificateV2InputRequestJSON contains the JSON
// metadata for the struct [InputValueTeamsDevicesClientCertificateV2InputRequest]
type inputValueTeamsDevicesClientCertificateV2InputRequestJSON struct {
	CertificateID    apijson.Field
	CheckPrivateKey  apijson.Field
	OperatingSystem  apijson.Field
	Cn               apijson.Field
	ExtendedKeyUsage apijson.Field
	Locations        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InputValueTeamsDevicesClientCertificateV2InputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesClientCertificateV2InputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputValueTeamsDevicesClientCertificateV2InputRequest) implementsInputValue() {}

// Operating System
type InputValueTeamsDevicesClientCertificateV2InputRequestOperatingSystem string

const (
	InputValueTeamsDevicesClientCertificateV2InputRequestOperatingSystemWindows InputValueTeamsDevicesClientCertificateV2InputRequestOperatingSystem = "windows"
	InputValueTeamsDevicesClientCertificateV2InputRequestOperatingSystemMac     InputValueTeamsDevicesClientCertificateV2InputRequestOperatingSystem = "mac"
	InputValueTeamsDevicesClientCertificateV2InputRequestOperatingSystemLinux   InputValueTeamsDevicesClientCertificateV2InputRequestOperatingSystem = "linux"
)

func (r InputValueTeamsDevicesClientCertificateV2InputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesClientCertificateV2InputRequestOperatingSystemWindows, InputValueTeamsDevicesClientCertificateV2InputRequestOperatingSystemMac, InputValueTeamsDevicesClientCertificateV2InputRequestOperatingSystemLinux:
		return true
	}
	return false
}

type InputValueTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsage string

const (
	InputValueTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsageClientAuth      InputValueTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsage = "clientAuth"
	InputValueTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsageEmailProtection InputValueTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsage = "emailProtection"
)

func (r InputValueTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsage) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsageClientAuth, InputValueTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsageEmailProtection:
		return true
	}
	return false
}

type InputValueTeamsDevicesClientCertificateV2InputRequestLocations struct {
	// List of paths to check for client certificate on linux.
	Paths []string `json:"paths"`
	// List of trust stores to check for client certificate.
	TrustStores []InputValueTeamsDevicesClientCertificateV2InputRequestLocationsTrustStore `json:"trust_stores"`
	JSON        inputValueTeamsDevicesClientCertificateV2InputRequestLocationsJSON         `json:"-"`
}

// inputValueTeamsDevicesClientCertificateV2InputRequestLocationsJSON contains the
// JSON metadata for the struct
// [InputValueTeamsDevicesClientCertificateV2InputRequestLocations]
type inputValueTeamsDevicesClientCertificateV2InputRequestLocationsJSON struct {
	Paths       apijson.Field
	TrustStores apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InputValueTeamsDevicesClientCertificateV2InputRequestLocations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesClientCertificateV2InputRequestLocationsJSON) RawJSON() string {
	return r.raw
}

type InputValueTeamsDevicesClientCertificateV2InputRequestLocationsTrustStore string

const (
	InputValueTeamsDevicesClientCertificateV2InputRequestLocationsTrustStoreSystem InputValueTeamsDevicesClientCertificateV2InputRequestLocationsTrustStore = "system"
	InputValueTeamsDevicesClientCertificateV2InputRequestLocationsTrustStoreUser   InputValueTeamsDevicesClientCertificateV2InputRequestLocationsTrustStore = "user"
)

func (r InputValueTeamsDevicesClientCertificateV2InputRequestLocationsTrustStore) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesClientCertificateV2InputRequestLocationsTrustStoreSystem, InputValueTeamsDevicesClientCertificateV2InputRequestLocationsTrustStoreUser:
		return true
	}
	return false
}

type InputValueTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus InputValueTeamsDevicesWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                             `json:"connection_id,required"`
	JSON         inputValueTeamsDevicesWorkspaceOneInputRequestJSON `json:"-"`
}

// inputValueTeamsDevicesWorkspaceOneInputRequestJSON contains the JSON metadata
// for the struct [InputValueTeamsDevicesWorkspaceOneInputRequest]
type inputValueTeamsDevicesWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InputValueTeamsDevicesWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesWorkspaceOneInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputValueTeamsDevicesWorkspaceOneInputRequest) implementsInputValue() {}

// Compliance Status
type InputValueTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	InputValueTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    InputValueTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	InputValueTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant InputValueTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	InputValueTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      InputValueTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

func (r InputValueTeamsDevicesWorkspaceOneInputRequestComplianceStatus) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant, InputValueTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant, InputValueTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown:
		return true
	}
	return false
}

type InputValueTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// Operator
	Operator InputValueTeamsDevicesCrowdstrikeInputRequestOperator `json:"operator"`
	// Os Version
	Os string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State InputValueTeamsDevicesCrowdstrikeInputRequestState `json:"state"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            inputValueTeamsDevicesCrowdstrikeInputRequestJSON            `json:"-"`
}

// inputValueTeamsDevicesCrowdstrikeInputRequestJSON contains the JSON metadata for
// the struct [InputValueTeamsDevicesCrowdstrikeInputRequest]
type inputValueTeamsDevicesCrowdstrikeInputRequestJSON struct {
	ConnectionID    apijson.Field
	LastSeen        apijson.Field
	Operator        apijson.Field
	Os              apijson.Field
	Overall         apijson.Field
	SensorConfig    apijson.Field
	State           apijson.Field
	Version         apijson.Field
	VersionOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InputValueTeamsDevicesCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesCrowdstrikeInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputValueTeamsDevicesCrowdstrikeInputRequest) implementsInputValue() {}

// Operator
type InputValueTeamsDevicesCrowdstrikeInputRequestOperator string

const (
	InputValueTeamsDevicesCrowdstrikeInputRequestOperatorLess            InputValueTeamsDevicesCrowdstrikeInputRequestOperator = "<"
	InputValueTeamsDevicesCrowdstrikeInputRequestOperatorLessOrEquals    InputValueTeamsDevicesCrowdstrikeInputRequestOperator = "<="
	InputValueTeamsDevicesCrowdstrikeInputRequestOperatorGreater         InputValueTeamsDevicesCrowdstrikeInputRequestOperator = ">"
	InputValueTeamsDevicesCrowdstrikeInputRequestOperatorGreaterOrEquals InputValueTeamsDevicesCrowdstrikeInputRequestOperator = ">="
	InputValueTeamsDevicesCrowdstrikeInputRequestOperatorEquals          InputValueTeamsDevicesCrowdstrikeInputRequestOperator = "=="
)

func (r InputValueTeamsDevicesCrowdstrikeInputRequestOperator) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesCrowdstrikeInputRequestOperatorLess, InputValueTeamsDevicesCrowdstrikeInputRequestOperatorLessOrEquals, InputValueTeamsDevicesCrowdstrikeInputRequestOperatorGreater, InputValueTeamsDevicesCrowdstrikeInputRequestOperatorGreaterOrEquals, InputValueTeamsDevicesCrowdstrikeInputRequestOperatorEquals:
		return true
	}
	return false
}

// For more details on state, please refer to the Crowdstrike documentation.
type InputValueTeamsDevicesCrowdstrikeInputRequestState string

const (
	InputValueTeamsDevicesCrowdstrikeInputRequestStateOnline  InputValueTeamsDevicesCrowdstrikeInputRequestState = "online"
	InputValueTeamsDevicesCrowdstrikeInputRequestStateOffline InputValueTeamsDevicesCrowdstrikeInputRequestState = "offline"
	InputValueTeamsDevicesCrowdstrikeInputRequestStateUnknown InputValueTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

func (r InputValueTeamsDevicesCrowdstrikeInputRequestState) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesCrowdstrikeInputRequestStateOnline, InputValueTeamsDevicesCrowdstrikeInputRequestStateOffline, InputValueTeamsDevicesCrowdstrikeInputRequestStateUnknown:
		return true
	}
	return false
}

// Version Operator
type InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperatorLess            InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperatorLessOrEquals    InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreater         InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreaterOrEquals InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperatorEquals          InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

func (r InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperator) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperatorLess, InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperatorLessOrEquals, InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreater, InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreaterOrEquals, InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperatorEquals:
		return true
	}
	return false
}

type InputValueTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus InputValueTeamsDevicesIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                       `json:"connection_id,required"`
	JSON         inputValueTeamsDevicesIntuneInputRequestJSON `json:"-"`
}

// inputValueTeamsDevicesIntuneInputRequestJSON contains the JSON metadata for the
// struct [InputValueTeamsDevicesIntuneInputRequest]
type inputValueTeamsDevicesIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InputValueTeamsDevicesIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesIntuneInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputValueTeamsDevicesIntuneInputRequest) implementsInputValue() {}

// Compliance Status
type InputValueTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	InputValueTeamsDevicesIntuneInputRequestComplianceStatusCompliant     InputValueTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	InputValueTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  InputValueTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	InputValueTeamsDevicesIntuneInputRequestComplianceStatusUnknown       InputValueTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	InputValueTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable InputValueTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	InputValueTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod InputValueTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	InputValueTeamsDevicesIntuneInputRequestComplianceStatusError         InputValueTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

func (r InputValueTeamsDevicesIntuneInputRequestComplianceStatus) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesIntuneInputRequestComplianceStatusCompliant, InputValueTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant, InputValueTeamsDevicesIntuneInputRequestComplianceStatusUnknown, InputValueTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable, InputValueTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod, InputValueTeamsDevicesIntuneInputRequestComplianceStatusError:
		return true
	}
	return false
}

type InputValueTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator InputValueTeamsDevicesKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                       `json:"issue_count,required"`
	JSON       inputValueTeamsDevicesKolideInputRequestJSON `json:"-"`
}

// inputValueTeamsDevicesKolideInputRequestJSON contains the JSON metadata for the
// struct [InputValueTeamsDevicesKolideInputRequest]
type inputValueTeamsDevicesKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InputValueTeamsDevicesKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesKolideInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputValueTeamsDevicesKolideInputRequest) implementsInputValue() {}

// Count Operator
type InputValueTeamsDevicesKolideInputRequestCountOperator string

const (
	InputValueTeamsDevicesKolideInputRequestCountOperatorLess            InputValueTeamsDevicesKolideInputRequestCountOperator = "<"
	InputValueTeamsDevicesKolideInputRequestCountOperatorLessOrEquals    InputValueTeamsDevicesKolideInputRequestCountOperator = "<="
	InputValueTeamsDevicesKolideInputRequestCountOperatorGreater         InputValueTeamsDevicesKolideInputRequestCountOperator = ">"
	InputValueTeamsDevicesKolideInputRequestCountOperatorGreaterOrEquals InputValueTeamsDevicesKolideInputRequestCountOperator = ">="
	InputValueTeamsDevicesKolideInputRequestCountOperatorEquals          InputValueTeamsDevicesKolideInputRequestCountOperator = "=="
)

func (r InputValueTeamsDevicesKolideInputRequestCountOperator) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesKolideInputRequestCountOperatorLess, InputValueTeamsDevicesKolideInputRequestCountOperatorLessOrEquals, InputValueTeamsDevicesKolideInputRequestCountOperatorGreater, InputValueTeamsDevicesKolideInputRequestCountOperatorGreaterOrEquals, InputValueTeamsDevicesKolideInputRequestCountOperatorEquals:
		return true
	}
	return false
}

type InputValueTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator InputValueTeamsDevicesTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel InputValueTeamsDevicesTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator InputValueTeamsDevicesTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                      `json:"total_score"`
	JSON       inputValueTeamsDevicesTaniumInputRequestJSON `json:"-"`
}

// inputValueTeamsDevicesTaniumInputRequestJSON contains the JSON metadata for the
// struct [InputValueTeamsDevicesTaniumInputRequest]
type inputValueTeamsDevicesTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InputValueTeamsDevicesTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesTaniumInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputValueTeamsDevicesTaniumInputRequest) implementsInputValue() {}

// Operator to evaluate risk_level or eid_last_seen.
type InputValueTeamsDevicesTaniumInputRequestOperator string

const (
	InputValueTeamsDevicesTaniumInputRequestOperatorLess            InputValueTeamsDevicesTaniumInputRequestOperator = "<"
	InputValueTeamsDevicesTaniumInputRequestOperatorLessOrEquals    InputValueTeamsDevicesTaniumInputRequestOperator = "<="
	InputValueTeamsDevicesTaniumInputRequestOperatorGreater         InputValueTeamsDevicesTaniumInputRequestOperator = ">"
	InputValueTeamsDevicesTaniumInputRequestOperatorGreaterOrEquals InputValueTeamsDevicesTaniumInputRequestOperator = ">="
	InputValueTeamsDevicesTaniumInputRequestOperatorEquals          InputValueTeamsDevicesTaniumInputRequestOperator = "=="
)

func (r InputValueTeamsDevicesTaniumInputRequestOperator) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesTaniumInputRequestOperatorLess, InputValueTeamsDevicesTaniumInputRequestOperatorLessOrEquals, InputValueTeamsDevicesTaniumInputRequestOperatorGreater, InputValueTeamsDevicesTaniumInputRequestOperatorGreaterOrEquals, InputValueTeamsDevicesTaniumInputRequestOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type InputValueTeamsDevicesTaniumInputRequestRiskLevel string

const (
	InputValueTeamsDevicesTaniumInputRequestRiskLevelLow      InputValueTeamsDevicesTaniumInputRequestRiskLevel = "low"
	InputValueTeamsDevicesTaniumInputRequestRiskLevelMedium   InputValueTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	InputValueTeamsDevicesTaniumInputRequestRiskLevelHigh     InputValueTeamsDevicesTaniumInputRequestRiskLevel = "high"
	InputValueTeamsDevicesTaniumInputRequestRiskLevelCritical InputValueTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

func (r InputValueTeamsDevicesTaniumInputRequestRiskLevel) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesTaniumInputRequestRiskLevelLow, InputValueTeamsDevicesTaniumInputRequestRiskLevelMedium, InputValueTeamsDevicesTaniumInputRequestRiskLevelHigh, InputValueTeamsDevicesTaniumInputRequestRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator
type InputValueTeamsDevicesTaniumInputRequestScoreOperator string

const (
	InputValueTeamsDevicesTaniumInputRequestScoreOperatorLess            InputValueTeamsDevicesTaniumInputRequestScoreOperator = "<"
	InputValueTeamsDevicesTaniumInputRequestScoreOperatorLessOrEquals    InputValueTeamsDevicesTaniumInputRequestScoreOperator = "<="
	InputValueTeamsDevicesTaniumInputRequestScoreOperatorGreater         InputValueTeamsDevicesTaniumInputRequestScoreOperator = ">"
	InputValueTeamsDevicesTaniumInputRequestScoreOperatorGreaterOrEquals InputValueTeamsDevicesTaniumInputRequestScoreOperator = ">="
	InputValueTeamsDevicesTaniumInputRequestScoreOperatorEquals          InputValueTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

func (r InputValueTeamsDevicesTaniumInputRequestScoreOperator) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesTaniumInputRequestScoreOperatorLess, InputValueTeamsDevicesTaniumInputRequestScoreOperatorLessOrEquals, InputValueTeamsDevicesTaniumInputRequestScoreOperatorGreater, InputValueTeamsDevicesTaniumInputRequestScoreOperatorGreaterOrEquals, InputValueTeamsDevicesTaniumInputRequestScoreOperatorEquals:
		return true
	}
	return false
}

type InputValueTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus InputValueTeamsDevicesSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// Agent operational state.
	OperationalState InputValueTeamsDevicesSentineloneS2sInputRequestOperationalState `json:"operational_state"`
	// operator
	Operator InputValueTeamsDevicesSentineloneS2sInputRequestOperator `json:"operator"`
	JSON     inputValueTeamsDevicesSentineloneS2sInputRequestJSON     `json:"-"`
}

// inputValueTeamsDevicesSentineloneS2sInputRequestJSON contains the JSON metadata
// for the struct [InputValueTeamsDevicesSentineloneS2sInputRequest]
type inputValueTeamsDevicesSentineloneS2sInputRequestJSON struct {
	ConnectionID     apijson.Field
	ActiveThreats    apijson.Field
	Infected         apijson.Field
	IsActive         apijson.Field
	NetworkStatus    apijson.Field
	OperationalState apijson.Field
	Operator         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InputValueTeamsDevicesSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesSentineloneS2sInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputValueTeamsDevicesSentineloneS2sInputRequest) implementsInputValue() {}

// Network status of device.
type InputValueTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	InputValueTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     InputValueTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	InputValueTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  InputValueTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	InputValueTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting InputValueTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	InputValueTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    InputValueTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

func (r InputValueTeamsDevicesSentineloneS2sInputRequestNetworkStatus) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected, InputValueTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected, InputValueTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting, InputValueTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting:
		return true
	}
	return false
}

// Agent operational state.
type InputValueTeamsDevicesSentineloneS2sInputRequestOperationalState string

const (
	InputValueTeamsDevicesSentineloneS2sInputRequestOperationalStateNa                    InputValueTeamsDevicesSentineloneS2sInputRequestOperationalState = "na"
	InputValueTeamsDevicesSentineloneS2sInputRequestOperationalStatePartiallyDisabled     InputValueTeamsDevicesSentineloneS2sInputRequestOperationalState = "partially_disabled"
	InputValueTeamsDevicesSentineloneS2sInputRequestOperationalStateAutoFullyDisabled     InputValueTeamsDevicesSentineloneS2sInputRequestOperationalState = "auto_fully_disabled"
	InputValueTeamsDevicesSentineloneS2sInputRequestOperationalStateFullyDisabled         InputValueTeamsDevicesSentineloneS2sInputRequestOperationalState = "fully_disabled"
	InputValueTeamsDevicesSentineloneS2sInputRequestOperationalStateAutoPartiallyDisabled InputValueTeamsDevicesSentineloneS2sInputRequestOperationalState = "auto_partially_disabled"
	InputValueTeamsDevicesSentineloneS2sInputRequestOperationalStateDisabledError         InputValueTeamsDevicesSentineloneS2sInputRequestOperationalState = "disabled_error"
	InputValueTeamsDevicesSentineloneS2sInputRequestOperationalStateDBCorruption          InputValueTeamsDevicesSentineloneS2sInputRequestOperationalState = "db_corruption"
)

func (r InputValueTeamsDevicesSentineloneS2sInputRequestOperationalState) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesSentineloneS2sInputRequestOperationalStateNa, InputValueTeamsDevicesSentineloneS2sInputRequestOperationalStatePartiallyDisabled, InputValueTeamsDevicesSentineloneS2sInputRequestOperationalStateAutoFullyDisabled, InputValueTeamsDevicesSentineloneS2sInputRequestOperationalStateFullyDisabled, InputValueTeamsDevicesSentineloneS2sInputRequestOperationalStateAutoPartiallyDisabled, InputValueTeamsDevicesSentineloneS2sInputRequestOperationalStateDisabledError, InputValueTeamsDevicesSentineloneS2sInputRequestOperationalStateDBCorruption:
		return true
	}
	return false
}

// operator
type InputValueTeamsDevicesSentineloneS2sInputRequestOperator string

const (
	InputValueTeamsDevicesSentineloneS2sInputRequestOperatorLess            InputValueTeamsDevicesSentineloneS2sInputRequestOperator = "<"
	InputValueTeamsDevicesSentineloneS2sInputRequestOperatorLessOrEquals    InputValueTeamsDevicesSentineloneS2sInputRequestOperator = "<="
	InputValueTeamsDevicesSentineloneS2sInputRequestOperatorGreater         InputValueTeamsDevicesSentineloneS2sInputRequestOperator = ">"
	InputValueTeamsDevicesSentineloneS2sInputRequestOperatorGreaterOrEquals InputValueTeamsDevicesSentineloneS2sInputRequestOperator = ">="
	InputValueTeamsDevicesSentineloneS2sInputRequestOperatorEquals          InputValueTeamsDevicesSentineloneS2sInputRequestOperator = "=="
)

func (r InputValueTeamsDevicesSentineloneS2sInputRequestOperator) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesSentineloneS2sInputRequestOperatorLess, InputValueTeamsDevicesSentineloneS2sInputRequestOperatorLessOrEquals, InputValueTeamsDevicesSentineloneS2sInputRequestOperatorGreater, InputValueTeamsDevicesSentineloneS2sInputRequestOperatorGreaterOrEquals, InputValueTeamsDevicesSentineloneS2sInputRequestOperatorEquals:
		return true
	}
	return false
}

type InputValueTeamsDevicesCustomS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// operator
	Operator InputValueTeamsDevicesCustomS2sInputRequestOperator `json:"operator,required"`
	// A value between 0-100 assigned to devices set by the 3rd party posture provider.
	Score float64                                         `json:"score,required"`
	JSON  inputValueTeamsDevicesCustomS2sInputRequestJSON `json:"-"`
}

// inputValueTeamsDevicesCustomS2sInputRequestJSON contains the JSON metadata for
// the struct [InputValueTeamsDevicesCustomS2sInputRequest]
type inputValueTeamsDevicesCustomS2sInputRequestJSON struct {
	ConnectionID apijson.Field
	Operator     apijson.Field
	Score        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InputValueTeamsDevicesCustomS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputValueTeamsDevicesCustomS2sInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r InputValueTeamsDevicesCustomS2sInputRequest) implementsInputValue() {}

// operator
type InputValueTeamsDevicesCustomS2sInputRequestOperator string

const (
	InputValueTeamsDevicesCustomS2sInputRequestOperatorLess            InputValueTeamsDevicesCustomS2sInputRequestOperator = "<"
	InputValueTeamsDevicesCustomS2sInputRequestOperatorLessOrEquals    InputValueTeamsDevicesCustomS2sInputRequestOperator = "<="
	InputValueTeamsDevicesCustomS2sInputRequestOperatorGreater         InputValueTeamsDevicesCustomS2sInputRequestOperator = ">"
	InputValueTeamsDevicesCustomS2sInputRequestOperatorGreaterOrEquals InputValueTeamsDevicesCustomS2sInputRequestOperator = ">="
	InputValueTeamsDevicesCustomS2sInputRequestOperatorEquals          InputValueTeamsDevicesCustomS2sInputRequestOperator = "=="
)

func (r InputValueTeamsDevicesCustomS2sInputRequestOperator) IsKnown() bool {
	switch r {
	case InputValueTeamsDevicesCustomS2sInputRequestOperatorLess, InputValueTeamsDevicesCustomS2sInputRequestOperatorLessOrEquals, InputValueTeamsDevicesCustomS2sInputRequestOperatorGreater, InputValueTeamsDevicesCustomS2sInputRequestOperatorGreaterOrEquals, InputValueTeamsDevicesCustomS2sInputRequestOperatorEquals:
		return true
	}
	return false
}

// Compliance Status
type InputValueComplianceStatus string

const (
	InputValueComplianceStatusCompliant     InputValueComplianceStatus = "compliant"
	InputValueComplianceStatusNoncompliant  InputValueComplianceStatus = "noncompliant"
	InputValueComplianceStatusUnknown       InputValueComplianceStatus = "unknown"
	InputValueComplianceStatusNotapplicable InputValueComplianceStatus = "notapplicable"
	InputValueComplianceStatusIngraceperiod InputValueComplianceStatus = "ingraceperiod"
	InputValueComplianceStatusError         InputValueComplianceStatus = "error"
)

func (r InputValueComplianceStatus) IsKnown() bool {
	switch r {
	case InputValueComplianceStatusCompliant, InputValueComplianceStatusNoncompliant, InputValueComplianceStatusUnknown, InputValueComplianceStatusNotapplicable, InputValueComplianceStatusIngraceperiod, InputValueComplianceStatusError:
		return true
	}
	return false
}

// Count Operator
type InputValueCountOperator string

const (
	InputValueCountOperatorLess            InputValueCountOperator = "<"
	InputValueCountOperatorLessOrEquals    InputValueCountOperator = "<="
	InputValueCountOperatorGreater         InputValueCountOperator = ">"
	InputValueCountOperatorGreaterOrEquals InputValueCountOperator = ">="
	InputValueCountOperatorEquals          InputValueCountOperator = "=="
)

func (r InputValueCountOperator) IsKnown() bool {
	switch r {
	case InputValueCountOperatorLess, InputValueCountOperatorLessOrEquals, InputValueCountOperatorGreater, InputValueCountOperatorGreaterOrEquals, InputValueCountOperatorEquals:
		return true
	}
	return false
}

// Network status of device.
type InputValueNetworkStatus string

const (
	InputValueNetworkStatusConnected     InputValueNetworkStatus = "connected"
	InputValueNetworkStatusDisconnected  InputValueNetworkStatus = "disconnected"
	InputValueNetworkStatusDisconnecting InputValueNetworkStatus = "disconnecting"
	InputValueNetworkStatusConnecting    InputValueNetworkStatus = "connecting"
)

func (r InputValueNetworkStatus) IsKnown() bool {
	switch r {
	case InputValueNetworkStatusConnected, InputValueNetworkStatusDisconnected, InputValueNetworkStatusDisconnecting, InputValueNetworkStatusConnecting:
		return true
	}
	return false
}

// Operating system
type InputValueOperatingSystem string

const (
	InputValueOperatingSystemWindows  InputValueOperatingSystem = "windows"
	InputValueOperatingSystemLinux    InputValueOperatingSystem = "linux"
	InputValueOperatingSystemMac      InputValueOperatingSystem = "mac"
	InputValueOperatingSystemAndroid  InputValueOperatingSystem = "android"
	InputValueOperatingSystemIos      InputValueOperatingSystem = "ios"
	InputValueOperatingSystemChromeos InputValueOperatingSystem = "chromeos"
)

func (r InputValueOperatingSystem) IsKnown() bool {
	switch r {
	case InputValueOperatingSystemWindows, InputValueOperatingSystemLinux, InputValueOperatingSystemMac, InputValueOperatingSystemAndroid, InputValueOperatingSystemIos, InputValueOperatingSystemChromeos:
		return true
	}
	return false
}

// Agent operational state.
type InputValueOperationalState string

const (
	InputValueOperationalStateNa                    InputValueOperationalState = "na"
	InputValueOperationalStatePartiallyDisabled     InputValueOperationalState = "partially_disabled"
	InputValueOperationalStateAutoFullyDisabled     InputValueOperationalState = "auto_fully_disabled"
	InputValueOperationalStateFullyDisabled         InputValueOperationalState = "fully_disabled"
	InputValueOperationalStateAutoPartiallyDisabled InputValueOperationalState = "auto_partially_disabled"
	InputValueOperationalStateDisabledError         InputValueOperationalState = "disabled_error"
	InputValueOperationalStateDBCorruption          InputValueOperationalState = "db_corruption"
)

func (r InputValueOperationalState) IsKnown() bool {
	switch r {
	case InputValueOperationalStateNa, InputValueOperationalStatePartiallyDisabled, InputValueOperationalStateAutoFullyDisabled, InputValueOperationalStateFullyDisabled, InputValueOperationalStateAutoPartiallyDisabled, InputValueOperationalStateDisabledError, InputValueOperationalStateDBCorruption:
		return true
	}
	return false
}

// Operator
type InputValueOperator string

const (
	InputValueOperatorLess            InputValueOperator = "<"
	InputValueOperatorLessOrEquals    InputValueOperator = "<="
	InputValueOperatorGreater         InputValueOperator = ">"
	InputValueOperatorGreaterOrEquals InputValueOperator = ">="
	InputValueOperatorEquals          InputValueOperator = "=="
)

func (r InputValueOperator) IsKnown() bool {
	switch r {
	case InputValueOperatorLess, InputValueOperatorLessOrEquals, InputValueOperatorGreater, InputValueOperatorGreaterOrEquals, InputValueOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type InputValueRiskLevel string

const (
	InputValueRiskLevelLow      InputValueRiskLevel = "low"
	InputValueRiskLevelMedium   InputValueRiskLevel = "medium"
	InputValueRiskLevelHigh     InputValueRiskLevel = "high"
	InputValueRiskLevelCritical InputValueRiskLevel = "critical"
)

func (r InputValueRiskLevel) IsKnown() bool {
	switch r {
	case InputValueRiskLevelLow, InputValueRiskLevelMedium, InputValueRiskLevelHigh, InputValueRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator
type InputValueScoreOperator string

const (
	InputValueScoreOperatorLess            InputValueScoreOperator = "<"
	InputValueScoreOperatorLessOrEquals    InputValueScoreOperator = "<="
	InputValueScoreOperatorGreater         InputValueScoreOperator = ">"
	InputValueScoreOperatorGreaterOrEquals InputValueScoreOperator = ">="
	InputValueScoreOperatorEquals          InputValueScoreOperator = "=="
)

func (r InputValueScoreOperator) IsKnown() bool {
	switch r {
	case InputValueScoreOperatorLess, InputValueScoreOperatorLessOrEquals, InputValueScoreOperatorGreater, InputValueScoreOperatorGreaterOrEquals, InputValueScoreOperatorEquals:
		return true
	}
	return false
}

// For more details on state, please refer to the Crowdstrike documentation.
type InputValueState string

const (
	InputValueStateOnline  InputValueState = "online"
	InputValueStateOffline InputValueState = "offline"
	InputValueStateUnknown InputValueState = "unknown"
)

func (r InputValueState) IsKnown() bool {
	switch r {
	case InputValueStateOnline, InputValueStateOffline, InputValueStateUnknown:
		return true
	}
	return false
}

// Version Operator
type InputValueVersionOperator string

const (
	InputValueVersionOperatorLess            InputValueVersionOperator = "<"
	InputValueVersionOperatorLessOrEquals    InputValueVersionOperator = "<="
	InputValueVersionOperatorGreater         InputValueVersionOperator = ">"
	InputValueVersionOperatorGreaterOrEquals InputValueVersionOperator = ">="
	InputValueVersionOperatorEquals          InputValueVersionOperator = "=="
)

func (r InputValueVersionOperator) IsKnown() bool {
	switch r {
	case InputValueVersionOperatorLess, InputValueVersionOperatorLessOrEquals, InputValueVersionOperatorGreater, InputValueVersionOperatorGreaterOrEquals, InputValueVersionOperatorEquals:
		return true
	}
	return false
}

// The value to be checked against.
type InputValueParam struct {
	// List ID.
	ID param.Field[string] `json:"id"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id"`
	// Confirm the certificate was not imported from another device. We recommend
	// keeping this enabled unless the certificate was deployed without a private key.
	CheckPrivateKey param.Field[bool]        `json:"check_private_key"`
	CheckDisks      param.Field[interface{}] `json:"checkDisks"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn"`
	// Compliance Status
	ComplianceStatus param.Field[InputValueComplianceStatus] `json:"compliance_status"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id"`
	// Count Operator
	CountOperator param.Field[InputValueCountOperator] `json:"countOperator"`
	// Domain
	Domain param.Field[string] `json:"domain"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// Enabled
	Enabled param.Field[bool] `json:"enabled"`
	// Whether or not file exists
	Exists           param.Field[bool]        `json:"exists"`
	ExtendedKeyUsage param.Field[interface{}] `json:"extended_key_usage"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen  param.Field[string]      `json:"last_seen"`
	Locations param.Field[interface{}] `json:"locations"`
	// Network status of device.
	NetworkStatus param.Field[InputValueNetworkStatus] `json:"network_status"`
	// Operating system
	OperatingSystem param.Field[InputValueOperatingSystem] `json:"operating_system"`
	// Agent operational state.
	OperationalState param.Field[InputValueOperationalState] `json:"operational_state"`
	// Operator
	Operator param.Field[InputValueOperator] `json:"operator"`
	// Os Version
	Os param.Field[string] `json:"os"`
	// Operating System Distribution Name (linux only)
	OsDistroName param.Field[string] `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision param.Field[string] `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Version Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OsVersionExtra param.Field[string] `json:"os_version_extra"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// File path.
	Path param.Field[string] `json:"path"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[InputValueRiskLevel] `json:"risk_level"`
	// A value between 0-100 assigned to devices set by the 3rd party posture provider.
	Score param.Field[float64] `json:"score"`
	// Score Operator
	ScoreOperator param.Field[InputValueScoreOperator] `json:"scoreOperator"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State param.Field[InputValueState] `json:"state"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
	// Version of OS
	Version param.Field[string] `json:"version"`
	// Version Operator
	VersionOperator param.Field[InputValueVersionOperator] `json:"versionOperator"`
}

func (r InputValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueParam) implementsInputValueUnionParam() {}

// The value to be checked against.
//
// Satisfied by [InputValueTeamsDevicesFileInputRequestParam],
// [InputValueTeamsDevicesUniqueClientIDInputRequestParam],
// [InputValueTeamsDevicesDomainJoinedInputRequestParam],
// [InputValueTeamsDevicesOsVersionInputRequestParam],
// [InputValueTeamsDevicesFirewallInputRequestParam],
// [InputValueTeamsDevicesSentineloneInputRequestParam],
// [InputValueTeamsDevicesCarbonblackInputRequestParam],
// [InputValueTeamsDevicesDiskEncryptionInputRequestParam],
// [InputValueTeamsDevicesApplicationInputRequestParam],
// [InputValueTeamsDevicesClientCertificateInputRequestParam],
// [InputValueTeamsDevicesClientCertificateV2InputRequestParam],
// [InputValueTeamsDevicesWorkspaceOneInputRequestParam],
// [InputValueTeamsDevicesCrowdstrikeInputRequestParam],
// [InputValueTeamsDevicesIntuneInputRequestParam],
// [InputValueTeamsDevicesKolideInputRequestParam],
// [InputValueTeamsDevicesTaniumInputRequestParam],
// [InputValueTeamsDevicesSentineloneS2sInputRequestParam],
// [InputValueTeamsDevicesCustomS2sInputRequestParam], [InputValueParam].
type InputValueUnionParam interface {
	implementsInputValueUnionParam()
}

type InputValueTeamsDevicesFileInputRequestParam struct {
	// Operating system
	OperatingSystem param.Field[InputValueTeamsDevicesFileInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r InputValueTeamsDevicesFileInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueTeamsDevicesFileInputRequestParam) implementsInputValueUnionParam() {}

type InputValueTeamsDevicesUniqueClientIDInputRequestParam struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System
	OperatingSystem param.Field[InputValueTeamsDevicesUniqueClientIDInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r InputValueTeamsDevicesUniqueClientIDInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueTeamsDevicesUniqueClientIDInputRequestParam) implementsInputValueUnionParam() {}

type InputValueTeamsDevicesDomainJoinedInputRequestParam struct {
	// Operating System
	OperatingSystem param.Field[InputValueTeamsDevicesDomainJoinedInputRequestOperatingSystem] `json:"operating_system,required"`
	// Domain
	Domain param.Field[string] `json:"domain"`
}

func (r InputValueTeamsDevicesDomainJoinedInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueTeamsDevicesDomainJoinedInputRequestParam) implementsInputValueUnionParam() {}

type InputValueTeamsDevicesOsVersionInputRequestParam struct {
	// Operating System
	OperatingSystem param.Field[InputValueTeamsDevicesOsVersionInputRequestOperatingSystem] `json:"operating_system,required"`
	// Operator
	Operator param.Field[InputValueTeamsDevicesOsVersionInputRequestOperator] `json:"operator,required"`
	// Version of OS
	Version param.Field[string] `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OsDistroName param.Field[string] `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OsDistroRevision param.Field[string] `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Version Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OsVersionExtra param.Field[string] `json:"os_version_extra"`
}

func (r InputValueTeamsDevicesOsVersionInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueTeamsDevicesOsVersionInputRequestParam) implementsInputValueUnionParam() {}

type InputValueTeamsDevicesFirewallInputRequestParam struct {
	// Enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System
	OperatingSystem param.Field[InputValueTeamsDevicesFirewallInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r InputValueTeamsDevicesFirewallInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueTeamsDevicesFirewallInputRequestParam) implementsInputValueUnionParam() {}

type InputValueTeamsDevicesSentineloneInputRequestParam struct {
	// Operating system
	OperatingSystem param.Field[InputValueTeamsDevicesSentineloneInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r InputValueTeamsDevicesSentineloneInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueTeamsDevicesSentineloneInputRequestParam) implementsInputValueUnionParam() {}

type InputValueTeamsDevicesCarbonblackInputRequestParam struct {
	// Operating system
	OperatingSystem param.Field[InputValueTeamsDevicesCarbonblackInputRequestOperatingSystem] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r InputValueTeamsDevicesCarbonblackInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueTeamsDevicesCarbonblackInputRequestParam) implementsInputValueUnionParam() {}

type InputValueTeamsDevicesDiskEncryptionInputRequestParam struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r InputValueTeamsDevicesDiskEncryptionInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueTeamsDevicesDiskEncryptionInputRequestParam) implementsInputValueUnionParam() {}

type InputValueTeamsDevicesApplicationInputRequestParam struct {
	// Operating system
	OperatingSystem param.Field[InputValueTeamsDevicesApplicationInputRequestOperatingSystem] `json:"operating_system,required"`
	// Path for the application.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r InputValueTeamsDevicesApplicationInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueTeamsDevicesApplicationInputRequestParam) implementsInputValueUnionParam() {}

type InputValueTeamsDevicesClientCertificateInputRequestParam struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn,required"`
}

func (r InputValueTeamsDevicesClientCertificateInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueTeamsDevicesClientCertificateInputRequestParam) implementsInputValueUnionParam() {}

type InputValueTeamsDevicesClientCertificateV2InputRequestParam struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Confirm the certificate was not imported from another device. We recommend
	// keeping this enabled unless the certificate was deployed without a private key.
	CheckPrivateKey param.Field[bool] `json:"check_private_key,required"`
	// Operating System
	OperatingSystem param.Field[InputValueTeamsDevicesClientCertificateV2InputRequestOperatingSystem] `json:"operating_system,required"`
	// Common Name that is protected by the client certificate. This may include one or
	// more variables in the ${ } notation. Only ${serial_number} and ${hostname} are
	// valid variables.
	Cn param.Field[string] `json:"cn"`
	// List of values indicating purposes for which the certificate public key can be
	// used
	ExtendedKeyUsage param.Field[[]InputValueTeamsDevicesClientCertificateV2InputRequestExtendedKeyUsage] `json:"extended_key_usage"`
	Locations        param.Field[InputValueTeamsDevicesClientCertificateV2InputRequestLocationsParam]     `json:"locations"`
}

func (r InputValueTeamsDevicesClientCertificateV2InputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueTeamsDevicesClientCertificateV2InputRequestParam) implementsInputValueUnionParam() {
}

type InputValueTeamsDevicesClientCertificateV2InputRequestLocationsParam struct {
	// List of paths to check for client certificate on linux.
	Paths param.Field[[]string] `json:"paths"`
	// List of trust stores to check for client certificate.
	TrustStores param.Field[[]InputValueTeamsDevicesClientCertificateV2InputRequestLocationsTrustStore] `json:"trust_stores"`
}

func (r InputValueTeamsDevicesClientCertificateV2InputRequestLocationsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InputValueTeamsDevicesWorkspaceOneInputRequestParam struct {
	// Compliance Status
	ComplianceStatus param.Field[InputValueTeamsDevicesWorkspaceOneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r InputValueTeamsDevicesWorkspaceOneInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueTeamsDevicesWorkspaceOneInputRequestParam) implementsInputValueUnionParam() {}

type InputValueTeamsDevicesCrowdstrikeInputRequestParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen param.Field[string] `json:"last_seen"`
	// Operator
	Operator param.Field[InputValueTeamsDevicesCrowdstrikeInputRequestOperator] `json:"operator"`
	// Os Version
	Os param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State param.Field[InputValueTeamsDevicesCrowdstrikeInputRequestState] `json:"state"`
	// Version
	Version param.Field[string] `json:"version"`
	// Version Operator
	VersionOperator param.Field[InputValueTeamsDevicesCrowdstrikeInputRequestVersionOperator] `json:"versionOperator"`
}

func (r InputValueTeamsDevicesCrowdstrikeInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueTeamsDevicesCrowdstrikeInputRequestParam) implementsInputValueUnionParam() {}

type InputValueTeamsDevicesIntuneInputRequestParam struct {
	// Compliance Status
	ComplianceStatus param.Field[InputValueTeamsDevicesIntuneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r InputValueTeamsDevicesIntuneInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueTeamsDevicesIntuneInputRequestParam) implementsInputValueUnionParam() {}

type InputValueTeamsDevicesKolideInputRequestParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Count Operator
	CountOperator param.Field[InputValueTeamsDevicesKolideInputRequestCountOperator] `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count,required"`
}

func (r InputValueTeamsDevicesKolideInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueTeamsDevicesKolideInputRequestParam) implementsInputValueUnionParam() {}

type InputValueTeamsDevicesTaniumInputRequestParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator param.Field[InputValueTeamsDevicesTaniumInputRequestOperator] `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[InputValueTeamsDevicesTaniumInputRequestRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[InputValueTeamsDevicesTaniumInputRequestScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
}

func (r InputValueTeamsDevicesTaniumInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueTeamsDevicesTaniumInputRequestParam) implementsInputValueUnionParam() {}

type InputValueTeamsDevicesSentineloneS2sInputRequestParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[InputValueTeamsDevicesSentineloneS2sInputRequestNetworkStatus] `json:"network_status"`
	// Agent operational state.
	OperationalState param.Field[InputValueTeamsDevicesSentineloneS2sInputRequestOperationalState] `json:"operational_state"`
	// operator
	Operator param.Field[InputValueTeamsDevicesSentineloneS2sInputRequestOperator] `json:"operator"`
}

func (r InputValueTeamsDevicesSentineloneS2sInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueTeamsDevicesSentineloneS2sInputRequestParam) implementsInputValueUnionParam() {}

type InputValueTeamsDevicesCustomS2sInputRequestParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// operator
	Operator param.Field[InputValueTeamsDevicesCustomS2sInputRequestOperator] `json:"operator,required"`
	// A value between 0-100 assigned to devices set by the 3rd party posture provider.
	Score param.Field[float64] `json:"score,required"`
}

func (r InputValueTeamsDevicesCustomS2sInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InputValueTeamsDevicesCustomS2sInputRequestParam) implementsInputValueUnionParam() {}

type MatchItem struct {
	Platform Platform      `json:"platform"`
	JSON     matchItemJSON `json:"-"`
}

// matchItemJSON contains the JSON metadata for the struct [MatchItem]
type matchItemJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MatchItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matchItemJSON) RawJSON() string {
	return r.raw
}

type MatchItemParam struct {
	Platform param.Field[Platform] `json:"platform"`
}

func (r MatchItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Platform string

const (
	PlatformWindows Platform = "windows"
	PlatformMac     Platform = "mac"
	PlatformLinux   Platform = "linux"
	PlatformAndroid Platform = "android"
	PlatformIos     Platform = "ios"
)

func (r Platform) IsKnown() bool {
	switch r {
	case PlatformWindows, PlatformMac, PlatformLinux, PlatformAndroid, PlatformIos:
		return true
	}
	return false
}

type SingleResponsePosture struct {
	Result DevicePostureRules        `json:"result"`
	JSON   singleResponsePostureJSON `json:"-"`
	APIResponseSingleTeamsDevices
}

// singleResponsePostureJSON contains the JSON metadata for the struct
// [SingleResponsePosture]
type singleResponsePostureJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponsePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponsePostureJSON) RawJSON() string {
	return r.raw
}

type AccountDevicePostureListResponse struct {
	Result []DevicePostureRules                 `json:"result"`
	JSON   accountDevicePostureListResponseJSON `json:"-"`
	APIResponseCollectionTeamsDevices
}

// accountDevicePostureListResponseJSON contains the JSON metadata for the struct
// [AccountDevicePostureListResponse]
type accountDevicePostureListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDevicePostureListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDevicePostureDeleteResponse struct {
	Result AccountDevicePostureDeleteResponseResult `json:"result"`
	JSON   accountDevicePostureDeleteResponseJSON   `json:"-"`
	APIResponseSingleTeamsDevices
}

// accountDevicePostureDeleteResponseJSON contains the JSON metadata for the struct
// [AccountDevicePostureDeleteResponse]
type accountDevicePostureDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDevicePostureDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDevicePostureDeleteResponseResult struct {
	// API UUID.
	ID   string                                       `json:"id"`
	JSON accountDevicePostureDeleteResponseResultJSON `json:"-"`
}

// accountDevicePostureDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountDevicePostureDeleteResponseResult]
type accountDevicePostureDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDevicePostureDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDevicePostureDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDevicePostureNewParams struct {
	// The name of the device posture rule.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture rule.
	Type param.Field[DeviceTypePostureRule] `json:"type,required"`
	// The description of the device posture rule.
	Description param.Field[string] `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration param.Field[string] `json:"expiration"`
	// The value to be checked against.
	Input param.Field[InputValueUnionParam] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]MatchItemParam] `json:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule param.Field[string] `json:"schedule"`
}

func (r AccountDevicePostureNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePostureUpdateParams struct {
	// The name of the device posture rule.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture rule.
	Type param.Field[DeviceTypePostureRule] `json:"type,required"`
	// The description of the device posture rule.
	Description param.Field[string] `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration param.Field[string] `json:"expiration"`
	// The value to be checked against.
	Input param.Field[InputValueUnionParam] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]MatchItemParam] `json:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule param.Field[string] `json:"schedule"`
}

func (r AccountDevicePostureUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePostureDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountDevicePostureDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
