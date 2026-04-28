// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountAIRunCfService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfService] method instead.
type AccountAIRunCfService struct {
	Options         []option.RequestOption
	Baai            *AccountAIRunCfBaaiService
	BlackForestLabs *AccountAIRunCfBlackForestLabService
	Bytedance       *AccountAIRunCfBytedanceService
	DeepseekAI      *AccountAIRunCfDeepseekAIService
	Defog           *AccountAIRunCfDefogService
	Facebook        *AccountAIRunCfFacebookService
	Fblgit          *AccountAIRunCfFblgitService
	Google          *AccountAIRunCfGoogleService
	Huggingface     *AccountAIRunCfHuggingfaceService
	Lykon           *AccountAIRunCfLykonService
	MetaLlama       *AccountAIRunCfMetaLlamaService
	Meta            *AccountAIRunCfMetaService
	Microsoft       *AccountAIRunCfMicrosoftService
	Mistral         *AccountAIRunCfMistralService
	MyshellAI       *AccountAIRunCfMyshellAIService
	OpenAI          *AccountAIRunCfOpenAIService
	Openchat        *AccountAIRunCfOpenchatService
	Qwen            *AccountAIRunCfQwenService
	Runwayml        *AccountAIRunCfRunwaymlService
	Stabilityai     *AccountAIRunCfStabilityaiService
	Thebloke        *AccountAIRunCfTheblokeService
	Tiiuae          *AccountAIRunCfTiiuaeService
	Tinyllama       *AccountAIRunCfTinyllamaService
}

// NewAccountAIRunCfService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountAIRunCfService(opts ...option.RequestOption) (r *AccountAIRunCfService) {
	r = &AccountAIRunCfService{}
	r.Options = opts
	r.Baai = NewAccountAIRunCfBaaiService(opts...)
	r.BlackForestLabs = NewAccountAIRunCfBlackForestLabService(opts...)
	r.Bytedance = NewAccountAIRunCfBytedanceService(opts...)
	r.DeepseekAI = NewAccountAIRunCfDeepseekAIService(opts...)
	r.Defog = NewAccountAIRunCfDefogService(opts...)
	r.Facebook = NewAccountAIRunCfFacebookService(opts...)
	r.Fblgit = NewAccountAIRunCfFblgitService(opts...)
	r.Google = NewAccountAIRunCfGoogleService(opts...)
	r.Huggingface = NewAccountAIRunCfHuggingfaceService(opts...)
	r.Lykon = NewAccountAIRunCfLykonService(opts...)
	r.MetaLlama = NewAccountAIRunCfMetaLlamaService(opts...)
	r.Meta = NewAccountAIRunCfMetaService(opts...)
	r.Microsoft = NewAccountAIRunCfMicrosoftService(opts...)
	r.Mistral = NewAccountAIRunCfMistralService(opts...)
	r.MyshellAI = NewAccountAIRunCfMyshellAIService(opts...)
	r.OpenAI = NewAccountAIRunCfOpenAIService(opts...)
	r.Openchat = NewAccountAIRunCfOpenchatService(opts...)
	r.Qwen = NewAccountAIRunCfQwenService(opts...)
	r.Runwayml = NewAccountAIRunCfRunwaymlService(opts...)
	r.Stabilityai = NewAccountAIRunCfStabilityaiService(opts...)
	r.Thebloke = NewAccountAIRunCfTheblokeService(opts...)
	r.Tiiuae = NewAccountAIRunCfTiiuaeService(opts...)
	r.Tinyllama = NewAccountAIRunCfTinyllamaService(opts...)
	return
}
