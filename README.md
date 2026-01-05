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
