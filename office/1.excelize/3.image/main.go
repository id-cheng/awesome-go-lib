package main

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, _ := excelize.OpenFile("Book.xlsx")
	defer f.Close()
	// Insert a picture.
	f.AddPicture("Sheet1", "A2", "image.png", nil)
	// Insert a picture to worksheet with scaling.
	f.AddPicture("Sheet1", "D2", "image.jpg",
		&excelize.GraphicOptions{ScaleX: 0.5, ScaleY: 0.5})
	// Insert a picture offset in the cell with printing support.
	enable, disable := true, false
	f.AddPicture("Sheet1", "H2", "image.gif",
		&excelize.GraphicOptions{
			PrintObject:     &enable,
			LockAspectRatio: false,
			OffsetX:         15,
			OffsetY:         10,
			Locked:          &disable,
		})
	// Save the spreadsheet with the origin path.
	f.Save()
}
