package main

import (
	"GuardMinecraft/cloudflare"
	"GuardMinecraft/console"
	_const "GuardMinecraft/const"
	"GuardMinecraft/proxy"
	"GuardMinecraft/speedtest"
	"GuardMinecraft/version"
	"fmt"
	"github.com/fatih/color"
	"runtime"
	"strings"
)

var Logo = "   ____                     _   __  __ _                            __ _   \n" +
	"  / ___|_   _  __ _ _ __ __| | |  \\/  (_)_ __   ___  ___ _ __ __ _ / _| |_ \n" +
	" | |  _| | | |/ _` | '__/ _` | | |\\/| | | '_ \\ / _ \\/ __| '__/ _` | |_| __|\n" +
	" | |_| | |_| | (_| | | | (_| | | |  | | | | | |  __/ (__| | | (_| |  _| |_ \n" +
	"  \\____|\\__,_|\\__,_|_|  \\__,_| |_|  |_|_|_| |_|\\___|\\___|_|  \\__,_|_|  \\__|\n"

func main() {
	console.SetTitle(fmt.Sprintf("Guard Minecraft %s (%s)", version.Version, version.CommitHash))
	console.Println(color.HiRedString(Logo))
	color.HiGreen("欢迎使用Minecraft专用代理工具 %s (%s)!\n", version.Version, version.CommitHash)
	color.HiCyan("项目地址: https://github.com/lycorisor/GuardMinecraft")
	color.HiBlack("Build Information: %s, %s/%s\n",
		runtime.Version(), runtime.GOOS, runtime.GOARCH)
	color.HiRed("开始自动注册Cloudflare Warp")
	data, err := cloudflare.Get()
	if err != nil {
		panic(err)
	}

	color.HiRed("注册成功, 获取到如下信息:")
	// print information
	{
		fmt.Println("device_id:", data.Response.ID)
		fmt.Println("token:", data.Response.Token)
		fmt.Println("account_id:", data.Response.Account.ID)
		fmt.Println("account_type:", data.Response.Account.AccountType)
		fmt.Println("license:", data.Response.Account.License)
		fmt.Println("private_key:", data.PrivateKey)
		fmt.Println("public_key:", data.Response.Config.Peers[0].PublicKey)
		fmt.Println("client_id:", data.Response.Config.ClientID)
		fmt.Println("reserved: [", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(data.Reversed)), ","), "[]"), "]")
		fmt.Println("v4:", data.Response.Config.Interface.Addresses.V4)
		fmt.Println("v6:", data.Response.Config.Interface.Addresses.V6)
		fmt.Println("endpoint:", data.Response.Config.Peers[0].Endpoint.Host)
	}

	color.HiRed("开始进行Cloudflare Warp IP优选")

	speedtest.InitRandSeed()
	speedtest.InitHandshakePacket()

	pingData := speedtest.NewWarping().Run().FilterDelay().FilterLossRate()
	color.HiRed("Cloudflare Warp IP优选结束")
	fmt.Println(color.RedString("最优节点: "), color.BlueString(pingData[0].IP.String()))

	proxy.Config = fmt.Sprintf(
		_const.ProxyDefaultConfig,
		pingData[0].IP.IP.String(),
		pingData[0].IP.Port,
		data.PrivateKey,
		data.Response.Config.Peers[0].PublicKey,
		"["+strings.Trim(strings.Join(strings.Fields(fmt.Sprint(data.Reversed)), ","), "[]")+"]",
	)

	color.HiRed("已启动代理程序")
	panic(proxy.Run())
}
