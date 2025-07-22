// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountDlpProfilePredefinedService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDlpProfilePredefinedService] method instead.
type AccountDlpProfilePredefinedService struct {
	Options []option.RequestOption
}

// NewAccountDlpProfilePredefinedService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDlpProfilePredefinedService(opts ...option.RequestOption) (r *AccountDlpProfilePredefinedService) {
	r = &AccountDlpProfilePredefinedService{}
	r.Options = opts
	return
}

// Fetches a predefined DLP profile by id.
func (r *AccountDlpProfilePredefinedService) Get(ctx context.Context, accountID string, profileID string, opts ...option.RequestOption) (res *AccountDlpProfilePredefinedGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", accountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a DLP predefined profile. Only supports enabling/disabling entries.
func (r *AccountDlpProfilePredefinedService) Update(ctx context.Context, accountID string, profileID string, body AccountDlpProfilePredefinedUpdateParams, opts ...option.RequestOption) (res *AccountDlpProfilePredefinedUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", accountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountDlpProfilePredefinedGetResponse struct {
	Result Profile                                    `json:"result"`
	JSON   accountDlpProfilePredefinedGetResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountDlpProfilePredefinedGetResponseJSON contains the JSON metadata for the
// struct [AccountDlpProfilePredefinedGetResponse]
type accountDlpProfilePredefinedGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfilePredefinedGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpProfilePredefinedGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpProfilePredefinedUpdateResponse struct {
	Result Profile                                       `json:"result"`
	JSON   accountDlpProfilePredefinedUpdateResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountDlpProfilePredefinedUpdateResponseJSON contains the JSON metadata for the
// struct [AccountDlpProfilePredefinedUpdateResponse]
type accountDlpProfilePredefinedUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfilePredefinedUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpProfilePredefinedUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpProfilePredefinedUpdateParams struct {
	Entries             param.Field[[]AccountDlpProfilePredefinedUpdateParamsEntry] `json:"entries,required"`
	AIContextEnabled    param.Field[bool]                                           `json:"ai_context_enabled"`
	AllowedMatchCount   param.Field[int64]                                          `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string]                                         `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	OcrEnabled       param.Field[bool]                  `json:"ocr_enabled"`
}

func (r AccountDlpProfilePredefinedUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDlpProfilePredefinedUpdateParamsEntry struct {
	ID      param.Field[string] `json:"id,required" format:"uuid"`
	Enabled param.Field[bool]   `json:"enabled,required"`
}

func (r AccountDlpProfilePredefinedUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
