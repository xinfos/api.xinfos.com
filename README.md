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
| ── app


### API list

| API Name                       | Request Method | Content-Type |   AUTH | Description |
| ------------------------------ | :-----------:| :---------------:| :------------:| :---------------|
| /backend/brand/list            | `POST`        | `application/json`|    NO  | Brand list interface |

