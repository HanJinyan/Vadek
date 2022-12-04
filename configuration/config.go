package configuration

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func NewConfiguration() *Config {
	var configFile string
	flag.StringVar(&configFile, "config", "", "")
	flag.Parse()
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType("yaml")
	if configFile != "" {
		viper.SetConfigType(configFile)
	} else {
		viper.AddConfigPath("./")
		viper.SetConfigName("config")
	}
	config := &Config{}
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}
	if config.Vadek.WorkDir == "" {
		pwd, err := os.Getwd()
		if err != nil {
			panic(errors.Wrap(err, "init config: get current dir"))

		}
		config.Vadek.WorkDir, _ = filepath.Abs(pwd)
	} else {
		workDir, err := filepath.Abs(config.Vadek.WorkDir)
		if err != nil {
			panic(err)
		}
		config.Vadek.WorkDir = workDir
	}
	normalizeDir := func(path *string, subDir string) {
		if *path == "" {
			*path = filepath.Join(config.Vadek.WorkDir, subDir)
		} else {
			temp, err := filepath.Abs(*path)
			if err != nil {
				panic(err)
			}
			*path = temp
		}
	}
	//TODO:后期记得修改目录问题
	normalizeDir(&config.Vadek.LogDir, "log")
	normalizeDir(&config.Vadek.TemplateDir, "resources/template")
	normalizeDir(&config.Vadek.AdminResourcesDir, "resources/admin")
	normalizeDir(&config.Vadek.UploadDir, "/upload")
	normalizeDir(&config.Vadek.ThemeDir, "resources/template/theme")
	if config.SQlite3 != nil && config.SQlite3.Enable {
		normalizeDir(&config.SQlite3.File, "Vadek.db")
	}
	if !FileIsExisted(config.Vadek.TemplateDir) {
		panic("template dir: " + config.Vadek.TemplateDir + " not exist")
	}
	if !FileIsExisted(config.Vadek.AdminResourcesDir) {
		panic("AdminResourcesDir dir: " + config.Vadek.AdminResourcesDir + " not exist")
	}
	if !FileIsExisted(config.Vadek.ThemeDir) {
		panic("Theme dir: " + config.Vadek.ThemeDir + " not exist")
	}
	initDirectory(config)
	mode = config.Vadek.Mode
	return config
}
func initDirectory(config *Config) {
	mkdirFunc := func(dir string, err error) error {
		if err == nil {
			if _, err = os.Stat(dir); os.IsNotExist(err) {
				err = os.MkdirAll(dir, os.ModePerm)
			}
		}
		return err
	}
	err := mkdirFunc(config.Vadek.LogDir, nil)
	err = mkdirFunc(config.Vadek.UploadDir, err)
	if err != nil {
		panic(fmt.Errorf("initDirectory err=%v", err))
	}
}

var mode string

func IsDev() bool {
	return mode == "development"
}
func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		existed = false
	}
	return existed
}
