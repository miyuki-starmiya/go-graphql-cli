package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"go-graphql-cli/domain/models/graphql"
	"go-graphql-cli/domain/repositories"
	"go-graphql-cli/usecase/converter"
)

type (
	EntriesUseCase interface {
		Fetch(arg string)
		GetEntries() ([]*graphql.Entry, error)
		GetEntry(id string) (*graphql.Entry, error)
	}
	entriesUseCaseImpl struct{
		converter converter.EntriesConverter
		repository repositories.EntryRepository
	}
	Entry              struct {
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

func (u *entriesUseCaseImpl) Fetch(arg string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessToken := os.Getenv("ACCESS_TOKEN")
	fmt.Printf("arg: %s\n", arg)

	// Define the URL with the access token
	url := fmt.Sprintf("https://cdn.contentful.com/spaces/2vskphwbz4oc/entries/%s?access_token=%s", arg, accessToken)

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Define an Entry variable to hold the unmarshaled data
	var entry Entry

	// Unmarshal the JSON data into the Entry struct
	err = json.Unmarshal(body, &entry)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Print the Entry struct
	fmt.Printf("ID: %s\n", entry.Sys.ID)
	fmt.Printf("Name: %s\n", entry.Fields.Name)
	fmt.Printf("CreatedAt: %s\n", entry.Sys.CreatedAt)

	// TODO: call repository
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
