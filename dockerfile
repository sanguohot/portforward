FROM sanguohot/cgo:v1.12.4
WORKDIR /opt/portforward
COPY . .
RUN GOPROXY=https://goproxy.cn go build --ldflags "-linkmode external -extldflags -static" -o portforward bin/app/app.go

FROM busybox:1.28
WORKDIR /root/
COPY ca-certificates.crt /etc/ssl/certs/
COPY etc/config.yaml ./etc/
COPY --from=0 /opt/portforward/portforward .
ENTRYPOINT ["./portforward"]