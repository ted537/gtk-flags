package main

import (
	"flag"
	"fmt"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

type boolFlag interface {
	flag.Value
	IsBoolFlag() bool
}

func widgetForStringFlag(
	f *flag.Flag,
) (gtk.IWidget, func() string) {
	widgetV := gtk.NewVBox(false, 0)
	label := gtk.NewLabel(f.Name)
	label.ModifyFontEasy("bold 24px")
	widgetV.Add(label)
	widget := gtk.NewHBox(false, 0)
	entry := gtk.NewEntry()
	entry.SetText(f.Value.String())
	widget.Add(entry)
	widgetV.Add(widget)
	//framebox1.PackStart(widgetV, false, false, 0)

	button := gtk.NewButtonWithLabel("Choose File/Folder")
	button.Clicked(func() {
		//--------------------------------------------------------
		// GtkFileChooserDialog
		//--------------------------------------------------------
		filechooserdialog := gtk.NewFileChooserDialog(
			"Choose File...",
			button.GetTopLevelAsWindow(),
			gtk.FILE_CHOOSER_ACTION_OPEN,
			gtk.STOCK_CANCEL,
			gtk.RESPONSE_CANCEL,
			gtk.STOCK_OK,
			gtk.RESPONSE_ACCEPT)

		rt := filechooserdialog.Run()
		if rt == gtk.RESPONSE_ACCEPT {
			entry.SetText(filechooserdialog.GetFilename())
		}
		filechooserdialog.Destroy()
	})
	widget.Add(button)

	desc := gtk.NewLabel(f.Usage)
	widgetV.Add(desc)

	return widgetV, entry.GetText
}

func widgetForBoolFlag(
	f *flag.Flag, value boolFlag,
) (gtk.IWidget, func() string) {
	widgetV := gtk.NewVBox(false, 0)
	widgetH := gtk.NewHBox(false, 0)

	label := gtk.NewLabel(f.Name)
	label.ModifyFontEasy("bold 24px")

	widgetV.PackStart(widgetH, false, false, 0)

	checkButton := gtk.NewCheckButton()
	widgetH.PackStart(checkButton, false, false, 20)
	widgetH.PackStart(label, false, false, 0)

	descRow := gtk.NewHBox(false, 0)
	widgetV.PackStart(descRow, false, false, 0)
	descRow.PackStart(gtk.NewLabel(f.Usage), false, false, 20)

	getter := func() string {
		if checkButton.GetActive() {
			return "true"
		} else {
			return "false"
		}
	}

	return widgetV, getter
}

func widgetForFlag(f *flag.Flag) (gtk.IWidget, func() string) {
	bf, ok := f.Value.(boolFlag)
	if ok && bf.IsBoolFlag() {
		return widgetForBoolFlag(f, bf)
	}

	return widgetForStringFlag(f)
}

type FlagValueGetters map[*flag.Flag]func() string

func RunGui(appName string, exec func() error) {
	flagValueGetters := make(FlagValueGetters)

	//var menuitem *gtk.MenuItem
	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle(appName)
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		println("got destroy!", ctx.Data().(string))
		gtk.MainQuit()
	}, "foo")

	//--------------------------------------------------------
	// GtkVBox
	//--------------------------------------------------------
	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)
	vbox := gtk.NewVBox(false, 1)
	swin.AddWithViewPort(vbox)

	//--------------------------------------------------------
	// GtkMenuBar
	//--------------------------------------------------------
	menubar := gtk.NewMenuBar()
	vbox.PackStart(menubar, false, false, 0)

	//--------------------------------------------------------
	// GtkVPaned
	//--------------------------------------------------------
	vpaned := gtk.NewVPaned()
	vbox.Add(vpaned)

	//--------------------------------------------------------
	// GtkFrame
	//--------------------------------------------------------
	frame1 := gtk.NewFrame("Inputs")
	framebox1 := gtk.NewVBox(false, 1)
	//framebox1.Add(gtk.NewFixed(0, 1, 1, 1))
	frame1.Add(framebox1)

	frame2 := gtk.NewFrame("Execute")
	framebox2 := gtk.NewVBox(false, 1)
	frame2.Add(framebox2)

	//runBoxV := gtk.NewVBox(false, 0)
	runBoxH := gtk.NewHBox(false, 0)
	runButton := gtk.NewButtonWithLabel("Run")
	runBoxH.Add(gtk.NewHBox(false, 0))
	runBoxH.Add(runButton)
	runBoxH.Add(gtk.NewHBox(false, 0))
	framebox2.PackStart(runBoxH, false, false, 0)

	vpaned.Pack1(frame1, false, false)
	vpaned.Pack2(frame2, false, false)

	flags := []*flag.Flag{}
	flag.VisitAll(func(someFlag *flag.Flag) {
		flags = append(flags, someFlag)
	})

	for idx, someFlag := range flags {
		widget, getter := widgetForFlag(someFlag)
		flagValueGetters[someFlag] = getter

		framebox1.PackStart(widget, false, false, 5)

		if idx+1 < len(flags)-1 {
			// sepBox := gtk.NewVBox(false, 20)
			// sep := gtk.NewHSeparator()
			// sepBox.Add(sep)
			// framebox1.PackStart(sepBox, false, true, 120)
		}
	}

	runButton.Clicked(func() {
		// Set flags
		for _, flag := range flags {
			getter, ok := flagValueGetters[flag]
			if !ok {
				panic("Could not find getter for flag " + flag.Name)
			}
			err := flag.Value.Set(getter())
			if err != nil {
				fmt.Println(err)
			}

		}

		// Execute program.
		err := exec()

		mt := gtk.MESSAGE_INFO
		msg := "Operation executed succesfully"
		if err != nil {
			msg = fmt.Sprintf("Error: %s", err.Error())
			mt = gtk.MESSAGE_ERROR
		}

		dlg := gtk.NewMessageDialog(
			window, gtk.DIALOG_DESTROY_WITH_PARENT, mt, gtk.BUTTONS_OK,
			msg,
		)
		dlg.Response(func() {
			dlg.Destroy()
		})
		dlg.Show()
	})

	//--------------------------------------------------------
	// Event
	//--------------------------------------------------------
	window.Add(swin)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
