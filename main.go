package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	// Create a new application
	myApp := app.NewWithID("999")

	// Create a simple icon programmatically
	icon := createSimpleIcon()

	// Create bound strings for the detailed view
	cpuStr := binding.NewString()
	memStr := binding.NewString()
	memDetailsStr := binding.NewString()
	timeStr := binding.NewString()

	// Initialize with placeholder text
	cpuStr.Set("CPU: Loading...")
	memStr.Set("Memory: Loading...")
	memDetailsStr.Set("Details: Loading...")
	timeStr.Set("Last updated: --")

	// Create the detailed window
	detailWindow := myApp.NewWindow("System Monitor - Details")

	title := widget.NewLabel("System Resource Monitor")
	title.TextStyle.Bold = true

	cpuLabel := widget.NewLabelWithData(cpuStr)
	memLabel := widget.NewLabelWithData(memStr)
	memDetailsLabel := widget.NewLabelWithData(memDetailsStr)
	timeLabel := widget.NewLabelWithData(timeStr)

	detailWindow.SetContent(container.NewVBox(
		title,
		// widget.NewSeparator(),
		widget.NewLabel("CPU Information:"),
		cpuLabel,
		widget.NewSeparator(),
		widget.NewLabel("Memory Information:"),
		memLabel,
		memDetailsLabel,
		widget.NewSeparator(),
		timeLabel,
	))

	detailWindow.Resize(fyne.NewSize(450, 300))

	// Don't show the window initially
	detailWindow.SetCloseIntercept(func() {
		detailWindow.Hide()
	})

	// Check if desktop features are supported
	if desk, ok := myApp.(desktop.App); ok {
		// Create system tray menu
		menu := fyne.NewMenu("System Monitor")

		// Create menu items
		cpuItem := fyne.NewMenuItem("CPU: --", nil)
		cpuItem.Disabled = true

		memItem := fyne.NewMenuItem("Memory: --", nil)
		memItem.Disabled = true

		showItem := fyne.NewMenuItem("Show Details", func() {
			detailWindow.Show()
		})

		quitItem := fyne.NewMenuItem("Quit", func() {
			myApp.Quit()
		})

		menu.Items = []*fyne.MenuItem{cpuItem, memItem, showItem, quitItem}

		// Set up the system tray
		desk.SetSystemTrayMenu(menu)
		desk.SetSystemTrayIcon(icon)

		// Update stats in background
		go func() {
			for {
				// Get CPU usage
				cpuPercent, err := cpu.Percent(1*time.Second, false)
				if err == nil && len(cpuPercent) > 0 {
					cpuItem.Label = fmt.Sprintf("CPU: %.1f%%", cpuPercent[0])
					cpuStr.Set(fmt.Sprintf("CPU Usage: %.1f%%", cpuPercent[0]))
					menu.Refresh()
				}

				// Get memory usage
				vmem, err := mem.VirtualMemory()
				if err == nil {
					usedGB := float64(vmem.Used) / 1024 / 1024 / 1024
					totalGB := float64(vmem.Total) / 1024 / 1024 / 1024
					availableGB := float64(vmem.Available) / 1024 / 1024 / 1024

					memItem.Label = fmt.Sprintf("Memory: %.1f%%", vmem.UsedPercent)
					memStr.Set(fmt.Sprintf("Memory Usage: %.1f%%", vmem.UsedPercent))
					memDetailsStr.Set(fmt.Sprintf("Used: %.2f GB / Total: %.2f GB\nAvailable: %.2f GB",
						usedGB, totalGB, availableGB))
					menu.Refresh()
				}

				// Update time
				currentTime := time.Now().Format("15:04:05")
				timeStr.Set(fmt.Sprintf("Last updated: %s", currentTime))

				time.Sleep(2 * time.Second)
			}
		}()
	}

	myApp.Run()
}

func createSimpleIcon() *fyne.StaticResource {
	// Create a 32x32 image
	img := image.NewRGBA(image.Rect(0, 0, 32, 32))

	// Fill with a color (blue square with white border)
	blue := color.RGBA{41, 128, 185, 255}
	white := color.RGBA{255, 255, 255, 255}

	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x++ {
			if x < 2 || x > 29 || y < 2 || y > 29 {
				img.Set(x, y, white)
			} else {
				img.Set(x, y, blue)
			}
		}
	}

	// Encode to PNG
	var buf bytes.Buffer
	png.Encode(&buf, img)

	return fyne.NewStaticResource("icon.png", buf.Bytes())
}
