# gripkit

gRPC go + web server toolkit. not a replacement but a wrapper to initialize grpc + grpc-web in a single shot.

```
$ go get -u github.com/roleypoly/gripkit
```

## API

Basic usage is 
```go
gk := gripkit.Create(
  // gripkit.WithHTTPOptions(gripkit.HTTPOptions{ Addr: "", TLSCertFile: "", TLSKeyFile: "" }), // HTTP(S) server settings.
  // gripkit.WithGrpcWeb( /* add grpcweb.Options... */ ), // turns on gRPC-Web, options optional.
  // gripkit.WithOptions( /* add grpc.Options... */ ), // adds gRPC options
)

proto.RegisterGreeterService(gk.Server, greeterService) // setup gRPC services based on gk.Server.

err := gk.Serve() // starts HTTP(S) server. If TLSCertFile/TLSKeyFile isn't set, will use HTTP, otherwise HTTPS.
if err != nil {
  panic(err)
}
```
