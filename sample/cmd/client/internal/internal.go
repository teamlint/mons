package internal

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/teamlint/mons/sample/application/command"
	"github.com/teamlint/mons/sample/application/query"
	"github.com/teamlint/mons/sample/application/service"
	grpcclient "github.com/teamlint/mons/sample/client/grpc"
	httpclient "github.com/teamlint/mons/sample/client/http"
	natsclient "github.com/teamlint/mons/sample/client/nats"
	"google.golang.org/grpc"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	kithttp "github.com/go-kit/kit/transport/http"
	kitnats "github.com/go-kit/kit/transport/nats"

	"github.com/nats-io/nats.go"
)

var (
	fs       = flag.NewFlagSet("client", flag.ExitOnError)
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
		client service.UserService
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
		elapse(start, "[%s] connec time", *trans)
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
		elapse(start, "[%s] connec time", *trans)
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
		elapse(start, "[%s] connec time", *trans)
	}

	begin := time.Now()
	// find user
	start = time.Now()
	ctx := context.Background()
	query := query.UserQuery{
		ID: "23f9db62-a145-11e9-ae8f-132379f358df",
	}
	user, err := client.Find(ctx, &query)
	if err != nil {
		log.Println(err)
		return
	}
	elapse(start, "[client.%s] find user result: %+v", *trans, *user)
	// update user
	start = time.Now()
	origin := strings.Split(user.Username, "-")[0]
	cmd := command.UpdateUserCommandFrom(user)
	cmd.Username = fmt.Sprintf("%v-%v", origin, time.Now().UnixNano()/1000)
	log.Printf("[cmd] UpdateUserCommand: %+v\n", *cmd)
	err = client.Update(ctx, cmd)
	elapse(start, "[client.%s] update user result: %v", *trans, err == nil)
	// load
	elapse(begin, "[client.%s] load time", *trans)
}

func elapse(start time.Time, format string, vals ...interface{}) {
	us := time.Since(start).Nanoseconds() / 1000
	ms := float64(us) / float64(1000)
	load := strconv.FormatFloat(ms, 'f', 3, 64)
	log.Printf("[%s ms] %v\n", load, fmt.Sprintf(format, vals...))
}
