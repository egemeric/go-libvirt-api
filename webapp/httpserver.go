package webapp

import (
	"libvirt-go-api/models"
	"libvirt-go-api/proc_parse"
	"libvirt-go-api/vm_query"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var HostCpu []models.ProcCpuinfo
var router *gin.Engine

func init() {
	HostCpu = proc_parse.GetHostCpuInfo()
	router = gin.Default()
	workingdir, _ := os.Getwd()
	router.StaticFS("/app", http.Dir(workingdir+"/webapp/html"))
	read_only_group := router.Group("/api/get")
	{
		read_only_group.Use(CORSMiddleware())
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
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func StartServer(portnum uint) {
	var p string = strconv.FormatUint(uint64(portnum), 10)
	router.Run(":" + p)
}
