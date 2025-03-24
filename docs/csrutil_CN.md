### Mac 权限问题 解决方案

> 🌐️ 中文 | [English](freeTrialsSolve.md)

> Mac 用户可以临时关闭SIP，再执行 cursor-vip 即可解决这个问题，解决后可再次开启SIP


#### MacOS Apple Silicon
开机时长按开机键，直到出现「设置」后松开，点击「选项」按钮继续，从上方的菜单栏点击「实用工具」选择「终端」，输入「csrutil disable」后回车，点击菜单栏  标志，选择「重新启动」

#### MacOS Intel CPU
开机时立即在键盘上按住 Command ⌘ + R，直到看到 Apple 标志或旋转的地球时松开，从上方的菜单栏点击「实用工具」选择「终端」，输入「csrutil disable」后回车，点击菜单栏  标志，选择「重新启动」


> 其他操作系统理论上直接运行 cursor-vip 即可，如果不行，可以联系我远程看看是什么问题
