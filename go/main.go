package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

func fancyMessage(level int, input string) {
	if level == 0 {
		fmt.Printf("  [\033[32m+\033[0m] INFO: %s\n", input)
	} else if level == 1 {
		fmt.Printf("  [\033[33m*\033[0m] WARNING: %s\n", input)
	} else if level == 2 {
		fmt.Printf("  [\033[31m!\033[0m] ERROR: %s\n", input)
		os.Exit(1)
	} else {
		fmt.Printf("Level does not exist: %d\n", level)
		os.Exit(1)
	}
}

func executeWrapper(name string, args ...string) string {
	out, err := exec.Command(name, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(string(out), "\n", "", -1)
}

func main() {
	fmt.Println("Rolling Rhino 🦏")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	help := flag.Bool("h", false, "displays help page")
	force := flag.Bool("f", false, "(force) skips confirmations")
	docker := flag.Bool("d", false, "skips desktop detection for Docker")
	flag.Parse()
	if *help {
		fmt.Printf("  Please visit the following for more information...\n  https://github.com/wimpysworld/rolling-rhino\n")
		os.Exit(0)
	}

	if runtime.GOOS != "linux" {
		fancyMessage(2, "This application only works on Ubuntu systems.")
	}
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	if user.Uid != "0" {
		fancyMessage(2, "You need to be root. Current Uid: "+user.Uid)
	}

	err = exec.Command("which", "lsb_release").Run()
	if err != nil {
		fancyMessage(2, "lsb_release not detected. Quitting.")
	}
	fancyMessage(0, "lsb_release detected.")

	contents, err := ioutil.ReadFile("/etc/apt/sources.list")
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(string(contents), "devel") {
		fancyMessage(2, "Already tracking the devel series. Nothing to do.")
	}

	stdout := executeWrapper("lsb_release", "--id", "--short")
	if stdout == "Ubuntu" {
		fancyMessage(0, "Ubuntu detected.")
	} else {
		fancyMessage(2, stdout+" detected, which is not supported.")
	}

	stdout = executeWrapper("lsb_release", "--description", "--short")
	if strings.Contains(stdout, "development branch") {
		fancyMessage(0, stdout+" detected.")
	} else if strings.Contains(stdout, "LTS") {
		fancyMessage(2, stdout+" detected. Switching an LTS release to the devel series directly is not supported.")
	} else {
		fancyMessage(2, stdout+" detected. Switching an interim release to the devel series directly is not supported.")
	}

	found := false
	if !*docker {
		desktops := [9]string{
			"kubuntu-desktop", "lubuntu-desktop", "ubuntu-desktop", "ubuntu-budgie-desktop", "ubuntukylin-desktop", "ubuntu-mate-desktop", "ubuntustudio-desktop", "xubuntu-desktop", "ubuntu-wsl",
		}
		for _, desktop := range desktops {
			stdout = executeWrapper("env", "LANG=C", "apt", "list", "--installed", desktop, "2>/dev/null")
			if strings.Contains(stdout, "installed") {
				fancyMessage(0, "Detected "+desktop)
				found = true
				break
			}
		}
		if !found {
			fancyMessage(2, "No installed desktop packages were detected. Quitting.")
		}
	}

	found = false
	files, err := ioutil.ReadDir("/etc/apt/sources.list.d/")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".list" {
			found = true
			break
		}
	}
	if found {
		fancyMessage(1, "PPAs detected, you're responsible for taking care of PPA migrations in the future.")
	} else {
		fancyMessage(0, "No PPAs detected, this is good.")
	}

	fancyMessage(0, "All checks passed.")
	if !*force {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("\nAre you sure you want to start tracking the devel series? [y/N] ")
		char, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		if char != 121 {
			os.Exit(1)
		}
	}

	contents, err = ioutil.ReadFile("/etc/apt/sources.list")
	if err != nil {
		log.Fatal(err)
	}
	oldList := fmt.Sprintf("/etc/apt/sources.list.%s", executeWrapper("lsb_release", "--codename", "--short"))
	err = ioutil.WriteFile(oldList, contents, 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("/etc/apt/sources.list", []byte(devel), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fancyMessage(0, "Switching to devel series.")
	tasks := [5]string{
		"autoclean", "clean", "update", "dist-upgrade", "autoremove",
	}
	for _, task := range tasks {
		err = exec.Command("apt", "-y", task).Run()
		if err != nil {
			fmt.Printf("Found error, but continuing: %s\n", err)
		}
		fancyMessage(0, "Finished task: "+task)
	}

	// TODO: Add logo printing without relying on path.
	fancyMessage(0, "Your Rolling Rhino is ready.")
}

var devel string = `
# See http://help.ubuntu.com/community/UpgradeNotes for how to upgrade to
# newer versions of the distribution.
deb http://archive.ubuntu.com/ubuntu devel main restricted
# deb-src http://archive.ubuntu.com/ubuntu devel main restricted
## Major bug fix updates produced after the final release of the
## distribution.
deb http://archive.ubuntu.com/ubuntu devel-updates main restricted
# deb-src http://archive.ubuntu.com/ubuntu devel-updates main restricted
## N.B. software from this repository is ENTIRELY UNSUPPORTED by the Ubuntu
## team. Also, please note that software in universe WILL NOT receive any
## review or updates from the Ubuntu security team.
deb http://archive.ubuntu.com/ubuntu devel universe
# deb-src http://archive.ubuntu.com/ubuntu devel universe
deb http://archive.ubuntu.com/ubuntu devel-updates universe
# deb-src http://archive.ubuntu.com/ubuntu devel-updates universe
## N.B. software from this repository is ENTIRELY UNSUPPORTED by the Ubuntu
## team, and may not be under a free licence. Please satisfy yourself as to
## your rights to use the software. Also, please note that software in
## multiverse WILL NOT receive any review or updates from the Ubuntu
## security team.
deb http://archive.ubuntu.com/ubuntu devel multiverse
# deb-src http://archive.ubuntu.com/ubuntu devel multiverse
deb http://archive.ubuntu.com/ubuntu devel-updates multiverse
# deb-src http://archive.ubuntu.com/ubuntu devel-updates multiverse
## N.B. software from this repository may not have been tested as
## extensively as that contained in the main release, although it includes
## newer versions of some applications which may provide useful features.
## Also, please note that software in backports WILL NOT receive any review
## or updates from the Ubuntu security team.
deb http://archive.ubuntu.com/ubuntu devel-backports main restricted universe multiverse
# deb-src http://archive.ubuntu.com/ubuntu devel-backports main restricted universe multiverse
## Uncomment the following two lines to add software from Canonical's
## 'partner' repository.
## This software is not part of Ubuntu, but is offered by Canonical and the
## respective vendors as a service to Ubuntu users.
# deb http://archive.canonical.com/ubuntu devel partner
# deb-src http://archive.canonical.com/ubuntu devel partner
deb http://security.ubuntu.com/ubuntu devel-security main restricted
# deb-src http://security.ubuntu.com/ubuntu devel-security main restricted
deb http://security.ubuntu.com/ubuntu devel-security universe
# deb-src http://security.ubuntu.com/ubuntu devel-security universe
deb http://security.ubuntu.com/ubuntu devel-security multiverse
# deb-src http://security.ubuntu.com/ubuntu devel-security multiverse
`
