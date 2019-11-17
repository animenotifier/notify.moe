package admin

import (
	"net/http"
	"runtime"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// Get admin page.
func Get(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil || (user.Role != "admin" && user.Role != "editor") {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/")
	}

	// // CPU
	cpuModel := ""
	cpuInfo, err := cpu.Info()

	if err == nil {
		cpuModel = cpuInfo[0].ModelName
	}

	cpuUsage := 0.0
	cpuUsages, err := cpu.Percent(0, false)

	if err == nil {
		cpuUsage = cpuUsages[0]
	}

	// Memory
	memUsage := 0.0
	memTotal := uint64(0)
	memInfo, err := mem.VirtualMemory()

	if err == nil {
		memUsage = memInfo.UsedPercent
		memTotal = memInfo.Total
	}

	// Disk
	diskUsage := 0.0
	diskTotal := uint64(0)
	diskInfo, err := disk.Usage("/")

	if err == nil {
		diskUsage = diskInfo.UsedPercent
		diskTotal = diskInfo.Total
	}

	// GC
	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)

	// Host
	platform, family, platformVersion, _ := host.PlatformInformation()
	kernelVersion, _ := host.KernelVersion()
	kernelVersion = strings.Replace(kernelVersion, "-generic", "", 1)

	return ctx.HTML(components.Admin(
		user,
		platform,
		family,
		platformVersion,
		kernelVersion,
		cpuUsage,
		memUsage,
		diskUsage,
		cpuModel,
		memTotal,
		diskTotal,
		memStats,
	))
}
