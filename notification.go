// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// NotificationService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationService] method instead.
type NotificationService struct {
	Options []option.RequestOption
}

// NewNotificationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNotificationService(opts ...option.RequestOption) (r *NotificationService) {
	r = &NotificationService{}
	r.Options = opts
	return
}

// Publish to ANS
func (r *NotificationService) Publish(ctx context.Context, opts ...option.RequestOption) (res *NotificationPublishResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "notification"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type NotificationPublishResponse struct {
	DefaultAlert NotificationPublishResponseDefaultAlert `json:"defaultAlert,required"`
	Message      string                                  `json:"message,required"`
	JSON         notificationPublishResponseJSON         `json:"-"`
}

// notificationPublishResponseJSON contains the JSON metadata for the struct
// [NotificationPublishResponse]
type notificationPublishResponseJSON struct {
	DefaultAlert apijson.Field
	Message      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NotificationPublishResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notificationPublishResponseJSON) RawJSON() string {
	return r.raw
}

type NotificationPublishResponseDefaultAlert struct {
	TypeName     string                                              `json:"$typeName,required"`
	AccountID    float64                                             `json:"accountId,required"`
	AccountTag   string                                              `json:"accountTag,required"`
	AlertContext NotificationPublishResponseDefaultAlertAlertContext `json:"alertContext,required"`
	AlertType    float64                                             `json:"alertType,required"`
	PolicyID     string                                              `json:"policyId,required"`
	Severity     float64                                             `json:"severity,required"`
	Source       string                                              `json:"source,required"`
	ZoneID       float64                                             `json:"zoneId,required"`
	ZoneTag      string                                              `json:"zoneTag,required"`
	JSON         notificationPublishResponseDefaultAlertJSON         `json:"-"`
}

// notificationPublishResponseDefaultAlertJSON contains the JSON metadata for the
// struct [NotificationPublishResponseDefaultAlert]
type notificationPublishResponseDefaultAlertJSON struct {
	TypeName     apijson.Field
	AccountID    apijson.Field
	AccountTag   apijson.Field
	AlertContext apijson.Field
	AlertType    apijson.Field
	PolicyID     apijson.Field
	Severity     apijson.Field
	Source       apijson.Field
	ZoneID       apijson.Field
	ZoneTag      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NotificationPublishResponseDefaultAlert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notificationPublishResponseDefaultAlertJSON) RawJSON() string {
	return r.raw
}

type NotificationPublishResponseDefaultAlertAlertContext struct {
	Case  string                                                   `json:"case,required"`
	Value NotificationPublishResponseDefaultAlertAlertContextValue `json:"value,required"`
	JSON  notificationPublishResponseDefaultAlertAlertContextJSON  `json:"-"`
}

// notificationPublishResponseDefaultAlertAlertContextJSON contains the JSON
// metadata for the struct [NotificationPublishResponseDefaultAlertAlertContext]
type notificationPublishResponseDefaultAlertAlertContextJSON struct {
	Case        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotificationPublishResponseDefaultAlertAlertContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notificationPublishResponseDefaultAlertAlertContextJSON) RawJSON() string {
	return r.raw
}

type NotificationPublishResponseDefaultAlertAlertContextValue struct {
	TypeName   string                                                         `json:"$typeName,required"`
	AccountTag string                                                         `json:"accountTag,required"`
	Ports      []NotificationPublishResponseDefaultAlertAlertContextValuePort `json:"ports,required"`
	JSON       notificationPublishResponseDefaultAlertAlertContextValueJSON   `json:"-"`
}

// notificationPublishResponseDefaultAlertAlertContextValueJSON contains the JSON
// metadata for the struct
// [NotificationPublishResponseDefaultAlertAlertContextValue]
type notificationPublishResponseDefaultAlertAlertContextValueJSON struct {
	TypeName    apijson.Field
	AccountTag  apijson.Field
	Ports       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotificationPublishResponseDefaultAlertAlertContextValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notificationPublishResponseDefaultAlertAlertContextValueJSON) RawJSON() string {
	return r.raw
}

type NotificationPublishResponseDefaultAlertAlertContextValuePort struct {
	TypeName string                                                           `json:"$typeName,required"`
	IP       string                                                           `json:"ip,required"`
	Number   float64                                                          `json:"number,required"`
	Status   float64                                                          `json:"status,required"`
	JSON     notificationPublishResponseDefaultAlertAlertContextValuePortJSON `json:"-"`
}

// notificationPublishResponseDefaultAlertAlertContextValuePortJSON contains the
// JSON metadata for the struct
// [NotificationPublishResponseDefaultAlertAlertContextValuePort]
type notificationPublishResponseDefaultAlertAlertContextValuePortJSON struct {
	TypeName    apijson.Field
	IP          apijson.Field
	Number      apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotificationPublishResponseDefaultAlertAlertContextValuePort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notificationPublishResponseDefaultAlertAlertContextValuePortJSON) RawJSON() string {
	return r.raw
}
