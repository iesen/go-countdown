package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"io/ioutil"
)

func main() {
	systray.Run(onReady, onExit)
	fmt.Println("Hey")
}

func onExit() {

}

func onReady() {
	iconBytes := getIcon("assets/countdown.ico")
	// time.Sleep(100 * time.Millisecond)
	systray.SetIcon(iconBytes)
	systray.SetTitle("Countdown")
	systray.SetTooltip("Countdown app")
	quitMenu := systray.AddMenuItem("Quit", "Quit app")
	go func() {
		<-quitMenu.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()
}

func getIcon(iconPath string) []byte {
	b, err := ioutil.ReadFile(iconPath)
	if err != nil {
		fmt.Print(err)
	}
	return b
}