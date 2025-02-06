package coinpaprika

import "github.com/stretchr/testify/suite"

type ChangeLogTestSuite struct {
	suite.Suite
	paprikaClient *Client
}

func (suite *ChangeLogTestSuite) SetupTest() {
	paprikaClient := NewClient(nil)
	suite.NotNil(paprikaClient)

	suite.paprikaClient = paprikaClient
}

func (suite *ChangeLogTestSuite) TestChangeLog() {
	options := &ChangeLogOptions{Page: 1}
	changeLog, err := suite.paprikaClient.ChangeLog.CoinsChangeLog(options)
	suite.NoError(err)
	suite.NotEmpty(changeLog)
}
