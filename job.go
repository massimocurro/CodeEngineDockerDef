package main


import (
	"fmt"
	"os"
        "net/http"
)

func main() {
        var target := os.Getenv("Key")
	fmt.Printf("Hi from a batch job! My index within the array of %s instance(s) is %s\n", os.Getenv("JOB_ARRAY_SIZE"), os.Getenv("JOB_INDEX"))
        http.Get("https://function-76.1at6rgz00yjr.eu-de.codeengine.appdomain.cloud")
}
