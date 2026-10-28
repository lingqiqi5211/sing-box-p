package libbox

import (
	"errors"
	"os"
	"testing"

	"github.com/sagernet/sing-box/adapter"
)

type testAutoRedirectPlatform struct {
	enabled bool
	inet6   bool
}

func (p *testAutoRedirectPlatform) UsePlatformAutoRedirect() bool {
	return p.enabled
}

func (p *testAutoRedirectPlatform) CreateAutoRedirectListener(inet6 bool) (int32, error) {
	p.inet6 = inet6
	return 42, nil
}

func (p *testAutoRedirectPlatform) CreateAutoRedirect([]byte, AutoRedirectHandler) (AutoRedirectSession, error) {
	return nil, os.ErrPermission
}

func TestOptionalPlatformAutoRedirect(t *testing.T) {
	wrapper := &platformInterfaceWrapper{}
	server := &CommandServer{platformWrapper: wrapper}
	if wrapper.UsePlatformAutoRedirect() {
		t.Fatal("client without root extension enabled auto-redirect")
	}
	if _, err := wrapper.CreateAutoRedirectListener(false); !errors.Is(err, os.ErrInvalid) {
		t.Fatalf("missing listener extension: %v", err)
	}
	if _, err := wrapper.CreateAutoRedirect(adapter.AutoRedirectOptions{}); !errors.Is(err, os.ErrInvalid) {
		t.Fatalf("missing redirect extension: %v", err)
	}

	platform := &testAutoRedirectPlatform{enabled: true}
	server.SetPlatformAutoRedirect(platform)
	if !wrapper.UsePlatformAutoRedirect() {
		t.Fatal("registered root extension not enabled")
	}
	fd, err := wrapper.CreateAutoRedirectListener(true)
	if err != nil || fd != 42 || !platform.inet6 {
		t.Fatalf("listener extension not forwarded: fd=%d inet6=%v err=%v", fd, platform.inet6, err)
	}
	platform.enabled = false
	if wrapper.UsePlatformAutoRedirect() {
		t.Fatal("disabled root extension enabled auto-redirect")
	}
	server.SetPlatformAutoRedirect(nil)
	if wrapper.UsePlatformAutoRedirect() {
		t.Fatal("cleared extension enabled auto-redirect")
	}
}
