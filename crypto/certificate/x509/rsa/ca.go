package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

// Generate self-signed certificate Root CA cert (ca.pem) and its key (ca-key.pem)
func main() {
	// Generate pub & priv key pair by RSA
	size := 2024
	priv, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}

	// Create CA certificate template
	ca := x509.Certificate{
		IsCA:         true,
		SerialNumber: big.NewInt(1234),
		Subject: pkix.Name{
			Country:      []string{"Japan"},
			Organization: []string{"TCNKSM CA Inc."},
		},

		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(24 * time.Hour),

		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// Create Certificate
	derBytes, err := x509.CreateCertificate(rand.Reader, &ca, &ca, &priv.PublicKey, priv)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}

	certOut, err := os.Create("certs/ca.pem")
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	defer certOut.Close()

	if err := pem.Encode(certOut, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	}); err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}

	keyOut, err := os.OpenFile("certs/ca-key.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	defer keyOut.Close()

	if err := pem.Encode(keyOut, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(priv),
	}); err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
}
