// Code generated by zanzibar
// @generated

package googlenow

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"github.com/uber-go/zap"
	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
)

// HandleCheckCredentialsRequest handles "/googlenow/check-credentials".
func HandleCheckCredentialsRequest(
	ctx context.Context,
	req *zanzibar.IncomingHTTPRequest,
	res *zanzibar.OutgoingHTTPResponse,
	clients *clients.Clients,
) {
	// Handle request headers.
	h := http.Header{}
	for _, header := range []string{"x-uuid", "x-token"} {
		h.Set(header, req.Header.Get(header))
	}

	// Handle request body.
	clientResp, err := clients.GoogleNow.CheckCredentials(ctx, h)
	if err != nil {
		req.Logger.Error("Could not make client request",
			zap.String("error", err.Error()),
		)
		res.SendError(500, errors.Wrap(err, "could not make client request:"))
		return
	}

	defer func() {
		if cerr := clientResp.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	// Handle client respnse.
	if !res.IsOKResponse(clientResp.StatusCode, []int{200, 202}) {
		req.Logger.Warn("Unknown response status code",
			zap.Int("status code", clientResp.StatusCode),
		)
	}
	res.WriteJSONBytes(clientResp.StatusCode, nil)
}
