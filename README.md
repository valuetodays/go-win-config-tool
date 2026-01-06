# go-win-config-tool

prerequirement:
  - go go1.23.4
  - node v22.14.0


go install github.com/wailsapp/wails/v2/cmd/wails@latest

wails init -n go-win-config-tool

wails dev

## goal

- [x] show directories
- [x] show envs
- [ ] show softwares
  + [ ] copy or unzip them
- [ ] show shortcut list

## design

/main.go 
/app.go
/app_reload.go 重新加载配置文件
/app_paths.go 对应 PathsTab 的请求入口
/app_settings.go 对应 SettingsTab 的请求入口
/service/path_service.go 由 /app_paths.go 调用的业务逻辑
/service/xxx_service.go  由 /xxx_paths.go 调用的业务逻辑
/domain/*  存放返回给前端的公共模型
/internal/* 存放内部实现
/config/* 存放配置文件相关