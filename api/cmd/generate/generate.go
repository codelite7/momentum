package generate

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/codelite7/momentum/api/common"
	"github.com/codelite7/momentum/api/config"
	"github.com/codelite7/momentum/api/ent"
	"github.com/codelite7/momentum/api/ent/user"
	"github.com/urfave/cli/v2"
	"strings"
)

var GenerateCommand = &cli.Command{
	Name:  "generate",
	Usage: "generates fake data",
	Flags: flags,
	Action: func(c *cli.Context) error {
		return generate()
	},
}

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "postgres-uri",
		Aliases:     []string{"pu"},
		Value:       "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
		Usage:       "postgres connection string",
		EnvVars:     []string{"POSTGRES_URI"},
		Destination: &config.PostgresUri,
	},
	&cli.StringFlag{
		Name:        "user-email",
		Aliases:     []string{"ue"},
		Usage:       "email address of user to generate data for",
		Destination: &userEmail,
	},
	&cli.IntFlag{
		Name:        "num-threads",
		Aliases:     []string{"nt"},
		Value:       100,
		Usage:       "Number of threads to generate",
		Destination: &numThreads,
	},
	&cli.IntFlag{
		Name:        "messages-per-thread",
		Aliases:     []string{"mpt"},
		Value:       50,
		Usage:       "Number of messages to generate in each thread",
		Destination: &numMessages,
	},
}

func generate() error {
	err := common.InitializeEntClient()
	if err != nil {
		return err
	}
	user, err := ensureUser()
	if err != nil {
		return err
	}
	agent, err := createAgent()
	if err != nil {
		return err
	}
	threads, err := createThreads(user)
	if err != nil {
		return err
	}
	messages, err := createMessages(user, threads)
	if err != nil {
		return err
	}

	_, err = createResponses(agent, messages)
	if err != nil {
		return err
	}

	return nil
}

func ensureUser() (*ent.User, error) {
	user, err := common.EntClient.User.Query().Where(user.Email(userEmail)).First(context.Background())
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return nil, err
	}
	if user == nil {
		return common.EntClient.User.Create().SetEmail(userEmail).Save(context.Background())
	}
	return user, nil
}

func createAgent() (*ent.Agent, error) {
	return common.EntClient.Agent.Create().SetModel("test").SetProvider("test").SetAPIKey("test").Save(context.Background())
}

func createThreads(user *ent.User) ([]*ent.Thread, error) {
	creates := []*ent.ThreadCreate{}
	for i := 0; i < numThreads; i++ {
		thread := generateFakeThread()
		create := common.EntClient.Thread.Create().SetCreatedBy(user).SetName(thread.Name)
		creates = append(creates, create)
	}
	return common.EntClient.Thread.CreateBulk(creates...).Save(context.Background())
}

func generateFakeThread() *ent.Thread {
	return &ent.Thread{
		Name: gofakeit.Name(),
	}
}

func createMessages(user *ent.User, threads []*ent.Thread) ([]*ent.Message, error) {
	creates := []*ent.MessageCreate{}
	for _, thread := range threads {
		for i := 0; i < numMessages; i++ {
			message := generateFakeMessage()
			create := common.EntClient.Message.Create().SetSentBy(user).SetThread(thread).SetContent(message.Content)
			creates = append(creates, create)
		}

	}
	return common.EntClient.Message.CreateBulk(creates...).Save(context.Background())
}

func generateFakeMessage() *ent.Message {
	return &ent.Message{
		Content: generateFakeContent(),
	}
}

func generateFakeContent() string {
	numParagraphs := gofakeit.Number(1, 3)
	numSentences := gofakeit.Number(3, 10)
	numWords := gofakeit.Number(10, 25)

	return gofakeit.Paragraph(numParagraphs, numSentences, numWords, "\n")
}

func createResponses(agent *ent.Agent, messages []*ent.Message) ([]*ent.Response, error) {
	creates := []*ent.ResponseCreate{}
	for _, message := range messages {
		response := generateFakeResponse()
		create := common.EntClient.Response.Create().SetMessage(message).SetSentBy(agent).SetContent(response.Content)
		creates = append(creates, create)
	}
	return common.EntClient.Response.CreateBulk(creates...).Save(context.Background())
}

func generateFakeResponse() *ent.Response {
	return &ent.Response{Content: generateFakeContent()}
}
