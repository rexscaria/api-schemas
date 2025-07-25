// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountMagicSiteACLService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicSiteACLService] method instead.
type AccountMagicSiteACLService struct {
	Options []option.RequestOption
}

// NewAccountMagicSiteACLService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMagicSiteACLService(opts ...option.RequestOption) (r *AccountMagicSiteACLService) {
	r = &AccountMagicSiteACLService{}
	r.Options = opts
	return
}

// Creates a new Site ACL.
func (r *AccountMagicSiteACLService) New(ctx context.Context, accountID string, siteID string, body AccountMagicSiteACLNewParams, opts ...option.RequestOption) (res *MagicACLSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/acls", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific Site ACL.
func (r *AccountMagicSiteACLService) Get(ctx context.Context, accountID string, siteID string, aclID string, opts ...option.RequestOption) (res *MagicACLSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/acls/%s", accountID, siteID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a specific Site ACL.
func (r *AccountMagicSiteACLService) Update(ctx context.Context, accountID string, siteID string, aclID string, body AccountMagicSiteACLUpdateParams, opts ...option.RequestOption) (res *MagicACLModifiedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/acls/%s", accountID, siteID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists Site ACLs associated with an account.
func (r *AccountMagicSiteACLService) List(ctx context.Context, accountID string, siteID string, opts ...option.RequestOption) (res *AccountMagicSiteACLListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/acls", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove a specific Site ACL.
func (r *AccountMagicSiteACLService) Delete(ctx context.Context, accountID string, siteID string, aclID string, opts ...option.RequestOption) (res *AccountMagicSiteACLDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/acls/%s", accountID, siteID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Patch a specific Site ACL.
func (r *AccountMagicSiteACLService) Patch(ctx context.Context, accountID string, siteID string, aclID string, body AccountMagicSiteACLPatchParams, opts ...option.RequestOption) (res *MagicACLModifiedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/acls/%s", accountID, siteID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Bidirectional ACL policy for network traffic within a site.
type MagicACL struct {
	// Identifier
	ID string `json:"id"`
	// Description for the ACL.
	Description string `json:"description"`
	// The desired forwarding action for this ACL policy. If set to "false", the policy
	// will forward traffic to Cloudflare. If set to "true", the policy will forward
	// traffic locally on the Magic Connector. If not included in request, will default
	// to false.
	ForwardLocally bool                     `json:"forward_locally"`
	Lan1           MagicLanACLConfiguration `json:"lan_1"`
	Lan2           MagicLanACLConfiguration `json:"lan_2"`
	// The name of the ACL.
	Name      string             `json:"name"`
	Protocols []MagicACLProtocol `json:"protocols"`
	// The desired traffic direction for this ACL policy. If set to "false", the policy
	// will allow bidirectional traffic. If set to "true", the policy will only allow
	// traffic in one direction. If not included in request, will default to false.
	Unidirectional bool         `json:"unidirectional"`
	JSON           magicACLJSON `json:"-"`
}

// magicACLJSON contains the JSON metadata for the struct [MagicACL]
type magicACLJSON struct {
	ID             apijson.Field
	Description    apijson.Field
	ForwardLocally apijson.Field
	Lan1           apijson.Field
	Lan2           apijson.Field
	Name           apijson.Field
	Protocols      apijson.Field
	Unidirectional apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *MagicACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicACLJSON) RawJSON() string {
	return r.raw
}

// Array of allowed communication protocols between configured LANs. If no
// protocols are provided, all protocols are allowed.
type MagicACLProtocol string

const (
	MagicACLProtocolTcp  MagicACLProtocol = "tcp"
	MagicACLProtocolUdp  MagicACLProtocol = "udp"
	MagicACLProtocolIcmp MagicACLProtocol = "icmp"
)

func (r MagicACLProtocol) IsKnown() bool {
	switch r {
	case MagicACLProtocolTcp, MagicACLProtocolUdp, MagicACLProtocolIcmp:
		return true
	}
	return false
}

type MagicACLModifiedResponse struct {
	Errors   []MagicACLModifiedResponseError   `json:"errors,required"`
	Messages []MagicACLModifiedResponseMessage `json:"messages,required"`
	// Bidirectional ACL policy for network traffic within a site.
	Result MagicACL `json:"result,required"`
	// Whether the API call was successful
	Success MagicACLModifiedResponseSuccess `json:"success,required"`
	JSON    magicACLModifiedResponseJSON    `json:"-"`
}

// magicACLModifiedResponseJSON contains the JSON metadata for the struct
// [MagicACLModifiedResponse]
type magicACLModifiedResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicACLModifiedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicACLModifiedResponseJSON) RawJSON() string {
	return r.raw
}

type MagicACLModifiedResponseError struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           MagicACLModifiedResponseErrorsSource `json:"source"`
	JSON             magicACLModifiedResponseErrorJSON    `json:"-"`
}

// magicACLModifiedResponseErrorJSON contains the JSON metadata for the struct
// [MagicACLModifiedResponseError]
type magicACLModifiedResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MagicACLModifiedResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicACLModifiedResponseErrorJSON) RawJSON() string {
	return r.raw
}

type MagicACLModifiedResponseErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    magicACLModifiedResponseErrorsSourceJSON `json:"-"`
}

// magicACLModifiedResponseErrorsSourceJSON contains the JSON metadata for the
// struct [MagicACLModifiedResponseErrorsSource]
type magicACLModifiedResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicACLModifiedResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicACLModifiedResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type MagicACLModifiedResponseMessage struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           MagicACLModifiedResponseMessagesSource `json:"source"`
	JSON             magicACLModifiedResponseMessageJSON    `json:"-"`
}

// magicACLModifiedResponseMessageJSON contains the JSON metadata for the struct
// [MagicACLModifiedResponseMessage]
type magicACLModifiedResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MagicACLModifiedResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicACLModifiedResponseMessageJSON) RawJSON() string {
	return r.raw
}

type MagicACLModifiedResponseMessagesSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    magicACLModifiedResponseMessagesSourceJSON `json:"-"`
}

// magicACLModifiedResponseMessagesSourceJSON contains the JSON metadata for the
// struct [MagicACLModifiedResponseMessagesSource]
type magicACLModifiedResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicACLModifiedResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicACLModifiedResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicACLModifiedResponseSuccess bool

const (
	MagicACLModifiedResponseSuccessTrue MagicACLModifiedResponseSuccess = true
)

func (r MagicACLModifiedResponseSuccess) IsKnown() bool {
	switch r {
	case MagicACLModifiedResponseSuccessTrue:
		return true
	}
	return false
}

type MagicACLSingleResponse struct {
	Errors   []MagicACLSingleResponseError   `json:"errors,required"`
	Messages []MagicACLSingleResponseMessage `json:"messages,required"`
	// Bidirectional ACL policy for network traffic within a site.
	Result MagicACL `json:"result,required"`
	// Whether the API call was successful
	Success MagicACLSingleResponseSuccess `json:"success,required"`
	JSON    magicACLSingleResponseJSON    `json:"-"`
}

// magicACLSingleResponseJSON contains the JSON metadata for the struct
// [MagicACLSingleResponse]
type magicACLSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicACLSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicACLSingleResponseJSON) RawJSON() string {
	return r.raw
}

type MagicACLSingleResponseError struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           MagicACLSingleResponseErrorsSource `json:"source"`
	JSON             magicACLSingleResponseErrorJSON    `json:"-"`
}

// magicACLSingleResponseErrorJSON contains the JSON metadata for the struct
// [MagicACLSingleResponseError]
type magicACLSingleResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MagicACLSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicACLSingleResponseErrorJSON) RawJSON() string {
	return r.raw
}

type MagicACLSingleResponseErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    magicACLSingleResponseErrorsSourceJSON `json:"-"`
}

// magicACLSingleResponseErrorsSourceJSON contains the JSON metadata for the struct
// [MagicACLSingleResponseErrorsSource]
type magicACLSingleResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicACLSingleResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicACLSingleResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type MagicACLSingleResponseMessage struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           MagicACLSingleResponseMessagesSource `json:"source"`
	JSON             magicACLSingleResponseMessageJSON    `json:"-"`
}

// magicACLSingleResponseMessageJSON contains the JSON metadata for the struct
// [MagicACLSingleResponseMessage]
type magicACLSingleResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MagicACLSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicACLSingleResponseMessageJSON) RawJSON() string {
	return r.raw
}

type MagicACLSingleResponseMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    magicACLSingleResponseMessagesSourceJSON `json:"-"`
}

// magicACLSingleResponseMessagesSourceJSON contains the JSON metadata for the
// struct [MagicACLSingleResponseMessagesSource]
type magicACLSingleResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicACLSingleResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicACLSingleResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicACLSingleResponseSuccess bool

const (
	MagicACLSingleResponseSuccessTrue MagicACLSingleResponseSuccess = true
)

func (r MagicACLSingleResponseSuccess) IsKnown() bool {
	switch r {
	case MagicACLSingleResponseSuccessTrue:
		return true
	}
	return false
}

type MagicACLUpdateRequestParam struct {
	// Description for the ACL.
	Description param.Field[string] `json:"description"`
	// The desired forwarding action for this ACL policy. If set to "false", the policy
	// will forward traffic to Cloudflare. If set to "true", the policy will forward
	// traffic locally on the Magic Connector. If not included in request, will default
	// to false.
	ForwardLocally param.Field[bool]                          `json:"forward_locally"`
	Lan1           param.Field[MagicLanACLConfigurationParam] `json:"lan_1"`
	Lan2           param.Field[MagicLanACLConfigurationParam] `json:"lan_2"`
	// The name of the ACL.
	Name      param.Field[string]                          `json:"name"`
	Protocols param.Field[[]MagicACLUpdateRequestProtocol] `json:"protocols"`
	// The desired traffic direction for this ACL policy. If set to "false", the policy
	// will allow bidirectional traffic. If set to "true", the policy will only allow
	// traffic in one direction. If not included in request, will default to false.
	Unidirectional param.Field[bool] `json:"unidirectional"`
}

func (r MagicACLUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Array of allowed communication protocols between configured LANs. If no
// protocols are provided, all protocols are allowed.
type MagicACLUpdateRequestProtocol string

const (
	MagicACLUpdateRequestProtocolTcp  MagicACLUpdateRequestProtocol = "tcp"
	MagicACLUpdateRequestProtocolUdp  MagicACLUpdateRequestProtocol = "udp"
	MagicACLUpdateRequestProtocolIcmp MagicACLUpdateRequestProtocol = "icmp"
)

func (r MagicACLUpdateRequestProtocol) IsKnown() bool {
	switch r {
	case MagicACLUpdateRequestProtocolTcp, MagicACLUpdateRequestProtocolUdp, MagicACLUpdateRequestProtocolIcmp:
		return true
	}
	return false
}

type MagicLanACLConfiguration struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LanID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LanName string `json:"lan_name"`
	// Array of port ranges on the provided LAN that will be included in the ACL. If no
	// ports or port rangess are provided, communication on any port on this LAN is
	// allowed.
	PortRanges []string `json:"port_ranges"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// or port ranges are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []string                     `json:"subnets"`
	JSON    magicLanACLConfigurationJSON `json:"-"`
}

// magicLanACLConfigurationJSON contains the JSON metadata for the struct
// [MagicLanACLConfiguration]
type magicLanACLConfigurationJSON struct {
	LanID       apijson.Field
	LanName     apijson.Field
	PortRanges  apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicLanACLConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicLanACLConfigurationJSON) RawJSON() string {
	return r.raw
}

type MagicLanACLConfigurationParam struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LanID param.Field[string] `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LanName param.Field[string] `json:"lan_name"`
	// Array of port ranges on the provided LAN that will be included in the ACL. If no
	// ports or port rangess are provided, communication on any port on this LAN is
	// allowed.
	PortRanges param.Field[[]string] `json:"port_ranges"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// or port ranges are provided, communication on any port on this LAN is allowed.
	Ports param.Field[[]int64] `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets param.Field[[]string] `json:"subnets"`
}

func (r MagicLanACLConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicSiteACLListResponse struct {
	Errors   []AccountMagicSiteACLListResponseError   `json:"errors,required"`
	Messages []AccountMagicSiteACLListResponseMessage `json:"messages,required"`
	Result   []MagicACL                               `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicSiteACLListResponseSuccess `json:"success,required"`
	JSON    accountMagicSiteACLListResponseJSON    `json:"-"`
}

// accountMagicSiteACLListResponseJSON contains the JSON metadata for the struct
// [AccountMagicSiteACLListResponse]
type accountMagicSiteACLListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicSiteACLListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteACLListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicSiteACLListResponseError struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           AccountMagicSiteACLListResponseErrorsSource `json:"source"`
	JSON             accountMagicSiteACLListResponseErrorJSON    `json:"-"`
}

// accountMagicSiteACLListResponseErrorJSON contains the JSON metadata for the
// struct [AccountMagicSiteACLListResponseError]
type accountMagicSiteACLListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountMagicSiteACLListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteACLListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountMagicSiteACLListResponseErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    accountMagicSiteACLListResponseErrorsSourceJSON `json:"-"`
}

// accountMagicSiteACLListResponseErrorsSourceJSON contains the JSON metadata for
// the struct [AccountMagicSiteACLListResponseErrorsSource]
type accountMagicSiteACLListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicSiteACLListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteACLListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountMagicSiteACLListResponseMessage struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           AccountMagicSiteACLListResponseMessagesSource `json:"source"`
	JSON             accountMagicSiteACLListResponseMessageJSON    `json:"-"`
}

// accountMagicSiteACLListResponseMessageJSON contains the JSON metadata for the
// struct [AccountMagicSiteACLListResponseMessage]
type accountMagicSiteACLListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountMagicSiteACLListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteACLListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountMagicSiteACLListResponseMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    accountMagicSiteACLListResponseMessagesSourceJSON `json:"-"`
}

// accountMagicSiteACLListResponseMessagesSourceJSON contains the JSON metadata for
// the struct [AccountMagicSiteACLListResponseMessagesSource]
type accountMagicSiteACLListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicSiteACLListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteACLListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountMagicSiteACLListResponseSuccess bool

const (
	AccountMagicSiteACLListResponseSuccessTrue AccountMagicSiteACLListResponseSuccess = true
)

func (r AccountMagicSiteACLListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicSiteACLListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicSiteACLDeleteResponse struct {
	Errors   []AccountMagicSiteACLDeleteResponseError   `json:"errors,required"`
	Messages []AccountMagicSiteACLDeleteResponseMessage `json:"messages,required"`
	// Bidirectional ACL policy for network traffic within a site.
	Result MagicACL `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicSiteACLDeleteResponseSuccess `json:"success,required"`
	JSON    accountMagicSiteACLDeleteResponseJSON    `json:"-"`
}

// accountMagicSiteACLDeleteResponseJSON contains the JSON metadata for the struct
// [AccountMagicSiteACLDeleteResponse]
type accountMagicSiteACLDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicSiteACLDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteACLDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicSiteACLDeleteResponseError struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           AccountMagicSiteACLDeleteResponseErrorsSource `json:"source"`
	JSON             accountMagicSiteACLDeleteResponseErrorJSON    `json:"-"`
}

// accountMagicSiteACLDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountMagicSiteACLDeleteResponseError]
type accountMagicSiteACLDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountMagicSiteACLDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteACLDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountMagicSiteACLDeleteResponseErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    accountMagicSiteACLDeleteResponseErrorsSourceJSON `json:"-"`
}

// accountMagicSiteACLDeleteResponseErrorsSourceJSON contains the JSON metadata for
// the struct [AccountMagicSiteACLDeleteResponseErrorsSource]
type accountMagicSiteACLDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicSiteACLDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteACLDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountMagicSiteACLDeleteResponseMessage struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           AccountMagicSiteACLDeleteResponseMessagesSource `json:"source"`
	JSON             accountMagicSiteACLDeleteResponseMessageJSON    `json:"-"`
}

// accountMagicSiteACLDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountMagicSiteACLDeleteResponseMessage]
type accountMagicSiteACLDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountMagicSiteACLDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteACLDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountMagicSiteACLDeleteResponseMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    accountMagicSiteACLDeleteResponseMessagesSourceJSON `json:"-"`
}

// accountMagicSiteACLDeleteResponseMessagesSourceJSON contains the JSON metadata
// for the struct [AccountMagicSiteACLDeleteResponseMessagesSource]
type accountMagicSiteACLDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicSiteACLDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteACLDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountMagicSiteACLDeleteResponseSuccess bool

const (
	AccountMagicSiteACLDeleteResponseSuccessTrue AccountMagicSiteACLDeleteResponseSuccess = true
)

func (r AccountMagicSiteACLDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicSiteACLDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicSiteACLNewParams struct {
	Lan1 param.Field[MagicLanACLConfigurationParam] `json:"lan_1,required"`
	Lan2 param.Field[MagicLanACLConfigurationParam] `json:"lan_2,required"`
	// The name of the ACL.
	Name param.Field[string] `json:"name,required"`
	// Description for the ACL.
	Description param.Field[string] `json:"description"`
	// The desired forwarding action for this ACL policy. If set to "false", the policy
	// will forward traffic to Cloudflare. If set to "true", the policy will forward
	// traffic locally on the Magic Connector. If not included in request, will default
	// to false.
	ForwardLocally param.Field[bool]                                   `json:"forward_locally"`
	Protocols      param.Field[[]AccountMagicSiteACLNewParamsProtocol] `json:"protocols"`
	// The desired traffic direction for this ACL policy. If set to "false", the policy
	// will allow bidirectional traffic. If set to "true", the policy will only allow
	// traffic in one direction. If not included in request, will default to false.
	Unidirectional param.Field[bool] `json:"unidirectional"`
}

func (r AccountMagicSiteACLNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Array of allowed communication protocols between configured LANs. If no
// protocols are provided, all protocols are allowed.
type AccountMagicSiteACLNewParamsProtocol string

const (
	AccountMagicSiteACLNewParamsProtocolTcp  AccountMagicSiteACLNewParamsProtocol = "tcp"
	AccountMagicSiteACLNewParamsProtocolUdp  AccountMagicSiteACLNewParamsProtocol = "udp"
	AccountMagicSiteACLNewParamsProtocolIcmp AccountMagicSiteACLNewParamsProtocol = "icmp"
)

func (r AccountMagicSiteACLNewParamsProtocol) IsKnown() bool {
	switch r {
	case AccountMagicSiteACLNewParamsProtocolTcp, AccountMagicSiteACLNewParamsProtocolUdp, AccountMagicSiteACLNewParamsProtocolIcmp:
		return true
	}
	return false
}

type AccountMagicSiteACLUpdateParams struct {
	MagicACLUpdateRequest MagicACLUpdateRequestParam `json:"magic_acl_update_request,required"`
}

func (r AccountMagicSiteACLUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MagicACLUpdateRequest)
}

type AccountMagicSiteACLPatchParams struct {
	MagicACLUpdateRequest MagicACLUpdateRequestParam `json:"magic_acl_update_request,required"`
}

func (r AccountMagicSiteACLPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MagicACLUpdateRequest)
}
