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

// ZoneSubscriptionService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSubscriptionService] method instead.
type ZoneSubscriptionService struct {
	Options []option.RequestOption
}

// NewZoneSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSubscriptionService(opts ...option.RequestOption) (r *ZoneSubscriptionService) {
	r = &ZoneSubscriptionService{}
	r.Options = opts
	return
}

// Create a zone subscription, either plan or add-ons.
func (r *ZoneSubscriptionService) New(ctx context.Context, identifier string, body ZoneSubscriptionNewParams, opts ...option.RequestOption) (res *ZoneSubscription, err error) {
	opts = append(r.Options[:], opts...)
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists zone subscription details.
func (r *ZoneSubscriptionService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *ZoneSubscription, err error) {
	opts = append(r.Options[:], opts...)
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates zone subscriptions, either plan or add-ons.
func (r *ZoneSubscriptionService) Update(ctx context.Context, identifier string, body ZoneSubscriptionUpdateParams, opts ...option.RequestOption) (res *ZoneSubscription, err error) {
	opts = append(r.Options[:], opts...)
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneSubscription struct {
	Result interface{}          `json:"result"`
	JSON   zoneSubscriptionJSON `json:"-"`
	APIResponseSingleBilling
}

// zoneSubscriptionJSON contains the JSON metadata for the struct
// [ZoneSubscription]
type zoneSubscriptionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSubscriptionJSON) RawJSON() string {
	return r.raw
}

type ZoneSubscriptionNewParams struct {
	SubscriptionV2 SubscriptionV2Param `json:"subscription_v2,required"`
}

func (r ZoneSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SubscriptionV2)
}

type ZoneSubscriptionUpdateParams struct {
	SubscriptionV2 SubscriptionV2Param `json:"subscription_v2,required"`
}

func (r ZoneSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SubscriptionV2)
}
