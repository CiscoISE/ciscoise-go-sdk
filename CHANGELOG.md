# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.1.1] - 2021-12-17
### Fixed
- Fix GuestUserCustomFields structs from interface{} to map[string]interface{}

## [1.1.0] - 2021-12-13

### Added
- Add support to ISE version 3.1.1.
- Add Licensing service.
- Add NodeServices service.
- Add Patching service.
- Add Proxy service.
- Add Telemetry service.
- Add aliases to previous structures and functions.

### Removed
- Remove ReplicationStatus service.
- Remove SyncIseNode service.

### Changed
- Change previous GetTaskStatus is GetAllTaskStatus.
- Change previous GetTaskStatusByID is GetTaskStatus.

## [1.0.0] - 2021-12-13

### Added
- Add os.SetEnv error handling
- Add RestyClient function 
- Add NbarApp service
- Add SgVnMapping service
- Add VirtualNetwork service
- Add VnVLANMapping service

### Changed
- Update examples
- Update documentation


## [0.0.8] - 2021-09-28
### Changed
- Update examples
- Update documentation
- Add os.SetEnv error handling


## [0.0.7] - 2021-09-27
### Fixed
- Update FileSize type
- Fix GetCertificateTemplateByName path

## [0.0.6] - 2021-09-27
### Fixed
- Change AvailableSSIDs from string to []string [3e430fc]
- Update GenerationID type for sgt [d459c4e]

## [0.0.5] - 2021-09-23
### Fixed
- Fix AllowedInterfaces for *_portal structs [13b25c2]
- Fix CustomAttributes for endpoint structs [8e0683f]

## [0.0.4] - 2021-09-10
### Fixed
- Add request body to consumer and provider functions 796e94b
- Fix UpdateList type df1e0b4
- Update DownloadableACLByID struct 9f9832d
- Update ResponseMiscGetMntVersion.TypeOfNode type 2fd9075

### Added
- Add pointer to response fields df1e0b4

### Changed
- Updates structures of responses 263a4c5 (certificates, device_administration_time_date_conditions and network_access_time_date_conditions)


[0.0.4]: https://github.com/CiscoISE/ciscoise-go-sdk/commits/v0.0.4
[0.0.5]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v0.0.4...v0.0.5
[0.0.6]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v0.0.5...v0.0.6
[0.0.7]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v0.0.6...v0.0.7
[0.0.8]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v0.0.7...v0.0.8
[1.0.0]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v0.0.8...v1.0.0
[1.1.0]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v1.0.0...v1.1.0
[1.1.1]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v1.1.0...v1.1.1
[Unreleased]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v1.1.1...main
