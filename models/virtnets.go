package models

type VirtNet struct {
	Name       string
	AutoStart  bool
	BridgeName string
	Uuid       string
	XmlDesc    string
	IpLeases   []DHCPLeases
}

type DHCPLeases struct {
	Hostname string
	IpAddr   string
	MacAddr  string
	ClientId string
}
