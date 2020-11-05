package tree

import (
	"bst/app"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

//вынесен отдельно, для удобства тестирования
var logger *logrus.Logger

func InitBST(log *logrus.Logger, initValues []int) app.Storage {
	if log != nil {
		logger = logrus.New()
		logger.SetFormatter(&logrus.JSONFormatter{})
		logger.SetOutput(os.Stdout)
	}
	newTree := Tree{
		Root:    nil,
		RWMutex: sync.RWMutex{},
	}
	if initValues != nil {
		for i := range initValues {
			newTree.Insert(initValues[i])
		}
	}
	return &newTree
}

type Tree struct {
	Root *Node
	sync.RWMutex
}
type Node struct {
	Value int
	left  *Node
	right *Node
}

func (t *Tree) clear() {
	t.Root = nil
}

func (t *Tree) Insert(value int) {
	t.Lock()
	defer t.Unlock()
	newNode := &Node{value, nil, nil}
	if t.Root == nil {
		logger.Info("bst пустое, вставка в корень")
		t.Root = newNode
	} else {
		logger.Info("bst не пустое, вставка через рекурсию")
		insertNode(t.Root, newNode)
	}
}

func insertNode(node, newNode *Node) {
	if newNode.Value < node.Value {
		if node.left == nil {
			logger.Info("левый потомок пустой. Вставка")
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			logger.Info("правый потомок пустой. Вставка")
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

func (t *Tree) Search(value int) bool {
	t.RLock()
	defer t.RUnlock()
	return search(t.Root, value)
}

func search(n *Node, value int) bool {
	if n == nil {
		logger.Info("узел пустой")
		return false
	}
	if value < n.Value {
		logger.Info("поиск по левой стороне")
		return search(n.left, value)
	}
	if value > n.Value {
		logger.Info("поиск по правой стороне")
		return search(n.right, value)
	}
	return true
}

func (t *Tree) Remove(value int) {
	t.Lock()
	defer t.Unlock()
	t.Root = remove(t.Root, value)
}

func remove(node *Node, value int) *Node {
	if node == nil {
		logger.Info("узел пустой")
		return nil
	}
	if value < node.Value {
		logger.Info("поиск элемента удаления по левой стороне")
		node.left = remove(node.left, value)
		return node
	}
	if value > node.Value {
		logger.Info("поиск элемента удаления по правой стороне")
		node.right = remove(node.right, value)
		return node
	}
	if node.left == nil && node.right == nil {
		logger.Info("потомков нет, удаление узла")
		node = nil
		return nil
	}
	if node.left == nil {
		logger.Info("потомков есть, создание новой связи между дочерним узлом и удаляемым родителем")
		node = node.right
		return node
	}
	if node.right == nil {
		logger.Info("потомков есть, создание новой связи между дочерним узлом и удаляемым родителем")
		node = node.left
		return node
	}
	rightSide := node.right
	logger.Info("есть оба потомка. Поиск элемента без левого потомка и его замена")
	for {
		if rightSide != nil && rightSide.left != nil {
			rightSide = rightSide.left
		} else {
			break
		}
	}
	node.Value, node.Value = rightSide.Value, rightSide.Value
	node.right = remove(node.right, node.Value)
	return node
}
