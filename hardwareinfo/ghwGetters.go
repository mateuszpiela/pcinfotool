package hardwareinfo

import (
	"fmt"

	"github.com/jaypipes/ghw"
)

func getCpu() *ghw.CPUInfo {
	cpu, err := ghw.CPU()

	if err != nil {
		fmt.Printf("Cannot get info about CPU: %v", err)
	}

	return cpu
}

func getMem() *ghw.MemoryInfo {
	mem, err := ghw.Memory()

	if err != nil {
		fmt.Printf("Cannot get info about Memory: %v", err)
	}

	return mem
}

func getBS() *ghw.BlockInfo {
	bs, err := ghw.Block()

	if err != nil {
		fmt.Printf("Cannot get info about Block Storage: %v", err)
	}

	return bs
}

func getNetwork() *ghw.NetworkInfo {
	nic, err := ghw.Network()

	if err != nil {
		fmt.Printf("Cannot get info about Network: %v", err)
	}

	return nic
}

func getPCI() *ghw.PCIInfo {
	pci, err := ghw.PCI()

	if err != nil {
		fmt.Printf("Cannot get info about PCI: %v", err)
	}

	return pci
}

func getGPU() *ghw.GPUInfo {
	gpu, err := ghw.GPU()

	if err != nil {
		fmt.Printf("Cannot get info about GPU: %v", err)
	}

	return gpu
}

func getBios() *ghw.BIOSInfo {
	bios, err := ghw.BIOS()

	if err != nil {
		fmt.Printf("Cannot get info about BIOS: %v", err)
	}

	return bios
}

func getBaseboard() *ghw.BaseboardInfo {
	baseboard, err := ghw.Baseboard()

	if err != nil {
		fmt.Printf("Cannot get info about Baseboard: %v", err)
	}

	return baseboard
}

func getProduct() *ghw.ProductInfo {
	product, err := ghw.Product()

	if err != nil {
		fmt.Printf("Cannot get info about Product: %v", err)
	}

	return product
}