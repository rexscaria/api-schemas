// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/option"
)

// AccountDeviceDexTestService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDeviceDexTestService] method instead.
type AccountDeviceDexTestService struct {
	Options []option.RequestOption
}

// NewAccountDeviceDexTestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDeviceDexTestService(opts ...option.RequestOption) (r *AccountDeviceDexTestService) {
	r = &AccountDeviceDexTestService{}
	r.Options = opts
	return
}

type MessagesDeviceTestsItems struct {
	Code             int64                          `json:"code,required"`
	Message          string                         `json:"message,required"`
	DocumentationURL string                         `json:"documentation_url"`
	Source           MessagesDeviceTestsItemsSource `json:"source"`
	JSON             messagesDeviceTestsItemsJSON   `json:"-"`
}

// messagesDeviceTestsItemsJSON contains the JSON metadata for the struct
// [MessagesDeviceTestsItems]
type messagesDeviceTestsItemsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MessagesDeviceTestsItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesDeviceTestsItemsJSON) RawJSON() string {
	return r.raw
}

type MessagesDeviceTestsItemsSource struct {
	Pointer string                             `json:"pointer"`
	JSON    messagesDeviceTestsItemsSourceJSON `json:"-"`
}

// messagesDeviceTestsItemsSourceJSON contains the JSON metadata for the struct
// [MessagesDeviceTestsItemsSource]
type messagesDeviceTestsItemsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesDeviceTestsItemsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesDeviceTestsItemsSourceJSON) RawJSON() string {
	return r.raw
}
