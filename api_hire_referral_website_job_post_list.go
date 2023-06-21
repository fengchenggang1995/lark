// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// GetHireReferralWebsiteJobPostList 获取内推官网下的职位列表。自定义数据暂不支持列表获取, 请从「获取内推官网下职位广告详情」接口获取。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/referral_website-job_post/list
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/get-candidates/referral/list
func (r *HireService) GetHireReferralWebsiteJobPostList(ctx context.Context, request *GetHireReferralWebsiteJobPostListReq, options ...MethodOptionFunc) (*GetHireReferralWebsiteJobPostListResp, *Response, error) {
	if r.cli.mock.mockHireGetHireReferralWebsiteJobPostList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#GetHireReferralWebsiteJobPostList mock enable")
		return r.cli.mock.mockHireGetHireReferralWebsiteJobPostList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireReferralWebsiteJobPostList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/referral_websites/job_posts",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireReferralWebsiteJobPostListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireGetHireReferralWebsiteJobPostList mock HireGetHireReferralWebsiteJobPostList method
func (r *Mock) MockHireGetHireReferralWebsiteJobPostList(f func(ctx context.Context, request *GetHireReferralWebsiteJobPostListReq, options ...MethodOptionFunc) (*GetHireReferralWebsiteJobPostListResp, *Response, error)) {
	r.mockHireGetHireReferralWebsiteJobPostList = f
}

// UnMockHireGetHireReferralWebsiteJobPostList un-mock HireGetHireReferralWebsiteJobPostList method
func (r *Mock) UnMockHireGetHireReferralWebsiteJobPostList() {
	r.mockHireGetHireReferralWebsiteJobPostList = nil
}

// GetHireReferralWebsiteJobPostListReq ...
type GetHireReferralWebsiteJobPostListReq struct {
	ProcessType      *int64            `query:"process_type" json:"-"`       // 招聘流程类型, 示例值: 1, 可选值有: 1: 社招, 2: 校招
	PageToken        *string           `query:"page_token" json:"-"`         // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: eyJvZmZzZXQiOjEwLCJ0aW1lc3RhbXAiOjE2Mjc1NTUyMjM2NzIsImlkIjpudWxsfQ[
	PageSize         *int64            `query:"page_size" json:"-"`          // 分页大小, 示例值: 10, 最大值: `10`
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门 ID 的类型, 示例值: open_department_id, 可选值有: open_department_id: 以 open_department_id 来标识部门, department_id: 以 department_id 来标识部门, 默认值: `open_department_id`
}

// GetHireReferralWebsiteJobPostListResp ...
type GetHireReferralWebsiteJobPostListResp struct {
	Items     []*GetHireReferralWebsiteJobPostListRespItem `json:"items,omitempty"`      // 列表
	HasMore   bool                                         `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                                       `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetHireReferralWebsiteJobPostListRespItem ...
type GetHireReferralWebsiteJobPostListRespItem struct {
	ID                 string                                                       `json:"id,omitempty"`                   // 职位广告 ID
	Title              string                                                       `json:"title,omitempty"`                // 标题
	JobID              string                                                       `json:"job_id,omitempty"`               // 职位 ID
	JobCode            string                                                       `json:"job_code,omitempty"`             // 职位编码
	JobExpireTime      string                                                       `json:"job_expire_time,omitempty"`      // 职位过期时间, 「null」代表「长期有效」
	JobActiveStatus    int64                                                        `json:"job_active_status,omitempty"`    // 职位状态, 可选值有: 1: 启用态, 2: 禁用态
	JobProcessType     int64                                                        `json:"job_process_type,omitempty"`     // 职位流程类型, 可选值有: 1: 社招, 2: 校招
	JobRecruitmentType *GetHireReferralWebsiteJobPostListRespItemJobRecruitmentType `json:"job_recruitment_type,omitempty"` // 职位雇佣类型
	JobDepartment      *GetHireReferralWebsiteJobPostListRespItemJobDepartment      `json:"job_department,omitempty"`       // 职位部门
	JobType            *GetHireReferralWebsiteJobPostListRespItemJobType            `json:"job_type,omitempty"`             // 职位类型
	Address            *GetHireReferralWebsiteJobPostListRespItemAddress            `json:"address,omitempty"`              // 职位地址
	MinSalary          string                                                       `json:"min_salary,omitempty"`           // 月薪范围-最低薪资
	MaxSalary          string                                                       `json:"max_salary,omitempty"`           // 月薪范围-最高薪资
	RequiredDegree     int64                                                        `json:"required_degree,omitempty"`      // 学历要求, 可选值有: 1: 小学及以上, 2: 初中及以上, 3: 专职及以上, 4: 高中及以上, 5: 大专及以上, 6: 本科及以上, 7: 硕士及以上, 8: 博士及以上, 20: 不限
	Experience         int64                                                        `json:"experience,omitempty"`           // 经验, 可选值有: 1: 不限, 2: 应届毕业生, 3: 1年以下, 4: 1-3年, 5: 3-5年, 6: 5-7年, 7: 7-10年, 8: 10年以上
	Headcount          int64                                                        `json:"headcount,omitempty"`            // 数量
	HighLightList      []*GetHireReferralWebsiteJobPostListRespItemHighLight        `json:"high_light_list,omitempty"`      // 职位亮点
	Description        string                                                       `json:"description,omitempty"`          // 职位描述
	Requirement        string                                                       `json:"requirement,omitempty"`          // 职位要求
	Creator            *GetHireReferralWebsiteJobPostListRespItemCreator            `json:"creator,omitempty"`              // 创建人
	CreateTime         string                                                       `json:"create_time,omitempty"`          // 创建时间
	ModifyTime         string                                                       `json:"modify_time,omitempty"`          // 修改时间
	CustomizedDataList []*GetHireReferralWebsiteJobPostListRespItemCustomizedData   `json:"customized_data_list,omitempty"` // 自定义字段
	AddressList        []*GetHireReferralWebsiteJobPostListRespItemAddress          `json:"address_list,omitempty"`         // 职位广告地址列表
}

// GetHireReferralWebsiteJobPostListRespItemAddress ...
type GetHireReferralWebsiteJobPostListRespItemAddress struct {
	ID       string                                                    `json:"id,omitempty"`       // ID
	Name     *GetHireReferralWebsiteJobPostListRespItemAddressName     `json:"name,omitempty"`     // 名称
	District *GetHireReferralWebsiteJobPostListRespItemAddressDistrict `json:"district,omitempty"` // 区域信息
	City     *GetHireReferralWebsiteJobPostListRespItemAddressCity     `json:"city,omitempty"`     // 城市信息
	State    *GetHireReferralWebsiteJobPostListRespItemAddressState    `json:"state,omitempty"`    // 省信息
	Country  *GetHireReferralWebsiteJobPostListRespItemAddressCountry  `json:"country,omitempty"`  // 国家信息
}

// GetHireReferralWebsiteJobPostListRespItemAddressCity ...
type GetHireReferralWebsiteJobPostListRespItemAddressCity struct {
	Code string                                                    `json:"code,omitempty"` // 编码
	Name *GetHireReferralWebsiteJobPostListRespItemAddressCityName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostListRespItemAddressCityName ...
type GetHireReferralWebsiteJobPostListRespItemAddressCityName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostListRespItemAddressCountry ...
type GetHireReferralWebsiteJobPostListRespItemAddressCountry struct {
	Code string                                                       `json:"code,omitempty"` // 编码
	Name *GetHireReferralWebsiteJobPostListRespItemAddressCountryName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostListRespItemAddressCountryName ...
type GetHireReferralWebsiteJobPostListRespItemAddressCountryName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostListRespItemAddressDistrict ...
type GetHireReferralWebsiteJobPostListRespItemAddressDistrict struct {
	Code string                                                        `json:"code,omitempty"` // 编码
	Name *GetHireReferralWebsiteJobPostListRespItemAddressDistrictName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostListRespItemAddressDistrictName ...
type GetHireReferralWebsiteJobPostListRespItemAddressDistrictName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostListRespItemAddressName ...
type GetHireReferralWebsiteJobPostListRespItemAddressName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostListRespItemAddressState ...
type GetHireReferralWebsiteJobPostListRespItemAddressState struct {
	Code string                                                     `json:"code,omitempty"` // 编码
	Name *GetHireReferralWebsiteJobPostListRespItemAddressStateName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostListRespItemAddressStateName ...
type GetHireReferralWebsiteJobPostListRespItemAddressStateName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostListRespItemCreator ...
type GetHireReferralWebsiteJobPostListRespItemCreator struct {
	ID   string                                                `json:"id,omitempty"`   // ID
	Name *GetHireReferralWebsiteJobPostListRespItemCreatorName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostListRespItemCreatorName ...
type GetHireReferralWebsiteJobPostListRespItemCreatorName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostListRespItemCustomizedData ...
type GetHireReferralWebsiteJobPostListRespItemCustomizedData struct {
	ObjectID   string                                                        `json:"object_id,omitempty"`   // 自定义字段 ID
	Name       *GetHireReferralWebsiteJobPostListRespItemCustomizedDataName  `json:"name,omitempty"`        // 字段名称
	ObjectType int64                                                         `json:"object_type,omitempty"` // 字段类型, 可选值有: 1: 单行文本, 2: 多行文本, 3: 单选, 4: 多选, 5: 日期, 6: 月份选择, 7: 年份选择, 8: 时间段, 9: 数字, 10: 默认字段
	Value      *GetHireReferralWebsiteJobPostListRespItemCustomizedDataValue `json:"value,omitempty"`       // 自定义字段值
}

// GetHireReferralWebsiteJobPostListRespItemCustomizedDataName ...
type GetHireReferralWebsiteJobPostListRespItemCustomizedDataName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostListRespItemCustomizedDataValue ...
type GetHireReferralWebsiteJobPostListRespItemCustomizedDataValue struct {
	Content    string                                                                 `json:"content,omitempty"`     // 当字段类型为单行文本、多行文本、模块、默认字段时, 从此字段取值
	Option     *GetHireReferralWebsiteJobPostListRespItemCustomizedDataValueOption    `json:"option,omitempty"`      // 当字段类型为单选时, 从此字段取值
	OptionList []*GetHireReferralWebsiteJobPostListRespItemCustomizedDataValueOption  `json:"option_list,omitempty"` // 当字段类型为多选时, 从此字段取值
	TimeRange  *GetHireReferralWebsiteJobPostListRespItemCustomizedDataValueTimeRange `json:"time_range,omitempty"`  // 当字段类型为时间段时, 从此字段取值
	Time       string                                                                 `json:"time,omitempty"`        // 当字段类型为日期选择、月份选择、年份选择时, 从此字段取值, 该字段是毫秒级时间戳
	Number     string                                                                 `json:"number,omitempty"`      // 当字段类型为数字时, 从此字段取值
}

// GetHireReferralWebsiteJobPostListRespItemCustomizedDataValueOption ...
type GetHireReferralWebsiteJobPostListRespItemCustomizedDataValueOption struct {
	Key  string                                                                  `json:"key,omitempty"`  // 选项 ID
	Name *GetHireReferralWebsiteJobPostListRespItemCustomizedDataValueOptionName `json:"name,omitempty"` // 选项名称
}

// GetHireReferralWebsiteJobPostListRespItemCustomizedDataValueOptionName ...
type GetHireReferralWebsiteJobPostListRespItemCustomizedDataValueOptionName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostListRespItemCustomizedDataValueTimeRange ...
type GetHireReferralWebsiteJobPostListRespItemCustomizedDataValueTimeRange struct {
	StartTime string `json:"start_time,omitempty"` // 开始时间
	EndTime   string `json:"end_time,omitempty"`   // 结束时间
}

// GetHireReferralWebsiteJobPostListRespItemHighLight ...
type GetHireReferralWebsiteJobPostListRespItemHighLight struct {
	ID   string                                                  `json:"id,omitempty"`   // ID
	Name *GetHireReferralWebsiteJobPostListRespItemHighLightName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostListRespItemHighLightName ...
type GetHireReferralWebsiteJobPostListRespItemHighLightName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostListRespItemJobDepartment ...
type GetHireReferralWebsiteJobPostListRespItemJobDepartment struct {
	ID   string                                                      `json:"id,omitempty"`   // ID
	Name *GetHireReferralWebsiteJobPostListRespItemJobDepartmentName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostListRespItemJobDepartmentName ...
type GetHireReferralWebsiteJobPostListRespItemJobDepartmentName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostListRespItemJobRecruitmentType ...
type GetHireReferralWebsiteJobPostListRespItemJobRecruitmentType struct {
	ID   string                                                           `json:"id,omitempty"`   // ID
	Name *GetHireReferralWebsiteJobPostListRespItemJobRecruitmentTypeName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostListRespItemJobRecruitmentTypeName ...
type GetHireReferralWebsiteJobPostListRespItemJobRecruitmentTypeName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostListRespItemJobType ...
type GetHireReferralWebsiteJobPostListRespItemJobType struct {
	ID   string                                                `json:"id,omitempty"`   // ID
	Name *GetHireReferralWebsiteJobPostListRespItemJobTypeName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostListRespItemJobTypeName ...
type GetHireReferralWebsiteJobPostListRespItemJobTypeName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// getHireReferralWebsiteJobPostListResp ...
type getHireReferralWebsiteJobPostListResp struct {
	Code int64                                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                 `json:"msg,omitempty"`  // 错误描述
	Data *GetHireReferralWebsiteJobPostListResp `json:"data,omitempty"`
}