package v3

import (
	"fmt"
	"strings"
)

type router struct {
	// trees key是HTTP方法
	trees map[string]*node
}

// node 代表路由树的节点
// 路由树的匹配顺序是：
// 1.静态完全匹配
// 2.通配符匹配
// 这是不回溯匹配
type node struct {
	path string
	// 字节点
	children map[string]*node
	// 命中路由后执行的逻辑
	handler HandleFunc

	// 通配符*表达的节点，任意匹配
	starChild *node
}

func newRouter() router {
	return router{
		trees: map[string]*node{},
	}
}

// addRoute 注册路由
// method 是HTTP方法
// path 必须是/开始并且结尾不能有/，中间也不允许有连续的/
func (r *router) addRoute(method, path string, handler HandleFunc) {
	if len(path) == 0 {
		panic("web: 路由是空字符串")
	}
	if path[0] != '/' {
		panic("web: 路由必须以 / 开头")
	}

	if path != "/" && path[len(path)-1] == '/' {
		panic("web: 路由不能以 / 结尾")
	}

	root, ok := r.trees[method]
	if !ok {
		root = &node{path: "/"}
		r.trees[method] = root
	}
	if path == "/" {
		if root.handler != nil {
			panic("web: 路由冲突[/]")
		}
		root.handler = handler
		return
	}

	segs := strings.Split(path[1:], "/")
	// 开始一段段处理
	for _, seg := range segs {
		if len(seg) == 0 {
			panic(fmt.Sprintf("web: 非法路由。不允许使用 //a/b, /a//b 之类的路由, [%s]", path))
		}
		root = root.childOrCreate(seg)
	}
	if root.handler != nil {
		panic(fmt.Sprintf("web: 路由冲突[%s]", path))
	}
	root.handler = handler
}

// findRoute 查找对应的节点
func (r *router) findRoute(method, path string) (*node, bool) {
	// 先判断有没有对应http method的路由
	root, ok := r.trees[method]
	if !ok {
		return nil, false
	}
	// 如果是跟路径，直接返回root节点
	if path == "/" {
		return root, true
	}

	// 广度优先地遍历路由树，找到叶子节点
	segs := strings.Split(strings.Trim(path, "/"), "/")
	for _, seg := range segs {
		root, ok = root.childOf(seg)
		if !ok {
			return nil, false
		}
	}
	return root, true
}

func (n *node) childOf(path string) (*node, bool) {
	if n.children == nil {
		return n.starChild, n.starChild == nil
	}
	res, ok := n.children[path]
	if !ok {
		return n.starChild, n.starChild == nil
	}
	return res, ok
}

// childOrCreate 查找字节点，如果字节点不存在就创建一个
func (n *node) childOrCreate(path string) *node {
	// 处理通配符路径
	if path == "*" {
		if n.starChild == nil {
			n.starChild = &node{path: "*"}
		}
		return n.starChild
	}
	if n.children == nil {
		n.children = make(map[string]*node)
	}
	child, ok := n.children[path]
	if !ok {
		child = &node{path: path}
		n.children[path] = child
	}
	return child
}
