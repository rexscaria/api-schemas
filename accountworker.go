// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountWorkerService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerService] method instead.
type AccountWorkerService struct {
	Options         []option.RequestOption
	AccountSettings *AccountWorkerAccountSettingService
	Assets          *AccountWorkerAssetService
	Dispatch        *AccountWorkerDispatchService
	Domains         *AccountWorkerDomainService
	DurableObjects  *AccountWorkerDurableObjectService
	Scripts         *AccountWorkerScriptService
	Services        *AccountWorkerServiceService
	Subdomain       *AccountWorkerSubdomainService
}

// NewAccountWorkerService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountWorkerService(opts ...option.RequestOption) (r *AccountWorkerService) {
	r = &AccountWorkerService{}
	r.Options = opts
	r.AccountSettings = NewAccountWorkerAccountSettingService(opts...)
	r.Assets = NewAccountWorkerAssetService(opts...)
	r.Dispatch = NewAccountWorkerDispatchService(opts...)
	r.Domains = NewAccountWorkerDomainService(opts...)
	r.DurableObjects = NewAccountWorkerDurableObjectService(opts...)
	r.Scripts = NewAccountWorkerScriptService(opts...)
	r.Services = NewAccountWorkerServiceService(opts...)
	r.Subdomain = NewAccountWorkerSubdomainService(opts...)
	return
}
