package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"ela/pkg/downloader"
	"ela/pkg/elra"
	"ela/pkg/i18n"
	"ela/pkg/settings"
	"ela/pkg/version"
)

type ELAUI struct {
	app      fyne.App
	window   fyne.Window
	settings *settings.SettingsManager
}

func NewELAUI(sm *settings.SettingsManager) *ELAUI {
	return &ELAUI{
		settings: sm,
	}
}

func (u *ELAUI) Start() {
	u.app = app.NewWithID("ela.launcher")
	title := "ELA - Efficient Linux Launcher"
	if i18n.GetLang() == "zh" {
		title = "ELA - 高效 Linux 启动器"
	}
	u.window = u.app.NewWindow(title)
	u.window.Resize(fyne.NewSize(620, 480))

	u.showLicense()
	u.window.ShowAndRun()
}

func (u *ELAUI) showLicense() {
	title := widget.NewLabel(i18n.T("license_title"))
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	text := widget.NewLabel(i18n.T("license_text"))
	text.Wrapping = fyne.TextWrapWord
	text.Alignment = fyne.TextAlignCenter

	enterBtn := widget.NewButton(i18n.T("license_accept"), func() {
		u.showMainMenu()
	})
	exitBtn := widget.NewButton(i18n.T("license_reject"), func() {
		u.app.Quit()
	})

	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		text,
		container.NewHBox(layout.NewSpacer(), enterBtn, exitBtn, layout.NewSpacer()),
	)

	u.window.SetContent(content)
}

func (u *ELAUI) showMainMenu() {
	title := widget.NewLabel(i18n.T("menu_title"))
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	ver := widget.NewLabel(i18n.T("menu_version") + ": " + version.String())
	ver.Alignment = fyne.TextAlignCenter

	btn1 := widget.NewButton(i18n.T("menu_download"), u.showDownloadMenu)
	btn2 := widget.NewButton(i18n.T("menu_existing"), u.showExistingMenu)
	btn3 := widget.NewButton(i18n.T("menu_usb"), u.showUSBMenu)
	btn4 := widget.NewButton(i18n.T("menu_registry"), u.showRegistryMenu)
	btn5 := widget.NewButton(i18n.T("menu_settings"), u.showSettingsMenu)
	btn6 := widget.NewButton(i18n.T("menu_exit"), func() { u.app.Quit() })

	content := container.NewVBox(
		title,
		ver,
		widget.NewSeparator(),
		btn1,
		btn2,
		btn3,
		btn4,
		btn5,
		widget.NewSeparator(),
		btn6,
	)

	u.window.SetContent(container.NewCenter(content))
}

// ============================================================
// Download
// ============================================================

func (u *ELAUI) showDownloadMenu() {
	dl := downloader.NewDownloader()
	distros := dl.ListDistros()

	title := widget.NewLabel(i18n.T("download_title"))
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	list := widget.NewList(
		func() int { return len(distros) },
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(i int, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(distros[i].Name + " " + distros[i].Version + " (" + distros[i].Arch + ")")
		},
	)

	var selected int = -1
	list.OnSelected = func(id int) {
		selected = id
		d := distros[id]
		msg := fmt.Sprintf(i18n.T("download_confirm_msg"),
			d.Name, d.Version, d.Arch, d.Size/1024/1024)
		dialog.ShowConfirm(i18n.T("download_confirm"), msg,
			func(confirm bool) {
				if confirm && selected >= 0 {
					u.doDownload(distros[selected])
				}
			}, u.window)
	}

	backBtn := widget.NewButton(i18n.T("registry_back"), u.showMainMenu)

	content := container.NewBorder(title, backBtn, nil, nil, list)
	u.window.SetContent(content)
}

func (u *ELAUI) doDownload(distro downloader.Distro) {
	progress := widget.NewProgressBar()
	label := widget.NewLabel(fmt.Sprintf(i18n.T("download_start"), distro.Name))
	label.Alignment = fyne.TextAlignCenter

	content := container.NewVBox(
		widget.NewLabel("Downloading..."),
		label,
		progress,
	)

	u.window.SetContent(content)

	cfg := u.settings.Get()
	progressChan := make(chan downloader.Progress)

	go func() {
		err := downloader.NewDownloader().Download(distro, cfg.InstallDir, progressChan)
		if err != nil {
			dialog.ShowError(fmt.Errorf("%s: %v", i18n.T("download_fail"), err), u.window)
			return
		}
	}()

	go func() {
		for p := range progressChan {
			progress.SetValue(p.Percent / 100)
			label.SetText(fmt.Sprintf(i18n.T("download_progress"),
				distro.Name, p.Percent, p.Speed/1024/1024))
		}
		dialog.ShowInformation(i18n.T("download_done"),
			fmt.Sprintf(i18n.T("download_done_msg"), distro.Name), u.window)
		u.showMainMenu()
	}()
}

// ============================================================
// Existing System
// ============================================================

func (u *ELAUI) showExistingMenu() {
	title := widget.NewLabel(i18n.T("existing_title"))
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	info := widget.NewLabel(i18n.T("existing_info"))
	info.Wrapping = fyne.TextWrapWord
	info.Alignment = fyne.TextAlignCenter

	entry := widget.NewEntry()
	entry.SetPlaceHolder(i18n.T("existing_browse"))

	loadBtn := widget.NewButton(i18n.T("existing_browse"), func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err == nil && reader != nil {
				entry.SetText(reader.URI().Path())
				reader.Close()
			}
		}, u.window)
	})

	confirmBtn := widget.NewButton(i18n.T("existing_load"), func() {
		if entry.Text != "" {
			dialog.ShowInformation(i18n.T("existing_selected"),
				fmt.Sprintf(i18n.T("existing_load_msg"), entry.Text), u.window)
			u.showMainMenu()
		}
	})

	backBtn := widget.NewButton(i18n.T("registry_back"), u.showMainMenu)

	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		info,
		entry,
		container.NewHBox(layout.NewSpacer(), loadBtn, confirmBtn, backBtn, layout.NewSpacer()),
	)

	u.window.SetContent(content)
}

// ============================================================
// USB
// ============================================================

func (u *ELAUI) showUSBMenu() {
	title := widget.NewLabel(i18n.T("usb_title"))
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	info := widget.NewLabel(i18n.T("usb_info"))
	info.Wrapping = fyne.TextWrapWord
	info.Alignment = fyne.TextAlignCenter

	usbDevices := []string{"D: (Ubuntu 22.04)", "E: (Arch Linux)"}
	list := widget.NewList(
		func() int { return len(usbDevices) },
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(i int, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(usbDevices[i])
		},
	)

	list.OnSelected = func(id int) {
		msg := fmt.Sprintf(i18n.T("usb_confirm_msg"), usbDevices[id])
		dialog.ShowConfirm(i18n.T("usb_confirm"), msg,
			func(confirm bool) {
				if confirm {
					dialog.ShowInformation(i18n.T("usb_confirm"),
						fmt.Sprintf(i18n.T("usb_start"), usbDevices[id]), u.window)
				}
			}, u.window)
	}

	refreshBtn := widget.NewButton(i18n.T("registry_refresh"), u.showUSBMenu)
	backBtn := widget.NewButton(i18n.T("registry_back"), u.showMainMenu)

	content := container.NewBorder(
		title,
		container.NewHBox(layout.NewSpacer(), refreshBtn, backBtn, layout.NewSpacer()),
		nil,
		nil,
		list,
	)

	u.window.SetContent(content)
}

// ============================================================
// Registry
// ============================================================

func (u *ELAUI) showRegistryMenu() {
	title := widget.NewLabel(i18n.T("registry_title"))
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	instances, _ := elra.ListInstances()
	if instances == nil {
		instances = []elra.Instance{}
	}

	list := widget.NewList(
		func() int { return len(instances) },
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(i int, obj fyne.CanvasObject) {
			inst := instances[i]
			obj.(*widget.Label).SetText(fmt.Sprintf("%s (%s) - %s", inst.Name, inst.Arch, inst.Status))
		},
	)

	list.OnSelected = func(id int) {
		inst := instances[id]
		msg := fmt.Sprintf(i18n.T("registry_delete_msg"), inst.Name, inst.Arch)
		dialog.ShowConfirm(i18n.T("registry_delete_confirm"), msg,
			func(confirm bool) {
				if confirm {
					elra.DeleteInstance(inst.ID)
					u.showRegistryMenu()
				}
			}, u.window)
	}

	newBtn := widget.NewButton(i18n.T("registry_create"), u.showCreateRecord)
	mergeBtn := widget.NewButton(i18n.T("registry_merge"), u.showMergeRecord)
	refreshBtn := widget.NewButton(i18n.T("registry_refresh"), u.showRegistryMenu)
	backBtn := widget.NewButton(i18n.T("registry_back"), u.showMainMenu)

	content := container.NewBorder(
		title,
		container.NewHBox(layout.NewSpacer(), newBtn, mergeBtn, refreshBtn, backBtn, layout.NewSpacer()),
		nil,
		nil,
		list,
	)

	u.window.SetContent(content)
}

// ============================================================
// Create Record
// ============================================================

func (u *ELAUI) showCreateRecord() {
	title := widget.NewLabel(i18n.T("create_record_title"))
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	instIDEntry := widget.NewEntry()
	instIDEntry.SetPlaceHolder(i18n.T("create_record_id"))

	modeEntry := widget.NewEntry()
	modeEntry.SetPlaceHolder(i18n.T("create_record_mode"))

	pathEntry := widget.NewEntry()
	pathEntry.SetPlaceHolder(i18n.T("create_record_path"))

	createBtn := widget.NewButton(i18n.T("create_record_create"), func() {
		if instIDEntry.Text == "" || modeEntry.Text == "" || pathEntry.Text == "" {
			dialog.ShowError(fmt.Errorf(i18n.T("err_fill_fields")), u.window)
			return
		}
		rec := &elra.RecordPackage{
			ID:         "REC-" + instIDEntry.Text,
			TargetID:   instIDEntry.Text,
			Mode:       modeEntry.Text,
			Path:       pathEntry.Text,
			AutoInject: true,
		}
		if elra.CreateRecord(rec) == 0 {
			dialog.ShowInformation(i18n.T("success"), i18n.T("settings_save_ok"), u.window)
			u.showRegistryMenu()
		} else {
			dialog.ShowError(fmt.Errorf(i18n.T("err_create_fail")), u.window)
		}
	})

	backBtn := widget.NewButton(i18n.T("registry_back"), u.showRegistryMenu)

	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		widget.NewLabel("ID:"), instIDEntry,
		widget.NewLabel("Mode:"), modeEntry,
		widget.NewLabel("Path:"), pathEntry,
		container.NewHBox(layout.NewSpacer(), createBtn, backBtn, layout.NewSpacer()),
	)

	u.window.SetContent(content)
}

// ============================================================
// Merge Record
// ============================================================

func (u *ELAUI) showMergeRecord() {
	title := widget.NewLabel(i18n.T("merge_record_title"))
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	instIDEntry := widget.NewEntry()
	instIDEntry.SetPlaceHolder(i18n.T("merge_record_id"))

	recIDEntry := widget.NewEntry()
	recIDEntry.SetPlaceHolder(i18n.T("merge_record_rid"))

	mergeBtn := widget.NewButton(i18n.T("merge_record_merge"), func() {
		if instIDEntry.Text == "" || recIDEntry.Text == "" {
			dialog.ShowError(fmt.Errorf(i18n.T("err_fill_fields")), u.window)
			return
		}
		if elra.MergeRecord(instIDEntry.Text, recIDEntry.Text) == 0 {
			dialog.ShowInformation(i18n.T("success"), i18n.T("settings_save_ok"), u.window)
			u.showRegistryMenu()
		} else {
			dialog.ShowError(fmt.Errorf(i18n.T("err_merge_fail")), u.window)
		}
	})

	backBtn := widget.NewButton(i18n.T("registry_back"), u.showRegistryMenu)

	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		widget.NewLabel("Instance ID:"), instIDEntry,
		widget.NewLabel("Record ID:"), recIDEntry,
		container.NewHBox(layout.NewSpacer(), mergeBtn, backBtn, layout.NewSpacer()),
	)

	u.window.SetContent(content)
}

// ============================================================
// Settings
// ============================================================

func (u *ELAUI) showSettingsMenu() {
	title := widget.NewLabel(i18n.T("settings_title"))
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	cfg := u.settings.Get()

	checkGUI := widget.NewCheck(i18n.T("settings_gui"), func(v bool) {
		_ = u.settings.Set("gui_preferred", v)
	})
	checkGUI.Checked = cfg.GUIPreferred

	checkSilent := widget.NewCheck(i18n.T("settings_silent"), func(v bool) {
		_ = u.settings.Set("silent_boot", v)
	})
	checkSilent.Checked = cfg.SilentBoot

	checkInject := widget.NewCheck(i18n.T("settings_auto_inject"), func(v bool) {
		_ = u.settings.Set("auto_inject", v)
	})
	checkInject.Checked = cfg.AutoInject

	checkNetwork := widget.NewCheck(i18n.T("settings_network"), func(v bool) {
		_ = u.settings.Set("network_direct", v)
	})
	checkNetwork.Checked = cfg.NetworkDirect

	saveBtn := widget.NewButton(i18n.T("settings_save"), func() {
		if err := u.settings.Save(); err != nil {
			dialog.ShowError(fmt.Errorf(i18n.T("err_save_fail"), err), u.window)
		} else {
			dialog.ShowInformation(i18n.T("success"), i18n.T("settings_save_ok"), u.window)
		}
	})

	resetBtn := widget.NewButton(i18n.T("settings_reset"), func() {
		if err := u.settings.Reset(); err == nil {
			dialog.ShowInformation(i18n.T("success"), i18n.T("settings_reset_ok"), u.window)
			u.showSettingsMenu()
		}
	})

	backBtn := widget.NewButton(i18n.T("registry_back"), u.showMainMenu)

	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		checkGUI,
		checkSilent,
		checkInject,
		checkNetwork,
		container.NewHBox(layout.NewSpacer(), saveBtn, resetBtn, backBtn, layout.NewSpacer()),
	)

	u.window.SetContent(content)
}
