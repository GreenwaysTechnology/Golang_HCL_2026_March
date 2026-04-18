package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// 1. Define the Suite
type AdvancedTestSuite struct {
	suite.Suite
	// Imagine this is a DB connection or Cache client
	MockResource string 
}

// 2. Setup (Runs before all tests in the suite)
func (s *AdvancedTestSuite) SetupSuite() {
	s.MockResource = "initialized"
}

// 3. Teardown (Runs after all tests)
func (s *AdvancedTestSuite) TearDownSuite() {
	s.MockResource = ""
}

// 4. Test HTTP Handler
func (s *AdvancedTestSuite) TestFibonacciHandler() {
	// Create a request to /fib?n=10
	req, _ := http.NewRequest("GET", "/fib?n=10", nil)
	
	// ResponseRecorder acts as the client-side buffer
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FibonacciHandler)

	handler.ServeHTTP(rr, req)

	// Assertions using testify
	assert.Equal(s.T(), http.StatusOK, rr.Code)
	assert.Equal(s.T(), "55", rr.Body.String())
}

// 5. Entry point to run the suite
func TestRunSuite(t *testing.T) {
	suite.Run(t, new(AdvancedTestSuite))
}