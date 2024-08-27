# cursor-vip

[English](./README.md)

`cursor-vip` 是一个用于 CURSOR 智能代码编辑器 无需账号登录即可享受VIP智能提示的工具服务。


### 使用
安装 cursor-vip：
```bash
bash <(curl https://mirror.ghproxy.com/https://github.com/kingparks/cursor-vip/releases/download/latest/install.sh) githubReadme
```

启动 cursor-vip：
```bash
cursor-vip
```
首次启动后需安装信任证书，证书会在首次启动命令后自动生成，路径为 `~/.cursor-vip/cursor-vip-ca-cert.pem`。
* MacOS: 在终端执行 `open ~/.cursor-vip`，双击 cursor-vip-ca-cert.pem 文件，弹出“钥匙串访问”窗口，选择证书，搜索 cursor-vip，双击 cursor-vip，展开信任，选择使用此证书时“始终信任”，关闭弹窗，输入密码确认，证书安装完成。
* Windows: 在windows搜索输入 cer,选择`管理用户证书`功能，展开`受信任的根证书颁发机构`，选中`证书`，右键`所有任务`，选择`导入`，下一步，输入`%homepath%\.cursor-vip\cursor-vip-ca-cert.pem`文件，一直下一步，完成; 重新打开浏览器。
* Linux: //TODO

享受 cursor-vip：
* 打开 CURSOR 智能代码编辑器，点击右上角设置图标，退出然后登录即可


