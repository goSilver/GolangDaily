package d_prototype

import (
	"encoding/json"
	"time"
)

// Keyword 搜索关键字
type Keyword struct {
	word      string
	visit     int
	UpdatedAt *time.Time
}

// clone 这里使用序列化与反序列化方式深拷贝
func (k *Keyword) clone() *Keyword {
	var newKeyword Keyword
	marshal, _ := json.Marshal(k)
	json.Unmarshal(marshal, &newKeyword)
	return &newKeyword
}

// Keywords 关键字map
type Keywords map[string]*Keyword

// Clone 复制一个新的 keywords
// updatedWords: 需要更新的关键词列表，由于从数据库中获取数据常常是数组的方式
func (words Keywords) Clone(updatedKeywords []*Keyword) Keywords {
	newKeywords := Keywords{}

	for k, v := range words {
		newKeywords[k] = v
	}

	for _, keyword := range updatedKeywords {
		newKeywords[keyword.word] = keyword.clone()
	}

	return newKeywords
}
