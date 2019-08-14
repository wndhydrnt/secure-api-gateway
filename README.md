# secure-api-gateway

## Requirements

* curl
* docker
* docker-compose
* go

## Usage

```
$ make package_oauth2_introspection
$ docker-compose up -d
Creating network "secure-api-gateway_default" with the default driver
Creating secure-api-gateway_app_1                  ... done
Creating secure-api-gateway_oathkeeper_1           ... done
Creating secure-api-gateway_oauth2-introspection_1 ... done
Creating secure-api-gateway_ambassador_1           ... done
$ curl -i -H'Host: app.local' -H'Authorization: Bearer unknown' -H'Accept: application/json' localhost:8080/message
HTTP/1.1 403 Forbidden
date: Sun, 11 Aug 2019 13:45:29 GMT
server: envoy
content-length: 0

$ curl -i -H'Host: app.local' -H'Authorization: Bearer abc123' -H'Accept: application/json' localhost:8080/message
HTTP/1.1 200 OK
server: envoy
date: Sun, 11 Aug 2019 13:45:02 GMT
content-type: text/plain
x-envoy-upstream-service-time: 1
transfer-encoding: chunked

CLIENT VALUES:
client_address=172.20.0.5
command=GET
real path=/message
query=nil
request_version=1.1
request_uri=http://app.local:8080/message

SERVER VALUES:
server_version=nginx: 1.10.0 - lua: 10001

HEADERS RECEIVED:
accept=application/json
authorization=Bearer abc123
content-length=0
host=app.local
user-agent=curl/7.54.0
x-envoy-expected-rq-timeout-ms=3000
x-envoy-internal=true
x-envoy-original-path=/message
x-forwarded-for=172.20.0.1
x-forwarded-proto=http
x-request-id=dae0ac12-a4c7-4e23-a1bd-85275530d737
x-subject=user-subject
x-user=a-user
BODY:
-no body in request-

```

## Useful links

### Ambassador

* [Deploying with docker-compose](https://www.getambassador.io/user-guide/docker-compose#create-the-initial-configuration)
* [Global configuration](https://www.getambassador.io/reference/core/ambassador)
* [Mapping reference](https://www.getambassador.io/reference/mappings)
* [Authentication Service config](https://www.getambassador.io/reference/services/auth-service/)
* [Diagnostics endpoint](http://localhost:8080/ambassador/v0/diag/)

### Ory Oathkeeper

* [Decision API](https://www.ory.sh/docs/oathkeeper/#access-control-decision-api)
* [Authenticator reference](https://www.ory.sh/docs/oathkeeper/pipeline/authn)
* [Authorizers reference](https://www.ory.sh/docs/oathkeeper/pipeline/authz)
* [Mutator reference](https://www.ory.sh/docs/oathkeeper/pipeline/mutator)

### Oauth2

* [Token Introspection](https://www.oauth.com/oauth2-servers/token-introspection-endpoint/)
