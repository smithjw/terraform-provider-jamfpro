# Changelog

## [0.14.0](https://github.com/deploymenttheory/terraform-provider-jamfpro/compare/v0.13.2...v0.14.0) (2025-03-25)


### Features

* added app installer global settings with example ([67021ed](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/67021ed10ad40e5211a5909a2919efb51f1610bb))
* added app installer global settings with examples ([a36273d](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/a36273da45d6d8ae1a0ac923a11a11405b19eb4e))


### Bug Fixes

* added debug logs ([6649eb6](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/6649eb653a7330f2bbb8b03faa9d41047806cf46))
* added plist root level PayloadUUID and PayloadIdentifier validation checker ([aecf0a2](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/aecf0a2ba896e55b8d64a08e2b553f7ab902b8eb))
* aligned mobile device config profiles with refined macos logic ([b0274f8](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/b0274f80fff9fedbefacc158af1ef9b6d0bbf6c7))
* final tweaks for mobile device plist handling ([6029e0b](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/6029e0b0e4346f9a145205c96691e110dfb02958))
* for '&' handling within plists for macos and mobile devices ([9704936](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/9704936df4019916e5eb4d13596f9b8e1ee523cd))
* removed not required html escaping from payloads for create operation ([a02b2a2](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/a02b2a2933b631e173be2849a856e12fd4c8d5a0))
* tidy up ([8eb572c](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/8eb572ca484a73d65b17f269dfee7b9d0d1f2e8f))
* updated failing test bugs ([84bad4f](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/84bad4f8ded71fda7e6e675cedcc3dfb1212f43b))

## [0.13.2](https://github.com/deploymenttheory/terraform-provider-jamfpro/compare/v0.13.1...v0.13.2) (2025-03-14)


### Bug Fixes

* dep to full version ([93a21e4](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/93a21e4c5b3b3a8a18d765d1e2e03198db0b01da))
* package uploads no longer timeout, metadata and upload are separate operations now. ([#621](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/621)) ([0058eaa](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/0058eaab4bc52f35fca73b1d0ca1343af755f28d))

## [0.13.1](https://github.com/deploymenttheory/terraform-provider-jamfpro/compare/v0.13.0...v0.13.1) (2025-03-12)


### Bug Fixes

* packages now have a custom 10 minute timeout ([#619](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/619)) ([8b49258](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/8b49258d57fe59542e4bc12b8bb433074bf39333))

## [0.13.0](https://github.com/deploymenttheory/terraform-provider-jamfpro/compare/v0.12.0...v0.13.0) (2025-03-04)


### Features

* API integration RefreshClientCredentialsByApiRoleID ([7fa1c2a](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/7fa1c2a53465a576bbaec2d5b71179ea8c020eab))
* centralized uuid handling ([9f9e10b](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/9f9e10bb1444a63440c2ddbda5d363369548807b))
* migrated restricted software from scope lists to scope sets meaning order of id's no longer matters ([6d2a37a](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/6d2a37a7b2e0fe8d756e4e87ba19bbc0b7f1792d))
* standardized mobile device plist behaviour with tf operational flow of macos config profiles ([ee83b36](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/ee83b36f0a18eb5da9b4d5f1c69bd778e3549f1e))
* updated restricted software example ([fabd224](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/fabd2245645134a9eb919d70476468c48f3c02e4))


### Bug Fixes

* fix:  ([b8e180f](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/b8e180f71c5b87aa7cc8f5321ba841249ce3a7cd))
* for plists that inserted escaped html characters during updates and removed validators for PlistPayloadDescription and PlistPayloadName. not required ([944b263](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/944b263d3e98e340ca74cb25f742f9db08e9a604))
* for plists that inserted escaped html characters during updates. removed validators for PlistPayloadDescription and PlistPayloadName. not required ([03c4a05](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/03c4a05cf1504f58d36ccaa8dbd35a21e403731e))
* migrated restricted software from scope lists to scope sets meaning the order of scope id's is no longer required to be ascending  in the hcl ([fc37373](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/fc373738c7deaf7a873c6c947e4cf4c915cce80c))
* reinstated PayloadDisplayName to diff suppression ([e613f80](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/e613f807e29b9f5e535cfe96d3022cf1eae51fdf))
* removed deletebyscriptname from the packages crud ([3258409](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/3258409c3c40c651653993880d1f96debd8b6c23))
* script deletions are now sequential ([#612](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/612)) ([add019d](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/add019de43121c5746bc00a43fce760eae74e8cd))
* standardized mobile device plists with macos ([19ac8da](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/19ac8da4842179673e55fe31e27ca7aee2ab0483))

## [0.12.0](https://github.com/deploymenttheory/terraform-provider-jamfpro/compare/v0.11.0...v0.12.0) (2025-02-21)


### Features

* added device communication settings resource ( ([507858b](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/507858bd700f61c436f3e03aa2e01b5007e60ab8))
* added enrollment prestage panes to enrolment customizations ([ceb79d0](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/ceb79d0fc65df219a7619d669b1c70ea79247d8e))
* added plist validation for PayloadName and PayloadDescription ([5709c85](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/5709c85ec991debcc44bdc077f59f5bfb67acac4))
* added resource inventory collection settings ([b4d43ee](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/b4d43eea5ff6597b7a1db4b2dafaa86f22c2327d))
* added SMTP server resource ([a21c0e0](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/a21c0e080729c0aeff40cf02805593062134c13c))
* migrated computer_checkin resource to client_checkin, fixes for missing api calls and migrated from classic to jamf pro api ([76ab64e](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/76ab64edd66c9a69509873faac24528f4672f483))


### Bug Fixes

* added comments and refactored update func ([c084c2e](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/c084c2e88149a3b031c7818e0faadb3c90c33398))
* added examples for enrollment customization ([83f87a4](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/83f87a4749613d8b5de0aa9934937e546dbbb23d))
* added resizing of enrollment customization images ([0ef2808](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/0ef28089f1f7f0b1ce4f772bcbd490403ee37529))
* removed redundant code ([f8d0af1](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/f8d0af15cd1c973da8bdd9b754c5bcc8f3acb3bf))
* removed redundant fmt.Sprintf ([8da1fa3](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/8da1fa31a2751d3eac93fdfc4911d577561cbeee))

## [0.11.0](https://github.com/deploymenttheory/terraform-provider-jamfpro/compare/v0.10.0...v0.11.0) (2025-01-24)


### Features

* add device enrollment data source for Jamf Pro ([eecd3f0](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/eecd3f0b36fc2b01954f57158770ad39b8ebc27f))
* add device enrollment data source for Jamf Pro ([#575](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/575)) ([8b78239](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/8b782392db59edd37fa832b76aadf7f38c37617a))
* enhance account and account group constructors with privilege validation and add fuzzy matching for invalid privileges ([465f85d](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/465f85dbee488f4d356f2e484acf5f09bc3c1498))


### Bug Fixes

* implemented dynamic lookup of api privileges  ([#572](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/572)) ([858be90](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/858be907564fcdf2f76d307d4629220f7b0f2d15))
* improve error messages and enhance example usage for Jamf Pro data sources ([5270fcc](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/5270fcc954aca36fab6689031499f1b919624dbc))
* refactored logic to define new pattern for data sources ([36ada49](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/36ada493b1aea8f419a3aa25b4e04ceba38cb2be))

## [0.10.0](https://github.com/deploymenttheory/terraform-provider-jamfpro/compare/v0.9.1...v0.10.0) (2025-01-20)


### Features

* added jamf connect to provider with examples ([#568](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/568)) ([d66439a](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/d66439a034d40550d10c8dafba9d37cd05902885))

## [0.9.1](https://github.com/deploymenttheory/terraform-provider-jamfpro/compare/v0.9.0...v0.9.1) (2025-01-06)


### Bug Fixes

* added computer_prestage_enrollment example ([#563](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/563)) ([b5fb211](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/b5fb2119276aae1e375ee191c840a97132512473))
* downgrade actions/setup-go to v5.2.0 in workflow configuration ([#564](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/564)) ([054da7c](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/054da7c00a5e1ab8250c0d54a90696e825f9cdca))
* macos plist handling for payloadUUID update operation ([#562](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/562)) ([a506164](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/a50616444e3335825ec082bafa57f4889d616f51))
* migrated doc gen to github action ([#554](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/554)) ([b5740eb](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/b5740eb7217bb354af7d1fbb772098fca6e9ff28))

## [0.9.0](https://github.com/deploymenttheory/terraform-provider-jamfpro/compare/v0.8.0...v0.9.0) (2024-12-23)


### Features

* add workflow_dispatch input for release version and update dependencies ([fe76925](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/fe769256ca701cc0c06a6068ade2f2d15039c35e))
* enhance security by adding Harden Runner to workflows and updat… ([#549](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/549)) ([9a4cd52](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/9a4cd529ceb32959d7917cb6ebdfc464f51c8bb9))
* enhance security by adding Harden Runner to workflows and update checkout action version ([9a4cd52](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/9a4cd529ceb32959d7917cb6ebdfc464f51c8bb9))
* pattern for data sources by name or by id with examples ([#540](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/540)) ([5838861](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/5838861eddeb97bce39b9199ba9802d1032ad140))
* update app installer retrieval method and add automated workflo… ([#544](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/544)) ([0d56b96](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/0d56b960ba9f3137a94028fb96ffbfb752f4cd94))
* update app installer retrieval method and add automated workflows for dependency management and documentation generation ([0d56b96](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/0d56b960ba9f3137a94028fb96ffbfb752f4cd94))
* updated data resources to support get by name with examples ([#553](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/553)) ([4049226](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/404922676b9dfe2ea8892b16db9e44e847237e2b))


### Bug Fixes

* added emoji's to runner titles for styling ([#551](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/551)) ([f41a421](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/f41a4216afa4b07a300234d24cce5ec40668b4ed))
* harden runners ([#550](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/550)) ([f0638b1](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/f0638b15a11f106ef66de7c31a22352876073964))
* pipeline testing for tf docs ([#546](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/546)) ([628b68b](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/628b68b790aab9561a391d85fbf2d6c8416a7fc0))

## [0.8.0](https://github.com/deploymenttheory/terraform-provider-jamfpro/compare/v0.7.1...v0.8.0) (2024-12-12)


### Features

* Add "Reinstall Button Text" field to self-service configuration ([00ba59d](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/00ba59d44d4b1db466ab08eecf4e79ceea717e0d))
* add data source lists for computer extension attributes, scripts, and webhooks ([223c1e0](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/223c1e0639d9602dfef3fa82716ebff5cd43bd5c))
* Add dependency on pre-release checks for Terraform provider rel… ([#509](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/509)) ([a1922d8](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/a1922d8bf6472cf771c8cafada3f70e6a2318bc7))
* Add dependency on pre-release checks for Terraform provider release workflow ([a1922d8](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/a1922d8bf6472cf771c8cafada3f70e6a2318bc7))
* add documentation for jamfpro_icon and jamfpro_managed_software_update resources ([403f1a4](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/403f1a4017cd0f0a91cb34ac7f00a37201f6899a))
* Add GitHub Actions workflow for automated release management ([45a555d](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/45a555dd75376f0f9e2259f3638be0876b532c48))
* add Jamf Pro icon resource with local and web source support ([0762db0](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/0762db0a789b9e8780537f7c1f5365c5becc2506))
* add list data sources for Jamf Pro scripts, webhooks, and computer extension attributes ([6bab5aa](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/6bab5aaf82c2a358f09b4a3c31f5415a9cbb7a69))
* Add new JSS settings and actions privileges ([6d411b0](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/6d411b0c967fc892a754845353c63f709afb75df))
* Add PayloadScope field to PayloadContent struct ([b49b423](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/b49b42306065636e34765d9ed8de7bade0902a6c))
* Add script to export Jamf Pro user account privileges ([cf4536c](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/cf4536c193ffc503e6aa2b5b227e3d856a2cd29c))
* add Terraform tests for independent computer extension attributes, sites, categories, scripts, and policies ([eaea1e8](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/eaea1e82427816e4ace25d3ac93ecfe4adfee264))
* add Terraform tests for independent computer extension attributes, sites, categories, scripts, and policies with dependencies ([b45466e](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/b45466eef1f7476a964f2d5a33b59a174012aa74))
* **docs:** add APP Installers resource details to README ([3ea8023](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/3ea8023499a553bc44045695e0efee9c338d714c))
* **docs:** add example usage for advanced search resources and create managed software update documentation ([8f79ba3](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/8f79ba37a0f8d29fa05674711479edcbbdb45c11))
* **docs:** update README with new resources and their management details ([a5e09a1](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/a5e09a103526725da26254a6e4521ac5d007ee22))
* enhance file cleanup logic for downloaded packages to ensure safety and improve logging ([7c065b8](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/7c065b8ae62352979afd3a538772af5f2cbd9d6f))
* enhance icon resource construction to prevent conflicting source specifications ([9feb4e3](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/9feb4e31191f2e1012d611aca2f8ff0f689c2041))
* implement DownloadFile function for downloading files from web sources ([2773589](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/2773589a87824af605b46c4af6821ef2f7801e36))
* Improve error message for macOS configuration profile level validation ([d113dbf](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/d113dbff96c5c31d62a57ab6187f43ce1755795b))
* **managedsoftwareupdates:** add managed software updates resource and documentation ([fcd71ce](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/fcd71ce5528ecfff846735e3b647944cf9ce0a8c))
* **managedsoftwareupdates:** add resource management for software updates ([455daca](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/455dacac7c28180301a0d78aff6ff492e9aef0ba))
* **managedsoftwareupdates:** refactor resource management and validation for software update plans ([3049c22](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/3049c221db9078e2291d228de34fd2f86e9ffb8f))
* Normalize payload XML in plist payload content ([0688e8d](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/0688e8d177d5734aa34fd9cc48307129167fd563))
* Normalize payload XML in plist payload content ([7b9e13c](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/7b9e13ca6e4a878c8f9f95b438df2967c05bdf59))
* **provider:** comment out managed software update resource in provider configuration ([0928ccc](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/0928ccc684dfe034b1933bedeb4a80c9a8f710a8))
* Refactor package create update crud ([#501](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/501)) ([2a4e463](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/2a4e463386a44be6f8291e1d7a0688732d29cfbb))
* remove outdated data source documentation for computer extension attributes, scripts, and webhooks ([398dd3a](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/398dd3a16db433ec616846ead910ca6e8106114c))
* **resources:** add advanced search resources for computers, mobile devices, users, and network segments to examples ([928d85c](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/928d85c2f4496d38a80ced6c8613da75ec25f0a8))
* **resources:** enhance managed software update configuration and validation logic ([5dcd1e1](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/5dcd1e167139022b9d9c1498507af381202d4d37))
* **resources:** enhance managed software update construction with conditional field assignments ([4eeb25d](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/4eeb25d4c89231fa6de1c1df6d7c5268253b558f))
* **resources:** enhance managed software update construction with default values for optional fields ([5f21d27](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/5f21d271c5d0ab7935ee3b9c765f51e71e8845d0))
* **resources:** enhance managed software update construction with device handling and improve validation error messages ([e88e60e](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/e88e60e43d3dba2cacd73bc4ed5ad1dd478a369f))
* **resources:** implement custom validation for managed software updates and enhance state management ([b244f05](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/b244f05618da1f3feed08e7b3df64aa64711565e))
* **resources:** refactor managed software update configuration to root level attributes and simplify state management ([161a805](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/161a8059948952d91efe75279c780dbfff6ae4ee))
* **resources:** streamline managed software update configuration by consolidating field assignments ([68206c6](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/68206c66965239ed6a9c325a1b28a8b4ae1edd78))
* Update Api Privileges data for Jamf Pro Version ${{ env.VERSION_DIR }} ([577f173](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/577f1732e8409258a4c13366a1caaeb656b1b762))
* Update Api Privileges data for Jamf Pro Version ${{ env.VERSION_DIR }} ([f285d5b](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/f285d5b7738d633c447a52b82629b25117ae29bd))
* Update branch name for Api Privileges and User Account Privileges maintenance workflows ([c1dc4d7](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/c1dc4d732ada817bf818c1270b315c45327657bf))
* Update go-api-sdk-jamfpro dependency to v1.15.4 and add no_exeute_on field to policy schema ([#510](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/510)) ([8577bff](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/8577bff646af404434b201abba5d8a409f06a9d2))
* update Jamf Pro Icon resource functions to improve file cleanup and logging ([5264d88](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/5264d8882f1fc836eb5a56c5d52a62be30b3d27d))
* Update privileges descriptions for Casper Admin ([195cbb2](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/195cbb275fed293c5feb22f2ec245ba939b1db27))
* Update User Account Privileges data for Jamf Pro ([5d069c8](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/5d069c8d7d1bb20871f48e99a2df33b6d61ebb81))


### Bug Fixes

* add default value for self_service_icon_id in MacOS configuration profiles resource ([bb6132e](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/bb6132ecab92330f821c4c634b20eac773df7418))
* add mutex locking to prevent concurrent creation of resources ([bb2d298](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/bb2d2982fbe6aa6d9f1aace5c4ad822ab4cb6966))
* change prestage_installed_profile_ids and custom_package_ids types from list to set for improved data handling and enforce ordering requirements ([a348a2b](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/a348a2b1964f7cbad6ccf215384f157ee7a18300))
* change self_service_categories type from list to set for improved data handling and consistency in MacOS configuration profiles ([825737c](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/825737ca0606d567ffbdf64162310919a163646b))
* change self_service_category type from list to set for improved data handling in MacOS configuration profiles and align with jamf pro behaviour ([fd9427d](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/fd9427ddbf653ac5f82725943d7092a7627a966d))
* correct spelling and improve clarity in Terraform test descriptions for script policies with dependencies ([b4e1952](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/b4e19523b742732dd821bf594bf0b46e7c1846f7))
* **deps:** update go-api-sdk-jamfpro to version 1.11.3 ([606b39c](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/606b39c56a3d3b9f40c1b175738018c2a5f5442e))
* enhance documentation for macOS and mobile device configuration profiles, improving clarity and detail on payload handling and diff suppression ([c46d7a9](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/c46d7a90aa9a7ea0414f03e1e06d4f565f4927e7))
* enhance validation and diff suppression for macOS configuration profiles, improving error handling and documentation ([ea147bc](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/ea147bce52dbceae6a80d688c55b1b1ce1f78b82))
* improve documentation clarity and formatting in macOS configuration profile plist ([d8afda1](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/d8afda1deb91083712d3c13c59dd187802953be3))
* improve documentation formatting and clarity in macOS configuration profiles resource ([6f11e4f](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/6f11e4f2098d31a0afc21f3964c31b9c70b939ab))
* improve logging in diff suppression for mobile device configuration profiles, enhancing clarity and error handling ([02b0a13](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/02b0a13ab5de92054dc19d31ad6c0389a957ee5e))
* incorrectly named SetId value in activationcode ([8d11c6e](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/8d11c6eb997c37cd5663bb192ca6384bcbf183bf))
* increase resource creation timeout from 70 to 120 seconds. added missing timeouts and standardised across all resource types apart from packages ([e5f8bbd](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/e5f8bbde27436b7b5ba6ca6875217e0b5320a901))
* refactor of state migration logic for Payload/UserInteraction ([#517](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/517)) ([ae80f2c](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/ae80f2c23dd763a7c527b0b0e7b214465cf4e752))
* release please with goreleaser ([#521](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/521)) ([1dcd6c0](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/1dcd6c0dd436514a838754a285ad13c4ed0c4b50))
* release please with goreleaser ([#524](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/524)) ([7c9d112](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/7c9d11209ec26616474f84b490975634a60c3e2c))
* removed mutex from all CRUD ops ([f187708](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/f187708d54f76286b9b23f4c56e8db25a409947b))
* ResourceJamfProActivationCode - add Sensitive attribute to "code" schema ([bbd9f31](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/bbd9f310645b92a2a0d5111f55e5d8d6da6f6e41))
* State migrator for policy schema tag typo fix ([#505](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/505)) ([61ffca9](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/61ffca9bce6377fba876474dcb9a6dee3a642716))
* **staticcomputergroups:** resource.Computers is always nil - do not gate creating it on this fact. It is also a pointer, so so set the value accordingly. ([22f0d0f](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/22f0d0fc2c97ff988dac48dedb49bb38931922e3))
* suppress diff for popup_menu_choices attribute in computerextensionattributes resource ([ed600c5](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/ed600c50e77a5584e656564b0374d2bc33b47205))
* temporarily incerased all CRUD context timeouts to 75 seconds as workaround for stating issue ([6da1790](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/6da17904dc64b0bbfb2196473980b77593b1b103))
* update activation code and policy documentation for clarity and sensitivity ([12c1bd2](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/12c1bd2549ff811bc717435fc186f85c3fb9e3ce))
* update advanced mobile device search resource schema with validation and improve payload handling in CRUD operations ([88a1fa5](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/88a1fa5c4b34adc22a674e7ba6488a78ce4baabc))
* update documentation for advanced mobile device search and computer prestage enrollment, changing list types to set for improved data handling and consistency ([b35dc71](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/b35dc71be81edfbe75bb20894f6f81e666dfb5d3))
* update error message for category construction failure in CRUD operations ([a12a9d9](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/a12a9d9a5e1d898cc6ecb32dcfa74eb66ffa0b40))
* update self-service display names and descriptions in multiple policy files ([2baaabe](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/2baaabe6b8951c599fcb9512e69578b52f918a48))
* update user interaction fields and improve XML logging in policy constructor ([599b4bb](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/599b4bbaf3d8d62215742162fb7b6c62087cb455))
* updated Policy resource datetime limitations key descriptions, validators + examples. ([#506](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/506)) ([cfaac1d](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/cfaac1d62807aa9cc2973aa0a52b4544460a44cf))

## [0.7.1](https://github.com/deploymenttheory/terraform-provider-jamfpro/compare/v0.7.0...v0.7.1) (2024-12-12)


### Bug Fixes

* release please with goreleaser ([#521](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/521)) ([1dcd6c0](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/1dcd6c0dd436514a838754a285ad13c4ed0c4b50))

## [0.7.0](https://github.com/deploymenttheory/terraform-provider-jamfpro/compare/v0.6.1...v0.7.0) (2024-12-12)


### Features

* Add dependency on pre-release checks for Terraform provider rel… ([#509](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/509)) ([a1922d8](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/a1922d8bf6472cf771c8cafada3f70e6a2318bc7))
* Add dependency on pre-release checks for Terraform provider release workflow ([a1922d8](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/a1922d8bf6472cf771c8cafada3f70e6a2318bc7))
* Update go-api-sdk-jamfpro dependency to v1.15.4 and add no_exeute_on field to policy schema ([#510](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/510)) ([8577bff](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/8577bff646af404434b201abba5d8a409f06a9d2))


### Bug Fixes

* refactor of state migration logic for Payload/UserInteraction ([#517](https://github.com/deploymenttheory/terraform-provider-jamfpro/issues/517)) ([ae80f2c](https://github.com/deploymenttheory/terraform-provider-jamfpro/commit/ae80f2c23dd763a7c527b0b0e7b214465cf4e752))

## 0.1.0 (Unreleased)

FEATURES:
