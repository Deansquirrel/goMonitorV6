package common

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/Deansquirrel/goMonitorV6/global"
	"github.com/Deansquirrel/goMonitorV6/object"
	"github.com/Deansquirrel/goToolCommon"
)

import log "github.com/Deansquirrel/goToolLog"

const SysConfigFile = "config.toml"

func UpdateParams() {
	if global.Args.LogStdOut {
		log.StdOut = true
	}
}

//加载系统配置
func LoadSysConfig() {
	path, err := goToolCommon.GetCurrPath()
	if err != nil {
		log.Error("获取程序所在路径失败")
		return
	}
	fileFullPath := path + "\\" + SysConfigFile
	b, err := goToolCommon.PathExists(fileFullPath)
	if err != nil {
		log.Error(fmt.Sprintf("检查配置文件是否存在时遇到错误：%s，Path：%s", err.Error(), fileFullPath))
		return
	}
	if b {
		var c object.SystemConfig
		_, err := toml.DecodeFile(fileFullPath, &c)
		if err != nil {
			log.Error(fmt.Sprintf("加载配置文件时遇到错误：%s，Paht：%s", err.Error(), fileFullPath))
			return
		}
		c.FormatConfig()
		global.SysConfig = &c
		//global.IsConfigExists = true
	} else {
		log.Warn("未找到配置文件 %s")
	}
}

//刷新系统配置
func RefreshSysConfig() {
	global.SysConfig.FormatConfig()

	setLogLevel(global.SysConfig.Total.LogLevel)
	log.StdOut = global.SysConfig.Total.StdOut || global.Args.LogStdOut
}

//设置日志级别
func setLogLevel(logLevel string) {
	switch logLevel {
	case "debug":
		log.Level = log.LevelDebug
		return
	case "info":
		log.Level = log.LevelInfo
		return
	case "warn":
		log.Level = log.LevelWarn
		return
	case "error":
		log.Level = log.LevelError
		return
	default:
		log.Level = log.LevelWarn
	}
}
