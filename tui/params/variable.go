package params

import (
	"github.com/unknwon/i18n"
	"io"
	"os"
)

var Mode int64           // 1模式1 2模式2
var CursorVersion string // cursor版本号
var Lang string
var ExclusiveToken string
var M3c string
var Promotion string
var DeviceID string
var MachineID string
var ColorOut io.Writer
var Sigs chan os.Signal
var SigCountDown chan int
var Trr *Tr

type Tr struct {
	i18n.Locale
}
