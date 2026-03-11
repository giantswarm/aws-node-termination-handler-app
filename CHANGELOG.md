# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Add `io.giantswarm.application.audience: giantswarm` annotation.
- Migrate chart metadata annotations to `io.giantswarm.application.*` format.

## [1.23.1] - 2026-02-12

### Fixed

- Push the chart to the default catalog.

## [1.23.0] - 2025-10-06

### Changed

- Upgrade to v1.25.2 to fix a [resource leak bug](https://github.com/aws/aws-node-termination-handler/issues/1172)

## [1.22.0] - 2025-10-02

### Changed

- Enable heartbeats by default (interval 30 seconds, maximum 1 hour)
- Sync `values.yaml` with upstream

## [1.21.0] - 2025-06-24

### Changed

- Upgrade image to 1.25.1
- Upgrade chart from upstream
- Allow metrics port in network policy

## [1.20.0] - 2024-10-30

### Changed

- Use vendir to pull from upstream repository https://github.com/giantswarm/aws-node-termination-handler-upstream
- Upgrade image to 1.22.1

## [1.19.0] - 2024-03-26

### Added

- Add `global.podSecurityStandards.enforced` value for PSS migration.

[Unreleased]: https://github.com/giantswarm/aws-node-termination-handler-app/compare/v1.23.1...HEAD
[1.23.1]: https://github.com/giantswarm/aws-node-termination-handler-app/compare/v1.23.0...v1.23.1
[1.23.0]: https://github.com/giantswarm/aws-node-termination-handler-app/compare/v1.22.0...v1.23.0
[1.22.0]: https://github.com/giantswarm/aws-node-termination-handler-app/compare/v1.21.0...v1.22.0
[1.21.0]: https://github.com/giantswarm/aws-node-termination-handler-app/compare/v1.20.0...v1.21.0
[1.20.0]: https://github.com/giantswarm/aws-node-termination-handler-app/compare/v1.19.0...v1.20.0
[1.19.0]: https://github.com/giantswarm/aws-node-termination-handler-app/releases/tag/v1.19.0
