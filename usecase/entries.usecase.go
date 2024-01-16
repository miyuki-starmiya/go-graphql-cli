package usecase

import (
	"fmt"
)

type (
    entriesUseCase interface {
        Fetch(arg string)
    }
    entriesUseCaseImpl struct {}
)

func NewEntriesUseCase() entriesUseCase {
    return &entriesUseCaseImpl{}
}

func (u *entriesUseCaseImpl) Fetch(arg string) {
    fmt.Printf("arg: %s\n", arg)
}
