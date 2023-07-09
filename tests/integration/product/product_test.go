package product

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestProductTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
