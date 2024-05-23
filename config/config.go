package config

import (
	"encoding/json"
	"os"
	"time"
	dlog "web/db_logger"
	"web/logger"
)

var Configure Configuration

type Configuration struct {
	AppSetting     App       `json:"app"`
	PostgreCfg     Postgre   `json:"postgre_cfg"`
	Log            dlog.Conf `json:"log"`
	ServerSetting  Server    `json:"server"`
	MerkleSetting  Merkle    `json:"merkle"`
	RuntimeSetting Runtime   `json:"runtime"`
}

type Postgre struct {
	Conf map[string]string `json:"conf"`
}

type Server struct {
	RunMode         string        `json:"run_mode"`
	HttpPort        int32         `json:"http_port"`
	ReadTimeout     time.Duration `json:"read_timeout"`
	WriteTimeout    time.Duration `json:"write_timeout"`
	ShutDownTimeout time.Duration `json:"shut_down_timeout"`
}

type App struct {
	ExpireTime      int64  `json:"expire_time"`
	RuntimeRootPath string `json:"runtime_rootPath"`
	LogSavePath     string `json:"log_save_path"`
	LogSaveName     string `json:"log_save_name"`
	LogFileExt      string `json:"log_file_ext"`
	LogLevel        string `json:"log_level"`
}

type Merkle struct {
	RemoteSource string `json:"remote_source"`
	RemotePath   string `json:"remote_path"`
	FileSource   string `json:"file_source"`
	FilePath     string `json:"file_path"`
	FileExt      string `json:"file_ext"`
}

type Runtime struct {
	RuntimePath string `json:"runtime_path"`
	RuntimeFile string `json:"runtime_file"`
}

func InitConfig(path string) {
	logger.Debugf(`Init config running.file path:[%s]`, path)
	file, err := os.Open(path)
	if err != nil {
		logger.Errorf("failed to load config file.[err=%v]", err)
		panic("load config file failed ")
	}

	defer file.Close()
	stat, _ := file.Stat()

	bytes := make([]byte, stat.Size())

	_, err = file.Read(bytes)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &Configure)

	if err != nil {
		logger.Errorf("failed to unmarshal config json")
		panic("wrong config json")
	}

	logger.Infof("Init config success")
}
