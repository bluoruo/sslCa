/*
	Copyright (c) 2024, Harry.Ren <harry@cmche.com>

	这是一个SSL免费证书申请工具，通过调用Let's Encrypt的API来申请SSL证书，支持自动续期。
	为了方便使用，我们提供了操作lego来申请和续约，也集成了部分效验和操作证书的方法。

	调用方法：
		1. 设置证书文件的存放路径
			sslca.SetCertFile("cert.pem")
		2. 获取证书信息
			status, info := sslca.GetCertInfo()
			if status == 0 { // 成功
				fmt.Println(info)
			} else {	// 失败
				fmt.Println(sslca.ErrMsg(status))
			}
	update:
	 		2021-08-24: v0.01

*/

package sslCa

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// CertInfo 证书信息
type CertInfo struct {
	SerialNumber       string // 证书序列号
	Issuer             string // 证书颁发者
	Subject            string // 证书主题
	NotBefore          string // 证书有效期开始时间
	NotAfter           string // 证书有效期结束时间
	SignatureAlgorithm string // 证书签名算法
}

// certFile 证书文件路径
var certFile string

// SetCertFile 设置证书文件的存放路径
func SetCertFile(filename string) {
	certFile = filename
}

// GetCertInfo 获取证书信息
func GetCertInfo() (status int, info CertInfo) {
	// 判断 certFile 是否为空
	if certFile == "" {
		return 1, CertInfo{}
	}
	// 判断证书文件是否存在
	if !checkCertExist() {
		return 2, CertInfo{}
	}
	// 获取证书信息
	status, info = getCertInfo()

	return status, info
}

// ErrMsg 获取错误信息
func ErrMsg(status int) string {
	switch status {
	case 0:
		return "成功"
	case 1:
		return "证书文件路径为空"
	case 2:
		return "证书文件不存在"
	case 3:
		return "读取证书文件失败"
	case 4:
		return "证书解码失败"
	case 5:
		return "证书转换x509失败"
	default:
		return "未知错误"
	}
}

// checkCertExist 检查证书文件是否存在
func checkCertExist() bool {
	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		return false
	}
	return true
}

// getCertInfo 获取证书信息
func getCertInfo() (status int, info CertInfo) {
	// 读取证书文件
	certPem, err := os.ReadFile(certFile)
	if err != nil {
		fmt.Println("读取证书文件失败", err)
		return 3, CertInfo{}
	}
	// 解码证书
	block, _ := pem.Decode(certPem)
	if block == nil {
		fmt.Println("证书解码失败")
		return 4, CertInfo{}
	}
	// 将证书转换为x509证书
	caCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("证书转换x509失败", err)
		return 5, CertInfo{}
	}

	// 返回证书信息
	return 0, CertInfo{
		SerialNumber:       caCert.SerialNumber.String(),
		Issuer:             caCert.Issuer.String(),
		Subject:            caCert.Subject.String(),
		NotBefore:          caCert.NotBefore.String(),
		NotAfter:           caCert.NotAfter.String(),
		SignatureAlgorithm: caCert.SignatureAlgorithm.String(),
	}
}
