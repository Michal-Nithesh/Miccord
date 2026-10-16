# Miccord

> **Your conversations. Your terminal.**

Miccord is an independent, community-built terminal client project inspired by terminal-based chat applications, designed for Linux enthusiasts, particularly on Arch Linux and Omarchy Linux with Hyprland.

Featuring a keyboard-first workflow and a refined dark theme with subtle green accents, Miccord lets you chat, manage channels, browse servers, and listen to voice channels without leaving your terminal.

![Preview](.github/preview.png)
![Picker](.github/picker.png)

## Highlights

- **Lightweight & Fast**: Built in Go using a high-performance terminal UI.
- **Omarchy & Hyprland Ready**: Keyboard-centric design, responsive layout adaptation, and clean contrast with subtle green accents.
- **Rich Feature Set**: Servers (guilds), text & voice channels, markdown formatting, autocomplete mentions, reactions, and attachments.
- **Flexible Authentication**: Log in via QR code, password, token prompt, or environment variable.
- **Backward-Compatible Migration**: Smoothly falls back to and migrates configuration and stored tokens from existing MichalTerm and Discordo installations.

## Installation

### Building from Source

Requires [Go](https://go.dev/dl) (version 1.27+). On Linux, voice support optionally requires a C compiler and `libpulse` or `libasound` (PulseAudio / ALSA).

```sh
# Clone repository
git clone https://github.com/Michal-Nithesh/miccord.git
cd miccord

# Build executable
go build -trimpath -ldflags="-s -w" -o miccord .

# Optional: Install to your local bin path
install -Dm755 miccord ~/.local/bin/miccord
```

#### Build Tags

- `-tags no_voice`: Build without voice audio dependencies.
- `-tags no_spoof_tls_fingerprint`: Build without TLS fingerprint spoofing.

## Usage

Start Miccord simply by running:

```sh
miccord
```

### Command Line Options

```text
Miccord - Your conversations. Your terminal.

Usage:
  miccord [options]

Options:
  -config-path string
        path of the configuration file (default: ~/.config/miccord/config.toml)
  -log-level string
        log level (debug, info, warn, error) (default "info")
  -log-path string
        path of the log file (default: ~/.cache/miccord/logs.txt)
  -version
        print version and exit
```

### Authentication Methods

1. **QR Code (Recommended for mobile users)**:
   - Run `miccord`.
   - Select the **QR** tab and scan the code using the Discord mobile app.
2. **Token (UI)**:
   - Run `miccord`.
   - Select the **Token** tab, enter your token, and click Login (the token is securely stored in your system keyring).
3. **Password**:
   - Run `miccord`.
   - Enter your email/phone and password in the **Password** tab.
4. **Environment Variable**:
   - Set `MICCORD_TOKEN` (or legacy `MICHALTERM_TOKEN` / `DISCORDO_TOKEN`):

```sh
export MICCORD_TOKEN="your_discord_token_here"
miccord
```

## Configuration

Miccord loads its configuration from:

- **Linux / Unix**: `$XDG_CONFIG_HOME/miccord/config.toml` or `~/.config/miccord/config.toml`
- **Darwin (macOS)**: `~/Library/Application Support/miccord/config.toml`
- **Windows**: `%AppData%/miccord/config.toml`

### Migration & Fallback

If no configuration file exists at `~/.config/miccord/config.toml`, Miccord will automatically check for and fall back to your existing configuration in the following order:
1. `~/.config/michalterm/config.toml`
2. `~/.config/discordo/config.toml`

Similarly, saved authentication tokens in your system keyring will automatically be migrated from `michalterm` or `discordo` to `miccord` without logging or exposing credentials.

For all configurable options, themes, keybindings, and layout settings, see the [Configuration Reference](./docs/config.md).

## Logs

Logs are written to:

- **Linux / Unix**: `$XDG_CACHE_HOME/miccord/logs.txt` or `~/.cache/miccord/logs.txt`
- **Darwin (macOS)**: `~/Library/Caches/miccord/logs.txt`
- **Windows**: `%LocalAppData%/miccord/logs.txt`

## License & Attribution

Miccord is open-source software licensed under the **GNU General Public License v3.0 (GPL-3.0)**.
See the [LICENSE](./LICENSE) file for complete details.

### Attribution

Miccord is based on and derived from [Discordo](https://github.com/ayn2op/discordo), originally authored by **ayn2op**:
- Copyright (C) 2025-present ayn2op
- Copyright (C) 2026 Michal Nithesh and Miccord contributors

All original notices, license grants, and architectural foundations are preserved in accordance with the GNU General Public License v3.0.

## Disclaimer

> [!IMPORTANT]
> Using automated user accounts ("self-bots") or third-party clients may be against Discord's Terms of Service. Miccord does not automate normal user account actions and connects as an interactive terminal client. Use responsibly and at your own discretion.
