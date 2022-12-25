[![Go Checks](https://github.com/hamnsk/vmware-exporter/actions/workflows/go_checks.yml/badge.svg?branch=main)](https://github.com/hamnsk/vmware-exporter/actions/workflows/go_checks.yml)
[![Release](https://github.com/hamnsk/vmware-exporter/actions/workflows/release.yml/badge.svg)](https://github.com/hamnsk/vmware-exporter/actions/workflows/release.yml)

# VMWare Exporter 
Collect metrics from ESXi Host and VM's performance metrics

> Attention!
> This exporter has not been tested with vCloud Director due to the fact that I do not have a license for it.
> If someone can test this functionality I will be glad. 
> It would be ideal to have access to it for further product development.

## Screenshots
### Cluster Overview
![Screen 1](./assets/screenshots/VMware%20Cluster%20Owerview%201.png "Cluster Overview")
![Screen 2](./assets/screenshots/VMware%20Cluster%20Owerview%202.png "Cluster Overview")
![Screen 3](./assets/screenshots/VMware%20Cluster%20Owerview%203.png "Cluster Overview")

## Road Map

* Create Grafana Dashboards
* Refactor and Optimize Code

## Build

### Binaries

```shell
$ GOOS=linux GOARCH=amd64 make build
```

### Docker
```shell
docker build -f ./Dockerfile.local -t vmware-exporter .
```

## Running this software

### From binaries
Download or build the most suitable binary from [the releases tab](https://github.com/hamnsk/vmware-exporter/releases)

Then:

```shell
$ ENV_VARS ./.bin/vmware-exporter
```
ENV_VARS - [see](#configuration) list of environment variables for configure exporter

### Using the docker image

```shell
$ docker run --rm \
  -p 9513/tcp \
  --name vmware_exporter \
  --env VMWARE_EXPORTER_BIND_ADDR=0.0.0.0:9513 \
  --env VMWARE_EXPORTER_VMWARE_USER=monitoring \
  --env VMWARE_EXPORTER_VMWARE_PASSWORD=strongPass \
  --env VMWARE_EXPORTER_LOG_LEVEL=info \
  --env VMWARE_EXPORTER_HTTP_WRITE_TIMEOUT=30s \
  --env VMWARE_EXPORTER_HTTP_READ_TIMEOUT=30s \
  --env VMWARE_EXPORTER_SCRAPE_TIMEOUT=30s \
  ghcr.io/hamnsk/vmware-exporter:latest
```


### Checking the results
Visiting [http://localhost:9513/probe?target=esxi.host.name.or.ipaddr
](http://localhost:9513/probe?target=esxi.host.name.or.ipaddr)
for collect metrics from ESXi host

Then visit [http://localhost:9513/metrics](http://localhost:9513/metrics)
will return metrics for a main app

## Prometheus Configuration

VMWare exporter implements the multi-target exporter pattern, so we advice
to read the guide [Understanding and using the multi-target exporter pattern
](https://prometheus.io/docs/guides/multi-target-exporter/) to get the general
idea about the configuration.

The vmware exporter needs to be passed the target as a parameter, this can be
done with relabelling.

Example config:

```yaml
scrape_configs:
  - job_name: 'vmware_exporter'
    metrics_path: /probe
    static_configs:
      - targets:
        - esxi0001.localdomain
        - esxi0002.tech
        - 192.168.1.17
    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_target
      - source_labels: [__param_target]
        target_label: instance
      - target_label: __address__
        replacement: 127.0.0.1:9513  # The vmware exporter's real hostname:port.
```

## Configuration

VMWare exporter is configured via a environment variables.

### Environment variables
| ENV VARIABLE                       | DEFAULT    | DESCRIPTION                                       |
|------------------------------------|------------|---------------------------------------------------|
| VMWARE_EXPORTER_BIND_ADDR          | :9513      | Exporter bind address in format XXX.XXX.XXX.XXX:PORT |
| VMWARE_EXPORTER_VMWARE_USER        | monitoring | ESXi user name                                    |         
| VMWARE_EXPORTER_VMWARE_PASSWORD    | password   | ESXi user password                                | 
| VMWARE_EXPORTER_LOG_LEVEL          | info       | Log level e.g. in info,warn,error,debug           |
| VMWARE_EXPORTER_HTTP_WRITE_TIMEOUT | 30s        | Time duration e.g. 30s or 1m | 
| VMWARE_EXPORTER_HTTP_READ_TIMEOUT  | 30s        | Time duration e.g. 30s or 1m |  
| VMWARE_EXPORTER_SCRAPE_TIMEOUT     | 60s        | Time duration e.g. 30s or 1m |  


