// Lute - 一款结构化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package test

import (
	"testing"

	"github.com/88250/lute/ast"

	"github.com/88250/lute"
)

var md2BlockDOMTests = []parseTest{

	{"34", "```\nfoo\n```", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeCodeBlock\" class=\"code-block\" updated=\"20060102150405\"><div class=\"protyle-action\"><span class=\"protyle-action--first protyle-action__language\" contenteditable=\"false\"></span><span class=\"fn__flex-1\"></span><span class=\"protyle-icon protyle-icon--first protyle-action__copy\"><svg><use xlink:href=\"#iconCopy\"></use></svg></span><span class=\"protyle-icon protyle-icon--last protyle-action__menu\"><svg><use xlink:href=\"#iconMore\"></use></svg></span></div><div contenteditable=\"true\" spellcheck=\"false\">foo\n</div><div class=\"protyle-attr\" contenteditable=\"false\">​</div></div>"},
	{"33", "{{{row\n{: id=\"20220426085736-e2v2fzx\"}\n\n{: id=\"20220426085738-q96jknf\"}\n\n{: id=\"20220426085739-6f0aec3\"}\n\n}}}\n{: id=\"20220426085744-ipdufm6\"}\n", "<div data-node-id=\"20220426085744-ipdufm6\" data-node-index=\"1\" data-type=\"NodeSuperBlock\" class=\"sb\" updated=\"20220426085744\" data-sb-layout=\"row\"><div data-node-id=\"20220426085736-e2v2fzx\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20220426085736\"><div contenteditable=\"true\" spellcheck=\"false\"></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20220426085738-q96jknf\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20220426085738\"><div contenteditable=\"true\" spellcheck=\"false\"></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20220426085739-6f0aec3\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20220426085739\"><div contenteditable=\"true\" spellcheck=\"false\"></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"32", "* [ ] # foo", "<div data-subtype=\"t\" data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeList\" class=\"list\" updated=\"20060102150405\"><div data-marker=\"*\" data-subtype=\"t\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li\" updated=\"20060102150405\"> <div class=\"protyle-action protyle-action--task\"><svg><use xlink:href=\"#iconUncheck\"></use></svg></div><div data-subtype=\"h1\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeHeading\" class=\"h1\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">foo</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"31", "#<a style=\"b\">foo</a>#", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\"><span data-type=\"tag\" data-content=\"&lt;a style=&quot;b&quot;&gt;foo&lt;/a&gt;\">&lt;a style=&quot;b&quot;&gt;foo&lt;/a&gt;</span></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"30", "<div>foo</div>", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeHTMLBlock\" class=\"render-node\" updated=\"20060102150405\" data-subtype=\"block\"><div class=\"protyle-icons\"><span class=\"protyle-icon protyle-icon--first protyle-action__edit\"><svg><use xlink:href=\"#iconEdit\"></use></svg></span><span class=\"protyle-icon protyle-action__menu protyle-icon--last\"><svg><use xlink:href=\"#iconMore\"></use></svg></span></div><div><protyle-html data-content=\"&lt;div&gt;foo&lt;/div&gt;\"></protyle-html><span style=\"position: absolute\">\u200b</span></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"29", "{{{\n###### 123\n{: custom-type=\"card\"}\n2022-02-18 23:22:27\n{: custom-type=\"datetime\"}\ncontent\n}}}", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeSuperBlock\" class=\"sb\" updated=\"20060102150405\" data-sb-layout=\"row\"><div data-subtype=\"h6\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeHeading\" class=\"h6\" custom-type=\"card\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">123</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" custom-type=\"datetime\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">2022-02-18 23:22:27</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">content</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"28", "|col1‸<br />{: colspan=\"1\" rowspan=\"3\"}|col2|col3|\n|{: class=\"fn__none\"}|||\n|{: class=\"fn__none\"}|||\n| ----------| ------| ------|\n{: id=\"20220216000010-g61t0xu\" updated=\"20220216000039\" colgroup=\"||\"}", "<div data-node-id=\"20220216000010-g61t0xu\" data-node-index=\"1\" data-type=\"NodeTable\" class=\"table\" updated=\"20220216000039\" colgroup=\"||\"><div contenteditable=\"false\"><table contenteditable=\"true\" spellcheck=\"false\"><colgroup><col /><col /><col /></colgroup><thead><tr><th colspan=\"1\" rowspan=\"3\">col1‸<br /></th><th>col2</th><th>col3</th></tr><tr><th class=\"fn__none\"></th><th></th><th></th></tr><tr><th class=\"fn__none\"></th><th></th><th></th></tr></thead><tbody></tbody></table><div class=\"protyle-action__table\"><div class=\"table__resize\"></div><div class=\"table__select\"></div></div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"27", "((20210510191408-b2n8h2c 'foo'))", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\"><span data-type=\"block-ref\" data-subtype=\"d\" data-id=\"20210510191408-b2n8h2c\">foo</span></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"26", "[foo](assets/bar#bar?test#bar)", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\"><span data-type=\"a\" data-href=\"assets/bar%23bar?test#bar\">foo</span></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"25", "[foo](assets/bar#bar&baz bazz)", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\"><span data-type=\"a\" data-href=\"assets/bar%23bar&amp;baz bazz\">foo</span></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"24", "[foo](bar#bar&baz bazz)", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\"><span data-type=\"a\" data-href=\"bar#bar&amp;baz bazz\">foo</span></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"23", "![a](\"<img src=xss onerror=alert(1)>)", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\"><span contenteditable=\"false\" data-type=\"img\" class=\"img\"><span> </span><span><span class=\"protyle-action protyle-icons\"><span class=\"protyle-icon protyle-icon--only\"><svg class=\"svg\"><use xlink:href=\"#iconMore\"></use></svg></span></span><img src=\"&#34;<img src=\"xss\">\" data-src=\"&#34;<img src=\"xss\">&#34; alt=&#34;a&#34; /&gt;<span class=\"protyle-action__drag\"></span><span class=\"img__net\"><svg><use xlink:href=\"#iconLanguage\"></use></svg></span><span class=\"protyle-action__title\"></span></span><span> </span></span></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"22", "* [ ] `foo`", "<div data-subtype=\"t\" data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeList\" class=\"list\" updated=\"20060102150405\"><div data-marker=\"*\" data-subtype=\"t\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li\" updated=\"20060102150405\"><div class=\"protyle-action protyle-action--task\"><svg><use xlink:href=\"#iconUncheck\"></use></svg></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\"><code>foo</code></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"21", "* [ ] **foo**", "<div data-subtype=\"t\" data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeList\" class=\"list\" updated=\"20060102150405\"><div data-marker=\"*\" data-subtype=\"t\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li\" updated=\"20060102150405\"><div class=\"protyle-action protyle-action--task\"><svg><use xlink:href=\"#iconUncheck\"></use></svg></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\"><strong>foo</strong></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"20", "* [ ]", "<div data-subtype=\"t\" data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeList\" class=\"list\" updated=\"20060102150405\"><div data-marker=\"*\" data-subtype=\"t\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li\" updated=\"20060102150405\"><div class=\"protyle-action protyle-action--task\"><svg><use xlink:href=\"#iconUncheck\"></use></svg></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\"></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"19", "[foo](<https://b3log.org/bar>", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">[foo](<span data-type=\"a\" data-href=\"https://b3log.org/bar\">https://b3log.org/bar</span></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"18", "<<assets/文件名-20210911230735-pzlpdtf.pdf/20210911230820-lhiaysx \"注解锚文本\">>", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\"><span data-type=\"file-annotation-ref\" data-subtype=\"s\" data-id=\"assets/文件名-20210911230735-pzlpdtf.pdf/20210911230820-lhiaysx\">注解锚文本</span></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"17", "```\n{{{col\n{{{row\nfoo\n{: style=\"color: var(--b3-card-error-color);background-color: var(--b3-card-error-background);\"}\n}}}\n{{{row\nbar\n}}}\n}}}\n```", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeCodeBlock\" class=\"code-block\" updated=\"20060102150405\"><div class=\"protyle-action\"><span class=\"protyle-action--first protyle-action__language\" contenteditable=\"false\"></span><span class=\"fn__flex-1\"></span><span class=\"protyle-icon protyle-icon--first protyle-action__copy\"><svg><use xlink:href=\"#iconCopy\"></use></svg></span><span class=\"protyle-icon protyle-icon--last protyle-action__menu\"><svg><use xlink:href=\"#iconMore\"></use></svg></span></div><div contenteditable=\"true\" spellcheck=\"false\">{{{col\n{{{row\nfoo\n{: style=&quot;color: var(--b3-card-error-color);background-color: var(--b3-card-error-background);&quot;}\n}}}\n{{{row\nbar\n}}}\n}}}\n</div><div class=\"protyle-attr\" contenteditable=\"false\">​</div></div>"},
	{"16", "<iframe src=\"https://b3log.org\" updated=\"20060102150405\" data-subtype=\"widget\"></iframe>", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeWidget\" class=\"iframe\" updated=\"20060102150405\" data-subtype=\"widget\"><div class=\"iframe-content\"><iframe src=\"https://b3log.org\" data-src=\"https://b3log.org\" updated=\"20060102150405\" data-subtype=\"widget\"></iframe><span class=\"protyle-action__drag\" contenteditable=\"false\"></span></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"15", "<iframe src=\"https://b3log.org\"></iframe>", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeIFrame\" class=\"iframe\" updated=\"20060102150405\"><div class=\"iframe-content\"><iframe src=\"https://b3log.org\" data-src=\"https://b3log.org\"></iframe><span class=\"protyle-action__drag\" contenteditable=\"false\"></span></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"14", "foo<span data-type=\"virtual-block-ref\">bar</span>baz\n", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">foo<span data-type=\"virtual-block-ref\">bar</span>baz</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"13", "1. foo\n\n    1. bar\n", "<div data-subtype=\"o\" data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeList\" class=\"list\" updated=\"20060102150405\"><div data-marker=\"1.\" data-subtype=\"o\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li\" updated=\"20060102150405\"><div class=\"protyle-action protyle-action--order\" contenteditable=\"false\" draggable=\"true\">1.</div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">foo</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-subtype=\"o\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeList\" class=\"list\" updated=\"20060102150405\"><div data-marker=\"1.\" data-subtype=\"o\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li\" updated=\"20060102150405\"><div class=\"protyle-action protyle-action--order\" contenteditable=\"false\" draggable=\"true\">1.</div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">bar</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"12", "![foo](assets/bar)", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\"><span contenteditable=\"false\" data-type=\"img\" class=\"img\"><span> </span><span><span class=\"protyle-action protyle-icons\"><span class=\"protyle-icon protyle-icon--only\"><svg class=\"svg\"><use xlink:href=\"#iconMore\"></use></svg></span></span><img src=\"assets/bar\" data-src=\"assets/bar\" alt=\"foo\" /><span class=\"protyle-action__drag\"></span><span class=\"protyle-action__title\"></span></span><span> </span></span></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"11", "| col1 | col2 | col3 |\n| ------ | ------ | ------ |\n|      |      |      |\n|      |      |      |\n\n<video controls=\"controls\" src=\"assets/RPReplay_Final16206653001-20210513170900-i8k37se.mp4\" data-src=\"assets/RPReplay_Final16206653001-20210513170900-i8k37se.mp4\"></video>\n\n11 ((20210513163831-8k81fw8))\n\n{{select * from blocks where id='20210513163831-8k81fw8'}}\n```js\nvar a = 1\n```\n\n![CleanShot_2021-05-07_at_18.30.11.gif](assets/CleanShot_2021-05-07_at_18.30.11-20210513163234-77jawpg.gif \"title\")\n\n<audio controls=\"controls\" src=\"assets/record1620894762009-20210513163242-toqc85e.wav\" data-src=\"assets/record1620894762009-20210513163242-toqc85e.wav\"></audio>\n", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeTable\" class=\"table\" updated=\"20060102150405\"><div contenteditable=\"false\"><table contenteditable=\"true\" spellcheck=\"false\"><colgroup><col /><col /><col /></colgroup><thead><tr><th>col1</th><th>col2</th><th>col3</th></tr></thead><tbody><tr><td></td><td></td><td></td></tr><tr><td></td><td></td><td></td></tr></tbody></table><div class=\"protyle-action__table\"><div class=\"table__resize\"></div><div class=\"table__select\"></div></div></div><div class=\"protyle-attr\" contenteditable=\"false\">​</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"2\" data-type=\"NodeVideo\" class=\"iframe\" updated=\"20060102150405\"><div class=\"iframe-content\">​<video controls=\"controls\" src=\"assets/RPReplay_Final16206653001-20210513170900-i8k37se.mp4\" data-src=\"assets/RPReplay_Final16206653001-20210513170900-i8k37se.mp4\"></video><span class=\"protyle-action__drag\" contenteditable=\"false\"></span></div><div class=\"protyle-attr\" contenteditable=\"false\">​</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"3\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">11 <span data-type=\"block-ref\" data-subtype=\"d\" data-id=\"20210513163831-8k81fw8\"></span></div><div class=\"protyle-attr\" contenteditable=\"false\">​</div></div><div data-content=\"select * from blocks where id='20210513163831-8k81fw8'\" data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"4\" data-type=\"NodeBlockQueryEmbed\" class=\"render-node\" updated=\"20060102150405\"><div class=\"protyle-attr\" contenteditable=\"false\">​</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"5\" data-type=\"NodeCodeBlock\" class=\"code-block\" updated=\"20060102150405\"><div class=\"protyle-action\"><span class=\"protyle-action--first protyle-action__language\" contenteditable=\"false\">js</span><span class=\"fn__flex-1\"></span><span class=\"protyle-icon protyle-icon--first protyle-action__copy\"><svg><use xlink:href=\"#iconCopy\"></use></svg></span><span class=\"protyle-icon protyle-icon--last protyle-action__menu\"><svg><use xlink:href=\"#iconMore\"></use></svg></span></div><div contenteditable=\"true\" spellcheck=\"false\">var a = 1\n</div><div class=\"protyle-attr\" contenteditable=\"false\">​</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"6\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\"><span contenteditable=\"false\" data-type=\"img\" class=\"img\"><span> </span><span><span class=\"protyle-action protyle-icons\"><span class=\"protyle-icon protyle-icon--only\"><svg class=\"svg\"><use xlink:href=\"#iconMore\"></use></svg></span></span><img src=\"assets/CleanShot_2021-05-07_at_18.30.11-20210513163234-77jawpg.gif\" data-src=\"assets/CleanShot_2021-05-07_at_18.30.11-20210513163234-77jawpg.gif\" alt=\"CleanShot_2021-05-07_at_18.30.11.gif\" title=\"title\" /><span class=\"protyle-action__drag\"></span><span class=\"protyle-action__title\">title</span></span><span> </span></span></div><div class=\"protyle-attr\" contenteditable=\"false\">​</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"7\" data-type=\"NodeAudio\" class=\"iframe\" updated=\"20060102150405\"><div class=\"iframe-content\"><audio controls=\"controls\" src=\"assets/record1620894762009-20210513163242-toqc85e.wav\" data-src=\"assets/record1620894762009-20210513163242-toqc85e.wav\"></audio>​</div><div class=\"protyle-attr\" contenteditable=\"false\">​</div></div>"},
	{"10", "((20210510191408-b2n8h2c \"foo\"))", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\"><span data-type=\"block-ref\" data-subtype=\"s\" data-id=\"20210510191408-b2n8h2c\">foo</span></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"9", "{{{\n\nfoo\n\n{{{\n\n* bar\n\n  {{{\n  \n  bar2\n  \n  bar3\n  \n  }}}\n\nbar4\n\n}}}\n\nbaz\n\n}}}\n\nbazz\n\n", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeSuperBlock\" class=\"sb\" updated=\"20060102150405\" data-sb-layout=\"row\"><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">foo</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeSuperBlock\" class=\"sb\" updated=\"20060102150405\" data-sb-layout=\"row\"><div data-subtype=\"u\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeList\" class=\"list\" updated=\"20060102150405\"><div data-marker=\"*\" data-subtype=\"u\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li\" updated=\"20060102150405\"><div class=\"protyle-action\" draggable=\"true\"><svg><use xlink:href=\"#iconDot\"></use></svg></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">bar</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeSuperBlock\" class=\"sb\" updated=\"20060102150405\" data-sb-layout=\"row\"><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">bar2</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">bar3</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">bar4</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">baz</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"2\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">bazz</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"8", "{{{\n\nfoo\n\n{{{\n\n* bar \n\nba2\n\n}}}\n\nbaz\n\n}}}\n\nbazz\n\n", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeSuperBlock\" class=\"sb\" updated=\"20060102150405\" data-sb-layout=\"row\"><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">foo</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeSuperBlock\" class=\"sb\" updated=\"20060102150405\" data-sb-layout=\"row\"><div data-subtype=\"u\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeList\" class=\"list\" updated=\"20060102150405\"><div data-marker=\"*\" data-subtype=\"u\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li\" updated=\"20060102150405\"><div class=\"protyle-action\" draggable=\"true\"><svg><use xlink:href=\"#iconDot\"></use></svg></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">bar</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">ba2</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">baz</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"2\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">bazz</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"7", "{{{\n\n* baz\n\n}}}\n\nbazz\n\n", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeSuperBlock\" class=\"sb\" updated=\"20060102150405\" data-sb-layout=\"row\"><div data-subtype=\"u\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeList\" class=\"list\" updated=\"20060102150405\"><div data-marker=\"*\" data-subtype=\"u\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li\" updated=\"20060102150405\"><div class=\"protyle-action\" draggable=\"true\"><svg><use xlink:href=\"#iconDot\"></use></svg></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">baz</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"2\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">bazz</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"6", "{{{col\n\n{{{\nfoo\n\nbar\n}}}\n\nbaz\n\n}}}\n\nbazz\n\n", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeSuperBlock\" class=\"sb\" updated=\"20060102150405\" data-sb-layout=\"col\"><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeSuperBlock\" class=\"sb\" updated=\"20060102150405\" data-sb-layout=\"row\"><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">foo</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">bar</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">baz</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"2\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">bazz</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"5", "{{{col\n\n{{{\nfoo\n}}}\n\n}}}\n\nbar\n\n", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeSuperBlock\" class=\"sb\" updated=\"20060102150405\" data-sb-layout=\"col\"><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeSuperBlock\" class=\"sb\" updated=\"20060102150405\" data-sb-layout=\"row\"><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">foo</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"2\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">bar</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"4", "<u>foo</u>", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\"><u>foo</u></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"3", "* [x] foo", "<div data-subtype=\"t\" data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeList\" class=\"list\" updated=\"20060102150405\"><div data-marker=\"*\" data-subtype=\"t\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li protyle-task--done\" updated=\"20060102150405\"><div class=\"protyle-action protyle-action--task\"><svg><use xlink:href=\"#iconCheck\"></use></svg></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\">foo</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"2", "{{select * from blocks}}", "<div data-content=\"select * from blocks\" data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeBlockQueryEmbed\" class=\"render-node\" updated=\"20060102150405\"><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"1", "<kbd>foo</kbd>", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20060102150405\"><div contenteditable=\"true\" spellcheck=\"false\"><kbd>foo</kbd></div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
	{"0", "<audio src=\"assets/foo\"></audio>", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeAudio\" class=\"iframe\" updated=\"20060102150405\"><div class=\"iframe-content\"><audio src=\"assets/foo\" data-src=\"assets/foo\"></audio>\u200b</div><div class=\"protyle-attr\" contenteditable=\"false\">\u200b</div></div>"},
}

func TestMd2BlockDOM(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.SetProtyleWYSIWYG(true)
	luteEngine.SetToC(true)
	luteEngine.ParseOptions.BlockRef = true
	luteEngine.ParseOptions.KramdownBlockIAL = true
	luteEngine.RenderOptions.KramdownBlockIAL = true
	luteEngine.ParseOptions.Tag = true
	luteEngine.ParseOptions.SuperBlock = true
	luteEngine.ParseOptions.GitConflict = true
	luteEngine.ParseOptions.LinkRef = false
	luteEngine.SetAutoSpace(true)
	luteEngine.SetParagraphBeginningSpace(true)
	luteEngine.SetFileAnnotationRef(true)
	luteEngine.SetImgPathAllowSpace(true)
	luteEngine.SetSanitize(true)

	ast.Testing = true
	for _, test := range md2BlockDOMTests {
		result := luteEngine.Md2BlockDOM(test.from)
		if test.to != result {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal html\n\t%q", test.name, test.to, result, test.from)
		}
	}
	ast.Testing = false
}
