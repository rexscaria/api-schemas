// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountAlertingV3DestinationPagerdutyConnectService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAlertingV3DestinationPagerdutyConnectService] method instead.
type AccountAlertingV3DestinationPagerdutyConnectService struct {
	Options []option.RequestOption
}

// NewAccountAlertingV3DestinationPagerdutyConnectService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewAccountAlertingV3DestinationPagerdutyConnectService(opts ...option.RequestOption) (r *AccountAlertingV3DestinationPagerdutyConnectService) {
	r = &AccountAlertingV3DestinationPagerdutyConnectService{}
	r.Options = opts
	return
}

// Creates a new token for integrating with PagerDuty.
func (r *AccountAlertingV3DestinationPagerdutyConnectService) NewToken(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAlertingV3DestinationPagerdutyConnectNewTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty/connect", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Links PagerDuty with the account using the integration token.
func (r *AccountAlertingV3DestinationPagerdutyConnectService) Link(ctx context.Context, accountID string, tokenID string, opts ...option.RequestOption) (res *IDResponseAlerting, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty/connect/%s", accountID, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type APIResponseSingleAlerting struct {
	Errors   []AaaMessage `json:"errors,required"`
	Messages []AaaMessage `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseSingleAlertingSuccess `json:"success,required"`
	JSON    apiResponseSingleAlertingJSON    `json:"-"`
}

// apiResponseSingleAlertingJSON contains the JSON metadata for the struct
// [APIResponseSingleAlerting]
type apiResponseSingleAlertingJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleAlerting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleAlertingJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseSingleAlertingSuccess bool

const (
	APIResponseSingleAlertingSuccessTrue APIResponseSingleAlertingSuccess = true
)

func (r APIResponseSingleAlertingSuccess) IsKnown() bool {
	switch r {
	case APIResponseSingleAlertingSuccessTrue:
		return true
	}
	return false
}

type IDResponseAlerting struct {
	Result IDResponseAlertingResult `json:"result"`
	JSON   idResponseAlertingJSON   `json:"-"`
	APIResponseSingleAlerting
}

// idResponseAlertingJSON contains the JSON metadata for the struct
// [IDResponseAlerting]
type idResponseAlertingJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseAlerting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseAlertingJSON) RawJSON() string {
	return r.raw
}

type IDResponseAlertingResult struct {
	// UUID
	ID   string                       `json:"id"`
	JSON idResponseAlertingResultJSON `json:"-"`
}

// idResponseAlertingResultJSON contains the JSON metadata for the struct
// [IDResponseAlertingResult]
type idResponseAlertingResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseAlertingResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseAlertingResultJSON) RawJSON() string {
	return r.raw
}

type AccountAlertingV3DestinationPagerdutyConnectNewTokenResponse struct {
	Result AccountAlertingV3DestinationPagerdutyConnectNewTokenResponseResult `json:"result"`
	JSON   accountAlertingV3DestinationPagerdutyConnectNewTokenResponseJSON   `json:"-"`
	APIResponseSingleAlerting
}

// accountAlertingV3DestinationPagerdutyConnectNewTokenResponseJSON contains the
// JSON metadata for the struct
// [AccountAlertingV3DestinationPagerdutyConnectNewTokenResponse]
type accountAlertingV3DestinationPagerdutyConnectNewTokenResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyConnectNewTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3DestinationPagerdutyConnectNewTokenResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAlertingV3DestinationPagerdutyConnectNewTokenResponseResult struct {
	// token in form of UUID
	ID   string                                                                 `json:"id"`
	JSON accountAlertingV3DestinationPagerdutyConnectNewTokenResponseResultJSON `json:"-"`
}

// accountAlertingV3DestinationPagerdutyConnectNewTokenResponseResultJSON contains
// the JSON metadata for the struct
// [AccountAlertingV3DestinationPagerdutyConnectNewTokenResponseResult]
type accountAlertingV3DestinationPagerdutyConnectNewTokenResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyConnectNewTokenResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3DestinationPagerdutyConnectNewTokenResponseResultJSON) RawJSON() string {
	return r.raw
}
