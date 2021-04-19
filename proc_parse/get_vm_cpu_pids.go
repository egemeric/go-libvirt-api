package proc_parse

import (
	"log"
	"regexp"

	libvirt "github.com/libvirt/libvirt-go"
)

func Vm_pids(dom *libvirt.Domain) (str_pids []string) {

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
