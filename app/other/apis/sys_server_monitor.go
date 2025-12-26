package apis

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v3/net"

	"go-admin/app/other/service/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

var excludeNetInterfaces = []string{
	"lo", "tun", "docker", "veth", "br-", "vmbr", "vnet", "kube",
}

type ServerMonitor struct {
	api.Api
}

// GetHourDiffer 获取相差时间
func GetHourDiffer(startTime, endTime string) int64 {
	t1, err1 := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
	t2, err2 := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
	if err1 != nil || err2 != nil || !t1.Before(t2) {
		return 0
	}
	return (t2.Unix() - t1.Unix()) / 3600
}

// ServerInfo 获取系统信息
// @Summary 获取系统信息
// @Description 获取系统信息
// @Tags 系统监控
// @Success 200 {object} response.Response{data=dto.GetServerMonitorInfoResp} "{"code": 200, "data": [...]}"
// @Router /lotus/api/v1/server/monitor [get]
// @Security Bearer
func (e ServerMonitor) ServerInfo(c *gin.Context) {
	e.Context = c

	osInfo := getOSInfo()
	memInfo := getMemoryInfo()
	swapInfo := getSwapInfo()
	cpuInfo := getCPUInfo()
	diskInfo := getDiskInfo()
	netInfo := getNetworkInfo()

	bootTime, _ := host.BootTime()
	cachedBootTime := time.Unix(int64(bootTime), 0)

	serverInfo := dto.ServerMonitorInfo{
		Code:     200,
		Os:       osInfo,
		Mem:      memInfo,
		Cpu:      cpuInfo,
		Disk:     diskInfo,
		Net:      netInfo,
		Swap:     swapInfo,
		Location: "Aliyun",
		BootTime: GetHourDiffer(cachedBootTime.Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05")),
	}

	e.OK(serverInfo, "获取成功")
}

func getOSInfo() dto.OSInfo {
	sysInfo, _ := host.Info()
	return dto.OSInfo{
		GoOs:         runtime.GOOS,
		Arch:         runtime.GOARCH,
		Mem:          runtime.MemProfileRate,
		Compiler:     runtime.Compiler,
		Version:      runtime.Version(),
		NumGoroutine: runtime.NumGoroutine(),
		ProjectDir:   pkg.GetCurrentPath(),
		HostName:     sysInfo.Hostname,
		Time:         time.Now().Format("2006-01-02 15:04:05"),
	}
}

func getMemoryInfo() dto.MemoryInfo {
	memory, _ := mem.VirtualMemory()
	return dto.MemoryInfo{
		Used:    memory.Used / MB,
		Total:   memory.Total / MB,
		Percent: pkg.Round(memory.UsedPercent, 2),
	}
}

func getSwapInfo() dto.SwapInfo {
	memory, _ := mem.VirtualMemory()
	return dto.SwapInfo{
		Used:  memory.SwapTotal - memory.SwapFree,
		Total: memory.SwapTotal,
	}
}

func getCPUInfo() dto.CPUInfo {
	cpuInfo, _ := cpu.Info()
	percent, _ := cpu.Percent(0, false)
	cpuNum, _ := cpu.Counts(false)
	return dto.CPUInfo{
		CpuInfo: cpuInfo,
		Percent: pkg.Round(percent[0], 2),
		CpuNum:  cpuNum,
	}
}

func getDiskInfo() dto.DiskInfo {
	var diskTotal, diskUsed, diskUsedPercent float64
	diskList := make([]disk.UsageStat, 0)

	diskInfo, err := disk.Partitions(true)
	if err == nil {
		for _, p := range diskInfo {
			diskDetail, err := disk.Usage(p.Mountpoint)
			if err == nil {
				diskDetail.UsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", diskDetail.UsedPercent), 64)
				diskDetail.Total /= MB
				diskDetail.Used /= MB
				diskDetail.Free /= MB
				diskList = append(diskList, *diskDetail)
			}
		}
	}

	d, _ := disk.Usage("/")
	diskTotal = float64(d.Total / GB)
	diskUsed = float64(d.Used / GB)
	diskUsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", d.UsedPercent), 64)

	return dto.DiskInfo{
		Total:   diskTotal,
		Used:    diskUsed,
		Percent: diskUsedPercent,
	}
}

func getNetworkInfo() dto.NetworkInfo {
	netInSpeed, netOutSpeed := trackNetworkSpeed()()
	return dto.NetworkInfo{
		In:  pkg.Round(float64(netInSpeed)/KB, 2),
		Out: pkg.Round(float64(netOutSpeed)/KB, 2),
	}
}

func trackNetworkSpeed() func() (uint64, uint64) {
	var lastIn uint64
	var lastOut uint64
	var lastTime uint64
	var mu sync.Mutex

	// 闭包，捕获外部变量
	return func() (uint64, uint64) {
		var in, out uint64
		mu.Lock()
		defer mu.Unlock()

		nc, err := net.IOCounters(true)
		if err != nil {
			return 0, 0
		}

		for _, v := range nc {
			if isListContainsStr(excludeNetInterfaces, v.Name) {
				continue
			}
			in += v.BytesRecv
			out += v.BytesSent
		}

		now := uint64(time.Now().Unix())
		diff := now - lastTime

		var inSpeed, outSpeed uint64
		if lastTime != 0 && diff > 0 {
			inSpeed = (in - lastIn) / diff
			outSpeed = (out - lastOut) / diff
		}

		// 更新闭包中的状态
		lastIn = in
		lastOut = out
		lastTime = now

		return inSpeed, outSpeed
	}
}

func isListContainsStr(list []string, str string) bool {
	for _, item := range list {
		if strings.Contains(str, item) {
			return true
		}
	}
	return false
}
