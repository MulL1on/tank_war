package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"encoding/pem"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/quic-go/quic-go"
	"google.golang.org/protobuf/proto"
	"math/big"
	"net"
	"strconv"
	"tank_war/server/cmd/game/config"
	"tank_war/server/cmd/game/handler"
	pb "tank_war/server/cmd/game/handler/pb/quic"
	"tank_war/server/cmd/game/initialize"
)

func main() {

	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitRdb()
	initialize.InitFlag()
	initialize.InitRegistry()

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	listener, err := quic.ListenAddr(net.JoinHostPort("127.0.0.1", strconv.Itoa(8888)), generateTLSConfig(), nil)
	if err != nil {
		klog.Fatalf("quic.ListenAddr failed: %v", err)
	}
	klog.Infof("server is listening on: %v", listener.Addr())
	defer listener.Close()

	for {
		conn, err := listener.Accept(context.Background())
		if err != nil {
			klog.Fatalf("listener.Accept failed: %v")
		}
		klog.Infof("client connected,ip:", conn.RemoteAddr().String())
		stream, err := conn.OpenStream()
		if err != nil {
			klog.Infof("conn.OpenStream failed: %v", err)
		}
		//write client id to client
		_, err = stream.Write([]byte{byte(1)})

		data := make([]byte, 1024)

		// read data length
		//获取加入房间请求
		_, err = stream.Read(data)
		if err != nil {
			klog.Infof("read err: %v")
			continue
		}
		buffer := bytes.NewBuffer(data)

		var length int32
		if err = binary.Read(buffer, binary.BigEndian, &length); err != nil {
			klog.Infof("read data length error", err)
		}

		// read data content
		data = make([]byte, length)
		if err = binary.Read(buffer, binary.BigEndian, &data); err != nil {
			klog.Infof("read data content error", err)
		}

		//解析请求
		var msg = &pb.JoinRoomReq{}
		err = proto.Unmarshal(data, msg)
		if err != nil {
			klog.Infof("unmarshal err: %v")
			continue
		}
		handler.NewClient(stream, msg.UserId, msg)
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
