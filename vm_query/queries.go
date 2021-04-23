package vm_query

import (
	"libvirt-go-api/connector"
	"libvirt-go-api/models"
	"log"
	"regexp"

	libvirt "github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

func GetDomsActive() (domains []models.Domains) {
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
func GetDomsInActive() (domains []models.Domains) {
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

func GetDomStats(Uuid string) (d *models.DomStatus, err error) {
	var i []models.DomInterfaces
	var pids []string
	dom, err := connector.Libvirt.Connection.LookupDomainByUUIDString(Uuid)
	if err != nil {
		return d, err
	}
	dom_status, err := dom.IsActive()
	dom_info, err := dom.GetInfo()
	dom_name, err := dom.GetName()
	//dom_os, _ := dom.GetGuestInfo(libvirt.DOMAIN_GUEST_INFO_OS, 0)
	if err != nil {
		return d, err
	}
	xml_desc, _ := dom.GetXMLDesc(0)
	domcfg := &libvirtxml.Domain{}
	err = domcfg.Unmarshal(xml_desc)
	if err != nil {
		return d, err
	}
	if dom_status {
		for _, mac_addr := range domcfg.Devices.Interfaces {
			i = append(i, models.DomInterfaces{Macaddr: mac_addr.MAC.Address, Name: mac_addr.Alias.Name})

		}

		pids = GetVmPids(dom)
	}
	d = &models.DomStatus{Status: dom_status,
		Name:       dom_name,
		Uuid:       Uuid,
		Cputime:    dom_info.CpuTime,
		Maxmemory:  dom_info.MaxMem,
		Memory:     dom_info.Memory,
		Virtcpuct:  dom_info.NrVirtCpu,
		Domxml:     xml_desc,
		Interfaces: i,
		P_ids:      pids,
	}
	return d, err
}

func GetAllDomains() (domains []models.Domains) {
	active := GetDomsActive()
	inactive := GetDomsInActive()
	domains = append(domains, active...)
	domains = append(domains, inactive...)
	return domains
}

func GetVmPids(dom *libvirt.Domain) (str_pids []string) {

	pids, err := dom.QemuMonitorCommand("info cpus", libvirt.DOMAIN_QEMU_MONITOR_COMMAND_HMP)
	if err != nil {
		log.Fatalln(err)
		return
	}
	regThreadID := regexp.MustCompile("thread_id=([0-9]*)\\s")
	threadIDsRaw := regThreadID.FindAllStringSubmatch(pids, -1)
	str_pids = make([]string, len(threadIDsRaw))
	for i, thread := range threadIDsRaw {
		str_pids[i] = thread[1]
	}
	return str_pids
}

func GetActiveNetworks() []models.VirtNet {
	var virtnets []models.VirtNet
	var DHCPLeases []models.DHCPLeases
	network, err := connector.Libvirt.Connection.ListAllNetworks(libvirt.CONNECT_LIST_NETWORKS_ACTIVE)
	if err != nil {
		log.Fatalln(err)
	}
	virtnets = make([]models.VirtNet, len(network))
	for i, net := range network {

		start, _ := net.GetAutostart()
		bname, _ := net.GetBridgeName()
		uuid, _ := net.GetUUIDString()
		xmldesc, _ := net.GetXMLDesc(libvirt.NETWORK_XML_INACTIVE)
		name, _ := net.GetName()
		ip_leases, _ := net.GetDHCPLeases()
		DHCPLeases = make([]models.DHCPLeases, len(ip_leases))
		for j, ip_leases := range ip_leases {
			DHCPLeases[j] = models.DHCPLeases{Hostname: ip_leases.Hostname,
				IpAddr:   ip_leases.IPaddr,
				MacAddr:  ip_leases.Mac,
				ClientId: ip_leases.Clientid}
		}
		virtnets[i] = models.VirtNet{Name: name,
			AutoStart:  start,
			BridgeName: bname,
			Uuid:       uuid,
			XmlDesc:    xmldesc,
			IpLeases:   DHCPLeases}
		net.Free()

	}
	return virtnets
}
