# UserService

## API 

### Service interfaces
- package
```go
github.com/teamlint/mons/sample2/services/user/api
```

- interface
```go
type UserServer interface {
	Find(context.Context, *FindUserRequest) (*FindUserReply, error)
	Update(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
}
```

### Message types
- FindUserRequest
- FindUserReply
- UpdateUserRequest
- UpdateUserReply

## Client

### HTTP

### gRPC

### NATS
- package
```go
import natsclient "github.com/teamlint/mons/sample2/services/user/client/nats"
```

- example
```go
package main

import (
	"context"
	"log"

    api "github.com/teamlint/mons/sample2/services/user/api"
    natsclient "github.com/teamlint/mons/sample2/services/user/client/nats"
    kitnats "github.com/go-kit/kit/transport/nats"

    "github.com/nats-io/nats.go"
)

func main(){
	var (
		client api.UserServer
		nc     *nats.Conn
		err    error
	)

    nc, err = nats.Connect(nats.DefaultURL)
    if err != nil {
        log.Fatalf("[nats] connect error: %v", err)
    }
    defer nc.Close()

    client, err = natsclient.New(nc, map[string][]kitnats.PublisherOption{})
    if err != nil {
        log.Fatalf("[nats] client instance error: %v", err)
    }
    
    // UserServer.Find
    ctx := context.Background()
    findReq := api.FindUserRequest{...}
    findResult, err := client.Find(ctx, &findReq)
    if err != nil {
        log.Fatalf("[Find] err: %v\n", err)
        return
    }
    log.Printf("[Find] result: %+v\n", *findResult)
}
```

## Build Information
- BuildDate     2019-09-27 18:43:17.366587 +0800 CST m=+0.007811720      
- BuildHostname venjiang 
- BuildUser     venjiang
- GoPWD         github.com/teamlint/mons/sample2/services/user
- PWD           /Users/venjiang/go/src/github.com/teamlint/mons/sample2/services/user  
- Debug         false  
- DestinationDir.
- File          user.proto  
- RawFilename   README.md.tmpl
- Filename      README.md.tmpl
- TemplateDir   ../../templates
- Service       name:"User" method:<name:"Find" input_type:".api.FindUserRequest" output_type:".api.FindUserReply" > method:<name:"Update" input_type:".api.UpdateUserRequest" output_type:".api.UpdateUserReply" > 
- Enum          []  

