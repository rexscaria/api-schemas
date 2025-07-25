// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"
)

type UnionTime time.Time

func (UnionTime) ImplementsUserListAuditLogsParamsBeforeUnion() {}
func (UnionTime) ImplementsUserListAuditLogsParamsSinceUnion()  {}

type UnionString string

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
func (UnionString) ImplementsAccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersValueUnion() {}
func (UnionString) ImplementsAccountAIGatewayGatewayLogListGatewayLogsParamsFiltersValueUnion()   {}
func (UnionString) ImplementsAccountAIGatewayGatewayLogPatchGatewayLogParamsMetadataUnion()       {}
func (UnionString) ImplementsAccountAIRunExecuteModelResponseResultUnion()                        {}
func (UnionString) ImplementsAccountAIRunExecuteModelParamsBodyTextEmbeddingsTextUnion()          {}
func (UnionString) ImplementsAccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectTextUnion()    {}
func (UnionString) ImplementsAccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequestsTextUnion() {
}
func (UnionString) ImplementsAccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectTextUnion() {}
func (UnionString) ImplementsAccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequestsTextUnion() {
}
func (UnionString) ImplementsAccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputEmbeddingTextUnion() {}
func (UnionString) ImplementsAccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputEmbeddingTextUnion() {
}
func (UnionString) ImplementsAccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectTextUnion() {}
func (UnionString) ImplementsAccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequestsTextUnion() {
}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyPromptImageUnion() {
}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentUnion() {
}
func (UnionString) ImplementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesImageUnion() {
}
func (UnionString) ImplementsAccountBrowserRenderingGetPdfParamsPdfOptionsHeightUnion()       {}
func (UnionString) ImplementsAccountBrowserRenderingGetPdfParamsPdfOptionsMarginBottomUnion() {}
func (UnionString) ImplementsAccountBrowserRenderingGetPdfParamsPdfOptionsMarginLeftUnion()   {}
func (UnionString) ImplementsAccountBrowserRenderingGetPdfParamsPdfOptionsMarginRightUnion()  {}
func (UnionString) ImplementsAccountBrowserRenderingGetPdfParamsPdfOptionsMarginTopUnion()    {}
func (UnionString) ImplementsAccountBrowserRenderingGetPdfParamsPdfOptionsWidthUnion()        {}
func (UnionString) ImplementsEmailRuleConditionValueUnionParam()                              {}
func (UnionString) ImplementsEmailRuleConditionValueUnion()                                   {}
func (UnionString) ImplementsMagicHealthCheckBaseTargetUnionParam()                           {}
func (UnionString) ImplementsMagicHealthCheckBaseTargetUnion()                                {}
func (UnionString) ImplementsMagicTunnelHealthCheckTargetUnionParam()                         {}
func (UnionString) ImplementsMagicTunnelHealthCheckTargetUnion()                              {}
func (UnionString) ImplementsAccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultValuesUnion() {
}
func (UnionString) ImplementsAccountWorkflowInstanceGetResponseResultOutputUnion()      {}
func (UnionString) ImplementsRadarRankingGetTimeseriesGroupsResponseResultSerie0Union() {}
func (UnionString) ImplementsRadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0Union() {
}
func (UnionString) ImplementsZoneAnalyticsSinceUnion()                                       {}
func (UnionString) ImplementsZoneAnalyticsUntilUnionParam()                                  {}
func (UnionString) ImplementsZoneAnalyticsUntilUnion()                                       {}
func (UnionString) ImplementsZoneAnalyticsListColosParamsSinceUnion()                        {}
func (UnionString) ImplementsZoneAnalyticsGetDashboardParamsSinceUnion()                     {}
func (UnionString) ImplementsTlsSettingValueUnionParam()                                     {}
func (UnionString) ImplementsTlsSettingValueUnion()                                          {}
func (UnionString) ImplementsZoneLogReceivedGetLogsParamsEndUnion()                          {}
func (UnionString) ImplementsZoneLogReceivedGetLogsParamsStartUnion()                        {}
func (UnionString) ImplementsZarazConfigReturnToolsZarazManagedComponentDefaultFieldsUnion() {}
func (UnionString) ImplementsZarazConfigReturnToolsZarazManagedComponentSettingsUnion()      {}
func (UnionString) ImplementsZarazConfigReturnToolsWorkerDefaultFieldsUnion()                {}
func (UnionString) ImplementsZarazConfigReturnToolsWorkerSettingsUnion()                     {}
func (UnionString) ImplementsZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentDefaultFieldsUnion() {
}
func (UnionString) ImplementsZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentSettingsUnion() {
}
func (UnionString) ImplementsZoneSettingZarazConfigUpdateParamsToolsWorkerDefaultFieldsUnion() {}
func (UnionString) ImplementsZoneSettingZarazConfigUpdateParamsToolsWorkerSettingsUnion()      {}
func (UnionString) ImplementsAppConfigOriginPortUnionParam()                                   {}
func (UnionString) ImplementsAppConfigOriginPortUnion()                                        {}

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
func (UnionBool) ImplementsAccountWorkerDispatchNamespaceScriptUploadParamsMetadataAssetsConfigRunWorkerFirstUnion() {
}
func (UnionBool) ImplementsAccountWorkerScriptUploadParamsMetadataAssetsConfigRunWorkerFirstUnion() {}
func (UnionBool) ImplementsZarazConfigReturnToolsZarazManagedComponentDefaultFieldsUnion()          {}
func (UnionBool) ImplementsZarazConfigReturnToolsZarazManagedComponentSettingsUnion()               {}
func (UnionBool) ImplementsZarazConfigReturnToolsWorkerDefaultFieldsUnion()                         {}
func (UnionBool) ImplementsZarazConfigReturnToolsWorkerSettingsUnion()                              {}
func (UnionBool) ImplementsZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentDefaultFieldsUnion() {
}
func (UnionBool) ImplementsZoneSettingZarazConfigUpdateParamsToolsZarazManagedComponentSettingsUnion() {
}
func (UnionBool) ImplementsZoneSettingZarazConfigUpdateParamsToolsWorkerDefaultFieldsUnion() {}
func (UnionBool) ImplementsZoneSettingZarazConfigUpdateParamsToolsWorkerSettingsUnion()      {}

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
func (UnionFloat) ImplementsAccountBrowserRenderingGetPdfParamsPdfOptionsHeightUnion()           {}
func (UnionFloat) ImplementsAccountBrowserRenderingGetPdfParamsPdfOptionsMarginBottomUnion()     {}
func (UnionFloat) ImplementsAccountBrowserRenderingGetPdfParamsPdfOptionsMarginLeftUnion()       {}
func (UnionFloat) ImplementsAccountBrowserRenderingGetPdfParamsPdfOptionsMarginRightUnion()      {}
func (UnionFloat) ImplementsAccountBrowserRenderingGetPdfParamsPdfOptionsMarginTopUnion()        {}
func (UnionFloat) ImplementsAccountBrowserRenderingGetPdfParamsPdfOptionsWidthUnion()            {}
func (UnionFloat) ImplementsAccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultValuesUnion() {
}
func (UnionFloat) ImplementsAccountWorkflowInstanceGetResponseResultOutputUnion()      {}
func (UnionFloat) ImplementsRadarRankingGetTimeseriesGroupsResponseResultSerie0Union() {}
func (UnionFloat) ImplementsRadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0Union() {
}
func (UnionFloat) ImplementsTlsSettingValueUnionParam()                         {}
func (UnionFloat) ImplementsTlsSettingValueUnion()                              {}
func (UnionFloat) ImplementsZoneSettingUpdateSettingParamsBodyValueValueUnion() {}
