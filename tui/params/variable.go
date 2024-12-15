package params

import (
	"github.com/unknwon/i18n"
	"io"
	"os"
)

var Mode int64 // 1普通模式 2代理模式
var Lang string
var Promotion string
var DeviceID string
var MachineID string
var ColorOut io.Writer
var Sigs chan os.Signal
var Trr *Tr

type Tr struct {
	i18n.Locale
}
