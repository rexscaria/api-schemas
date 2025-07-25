// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountAccessUserActiveSessionService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessUserActiveSessionService] method instead.
type AccountAccessUserActiveSessionService struct {
	Options []option.RequestOption
}

// NewAccountAccessUserActiveSessionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessUserActiveSessionService(opts ...option.RequestOption) (r *AccountAccessUserActiveSessionService) {
	r = &AccountAccessUserActiveSessionService{}
	r.Options = opts
	return
}

// Get an active session for a single user.
func (r *AccountAccessUserActiveSessionService) Get(ctx context.Context, accountID string, userID string, nonce string, opts ...option.RequestOption) (res *AccountAccessUserActiveSessionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if nonce == "" {
		err = errors.New("missing required nonce parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/users/%s/active_sessions/%s", accountID, userID, nonce)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get active sessions for a single user.
func (r *AccountAccessUserActiveSessionService) List(ctx context.Context, accountID string, userID string, opts ...option.RequestOption) (res *AccountAccessUserActiveSessionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/users/%s/active_sessions", accountID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Identity struct {
	AccountID          string                           `json:"account_id"`
	AuthStatus         string                           `json:"auth_status"`
	CommonName         string                           `json:"common_name"`
	DeviceID           string                           `json:"device_id"`
	DeviceSessions     map[string]IdentityDeviceSession `json:"device_sessions"`
	DevicePosture      map[string]IdentityDevicePosture `json:"devicePosture"`
	Email              string                           `json:"email"`
	Geo                IdentityGeo                      `json:"geo"`
	Iat                float64                          `json:"iat"`
	Idp                IdentityIdp                      `json:"idp"`
	IP                 string                           `json:"ip"`
	IsGateway          bool                             `json:"is_gateway"`
	IsWarp             bool                             `json:"is_warp"`
	MtlsAuth           IdentityMtlsAuth                 `json:"mtls_auth"`
	ServiceTokenID     string                           `json:"service_token_id"`
	ServiceTokenStatus bool                             `json:"service_token_status"`
	UserUuid           string                           `json:"user_uuid"`
	Version            float64                          `json:"version"`
	JSON               identityJSON                     `json:"-"`
}

// identityJSON contains the JSON metadata for the struct [Identity]
type identityJSON struct {
	AccountID          apijson.Field
	AuthStatus         apijson.Field
	CommonName         apijson.Field
	DeviceID           apijson.Field
	DeviceSessions     apijson.Field
	DevicePosture      apijson.Field
	Email              apijson.Field
	Geo                apijson.Field
	Iat                apijson.Field
	Idp                apijson.Field
	IP                 apijson.Field
	IsGateway          apijson.Field
	IsWarp             apijson.Field
	MtlsAuth           apijson.Field
	ServiceTokenID     apijson.Field
	ServiceTokenStatus apijson.Field
	UserUuid           apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Identity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityJSON) RawJSON() string {
	return r.raw
}

type IdentityDeviceSession struct {
	LastAuthenticated float64                   `json:"last_authenticated"`
	JSON              identityDeviceSessionJSON `json:"-"`
}

// identityDeviceSessionJSON contains the JSON metadata for the struct
// [IdentityDeviceSession]
type identityDeviceSessionJSON struct {
	LastAuthenticated apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *IdentityDeviceSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityDeviceSessionJSON) RawJSON() string {
	return r.raw
}

type IdentityDevicePosture struct {
	ID          string                     `json:"id"`
	Check       IdentityDevicePostureCheck `json:"check"`
	Data        interface{}                `json:"data"`
	Description string                     `json:"description"`
	Error       string                     `json:"error"`
	RuleName    string                     `json:"rule_name"`
	Success     bool                       `json:"success"`
	Timestamp   string                     `json:"timestamp"`
	Type        string                     `json:"type"`
	JSON        identityDevicePostureJSON  `json:"-"`
}

// identityDevicePostureJSON contains the JSON metadata for the struct
// [IdentityDevicePosture]
type identityDevicePostureJSON struct {
	ID          apijson.Field
	Check       apijson.Field
	Data        apijson.Field
	Description apijson.Field
	Error       apijson.Field
	RuleName    apijson.Field
	Success     apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityDevicePostureJSON) RawJSON() string {
	return r.raw
}

type IdentityDevicePostureCheck struct {
	Exists bool                           `json:"exists"`
	Path   string                         `json:"path"`
	JSON   identityDevicePostureCheckJSON `json:"-"`
}

// identityDevicePostureCheckJSON contains the JSON metadata for the struct
// [IdentityDevicePostureCheck]
type identityDevicePostureCheckJSON struct {
	Exists      apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityDevicePostureCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityDevicePostureCheckJSON) RawJSON() string {
	return r.raw
}

type IdentityGeo struct {
	Country string          `json:"country"`
	JSON    identityGeoJSON `json:"-"`
}

// identityGeoJSON contains the JSON metadata for the struct [IdentityGeo]
type identityGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityGeoJSON) RawJSON() string {
	return r.raw
}

type IdentityIdp struct {
	ID   string          `json:"id"`
	Type string          `json:"type"`
	JSON identityIdpJSON `json:"-"`
}

// identityIdpJSON contains the JSON metadata for the struct [IdentityIdp]
type identityIdpJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityIdp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityIdpJSON) RawJSON() string {
	return r.raw
}

type IdentityMtlsAuth struct {
	AuthStatus    string               `json:"auth_status"`
	CertIssuerDn  string               `json:"cert_issuer_dn"`
	CertIssuerSki string               `json:"cert_issuer_ski"`
	CertPresented bool                 `json:"cert_presented"`
	CertSerial    string               `json:"cert_serial"`
	JSON          identityMtlsAuthJSON `json:"-"`
}

// identityMtlsAuthJSON contains the JSON metadata for the struct
// [IdentityMtlsAuth]
type identityMtlsAuthJSON struct {
	AuthStatus    apijson.Field
	CertIssuerDn  apijson.Field
	CertIssuerSki apijson.Field
	CertPresented apijson.Field
	CertSerial    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IdentityMtlsAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityMtlsAuthJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUserActiveSessionGetResponse struct {
	Errors   []MessagesAccessItem `json:"errors,required"`
	Messages []MessagesAccessItem `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountAccessUserActiveSessionGetResponseSuccess `json:"success,required"`
	Result  AccountAccessUserActiveSessionGetResponseResult  `json:"result"`
	JSON    accountAccessUserActiveSessionGetResponseJSON    `json:"-"`
}

// accountAccessUserActiveSessionGetResponseJSON contains the JSON metadata for the
// struct [AccountAccessUserActiveSessionGetResponse]
type accountAccessUserActiveSessionGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserActiveSessionGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountAccessUserActiveSessionGetResponseSuccess bool

const (
	AccountAccessUserActiveSessionGetResponseSuccessTrue AccountAccessUserActiveSessionGetResponseSuccess = true
)

func (r AccountAccessUserActiveSessionGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountAccessUserActiveSessionGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountAccessUserActiveSessionGetResponseResult struct {
	AccountID          string                                                                  `json:"account_id"`
	AuthStatus         string                                                                  `json:"auth_status"`
	CommonName         string                                                                  `json:"common_name"`
	DeviceID           string                                                                  `json:"device_id"`
	DeviceSessions     map[string]AccountAccessUserActiveSessionGetResponseResultDeviceSession `json:"device_sessions"`
	DevicePosture      map[string]AccountAccessUserActiveSessionGetResponseResultDevicePosture `json:"devicePosture"`
	Email              string                                                                  `json:"email"`
	Geo                AccountAccessUserActiveSessionGetResponseResultGeo                      `json:"geo"`
	Iat                float64                                                                 `json:"iat"`
	Idp                AccountAccessUserActiveSessionGetResponseResultIdp                      `json:"idp"`
	IP                 string                                                                  `json:"ip"`
	IsGateway          bool                                                                    `json:"is_gateway"`
	IsWarp             bool                                                                    `json:"is_warp"`
	IsActive           bool                                                                    `json:"isActive"`
	MtlsAuth           AccountAccessUserActiveSessionGetResponseResultMtlsAuth                 `json:"mtls_auth"`
	ServiceTokenID     string                                                                  `json:"service_token_id"`
	ServiceTokenStatus bool                                                                    `json:"service_token_status"`
	UserUuid           string                                                                  `json:"user_uuid"`
	Version            float64                                                                 `json:"version"`
	JSON               accountAccessUserActiveSessionGetResponseResultJSON                     `json:"-"`
}

// accountAccessUserActiveSessionGetResponseResultJSON contains the JSON metadata
// for the struct [AccountAccessUserActiveSessionGetResponseResult]
type accountAccessUserActiveSessionGetResponseResultJSON struct {
	AccountID          apijson.Field
	AuthStatus         apijson.Field
	CommonName         apijson.Field
	DeviceID           apijson.Field
	DeviceSessions     apijson.Field
	DevicePosture      apijson.Field
	Email              apijson.Field
	Geo                apijson.Field
	Iat                apijson.Field
	Idp                apijson.Field
	IP                 apijson.Field
	IsGateway          apijson.Field
	IsWarp             apijson.Field
	IsActive           apijson.Field
	MtlsAuth           apijson.Field
	ServiceTokenID     apijson.Field
	ServiceTokenStatus apijson.Field
	UserUuid           apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserActiveSessionGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUserActiveSessionGetResponseResultDeviceSession struct {
	LastAuthenticated float64                                                          `json:"last_authenticated"`
	JSON              accountAccessUserActiveSessionGetResponseResultDeviceSessionJSON `json:"-"`
}

// accountAccessUserActiveSessionGetResponseResultDeviceSessionJSON contains the
// JSON metadata for the struct
// [AccountAccessUserActiveSessionGetResponseResultDeviceSession]
type accountAccessUserActiveSessionGetResponseResultDeviceSessionJSON struct {
	LastAuthenticated apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionGetResponseResultDeviceSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserActiveSessionGetResponseResultDeviceSessionJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUserActiveSessionGetResponseResultDevicePosture struct {
	ID          string                                                            `json:"id"`
	Check       AccountAccessUserActiveSessionGetResponseResultDevicePostureCheck `json:"check"`
	Data        interface{}                                                       `json:"data"`
	Description string                                                            `json:"description"`
	Error       string                                                            `json:"error"`
	RuleName    string                                                            `json:"rule_name"`
	Success     bool                                                              `json:"success"`
	Timestamp   string                                                            `json:"timestamp"`
	Type        string                                                            `json:"type"`
	JSON        accountAccessUserActiveSessionGetResponseResultDevicePostureJSON  `json:"-"`
}

// accountAccessUserActiveSessionGetResponseResultDevicePostureJSON contains the
// JSON metadata for the struct
// [AccountAccessUserActiveSessionGetResponseResultDevicePosture]
type accountAccessUserActiveSessionGetResponseResultDevicePostureJSON struct {
	ID          apijson.Field
	Check       apijson.Field
	Data        apijson.Field
	Description apijson.Field
	Error       apijson.Field
	RuleName    apijson.Field
	Success     apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionGetResponseResultDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserActiveSessionGetResponseResultDevicePostureJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUserActiveSessionGetResponseResultDevicePostureCheck struct {
	Exists bool                                                                  `json:"exists"`
	Path   string                                                                `json:"path"`
	JSON   accountAccessUserActiveSessionGetResponseResultDevicePostureCheckJSON `json:"-"`
}

// accountAccessUserActiveSessionGetResponseResultDevicePostureCheckJSON contains
// the JSON metadata for the struct
// [AccountAccessUserActiveSessionGetResponseResultDevicePostureCheck]
type accountAccessUserActiveSessionGetResponseResultDevicePostureCheckJSON struct {
	Exists      apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionGetResponseResultDevicePostureCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserActiveSessionGetResponseResultDevicePostureCheckJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUserActiveSessionGetResponseResultGeo struct {
	Country string                                                 `json:"country"`
	JSON    accountAccessUserActiveSessionGetResponseResultGeoJSON `json:"-"`
}

// accountAccessUserActiveSessionGetResponseResultGeoJSON contains the JSON
// metadata for the struct [AccountAccessUserActiveSessionGetResponseResultGeo]
type accountAccessUserActiveSessionGetResponseResultGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionGetResponseResultGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserActiveSessionGetResponseResultGeoJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUserActiveSessionGetResponseResultIdp struct {
	ID   string                                                 `json:"id"`
	Type string                                                 `json:"type"`
	JSON accountAccessUserActiveSessionGetResponseResultIdpJSON `json:"-"`
}

// accountAccessUserActiveSessionGetResponseResultIdpJSON contains the JSON
// metadata for the struct [AccountAccessUserActiveSessionGetResponseResultIdp]
type accountAccessUserActiveSessionGetResponseResultIdpJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionGetResponseResultIdp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserActiveSessionGetResponseResultIdpJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUserActiveSessionGetResponseResultMtlsAuth struct {
	AuthStatus    string                                                      `json:"auth_status"`
	CertIssuerDn  string                                                      `json:"cert_issuer_dn"`
	CertIssuerSki string                                                      `json:"cert_issuer_ski"`
	CertPresented bool                                                        `json:"cert_presented"`
	CertSerial    string                                                      `json:"cert_serial"`
	JSON          accountAccessUserActiveSessionGetResponseResultMtlsAuthJSON `json:"-"`
}

// accountAccessUserActiveSessionGetResponseResultMtlsAuthJSON contains the JSON
// metadata for the struct
// [AccountAccessUserActiveSessionGetResponseResultMtlsAuth]
type accountAccessUserActiveSessionGetResponseResultMtlsAuthJSON struct {
	AuthStatus    apijson.Field
	CertIssuerDn  apijson.Field
	CertIssuerSki apijson.Field
	CertPresented apijson.Field
	CertSerial    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionGetResponseResultMtlsAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserActiveSessionGetResponseResultMtlsAuthJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUserActiveSessionListResponse struct {
	Errors   []MessagesAccessItem `json:"errors,required"`
	Messages []MessagesAccessItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    AccountAccessUserActiveSessionListResponseSuccess    `json:"success,required"`
	Result     []AccountAccessUserActiveSessionListResponseResult   `json:"result"`
	ResultInfo AccountAccessUserActiveSessionListResponseResultInfo `json:"result_info"`
	JSON       accountAccessUserActiveSessionListResponseJSON       `json:"-"`
}

// accountAccessUserActiveSessionListResponseJSON contains the JSON metadata for
// the struct [AccountAccessUserActiveSessionListResponse]
type accountAccessUserActiveSessionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserActiveSessionListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountAccessUserActiveSessionListResponseSuccess bool

const (
	AccountAccessUserActiveSessionListResponseSuccessTrue AccountAccessUserActiveSessionListResponseSuccess = true
)

func (r AccountAccessUserActiveSessionListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountAccessUserActiveSessionListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountAccessUserActiveSessionListResponseResult struct {
	Expiration int64                                                    `json:"expiration"`
	Metadata   AccountAccessUserActiveSessionListResponseResultMetadata `json:"metadata"`
	Name       string                                                   `json:"name"`
	JSON       accountAccessUserActiveSessionListResponseResultJSON     `json:"-"`
}

// accountAccessUserActiveSessionListResponseResultJSON contains the JSON metadata
// for the struct [AccountAccessUserActiveSessionListResponseResult]
type accountAccessUserActiveSessionListResponseResultJSON struct {
	Expiration  apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserActiveSessionListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUserActiveSessionListResponseResultMetadata struct {
	Apps    map[string]AccountAccessUserActiveSessionListResponseResultMetadataApp `json:"apps"`
	Expires int64                                                                  `json:"expires"`
	Iat     int64                                                                  `json:"iat"`
	Nonce   string                                                                 `json:"nonce"`
	Ttl     int64                                                                  `json:"ttl"`
	JSON    accountAccessUserActiveSessionListResponseResultMetadataJSON           `json:"-"`
}

// accountAccessUserActiveSessionListResponseResultMetadataJSON contains the JSON
// metadata for the struct
// [AccountAccessUserActiveSessionListResponseResultMetadata]
type accountAccessUserActiveSessionListResponseResultMetadataJSON struct {
	Apps        apijson.Field
	Expires     apijson.Field
	Iat         apijson.Field
	Nonce       apijson.Field
	Ttl         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionListResponseResultMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserActiveSessionListResponseResultMetadataJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUserActiveSessionListResponseResultMetadataApp struct {
	Hostname string                                                          `json:"hostname"`
	Name     string                                                          `json:"name"`
	Type     string                                                          `json:"type"`
	Uid      string                                                          `json:"uid"`
	JSON     accountAccessUserActiveSessionListResponseResultMetadataAppJSON `json:"-"`
}

// accountAccessUserActiveSessionListResponseResultMetadataAppJSON contains the
// JSON metadata for the struct
// [AccountAccessUserActiveSessionListResponseResultMetadataApp]
type accountAccessUserActiveSessionListResponseResultMetadataAppJSON struct {
	Hostname    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionListResponseResultMetadataApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserActiveSessionListResponseResultMetadataAppJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUserActiveSessionListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                                  `json:"total_count"`
	JSON       accountAccessUserActiveSessionListResponseResultInfoJSON `json:"-"`
}

// accountAccessUserActiveSessionListResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountAccessUserActiveSessionListResponseResultInfo]
type accountAccessUserActiveSessionListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserActiveSessionListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserActiveSessionListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}
