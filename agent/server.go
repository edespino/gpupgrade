package agent

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"

	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/utils/daemon"
	"github.com/greenplum-db/gpupgrade/utils/log"

	"github.com/greenplum-db/gp-common-go-libs/cluster"
	"github.com/greenplum-db/gp-common-go-libs/gplog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	executor cluster.Executor
	conf     Config

	mu      sync.Mutex
	server  *grpc.Server
	lis     net.Listener
	stopped chan struct{}
	daemon  bool
}

type Config struct {
	Port     int
	StateDir string
}

func NewServer(executor cluster.Executor, conf Config) *Server {
	return &Server{
		executor: executor,
		conf:     conf,
		stopped:  make(chan struct{}, 1),
	}
}

// MakeDaemon tells the Server to disconnect its stdout/stderr streams after
// successfully starting up.
func (a *Server) MakeDaemon() {
	a.daemon = true
}

func (a *Server) Start() {
	createIfNotExists(a.conf.StateDir)
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(a.conf.Port))
	if err != nil {
		gplog.Fatal(err, "failed to listen")
	}

	// Set up an interceptor function to log any panics we get from request
	// handlers.
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer log.WritePanics()
		return handler(ctx, req)
	}
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor))

	a.mu.Lock()
	a.server = server
	a.lis = lis
	a.mu.Unlock()

	idl.RegisterAgentServer(server, a)
	reflection.Register(server)

	if a.daemon {
		// Send an identifier string back to the hub, and log it locally for
		// easier debugging.
		info := fmt.Sprintf("Agent started on port %d (pid %d)", a.conf.Port, os.Getpid())

		fmt.Println(info)
		daemon.Daemonize()
		gplog.Info(info)
	}

	err = server.Serve(lis)
	if err != nil {
		gplog.Fatal(err, "failed to serve: %s", err)
	}

	a.stopped <- struct{}{}
}

func (a *Server) StopAgent(ctx context.Context, in *idl.StopAgentRequest) (*idl.StopAgentReply, error) {
	a.Stop()
	return &idl.StopAgentReply{}, nil
}

func (a *Server) Stop() {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.server != nil {
		a.server.Stop()
		<-a.stopped
	}
}

func createIfNotExists(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0777)
	}
}