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

// ZoneRulesetPhaseEntrypointVersionService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneRulesetPhaseEntrypointVersionService] method instead.
type ZoneRulesetPhaseEntrypointVersionService struct {
	Options []option.RequestOption
}

// NewZoneRulesetPhaseEntrypointVersionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneRulesetPhaseEntrypointVersionService(opts ...option.RequestOption) (r *ZoneRulesetPhaseEntrypointVersionService) {
	r = &ZoneRulesetPhaseEntrypointVersionService{}
	r.Options = opts
	return
}

// Fetches a specific version of a zone entry point ruleset.
func (r *ZoneRulesetPhaseEntrypointVersionService) Get(ctx context.Context, zoneID string, rulesetPhase RulesetPhase, rulesetVersion string, opts ...option.RequestOption) (res *ZoneRulesetPhaseEntrypointVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if rulesetVersion == "" {
		err = errors.New("missing required ruleset_version parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rulesets/phases/%v/entrypoint/versions/%s", zoneID, rulesetPhase, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches the versions of a zone entry point ruleset.
func (r *ZoneRulesetPhaseEntrypointVersionService) List(ctx context.Context, zoneID string, rulesetPhase RulesetPhase, opts ...option.RequestOption) (res *ZoneRulesetPhaseEntrypointVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rulesets/phases/%v/entrypoint/versions", zoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneRulesetPhaseEntrypointVersionGetResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneRulesetPhaseEntrypointVersionGetResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneRulesetPhaseEntrypointVersionGetResponseSuccess `json:"success,required"`
	JSON    zoneRulesetPhaseEntrypointVersionGetResponseJSON    `json:"-"`
}

// zoneRulesetPhaseEntrypointVersionGetResponseJSON contains the JSON metadata for
// the struct [ZoneRulesetPhaseEntrypointVersionGetResponse]
type zoneRulesetPhaseEntrypointVersionGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetPhaseEntrypointVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneRulesetPhaseEntrypointVersionGetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetPhaseEntrypointVersionGetResponseMessagesSource `json:"source"`
	JSON   zoneRulesetPhaseEntrypointVersionGetResponseMessageJSON    `json:"-"`
}

// zoneRulesetPhaseEntrypointVersionGetResponseMessageJSON contains the JSON
// metadata for the struct [ZoneRulesetPhaseEntrypointVersionGetResponseMessage]
type zoneRulesetPhaseEntrypointVersionGetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointVersionGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetPhaseEntrypointVersionGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneRulesetPhaseEntrypointVersionGetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                         `json:"pointer,required"`
	JSON    zoneRulesetPhaseEntrypointVersionGetResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetPhaseEntrypointVersionGetResponseMessagesSourceJSON contains the JSON
// metadata for the struct
// [ZoneRulesetPhaseEntrypointVersionGetResponseMessagesSource]
type zoneRulesetPhaseEntrypointVersionGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointVersionGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetPhaseEntrypointVersionGetResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneRulesetPhaseEntrypointVersionGetResponseSuccess bool

const (
	ZoneRulesetPhaseEntrypointVersionGetResponseSuccessTrue ZoneRulesetPhaseEntrypointVersionGetResponseSuccess = true
)

func (r ZoneRulesetPhaseEntrypointVersionGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneRulesetPhaseEntrypointVersionGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneRulesetPhaseEntrypointVersionListResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneRulesetPhaseEntrypointVersionListResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneRulesetPhaseEntrypointVersionListResponseSuccess `json:"success,required"`
	// Cursors information to navigate the results.
	ResultInfo ZoneRulesetPhaseEntrypointVersionListResponseResultInfo `json:"result_info"`
	JSON       zoneRulesetPhaseEntrypointVersionListResponseJSON       `json:"-"`
}

// zoneRulesetPhaseEntrypointVersionListResponseJSON contains the JSON metadata for
// the struct [ZoneRulesetPhaseEntrypointVersionListResponse]
type zoneRulesetPhaseEntrypointVersionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetPhaseEntrypointVersionListResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneRulesetPhaseEntrypointVersionListResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetPhaseEntrypointVersionListResponseMessagesSource `json:"source"`
	JSON   zoneRulesetPhaseEntrypointVersionListResponseMessageJSON    `json:"-"`
}

// zoneRulesetPhaseEntrypointVersionListResponseMessageJSON contains the JSON
// metadata for the struct [ZoneRulesetPhaseEntrypointVersionListResponseMessage]
type zoneRulesetPhaseEntrypointVersionListResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointVersionListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetPhaseEntrypointVersionListResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneRulesetPhaseEntrypointVersionListResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                          `json:"pointer,required"`
	JSON    zoneRulesetPhaseEntrypointVersionListResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetPhaseEntrypointVersionListResponseMessagesSourceJSON contains the
// JSON metadata for the struct
// [ZoneRulesetPhaseEntrypointVersionListResponseMessagesSource]
type zoneRulesetPhaseEntrypointVersionListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointVersionListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetPhaseEntrypointVersionListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneRulesetPhaseEntrypointVersionListResponseSuccess bool

const (
	ZoneRulesetPhaseEntrypointVersionListResponseSuccessTrue ZoneRulesetPhaseEntrypointVersionListResponseSuccess = true
)

func (r ZoneRulesetPhaseEntrypointVersionListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneRulesetPhaseEntrypointVersionListResponseSuccessTrue:
		return true
	}
	return false
}

// Cursors information to navigate the results.
type ZoneRulesetPhaseEntrypointVersionListResponseResultInfo struct {
	// Set of cursors.
	Cursors ZoneRulesetPhaseEntrypointVersionListResponseResultInfoCursors `json:"cursors"`
	JSON    zoneRulesetPhaseEntrypointVersionListResponseResultInfoJSON    `json:"-"`
}

// zoneRulesetPhaseEntrypointVersionListResponseResultInfoJSON contains the JSON
// metadata for the struct
// [ZoneRulesetPhaseEntrypointVersionListResponseResultInfo]
type zoneRulesetPhaseEntrypointVersionListResponseResultInfoJSON struct {
	Cursors     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointVersionListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetPhaseEntrypointVersionListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

// Set of cursors.
type ZoneRulesetPhaseEntrypointVersionListResponseResultInfoCursors struct {
	// Cursor to use for the next page.
	After string                                                             `json:"after"`
	JSON  zoneRulesetPhaseEntrypointVersionListResponseResultInfoCursorsJSON `json:"-"`
}

// zoneRulesetPhaseEntrypointVersionListResponseResultInfoCursorsJSON contains the
// JSON metadata for the struct
// [ZoneRulesetPhaseEntrypointVersionListResponseResultInfoCursors]
type zoneRulesetPhaseEntrypointVersionListResponseResultInfoCursorsJSON struct {
	After       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointVersionListResponseResultInfoCursors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetPhaseEntrypointVersionListResponseResultInfoCursorsJSON) RawJSON() string {
	return r.raw
}
