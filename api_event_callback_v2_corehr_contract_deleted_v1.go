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

// EventV2CorehrContractDeletedV1 合同删除{使用示例}(url=/api/tools/api_explore/api_explore_config?project=corehr&version=v1&resource=contract&event=deleted)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/contract/events/deleted
func (r *EventCallbackService) HandlerEventV2CorehrContractDeletedV1(f EventV2CorehrContractDeletedV1Handler) {
	r.cli.eventHandler.eventV2CorehrContractDeletedV1Handler = f
}

// EventV2CorehrContractDeletedV1Handler event EventV2CorehrContractDeletedV1 handler
type EventV2CorehrContractDeletedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2CorehrContractDeletedV1) (string, error)

// EventV2CorehrContractDeletedV1 ...
type EventV2CorehrContractDeletedV1 struct {
	ContractID string `json:"contract_id,omitempty"` // ID
}