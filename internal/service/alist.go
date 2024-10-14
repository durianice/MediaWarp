package service

import (
	"MediaWarp/internal/config"
	"MediaWarp/internal/service/alist"

	"MediaWarp/pkg"
	"sync"
)

var (
	alistSeverMap sync.Map
	cfg           *config.ConfigManager
)

func init() {
	cfg = config.GetConfig()
	if cfg.AlistStrm.Enable {
		for _, alist := range cfg.AlistStrm.List {
			registerAlistServer(alist.ADDR, alist.Username, alist.Password)
		}
	}
}

// 注册Alist服务器
//
// 将Alist服务器注册到全局Map中
func registerAlistServer(addr string, username string, password string) {
	alistServer := alist.New(addr, username, password)
	alistServer.Init()
	alistSeverMap.Store(alistServer.GetEndpoint(), alistServer)
}

// 获取Alist服务器
//
// 从全局Map中获取Alist服务器
// 若未找到则抛出panic
func GetAlistServer(addr string) *alist.AlistServer {
	endpoint := pkg.GetEndpoint(addr)
	if server, ok := alistSeverMap.Load(endpoint); ok {
		return server.(*alist.AlistServer)
	}
	panic("Alist服务器：" + endpoint + " 未注册")
}