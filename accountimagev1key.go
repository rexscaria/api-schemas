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

// AccountImageV1KeyService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountImageV1KeyService] method instead.
type AccountImageV1KeyService struct {
	Options []option.RequestOption
}

// NewAccountImageV1KeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountImageV1KeyService(opts ...option.RequestOption) (r *AccountImageV1KeyService) {
	r = &AccountImageV1KeyService{}
	r.Options = opts
	return
}

// Create a new signing key with specified name. Returns all keys available.
func (r *AccountImageV1KeyService) New(ctx context.Context, accountID string, signingKeyName string, opts ...option.RequestOption) (res *ImageKeyResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if signingKeyName == "" {
		err = errors.New("missing required signing_key_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/keys/%s", accountID, signingKeyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Lists your signing keys. These can be found on your Cloudflare Images dashboard.
func (r *AccountImageV1KeyService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *ImageKeyResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/keys", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete signing key with specified name. Returns all keys available. When last
// key is removed, a new default signing key will be generated.
func (r *AccountImageV1KeyService) Delete(ctx context.Context, accountID string, signingKeyName string, opts ...option.RequestOption) (res *ImageKeyResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if signingKeyName == "" {
		err = errors.New("missing required signing_key_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/keys/%s", accountID, signingKeyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ImageKeyResponseCollection struct {
	Errors   []ImageKeyResponseCollectionError   `json:"errors,required"`
	Messages []ImageKeyResponseCollectionMessage `json:"messages,required"`
	Result   ImageKeyResponseCollectionResult    `json:"result,required"`
	// Whether the API call was successful
	Success ImageKeyResponseCollectionSuccess `json:"success,required"`
	JSON    imageKeyResponseCollectionJSON    `json:"-"`
}

// imageKeyResponseCollectionJSON contains the JSON metadata for the struct
// [ImageKeyResponseCollection]
type imageKeyResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageKeyResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageKeyResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type ImageKeyResponseCollectionError struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           ImageKeyResponseCollectionErrorsSource `json:"source"`
	JSON             imageKeyResponseCollectionErrorJSON    `json:"-"`
}

// imageKeyResponseCollectionErrorJSON contains the JSON metadata for the struct
// [ImageKeyResponseCollectionError]
type imageKeyResponseCollectionErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ImageKeyResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageKeyResponseCollectionErrorJSON) RawJSON() string {
	return r.raw
}

type ImageKeyResponseCollectionErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    imageKeyResponseCollectionErrorsSourceJSON `json:"-"`
}

// imageKeyResponseCollectionErrorsSourceJSON contains the JSON metadata for the
// struct [ImageKeyResponseCollectionErrorsSource]
type imageKeyResponseCollectionErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageKeyResponseCollectionErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageKeyResponseCollectionErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ImageKeyResponseCollectionMessage struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           ImageKeyResponseCollectionMessagesSource `json:"source"`
	JSON             imageKeyResponseCollectionMessageJSON    `json:"-"`
}

// imageKeyResponseCollectionMessageJSON contains the JSON metadata for the struct
// [ImageKeyResponseCollectionMessage]
type imageKeyResponseCollectionMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ImageKeyResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageKeyResponseCollectionMessageJSON) RawJSON() string {
	return r.raw
}

type ImageKeyResponseCollectionMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    imageKeyResponseCollectionMessagesSourceJSON `json:"-"`
}

// imageKeyResponseCollectionMessagesSourceJSON contains the JSON metadata for the
// struct [ImageKeyResponseCollectionMessagesSource]
type imageKeyResponseCollectionMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageKeyResponseCollectionMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageKeyResponseCollectionMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type ImageKeyResponseCollectionResult struct {
	Keys []ImageKeyResponseCollectionResultKey `json:"keys"`
	JSON imageKeyResponseCollectionResultJSON  `json:"-"`
}

// imageKeyResponseCollectionResultJSON contains the JSON metadata for the struct
// [ImageKeyResponseCollectionResult]
type imageKeyResponseCollectionResultJSON struct {
	Keys        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageKeyResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageKeyResponseCollectionResultJSON) RawJSON() string {
	return r.raw
}

type ImageKeyResponseCollectionResultKey struct {
	// Key name.
	Name string `json:"name"`
	// Key value.
	Value string                                  `json:"value"`
	JSON  imageKeyResponseCollectionResultKeyJSON `json:"-"`
}

// imageKeyResponseCollectionResultKeyJSON contains the JSON metadata for the
// struct [ImageKeyResponseCollectionResultKey]
type imageKeyResponseCollectionResultKeyJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageKeyResponseCollectionResultKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageKeyResponseCollectionResultKeyJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ImageKeyResponseCollectionSuccess bool

const (
	ImageKeyResponseCollectionSuccessTrue ImageKeyResponseCollectionSuccess = true
)

func (r ImageKeyResponseCollectionSuccess) IsKnown() bool {
	switch r {
	case ImageKeyResponseCollectionSuccessTrue:
		return true
	}
	return false
}
