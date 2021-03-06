package utils

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"time"
)

var (
	// NonceStr 生成一个长度为 2*n 的随机字符串，可 mock
	NonceStr = func(n int) string {
		buf := make([]byte, n)
		_, err := rand.Read(buf)
		if err != nil {
			panic(err)
		}
		return hex.EncodeToString(buf)
	}
)

var (
	// Now 返回当前时间，可 mock
	Now = time.Now
)

// TLSConfig 根据 cert/key 以及（可选的）ca 创建一个 tls.Config，例如可用于配置 https:
//
// 	tlsConfig, err := TLSConfig(certPEMBlock, keyPEMBlock, caPEMBlock)
// 	if err != nil {
// 		...
// 	}
// 	client := &http.Client{
// 		Transport: &http.Transport{
// 			TLSClientConfig: tlsConfig,
// 		},
// 	}
func TLSConfig(certPEMBlock, keyPEMBlock, caPEMBlock []byte) (*tls.Config, error) {
	// key/cert pair
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		return nil, err
	}

	// 如果有 ca 提供，否则使用 host 本身的 ca
	var caCertPool *x509.CertPool
	if len(caPEMBlock) != 0 {
		caCertPool = x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caPEMBlock)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}, nil
}

// MustTLSConfig 是 must 版 TLSConfig
func MustTLSConfig(certPEMBlock, keyPEMBlock, caPEMBlock []byte) *tls.Config {
	tlsConfig, err := TLSConfig(certPEMBlock, keyPEMBlock, caPEMBlock)
	if err != nil {
		panic(err)
	}
	return tlsConfig
}
