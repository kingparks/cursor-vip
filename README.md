# cursor-vip

> 🌐️ English | [中文](README_CN.md)

`cursor-vip` is a tool service for CURSOR smart code editor to enjoy VIP intelligent prompts without logging in.


### Usage

Open the terminal on MacOS/Linux; Open Git Bash on Windows. Then execute the following command to install:
>some computers may report false positives, need to close the antivirus software/computer housekeeper/security protection and then proceed

Method 1: Install via GitHub script
```bash
bash <(curl -Lk https://github.com/kingparks/cursor-vip/releases/download/latest/i.sh) githubReadme
```
Method 2: Install via ghp.ci proxy script
```bash
bash <(curl -Lk https://ghp.ci/https://github.com/kingparks/cursor-vip/releases/download/latest/install.sh) githubReadme
```
Method 3: Manually download the binary file
> download the binary file for the corresponding operating system from the [release](https://github.com/kingparks/cursor-vip/releases) page
 ```shell
# MaxOS/Linux
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

<!--
<details>
<summary>Precautions for using strong proxy mode</summary>

Strong proxy mode For the first time after starting, you need to install the trusted certificate. The certificate will be automatically generated after the first start command, and the path is `~/.cursor-vip/cursor-vip-ca-cert.pem`.
* MacOS: Execute `open ~/.cursor-vip` in the terminal, double-click the cursor-vip-ca-cert.pem file, pop up the "Keychain Access" window, select the certificate, search for cursor-vip, double-click cursor-vip, expand trust, select "Always trust when using this certificate", close the pop-up window, enter the password to confirm, and the certificate is installed.
* Windows: Search for cer in windows, select the `certmgr.msc` function, expand `Trusted Root Certification Authorities`, select `Certificates`, right-click `All Tasks`, select `Import...`, next, enter the `%homepath%\.cursor-vip\cursor-vip-ca-cert.pem` file, next all the way, complete; reopen the browser.
* Linux: //TODO linux currently only supports minimalist mode

</details>
-->

Enjoy cursor-vip:
* Open the CURSOR code editor, click the settings icon in the upper right corner, and you can see that you are already a VIP account.
* Enjoy code prompts and chat functions.

---
### Promotion Guide
By promoting cursor-vip, you can automatically get up to an unlimited amount of free usage time, and the other party can also get 15 days of free (self-registration for 7 days), you can refer to the following methods:
1. After cloning this repository, replace `githubReadme` in the `README` with your DeviceID.
2. Post articles on platforms such as SegmentFault/Stack Overflow/GitHub Discussions/Reddit/Dev.to/Hacker News/DZone/V2EX/Quora/WeChat Official Account/Zhihu/Bilibili/CSDN/, promoting the use of cursor pro in this way, and attach your own promotional installation command or cloned github address.
3. Comment under cursor related articles or videos, for example: "Cooperate with Github's [Your Github Name]/cursor-vip for low-priced or freelancing Pro"
---
### Q & A

* I have two computers at home and at work, and the usage time will not overlap. Can I pay for one and extend the usage period of two devices?
> No, because the cursor official is used to prompt Too many computers used according to the device detection, and the platform currently shares a Pro account for every 8 devices

* Slow request, add requests here...
* Global Rate Limit Hit - Server is Busy...
* Unable to reach Anthropic...
> cursor official normal busy period advanced model queue, change small model can solve, or try to delete the cache:
> Mac: rm ~/Library/Application\ Support/Cursor
> Windows: rd -r %UserProfile%\AppData\Roaming\Cursor\Cache

* Using composer prompts: We're currently receiving a large number of slow requests and could not queue yours
> There is indeed this problem, currently no solution, at this time, use chat

* Connection failed. check your internet connection or VPN...
> The problem of computer network, such as setting a proxy, but the proxy service is not started or abnormal, starting or closing the proxy service can generally solve the problem, and if it still does not work, you can try to restart the computer

* Is it possible not to display the window after running?
> No, you need to keep a window open to maintain a service for cursor use, but you can minimize it

* After promotion, the increase in the number of people promoted is not displayed
> The person being recommended needs to use the complete promotion command to install, and it only takes effect for the first installation of the device, and virtual machines such as VMware will not increase the number of people

---
### Star History
<a href="https://star-history.com/#kingparks/cursor-vip&Date">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=kingparks/cursor-vip&type=Date&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=kingparks/cursor-vip&type=Date" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=kingparks/cursor-vip&type=Date" />
 </picture>
</a>