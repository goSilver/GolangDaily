package b_factory

import "errors"

// ConfigParse 配置解析接口
type ConfigParse interface {
	Parse(path string) string
}

// JsonConfigParse Json解析类
type JsonConfigParse struct {
}

func (json *JsonConfigParse) Parse(path string) string {
	return "Json" + path
}

// XmlConfigParse xml解析类
type XmlConfigParse struct {
}

func (xml *XmlConfigParse) Parse(path string) string {
	return "XML" + path
}

// SimpleParseFactory 简单工厂
// 这种实现每次都是创建一个新的对象
type SimpleParseFactory struct {
}

func (simple *SimpleParseFactory) create(ext string) (ConfigParse, error) {
	switch ext {
	case "json":
		// 这里是简单构造逻辑
		return &JsonConfigParse{}, nil
	case "xml":
		// 这里是简单构造逻辑
		return &XmlConfigParse{}, nil
	}
	return nil, errors.New("未知文件类型")
}

// SimpleParseMapFactory 简单工厂
// 这种实现提前将对象创建好，缓存在一个map中，无须每次重复创建
type SimpleParseMapFactory struct {
	parseMap map[string]ConfigParse
}

func NewSimpleParseMapFactory() SimpleParseMapFactory {
	return SimpleParseMapFactory{
		parseMap: map[string]ConfigParse{
			"json": &JsonConfigParse{},
			"xml":  &XmlConfigParse{},
		},
	}
}

func (simple *SimpleParseMapFactory) create(ext string) ConfigParse {
	return simple.parseMap[ext]
}
