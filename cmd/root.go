package cmd

import (
	"flag"
	"fmt"
	"log/slog"
	"strings"

	"github.com/Michal-Nithesh/miccord/internal/config"
	"github.com/Michal-Nithesh/miccord/internal/consts"
	"github.com/Michal-Nithesh/miccord/internal/logger"
	"github.com/Michal-Nithesh/miccord/internal/ui/root"
	"github.com/ayn2op/arikawa/v3/utils/ws"
	"github.com/ayn2op/tview"
	"github.com/gdamore/tcell/v3"
)

func Run() error {
	flag.Usage = func() {
		out := flag.CommandLine.Output()
		_, _ = fmt.Fprintf(out, "%s - %s\n\n", consts.DisplayName, consts.Tagline) //nolint:errcheck
		_, _ = fmt.Fprintf(out, "Usage:\n  %s [options]\n\n", consts.Name)         //nolint:errcheck
		_, _ = fmt.Fprintf(out, "Options:\n")                                      //nolint:errcheck
		flag.PrintDefaults()
	}

	configPath := flag.String("config-path", config.DefaultPath(), "path of the configuration file")
	logPath := flag.String("log-path", logger.DefaultPath(), "path of the log file")
	logLevel := flag.String("log-level", "info", "log level")
	themeFlag := flag.String("theme", "", "color theme mode: 'dark' (Night Owl) or 'light'")
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Println(buildVersion())
		return nil
	}

	var level slog.Level
	if err := level.UnmarshalText([]byte(*logLevel)); err != nil {
		return fmt.Errorf("invalid log level: %w", err)
	}
	ws.EnableRawEvents = level == slog.LevelDebug

	logFile, err := logger.Load(*logPath, level)
	if err != nil {
		return fmt.Errorf("failed to load logger: %w", err)
	}
	defer logFile.Close() //nolint:errcheck

	cfg, err := config.Load(*configPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	if *themeFlag != "" {
		if strings.EqualFold(*themeFlag, "light") || strings.EqualFold(*themeFlag, "dark") {
			cfg.SetThemeMode(*themeFlag)
		} else {
			return fmt.Errorf("invalid theme mode %q; supported modes are 'dark' and 'light'", *themeFlag)
		}
	}

	screen, err := tcell.NewScreen()
	if err != nil {
		return fmt.Errorf("failed to create screen: %w", err)
	}

	if err := screen.Init(); err != nil {
		return fmt.Errorf("failed to init screen: %w", err)
	}

	if cfg.Mouse {
		screen.EnableMouse()
	}
	screen.EnablePaste()
	if cfg.Notifications.WhenUnfocused {
		screen.EnableFocus()
	}

	return tview.NewApplication(root.NewModel(cfg), tview.WithScreen(screen)).Run()
}
