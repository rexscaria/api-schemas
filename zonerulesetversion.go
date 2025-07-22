// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneRulesetVersionService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneRulesetVersionService] method instead.
type ZoneRulesetVersionService struct {
	Options []option.RequestOption
}

// NewZoneRulesetVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneRulesetVersionService(opts ...option.RequestOption) (r *ZoneRulesetVersionService) {
	r = &ZoneRulesetVersionService{}
	r.Options = opts
	return
}

// Fetches a specific version of a zone ruleset.
func (r *ZoneRulesetVersionService) Get(ctx context.Context, zoneID string, rulesetID string, rulesetVersion string, opts ...option.RequestOption) (res *ZoneRulesetVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	if rulesetVersion == "" {
		err = errors.New("missing required ruleset_version parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rulesets/%s/versions/%s", zoneID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches the versions of a zone ruleset.
func (r *ZoneRulesetVersionService) List(ctx context.Context, zoneID string, rulesetID string, opts ...option.RequestOption) (res *ZoneRulesetVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rulesets/%s/versions", zoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an existing version of a zone ruleset.
func (r *ZoneRulesetVersionService) Delete(ctx context.Context, zoneID string, rulesetID string, rulesetVersion string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	if rulesetVersion == "" {
		err = errors.New("missing required ruleset_version parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rulesets/%s/versions/%s", zoneID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetches the rules of a managed zone ruleset version for a given tag.
func (r *ZoneRulesetVersionService) ListByTag(ctx context.Context, zoneID string, rulesetID string, rulesetVersion string, ruleTag string, opts ...option.RequestOption) (res *ZoneRulesetVersionListByTagResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	if rulesetVersion == "" {
		err = errors.New("missing required ruleset_version parameter")
		return
	}
	if ruleTag == "" {
		err = errors.New("missing required rule_tag parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rulesets/%s/versions/%s/by_tag/%s", zoneID, rulesetID, rulesetVersion, ruleTag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneRulesetVersionGetResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneRulesetVersionGetResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneRulesetVersionGetResponseSuccess `json:"success,required"`
	JSON    zoneRulesetVersionGetResponseJSON    `json:"-"`
}

// zoneRulesetVersionGetResponseJSON contains the JSON metadata for the struct
// [ZoneRulesetVersionGetResponse]
type zoneRulesetVersionGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneRulesetVersionGetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetVersionGetResponseMessagesSource `json:"source"`
	JSON   zoneRulesetVersionGetResponseMessageJSON    `json:"-"`
}

// zoneRulesetVersionGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneRulesetVersionGetResponseMessage]
type zoneRulesetVersionGetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetVersionGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetVersionGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneRulesetVersionGetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    zoneRulesetVersionGetResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetVersionGetResponseMessagesSourceJSON contains the JSON metadata for
// the struct [ZoneRulesetVersionGetResponseMessagesSource]
type zoneRulesetVersionGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetVersionGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetVersionGetResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneRulesetVersionGetResponseSuccess bool

const (
	ZoneRulesetVersionGetResponseSuccessTrue ZoneRulesetVersionGetResponseSuccess = true
)

func (r ZoneRulesetVersionGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneRulesetVersionGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneRulesetVersionListResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneRulesetVersionListResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneRulesetVersionListResponseSuccess `json:"success,required"`
	// Cursors information to navigate the results.
	ResultInfo ZoneRulesetVersionListResponseResultInfo `json:"result_info"`
	JSON       zoneRulesetVersionListResponseJSON       `json:"-"`
}

// zoneRulesetVersionListResponseJSON contains the JSON metadata for the struct
// [ZoneRulesetVersionListResponse]
type zoneRulesetVersionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetVersionListResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneRulesetVersionListResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetVersionListResponseMessagesSource `json:"source"`
	JSON   zoneRulesetVersionListResponseMessageJSON    `json:"-"`
}

// zoneRulesetVersionListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneRulesetVersionListResponseMessage]
type zoneRulesetVersionListResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetVersionListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetVersionListResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneRulesetVersionListResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                           `json:"pointer,required"`
	JSON    zoneRulesetVersionListResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetVersionListResponseMessagesSourceJSON contains the JSON metadata for
// the struct [ZoneRulesetVersionListResponseMessagesSource]
type zoneRulesetVersionListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetVersionListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetVersionListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneRulesetVersionListResponseSuccess bool

const (
	ZoneRulesetVersionListResponseSuccessTrue ZoneRulesetVersionListResponseSuccess = true
)

func (r ZoneRulesetVersionListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneRulesetVersionListResponseSuccessTrue:
		return true
	}
	return false
}

// Cursors information to navigate the results.
type ZoneRulesetVersionListResponseResultInfo struct {
	// Set of cursors.
	Cursors ZoneRulesetVersionListResponseResultInfoCursors `json:"cursors"`
	JSON    zoneRulesetVersionListResponseResultInfoJSON    `json:"-"`
}

// zoneRulesetVersionListResponseResultInfoJSON contains the JSON metadata for the
// struct [ZoneRulesetVersionListResponseResultInfo]
type zoneRulesetVersionListResponseResultInfoJSON struct {
	Cursors     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetVersionListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetVersionListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

// Set of cursors.
type ZoneRulesetVersionListResponseResultInfoCursors struct {
	// Cursor to use for the next page.
	After string                                              `json:"after"`
	JSON  zoneRulesetVersionListResponseResultInfoCursorsJSON `json:"-"`
}

// zoneRulesetVersionListResponseResultInfoCursorsJSON contains the JSON metadata
// for the struct [ZoneRulesetVersionListResponseResultInfoCursors]
type zoneRulesetVersionListResponseResultInfoCursorsJSON struct {
	After       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetVersionListResponseResultInfoCursors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetVersionListResponseResultInfoCursorsJSON) RawJSON() string {
	return r.raw
}

type ZoneRulesetVersionListByTagResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneRulesetVersionListByTagResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneRulesetVersionListByTagResponseSuccess `json:"success,required"`
	JSON    zoneRulesetVersionListByTagResponseJSON    `json:"-"`
}

// zoneRulesetVersionListByTagResponseJSON contains the JSON metadata for the
// struct [ZoneRulesetVersionListByTagResponse]
type zoneRulesetVersionListByTagResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetVersionListByTagResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetVersionListByTagResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneRulesetVersionListByTagResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetVersionListByTagResponseMessagesSource `json:"source"`
	JSON   zoneRulesetVersionListByTagResponseMessageJSON    `json:"-"`
}

// zoneRulesetVersionListByTagResponseMessageJSON contains the JSON metadata for
// the struct [ZoneRulesetVersionListByTagResponseMessage]
type zoneRulesetVersionListByTagResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetVersionListByTagResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetVersionListByTagResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneRulesetVersionListByTagResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                `json:"pointer,required"`
	JSON    zoneRulesetVersionListByTagResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetVersionListByTagResponseMessagesSourceJSON contains the JSON metadata
// for the struct [ZoneRulesetVersionListByTagResponseMessagesSource]
type zoneRulesetVersionListByTagResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetVersionListByTagResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetVersionListByTagResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneRulesetVersionListByTagResponseSuccess bool

const (
	ZoneRulesetVersionListByTagResponseSuccessTrue ZoneRulesetVersionListByTagResponseSuccess = true
)

func (r ZoneRulesetVersionListByTagResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneRulesetVersionListByTagResponseSuccessTrue:
		return true
	}
	return false
}
