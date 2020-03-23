package ruisBitswap

import (
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-peer"
)

var MFilter IFilter

type IFilter interface {
	CheckWant(p peer.ID, d cid.Cid) bool
	CheckSend(p peer.ID, ds cid.Cid) bool
	GetRecv(blk blocks.Block, from peer.ID)
}
