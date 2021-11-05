/*
CONS3RT API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: support@cons3rt.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gocons3rt

import (
	"crypto/tls"
	"encoding/pem"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"log"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

func NewTLSClientCertificate(certFile string, certPassword string) (*tls.Config, error) {
	p12, err := ioutil.ReadFile(certFile)
	if err != nil {
		log.Fatal(err)
	}
	blocks, err := pkcs12.ToPEM(p12, certPassword)
	if err != nil {
		panic(err)
	}

	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	// then use PEM data for tls to construct tls certificate:
	cert, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		panic(err)
	}

	config := tls.Config{Certificates: []tls.Certificate{cert}, Renegotiation: tls.RenegotiateOnceAsClient}
	return &config, err
}