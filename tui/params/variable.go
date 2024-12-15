package params

import (
	"github.com/unknwon/i18n"
	"os"
)

var Mode int64 // 1普通模式 2代理模式
var Lang string
var DeviceID string
var MachineID string
var Sigs chan os.Signal
var Trr *Tr

type Tr struct {
	i18n.Locale
}
