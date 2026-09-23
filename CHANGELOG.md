# Changelog

All notable changes to this project will be documented in this file.

## Unreleased

### Added

- Structured logging functions `Debugw`, `Infow`, `Warnw`, `Errorw`, `Panicw` and `Fatalw`, and `With` to attach the fields of a unit of work once.
  `Nop` implements them too.

### Fixed

- `Logger()` reports the caller of the line correctly when used directly.
