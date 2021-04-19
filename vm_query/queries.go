package vm_query

import (
	"libvirt-go-api/connector"
	"libvirt-go-api/models"
	"libvirt-go-api/proc_parse"
	"log"

	libvirt "github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

func ListActive() (domains []models.Domains) {
	doms, err := connector.Libvirt.Connection.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		log.Fatal(err)
	}
	domains = make([]models.Domains, len(doms))
	for i, dom := range doms {
		name, err := dom.GetName()
		dom_id, err := dom.GetUUIDString()
		status, err := dom.IsActive()
		if err == nil {
			domains[i] = models.Domains{Name: name, Uuid: dom_id, Status: status}
		}
		dom.Free()
	}
	return domains

}
func ListInActive() (domains []models.Domains) {
	doms, err := connector.Libvirt.Connection.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_INACTIVE)
	if err != nil {
		log.Fatalln(err)
	}
	domains = make([]models.Domains, len(doms))
	for i, dom := range doms {
		name, err := dom.GetName()
		dom_id, err := dom.GetUUIDString()
		status, err := dom.IsActive()
		if err == nil {
			domains[i] = models.Domains{Name: name, Uuid: dom_id, Status: status}
		}
		dom.Free()
	}
	return domains
}

func DomStats(Uuid string) (d *models.DomStatus) {
	var i []models.DomInterfaces
	dom, err := connector.Libvirt.Connection.LookupDomainByUUIDString(Uuid)
	if err != nil {
		log.Fatalln(err)
	}
	dom_status, err := dom.IsActive()
	dom_info, err := dom.GetInfo()
	dom_name, err := dom.GetName()
	if err != nil {
		log.Fatal(err)
	}
	xml_desc, _ := dom.GetXMLDesc(0)
	domcfg := &libvirtxml.Domain{}
	err = domcfg.Unmarshal(xml_desc)
	if err != nil {
		log.Fatal(err)
	}

	for _, mac_addr := range domcfg.Devices.Interfaces {
		i = append(i, models.DomInterfaces{Macaddr: mac_addr.MAC.Address, Name: mac_addr.Alias.Name})

	}
	pids := proc_parse.Vm_pids(dom)
	d = &models.DomStatus{Status: dom_status, Name: dom_name, Uuid: Uuid, Cputime: dom_info.CpuTime, Maxmemory: dom_info.MaxMem, Memory: dom_info.Memory, Virtcpuct: dom_info.NrVirtCpu, Domxml: xml_desc, Interfaces: i, P_ids: pids}
	return d
}

func GetAllDomains() (domains []models.Domains) {
	active := ListActive()
	inactive := ListInActive()
	domains = append(domains, active...)
	domains = append(domains, inactive...)
	return domains
}
