package service_test

import (
	. "github.com/maurofran/go-ddd-identityaccess/infrastructure/service"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"
)

var _ = Describe("BcryptEncryptionService", func() {
	fixture := new(BCryptEncryptionService)

	It("Should encrypt the value using bcrypt algorithm", func() {
		res, err := fixture.EncryptValue("TestValue")

		Expect(err).ToNot(HaveOccurred())
		Expect(res).ToNot(BeEmpty())
		Expect(bcrypt.CompareHashAndPassword([]byte(res), []byte("TestValue"))).ToNot(HaveOccurred())
	})
})
