package i18n

var Zh = map[string]string{
// 许可证
"license_title":     "ELA 许可证协议",
"license_text":      "本软件按 AS IS 提供，不提供任何担保。\n使用本软件导致的任何损失，开发者概不负责。\n\n源代码：https://github.com/allow-make/Efficient-Linux-Launch-ELA-\n\n按 Enter 同意并继续，按 Delete 退出",
"license_accept":    "同意并继续",
"license_reject":    "退出",

// 主菜单
"menu_title":        "ELA 启动向导",
"menu_download":     "下载 Linux 系统",
"menu_existing":     "使用已有系统",
"menu_usb":          "从 U 盘启动",
"menu_registry":     "注册表管理",
"menu_settings":     "设置",
"menu_exit":         "退出",
"menu_version":      "版本",

// 下载
"download_title":       "选择要下载的发行版",
"download_confirm":     "确认下载",
"download_confirm_msg": "下载 %s %s (%s)?\n大小: %d MB",
"download_start":       "开始下载 %s...",
"download_progress":    "%s %.1f%% (%.2f MB/s)",
"download_done":        "完成",
"download_done_msg":    "%s 下载完成！",
"download_fail":        "下载失败",

// 已有系统
"existing_title":       "使用已有系统",
"existing_info":        "请选择已存在的 Linux 镜像文件或 .elra 包",
"existing_browse":      "浏览",
"existing_load":        "加载",
"existing_selected":    "已选择",
"existing_load_msg":    "加载: %s",

// USB
"usb_title":            "从 U 盘启动",
"usb_info":             "检测到的可引导 USB 设备：",
"usb_confirm":          "从 U 盘启动",
"usb_confirm_msg":      "确定从 %s 启动吗？",
"usb_start":            "正在从 %s 启动...",

// 注册表
"registry_title":       "注册表管理",
"registry_delete_confirm": "删除实例",
"registry_delete_msg":  "确定删除实例 %s (%s) 吗？",
"registry_create":      "创建记录包",
"registry_merge":       "合并记录包",
"registry_refresh":     "刷新",
"registry_back":        "返回",

// 创建记录包
"create_record_title":  "创建记录包",
"create_record_id":     "实例 ID (如 INS-001)",
"create_record_mode":   "模式 (hyper/proot/website)",
"create_record_path":   "路径 (/ela/instances/arch)",
"create_record_create": "创建",

// 合并记录包
"merge_record_title":   "合并记录包",
"merge_record_id":      "实例 ID",
"merge_record_rid":     "记录包 ID",
"merge_record_merge":   "合并",

// 设置
"settings_title":         "设置",
"settings_gui":           "图形界面优先启动",
"settings_silent":        "静默启动",
"settings_auto_inject":   "自动注入驱动",
"settings_network":       "网络直连模式",
"settings_save":          "保存设置",
"settings_save_ok":       "设置已保存",
"settings_reset":         "重置为默认",
"settings_reset_ok":      "已重置为默认设置",

// 通用
"ok":     "确定",
"cancel": "取消",
"error":  "错误",
"success": "成功",
"warning": "警告",
"info":    "提示",

// 错误
"err_fill_fields":  "请填写所有字段",
"err_create_fail":  "创建失败",
"err_merge_fail":   "合并失败",
"err_save_fail":    "保存失败: %v",
"err_not_found":    "未找到",
}
