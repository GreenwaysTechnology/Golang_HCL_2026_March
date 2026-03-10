package main

import "fmt"

func main() {
	var name = "Subramanian Murgan"
	fmt.Println("name:", name)
	var htmlDoc = `<html>
	<head><meta charset="utf-8" />
<title>Murgan</title>
</head>
<body>
<h1>Murgan</h1>
</body>
</html>`
	fmt.Println(htmlDoc)
}
