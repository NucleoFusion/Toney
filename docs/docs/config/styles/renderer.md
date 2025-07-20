# Markdown Renderer

With this config option you can configurate your markdown renderer. We use [glamour](https://charm.sh/glamour) so you can get an even more extensive config description there.

The toml tag is `[styles.renderer]`.

## Item Prefix

can change the item prefix using the `item_prefix` toml key.

```toml
item_prefix = "• "
```
## Enumeration Prefix 

can change the item prefix using the `enumeration_prefix` toml key.

```toml
enumeration_prefix = "• "
```

## Document 

Used to configure the styles for the whole document, and also default ones that are inherited by all children in the document.
Configure using the `[styles.renderer.document]`.

Document uses the BlockStyle, for all available fields see [here](#block-style).

Default config is:
```toml
[styles.renderer.document]
block_prefix = "\n"
block_suffix = "\n"
margin = 1
color = "#b4befe"
```

## Block Quote 

Used to configure the styles for the quotations ('>').
Configure using the `[styles.renderer.block_quote]`.

Block Quote uses the BlockStyle, for all available fields see [here](#block-style).

Default config is:
```toml
[styles.renderer.block_quote]
indent = 1
indent_token = "| "
```

## Headings

Used to configure the styles for the headings('#').

### Base

Default heading configuration, use the `[styles.renderer.heading.base]` toml key.

Base uses the BlockStyle, for all available fields see [here](#block-style).

Default config is:
```toml
[styles.renderer.heading.base]
block_suffix = "\n"
color = "#94e2d5"
bold = true
```

### Levels

Is an array holding custom configuraion for each heading level, **ORDER MATTERS**.
To add use the `[[styles.renderer.heading.levels]]` for each heading level.

Levels uses the BlockStyle, for all available fields see [here](#block-style).

Default config is:
```toml
[[styles.renderer.heading.levels]] # h1
prefix = " "
suffic = " "
color = "#74c7ec"
background = "#313244"
bold = true

[[styles.renderer.heading.levels]] # h2
prefix = "## "

[[styles.renderer.heading.levels]] # h3
prefix = "### "

[[styles.renderer.heading.levels]] # h4
prefix = "#### "

[[styles.renderer.heading.levels]] # h5
prefix = "##### "

[[styles.renderer.heading.levels]] # h6
prefix = "###### "
color = "#9399b2"
bold = false
```

## Horizontal Rule

Used to configure the styles for the horizontal rule('---').
Configure using the `[styles.renderer.horizontal_rule]`.

Horizontal Rule uses the InlineStyle, for all available fields see [here](#inline-style).

Default config is:
```toml
[styles.renderer.horizontal_rule]
color = "#585b70"
format = "\n────────\n"
```

## List 

Used to configure the styles for the list('- ' or '1. '), using the `[styles.renderer.list]` toml key.

### Level Indent 

can change the indentation of list items using the `level_indent` toml key.

Default config is:
```toml
level_indent = 2 
```

### Styles 

used for custom configuration of list style.
To add use the `[styles.renderer.list.styles]` toml tag.

Styles uses the InlineStyle, for all available fields see [here](#inline-style).

Default config is:
```toml
[[styles.renderer.heading.levels]] # h1
suffix = "\n"
color = "#cdd6f4"
bold = false
```

### Task 
<!-- TODO:  -->

Used to configure the styles for tasks('[ ]' or '[X]'), using the `[styles.renderer.list.task]` toml key.

**To be added, not yet supported.** you can change the content as strings, but colors would be lost, since color support is not yet added.
If you _do_ want to change, use 

`ticked = "• [✓] "` or `unticked = "• [ ] "`

## Link 

Used to configure the styles for links('\[x](url)').
Configure using the `[styles.renderer.link]`.

Link uses the InlineStyle, for all available fields see [here](#inline-style).

Default config is:
```toml
[styles.renderer.link]
color = "#89b4fa"
underline = true
```

## Image 

Used to configure the styles for images('!\[x](url)').
Configure using the `[styles.renderer.image]`.

Image uses the InlineStyle, for all available fields see [here](#inline-style).

Default config is:
```toml
[styles.renderer.image]
color = "#fab387"
underline = true
```

## Emph 

Used to configure the styles for links('* *' or '_ _').
Configure using the `[styles.renderer.emph]`.

Emph uses the InlineStyle, for all available fields see [here](#inline-style).

Default config is:
```toml
[styles.renderer.emph]
color = "#89b4fa"
italic = true
```

## Strong 

Used to configure the styles for strong('** **').
Configure using the `[styles.renderer.strong]`.

Strong uses the InlineStyle, for all available fields see [here](#inline-style).

Default config is:
```toml
[styles.renderer.strong]
color = "#89b4fa"
bold = true
```

## Strikethrough

Used to configure the styles for strikethroughs('~~ ~~').
Configure using the `[styles.renderer.strikethrough]`.

Strikethrough uses the InlineStyle, for all available fields see [here](#inline-style).

Default config is:
```toml
[styles.renderer.strikethrough]
color = "#89b4fa"
italic = true
```

## Code 

Used to configure the styles for the inline code block (``), configured using the `[styles.renderer.code]`.

Code uses the BlockStyle, for all available fields see [here](#block-style).

Default config is:
```toml
[styles.renderer.code]
prefix = " "
suffix = " "
color = "#89dceb"
background = "#11111b"
```

## Code Block 

<!-- TODO: Add all CodeBlockStyle fields -->

Used to configure the styles for the code block('\``` ```'). find all fields [here](#code-block-style)

### Code Block Style

used for custom configuration of text.
To add use the `[styles.renderer.code_block.block_style]` toml tag.

Block Styles uses the BlockStyle, for all available fields see [here](#block-style).

Default config is:
```toml
[styles.renderer.code_block.text]
background = "#11111b"
color = "#89dceb"
margin = 1 
```

### Text 

used for custom configuration of text.
To add use the `[styles.renderer.code_block.text]` toml tag.

Styles uses the InlineStyle, for all available fields see [here](#inline-style).

Default config is:
```toml
[styles.renderer.code_block.text]
color = "#cdd6f4"
```

## Table 

Used to configure the styles for tables, using the `[styles.renderer.table]` toml key.

### Center Seperator 

can change the center seperator for the table using the `center_seperator` toml key.

Default config is:
```toml
center_seperator = "|" 
```

### Column Seperator 

can change the column seperator for the table using the `column_seperator` toml key.

Default config is:
```toml
column_seperator = "|" 
```

### Row Seperator

can change the row seperator for the table using the `row_seperator` toml key.

Default config is:
```toml
row_seperator = "-" 
```

### Table Block Style 

used for custom configuration of table style.
To add use the `[styles.renderer.table.block_style]` toml tag.

Block Style uses the BlockStyle, for all available fields see [here](#block-style).

Default config is:
```toml
[styles.renderer.heading.block_style]
color = "#cdd6f4"
```

### Header 

used for custom configuration of table header.
To add use the `[styles.renderer.table.header]` toml tag.

Header uses the BlockStyle, for all available fields see [here](#block-style).

Default config is:
```toml
[styles.renderer.heading.header]
color = "#b4befe"
background = "#1e1e2e"
bold = true
```

### Cell 

used for custom configuration of each table cell.
To add use the `[styles.renderer.table.cell]` toml tag.

Cell uses the BlockStyle, for all available fields see [here](#block-style).

Default config is:
```toml
[styles.renderer.heading.cell]
color = "#cdd6f4"
background = "#1e1e2e"
```

## Block Style

Used to configure the styles for block elements, using the `[styles.renderer.block]` toml key.

### Format

Can change the format string for the block using the `format` TOML key.

```toml
format = ""  # **no default value**
```

### Block Prefix

Can change the block prefix using the `block_prefix` TOML key.

```toml
block_prefix = ""  # **no default value**
```

### Block Suffix

Can change the block suffix using the `block_suffix` TOML key.

```toml
block_suffix = ""  # **no default value**
```

### Prefix

Can change the prefix before the block using the `prefix` TOML key.

```toml
prefix = ""  # **no default value**
```

### Suffix

Can change the suffix after the block using the `suffix` TOML key.

```toml
suffix = ""  # **no default value**
```

### Indent Token

Can change the character used for indentation using the `indent_token` TOML key.

```toml
indent_token = ""  # **no default value**
```

### Margin

Can change the margin around the block using the `margin` TOML key.

```toml
margin = 0  # **no default value**
```

### Padding

Can change the padding inside the block using the `padding` TOML key.

```toml
padding = 0  # **no default value**
```

### Indent

Can set the number of indentation levels using the `indent` TOML key.

```toml
indent = 0  # **no default value**
```

### Color

Can change the text color using the `color` TOML key.

```toml
color = ""  # **no default value**
```

### Background

Can change the background color using the `background` TOML key.

```toml
background = ""  # **no default value**
```

### Bold

Can toggle bold formatting using the `bold` TOML key.

```toml
bold = false  # **no default value**
```

### Italic

Can toggle italic formatting using the `italic` TOML key.

```toml
italic = false  # **no default value**
```

### Underline

Can toggle underline formatting using the `underline` TOML key.

```toml
underline = false  # **no default value**
```

## Inline Style

Used to configure the styles for inline elements, using the `[styles.renderer.inline]` TOML key.

### Block Prefix

Can change the block prefix using the `block_prefix` TOML key.

```toml
block_prefix = ""  # **no default value**
```

### Block Suffix

Can change the block suffix using the `block_suffix` TOML key.

```toml
block_suffix = ""  # **no default value**
```

### Prefix

Can change the prefix before the inline element using the `prefix` TOML key.

```toml
prefix = ""  # **no default value**
```

### Suffix

Can change the suffix after the inline element using the `suffix` TOML key.

```toml
suffix = ""  # **no default value**
```

### Color

Can change the text color using the `color` TOML key.

```toml
color = ""  # **no default value**
```

### Background

Can change the background color using the `background` TOML key.

```toml
background = ""  # **no default value**
```

### Bold

Can toggle bold formatting using the `bold` TOML key.

```toml
bold = false  # **no default value**
```

### Italic

Can toggle italic formatting using the `italic` TOML key.

```toml
italic = false  # **no default value**
```

### Underline

Can toggle underline formatting using the `underline` TOML key.

```toml
underline = false  # **no default value**
```

### Strikethrough

Can toggle strikethrough formatting using the `strikethrough` TOML key.

```toml
strikethrough = false  # **no default value**
```

### Format

Can set a format string using the `format` TOML key.

```toml
format = ""  # **no default value**
```
