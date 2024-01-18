package resolvers

import (
	"gorm.io/gorm"

	"go-graphql-cli/domain/repositories"
	"go-graphql-cli/usecase"
	"go-graphql-cli/usecase/converter"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
    EntriesUseCase usecase.EntriesUseCase
}

func NewResolver(db *gorm.DB) *Resolver {
    return &Resolver{
        EntriesUseCase: usecase.NewEntriesUseCase(
            converter.NewEntriesConverter(),
            repositories.NewEntryRepository(db),
        ),
    }
}
