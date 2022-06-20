# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
[markdownlint](https://dlaa.me/markdownlint/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## [0.0.2] - 2022-06-20

### Added in 0.0.2

- Migrated from `xyzzygoapi` to `g2-sdk-go`
- G2diagnosticImpl complete
- G2engineImpl
    - AddRecord
    - AddRecordWithInfo
    - ClearLastException
    - DeleteRecord
    - DeleteRecordWithInfo
    - Destroy
    - GetLastException
    - Init
    - Stats

## [0.0.1] - 2022-04-23

### Added in 0.0.1

- A single working call to Senzing.
- G2diagnosticImpl:
    - GetAvailableMemory
    - ClearLastException
    - GetLogicalCores
    - GetPhysicalCores
    - GetTotalSystemMemory
