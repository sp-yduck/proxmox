package cluster

import (
	"fmt"
	"strconv"
	"strings"
)

// /cluster

type Cluster struct {
	Client  Client
	Version int
	Quorate int
	Name    string
	ID      string
}

type Client interface {
	Get(p string, v interface{}) error
	Post(p string, d interface{}, v interface{}) error
	Put(p string, d interface{}, v interface{}) error
	Delete(p string, v interface{}) error
}

type FirewallSecurityGroup struct {
	Client  Client
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

type ClusterResources []*ClusterResource

type ClusterResource struct {
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

func (cl *Cluster) NextID() (int, error) {
	var ret string
	cl.Client.Get("/cluster/nextid", &ret)
	return strconv.Atoi(ret)
}

// Resources retrieves a summary list of all resources in the cluster.
// It calls /cluster/resources api v2 endpoint with an optional "type" parameter
// to filter searched values.
// It returns a list of ClusterResources.
func (cl *Cluster) Resources(filters ...string) (rs ClusterResources, err error) {
	url := "/cluster/resources"

	// filters are variadic because they're optional, munging everything passed into one big string to make
	// a good request and the api will error out if there's an issue
	if f := strings.Replace(strings.Join(filters, ""), " ", "", -1); f != "" {
		url = fmt.Sprintf("%s?type=%s", url, f)
	}

	return rs, cl.Client.Get(url, &rs)
}

func (cl *Cluster) FWGroups() (groups []*FirewallSecurityGroup, err error) {
	err = cl.Client.Get("/cluster/firewall/groups", &groups)

	if nil == err {
		for _, g := range groups {
			g.Client = cl.Client
		}
	}
	return
}

func (cl *Cluster) FWGroup(name string) (group *FirewallSecurityGroup, err error) {
	group = &FirewallSecurityGroup{}
	err = cl.Client.Get(fmt.Sprintf("/cluster/firewall/groups/%s", name), &group.Rules)
	if nil == err {
		group.Group = name
		group.Client = cl.Client
	}
	return
}

func (cl *Cluster) NewFWGroup(group *FirewallSecurityGroup) error {
	return cl.Client.Post(fmt.Sprintf("/cluster/firewall/groups"), group, &group)
}

func (g *FirewallSecurityGroup) GetRules() ([]*FirewallRule, error) {
	return g.Rules, g.Client.Get(fmt.Sprintf("/cluster/firewall/groups/%s", g.Group), &g.Rules)
}

func (g *FirewallSecurityGroup) Delete() error {
	return g.Client.Delete(fmt.Sprintf("/cluster/firewall/groups/%s", g.Group), nil)
}

func (g *FirewallSecurityGroup) RuleCreate(rule *FirewallRule) error {
	return g.Client.Post(fmt.Sprintf("/cluster/firewall/groups/%s", g.Group), rule, nil)
}

func (g *FirewallSecurityGroup) RuleUpdate(rule *FirewallRule) error {
	return g.Client.Put(fmt.Sprintf("/cluster/firewall/groups/%s/%d", g.Group, rule.Pos), rule, nil)
}

func (g *FirewallSecurityGroup) RuleDelete(rulePos int) error {
	return g.Client.Delete(fmt.Sprintf("/cluster/firewall/groups/%s/%d", g.Group, rulePos), nil)
}
