package models

type Tables struct {
	Systems []Systems `json:"systems,omitempty"`
	SysCtl []SysCtl `json:"sysctl,omitempty"`
	Meminfo []Meminfo `json:"meminfo,omitempty"`
	Mounts []Mounts `json:"mounts,omitempty"`
	SwapFs []SwapFs `json:"swapfs,omitempty"`
}

type SysCtl struct {
	ID             int    `json:"id,omitempty"`
	Machine_ID     string `json:"machine_id,omitempty"`
	Setting  string `json:"setting,omitempty"`
	Value string `json:"value,omitempty"`
	Scan_Timestamp string `json:"scan_timestamp,omitempty"`
}
type Mounts struct {
	ID             int    `json:"id,omitempty"`
	Machine_ID     string `json:"machine_id,omitempty"`
	MountPoint  string `json:"mount_point,omitempty"`
	MountSource  string `json:"mount_source,omitempty"`
	FsType string `json:"fs_type,omitempty"`
	FsMntOpts string `json:"fs_mntops,omitempty"`
	FsFreq string `json:"fs_greq,omitempty"`
	FsPassno string `json:"fs_passno,omitempty"`
	Scan_Timestamp string `json:"scan_timestamp,omitempty"`
}
type Meminfo struct {
	ID             int    `json:"id,omitempty"`
	Machine_ID     string `json:"machine_id,omitempty"`
	Setting  string `json:"setting,omitempty"`
	Unit int `json:"unit,omitempty"`
	UnitSymbol  string `json:"unit_symbol,omitempty"`
	Scan_Timestamp string `json:"scan_timestamp,omitempty"`
}
type SwapFs struct {
	ID             int    `json:"id,omitempty"`
	Machine_ID     string `json:"machine_id,omitempty"`
	SwapDevice  string `json:"swap_device,omitempty"`
	SwapType  string `json:"swap_type,omitempty"`
	Size int `json:"size,omitempty"`
	Used int `json:"used,omitempty"`
	Priority  string `json:"priority,omitempty"`
	Scan_Timestamp string `json:"scan_timestamp,omitempty"`
}
type Systems struct {
	ID             int    `json:"id,omitempty"`
	Hostname       string `json:"hostname,omitempty"`
	Domain         string `json:"domain,omitempty"`
	Device_Type    string `json:"device_type,omitempty"`
	Machine_ID     string `json:"machine_id,omitempty"`
	Device_Status  string `json:"device_status,omitempty"`
	Scan_Timestamp string `json:"scan_timestamp,omitempty"`
}
