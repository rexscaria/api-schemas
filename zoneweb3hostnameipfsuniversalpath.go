// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneWeb3HostnameIpfsUniversalPathService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneWeb3HostnameIpfsUniversalPathService] method instead.
type ZoneWeb3HostnameIpfsUniversalPathService struct {
	Options     []option.RequestOption
	ContentList *ZoneWeb3HostnameIpfsUniversalPathContentListService
}

// NewZoneWeb3HostnameIpfsUniversalPathService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneWeb3HostnameIpfsUniversalPathService(opts ...option.RequestOption) (r *ZoneWeb3HostnameIpfsUniversalPathService) {
	r = &ZoneWeb3HostnameIpfsUniversalPathService{}
	r.Options = opts
	r.ContentList = NewZoneWeb3HostnameIpfsUniversalPathContentListService(opts...)
	return
}
