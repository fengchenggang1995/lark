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
package larkext

import (
	"context"

	"github.com/chyroc/lark"
)

// BitableTable is a table that can be used to store bitable data.
type BitableTable struct {
	larkClient *lark.Lark
	appToken   string
	tableID    string
}

// NewBitableTable new a bitable table.
func NewBitableTable(larkClient *lark.Lark, appToken string, tableID string) *BitableTable {
	return &BitableTable{larkClient: larkClient, appToken: appToken, tableID: tableID}
}

// Delete delete bitable table
func (r *BitableTable) Delete(ctx context.Context) error {
	return r.delete(ctx)
}
