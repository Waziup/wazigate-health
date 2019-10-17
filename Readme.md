# Wazigate Health

![Wazigate Health](www/icons/health.svg)

[![Docker Pulls](https://img.shields.io/docker/pulls/waziup/wazigate-health?style=flat-square)](https://hub.docker.com/r/waziup/wazigate-health)
[![godoc reference](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/Waziup/wazigate-health/health)

The Wazigate Health service is a Wazigate App that submits gateway health data (CPU usage, memory usage, ...) and telemetry data (OS version, ...) to the Gateway API.
This creates a sensors called "Health".

## Health Data

| Field | Description |
| --- | --- |
| `cpu` | CPU usage, range 0 (0%) to 1 (100%). Example: `0.1708` (= 17%) |
| `cpuTemp` | CPU temperature in °C. Example: `55.844` |
| `mem` | Memory info. Example: `{used: 232036, all: 949448}` (Total Size ~1GB, Used ~200MB) |
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
| `/proc/meminfo` | Memory Info. |
| `/etc/os-release` | Telemetry: OS information. |
| `/sys/class/thermal/thermal_zone0/temp` | CPU temperature. |

## Commands

To download the latest image and run it:

```bash
# Download the image from docker
docker pull waziup/wazigate-health

# Run the wazigate-health docker image
docker run --rm -v /proc:/proc -v /etc/os-release:/etc/os-release  -v /sys/class/thermal/thermal_zone0/temp:/sys/class/thermal/thermal_zone0/temp --network host -it wazigate-health:latest
```

To build this service from scratch:

```bash
# Clone this repository from github
git clone https://github.com/Waziup/wazigate-health.git
cd wazigate-health

# Build the wazigate-health docker image
docker build -f "Dockerfile" -t wazigate-health:latest .

# Run the wazigate-health docker image
docker run --rm -v /proc:/proc -v /etc/os-release:/etc/os-release  -v /sys/class/thermal/thermal_zone0/temp:/sys/class/thermal/thermal_zone0/temp --network host -it wazigate-health:latest
```
