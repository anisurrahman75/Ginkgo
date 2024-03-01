package e2e_test

import (
	"github.com/onsi/ginkgo/v2/types"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	_ "github.com/anisurrahman75/Ginkgo/tests/example"
	"gomodules.xyz/logs"
)

const (
	TIMEOUT = 20 * time.Minute
)

// ginkgo -r --v -race  --trace
func TestE2e(t *testing.T) {
	logs.InitLogs()
	RegisterFailHandler(Fail)
	SetDefaultEventuallyTimeout(TIMEOUT)

	reporterConfig := types.NewDefaultReporterConfig()
	reporterConfig.JUnitReport = "junit.xml"
	reporterConfig.JSONReport = "report.json"
	reporterConfig.Verbose = true
	RunSpecs(t, "e2e Suite", Label("test"), reporterConfig)
}
