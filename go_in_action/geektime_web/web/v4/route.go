package v4

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

	paramChild *node
}

type matchInfo struct {
	n          *node
	pathParams map[string]string
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
func (r *router) findRoute(method, path string) (*matchInfo, bool) {
	// 先判断有没有对应http method的路由
	root, ok := r.trees[method]
	if !ok {
		return nil, false
	}
	// 如果是跟路径，直接返回root节点
	if path == "/" {
		return &matchInfo{n: root}, true
	}

	// 广度优先地遍历路由树，找到叶子节点
	segs := strings.Split(strings.Trim(path, "/"), "/")
	mi := &matchInfo{}
	for _, seg := range segs {
		var matchParam bool
		root, matchParam, ok = root.childOf(seg)
		if !ok {
			return nil, false
		}
		if matchParam {
			mi.addValue(root.path[1:], seg)
		}
	}
	mi.n = root
	return mi, true
}

func (n *node) childOf(path string) (*node, bool, bool) {
	if n.children == nil {
		if n.paramChild != nil {
			return n.paramChild, true, true
		}
		return n.starChild, false, n.starChild == nil
	}
	res, ok := n.children[path]
	if !ok {
		if n.paramChild != nil {
			return n.paramChild, true, true
		}
		return n.starChild, false, n.starChild == nil
	}
	return res, false, ok
}

// childOrCreate 查找字节点，如果字节点不存在就创建一个
// 首先会判断 path 是不是通配符路径
// 其次判断 path 是不是参数路径，即以 : 开头的路径
// 最后会从 children 里面查找，
// 如果没有找到，那么会创建一个新的节点，并且保存在 node 里面
func (n *node) childOrCreate(path string) *node {
	// 处理通配符路径
	if path == "*" {
		if n.paramChild != nil {
			panic(fmt.Sprintf("web: 非法路由，已有路径参数路由。不允许同时注册通配符路由和参数路由 [%s]", path))
		}
		if n.starChild == nil {
			n.starChild = &node{path: "*"}
		}
		return n.starChild
	}

	// 以 : 开头，我们认为是参数路由
	if path[0] == ':' {
		if n.starChild != nil {
			panic(fmt.Sprintf("web: 非法路由，已有通配符路由。不允许同时注册通配符路由和参数路由 [%s]", path))
		}
		if n.paramChild != nil {
			if n.paramChild.path != path {
				panic(fmt.Sprintf("web: 路由冲突，参数路由冲突，已有 %s，新注册 %s", n.paramChild.path, path))
			}
		} else {
			n.paramChild = &node{path: path}
		}
		return n.paramChild
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

func (m *matchInfo) addValue(key string, value string) {
	if m.pathParams == nil {
		// 大多数情况，参数路径只会有一段
		m.pathParams = map[string]string{key: value}
	}
	m.pathParams[key] = value
}
