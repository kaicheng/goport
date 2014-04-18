package engineio

import (
    "github.com/kaicheng/goport/engineio/parser"
    "github.com/kaicheng/goport/events"
)

type websocketTransport struct {
    events.EventEmitter

    socket *Socket
    writable bool
}

func (trans *websocketTransport)) onData(data []byte) {
    trans.onPacket(parser.DecodePacket(data))
}

func (trans *websocketTransport) send(packets []*parser.Packet) {
    for i, pkt := range packets {
        parser.EncodePacket(pkt, trans.supportsbinary, func(data []byte) {
            trans.writable = false
            trans.socket.send(data, function(err){
                if err {
                    trans.onError("write error", err)
                }
                trans.writable = true
                trans.Emit("drain")
            })
        })
    }
}

func (trans *websocketTransport) doClose(fn func()) {
    trans.socket.close()
    if fn != nil {
        fn()
    }
}

func (trans *websocketTransport) onClose() {
    trans.readyState = "closed"
    trans.Emit("close")
}

func (trans *websocketTransport) onPacket(pkt *parser.Packet) {
    trans.Emit("packet", pkt)
}

func (trans *websocketTransport) onError(msg, desc string) {
    err := Error{"TransportError", msg, desc}
    trans.Emit("error", err)
}

func (trans *websocketTransport) close (fn func()) {
    trans.readyState = "closing"
    trans.doClose(fn)
}

func (trans *Transport) onRequest(req *Request) {
    trans.req = req
}
