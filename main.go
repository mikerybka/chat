package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mikerybka/openai"
	"github.com/mikerybka/util"
)

func main() {
	r, err := openai.NewClient(util.RequireEnvVar("OPENAI_API_KEY")).GetResponse("gpt-4o", strings.Join(os.Args[1:], " "))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(r.Output) < 1 {
		b, _ := json.MarshalIndent(r, "", "  ")
		fmt.Println(string(b))
		return
	}
	fmt.Println(r.Output[0].Content[0].Text)
}
