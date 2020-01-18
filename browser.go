package browser

import (
	"errors"
	"log"
	"os/exec"
	"runtime"
)


func getBrowserHelper() (string, error) {
	var helpers = []string{
		"xdg-open",
		"sensible-browser",
		"gnome-open",
		"kde-open",
		"epiphany",
		"firefox",
		"mozilla",
		"konqueror",
		"opera",
	}

	for _, h := range helpers {
		err := exec.Command("which", h).Start()
		if err == nil {
			return h, nil
		}
	}
	return "", errors.New("unable to find a suitable browser on your OS")
}

func OpenUrl(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		h, err := getBrowserHelper()
		if err != nil {
			log.Printf("%v\nPlease open your browser manually.\n", err)
			return
		}
		err = exec.Command(h, url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		log.Printf("unsupported platform: %v. Please open your browser manually.\n", runtime.GOOS)
		return
	}
	if err != nil {
		log.Fatal(err)
	}
}
