package admin

import (
	"time"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// Get admin page.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil || (user.Role != "admin" && user.Role != "editor") {
		return ctx.Redirect("/")
	}

	// CPU
	cpuUsage := 0.0
	cpuUsages, err := cpu.Percent(1*time.Second, false)

	if err == nil {
		cpuUsage = cpuUsages[0]
	}

	// Memory
	memUsage := 0.0
	memInfo, _ := mem.VirtualMemory()

	if err == nil {
		memUsage = memInfo.UsedPercent
	}

	// Disk
	diskUsage := 0.0
	diskInfo, err := disk.Usage("/")

	if err == nil {
		diskUsage = diskInfo.UsedPercent
	}

	// Host
	platform, family, platformVersion, _ := host.PlatformInformation()
	kernelVersion, err := host.KernelVersion()

	return ctx.HTML(components.Admin(user, cpuUsage, memUsage, diskUsage, platform, family, platformVersion, kernelVersion))
}

func average(floatSlice []float64) float64 {
	if len(floatSlice) == 0 {
		return arn.DefaultAverageRating
	}

	var sum float64

	for _, value := range floatSlice {
		sum += value
	}

	return sum / float64(len(floatSlice))
}
