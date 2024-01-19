package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"go-graphql-cli/domain/models/entities"
	"go-graphql-cli/domain/models/graphql"
	"go-graphql-cli/domain/repositories"
	"go-graphql-cli/usecase/converter"
)

type (
	EntriesUseCase interface {
		Fetch(arg string) (*entities.Entry, error)
		GetEntries() ([]*graphql.Entry, error)
		GetEntry(id string) (*graphql.Entry, error)
		CreateEntry(entry *entities.Entry) error
	}
	entriesUseCaseImpl struct{
		converter converter.EntriesConverter
		repository repositories.EntryRepository
	}
	EntryRes              struct {
		Sys    Sys    `json:"sys"`
		Fields Fields `json:"fields"`
	}
	Sys struct {
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
	}
	Fields struct {
		Name string `json:"name"`
	}
)

func NewEntriesUseCase(
	c converter.EntriesConverter,
	r repositories.EntryRepository,
) EntriesUseCase {
	return &entriesUseCaseImpl{
		converter: c,
		repository: r,
	}
}

func (u *entriesUseCaseImpl) Fetch(arg string) (*entities.Entry, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %w", err)
	}

	accessToken := os.Getenv("ACCESS_TOKEN")
	fmt.Printf("arg: %s\n", arg)
	url := fmt.Sprintf("https://cdn.contentful.com/spaces/2vskphwbz4oc/entries/%s?access_token=%s", arg, accessToken)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error making GET request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %w", err)
	}

	var entryRes EntryRes

	err = json.Unmarshal(body, &entryRes)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling JSON data: %w", err)
	}

	// Print the Entry struct
	entry := &entities.Entry{
		ID:        entryRes.Sys.ID,
		Name:      entryRes.Fields.Name,
		CreatedAt: entryRes.Sys.CreatedAt,
	}

	return entry, nil
}

func (u *entriesUseCaseImpl) CreateEntry(entry *entities.Entry) error {
	if err := u.repository.CreateEntry(entry); err != nil {
		return err
	}
	return nil
}

func (u *entriesUseCaseImpl) GetEntries() ([]*graphql.Entry, error) {
	es, err := u.repository.GetEntries()
	if err != nil {
		return nil, err
	}
	var entries []*graphql.Entry
	for _, e := range es {
		entries = append(entries, u.converter.ConvertEntityToGraphQLType(&e))
	}
	return entries, nil
}

func (u *entriesUseCaseImpl) GetEntry(id string) (*graphql.Entry, error) {
	e, err := u.repository.GetEntry(id)
	if err != nil {
		return nil, err
	}
	return u.converter.ConvertEntityToGraphQLType(e), nil
}
