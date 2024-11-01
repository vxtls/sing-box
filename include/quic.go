//go:build with_quic

package include

import (
	"github.com/sagernet/sing-box/adapter/outbound"
	"github.com/sagernet/sing-box/protocol/hysteria"
	"github.com/sagernet/sing-box/protocol/hysteria2"
	"github.com/sagernet/sing-box/protocol/tuic"
	_ "github.com/sagernet/sing-box/transport/v2rayquic"
	_ "github.com/sagernet/sing-dns/quic"
)

func registerQUICOutbounds(registry *outbound.Registry) {
	hysteria.RegisterOutbound(registry)
	tuic.RegisterOutbound(registry)
	hysteria2.RegisterOutbound(registry)
}
