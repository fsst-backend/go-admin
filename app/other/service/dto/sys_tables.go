package dto

type SysTableSearch struct {
	TBName       string `form:"tableName" search:"type:exact;column:table_name;table:table_name"`
	TableComment string `form:"tableComment" search:"type:icontains;column:table_comment;table:table_comment"`
}

type OSInfo struct {
	GoOs         string `json:"goOs"`
	Arch         string `json:"arch"`
	Mem          int    `json:"mem"`
	Compiler     string `json:"compiler"`
	Version      string `json:"version"`
	NumGoroutine int    `json:"numGoroutine"`
	Ip           string `json:"ip"`
	ProjectDir   string `json:"projectDir"`
	HostName     string `json:"hostName"`
	Time         string `json:"time"`
}

type MemoryInfo struct {
	Used    uint64  `json:"used"`
	Total   uint64  `json:"total"`
	Percent float64 `json:"percent"`
}

type SwapInfo struct {
	Used  uint64 `json:"used"`
	Total uint64 `json:"total"`
}

type CPUInfo struct {
	CpuInfo interface{} `json:"cpuInfo"`
	Percent float64     `json:"percent"`
	CpuNum  int         `json:"cpuNum"`
}

type DiskInfo struct {
	Total   float64 `json:"total"`
	Used    float64 `json:"used"`
	Percent float64 `json:"percent"`
}

type NetworkInfo struct {
	In  float64 `json:"in"`
	Out float64 `json:"out"`
}

type ServerMonitorInfo struct {
	Code     int         `json:"code"`
	Os       OSInfo      `json:"os"`
	Mem      MemoryInfo  `json:"mem"`
	Cpu      CPUInfo     `json:"cpu"`
	Disk     DiskInfo    `json:"disk"`
	Net      NetworkInfo `json:"net"`
	Swap     SwapInfo    `json:"swap"`
	Location string      `json:"location"`
	BootTime int64       `json:"bootTime"`
}

type GetServerMonitorInfoResp struct {
	ServerMonitorInfo `json:"data"`
}
