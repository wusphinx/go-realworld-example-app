package testcase

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ChanTestSuite struct {
	suite.Suite
}

func (suite *ChanTestSuite) TestReadClosedChan1() {
	ic := make(chan int)
	close(ic)
	select {
	case a := <-ic:
		suite.Equal(0, a)
	}
}

func (suite *ChanTestSuite) TestReadClosedChan2() {
	ic := make(chan int)
	close(ic)
	select {
	// 用这种方式比较安全
	case a, ok := <-ic:
		suite.Equal(0, a)
		suite.Equal(false, ok)
	}
}

func (suite *ChanTestSuite) TestReadClosedChan3() {
	ic := make(chan int)
	close(ic)
	select {
	case a := <-ic:
		suite.Equal(0, a)
	}
}

func (suite *ChanTestSuite) TestWriteClosedChan() {

}

func TestCChanTestSuite(t *testing.T) {
	suite.Run(t, new(ChanTestSuite))
}
