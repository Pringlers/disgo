package httpserver

import (
	"context"
	"encoding/hex"
	"net/http"
)

var _ Server = (*serverImpl)(nil)

// New creates a new Server with the given publicKey eventHandlerFunc and ConfigOpt(s)
func New(publicKey string, eventHandlerFunc EventHandlerFunc, opts ...ConfigOpt) Server {
	config := DefaultConfig()
	config.Apply(opts)

	hexDecodedKey, err := hex.DecodeString(publicKey)
	if err != nil {
		config.Logger.Errorf("error while decoding hex string: %s", err)
	}

	return &serverImpl{
		config:           *config,
		publicKey:        hexDecodedKey,
		eventHandlerFunc: eventHandlerFunc,
	}
}

type serverImpl struct {
	config           Config
	publicKey        PublicKey
	eventHandlerFunc EventHandlerFunc
}

func (s *serverImpl) Start() {
	s.config.ServeMux.Handle(s.config.URL, HandleInteraction(s.publicKey, s.config.Logger, s.eventHandlerFunc))
	s.config.HTTPServer.Addr = s.config.Address
	s.config.HTTPServer.Handler = s.config.ServeMux

	go func() {
		var err error
		if s.config.CertFile != "" && s.config.KeyFile != "" {
			err = s.config.HTTPServer.ListenAndServeTLS(s.config.CertFile, s.config.KeyFile)
		} else {
			err = s.config.HTTPServer.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			s.config.Logger.Error("error while running http server: ", err)
		}
	}()
}

func (s *serverImpl) Close(ctx context.Context) {
	_ = s.config.HTTPServer.Shutdown(ctx)
}
