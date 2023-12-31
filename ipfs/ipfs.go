package ipfs

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/bulutcan99/go_ipfs_chain_builder/internal/aggregate"
)

type Node struct {
	Data     aggregate.AggregatedData
	Hash     string
	PrevHash string
	Next     *Node
}

func NewNode(data aggregate.AggregatedData, prevHash string) *Node {
	node := &Node{
		Data:     data,
		PrevHash: prevHash,
	}
	node.Hash = node.calculateHash()
	return node
}

func (n *Node) calculateHash() string {
	bytes, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(bytes)
	return hex.EncodeToString(hash[:])
}
