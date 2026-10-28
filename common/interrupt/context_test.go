package interrupt

import (
	"context"
	"testing"
)

func TestConnectionContextFlagsCanBeCleared(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = ContextWithIsExternalConnection(ctx, true)
	ctx = ContextWithIsResourceDownload(ctx, true)
	if !IsExternalConnectionFromContext(ctx) || !IsResourceDownloadFromContext(ctx) {
		t.Fatal("connection flags were not set")
	}
	ctx = ContextWithIsExternalConnection(ctx, false)
	ctx = ContextWithIsResourceDownload(ctx, false)
	if IsExternalConnectionFromContext(ctx) || IsResourceDownloadFromContext(ctx) {
		t.Fatal("false must override an inherited true flag")
	}
	cancel()
	if ctx.Err() != context.Canceled {
		t.Fatal("connection context must preserve parent cancellation")
	}
}
