package group

import (
	"testing"

	"github.com/sagernet/sing-box/option"

	"github.com/stretchr/testify/require"
)

func TestGroupRejectsDuplicateOutboundTags(t *testing.T) {
	_, err := NewSelector(t.Context(), nil, nil, "selector", option.SelectorOutboundOptions{GroupCommonOption: option.GroupCommonOption{Outbounds: []string{"same", "same"}}})
	require.EqualError(t, err, "duplicate tag: same")
	_, err = NewLoadBalance(t.Context(), nil, nil, "loadbalance", option.LoadBalanceOutboundOptions{GroupCommonOption: option.GroupCommonOption{Outbounds: []string{"same", "same"}}})
	require.EqualError(t, err, "duplicate tag: same")
}
