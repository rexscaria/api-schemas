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

// AccountAITaskService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAITaskService] method instead.
type AccountAITaskService struct {
	Options []option.RequestOption
}

// NewAccountAITaskService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountAITaskService(opts ...option.RequestOption) (r *AccountAITaskService) {
	r = &AccountAITaskService{}
	r.Options = opts
	return
}

// Task Search
func (r *AccountAITaskService) Search(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAITaskSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/tasks/search", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAITaskSearchResponse struct {
	Errors   []interface{}                   `json:"errors,required"`
	Messages []string                        `json:"messages,required"`
	Result   []interface{}                   `json:"result,required"`
	Success  bool                            `json:"success,required"`
	JSON     accountAITaskSearchResponseJSON `json:"-"`
}

// accountAITaskSearchResponseJSON contains the JSON metadata for the struct
// [AccountAITaskSearchResponse]
type accountAITaskSearchResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAITaskSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAITaskSearchResponseJSON) RawJSON() string {
	return r.raw
}
