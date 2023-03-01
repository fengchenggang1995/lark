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

// BatchDeletePublicMailboxMember 一次请求可以删除一个公共邮箱中的多个成员。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/batch_delete
func (r *MailService) BatchDeletePublicMailboxMember(ctx context.Context, request *BatchDeletePublicMailboxMemberReq, options ...MethodOptionFunc) (*BatchDeletePublicMailboxMemberResp, *Response, error) {
	if r.cli.mock.mockMailBatchDeletePublicMailboxMember != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#BatchDeletePublicMailboxMember mock enable")
		return r.cli.mock.mockMailBatchDeletePublicMailboxMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "BatchDeletePublicMailboxMember",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/batch_delete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchDeletePublicMailboxMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailBatchDeletePublicMailboxMember mock MailBatchDeletePublicMailboxMember method
func (r *Mock) MockMailBatchDeletePublicMailboxMember(f func(ctx context.Context, request *BatchDeletePublicMailboxMemberReq, options ...MethodOptionFunc) (*BatchDeletePublicMailboxMemberResp, *Response, error)) {
	r.mockMailBatchDeletePublicMailboxMember = f
}

// UnMockMailBatchDeletePublicMailboxMember un-mock MailBatchDeletePublicMailboxMember method
func (r *Mock) UnMockMailBatchDeletePublicMailboxMember() {
	r.mockMailBatchDeletePublicMailboxMember = nil
}

// BatchDeletePublicMailboxMemberReq ...
type BatchDeletePublicMailboxMemberReq struct {
	PublicMailboxID string   `path:"public_mailbox_id" json:"-"` // The unique ID or email address of a public mailbox, 示例值: "xxxxxxxxxxxxxxx or test_public_mailbox@xxx.xx"
	MemberIDList    []string `json:"member_id_list,omitempty"`   // 本次调用删除的公共邮箱成员ID列表, 示例值: ["xxxxxxx", "yyyyyyy"], 长度范围: `1` ～ `200`
}

// BatchDeletePublicMailboxMemberResp ...
type BatchDeletePublicMailboxMemberResp struct {
}

// batchDeletePublicMailboxMemberResp ...
type batchDeletePublicMailboxMemberResp struct {
	Code int64                               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                              `json:"msg,omitempty"`  // 错误描述
	Data *BatchDeletePublicMailboxMemberResp `json:"data,omitempty"`
}