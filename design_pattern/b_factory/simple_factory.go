package b_factory

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
type SimpleParseFactory struct {
}

func (simple *SimpleParseFactory) create(ext string) ConfigParse {
	switch ext {
	case "json":
		// 这里是简单构造逻辑
		return &JsonConfigParse{}
	case "xml":
		// 这里是简单构造逻辑
		return &XmlConfigParse{}
	}
	return nil
}
