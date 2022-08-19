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

// BatchGetUser 为了更好地提升该接口的安全性, 我们对其进行了升级, 请尽快迁移至[新版本>>](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/get)
//
// 该接口用于批量获取用户详细信息。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uIzNz4iM3MjLyczM
//
// Deprecated
func (r *ContactService) BatchGetUser(ctx context.Context, request *BatchGetUserReq, options ...MethodOptionFunc) (*BatchGetUserResp, *Response, error) {
	if r.cli.mock.mockContactBatchGetUser != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#BatchGetUser mock enable")
		return r.cli.mock.mockContactBatchGetUser(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "BatchGetUser",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v1/user/batch_get",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchGetUserResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactBatchGetUser mock ContactBatchGetUser method
func (r *Mock) MockContactBatchGetUser(f func(ctx context.Context, request *BatchGetUserReq, options ...MethodOptionFunc) (*BatchGetUserResp, *Response, error)) {
	r.mockContactBatchGetUser = f
}

// UnMockContactBatchGetUser un-mock ContactBatchGetUser method
func (r *Mock) UnMockContactBatchGetUser() {
	r.mockContactBatchGetUser = nil
}

// BatchGetUserReq ...
type BatchGetUserReq struct {
	EmployeeIDs []string `query:"employee_ids" json:"-"` // 支持通过 open_id 或者 employee_id 查询用户信息, 不支持混合两种 ID 进行查询, 单次请求支持的最大用户数量为100
	OpenIDs     []string `query:"open_ids" json:"-"`     // 支持通过 open_id 或者 employee_id 查询用户信息, 不支持混合两种 ID 进行查询, 单次请求支持的最大用户数量为100
}

// BatchGetUserResp ...
type BatchGetUserResp struct {
	UserInfos []*BatchGetUserRespUserInfo `json:"user_infos,omitempty"` // 用户信息
}

// BatchGetUserRespUserInfo ...
type BatchGetUserRespUserInfo struct {
	Name             string                 `json:"name,omitempty"`               // 用户名
	NamePy           string                 `json:"name_py,omitempty"`            // 用户名拼音
	EnName           string                 `json:"en_name,omitempty"`            // 英文名
	EmployeeID       string                 `json:"employee_id,omitempty"`        // 用户的 employee_id, 申请了"获取用户 user_id"权限的应用返回该字段
	EmployeeNo       string                 `json:"employee_no,omitempty"`        // 工号
	OpenID           string                 `json:"open_id,omitempty"`            // 用户的 open_id
	UnionID          string                 `json:"union_id,omitempty"`           // 用户的 union_id
	Status           int64                  `json:"status,omitempty"`             // 用户状态, bit0(最低位): 1冻结, 0未冻结；bit1:1离职, 0在职；bit2:1未激活, 0已激活
	EmployeeType     int64                  `json:"employee_type,omitempty"`      // 员工类型。1:正式员工；2:实习生；3:外包；4:劳务；5:顾问
	Avatar72         string                 `json:"avatar_72,omitempty"`          // 用户头像, 72*72px
	Avatar240        string                 `json:"avatar_240,omitempty"`         // 用户头像, 240*240px
	Avatar640        string                 `json:"avatar_640,omitempty"`         // 用户头像, 640*640px
	AvatarURL        string                 `json:"avatar_url,omitempty"`         // 用户头像, 原始大小
	Gender           int64                  `json:"gender,omitempty"`             // 性别, 未设置不返回该字段。1:男；2:女
	Email            string                 `json:"email,omitempty"`              // 用户邮箱地址, 已申请"获取用户邮箱"权限返回该字段
	Mobile           string                 `json:"mobile,omitempty"`             // 用户手机号, 已申请"获取用户手机号"权限的企业自建应用返回该字段
	Description      string                 `json:"description,omitempty"`        // 用户个人签名
	Country          string                 `json:"country,omitempty"`            // 用户所在国家
	City             string                 `json:"city,omitempty"`               // 用户所在城市
	WorkStation      string                 `json:"work_station,omitempty"`       // 工位
	IsTenantManager  bool                   `json:"is_tenant_manager,omitempty"`  // 是否是企业超级管理员
	JoinTime         int64                  `json:"join_time,omitempty"`          // 入职时间, 未设置不返回该字段
	UpdateTime       int64                  `json:"update_time,omitempty"`        // 更新时间
	LeaderEmployeeID string                 `json:"leader_employee_id,omitempty"` // 用户直接领导的 employee_id, 企业自建应用返回, 应用商店应用申请了 employee_id 权限时才返回
	LeaderOpenID     string                 `json:"leader_open_id,omitempty"`     // 用户直接领导的 open_id
	LeaderUnionID    string                 `json:"leader_union_id,omitempty"`    // 用户直接领导的 union_id
	Departments      []string               `json:"departments,omitempty"`        // 用户所在部门自定义 ID列表, 用户可能同时存在于多个部门
	OpenDepartments  []string               `json:"open_departments,omitempty"`   // 用户所在部门 openID 列表, 用户可能同时存在于多个部门
	CustomAttrs      map[string]interface{} `json:"custom_attrs,omitempty"`       // 用户的自定义属性信息。 该字段返回的每一个属性包括自定义属性 ID 和自定义属性值。  企业开放了自定义用户属性且为该用户设置了自定义属性的值, 才会返回该字段
}

// batchGetUserResp ...
type batchGetUserResp struct {
	Code int64             `json:"code,omitempty"` // 返回码, 非 0 表示失败
	Msg  string            `json:"msg,omitempty"`  // 返回码的描述
	Data *BatchGetUserResp `json:"data,omitempty"` // 返回业务数据
}