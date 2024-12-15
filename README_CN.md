# cursor-vip

> 🌐️ 中文 | [English](README.md)

`cursor-vip` 是一个用于 CURSOR 智能代码编辑器 无需账号登录即可享受VIP智能提示的工具服务。


### 使用方式

在 MacOS/Linux 中，请打开终端；在 Windows 中，请打开 Git Bash。然后执行以下命令来安装：
> 部分电脑可能会误报毒，需要关闭杀毒软件/电脑管家/安全防护再进行

方式1：通过 ghp.ci 代理脚本
```bash
bash <(curl -Lk https://ghp.ci/https://github.com/kingparks/cursor-vip/releases/download/latest/install.sh) githubReadme
```
方式2：通过 GitHub 脚本
```bash
bash <(curl -Lk https://github.com/kingparks/cursor-vip/releases/download/latest/i.sh) githubReadme
```
方式3：手动下载二进制文件
> 从 [release](https://github.com/kingparks/cursor-vip/releases) 页下载对应操作系统的二进制文件
```shell
# MaxOS/Linux
sudo mv cursor-vip_xx_xxx /usr/local/bin/cursor-vip;
chmod +x /usr/local/bin/cursor-vip;
cursor-vip githubReadme;
# Windows
# 双击 cursor-vip_xx_xxx.exe
```

启动 cursor-vip：
```bash
# MaxOS/Linux
cursor-vip
# Windows
# 双击 cursor-vip_xx_xxx.exe
```

[代理模式的额外设置](docs/proxyMode_CN.md)

享受 cursor-vip：
* 打开 CURSOR 代码编辑器，点击右上角设置图标，可看到已是VIP账号。
* 尽情享受代码提示和聊天功能。
---

### 推广指南
[推广指南](docs/promotion_CN.md)

---

### 提问与回答
[提问与回答](docs/Q&A_CN.md)

---
### Star History
<a href="https://star-history.com/#kingparks/cursor-vip&Date">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=kingparks/cursor-vip&type=Date&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=kingparks/cursor-vip&type=Date" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=kingparks/cursor-vip&type=Date" />
 </picture>
</a>