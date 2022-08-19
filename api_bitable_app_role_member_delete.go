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

// DeleteBitableAppRoleMember 删除自定义角色的协作者
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-role-member/delete
func (r *BitableService) DeleteBitableAppRoleMember(ctx context.Context, request *DeleteBitableAppRoleMemberReq, options ...MethodOptionFunc) (*DeleteBitableAppRoleMemberResp, *Response, error) {
	if r.cli.mock.mockBitableDeleteBitableAppRoleMember != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#DeleteBitableAppRoleMember mock enable")
		return r.cli.mock.mockBitableDeleteBitableAppRoleMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "DeleteBitableAppRoleMember",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/roles/:role_id/members/:member_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteBitableAppRoleMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableDeleteBitableAppRoleMember mock BitableDeleteBitableAppRoleMember method
func (r *Mock) MockBitableDeleteBitableAppRoleMember(f func(ctx context.Context, request *DeleteBitableAppRoleMemberReq, options ...MethodOptionFunc) (*DeleteBitableAppRoleMemberResp, *Response, error)) {
	r.mockBitableDeleteBitableAppRoleMember = f
}

// UnMockBitableDeleteBitableAppRoleMember un-mock BitableDeleteBitableAppRoleMember method
func (r *Mock) UnMockBitableDeleteBitableAppRoleMember() {
	r.mockBitableDeleteBitableAppRoleMember = nil
}

// DeleteBitableAppRoleMemberReq ...
type DeleteBitableAppRoleMemberReq struct {
	AppToken     string  `path:"app_token" json:"-"`       // bitable app token, 示例值: "appbcbWCzen6D8dezhoCH2RpMAh"
	RoleID       string  `path:"role_id" json:"-"`         // 自定义角色的id, 示例值: "roljRpwIUt"
	MemberID     string  `path:"member_id" json:"-"`       // 协作者id, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad53uew2"
	MemberIDType *IDType `query:"member_id_type" json:"-"` // 协作者id类型, 与请求体中的member_id要对应, 示例值: "open_id", 可选值有: open_id: 以open_id来识别协作者, union_id: 以union_id来识别协作者, user_id: 以user_id来识别协作者, chat_id: 以chat_id来识别协作者, department_id: 以department_id来识别协作者, open_department_id: 以open_department_id来识别协作者, 默认值: `open_id`
}

// DeleteBitableAppRoleMemberResp ...
type DeleteBitableAppRoleMemberResp struct {
}

// deleteBitableAppRoleMemberResp ...
type deleteBitableAppRoleMemberResp struct {
	Code int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *DeleteBitableAppRoleMemberResp `json:"data,omitempty"`
}