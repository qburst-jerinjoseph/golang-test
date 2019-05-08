package app

import (
	"context"
	"lazy-go/internal/data"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

// Server is the top level shohin server application object.
type Server struct {
	*logrus.Logger
	data.Repo
	Name string
}

// Shutdown for gracefull shutdown
func (s *Server) Shutdown(c context.Context, t time.Duration) {
	s.Info("lazygo server shut down gracefully")
}

// Serve is responsible for starting services built using gep
func Serve(s Server, port string, h http.Handler) {
	go func() {
		if !strings.HasPrefix(port, ":") {
			port = ":" + port
		}
		s.Infof("Listening on port %v...\n", port)
		err := http.ListenAndServe(port, h)
		if err != http.ErrServerClosed {
			log.Fatal("failed to start server: ", err)
		}
	}()
	graceful(s, 5*time.Second)
}

// graceful provides the top level app context with a chance to
// release long-lived assets prior to app shutdown
func graceful(s Server, timeout time.Duration) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	s.Shutdown(ctx, timeout)
}
