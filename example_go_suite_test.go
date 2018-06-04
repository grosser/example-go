package example

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestAwesome(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Shopping Cart Suite")
}