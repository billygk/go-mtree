package main

import (
	"fmt"

	"github.com/billygk/go-mtree/merkle"
)

//var mtree *merkle.MerkleTree

func main() {
	var dataList []string
	dataList = append(dataList, "Node 0")
	dataList = append(dataList, "Node 1")
	dataList = append(dataList, "Node 2")
	dataList = append(dataList, "Node 3")
	dataList = append(dataList, "Node 4")
	dataList = append(dataList, "Node 5")
	dataList = append(dataList, "Node 6")
	dataList = append(dataList, "Node 7")
	dataList = append(dataList, "Node 8")
	dataList = append(dataList, "Node 9")

	mtree := merkle.New(dataList)
	fmt.Println(mtree.GetRootHash())
}
