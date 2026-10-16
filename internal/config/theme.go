package config

import (
	"encoding/json"
	"errors"
	"slices"
	"strings"

	"github.com/ayn2op/arikawa/v3/discord"
	"github.com/ayn2op/tview"
	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

var errInvalidType = errors.New("invalid type")

var (
	attributes = enum[tcell.AttrMask]{
		{"bold", tcell.AttrBold},
		{"dim", tcell.AttrDim},
		{"italic", tcell.AttrItalic},
		{"blink", tcell.AttrBlink},
		{"reverse", tcell.AttrReverse},
		{"strikethrough", tcell.AttrStrikeThrough},
	}
	underlines = enum[tcell.UnderlineStyle]{
		{"", tcell.UnderlineStyleNone},
		{"solid", tcell.UnderlineStyleSolid},
		{"double", tcell.UnderlineStyleDouble},
		{"curly", tcell.UnderlineStyleCurly},
		{"dotted", tcell.UnderlineStyleDotted},
		{"dashed", tcell.UnderlineStyleDashed},
	}
)

type StyleWrapper struct{ tview.Style }

func (sw *StyleWrapper) UnmarshalTOML(v any) error {
	m, ok := v.(map[string]any)
	if !ok {
		return errInvalidType
	}

	// Reset on new styles
	sw.Style = tcell.StyleDefault

	for key, val := range m {
		switch key {
		case "foreground":
			if s, ok := val.(string); ok {
				sw.Style = sw.Foreground(tcell.GetColor(s))
			}
		case "background":
			if s, ok := val.(string); ok {
				sw.Style = sw.Background(tcell.GetColor(s))
			}
		case "attributes":
			switch val := val.(type) {
			case string:
				sw.parseAttr(val)
			case []any:
				for _, attr := range val {
					if s, ok := attr.(string); ok {
						sw.parseAttr(s)
					}
				}
			}
		case "underline":
			if s, ok := val.(string); ok {
				underline, err := underlines.parse(s)
				if err != nil {
					return err
				}
				sw.Style = sw.Underline(underline)
			}
		case "underline_color":
			if s, ok := val.(string); ok {
				sw.Style = sw.Underline(tcell.GetColor(s))
			}
		}
	}

	return nil
}

func (sw *StyleWrapper) parseAttr(s string) {
	if s == "underline" {
		sw.Style = sw.Underline(true)
		return
	}
	// An unknown attribute parses to no attribute.
	attr, _ := attributes.parse(s)
	sw.Style = sw.Attributes(sw.GetAttributes() | attr) //nolint:staticcheck
}

// MarshalTOML writes the style as the inline table UnmarshalTOML reads.
func (sw StyleWrapper) MarshalTOML() ([]byte, error) {
	var fields []string
	add := func(key string, value any) {
		b, _ := json.Marshal(value)
		fields = append(fields, key+" = "+string(b))
	}
	if c := sw.GetForeground(); c != tcell.ColorDefault {
		add("foreground", colorName(c))
	}
	if c := sw.GetBackground(); c != tcell.ColorDefault {
		add("background", colorName(c))
	}
	var attrs []string
	for _, a := range attributes {
		if sw.GetAttributes()&a.value != 0 {
			attrs = append(attrs, a.name)
		}
	}
	switch len(attrs) {
	case 0:
	case 1:
		add("attributes", attrs[0])
	default:
		add("attributes", attrs)
	}
	if u := sw.GetUnderlineStyle(); u != tcell.UnderlineStyleNone {
		name, _ := underlines.text(u)
		add("underline", string(name))
	}
	if c := sw.GetUnderlineColor(); c != tcell.ColorDefault {
		add("underline_color", colorName(c))
	}
	if fields == nil {
		return []byte("{}"), nil
	}
	return []byte("{ " + strings.Join(fields, ", ") + " }"), nil
}

// colorName returns the first of the names of c, or its hex code if it has none.
func colorName(c tcell.Color) string {
	var names []string
	for name, color := range tcell.ColorNames {
		if color == c {
			names = append(names, name)
		}
	}
	if names == nil {
		return c.CSS()
	}
	return slices.Min(names)
}

type (
	HelpThemeConfig struct {
		ShortKeyStyle  StyleWrapper `toml:"short_key_style"`
		ShortDescStyle StyleWrapper `toml:"short_desc_style"`
		FullKeyStyle   StyleWrapper `toml:"full_key_style"`
		FullDescStyle  StyleWrapper `toml:"full_desc_style"`
	}

	StyleThemeConfig struct {
		NormalStyle StyleWrapper `toml:"normal_style"`
		// The style used while the widget is focused.
		ActiveStyle StyleWrapper `toml:"active_style"`
	}

	MessagesListThemeConfig struct {
		AuthorStyle      StyleWrapper `toml:"author_style"`
		MentionStyle     StyleWrapper `toml:"mention_style"`
		EmojiStyle       StyleWrapper `toml:"emoji_style"`
		URLStyle         StyleWrapper `toml:"url_style"`
		AttachmentStyle  StyleWrapper `toml:"attachment_style"`
		ReactionStyle    StyleWrapper `toml:"reaction_style"`
		OwnReactionStyle StyleWrapper `toml:"own_reaction_style"`

		MessageStyle         StyleWrapper `toml:"message_style"`
		SelectedMessageStyle StyleWrapper `toml:"selected_message_style"`

		Embeds MessagesListEmbedsThemeConfig `toml:"embeds"`
	}

	MessagesListEmbedsThemeConfig struct {
		ProviderStyle    StyleWrapper `toml:"provider_style"`
		AuthorStyle      StyleWrapper `toml:"author_style"`
		TitleStyle       StyleWrapper `toml:"title_style"`
		DescriptionStyle StyleWrapper `toml:"description_style"`
		FieldNameStyle   StyleWrapper `toml:"field_name_style"`
		FieldValueStyle  StyleWrapper `toml:"field_value_style"`
		FooterStyle      StyleWrapper `toml:"footer_style"`
		URLStyle         StyleWrapper `toml:"url_style"`
	}

	MentionsListThemeConfig struct {
		SelectedStyle StyleWrapper `toml:"selected_style"`
	}

	DialogThemeConfig struct {
		Style StyleWrapper `toml:"style"`
		// The style of everything behind the dialog.
		BackgroundStyle StyleWrapper `toml:"background_style"`
	}

	ScrollBarThemeConfig struct {
		TrackStyle StyleWrapper `toml:"track_style"`
		ThumbStyle StyleWrapper `toml:"thumb_style"`
	}

	ThemeConfig struct {
		// Theme mode: "dark" (Night Owl) or "light" (Night Owl Light).
		Mode string `toml:"mode"`

		Title  StyleThemeConfig `toml:"title"`
		Footer StyleThemeConfig `toml:"footer"`
		Border StyleThemeConfig `toml:"border"`
		// The status styles are those of direct-message recipients.
		GuildsTree GuildsTreeThemeConfig `toml:"guilds_tree"`
		// The status styles are those of the dot before each member.
		MembersTree  TreeThemeConfig         `toml:"members_tree"`
		ScrollBar    ScrollBarThemeConfig    `toml:"scroll_bar"`
		MessagesList MessagesListThemeConfig `toml:"messages_list"`
		MentionsList MentionsListThemeConfig `toml:"mentions_list"`
		Dialog       DialogThemeConfig       `toml:"dialog"`
		Help         HelpThemeConfig         `toml:"help"`
	}
)

type GuildsTreeThemeConfig struct {
	TreeThemeConfig
	// The style of whoever is speaking in the voice channel joined.
	SpeakingStyle StyleWrapper `toml:"speaking_style"`
	// The style of how full a voice channel that has a user limit is.
	VoiceLimitStyle StyleWrapper `toml:"voice_limit_style"`
}

type TreeThemeConfig struct {
	// The color of the lines that give the tree its shape.
	GraphicsColor string `toml:"graphics_color"`

	OnlineStyle  StyleWrapper `toml:"online_style"`
	IdleStyle    StyleWrapper `toml:"idle_style"`
	DNDStyle     StyleWrapper `toml:"dnd_style"`
	OfflineStyle StyleWrapper `toml:"offline_style"`
}

// StatusStyle returns the style for a user with status.
func (ctt TreeThemeConfig) StatusStyle(status discord.Status) tview.Style {
	switch status {
	case discord.OnlineStatus:
		return ctt.OnlineStyle.Style
	case discord.IdleStatus:
		return ctt.IdleStyle.Style
	case discord.DoNotDisturbStatus:
		return ctt.DNDStyle.Style
	default:
		return ctt.OfflineStyle.Style
	}
}

// NightOwlTheme returns the dark Night Owl color theme.
func NightOwlTheme() ThemeConfig {
	var (
		plain  = tcell.StyleDefault
		bold   = plain.Bold(true)
		italic = plain.Italic(true)

		fg       = color.GetColor("#d6deeb")
		cyan     = color.GetColor("#7fdbca")
		blue     = color.GetColor("#82aaff")
		green    = color.GetColor("#addb67")
		yellow   = color.GetColor("#ecc48d")
		red      = color.GetColor("#ef5350")
		purple   = color.GetColor("#c792ea")
		muted    = color.GetColor("#5f7e97")
		border   = color.GetColor("#1d3b53")
		dialogBg = color.GetColor("#0b2942")

		focus = StyleThemeConfig{
			NormalStyle: StyleWrapper{plain.Foreground(border)},
			ActiveStyle: StyleWrapper{bold.Foreground(cyan)},
		}
		title = StyleThemeConfig{
			NormalStyle: StyleWrapper{plain.Foreground(blue)},
			ActiveStyle: StyleWrapper{bold.Foreground(cyan)},
		}
		footer = StyleThemeConfig{
			NormalStyle: StyleWrapper{plain.Foreground(muted)},
			ActiveStyle: StyleWrapper{bold.Foreground(cyan)},
		}
		tree = TreeThemeConfig{
			GraphicsColor: "#1d3b53",
			OnlineStyle:   StyleWrapper{plain.Foreground(green)},
			IdleStyle:     StyleWrapper{plain.Foreground(yellow)},
			DNDStyle:      StyleWrapper{plain.Foreground(red)},
			OfflineStyle:  StyleWrapper{plain.Foreground(color.GetColor("#637777"))},
		}
	)

	return ThemeConfig{
		Mode:   "dark",
		Title:  title,
		Footer: footer,
		Border: focus,
		GuildsTree: GuildsTreeThemeConfig{
			TreeThemeConfig: tree,
			SpeakingStyle:   StyleWrapper{bold.Foreground(green)},
			VoiceLimitStyle: StyleWrapper{plain.Foreground(blue)},
		},
		MembersTree: tree,
		ScrollBar: ScrollBarThemeConfig{
			TrackStyle: StyleWrapper{plain.Foreground(color.GetColor("#0b2942"))},
			ThumbStyle: StyleWrapper{plain.Foreground(cyan)},
		},
		MessagesList: MessagesListThemeConfig{
			AuthorStyle:          StyleWrapper{bold.Foreground(blue)},
			MentionStyle:         StyleWrapper{bold.Foreground(yellow)},
			EmojiStyle:           StyleWrapper{plain.Foreground(green)},
			URLStyle:             StyleWrapper{plain.Foreground(blue).Underline(true)},
			AttachmentStyle:      StyleWrapper{plain.Foreground(green)},
			ReactionStyle:        StyleWrapper{plain.Foreground(muted)},
			OwnReactionStyle:     StyleWrapper{bold.Foreground(cyan)},
			MessageStyle:         StyleWrapper{plain.Foreground(fg)},
			SelectedMessageStyle: StyleWrapper{plain.Foreground(color.GetColor("#ffffff")).Background(border)},
			Embeds: MessagesListEmbedsThemeConfig{
				ProviderStyle:    StyleWrapper{italic.Foreground(muted)},
				AuthorStyle:      StyleWrapper{italic.Foreground(purple)},
				TitleStyle:       StyleWrapper{bold.Foreground(blue)},
				DescriptionStyle: StyleWrapper{plain.Foreground(fg)},
				FieldNameStyle:   StyleWrapper{bold.Foreground(cyan).Underline(true)},
				FieldValueStyle:  StyleWrapper{plain.Foreground(fg)},
				FooterStyle:      StyleWrapper{italic.Foreground(muted)},
				URLStyle:         StyleWrapper{plain.Foreground(blue).Underline(true)},
			},
		},
		MentionsList: MentionsListThemeConfig{
			SelectedStyle: StyleWrapper{bold.Foreground(cyan).Background(border)},
		},
		Dialog: DialogThemeConfig{
			Style:           StyleWrapper{plain.Foreground(fg).Background(dialogBg)},
			BackgroundStyle: StyleWrapper{plain.Dim(true).Background(color.GetColor("#01111d"))},
		},
		Help: HelpThemeConfig{
			ShortKeyStyle:  StyleWrapper{bold.Foreground(cyan)},
			ShortDescStyle: StyleWrapper{plain.Foreground(muted)},
			FullKeyStyle:   StyleWrapper{bold.Foreground(blue)},
			FullDescStyle:  StyleWrapper{plain.Foreground(fg)},
		},
	}
}

// LightOwlTheme returns the light color theme (Night Owl Light).
func LightOwlTheme() ThemeConfig {
	var (
		plain  = tcell.StyleDefault
		bold   = plain.Bold(true)
		italic = plain.Italic(true)

		fg       = color.GetColor("#403f53")
		teal     = color.GetColor("#0c969b")
		blue     = color.GetColor("#4876d6")
		green    = color.GetColor("#2aa298")
		amber    = color.GetColor("#c96b00")
		red      = color.GetColor("#e05666")
		purple   = color.GetColor("#9042f5")
		muted    = color.GetColor("#7e8a99")
		border   = color.GetColor("#d3d8df")
		selBg    = color.GetColor("#d6deeb")
		dialogBg = color.GetColor("#f0f0f0")

		focus = StyleThemeConfig{
			NormalStyle: StyleWrapper{plain.Foreground(border)},
			ActiveStyle: StyleWrapper{bold.Foreground(teal)},
		}
		title = StyleThemeConfig{
			NormalStyle: StyleWrapper{plain.Foreground(blue)},
			ActiveStyle: StyleWrapper{bold.Foreground(teal)},
		}
		footer = StyleThemeConfig{
			NormalStyle: StyleWrapper{plain.Foreground(muted)},
			ActiveStyle: StyleWrapper{bold.Foreground(teal)},
		}
		tree = TreeThemeConfig{
			GraphicsColor: "#d3d8df",
			OnlineStyle:   StyleWrapper{plain.Foreground(green)},
			IdleStyle:     StyleWrapper{plain.Foreground(amber)},
			DNDStyle:      StyleWrapper{plain.Foreground(red)},
			OfflineStyle:  StyleWrapper{plain.Foreground(color.GetColor("#90a4ae"))},
		}
	)

	return ThemeConfig{
		Mode:   "light",
		Title:  title,
		Footer: footer,
		Border: focus,
		GuildsTree: GuildsTreeThemeConfig{
			TreeThemeConfig: tree,
			SpeakingStyle:   StyleWrapper{bold.Foreground(green)},
			VoiceLimitStyle: StyleWrapper{plain.Foreground(blue)},
		},
		MembersTree: tree,
		ScrollBar: ScrollBarThemeConfig{
			TrackStyle: StyleWrapper{plain.Foreground(color.GetColor("#e0e5eb"))},
			ThumbStyle: StyleWrapper{plain.Foreground(teal)},
		},
		MessagesList: MessagesListThemeConfig{
			AuthorStyle:          StyleWrapper{bold.Foreground(blue)},
			MentionStyle:         StyleWrapper{bold.Foreground(amber)},
			EmojiStyle:           StyleWrapper{plain.Foreground(green)},
			URLStyle:             StyleWrapper{plain.Foreground(blue).Underline(true)},
			AttachmentStyle:      StyleWrapper{plain.Foreground(color.GetColor("#498205"))},
			ReactionStyle:        StyleWrapper{plain.Foreground(muted)},
			OwnReactionStyle:     StyleWrapper{bold.Foreground(teal)},
			MessageStyle:         StyleWrapper{plain.Foreground(fg)},
			SelectedMessageStyle: StyleWrapper{plain.Foreground(color.GetColor("#011627")).Background(selBg)},
			Embeds: MessagesListEmbedsThemeConfig{
				ProviderStyle:    StyleWrapper{italic.Foreground(muted)},
				AuthorStyle:      StyleWrapper{italic.Foreground(purple)},
				TitleStyle:       StyleWrapper{bold.Foreground(blue)},
				DescriptionStyle: StyleWrapper{plain.Foreground(fg)},
				FieldNameStyle:   StyleWrapper{bold.Foreground(teal).Underline(true)},
				FieldValueStyle:  StyleWrapper{plain.Foreground(fg)},
				FooterStyle:      StyleWrapper{italic.Foreground(muted)},
				URLStyle:         StyleWrapper{plain.Foreground(blue).Underline(true)},
			},
		},
		MentionsList: MentionsListThemeConfig{
			SelectedStyle: StyleWrapper{bold.Foreground(teal).Background(selBg)},
		},
		Dialog: DialogThemeConfig{
			Style:           StyleWrapper{plain.Foreground(fg).Background(dialogBg)},
			BackgroundStyle: StyleWrapper{plain.Background(color.GetColor("#ffffff")).Dim(true)},
		},
		Help: HelpThemeConfig{
			ShortKeyStyle:  StyleWrapper{bold.Foreground(teal)},
			ShortDescStyle: StyleWrapper{plain.Foreground(muted)},
			FullKeyStyle:   StyleWrapper{bold.Foreground(blue)},
			FullDescStyle:  StyleWrapper{plain.Foreground(fg)},
		},
	}
}

func defaultTheme() ThemeConfig {
	return NightOwlTheme()
}
