package server

import (
	"crypto/tls"
//	"io/ioutil"
	"github.com/GeertJohan/go.rice"

	"github.com/agurha/tunnel/pkg/log"
)

func LoadTLSConfig(crtPath string, keyPath string) (tlsConfig *tls.Config, err error) {
//	fileOrAsset := func(path string, default_path string) ([]byte, error) {
//		loadFn := ioutil.ReadFile
//		if path == "" {
//			loadFn = assets.Asset
//			path = default_path
//		}
//
//		return loadFn(path)
//	}

	box, err := rice.FindBox("assets/tls")

	var (
		crt  []byte
		key  []byte
		cert tls.Certificate
	)

	if crt, err = box.Bytes("device.crt"); err != nil {
		return
	}

	if key, err = box.Bytes("device.key"); err != nil {
		return
	}

	if cert, err = tls.X509KeyPair(crt, key); err != nil {
		return
	}

	tlsConfig = &tls.Config{ Certificates: []tls.Certificate{cert},InsecureSkipVerify: true }

	log.Info("tls config is %s", tlsConfig)

	return
}
