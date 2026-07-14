package settings

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
)

type Settings struct {
	GUIPreferred   bool   `json:"gui_preferred"`
	SilentBoot     bool   `json:"silent_boot"`
	AutoInject     bool   `json:"auto_inject"`
	NetworkDirect  bool   `json:"network_direct"`
	DefaultArch    string `json:"default_arch"`
	DefaultMode    string `json:"default_mode"`
	InstallDir     string `json:"install_dir"`
	RegistryPath   string `json:"registry_path"`
	ChannelPort    int    `json:"channel_port"`
	DownloadRetries int   `json:"download_retries"`
	ProxyURL       string `json:"proxy_url"`
	LogLevel       string `json:"log_level"`
	Theme          string `json:"theme"`
	FontSize       int    `json:"font_size"`
	ConfirmExit    bool   `json:"confirm_exit"`
	SaveSession    bool   `json:"save_session"`
	MaxInstances   int    `json:"max_instances"`
	CPULimit       int    `json:"cpu_limit"`
	MemoryLimit    int    `json:"memory_limit"`
	AutoUpdate     bool   `json:"auto_update"`
}

type SettingsManager struct {
	mu       sync.RWMutex
	settings *Settings
	path     string
}

func NewSettingsManager(configPath string) (*SettingsManager, error) {
	if configPath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		configPath = filepath.Join(home, ".ela", "settings.json")
	}
	sm := &SettingsManager{
		path: configPath,
		settings: &Settings{
			GUIPreferred:    false,
			SilentBoot:      false,
			AutoInject:      true,
			NetworkDirect:   false,
			DefaultArch:     "x86_64",
			DefaultMode:     "hyper",
			InstallDir:      filepath.Join(os.Getenv("HOME"), ".ela", "instances"),
			RegistryPath:    filepath.Join(os.Getenv("HOME"), ".ela", "registry.db"),
			ChannelPort:     9999,
			DownloadRetries: 3,
			ProxyURL:        "",
			LogLevel:        "info",
			Theme:           "default",
			FontSize:        14,
			ConfirmExit:     true,
			SaveSession:     true,
			MaxInstances:    10,
			CPULimit:        0,
			MemoryLimit:     0,
			AutoUpdate:      true,
		},
	}
	_ = sm.Load()
	return sm, nil
}

func (sm *SettingsManager) Load() error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	data, err := os.ReadFile(sm.path)
	if err != nil {
		if os.IsNotExist(err) {
			return sm.save()
		}
		return err
	}
	var s Settings
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	sm.settings = &s
	return nil
}

func (sm *SettingsManager) save() error {
	data, err := json.MarshalIndent(sm.settings, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(sm.path), 0755); err != nil {
		return err
	}
	return os.WriteFile(sm.path, data, 0644)
}

func (sm *SettingsManager) Save() error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.save()
}

func (sm *SettingsManager) Get() *Settings {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	s := *sm.settings
	return &s
}

func (sm *SettingsManager) Set(key string, value interface{}) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	switch key {
	case "gui_preferred":
		if v, ok := value.(bool); ok {
			sm.settings.GUIPreferred = v
		}
	case "silent_boot":
		if v, ok := value.(bool); ok {
			sm.settings.SilentBoot = v
		}
	case "auto_inject":
		if v, ok := value.(bool); ok {
			sm.settings.AutoInject = v
		}
	case "network_direct":
		if v, ok := value.(bool); ok {
			sm.settings.NetworkDirect = v
		}
	case "default_arch":
		if v, ok := value.(string); ok {
			sm.settings.DefaultArch = v
		}
	case "default_mode":
		if v, ok := value.(string); ok {
			sm.settings.DefaultMode = v
		}
	case "install_dir":
		if v, ok := value.(string); ok {
			sm.settings.InstallDir = v
		}
	case "registry_path":
		if v, ok := value.(string); ok {
			sm.settings.RegistryPath = v
		}
	case "channel_port":
		if v, ok := value.(int); ok {
			sm.settings.ChannelPort = v
		}
	case "download_retries":
		if v, ok := value.(int); ok {
			sm.settings.DownloadRetries = v
		}
	case "proxy_url":
		if v, ok := value.(string); ok {
			sm.settings.ProxyURL = v
		}
	case "log_level":
		if v, ok := value.(string); ok {
			sm.settings.LogLevel = v
		}
	case "theme":
		if v, ok := value.(string); ok {
			sm.settings.Theme = v
		}
	case "font_size":
		if v, ok := value.(int); ok {
			sm.settings.FontSize = v
		}
	case "confirm_exit":
		if v, ok := value.(bool); ok {
			sm.settings.ConfirmExit = v
		}
	case "save_session":
		if v, ok := value.(bool); ok {
			sm.settings.SaveSession = v
		}
	case "max_instances":
		if v, ok := value.(int); ok {
			sm.settings.MaxInstances = v
		}
	case "cpu_limit":
		if v, ok := value.(int); ok {
			sm.settings.CPULimit = v
		}
	case "memory_limit":
		if v, ok := value.(int); ok {
			sm.settings.MemoryLimit = v
		}
	case "auto_update":
		if v, ok := value.(bool); ok {
			sm.settings.AutoUpdate = v
		}
	default:
		return errors.New("unknown setting: " + key)
	}
	return sm.save()
}

func (sm *SettingsManager) Reset() error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	home, _ := os.UserHomeDir()
	sm.settings = &Settings{
		GUIPreferred:    false,
		SilentBoot:      false,
		AutoInject:      true,
		NetworkDirect:   false,
		DefaultArch:     "x86_64",
		DefaultMode:     "hyper",
		InstallDir:      filepath.Join(home, ".ela", "instances"),
		RegistryPath:    filepath.Join(home, ".ela", "registry.db"),
		ChannelPort:     9999,
		DownloadRetries: 3,
		ProxyURL:        "",
		LogLevel:        "info",
		Theme:           "default",
		FontSize:        14,
		ConfirmExit:     true,
		SaveSession:     true,
		MaxInstances:    10,
		CPULimit:        0,
		MemoryLimit:     0,
		AutoUpdate:      true,
	}
	return sm.save()
}
