package coinpaprika

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// ContractsTestSuite is a test suite for testing the CoinPaprika API client.
// It embeds suite.Suite from the testify package and includes a paprikaClient
// field which is an instance of the Client used to interact with the CoinPaprika API.
type ContractsTestSuite struct {
	suite.Suite
	paprikaClient *Client
}

// SetupTest initializes the test suite by creating a new instance of the
// CoinPaprika API client and assigning it to the suite's paprikaClient field.
// It also ensures that the client instance is not nil.
func (suite *ContractsTestSuite) SetupTest() {
	paprikaClient := NewClient(nil)
	suite.NotNil(paprikaClient)

	suite.paprikaClient = paprikaClient
}

// TestList verifies that the Contracts.List() method of the paprikaClient
// returns a non-empty list of contracts without any errors.
func (suite *ContractsTestSuite) TestList() {
	contracts, err := suite.paprikaClient.Contracts.List()
	suite.NoError(err)
	suite.NotEmpty(contracts)
}

// TestGetTickerByContractAddress tests the retrieval of ticker information
// for a given contract address using the CoinPaprika API client.
// It verifies that no error occurs and that the returned tickers are not empty.
func (suite *ContractsTestSuite) TestGetTickerByContractAddress() {
	tickers, err := suite.paprikaClient.Contracts.GetTickerByContractAddress("eth-ethereum", "0xd26114cd6ee289accf82350c8d8487fedb8a0c07")
	suite.NoError(err)
	suite.NotEmpty(tickers)
}

// TestGetContractaddressessForPlatform tests the retrieval of contract addresses for a given platform
// using the GetContractaddressessForPlatform method of the paprikaClient. It ensures that no error
// is returned and that the list of platforms is not empty.
func (suite *ContractsTestSuite) TestGetContractaddressessForPlatform() {
	platforms, err := suite.paprikaClient.Contracts.GetContractaddressessForPlatform("eth-ethereum")
	suite.NoError(err)
	suite.NotEmpty(platforms)
}

// TestContractsTestSuite runs the test suite for the ContractsTestSuite.
// It uses the suite.Run function from the testify package to execute the tests.
//
// Parameters:
// - t (*testing.T): The testing framework's instance used to run the test suite.
func TestContractsTestSuite(t *testing.T) {
	suite.Run(t, new(ContractsTestSuite))
}
