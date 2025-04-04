// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountLogpushDatasetService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountLogpushDatasetService] method instead.
type AccountLogpushDatasetService struct {
	Options []option.RequestOption
}

// NewAccountLogpushDatasetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountLogpushDatasetService(opts ...option.RequestOption) (r *AccountLogpushDatasetService) {
	r = &AccountLogpushDatasetService{}
	r.Options = opts
	return
}

// Lists all fields available for a dataset. The response result is an object with
// key-value pairs, where keys are field names, and values are descriptions.
func (r *AccountLogpushDatasetService) ListFields(ctx context.Context, accountID string, datasetID string, opts ...option.RequestOption) (res *FieldResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/logpush/datasets/%s/fields", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists Logpush jobs for an account for a dataset.
func (r *AccountLogpushDatasetService) ListJobs(ctx context.Context, accountID string, datasetID string, opts ...option.RequestOption) (res *JobResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/logpush/datasets/%s/jobs", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type FieldResponseCollection struct {
	Result interface{}                 `json:"result"`
	JSON   fieldResponseCollectionJSON `json:"-"`
	CommonResponseLogPush
}

// fieldResponseCollectionJSON contains the JSON metadata for the struct
// [FieldResponseCollection]
type fieldResponseCollectionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type JobResponseCollection struct {
	Result []LogpushJob              `json:"result"`
	JSON   jobResponseCollectionJSON `json:"-"`
	CommonResponseLogPush
}

// jobResponseCollectionJSON contains the JSON metadata for the struct
// [JobResponseCollection]
type jobResponseCollectionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobResponseCollectionJSON) RawJSON() string {
	return r.raw
}
