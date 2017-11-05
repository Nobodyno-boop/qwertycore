package qwertycore

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

// User
type User interface {
	mkdir(i string)
	del(i string)
}

// Oss
type Oss struct {
	Os   string
	Path string
	Qw   string
}

func (o Oss) createFile(i string, b []byte) {
	if o.Os == "windows" {
		err := ioutil.WriteFile(o.Qw+`\`+i, b, 0644)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		err := ioutil.WriteFile(o.Qw+"/"+i, b, 0644)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

//c
func (o Oss) mkdir(i string) {
	if o.Os == "windows" {
		strings.Replace(i, `/`, `\`, -1)
	}
	if _, err := os.Stat(o.Qw + i); os.IsNotExist(err) {
		os.Mkdir(o.Qw+i, 0777)
	}
}

func (o Oss) del(i string) {
	if o.Os == "windows" {
		strings.Replace(i, `/`, `\`, -1)
	}
	var err = os.Remove(o.Qw + i)

	if isError(err) {
		return
	}
}

/**
@return User
*/
func CreateUser() *Oss {
	switch runtime.GOOS {
	case "windows":
		return &Oss{Os: "windows", Path: os.Getenv("USERPROFIL"), Qw: os.Getenv("USERPROFILE") + `\Documents\qwertycore`}
	case "linux":
		return &Oss{Os: "linux", Path: os.Getenv("HOME"), Qw: os.Getenv("HOME") + "/qwertycore"}
	case "darwin":
		fmt.Println("! MAC IS NOT SUPPORT !")
		break
	}

	return &Oss{}
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
