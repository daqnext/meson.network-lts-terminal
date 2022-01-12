package commonmsg

type MachineStateBaseMsg struct {
	MacAddr      string  //`json:"mac_addr"`
	MemTotal     uint64  //`json:"mem_total"` // uint: byte
	MemAvailable uint64  //`json:"mem_avail"`
	Version      string  //`json:"version"`
	CpuUsage     float64 //`json:"cpu_usage"`
}

type TerminalStatesMsg struct {
	OS               string     //`json:"os"`
	CPU              string     //`json:"cpu"` // cpu model name
	Port             string     //`json:"port"`
	NetInMbs         [5]float64 //`json:"net_in_mbs"`
	NetOutMbs        [5]float64 //`json:"net_out_mbs"`
	CdnDiskTotal     uint64     //`json:"cdn_disk_total"`
	CdnDiskAvailable uint64     //`json:"cdn_disk_avail"`
	MachineSetupTime string     //`json:"machine_setup_time"`
	SequenceId       int        //`json:"sequence_id"`
	MachineStateBaseMsg
}
