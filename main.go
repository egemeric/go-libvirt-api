package main

import (
	"fmt"
	"libvirt-go-api/connector"
	"libvirt-go-api/models"
	"libvirt-go-api/vm_query"
	"log"
)

func init() {
	connector.Libvirt.ConenctionURI = "qemu:///system"
}

func main() {
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
