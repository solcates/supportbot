package backends

import (
	"github.com/flashmob/go-guerrilla/backends"
)

type FireStore struct {

}

func (f *FireStore) Process(*interface{}) backends.Result {
	panic("implement me")
}

func (f *FireStore) ValidateRcpt(e *interface{}) backends.RcptError {
	panic("implement me")
}

func (f *FireStore) Initialize(backends.BackendConfig) error {
	panic("implement me")
}

func (f *FireStore) Reinitialize() error {
	panic("implement me")
}

func (f *FireStore) Shutdown() error {
	panic("implement me")
}

func (f *FireStore) Start() error {
	panic("implement me")
}

