package self_dependency_consumer

import (
	"fmt"
	"github.com/my-org-for-test/go-repo-three/internal"
)

func PrintDependencyMessage() {
	fmt.Println("go-repo-three. Release 2: Dependency Consumer")
    internal.InternalFunction()
}
