# Changelog

All notable changes to this project will be documented in this file.

## Unreleased

### Added

- `format` selects the output encoding: `console` (default) for humans, `json` for log shippers with one object per line and lowercase level names.
- Structured logging functions `Debugw`, `Infow`, `Warnw`, `Errorw`, `Panicw` and `Fatalw`, and `With` to attach the fields of a unit of work once.
  `Nop` implements them too.

### Changed

- Timestamps are ISO 8601 in UTC with millisecond precision in both formats.
- The console format is coloured only when standard output is a terminal, so container logs are plain text.
- Standard output is always written. `Console` is ignored.
- `File` and the rotation options are deprecated and log a warning when set. The file is still written this release.
- `SyncFileLogger` no longer logs.

### Fixed

- `Logger()` reports the caller of the line correctly when used directly.
