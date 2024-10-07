package gohttp

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"sync"
)

var (
	mockupServer = mockServer{mocks: make(map[string]*Mock)}
)

type mockServer struct {
	enable      bool
	serverMutex sync.Mutex
	mocks       map[string]*Mock
}

func StartMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()
	mockupServer.enable = true
}

func StopMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()
	mockupServer.enable = false
}

func FlushMocks() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()
	mockupServer.mocks = make(map[string]*Mock)
}

func AddMock(mock Mock) {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()
	key := mockupServer.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	mockupServer.mocks[key] = &mock
}

func (m *mockServer) getMockKey(method, url, body string) string {
	hasher := md5.New()
	hasher.Write([]byte(method + url + m.cleanBody(body)))
	kye := hex.EncodeToString(hasher.Sum(nil))
	return kye
}

func (m *mockServer) cleanBody(body string) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}
	body = strings.ReplaceAll(body, "\t", "")
	body = strings.ReplaceAll(body, "\n", "")
	return body
}

func (m *mockServer) getMock(method, url, body string) *Mock {
	if !m.enable {
		return nil
	}
	key := mockupServer.getMockKey(method, url, body)
	if mock := mockupServer.mocks[key]; mock != nil {
		return mock
	}
	return &Mock{
		Error: errors.New(fmt.Sprintf("no mock matching %s from %s with given body", method, url)),
	}

}
