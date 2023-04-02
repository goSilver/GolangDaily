package b_factory

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	testCases := []struct {
		name    string
		ext     string
		path    string
		wantRes string
		wantErr error
	}{
		{
			name:    "XML格式",
			ext:     "xml",
			path:    "/local",
			wantRes: "XML/local",
			wantErr: nil,
		},
		{
			name:    "JSON格式",
			ext:     "json",
			path:    "/local",
			wantRes: "Json/local",
			wantErr: nil,
		},
		{
			name:    "未知格式",
			ext:     "unknown",
			path:    "/local",
			wantRes: "",
			wantErr: errors.New("未知文件类型"),
		},
	}

	factory := SimpleParseFactory{}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			xmlParser, err := factory.create(testCase.ext)
			if err != nil {
				assert.Equal(t, err, testCase.wantErr)
				return
			}
			res := xmlParser.Parse(testCase.path)
			assert.Equal(t, res, testCase.wantRes)
		})
	}

}
