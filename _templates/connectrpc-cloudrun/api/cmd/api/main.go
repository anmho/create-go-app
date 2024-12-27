package main

import (
	_ "embed"
	"errors"
	"net/http"

	"connectrpc.com/grpcreflect"
	"{{ .ModuleName }}/gen/protos/greet/v1/greetv1connect"
	"{{ .ModuleName }}/pkg/server"
	_ "github.com/jackc/pgx/v5/stdlib" // Import the pgx driver for database/sql
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Config struct {
	DbHost         string `env:"DB_HOST"`
	DbUser         string `env:"DB_USER"`
	DbPass         string `env:"DB_PASS"`
	DbName         string `env:"DB_NAME"`
	DbPort         string `env:"DB_PORT"`
	ClerkSecretKey string `env:"CLERK_SECRET_KEY"`
	Port           int    `env:"PORT" envDefault:"8080"`
}


type Stage = string

const (
	Development Stage = "development"
	Production Stage = "production"
)

func main() {
	greeter := &server.GreetServer{}
	mux := http.NewServeMux()

	mux.Handle(greetv1connect.NewGreetServiceHandler(greeter))
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("ok")); err != nil {
			http.Error(w, "unexpected error", http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("hello world")); err != nil {
			http.Error(w, "unexpected error", http.StatusInternalServerError)
		}
	})

	reflector := grpcreflect.NewStaticReflector(
		"protos.greet.v1.GreetService",
		// protoc-gen-connect-go generates package-level constants
		// for these fully-qualified protobuf service names, so you'd more likely
		// reference userv1.UserServiceName and groupv1.GroupServiceName.
	  )

	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	// Many tools still expect the older version of the server reflection API, so
	// most servers should mount both handlers.
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	err := http.ListenAndServe(
		":8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
		// mux,
	)
	if !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}