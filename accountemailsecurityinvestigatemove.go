// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountEmailSecurityInvestigateMoveService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountEmailSecurityInvestigateMoveService] method instead.
type AccountEmailSecurityInvestigateMoveService struct {
	Options []option.RequestOption
}

// NewAccountEmailSecurityInvestigateMoveService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountEmailSecurityInvestigateMoveService(opts ...option.RequestOption) (r *AccountEmailSecurityInvestigateMoveService) {
	r = &AccountEmailSecurityInvestigateMoveService{}
	r.Options = opts
	return
}

type RetractionResponseItem struct {
	CompletedTimestamp time.Time                  `json:"completed_timestamp,required" format:"date-time"`
	Destination        string                     `json:"destination,required"`
	ItemCount          int64                      `json:"item_count,required"`
	MessageID          string                     `json:"message_id,required"`
	Operation          string                     `json:"operation,required"`
	Recipient          string                     `json:"recipient,required"`
	Status             string                     `json:"status,required"`
	JSON               retractionResponseItemJSON `json:"-"`
}

// retractionResponseItemJSON contains the JSON metadata for the struct
// [RetractionResponseItem]
type retractionResponseItemJSON struct {
	CompletedTimestamp apijson.Field
	Destination        apijson.Field
	ItemCount          apijson.Field
	MessageID          apijson.Field
	Operation          apijson.Field
	Recipient          apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RetractionResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r retractionResponseItemJSON) RawJSON() string {
	return r.raw
}
