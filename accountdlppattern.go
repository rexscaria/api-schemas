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

// AccountDlpPatternService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDlpPatternService] method instead.
type AccountDlpPatternService struct {
	Options []option.RequestOption
}

// NewAccountDlpPatternService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDlpPatternService(opts ...option.RequestOption) (r *AccountDlpPatternService) {
	r = &AccountDlpPatternService{}
	r.Options = opts
	return
}

// Validates whether this pattern is a valid regular expression. Rejects it if the
// regular expression is too complex or can match an unbounded-length string. The
// regex will be rejected if it uses `*` or `+`. Bound the maximum number of
// characters that can be matched using a range, e.g. `{1,100}`.
func (r *AccountDlpPatternService) Validate(ctx context.Context, accountID string, body AccountDlpPatternValidateParams, opts ...option.RequestOption) (res *AccountDlpPatternValidateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/patterns/validate", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountDlpPatternValidateResponse struct {
	Result AccountDlpPatternValidateResponseResult `json:"result"`
	JSON   accountDlpPatternValidateResponseJSON   `json:"-"`
	APIResponseSingleDlp
}

// accountDlpPatternValidateResponseJSON contains the JSON metadata for the struct
// [AccountDlpPatternValidateResponse]
type accountDlpPatternValidateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpPatternValidateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpPatternValidateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDlpPatternValidateResponseResult struct {
	Valid bool                                        `json:"valid,required"`
	JSON  accountDlpPatternValidateResponseResultJSON `json:"-"`
}

// accountDlpPatternValidateResponseResultJSON contains the JSON metadata for the
// struct [AccountDlpPatternValidateResponseResult]
type accountDlpPatternValidateResponseResultJSON struct {
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpPatternValidateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDlpPatternValidateResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDlpPatternValidateParams struct {
	Regex param.Field[string] `json:"regex,required"`
	// Maximum number of bytes that the regular expression can match.
	//
	// If this is `null` then there is no limit on the length. Patterns can use `*` and
	// `+`. Otherwise repeats should use a range `{m,n}` to restrict patterns to the
	// length. If this field is missing, then a default length limit is used.
	//
	// Note that the length is specified in bytes. Since regular expressions use UTF-8
	// the pattern `.` can match up to 4 bytes. Hence `.{1,256}` has a maximum length
	// of 1024 bytes.
	MaxMatchBytes param.Field[int64] `json:"max_match_bytes"`
}

func (r AccountDlpPatternValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
