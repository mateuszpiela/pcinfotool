package hardwareinfo

import (
	"github.com/jaypipes/ghw"
)

type HwInfoReport struct {
	Product *ghw.ProductInfo `json:"product"`
	Baseboard *ghw.BaseboardInfo `json:"baseboard"`
	Bios *ghw.BIOSInfo `json:"bios"`
	Cpu *ghw.CPUInfo `json:"cpu"`
	Gpu *ghw.GPUInfo `json:"gpu"`
	Mem *ghw.MemoryInfo `json:"mem"`
	BlockStorage *ghw.BlockInfo `json:"blockstorage"`
	Network *ghw.NetworkInfo `json:"network"`
	Pci *ghw.PCIInfo `json:"pci"`
}


func GenerateHardwareReport() *HwInfoReport {
	report := HwInfoReport{}
	report.Product = getProduct()
	report.Baseboard = getBaseboard()
	report.Bios = getBios()
	report.Cpu = getCpu()
	report.Gpu = getGPU()
	report.Mem = getMem()
	report.BlockStorage = getBS()
	report.Network = getNetwork()
	report.Pci = getPCI()

	return &report
}

