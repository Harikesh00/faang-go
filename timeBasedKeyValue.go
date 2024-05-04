package main

/*
type TimeMap struct {
	data map[string]*TreeMap
}

type Value struct {
	timestamp int
	value     string
}

type TreeNode struct {
	value Value
	left  *TreeNode
	right *TreeNode
}

type TreeMap struct {
	root *TreeNode
}

func Constructor() TimeMap {
	return TimeMap{data: make(map[string]*TreeMap)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.data[key]; !ok {
		this.data[key] = &TreeMap{}
	}
	this.data[key].insert(Value{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if tree, ok := this.data[key]; ok {
		node := tree.search(timestamp)
		if node != nil {
			return node.value.value
		}
	}
	return ""
}

func (t *TreeMap) insert(value Value) {
	if t.root == nil {
		t.root = &TreeNode{value: value}
		return
	}
	t.root = t.insertNode(t.root, value)
}

func (t *TreeMap) insertNode(node *TreeNode, value Value) *TreeNode {
	if node == nil {
		return &TreeNode{value: value}
	}
	if value.timestamp <= node.value.timestamp {
		node.left = t.insertNode(node.left, value)
	} else {
		node.right = t.insertNode(node.right, value)
	}
	return node
}

func (t *TreeMap) search(timestamp int) *TreeNode {
	return t.searchNode(t.root, timestamp)
}

func (t *TreeMap) searchNode(node *TreeNode, timestamp int) *TreeNode {
	if node == nil {
		return nil
	}
	if timestamp == node.value.timestamp {
		return node
	} else if timestamp < node.value.timestamp {
		return t.searchNode(node.left, timestamp)
	} else {
		right := t.searchNode(node.right, timestamp)
		if right != nil {
			return right
		}
		return node
	}
}

func main() {
	kv := Constructor()
	kv.Set("foo", "bar1", 1)
	kv.Set("foo", "bar2", 2)
	fmt.Println(kv.Get("foo", 1)) // Output: "bar1"
	fmt.Println(kv.Get("foo", 3)) // Output: "bar2"
}
*/
