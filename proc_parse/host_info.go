package proc_parse

import (
	"bufio"
	"fmt"
	"libvirt-go-api/models"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetHostCpuInfo() []models.ProcCpuinfo {
	core := []models.ProcCpuinfo{}
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	core_ct := 0
	for scanner.Scan() {
		row := scanner.Text()
		row_split := strings.Split(row, ":")
		if len(core) <= core_ct {
			core = append(core, models.ProcCpuinfo{})
		}
		if len(row_split) == 1 {
			core_ct++
			continue
		}
		key := strings.TrimSpace(row_split[0])
		value := strings.TrimSpace(row_split[1])
		switch key {
		case "processor":
			core[core_ct].Processor, _ = strconv.Atoi(value)
		case "vendor_id":
			core[core_ct].VendorID = value
		case "cpu family":
			core[core_ct].CPUFamily, _ = strconv.Atoi(value)
		case "model":
			core[core_ct].Model, _ = strconv.Atoi(value)
		case "model name":
			core[core_ct].ModelName = value
		case "stepping":
			core[core_ct].Stepping, _ = strconv.Atoi(value)
		case "microcode":
			core[core_ct].Microcode = value
		case "cpu MHz":
			tmp, _ := strconv.ParseFloat(value, 32)
			core[core_ct].CPUMhz = float32(tmp)
		case "cache size":
			core[core_ct].CacheSize = value
		case "physical id":
			core[core_ct].PhysicalID, _ = strconv.Atoi(value)
		case "siblings":
			core[core_ct].Siblings, _ = strconv.Atoi(value)
		case "core id":
			core[core_ct].CoreID, _ = strconv.Atoi(value)
		case "cpu cores":
			core[core_ct].CPUCores, _ = strconv.Atoi(value)
		case "apicid":
			core[core_ct].ApicID, _ = strconv.Atoi(value)
		case "initial apicid":
			core[core_ct].InitialApicID, _ = strconv.Atoi(value)
		case "fpu":
			core[core_ct].Fpu = value
		case "fpu_exception":
			core[core_ct].FpuException = value
		case "cpuid level":
			core[core_ct].CpuidLevel, _ = strconv.Atoi(value)
		case "wp":
			core[core_ct].Wp = value
		case "flags":
			flags := strings.Split(value, " ")
			fmt.Println(len(flags))
			for _, flag := range flags {
				core[core_ct].Flags = append(core[core_ct].Flags, flag)

			}

		case "bugs":
			core[core_ct].Bugs = value
		case "bogomips":
			tmp, _ := strconv.ParseFloat(value, 32)
			core[core_ct].Bogomips = float32(tmp)
		case "clflush size":
			core[core_ct].ClflushSize, _ = strconv.Atoi(value)
		case "cache_alignment":
			core[core_ct].CacheAlignment, _ = strconv.Atoi(value)
		case "address sizes":
			core[core_ct].AddressSizes = value
		case "power management":
			core[core_ct].PowerManagement = value
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return core
}

func GetHostMemoryInfo() models.ProcMeminfo {
	stats := models.ProcMeminfo{}
	file, _ := os.Open("/proc/meminfo")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		row := scanner.Text()
		rowfields := strings.Split(row, ":")

		key := strings.TrimSpace(rowfields[0])
		valueRaw := strings.TrimSpace(rowfields[1])
		valueParts := strings.Split(valueRaw, " ")
		value := valueParts[0]

		switch key {
		case "MemTotal":
			stats.MemTotal, _ = strconv.Atoi(value)
		case "MemFree":
			stats.MemFree, _ = strconv.Atoi(value)
		case "MemAvailable":
			stats.MemAvailable, _ = strconv.Atoi(value)
		case "Buffers":
			stats.Buffers, _ = strconv.Atoi(value)
		case "Cached":
			stats.Cached, _ = strconv.Atoi(value)
		case "SwapCached":
			stats.SwapCached, _ = strconv.Atoi(value)
		case "Active":
			stats.Active, _ = strconv.Atoi(value)
		case "Inactive":
			stats.Inactive, _ = strconv.Atoi(value)
		case "Active(anon)":
			stats.ActiveAanon, _ = strconv.Atoi(value)
		case "Inactive(anon)":
			stats.InactiveAanon, _ = strconv.Atoi(value)
		case "Active(file)":
			stats.ActiveFile, _ = strconv.Atoi(value)
		case "Inactive(file)":
			stats.InactiveFile, _ = strconv.Atoi(value)
		case "Unevictable":
			stats.Unevictable, _ = strconv.Atoi(value)
		case "Mlocked":
			stats.Mlocked, _ = strconv.Atoi(value)
		case "SwapTotal":
			stats.SwapTotal, _ = strconv.Atoi(value)
		case "SwapFree":
			stats.SwapFree, _ = strconv.Atoi(value)
		case "Dirty":
			stats.Dirty, _ = strconv.Atoi(value)
		case "Writeback":
			stats.Writeback, _ = strconv.Atoi(value)
		case "AnonPages":
			stats.AnonPages, _ = strconv.Atoi(value)
		case "Mapped":
			stats.Mapped, _ = strconv.Atoi(value)
		case "Shmem":
			stats.Shmem, _ = strconv.Atoi(value)
		case "Slab":
			stats.Slab, _ = strconv.Atoi(value)
		case "SReclaimable":
			stats.SReclaimable, _ = strconv.Atoi(value)
		case "SUnreclaim":
			stats.SUnreclaim, _ = strconv.Atoi(value)
		case "KernelStack":
			stats.KernelStack, _ = strconv.Atoi(value)
		case "PageTables":
			stats.PageTables, _ = strconv.Atoi(value)
		case "NFS_Unstable":
			stats.NFSUnstable, _ = strconv.Atoi(value)
		case "Bounce":
			stats.Bounce, _ = strconv.Atoi(value)
		case "WritebackTmp":
			stats.WritebackTmp, _ = strconv.Atoi(value)
		case "CommitLimit":
			stats.CommitLimit, _ = strconv.Atoi(value)
		case "Committed_AS":
			stats.CommittedAS, _ = strconv.Atoi(value)
		case "VmallocTotal":
			stats.VmallocTotal, _ = strconv.Atoi(value)
		case "VmallocUsed":
			stats.VmallocUsed, _ = strconv.Atoi(value)
		case "VmallocChunk":
			stats.VmallocChunk, _ = strconv.Atoi(value)
		case "HardwareCorrupted":
			stats.HardwareCorrupted, _ = strconv.Atoi(value)
		case "AnonHugePages":
			stats.AnonHugePages, _ = strconv.Atoi(value)
		case "ShmemHugePages":
			stats.ShmemHugePages, _ = strconv.Atoi(value)
		case "ShmemPmdMapped":
			stats.ShmemPmdMapped, _ = strconv.Atoi(value)
		case "HugePages_Total":
			stats.HugePagesTotal, _ = strconv.Atoi(value)
		case "HugePages_Free":
			stats.HugePagesFree, _ = strconv.Atoi(value)
		case "HugePages_Rsvd":
			stats.HugePagesRsvd, _ = strconv.Atoi(value)
		case "HugePages_Surp":
			stats.HugePagesSurp, _ = strconv.Atoi(value)
		case "Hugepagesize":
			stats.Hugepagesize, _ = strconv.Atoi(value)
		case "Hugetlb":
			stats.Hugetlb, _ = strconv.Atoi(value)
		case "DirectMap4k":
			stats.DirectMap4k, _ = strconv.Atoi(value)
		case "DirectMap2M":
			stats.DirectMap2M, _ = strconv.Atoi(value)
		case "DirectMap1G":
			stats.DirectMap1G, _ = strconv.Atoi(value)
		}
	}

	return stats
}
