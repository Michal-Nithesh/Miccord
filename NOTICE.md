# MichalTerm Notice & Attribution

This project, **MichalTerm**, is a derivative work of **Discordo** (https://github.com/ayn2op/discordo).

## Copyright Notices

- **Discordo** (Original codebase):
  Copyright (C) 2025-present ayn2op

- **MichalTerm** (Modifications & Rebranding):
  Copyright (C) 2026 Michal Nithesh and contributors

## License

This software is licensed under the **GNU General Public License Version 3** (GPL-3.0).
See the [LICENSE](./LICENSE) file for the complete terms and conditions.

## Summary of Modifications (October 2026)

In accordance with Section 5(a) of the GNU General Public License v3:
1. **Rebranding**: Rebranded product to MichalTerm ("Your conversations. Your terminal.") with executable `michalterm`.
2. **Configuration & Storage Paths**: Updated configuration path to `$XDG_CONFIG_HOME/michalterm/config.toml` (`~/.config/michalterm/config.toml`) and cache path to `~/.cache/michalterm/`, with backward-compatible fallback to existing Discordo configuration.
3. **Authentication & Keyring**: Updated system keyring service to `michalterm` (with automatic migration from `discordo`) and primary environment variable to `MICHALTERM_TOKEN` (with fallback to `DISCORDO_TOKEN`).
4. **Terminal Interface & Theme**: Refined default theme with subtle green accents and high-contrast styling optimized for Omarchy Linux with Hyprland.
5. **CLI & Documentation**: Added custom command-line help banner, formatted version information, and updated documentation.
