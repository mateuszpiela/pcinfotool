package reportmaker

import (
	"github.com/mateuszpiela/pcinfotool/hardwareinfo"
	"github.com/mateuszpiela/pcinfotool/osinfo"
)

type ReportMaker struct {
	ReportCreatedAt string
	ReportVersion string
	ReportSoftware string
	HardwareReport *hardwareinfo.HwInfoReport
	OsReport *osinfo.OsReport
}

func GenerateReport() {

}