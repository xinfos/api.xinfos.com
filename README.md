# api.xinfos.com
This is the basic API service of xinfos Mall, which mainly involves the backstage management of API, front-end mall API. The whole API service is developed using Golang, and the Gin framework is selected as the service infrastructure.

### Contents
- [Installation](#installation)
- [Quick start](#quick-start)
- [Directory](#directory)
- [API list](#api-list)

### Installation

1. Download the source code, [Go](https://golang.org/) installed (**version 1.13+ is required**), then you can use the below Go command to run.

```sh
$ git clone https://github.com/xinfos/api.xinfos.com.git
```

### Quick start

```sh
$ go run cmd/main.go
```

### Directory
```sh
├── api                 # OpenAPI JSON schema files, protocol definition files.
│   ├── backend         # backend API
│   ├── v1              # V1 API 
├── apidoc              # API document
├── cmd                 # Main applications for this project.
├── configs             # Configuration file templates or default configs.
│   ├── config.go       # Configuration templates struct
│   ├── config.yaml     # Configuration file
├── driver              # Drivers of various middleware, such as MySQL,Redis
├── internal            # Private application and library code.
│   ├── model           # Data model
│   ├── repository      # Data operation encapsulation
│   │   └── cache       # Data cache
│   ├── service         # Bu
├── pkg                 # Packaging and Continuous Integration.
│   ├── metrics         # metrics
├── tools               # Supporting tools for this project.
├── vendor              # Application dependencies
├── go.mod              # go.mod
├── go.sum              # go.sum
└── README.md           # README
```


### API list

1、API documentation service can be started quickly.

```sh
go run apidoc/main.go
```

2、API 

| API Name                       | Request Method | Content-Type |   AUTH | Description |
| ------------------------------ | :-----------:| :---------------:| :------------:| :---------------|
| /backend/brand/list            | `POST`        | `application/json`|    NO  | Brand list interface |

