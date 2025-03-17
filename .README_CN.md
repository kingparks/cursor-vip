# cursor-vip

> 🌐️ 中文 | [English](README.md)

⚠️ 如果您的地区支持官方 cursor 支付，在您资金允许的情况下,请支持官方购买，cursor 是一个不错的编辑器。

⚠️ 本项目仅为了支付不方便的地区提供一个解决方案，目前每8个设备共享一个 Pro 账号，Pro 账号都是我们花钱购买的，不要滥用，谢谢！

⚠️ 使用本项目所产生花费及收益比与官方一样，本项目没有更便宜，只是因为8个人分摊了，所以感觉花费更少。但是稳定性和速度都不如官方。

⚠️ 目前的模式稳定性排行为：模式3 > 模式2 > 模式4。

⚠️ 我们后期计划与 cursor 官方合作，为能够促进世界文明进步的项目提供支持。


### 使用方式

在 MacOS/Linux 中，请打开终端；在 Windows 中，请打开 Git Bash(如果没有GitBash,请选择方式3手动下载，不支持wsl)。然后执行以下命令来安装：
> 部分电脑可能会误报毒，需要关闭杀毒软件/电脑管家/安全防护再进行

方式1：通过 Gitee 脚本
```bash
bash <(curl -Lk https://gitee.com/kingparks/cursor-vip/releases/download/latest/ic.sh) githubReadme
```
方式2：通过 GitHub 脚本
```bash
bash <(curl -Lk https://github.com/kingparks/cursor-vip/releases/download/latest/i.sh) githubReadme
```
方式3：手动下载二进制文件
> 从 [release](https://github.com/kingparks/cursor-vip/releases) 页下载对应操作系统的二进制文件
```shell
# MaxOS/Linux (linux 二进制安装还需额外执行 i.sh 的内容)
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

[模式2的额外设置](docs/proxyMode_CN.md)

享受 cursor-vip：
* 打开 CURSOR 代码编辑器，点击右上角设置图标，可看到已是VIP账号。
* 尽情享受代码提示和聊天功能。
---

```md
# 手动修改配置
# 配置文件位于 ~/.cursor-viprc 是个json文件
# 配置项：
# lang - 语言 - String ，可选值：en英语 zh中文 nl荷兰语 ru俄语 hu匈牙利语 tr土耳其语 es西班牙语
# mode - 模式 - int ，可选值：1模式1 2模式2
```

### 推广指南
[推广指南](docs/promotion_CN.md)

---

### 如何使用自定义模型
[Gemini 2.0](docs/models-gemini-2.0_CN.md)

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