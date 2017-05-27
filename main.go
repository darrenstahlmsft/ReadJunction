package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	junctionPath := "test_junction"

	// Replace this volume GUID with the guid of your C:\ drive found using `mountvol.exe`
	folderPath := `\\?\Volume{7d396648-0527-4fed-9cae-5d961ec7311a}\`

	os.Remove(junctionPath)

	cmd := exec.Command("cmd", "/c", "mklink", "/J", junctionPath, folderPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))

	dest, err := os.Readlink(junctionPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Readlink returned:", dest)
}
