package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

const formatString string = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <title>Top</title>
  <meta name="viewport" content="width=device-width,initial-scale=1" />
  <meta name="description" content="" />
</head>
<body style="background-color:black;color:white;">
<pre>%s</pre>
</body>
</html>
`

func home(w http.ResponseWriter, r *http.Request) {
	output := ""

	if _, err := exec.LookPath("fastfetch"); err == nil {
		cmd2 := exec.Command("sh", "-c", "fastfetch | aha --no-header")
		cmdOutput, err := cmd2.Output()
		if err != nil {
			log.Fatal(err)
		}
		output = strings.TrimPrefix(string(cmdOutput), "\n") + output
	} else {
		// default to neofetch
		cmd2 := exec.Command("sh", "-c", "neofetch | aha --no-header")
		cmdOutput, err := cmd2.Output()
		if err != nil {
			log.Fatal(err)
		}
		output = strings.TrimPrefix(string(cmdOutput), "\n") + output
	}

	cmd := exec.Command("top", "-bn1")
	cmdOutput, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	output += "\n" + string(cmdOutput)
	pattern := regexp.MustCompile("(\\w+)@(\\w+)")
	output = pattern.ReplaceAllString(output, "\n\n$1@$2")

	html := fmt.Sprintf(formatString, output)
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Println("...STARTING SERVER...")
	log.Fatal(http.ListenAndServe(":8880", nil))
}
