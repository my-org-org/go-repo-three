module github.com/my-org-for-test/go-repo-three

go 1.22

require github.com/my-org-for-test/go-repo-two v1.0.4
require github.com/my-org-for-test/go-repo-one v1.0.4

replace github.com/my-org-for-test/go-repo-two v1.0.4 => github.com/my-org-org/go-repo-two v1.0.4
replace github.com/my-org-for-test/go-repo-one v1.0.4 => github.com/my-org-org/go-repo-one v1.0.4