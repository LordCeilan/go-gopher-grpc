# Gopher gRPC example

In this project we will create a gRPC implementation with CLI application using cobra.
08
The propper order is to create a new project in the folder we will use.

```bash
go mod init github.com/{userName}/{repoName}
```

this will create a new model with the name previosly choosed. 

For this to work first we'll need to have installed Go.

Then install protocol buffer in the next link:

[Protocol buffers](https://github.com/protocolbuffers/protobuf/tags)

in your C:/ folder if you are a Windows peasant as me and don't forget to include it in the PATH

Go plugins need to be also installed so use

## Protoc

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

### Protoc gRPC

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go-grpc
```

Then we should init the project as usual.

```bash
go mod init github.com/{userName}/{repository}
```

Then install grpc-go package

```bash
go get -u google.golang.org/grpc
```

Now to create a CLI application we must use:

```bash
go get -u github.com/spf13@latest
```

if needed search for cobra cli and see it's propper implementation.

Generate a CLI application structure and imports.

```bash
    cobra init --pkg-name github.com/{userName}/{repositoryName}
```

```bash
cobra-cli init 
```

Here we'll use **Viper** so install it: 

```bash
go get github.com/spf13/viper@v1.8.1
```

Add the next services:

```bash
cobra add client
```

```bash
cobra add server
```

In the `go.mod` file should have the following imports: 

```go
module github.com/LordCeilan/go-gopher-grpc

go 1.19

require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/pelletier/go-toml v1.9.3 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.8.1 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	golang.org/x/text v0.3.5 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
```

In the created files must be alocated a `root.go` file and edit it this way: 

```go
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "go-gopher-grpc",
    Short: "gRPC app in Go",
    Long:  `gRPC application written in Go.`,
}
```

## Execute it

To execute the API use:

```bash
go run main.go
```

```bash
go run main.go client
```


```bash
go run main.go server
```

## Proto Files

Lets create a `.proto` file which uses the needed info to interact with our gRPC server

```proto
syntax = "proto3";
package gopher;

option go_package = "github.com/scraly/learning-by-examples/go-gopher-grpc";

// The gopher service definition.
service Gopher {
  // Get Gopher URL
  rpc GetGopher (GopherRequest) returns (GopherReply) {}
}

// The request message containing the user's name.
message GopherRequest {
  string name = 1;
}

// The response message containing the greetings
message GopherReply {
  string message = 1;
}

```

generate it's respective code from pkg/gopher file :

```bash
protoc --proto_path=. *.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative 
```

it will create a `pb.go` and a `gopher_grpc.pb.go` file with the name of the proto file selected. Remember to change the `*` symbol with the name if only want to create a single pair of `_grpc.pb.go` & `pb.go`

In our case`gopher.pb.go` & `gopher_grpc.pb.go` will contain generated code that we'll import in our `server.go` file in order to register our gRPC server to `Gopher` service

the dependecies implemented in `.pg.go` are the ones used in the `grpc.pb.go` file and then in the `.go` file.

## Server

The server we we'll need to initialize will use the next imports

```go
    import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	pb "github.com/LordCeilan/go-gopher-grpc/pkg/gopher"
	"google.golang.org/grpc"
)
```

`pb` will be the alias we'll use to access our generated `.proto` code

constants will be initialized:

```go
const (
    port = ":9000"
    KuteGoAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"
)
```

