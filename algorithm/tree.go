package main

import (
	"container/list"
)

// 树节点
type treeNode[E comparable] struct {
	e     E
	left  *treeNode[E]
	right *treeNode[E]
	size  int
}

// 创建树节点
func newTreeNode[E comparable](e E) *treeNode[E] {
	return &treeNode[E]{e, nil, nil, 1}
}

// BST 二分搜索树, 不包含重复元素
type BST[E comparable] struct {
	root    *treeNode[E]
	size    int
	compare func(a, b E) int // 元素 e 必须是可比较的
}

// NewBST 创建二分搜索树
func NewBST[E comparable](compare func(a, b E) int) *BST[E] {
	return &BST[E]{nil, 0, compare}
}

// Clear 清空二分搜索树
func (bst *BST[E]) Clear() {
	bst.root = nil
	bst.size = 0
}

// GetSize 返回元素数量
func (bst *BST[E]) GetSize() int {
	return bst.size
}

// 返回以 node 为根节点的二分搜索树的元素个数
func (bst *BST[E]) sizeOf(node *treeNode[E]) int {
	if node == nil {
		return 0
	}
	return node.size
}

// IsEmpty 判断树是否为空
func (bst *BST[E]) IsEmpty() bool {
	return bst.size == 0
}

// Add 向树中添加元素
func (bst *BST[E]) Add(e E) {
	bst.root = bst.add(bst.root, e)
}

// 向以 node 为根节点的二分搜索树中添加元素 e, 并返回新的根节点
func (bst *BST[E]) add(node *treeNode[E], e E) (root *treeNode[E]) {
	if node == nil {
		bst.size++
		return newTreeNode[E](e)
	}

	res := bst.compare(e, node.e)
	if res < 0 {
		node.left = bst.add(node.left, e)
	} else if res > 0 {
		node.right = bst.add(node.right, e)
	}

	node.size = bst.sizeOf(node.left) + bst.sizeOf(node.right) + 1
	return node
}

// Contains 判断元素是否存在
func (bst *BST[E]) Contains(e E) bool {
	return bst.contains(bst.root, e)
}

// 在以 node 为根节点的二分搜索树中搜索是否存在元素 e
func (bst *BST[E]) contains(node *treeNode[E], e E) bool {
	if node == nil {
		return false
	}

	res := bst.compare(e, node.e)
	if res == 0 {
		return true
	}

	if res < 0 {
		return bst.contains(node.left, e)
	}
	return bst.contains(node.right, e)
}

// Minimum 返回二分搜索树中的最小元素
func (bst *BST[E]) Minimum() E {
	if bst.IsEmpty() {
		panic("bst is empty")
	}
	return bst.minimum(bst.root).e
}

// 返回以 node 为根节点的二分搜索树中的最小元素所在的节点
func (bst *BST[E]) minimum(node *treeNode[E]) (minNode *treeNode[E]) {
	if node.left == nil {
		return node
	}
	return bst.minimum(node.left)
}

// Maximum 返回二分搜索树中的最大元素
func (bst *BST[E]) Maximum() E {
	if bst.IsEmpty() {
		panic("bst is empty")
	}
	return bst.maximum(bst.root).e
}

// 返回以 node 为根节点的二分搜索树中的最大元素所在的节点
func (bst *BST[E]) maximum(node *treeNode[E]) (maxNode *treeNode[E]) {
	if node.right == nil {
		return node
	}
	return bst.maximum(node.right)
}

// RemoveMin 删除二分搜索树中的最小元素并返回
func (bst *BST[E]) RemoveMin() E {
	minimum := bst.Minimum()
	bst.root = bst.removeMin(bst.root)
	return minimum
}

// 删除以 node 为根节点的二分搜索树的最小元素所在的节点, 并返回新的根节点
func (bst *BST[E]) removeMin(node *treeNode[E]) (root *treeNode[E]) {
	if node.left == nil {
		rightNode := node.right
		bst.size--
		return rightNode
	}

	node.left = bst.removeMin(node.left)
	node.size = bst.sizeOf(node.left) + bst.sizeOf(node.right) + 1
	return node
}

// RemoveMax 删除二分搜索树中的最大元素并返回
func (bst *BST[E]) RemoveMax() E {
	maximum := bst.Maximum()
	bst.root = bst.removeMax(bst.root)
	return maximum
}

// 删除以 node 为根节点的二分搜索树的最大元素所在的节点, 并返回新的根节点
func (bst *BST[E]) removeMax(node *treeNode[E]) (root *treeNode[E]) {
	if node.right == nil {
		leftNode := node.left
		bst.size--
		return leftNode
	}

	node.right = bst.removeMax(node.right)
	node.size = bst.sizeOf(node.left) + bst.sizeOf(node.right) + 1
	return node
}

// Remove 删除二分搜索树中的指定元素
func (bst *BST[E]) Remove(e E) {
	bst.root = bst.remove(bst.root, e)
}

// 以 node 为根节点的二分搜索树, 删除指定元素 e 所在的节点, 并返回新的根节点
func (bst *BST[E]) remove(node *treeNode[E], e E) (root *treeNode[E]) {
	if node == nil {
		return nil
	}

	var retNode *treeNode[E]
	res := bst.compare(e, node.e)
	if res < 0 {
		node.left = bst.remove(node.left, e)
		retNode = node
	} else if res > 0 {
		node.right = bst.remove(node.right, e)
		retNode = node
	} else {
		if node.left == nil {
			rightNode := node.right
			node.right = nil // help gc
			bst.size--
			retNode = rightNode
		} else if node.right == nil {
			leftNode := node.left
			node.left = nil // help gc
			bst.size--
			retNode = leftNode
		} else {
			successor := bst.minimum(node.right)
			successor.right = bst.removeMin(node.right) // 已经 size-- 了
			successor.left = node.left
			// help gc
			node.left = nil
			node.right = nil
			retNode = successor
		}
	}

	// 删除叶子节点后, 返回的 retNode 就为 null
	if retNode == nil {
		return nil
	}
	retNode.size = bst.sizeOf(retNode.left) + bst.sizeOf(retNode.right) + 1
	return retNode
}

// PreOrder 前序遍历
func (bst *BST[E]) PreOrder() []E {
	var result []E
	bst.preOrder(bst.root, &result)
	return result
}

// 对以 node 为根节点的二分搜索树进行前序遍历
func (bst *BST[E]) preOrder(node *treeNode[E], result *[]E) {
	if node == nil {
		return
	}
	*result = append(*result, node.e)
	bst.preOrder(node.left, result)
	bst.preOrder(node.right, result)
}

// InOrder 中序遍历
func (bst *BST[E]) InOrder() []E {
	var result []E
	bst.inOrder(bst.root, &result)
	return result
}

// 对以 node 为根节点的二分搜索树进行中序遍历
func (bst *BST[E]) inOrder(node *treeNode[E], result *[]E) {
	if node == nil {
		return
	}
	bst.inOrder(node.left, result)
	*result = append(*result, node.e)
	bst.inOrder(node.right, result)
}

// PostOrder 后序遍历
func (bst *BST[E]) PostOrder() []E {
	var result []E
	bst.postOrder(bst.root, &result)
	return result
}

// 对以 node 为根节点的二分搜索树进行后序遍历
func (bst *BST[E]) postOrder(node *treeNode[E], result *[]E) {
	if node == nil {
		return
	}
	bst.postOrder(node.left, result)
	bst.postOrder(node.right, result)
	*result = append(*result, node.e)
}

// LevelOrder 层序遍历
func (bst *BST[E]) LevelOrder() []E {
	var result []E
	if bst.IsEmpty() {
		return result
	}

	queue := list.New()
	queue.PushBack(bst.root)

	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(*treeNode[E])
		result = append(result, cur.e)

		if cur.left != nil {
			queue.PushBack(cur.left)
		}
		if cur.right != nil {
			queue.PushBack(cur.right)
		}
	}

	return result
}

// Floor 寻找 e 的 floor 值 (<= e 的最大值)
func (bst *BST[E]) Floor(e E) (E, bool) {
	if bst.IsEmpty() {
		panic("bst is empty")
	}
	if bst.compare(e, bst.Minimum()) < 0 {
		var empty E
		return empty, false
	}
	return bst.floor(bst.root, e).e, true
}

// 在以 node 为根节点的二分搜索树中搜索元素 e 的 floor 节点
func (bst *BST[E]) floor(node *treeNode[E], e E) *treeNode[E] {
	if node == nil {
		return nil
	}
	res := bst.compare(node.e, e)

	// node.e == e
	// 则 node 本身就是 e 的 floor 节点
	if res == 0 {
		return node
	}

	// node.e > e
	// 则要寻找的 e 的 floor 节点一定在 node 的左子树中
	if res > 0 {
		return bst.floor(node.left, e)
	}

	// node.e < e
	// 则 node 有可能是 e 的 floor 节点, 也有可能不是(存在比 node.e 大但是小于 e 的其余节点)
	// 需要尝试向 node 的右子树寻找一下
	tempNode := bst.floor(node.right, e)
	if tempNode != nil {
		return tempNode
	}
	return node
}

// Ceil 寻找 e 的 ceil 值  (>= e 的最小值)
func (bst *BST[E]) Ceil(e E) (E, bool) {
	if bst.IsEmpty() {
		panic("bst is empty")
	}
	if bst.compare(e, bst.Maximum()) > 0 {
		var empty E
		return empty, false
	}
	return bst.ceil(bst.root, e).e, true
}

// 在以 node 为根节点的二分搜索树中搜索元素 e 的 ceil 节点
func (bst *BST[E]) ceil(node *treeNode[E], e E) *treeNode[E] {
	if node == nil {
		return nil
	}
	res := bst.compare(node.e, e)

	// node.e == e
	// 则 node 本身就是 e 的 ceil 节点
	if res == 0 {
		return node
	}

	// node.e < e
	// 则要寻找的 e 的 ceil 节点一定在 node 的右子树中
	if res < 0 {
		return bst.ceil(node.right, e)
	}

	// node.e > e
	// 则 node 有可能是 e 的 ceil 节点, 也有可能不是(存在比 node.e 小但是大于 e 的其余节点)
	// 需要尝试向 node 的左子树寻找一下
	tempNode := bst.ceil(node.left, e)
	if tempNode != nil {
		return tempNode
	}
	return node
}

// Rank 寻找 e 的 rank 值 (0-based)
func (bst *BST[E]) Rank(e E) int {
	if !bst.Contains(e) {
		panic("element not found")
	}
	return bst.rank(bst.root, e)
}

// 在以 node 为根节点的二分搜索树中搜索元素 e 的 rank 值
func (bst *BST[E]) rank(node *treeNode[E], e E) int {
	res := bst.compare(e, node.e)
	if res == 0 {
		return bst.sizeOf(node.left)
	}
	if res < 0 {
		return bst.rank(node.left, e)
	}
	return bst.sizeOf(node.left) + 1 + bst.rank(node.right, e)
}

// Select 寻找排名为 index 的元素 (0-based)
func (bst *BST[E]) Select(index int) E {
	if index < 0 || index >= bst.GetSize() {
		panic("index out of range, need 0 <= index < size")
	}
	return bst.selectNode(bst.root, index).e
}

// 在以 node 为根节点的二分搜索树中搜索排名为 index 的元素 (0-based)
// 排名为 index 的元素, 它的前面有 index 个比它小的元素
func (bst *BST[E]) selectNode(node *treeNode[E], index int) *treeNode[E] {
	leftSize := bst.sizeOf(node.left)
	if index == leftSize {
		return node
	}
	if index < leftSize {
		return bst.selectNode(node.left, index)
	}
	return bst.selectNode(node.right, index-leftSize-1)
}
