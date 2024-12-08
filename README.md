[![CircleCI](https://circleci.com/gh/alexfalkowski/idpd.svg?style=svg)](https://circleci.com/gh/alexfalkowski/idpd)
[![codecov](https://codecov.io/gh/alexfalkowski/idpd/graph/badge.svg?token=S9SPVVYQAY)](https://codecov.io/gh/alexfalkowski/idpd)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexfalkowski/idpd)](https://goreportcard.com/report/github.com/alexfalkowski/idpd)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/idpd.svg)](https://pkg.go.dev/github.com/alexfalkowski/idpd)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# Internal Developer Platform

An [Internal Developer Platform](https://internaldeveloperplatform.org/what-is-an-internal-developer-platform/) (IDP) is built by a platform team to build golden paths and enable developer self-service.

### Why a service?

Internal developer portals serve as the interface through which developers can discover and access internal developer platform capabilities.

## Server

The server is composed of [RESTful](https://aws.amazon.com/what-is/restful-api/) API.

Each of the endpoints is documented using the following [style guide](https://docs.gitlab.com/ee/development/documentation/restful_api_styleguide.html).

### Create a pipeline

This endpoint takes pipeline definition and persists it.

```plaintext
POST /pipelines
```
Example request:

```shell
curl --header "Authorization: Bearer <token>"  --header "Content-Type: application/json" --request POST --data '{ "pipeline": {"name":"test","jobs":[{"name":"test","steps":["test","test2"]}]}' --url "http://localhost:11000/pipelines"
```

Example response:

```json
{
   "meta":{
      "ipAddr":"127.0.0.1",
      "ipAddrKind":"remote",
      "requestId":"ae429023-e070-433b-846c-47cf8a209b42",
      "traceId":"893bf1648d88427b14a9ebd8f8f73437",
      "userAgent":"IDP-ruby-client/1.0 HTTP/1.0"
   },
   "pipeline":{
      "id":1,
      "name":"test",
      "jobs":[
         {
            "name":"test",
            "steps":[
               "test",
               "test2"
            ]
         }
      ]
   }
}
```

### Get a pipeline

This endpoint gets a pipeline by an ID.

```plaintext
GET /pipelines/{id}
```
Example request:

```shell
curl --header "Authorization: Bearer <token>"  --header "Content-Type: application/json" --request GET --url "http://localhost:11000/pipelines/1"
```

Example response:

```json
{
   "meta":{
      "ipAddr":"127.0.0.1",
      "ipAddrKind":"remote",
      "requestId":"ae429023-e070-433b-846c-47cf8a209b42",
      "traceId":"893bf1648d88427b14a9ebd8f8f73437",
      "userAgent":"IDP-ruby-client/1.0 HTTP/1.0"
   },
   "pipeline":{
      "id":1,
      "name":"test",
      "jobs":[
         {
            "name":"test",
            "steps":[
               "test",
               "test2"
            ]
         }
      ]
   }
}
```

### Update a Pipeline

This endpoint updates a pipeline by an ID.

```plaintext
PUT /pipelines/{id}
```
Example request:

```shell
curl --header "Authorization: Bearer <token>"  --header "Content-Type: application/json" --request PUT --data '{ "pipeline": {"name":"test","jobs":[{"name":"test","steps":["test","test2"]}]}' --url "http://localhost:11000/pipelines/1"

Example response:

```json
{
   "meta":{
      "ipAddr":"127.0.0.1",
      "ipAddrKind":"remote",
      "requestId":"ae429023-e070-433b-846c-47cf8a209b42",
      "traceId":"893bf1648d88427b14a9ebd8f8f73437",
      "userAgent":"IDP-ruby-client/1.0 HTTP/1.0"
   },
   "pipeline":{
      "id":1,
      "name":"test",
      "jobs":[
         {
            "name":"test",
            "steps":[
               "test",
               "test2"
            ]
         }
      ]
   }
}
```

### Delete a Pipeline

This endpoint deletes a pipeline by an ID.

```plaintext
DELETE /pipelines/{id}
```
Example request:

```shell
curl --header "Authorization: Bearer <token>"  --header "Content-Type: application/json" --request DELETE --url "http://localhost:11000/pipelines/1"

Example response:

```json
{
   "meta":{
      "ipAddr":"127.0.0.1",
      "ipAddrKind":"remote",
      "requestId":"ae429023-e070-433b-846c-47cf8a209b42",
      "traceId":"893bf1648d88427b14a9ebd8f8f73437",
      "userAgent":"IDP-ruby-client/1.0 HTTP/1.0"
   },
   "pipeline":{
      "id":1,
      "name":"test",
      "jobs":[
         {
            "name":"test",
            "steps":[
               "test",
               "test2"
            ]
         }
      ]
   }
}
```

### Trigger a Pipeline

This endpoint triggers a pipeline by an ID.

```plaintext
POST /pipelines/{id}/triggers
```
Example request:

```shell
curl --header "Authorization: Bearer <token>"  --header "Content-Type: application/json" --request POST --url "http://localhost:11000/pipelines/1/triger"

Example response:

```json
{
   "meta":{
      "ipAddr":"127.0.0.1",
      "ipAddrKind":"remote",
      "requestId":"ae429023-e070-433b-846c-47cf8a209b42",
      "traceId":"893bf1648d88427b14a9ebd8f8f73437",
      "userAgent":"IDP-ruby-client/1.0 HTTP/1.0"
   },
   "pipeline":{
      "id":1,
      "name":"test",
      "jobs":[
         {
            "name":"test",
            "steps":[
               "output of command",
               "output of command"
            ]
         }
      ]
   }
}
```

## Health

The system defines a way to monitor all of it's dependencies.

To configure we just need the have the following configuration:

```yaml
health:
  duration: 1s (how often to check)
  timeout: 1s (when we should timeout the check)
```

## Design

The design is heavily influenced by [go-service](https://github.com/alexfalkowski/go-service), using the rest package.

The [api](api) package calls the [pipeline](pipeline) package.

### API

This package handles encoding and decoding of the API and making sure the correct status codes are set.

### Pipeline

This package contains all the domain knowledge that deals with pipelines, from validation to execution.

## Other Systems

Check out the available [tooling](https://internaldeveloperplatform.org/platform-tooling/).

## Development

If you would like to contribute, here is how you can get started.

### Structure

The project follows the structure in [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

### Dependencies

Please make sure that you have the following installed:
- [Ruby](.ruby-version)
- [Golang](go.mod)

### Style

This project favours the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

### Setup

Check out [CI](.circleci/config.yml).

### Changes

To see what has changed, please have a look at `CHANGELOG.md`
