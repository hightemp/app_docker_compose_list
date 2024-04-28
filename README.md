# app_docker_compose_list

Console app for printing containers which are used in docker-compose.

```
~/Projects/app_docker_compose_list$ go run main.go
+--------+---------------------------------------------+----------------------------------+---------+------------------------+-----------------+----------------------------------------------------------------------+
| ID     | IMAGE                                       | NAME                             | STATE   | STATUS                 | PORTS           | DOCKER COMPOSE PATH                                                  |
+--------+---------------------------------------------+----------------------------------+---------+------------------------+-----------------+----------------------------------------------------------------------+
| 0216cc | postgres:15.4                               | [/standalone_teable-db_1]        | running | Up 2 hours (healthy)   |  42345          | /home/hightemp/Projects/10.teable/teable/dockers/examples/standalone |
| e13c69 | ser_proxy-proxy                             | [/ser_proxy-proxy-1]             | running | Up 4 weeks             |  8080 1080      | /home/hightemp/Projects/ser_proxy                                    |
| 4f599b | ghcr.io/dreamacro/clash                     | [/clash-clash-1]                 | running | Up 4 weeks             |  7890 7891      | /home/hightemp/ForTesting/clash                                      |
| 5b36b0 | budibase/proxy                              | [/bbproxy]                       | running | Up 4 weeks             |  10001          | /home/hightemp/ForTesting/budibase                                   |
| c754f5 | minio/minio                                 | [/budibase_minio-service_1]      | running | Up 4 weeks (healthy)   |                 | /home/hightemp/ForTesting/budibase                                   |
| ba224a | redis                                       | [/budibase_redis-service_1]      | running | Up 4 weeks             |                 | /home/hightemp/ForTesting/budibase                                   |
| b8b845 | ibmcom/couchdb3                             | [/budibase_couchdb-service_1]    | running | Up 4 weeks             |                 | /home/hightemp/ForTesting/budibase                                   |
| c26224 | containrrr/watchtower                       | [/budibase_watchtower-service_1] | running | Up 4 weeks             |                 | /home/hightemp/ForTesting/budibase                                   |
| 9f57e8 | budibase/couchdb                            | [/budi-couchdb3-dev]             | running | Up 4 weeks             |  4005           | /home/hightemp/ForTesting/budibase/hosting                           |
| 2b9176 | docker.io/apitable/backend-server:latest    | [/apitable-backend-server-1]     | running | Up 4 weeks (unhealthy) |                 | /home/hightemp/ForTesting/apitable/apitable                          |
| 7e0ee1 | docker.io/apitable/room-server:latest       | [/apitable-room-server-1]        | running | Up 4 weeks             |                 | /home/hightemp/ForTesting/apitable/apitable                          |
| ddc33f | minio/minio:RELEASE.2023-01-25T00-19-54Z    | [/apitable-minio-1]              | running | Up 4 weeks (unhealthy) |                 | /home/hightemp/ForTesting/apitable/apitable                          |
| fa3951 | redis:7.0.8                                 | [/apitable-redis-1]              | running | Up 4 weeks             |                 | /home/hightemp/ForTesting/apitable/apitable                          |
| 27ad2b | docker.io/apitable/web-server:latest        | [/apitable-web-server-1]         | running | Up 4 weeks             |                 | /home/hightemp/ForTesting/apitable/apitable                          |
| 152f7b | rabbitmq:3.11.9-management                  | [/apitable-rabbitmq-1]           | running | Up 4 weeks             |                 | /home/hightemp/ForTesting/apitable/apitable                          |
| c38f5a | mysql:8.0.32                                | [/apitable-mysql-1]              | running | Up 4 weeks (healthy)   |                 | /home/hightemp/ForTesting/apitable/apitable                          |
| b370cd | docker.io/apitable/imageproxy-server:latest | [/apitable-imageproxy-server-1]  | running | Up 4 weeks             |                 | /home/hightemp/ForTesting/apitable/apitable                          |
| 3447f3 | agentgpt_platform                           | [/platform]                      | running | Up 4 weeks             |  8002           | /home/hightemp/ForTesting/AgentGPT                                   |
| 021a5a | mysql:8.0                                   | [/db]                            | running | Up 4 weeks             |  3307           | /home/hightemp/ForTesting/AgentGPT                                   |
| e52396 | portainer/portainer-ce:latest               | [/portainer]                     | running | Up 4 weeks             |  8000 9000 9443 | /home/hightemp/ForTesting/portainer                                  |
| d66fc7 | louislam/uptime-kuma:1                      | [/uptime-kuma]                   | running | Up 4 weeks (healthy)   |  3001           | /home/hightemp/ForTesting/uptime-kuma                                |
+--------+---------------------------------------------+----------------------------------+---------+------------------------+-----------------+----------------------------------------------------------------------+
```