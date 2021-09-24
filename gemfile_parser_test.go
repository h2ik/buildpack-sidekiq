package sidekiq_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/paketo-buildpacks/sidekiq"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testGemfileParser(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		path   string
		parser sidekiq.GemfileParser
	)

	it.Before(func() {
		file, err := ioutil.TempFile("", "Gemfile")
		Expect(err).NotTo(HaveOccurred())
		defer file.Close()

		path = file.Name()

		parser = sidekiq.NewGemfileParser()
	})

	it.After(func() {
		Expect(os.RemoveAll(path)).To(Succeed())
	})

	context("Parse", func() {
		context("when using sidekiq", func() {
			it("parses correctly without spaces", func() {
				const GEMFILE_CONTENTS = `
source 'https://rubygems.org'

gem 'sidekiq'
`

				Expect(ioutil.WriteFile(path, []byte(GEMFILE_CONTENTS), 0644)).To(Succeed())

				hasSidekiq, err := parser.Parse(path, "sidekiq")
				Expect(err).NotTo(HaveOccurred())
				Expect(hasSidekiq).To(Equal(true))
			})

			it("parses correctly with spaces", func() {
				const GEMFILE_CONTENTS = `
source 'https://rubygems.org' do
	gem 'sidekiq'
end
`

				Expect(ioutil.WriteFile(path, []byte(GEMFILE_CONTENTS), 0644)).To(Succeed())

				hasSidekiq, err := parser.Parse(path, "sidekiq")
				Expect(err).NotTo(HaveOccurred())
				Expect(hasSidekiq).To(Equal(true))
			})
		})

		context("when sidekiq is commented out", func() {
			it("parses correctly", func() {
				const GEMFILE_CONTENTS = `source 'https://rubygems.org'

				#gem 'sidekiq`

				Expect(ioutil.WriteFile(path, []byte(GEMFILE_CONTENTS), 0644)).To(Succeed())

				hasSidekiq, err := parser.Parse(path, "sidekiq")
				Expect(err).NotTo(HaveOccurred())
				Expect(hasSidekiq).To(Equal(false))
			})
		})

		context("when not using sidekiq", func() {
			it("parses correctly", func() {
				const GEMFILE_CONTENTS = `source 'https://rubygems.org'`

				Expect(ioutil.WriteFile(path, []byte(GEMFILE_CONTENTS), 0644)).To(Succeed())

				hasSidekiq, err := parser.Parse(path, "sidekiq")
				Expect(err).NotTo(HaveOccurred())
				Expect(hasSidekiq).To(Equal(false))
			})
		})

		context("when the Gemfile file does not exist", func() {
			it.Before(func() {
				Expect(os.Remove(path)).To(Succeed())
			})

			it("returns all false", func() {
				hasSidekiq, err := parser.Parse(path, "sidekiq")
				Expect(err).NotTo(HaveOccurred())
				Expect(hasSidekiq).To(Equal(false))
			})
		})

		context("failure cases", func() {
			context("when the Gemfile cannot be opened", func() {
				it.Before(func() {
					Expect(os.Chmod(path, 0000)).To(Succeed())
				})

				it("returns an error", func() {
					_, err := parser.Parse(path, "sidekiq")
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError(ContainSubstring("failed to parse Gemfile:")))
					Expect(err).To(MatchError(ContainSubstring("permission denied")))
				})
			})
		})
	})
}
