// Package ipa 存放和ipa相关的内容
package ipa

import (
	"archive/zip"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"howett.net/plist"
)

// IsPayloadAppInfoPlist 是Payload目录下app包中的Info.plist文件
func IsPayloadAppInfoPlist(file *zip.File) bool {
	if filepath.Base(file.Name) == "Info.plist" {
		log.Printf("found an Info.plist file->%s", file.Name)
		nameSplit := strings.Split(file.Name, "/")
		// 筛选Payload/**.app/Info.plist
		if len(nameSplit) == 3 {
			log.Printf("the target Info.plist file is->%s", file.Name)
			return true
		}
	}
	return false
}

// GetInfoPlistFileContent 获取Info.plist文件内容
func GetInfoPlistFileContent(file *zip.File) (map[string]interface{}, error) {
	log.Println("start to get the content of the Info.plist file")
	var result map[string]interface{}
	fp, err := file.Open()
	if err != nil {
		return result, fmt.Errorf("open file error->%+v", err)
	}
	bs, err := ioutil.ReadAll(fp)
	if err != nil {
		return result, fmt.Errorf("read file content error->%+v", err)
	}
	defer fp.Close()
	if _, err = plist.Unmarshal(bs, &result); err != nil {
		return result, fmt.Errorf("unmarshal Info.plist content error->%+v", err)
	}
	log.Printf("show the content of Info.plist->%+v", result)
	return result, nil
}

// GetBuildNumberFromIPA 从ipa中获取构建号
func GetBuildNumberFromIPA(ipaLocalPath string) (string, error) {
	// 读取ipa内容
	zr, err := zip.OpenReader(ipaLocalPath)
	if err != nil {
		return "", fmt.Errorf("read ipa error->%+v", err)
	}
	defer zr.Close()
	// 遍历ipa中的文件
	for _, file := range zr.File {
		// 过滤目标Info.plist文件
		if !IsPayloadAppInfoPlist(file) {
			continue
		}
		// 获取Info.plist内容
		content, err := GetInfoPlistFileContent(file)
		if err != nil {
			return "", fmt.Errorf("get Info.plist content error->%+v", err)
		}
		// 获取构建号
		buildNumber, okB := content["CFBundleVersion"]
		if okB {
			buildNumberStr := fmt.Sprintf("%+v", buildNumber)
			// 为了与AppStore中的构建号保持一致需要去掉Info.plist中CFBundleVersion的前缀"0"
			buildNumberStrWithoutPrefix := strings.TrimPrefix(buildNumberStr, "0")
			log.Printf("show the buildNumberStr from Info.plist->%s", buildNumberStr)
			log.Printf("show the buildNumberStrWithoutPrefix from Info.plist->%s", buildNumberStrWithoutPrefix)
			return buildNumberStrWithoutPrefix, nil
		}
	}
	return "", errors.New("no build number found")
}
