package helper

import (
	"os/exec"
	"strings"
)

func GetUserHardwareSerial() string {
	cmd := exec.Command("wmic", "cpu", "get", "ProcessorId")

	cpuInfo, _ := cmd.CombinedOutput()

	splitExtraCpuInfo := strings.Split(string(cpuInfo), "ProcessorId")

	cpuSerial := strings.TrimSpace(splitExtraCpuInfo[1])

	return cpuSerial
}
