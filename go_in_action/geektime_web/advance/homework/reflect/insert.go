package reflect

import (
	"errors"
	"reflect"
	"strings"
)

var errInvalidEntity = errors.New("invalid entity")

// InsertStmt 作业里面我们这个只是生成 SQL，所以在处理 sql.NullString 之类的接口
// 只需要判断有没有实现 driver.Valuer 就可以了
func InsertStmt(entity interface{}) (string, []interface{}, error) {
	if entity == nil {
		return "", nil, errInvalidEntity
	}
	val := reflect.ValueOf(entity)
	typ := val.Type()
	// 检测 entity 是否符合我们的要求
	// 我们只支持有限的几种输入
	if typ.Kind() == reflect.Ptr { // 处理一级指针
		typ = typ.Elem()
		val = val.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return "", nil, errInvalidEntity
	}

	// 使用 strings.builder 来拼接 字符串
	// bd := strings.builder{}
	sb := strings.Builder{}
	_, _ = sb.WriteString("INSERT INTO `")

	// 构造 INSERT INTO XXX，XXX 是你的表名，这里我们直接用结构体名字
	sb.WriteString(typ.Name())
	sb.WriteString("`(")

	// 遍历所有的字段，构造出来的是 INSERT INTO XXX(col1, col2, col3)
	// 在这个遍历的过程中，你就可以把参数构造出来
	// 如果你打算支持组合，那么这里你要深入解析每一个组合的结构体
	// 并且层层深入进去
	fields, values := fieldNameAndValue(val)

	// 拼接 VALUES，达成 INSERT INTO XXX(col1, col2, col3) VALUES
	for i, name := range fields {
		if i > 0 {
			sb.WriteRune(',')
		}
		sb.WriteRune('`')
		sb.WriteString(name)
		sb.WriteRune('`')
	}
	sb.WriteString(") VALUES (")
	args := make([]interface{}, 0, len(values))
	// 再一次遍历所有的字段，要拼接成 INSERT INTO XXX(col1, col2, col3) VALUES(?,?,?)
	// 注意，在第一次遍历的时候我们就已经拿到了参数的值，所以这里就是简单拼接 ?,?,?
	for i, field := range fields {
		if i > 0 {
			sb.WriteRune(',')
		}
		sb.WriteRune('?')
		args = append(args, values[field])
	}
	if len(args) == 0 {
		return "", nil, errInvalidEntity
	}
	sb.WriteString(");")
	return sb.String(), args, nil
}

func fieldNameAndValue(val reflect.Value) ([]string, map[string]interface{}) {
	typ := val.Type()
	fieldNum := val.NumField()
	fields := make([]string, 0, fieldNum)
	values := make(map[string]interface{}, fieldNum)
	for i := 0; i < fieldNum; i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		// 处理组合
		if field.Type.Kind() == reflect.Struct && field.Anonymous {
			subFields, subValues := fieldNameAndValue(fieldVal)
			for _, k := range subFields {
				if _, ok := values[k]; ok {
					continue // 组合情况下，忽略重复字段
				}
				fields = append(fields, k)
				values[k] = subValues[k]
			}
			continue
		}

		fields = append(fields, field.Name)
		values[field.Name] = fieldVal.Interface()
	}
	return fields, values
}
