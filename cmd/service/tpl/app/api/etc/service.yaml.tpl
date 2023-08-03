Name: {{ .serviceName }}.api
Host: 0.0.0.0
Port: {{ .apiPort }}

Log:
  Encoding: plain
  TimeFormat: "2006-01-02 15:04:05 MST -07"

DB:
  Host: mysql
  Port: 3306
  Name: {{ .serviceName }}_service
  User: root
  Password: root
  LogLevel: 4
  SlowThreshold: 1

CorsOrigins:
  - "http://localhost:{{ .docPort }}"
  - "https://localhost:{{ .docPort }}"

JWTXRpc:
  Etcd:
    Hosts:
      - etcd:2379
    Key: jwtx.rpc
