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

type IDResponseAlerting struct {
	Errors   []IDResponseAlertingError   `json:"errors,required"`
	Messages []IDResponseAlertingMessage `json:"messages,required"`
	// Whether the API call was successful
	Success IDResponseAlertingSuccess `json:"success,required"`
	Result  IDResponseAlertingResult  `json:"result"`
	JSON    idResponseAlertingJSON    `json:"-"`
}

// idResponseAlertingJSON contains the JSON metadata for the struct
// [IDResponseAlerting]
type idResponseAlertingJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

type IDResponseAlertingError struct {
	Code             int64                          `json:"code,required"`
	Message          string                         `json:"message,required"`
	DocumentationURL string                         `json:"documentation_url"`
	Source           IDResponseAlertingErrorsSource `json:"source"`
	JSON             idResponseAlertingErrorJSON    `json:"-"`
}

// idResponseAlertingErrorJSON contains the JSON metadata for the struct
// [IDResponseAlertingError]
type idResponseAlertingErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IDResponseAlertingError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseAlertingErrorJSON) RawJSON() string {
	return r.raw
}

type IDResponseAlertingErrorsSource struct {
	Pointer string                             `json:"pointer"`
	JSON    idResponseAlertingErrorsSourceJSON `json:"-"`
}

// idResponseAlertingErrorsSourceJSON contains the JSON metadata for the struct
// [IDResponseAlertingErrorsSource]
type idResponseAlertingErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseAlertingErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseAlertingErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IDResponseAlertingMessage struct {
	Code             int64                            `json:"code,required"`
	Message          string                           `json:"message,required"`
	DocumentationURL string                           `json:"documentation_url"`
	Source           IDResponseAlertingMessagesSource `json:"source"`
	JSON             idResponseAlertingMessageJSON    `json:"-"`
}

// idResponseAlertingMessageJSON contains the JSON metadata for the struct
// [IDResponseAlertingMessage]
type idResponseAlertingMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IDResponseAlertingMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseAlertingMessageJSON) RawJSON() string {
	return r.raw
}

type IDResponseAlertingMessagesSource struct {
	Pointer string                               `json:"pointer"`
	JSON    idResponseAlertingMessagesSourceJSON `json:"-"`
}

// idResponseAlertingMessagesSourceJSON contains the JSON metadata for the struct
// [IDResponseAlertingMessagesSource]
type idResponseAlertingMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseAlertingMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseAlertingMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IDResponseAlertingSuccess bool

const (
	IDResponseAlertingSuccessTrue IDResponseAlertingSuccess = true
)

func (r IDResponseAlertingSuccess) IsKnown() bool {
	switch r {
	case IDResponseAlertingSuccessTrue:
		return true
	}
	return false
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
	Errors   []AaaMessage `json:"errors,required"`
	Messages []AaaMessage `json:"messages,required"`
	// Whether the API call was successful
	Success AccountAlertingV3DestinationPagerdutyConnectNewTokenResponseSuccess `json:"success,required"`
	Result  AccountAlertingV3DestinationPagerdutyConnectNewTokenResponseResult  `json:"result"`
	JSON    accountAlertingV3DestinationPagerdutyConnectNewTokenResponseJSON    `json:"-"`
}

// accountAlertingV3DestinationPagerdutyConnectNewTokenResponseJSON contains the
// JSON metadata for the struct
// [AccountAlertingV3DestinationPagerdutyConnectNewTokenResponse]
type accountAlertingV3DestinationPagerdutyConnectNewTokenResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful
type AccountAlertingV3DestinationPagerdutyConnectNewTokenResponseSuccess bool

const (
	AccountAlertingV3DestinationPagerdutyConnectNewTokenResponseSuccessTrue AccountAlertingV3DestinationPagerdutyConnectNewTokenResponseSuccess = true
)

func (r AccountAlertingV3DestinationPagerdutyConnectNewTokenResponseSuccess) IsKnown() bool {
	switch r {
	case AccountAlertingV3DestinationPagerdutyConnectNewTokenResponseSuccessTrue:
		return true
	}
	return false
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
