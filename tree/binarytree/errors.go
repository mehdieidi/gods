package binarytree

import "errors"

var (
	TreeNotEmptyErr     = errors.New("tree is not empty")
	LeftChildExistsErr  = errors.New("left child already exists")
	RightChildExistsErr = errors.New("right child already exists")
	MustBeLeafErr       = errors.New("node must be a leaf")
	HasTwoChildrenErr   = errors.New("node has two children")
)
