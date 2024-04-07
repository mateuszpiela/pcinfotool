package reportmaker

import (
	"github.com/mateuszpiela/pcinfotool/hardwareinfo"
)

type ReportMaker struct {
	ReportVersion string
	ReportSoftware string
	HardwareReport *hardwareinfo.HwInfoReport
}

func GenerateReport() {

}