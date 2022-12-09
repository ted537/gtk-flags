package main

import (
	"flag"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func widgetForFlag(f *flag.Flag) gtk.IWidget {
	widgetV := gtk.NewVBox(false, 0)
	label := gtk.NewLabel(f.Name)
	label.ModifyFontEasy("bold")
	widgetV.Add(label)
	widget := gtk.NewHBox(false, 0)
	entry := gtk.NewEntry()
	entry.SetText(f.Value.String())
	widget.Add(entry)
	widgetV.Add(widget)
	//framebox1.PackStart(widgetV, false, false, 0)

	button := gtk.NewButtonWithLabel("Choose File/Folder")
	button.Clicked(func() {
		println("Dialog OK!")

		//--------------------------------------------------------
		// GtkFileChooserDialog
		//--------------------------------------------------------
		filechooserdialog := gtk.NewFileChooserDialog(
			"Choose File...",
			button.GetTopLevelAsWindow(),
			gtk.FILE_CHOOSER_ACTION_OPEN,
			gtk.STOCK_OK,
			gtk.RESPONSE_ACCEPT)
		filter := gtk.NewFileFilter()
		filter.AddPattern("*")
		filechooserdialog.AddFilter(filter)
		filechooserdialog.Response(func() {
			entry.SetText(filechooserdialog.GetFilename())
			filechooserdialog.Destroy()
		})
		filechooserdialog.Run()
	})
	widget.Add(button)

	desc := gtk.NewLabel(f.Usage)
	widgetV.Add(desc)

	//f.Value

	return widgetV
}

func RunGui() {
	//var menuitem *gtk.MenuItem
	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("GTK Go!")
	window.SetIconName("gtk-dialog-info")
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

	frame2 := gtk.NewFrame("Output")
	framebox2 := gtk.NewVBox(false, 1)
	frame2.Add(framebox2)

	//runBoxV := gtk.NewVBox(false, 0)
	runBoxH := gtk.NewHBox(false, 0)
	runLabel := gtk.NewButtonWithLabel("Run")
	runBoxH.Add(gtk.NewHBox(false, 0))
	runBoxH.Add(runLabel)
	runBoxH.Add(gtk.NewHBox(false, 0))
	framebox2.PackStart(runBoxH, false, false, 0)

	vpaned.Pack1(frame1, false, false)
	vpaned.Pack2(frame2, false, false)

	flags := []*flag.Flag{}
	flag.VisitAll(func(someFlag *flag.Flag) {
		flags = append(flags, someFlag)
	})

	for idx, someFlag := range flags {
		for c := 0; c < 10; c++ {
			widget := widgetForFlag(someFlag)
			framebox1.PackStart(widget, false, false, 10)

			if idx+1 < len(flags)-1 {
				//framebox1.PackStart(gtk.NewVSeparator(), true, false, 5)
			}
		}
	}

	//--------------------------------------------------------
	// Event
	//--------------------------------------------------------
	window.Add(swin)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
