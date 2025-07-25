// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountWorkerServiceEnvironmentContentService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerServiceEnvironmentContentService] method instead.
type AccountWorkerServiceEnvironmentContentService struct {
	Options []option.RequestOption
}

// NewAccountWorkerServiceEnvironmentContentService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountWorkerServiceEnvironmentContentService(opts ...option.RequestOption) (r *AccountWorkerServiceEnvironmentContentService) {
	r = &AccountWorkerServiceEnvironmentContentService{}
	r.Options = opts
	return
}

// Get script content from a worker with an environment.
func (r *AccountWorkerServiceEnvironmentContentService) Get(ctx context.Context, accountID string, serviceName string, environmentName string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "string")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if serviceName == "" {
		err = errors.New("missing required service_name parameter")
		return
	}
	if environmentName == "" {
		err = errors.New("missing required environment_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/content", accountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Put script content from a worker with an environment.
func (r *AccountWorkerServiceEnvironmentContentService) Put(ctx context.Context, accountID string, serviceName string, environmentName string, params AccountWorkerServiceEnvironmentContentPutParams, opts ...option.RequestOption) (res *SingleScriptResponse, err error) {
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
	if serviceName == "" {
		err = errors.New("missing required service_name parameter")
		return
	}
	if environmentName == "" {
		err = errors.New("missing required environment_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/content", accountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type AccountWorkerServiceEnvironmentContentPutParams struct {
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[AccountWorkerServiceEnvironmentContentPutParamsMetadata] `json:"metadata,required"`
	// An array of modules (often JavaScript files) comprising a Worker script. At
	// least one module must be present and referenced in the metadata as `main_module`
	// or `body_part` by filename.<br/>Possible Content-Type(s) are:
	// `application/javascript+module`, `text/javascript+module`,
	// `application/javascript`, `text/javascript`, `application/wasm`, `text/plain`,
	// `application/octet-stream`, `application/source-map`.
	Files                  param.Field[[]io.Reader] `json:"files" format:"binary"`
	CfWorkerBodyPart       param.Field[string]      `header:"CF-WORKER-BODY-PART"`
	CfWorkerMainModulePart param.Field[string]      `header:"CF-WORKER-MAIN-MODULE-PART"`
}

func (r AccountWorkerServiceEnvironmentContentPutParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type AccountWorkerServiceEnvironmentContentPutParamsMetadata struct {
	// Name of the part in the multipart request that contains the script (e.g. the
	// file adding a listener to the `fetch` event). Indicates a
	// `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
}

func (r AccountWorkerServiceEnvironmentContentPutParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
