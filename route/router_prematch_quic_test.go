package route

import (
	"context"
	"testing"
	"time"

	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/log"
	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"
)

// newPreMatchQUICRouter builds a minimal router for pre-match tests. The QUIC
// sniff cache belongs to the sniffer and is intentionally absent here.
func newPreMatchQUICRouter(t *testing.T, _ time.Duration) (*Router, adapter.InboundContext) {
	t.Helper()
	router := &Router{
		ctx:          context.Background(),
		logger:       log.NewNOPFactory().NewLogger("router"),
		dnsTransport: &preMatchQUICDNSManager{},
	}
	metadata := adapter.InboundContext{
		Source:      M.ParseSocksaddr("127.0.0.1:10000"),
		Destination: M.ParseSocksaddr("1.1.1.1:443"),
		Network:     N.NetworkUDP,
		Domain:      "reverse.example",
	}
	return router, metadata
}

type preMatchQUICDNSManager struct{ adapter.DNSTransportManager }

func (*preMatchQUICDNSManager) FakeIP() adapter.FakeIPTransport { return nil }

type preMatchQUICRule struct {
	adapter.Rule
	action   adapter.RuleAction
	disabled bool
}

func (r *preMatchQUICRule) Action() adapter.RuleAction { return r.action }

func (r *preMatchQUICRule) Disabled() bool { return r.disabled }

func (*preMatchQUICRule) String() string { return "" }

func (*preMatchQUICRule) Match(*adapter.InboundContext) bool { return true }
