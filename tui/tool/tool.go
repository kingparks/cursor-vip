package tool

import (
	"crypto/md5"
	"fmt"
	"github.com/denisbrodbeck/machineid"
	"github.com/gofrs/flock"
	"github.com/kingparks/cursor-vip/tui/params"
	"github.com/tidwall/gjson"
	"howett.net/plist"
	"net"
	"os"
	"os/exec"
	"path"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// 获取本地语言
func GetLocale() (langRes, locRes string) {
	osHost := runtime.GOOS
	langRes = "en"
	locRes = "US"
	switch osHost {
	case "windows":
		// Exec powershell Get-Culture on Windows.
		cmd := exec.Command("powershell", "Get-Culture | select -exp Name")
		output, err := cmd.Output()
		if err == nil {
			langLocRaw := strings.TrimSpace(string(output))
			langLoc := strings.Split(langLocRaw, "-")
			langRes = langLoc[0]
			langRes = strings.Split(langRes, "-")[0]
			locRes = langLoc[1]
			return
		}
	case "darwin":
		// Exec shell Get-Culture on MacOS.
		cmd := exec.Command("sh", "osascript -e 'user locale of (get system info)'")
		output, err := cmd.Output()
		if err == nil {
			langLocRaw := strings.TrimSpace(string(output))
			langLoc := strings.Split(langLocRaw, "_")
			langRes = langLoc[0]
			langRes = strings.Split(langRes, "-")[0]
			if len(langLoc) == 1 {
				return
			}
			locRes = langLoc[1]
			return
		}
		plistB, err := os.ReadFile(os.Getenv("HOME") + "/Library/Preferences/.GlobalPreferences.plist")
		if err != nil {
			fmt.Println(err)
			_, _ = fmt.Scanln()
			panic(err)
		}
		var a map[string]interface{}
		_, err = plist.Unmarshal(plistB, &a)
		if err != nil {
			fmt.Println(err)
			_, _ = fmt.Scanln()
			panic(err)
		}
		langLocRaw, ok := a["AppleLocale"].(string)
		if !ok {
			//fmt.Println("AppleLocale key not found or not a string")
			return
		}
		langLoc := strings.Split(langLocRaw, "_")
		langRes = langLoc[0]
		langRes = strings.Split(langRes, "-")[0]
		if len(langLoc) == 1 {
			return
		}
		locRes = langLoc[1]
		return
	case "linux":
		envlang, ok := os.LookupEnv("LANG")
		if ok {
			langLocRaw := strings.TrimSpace(envlang)
			langLocRaw = strings.Split(envlang, ".")[0]
			langLoc := strings.Split(langLocRaw, "_")
			langRes = langLoc[0]
			langRes = strings.Split(langRes, "-")[0]
			if len(langLoc) == 1 {
				return
			}
			locRes = langLoc[1]
			return
		}
	}
	if langRes == "" {
		langLocRaw := os.Getenv("LC_CTYPE")
		langLocRaw = strings.Split(langLocRaw, ".")[0]
		langLoc := strings.Split(langLocRaw, "_")
		langRes = langLoc[0]
		langRes = strings.Split(langRes, "-")[0]
		if len(langLoc) == 1 {
			return
		}
		locRes = langLoc[1]
		return
	}
	return
}

// 获取配置
func GetConfig() (lang, promotion string, mode int64) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}
	b, _ := os.ReadFile(homeDir + "/.cursor-viprc")
	s := string(b)
	lang = gjson.Get(s, "lang").String()
	mode = gjson.Get(s, "mode").Int()
	promotion = gjson.Get(s, "promotion").String()
	if lang == "" {
		lang, _ = GetLocale()
	}
	if mode == 0 {
		mode = 2
	}
	if params.IsOnlyMod2 {
		mode = 2
	}
	return
}

// 设置配置
func SetConfig(lang string, mode int64) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}
	config := fmt.Sprintf(`{"lang":"%s","promotion":"%s","mode":%d}`, lang, params.Promotion, mode)
	_ = os.WriteFile(homeDir+"/.cursor-viprc", []byte(config), 0644)
}

func GetMac_241018() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}

	var macError []string
	for _, inter := range interfaces {
		hardwareAddr := inter.HardwareAddr.String()
		if hardwareAddr == "" {
			continue
		}
		macError = append(macError, inter.Name+": "+hardwareAddr)
	}
	sort.Strings(macError)
	return strings.Join(macError, "\n")
}
func GetDeviceID() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}

	var macAddress, bluetoothAddress []string
	virtualMacPrefixes := []string{
		"00:05:69", "00:0C:29", "00:1C:14", "00:50:56", // VMware
		"00:15:5D",             // Hyper-V
		"08:00:27", "0A:00:27", // VirtualBox
	}

	for _, inter := range interfaces {
		hardwareAddr := inter.HardwareAddr.String()
		if hardwareAddr == "" {
			continue
		}

		isVirtual := false
		for _, prefix := range virtualMacPrefixes {
			if strings.HasPrefix(hardwareAddr, strings.ToLower(prefix)) {
				isVirtual = true
				break
			}
		}
		if isVirtual {
			continue
		}

		switch {
		case strings.HasPrefix(inter.Name, "en"), strings.HasPrefix(inter.Name, "Ethernet"), strings.HasPrefix(inter.Name, "以太网"):
			macAddress = append(macAddress, hardwareAddr)
		case strings.HasPrefix(inter.Name, "Bluetooth"), strings.HasSuffix(inter.Name, "Bluetooth"), strings.HasPrefix(inter.Name, "蓝牙网络连接"):
			bluetoothAddress = append(bluetoothAddress, hardwareAddr)
		}
	}

	if len(macAddress) == 0 {
		macAddress = append(macAddress, bluetoothAddress...)
		if len(macAddress) == 0 {
			//_, _ = fmt.Fprintf(params.ColorOut, red, "no mac address found,Please contact customer service")
			//_, _ = fmt.Scanln()
			//return macErrorStr
			return GetMachineID()
		}
	}
	sort.Strings(macAddress)
	return fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(macAddress, ","))))
}
func GetMachineID() string {
	id, err := machineid.ID()
	if err != nil {
		return err.Error()
	}
	id = strings.ToLower(id)
	id = strings.ReplaceAll(id, "-", "")
	return id
}

func CountDown(seconds int) {
	go func(seconds int) {
		countdown := seconds // Countdown in seconds
		for countdown >= 0 {
			days := countdown / (24 * 3600)
			hours := (countdown % (24 * 3600)) / 3600
			minutes := (countdown % 3600) / 60
			seconds := countdown % 60

			_, _ = fmt.Fprintf(params.ColorOut, "\r%dd %dh %dm %ds", days, hours, minutes, seconds)
			time.Sleep(1 * time.Second)
			countdown--
		}
		// 发送退出信号
		params.Sigs <- syscall.SIGTERM
	}(seconds)
}

// EnsureSingleInstance 确保程序的单实例运行
func EnsureSingleInstance(name string) (*flock.Flock, string, error) {
	homeDir, _ := os.UserHomeDir()
	_ = os.MkdirAll(path.Join(homeDir, "."+name), 0755)
	lockFilePath := path.Join(homeDir, "."+name, name+".lock")
	pidFilePath := path.Join(homeDir, "."+name, name+".pid")
	// 创建文件锁
	lock := flock.New(lockFilePath)

	// 尝试获取锁
	locked, err := lock.TryLock()
	if err != nil {
		return nil, pidFilePath, fmt.Errorf("无法创建锁文件: %w", err)
	}

	if !locked {
		// 读取 PID 文件
		pidData, err := os.ReadFile(pidFilePath)
		if err != nil {
			return nil, pidFilePath, fmt.Errorf("无法读取 PID 文件: %w", err)
		}

		pid, err := strconv.Atoi(string(pidData))
		if err != nil {
			return nil, pidFilePath, fmt.Errorf("PID 文件格式错误: %w", err)
		}

		// 尝试终止旧进程
		if err := terminateProcess(pid); err != nil {
			return nil, pidFilePath, err
		}

		// 等待旧实例退出
		time.Sleep(1 * time.Second)

		// 再次尝试获取锁
		locked, err = lock.TryLock()
		if err != nil {
			return nil, pidFilePath, fmt.Errorf("无法重新获取锁文件: %w", err)
		}
		if !locked {
			return nil, pidFilePath, fmt.Errorf("旧实例仍在运行，无法获取锁")
		}
	}

	// 记录当前进程的 PID
	pid := os.Getpid()
	err = os.WriteFile(pidFilePath, []byte(strconv.Itoa(pid)), 0644)
	if err != nil {
		return nil, pidFilePath, fmt.Errorf("无法写入 PID 文件: %w", err)
	}
	return lock, pidFilePath, nil
}

// terminateProcess 在 Windows 上终止指定 PID 的进程
func terminateProcess(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("无法找到进程 PID=%d: %w", pid, err)
	}
	// 调用 Kill 方法终止进程
	if err := process.Kill(); err != nil {
		return fmt.Errorf("无法终止进程 PID=%d: %w", pid, err)
	}
	return nil
}

func OpenNewTerminal() {
	execPath, err := os.Executable()
	if err != nil {
		return
	}
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "cmd", "/k", execPath)
	case "darwin":
		cmd = exec.Command("open", "-n", "-F", execPath)
	case "linux":
		cmd = exec.Command("x-terminal-emulator", "-e", "bash", "-c", execPath)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Start()
	if err != nil {
		return
	}
	params.Sigs <- syscall.SIGTERM
}
