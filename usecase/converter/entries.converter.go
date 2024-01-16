package converter

import (
	"go-graphql-cli/domain/models/entities"
	"go-graphql-cli/domain/models/graphql"
	"time"
)

type (
    EntriesConverter interface {
        ConvertEntitiesToGraphQLType(entities []entities.Entry) []*graphql.Entry
    }
    entriesConverterImpl struct{}
)

func NewEntriesConverter() EntriesConverter {
    return &entriesConverterImpl{}
}

func (c *entriesConverterImpl) ConvertEntitiesToGraphQLType(entities []entities.Entry) []*graphql.Entry {
    var entries []*graphql.Entry
    for _, entity := range entities {
        createdAtStr := entity.CreatedAt.Format(time.RFC3339)
        entry := &graphql.Entry{
            ID:        entity.ID,
            Name:      entity.Name,
            CreatedAt: createdAtStr,
        }
        entries = append(entries, entry)
    }
    return entries
}
