// Package nethttp manages application http interface.
package nethttp

import (
	"net/http"

	"github.com/acme-corp-tech/brick"
	"github.com/acme-corp-tech/auth/internal/infra/nethttp/ui"
	"github.com/acme-corp-tech/auth/internal/infra/service"
	"github.com/acme-corp-tech/auth/internal/usecase"
	"github.com/swaggest/rest/nethttp"
)

// NewRouter creates an instance of router filled with handlers and docs.
func NewRouter(deps *service.Locator) http.Handler {
	r := brick.NewBaseRouter(deps.BaseLocator)

	r.Method(http.MethodGet, "/hello", nethttp.NewHandler(usecase.HelloWorld(deps)))

	r.Method(http.MethodGet, "/", ui.Index())
	r.Mount("/static/", http.StripPrefix("/static", ui.Static))

	return r
}
