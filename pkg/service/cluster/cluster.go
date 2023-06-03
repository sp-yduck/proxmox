package cluster

import (
	"fmt"
	"strings"
)

func (c *Cluster) NextID() (int, error) {
	var nextid int
	if err := c.Client.Get("/cluster/nextid", &nextid); err != nil {
		return 0, err
	}
	return nextid, nil
}

// Resources retrieves a summary list of all resources in the cluster.
// It calls /cluster/resources api v2 endpoint with an optional "type" parameter
// to filter searched values.
// It returns a list of ClusterResources.
func (c *Cluster) Resources(filters ...string) (rs []*Resource, err error) {
	url := "/cluster/resources"

	// filters are variadic because they're optional, munging everything passed into one big string to make
	// a good request and the api will error out if there's an issue
	if f := strings.Replace(strings.Join(filters, ""), " ", "", -1); f != "" {
		url = fmt.Sprintf("%s?type=%s", url, f)
	}

	return rs, c.Client.Get(url, &rs)
}

func (c *Cluster) FWGroups() (groups []*FirewallSecurityGroup, err error) {
	err = c.Client.Get("/cluster/firewall/groups", &groups)

	if nil == err {
		for _, g := range groups {
			g.Client = c.Client
		}
	}
	return
}

func (c *Cluster) FWGroup(name string) (group *FirewallSecurityGroup, err error) {
	group = &FirewallSecurityGroup{}
	err = c.Client.Get(fmt.Sprintf("/cluster/firewall/groups/%s", name), &group.Rules)
	if nil == err {
		group.Group = name
		group.Client = c.Client
	}
	return
}

func (c *Cluster) NewFWGroup(group *FirewallSecurityGroup) error {
	return c.Client.Post(fmt.Sprintf("/cluster/firewall/groups"), group, &group)
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
