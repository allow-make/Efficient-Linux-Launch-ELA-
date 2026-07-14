package i18n

var En = map[string]string{
// License
"license_title":     "ELA License Agreement",
"license_text":      "This software is provided AS IS, without any warranty.\nThe developer is not responsible for any loss caused by using this software.\n\nSource: https://github.com/allow-make/Efficient-Linux-Launch-ELA-\n\nPress Enter to agree and continue, Delete to exit",
"license_accept":    "Agree and Continue",
"license_reject":    "Exit",

// Main Menu
"menu_title":        "ELA Launcher",
"menu_download":     "Download Linux System",
"menu_existing":     "Use Existing System",
"menu_usb":          "Boot from USB",
"menu_registry":     "Registry Management",
"menu_settings":     "Settings",
"menu_exit":         "Exit",
"menu_version":      "Version",

// Download
"download_title":       "Select Distribution to Download",
"download_confirm":     "Confirm Download",
"download_confirm_msg": "Download %s %s (%s)?\nSize: %d MB",
"download_start":       "Downloading %s...",
"download_progress":    "%s %.1f%% (%.2f MB/s)",
"download_done":        "Done",
"download_done_msg":    "%s download completed!",
"download_fail":        "Download failed",

// Existing System
"existing_title":       "Use Existing System",
"existing_info":        "Select an existing Linux image or .elra package",
"existing_browse":      "Browse",
"existing_load":        "Load",
"existing_selected":    "Selected",
"existing_load_msg":    "Loading: %s",

// USB
"usb_title":            "Boot from USB",
"usb_info":             "Detected bootable USB devices:",
"usb_confirm":          "Boot from USB",
"usb_confirm_msg":      "Boot from %s?",
"usb_start":            "Booting from %s...",

// Registry
"registry_title":       "Registry Management",
"registry_delete_confirm": "Delete Instance",
"registry_delete_msg":  "Delete instance %s (%s)?",
"registry_create":      "Create Record Package",
"registry_merge":       "Merge Record Package",
"registry_refresh":     "Refresh",
"registry_back":        "Back",

// Create Record
"create_record_title":  "Create Record Package",
"create_record_id":     "Instance ID (e.g. INS-001)",
"create_record_mode":   "Mode (hyper/proot/website)",
"create_record_path":   "Path (/ela/instances/arch)",
"create_record_create": "Create",

// Merge Record
"merge_record_title":   "Merge Record Package",
"merge_record_id":      "Instance ID",
"merge_record_rid":     "Record ID",
"merge_record_merge":   "Merge",

// Settings
"settings_title":         "Settings",
"settings_gui":           "GUI Preferred",
"settings_silent":        "Silent Boot",
"settings_auto_inject":   "Auto Inject Driver",
"settings_network":       "Network Direct Mode",
"settings_save":          "Save Settings",
"settings_save_ok":       "Settings saved",
"settings_reset":         "Reset to Default",
"settings_reset_ok":      "Reset to default settings",

// Common
"ok":     "OK",
"cancel": "Cancel",
"error":  "Error",
"success": "Success",
"warning": "Warning",
"info":    "Info",

// Errors
"err_fill_fields":  "Please fill in all fields",
"err_create_fail":  "Create failed",
"err_merge_fail":   "Merge failed",
"err_save_fail":    "Save failed: %v",
"err_not_found":    "Not found",
}
