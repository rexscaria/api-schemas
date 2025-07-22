// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountUrlscannerV2ScreenshotService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountUrlscannerV2ScreenshotService] method instead.
type AccountUrlscannerV2ScreenshotService struct {
	Options []option.RequestOption
}

// NewAccountUrlscannerV2ScreenshotService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountUrlscannerV2ScreenshotService(opts ...option.RequestOption) (r *AccountUrlscannerV2ScreenshotService) {
	r = &AccountUrlscannerV2ScreenshotService{}
	r.Options = opts
	return
}

// Get scan's screenshot by resolution (desktop/mobile/tablet).
func (r *AccountUrlscannerV2ScreenshotService) GetScreenshot(ctx context.Context, accountID string, scanID string, query AccountUrlscannerV2ScreenshotGetScreenshotParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "image/png")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/screenshots/%s.png", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountUrlscannerV2ScreenshotGetScreenshotParams struct {
	// Target device type.
	Resolution param.Field[AccountUrlscannerV2ScreenshotGetScreenshotParamsResolution] `query:"resolution"`
}

// URLQuery serializes [AccountUrlscannerV2ScreenshotGetScreenshotParams]'s query
// parameters as `url.Values`.
func (r AccountUrlscannerV2ScreenshotGetScreenshotParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Target device type.
type AccountUrlscannerV2ScreenshotGetScreenshotParamsResolution string

const (
	AccountUrlscannerV2ScreenshotGetScreenshotParamsResolutionDesktop AccountUrlscannerV2ScreenshotGetScreenshotParamsResolution = "desktop"
	AccountUrlscannerV2ScreenshotGetScreenshotParamsResolutionMobile  AccountUrlscannerV2ScreenshotGetScreenshotParamsResolution = "mobile"
	AccountUrlscannerV2ScreenshotGetScreenshotParamsResolutionTablet  AccountUrlscannerV2ScreenshotGetScreenshotParamsResolution = "tablet"
)

func (r AccountUrlscannerV2ScreenshotGetScreenshotParamsResolution) IsKnown() bool {
	switch r {
	case AccountUrlscannerV2ScreenshotGetScreenshotParamsResolutionDesktop, AccountUrlscannerV2ScreenshotGetScreenshotParamsResolutionMobile, AccountUrlscannerV2ScreenshotGetScreenshotParamsResolutionTablet:
		return true
	}
	return false
}
