package coinpaprika

import "github.com/stretchr/testify/suite"

// ChangeLogTestSuite is a test suite for testing the ChangeLog-related functionality.
// It uses the suite package to provide structured and reusable test cases.
//
// Fields:
//   - suite.Suite: Embedded struct from the suite package to enable test suite functionality.
//   - paprikaClient: A client instance used to interact with the API during tests.
type ChangeLogTestSuite struct {
	suite.Suite
	paprikaClient *Client
}

// SetupTest initializes the test environment before each test case is executed.
// It creates a new API client and ensures it is not nil.
func (suite *ChangeLogTestSuite) SetupTest() {
	paprikaClient := NewClient(nil)
	suite.NotNil(paprikaClient)

	suite.paprikaClient = paprikaClient
}

// TestChangeLog tests the CoinsChangeLog method of the ChangeLog service.
// It verifies that the method returns a non-empty change log and no error when provided with valid options.
func (suite *ChangeLogTestSuite) TestChangeLog() {
	options := &ChangeLogOptions{Page: 1}
	changeLog, err := suite.paprikaClient.ChangeLog.CoinsChangeLog(options)
	suite.NoError(err)
	suite.NotEmpty(changeLog)
}
