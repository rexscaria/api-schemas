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

// AccountCloudforceOneScanService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCloudforceOneScanService] method instead.
type AccountCloudforceOneScanService struct {
	Options []option.RequestOption
	Config  *AccountCloudforceOneScanConfigService
}

// NewAccountCloudforceOneScanService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCloudforceOneScanService(opts ...option.RequestOption) (r *AccountCloudforceOneScanService) {
	r = &AccountCloudforceOneScanService{}
	r.Options = opts
	r.Config = NewAccountCloudforceOneScanConfigService(opts...)
	return
}

// Get the Latest Scan Result
func (r *AccountCloudforceOneScanService) GetResults(ctx context.Context, accountID string, configID string, opts ...option.RequestOption) (res *AccountCloudforceOneScanGetResultsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/scans/results/%s", accountID, configID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountCloudforceOneScanGetResultsResponse struct {
	Errors   []string                                         `json:"errors,required"`
	Messages []string                                         `json:"messages,required"`
	Result   AccountCloudforceOneScanGetResultsResponseResult `json:"result,required"`
	Success  bool                                             `json:"success,required"`
	JSON     accountCloudforceOneScanGetResultsResponseJSON   `json:"-"`
}

// accountCloudforceOneScanGetResultsResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneScanGetResultsResponse]
type accountCloudforceOneScanGetResultsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneScanGetResultsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneScanGetResultsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneScanGetResultsResponseResult struct {
	OneOneOneOne []AccountCloudforceOneScanGetResultsResponseResult1_1_1_1 `json:"1.1.1.1,required"`
	JSON         accountCloudforceOneScanGetResultsResponseResultJSON      `json:"-"`
}

// accountCloudforceOneScanGetResultsResponseResultJSON contains the JSON metadata
// for the struct [AccountCloudforceOneScanGetResultsResponseResult]
type accountCloudforceOneScanGetResultsResponseResultJSON struct {
	OneOneOneOne apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountCloudforceOneScanGetResultsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneScanGetResultsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneScanGetResultsResponseResult1_1_1_1 struct {
	Number float64                                                     `json:"number"`
	Proto  string                                                      `json:"proto"`
	Status string                                                      `json:"status"`
	JSON   accountCloudforceOneScanGetResultsResponseResult1_1_1_1JSON `json:"-"`
}

// accountCloudforceOneScanGetResultsResponseResult1_1_1_1JSON contains the JSON
// metadata for the struct
// [AccountCloudforceOneScanGetResultsResponseResult1_1_1_1]
type accountCloudforceOneScanGetResultsResponseResult1_1_1_1JSON struct {
	Number      apijson.Field
	Proto       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneScanGetResultsResponseResult1_1_1_1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneScanGetResultsResponseResult1_1_1_1JSON) RawJSON() string {
	return r.raw
}
