package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/quic-go/quic-go"
	"log"
	"math/big"
	"tank_war/server/cmd/game/initialize"
	"tank_war/server/cmd/game/room"
)

func main() {

	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitConfig()

	listener, err := quic.ListenAddr("localhost:8888", generateTLSConfig(), nil)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	clientId := 1
	for {
		conn, err := listener.Accept(context.Background())
		if err != nil {
			klog.Fatalf("listener.Accept failed: %v")
		}
		klog.Infof("client connected,ip:", conn.RemoteAddr().String())
		stream, err := conn.OpenStream()
		if err != nil {
			log.Println(err)
		}

		//write client id to client
		_, err = stream.Write([]byte{byte(clientId)})
		log.Println("client id:", clientId)

		//read ok from client
		data := make([]byte, 1)
		_, err = stream.Read(data)
		if err != nil {
			log.Println(err)
		}

		//check ok
		if data[0] != byte(clientId) {
			log.Println("client id not match")
			stream.Close()
			continue
		}
		room.NewClient(stream, int32(clientId))

		clientId++
	}
}

func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"tank_war"},
	}
}
