# cursor-vip

[中文](./README_CN.md)

`cursor-vip` is a tool service for CURSOR smart code editor to enjoy VIP intelligent prompts without logging in.


### Usage
Install cursor-vip:
```bash
bash <(curl https://mirror.ghproxy.com/https://github.com/kingparks/cursor-vip/releases/download/latest/install.sh) githubReadme
```

Launch cursor-vip:
```bash
cursor-vip
```
For the first time after starting, you need to install the trusted certificate. The certificate will be automatically generated after the first start command, and the path is `~/.cursor-vip/cursor-vip-ca-cert.pem`.
* MacOS: Execute `open ~/.cursor-vip` in the terminal, double-click the cursor-vip-ca-cert.pem file, pop up the "Keychain Access" window, select the certificate, search for cursor-vip, double-click cursor-vip, expand trust, select "Always trust when using this certificate", close the pop-up window, enter the password to confirm, and the certificate is installed.
* Windows: Search for cer in windows, select the `certmgr.msc` function, expand `Trusted Root Certification Authorities`, select `Certificates`, right-click `All Tasks`, select `Import`, next, enter the `%homepath%\.cursor-vip\cursor-vip-ca-cert.pem` file, next all the way, complete; reopen the browser.
* Linux: //TODO

Enjoy cursor-vip:
* Open the CURSOR code editor, click the settings icon in the upper right corner, exit and log in to enjoy


