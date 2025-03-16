package slack

type Message struct {
	Text   string  `json:"text"`
	Blocks []Block `json:"blocks,omitempty"`
}

type TextType string

const (
	TextTypePlainText TextType = "plain_text"
	TextTypeMarkdown  TextType = "mrkdwn"
)

type Text struct {
	Type  TextType `json:"type"`
	Text  string   `json:"text"`
	Emoji bool     `json:"emoji,omitempty"`
}

type BlockType string

const (
	BlockTypeHeader  BlockType = "header"
	BlockTypeSection BlockType = "section"
)

type Block struct {
	Type   BlockType `json:"type"`
	Text   *Text     `json:"text,omitempty"`
	Fields []Text    `json:"fields,omitempty"`
}
