package main

import (
	"fmt"
	"github.com/my-org-org/go-repo-two/dependency_consumer"
	"github.com/my-org-org/go-repo-three/self_dependency_consumer"
)

func main() {
    fmt.Println("========> go-repo-three: Release 7777")

	fmt.Println("go-repo-three: Dependency on go-repo-two:")
	dependency_consumer.PrintDependencyMessage()

    fmt.Println("go-repo-three: Self dependency")
	self_dependency_consumer.PrintDependencyMessage()
}
