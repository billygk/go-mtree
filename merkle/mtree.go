package merkle

import (
	"crypto/sha256"
	"crypto/sha512"
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

// hashType [sha256|sha512] to be used. Defaults to sha256
func New(dataList []string, hashType string) MerkleTree {
	hType := "sha256"
	if hashType == "sha512" {
		hType = "sha512"
	}
	tree := &merkleTree{
		hashType: hType,
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
	tmpHashList := m.getNewHashList(nodes)
	for {
		if len(tmpHashList) == 1 {
			break
		}
		tmpHashList = m.getNewHashList(tmpHashList)
	}
	m.root = tmpHashList[0]
	return m.root
}

func (m *merkleTree) getNewHashList(list []string) []string {
	var newHashList []string
	for i := 0; i < len(list); i = i + 2 {
		left := list[i]
		right := ""
		if i+1 != len(list) {
			right = list[i+1]
		}

		hashLeftAndRight := ""
		if m.hashType == "sha256" {
			hashByte := sha256.Sum256([]byte(left + right))
			hashLeftAndRight = hex.EncodeToString(hashByte[:])
		} else {
			hashByte := sha512.Sum512([]byte(left + right))
			hashLeftAndRight = hex.EncodeToString(hashByte[:])
		}

		newHashList = append(newHashList, hashLeftAndRight)
	}
	return newHashList
}
