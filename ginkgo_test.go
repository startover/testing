package stringutil

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StringutilTest", func() {
	var in, want string

	BeforeEach(func() {
		const in, want = "Hello, world", "dlrow ,olleH"
	})

	Describe("With ASCII and UTF8 strings defined", func() {
		Context("Reverse the give strings", func() {
			It("should be reversed", func() {
				got := Reverse(in)
				Expect(got).To(Equal(want))
			})
		})
	})
})

var _ = Describe("StringutilTableTest", func() {
	var (
		strs []struct {
			in, want string
		}
	)

	BeforeEach(func() {
		strs = []struct {
			in, want string
		}{
			{"Hello, world", "dlrow ,olleH"},
			{"Hello, 世界", "界世 ,olleH"},
			{"", ""},
		}
	})

	Describe("With ASCII and UTF8 strings defined", func() {
		Context("Reverse the give strings", func() {
			It("should be reversed", func() {
				for _, c := range strs {
					got := Reverse(c.in)
					Expect(got).To(Equal(c.want))
				}
			})
		})
	})
})
