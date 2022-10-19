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

// GetBitableTableList 根据  app_token, 获取多维表格下的所有数据表
//
// 该接口支持调用频率上限为 20 QPS
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/list
func (r *BitableService) GetBitableTableList(ctx context.Context, request *GetBitableTableListReq, options ...MethodOptionFunc) (*GetBitableTableListResp, *Response, error) {
	if r.cli.mock.mockBitableGetBitableTableList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#GetBitableTableList mock enable")
		return r.cli.mock.mockBitableGetBitableTableList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "GetBitableTableList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getBitableTableListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableGetBitableTableList mock BitableGetBitableTableList method
func (r *Mock) MockBitableGetBitableTableList(f func(ctx context.Context, request *GetBitableTableListReq, options ...MethodOptionFunc) (*GetBitableTableListResp, *Response, error)) {
	r.mockBitableGetBitableTableList = f
}

// UnMockBitableGetBitableTableList un-mock BitableGetBitableTableList method
func (r *Mock) UnMockBitableGetBitableTableList() {
	r.mockBitableGetBitableTableList = nil
}

// GetBitableTableListReq ...
type GetBitableTableListReq struct {
	AppToken  string  `path:"app_token" json:"-"`   // bitable app token, 示例值: "appbcbWCzen6D8dezhoCH2RpMAh"
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "tblsRc9GRRXKqhvW"
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 10, 默认值: `20`, 最大值: `100`
}

// GetBitableTableListResp ...
type GetBitableTableListResp struct {
	HasMore   bool                           `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                         `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Total     int64                          `json:"total,omitempty"`      // 总数
	Items     []*GetBitableTableListRespItem `json:"items,omitempty"`      // 数据表信息
}

// GetBitableTableListRespItem ...
type GetBitableTableListRespItem struct {
	TableID  string `json:"table_id,omitempty"` // 数据表 id
	Revision int64  `json:"revision,omitempty"` // 数据表的版本号
	Name     string `json:"name,omitempty"`     // 数据表名字
}

// getBitableTableListResp ...
type getBitableTableListResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *GetBitableTableListResp `json:"data,omitempty"`
}
