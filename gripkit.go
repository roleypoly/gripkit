// Package gripkit provides wrappers and helpers for
package gripkit // import "github.com/roleypoly/gripkit"

import (
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

type Gripkit struct {
	Server        *grpc.Server
	httpHandler   *func(w http.ResponseWriter, r *http.Request)
	options       *options
	grpcWebServer *grpcweb.WrappedGrpcServer
}

func Create(options ...Option) *Gripkit {
	gkOptions := evaluateOptions(options...)
	grpcServer := grpc.NewServer(gkOptions.grpcOptions...)

	var grpcWrapper *grpcweb.WrappedGrpcServer
	if gkOptions.wrapGrpcWeb {
		grpcWrapper = grpcweb.WrapServer(grpcServer, gkOptions.grpcWebOptions...)
	}

	return &Gripkit{
		Server:        grpcServer,
		grpcWebServer: grpcWrapper,
		options:       gkOptions,
	}
}

func (gk *Gripkit) Serve() error {
	handler := gk.Server.ServeHTTP
	if gk.options.wrapGrpcWeb {
		handler = gk.grpcWebServer.ServeHTTP
	}

	httpHandler := http.HandlerFunc(handler)

	if gk.options.httpOptions.TLSCertPath == "" || gk.options.httpOptions.TLSKeyPath == "" {
		return http.ListenAndServe(
			gk.options.httpOptions.Addr,
			httpHandler,
		)
	}

	return http.ListenAndServeTLS(
		gk.options.httpOptions.Addr,
		gk.options.httpOptions.TLSCertPath,
		gk.options.httpOptions.TLSKeyPath,
		httpHandler,
	)
}
