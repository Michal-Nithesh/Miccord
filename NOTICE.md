# Miccord Notice & Attribution

This project, **Miccord**, is an independent derivative work of **Discordo** (https://github.com/ayn2op/discordo), previously developed as MichalTerm.

## Copyright Notices

- **Discordo** (Original codebase):
  Copyright (C) 2025-present ayn2op

- **Miccord / MichalTerm** (Modifications & Rebranding):
  Copyright (C) 2026 Michal Nithesh and contributors

## License

This software is licensed under the **GNU General Public License Version 3** (GPL-3.0).
See the [LICENSE](./LICENSE) file for the complete terms and conditions.

## Summary of Modifications (October 2026)

In accordance with Section 5(a) of the GNU General Public License v3:
1. **Rebranding**: Rebranded product to **Miccord** ("Your conversations. Your terminal.") with executable `miccord`.
2. **Configuration & Storage Paths**: Updated configuration path to `$XDG_CONFIG_HOME/miccord/config.toml` (`~/.config/miccord/config.toml`) and cache path to `~/.cache/miccord/`, with backward-compatible fallback to existing MichalTerm and Discordo configurations.
3. **Authentication & Keyring**: Updated system keyring service to `miccord` (with automatic migration from `michalterm` and `discordo` without token exposure) and primary environment variable to `MICCORD_TOKEN` (with fallback to `MICHALTERM_TOKEN` and `DISCORDO_TOKEN`).
4. **Terminal Interface & Theme**: Refined default theme with subtle green accents and high-contrast styling optimized for Omarchy Linux with Hyprland.
5. **CLI & Documentation**: Added custom command-line help banner, formatted version information, and updated documentation.
