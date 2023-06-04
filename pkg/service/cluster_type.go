package service

type Resource struct {
	ID         string  `jsont:"id"`
	Type       string  `json:"type"`
	Content    string  `json:",omitempty"`
	CPU        float64 `json:",omitempty"`
	Disk       uint64  `json:",omitempty"` // documented as string but this is an int
	HAstate    string  `json:",omitempty"`
	Level      string  `json:",omitempty"`
	MaxCPU     uint64  `json:",omitempty"`
	MaxDisk    uint64  `json:",omitempty"`
	MaxMem     uint64  `json:",omitempty"`
	Mem        uint64  `json:",omitempty"` // documented as string but this is an int
	Name       string  `json:",omitempty"`
	Node       string  `json:",omitempty"`
	PluginType string  `json:",omitempty"`
	Pool       string  `json:",omitempty"`
	Status     string  `json:",omitempty"`
	Storage    string  `json:",omitempty"`
	Uptime     uint64  `json:",omitempty"`
}

type FirewallSecurityGroup struct {
	Group   string          `json:"group,omitempty"`
	Comment string          `json:"comment,omitempty"`
	Rules   []*FirewallRule `json:"rules,omitempty"`
}
type FirewallRule struct {
	Type     string `json:"type,omitempty"`
	Action   string `json:"action,omitempty"`
	Pos      int    `json:"pos,omitempty"`
	Comment  string `json:"comment,omitempty"`
	Dest     string `json:"dest,omitempty"`
	Dport    string `json:"dport,omitempty"`
	Enable   int    `json:"enable,omitempty"`
	IcmpType string `json:"icmp_type,omitempty"`
	Iface    string `json:"iface,omitempty"`
	Log      string `json:"log,omitempty"`
	Macro    string `json:"macro,omitempty"`
	Proto    string `json:"proto,omitempty"`
	Source   string `json:"source,omitempty"`
	Sport    string `json:"sport,omitempty"`
}

func (r *FirewallRule) IsEnable() bool {
	return 1 == r.Enable
}

type FirewallNodeOption struct {
	Enable                           bool   `json:"enable,omitempty"`
	LogLevelIn                       string `json:"log_level_in,omitempty"`
	LogLevelOut                      string `json:"log_level_out,omitempty"`
	LogNfConntrack                   bool   `json:"log_nf_conntrack,omitempty"`
	Ntp                              bool   `json:"ntp,omitempty"`
	NFConntrackAllowInvalid          bool   `json:"nf_conntrack_allow_invalid,omitempty"`
	NFConntrackMax                   int    `json:"nf_conntrack_max,omitempty"`
	NFConntrackTCPTimeoutEstablished int    `json:"nf_conntrack_tcp_timeout_established,omitempty"`
	NFConntrackTCPTimeoutSynRecv     int    `json:"nf_conntrack_tcp_timeout_syn_recv,omitempty"`
	Nosmurfs                         bool   `json:"nosmurfs,omitempty"`
	ProtectionSynflood               bool   `json:"protection_synflood,omitempty"`
	ProtectionSynfloodBurst          int    `json:"protection_synflood_burst,omitempty"`
	ProtectionSynfloodRate           int    `json:"protection_synflood_rate,omitempty"`
	SmurfLogLevel                    string `json:"smurf_log_level,omitempty"`
	TCPFlagsLogLevel                 string `json:"tcp_flags_log_level,omitempty"`
	TCPflags                         bool   `json:"tcpflags,omitempty"`
}

type FirewallVirtualMachineOption struct {
	Enable      bool   `json:"enable,omitempty"`
	Dhcp        bool   `json:"dhcp,omitempty"`
	Ipfilter    bool   `json:"ipfilter,omitempty"`
	LogLevelIn  string `json:"log_level_in,omitempty"`
	LogLevelOut string `json:"log_level_out,omitempty"`
	Macfilter   bool   `json:"macfilter,omitempty"`
	Ntp         bool   `json:"ntp,omitempty"`
	PolicyIn    string `json:"policy_in,omitempty"`
	PolicyOut   string `json:"policy_out,omitempty"`
	Radv        bool   `json:"radv,omitempty"`
}
