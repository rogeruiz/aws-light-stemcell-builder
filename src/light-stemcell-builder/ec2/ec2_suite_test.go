package ec2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEc2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ec2 Suite")
}
