// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// PreviewApprovalInstance
//
// 提交审批前，预览审批流程。或者发起审批后，在某一审批节点预览后续流程
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ukTM5UjL5ETO14SOxkTN/approval-preview
func (r *ApprovalService) PreviewApprovalInstance(ctx context.Context, request *PreviewApprovalInstanceReq, options ...MethodOptionFunc) (*PreviewApprovalInstanceResp, *Response, error) {
	if r.cli.mock.mockApprovalPreviewApprovalInstance != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#PreviewApprovalInstance mock enable")
		return r.cli.mock.mockApprovalPreviewApprovalInstance(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "PreviewApprovalInstance",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/approval/v4/instances/preview",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(previewApprovalInstanceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockApprovalPreviewApprovalInstance(f func(ctx context.Context, request *PreviewApprovalInstanceReq, options ...MethodOptionFunc) (*PreviewApprovalInstanceResp, *Response, error)) {
	r.mockApprovalPreviewApprovalInstance = f
}

func (r *Mock) UnMockApprovalPreviewApprovalInstance() {
	r.mockApprovalPreviewApprovalInstance = nil
}

type PreviewApprovalInstanceReq struct {
	UserIDType   *IDType                         `query:"user_id_type" json:"-"`  // 默认为open_id, 对于open_id(ou_开头)类型，user_id_type为open_id, 对于employeeID(8位字符串，如f7cb567e)类型，user_id_type为user_id
	ApprovalCode *string                         `json:"approval_code,omitempty"` // 审批定义 Code
	UserID       string                          `json:"user_id,omitempty"`       // 发起审批用户，employeid或者openid
	DepartmentID *string                         `json:"department_id,omitempty"` // 发起审批用户部门，如果用户只属于一个部门，可以不填，如果属于多个部门，必须填其中一个部门
	Form         *PreviewApprovalInstanceReqForm `json:"form,omitempty"`          // json 数组，控件值。提交审批之前，查看预览流程时，该字段必填
	InstanceCode *string                         `json:"instance_code,omitempty"` // 审批实例code
	TaskID       *string                         `json:"task_id,omitempty"`       // 若审批实例已存在，则传递当前审批任务对应的task_id
}

type PreviewApprovalInstanceReqForm struct {
	ID    string `json:"id,omitempty"`    // 控件ID，也可以使用自定义 ID custom_id 的值
	Type  string `json:"type,omitempty"`  // 控件类型
	Value string `json:"value,omitempty"` // 控件值，不同类型的值格式不一样
}

type previewApprovalInstanceResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码，非0表示失败
	Msg  string                       `json:"msg,omitempty"`  // 返回码的描述
	Data *PreviewApprovalInstanceResp `json:"data,omitempty"` // 返回业务信息
}

type PreviewApprovalInstanceResp struct {
	PreviewNodes []string `json:"preview_nodes,omitempty"`  // 预览节点信息
	UserIDList   []string `json:"user_id_list,omitempty"`   // 审批人id列表
	EndCcIDList  []string `json:"end_cc_id_list,omitempty"` // 审批结束抄送人ie列表
	NodeID       string   `json:"node_id,omitempty"`        // 节点id
	NodeName     string   `json:"node_name,omitempty"`      // 节点名称
	NodeType     string   `json:"node_type,omitempty"`      // 节点类型：AND：会签，OR: 或签
	CustomNodeID string   `json:"custom_node_id,omitempty"` // 用户自定义节点id
	Comments     []string `json:"comments,omitempty"`       // 节点的说明信息
}