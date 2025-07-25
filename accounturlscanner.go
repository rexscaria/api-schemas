// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountUrlscannerService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountUrlscannerService] method instead.
type AccountUrlscannerService struct {
	Options []option.RequestOption
	Scan    *AccountUrlscannerScanService
	V2      *AccountUrlscannerV2Service
}

// NewAccountUrlscannerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountUrlscannerService(opts ...option.RequestOption) (r *AccountUrlscannerService) {
	r = &AccountUrlscannerService{}
	r.Options = opts
	r.Scan = NewAccountUrlscannerScanService(opts...)
	r.V2 = NewAccountUrlscannerV2Service(opts...)
	return
}

// Returns the plain response of the network request.
//
// Deprecated: Use
// [V2](https://developers.cloudflare.com/api/resources/url_scanner/subresources/responses/methods/get/)
// instead.
func (r *AccountUrlscannerService) GetRawResponse(ctx context.Context, accountID string, responseID string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/response/%s", accountID, responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
