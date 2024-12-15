### Q & A

> 🌐️ English | [中文](Q&A_CN.md)

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

* Prompts: Connection failed. check your internet connection or VPN...
> The problem of computer network, such as setting a proxy, but the proxy service is not started or abnormal, starting or closing the proxy service can generally solve the problem, and if it still does not work, you can try to restart the computer

* Prompts: Is it possible not to display the window after running?
> No, you need to keep a window open to maintain a service for cursor use, but you can minimize it

* Prompts：Too many computers used within the last 24 hours
> Currently, a scheduled task is set. If this happens, the server will automatically change the account, but you need to wait for a period of time. The client will also check every 7 minutes until cursor-vip displays "Encounter a problem? Press Enter to restart cursor to solve the problem" Press Enter to restart at this time

* After promotion, the increase in the number of people promoted is not displayed
> The person being recommended needs to use the complete promotion command to install, and it only takes effect for the first installation of the device, and virtual machines such as VMware will not increase the number of people
