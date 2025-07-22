// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"
)

type UnionTime time.Time

func (UnionTime) ImplementsUserListAuditLogsParamsBeforeUnion() {}
func (UnionTime) ImplementsUserListAuditLogsParamsSinceUnion()  {}

type UnionString string

func (UnionString) ImplementsEmbeddedPoliciesPoliciesUnionParam()            {}
func (UnionString) ImplementsAccountAccessPolicyTestStartParamsPolicyUnion() {}
func (UnionString) ImplementsAccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersValueUnion() {
}
func (UnionString) ImplementsAccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersValueUnion() {
}
func (UnionString) ImplementsAccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersValueUnion() {
}
func (UnionString) ImplementsAccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersValueUnion() {
}
func (UnionString) ImplementsAccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersValueUnion() {
}
func (UnionString) ImplementsAccountAIGatewayGatewayDatasetNewDatasetParamsFiltersValueUnion()    {}
func (UnionString) ImplementsAccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersValueUnion() {}
func (UnionString) ImplementsAccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersValueUnion() {
}
func (UnionString) ImplementsAccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersValueUnion() {
}
func (UnionString) ImplementsAccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersValueUnion() {
}
func (UnionString) ImplementsAccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersValueUnion() {
}
func (UnionString) ImplementsAccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersValueUnion()   {}
func (UnionString) ImplementsAccountAIGatewayGatewayLogListGatewayLogsParamsFiltersValueUnion()     {}
func (UnionString) ImplementsAccountAIGatewayGatewayLogPatchGatewayLogParamsMetadataUnion()         {}
func (UnionString) ImplementsAccountAIRunExecuteModelResponseResultUnion()                          {}
func (UnionString) ImplementsAccountAIRunExecuteModelParamsBodyUnion()                              {}
func (UnionString) ImplementsAccountAIRunExecuteModelParamsBodyTextEmbeddingsTextUnion()            {}
func (UnionString) ImplementsAccountAIRunExecuteModelParamsBodyObjectImageUnion()                   {}
func (UnionString) ImplementsAccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsTextUnion()                {}
func (UnionString) ImplementsAccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsTextUnion()               {}
func (UnionString) ImplementsAccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputEmbeddingTextUnion() {}
func (UnionString) ImplementsAccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsTextUnion()               {}
func (UnionString) ImplementsAccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructResponseResultUnion() {
}
func (UnionString) ImplementsAccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bResponseResultUnion() {
}
func (UnionString) ImplementsAccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultUnion()             {}
func (UnionString) ImplementsAccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ResponseResultUnion()   {}
func (UnionString) ImplementsAccountAIRunCfGoogleExecuteGemma2bItLoraResponseResultUnion()          {}
func (UnionString) ImplementsAccountAIRunCfGoogleExecuteGemma7bItLoraResponseResultUnion()          {}
func (UnionString) ImplementsAccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraResponseResultUnion() {}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultUnion()        {}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultUnion()        {}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultUnion()     {}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultUnion() {
}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultUnion()     {}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultUnion()     {}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultUnion()  {}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultUnion() {}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultUnion()  {}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultUnion()      {}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultUnion() {
}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultUnion() {}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultUnion() {}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultUnion() {
}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultUnion()      {}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultUnion()   {}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultResponseUnion() {}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyPromptImageUnion() {
}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesImageUnion() {
}
func (UnionString) ImplementsAccountAIRunCfMicrosoftExecutePhi2ResponseResultUnion()                {}
func (UnionString) ImplementsAccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultUnion() {}
func (UnionString) ImplementsAccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultUnion() {
}
func (UnionString) ImplementsAccountAIRunCfMyshellAIExecuteMelottsResponseResultUnion()             {}
func (UnionString) ImplementsAccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultUnion()     {}
func (UnionString) ImplementsAccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultUnion()         {}
func (UnionString) ImplementsAccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultUnion()         {}
func (UnionString) ImplementsAccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultUnion()       {}
func (UnionString) ImplementsAccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultUnion()        {}
func (UnionString) ImplementsAccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqResponseResultUnion() {}
func (UnionString) ImplementsAccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultUnion()       {}
func (UnionString) ImplementsAccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultUnion() {
}
func (UnionString) ImplementsAccountAIRunHfGoogleExecuteGemma7bItResponseResultUnion() {}
func (UnionString) ImplementsAccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultUnion() {
}
func (UnionString) ImplementsAccountAIRunHfMistralExecuteMistral7bInstructV0_2ResponseResultUnion() {}
func (UnionString) ImplementsAccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultUnion() {
}
func (UnionString) ImplementsAccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultUnion() {}
func (UnionString) ImplementsAccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultUnion() {
}
func (UnionString) ImplementsAccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqResponseResultUnion() {
}
func (UnionString) ImplementsAccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqResponseResultUnion() {
}
func (UnionString) ImplementsAccountAIRunHfTheblokeExecuteLlama2_13bChatAwqResponseResultUnion() {}
func (UnionString) ImplementsAccountAIRunHfTheblokeExecuteLlamaguard7bAwqResponseResultUnion()   {}
func (UnionString) ImplementsAccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqResponseResultUnion() {
}
func (UnionString) ImplementsAccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqResponseResultUnion() {}
func (UnionString) ImplementsAccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqResponseResultUnion() {
}
func (UnionString) ImplementsAccountAIRunHfTheblokeExecuteZephyr7bBetaAwqResponseResultUnion() {}
func (UnionString) ImplementsAPIResponseTunnelResultUnion()                                    {}
func (UnionString) ImplementsAPIResponseTeamsDevicesResultUnion()                              {}
func (UnionString) ImplementsAccountDiagnosticRunTracerouteResponseResultUnion()               {}
func (UnionString) ImplementsEmailRuleConditionValueUnionParam()                               {}
func (UnionString) ImplementsEmailRuleConditionValueUnion()                                    {}
func (UnionString) ImplementsFirewallAPIResponseCommonResultUnion()                            {}
func (UnionString) ImplementsAPIResponseImagesResultUnion()                                    {}
func (UnionString) ImplementsCollectionResponseResultUnion()                                   {}
func (UnionString) ImplementsMagicAPIResponseSingleResultUnion()                               {}
func (UnionString) ImplementsMagicHealthCheckBaseTargetUnionParam()                            {}
func (UnionString) ImplementsMagicHealthCheckBaseTargetUnion()                                 {}
func (UnionString) ImplementsAPIResponseMagicVisibilityMnmResultUnion()                        {}
func (UnionString) ImplementsAPIResponseMagicVisibilityPcapsResultUnion()                      {}
func (UnionString) ImplementsRegistrarApiapiResponseCommonResultUnion()                        {}
func (UnionString) ImplementsAPIResponseResourceSharingResultUnion()                           {}
func (UnionString) ImplementsAccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultValuesUnion() {
}
func (UnionString) ImplementsAccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadataValuesValueUnion() {
}
func (UnionString) ImplementsAPIResponseBillingResultUnion()                       {}
func (UnionString) ImplementsVectorizeAPIResponseCommonResultUnion()               {}
func (UnionString) ImplementsAccountWorkflowInstanceGetResponseResultOutputUnion() {}
func (UnionString) ImplementsAccountWorkflowInstanceGetResponseResultStepsObjectConfigRetriesDelayUnion() {
}
func (UnionString) ImplementsAccountWorkflowInstanceGetResponseResultStepsObjectConfigTimeoutUnion() {
}
func (UnionString) ImplementsAPIResponseCustomPagesResultUnion()                        {}
func (UnionString) ImplementsRadarRankingGetTimeseriesGroupsResponseResultSerie0Union() {}
func (UnionString) ImplementsRadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0Union() {
}
func (UnionString) ImplementsZoneAnalyticsAPIResponseResultUnion()                             {}
func (UnionString) ImplementsZoneAnalyticsSinceUnion()                                         {}
func (UnionString) ImplementsZoneAnalyticsUntilUnionParam()                                    {}
func (UnionString) ImplementsZoneAnalyticsUntilUnion()                                         {}
func (UnionString) ImplementsZoneAnalyticsListColosParamsSinceUnion()                          {}
func (UnionString) ImplementsZoneAnalyticsGetDashboardParamsSinceUnion()                       {}
func (UnionString) ImplementsAPIResponseCommonResultUnion()                                    {}
func (UnionString) ImplementsWafManagedRulesAPIResponseCommonResultUnion()                     {}
func (UnionString) ImplementsHealthcheckAPICommonResultUnion()                                 {}
func (UnionString) ImplementsAPIResponseSingleZonesSchemasResultUnion()                        {}
func (UnionString) ImplementsTlsSettingValueUnionParam()                                       {}
func (UnionString) ImplementsTlsSettingValueUnion()                                            {}
func (UnionString) ImplementsAPIResponseWafProductBundleResultUnion()                          {}
func (UnionString) ImplementsZoneLogReceivedGetLogsParamsEndUnion()                            {}
func (UnionString) ImplementsZoneLogReceivedGetLogsParamsStartUnion()                          {}
func (UnionString) ImplementsZarazConfigReturnToolsObjectDefaultFieldsUnion()                  {}
func (UnionString) ImplementsZarazConfigReturnToolsObjectSettingsUnion()                       {}
func (UnionString) ImplementsZoneSettingZarazConfigUpdateParamsToolsObjectDefaultFieldsUnion() {}
func (UnionString) ImplementsZoneSettingZarazConfigUpdateParamsToolsObjectSettingsUnion()      {}
func (UnionString) ImplementsAppConfigOriginPortUnionParam()                                   {}
func (UnionString) ImplementsAppConfigOriginPortUnion()                                        {}
func (UnionString) ImplementsAPIResponseWeb3ResultUnion()                                      {}

type UnionBool bool

func (UnionBool) ImplementsAccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersValueUnion() {
}
func (UnionBool) ImplementsAccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersValueUnion() {
}
func (UnionBool) ImplementsAccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersValueUnion() {
}
func (UnionBool) ImplementsAccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersValueUnion() {
}
func (UnionBool) ImplementsAccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersValueUnion() {
}
func (UnionBool) ImplementsAccountAIGatewayGatewayDatasetNewDatasetParamsFiltersValueUnion()    {}
func (UnionBool) ImplementsAccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersValueUnion() {}
func (UnionBool) ImplementsAccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersValueUnion() {
}
func (UnionBool) ImplementsAccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersValueUnion() {
}
func (UnionBool) ImplementsAccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersValueUnion() {
}
func (UnionBool) ImplementsAccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersValueUnion() {
}
func (UnionBool) ImplementsAccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersValueUnion() {}
func (UnionBool) ImplementsAccountAIGatewayGatewayLogListGatewayLogsParamsFiltersValueUnion()   {}
func (UnionBool) ImplementsAccountAIGatewayGatewayLogPatchGatewayLogParamsMetadataUnion()       {}
func (UnionBool) ImplementsAccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultValuesUnion() {
}
func (UnionBool) ImplementsAccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadataValuesValueUnion() {
}
func (UnionBool) ImplementsZarazConfigReturnToolsObjectDefaultFieldsUnion()                  {}
func (UnionBool) ImplementsZarazConfigReturnToolsObjectSettingsUnion()                       {}
func (UnionBool) ImplementsZoneSettingZarazConfigUpdateParamsToolsObjectDefaultFieldsUnion() {}
func (UnionBool) ImplementsZoneSettingZarazConfigUpdateParamsToolsObjectSettingsUnion()      {}

type UnionInt int64

func (UnionInt) ImplementsZoneAnalyticsSinceUnion()                   {}
func (UnionInt) ImplementsZoneAnalyticsUntilUnionParam()              {}
func (UnionInt) ImplementsZoneAnalyticsUntilUnion()                   {}
func (UnionInt) ImplementsZoneAnalyticsListColosParamsSinceUnion()    {}
func (UnionInt) ImplementsZoneAnalyticsGetDashboardParamsSinceUnion() {}
func (UnionInt) ImplementsZoneLogReceivedGetLogsParamsEndUnion()      {}
func (UnionInt) ImplementsZoneLogReceivedGetLogsParamsStartUnion()    {}
func (UnionInt) ImplementsZoneActionCacheTtlByStatusValueUnionParam() {}
func (UnionInt) ImplementsZoneActionCacheTtlByStatusValueUnion()      {}
func (UnionInt) ImplementsAppConfigOriginPortUnionParam()             {}
func (UnionInt) ImplementsAppConfigOriginPortUnion()                  {}

type UnionFloat float64

func (UnionFloat) ImplementsAccountAIGatewayGatewayDatasetNewDatasetResponseResultFiltersValueUnion() {
}
func (UnionFloat) ImplementsAccountAIGatewayGatewayDatasetDeleteDatasetResponseResultFiltersValueUnion() {
}
func (UnionFloat) ImplementsAccountAIGatewayGatewayDatasetFetchDatasetResponseResultFiltersValueUnion() {
}
func (UnionFloat) ImplementsAccountAIGatewayGatewayDatasetListDatasetsResponseResultFiltersValueUnion() {
}
func (UnionFloat) ImplementsAccountAIGatewayGatewayDatasetUpdateDatasetResponseResultFiltersValueUnion() {
}
func (UnionFloat) ImplementsAccountAIGatewayGatewayDatasetNewDatasetParamsFiltersValueUnion()    {}
func (UnionFloat) ImplementsAccountAIGatewayGatewayDatasetUpdateDatasetParamsFiltersValueUnion() {}
func (UnionFloat) ImplementsAccountAIGatewayGatewayEvaluationNewEvaluationResponseResultDatasetsFiltersValueUnion() {
}
func (UnionFloat) ImplementsAccountAIGatewayGatewayEvaluationDeleteEvaluationResponseResultDatasetsFiltersValueUnion() {
}
func (UnionFloat) ImplementsAccountAIGatewayGatewayEvaluationFetchEvaluationResponseResultDatasetsFiltersValueUnion() {
}
func (UnionFloat) ImplementsAccountAIGatewayGatewayEvaluationListEvaluationsResponseResultDatasetsFiltersValueUnion() {
}
func (UnionFloat) ImplementsAccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersValueUnion() {}
func (UnionFloat) ImplementsAccountAIGatewayGatewayLogListGatewayLogsParamsFiltersValueUnion()   {}
func (UnionFloat) ImplementsAccountAIGatewayGatewayLogPatchGatewayLogParamsMetadataUnion()       {}
func (UnionFloat) ImplementsAccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultValuesUnion() {
}
func (UnionFloat) ImplementsAccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadataValuesValueUnion() {
}
func (UnionFloat) ImplementsAccountWorkflowInstanceGetResponseResultOutputUnion() {}
func (UnionFloat) ImplementsAccountWorkflowInstanceGetResponseResultStepsObjectConfigRetriesDelayUnion() {
}
func (UnionFloat) ImplementsAccountWorkflowInstanceGetResponseResultStepsObjectConfigTimeoutUnion() {}
func (UnionFloat) ImplementsRadarRankingGetTimeseriesGroupsResponseResultSerie0Union()              {}
func (UnionFloat) ImplementsRadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0Union() {
}
func (UnionFloat) ImplementsTlsSettingValueUnionParam() {}
func (UnionFloat) ImplementsTlsSettingValueUnion()      {}
