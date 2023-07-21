package urlpath

import (
	"strings"
	"unicode/utf8"
	"utils/ascii"
)

type multiBuilder interface {
	toLines(prefix string) []string
}

func addSlash(prev string) string {
	if (0 == len(prev)) || strings.HasSuffix(prev, "/") || strings.HasSuffix(prev, "┤") {
		return ""
	}
	return "/"
}

type LinearBuilder struct {
	line          string
	firstRaw      bool
	multiBuilders []multiBuilder
}

func (b *LinearBuilder) toLines(prefix string) (lines []string) {
	if !b.firstRaw {
		prefix += addSlash(prefix)
	}
	prefix += b.line
	if 0 == len(b.multiBuilders) {
		return append(lines, prefix)
	}
	subsequentPrefix := createSubsequentPrefix(prefix)
	for _, builder := range b.multiBuilders {
		lines = append(lines, builder.toLines(prefix)...)
		prefix = subsequentPrefix
	}
	return
}

func (b *LinearBuilder) AddRawText(text string) *LinearBuilder {
	if 0 == len(b.line) {
		b.firstRaw = true
	}
	b.line += text
	return b
}

func (b *LinearBuilder) AddPathEntry(text string) *LinearBuilder {
	if 0 != len(b.line) {
		b.line += "/"
	}
	if 0 == len(text) {
		text = "?empty¿"
	}
	b.line += text
	return b
}

func (b *LinearBuilder) AddOr() *OrBuilder {
	builder := &OrBuilder{}
	b.multiBuilders = append(b.multiBuilders, builder)
	return builder
}

func (b *LinearBuilder) AddMap() *MapBuilder {
	builder := &MapBuilder{options: make(map[string]*LinearBuilder)}
	b.multiBuilders = append(b.multiBuilders, builder)
	return builder
}

type OrBuilder struct {
	options []*LinearBuilder
}

func (b *OrBuilder) NextOption() *LinearBuilder {
	lb := &LinearBuilder{}
	b.options = append(b.options, lb)
	return lb
}

func (b *OrBuilder) toLines(prefix string) (lines []string) {
	prefix += addSlash(prefix)
	subsequentPrefix := createSubsequentPrefix(prefix)
	for _, option := range b.options {
		lines = append(lines, option.toLines(prefix+"┤")...)
		prefix = subsequentPrefix
	}
	return
}

type MapBuilder struct {
	orderedKeys []string
	options     map[string]*LinearBuilder
}

func (b *MapBuilder) AddOption(option string) *LinearBuilder {
	if _, ok := b.options[option]; ok {
		panic("attempt to add duplicate option '" + option + "' to (Graph) MapBuilder")
	}
	b.orderedKeys = append(b.orderedKeys, option)
	builder := &LinearBuilder{}
	b.options[option] = builder
	return builder
}

func (b *MapBuilder) toLines(prefix string) (lines []string) { // upside down question mark indicates Map entry
	prefix += addSlash(prefix)
	subsequentPrefix := createSubsequentPrefix(prefix)
	for _, key := range b.orderedKeys {
		lines = append(lines, b.options[key].toLines(prefix+"¿"+key)...)
		prefix = subsequentPrefix
	}
	return
}

func createSubsequentPrefix(prefix string) string { // simplistic alignment for subsequent rows
	return ascii.CharsN(' ', utf8.RuneCountInString(prefix)) // any rune not standard Ascii width won't align
}

type GraphBuilder struct {
	builder *LinearBuilder
}

func NewGraphBuilder() *GraphBuilder {
	pgb := &GraphBuilder{builder: &LinearBuilder{}}
	return pgb
}

func (b *GraphBuilder) GetLinearBuilder() *LinearBuilder {
	return b.builder
}

func (b *GraphBuilder) String() (result string) {
	lines := b.builder.toLines("/")
	lineCount := len(lines)
	if lineCount != 0 {
		result = lines[0]
		for i := 1; i < lineCount; i++ {
			result += "\n" + lines[i]
		}
	}
	return
}
