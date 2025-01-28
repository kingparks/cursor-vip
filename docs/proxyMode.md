### Proxy mode

> 🌐️ English | [中文](proxyMode_CN.md)

> Notice that trusting certificates is a very dangerous thing. If you set the proxy to the system proxy, the proxy service can capture all network requests on your computer. Therefore, for untrusted programs, never trust certificates and set the proxy to the system proxy.

> The mode 2,3,4 is that cursor-vip starts a proxy service locally on your computer, only the proxy configuration of the cursor client points to the proxy service of cursor-vip, and other software does not use this proxy. We promise not to record any of your network requests.

Proxy mode For the first time after starting, you need to install the trusted certificate. The certificate will be automatically generated after the first start command, and the path is `~/.cursor-vip/`.

#### MacOS:
Command method
```sh
sudo security add-trusted-cert -d -p ssl -p basic -k /Library/Keychains/System.keychain ~/.cursor-vip/i-need-to-trust-ca-cert.pem
```
Graphical method
> Execute open ~/.cursor-vip in the terminal, double-click the i-need-to-trust-ca-cert.pem file, pop up the "Keychain Access" window, select the certificate, search for cursor-vip, double-click cursor-vip, expand trust, select X.509 Basic Policy "Always trust", close the pop-up window, enter the password to confirm, and the certificate is installed.
#### Windows:
Command method (gitBash)
```sh
certutil -addstore root ~/.cursor-vip/ca-cert.cer
```
Command method（cmd）
```sh
certutil -addstore root %USERPROFILE%\.cursor-vip\ca-cert.cer
```
Graphical method
> Search for cer in windows, select the certmgr.msc function, expand Trusted Root Certification Authorities, select Certificates, right-click All Tasks, select Import..., next, enter the %homepath%\.cursor-vip\ca-cert.cer file, next all the way, complete;
#### Linux
> [Ubuntu/Debian](https://askubuntu.com/questions/73287/how-do-i-install-a-root-certificate/94861#94861)、[Fedora](https://docs.fedoraproject.org/en-US/quick-docs/using-shared-system-certificates/#proc_adding-new-certificates)
