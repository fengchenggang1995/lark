// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetSheetValue
//
// 该接口用于根据 spreadsheetToken 和 range 读取表格单个范围的值，返回数据限制为10M。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ugTMzUjL4EzM14COxMTN
func (r *DriveService) GetSheetValue(ctx context.Context, request *GetSheetValueReq, options ...MethodOptionFunc) (*GetSheetValueResp, *Response, error) {
	if r.cli.mock.mockDriveGetSheetValue != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetSheetValue mock enable")
		return r.cli.mock.mockDriveGetSheetValue(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetSheetValue",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values/:range",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getSheetValueResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveGetSheetValue(f func(ctx context.Context, request *GetSheetValueReq, options ...MethodOptionFunc) (*GetSheetValueResp, *Response, error)) {
	r.mockDriveGetSheetValue = f
}

func (r *Mock) UnMockDriveGetSheetValue() {
	r.mockDriveGetSheetValue = nil
}

type GetSheetValueReq struct {
	ValueRenderOption    *string `query:"valueRenderOption" json:"-"`    // valueRenderOption=ToString 可返回纯文本的值；valueRenderOption=FormattedValue 计算并格式化单元格；valueRenderOption=Formula单元格中含有公式时返回公式本身；valueRenderOption=UnformattedValue计算但不对单元格进行格式化。
	DateTimeRenderOption *string `query:"dateTimeRenderOption" json:"-"` // dateTimeRenderOption=FormattedString 计算并对时间日期按照其格式进行格式化，但不会对数字进行格式化，返回格式化后的字符串。
	SpreadSheetToken     string  `path:"spreadsheetToken" json:"-"`      // spreadsheet 的 token，详见[在线表格开发指南](/ssl:ttdoc/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	Range                string  `path:"range" json:"-"`                 // 查询范围，包含 sheetId 与单元格范围两部分，目前支持四种索引方式，详见[在线表格开发指南](/ssl:ttdoc/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
}

type getSheetValueResp struct {
	Code int64              `json:"code,omitempty"`
	Msg  string             `json:"msg,omitempty"`
	Data *GetSheetValueResp `json:"data,omitempty"`
}

type GetSheetValueResp struct {
	Revision         int64                        `json:"revision,omitempty"`         // sheet 的版本号
	SpreadSheetToken string                       `json:"spreadsheetToken,omitempty"` // spreadsheet 的 token，详见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN) 的第 4 项
	ValueRange       *GetSheetValueRespValueRange `json:"valueRange,omitempty"`       // 值与范围
}

type GetSheetValueRespValueRange struct {
	MajorDimension string        `json:"majorDimension,omitempty"` // 插入维度
	Range          string        `json:"range,omitempty"`          // 返回数据的范围，为空时表示查询范围没有数据
	Revision       int64         `json:"revision,omitempty"`       // sheet 的版本号
	Values         []interface{} `json:"values,omitempty"`         // 查询得到的值
}
