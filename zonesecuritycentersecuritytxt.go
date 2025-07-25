// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneSecurityCenterSecuritytxtService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSecurityCenterSecuritytxtService] method instead.
type ZoneSecurityCenterSecuritytxtService struct {
	Options []option.RequestOption
}

// NewZoneSecurityCenterSecuritytxtService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSecurityCenterSecuritytxtService(opts ...option.RequestOption) (r *ZoneSecurityCenterSecuritytxtService) {
	r = &ZoneSecurityCenterSecuritytxtService{}
	r.Options = opts
	return
}

// Get security.txt
func (r *ZoneSecurityCenterSecuritytxtService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneSecurityCenterSecuritytxtGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/security-center/securitytxt", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update security.txt
func (r *ZoneSecurityCenterSecuritytxtService) Update(ctx context.Context, zoneID string, body ZoneSecurityCenterSecuritytxtUpdateParams, opts ...option.RequestOption) (res *SingleResponseReport, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/security-center/securitytxt", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete security.txt
func (r *ZoneSecurityCenterSecuritytxtService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SingleResponseReport, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/security-center/securitytxt", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type SecurityTxt struct {
	Acknowledgments    []string        `json:"acknowledgments" format:"uri"`
	Canonical          []string        `json:"canonical" format:"uri"`
	Contact            []string        `json:"contact" format:"uri"`
	Enabled            bool            `json:"enabled"`
	Encryption         []string        `json:"encryption" format:"uri"`
	Expires            time.Time       `json:"expires" format:"date-time"`
	Hiring             []string        `json:"hiring" format:"uri"`
	Policy             []string        `json:"policy" format:"uri"`
	PreferredLanguages string          `json:"preferredLanguages"`
	JSON               securityTxtJSON `json:"-"`
}

// securityTxtJSON contains the JSON metadata for the struct [SecurityTxt]
type securityTxtJSON struct {
	Acknowledgments    apijson.Field
	Canonical          apijson.Field
	Contact            apijson.Field
	Enabled            apijson.Field
	Encryption         apijson.Field
	Expires            apijson.Field
	Hiring             apijson.Field
	Policy             apijson.Field
	PreferredLanguages apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SecurityTxt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityTxtJSON) RawJSON() string {
	return r.raw
}

type SecurityTxtParam struct {
	Acknowledgments    param.Field[[]string]  `json:"acknowledgments" format:"uri"`
	Canonical          param.Field[[]string]  `json:"canonical" format:"uri"`
	Contact            param.Field[[]string]  `json:"contact" format:"uri"`
	Enabled            param.Field[bool]      `json:"enabled"`
	Encryption         param.Field[[]string]  `json:"encryption" format:"uri"`
	Expires            param.Field[time.Time] `json:"expires" format:"date-time"`
	Hiring             param.Field[[]string]  `json:"hiring" format:"uri"`
	Policy             param.Field[[]string]  `json:"policy" format:"uri"`
	PreferredLanguages param.Field[string]    `json:"preferredLanguages"`
}

func (r SecurityTxtParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSecurityCenterSecuritytxtGetResponse struct {
	Errors   []AttackSurfaceReportMessage `json:"errors,required"`
	Messages []AttackSurfaceReportMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneSecurityCenterSecuritytxtGetResponseSuccess `json:"success,required"`
	Result  SecurityTxt                                     `json:"result"`
	JSON    zoneSecurityCenterSecuritytxtGetResponseJSON    `json:"-"`
}

// zoneSecurityCenterSecuritytxtGetResponseJSON contains the JSON metadata for the
// struct [ZoneSecurityCenterSecuritytxtGetResponse]
type zoneSecurityCenterSecuritytxtGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecurityCenterSecuritytxtGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSecurityCenterSecuritytxtGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneSecurityCenterSecuritytxtGetResponseSuccess bool

const (
	ZoneSecurityCenterSecuritytxtGetResponseSuccessTrue ZoneSecurityCenterSecuritytxtGetResponseSuccess = true
)

func (r ZoneSecurityCenterSecuritytxtGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSecurityCenterSecuritytxtGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneSecurityCenterSecuritytxtUpdateParams struct {
	SecurityTxt SecurityTxtParam `json:"security_txt,required"`
}

func (r ZoneSecurityCenterSecuritytxtUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SecurityTxt)
}
