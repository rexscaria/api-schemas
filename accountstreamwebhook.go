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

// AccountStreamWebhookService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountStreamWebhookService] method instead.
type AccountStreamWebhookService struct {
	Options []option.RequestOption
}

// NewAccountStreamWebhookService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamWebhookService(opts ...option.RequestOption) (r *AccountStreamWebhookService) {
	r = &AccountStreamWebhookService{}
	r.Options = opts
	return
}

// Creates a webhook notification.
func (r *AccountStreamWebhookService) New(ctx context.Context, accountID string, body AccountStreamWebhookNewParams, opts ...option.RequestOption) (res *WebhookResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/webhook", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves a list of webhooks.
func (r *AccountStreamWebhookService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *WebhookResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/webhook", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a webhook.
func (r *AccountStreamWebhookService) Delete(ctx context.Context, accountID string, opts ...option.RequestOption) (res *DeletedStreamResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/webhook", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type WebhookResponseSingle struct {
	Errors   []StreamMessages `json:"errors,required"`
	Messages []StreamMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success WebhookResponseSingleSuccess `json:"success,required"`
	Result  interface{}                  `json:"result"`
	JSON    webhookResponseSingleJSON    `json:"-"`
}

// webhookResponseSingleJSON contains the JSON metadata for the struct
// [WebhookResponseSingle]
type webhookResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookResponseSingleJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type WebhookResponseSingleSuccess bool

const (
	WebhookResponseSingleSuccessTrue WebhookResponseSingleSuccess = true
)

func (r WebhookResponseSingleSuccess) IsKnown() bool {
	switch r {
	case WebhookResponseSingleSuccessTrue:
		return true
	}
	return false
}

type AccountStreamWebhookNewParams struct {
	// The URL where webhooks will be sent.
	NotificationURL param.Field[string] `json:"notificationUrl,required" format:"uri"`
}

func (r AccountStreamWebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
