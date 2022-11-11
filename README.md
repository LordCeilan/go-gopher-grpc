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
    KuteGoAPIURL = "127.0.0.1:8080"
)
```

In our case we'll define the KuteGoAPIURL this way due to the local app is running in the port given.

Then create the structs of our server in our Gopher data.

```go
type Server struct {
 pb.UnimplementedGopherServer
}
//Which is imported from gopher.pb.go

type Gopher struct {
 URL string `json: "url"`
}
```

Now the serverCmd is improved which inits a gRPC server, service to RPC service and start our server.

```go
var serverCmd = &cobra.Command{
    Use: "server",
    Short: "Starts the Schema gRPC server", 

    Run: func(cmd *cobra.Command, args[]string) {
        lis, err := net.Listen("tcp", port)

        if err != nil {
            log.Fatalf("failed to listen %v", err)
        }

        grpcServer := grpc.NewServer()

        pb.RegisterGopherServer(grpcServer, &Server{})

        log.Printf("Grpc server listening on %v")

        if err := grpcServer.Serve(lis); err != nil {
            log.Fatalf("failed to serve: %v", err)
        }
    },
}
```

Finally we'll implement `GetGopher` method.

This will query into an external API located in this repo, this has to do with the `KuteGoAPIURL` value given before.

- [KuteGo API](https://github.com/gaelleacas/kutego-api)

Follow the instructions in the README.md to run this project.

## Implementation

Now jump into the `server.go` file and add the next code:

```go
func (s *Server) GetGopher(ctx context.Context, req *pb.GopherRequest) (*pb.GopherReply, error) {

    res := &pb.GopherReply{}

        if req == nil {
            fmt.Println("request must not be nil")
            return res, xerrors.Errorf("request must not be nil")
        }

        if req.Name == "" {
            fmt.Println("name must not be empty in the request")
            return res, xerrors.Errorf("name must not be empty in the request")
        }

        log.Printf("Received: %v", req.GetName())

        parsedUrl, err := url.Parse("http://" + KuteGoAPIURL + "/gophers?name=" + req.GetName())

        if err != nil {
            log.Fatalf("url incorrect %v or not authorized", parsedUrl)
        }

        // fmt.Printf("%q", parsedUrl)
        // fmt.Printf("%s", parsedUrl.String())

        response, err := http.Get(parsedUrl.String())

        if err != nil {
            log.Fatalf("failed to call KuteGoAPI: %v", err)
        }

        defer response.Body.Close()

        if response.Status == "200 OK" {
            body, err := ioutil.ReadAll(response.Body)

            if err != nil {
                log.Fatalf("failed to read response body: %v", err)
            }

            var data []Gopher
            err = json.Unmarshal(body, &data)
            if err != nil {
                log.Fatalf("failed to unmarshal JSON: %v", err)
            }

            var gophers strings.Builder
            for _, gopher := range data {
                gophers.WriteString(gopher.URL + "\n")
            }

            res.Message = gophers.String()

        } else {
            log.Fatal("Can't get the Gopher :c")
        }

    return res, nil
}
```

Lastly install if needed the external dependencies.

```bash
go get google.golang.org/grpc
go get golang.org/x/xerrors
```

## gRPC Client

For the client let's edit the `client.go` file

We'll initialize the package caled cmdm and all dependencies we'll import.

```go
package cmd

import (
    "context"
    "log"
    "os"
    "time"
    
    "google.golang.org/grpc"

    pb "github.com/LordCeilan/go-gopher-grpc/pkg/gopher"

    "github.com/spf13/cobra"
)

```

const (
    address = "localhost:9000"
    defaultName = "deez-nuts"
)

Improving the clientCmd run function:

- Initialize a gRPC client
- Connect to gRPC server
- Call the GetGopher function with the Gopher's name from console
- return "URL:" + the message returned by the gRPC call

```go
var clientCmd = &cobra.Command{

    Use: "client",
    Short: "Query the gRPC server",

    Run: func(cmd *cobra,Command, args[]string) {
        var conn *grpc.ClientConn
        conn, err := grpc.Dial(address, grpc.WithInsecure())
        if err != nil {
            log.Fatalf("did not connect: %s", err)
        }

        client := pb.NewGopherClient(conn)

        var name string

        if len(os.Args) > 2 {
            name = os.Args[2]
        }

        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()

        r, err := client.GetGopher(ctx, &pb.GopherRequest{Name: name})
        if err != nil {
            log.Fatalf("could not greet: %v", err)
        }

        log.Printf("URL: %s", r.GetMessage())
    },
}
```

## Test it

To start the server use

```bash
go run main.go server
```

And open another terminal and call the client command with the "gandalf" parameter:

```bash
go run main.go client gandalf
```

The application properly answers the url and the URL

`2022/11/11 15:03:04 URL: https://raw.githubusercontent.com/scraly/gophers/main/gandalf.png`

## Build it

For building process we'll use Taskfile, for it's installation we'll use `Chocolatey` and for it's installation just follow the instructions bellow:

1 [Chocolatey](https://chocolatey.org/install)
2 In an elevated `Powershell` prompt use:

```bash
choco install go-task
```

For building, create a `Taskfile.yml` this file will work step by step the proper ones to generate compilated files for it's usage.

```yml
version: "3"

task:
    build:
        desc: Build the app
        cmds: 
        - GOFLAGS =-mod=mod go build -o bin/gopher-grpc main.go

    run:
        desc: Run the app
        cmds: 
        - GOFLAGS =-mod=mod go run main.go

    generate:
        desc: Generate Go code from protobuf
        cmds:
        - protoc --proto_path=. pkg/gopher/*.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative

    test:
        desc: Execute Unit Test
        cmds: 
        - gotestsum --junitfile test-result/unit-test.xml -- -short -race -cover -coverprofile test-results/cover.out ./...
```

Then if you run:

```bash
task build
```

and testing it with our fresh executable binary:

```bash
./bin/gopher-grpc server
```

and in the another tab of the terminal the client:

```bash
./bin/gopher-grpc client yoda-gopher
```

this will execute the same way as we were running the but in a prod environment is not useful u know

```bash
go run main.go client
```

## TESTINGS

Dude, this is my less favorite part, but it's a must!
So in order accomplish it we'll use some integrated tools in golang.

```bash
go test
```

obviusly it will show no testing were made but we'll discover a useful tool gotestsum

## Gotestsum

We'll implement testing with this tool. It runs with `go test` and improves the display of results. Making them more human-readable, practical report with possible output directly in JUnit format.

Install it by using:

```bash
go get gotest.tools/gotestsum
```
