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

// AccountBotnetFeedConfigAsnService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountBotnetFeedConfigAsnService] method instead.
type AccountBotnetFeedConfigAsnService struct {
	Options []option.RequestOption
}

// NewAccountBotnetFeedConfigAsnService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountBotnetFeedConfigAsnService(opts ...option.RequestOption) (r *AccountBotnetFeedConfigAsnService) {
	r = &AccountBotnetFeedConfigAsnService{}
	r.Options = opts
	return
}

// Delete an ASN from botnet threat feed for a given user.
func (r *AccountBotnetFeedConfigAsnService) DeleteAsn(ctx context.Context, accountID string, asnID int64, opts ...option.RequestOption) (res *AccountBotnetFeedConfigAsnDeleteAsnResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/botnet_feed/configs/asn/%v", accountID, asnID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Gets a list of all ASNs registered for a user for the DDoS Botnet Feed API.
func (r *AccountBotnetFeedConfigAsnService) ListAsns(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountBotnetFeedConfigAsnListAsnsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/botnet_feed/configs/asn", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountBotnetFeedConfigAsnDeleteAsnResponse struct {
	Result AccountBotnetFeedConfigAsnDeleteAsnResponseResult `json:"result"`
	JSON   accountBotnetFeedConfigAsnDeleteAsnResponseJSON   `json:"-"`
	CommonResponseBotnetAsn
}

// accountBotnetFeedConfigAsnDeleteAsnResponseJSON contains the JSON metadata for
// the struct [AccountBotnetFeedConfigAsnDeleteAsnResponse]
type accountBotnetFeedConfigAsnDeleteAsnResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBotnetFeedConfigAsnDeleteAsnResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBotnetFeedConfigAsnDeleteAsnResponseJSON) RawJSON() string {
	return r.raw
}

type AccountBotnetFeedConfigAsnDeleteAsnResponseResult struct {
	Asn  int64                                                 `json:"asn"`
	JSON accountBotnetFeedConfigAsnDeleteAsnResponseResultJSON `json:"-"`
}

// accountBotnetFeedConfigAsnDeleteAsnResponseResultJSON contains the JSON metadata
// for the struct [AccountBotnetFeedConfigAsnDeleteAsnResponseResult]
type accountBotnetFeedConfigAsnDeleteAsnResponseResultJSON struct {
	Asn         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBotnetFeedConfigAsnDeleteAsnResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBotnetFeedConfigAsnDeleteAsnResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountBotnetFeedConfigAsnListAsnsResponse struct {
	Result AccountBotnetFeedConfigAsnListAsnsResponseResult `json:"result"`
	JSON   accountBotnetFeedConfigAsnListAsnsResponseJSON   `json:"-"`
	CommonResponseBotnetAsn
}

// accountBotnetFeedConfigAsnListAsnsResponseJSON contains the JSON metadata for
// the struct [AccountBotnetFeedConfigAsnListAsnsResponse]
type accountBotnetFeedConfigAsnListAsnsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBotnetFeedConfigAsnListAsnsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBotnetFeedConfigAsnListAsnsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountBotnetFeedConfigAsnListAsnsResponseResult struct {
	Asn  int64                                                `json:"asn"`
	JSON accountBotnetFeedConfigAsnListAsnsResponseResultJSON `json:"-"`
}

// accountBotnetFeedConfigAsnListAsnsResponseResultJSON contains the JSON metadata
// for the struct [AccountBotnetFeedConfigAsnListAsnsResponseResult]
type accountBotnetFeedConfigAsnListAsnsResponseResultJSON struct {
	Asn         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBotnetFeedConfigAsnListAsnsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBotnetFeedConfigAsnListAsnsResponseResultJSON) RawJSON() string {
	return r.raw
}
