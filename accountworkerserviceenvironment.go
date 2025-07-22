// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountWorkerServiceEnvironmentService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerServiceEnvironmentService] method instead.
type AccountWorkerServiceEnvironmentService struct {
	Options  []option.RequestOption
	Content  *AccountWorkerServiceEnvironmentContentService
	Settings *AccountWorkerServiceEnvironmentSettingService
}

// NewAccountWorkerServiceEnvironmentService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerServiceEnvironmentService(opts ...option.RequestOption) (r *AccountWorkerServiceEnvironmentService) {
	r = &AccountWorkerServiceEnvironmentService{}
	r.Options = opts
	r.Content = NewAccountWorkerServiceEnvironmentContentService(opts...)
	r.Settings = NewAccountWorkerServiceEnvironmentSettingService(opts...)
	return
}
