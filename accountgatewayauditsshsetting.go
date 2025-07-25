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

// AccountGatewayAuditSSHSettingService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountGatewayAuditSSHSettingService] method instead.
type AccountGatewayAuditSSHSettingService struct {
	Options []option.RequestOption
}

// NewAccountGatewayAuditSSHSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountGatewayAuditSSHSettingService(opts ...option.RequestOption) (r *AccountGatewayAuditSSHSettingService) {
	r = &AccountGatewayAuditSSHSettingService{}
	r.Options = opts
	return
}

// Gets all Zero Trust Audit SSH and SSH with Access for Infrastructure settings
// for an account.
func (r *AccountGatewayAuditSSHSettingService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *SingleResponseAudit, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/audit_ssh_settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates Zero Trust Audit SSH and SSH with Access for Infrastructure settings for
// an account.
func (r *AccountGatewayAuditSSHSettingService) Update(ctx context.Context, accountID string, body AccountGatewayAuditSSHSettingUpdateParams, opts ...option.RequestOption) (res *SingleResponseAudit, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/audit_ssh_settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Rotates the SSH account seed that is used for generating the host key identity
// when connecting through the Cloudflare SSH Proxy.
func (r *AccountGatewayAuditSSHSettingService) RotateSeed(ctx context.Context, accountID string, opts ...option.RequestOption) (res *SingleResponseAudit, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/audit_ssh_settings/rotate_seed", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type SingleResponseAudit struct {
	Errors   []SingleResponseAuditError   `json:"errors,required"`
	Messages []SingleResponseAuditMessage `json:"messages,required"`
	// Whether the API call was successful
	Success SingleResponseAuditSuccess `json:"success,required"`
	Result  SingleResponseAuditResult  `json:"result"`
	JSON    singleResponseAuditJSON    `json:"-"`
}

// singleResponseAuditJSON contains the JSON metadata for the struct
// [SingleResponseAudit]
type singleResponseAuditJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseAudit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseAuditJSON) RawJSON() string {
	return r.raw
}

type SingleResponseAuditError struct {
	Code             int64                           `json:"code,required"`
	Message          string                          `json:"message,required"`
	DocumentationURL string                          `json:"documentation_url"`
	Source           SingleResponseAuditErrorsSource `json:"source"`
	JSON             singleResponseAuditErrorJSON    `json:"-"`
}

// singleResponseAuditErrorJSON contains the JSON metadata for the struct
// [SingleResponseAuditError]
type singleResponseAuditErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SingleResponseAuditError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseAuditErrorJSON) RawJSON() string {
	return r.raw
}

type SingleResponseAuditErrorsSource struct {
	Pointer string                              `json:"pointer"`
	JSON    singleResponseAuditErrorsSourceJSON `json:"-"`
}

// singleResponseAuditErrorsSourceJSON contains the JSON metadata for the struct
// [SingleResponseAuditErrorsSource]
type singleResponseAuditErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseAuditErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseAuditErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SingleResponseAuditMessage struct {
	Code             int64                             `json:"code,required"`
	Message          string                            `json:"message,required"`
	DocumentationURL string                            `json:"documentation_url"`
	Source           SingleResponseAuditMessagesSource `json:"source"`
	JSON             singleResponseAuditMessageJSON    `json:"-"`
}

// singleResponseAuditMessageJSON contains the JSON metadata for the struct
// [SingleResponseAuditMessage]
type singleResponseAuditMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SingleResponseAuditMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseAuditMessageJSON) RawJSON() string {
	return r.raw
}

type SingleResponseAuditMessagesSource struct {
	Pointer string                                `json:"pointer"`
	JSON    singleResponseAuditMessagesSourceJSON `json:"-"`
}

// singleResponseAuditMessagesSourceJSON contains the JSON metadata for the struct
// [SingleResponseAuditMessagesSource]
type singleResponseAuditMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseAuditMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseAuditMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SingleResponseAuditSuccess bool

const (
	SingleResponseAuditSuccessTrue SingleResponseAuditSuccess = true
)

func (r SingleResponseAuditSuccess) IsKnown() bool {
	switch r {
	case SingleResponseAuditSuccessTrue:
		return true
	}
	return false
}

type SingleResponseAuditResult struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Base64 encoded HPKE public key used to encrypt all your ssh session logs.
	// https://developers.cloudflare.com/cloudflare-one/connections/connect-networks/use-cases/ssh/ssh-infrastructure-access/#enable-ssh-command-logging
	PublicKey string `json:"public_key"`
	// Seed ID
	SeedID    string                        `json:"seed_id"`
	UpdatedAt time.Time                     `json:"updated_at" format:"date-time"`
	JSON      singleResponseAuditResultJSON `json:"-"`
}

// singleResponseAuditResultJSON contains the JSON metadata for the struct
// [SingleResponseAuditResult]
type singleResponseAuditResultJSON struct {
	CreatedAt   apijson.Field
	PublicKey   apijson.Field
	SeedID      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseAuditResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseAuditResultJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayAuditSSHSettingUpdateParams struct {
	// Base64 encoded HPKE public key used to encrypt all your ssh session logs.
	// https://developers.cloudflare.com/cloudflare-one/connections/connect-networks/use-cases/ssh/ssh-infrastructure-access/#enable-ssh-command-logging
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r AccountGatewayAuditSSHSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
