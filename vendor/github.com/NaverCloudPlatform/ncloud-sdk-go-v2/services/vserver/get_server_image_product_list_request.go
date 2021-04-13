/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type GetServerImageProductListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// REGION코드
BlockStorageSize *int32 `json:"blockStorageSize,omitempty"`

	// 제외할상품코드
ExclusionProductCode *string `json:"exclusionProductCode,omitempty"`

	// 상품코드
ProductCode *string `json:"productCode,omitempty"`

	// 플랫폼유형코드리스트
PlatformTypeCodeList []*string `json:"platformTypeCodeList,omitempty"`
}
