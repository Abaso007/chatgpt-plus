package utils

import (
	"bytes"
	logger2 "chatplus/logger"
	"chatplus/types"
	"github.com/BurntSushi/toml"
	"net/http"
	"os"
)

func NewDefaultConfig() *types.Config {
	return &types.Config{
		Title:        "Chat-Plus AI 助手",
		ConsoleTitle: "Chat-Plus 控制台",
		Listen:       "0.0.0.0:5678",
		ProxyURL:     make([]string, 0),
		ImgURL:       types.ImgURL{},
		Manager:      types.Manager{Username: "admin", Password: "admin123"},
		AccessKey:    RandString(42),

		Session: types.Session{
			SecretKey: RandString(64),
			Name:      "CHAT_SESSION_ID",
			Domain:    "",
			Path:      "/",
			MaxAge:    86400,
			Secure:    true,
			HttpOnly:  false,
			SameSite:  http.SameSiteLaxMode,
		},
		Chat: types.Chat{
			ApiURL:                "https://api.openai.com/v1/chat/completions",
			ApiKeys:               make([]types.APIKey, 0),
			Model:                 "gpt-3.5-turbo",
			MaxTokens:             1024,
			Temperature:           0.9,
			EnableContext:         true,
			ChatContextExpireTime: 3600,
		},
	}
}

var logger = logger2.GetLogger()

func LoadConfig(configFile string) (*types.Config, error) {
	var config *types.Config
	_, err := os.Stat(configFile)
	if err != nil {
		logger.Info("creating new config file: ", configFile)
		config = NewDefaultConfig()
		// save config
		err := SaveConfig(config, configFile)
		if err != nil {
			return nil, err
		}

		return config, nil
	}
	_, err = toml.DecodeFile(configFile, &config)
	if err != nil {
		return nil, err
	}

	return config, err
}

func SaveConfig(config *types.Config, configFile string) error {
	buf := new(bytes.Buffer)
	encoder := toml.NewEncoder(buf)
	if err := encoder.Encode(&config); err != nil {
		return err
	}

	return os.WriteFile(configFile, buf.Bytes(), 0644)
}
