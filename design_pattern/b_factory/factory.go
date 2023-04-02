package b_factory

import "errors"

// NormalParseFactory 工厂方法
type NormalParseFactory interface {
	createParse() ConfigParse
}

type JsonNormalParseFactory struct {
}

func (json *JsonNormalParseFactory) createParse() ConfigParse {
	// 假装此处有复杂的构造逻辑
	return &JsonConfigParse{}
}

type XmlNormalParseFactory struct {
}

func (xml *XmlNormalParseFactory) createParse() ConfigParse {
	// 假装此处有复杂的构造逻辑
	return &XmlConfigParse{}
}

// createFactory 创建工厂
func createFactory(ext string) (NormalParseFactory, error) {
	switch ext {
	case "json":
		return &JsonNormalParseFactory{}, nil
	case "xml":
		return &XmlNormalParseFactory{}, nil
	}
	return nil, errors.New("未知格式类型")
}
