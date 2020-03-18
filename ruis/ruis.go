package ruisBitswap

import (
	"github.com/ipfs/go-bitswap/decision"
	blocks "github.com/ipfs/go-block-format"
	"github.com/libp2p/go-libp2p-peer"
)

var MFilter IFilter

type IFilter interface {
	CheckSend(e *decision.Envelope) bool
	GetRecv(blk blocks.Block, from peer.ID)
}
