[![Go Checks](https://github.com/hamnsk/vmware-exporter/actions/workflows/go_checks.yml/badge.svg?branch=main)](https://github.com/hamnsk/vmware-exporter/actions/workflows/go_checks.yml)
[![Release](https://github.com/hamnsk/vmware-exporter/actions/workflows/release.yml/badge.svg)](https://github.com/hamnsk/vmware-exporter/actions/workflows/release.yml)

# VMWare Exporter 
Collect metrics from ESXi Host and VM's performance metrics

> Attention!
> This exporter has not been tested with vCloud Director due to the fact that I do not have a license for it.
> If someone can test this functionality I will be glad. 
> It would be ideal to have access to it for further product development.

## Features
* Get hypervisor metrics
* Get performance metrics for virtual machines
* **Using Hashicorp Vault to Store Authorization Data on Hypervisors**
* Can work without a Hashicorp Vault if the login and password for all hypervisors are the same,
and you are not embarrassed by passing them through environment variables
* Works in multi-target mode

## Screenshots
### Cluster Overview
![Screen 1](./assets/screenshots/VMware%20Cluster%20Owerview%201.png "Cluster Overview")
![Screen 2](./assets/screenshots/VMware%20Cluster%20Owerview%202.png "Cluster Overview")
![Screen 3](./assets/screenshots/VMware%20Cluster%20Owerview%203.png "Cluster Overview")

### ESXi Hosts Information
![Screen 1](./assets/screenshots/VMware%20ESX%20Hosts%20Information%201.png "ESXi Hosts Information")
![Screen 2](./assets/screenshots/VMware%20ESX%20Hosts%20Information%202.png "ESXi Hosts Information")
![Screen 3](./assets/screenshots/VMware%20ESX%20Hosts%20Information%203.png "ESXi Hosts Information")

### VM Information
![Screen 1](./assets/screenshots/VMware%20VM%20Information%201.png "VM Information")
![Screen 2](./assets/screenshots/VMware%20VM%20Information%202.png "VM Information")

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
$ docker build -f ./Dockerfile.local -t vmware-exporter .
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

### Usage with Hashicorp Vault as credentials storage
For authorization in Vault, the App Role method is used.
Refer to the Vault documentation to set up this method.

```shell
$ vault auth enable approle
```

For authorization on the destination ESXI host,
the exporter will find two secrets "username" and "password" in the path you specified.
The path must be specified with the target host.

```shell
$ vault kv put {{kv2_engine_name}}/{{path}}/{{target}} username="monitoring"
$ vault kv put {{kv2_engine_name}}/{{path}}/{{target}} password="monitoring"
```

For example if we want to write authorization data for the esxi0001 host,
we need to specify to execute the following commands:

```shell
$ vault kv put secrets/esxi/monitoring/esxi0001 username="monitoring"
$ vault kv patch secrets/esxi/monitoring/esxi0001 password="monitoring"
```

Then create a key tree access policy,
we need read access to the keys, and a listing of available targets

```shell
$ vault policy write -tls-skip-verify vmwere-exporter -<<EOF
  path "secrets/esxi/monitoring/*" {
  capabilities = [ "read", "list" ]
  }
  EOF
```

Then we need to install the policy created in the previous step for our role.
The lifetime of the token and secret can be set to your taste,
but the Vault developers recommend setting a short lifetime.
The exporter can automatically extend the lifetime of the token.
The "bind_secret_id" parameter must be set to false,
this is necessary in order to pass only the "role_id" parameter
when configuring the exporter, the exporter will do the rest.
It is also recommended to set the CIDR of the subnets
where the exporter will be launched,
in order to exclude the possibility of obtaining a secret and a token from another place.
Can be limited to 1 host, such as where the exporter is running.

```shell
$ vault write -tls-skip-verify auth/approle/role/vmware-exporter \
  token_policies="vmwere-exporter" \
  token_ttl=5m \
  token_max_ttl=10m \
  secret_id_bound_cidrs="0.0.0.0/0","127.0.0.1/32" \
  token_bound_cidrs="0.0.0.0/0","127.0.0.1/32" \
  secret_id_ttl=5m policies="vmwere-exporter"\
  bind_secret_id=false
```

Then check your App Role

```shell
$ vault read -tls-skip-verify auth/approle/role/vmware-exporter-demo
```

If all is well, then get the "role_id" parameter
to set the value of the environment variable VMWARE_EXPORTER_VAULT_ROLE_ID.

```shell
$ vault read -tls-skip-verify auth/approle/role/vmware-exporter/role-id
```

### Finally, result screenshots
![Screen 1](./assets/screenshots/Vault%20Settings%201.png "Vault Settings")
![Screen 2](./assets/screenshots/Vault%20Settings%202.png "Vault Settings")


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
| ENV VARIABLE                            | DEFAULT    | DESCRIPTION                                                 |
|-----------------------------------------|------------|-------------------------------------------------------------|
| VMWARE_EXPORTER_BIND_ADDR               | :9513      | Exporter bind address in format XXX.XXX.XXX.XXX:PORT        |
| VMWARE_EXPORTER_VMWARE_USER             | monitoring | ESXi user name when not using a Vault                       |         
| VMWARE_EXPORTER_VMWARE_PASSWORD         | password   | ESXi user password when not using a Vault                   | 
| VMWARE_EXPORTER_LOG_LEVEL               | info       | Log level e.g. in info,warn,error,debug                     |
| VMWARE_EXPORTER_HTTP_WRITE_TIMEOUT      | 30s        | Time duration e.g. 30s or 1m                                | 
| VMWARE_EXPORTER_HTTP_READ_TIMEOUT       | 30s        | Time duration e.g. 30s or 1m                                |  
| VMWARE_EXPORTER_USE_VAULT               | not set    | Bool flag use or not Hashicorp Vault as credentials storage |  
| VMWARE_EXPORTER_VAULT_ADDR              | not set    | Address of Hashicorp Vault                                  |  
| VMWARE_EXPORTER_VAULT_AUTH_NAME         | not set    | Name of authorization method APP Role for example: appauth  |  
| VMWARE_EXPORTER_VAULT_ROLE_ID           | not set    | App Role ID                                                 |  
| VMWARE_EXPORTER_VAULT_SECRET_STORE_NAME | not set    | Name of kv2 secret store for example: secrets               |  
| VMWARE_EXPORTER_VAULT_SECRET_STORE_PATH | not set    | Path to root of credentials for example esxi/monitoring     |  


## VMWare Api Reference

* [vSphere Web Services API](https://developer.vmware.com/apis/704/vsphere/api_versions_all_index.html#dataObjects)
