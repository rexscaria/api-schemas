// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountAlertingV3DestinationWebhookService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAlertingV3DestinationWebhookService] method instead.
type AccountAlertingV3DestinationWebhookService struct {
	Options []option.RequestOption
}

// NewAccountAlertingV3DestinationWebhookService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountAlertingV3DestinationWebhookService(opts ...option.RequestOption) (r *AccountAlertingV3DestinationWebhookService) {
	r = &AccountAlertingV3DestinationWebhookService{}
	r.Options = opts
	return
}

// Creates a new webhook destination.
func (r *AccountAlertingV3DestinationWebhookService) New(ctx context.Context, accountID string, body AccountAlertingV3DestinationWebhookNewParams, opts ...option.RequestOption) (res *IDResponseAlerting, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details for a single webhooks destination.
func (r *AccountAlertingV3DestinationWebhookService) Get(ctx context.Context, accountID string, webhookID string, opts ...option.RequestOption) (res *AccountAlertingV3DestinationWebhookGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", accountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a webhook destination.
func (r *AccountAlertingV3DestinationWebhookService) Update(ctx context.Context, accountID string, webhookID string, body AccountAlertingV3DestinationWebhookUpdateParams, opts ...option.RequestOption) (res *IDResponseAlerting, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", accountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Gets a list of all configured webhook destinations.
func (r *AccountAlertingV3DestinationWebhookService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAlertingV3DestinationWebhookListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a configured webhook destination.
func (r *AccountAlertingV3DestinationWebhookService) Delete(ctx context.Context, accountID string, webhookID string, opts ...option.RequestOption) (res *APIResponseCollectionAlerting, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", accountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Webhooks struct {
	// The unique identifier of a webhook
	ID string `json:"id"`
	// Timestamp of when the webhook destination was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of the last time an attempt to dispatch a notification to this webhook
	// failed.
	LastFailure time.Time `json:"last_failure" format:"date-time"`
	// Timestamp of the last time Cloudflare was able to successfully dispatch a
	// notification using this webhook.
	LastSuccess time.Time `json:"last_success" format:"date-time"`
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name string `json:"name"`
	// Type of webhook endpoint.
	Type WebhooksType `json:"type"`
	// The POST endpoint to call when dispatching a notification.
	URL  string       `json:"url"`
	JSON webhooksJSON `json:"-"`
}

// webhooksJSON contains the JSON metadata for the struct [Webhooks]
type webhooksJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	LastFailure apijson.Field
	LastSuccess apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Webhooks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhooksJSON) RawJSON() string {
	return r.raw
}

// Type of webhook endpoint.
type WebhooksType string

const (
	WebhooksTypeSlack   WebhooksType = "slack"
	WebhooksTypeGeneric WebhooksType = "generic"
	WebhooksTypeGchat   WebhooksType = "gchat"
)

func (r WebhooksType) IsKnown() bool {
	switch r {
	case WebhooksTypeSlack, WebhooksTypeGeneric, WebhooksTypeGchat:
		return true
	}
	return false
}

type AccountAlertingV3DestinationWebhookGetResponse struct {
	Result Webhooks                                           `json:"result"`
	JSON   accountAlertingV3DestinationWebhookGetResponseJSON `json:"-"`
	APIResponseSingleAlerting
}

// accountAlertingV3DestinationWebhookGetResponseJSON contains the JSON metadata
// for the struct [AccountAlertingV3DestinationWebhookGetResponse]
type accountAlertingV3DestinationWebhookGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3DestinationWebhookGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAlertingV3DestinationWebhookListResponse struct {
	Result []Webhooks                                          `json:"result"`
	JSON   accountAlertingV3DestinationWebhookListResponseJSON `json:"-"`
	APIResponseCollectionAlerting
}

// accountAlertingV3DestinationWebhookListResponseJSON contains the JSON metadata
// for the struct [AccountAlertingV3DestinationWebhookListResponse]
type accountAlertingV3DestinationWebhookListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3DestinationWebhookListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAlertingV3DestinationWebhookNewParams struct {
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name param.Field[string] `json:"name,required"`
	// The POST endpoint to call when dispatching a notification.
	URL param.Field[string] `json:"url,required"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching generic webhook notifications or formatted for supported
	// destinations. Secrets are not returned in any API response body.
	Secret param.Field[string] `json:"secret"`
}

func (r AccountAlertingV3DestinationWebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAlertingV3DestinationWebhookUpdateParams struct {
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name param.Field[string] `json:"name,required"`
	// The POST endpoint to call when dispatching a notification.
	URL param.Field[string] `json:"url,required"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching generic webhook notifications or formatted for supported
	// destinations. Secrets are not returned in any API response body.
	Secret param.Field[string] `json:"secret"`
}

func (r AccountAlertingV3DestinationWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
