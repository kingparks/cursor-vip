# cursor-vip

> 🌐️ English | [中文](README_CN.md)




### Usage
Open the terminal on MacOS/Linux; Open Git Bash on Windows(If you don't have GitBash, please choose method 3 to download manually, wsl is not supported). Then execute the following command to install:
>some computers may report false positives, need to close the antivirus software/computer housekeeper/security protection and then proceed

Method 1: Install via GitHub script
```bash
bash <(curl -Lk https://github.com/kingparks/cursor-vip/releases/download/latest/i.sh) githubReadme
```
Method 2: Install via Gitee script
```bash
bash <(curl -Lk https://gitee.com/kingparks/cursor-vip/releases/download/latest/ic.sh) githubReadme
```
Method 3: Manually download the binary file
> download the binary file for the corresponding operating system from the [release](https://github.com/kingparks/cursor-vip/releases) page
 ```shell
# MaxOS/Linux (linux binary installation also requires additional execution of the content of i.sh)
sudo mv cursor-vip_xx_xxx /usr/local/bin/cursor-vip;
chmod +x /usr/local/bin/cursor-vip;
cursor-vip githubReadme;
# Windows 
# double click cursor-vip_xx_xxx.exe
```

Launch cursor-vip：
```bash
# MaxOS/Linux
cursor-vip
# Windows
# double click cursor-vip_xx_xxx.exe
```

[Additional settings for proxy mode](docs/proxyMode.md)

Enjoy cursor-vip:
* Open the CURSOR code editor, click the settings icon in the upper right corner, and you can see that you are already a VIP account.
* Enjoy code prompts and chat functions.

---

```md
# Change configuration manually
# Configuration file is located at ~/.cursor-viprc is a json file
# Configuration items:
# lang - language - String, optional values: en English; zh Chinese; nl Dutch; ru Russian; hu Hungarian; tr Turkish; es Spanish;
# mode - mode - int, optional values: 1 normal mode; 2 proxy mode;
```

### Promotion Guide
[Promotion Guide](docs/promotion.md)

---

### How to use custom models
[Gemini 2.0](docs/models-gemini-2.0.md)

---

### Q & A
[Q & A](docs/Q&A.md)

---
### Star History
<a href="https://star-history.com/#kingparks/cursor-vip&Date">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=kingparks/cursor-vip&type=Date&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=kingparks/cursor-vip&type=Date" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=kingparks/cursor-vip&type=Date" />
 </picture>
</a>
