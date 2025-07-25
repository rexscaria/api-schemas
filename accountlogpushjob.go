// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/option"
)

// AccountLogpushJobService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountLogpushJobService] method instead.
type AccountLogpushJobService struct {
	Options []option.RequestOption
}

// NewAccountLogpushJobService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountLogpushJobService(opts ...option.RequestOption) (r *AccountLogpushJobService) {
	r = &AccountLogpushJobService{}
	r.Options = opts
	return
}

type MessagesLogpushItem struct {
	Code             int64                     `json:"code,required"`
	Message          string                    `json:"message,required"`
	DocumentationURL string                    `json:"documentation_url"`
	Source           MessagesLogpushItemSource `json:"source"`
	JSON             messagesLogpushItemJSON   `json:"-"`
}

// messagesLogpushItemJSON contains the JSON metadata for the struct
// [MessagesLogpushItem]
type messagesLogpushItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MessagesLogpushItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesLogpushItemJSON) RawJSON() string {
	return r.raw
}

type MessagesLogpushItemSource struct {
	Pointer string                        `json:"pointer"`
	JSON    messagesLogpushItemSourceJSON `json:"-"`
}

// messagesLogpushItemSourceJSON contains the JSON metadata for the struct
// [MessagesLogpushItemSource]
type messagesLogpushItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesLogpushItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesLogpushItemSourceJSON) RawJSON() string {
	return r.raw
}
