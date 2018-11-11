package fonerotoml

import "github.com/stretchr/testify/mock"

// MockClient is a mockable fonerotoml client.
type MockClient struct {
	mock.Mock
}

// GetFoneroToml is a mocking a method
func (m *MockClient) GetFoneroToml(domain string) (*Response, error) {
	a := m.Called(domain)
	return a.Get(0).(*Response), a.Error(1)
}

// GetFoneroTomlByAddress is a mocking a method
func (m *MockClient) GetFoneroTomlByAddress(address string) (*Response, error) {
	a := m.Called(address)
	return a.Get(0).(*Response), a.Error(1)
}
