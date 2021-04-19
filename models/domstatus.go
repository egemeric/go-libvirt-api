package models

type DomStatus struct {
	Status     bool
	Name       string
	Uuid       string
	Cputime    uint64
	Maxmemory  uint64
	Memory     uint64
	Virtcpuct  uint
	Domxml     string
	Interfaces []DomInterfaces
	P_ids      []string
}
type DomInterfaces struct {
	Name    string
	Macaddr string
}

type Domains struct {
	Name   string
	Uuid   string
	Status bool
}
