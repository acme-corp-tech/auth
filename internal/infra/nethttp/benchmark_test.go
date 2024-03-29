package nethttp_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/acme-corp-tech/brick/runtime"
	"github.com/acme-corp-tech/auth/internal/domain/greeting"
	"github.com/acme-corp-tech/auth/internal/infra"
	"github.com/acme-corp-tech/auth/internal/infra/nethttp"
	"github.com/acme-corp-tech/auth/internal/infra/service"
	"github.com/bool64/httptestbench"
	"github.com/stretchr/testify/require"
)

func Benchmark_hello(b *testing.B) {
	log.SetOutput(ioutil.Discard)

	cfg := service.Config{}
	cfg.Log.Output = ioutil.Discard
	cfg.ShutdownTimeout = time.Second
	l, err := infra.NewServiceLocator(cfg)
	require.NoError(b, err)

	l.GreetingMakerProvider = &greeting.SimpleMaker{}

	r := nethttp.NewRouter(l)

	httptestbench.ServeHTTP(b, 50, r,
		func(i int) *http.Request {
			req, err := http.NewRequest(http.MethodGet, "/hello?name=Jack&locale=en-US", nil)
			if err != nil {
				b.Fatal(err)
			}

			return req
		},
		func(i int, resp *httptest.ResponseRecorder) bool {
			return resp.Code == http.StatusOK
		},
	)

	b.StopTimer()
	b.ReportMetric(float64(runtime.StableHeapInUse())/float64(1024*1024), "MB/inuse")
	l.Shutdown()
	require.NoError(b, <-l.Wait())
}
