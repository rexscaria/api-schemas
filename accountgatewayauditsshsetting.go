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
	Result SingleResponseAuditResult `json:"result"`
	JSON   singleResponseAuditJSON   `json:"-"`
	APIResponseSingleZeroTrustGateway
}

// singleResponseAuditJSON contains the JSON metadata for the struct
// [SingleResponseAudit]
type singleResponseAuditJSON struct {
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
