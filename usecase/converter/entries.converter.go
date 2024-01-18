package converter

import (
	"go-graphql-cli/domain/models/entities"
	"go-graphql-cli/domain/models/graphql"
	"time"
)

type (
    EntriesConverter interface {
        ConvertEntityToGraphQLType(entity *entities.Entry) *graphql.Entry
    }
    entriesConverterImpl struct{}
)

func NewEntriesConverter() EntriesConverter {
    return &entriesConverterImpl{}
}

func (c *entriesConverterImpl) ConvertEntityToGraphQLType(entity *entities.Entry) *graphql.Entry {
    createdAtStr := entity.CreatedAt.Format(time.RFC3339)
    entry := &graphql.Entry{
        ID:        entity.ID,
        Name:      entity.Name,
        CreatedAt: createdAtStr,
    }
    return entry
}
