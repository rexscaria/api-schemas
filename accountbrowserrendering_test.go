// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/rexscaria/api-schemas"
	"github.com/rexscaria/api-schemas/internal/testutil"
	"github.com/rexscaria/api-schemas/option"
)

func TestAccountBrowserRenderingGetHTMLContentWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cfrex.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIEmail("My API Email"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Accounts.BrowserRendering.GetHTMLContent(
		context.TODO(),
		"account_id",
		cfrex.AccountBrowserRenderingGetHTMLContentParams{
			CacheTtl: cfrex.F(86400.000000),
			AddScriptTag: cfrex.F([]cfrex.AccountBrowserRenderingGetHTMLContentParamsAddScriptTag{{
				ID:      cfrex.F("id"),
				Content: cfrex.F("content"),
				Type:    cfrex.F("type"),
				URL:     cfrex.F("url"),
			}}),
			AddStyleTag: cfrex.F([]cfrex.AccountBrowserRenderingGetHTMLContentParamsAddStyleTag{{
				Content: cfrex.F("content"),
				URL:     cfrex.F("url"),
			}}),
			AllowRequestPattern: cfrex.F([]string{"string"}),
			AllowResourceTypes:  cfrex.F([]cfrex.AccountBrowserRenderingGetHTMLContentParamsAllowResourceType{cfrex.AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeDocument}),
			Authenticate: cfrex.F(cfrex.AccountBrowserRenderingGetHTMLContentParamsAuthenticate{
				Password: cfrex.F("x"),
				Username: cfrex.F("x"),
			}),
			BestAttempt: cfrex.F(true),
			Cookies: cfrex.F([]cfrex.AccountBrowserRenderingGetHTMLContentParamsCookie{{
				Name:         cfrex.F("name"),
				Value:        cfrex.F("value"),
				Domain:       cfrex.F("domain"),
				Expires:      cfrex.F(0.000000),
				HTTPOnly:     cfrex.F(true),
				PartitionKey: cfrex.F("partitionKey"),
				Path:         cfrex.F("path"),
				Priority:     cfrex.F(cfrex.AccountBrowserRenderingGetHTMLContentParamsCookiesPriorityLow),
				SameParty:    cfrex.F(true),
				SameSite:     cfrex.F(cfrex.AccountBrowserRenderingGetHTMLContentParamsCookiesSameSiteStrict),
				Secure:       cfrex.F(true),
				SourcePort:   cfrex.F(0.000000),
				SourceScheme: cfrex.F(cfrex.AccountBrowserRenderingGetHTMLContentParamsCookiesSourceSchemeUnset),
				URL:          cfrex.F("url"),
			}}),
			EmulateMediaType: cfrex.F("emulateMediaType"),
			GotoOptions: cfrex.F(cfrex.AccountBrowserRenderingGetHTMLContentParamsGotoOptions{
				Referer:        cfrex.F("referer"),
				ReferrerPolicy: cfrex.F("referrerPolicy"),
				Timeout:        cfrex.F(60000.000000),
				WaitUntil:      cfrex.F[cfrex.AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilUnion](cfrex.AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilString(cfrex.AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilStringLoad)),
			}),
			HTML:                 cfrex.F("x"),
			RejectRequestPattern: cfrex.F([]string{"string"}),
			RejectResourceTypes:  cfrex.F([]cfrex.AccountBrowserRenderingGetHTMLContentParamsRejectResourceType{cfrex.AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeDocument}),
			SetExtraHTTPHeaders: cfrex.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cfrex.F(true),
			URL:                  cfrex.F("https://example.com"),
			UserAgent:            cfrex.F("userAgent"),
			Viewport: cfrex.F(cfrex.AccountBrowserRenderingGetHTMLContentParamsViewport{
				Height:            cfrex.F(0.000000),
				Width:             cfrex.F(0.000000),
				DeviceScaleFactor: cfrex.F(0.000000),
				HasTouch:          cfrex.F(true),
				IsLandscape:       cfrex.F(true),
				IsMobile:          cfrex.F(true),
			}),
			WaitForSelector: cfrex.F(cfrex.AccountBrowserRenderingGetHTMLContentParamsWaitForSelector{
				Selector: cfrex.F("selector"),
				Hidden:   cfrex.F(cfrex.AccountBrowserRenderingGetHTMLContentParamsWaitForSelectorHiddenTrue),
				Timeout:  cfrex.F(60000.000000),
				Visible:  cfrex.F(cfrex.AccountBrowserRenderingGetHTMLContentParamsWaitForSelectorVisibleTrue),
			}),
			WaitForTimeout: cfrex.F(60000.000000),
		},
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountBrowserRenderingGetJsonWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cfrex.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIEmail("My API Email"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Accounts.BrowserRendering.GetJson(
		context.TODO(),
		"account_id",
		cfrex.AccountBrowserRenderingGetJsonParams{
			CacheTtl: cfrex.F(86400.000000),
			AddScriptTag: cfrex.F([]cfrex.AccountBrowserRenderingGetJsonParamsAddScriptTag{{
				ID:      cfrex.F("id"),
				Content: cfrex.F("content"),
				Type:    cfrex.F("type"),
				URL:     cfrex.F("url"),
			}}),
			AddStyleTag: cfrex.F([]cfrex.AccountBrowserRenderingGetJsonParamsAddStyleTag{{
				Content: cfrex.F("content"),
				URL:     cfrex.F("url"),
			}}),
			AllowRequestPattern: cfrex.F([]string{"string"}),
			AllowResourceTypes:  cfrex.F([]cfrex.AccountBrowserRenderingGetJsonParamsAllowResourceType{cfrex.AccountBrowserRenderingGetJsonParamsAllowResourceTypeDocument}),
			Authenticate: cfrex.F(cfrex.AccountBrowserRenderingGetJsonParamsAuthenticate{
				Password: cfrex.F("x"),
				Username: cfrex.F("x"),
			}),
			BestAttempt: cfrex.F(true),
			Cookies: cfrex.F([]cfrex.AccountBrowserRenderingGetJsonParamsCookie{{
				Name:         cfrex.F("name"),
				Value:        cfrex.F("value"),
				Domain:       cfrex.F("domain"),
				Expires:      cfrex.F(0.000000),
				HTTPOnly:     cfrex.F(true),
				PartitionKey: cfrex.F("partitionKey"),
				Path:         cfrex.F("path"),
				Priority:     cfrex.F(cfrex.AccountBrowserRenderingGetJsonParamsCookiesPriorityLow),
				SameParty:    cfrex.F(true),
				SameSite:     cfrex.F(cfrex.AccountBrowserRenderingGetJsonParamsCookiesSameSiteStrict),
				Secure:       cfrex.F(true),
				SourcePort:   cfrex.F(0.000000),
				SourceScheme: cfrex.F(cfrex.AccountBrowserRenderingGetJsonParamsCookiesSourceSchemeUnset),
				URL:          cfrex.F("url"),
			}}),
			EmulateMediaType: cfrex.F("emulateMediaType"),
			GotoOptions: cfrex.F(cfrex.AccountBrowserRenderingGetJsonParamsGotoOptions{
				Referer:        cfrex.F("referer"),
				ReferrerPolicy: cfrex.F("referrerPolicy"),
				Timeout:        cfrex.F(60000.000000),
				WaitUntil:      cfrex.F[cfrex.AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilUnion](cfrex.AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilString(cfrex.AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilStringLoad)),
			}),
			HTML:                 cfrex.F("x"),
			Prompt:               cfrex.F("prompt"),
			RejectRequestPattern: cfrex.F([]string{"string"}),
			RejectResourceTypes:  cfrex.F([]cfrex.AccountBrowserRenderingGetJsonParamsRejectResourceType{cfrex.AccountBrowserRenderingGetJsonParamsRejectResourceTypeDocument}),
			ResponseFormat: cfrex.F(cfrex.AccountBrowserRenderingGetJsonParamsResponseFormat{
				Type: cfrex.F("type"),
				Schema: cfrex.F(map[string]interface{}{
					"foo": map[string]interface{}{},
				}),
			}),
			SetExtraHTTPHeaders: cfrex.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cfrex.F(true),
			URL:                  cfrex.F("https://example.com"),
			UserAgent:            cfrex.F("userAgent"),
			Viewport: cfrex.F(cfrex.AccountBrowserRenderingGetJsonParamsViewport{
				Height:            cfrex.F(0.000000),
				Width:             cfrex.F(0.000000),
				DeviceScaleFactor: cfrex.F(0.000000),
				HasTouch:          cfrex.F(true),
				IsLandscape:       cfrex.F(true),
				IsMobile:          cfrex.F(true),
			}),
			WaitForSelector: cfrex.F(cfrex.AccountBrowserRenderingGetJsonParamsWaitForSelector{
				Selector: cfrex.F("selector"),
				Hidden:   cfrex.F(cfrex.AccountBrowserRenderingGetJsonParamsWaitForSelectorHiddenTrue),
				Timeout:  cfrex.F(60000.000000),
				Visible:  cfrex.F(cfrex.AccountBrowserRenderingGetJsonParamsWaitForSelectorVisibleTrue),
			}),
			WaitForTimeout: cfrex.F(60000.000000),
		},
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountBrowserRenderingGetLinksWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cfrex.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIEmail("My API Email"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Accounts.BrowserRendering.GetLinks(
		context.TODO(),
		"account_id",
		cfrex.AccountBrowserRenderingGetLinksParams{
			CacheTtl: cfrex.F(86400.000000),
			AddScriptTag: cfrex.F([]cfrex.AccountBrowserRenderingGetLinksParamsAddScriptTag{{
				ID:      cfrex.F("id"),
				Content: cfrex.F("content"),
				Type:    cfrex.F("type"),
				URL:     cfrex.F("url"),
			}}),
			AddStyleTag: cfrex.F([]cfrex.AccountBrowserRenderingGetLinksParamsAddStyleTag{{
				Content: cfrex.F("content"),
				URL:     cfrex.F("url"),
			}}),
			AllowRequestPattern: cfrex.F([]string{"string"}),
			AllowResourceTypes:  cfrex.F([]cfrex.AccountBrowserRenderingGetLinksParamsAllowResourceType{cfrex.AccountBrowserRenderingGetLinksParamsAllowResourceTypeDocument}),
			Authenticate: cfrex.F(cfrex.AccountBrowserRenderingGetLinksParamsAuthenticate{
				Password: cfrex.F("x"),
				Username: cfrex.F("x"),
			}),
			BestAttempt: cfrex.F(true),
			Cookies: cfrex.F([]cfrex.AccountBrowserRenderingGetLinksParamsCookie{{
				Name:         cfrex.F("name"),
				Value:        cfrex.F("value"),
				Domain:       cfrex.F("domain"),
				Expires:      cfrex.F(0.000000),
				HTTPOnly:     cfrex.F(true),
				PartitionKey: cfrex.F("partitionKey"),
				Path:         cfrex.F("path"),
				Priority:     cfrex.F(cfrex.AccountBrowserRenderingGetLinksParamsCookiesPriorityLow),
				SameParty:    cfrex.F(true),
				SameSite:     cfrex.F(cfrex.AccountBrowserRenderingGetLinksParamsCookiesSameSiteStrict),
				Secure:       cfrex.F(true),
				SourcePort:   cfrex.F(0.000000),
				SourceScheme: cfrex.F(cfrex.AccountBrowserRenderingGetLinksParamsCookiesSourceSchemeUnset),
				URL:          cfrex.F("url"),
			}}),
			EmulateMediaType: cfrex.F("emulateMediaType"),
			GotoOptions: cfrex.F(cfrex.AccountBrowserRenderingGetLinksParamsGotoOptions{
				Referer:        cfrex.F("referer"),
				ReferrerPolicy: cfrex.F("referrerPolicy"),
				Timeout:        cfrex.F(60000.000000),
				WaitUntil:      cfrex.F[cfrex.AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilUnion](cfrex.AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilString(cfrex.AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilStringLoad)),
			}),
			HTML:                 cfrex.F("x"),
			RejectRequestPattern: cfrex.F([]string{"string"}),
			RejectResourceTypes:  cfrex.F([]cfrex.AccountBrowserRenderingGetLinksParamsRejectResourceType{cfrex.AccountBrowserRenderingGetLinksParamsRejectResourceTypeDocument}),
			SetExtraHTTPHeaders: cfrex.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cfrex.F(true),
			URL:                  cfrex.F("https://example.com"),
			UserAgent:            cfrex.F("userAgent"),
			Viewport: cfrex.F(cfrex.AccountBrowserRenderingGetLinksParamsViewport{
				Height:            cfrex.F(0.000000),
				Width:             cfrex.F(0.000000),
				DeviceScaleFactor: cfrex.F(0.000000),
				HasTouch:          cfrex.F(true),
				IsLandscape:       cfrex.F(true),
				IsMobile:          cfrex.F(true),
			}),
			VisibleLinksOnly: cfrex.F(true),
			WaitForSelector: cfrex.F(cfrex.AccountBrowserRenderingGetLinksParamsWaitForSelector{
				Selector: cfrex.F("selector"),
				Hidden:   cfrex.F(cfrex.AccountBrowserRenderingGetLinksParamsWaitForSelectorHiddenTrue),
				Timeout:  cfrex.F(60000.000000),
				Visible:  cfrex.F(cfrex.AccountBrowserRenderingGetLinksParamsWaitForSelectorVisibleTrue),
			}),
			WaitForTimeout: cfrex.F(60000.000000),
		},
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountBrowserRenderingGetMarkdownWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cfrex.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIEmail("My API Email"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Accounts.BrowserRendering.GetMarkdown(
		context.TODO(),
		"account_id",
		cfrex.AccountBrowserRenderingGetMarkdownParams{
			CacheTtl: cfrex.F(86400.000000),
			AddScriptTag: cfrex.F([]cfrex.AccountBrowserRenderingGetMarkdownParamsAddScriptTag{{
				ID:      cfrex.F("id"),
				Content: cfrex.F("content"),
				Type:    cfrex.F("type"),
				URL:     cfrex.F("url"),
			}}),
			AddStyleTag: cfrex.F([]cfrex.AccountBrowserRenderingGetMarkdownParamsAddStyleTag{{
				Content: cfrex.F("content"),
				URL:     cfrex.F("url"),
			}}),
			AllowRequestPattern: cfrex.F([]string{"string"}),
			AllowResourceTypes:  cfrex.F([]cfrex.AccountBrowserRenderingGetMarkdownParamsAllowResourceType{cfrex.AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeDocument}),
			Authenticate: cfrex.F(cfrex.AccountBrowserRenderingGetMarkdownParamsAuthenticate{
				Password: cfrex.F("x"),
				Username: cfrex.F("x"),
			}),
			BestAttempt: cfrex.F(true),
			Cookies: cfrex.F([]cfrex.AccountBrowserRenderingGetMarkdownParamsCookie{{
				Name:         cfrex.F("name"),
				Value:        cfrex.F("value"),
				Domain:       cfrex.F("domain"),
				Expires:      cfrex.F(0.000000),
				HTTPOnly:     cfrex.F(true),
				PartitionKey: cfrex.F("partitionKey"),
				Path:         cfrex.F("path"),
				Priority:     cfrex.F(cfrex.AccountBrowserRenderingGetMarkdownParamsCookiesPriorityLow),
				SameParty:    cfrex.F(true),
				SameSite:     cfrex.F(cfrex.AccountBrowserRenderingGetMarkdownParamsCookiesSameSiteStrict),
				Secure:       cfrex.F(true),
				SourcePort:   cfrex.F(0.000000),
				SourceScheme: cfrex.F(cfrex.AccountBrowserRenderingGetMarkdownParamsCookiesSourceSchemeUnset),
				URL:          cfrex.F("url"),
			}}),
			EmulateMediaType: cfrex.F("emulateMediaType"),
			GotoOptions: cfrex.F(cfrex.AccountBrowserRenderingGetMarkdownParamsGotoOptions{
				Referer:        cfrex.F("referer"),
				ReferrerPolicy: cfrex.F("referrerPolicy"),
				Timeout:        cfrex.F(60000.000000),
				WaitUntil:      cfrex.F[cfrex.AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilUnion](cfrex.AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilString(cfrex.AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilStringLoad)),
			}),
			HTML:                 cfrex.F("x"),
			RejectRequestPattern: cfrex.F([]string{"string"}),
			RejectResourceTypes:  cfrex.F([]cfrex.AccountBrowserRenderingGetMarkdownParamsRejectResourceType{cfrex.AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeDocument}),
			SetExtraHTTPHeaders: cfrex.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cfrex.F(true),
			URL:                  cfrex.F("https://example.com"),
			UserAgent:            cfrex.F("userAgent"),
			Viewport: cfrex.F(cfrex.AccountBrowserRenderingGetMarkdownParamsViewport{
				Height:            cfrex.F(0.000000),
				Width:             cfrex.F(0.000000),
				DeviceScaleFactor: cfrex.F(0.000000),
				HasTouch:          cfrex.F(true),
				IsLandscape:       cfrex.F(true),
				IsMobile:          cfrex.F(true),
			}),
			WaitForSelector: cfrex.F(cfrex.AccountBrowserRenderingGetMarkdownParamsWaitForSelector{
				Selector: cfrex.F("selector"),
				Hidden:   cfrex.F(cfrex.AccountBrowserRenderingGetMarkdownParamsWaitForSelectorHiddenTrue),
				Timeout:  cfrex.F(60000.000000),
				Visible:  cfrex.F(cfrex.AccountBrowserRenderingGetMarkdownParamsWaitForSelectorVisibleTrue),
			}),
			WaitForTimeout: cfrex.F(60000.000000),
		},
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountBrowserRenderingGetPdfWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := cfrex.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIEmail("My API Email"),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.Accounts.BrowserRendering.GetPdf(
		context.TODO(),
		"account_id",
		cfrex.AccountBrowserRenderingGetPdfParams{
			CacheTtl: cfrex.F(86400.000000),
			AddScriptTag: cfrex.F([]cfrex.AccountBrowserRenderingGetPdfParamsAddScriptTag{{
				ID:      cfrex.F("id"),
				Content: cfrex.F("content"),
				Type:    cfrex.F("type"),
				URL:     cfrex.F("url"),
			}}),
			AddStyleTag: cfrex.F([]cfrex.AccountBrowserRenderingGetPdfParamsAddStyleTag{{
				Content: cfrex.F("content"),
				URL:     cfrex.F("url"),
			}}),
			AllowRequestPattern: cfrex.F([]string{"string"}),
			AllowResourceTypes:  cfrex.F([]cfrex.AccountBrowserRenderingGetPdfParamsAllowResourceType{cfrex.AccountBrowserRenderingGetPdfParamsAllowResourceTypeDocument}),
			Authenticate: cfrex.F(cfrex.AccountBrowserRenderingGetPdfParamsAuthenticate{
				Password: cfrex.F("x"),
				Username: cfrex.F("x"),
			}),
			BestAttempt: cfrex.F(true),
			Cookies: cfrex.F([]cfrex.AccountBrowserRenderingGetPdfParamsCookie{{
				Name:         cfrex.F("name"),
				Value:        cfrex.F("value"),
				Domain:       cfrex.F("domain"),
				Expires:      cfrex.F(0.000000),
				HTTPOnly:     cfrex.F(true),
				PartitionKey: cfrex.F("partitionKey"),
				Path:         cfrex.F("path"),
				Priority:     cfrex.F(cfrex.AccountBrowserRenderingGetPdfParamsCookiesPriorityLow),
				SameParty:    cfrex.F(true),
				SameSite:     cfrex.F(cfrex.AccountBrowserRenderingGetPdfParamsCookiesSameSiteStrict),
				Secure:       cfrex.F(true),
				SourcePort:   cfrex.F(0.000000),
				SourceScheme: cfrex.F(cfrex.AccountBrowserRenderingGetPdfParamsCookiesSourceSchemeUnset),
				URL:          cfrex.F("url"),
			}}),
			EmulateMediaType: cfrex.F("emulateMediaType"),
			GotoOptions: cfrex.F(cfrex.AccountBrowserRenderingGetPdfParamsGotoOptions{
				Referer:        cfrex.F("referer"),
				ReferrerPolicy: cfrex.F("referrerPolicy"),
				Timeout:        cfrex.F(60000.000000),
				WaitUntil:      cfrex.F[cfrex.AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilUnion](cfrex.AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilString(cfrex.AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilStringLoad)),
			}),
			HTML:                 cfrex.F("x"),
			RejectRequestPattern: cfrex.F([]string{"string"}),
			RejectResourceTypes:  cfrex.F([]cfrex.AccountBrowserRenderingGetPdfParamsRejectResourceType{cfrex.AccountBrowserRenderingGetPdfParamsRejectResourceTypeDocument}),
			SetExtraHTTPHeaders: cfrex.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cfrex.F(true),
			URL:                  cfrex.F("https://example.com"),
			UserAgent:            cfrex.F("userAgent"),
			Viewport: cfrex.F(cfrex.AccountBrowserRenderingGetPdfParamsViewport{
				Height:            cfrex.F(0.000000),
				Width:             cfrex.F(0.000000),
				DeviceScaleFactor: cfrex.F(0.000000),
				HasTouch:          cfrex.F(true),
				IsLandscape:       cfrex.F(true),
				IsMobile:          cfrex.F(true),
			}),
			WaitForSelector: cfrex.F(cfrex.AccountBrowserRenderingGetPdfParamsWaitForSelector{
				Selector: cfrex.F("selector"),
				Hidden:   cfrex.F(cfrex.AccountBrowserRenderingGetPdfParamsWaitForSelectorHiddenTrue),
				Timeout:  cfrex.F(60000.000000),
				Visible:  cfrex.F(cfrex.AccountBrowserRenderingGetPdfParamsWaitForSelectorVisibleTrue),
			}),
			WaitForTimeout: cfrex.F(60000.000000),
		},
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestAccountBrowserRenderingGetScreenshotWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cfrex.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIEmail("My API Email"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Accounts.BrowserRendering.GetScreenshot(
		context.TODO(),
		"account_id",
		cfrex.AccountBrowserRenderingGetScreenshotParams{
			CacheTtl: cfrex.F(86400.000000),
			AddScriptTag: cfrex.F([]cfrex.AccountBrowserRenderingGetScreenshotParamsAddScriptTag{{
				ID:      cfrex.F("id"),
				Content: cfrex.F("content"),
				Type:    cfrex.F("type"),
				URL:     cfrex.F("url"),
			}}),
			AddStyleTag: cfrex.F([]cfrex.AccountBrowserRenderingGetScreenshotParamsAddStyleTag{{
				Content: cfrex.F("content"),
				URL:     cfrex.F("url"),
			}}),
			AllowRequestPattern: cfrex.F([]string{"string"}),
			AllowResourceTypes:  cfrex.F([]cfrex.AccountBrowserRenderingGetScreenshotParamsAllowResourceType{cfrex.AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeDocument}),
			Authenticate: cfrex.F(cfrex.AccountBrowserRenderingGetScreenshotParamsAuthenticate{
				Password: cfrex.F("x"),
				Username: cfrex.F("x"),
			}),
			BestAttempt: cfrex.F(true),
			Cookies: cfrex.F([]cfrex.AccountBrowserRenderingGetScreenshotParamsCookie{{
				Name:         cfrex.F("name"),
				Value:        cfrex.F("value"),
				Domain:       cfrex.F("domain"),
				Expires:      cfrex.F(0.000000),
				HTTPOnly:     cfrex.F(true),
				PartitionKey: cfrex.F("partitionKey"),
				Path:         cfrex.F("path"),
				Priority:     cfrex.F(cfrex.AccountBrowserRenderingGetScreenshotParamsCookiesPriorityLow),
				SameParty:    cfrex.F(true),
				SameSite:     cfrex.F(cfrex.AccountBrowserRenderingGetScreenshotParamsCookiesSameSiteStrict),
				Secure:       cfrex.F(true),
				SourcePort:   cfrex.F(0.000000),
				SourceScheme: cfrex.F(cfrex.AccountBrowserRenderingGetScreenshotParamsCookiesSourceSchemeUnset),
				URL:          cfrex.F("url"),
			}}),
			EmulateMediaType: cfrex.F("emulateMediaType"),
			GotoOptions: cfrex.F(cfrex.AccountBrowserRenderingGetScreenshotParamsGotoOptions{
				Referer:        cfrex.F("referer"),
				ReferrerPolicy: cfrex.F("referrerPolicy"),
				Timeout:        cfrex.F(60000.000000),
				WaitUntil:      cfrex.F[cfrex.AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilUnion](cfrex.AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilString(cfrex.AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilStringLoad)),
			}),
			HTML:                 cfrex.F("x"),
			RejectRequestPattern: cfrex.F([]string{"string"}),
			RejectResourceTypes:  cfrex.F([]cfrex.AccountBrowserRenderingGetScreenshotParamsRejectResourceType{cfrex.AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeDocument}),
			ScreenshotOptions: cfrex.F(cfrex.AccountBrowserRenderingGetScreenshotParamsScreenshotOptions{
				CaptureBeyondViewport: cfrex.F(true),
				Clip: cfrex.F(cfrex.AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsClip{
					Height: cfrex.F(0.000000),
					Width:  cfrex.F(0.000000),
					X:      cfrex.F(0.000000),
					Y:      cfrex.F(0.000000),
					Scale:  cfrex.F(0.000000),
				}),
				Encoding:         cfrex.F(cfrex.AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsEncodingBinary),
				FromSurface:      cfrex.F(true),
				FullPage:         cfrex.F(true),
				OmitBackground:   cfrex.F(true),
				OptimizeForSpeed: cfrex.F(true),
				Quality:          cfrex.F(0.000000),
				Type:             cfrex.F(cfrex.AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsTypePng),
			}),
			ScrollPage: cfrex.F(true),
			Selector:   cfrex.F("selector"),
			SetExtraHTTPHeaders: cfrex.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cfrex.F(true),
			URL:                  cfrex.F("https://example.com"),
			UserAgent:            cfrex.F("userAgent"),
			Viewport: cfrex.F(cfrex.AccountBrowserRenderingGetScreenshotParamsViewport{
				Height:            cfrex.F(0.000000),
				Width:             cfrex.F(0.000000),
				DeviceScaleFactor: cfrex.F(0.000000),
				HasTouch:          cfrex.F(true),
				IsLandscape:       cfrex.F(true),
				IsMobile:          cfrex.F(true),
			}),
			WaitForSelector: cfrex.F(cfrex.AccountBrowserRenderingGetScreenshotParamsWaitForSelector{
				Selector: cfrex.F("selector"),
				Hidden:   cfrex.F(cfrex.AccountBrowserRenderingGetScreenshotParamsWaitForSelectorHiddenTrue),
				Timeout:  cfrex.F(60000.000000),
				Visible:  cfrex.F(cfrex.AccountBrowserRenderingGetScreenshotParamsWaitForSelectorVisibleTrue),
			}),
			WaitForTimeout: cfrex.F(60000.000000),
		},
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountBrowserRenderingGetSnapshotWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cfrex.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIEmail("My API Email"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Accounts.BrowserRendering.GetSnapshot(
		context.TODO(),
		"account_id",
		cfrex.AccountBrowserRenderingGetSnapshotParams{
			CacheTtl: cfrex.F(86400.000000),
			AddScriptTag: cfrex.F([]cfrex.AccountBrowserRenderingGetSnapshotParamsAddScriptTag{{
				ID:      cfrex.F("id"),
				Content: cfrex.F("content"),
				Type:    cfrex.F("type"),
				URL:     cfrex.F("url"),
			}}),
			AddStyleTag: cfrex.F([]cfrex.AccountBrowserRenderingGetSnapshotParamsAddStyleTag{{
				Content: cfrex.F("content"),
				URL:     cfrex.F("url"),
			}}),
			AllowRequestPattern: cfrex.F([]string{"string"}),
			AllowResourceTypes:  cfrex.F([]cfrex.AccountBrowserRenderingGetSnapshotParamsAllowResourceType{cfrex.AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeDocument}),
			Authenticate: cfrex.F(cfrex.AccountBrowserRenderingGetSnapshotParamsAuthenticate{
				Password: cfrex.F("x"),
				Username: cfrex.F("x"),
			}),
			BestAttempt: cfrex.F(true),
			Cookies: cfrex.F([]cfrex.AccountBrowserRenderingGetSnapshotParamsCookie{{
				Name:         cfrex.F("name"),
				Value:        cfrex.F("value"),
				Domain:       cfrex.F("domain"),
				Expires:      cfrex.F(0.000000),
				HTTPOnly:     cfrex.F(true),
				PartitionKey: cfrex.F("partitionKey"),
				Path:         cfrex.F("path"),
				Priority:     cfrex.F(cfrex.AccountBrowserRenderingGetSnapshotParamsCookiesPriorityLow),
				SameParty:    cfrex.F(true),
				SameSite:     cfrex.F(cfrex.AccountBrowserRenderingGetSnapshotParamsCookiesSameSiteStrict),
				Secure:       cfrex.F(true),
				SourcePort:   cfrex.F(0.000000),
				SourceScheme: cfrex.F(cfrex.AccountBrowserRenderingGetSnapshotParamsCookiesSourceSchemeUnset),
				URL:          cfrex.F("url"),
			}}),
			EmulateMediaType: cfrex.F("emulateMediaType"),
			GotoOptions: cfrex.F(cfrex.AccountBrowserRenderingGetSnapshotParamsGotoOptions{
				Referer:        cfrex.F("referer"),
				ReferrerPolicy: cfrex.F("referrerPolicy"),
				Timeout:        cfrex.F(60000.000000),
				WaitUntil:      cfrex.F[cfrex.AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilUnion](cfrex.AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilString(cfrex.AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilStringLoad)),
			}),
			HTML:                 cfrex.F("x"),
			RejectRequestPattern: cfrex.F([]string{"string"}),
			RejectResourceTypes:  cfrex.F([]cfrex.AccountBrowserRenderingGetSnapshotParamsRejectResourceType{cfrex.AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeDocument}),
			ScreenshotOptions: cfrex.F(cfrex.AccountBrowserRenderingGetSnapshotParamsScreenshotOptions{
				CaptureBeyondViewport: cfrex.F(true),
				Clip: cfrex.F(cfrex.AccountBrowserRenderingGetSnapshotParamsScreenshotOptionsClip{
					Height: cfrex.F(0.000000),
					Width:  cfrex.F(0.000000),
					X:      cfrex.F(0.000000),
					Y:      cfrex.F(0.000000),
					Scale:  cfrex.F(0.000000),
				}),
				FromSurface:      cfrex.F(true),
				FullPage:         cfrex.F(true),
				OmitBackground:   cfrex.F(true),
				OptimizeForSpeed: cfrex.F(true),
				Quality:          cfrex.F(0.000000),
				Type:             cfrex.F(cfrex.AccountBrowserRenderingGetSnapshotParamsScreenshotOptionsTypePng),
			}),
			SetExtraHTTPHeaders: cfrex.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cfrex.F(true),
			URL:                  cfrex.F("https://example.com"),
			UserAgent:            cfrex.F("userAgent"),
			Viewport: cfrex.F(cfrex.AccountBrowserRenderingGetSnapshotParamsViewport{
				Height:            cfrex.F(0.000000),
				Width:             cfrex.F(0.000000),
				DeviceScaleFactor: cfrex.F(0.000000),
				HasTouch:          cfrex.F(true),
				IsLandscape:       cfrex.F(true),
				IsMobile:          cfrex.F(true),
			}),
			WaitForSelector: cfrex.F(cfrex.AccountBrowserRenderingGetSnapshotParamsWaitForSelector{
				Selector: cfrex.F("selector"),
				Hidden:   cfrex.F(cfrex.AccountBrowserRenderingGetSnapshotParamsWaitForSelectorHiddenTrue),
				Timeout:  cfrex.F(60000.000000),
				Visible:  cfrex.F(cfrex.AccountBrowserRenderingGetSnapshotParamsWaitForSelectorVisibleTrue),
			}),
			WaitForTimeout: cfrex.F(60000.000000),
		},
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountBrowserRenderingScrapeElementsWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cfrex.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIEmail("My API Email"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Accounts.BrowserRendering.ScrapeElements(
		context.TODO(),
		"account_id",
		cfrex.AccountBrowserRenderingScrapeElementsParams{
			Elements: cfrex.F([]cfrex.AccountBrowserRenderingScrapeElementsParamsElement{{
				Selector: cfrex.F("selector"),
			}}),
			CacheTtl: cfrex.F(86400.000000),
			AddScriptTag: cfrex.F([]cfrex.AccountBrowserRenderingScrapeElementsParamsAddScriptTag{{
				ID:      cfrex.F("id"),
				Content: cfrex.F("content"),
				Type:    cfrex.F("type"),
				URL:     cfrex.F("url"),
			}}),
			AddStyleTag: cfrex.F([]cfrex.AccountBrowserRenderingScrapeElementsParamsAddStyleTag{{
				Content: cfrex.F("content"),
				URL:     cfrex.F("url"),
			}}),
			AllowRequestPattern: cfrex.F([]string{"string"}),
			AllowResourceTypes:  cfrex.F([]cfrex.AccountBrowserRenderingScrapeElementsParamsAllowResourceType{cfrex.AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeDocument}),
			Authenticate: cfrex.F(cfrex.AccountBrowserRenderingScrapeElementsParamsAuthenticate{
				Password: cfrex.F("x"),
				Username: cfrex.F("x"),
			}),
			BestAttempt: cfrex.F(true),
			Cookies: cfrex.F([]cfrex.AccountBrowserRenderingScrapeElementsParamsCookie{{
				Name:         cfrex.F("name"),
				Value:        cfrex.F("value"),
				Domain:       cfrex.F("domain"),
				Expires:      cfrex.F(0.000000),
				HTTPOnly:     cfrex.F(true),
				PartitionKey: cfrex.F("partitionKey"),
				Path:         cfrex.F("path"),
				Priority:     cfrex.F(cfrex.AccountBrowserRenderingScrapeElementsParamsCookiesPriorityLow),
				SameParty:    cfrex.F(true),
				SameSite:     cfrex.F(cfrex.AccountBrowserRenderingScrapeElementsParamsCookiesSameSiteStrict),
				Secure:       cfrex.F(true),
				SourcePort:   cfrex.F(0.000000),
				SourceScheme: cfrex.F(cfrex.AccountBrowserRenderingScrapeElementsParamsCookiesSourceSchemeUnset),
				URL:          cfrex.F("url"),
			}}),
			EmulateMediaType: cfrex.F("emulateMediaType"),
			GotoOptions: cfrex.F(cfrex.AccountBrowserRenderingScrapeElementsParamsGotoOptions{
				Referer:        cfrex.F("referer"),
				ReferrerPolicy: cfrex.F("referrerPolicy"),
				Timeout:        cfrex.F(60000.000000),
				WaitUntil:      cfrex.F[cfrex.AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilUnion](cfrex.AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilString(cfrex.AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilStringLoad)),
			}),
			HTML:                 cfrex.F("x"),
			RejectRequestPattern: cfrex.F([]string{"string"}),
			RejectResourceTypes:  cfrex.F([]cfrex.AccountBrowserRenderingScrapeElementsParamsRejectResourceType{cfrex.AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeDocument}),
			SetExtraHTTPHeaders: cfrex.F(map[string]string{
				"foo": "string",
			}),
			SetJavaScriptEnabled: cfrex.F(true),
			URL:                  cfrex.F("https://example.com"),
			UserAgent:            cfrex.F("userAgent"),
			Viewport: cfrex.F(cfrex.AccountBrowserRenderingScrapeElementsParamsViewport{
				Height:            cfrex.F(0.000000),
				Width:             cfrex.F(0.000000),
				DeviceScaleFactor: cfrex.F(0.000000),
				HasTouch:          cfrex.F(true),
				IsLandscape:       cfrex.F(true),
				IsMobile:          cfrex.F(true),
			}),
			WaitForSelector: cfrex.F(cfrex.AccountBrowserRenderingScrapeElementsParamsWaitForSelector{
				Selector: cfrex.F("selector"),
				Hidden:   cfrex.F(cfrex.AccountBrowserRenderingScrapeElementsParamsWaitForSelectorHiddenTrue),
				Timeout:  cfrex.F(60000.000000),
				Visible:  cfrex.F(cfrex.AccountBrowserRenderingScrapeElementsParamsWaitForSelectorVisibleTrue),
			}),
			WaitForTimeout: cfrex.F(60000.000000),
		},
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
