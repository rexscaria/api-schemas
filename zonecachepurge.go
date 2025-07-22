// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneCachePurgeService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCachePurgeService] method instead.
type ZoneCachePurgeService struct {
	Options []option.RequestOption
}

// NewZoneCachePurgeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneCachePurgeService(opts ...option.RequestOption) (r *ZoneCachePurgeService) {
	r = &ZoneCachePurgeService{}
	r.Options = opts
	return
}

type MessageItem struct {
	Code    int64           `json:"code,required"`
	Message string          `json:"message,required"`
	JSON    messageItemJSON `json:"-"`
}

// messageItemJSON contains the JSON metadata for the struct [MessageItem]
type messageItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageItemJSON) RawJSON() string {
	return r.raw
}
