package tree

import (
	"testing"

	"go-leetcode/src/datastructure"
)

var l31OriginLeft = &datastructure.ConnectNode{4, nil, nil, nil}
var l31OriginRight = &datastructure.ConnectNode{5, nil, nil, nil}
var l32OriginLeft = &datastructure.ConnectNode{6, nil, nil, nil}
var l32OriginRight = &datastructure.ConnectNode{7, nil, nil, nil}
var l2OriginLeft = &datastructure.ConnectNode{2, l31OriginLeft, l31OriginRight, nil}
var l2OriginRight = &datastructure.ConnectNode{3, l32OriginLeft, l32OriginRight, nil}

var l32Right = &datastructure.ConnectNode{7, nil, nil, nil}
var l32Left = &datastructure.ConnectNode{6, nil, nil, l32Right}
var l31Right = &datastructure.ConnectNode{5, nil, nil, l32Left}
var l31Left = &datastructure.ConnectNode{4, nil, nil, l31Right}
var l2Right = &datastructure.ConnectNode{3, l32Left, l32Right, nil}
var l2Left = &datastructure.ConnectNode{2, l31Left, l31Right, l2Right}

var Input = &datastructure.ConnectNode{
	1,
	l2OriginLeft, l2OriginRight, nil,
}

var Output = &datastructure.ConnectNode{
	1,
	l2Left, l2Right, nil,
}

func TestConnectRecursive(t *testing.T) {
	var tests = []struct {
		in       *datastructure.ConnectNode
		expected *datastructure.ConnectNode
	}{
		// Add Testcase Here
		{Input, Output},
	}
	for _, test := range tests {
		result := connectRecursive(test.in)
		if result != test.expected {
			t.Fatalf("")
		}
	}
}

func TestConnectIterative(t *testing.T) {
	var tests = []struct {
		in       *datastructure.ConnectNode
		expected *datastructure.ConnectNode
	}{
		// Add Testcase Here
		{Input, Output},
	}
	for _, test := range tests {
		result := connectRecursive(test.in)
		if result != test.expected {
			t.Fatalf("")
		}
	}
}
