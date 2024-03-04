package main

import (
	"fmt"
	"github.com/my-org-for-test/go-repo-two/dependency_consumer"
	"github.com/my-org-for-test/go-repo-three/self_dependency_consumer"
)

func main() {
    fmt.Println("========> go-repo-three: Release 3")

	fmt.Println("go-repo-three: Dependency on go-repo-two:")
	dependency_consumer.PrintDependencyMessage()

    fmt.Println("go-repo-three: Self dependency")
	self_dependency_consumer.PrintDependencyMessage()
}
