// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountWorkerScriptContentService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerScriptContentService] method instead.
type AccountWorkerScriptContentService struct {
	Options []option.RequestOption
}

// NewAccountWorkerScriptContentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerScriptContentService(opts ...option.RequestOption) (r *AccountWorkerScriptContentService) {
	r = &AccountWorkerScriptContentService{}
	r.Options = opts
	return
}

// Fetch script content only
func (r *AccountWorkerScriptContentService) GetV2(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "string")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/content/v2", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Put script content without touching config or metadata
func (r *AccountWorkerScriptContentService) Put(ctx context.Context, accountID string, scriptName string, params AccountWorkerScriptContentPutParams, opts ...option.RequestOption) (res *SingleScriptResponse, err error) {
	if params.CfWorkerBodyPart.Present {
		opts = append(opts, option.WithHeader("CF-WORKER-BODY-PART", fmt.Sprintf("%s", params.CfWorkerBodyPart)))
	}
	if params.CfWorkerMainModulePart.Present {
		opts = append(opts, option.WithHeader("CF-WORKER-MAIN-MODULE-PART", fmt.Sprintf("%s", params.CfWorkerMainModulePart)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/content", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type AccountWorkerScriptContentPutParams struct {
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata               param.Field[AccountWorkerScriptContentPutParamsMetadata] `json:"metadata,required"`
	CfWorkerBodyPart       param.Field[string]                                      `header:"CF-WORKER-BODY-PART"`
	CfWorkerMainModulePart param.Field[string]                                      `header:"CF-WORKER-MAIN-MODULE-PART"`
}

func (r AccountWorkerScriptContentPutParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// JSON encoded metadata about the uploaded parts and Worker configuration.
type AccountWorkerScriptContentPutParamsMetadata struct {
	// Name of the part in the multipart request that contains the script (e.g. the
	// file adding a listener to the `fetch` event). Indicates a
	// `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
}

func (r AccountWorkerScriptContentPutParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
