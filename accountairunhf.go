// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountAIRunHfService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunHfService] method instead.
type AccountAIRunHfService struct {
	Options      []option.RequestOption
	Google       *AccountAIRunHfGoogleService
	MetaLlama    *AccountAIRunHfMetaLlamaService
	Mistral      *AccountAIRunHfMistralService
	Mistralai    *AccountAIRunHfMistralaiService
	Nexusflow    *AccountAIRunHfNexusflowService
	Nousresearch *AccountAIRunHfNousresearchService
	Thebloke     *AccountAIRunHfTheblokeService
}

// NewAccountAIRunHfService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountAIRunHfService(opts ...option.RequestOption) (r *AccountAIRunHfService) {
	r = &AccountAIRunHfService{}
	r.Options = opts
	r.Google = NewAccountAIRunHfGoogleService(opts...)
	r.MetaLlama = NewAccountAIRunHfMetaLlamaService(opts...)
	r.Mistral = NewAccountAIRunHfMistralService(opts...)
	r.Mistralai = NewAccountAIRunHfMistralaiService(opts...)
	r.Nexusflow = NewAccountAIRunHfNexusflowService(opts...)
	r.Nousresearch = NewAccountAIRunHfNousresearchService(opts...)
	r.Thebloke = NewAccountAIRunHfTheblokeService(opts...)
	return
}
