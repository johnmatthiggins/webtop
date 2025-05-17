package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
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
<body>
<pre>%s</pre>
</body>
</html>
`

func home(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("top", "-bn1")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	html := fmt.Sprintf(formatString, string(output))

	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Println("STARTING SERVER...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
