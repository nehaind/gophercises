package main

import (
	"testing"

	"github.ibm.com/dash/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	main()
	dashtest.ControlCoverage(m)
}
