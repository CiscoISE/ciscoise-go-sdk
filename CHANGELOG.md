# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
## [1.1.18] - 2023-02-10
Parameter ID added to:
- `ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionChildren`

## [1.1.17] - 2023-02-10
Parameter ID added to:
- `ResponseNetworkAccessAuthorizationRulesGetNetworkAccessAuthorizationRuleByIDResponseRuleConditionChildren`
- `ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionChildren`
- `ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionChildren`


## [1.1.16] - 2023-01-30
### Added
Following parameters were added: `ID` in following services:
- `network_access_authentication_rules`
- `network_access_authorization_rules`
- `network_access_policy_set`

## [1.1.15] - 2023-01-30
### Added
Following parameters were added: `ID` in following services:
- `network_access_authorization_rules`

## [1.1.14] - 2022-10-27
### Added
Following parameters were added: `DictionaryName`,`AttributeName`, `Operator`, `AttributeValue` in following services:
- `device_administration_authentication_rules`
- `device_administration_authorization_exception_rules`
- `device_administration_authorization_global_exception_rules`
- `device_administration_authorization_rules`
- `network_access_authorization_rules`
- `network_access_policy_set`

## [1.1.13] - 2022-10-26
### Added
Following parameters were added to `RequestDeviceAdministrationPolicySetUpdateDeviceAdminPolicySetByIDConditionChildren` struct: `DictionaryName`,`AttributeName`, `Operator`, `AttributeValue`.

## [1.1.12] - 2022-10-26
### Added
Following parameters were added to `RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleCondition` struct: `DictionaryName`,`AttributeName`, `Operator`, `AttributeValue`.

## [1.1.11] - 2022-10-21
### Added 
- On `network_device` change `Othername` parameter to `Ndgtype`.

## [1.1.10] - 2022-10-21
### Added 
- On `active_directory` change `GroupName` parameter to `Name`.

## [1.1.9] - 2022-10-03
### Added 
- Add parameter `description` to `RequestNetworkAccessPolicySetUpdateNetworkAccessPolicySetByIDCondition` struct on `NetworkAccessPolicySetService`.

## [1.1.8] - 2022-09-30
### Added 
- Add parameter `description` to `RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetCondition` struct on `NetworkAccessPolicySetService`.

## [1.1.7] - 2022-09-29
### Added 
- Add parameters `ID` and `Name` to `RequestNetworkAccessPolicySetCreateNetworkAccessPolicySetCondition` struct on `NetworkAccessPolicySetService`.
  
## [1.1.6] - 2022-06-07
### Fixed
- Fixed `network_access_policy_set` functions.

## [1.1.5] - 2022-06-07
### Fixed
- Fixed ISE version 3.1.1 to 3.1_Patch_1 which is the correct version name.

## [1.1.4] - 2022-05-11
## Changed
-  Parameter `Arguments` of `RequestTacacsCommandSets` structs now can be empty.

## [1.1.3] - 2022-02-10

### Changed
- Update go-resty/resty/v2 from 2.6.0 to 2.7.0.

## [1.1.2] - 2021-12-17
### Fixed
- Fix return values for POST/PUT operations in cases where the response is an empty string rather than the expected structure.


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
[1.1.2]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v1.1.1...v1.1.2
[1.1.3]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v1.1.2...v1.1.3
[1.1.4]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v1.1.3...v1.1.4
[1.1.5]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v1.1.4...v1.1.5
[1.1.6]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v1.1.5...v1.1.6
[1.1.7]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v1.1.6...v1.1.7
[1.1.8]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v1.1.7...v1.1.8
[1.1.9]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v1.1.8...v1.1.9
[1.1.10]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v1.1.9...v1.1.10
[1.1.11]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v1.1.10...v1.1.11
[1.1.12]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v1.1.11...v1.1.12
[1.1.13]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v1.1.11...v1.1.12
[1.1.14]: https://github.com/CiscoISE/ciscoise-go-sdk/compare/v1.1.11...v1.1.12
[Unreleased]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v1.1.14...main
