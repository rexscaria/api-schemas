// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/option"
)

// AccountAccessCertificateService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessCertificateService] method instead.
type AccountAccessCertificateService struct {
	Options  []option.RequestOption
	Settings *AccountAccessCertificateSettingService
}

// NewAccountAccessCertificateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessCertificateService(opts ...option.RequestOption) (r *AccountAccessCertificateService) {
	r = &AccountAccessCertificateService{}
	r.Options = opts
	r.Settings = NewAccountAccessCertificateSettingService(opts...)
	return
}

type IDResponseCertificates struct {
	Errors   []MessagesAccessItem `json:"errors,required"`
	Messages []MessagesAccessItem `json:"messages,required"`
	// Whether the API call was successful.
	Success IDResponseCertificatesSuccess `json:"success,required"`
	Result  IDResponseCertificatesResult  `json:"result"`
	JSON    idResponseCertificatesJSON    `json:"-"`
}

// idResponseCertificatesJSON contains the JSON metadata for the struct
// [IDResponseCertificates]
type idResponseCertificatesJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseCertificatesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IDResponseCertificatesSuccess bool

const (
	IDResponseCertificatesSuccessTrue IDResponseCertificatesSuccess = true
)

func (r IDResponseCertificatesSuccess) IsKnown() bool {
	switch r {
	case IDResponseCertificatesSuccessTrue:
		return true
	}
	return false
}

type IDResponseCertificatesResult struct {
	// UUID.
	ID   string                           `json:"id"`
	JSON idResponseCertificatesResultJSON `json:"-"`
}

// idResponseCertificatesResultJSON contains the JSON metadata for the struct
// [IDResponseCertificatesResult]
type idResponseCertificatesResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseCertificatesResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseCertificatesResultJSON) RawJSON() string {
	return r.raw
}

type MessagesAccessItem struct {
	Code             int64                    `json:"code,required"`
	Message          string                   `json:"message,required"`
	DocumentationURL string                   `json:"documentation_url"`
	Source           MessagesAccessItemSource `json:"source"`
	JSON             messagesAccessItemJSON   `json:"-"`
}

// messagesAccessItemJSON contains the JSON metadata for the struct
// [MessagesAccessItem]
type messagesAccessItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MessagesAccessItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesAccessItemJSON) RawJSON() string {
	return r.raw
}

type MessagesAccessItemSource struct {
	Pointer string                       `json:"pointer"`
	JSON    messagesAccessItemSourceJSON `json:"-"`
}

// messagesAccessItemSourceJSON contains the JSON metadata for the struct
// [MessagesAccessItemSource]
type messagesAccessItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesAccessItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesAccessItemSourceJSON) RawJSON() string {
	return r.raw
}
