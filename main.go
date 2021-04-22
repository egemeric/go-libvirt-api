package main

import (
	"fmt"
	"libvirt-go-api/connector"
	"libvirt-go-api/proc_parse"
	"libvirt-go-api/vm_query"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	connector.Libvirt.ConenctionURI = "qemu:///system"
}

func main() {
	fmt.Println("Host Total Mem:", proc_parse.GetHostMemoryInfo().MemTotal/1000000, "GB")
	err := connector.StartConn()
	if err != nil {
		log.Fatalln(err)
	}
	StartServer()
	err = connector.CloseConn()
	if err != nil {
		log.Fatalln(err)
		return
	}

}

func StartServer() {
	r := gin.Default()
	r.GET("/virtnetworks", func(c *gin.Context) {
		c.JSON(200, vm_query.GetActiveNetworks())
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, vm_query.GetAllDomains())
	})
	r.GET("/hostinfo", func(c *gin.Context) {
		c.JSON(200, proc_parse.GetHostCpuInfo())
	})
	r.GET("/hostinfomem", func(c *gin.Context) {
		c.JSON(200, proc_parse.GetHostMemoryInfo())
	})

	r.Run(":9000")
}
