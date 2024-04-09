package reportmaker

import (
	"time"

	"github.com/mateuszpiela/pcinfotool/assembly"
	"github.com/mateuszpiela/pcinfotool/hardwareinfo"
	"github.com/mateuszpiela/pcinfotool/osinfo"
	"github.com/mateuszpiela/pcinfotool/servicesinfo"
)

type MainReportStruct struct {
	ReportCreatedAt string `json:"report_createdat"`
	ReportVersion string `json:"report_version"`
	ReportSoftware string `json:"report_software"`
	HardwareReport *hardwareinfo.HwInfoReport `json:"hw_report"`
	OsReport *osinfo.OsReport `json:"os_report"`
	ServicesReport servicesinfo.ServiceReport `json:"win_services"`
}

func GenerateReport() MainReportStruct {
	report := MainReportStruct{
		ReportCreatedAt: getTodayDate(),
		ReportVersion: assembly.Version,
		ReportSoftware: "MP_PCInfoTool",
		HardwareReport: hardwareinfo.GenerateHardwareReport(),
		OsReport: osinfo.GetOSInformation(),
		ServicesReport: servicesinfo.GenerateServicesReport(),
	}
	return report
}

func getTodayDate() string {
	now := time.Now()
	dateString := now.Format("2006-01-02 15:04:05")

	return dateString
}