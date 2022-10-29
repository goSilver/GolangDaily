package v2

import (
	"fmt"
	"strings"
)

type router struct {
	// trees key是HTTP方法
	trees map[string]*node
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
	root, ok := r.trees[method]
	if !ok {
		return nil, false
	}
	if path == "/" {
		return root, true
	}
	segs := strings.Split(strings.Trim(path, "/"), "/")
	for _, seg := range segs {
		root, ok = root.childOf(seg)
		if !ok {
			return nil, false
		}
	}
	return root, true
}

type node struct {
	path string
	// 字节点
	children map[string]*node
	// 命中路由后执行的逻辑
	handler HandleFunc
}

func (n *node) childOf(path string) (*node, bool) {
	if n.children == nil {
		return nil, false
	}
	res, ok := n.children[path]
	return res, ok
}

// childOrCreate 查找字节点，如果字节点不存在就创建一个
func (n *node) childOrCreate(path string) *node {
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
