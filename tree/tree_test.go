package tree

import (
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
	"sync"
	"testing"
)

var testTree Tree

func init() {
	if logger == nil {
		logger = logrus.New()
		logger.SetFormatter(&logrus.JSONFormatter{})
		logger.SetOutput(os.Stdout)
	}
}
func TestSearch(t1 *testing.T) {
	type fields struct {
		root    *Node
		RWMutex *sync.RWMutex
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "negative_search_est", fields: fields{
			root: &Node{
				Value: 18,
				left:  nil,
				right: &Node{
					Value: 68,
					left:  nil,
					right: nil,
				},
			},
			RWMutex: &sync.RWMutex{},
		}, args: args{value: 190}, want: false},
		{name: "positive_search_test", fields: fields{
			root: &Node{
				Value: 18,
				left:  nil,
				right: &Node{
					Value: 68,
					left:  nil,
					right: &Node{
						Value: 115,
						left:  nil,
						right: nil,
					},
				},
			},
			RWMutex: &sync.RWMutex{},
		}, args: args{value: 115},
			want: true,
		},
		{name: "left_positive_search_test", fields: fields{
			root: &Node{
				Value: 18,
				left: &Node{
					Value: 3,
					left:  nil,
					right: nil,
				},
				right: &Node{
					Value: 68,
					left:  nil,
					right: &Node{
						Value: 115,
						left:  nil,
						right: nil,
					},
				},
			},
			RWMutex: &sync.RWMutex{},
		}, args: args{value: 3}, want: true},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				Root:    tt.fields.root,
				RWMutex: *tt.fields.RWMutex,
			}
			if got := t.Search(tt.args.value); got != tt.want {
				t1.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Remove(t1 *testing.T) {

	testRemoveTree := Tree{
		Root: &Node{
			Value: 18,
			left:  nil,
			right: nil,
		},
		RWMutex: sync.RWMutex{},
	}
	expectedTree := Tree{
		RWMutex: sync.RWMutex{},
	}
	testRemoveTree.Remove(18)
	if testRemoveTree.Root != expectedTree.Root {
		t1.Errorf("Expected tree with 42, got %v", testRemoveTree.Root.Value)
	}
}

func TestNotExistValInTree_Remove(t1 *testing.T) {

	testRemoveTree := Tree{
		Root: &Node{
			Value: 18,
			left:  nil,
			right: nil,
		},
		RWMutex: sync.RWMutex{},
	}
	expectedTree := Tree{
		Root: &Node{
			Value: 18,
			left:  nil,
			right: nil,
		},
		RWMutex: sync.RWMutex{},
	}
	testRemoveTree.Remove(19)
	if testRemoveTree.Root.Value != expectedTree.Root.Value {
		t1.Errorf("Expected tree with 18, got %v", testRemoveTree.Root.Value)
	}
}

func TestTree_clear(t1 *testing.T) {
	testClearTree := Tree{
		Root: &Node{
			Value: 18,
			left:  nil,
			right: nil,
		},
		RWMutex: sync.RWMutex{},
	}
	testClearTree.clear()
	if testClearTree.Root != nil {
		t1.Errorf("Expected empty tree but got %v", testClearTree.Root)
	}
}

func Test_remove(t *testing.T) {
	type args struct {
		node  *Node
		value int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{name: "left_shift", args: args{
			node: &Node{
				Value: 18,
				left:  nil,
				right: &Node{
					Value: 68,
					left: &Node{
						Value: 59,
						left:  nil,
						right: nil,
					},
					right: &Node{
						Value: 115,
						left:  nil,
						right: nil,
					},
				},
			},
			value: 18,
		},
			want: &Node{
				Value: 68,
				left: &Node{
					Value: 59,
					left:  nil,
					right: nil,
				},
				right: &Node{
					Value: 115,
					left:  nil,
					right: nil,
				},
			},
		},
		{name: "max_val_remove", args: args{
			node: &Node{
				Value: 18,
				left:  nil,
				right: &Node{
					Value: 68,
					left: &Node{
						Value: 59,
						left:  nil,
						right: nil,
					},
					right: &Node{
						Value: 115,
						left:  nil,
						right: nil,
					},
				},
			},
			value: 115,
		},
			want: &Node{
				Value: 18,
				left:  nil,
				right: &Node{
					Value: 68,
					left: &Node{
						Value: 59,
						left:  nil,
						right: nil,
					},
					right: nil,
				},
			},
		},
		{name: "left_leaf_remove", args: args{
			node: &Node{
				Value: 18,
				left: &Node{
					Value: 9,
					left:  nil,
					right: nil,
				},
				right: &Node{
					Value: 68,
					left: &Node{
						Value: 59,
						left:  nil,
						right: nil,
					},
					right: &Node{
						Value: 115,
						left:  nil,
						right: nil,
					},
				},
			},
			value: 9,
		},
			want: &Node{
				Value: 18,
				left:  nil,
				right: &Node{
					Value: 68,
					left: &Node{
						Value: 59,
						left:  nil,
						right: nil,
					},
					right: &Node{
						Value: 115,
						left:  nil,
						right: nil,
					},
				},
			},
		},
		{name: "right_leaf_nil", args: args{
			node: &Node{
				Value: 18,
				left: &Node{
					Value: 9,
					left:  nil,
					right: nil,
				},
			},
			value: 18,
		},
			want: &Node{
				Value: 9,
				left:  nil,
				right: nil,
			},
		},
		{name: "remove_from_middle", args: args{
			node: &Node{
				Value: 18,
				left:  nil,
				right: &Node{
					Value: 68,
					left: &Node{
						Value: 59,
						left:  nil,
						right: nil,
					},
					right: &Node{
						Value: 115,
						left: &Node{
							Value: 99,
							left:  nil,
							right: nil,
						},
						right: &Node{
							Value: 125,
							left:  nil,
							right: nil,
						},
					},
				},
			},
			value: 68,
		},
			want: &Node{
				Value: 18,
				left:  nil,
				right: &Node{
					Value: 99,
					left: &Node{
						Value: 59,
						left:  nil,
						right: nil,
					},
					right: &Node{
						Value: 115,
						left:  nil,
						right: &Node{
							Value: 125,
							left:  nil,
							right: nil,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := remove(tt.args.node, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertNode(t *testing.T) {
	type args struct {
		node    *Node
		newNode *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{name: "left_node_insert", args: args{
			node: &Node{
				Value: 25,
				left:  nil,
				right: nil,
			},
			newNode: &Node{
				Value: 13,
				left:  nil,
				right: nil,
			},
		}, want: &Node{
			Value: 25,
			left: &Node{
				Value: 13,
				left:  nil,
				right: nil,
			},
			right: nil,
		}},
		{name: "right_node_insert", args: args{
			node: &Node{
				Value: 25,
				left:  nil,
				right: nil,
			},
			newNode: &Node{
				Value: 30,
				left:  nil,
				right: nil,
			},
		}, want: &Node{
			Value: 25,
			left:  nil,
			right: &Node{
				Value: 30,
				left:  nil,
				right: nil,
			},
		},
		},

		{name: "right_node_not_nil", args: args{
			node: &Node{
				Value: 25,
				left:  nil,
				right: &Node{
					Value: 30,
					left:  nil,
					right: nil,
				},
			},
			newNode: &Node{
				Value: 45,
				left:  nil,
				right: nil,
			},
		}, want: &Node{
			Value: 25,
			left:  nil,
			right: &Node{
				Value: 30,
				left:  nil,
				right: &Node{
					Value: 45,
					left:  nil,
					right: nil,
				},
			},
		},
		},
		{name: "left_node_not_nil", args: args{
			node: &Node{
				Value: 25,
				left: &Node{
					Value: 18,
					left:  nil,
					right: nil,
				},
				right: &Node{
					Value: 30,
					left:  nil,
					right: nil,
				},
			},
			newNode: &Node{
				Value: 9,
				left:  nil,
				right: nil,
			},
		}, want: &Node{
			Value: 25,
			left: &Node{
				Value: 18,
				left: &Node{
					Value: 9,
					left:  nil,
					right: nil,
				},
				right: nil,
			},
			right: &Node{
				Value: 30,
				left:  nil,
				right: nil,
			},
		},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertNode(tt.args.node, tt.args.newNode)
			if !reflect.DeepEqual(tt.args.node, tt.want) {
				t.Errorf("remove() = %v, want %v", tt.args.node, tt.want)
			}
		})
	}

}

func TestTree_Insert(t1 *testing.T) {
	type fields struct {
		Root    *Node
		RWMutex *sync.RWMutex
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Tree
	}{
		{name: "empty_tree_insert", fields: fields{
			Root:    nil,
			RWMutex: &sync.RWMutex{},
		}, args: args{15},
			want: &Tree{
				Root: &Node{
					Value: 15,
					left:  nil,
					right: nil,
				},
				RWMutex: sync.RWMutex{},
			},
		},
		{name: "not_empty_tree_insert", fields: fields{
			Root: &Node{
				Value: 15,
				left:  nil,
				right: nil,
			},
			RWMutex: &sync.RWMutex{},
		}, args: args{7},
			want: &Tree{
				Root: &Node{
					Value: 15,
					left: &Node{
						Value: 7,
						left:  nil,
						right: nil,
					},
					right: nil,
				},
				RWMutex: sync.RWMutex{},
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				Root:    tt.fields.Root,
				RWMutex: *tt.fields.RWMutex,
			}
			t.Insert(tt.args.value)
			if !reflect.DeepEqual(t, tt.want) {
				t1.Errorf("insert() = %v, want %v", t, tt.want)
			}
		})
	}
}
