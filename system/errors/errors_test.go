package errors_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/lucsky/cuid"
	systemErrors "github.com/micro-business/go-core/system/errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SystemErrors Tests", func() {
	Describe("ArgumentNilError", func() {
		Context("argumentName is provided", func() {
			When("error is instantiated", func() {
				var (
					argumentName string
					err          error
				)

				BeforeEach(func() {
					argumentName = cuid.New()
					err = systemErrors.NewArgumentNilError(argumentName, cuid.New())
				})

				It("should set the internal ArgumentName property", func() {
					var argumentNilErr systemErrors.ArgumentNilError

					errors.As(err, &argumentNilErr)
					Ω(errors.As(err, &argumentNilErr)).Should(BeTrue())
					Ω(argumentNilErr.ArgumentName).Should(Equal(argumentName))
				})

				It("should include argument name in the output error string", func() {
					Ω(strings.Contains(err.Error(), argumentName)).Should(BeTrue())
				})
			})
		})

		Context("message is provided", func() {
			When("error is instantiated", func() {
				var (
					message string
					err     error
				)

				BeforeEach(func() {
					message = cuid.New()
					err = systemErrors.NewArgumentNilError(cuid.New(), message)
				})

				It("should set the internal Message property", func() {
					var argumentNilErr systemErrors.ArgumentNilError

					errors.As(err, &argumentNilErr)
					Ω(errors.As(err, &argumentNilErr)).Should(BeTrue())
					Ω(argumentNilErr.Message).Should(Equal(message))
				})

				It("should include message in the output error string", func() {
					Ω(strings.Contains(err.Error(), message)).Should(BeTrue())
				})
			})
		})

		Context("err is provided", func() {
			Context("err is not provided", func() {
				When("error is instantiated", func() {
					var (
						err error
					)

					BeforeEach(func() {
						err = systemErrors.NewArgumentNilError(cuid.New(), cuid.New())
					})

					It("should unwrap the internal error adn return nil", func() {
						Ω(errors.Unwrap(err)).Should(BeNil())
					})

					It("should not include error message in the output error string", func() {
						Ω(strings.Contains(err.Error(), "Error: ")).Should(BeFalse())
					})
				})
			})
			When("error is instantiated", func() {
				var (
					expectedErr error
					err         error
				)

				BeforeEach(func() {
					expectedErr = errors.New(cuid.New())
					err = systemErrors.NewArgumentNilErrorWithError(cuid.New(), cuid.New(), expectedErr)
				})

				It("should unwrap the internal error", func() {
					Ω(errors.Unwrap(err)).Should(Equal(expectedErr))
				})

				It("should include error message in the output error string", func() {
					Ω(strings.Contains(err.Error(), expectedErr.Error())).Should(BeTrue())
				})
			})
		})

		Context("err is not provided", func() {
			When("error is instantiated", func() {
				var (
					err error
				)

				BeforeEach(func() {
					err = systemErrors.NewArgumentNilError(cuid.New(), cuid.New())
				})

				It("should unwrap the internal error adn return nil", func() {
					Ω(errors.Unwrap(err)).Should(BeNil())
				})

				It("should not include error message in the output error string", func() {
					Ω(strings.Contains(err.Error(), "Error: ")).Should(BeFalse())
				})
			})
		})

		Context("err is instantiated", func() {
			When("IsArgumentNilError function is called with ArgumentNilError", func() {
				It("should return true", func() {
					err := systemErrors.NewArgumentNilError(cuid.New(), cuid.New())
					Ω(systemErrors.IsArgumentNilError(err)).Should(BeTrue())
				})
			})

			When("IsArgumentNilError function is called with different error", func() {
				It("should return false", func() {
					Ω(systemErrors.IsArgumentNilError(errors.New(cuid.New()))).Should(BeFalse())
				})
			})

		})
	})

	Describe("ArgumentError", func() {
		Context("argumentName is provided", func() {
			When("error is instantiated", func() {
				var (
					argumentName string
					err          error
				)

				BeforeEach(func() {
					argumentName = cuid.New()
					err = systemErrors.NewArgumentError(argumentName, cuid.New())
				})

				It("should set the internal ArgumentName property", func() {
					var argumentNilErr systemErrors.ArgumentError

					errors.As(err, &argumentNilErr)
					Ω(errors.As(err, &argumentNilErr)).Should(BeTrue())
					Ω(argumentNilErr.ArgumentName).Should(Equal(argumentName))
				})

				It("should include argument name in the output error string", func() {
					Ω(strings.Contains(err.Error(), argumentName)).Should(BeTrue())
				})
			})
		})

		Context("message is provided", func() {
			When("error is instantiated", func() {
				var (
					message string
					err     error
				)

				BeforeEach(func() {
					message = cuid.New()
					err = systemErrors.NewArgumentError(cuid.New(), message)
				})

				It("should set the internal Message property", func() {
					var argumentNilErr systemErrors.ArgumentError

					errors.As(err, &argumentNilErr)
					Ω(errors.As(err, &argumentNilErr)).Should(BeTrue())
					Ω(argumentNilErr.Message).Should(Equal(message))
				})

				It("should include message in the output error string", func() {
					Ω(strings.Contains(err.Error(), message)).Should(BeTrue())
				})
			})
		})

		Context("err is provided", func() {
			Context("err is not provided", func() {
				When("error is instantiated", func() {
					var (
						err error
					)

					BeforeEach(func() {
						err = systemErrors.NewArgumentError(cuid.New(), cuid.New())
					})

					It("should unwrap the internal error adn return nil", func() {
						Ω(errors.Unwrap(err)).Should(BeNil())
					})

					It("should not include error message in the output error string", func() {
						Ω(strings.Contains(err.Error(), "Error: ")).Should(BeFalse())
					})
				})
			})
			When("error is instantiated", func() {
				var (
					expectedErr error
					err         error
				)

				BeforeEach(func() {
					expectedErr = errors.New(cuid.New())
					err = systemErrors.NewArgumentErrorWithError(cuid.New(), cuid.New(), expectedErr)
				})

				It("should unwrap the internal error", func() {
					Ω(errors.Unwrap(err)).Should(Equal(expectedErr))
				})

				It("should include error message in the output error string", func() {
					Ω(strings.Contains(err.Error(), expectedErr.Error())).Should(BeTrue())
				})
			})
		})

		Context("err is not provided", func() {
			When("error is instantiated", func() {
				var (
					err error
				)

				BeforeEach(func() {
					err = systemErrors.NewArgumentError(cuid.New(), cuid.New())
				})

				It("should unwrap the internal error adn return nil", func() {
					Ω(errors.Unwrap(err)).Should(BeNil())
				})

				It("should not include error message in the output error string", func() {
					Ω(strings.Contains(err.Error(), "Error: ")).Should(BeFalse())
				})
			})
		})

		Context("err is instantiated", func() {
			When("IsArgumentError function is called with ArgumentError", func() {
				It("should return true", func() {
					err := systemErrors.NewArgumentError(cuid.New(), cuid.New())

					Ω(systemErrors.IsArgumentError(err)).Should(BeTrue())
				})
			})

			When("IsArgumentError function is called with different error", func() {
				It("should return false", func() {
					Ω(systemErrors.IsArgumentError(errors.New(cuid.New()))).Should(BeFalse())
				})
			})

		})
	})
})

func TestSystemErrors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SystemErrors Tests")
}
