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

func Example3() {
	content := card.Card(
		card.ModuleImage("img_v2_bfd72a81-1533-4699-995d-12a675708d0g"),
		card.Div().SetText(card.MarkdownText("你是否曾因为一本书而产生心灵共振，开始感悟人生？\\n你有哪些想极力推荐给他人的珍藏好书？\\n\\n加入 **4·23 飞书读书节**，分享你的**挚爱书单**及**读书笔记**，**赢取千元读书礼**！\\n\\n📬 填写问卷，晒出你的珍藏好书\\n😍 想知道其他人都推荐了哪些好书？马上[入群围观](https://open.larksuite.com/)\\n📝 用[读书笔记模板](https://open.larksuite.com/)（桌面端打开），记录你的心得体会\\n🙌 更有惊喜特邀嘉宾 4月12日起带你共读")),
		card.Action(
			card.LinkButton("立即推荐好书", "https://open.larksuite.com/").SetPrimary(),
			card.LinkButton("查看活动指南", "https://open.larksuite.com/").SetDefault(),
		),
	).
		SetHeader(card.Header("📚晒挚爱好书，赢读书礼金").SetTurquoise())
	fmt.Println(content.String())
}
