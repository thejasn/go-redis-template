# go-redis-template

```
Sample application.yaml

redis:
  cluster:
    enabled: ${REDIS_CLUSTER_ENABLED:true}
    # Comma separated list of cluster nodes.
    node-addresses: ${REDIS_CLUSTER_NODES:172.16.0.32:30001,172.16.0.32:30002,172.16.0.32:30003,172.16.0.32:30004,172.16.0.32:30005,172.16.0.32:30006}
  single:
    # Host and port for the single redis instance.
    host: ${REDIS_HOST:172.16.0.32}
    port: ${REDIS_PORT:6379}
  idle-connection-timeout: 20000
  min-idle-connections: 10
  connect-timeout: 20000
  read-timeout: 2000
  write-timeout: 3000
  max-redirects: 2
  max-retries: 10
  pool-size: 128
  pool-timeout: 20000

```
