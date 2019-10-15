# Wazigate Health

![Wazigate Health](www/icons/health.svg)

The Wazigate Health service is a Wazigate App that submits gateway health data (CPU usage, memory usage, ...) and telemetry data (OS version, ...) to the Gateway API.
This creates a sensors called "Health".

## Health Data

| Field | Description |
| --- | --- |
| `cpu` | CPU usage, range 0 (0%) to 1 (100%). Example: `0.1708` (= 17%) |
| `boottime` | Boottime. Time the device booted. Example: `2019-10-14T10:55:43Z` |
| `telemetry` | Telemetry Data. See below for more information. |
| `disk` | Disk (SD Card) size and used bytes. Example: `{used: 6503288832, all: 7748247552}` (Total Size ~7GB, Used ~6GB) |

Telemetry Object:

| Field | Description |
| --- | --- |
| `osVersion` | OS Version. Example: `8 (jessie)` |
| `osVersionID` | OS Version ID. Example: `3.10.2` |
| `osName` | OS Name. Example: `Alpine Linux` |
| `osID` | OS Name. Example: `alpine` |

## Environment Variables

| Name | Description | Default Value |
| --- | --- | --- |
| `WAZIGATE_ADDR` | Local Address of the Wazigate Edge API. | `127.0.0.1:880` |

## Docker Hub

This service is available at the docker hub as [waziup/wazigat-health](https://hub.docker.com/r/waziup/wazigate-health).

Required Mappings:

| Name | Data |
| --- | --- |
| `/proc/stat` | CPU and Boottime. |
| `/etc/os-release` | Telemetry: OS information. |

