package main

import (
	"log"
	"net/http"

	"github.com/acme-corp-tech/brick"
	"github.com/acme-corp-tech/service-starter-kit/internal/infra"
	"github.com/acme-corp-tech/service-starter-kit/internal/infra/nethttp"
	"github.com/acme-corp-tech/service-starter-kit/internal/infra/service"
)

func main() {
	var cfg service.Config

	brick.Start("", &cfg, func(docsMode bool) (*brick.BaseLocator, http.Handler) {
		// Initialize application resources.
		sl, err := infra.NewServiceLocator(cfg)
		if err != nil {
			log.Fatalf("failed to init service: %v", err)
		}

		return sl.BaseLocator, nethttp.NewRouter(sl)
	})
}