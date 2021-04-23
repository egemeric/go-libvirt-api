package main

import (
	"libvirt-go-api/connector"
	"libvirt-go-api/webapp"
	"log"
)

func init() {
	connector.Libvirt.ConenctionURI = "qemu:///system"
}

func main() {
	err := connector.StartConn()
	if err != nil {
		log.Fatalln(err)
	}
	webapp.StartServer(9000)
	err = connector.CloseConn()
	if err != nil {
		log.Fatalln(err)
		return
	}

}
