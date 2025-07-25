// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountWorkerDispatchNamespaceScriptSettingService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerDispatchNamespaceScriptSettingService] method instead.
type AccountWorkerDispatchNamespaceScriptSettingService struct {
	Options []option.RequestOption
}

// NewAccountWorkerDispatchNamespaceScriptSettingService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewAccountWorkerDispatchNamespaceScriptSettingService(opts ...option.RequestOption) (r *AccountWorkerDispatchNamespaceScriptSettingService) {
	r = &AccountWorkerDispatchNamespaceScriptSettingService{}
	r.Options = opts
	return
}

// Get script settings from a script uploaded to a Workers for Platforms namespace.
func (r *AccountWorkerDispatchNamespaceScriptSettingService) Get(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, opts ...option.RequestOption) (res *AccountWorkerDispatchNamespaceScriptSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patch script metadata, such as bindings.
func (r *AccountWorkerDispatchNamespaceScriptSettingService) Patch(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, body AccountWorkerDispatchNamespaceScriptSettingPatchParams, opts ...option.RequestOption) (res *AccountWorkerDispatchNamespaceScriptSettingPatchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type MigrationStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes to create Durable Object namespaces with SQLite from.
	NewSqliteClasses []string `json:"new_sqlite_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []MigrationStepRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []MigrationStepTransferredClass `json:"transferred_classes"`
	JSON               migrationStepJSON               `json:"-"`
}

// migrationStepJSON contains the JSON metadata for the struct [MigrationStep]
type migrationStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewSqliteClasses   apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MigrationStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r migrationStepJSON) RawJSON() string {
	return r.raw
}

type MigrationStepRenamedClass struct {
	From string                        `json:"from"`
	To   string                        `json:"to"`
	JSON migrationStepRenamedClassJSON `json:"-"`
}

// migrationStepRenamedClassJSON contains the JSON metadata for the struct
// [MigrationStepRenamedClass]
type migrationStepRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MigrationStepRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r migrationStepRenamedClassJSON) RawJSON() string {
	return r.raw
}

type MigrationStepTransferredClass struct {
	From       string                            `json:"from"`
	FromScript string                            `json:"from_script"`
	To         string                            `json:"to"`
	JSON       migrationStepTransferredClassJSON `json:"-"`
}

// migrationStepTransferredClassJSON contains the JSON metadata for the struct
// [MigrationStepTransferredClass]
type migrationStepTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MigrationStepTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r migrationStepTransferredClassJSON) RawJSON() string {
	return r.raw
}

type MigrationStepParam struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes to create Durable Object namespaces with SQLite from.
	NewSqliteClasses param.Field[[]string] `json:"new_sqlite_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]MigrationStepRenamedClassParam] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]MigrationStepTransferredClassParam] `json:"transferred_classes"`
}

func (r MigrationStepParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MigrationStepRenamedClassParam struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r MigrationStepRenamedClassParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MigrationStepTransferredClassParam struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r MigrationStepTransferredClassParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type PlacementMode string

const (
	PlacementModeSmart PlacementMode = "smart"
)

func (r PlacementMode) IsKnown() bool {
	switch r {
	case PlacementModeSmart:
		return true
	}
	return false
}

type ScriptVersionItem struct {
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings []BindingItem `json:"bindings"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate string `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Limits to apply for this Worker.
	Limits ScriptVersionItemLimits `json:"limits"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Observability settings for the Worker.
	Observability Observability `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement ScriptVersionItemPlacement `json:"placement"`
	// Tags to help you manage your Workers.
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []TailConsumersScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel UsageModel            `json:"usage_model"`
	JSON       scriptVersionItemJSON `json:"-"`
}

// scriptVersionItemJSON contains the JSON metadata for the struct
// [ScriptVersionItem]
type scriptVersionItemJSON struct {
	Bindings           apijson.Field
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	Limits             apijson.Field
	Logpush            apijson.Field
	Observability      apijson.Field
	Placement          apijson.Field
	Tags               apijson.Field
	TailConsumers      apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptVersionItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionItemJSON) RawJSON() string {
	return r.raw
}

// Limits to apply for this Worker.
type ScriptVersionItemLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	CPUMs int64                       `json:"cpu_ms"`
	JSON  scriptVersionItemLimitsJSON `json:"-"`
}

// scriptVersionItemLimitsJSON contains the JSON metadata for the struct
// [ScriptVersionItemLimits]
type scriptVersionItemLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionItemLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionItemLimitsJSON) RawJSON() string {
	return r.raw
}

// Migrations to apply for Durable Objects associated with this Worker.
type ScriptVersionItemMigrations struct {
	// This field can have the runtime type of [[]string].
	DeletedClasses interface{} `json:"deleted_classes"`
	// This field can have the runtime type of [[]string].
	NewClasses interface{} `json:"new_classes"`
	// This field can have the runtime type of [[]string].
	NewSqliteClasses interface{} `json:"new_sqlite_classes"`
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// This field can have the runtime type of
	// [[]ScriptVersionItemMigrationsWorkersSingleStepMigrationsRenamedClass].
	RenamedClasses interface{} `json:"renamed_classes"`
	// This field can have the runtime type of [[]MigrationStep].
	Steps interface{} `json:"steps"`
	// This field can have the runtime type of
	// [[]ScriptVersionItemMigrationsWorkersSingleStepMigrationsTransferredClass].
	TransferredClasses interface{}                     `json:"transferred_classes"`
	JSON               scriptVersionItemMigrationsJSON `json:"-"`
	union              ScriptVersionItemMigrationsUnion
}

// scriptVersionItemMigrationsJSON contains the JSON metadata for the struct
// [ScriptVersionItemMigrations]
type scriptVersionItemMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewSqliteClasses   apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	Steps              apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r scriptVersionItemMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptVersionItemMigrations) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptVersionItemMigrations{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptVersionItemMigrationsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ScriptVersionItemMigrationsWorkersSingleStepMigrations],
// [ScriptVersionItemMigrationsWorkersMultipleStepMigrations].
func (r ScriptVersionItemMigrations) AsUnion() ScriptVersionItemMigrationsUnion {
	return r.union
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by [ScriptVersionItemMigrationsWorkersSingleStepMigrations] or
// [ScriptVersionItemMigrationsWorkersMultipleStepMigrations].
type ScriptVersionItemMigrationsUnion interface {
	implementsScriptVersionItemMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptVersionItemMigrationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionItemMigrationsWorkersSingleStepMigrations{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionItemMigrationsWorkersMultipleStepMigrations{}),
		},
	)
}

// A single set of migrations to apply.
type ScriptVersionItemMigrationsWorkersSingleStepMigrations struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes to create Durable Object namespaces with SQLite from.
	NewSqliteClasses []string `json:"new_sqlite_classes"`
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []ScriptVersionItemMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []ScriptVersionItemMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               scriptVersionItemMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// scriptVersionItemMigrationsWorkersSingleStepMigrationsJSON contains the JSON
// metadata for the struct [ScriptVersionItemMigrationsWorkersSingleStepMigrations]
type scriptVersionItemMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewSqliteClasses   apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptVersionItemMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionItemMigrationsWorkersSingleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionItemMigrationsWorkersSingleStepMigrations) implementsScriptVersionItemMigrations() {
}

type ScriptVersionItemMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                 `json:"from"`
	To   string                                                                 `json:"to"`
	JSON scriptVersionItemMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// scriptVersionItemMigrationsWorkersSingleStepMigrationsRenamedClassJSON contains
// the JSON metadata for the struct
// [ScriptVersionItemMigrationsWorkersSingleStepMigrationsRenamedClass]
type scriptVersionItemMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionItemMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionItemMigrationsWorkersSingleStepMigrationsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionItemMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                     `json:"from"`
	FromScript string                                                                     `json:"from_script"`
	To         string                                                                     `json:"to"`
	JSON       scriptVersionItemMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// scriptVersionItemMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [ScriptVersionItemMigrationsWorkersSingleStepMigrationsTransferredClass]
type scriptVersionItemMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionItemMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionItemMigrationsWorkersSingleStepMigrationsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionItemMigrationsWorkersMultipleStepMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []MigrationStep                                              `json:"steps"`
	JSON  scriptVersionItemMigrationsWorkersMultipleStepMigrationsJSON `json:"-"`
}

// scriptVersionItemMigrationsWorkersMultipleStepMigrationsJSON contains the JSON
// metadata for the struct
// [ScriptVersionItemMigrationsWorkersMultipleStepMigrations]
type scriptVersionItemMigrationsWorkersMultipleStepMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionItemMigrationsWorkersMultipleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionItemMigrationsWorkersMultipleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionItemMigrationsWorkersMultipleStepMigrations) implementsScriptVersionItemMigrations() {
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptVersionItemPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode PlacementMode                  `json:"mode"`
	JSON scriptVersionItemPlacementJSON `json:"-"`
}

// scriptVersionItemPlacementJSON contains the JSON metadata for the struct
// [ScriptVersionItemPlacement]
type scriptVersionItemPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionItemPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionItemPlacementJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionItemParam struct {
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]BindingItemUnionParam] `json:"bindings"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Limits to apply for this Worker.
	Limits param.Field[ScriptVersionItemLimitsParam] `json:"limits"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[ScriptVersionItemMigrationsUnionParam] `json:"migrations"`
	// Observability settings for the Worker.
	Observability param.Field[ObservabilityParam] `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement param.Field[ScriptVersionItemPlacementParam] `json:"placement"`
	// Tags to help you manage your Workers.
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]TailConsumersScriptParam] `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[UsageModel] `json:"usage_model"`
}

func (r ScriptVersionItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Limits to apply for this Worker.
type ScriptVersionItemLimitsParam struct {
	// The amount of CPU time this Worker can use in milliseconds.
	CPUMs param.Field[int64] `json:"cpu_ms"`
}

func (r ScriptVersionItemLimitsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations to apply for Durable Objects associated with this Worker.
type ScriptVersionItemMigrationsParam struct {
	DeletedClasses   param.Field[interface{}] `json:"deleted_classes"`
	NewClasses       param.Field[interface{}] `json:"new_classes"`
	NewSqliteClasses param.Field[interface{}] `json:"new_sqlite_classes"`
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag             param.Field[string]      `json:"old_tag"`
	RenamedClasses     param.Field[interface{}] `json:"renamed_classes"`
	Steps              param.Field[interface{}] `json:"steps"`
	TransferredClasses param.Field[interface{}] `json:"transferred_classes"`
}

func (r ScriptVersionItemMigrationsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionItemMigrationsParam) implementsScriptVersionItemMigrationsUnionParam() {}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by [ScriptVersionItemMigrationsWorkersSingleStepMigrationsParam],
// [ScriptVersionItemMigrationsWorkersMultipleStepMigrationsParam],
// [ScriptVersionItemMigrationsParam].
type ScriptVersionItemMigrationsUnionParam interface {
	implementsScriptVersionItemMigrationsUnionParam()
}

// A single set of migrations to apply.
type ScriptVersionItemMigrationsWorkersSingleStepMigrationsParam struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes to create Durable Object namespaces with SQLite from.
	NewSqliteClasses param.Field[[]string] `json:"new_sqlite_classes"`
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]ScriptVersionItemMigrationsWorkersSingleStepMigrationsRenamedClassParam] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]ScriptVersionItemMigrationsWorkersSingleStepMigrationsTransferredClassParam] `json:"transferred_classes"`
}

func (r ScriptVersionItemMigrationsWorkersSingleStepMigrationsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionItemMigrationsWorkersSingleStepMigrationsParam) implementsScriptVersionItemMigrationsUnionParam() {
}

type ScriptVersionItemMigrationsWorkersSingleStepMigrationsRenamedClassParam struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r ScriptVersionItemMigrationsWorkersSingleStepMigrationsRenamedClassParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionItemMigrationsWorkersSingleStepMigrationsTransferredClassParam struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r ScriptVersionItemMigrationsWorkersSingleStepMigrationsTransferredClassParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionItemMigrationsWorkersMultipleStepMigrationsParam struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]MigrationStepParam] `json:"steps"`
}

func (r ScriptVersionItemMigrationsWorkersMultipleStepMigrationsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionItemMigrationsWorkersMultipleStepMigrationsParam) implementsScriptVersionItemMigrationsUnionParam() {
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptVersionItemPlacementParam struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode param.Field[PlacementMode] `json:"mode"`
}

func (r ScriptVersionItemPlacementParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Usage model for the Worker invocations.
type UsageModel string

const (
	UsageModelStandard UsageModel = "standard"
)

func (r UsageModel) IsKnown() bool {
	switch r {
	case UsageModelStandard:
		return true
	}
	return false
}

type AccountWorkerDispatchNamespaceScriptSettingGetResponse struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountWorkerDispatchNamespaceScriptSettingGetResponseSuccess `json:"success,required"`
	Result  ScriptVersionItem                                             `json:"result"`
	JSON    accountWorkerDispatchNamespaceScriptSettingGetResponseJSON    `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseJSON contains the JSON
// metadata for the struct [AccountWorkerDispatchNamespaceScriptSettingGetResponse]
type accountWorkerDispatchNamespaceScriptSettingGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDispatchNamespaceScriptSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerDispatchNamespaceScriptSettingGetResponseSuccess bool

const (
	AccountWorkerDispatchNamespaceScriptSettingGetResponseSuccessTrue AccountWorkerDispatchNamespaceScriptSettingGetResponseSuccess = true
)

func (r AccountWorkerDispatchNamespaceScriptSettingGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerDispatchNamespaceScriptSettingGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerDispatchNamespaceScriptSettingPatchResponse struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountWorkerDispatchNamespaceScriptSettingPatchResponseSuccess `json:"success,required"`
	Result  ScriptVersionItem                                               `json:"result"`
	JSON    accountWorkerDispatchNamespaceScriptSettingPatchResponseJSON    `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingPatchResponseJSON contains the JSON
// metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingPatchResponse]
type accountWorkerDispatchNamespaceScriptSettingPatchResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingPatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDispatchNamespaceScriptSettingPatchResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerDispatchNamespaceScriptSettingPatchResponseSuccess bool

const (
	AccountWorkerDispatchNamespaceScriptSettingPatchResponseSuccessTrue AccountWorkerDispatchNamespaceScriptSettingPatchResponseSuccess = true
)

func (r AccountWorkerDispatchNamespaceScriptSettingPatchResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerDispatchNamespaceScriptSettingPatchResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerDispatchNamespaceScriptSettingPatchParams struct {
	Settings param.Field[ScriptVersionItemParam] `json:"settings"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingPatchParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
