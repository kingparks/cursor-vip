# cursor-vip

[中文](./README_CN.md)

`cursor-vip` is a tool service for CURSOR smart code editor to enjoy VIP intelligent prompts without logging in.


### Usage
Install cursor-vip (download directly from GitHub):
```bash
bash <(curl https://github.com/kingparks/cursor-vip/releases/download/latest/i.sh) githubReadme
```
Install cursor-vip (through ghp.ci proxy):
```bash
bash <(curl https://ghp.ci/https://github.com/kingparks/cursor-vip/releases/download/latest/install.sh) githubReadme
```
Install cursor-vip (manually):
> download the binary file for the corresponding operating system from the [release](https://github.com/kingparks/cursor-vip/releases) page

Launch cursor-vip:
```bash
cursor-vip
```
<details>
  <summary>Precautions for using strong proxy mode</summary>

Strong proxy mode For the first time after starting, you need to install the trusted certificate. The certificate will be automatically generated after the first start command, and the path is `~/.cursor-vip/cursor-vip-ca-cert.pem`.
* MacOS: Execute `open ~/.cursor-vip` in the terminal, double-click the cursor-vip-ca-cert.pem file, pop up the "Keychain Access" window, select the certificate, search for cursor-vip, double-click cursor-vip, expand trust, select "Always trust when using this certificate", close the pop-up window, enter the password to confirm, and the certificate is installed.
* Windows: Search for cer in windows, select the `certmgr.msc` function, expand `Trusted Root Certification Authorities`, select `Certificates`, right-click `All Tasks`, select `Import...`, next, enter the `%homepath%\.cursor-vip\cursor-vip-ca-cert.pem` file, next all the way, complete; reopen the browser.
* Linux: //TODO linux currently only supports minimalist mode

</details>

Enjoy cursor-vip:
* Open the CURSOR code editor, click the settings icon in the upper right corner, and you can see that you are already a VIP account.
* Enjoy code prompts and chat functions.

### Star History
![Star History Chart](https://api.star-history.com/svg?repos=kingparks/cursor-vip&type=Date)
