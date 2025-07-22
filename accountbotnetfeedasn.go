// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountBotnetFeedAsnService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountBotnetFeedAsnService] method instead.
type AccountBotnetFeedAsnService struct {
	Options []option.RequestOption
}

// NewAccountBotnetFeedAsnService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountBotnetFeedAsnService(opts ...option.RequestOption) (r *AccountBotnetFeedAsnService) {
	r = &AccountBotnetFeedAsnService{}
	r.Options = opts
	return
}

// Gets all the data the botnet tracking database has for a given ASN registered to
// user account for given date. If no date is given, it will return results for the
// previous day.
func (r *AccountBotnetFeedAsnService) GetDailyReport(ctx context.Context, accountID string, asnID int64, query AccountBotnetFeedAsnGetDailyReportParams, opts ...option.RequestOption) (res *AccountBotnetFeedAsnGetDailyReportResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/botnet_feed/asn/%v/day_report", accountID, asnID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Gets all the data the botnet threat feed tracking database has for a given ASN
// registered to user account.
func (r *AccountBotnetFeedAsnService) GetFullReport(ctx context.Context, accountID string, asnID int64, opts ...option.RequestOption) (res *AccountBotnetFeedAsnGetFullReportResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/botnet_feed/asn/%v/full_report", accountID, asnID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CommonResponseBotnetAsn struct {
	Errors   []DosMessages `json:"errors,required"`
	Messages []DosMessages `json:"messages,required"`
	// Whether the API call was successful
	Success CommonResponseBotnetAsnSuccess `json:"success,required"`
	JSON    commonResponseBotnetAsnJSON    `json:"-"`
}

// commonResponseBotnetAsnJSON contains the JSON metadata for the struct
// [CommonResponseBotnetAsn]
type commonResponseBotnetAsnJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommonResponseBotnetAsn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commonResponseBotnetAsnJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CommonResponseBotnetAsnSuccess bool

const (
	CommonResponseBotnetAsnSuccessTrue CommonResponseBotnetAsnSuccess = true
)

func (r CommonResponseBotnetAsnSuccess) IsKnown() bool {
	switch r {
	case CommonResponseBotnetAsnSuccessTrue:
		return true
	}
	return false
}

type DosMessages struct {
	Code    int64           `json:"code,required"`
	Message string          `json:"message,required"`
	JSON    dosMessagesJSON `json:"-"`
}

// dosMessagesJSON contains the JSON metadata for the struct [DosMessages]
type dosMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DosMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dosMessagesJSON) RawJSON() string {
	return r.raw
}

type AccountBotnetFeedAsnGetDailyReportResponse struct {
	Result AccountBotnetFeedAsnGetDailyReportResponseResult `json:"result"`
	JSON   accountBotnetFeedAsnGetDailyReportResponseJSON   `json:"-"`
	CommonResponseBotnetAsn
}

// accountBotnetFeedAsnGetDailyReportResponseJSON contains the JSON metadata for
// the struct [AccountBotnetFeedAsnGetDailyReportResponse]
type accountBotnetFeedAsnGetDailyReportResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBotnetFeedAsnGetDailyReportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBotnetFeedAsnGetDailyReportResponseJSON) RawJSON() string {
	return r.raw
}

type AccountBotnetFeedAsnGetDailyReportResponseResult struct {
	Cidr         string                                               `json:"cidr"`
	Date         time.Time                                            `json:"date" format:"date-time"`
	OffenseCount int64                                                `json:"offense_count"`
	JSON         accountBotnetFeedAsnGetDailyReportResponseResultJSON `json:"-"`
}

// accountBotnetFeedAsnGetDailyReportResponseResultJSON contains the JSON metadata
// for the struct [AccountBotnetFeedAsnGetDailyReportResponseResult]
type accountBotnetFeedAsnGetDailyReportResponseResultJSON struct {
	Cidr         apijson.Field
	Date         apijson.Field
	OffenseCount apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountBotnetFeedAsnGetDailyReportResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBotnetFeedAsnGetDailyReportResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountBotnetFeedAsnGetFullReportResponse struct {
	Result AccountBotnetFeedAsnGetFullReportResponseResult `json:"result"`
	JSON   accountBotnetFeedAsnGetFullReportResponseJSON   `json:"-"`
	CommonResponseBotnetAsn
}

// accountBotnetFeedAsnGetFullReportResponseJSON contains the JSON metadata for the
// struct [AccountBotnetFeedAsnGetFullReportResponse]
type accountBotnetFeedAsnGetFullReportResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBotnetFeedAsnGetFullReportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBotnetFeedAsnGetFullReportResponseJSON) RawJSON() string {
	return r.raw
}

type AccountBotnetFeedAsnGetFullReportResponseResult struct {
	Cidr         string                                              `json:"cidr"`
	Date         time.Time                                           `json:"date" format:"date-time"`
	OffenseCount int64                                               `json:"offense_count"`
	JSON         accountBotnetFeedAsnGetFullReportResponseResultJSON `json:"-"`
}

// accountBotnetFeedAsnGetFullReportResponseResultJSON contains the JSON metadata
// for the struct [AccountBotnetFeedAsnGetFullReportResponseResult]
type accountBotnetFeedAsnGetFullReportResponseResultJSON struct {
	Cidr         apijson.Field
	Date         apijson.Field
	OffenseCount apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountBotnetFeedAsnGetFullReportResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBotnetFeedAsnGetFullReportResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountBotnetFeedAsnGetDailyReportParams struct {
	Date param.Field[time.Time] `query:"date" format:"date-time"`
}

// URLQuery serializes [AccountBotnetFeedAsnGetDailyReportParams]'s query
// parameters as `url.Values`.
func (r AccountBotnetFeedAsnGetDailyReportParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
