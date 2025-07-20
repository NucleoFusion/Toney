package config

import "github.com/charmbracelet/glamour/ansi"

func ToGlamourStyle(styles RendererConfig) ansi.StyleConfig {
	cfg := ansi.StyleConfig{
		Document:       toStyle(styles.Document),
		BlockQuote:     toStyle(styles.BlockQuote),
		Paragraph:      toStyle(styles.Paragraph),
		Heading:        toHeading(styles.Heading, 0),
		H1:             toHeading(styles.Heading, 1),
		H2:             toHeading(styles.Heading, 2),
		H3:             toHeading(styles.Heading, 3),
		H4:             toHeading(styles.Heading, 4),
		H5:             toHeading(styles.Heading, 5),
		H6:             toHeading(styles.Heading, 6),
		HorizontalRule: toInline(styles.HorizontalRule),
		List:           toList(styles.List),
		Enumeration:    toInline(styles.Enumeration),
		Link:           toInline(styles.Link),
		Image:          toInline(styles.Image),
		Emph:           toInline(styles.Emph),
		Strong:         toInline(styles.Strong),
		Strikethrough:  toInline(styles.Strikethrough),
		Code:           toStyle(styles.Code),
		CodeBlock:      toCodeBlockStyle(styles.CodeBlock),
		Table:          toTableStyle(styles.Table),
	}

	cfg.Item = ansi.StylePrimitive{
		BlockPrefix: styles.ItemPrefix,
	}
	cfg.Enumeration = ansi.StylePrimitive{
		BlockPrefix: styles.EnumerationPrefix,
	}

	if styles.List.Task != nil {
		cfg.Task = ansi.StyleTask{
			StylePrimitive: ansi.StylePrimitive{},
			Ticked:         styles.List.Task.Ticked,
			Unticked:       styles.List.Task.Unticked,
		}
	}

	return cfg
}

func toTableStyle(cfg TableStyle) ansi.StyleTable {
	return ansi.StyleTable{
		StyleBlock:      toStyle(cfg.BlockStyle),
		CenterSeparator: stringPtr(cfg.CenterSeparator),
		ColumnSeparator: stringPtr(cfg.ColumnSeparator),
		RowSeparator:    stringPtr(cfg.RowSeparator),
	}
}

func toCodeBlockStyle(cfg CodeBlockStyle) ansi.StyleCodeBlock {
	var chroma ansi.Chroma

	chroma.Text = toInline(cfg.Text)
	chroma.Error = toInline(cfg.Error)
	chroma.Comment = toInline(cfg.Comment)
	chroma.CommentPreproc = toInline(cfg.CommentPreproc)
	chroma.Keyword = toInline(cfg.Keyword)
	chroma.KeywordReserved = toInline(cfg.KeywordReserved)
	chroma.KeywordNamespace = toInline(cfg.KeywordNamespace)
	chroma.KeywordType = toInline(cfg.KeywordType)
	chroma.Operator = toInline(cfg.Operator)
	chroma.Punctuation = toInline(cfg.Punctuation)
	chroma.Name = toInline(cfg.Name)
	chroma.NameBuiltin = toInline(cfg.NameBuiltin)
	chroma.NameTag = toInline(cfg.NameTag)
	chroma.NameAttribute = toInline(cfg.NameAttribute)
	chroma.NameClass = toInline(cfg.NameClass)
	chroma.NameConstant = toInline(cfg.NameConstant)
	chroma.NameDecorator = toInline(cfg.NameDecorator)
	chroma.NameException = toInline(cfg.NameException)
	chroma.NameFunction = toInline(cfg.NameFunction)
	chroma.NameOther = toInline(cfg.NameOther)
	chroma.Literal = toInline(cfg.Literal)
	chroma.LiteralNumber = toInline(cfg.LiteralNumber)
	chroma.LiteralDate = toInline(cfg.LiteralDate)
	chroma.LiteralString = toInline(cfg.LiteralString)
	chroma.LiteralStringEscape = toInline(cfg.LiteralStringEscape)
	chroma.GenericDeleted = toInline(cfg.GenericDeleted)
	chroma.GenericEmph = toInline(cfg.GenericEmph)
	chroma.GenericInserted = toInline(cfg.GenericInserted)
	chroma.GenericStrong = toInline(cfg.GenericStrong)
	chroma.GenericSubheading = toInline(cfg.GenericSubheading)
	chroma.Background = toInline(cfg.Background)

	return ansi.StyleCodeBlock{
		StyleBlock: toStyle(cfg.BlockStyle),
		Theme:      cfg.Theme,
		Chroma:     &chroma,
	}
}

func toHeading(cfg HeadingStyle, level int) ansi.StyleBlock {
	base := toStyle(cfg.Base)

	if level >= 1 && level <= len(cfg.Levels) {
		override := toStyle(cfg.Levels[level-1])
		return mergeBlocks(base, override)
	}

	return base
}

func toList(l ListStyle) ansi.StyleList {
	prim := toInline(l.Styles)

	return ansi.StyleList{
		StyleBlock: ansi.StyleBlock{
			StylePrimitive: prim,
			Indent:         toUintPtr(l.LevelIndent),
		},
		LevelIndent: uint(l.LevelIndent),
	}
}

func toInline(i InlineStyle) ansi.StylePrimitive {
	return ansi.StylePrimitive{
		BlockPrefix:     i.BlockPrefix,
		BlockSuffix:     i.BlockSuffix,
		Prefix:          i.Prefix,
		Suffix:          i.Suffix,
		Color:           stringPtr(i.Color),
		BackgroundColor: stringPtr(i.Background),
		Bold:            boolPtr(i.Bold),
		Italic:          boolPtr(i.Italic),
		Underline:       boolPtr(i.Underline),
		CrossedOut:      boolPtr(i.Strikethrough),
		Format:          i.Format,
	}
}

func toStyle(b BlockStyle) ansi.StyleBlock {
	prim := ansi.StylePrimitive{
		BlockPrefix:     b.BlockPrefix,
		BlockSuffix:     b.BlockSuffix,
		Prefix:          b.Prefix,
		Suffix:          b.Suffix,
		Color:           stringPtr(b.Color),
		BackgroundColor: stringPtr(b.Background),
		Bold:            boolPtr(b.Bold),
		Italic:          boolPtr(b.Italic),
		Underline:       boolPtr(b.Underline),
		CrossedOut:      nil,
		Faint:           nil,
		Conceal:         nil,
		Overlined:       nil,
		Inverse:         nil,
		Blink:           nil,
		Format:          b.Format,
	}
	return ansi.StyleBlock{
		StylePrimitive: prim,
		Indent:         toUintPtr(b.Indent),
		IndentToken:    stringPtr(b.IndentToken),
		Margin:         toUintPtr(b.Margin),
	}
}

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func boolPtr(b bool) *bool {
	if !b {
		return nil
	}
	return &b
}

func toUintPtr(v int) *uint {
	u := uint(v)
	return &u
}

func mergeBlocks(a, b ansi.StyleBlock) ansi.StyleBlock {
	prim := a.StylePrimitive
	op := b.StylePrimitive

	if op.Prefix != "" {
		prim.Prefix = op.Prefix
	}
	if op.Suffix != "" {
		prim.Suffix = op.Suffix
	}
	if op.Color != nil {
		prim.Color = op.Color
	}
	if op.BackgroundColor != nil {
		prim.BackgroundColor = op.BackgroundColor
	}
	if op.Bold != nil {
		prim.Bold = op.Bold
	}
	if op.Italic != nil {
		prim.Italic = op.Italic
	}
	if op.Underline != nil {
		prim.Underline = op.Underline
	}
	if op.CrossedOut != nil {
		prim.CrossedOut = op.CrossedOut
	}
	if op.Faint != nil {
		prim.Faint = op.Faint
	}
	if op.Conceal != nil {
		prim.Conceal = op.Conceal
	}
	if op.Overlined != nil {
		prim.Overlined = op.Overlined
	}
	if op.Blink != nil {
		prim.Blink = op.Blink
	}
	if op.Inverse != nil {
		prim.Inverse = op.Inverse
	}
	if op.Format != "" {
		prim.Format = op.Format
	}
	if op.BlockPrefix != "" {
		prim.BlockPrefix = op.BlockPrefix
	}
	if op.BlockSuffix != "" {
		prim.BlockSuffix = op.BlockSuffix
	}

	if b.Indent != nil {
		a.Indent = b.Indent
	}

	if b.Margin != nil {
		a.Margin = b.Margin
	}

	if b.IndentToken != nil {
		a.IndentToken = b.IndentToken
	}

	a.StylePrimitive = prim
	return a
}
