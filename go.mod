module github.com/daqnext/meson.network-lts-terminal

go 1.15

require (
	github.com/daqnext/BGJOB_GO v1.1.3
	github.com/daqnext/LocalLog v0.2.4
	github.com/daqnext/SPR-go v1.1.3
	github.com/daqnext/diskmgr v0.0.0-20211205133327-aedf81e9a43b
	github.com/daqnext/downloadmgr v0.0.0-20211124030451-426bcbdad734
	github.com/daqnext/fastjson v1.0.0
	github.com/daqnext/go-fast-cache v1.0.6
	github.com/daqnext/go-smart-routine v0.1.5
	github.com/daqnext/meson.network-lts-http-server v0.1.0
	github.com/daqnext/utils v0.1.2
	github.com/imroc/req v0.3.2
	github.com/kr/pretty v0.1.0 // indirect
	github.com/labstack/echo/v4 v4.6.1
	github.com/shirou/gopsutil/v3 v3.21.11
	github.com/takama/daemon v1.0.0
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

replace (
	github.com/daqnext/diskmgr => /Users/zhangzhenbo/workspace/go/project/diskmgr
	github.com/daqnext/downloadmgr => /Users/zhangzhenbo/workspace/go/project/downloadmgr
	github.com/daqnext/meson.network-lts-http-server => /Users/zhangzhenbo/workspace/go/project/meson.network-lts-http-server
	github.com/labstack/echo/v4 => /Users/zhangzhenbo/workspace/go/project/echo
)
