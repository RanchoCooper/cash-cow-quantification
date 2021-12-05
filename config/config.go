package config

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "path/filepath"

    "gopkg.in/yaml.v3"

    "quants/util"
)

const (
    configFilePath        = "/config.yaml"
    privateConfigFilePath = "/config.private.yaml"
)

type binanceConfig struct {
    Key    string `yaml:"key"`
    Secret string `yaml:"secret"`
}

type dingDingConfig struct {
    AccessToken string `yaml:"access_token"`
}

type logConfig struct {
    LogSavePath string `yaml:"log_save_path"`
    LogFileName string `yaml:"log_file_name"`
    LogFileExt  string `yaml:"log_file_ext"`
}

type mysqlConfig struct {
    User         string `yaml:"user"`
    Password     string `yaml:"password"`
    Host         string `yaml:"host"`
    Port         int    `yaml:"port"`
    Database     string `yaml:"database"`
    MaxIdleConns int    `yaml:"max_idle_conns"`
    MaxOpenConns int    `yaml:"max_open_conns"`
    MaxLifeTime  string `yaml:"max_life_time"`
    MaxIdleTime  string `yaml:"max_idle_time"`
    CharSet      string `yaml:"char_set"`
    ParseTime    bool   `yaml:"parse_time"`
    TimeZone     string `yaml:"time_zone"`
}

type config struct {
    Log      *logConfig      `yaml:"log"`
    Binance  *binanceConfig  `yaml:"binance"`
    DingDing *dingDingConfig `yaml:"dingding"`
    MySQL    *mysqlConfig    `yaml:"mysql"`
}

var Config = &config{}

func readYamlConfig(configPath string) {
    yamlFile, err := filepath.Abs(configPath)
    if err != nil {
        log.Fatalf("invalid config file path, err: %v", err)
    }
    content, err := ioutil.ReadFile(yamlFile)
    if err != nil {
        log.Printf("read config file fail, err: %v", err)
    }
    err = yaml.Unmarshal(content, Config)
    if err != nil {
        log.Printf("config file unmarshal fail, err: %v", err)
    }
}

func init() {
    configPath := util.GetCurrentPath()

    readYamlConfig(configPath + configFilePath)

    // read private sensitive configs
    if Config.Binance.Key == "" || Config.Binance.Secret == "" {
        // read private config
        readYamlConfig(configPath + privateConfigFilePath)
    }
    bf, _ := json.MarshalIndent(Config, "", "    ")
    fmt.Printf("Config:\n%s\n", string(bf))
}
