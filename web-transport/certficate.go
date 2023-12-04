package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"math/big"
	"time"
	// "encoding/base64"
)

func generateCertificate(subject []string) (tls.Certificate, error) {
	var certificate tls.Certificate
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return certificate, err
	}
	serialNumber, err := rand.Int(rand.Reader, big.NewInt(0x7FFFFFFF))
	if err != nil {
		return certificate, err
	}
	var certTemplate = x509.Certificate{
		SignatureAlgorithm: x509.ECDSAWithSHA256,
		PublicKeyAlgorithm: x509.ECDSA,
		NotAfter:           time.Now().Add(24 * time.Hour),
		DNSNames:           subject,
		SerialNumber:       serialNumber,
	}
	rawCertificate, err := x509.CreateCertificate(rand.Reader, &certTemplate, &certTemplate, &privateKey.PublicKey, privateKey)
	if err != nil {
		return certificate, err
	}
	// cert, err := x509.ParseCertificate(rawCertificate)
	// log.Println(base64.StdEncoding.EncodeToString(cert.Signature))
	return tls.Certificate{
		Certificate:                  [][]byte{rawCertificate},
		PrivateKey:                   privateKey,
		SupportedSignatureAlgorithms: []tls.SignatureScheme{tls.ECDSAWithP256AndSHA256},
	}, nil
}
