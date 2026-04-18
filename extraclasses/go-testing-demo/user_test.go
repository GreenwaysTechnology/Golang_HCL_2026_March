package main

import (
    "testing"
    "go-testing-demo/mocks"
    "github.com/stretchr/testify/suite"
    "github.com/stretchr/testify/assert"
    "go.uber.org/mock/gomock"
)

// Define a suite struct
type UserTestSuite struct {
    suite.Suite
    ctrl     *gomock.Controller
    mockRepo *mocks.MockRepository
    service  *UserService
}

// SetupTest runs BEFORE EACH test
func (s *UserTestSuite) SetupTest() {
    s.ctrl = gomock.NewController(s.T())
    s.mockRepo = mocks.NewMockRepository(s.ctrl)
    s.service = &UserService{repo: s.mockRepo}
}

// TearDownTest runs AFTER EACH test
func (s *UserTestSuite) TearDownTest() {
    s.ctrl.Finish()
}

// The actual test case
func (s *UserTestSuite) TestWelcomeUser() {
    // 1. Mock Behavior: Expect ID 1, return "Alice"
    s.mockRepo.EXPECT().GetUserName(1).Return("Alice")

    // 2. Execute
    result := s.service.WelcomeUser(1)

    // 3. Assert using testify/assert
    assert.Equal(s.T(), "Hello, Alice", result, "The greeting should match the username")
}

// Entry point for the suite
func TestUserTestSuite(t *testing.T) {
    suite.Run(t, new(UserTestSuite))
}