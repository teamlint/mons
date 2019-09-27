package internal

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
    
	"github.com/fatih/structs"
    api "github.com/teamlint/mons/sample2/services/user/api"
	grpcclient "github.com/teamlint/mons/sample2/services/user/client/grpc"
	httpclient "github.com/teamlint/mons/sample2/services/user/client/http"
    natsclient "github.com/teamlint/mons/sample2/services/user/client/nats"
	"google.golang.org/grpc"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	kithttp "github.com/go-kit/kit/transport/http"
	kitnats "github.com/go-kit/kit/transport/nats"

	"github.com/nats-io/nats.go"
)

var (
	fs       = flag.NewFlagSet("User-Client", flag.ExitOnError)
	trans    = fs.String("ts", "nats", "client transport protocol [http,nats,grpc]")
	natsAddr = fs.String("nats-addr", nats.DefaultURL, "NATS listen address")
	httpAddr = fs.String("http-addr", ":8081", "HTTP listen address")
	grpcAddr = fs.String("grpc-addr", ":8082", "GRPC listen address")
)

func Run() {
	// args
	fs.Parse(os.Args[1:])
	// demo
	var (
		client api.UserServer
		nc     *nats.Conn
		err    error
	)
	start := time.Now()
	// client
	switch *trans {
	case "http":
		log.Println("== [http client] ===")
		client, err = httpclient.New(*httpAddr, map[string][]kithttp.ClientOption{})
		if err != nil {
			log.Fatalf("[http] client instance error: %v", err)
		}
		elapse(start, "[%s] connect time", *trans)
	case "grpc":
		log.Println("== [grpc client] ===")
		conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("[grpc] connect error: %v", err)
		}
		defer conn.Close()
		client, err = grpcclient.New(conn, map[string][]kitgrpc.ClientOption{})
		if err != nil {
			log.Fatalf("[http] client instance error: %v", err)
		}
		elapse(start, "[%s] connect time", *trans)
	default:
		log.Println("=== [nats client] ===")
		nc, err = nats.Connect(*natsAddr)
		if err != nil {
			log.Fatalf("[nats] connect error: %v", err)
		}
		defer nc.Close()
		client, err = natsclient.New(nc, map[string][]kitnats.PublisherOption{})
		if err != nil {
			log.Fatalf("[nats] client instance error: %v", err)
		}
		elapse(start, "[%s] connect time", *trans)
	}

	begin := time.Now()
	ctx := context.Background()
    
    // User.Find
	start = time.Now()
	findReq := api.FindUserRequest{
		Id: "123",
	}
	findResult, err := client.Find(ctx, &findReq)
	if err != nil {
		log.Printf("[Find] err: %v\n", err)
		return
	}
	findErr, ok := structs.New(findResult).FieldOk("Error")
    // if has Error field, you can direct use struct.Error
	if ok && findErr.Value() != nil {
		log.Printf("[Find] business err: %v\n", findErr.Value())
	}
	elapse(start, "[client.%s] Find result: %+v", *trans, *findResult)
    
    // User.Update
	start = time.Now()
	updateReq := api.UpdateUserRequest{
		Id: "123",
	}
	updateResult, err := client.Update(ctx, &updateReq)
	if err != nil {
		log.Printf("[Update] err: %v\n", err)
		return
	}
	updateErr, ok := structs.New(updateResult).FieldOk("Error")
    // if has Error field, you can direct use struct.Error
	if ok && updateErr.Value() != nil {
		log.Printf("[Update] business err: %v\n", updateErr.Value())
	}
	elapse(start, "[client.%s] Update result: %+v", *trans, *updateResult)
    
	// load
	elapse(begin, "[client.%s] load time", *trans)
}

func elapse(start time.Time, format string, vals ...interface{}) {
	us := time.Since(start).Nanoseconds() / 1000
	ms := float64(us) / float64(1000)
	load := strconv.FormatFloat(ms, 'f', 3, 64)
	log.Printf("[%s ms] %v\n", load, fmt.Sprintf(format, vals...))
}
