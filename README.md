# GoRAT - A Lightweight Red Teaming Tool in Go

![GoRAT](https://img.shields.io/badge/Go-1.24-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)
 
GoRAT is a **lightweight remote access and system reconnaissance tool** written in **Go**, designed for **red teamers**, **penetration testers**, and **security researchers**. It allows remote execution of commands, system reconnaissance, and clipboard data retrieval through a **Telegram C2**.

---

## 🚀 Features

✅ **System Reconnaissance** – Gather OS details, environment variables, and running processes  
✅ **Remote Command Execution** – Run arbitrary shell commands remotely  
✅ **Clipboard Data Extraction** – Capture clipboard contents  
✅ **File & Directory Enumeration** – List files and directories  
✅ **System Resource Monitoring** – Fetch CPU, memory, disk, and network statistics  
✅ **Telegram C2 Integration** – Receive and execute commands via Telegram Bot  

---

## 📌 Installation

### Prerequisites:
- **Go 1.21+** installed on your system
- A **Telegram bot** with an API token
- Your Telegram **Chat ID**

### Clone the Repository:
```sh
git clone https://github.com/viralvaghela/GoRAT.git
cd GoRAT
```

### Edit Configuration:
Modify the `TOKEN` and `CHATID` variables in the source code to your Telegram bot credentials.

```go
const (
    TOKEN  = "YOUR_TELEGRAM_BOT_TOKEN"
    CHATID = "YOUR_CHAT_ID"
)
```

### Build the Binary:
```sh
go build -o gorat main.go
```

### Run the Tool:
```sh
./gorat
```

---

## 🎯 Usage

Once the bot is running, send the following commands via Telegram:

| Command | Description |
|---------|-------------|
| `info` | Get system information (OS, Arch, User) |
| `cpu` | Get CPU usage percentage |
| `ram` | Get memory usage percentage |
| `disk` | Get disk usage percentage |
| `network` | Get network interface details |
| `processes` | List running processes |
| `clipboard` | Get clipboard contents |
| `env` | Retrieve environment variables |
| `ls` | List files in the current directory |
| `<custom_command>` | Execute any shell command remotely |

---

## 🔥 To-Do List (Planned Enhancements)

- [ ] **AV & EDR Evasion** – Implement obfuscation, encryption, and AMSI bypass
- [ ] **Persistence Techniques** – Add registry modifications, system services, and cron jobs
- [ ] **Network & Port Scanning** – Expand reconnaissance capabilities
- [ ] **Payload Size Reduction** – Optimize and shrink the binary for stealth
- [ ] **Additional Data Exfiltration** – Extract more sensitive information

---

## ⚠️ Disclaimer
This tool is intended for **educational and research purposes only**. Unauthorized use of this tool for malicious activities is strictly prohibited. The author is not responsible for any misuse or damage caused by this software.

---

## 📜 License
This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.

---

## 🔗 Connect
If you have suggestions or improvements, feel free to **open an issue** or **contribute** to the project.

📌 **GitHub Repository:** [GoRAT](https://github.com/viralvaghela/GoRAT)
 
