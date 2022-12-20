package main

import (
	"github.com/nilic/kubectl-unlimited/cmd"

	// Import to initialize client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	cmd.Execute()
}
