package merkle

import (
	"crypto/sha256"
	"encoding/hex"
)

type MerkleTree interface {
	CalculateRootHash() string
	GetRootHash() string
}

type merkleTree struct {
	leafs    []string
	root     string
	hashType string
}

func New(dataList []string) MerkleTree {
	tree := &merkleTree{
		hashType: "sha256",
		leafs:    dataList,
		root:     "",
	}
	tree.CalculateRootHash()
	return tree
}

func (m *merkleTree) GetRootHash() string {
	return m.root
}

func (m *merkleTree) CalculateRootHash() string {
	nodes := m.leafs
	tmpHashList := getNewHashList(nodes)
	for {
		if len(tmpHashList) == 1 {
			break
		}
		tmpHashList = getNewHashList(tmpHashList)
	}
	m.root = tmpHashList[0]
	return m.root
}

func getNewHashList(list []string) []string {
	var newHashList []string
	for i := 0; i < len(list); i = i + 2 {
		idxLeft := i
		idxRight := i + 1
		left := list[idxLeft]

		right := ""
		if idxRight != len(list) {
			right = list[idxRight]
		}
		hashByte := sha256.Sum256([]byte(left + right))
		hashLeftAndRight := hex.EncodeToString(hashByte[:])
		newHashList = append(newHashList, hashLeftAndRight)
	}
	return newHashList
}
