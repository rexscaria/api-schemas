// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/rexscaria/api-schemas"
	"github.com/rexscaria/api-schemas/internal/testutil"
	"github.com/rexscaria/api-schemas/option"
)

func TestAccountPageProjectNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Pages.Projects.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cfrex.AccountPageProjectNewParams{
			Project: cfrex.ProjectParam{
				BuildConfig: cfrex.F(cfrex.BuildConfigParam{
					BuildCaching:      cfrex.F(true),
					BuildCommand:      cfrex.F("npm run build"),
					DestinationDir:    cfrex.F("build"),
					RootDir:           cfrex.F("/"),
					WebAnalyticsTag:   cfrex.F("cee1c73f6e4743d0b5e6bb1a0bcaabcc"),
					WebAnalyticsToken: cfrex.F("021e1057c18547eca7b79f2516f06o7x"),
				}),
				CanonicalDeployment: cfrex.F(cfrex.ProjectCanonicalDeploymentParam{
					DeploymentsParam: cfrex.DeploymentsParam{
						BuildConfig: cfrex.F(cfrex.BuildConfigParam{
							BuildCaching:      cfrex.F(true),
							BuildCommand:      cfrex.F("npm run build"),
							DestinationDir:    cfrex.F("build"),
							RootDir:           cfrex.F("/"),
							WebAnalyticsTag:   cfrex.F("cee1c73f6e4743d0b5e6bb1a0bcaabcc"),
							WebAnalyticsToken: cfrex.F("021e1057c18547eca7b79f2516f06o7x"),
						}),
						EnvVars: cfrex.F(map[string]cfrex.DeploymentsEnvVarsUnionParam{
							"foo": cfrex.DeploymentsEnvVarsPagesPlainTextEnvVarParam{
								Type:  cfrex.F(cfrex.DeploymentsEnvVarsPagesPlainTextEnvVarTypePlainText),
								Value: cfrex.F("hello world"),
							},
						}),
					},
				}),
				DeploymentConfigs: cfrex.F(cfrex.ProjectDeploymentConfigsParam{
					Preview: cfrex.F(cfrex.DeploymentConfigsValuesParam{
						AIBindings: cfrex.F(map[string]cfrex.DeploymentConfigsValuesAIBindingParam{
							"AI_BINDING": {
								ProjectID: cfrex.F("some-project-id"),
							},
						}),
						AnalyticsEngineDatasets: cfrex.F(map[string]cfrex.DeploymentConfigsValuesAnalyticsEngineDatasetParam{
							"ANALYTICS_ENGINE_BINDING": {
								Dataset: cfrex.F("api_analytics"),
							},
						}),
						Browsers: cfrex.F(map[string]cfrex.DeploymentConfigsValuesBrowserParam{
							"BROWSER": {},
						}),
						CompatibilityDate:  cfrex.F("2022-01-01"),
						CompatibilityFlags: cfrex.F([]string{"url_standard"}),
						D1Databases: cfrex.F(map[string]cfrex.DeploymentConfigsValuesD1DatabaseParam{
							"D1_BINDING": {
								ID: cfrex.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
							},
						}),
						DurableObjectNamespaces: cfrex.F(map[string]cfrex.DeploymentConfigsValuesDurableObjectNamespaceParam{
							"DO_BINDING": {
								NamespaceID: cfrex.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
							},
						}),
						EnvVars: cfrex.F(map[string]cfrex.DeploymentConfigsValuesEnvVarsUnionParam{
							"foo": cfrex.DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarParam{
								Type:  cfrex.F(cfrex.DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarTypePlainText),
								Value: cfrex.F("hello world"),
							},
						}),
						HyperdriveBindings: cfrex.F(map[string]cfrex.DeploymentConfigsValuesHyperdriveBindingParam{
							"HYPERDRIVE": {
								ID: cfrex.F("a76a99bc342644deb02c38d66082262a"),
							},
						}),
						KvNamespaces: cfrex.F(map[string]cfrex.DeploymentConfigsValuesKvNamespaceParam{
							"KV_BINDING": {
								NamespaceID: cfrex.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
							},
						}),
						MtlsCertificates: cfrex.F(map[string]cfrex.DeploymentConfigsValuesMtlsCertificateParam{
							"MTLS": {
								CertificateID: cfrex.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
							},
						}),
						Placement: cfrex.F(cfrex.DeploymentConfigsValuesPlacementParam{
							Mode: cfrex.F("smart"),
						}),
						QueueProducers: cfrex.F(map[string]cfrex.DeploymentConfigsValuesQueueProducerParam{
							"QUEUE_PRODUCER_BINDING": {
								Name: cfrex.F("some-queue"),
							},
						}),
						R2Buckets: cfrex.F(map[string]cfrex.DeploymentConfigsValuesR2BucketParam{
							"R2_BINDING": {
								Jurisdiction: cfrex.F("eu"),
								Name:         cfrex.F("some-bucket"),
							},
						}),
						Services: cfrex.F(map[string]cfrex.DeploymentConfigsValuesServiceParam{
							"SERVICE_BINDING": {
								Entrypoint:  cfrex.F("MyHandler"),
								Environment: cfrex.F("production"),
								Service:     cfrex.F("example-worker"),
							},
						}),
						VectorizeBindings: cfrex.F(map[string]cfrex.DeploymentConfigsValuesVectorizeBindingParam{
							"VECTORIZE": {
								IndexName: cfrex.F("my_index"),
							},
						}),
					}),
					Production: cfrex.F(cfrex.DeploymentConfigsValuesParam{
						AIBindings: cfrex.F(map[string]cfrex.DeploymentConfigsValuesAIBindingParam{
							"AI_BINDING": {
								ProjectID: cfrex.F("some-project-id"),
							},
						}),
						AnalyticsEngineDatasets: cfrex.F(map[string]cfrex.DeploymentConfigsValuesAnalyticsEngineDatasetParam{
							"ANALYTICS_ENGINE_BINDING": {
								Dataset: cfrex.F("api_analytics"),
							},
						}),
						Browsers: cfrex.F(map[string]cfrex.DeploymentConfigsValuesBrowserParam{
							"BROWSER": {},
						}),
						CompatibilityDate:  cfrex.F("2022-01-01"),
						CompatibilityFlags: cfrex.F([]string{"url_standard"}),
						D1Databases: cfrex.F(map[string]cfrex.DeploymentConfigsValuesD1DatabaseParam{
							"D1_BINDING": {
								ID: cfrex.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
							},
						}),
						DurableObjectNamespaces: cfrex.F(map[string]cfrex.DeploymentConfigsValuesDurableObjectNamespaceParam{
							"DO_BINDING": {
								NamespaceID: cfrex.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
							},
						}),
						EnvVars: cfrex.F(map[string]cfrex.DeploymentConfigsValuesEnvVarsUnionParam{
							"foo": cfrex.DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarParam{
								Type:  cfrex.F(cfrex.DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarTypePlainText),
								Value: cfrex.F("hello world"),
							},
						}),
						HyperdriveBindings: cfrex.F(map[string]cfrex.DeploymentConfigsValuesHyperdriveBindingParam{
							"HYPERDRIVE": {
								ID: cfrex.F("a76a99bc342644deb02c38d66082262a"),
							},
						}),
						KvNamespaces: cfrex.F(map[string]cfrex.DeploymentConfigsValuesKvNamespaceParam{
							"KV_BINDING": {
								NamespaceID: cfrex.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
							},
						}),
						MtlsCertificates: cfrex.F(map[string]cfrex.DeploymentConfigsValuesMtlsCertificateParam{
							"MTLS": {
								CertificateID: cfrex.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
							},
						}),
						Placement: cfrex.F(cfrex.DeploymentConfigsValuesPlacementParam{
							Mode: cfrex.F("smart"),
						}),
						QueueProducers: cfrex.F(map[string]cfrex.DeploymentConfigsValuesQueueProducerParam{
							"QUEUE_PRODUCER_BINDING": {
								Name: cfrex.F("some-queue"),
							},
						}),
						R2Buckets: cfrex.F(map[string]cfrex.DeploymentConfigsValuesR2BucketParam{
							"R2_BINDING": {
								Jurisdiction: cfrex.F("eu"),
								Name:         cfrex.F("some-bucket"),
							},
						}),
						Services: cfrex.F(map[string]cfrex.DeploymentConfigsValuesServiceParam{
							"SERVICE_BINDING": {
								Entrypoint:  cfrex.F("MyHandler"),
								Environment: cfrex.F("production"),
								Service:     cfrex.F("example-worker"),
							},
						}),
						VectorizeBindings: cfrex.F(map[string]cfrex.DeploymentConfigsValuesVectorizeBindingParam{
							"VECTORIZE": {
								IndexName: cfrex.F("my_index"),
							},
						}),
					}),
				}),
				LatestDeployment: cfrex.F(cfrex.ProjectLatestDeploymentParam{
					DeploymentsParam: cfrex.DeploymentsParam{
						BuildConfig: cfrex.F(cfrex.BuildConfigParam{
							BuildCaching:      cfrex.F(true),
							BuildCommand:      cfrex.F("npm run build"),
							DestinationDir:    cfrex.F("build"),
							RootDir:           cfrex.F("/"),
							WebAnalyticsTag:   cfrex.F("cee1c73f6e4743d0b5e6bb1a0bcaabcc"),
							WebAnalyticsToken: cfrex.F("021e1057c18547eca7b79f2516f06o7x"),
						}),
						EnvVars: cfrex.F(map[string]cfrex.DeploymentsEnvVarsUnionParam{
							"foo": cfrex.DeploymentsEnvVarsPagesPlainTextEnvVarParam{
								Type:  cfrex.F(cfrex.DeploymentsEnvVarsPagesPlainTextEnvVarTypePlainText),
								Value: cfrex.F("hello world"),
							},
						}),
					},
				}),
				Name:             cfrex.F("NextJS Blog"),
				ProductionBranch: cfrex.F("main"),
			},
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

func TestAccountPageProjectGet(t *testing.T) {
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
	_, err := client.Accounts.Pages.Projects.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is-my-project-01",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountPageProjectUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Pages.Projects.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is-my-project-01",
		cfrex.AccountPageProjectUpdateParams{
			Body: cfrex.AccountPageProjectUpdateParamsBody{
				ProjectParam: cfrex.ProjectParam{
					BuildConfig: cfrex.F(cfrex.BuildConfigParam{
						BuildCaching:      cfrex.F(true),
						BuildCommand:      cfrex.F("npm run build"),
						DestinationDir:    cfrex.F("build"),
						RootDir:           cfrex.F("/"),
						WebAnalyticsTag:   cfrex.F("cee1c73f6e4743d0b5e6bb1a0bcaabcc"),
						WebAnalyticsToken: cfrex.F("021e1057c18547eca7b79f2516f06o7x"),
					}),
					CanonicalDeployment: cfrex.F(cfrex.ProjectCanonicalDeploymentParam{
						DeploymentsParam: cfrex.DeploymentsParam{
							BuildConfig: cfrex.F(cfrex.BuildConfigParam{
								BuildCaching:      cfrex.F(true),
								BuildCommand:      cfrex.F("npm run build"),
								DestinationDir:    cfrex.F("build"),
								RootDir:           cfrex.F("/"),
								WebAnalyticsTag:   cfrex.F("cee1c73f6e4743d0b5e6bb1a0bcaabcc"),
								WebAnalyticsToken: cfrex.F("021e1057c18547eca7b79f2516f06o7x"),
							}),
							EnvVars: cfrex.F(map[string]cfrex.DeploymentsEnvVarsUnionParam{
								"foo": cfrex.DeploymentsEnvVarsPagesPlainTextEnvVarParam{
									Type:  cfrex.F(cfrex.DeploymentsEnvVarsPagesPlainTextEnvVarTypePlainText),
									Value: cfrex.F("hello world"),
								},
							}),
						},
					}),
					DeploymentConfigs: cfrex.F(cfrex.ProjectDeploymentConfigsParam{
						Preview: cfrex.F(cfrex.DeploymentConfigsValuesParam{
							AIBindings: cfrex.F(map[string]cfrex.DeploymentConfigsValuesAIBindingParam{
								"AI_BINDING": {
									ProjectID: cfrex.F("some-project-id"),
								},
							}),
							AnalyticsEngineDatasets: cfrex.F(map[string]cfrex.DeploymentConfigsValuesAnalyticsEngineDatasetParam{
								"ANALYTICS_ENGINE_BINDING": {
									Dataset: cfrex.F("api_analytics"),
								},
							}),
							Browsers: cfrex.F(map[string]cfrex.DeploymentConfigsValuesBrowserParam{
								"BROWSER": {},
							}),
							CompatibilityDate:  cfrex.F("2022-01-01"),
							CompatibilityFlags: cfrex.F([]string{"url_standard"}),
							D1Databases: cfrex.F(map[string]cfrex.DeploymentConfigsValuesD1DatabaseParam{
								"D1_BINDING": {
									ID: cfrex.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
								},
							}),
							DurableObjectNamespaces: cfrex.F(map[string]cfrex.DeploymentConfigsValuesDurableObjectNamespaceParam{
								"DO_BINDING": {
									NamespaceID: cfrex.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
								},
							}),
							EnvVars: cfrex.F(map[string]cfrex.DeploymentConfigsValuesEnvVarsUnionParam{
								"foo": cfrex.DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarParam{
									Type:  cfrex.F(cfrex.DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarTypePlainText),
									Value: cfrex.F("hello world"),
								},
							}),
							HyperdriveBindings: cfrex.F(map[string]cfrex.DeploymentConfigsValuesHyperdriveBindingParam{
								"HYPERDRIVE": {
									ID: cfrex.F("a76a99bc342644deb02c38d66082262a"),
								},
							}),
							KvNamespaces: cfrex.F(map[string]cfrex.DeploymentConfigsValuesKvNamespaceParam{
								"KV_BINDING": {
									NamespaceID: cfrex.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
								},
							}),
							MtlsCertificates: cfrex.F(map[string]cfrex.DeploymentConfigsValuesMtlsCertificateParam{
								"MTLS": {
									CertificateID: cfrex.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
								},
							}),
							Placement: cfrex.F(cfrex.DeploymentConfigsValuesPlacementParam{
								Mode: cfrex.F("smart"),
							}),
							QueueProducers: cfrex.F(map[string]cfrex.DeploymentConfigsValuesQueueProducerParam{
								"QUEUE_PRODUCER_BINDING": {
									Name: cfrex.F("some-queue"),
								},
							}),
							R2Buckets: cfrex.F(map[string]cfrex.DeploymentConfigsValuesR2BucketParam{
								"R2_BINDING": {
									Jurisdiction: cfrex.F("eu"),
									Name:         cfrex.F("some-bucket"),
								},
							}),
							Services: cfrex.F(map[string]cfrex.DeploymentConfigsValuesServiceParam{
								"SERVICE_BINDING": {
									Entrypoint:  cfrex.F("MyHandler"),
									Environment: cfrex.F("production"),
									Service:     cfrex.F("example-worker"),
								},
							}),
							VectorizeBindings: cfrex.F(map[string]cfrex.DeploymentConfigsValuesVectorizeBindingParam{
								"VECTORIZE": {
									IndexName: cfrex.F("my_index"),
								},
							}),
						}),
						Production: cfrex.F(cfrex.DeploymentConfigsValuesParam{
							AIBindings: cfrex.F(map[string]cfrex.DeploymentConfigsValuesAIBindingParam{
								"AI_BINDING": {
									ProjectID: cfrex.F("some-project-id"),
								},
							}),
							AnalyticsEngineDatasets: cfrex.F(map[string]cfrex.DeploymentConfigsValuesAnalyticsEngineDatasetParam{
								"ANALYTICS_ENGINE_BINDING": {
									Dataset: cfrex.F("api_analytics"),
								},
							}),
							Browsers: cfrex.F(map[string]cfrex.DeploymentConfigsValuesBrowserParam{
								"BROWSER": {},
							}),
							CompatibilityDate:  cfrex.F("2022-01-01"),
							CompatibilityFlags: cfrex.F([]string{"url_standard"}),
							D1Databases: cfrex.F(map[string]cfrex.DeploymentConfigsValuesD1DatabaseParam{
								"D1_BINDING": {
									ID: cfrex.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
								},
							}),
							DurableObjectNamespaces: cfrex.F(map[string]cfrex.DeploymentConfigsValuesDurableObjectNamespaceParam{
								"DO_BINDING": {
									NamespaceID: cfrex.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
								},
							}),
							EnvVars: cfrex.F(map[string]cfrex.DeploymentConfigsValuesEnvVarsUnionParam{
								"foo": cfrex.DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarParam{
									Type:  cfrex.F(cfrex.DeploymentConfigsValuesEnvVarsPagesPlainTextEnvVarTypePlainText),
									Value: cfrex.F("hello world"),
								},
							}),
							HyperdriveBindings: cfrex.F(map[string]cfrex.DeploymentConfigsValuesHyperdriveBindingParam{
								"HYPERDRIVE": {
									ID: cfrex.F("a76a99bc342644deb02c38d66082262a"),
								},
							}),
							KvNamespaces: cfrex.F(map[string]cfrex.DeploymentConfigsValuesKvNamespaceParam{
								"KV_BINDING": {
									NamespaceID: cfrex.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
								},
							}),
							MtlsCertificates: cfrex.F(map[string]cfrex.DeploymentConfigsValuesMtlsCertificateParam{
								"MTLS": {
									CertificateID: cfrex.F("d7cdd17c-916f-4cb7-aabe-585eb382ec4e"),
								},
							}),
							Placement: cfrex.F(cfrex.DeploymentConfigsValuesPlacementParam{
								Mode: cfrex.F("smart"),
							}),
							QueueProducers: cfrex.F(map[string]cfrex.DeploymentConfigsValuesQueueProducerParam{
								"QUEUE_PRODUCER_BINDING": {
									Name: cfrex.F("some-queue"),
								},
							}),
							R2Buckets: cfrex.F(map[string]cfrex.DeploymentConfigsValuesR2BucketParam{
								"R2_BINDING": {
									Jurisdiction: cfrex.F("eu"),
									Name:         cfrex.F("some-bucket"),
								},
							}),
							Services: cfrex.F(map[string]cfrex.DeploymentConfigsValuesServiceParam{
								"SERVICE_BINDING": {
									Entrypoint:  cfrex.F("MyHandler"),
									Environment: cfrex.F("production"),
									Service:     cfrex.F("example-worker"),
								},
							}),
							VectorizeBindings: cfrex.F(map[string]cfrex.DeploymentConfigsValuesVectorizeBindingParam{
								"VECTORIZE": {
									IndexName: cfrex.F("my_index"),
								},
							}),
						}),
					}),
					LatestDeployment: cfrex.F(cfrex.ProjectLatestDeploymentParam{
						DeploymentsParam: cfrex.DeploymentsParam{
							BuildConfig: cfrex.F(cfrex.BuildConfigParam{
								BuildCaching:      cfrex.F(true),
								BuildCommand:      cfrex.F("npm run build"),
								DestinationDir:    cfrex.F("build"),
								RootDir:           cfrex.F("/"),
								WebAnalyticsTag:   cfrex.F("cee1c73f6e4743d0b5e6bb1a0bcaabcc"),
								WebAnalyticsToken: cfrex.F("021e1057c18547eca7b79f2516f06o7x"),
							}),
							EnvVars: cfrex.F(map[string]cfrex.DeploymentsEnvVarsUnionParam{
								"foo": cfrex.DeploymentsEnvVarsPagesPlainTextEnvVarParam{
									Type:  cfrex.F(cfrex.DeploymentsEnvVarsPagesPlainTextEnvVarTypePlainText),
									Value: cfrex.F("hello world"),
								},
							}),
						},
					}),
					Name:             cfrex.F("NextJS Blog"),
					ProductionBranch: cfrex.F("main"),
				},
			},
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

func TestAccountPageProjectList(t *testing.T) {
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
	_, err := client.Accounts.Pages.Projects.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountPageProjectDelete(t *testing.T) {
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
	_, err := client.Accounts.Pages.Projects.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is-my-project-01",
		cfrex.AccountPageProjectDeleteParams{
			Body: map[string]interface{}{},
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

func TestAccountPageProjectPurgeBuildCache(t *testing.T) {
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
	_, err := client.Accounts.Pages.Projects.PurgeBuildCache(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is-my-project-01",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
