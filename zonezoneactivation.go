// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneZoneActivationService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneZoneActivationService] method instead.
type ZoneZoneActivationService struct {
	Options []option.RequestOption
}

// NewZoneZoneActivationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneZoneActivationService(opts ...option.RequestOption) (r *ZoneZoneActivationService) {
	r = &ZoneZoneActivationService{}
	r.Options = opts
	return
}

type MessageItem struct {
	Code             int64             `json:"code,required"`
	Message          string            `json:"message,required"`
	DocumentationURL string            `json:"documentation_url"`
	Source           MessageItemSource `json:"source"`
	JSON             messageItemJSON   `json:"-"`
}

// messageItemJSON contains the JSON metadata for the struct [MessageItem]
type messageItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MessageItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageItemJSON) RawJSON() string {
	return r.raw
}

type MessageItemSource struct {
	Pointer string                `json:"pointer"`
	JSON    messageItemSourceJSON `json:"-"`
}

// messageItemSourceJSON contains the JSON metadata for the struct
// [MessageItemSource]
type messageItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageItemSourceJSON) RawJSON() string {
	return r.raw
}
