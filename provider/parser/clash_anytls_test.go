package parser

import (
	"context"
	"testing"

	"github.com/sagernet/sing-box/option"
	"github.com/stretchr/testify/require"
)

func TestParseClashAnyTLSDisableReuse(t *testing.T) {
	outbounds, endpoints, err := ParseClashSubscription(context.Background(), `
proxies:
  - name: anytls-out
    type: anytls
    server: 127.0.0.1
    port: 443
    password: password
    disable-reuse: true
`)
	require.NoError(t, err)
	require.Empty(t, endpoints)
	require.Len(t, outbounds, 1)

	anyTLSOptions, ok := outbounds[0].Options.(*option.AnyTLSOutboundOptions)
	require.True(t, ok)
	require.True(t, anyTLSOptions.DisableReuse)
}
