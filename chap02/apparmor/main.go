package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// In this section, you will look at a different way of reading system-level information.
const (
	appArmorEnablePath = "/sys/module/apparmor/parameters/enabled"
	appArmorModePath   = "/sys/module/apparmor/parameters/mode"
)

func appArmorMode() (mode string) {
	content, err := os.ReadFile(appArmorModePath)
	if err != nil {
		log.Fatal("error (Armor Mode) : ", err)
	}
	return strings.TrimSpace(string(content))
}

func appArmorEnabled() (support bool) {
	content, err := os.ReadFile(appArmorEnablePath)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return
		}
		log.Fatal("error (Armor Enabled)")
		return
	}
	return strings.TrimSpace(string(content)) == "Y"
}

func main() {
	fmt.Println("AppArmor mode : ", appArmorMode())
	fmt.Println("ArrArmor is enabled : ", appArmorEnabled())
}
