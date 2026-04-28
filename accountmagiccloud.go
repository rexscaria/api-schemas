// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountMagicCloudService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicCloudService] method instead.
type AccountMagicCloudService struct {
	Options      []option.RequestOption
	CatalogSyncs *AccountMagicCloudCatalogSyncService
	Onramps      *AccountMagicCloudOnrampService
	Providers    *AccountMagicCloudProviderService
	Resources    *AccountMagicCloudResourceService
}

// NewAccountMagicCloudService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMagicCloudService(opts ...option.RequestOption) (r *AccountMagicCloudService) {
	r = &AccountMagicCloudService{}
	r.Options = opts
	r.CatalogSyncs = NewAccountMagicCloudCatalogSyncService(opts...)
	r.Onramps = NewAccountMagicCloudOnrampService(opts...)
	r.Providers = NewAccountMagicCloudProviderService(opts...)
	r.Resources = NewAccountMagicCloudResourceService(opts...)
	return
}
