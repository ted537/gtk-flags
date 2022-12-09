package main

import (
	"flag"
	"fmt"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

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
	vbox := gtk.NewVBox(false, 1)

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

	vpaned.Pack1(frame1, false, false)
	vpaned.Pack2(frame2, false, false)

	//--------------------------------------------------------
	// GtkImage
	//--------------------------------------------------------
	//dir, _ := path.Split(os.Args[0])
	//imagefile := path.Join(dir, "../../data/go-gtk-logo.png")

	//label := gtk.NewLabel("Go Binding for GTK")
	//label.ModifyFontEasy("DejaVu Serif 15")
	//framebox1.PackStart(label, false, true, 0)

	//--------------------------------------------------------
	// GtkEntry
	//--------------------------------------------------------
	entry := gtk.NewEntry()
	entry.SetText("Hello world")
	widget := gtk.NewHBox(false, 0)
	widget.Add(entry)
	framebox1.PackStart(widget, false, false, 0)

	//image := gtk.NewImageFromFile(imagefile)
	//framebox1.Add(image)

	// //--------------------------------------------------------
	// // GtkScale
	// //--------------------------------------------------------
	// scale := gtk.NewHScaleWithRange(0, 100, 1)
	// scale.Connect("value-changed", func() {
	// 	println("scale:", int(scale.GetValue()))
	// })
	// framebox2.Add(scale)

	// //--------------------------------------------------------
	// // GtkHBox
	// //--------------------------------------------------------
	//buttons := gtk.NewHBox(false, 1)

	//--------------------------------------------------------
	// GtkButton
	//--------------------------------------------------------
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
	//vpaned.Add(buttons)

	// //--------------------------------------------------------
	// // GtkFontButton
	// //--------------------------------------------------------
	// fontbutton := gtk.NewFontButton()
	// fontbutton.Connect("font-set", func() {
	// 	println("title:", fontbutton.GetTitle())
	// 	println("fontname:", fontbutton.GetFontName())
	// 	println("use_size:", fontbutton.GetUseSize())
	// 	println("show_size:", fontbutton.GetShowSize())
	// })
	// buttons.Add(fontbutton)
	// framebox2.PackStart(buttons, false, false, 0)

	// buttons = gtk.NewHBox(false, 1)

	// //--------------------------------------------------------
	// // GtkToggleButton
	// //--------------------------------------------------------
	// togglebutton := gtk.NewToggleButtonWithLabel("ToggleButton with label")
	// togglebutton.Connect("toggled", func() {
	// 	if togglebutton.GetActive() {
	// 		togglebutton.SetLabel("ToggleButton ON!")
	// 	} else {
	// 		togglebutton.SetLabel("ToggleButton OFF!")
	// 	}
	// })
	// buttons.Add(togglebutton)

	// //--------------------------------------------------------
	// // GtkCheckButton
	// //--------------------------------------------------------
	// checkbutton := gtk.NewCheckButtonWithLabel("CheckButton with label")
	// checkbutton.Connect("toggled", func() {
	// 	if checkbutton.GetActive() {
	// 		checkbutton.SetLabel("CheckButton CHECKED!")
	// 	} else {
	// 		checkbutton.SetLabel("CheckButton UNCHECKED!")
	// 	}
	// })
	// buttons.Add(checkbutton)

	// //--------------------------------------------------------
	// // GtkRadioButton
	// //--------------------------------------------------------
	// buttonbox := gtk.NewVBox(false, 1)
	// radiofirst := gtk.NewRadioButtonWithLabel(nil, "Radio1")
	// buttonbox.Add(radiofirst)
	// buttonbox.Add(gtk.NewRadioButtonWithLabel(radiofirst.GetGroup(), "Radio2"))
	// buttonbox.Add(gtk.NewRadioButtonWithLabel(radiofirst.GetGroup(), "Radio3"))
	// buttons.Add(buttonbox)
	// //radiobutton.SetMode(false);
	// radiofirst.SetActive(true)

	// framebox2.PackStart(buttons, false, false, 0)

	// //--------------------------------------------------------
	// // GtkVSeparator
	// //--------------------------------------------------------
	// vsep := gtk.NewVSeparator()
	// framebox2.PackStart(vsep, false, false, 0)

	// //--------------------------------------------------------
	// // GtkComboBoxEntry
	// //--------------------------------------------------------
	// combos := gtk.NewHBox(false, 1)
	// comboboxentry := gtk.NewComboBoxEntryNewText()
	// comboboxentry.AppendText("Monkey")
	// comboboxentry.AppendText("Tiger")
	// comboboxentry.AppendText("Elephant")
	// comboboxentry.Connect("changed", func() {
	// 	println("value:", comboboxentry.GetActiveText())
	// })
	// combos.Add(comboboxentry)

	// //--------------------------------------------------------
	// // GtkComboBox
	// //--------------------------------------------------------
	// combobox := gtk.NewComboBoxNewText()
	// combobox.AppendText("Peach")
	// combobox.AppendText("Banana")
	// combobox.AppendText("Apple")
	// combobox.SetActive(1)
	// combobox.Connect("changed", func() {
	// 	println("value:", combobox.GetActiveText())
	// })
	// combos.Add(combobox)

	// framebox2.PackStart(combos, false, false, 0)

	// //--------------------------------------------------------
	// // GtkTextView
	// //--------------------------------------------------------
	// swin := gtk.NewScrolledWindow(nil, nil)
	// swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	// swin.SetShadowType(gtk.SHADOW_IN)
	// textview := gtk.NewTextView()
	// var start, end gtk.TextIter
	// buffer := textview.GetBuffer()
	// buffer.GetStartIter(&start)
	// buffer.Insert(&start, "Hello ")
	// buffer.GetEndIter(&end)
	// buffer.Insert(&end, "World!")
	// var props map[string]interface{} = map[string]interface{}{
	// 	"background": "#FF0000", "weight": "700"}
	// tag := buffer.CreateTag("bold", props)
	// buffer.GetStartIter(&start)
	// buffer.GetEndIter(&end)
	// buffer.ApplyTag(tag, &start, &end)
	// swin.Add(textview)
	// framebox2.Add(swin)

	// buffer.Connect("changed", func() {
	// 	println("changed")
	// })

	// //--------------------------------------------------------
	// // GtkMenuItem
	// //--------------------------------------------------------
	// cascademenu := gtk.NewMenuItemWithMnemonic("_File")
	// menubar.Append(cascademenu)
	// submenu := gtk.NewMenu()
	// cascademenu.SetSubmenu(submenu)

	// menuitem = gtk.NewMenuItemWithMnemonic("E_xit")
	// menuitem.Connect("activate", func() {
	// 	gtk.MainQuit()
	// })
	// submenu.Append(menuitem)

	// cascademenu = gtk.NewMenuItemWithMnemonic("_View")
	// menubar.Append(cascademenu)
	// submenu = gtk.NewMenu()
	// cascademenu.SetSubmenu(submenu)

	// checkmenuitem := gtk.NewCheckMenuItemWithMnemonic("_Disable")
	// checkmenuitem.Connect("activate", func() {
	// 	vpaned.SetSensitive(!checkmenuitem.GetActive())
	// })
	// submenu.Append(checkmenuitem)

	// menuitem = gtk.NewMenuItemWithMnemonic("_Font")
	// menuitem.Connect("activate", func() {
	// 	fsd := gtk.NewFontSelectionDialog("Font")
	// 	fsd.SetFontName(fontbutton.GetFontName())
	// 	fsd.Response(func() {
	// 		println(fsd.GetFontName())
	// 		fontbutton.SetFontName(fsd.GetFontName())
	// 		fsd.Destroy()
	// 	})
	// 	fsd.SetTransientFor(window)
	// 	fsd.Run()
	// })
	// submenu.Append(menuitem)

	// cascademenu = gtk.NewMenuItemWithMnemonic("_Help")
	// menubar.Append(cascademenu)
	// submenu = gtk.NewMenu()
	// cascademenu.SetSubmenu(submenu)

	// menuitem = gtk.NewMenuItemWithMnemonic("_About")
	// menuitem.Connect("activate", func() {
	// 	dialog := gtk.NewAboutDialog()
	// 	dialog.SetName("Go-Gtk Demo!")
	// 	dialog.SetProgramName("demo")
	// 	dialog.SetAuthors(authors())
	// 	dir, _ := path.Split(os.Args[0])
	// 	imagefile := path.Join(dir, "../../data/mattn-logo.png")
	// 	pixbuf, _ := gdkpixbuf.NewPixbufFromFile(imagefile)
	// 	dialog.SetLogo(pixbuf)
	// 	dialog.SetLicense("The library is available under the same terms and conditions as the Go, the BSD style license, and the LGPL (Lesser GNU Public License). The idea is that if you can use Go (and Gtk) in a project, you should also be able to use go-gtk.")
	// 	dialog.SetWrapLicense(true)
	// 	dialog.Run()
	// 	dialog.Destroy()
	// })
	// submenu.Append(menuitem)

	// //--------------------------------------------------------
	// // GtkStatusbar
	// //--------------------------------------------------------
	// statusbar := gtk.NewStatusbar()
	// context_id := statusbar.GetContextId("go-gtk")
	// statusbar.Push(context_id, "GTK binding for Go!")

	//framebox2.PackStart(statusbar, false, false, 0)

	flag.VisitAll(func(someFlag *flag.Flag) {
		fmt.Println(someFlag.Name)
		fmt.Println(someFlag.Value)
		//fmt.Println(someFlag.)
	})

	//--------------------------------------------------------
	// Event
	//--------------------------------------------------------
	window.Add(vbox)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
