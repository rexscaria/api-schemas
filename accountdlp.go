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

// AccountDlpService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDlpService] method instead.
type AccountDlpService struct {
	Options    []option.RequestOption
	Datasets   *AccountDlpDatasetService
	Email      *AccountDlpEmailService
	Entries    *AccountDlpEntryService
	Patterns   *AccountDlpPatternService
	PayloadLog *AccountDlpPayloadLogService
	Profiles   *AccountDlpProfileService
}

// NewAccountDlpService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountDlpService(opts ...option.RequestOption) (r *AccountDlpService) {
	r = &AccountDlpService{}
	r.Options = opts
	r.Datasets = NewAccountDlpDatasetService(opts...)
	r.Email = NewAccountDlpEmailService(opts...)
	r.Entries = NewAccountDlpEntryService(opts...)
	r.Patterns = NewAccountDlpPatternService(opts...)
	r.PayloadLog = NewAccountDlpPayloadLogService(opts...)
	r.Profiles = NewAccountDlpProfileService(opts...)
	return
}

// Fetch limits associated with DLP for account
func (r *AccountDlpService) GetLimits(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountDlpGetLimitsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/limits", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountDlpGetLimitsResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDlpGetLimitsResponseSuccess `json:"success,required"`
	Result  AccountDlpGetLimitsResponseResult  `json:"result"`
	JSON    accountDlpGetLimitsResponseJSON    `json:"-"`
}

// accountDlpGetLimitsResponseJSON contains the JSON metadata for the struct
// [AccountDlpGetLimitsResponse]
type accountDlpGetLimitsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpGetLimitsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpGetLimitsResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDlpGetLimitsResponseSuccess bool

const (
	AccountDlpGetLimitsResponseSuccessTrue AccountDlpGetLimitsResponseSuccess = true
)

func (r AccountDlpGetLimitsResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDlpGetLimitsResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDlpGetLimitsResponseResult struct {
	MaxDatasetCells int64                                 `json:"max_dataset_cells,required"`
	JSON            accountDlpGetLimitsResponseResultJSON `json:"-"`
}

// accountDlpGetLimitsResponseResultJSON contains the JSON metadata for the struct
// [AccountDlpGetLimitsResponseResult]
type accountDlpGetLimitsResponseResultJSON struct {
	MaxDatasetCells apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDlpGetLimitsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpGetLimitsResponseResultJSON) RawJSON() string {
	return r.raw
}
