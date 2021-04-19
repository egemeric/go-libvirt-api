package connector

import (
	"log"

	libvirt "github.com/libvirt/libvirt-go"
)

var Libvirt struct {
	ConenctionURI string
	Connection    *libvirt.Connect
}

func StartConn() error {
	conn, err := libvirt.NewConnect(Libvirt.ConenctionURI)
	if err != nil {
		log.Printf("Failed to connect to libvirt. %+v", err)
		return err
	}
	Libvirt.Connection = conn
	return nil
}

func CloseConn() error {
	_, err := Libvirt.Connection.Close()
	if err != nil {
		log.Printf("Failed to close connection to libvirt. %+v", err)
		return err
	}
	return nil
}
