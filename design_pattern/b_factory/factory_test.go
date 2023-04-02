package b_factory

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactory(t *testing.T) {
	testCases := []struct {
		name    string
		ext     string
		path    string
		wantRes string
		wantErr error
	}{
		{
			name:    "JSON格式",
			ext:     "json",
			path:    "/local",
			wantRes: "Json/local",
			wantErr: nil,
		},
		{
			name:    "XML格式",
			ext:     "xml",
			path:    "/local",
			wantRes: "XML/local",
			wantErr: nil,
		},
		{
			name:    "unknown",
			ext:     "unknown",
			path:    "/local",
			wantRes: "",
			wantErr: errors.New("未知格式类型"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			factory, err := createFactory(testCase.ext)
			if err != nil {
				assert.Equal(t, err, testCase.wantErr)
				return
			}
			parse := factory.createParse()
			res := parse.Parse(testCase.path)
			assert.Equal(t, res, testCase.wantRes)
		})
	}

}
