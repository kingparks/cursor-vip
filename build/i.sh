set -e
URLS=("https://github.com/kingparks/cursor-vip/releases/download/latest/")
url=${URLS[0]}
lc_type=$(echo $LC_CTYPE | cut -c 1-2)
if [ -z $lc_type ] || [ "$lc_type" = "UT" ]; then
  lc_type=$(echo $LANG | cut -c 1-2)
fi

if [ "$lc_type" = "zh" ]; then
  echo "正在安装..."
else
  echo "Installing..."
fi

for url0 in ${URLS[@]}; do
  if curl -Is --connect-timeout 4 "$url0" | grep -q "HTTP/1.1 404"; then
    url=$url0
    break
  fi
done

os_name=$(uname -s | tr '[:upper:]' '[:lower:]')
if [[ $os_name == *"mingw"* ]]; then
  os_name="windows"
fi
raw_hw_name=$(uname -m)
case "$raw_hw_name" in
"amd64")
  hw_name="amd64"
  ;;
"x86_64")
  hw_name="amd64"
  ;;
"arm64")
  hw_name="arm64"
  ;;
"aarch64")
  hw_name="arm64"
  ;;
"i686")
  hw_name="386"
  ;;
"armv7l")
  hw_name="arm"
  ;;
*)
  echo "Unsupported hardware: $raw_hw_name"
  exit 1
  ;;
esac

if [ "$lc_type" = "zh" ]; then
  echo "当前系统为 ${os_name} ${hw_name}"
else
  echo "Current system is ${os_name} ${hw_name}"
fi

if [ ! -z $1 ]; then
   echo "{\"promotion\":\"$1\"}" >~/.cursor-viprc
fi

# 如果是mac或者linux系统
if [[ $os_name == "darwin" || $os_name == "linux" ]]; then
  if [ "$lc_type" = "zh" ]; then
    echo "请输入开机密码"
  else
    echo "Please enter the boot password"
  fi;
  # 安装
  sudo mkdir -p /usr/local/bin;
  sudo rm -f /usr/local/bin/cursor-vip;
  # 停掉正在运行的cursor-vip
  pkill cursor-vip > /dev/null || true
  sudo curl -Lko /usr/local/bin/cursor-vip ${url}/cursor-vip_${os_name}_${hw_name}
  sudo chmod +x /usr/local/bin/cursor-vip

  # linux 系统，接收用户输入 cursor.AppImage 文件路径，进行 --appimage-extract 操作生成 squashfs-root 目录；
  # 进入 squashfs-root 目录，执行 sudo chown -R root:root usr/share/cursor/chrome-sandbox ,执行 sudo chmod 4755 usr/share/cursor/chrome-sandbox
  # 回退到 squashfs-root 目录上一级，将 squashfs-root 移动到 ~/cursor
  if [[ $os_name == "linux" ]]; then
    if [ "$lc_type" = "zh" ]; then
      echo "请输入 cursor-xxx.AppImage 文件路径"
    else
      echo "Please enter the cursor-xxx.AppImage file path"
    fi;
    read -p "cursor-xxx.AppImage file path: " appimage_path
    if [ ! -f "$appimage_path" ]; then
      if [ "$lc_type" = "zh" ]; then
        echo "文件不存在"
      else
        echo "File does not exist"
      fi;
      exit 1
    fi
    appimage_path=$(realpath $appimage_path)
    appimage_name=$(basename $appimage_path)
    appimage_dir=$(dirname $appimage_path)
    cd $appimage_dir;
    chmod +x $appimage_name;
    killall -9 cursor > /dev/null 2>&1 || true;
    sudo rm -rf ./squashfs-root;
    ./$appimage_name --appimage-extract;
    cd squashfs-root;
    sudo chown -R root:root usr/share/cursor/chrome-sandbox > /dev/null 2>&1 || true;
    sudo chmod 4755 usr/share/cursor/chrome-sandbox > /dev/null 2>&1 || true;
    cd ..;
    rm -rf ~/cursor;
    mv squashfs-root ~/cursor;
    if [ "$lc_type" = "zh" ]; then
      echo "跳过登录后关闭 cursor"
    else
      echo "Skip logging in then close the cursor"
    fi;
    ~/cursor/AppRun;
  fi

  if [ "$lc_type" = "zh" ]; then
    echo "安装完成！自动运行；下次可直接输入 cursor-vip 并回车来运行程序"
  else
    echo "Installation completed! Automatically run; you can run the program by entering cursor-vip and pressing Enter next time"
  fi;

  echo ""
  cursor-vip
fi;
# 如果是windows系统
if [[ $os_name == "windows" ]]; then
  # 停掉正在运行cursor-vip
  if [ -n "$MSYSTEM" ]; then
    # Git Bash 环境
    taskkill -f -im cursor-vip.exe > /dev/null 2>&1 || true
  else
    # CMD 或 PowerShell 环境
    taskkill -f -im cursor-vip.exe > nul 2>&1 || true
  fi
  # 判断如果有powershell，则通过powershell来获取桌面路径,否则通过cmd来获取桌面路径
  desktop_dir="${USERPROFILE}/Desktop"
  if command -v powershell > /dev/null; then
    desktop_dir=$(powershell -Command "[Console]::OutputEncoding = [System.Text.Encoding]::UTF8; [Environment]::GetFolderPath('Desktop')")
  else
      if [ -d "${USERPROFILE}/Desktop" ]; then
        desktop_dir="${USERPROFILE}/Desktop"
      else
        desktop_dir="${USERPROFILE}/OneDrive/Desktop"
      fi
  fi
  # 安装
  curl -Lko ${desktop_dir}/cursor-vip.exe ${url}/cursor-vip_${os_name}_${hw_name}.exe
  if [ "$lc_type" = "zh" ]; then
    echo "安装完成！自动运行; 下次可直接输入 ./cursor-vip.exe 并回车来运行程序"
    echo "运行后如果360等杀毒软件误报木马，添加信任后，重新输入./cursor-vip.exe 并回车来运行程序"
  else
    echo "Installation completed! Automatically run; you can run the program by entering ./cursor-vip.exe and press Enter next time"
    echo "After running, if 360 antivirus software reports a Trojan horse, add trust, and then re-enter ./cursor-vip.exe and press Enter to run the program"
  fi

  echo ""
  chmod +x ${desktop_dir}/cursor-vip.exe
#  powershell -Command "Start-Process -FilePath '${desktop_dir}/cursor-vip.exe' -Verb RunAs"
  ${desktop_dir}/cursor-vip.exe
fi
