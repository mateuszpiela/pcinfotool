//go:build windows
// +build windows

package servicesinfo

import (
	"syscall"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc/mgr"
)

func GenerateServicesReport() (ServiceReport) {
	host, err := syscall.UTF16PtrFromString("localhost")
	if err != nil {
		panic(err)
	}

	handle, err := windows.OpenSCManager(host, nil, windows.SC_MANAGER_ENUMERATE_SERVICE)
	if err != nil {
		panic(err)
	}
	defer windows.CloseServiceHandle(handle)

	scMgr := &mgr.Mgr{Handle: handle}

	servicesName, err := scMgr.ListServices()

	if err != nil {
		panic(err)
	}

	services := make([]ServiceStruct, 0, len(servicesName))

	for _, srvName := range servicesName {
		srvNameUINT, err := syscall.UTF16PtrFromString(srvName)
		if err != nil {
			panic(err)
		}

		serviceHandle, err := windows.OpenService(handle, srvNameUINT, windows.SERVICE_QUERY_STATUS|windows.SERVICE_QUERY_CONFIG)
		if err != nil {
			panic(err)
		}

		serviceStatus := windows.SERVICE_STATUS{}
		err = windows.QueryServiceStatus(serviceHandle, &serviceStatus)
		if err != nil {
			panic(err)
		}

		serviceConfig := windows.QUERY_SERVICE_CONFIG{}
		neededForSC := uint32(0)
		windows.QueryServiceConfig(serviceHandle, &serviceConfig, neededForSC, &neededForSC)
		if err != nil {
			panic(err)
		}
		serviceConfigSize := neededForSC
		err = windows.QueryServiceConfig(serviceHandle, &serviceConfig, serviceConfigSize, &neededForSC)
		if err != nil {
			panic(err)
		}

		service := ServiceStruct{
			Name:      srvName,
			State:     serviceStateToString(serviceStatus.CurrentState),
			StartType: serviceStartTypeToString(serviceConfig.StartType),
		}

		windows.CloseServiceHandle(serviceHandle)

		services = appendToArray(services, service)
	}

	serviceReport := ServiceReport{Services: services} 
	return serviceReport
}

func appendToArray(services []ServiceStruct, service ServiceStruct) []ServiceStruct {
    // Check if the slice is full
    if len(services) >= cap(services) {
        // If it's full, return the same slice
        return services
    }
    // Otherwise, create a new slice with one more element
    newServices := make([]ServiceStruct, len(services)+1)
    // Copy the elements up to the last index
    copy(newServices, services)
    // Insert the new service at the end
    newServices[len(services)] = service
    return newServices
}

func serviceStartTypeToString(StartType uint32) string {
	var startStr string

	switch start := StartType; start {
	case windows.SERVICE_BOOT_START:
		startStr = "BOOT"
	case windows.SERVICE_SYSTEM_START:
		startStr = "SYSTEM"
	case windows.SERVICE_AUTO_START:
		startStr = "AUTO"
	case windows.SERVICE_DEMAND_START:
		startStr = "DEMAND"
	case windows.SERVICE_DISABLED:
		startStr = "DISABLED"
	}

	return startStr
}

func serviceStateToString(CurrentState uint32) string {
	var stateStr string

	switch status := CurrentState; status {
	case windows.SERVICE_RUNNING:
		stateStr = "RUNNING"
	case windows.SERVICE_START_PENDING:
		stateStr = "STARTING"
	case windows.SERVICE_STOP_PENDING:
		stateStr = "STOPPING"
	case windows.SERVICE_STOPPED:
		stateStr = "STOPPED"
	case windows.SERVICE_CONTINUE_PENDING:
		stateStr = "CONTINUE"
	case windows.SERVICE_PAUSE_PENDING:
		stateStr = "PAUSING"
	case windows.SERVICE_PAUSED:
		stateStr = "PAUSED"
	}

	return stateStr
}

