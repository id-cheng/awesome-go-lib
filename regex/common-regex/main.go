package main

import (
	"fmt"
	cregex "github.com/mingrammer/commonregex"
)

func main() {
	text := `John, please get that article on www.linkedin.com to me by 5:00PM on Jan 9th 2012. 4:00 would be ideal, actually. If you have any questions, You can reach me at (519)-236-2723x341 or get in touch with my associate at harold.smith@gmail.com`
	dateList := cregex.Date(text)
	// ['Jan 9th 2012']
	timeList := cregex.Time(text)
	// ['5:00PM', '4:00']
	linkList := cregex.Links(text)
	// ['www.linkedin.com', 'harold.smith@gmail.com']
	phoneList := cregex.PhonesWithExts(text)
	// ['(519)-236-2723x341']
	emailList := cregex.Emails(text)
	// ['harold.smith@gmail.com']
	fmt.Println(dateList)
	fmt.Println(timeList)
	fmt.Println(linkList)
	fmt.Println(phoneList)
	fmt.Println(emailList)
}
