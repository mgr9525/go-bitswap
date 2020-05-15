package ruisBitswap

import (
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
)

var MFilter IFilter

type IFilter interface {
	CheckWant(p peer.ID, d cid.Cid) bool
	GetSent(p peer.ID, ds cid.Cid)
	GetRecv(blk blocks.Block, from peer.ID)
}
