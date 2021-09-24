package fakes

import "sync"

type Parser struct {
	ParseCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Path string
		}
		Returns struct {
			HasSidekiq bool
			Err        error
		}
		Stub func(string) (bool, error)
	}
}

func (f *Parser) Parse(param1 string) (bool, error) {
	f.ParseCall.Lock()
	defer f.ParseCall.Unlock()
	f.ParseCall.CallCount++
	f.ParseCall.Receives.Path = param1
	if f.ParseCall.Stub != nil {
		return f.ParseCall.Stub(param1)
	}
	return f.ParseCall.Returns.HasSidekiq, f.ParseCall.Returns.Err
}