package main

import (
	"context"
	"fmt"
	"github.com/Deansquirrel/goMonitorV6/common"
	"github.com/Deansquirrel/goMonitorV6/global"
	"github.com/Deansquirrel/goMonitorV6/object"
	myService "github.com/Deansquirrel/goMonitorV6/service"
	"github.com/kardianos/service"
	"os"
	"time"
)

import log "github.com/Deansquirrel/goToolLog"

//初始化
func init() {
	global.Args = &object.ProgramArgs{}
	global.SysConfig = &object.SystemConfig{}

	global.Ctx, global.Cancel = context.WithCancel(context.Background())
}

func main() {
	//解析命令行参数
	{
		global.Args.Definition()
		global.Args.Parse()
		err := global.Args.Check()
		if err != nil {
			fmt.Print(err.Error())
			log.Error(err.Error())
			return
		}
		common.UpdateParams()
	}
	//加载系统参数
	{
		common.LoadSysConfig()
		common.RefreshSysConfig()
	}
	//安装、卸载或运行程序
	{
		svcConfig := &service.Config{
			Name:        global.SysConfig.Service.Name,
			DisplayName: global.SysConfig.Service.DisplayName,
			Description: global.SysConfig.Service.Description,
		}
		prg := &program{}
		s, err := service.New(prg, svcConfig)
		if err != nil {
			log.Error("定义服务配置时遇到错误：" + err.Error())
			return
		}

		if global.Args.IsInstall {
			err = s.Install()
			if err != nil {
				log.Error("安装为服务时遇到错误：" + err.Error())
			} else {
				fmt.Println(fmt.Sprintf("服务 %s 安装成功", global.SysConfig.Service.Name))
			}
			return
		}
		if global.Args.IsUninstall {
			err = s.Uninstall()
			if err != nil {
				log.Error("卸载服务时遇到错误：" + err.Error())
			} else {
				fmt.Println(fmt.Sprintf("服务 %s 卸载成功", global.SysConfig.Service.Name))
			}
			return
		}
		_ = s.Run()
	}
}

type program struct {
	isRunning bool
}

func (p *program) Start(s service.Service) error {
	if p.isRunning {
		return nil
	}
	log.Warn("Service Starting")
	defer log.Warn("Service Start Complete")
	err := p.run()
	if err != nil {
		log.Error(fmt.Sprintf("服务启动时遇到错误：%s", err.Error()))
	} else {
		p.isRunning = true
	}
	go func() {
		select {
		case <-global.Ctx.Done():
			err := p.Stop(s)
			if err != nil {
				fmt.Println(err.Error())
			}
			time.Sleep(time.Second * 3)
			os.Exit(0)
		}
	}()
	return err
}

func (p *program) run() error {
	//服务所执行的代码
	{
		return myService.Start()
	}
}

func (p *program) Stop(s service.Service) error {
	if !p.isRunning {
		return nil
	}
	p.isRunning = false
	log.Warn("Service Stopping")
	defer log.Warn("Service Stopped")
	{
		//TODO 停止服务时清理内容
	}
	return nil
}
