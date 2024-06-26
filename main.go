package main

import (
	"embed"
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

func HelpUpdate(c *menu.CallbackData) {
	fmt.Println("----", c.MenuItem.Label)
}

func main() {
	width, height := robotgo.GetScreenSize()

	// Create an instance of the app structure
	app := NewApp()

	//AppMenu := menu.NewMenu()
	//FileMenu := AppMenu.AddSubmenu("File")
	//FileMenu.AddText("&Open", keys.CmdOrCtrl("o"), HelpUpdate)
	//FileMenu.AddSeparator()
	//FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
	//	runtime.Quit(app.ctx)
	//})

	appOptions := options.App{
		/* 窗口标题栏中显示的文本 */
		Title:         "template",
		Width:         width * 8 / 10,
		Height:        height * 8 / 10,
		DisableResize: false,
		Fullscreen:    false,
		/* true:窗口将没有边框或标题栏 */
		Frameless: true,
		MinWidth:  width * 618 / 1000,
		MinHeight: height * 618 / 1000,
		//MaxWidth: 0,
		//MaxHeight: 0,
		/* 应用程序将被隐藏 */
		StartHidden: false,
		/* 关闭窗口将关闭应用程序。 将此设置为 true 意味着关闭窗口将隐藏窗口。 */
		HideWindowOnClose: false,
		/* 窗口在失去焦点时应保持在其他窗口之上 */
		AlwaysOnTop: false,
		/* 背景颜色, options.NewRGBA(255,0,0,128) - 红色，透明度为 50% */
		BackgroundColour: &options.RGBA{R: 255, G: 0, B: 0, A: 0},
		AssetServer: &assetserver.Options{
			Assets: assets,
			//Handler:    assetsHandler,
			//Middleware: assetsMidldeware,
		},
		/* 应用程序要使用的菜单 */
		Menu:               nil,
		Logger:             nil,
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.ERROR,
		/* 此回调在前端创建之后调用，但在 index.html 加载之前调用。 它提供了应用程序上下文。 */
		OnStartup: app.startup,
		/* 在前端加载完毕 index.html 及其资源后调用此回调。 它提供了应用程序上下文。 */
		OnDomReady: app.domReady,
		/* 在前端被销毁之后，应用程序终止之前，调用此回调。 它提供了应用程序上下文。 */
		OnShutdown: app.shutdown,
		/* 如果设置了此回调，它将在通过单击窗口关闭按钮或调用runtime.Quit即将退出应用程序时被调用.
		返回 true 将导致应用程序继续，false 将继续正常关闭。 返回 true 将导致应用程序继续，false
		将继续正常关闭。 这有助于与用户确认他们希望退出程序。*/
		OnBeforeClose: nil,
		/* 定义需要绑定到前端的方法的结构实例切片 */
		Bind: []interface{}{
			app,
		},
		EnumBind:         nil,
		WindowStartState: options.Normal,
		ErrorFormatter:   func(err error) any { return err.Error() },
		/* 指示用于标识哪些元素可用于拖动窗口的 CSS 属性 */
		CSSDragProperty: "--wails-draggable",
		/* 指示 CSSDragProperty 样式应该具有什么值才能拖动窗口 */
		CSSDragValue:             "drag",
		EnableDefaultContextMenu: false,
		/* 启用针对欺诈内容（例如恶意软件或网络钓鱼尝试）的扫描服务。 这些服务可能会从你的应用中发送信息，比如导航到苹果和微软的云服务的url和其他内容。 */
		EnableFraudulentWebsiteDetection: false,
		SingleInstanceLock:               nil,

		Windows: &windows.Options{
			/* 当使用 alpha 值 0 时，将此设置为 true 将使 webview 背景透明。 这意味着如果您在 CSS 中使用 rgba(0,0,0,0)
			作为 background-color，则主机窗口将显示出来。 通常与 窗口半透明 结合使用以制作看起来冷冰冰的应用程序。 */
			WebviewIsTransparent: true,
			/* 将此设置为 true 将使窗口半透明。 通常与 WebviewIsTransparent 结合使用。 */
			WindowIsTranslucent: true,
			/* 将此设置为 true 将删除标题栏左上角的图标。 */
			DisableWindowIcon:    false,
			IsZoomControlEnabled: false,
			ZoomFactor:           0,
			DisablePinchZoom:     false,
			/* 将此设置为 true 将移除 无边框 模式下的窗口装饰。 这意味着将不会有Aero 阴影 和 圆角显示在窗口上。 请注意，'圆角' 只在 Windows 11 上支持。 */
			DisableFramelessWindowDecorations: false,
			/* 这定义了 WebView2 存储用户数据的路径。 如果为空将使用 %APPDATA%\[BinaryName.exe] */
			WebviewUserDataPath: "",
			/* 这定义了带有 WebView2 可执行文件和库的目录的路径 如果为空，则使用系统中安装的 webview2 */
			WebviewBrowserPath: "",
			Theme:              windows.Dark,
			/* 允许您为浅色和深色模式以及窗口处于活动或非活动状态的 TitleBar、TitleText 和 Border 指定自定义颜色。 */
			CustomTheme:  nil,
			BackdropType: 0,
			Messages:     nil,
			/* ResizeDebounceMS 是调整窗口大小时去抖动 webview2 重绘的时间量。 默认值 (0) 将尽可能快地执行重绘 */
			ResizeDebounceMS: 0,
			/* 如果设置，当 Windows 启动切换到低功耗模式（挂起/休眠）时将调用此函数 */
			OnSuspend: nil,
			/* 如果设置，当 Windows 从低功耗模式（挂起/休眠）恢复时将调用此函数 */
			OnResume: nil,
			/* 设置为 true 将禁用 webview 的 GPU 硬件加速。 */
			WebviewGpuIsDisabled:                false,
			WebviewDisableRendererCodeIntegrity: false,
			EnableSwipeGestures:                 false,
		},

		Mac: nil,

		Linux: nil,

		Experimental: nil,
		Debug:        options.Debug{},
	}

	// Create application with options
	err := wails.Run(&appOptions)
	if err != nil {
		log.Fatal(err)
	}
}
