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

func TestZoneLoadBalancerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.LoadBalancers.New(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cfrex.ZoneLoadBalancerNewParams{
			DefaultPools: cfrex.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
			FallbackPool: cfrex.F("fallback_pool"),
			Name:         cfrex.F("www.example.com"),
			AdaptiveRouting: cfrex.F(cfrex.AdaptiveRoutingParam{
				FailoverAcrossPools: cfrex.F(true),
			}),
			CountryPools: cfrex.F(map[string][]string{
				"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
				"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
			}),
			Description: cfrex.F("Load Balancer for www.example.com"),
			LocationStrategy: cfrex.F(cfrex.LocationStrategyParam{
				Mode:      cfrex.F(cfrex.LocationStrategyModeResolverIP),
				PreferEcs: cfrex.F(cfrex.LocationStrategyPreferEcsAlways),
			}),
			Networks: cfrex.F([]string{"string"}),
			PopPools: cfrex.F(map[string][]string{
				"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
				"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
				"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
			}),
			Proxied: cfrex.F(true),
			RandomSteering: cfrex.F(cfrex.RandomSteeringParam{
				DefaultWeight: cfrex.F(0.200000),
				PoolWeights: cfrex.F(map[string]float64{
					"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
					"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
				}),
			}),
			RegionPools: cfrex.F(map[string][]string{
				"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
				"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
			}),
			Rules: cfrex.F([]cfrex.LoadBalancingRuleParam{{
				Condition: cfrex.F(`http.request.uri.path contains "/testing"`),
				Disabled:  cfrex.F(true),
				FixedResponse: cfrex.F(cfrex.LoadBalancingRuleFixedResponseParam{
					ContentType: cfrex.F("application/json"),
					Location:    cfrex.F("www.example.com"),
					MessageBody: cfrex.F("Testing Hello"),
					StatusCode:  cfrex.F(int64(0)),
				}),
				Name: cfrex.F("route the path /testing to testing datacenter."),
				Overrides: cfrex.F(cfrex.LoadBalancingRuleOverridesParam{
					AdaptiveRouting: cfrex.F(cfrex.AdaptiveRoutingParam{
						FailoverAcrossPools: cfrex.F(true),
					}),
					CountryPools: cfrex.F(map[string][]string{
						"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
						"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					DefaultPools: cfrex.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cfrex.F("fallback_pool"),
					LocationStrategy: cfrex.F(cfrex.LocationStrategyParam{
						Mode:      cfrex.F(cfrex.LocationStrategyModeResolverIP),
						PreferEcs: cfrex.F(cfrex.LocationStrategyPreferEcsAlways),
					}),
					PopPools: cfrex.F(map[string][]string{
						"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
						"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
						"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					RandomSteering: cfrex.F(cfrex.RandomSteeringParam{
						DefaultWeight: cfrex.F(0.200000),
						PoolWeights: cfrex.F(map[string]float64{
							"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
							"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
						}),
					}),
					RegionPools: cfrex.F(map[string][]string{
						"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
						"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
					}),
					SessionAffinity: cfrex.F(cfrex.SessionAffinityCookie),
					SessionAffinityAttributes: cfrex.F(cfrex.SessionAffinityAttributesParam{
						DrainDuration:        cfrex.F(100.000000),
						Headers:              cfrex.F([]string{"_1K--W2kIFj1"}),
						RequireAllHeaders:    cfrex.F(true),
						Samesite:             cfrex.F(cfrex.SessionAffinityAttributesSamesiteAuto),
						Secure:               cfrex.F(cfrex.SessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cfrex.F(cfrex.SessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTtl: cfrex.F(1800.000000),
					SteeringPolicy:     cfrex.F(cfrex.SteeringPolicyDynamicLatency),
					Ttl:                cfrex.F(30.000000),
				}),
				Priority:   cfrex.F(int64(0)),
				Terminates: cfrex.F(true),
			}}),
			SessionAffinity: cfrex.F(cfrex.SessionAffinityCookie),
			SessionAffinityAttributes: cfrex.F(cfrex.SessionAffinityAttributesParam{
				DrainDuration:        cfrex.F(100.000000),
				Headers:              cfrex.F([]string{"_1K--W2kIFj1"}),
				RequireAllHeaders:    cfrex.F(true),
				Samesite:             cfrex.F(cfrex.SessionAffinityAttributesSamesiteAuto),
				Secure:               cfrex.F(cfrex.SessionAffinityAttributesSecureAuto),
				ZeroDowntimeFailover: cfrex.F(cfrex.SessionAffinityAttributesZeroDowntimeFailoverSticky),
			}),
			SessionAffinityTtl: cfrex.F(1800.000000),
			SteeringPolicy:     cfrex.F(cfrex.SteeringPolicyDynamicLatency),
			Ttl:                cfrex.F(30.000000),
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

func TestZoneLoadBalancerGet(t *testing.T) {
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
	_, err := client.Zones.LoadBalancers.Get(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"699d98642c564d2e855e9661899b7252",
	)
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneLoadBalancerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.LoadBalancers.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"699d98642c564d2e855e9661899b7252",
		cfrex.ZoneLoadBalancerUpdateParams{
			DefaultPools: cfrex.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
			FallbackPool: cfrex.F("fallback_pool"),
			Name:         cfrex.F("www.example.com"),
			AdaptiveRouting: cfrex.F(cfrex.AdaptiveRoutingParam{
				FailoverAcrossPools: cfrex.F(true),
			}),
			CountryPools: cfrex.F(map[string][]string{
				"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
				"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
			}),
			Description: cfrex.F("Load Balancer for www.example.com"),
			Enabled:     cfrex.F(true),
			LocationStrategy: cfrex.F(cfrex.LocationStrategyParam{
				Mode:      cfrex.F(cfrex.LocationStrategyModeResolverIP),
				PreferEcs: cfrex.F(cfrex.LocationStrategyPreferEcsAlways),
			}),
			Networks: cfrex.F([]string{"string"}),
			PopPools: cfrex.F(map[string][]string{
				"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
				"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
				"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
			}),
			Proxied: cfrex.F(true),
			RandomSteering: cfrex.F(cfrex.RandomSteeringParam{
				DefaultWeight: cfrex.F(0.200000),
				PoolWeights: cfrex.F(map[string]float64{
					"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
					"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
				}),
			}),
			RegionPools: cfrex.F(map[string][]string{
				"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
				"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
			}),
			Rules: cfrex.F([]cfrex.LoadBalancingRuleParam{{
				Condition: cfrex.F(`http.request.uri.path contains "/testing"`),
				Disabled:  cfrex.F(true),
				FixedResponse: cfrex.F(cfrex.LoadBalancingRuleFixedResponseParam{
					ContentType: cfrex.F("application/json"),
					Location:    cfrex.F("www.example.com"),
					MessageBody: cfrex.F("Testing Hello"),
					StatusCode:  cfrex.F(int64(0)),
				}),
				Name: cfrex.F("route the path /testing to testing datacenter."),
				Overrides: cfrex.F(cfrex.LoadBalancingRuleOverridesParam{
					AdaptiveRouting: cfrex.F(cfrex.AdaptiveRoutingParam{
						FailoverAcrossPools: cfrex.F(true),
					}),
					CountryPools: cfrex.F(map[string][]string{
						"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
						"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					DefaultPools: cfrex.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cfrex.F("fallback_pool"),
					LocationStrategy: cfrex.F(cfrex.LocationStrategyParam{
						Mode:      cfrex.F(cfrex.LocationStrategyModeResolverIP),
						PreferEcs: cfrex.F(cfrex.LocationStrategyPreferEcsAlways),
					}),
					PopPools: cfrex.F(map[string][]string{
						"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
						"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
						"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					RandomSteering: cfrex.F(cfrex.RandomSteeringParam{
						DefaultWeight: cfrex.F(0.200000),
						PoolWeights: cfrex.F(map[string]float64{
							"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
							"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
						}),
					}),
					RegionPools: cfrex.F(map[string][]string{
						"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
						"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
					}),
					SessionAffinity: cfrex.F(cfrex.SessionAffinityCookie),
					SessionAffinityAttributes: cfrex.F(cfrex.SessionAffinityAttributesParam{
						DrainDuration:        cfrex.F(100.000000),
						Headers:              cfrex.F([]string{"_1K--W2kIFj1"}),
						RequireAllHeaders:    cfrex.F(true),
						Samesite:             cfrex.F(cfrex.SessionAffinityAttributesSamesiteAuto),
						Secure:               cfrex.F(cfrex.SessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cfrex.F(cfrex.SessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTtl: cfrex.F(1800.000000),
					SteeringPolicy:     cfrex.F(cfrex.SteeringPolicyDynamicLatency),
					Ttl:                cfrex.F(30.000000),
				}),
				Priority:   cfrex.F(int64(0)),
				Terminates: cfrex.F(true),
			}}),
			SessionAffinity: cfrex.F(cfrex.SessionAffinityCookie),
			SessionAffinityAttributes: cfrex.F(cfrex.SessionAffinityAttributesParam{
				DrainDuration:        cfrex.F(100.000000),
				Headers:              cfrex.F([]string{"_1K--W2kIFj1"}),
				RequireAllHeaders:    cfrex.F(true),
				Samesite:             cfrex.F(cfrex.SessionAffinityAttributesSamesiteAuto),
				Secure:               cfrex.F(cfrex.SessionAffinityAttributesSecureAuto),
				ZeroDowntimeFailover: cfrex.F(cfrex.SessionAffinityAttributesZeroDowntimeFailoverSticky),
			}),
			SessionAffinityTtl: cfrex.F(1800.000000),
			SteeringPolicy:     cfrex.F(cfrex.SteeringPolicyDynamicLatency),
			Ttl:                cfrex.F(30.000000),
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

func TestZoneLoadBalancerList(t *testing.T) {
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
	_, err := client.Zones.LoadBalancers.List(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cfrex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneLoadBalancerDelete(t *testing.T) {
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
	_, err := client.Zones.LoadBalancers.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"699d98642c564d2e855e9661899b7252",
		cfrex.ZoneLoadBalancerDeleteParams{
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

func TestZoneLoadBalancerPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.LoadBalancers.Patch(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"699d98642c564d2e855e9661899b7252",
		cfrex.ZoneLoadBalancerPatchParams{
			AdaptiveRouting: cfrex.F(cfrex.AdaptiveRoutingParam{
				FailoverAcrossPools: cfrex.F(true),
			}),
			CountryPools: cfrex.F(map[string][]string{
				"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
				"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
			}),
			DefaultPools: cfrex.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
			Description:  cfrex.F("Load Balancer for www.example.com"),
			Enabled:      cfrex.F(true),
			FallbackPool: cfrex.F("fallback_pool"),
			LocationStrategy: cfrex.F(cfrex.LocationStrategyParam{
				Mode:      cfrex.F(cfrex.LocationStrategyModeResolverIP),
				PreferEcs: cfrex.F(cfrex.LocationStrategyPreferEcsAlways),
			}),
			Name: cfrex.F("www.example.com"),
			PopPools: cfrex.F(map[string][]string{
				"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
				"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
				"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
			}),
			Proxied: cfrex.F(true),
			RandomSteering: cfrex.F(cfrex.RandomSteeringParam{
				DefaultWeight: cfrex.F(0.200000),
				PoolWeights: cfrex.F(map[string]float64{
					"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
					"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
				}),
			}),
			RegionPools: cfrex.F(map[string][]string{
				"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
				"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
			}),
			Rules: cfrex.F([]cfrex.LoadBalancingRuleParam{{
				Condition: cfrex.F(`http.request.uri.path contains "/testing"`),
				Disabled:  cfrex.F(true),
				FixedResponse: cfrex.F(cfrex.LoadBalancingRuleFixedResponseParam{
					ContentType: cfrex.F("application/json"),
					Location:    cfrex.F("www.example.com"),
					MessageBody: cfrex.F("Testing Hello"),
					StatusCode:  cfrex.F(int64(0)),
				}),
				Name: cfrex.F("route the path /testing to testing datacenter."),
				Overrides: cfrex.F(cfrex.LoadBalancingRuleOverridesParam{
					AdaptiveRouting: cfrex.F(cfrex.AdaptiveRoutingParam{
						FailoverAcrossPools: cfrex.F(true),
					}),
					CountryPools: cfrex.F(map[string][]string{
						"GB": {"abd90f38ced07c2e2f4df50b1f61d4194"},
						"US": {"de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					DefaultPools: cfrex.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cfrex.F("fallback_pool"),
					LocationStrategy: cfrex.F(cfrex.LocationStrategyParam{
						Mode:      cfrex.F(cfrex.LocationStrategyModeResolverIP),
						PreferEcs: cfrex.F(cfrex.LocationStrategyPreferEcsAlways),
					}),
					PopPools: cfrex.F(map[string][]string{
						"LAX": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
						"LHR": {"abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"},
						"SJC": {"00920f38ce07c2e2f4df50b1f61d4194"},
					}),
					RandomSteering: cfrex.F(cfrex.RandomSteeringParam{
						DefaultWeight: cfrex.F(0.200000),
						PoolWeights: cfrex.F(map[string]float64{
							"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
							"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
						}),
					}),
					RegionPools: cfrex.F(map[string][]string{
						"ENAM": {"00920f38ce07c2e2f4df50b1f61d4194"},
						"WNAM": {"de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"},
					}),
					SessionAffinity: cfrex.F(cfrex.SessionAffinityCookie),
					SessionAffinityAttributes: cfrex.F(cfrex.SessionAffinityAttributesParam{
						DrainDuration:        cfrex.F(100.000000),
						Headers:              cfrex.F([]string{"_1K--W2kIFj1"}),
						RequireAllHeaders:    cfrex.F(true),
						Samesite:             cfrex.F(cfrex.SessionAffinityAttributesSamesiteAuto),
						Secure:               cfrex.F(cfrex.SessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cfrex.F(cfrex.SessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTtl: cfrex.F(1800.000000),
					SteeringPolicy:     cfrex.F(cfrex.SteeringPolicyDynamicLatency),
					Ttl:                cfrex.F(30.000000),
				}),
				Priority:   cfrex.F(int64(0)),
				Terminates: cfrex.F(true),
			}}),
			SessionAffinity: cfrex.F(cfrex.SessionAffinityCookie),
			SessionAffinityAttributes: cfrex.F(cfrex.SessionAffinityAttributesParam{
				DrainDuration:        cfrex.F(100.000000),
				Headers:              cfrex.F([]string{"_1K--W2kIFj1"}),
				RequireAllHeaders:    cfrex.F(true),
				Samesite:             cfrex.F(cfrex.SessionAffinityAttributesSamesiteAuto),
				Secure:               cfrex.F(cfrex.SessionAffinityAttributesSecureAuto),
				ZeroDowntimeFailover: cfrex.F(cfrex.SessionAffinityAttributesZeroDowntimeFailoverSticky),
			}),
			SessionAffinityTtl: cfrex.F(1800.000000),
			SteeringPolicy:     cfrex.F(cfrex.SteeringPolicyDynamicLatency),
			Ttl:                cfrex.F(30.000000),
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
