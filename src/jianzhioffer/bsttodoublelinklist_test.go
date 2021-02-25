package jianzhioffer

import (
	"go-leetcode/src/datastructure"
	"reflect"
	"testing"
)

func Test_bstTODoubleLinkList(t *testing.T) {
	type args struct {
		root *datastructure.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *datastructure.DoubleLinkListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bstTODoubleLinkList(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstTODoubleLinkList() = %v, want %v", got, tt.want)
			}
		})
	}
}
