### 提问与回答

> 🌐️ 中文 | [English](Q&A.md)

* 我有家庭和公司两个电脑，使用的时间不会重叠，能够付费一个，增加两个设备的使用期限吗？
> 不能，因为 cursor 官方是按照设备检测用来提示 Too many computers used , 目前该平台为每8个设备共享一个Pro账号

* Slow request, add requests here...
* Global Rate Limit Hit - Server is Busy...
* Unable to reach Anthropic...
> cursor 官方正常的繁忙期高级模型排队，换小模型可解决，或者尝试下删除缓存：
> Mac: rm ~/Library/Application\ Support/Cursor
> Windows: rd -r %UserProfile%\AppData\Roaming\Cursor\Cache

* 使用 composer 提示：We're currently receiving a large number of slow requests and could not queue yours
> 存在这个问题，目前没有解决方案，这时候先将就用 chat

* 提示：Connection failed. check your internet connection or VPN...
> 电脑网络的问题，例如设置了代理，但是代理服务没有启动或者异常，启动或关闭代理服务一般可解决，还不行可以尝试重启电脑

* 提示：Too many computers used within the last 24 hours
> 目前设置了定时任务，如果出现此类情况，服务端会自动换账号，但是需要等待一段时间，客户端段也会每7分钟检测一次，直到 cursor-vip 显示出 “遇到问题？按回车键将重启 cursor 解决问题” 这时候按回车键重启即可 

* 运行后可以不显示窗口吗
> 不行，需要开着窗口维持一个服务供 cursor 使用，但是可以最小化

* 推广后没有显示推广人数的增加
> 被推荐者需要用完整的推广命令来安装，且只对该设备的首次安装生效，另外例如 VMware 等虚拟机不会增加人数