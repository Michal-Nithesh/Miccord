package config

import (
	"testing"

	"github.com/BurntSushi/toml"
)

func TestThemePresets(t *testing.T) {
	t.Run("NightOwlTheme defaults", func(t *testing.T) {
		theme := NightOwlTheme()
		if theme.Mode != "dark" {
			t.Fatalf("expected Mode = 'dark', got %q", theme.Mode)
		}
		if theme.GuildsTree.GraphicsColor != "#1d3b53" {
			t.Fatalf("expected graphics_color = '#1d3b53', got %q", theme.GuildsTree.GraphicsColor)
		}

		data, err := toml.Marshal(theme)
		if err != nil {
			t.Fatalf("failed to marshal NightOwlTheme: %v", err)
		}
		var decoded ThemeConfig
		if err := toml.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("failed to unmarshal NightOwlTheme: %v", err)
		}
		if decoded.Mode != "dark" {
			t.Fatalf("expected decoded Mode = 'dark', got %q", decoded.Mode)
		}
	})

	t.Run("LightOwlTheme defaults", func(t *testing.T) {
		theme := LightOwlTheme()
		if theme.Mode != "light" {
			t.Fatalf("expected Mode = 'light', got %q", theme.Mode)
		}
		if theme.GuildsTree.GraphicsColor != "#d3d8df" {
			t.Fatalf("expected graphics_color = '#d3d8df', got %q", theme.GuildsTree.GraphicsColor)
		}

		data, err := toml.Marshal(theme)
		if err != nil {
			t.Fatalf("failed to marshal LightOwlTheme: %v", err)
		}
		var decoded ThemeConfig
		if err := toml.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("failed to unmarshal LightOwlTheme: %v", err)
		}
		if decoded.Mode != "light" {
			t.Fatalf("expected decoded Mode = 'light', got %q", decoded.Mode)
		}
	})

	t.Run("SetThemeMode switching", func(t *testing.T) {
		cfg := Default()
		cfg.SetThemeMode("light")
		if cfg.Theme.Mode != "light" {
			t.Fatalf("expected Theme.Mode = 'light', got %q", cfg.Theme.Mode)
		}
		if cfg.Markdown.Theme != "github" {
			t.Fatalf("expected Markdown.Theme = 'github', got %q", cfg.Markdown.Theme)
		}

		cfg.SetThemeMode("dark")
		if cfg.Theme.Mode != "dark" {
			t.Fatalf("expected Theme.Mode = 'dark', got %q", cfg.Theme.Mode)
		}
		if cfg.Markdown.Theme != "monokai" {
			t.Fatalf("expected Markdown.Theme = 'monokai', got %q", cfg.Markdown.Theme)
		}
	})
}
