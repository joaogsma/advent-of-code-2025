package main

import (
	"golang.org/x/exp/constraints"
)

type Partitionable[T Number] interface {
	comparable

	IsLeftOf(partition T) bool
	IsRightOf(partition T) bool

	String() string
}

type Number interface {
	constraints.Integer | constraints.Float
}

type BspNode[N Number, O Partitionable[N]] struct {
	Partition N
	Objects   []O
	Left      *BspNode[N, O]
	Right     *BspNode[N, O]
}

func BuildBspTree[N Number, O Partitionable[N]](data []O, capacity int, splitFn func([]O) N) BspNode[N, O] {
	if len(data) <= capacity {
		var zeroPartition N
		return BspNode[N, O]{zeroPartition, data, nil, nil}
	}
	partition := splitFn(data)

	var leftData, rightData []O
	for _, obj := range data {
		if obj.IsLeftOf(partition) {
			leftData = append(leftData, obj)
		}
		if obj.IsRightOf(partition) {
			rightData = append(rightData, obj)
		}
	}

	if len(leftData) == len(data) || len(rightData) == len(data) {
		return BspNode[N, O]{partition, data, nil, nil}
	}
	leftChild := BuildBspTree(leftData, capacity, splitFn)
	rightChild := BuildBspTree(rightData, capacity, splitFn)
	return BspNode[N, O]{partition, []O{}, &leftChild, &rightChild}
}

func (node BspNode[N, O]) Search(value N, matchFn func(O, N) bool) *Set[O] {
	if len(node.Objects) > 0 {
		return node.SearchLeaf(value, matchFn)
	}
	foundLeft := node.Left.Search(value, matchFn)
	foundRight := node.Right.Search(value, matchFn)
	return foundLeft.Union(foundRight)
}

func (node BspNode[N, O]) SearchLeaf(value N, matchFn func(O, N) bool) *Set[O] {
	foundObjects := NewSet[O]()
	for _, obj := range node.Objects {
		if matchFn(obj, value) {
			foundObjects.Add(obj)
		}
	}
	return foundObjects
}
