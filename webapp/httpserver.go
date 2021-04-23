package webapp

import (
	"libvirt-go-api/models"
	"libvirt-go-api/proc_parse"
	"libvirt-go-api/vm_query"

	"github.com/gin-gonic/gin"
)

var HostCpu []models.ProcCpuinfo
var router *gin.Engine

func init() {
	HostCpu = proc_parse.GetHostCpuInfo()
	router = gin.Default()
	read_only_group := router.Group("/api/get")
	{
		read_only_group.GET("/hostcpuinfo", func(c *gin.Context) {
			c.JSON(200, HostCpu)
		})
		read_only_group.GET("/hostmeminfo", func(c *gin.Context) {
			c.JSON(200, proc_parse.GetHostMemoryInfo())
		})
		read_only_group.GET("/getalldomains", func(c *gin.Context) {
			c.JSON(200, vm_query.GetAllDomains())
		})
		read_only_group.GET("/getallnetworks", func(c *gin.Context) {
			c.JSON(200, vm_query.GetActiveNetworks())
		})
		read_only_group.GET(("/domain/:uuid"), func(c *gin.Context) {
			uuid := c.Param("uuid")
			d, err := vm_query.GetDomStats(uuid)
			if err != nil {
				c.JSON(404, err)
			} else {
				c.JSON(200, d)
			}
		})

	}

}

func StartServer(portnum uint) {

	router.Run(":9000")
}
