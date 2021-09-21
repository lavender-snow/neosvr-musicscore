package config

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

type configList struct {
	PortNo    string
	RunMode   string
	CertFile  string
	KeyFile   string
	SecretKey string
}

const (
	configIni = "config.ini"
)

// Config は設定情報を管理します
var Config configList

func init() {
	cfg, err := ini.Load(configIni)
	if err != nil {
		log.Fatalf("iniファイルの読込に失敗しました: %v\n", err.Error())
	}

	secretKey := cfg.Section("session").Key("secret_key").String()

	if secretKey == "" {
		log.Fatalln("iniファイルにsecretKeyが設定されていません")
	}

	portNo := fmt.Sprintf(":%s", cfg.Section("gin").Key("port_no").String())

	if len(portNo) == 0 {
		portNo = ":8080"
		log.Println("ポート番号の設定がされていないため、8080として設定します")
	}

	runMode := cfg.Section("gin").Key("run_mode").String()

	if runMode != "https" && runMode != "http" {
		runMode = "http"
		log.Println("ginの実行モードが設定されていないため、httpとして設定します")
	}

	certFile := cfg.Section("gin").Key("cert_file").String()
	keyFile := cfg.Section("gin").Key("key_file").String()

	if runMode == "https" && (len(certFile) == 0 || len(keyFile) == 0) {
		log.Fatalln("実行モードがhttpsに設定されていますが、SSL証明書の情報が設定されていません")
	}

	Config = configList{
		SecretKey: secretKey,
		PortNo:    portNo,
		RunMode:   runMode,
		CertFile:  certFile,
		KeyFile:   keyFile,
	}
}
