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
package card

import (
	"fmt"

	"github.com/chyroc/lark/card"
)

func Example6() {
	content := card.Card(
		card.Div().
			SetFields(
				card.FieldMarkdown("**时间**\\n2021-07-25 15:35:00").SetIsShort(true),
				card.FieldMarkdown("**地点**\\n江苏省、浙江省、上海市").SetIsShort(true),
			),
		card.Div().
			SetText(card.MarkdownText("亲爱的同事们，\\n气象局发布台风橙色预警，7月26日江浙沪地区预计平均风力可达10级以上。\\n建议江浙沪地区同学明日居家办公。如有值班等特殊情况，请各部门视情况安排。\\n请同学们关好门窗，妥善安置室外用品，停止一切户外活动，注意保护自身安全。\\n如有疑问，请联系[值班号](https://open.larksuite.com/)。")),
		card.Action(
			card.Button("我已知悉", map[string]interface{}{"key": "value"}).SetPrimary(),
		),
		card.HR(),
		card.Note(
			card.MarkdownText("[来自应急通知](https://www.open.larksuite.com/)"),
		),
	).
		SetHeader(card.Header("【应急通知】7月26日江浙沪地区居家办公通知").SetRed())
	fmt.Println(content.String())
}
