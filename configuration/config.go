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
	conf := &Config{}
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(conf); err != nil {
		panic(err)
	}
	if conf.Vadek.WorkDir == "" {
		pwd, err := os.Getwd()
		if err != nil {
			panic(errors.Wrap(err, "init config: get current dir"))

		}
		conf.Vadek.WorkDir, _ = filepath.Abs(pwd)
	} else {
		workDir, err := filepath.Abs(conf.Vadek.WorkDir)
		if err != nil {
			panic(err)
		}
		conf.Vadek.WorkDir = workDir
	}
	normalizeDir := func(path *string, subDir string) {
		if *path == "" {
			*path = filepath.Join(conf.Vadek.WorkDir, subDir)
		} else {
			temp, err := filepath.Abs(*path)
			if err != nil {
				panic(err)
			}
			*path = temp
		}
	}
	//TODO:后期记得修改目录问题
	normalizeDir(&conf.Vadek.LogDir, "logs")
	normalizeDir(&conf.Vadek.TemplateDir, "resources/template")
	normalizeDir(&conf.Vadek.AdminResourcesDir, "resources/admin")
	normalizeDir(&conf.Vadek.UploadDir, "/upload")
	normalizeDir(&conf.Vadek.ThemeDir, "resources/template/theme")
	if conf.SQLite3 != nil && conf.SQLite3.Enable {
		normalizeDir(&conf.SQLite3.File, "Vadek.db")
	}
	if !FileIsExisted(conf.Vadek.TemplateDir) {
		panic("template dir: " + conf.Vadek.TemplateDir + " not exist")
	}
	if !FileIsExisted(conf.Vadek.AdminResourcesDir) {
		panic("AdminResourcesDir dir: " + conf.Vadek.AdminResourcesDir + " not exist")
	}
	if !FileIsExisted(conf.Vadek.ThemeDir) {
		panic("Theme dir: " + conf.Vadek.ThemeDir + " not exist")
	}
	initDirectory(conf)
	mode = conf.Vadek.Mode
	return conf
}
func initDirectory(conf *Config) {
	mkdirFunc := func(dir string, err error) error {
		if err == nil {
			if _, err = os.Stat(dir); os.IsNotExist(err) {
				err = os.MkdirAll(dir, os.ModePerm)
			}
		}
		return err
	}
	err := mkdirFunc(conf.Vadek.LogDir, nil)
	err = mkdirFunc(conf.Vadek.UploadDir, err)
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
