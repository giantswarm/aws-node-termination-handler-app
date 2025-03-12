# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Upgrade image to 1.25.0
- Upgrade chart from upstream
- Allow metrics port in network policy

## [1.20.0] - 2024-10-30

### Changed

- Use vendir to pull from upstream repository https://github.com/giantswarm/aws-node-termination-handler-upstream
- Upgrade image to 1.22.1

## [1.19.0] - 2024-03-26

### Added

- Add `global.podSecurityStandards.enforced` value for PSS migration.

## [1.18.0] - 2023-07-13

### Fixed

- Add required values for pss policies.

## [1.17.2-gs2] - 2023-07-05

### Added

- Add Service Monitor.

### Changed

 - Defined the use of the RuntimeDefault Seccompprofiles in the pod and container security context.

## [1.17.2-gs1] - 2022-09-14

### Changed

- Bump to upstream version 1.17.2 and related chart template.

## [1.16.5-gs1] - 2022-07-01

### Changed

- Bump to upstream version 1.16.5 and related chart template.

## [0.4.0] - 2022-06-15

### Changed

- Remove `imagePullSecrets`

## [0.3.1] - 2022-04-20

### Changed

- Push to `giantswarm` catalog instead of `control-plane` one.

## [0.3.0] - 2022-04-20

### Changed

- Allow customizing the container image registry domain.

## [0.2.0] - 2022-03-21

### Added

- Add VerticalPodAutoscaler CR.

## [0.1.0] - 2022-02-28

### Added

- First version.

[Unreleased]: https://github.com/giantswarm/aws-node-termination-handler-app/compare/v1.20.0...HEAD
[1.20.0]: https://github.com/giantswarm/aws-node-termination-handler-app/compare/v1.19.0...v1.20.0
[1.19.0]: https://github.com/giantswarm/aws-node-termination-handler-app/compare/v1.18.0...v1.19.0
[1.18.0]: https://github.com/giantswarm/aws-node-termination-handler-app/compare/v1.17.2-gs2...v1.18.0
[1.17.2-gs2]: https://github.com/giantswarm/aws-node-termination-handler-app/compare/v1.17.2-gs1...v1.17.2-gs2
[1.17.2-gs1]: https://github.com/giantswarm/aws-node-termination-handler-app/compare/v1.16.5-gs1...v1.17.2-gs1
[1.16.5-gs1]: https://github.com/giantswarm/aws-node-termination-handler-app/compare/v0.4.0...v1.16.5-gs1
[0.4.0]: https://github.com/giantswarm/aws-node-termination-handler-app/compare/v0.3.1...v0.4.0
[0.3.1]: https://github.com/giantswarm/aws-node-termination-handler-app/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/giantswarm/aws-node-termination-handler-app/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/aws-node-termination-handler-app/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/aws-node-termination-handler-app/compare/v0.0.0...v0.1.0
