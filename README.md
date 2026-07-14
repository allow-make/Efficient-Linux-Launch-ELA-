# ELA - Efficient Linux Launcher

> A fast Linux launcher built with C, C++ and Go, designed for managing the state, instances, and runtime of multiple Linux distributions.

[![License: GPL v2](https://img.shields.io/badge/License-GPL%20v2-blue.svg)](https://www.gnu.org/licenses/old-licenses/gpl-2.0.en.html)
[![PR Checks](https://github.com/allow-make/Efficient-Linux-Launch-ELA-/actions/workflows/pr-checks.yml/badge.svg)](https://github.com/allow-make/Efficient-Linux-Launch-ELA-/actions/workflows/pr-checks.yml)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Haxe Version](https://img.shields.io/badge/Haxe-4.3+-EA8220?style=flat&logo=haxe)](https://haxe.org/)

---

## 🚀 What is ELA?

ELA (**E**fficient **L**inux **A**uncher) is a fast, lightweight Linux launcher for Windows.

It allows you to:
- **Launch multiple Linux distributions** (Ubuntu, Debian, Arch, Kali, etc.)
- **Switch between terminal and GUI projection** modes
- **Manage instances** via the ELRA registry
- **Package instances** into portable `.elra` files
- **Run on both x86_64 and ARM64** architectures

> **ELA does not draw the terminal or the GUI.**  
> It projects the real Linux output (text or framebuffer) into a window.

---

## 🧠 Architecture

```

Linux Kernel
↓
LS (Linux Script) – captures output
↓
Lhyper – transport channel
↓
HL (Hyper Link) – sequence protocol
↓
ELA (Windows launcher) – displays output
↓
ELRA – registry management

```

### Key Components

| Component | Language | Role |
|-----------|----------|------|
| **ELA.exe** | C + Go + Haxe | Main launcher, UI, terminal rendering |
| **LS** | C | Linux output capture |
| **Lhyper** | Go | Transport channel |
| **HL** | Protocol | Sequence language (text, API, bridge) |
| **ELRA** | C + SQLite | Instance registry |

---

## 📦 Supported Distributions

| Category | Distributions |
|----------|---------------|
| **Mainstream** | Ubuntu, Debian, Fedora, openSUSE |
| **Geek** | Arch Linux, Gentoo, NixOS |
| **Security** | Kali Linux |
| **Lightweight** | Alpine Linux |
| **Chinese** | Deepin |

> **More distributions can be added via custom `.elra` packages.**

---

## 🛠️ Build from Source

### Prerequisites

- Go 1.21+
- Haxe 4.3+
- C compiler (GCC/MinGW)
- OpenGL development libraries
- SQLite3

### Build

```bash
# Clone the repository
git clone https://github.com/allow-make/Efficient-Linux-Launch-ELA-.git
cd Efficient-Linux-Launch-ELA-

# Build Go components
go build -o bin/ela.exe ./cmd/ela/...

# Build Haxe components
haxe build-cpp.hxml
haxe build-jvm.hxml

# Build C components
mkdir build && cd build
cmake .. -DCMAKE_BUILD_TYPE=Release
make -j$(nproc)

# The output will be in bin/
```

---

📖 Usage

First Start

1. Double-click ELA.exe
2. Read and accept the 10-page license agreement (Enter to proceed)
3. Choose a Linux distribution to download
4. Launch and start using it

Switching Instances

```bash
# In the ELA terminal
ela exit
```

This saves the current session and returns to the launcher wizard.

Registry Commands (via HL)

```hl
# Create an empty instance
\n;api;"ela";"create_instance";"{name:\"arch\",arch:\"x86_64\"}"

# Create a record package
\n;api;"ela";"create_record";"{target_id:\"INS-002\",mode:\"hyper\",path:\"/ela/instances/arch\"}"

# Merge record into instance
\n;api;"ela";"merge_record";"{instance_id:\"INS-002\",record_id:\"REC-001\"}"
```

---

🔧 Configuration

Setting Description
GUI Preferred Boot into GUI mode if supported
Silent Boot Auto-start Linux on Windows boot
Auto Inject Auto-inject LS + Lhyper into registered instances
Network Direct Linux uses host network directly

---

🤝 Contributing

We welcome contributions! Please read our Contributing Guidelines before submitting a PR.

Language Policy

Language Status
C ✅ Strongly Preferred
Go ✅ Strongly Preferred
Haxe (.hx) ✅ Preferred
C++ 🟡 Accepted (Strict)
JVM Languages ❌ Not Recommended (JNI is slow)
Scripting Languages ❌ Avoid

All PRs go through strict review. Please check the PR Checks before submitting.

---

📄 License

ELA is licensed under the GNU General Public License v2.0.

See the LICENSE file for details.

---

📬 Contact

· Issues: GitHub Issues
· Discussions: GitHub Discussions

---

Made with ❤️ for Linux enthusiasts everywhere.
