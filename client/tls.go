package client

import (
	_ "crypto/sha512"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/GeertJohan/go.rice"

	"github.com/agurha/tunnel/pkg/log"

)

func LoadTLSConfig(rootCertPaths []string) (*tls.Config, error) {
	pool := x509.NewCertPool()

	box, err := rice.FindBox("assets/tls")

	boxHTTPFiles, err := rice.FindBox("assets/public")

	log.Info("info about http files %s", boxHTTPFiles)


	if err != nil {
		log.Error("error opening rice.Box: %s\n", err)
	}

	for _, certPath := range rootCertPaths {

		log.Info("cert path is %s", certPath)
		rootCrt, err := box.Bytes(certPath) // assets.Asset(certPath)

		log.Info("rootCrt value is %s", rootCrt)

		if err != nil {
			return nil, err
		}

		pemBlock, _ := pem.Decode(rootCrt)
		if pemBlock == nil {
			return nil, fmt.Errorf("Bad PEM data")
		}

		certs, err := x509.ParseCertificates(pemBlock.Bytes)
		if err != nil {
			return nil, err
		}

		pool.AddCert(certs[0])
	}

	return &tls.Config{RootCAs: pool, InsecureSkipVerify: true}, nil
}
