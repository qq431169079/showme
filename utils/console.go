package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

func GetNowTime() string {
	return fmt.Sprintf("%s", time.Now().Format("15:04:05"))
}

func floatToString(x float64, f int) string {
	rs := strconv.FormatFloat(x, 'f', f, 64)
	return rs
}

// lambda for check length of Repeat " "
func parseRepeatSpace(info string, lens int) string {
	replace := "oops"
	if len(info) > lens {
		info = replace
	}
	return strings.Repeat(" ", lens-len(info)) + info
}

func CollectEasy() []string {
	rs := []string{}
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Info()
	load, _ := load.Avg()
	// cc, _ := cpu.Percent(time.Second, false)
	d, _ := disk.Usage("/")
	n, _ := host.Info()
	boottime, _ := host.BootTime()
	btime := time.Unix(int64(boottime), 0).Format("2006-01-02 15:04:05")
	rs = append(rs, fmt.Sprintf("%s: %v MB  Free: %v MB Used:%v Usage:%f%%", Colorize("        Mem       ", "white", "red", true, true), v.Total/1024/1024, v.Available/1024/1024, v.Used/1024/1024, v.UsedPercent))
	// fmt.Printf("        Mem       : %v MB  Free: %v MB Used:%v Usage:%f%%", v.Total/1024/1024, v.Available/1024/1024, v.Used/1024/1024, v.UsedPercent)
	if len(c) > 1 {
		for _, sub_cpu := range c {
			modelname := sub_cpu.ModelName
			cores := sub_cpu.Cores
			// fmt.Printf("        CPU       : %v   %v cores ", modelname, cores)
			rs = append(rs, fmt.Sprintf("%s: %v   %v cores ", Colorize("        CPU       ", "white", "red", true, true), modelname, cores))
		}
	} else {
		sub_cpu := c[0]
		modelname := sub_cpu.ModelName
		cores := sub_cpu.Cores
		rs = append(rs, fmt.Sprintf("%s: %v   %v cores ", Colorize("        CPU       ", "white", "red", true, true), modelname, cores))
	}
	rs = append(rs, fmt.Sprintf("%s: %.2f %.2f %.2f ", Colorize("        LOAD      ", "white", "red", true, true), load.Load1, load.Load5, load.Load15))
	rs = append(rs, fmt.Sprintf("%s: %v", Colorize("        SystemBoot", "white", "red", true, true), btime))
	// rs = append(rs, fmt.Sprintf("        CPU Used    : used %f%% ", cc[0]))
	rs = append(rs, fmt.Sprintf("%s: %v GB  Free: %v GB Usage:%f%%", Colorize("        HD        ", "white", "red", true, true), d.Total/1024/1024/1024, d.Free/1024/1024/1024, d.UsedPercent))
	rs = append(rs, fmt.Sprintf("%s: %v %v(%v)   %v", Colorize("        OS        ", "white", "red", true, true), n.OS, n.Platform, n.PlatformFamily, n.PlatformVersion))
	rs = append(rs, fmt.Sprintf("%s: %v ", Colorize("        Kernel    ", "white", "red", true, true), n.KernelVersion))
	rs = append(rs, fmt.Sprintf("%s: %v ", Colorize("        HostID    ", "white", "red", true, true), n.HostID))
	rs = append(rs, fmt.Sprintf("%s: %v ", Colorize("        Procs     ", "white", "red", true, true), n.Procs))
	rs = append(rs, fmt.Sprintf("%s: %v  ", Colorize("        Hostname  ", "white", "red", true, true), n.Hostname))
	rs = append(rs, fmt.Sprintf("%s: %s\n", Colorize("        IpLists   ", "white", "red", true, true), strings.Join(GetIps(), ",")))
	return rs
}

// 文字字体 参数介绍：text->文本内容 status->文字颜色 background->背景颜色 underline->是否下划线 highshow->是否高亮
// http://www.cnblogs.com/frydsh/p/4139922.html
func Colorize(text string, status string, background string, underline bool, highshow bool) string {
	out_one := "\033["
	out_two := ""
	out_three := ""
	out_four := ""
	//可动态配置字体颜色 背景色 高亮
	// 显示：0(默认)、1(粗体/高亮)、22(非粗体)、4(单条下划线)、24(无下划线)、5(闪烁)、25(无闪烁)、7(反显、翻转前景色和背景色)、27(无反显)
	// 颜色：0(黑)、1(红)、2(绿)、 3(黄)、4(蓝)、5(洋红)、6(青)、7(白)
	//  前景色为30+颜色值，如31表示前景色为红色；背景色为40+颜色值，如41表示背景色为红色。
	if underline == true && highshow == true {
		out_four = ";1;4m" //高亮
	} else if underline != true && highshow == true {
		out_four = ";1m"
	} else if underline == true && highshow != true {
		out_four = ";4m"
	} else {
		out_four = ";22m"
	}

	switch status {
	case "black":
		out_two = "30"
	case "red":
		out_two = "31"
	case "green":
		out_two = "32"
	case "yellow":
		out_two = "33"
	case "blue":
		out_two = "34"
	case "purple":
		out_two = "35"
	case "dgreen":
		out_two = "36"
	case "white":
		out_two = "37"
	default:
		out_two = ""
	}

	switch background {
	case "black":
		out_three = "40;"
	case "red":
		out_three = "41;"
	case "green":
		out_three = "42;"
	case "yellow":
		out_three = "43;"
	case "blue":
		out_three = "44;"
	case "purple":
		out_three = "45;"
	case "dgreen":
		out_three = "46;"
	case "white":
		out_three = "47;"
	default:
		out_three = ""
	}
	return out_one + out_three + out_two + out_four + text + "\033[0m"
}
