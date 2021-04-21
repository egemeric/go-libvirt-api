package main

import (
	"fmt"
	"libvirt-go-api/connector"
	"libvirt-go-api/models"
	"libvirt-go-api/proc_parse"
	"libvirt-go-api/vm_query"
	"log"
)

func init() {
	connector.Libvirt.ConenctionURI = "qemu:///system"
}

func main() {
	//cores := proc_parse.GetHostCpuInfo()
	//fmt.Println(cores[0])

	fmt.Println("Host Total Mem:", proc_parse.GetHostMemoryInfo().MemTotal/1000000, "GB")

	err := connector.StartConn()
	var domains []models.Domains
	if err != nil {
		log.Fatalln(err)
	}
	domains = vm_query.GetAllDomains()
	fmt.Println(domains)
	err = connector.CloseConn()

	if err != nil {
		log.Fatalln(err)
		return
	}

}
