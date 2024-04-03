package myweb

import "strings"

// trie树的节点定义
type node struct {
	pattern  string  //待匹配路由，如/p/:lang
	part     string  //路由中一部分，如:lang
	children []*node //子节点，如[doc,tutorial,intro]
	isWild   bool    //是否精确匹配，part包含:或*时为true
}

// 返回第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part {
			return child
		}
	}
	return nil
}

// 返回所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	childRens := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			childRens = append(childRens, child)
		}
	}
	return childRens
}

// 递归插入节点,pattern表示待匹配路由，parts表示分段部分，height是索引
func (n *node) insert(pattern string, parts []string, height int) {
	//注意如/p/:lang/doc只在第三曾节点pattern才设置pattern为/p/:lang/doc,p和:lang节点pattern属性都为空
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// 递归查找匹配，parts表示分段部分，height表示索引
func (n *node) search(parts []string, height int) *node {
	//匹配到最后一层或*退出
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
