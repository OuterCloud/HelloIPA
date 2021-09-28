package ipa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBuildNumberFromIPA(t *testing.T) {
	// 定义测试用例参数
	type args struct {
		ipaLocalPath string
	}
	// 定义测试用例预期结果
	type want struct {
		buildNumber string
	}
	// 定义并初始化测试用例
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Get the CFBundleVersion from Info.plist successfully",
			args: args{
				ipaLocalPath: "./test.ipa",
			},
			want: want{
				buildNumber: "3412",
			},
		},
	}
	// 遍历用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 执行待测方法
			buildNumber, err := GetBuildNumberFromIPA(tt.args.ipaLocalPath)
			t.Logf("build number is->%s", buildNumber)
			// 断言
			assert.Nil(t, err)
			assert.Equal(t, tt.want.buildNumber, buildNumber)
		})
	}
}
