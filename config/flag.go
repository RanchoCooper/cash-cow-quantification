package config

import (
    "flag"
)

var (
    configPathFromFlag string
)

func init() {
    flag.StringVar(&configPathFromFlag, "cf", "./config/config.yaml", "path of config file")
}
