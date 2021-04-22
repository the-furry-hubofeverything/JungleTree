package handlers

import (
    "encoding/json"
    "github.com/junglemc/mc"
    "github.com/junglemc/mc/status"
    "github.com/junglemc/net"
    "github.com/junglemc/net/codec"
    "github.com/junglemc/net/packet"
)

func statusRequest(c *net.Client, p codec.Packet) (err error) {
    response := status.ServerListResponse{
        Version: status.GameVersion{
            Name:     "1.16.5",
            Protocol: 754,
        },
        Players: status.ServerListPlayers{
            Max:    10,
            Online: 0,
            Sample: []status.ServerListPlayer{},
        },
        Description: mc.Chat{Text: "A JungleTree Server"},
    }

    data, err := json.Marshal(response)
    if err != nil {
        return
    }

    responsePkt := &packet.ClientboundStatusResponsePacket{Response: string(data)}
    return c.Send(responsePkt)
}

func statusPing(c *net.Client, p codec.Packet) (err error) {
    return c.Send(&packet.ClientboundStatusPongPacket{Time: p.(packet.ServerboundStatusPingPacket).Time})
}
