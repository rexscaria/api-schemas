// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneSnippetSnippetRuleService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSnippetSnippetRuleService] method instead.
type ZoneSnippetSnippetRuleService struct {
	Options []option.RequestOption
}

// NewZoneSnippetSnippetRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSnippetSnippetRuleService(opts ...option.RequestOption) (r *ZoneSnippetSnippetRuleService) {
	r = &ZoneSnippetSnippetRuleService{}
	r.Options = opts
	return
}

// Put Rules
func (r *ZoneSnippetSnippetRuleService) Update(ctx context.Context, zoneID string, body ZoneSnippetSnippetRuleUpdateParams, opts ...option.RequestOption) (res *ZoneSnippetSnippetRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Rules
func (r *ZoneSnippetSnippetRuleService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneSnippetSnippetRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete All Rules
func (r *ZoneSnippetSnippetRuleService) DeleteAll(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *APIResponseSnippets, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type SnippetRuleItem struct {
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Expression  string `json:"expression"`
	// Snippet identifying name
	SnippetName string              `json:"snippet_name"`
	JSON        snippetRuleItemJSON `json:"-"`
}

// snippetRuleItemJSON contains the JSON metadata for the struct [SnippetRuleItem]
type snippetRuleItemJSON struct {
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetRuleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetRuleItemJSON) RawJSON() string {
	return r.raw
}

type SnippetRuleItemParam struct {
	Description param.Field[string] `json:"description"`
	Enabled     param.Field[bool]   `json:"enabled"`
	Expression  param.Field[string] `json:"expression"`
	// Snippet identifying name
	SnippetName param.Field[string] `json:"snippet_name"`
}

func (r SnippetRuleItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSnippetSnippetRuleUpdateResponse struct {
	// List of snippet rules
	Result []SnippetRuleItem                        `json:"result"`
	JSON   zoneSnippetSnippetRuleUpdateResponseJSON `json:"-"`
	APIResponseSnippets
}

// zoneSnippetSnippetRuleUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSnippetSnippetRuleUpdateResponse]
type zoneSnippetSnippetRuleUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetSnippetRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetSnippetRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSnippetSnippetRuleListResponse struct {
	// List of snippet rules
	Result []SnippetRuleItem                      `json:"result"`
	JSON   zoneSnippetSnippetRuleListResponseJSON `json:"-"`
	APIResponseSnippets
}

// zoneSnippetSnippetRuleListResponseJSON contains the JSON metadata for the struct
// [ZoneSnippetSnippetRuleListResponse]
type zoneSnippetSnippetRuleListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetSnippetRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetSnippetRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSnippetSnippetRuleUpdateParams struct {
	// List of snippet rules
	Rules param.Field[[]SnippetRuleItemParam] `json:"rules"`
}

func (r ZoneSnippetSnippetRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
