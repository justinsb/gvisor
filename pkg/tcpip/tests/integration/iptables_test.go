// Copyright 2021 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package integration_test

import (
	"testing"

	"gvisor.dev/gvisor/pkg/tcpip"
	"gvisor.dev/gvisor/pkg/tcpip/buffer"
	"gvisor.dev/gvisor/pkg/tcpip/header"
	"gvisor.dev/gvisor/pkg/tcpip/link/channel"
	"gvisor.dev/gvisor/pkg/tcpip/network/ipv4"
	"gvisor.dev/gvisor/pkg/tcpip/network/ipv6"
	"gvisor.dev/gvisor/pkg/tcpip/stack"
)

// ifNameMatcher is an iptables matcher that matches the input interface name.
type ifNameMatcher struct {
	inputIfName string
}

// Name implements Matcher.Name.
func (*ifNameMatcher) Name() string {
	return "ifNameMatcher"
}

// Match implements Matcher.Match.
func (im *ifNameMatcher) Match(_ stack.Hook, _ *stack.PacketBuffer, inNicName string, _ string) (bool, bool) {
	if im.inputIfName != "" && im.inputIfName == inNicName {
		return true, false
	}
	return false, false
}

func TestIPTablesStatsForInput(t *testing.T) {
	const (
		nicID          = 1
		nicName        = "nic1"
		anotherNicName = "nic2"
		linkAddr       = tcpip.LinkAddress("\x0a\x0b\x0c\x0d\x0e\x0e")
		dstAddrV4      = "\x0a\x00\x00\x02"
		dstAddrV6      = "\x0a\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02"
	)

	tests := []struct {
		name               string
		setup              func(*testing.T, *stack.Stack, bool)
		expectReceived     int
		expectInputDropped int
	}{
		{
			name: "Accept",
			// No setup needed, tables accept everything by default.
			setup:              func(*testing.T, *stack.Stack, bool) {},
			expectReceived:     1,
			expectInputDropped: 0,
		},
		{
			name: "Drop - Filter.InputInterface matches and a Matcher returns true",
			setup: func(t *testing.T, s *stack.Stack, v6 bool) {
				t.Helper()
				ipt := s.IPTables()
				filter := ipt.GetTable(stack.FilterID, v6)
				ruleIdx := filter.BuiltinChains[stack.Input]
				filter.Rules[ruleIdx].Filter = stack.IPHeaderFilter{InputInterface: nicName}
				filter.Rules[ruleIdx].Target = &stack.DropTarget{}
				filter.Rules[ruleIdx].Matchers = []stack.Matcher{&ifNameMatcher{
					inputIfName: nicName,
				}}
				// Make sure the next rule is ACCEPT.
				filter.Rules[ruleIdx+1].Target = &stack.AcceptTarget{}
				if err := ipt.ReplaceTable(stack.FilterID, filter, v6); err != nil {
					t.Fatalf("failed to replace table: %s", err)
				}
			},
			expectReceived:     1,
			expectInputDropped: 1,
		},
		{
			name: "Filter.InputInterface is different",
			setup: func(t *testing.T, s *stack.Stack, v6 bool) {
				t.Helper()
				ipt := s.IPTables()
				filter := ipt.GetTable(stack.FilterID, v6)
				ruleIdx := filter.BuiltinChains[stack.Input]
				filter.Rules[ruleIdx].Filter = stack.IPHeaderFilter{InputInterface: anotherNicName}
				filter.Rules[ruleIdx].Target = &stack.DropTarget{}
				filter.Rules[ruleIdx+1].Target = &stack.AcceptTarget{}
				if err := ipt.ReplaceTable(stack.FilterID, filter, v6); err != nil {
					t.Fatalf("failed to replace table: %s", err)
				}
			},
			expectReceived:     1,
			expectInputDropped: 0,
		},
		{
			name: "Filter.InputInterface matches but invert flag is true",
			setup: func(t *testing.T, s *stack.Stack, v6 bool) {
				t.Helper()
				ipt := s.IPTables()
				filter := ipt.GetTable(stack.FilterID, v6)
				ruleIdx := filter.BuiltinChains[stack.Input]
				filter.Rules[ruleIdx].Filter = stack.IPHeaderFilter{
					InputInterface:       anotherNicName,
					InputInterfaceInvert: true,
				}
				filter.Rules[ruleIdx].Target = &stack.DropTarget{}
				filter.Rules[ruleIdx+1].Target = &stack.AcceptTarget{}
				if err := ipt.ReplaceTable(stack.FilterID, filter, v6); err != nil {
					t.Fatalf("failed to replace table: %s", err)
				}
			},
			expectReceived:     1,
			expectInputDropped: 1,
		},
		{
			name: "Input hook - a Matcher returns false",
			setup: func(t *testing.T, s *stack.Stack, v6 bool) {
				t.Helper()
				ipt := s.IPTables()
				filter := ipt.GetTable(stack.FilterID, v6)
				ruleIdx := filter.BuiltinChains[stack.Input]
				filter.Rules[ruleIdx].Target = &stack.DropTarget{}
				filter.Rules[ruleIdx].Matchers = []stack.Matcher{&ifNameMatcher{anotherNicName}}
				filter.Rules[ruleIdx+1].Target = &stack.AcceptTarget{}
				if err := ipt.ReplaceTable(stack.FilterID, filter, v6); err != nil {
					t.Fatalf("failed to replace table: %s", err)
				}
			},
			expectReceived:     1,
			expectInputDropped: 0,
		},
	}

	genPacket := func(v6 bool) *stack.PacketBuffer {
		const payloadSize = 20

		var pktSize int
		if v6 {
			pktSize = header.IPv6MinimumSize + payloadSize
		} else {
			pktSize = header.IPv4MinimumSize + payloadSize
		}
		hdr := buffer.NewPrependable(pktSize)
		if v6 {
			ip := header.IPv6(hdr.Prepend(pktSize))
			ip.Encode(&header.IPv6Fields{
				PayloadLength:     payloadSize,
				TransportProtocol: 99,
				HopLimit:          255,
				SrcAddr:           "\x0a\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01",
				DstAddr:           dstAddrV6,
			})
		} else {
			ip := header.IPv4(hdr.Prepend(pktSize))
			ip.Encode(&header.IPv4Fields{
				TOS:            0,
				TotalLength:    uint16(pktSize),
				ID:             1,
				Flags:          0,
				FragmentOffset: 16,
				TTL:            48,
				Protocol:       99,
				SrcAddr:        "\x0a\x00\x00\x01",
				DstAddr:        dstAddrV4,
			})
			ip.SetChecksum(0)
			ip.SetChecksum(^ip.CalculateChecksum())
		}
		vv := hdr.View().ToVectorisedView()
		return stack.NewPacketBuffer(stack.PacketBufferOptions{Data: vv})
	}

	for _, test := range tests {
		for _, v6 := range []bool{true, false} {
			var (
				name   string
				protoF stack.NetworkProtocolFactory
				proto  tcpip.NetworkProtocolNumber
				mtu    uint32
				addr   tcpip.Address
			)
			if v6 {
				name = test.name + "-IPv6"
				protoF = ipv6.NewProtocol
				proto = header.IPv6ProtocolNumber
				mtu = header.IPv6MinimumMTU
				addr = dstAddrV6
			} else {
				name = test.name + "-IPv4"
				protoF = ipv4.NewProtocol
				proto = header.IPv4ProtocolNumber
				mtu = header.IPv4MinimumMTU
				addr = dstAddrV4
			}
			t.Run(name, func(t *testing.T) {
				s := stack.New(stack.Options{
					NetworkProtocols: []stack.NetworkProtocolFactory{protoF},
				})
				e := channel.New(0, mtu, linkAddr)
				if err := s.CreateNICWithOptions(nicID, e, stack.NICOptions{Name: nicName}); err != nil {
					t.Fatalf("CreateNIC(%d, _) = %s", nicID, err)
				}
				if err := s.AddAddress(nicID, proto, addr); err != nil {
					t.Fatalf("AddAddress(%d, %d, %s) = %s", nicID, proto, addr, err)
				}

				test.setup(t, s, v6)
				e.InjectInbound(proto, genPacket(v6))

				if got := int(s.Stats().IP.PacketsReceived.Value()); got != test.expectReceived {
					t.Errorf("got PacketReceived = %d, want = %d", got, test.expectReceived)
				}
				if got := int(s.Stats().IP.IPTablesInputDropped.Value()); got != test.expectInputDropped {
					t.Errorf("got IPTablesInputDropped = %d, want = %d", got, test.expectInputDropped)
				}
			})
		}
	}
}
