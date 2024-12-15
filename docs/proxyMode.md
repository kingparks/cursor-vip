### Proxy mode

> 🌐️ English | [中文](proxyMode_CN.md)

> Notice that trusting certificates is a very dangerous thing. If you set the proxy to the system proxy, the proxy service can capture all network requests on your computer. Therefore, for untrusted programs, never trust certificates and set the proxy to the system proxy.

> The proxy mode is that cursor-vip starts a proxy service locally on your computer, only the proxy configuration of the cursor client points to the proxy service of cursor-vip, and other software does not use this proxy. We promise not to record any of your network requests.

Proxy mode For the first time after starting, you need to install the trusted certificate. The certificate will be automatically generated after the first start command, and the path is `~/.cursor-vip/i-need-to-trust-ca-cert.pem`.
* MacOS: Execute `open ~/.cursor-vip` in the terminal, double-click the i-need-to-trust-ca-cert.pem file, pop up the "Keychain Access" window, select the certificate, search for cursor-vip, double-click cursor-vip, expand trust, select X.509 Basic Policy "Always trust", close the pop-up window, enter the password to confirm, and the certificate is installed.
* Windows: Search for cer in windows, select the `certmgr.msc` function, expand `Trusted Root Certification Authorities`, select `Certificates`, right-click `All Tasks`, select `Import...`, next, enter the `%homepath%\.cursor-vip\i-need-to-trust-ca-cert.pem` file, next all the way, complete;
* Linux: //TODO 
