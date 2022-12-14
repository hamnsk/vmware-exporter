FROM golang:1.18-alpine3.14 AS certs

FROM scratch
ENV VMWARE_EXPORTER_BIND_ADDR=0.0.0.0:9513
ENV VMWARE_EXPORTER_VMWARE_USER=monitoring
ENV VMWARE_EXPORTER_VMWARE_PASSWORD=strongPass
ENV VMWARE_EXPORTER_LOG_LEVEL=info
ENV VMWARE_EXPORTER_HTTP_WRITE_TIMEOUT=30s
ENV VMWARE_EXPORTER_HTTP_READ_TIMEOUT=30s
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
COPY vmware-exporter /app/vmware-exporter

CMD ["/app/vmware-exporter"]