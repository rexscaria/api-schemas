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

// AccountMnmVpcFlowService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMnmVpcFlowService] method instead.
type AccountMnmVpcFlowService struct {
	Options []option.RequestOption
}

// NewAccountMnmVpcFlowService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMnmVpcFlowService(opts ...option.RequestOption) (r *AccountMnmVpcFlowService) {
	r = &AccountMnmVpcFlowService{}
	r.Options = opts
	return
}

// Generate authentication token for VPC flow logs export.
func (r *AccountMnmVpcFlowService) GenerateToken(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountMnmVpcFlowGenerateTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/vpc-flows/token", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AccountMnmVpcFlowGenerateTokenResponse struct {
	Errors   []MessagesMagicVisibilityMnmItem `json:"errors,required"`
	Messages []MessagesMagicVisibilityMnmItem `json:"messages,required"`
	// Authentication token to be used for VPC Flows export authentication.
	Result string `json:"result,required"`
	// Whether the API call was successful
	Success AccountMnmVpcFlowGenerateTokenResponseSuccess `json:"success,required"`
	JSON    accountMnmVpcFlowGenerateTokenResponseJSON    `json:"-"`
}

// accountMnmVpcFlowGenerateTokenResponseJSON contains the JSON metadata for the
// struct [AccountMnmVpcFlowGenerateTokenResponse]
type accountMnmVpcFlowGenerateTokenResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmVpcFlowGenerateTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMnmVpcFlowGenerateTokenResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountMnmVpcFlowGenerateTokenResponseSuccess bool

const (
	AccountMnmVpcFlowGenerateTokenResponseSuccessTrue AccountMnmVpcFlowGenerateTokenResponseSuccess = true
)

func (r AccountMnmVpcFlowGenerateTokenResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMnmVpcFlowGenerateTokenResponseSuccessTrue:
		return true
	}
	return false
}
