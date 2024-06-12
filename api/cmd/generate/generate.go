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
		Name:        "min-messages-per-thread",
		Aliases:     []string{"minmpt"},
		Value:       1,
		Usage:       "Minimum number of messages to generate in each thread",
		Destination: &minNumMessages,
	},
	&cli.IntFlag{
		Name:        "max-messages-per-thread",
		Aliases:     []string{"maxmpt"},
		Value:       50,
		Usage:       "Maximum number of messages to generate in each thread",
		Destination: &maxNumMessages,
	},
	&cli.IntFlag{
		Name:        "min-bookmarks-per-thread",
		Aliases:     []string{"minbpt"},
		Value:       0,
		Usage:       "Minimum number of bookmarks per thread",
		Destination: &minBookmarksPerThread,
	}, &cli.IntFlag{
		Name:        "max-bookmarks-per-thread",
		Aliases:     []string{"maxbpt"},
		Value:       10,
		Usage:       "Maximum number of bookmarks per thread",
		Destination: &maxBookmarksPerThread,
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
	_, err = createAgent()
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
	_, err = createBookmarks(user, messages)
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
		numMessages := gofakeit.Number(minNumMessages, maxNumMessages)
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

func createBookmarks(user *ent.User, messages []*ent.Message) ([]*ent.Bookmark, error) {
	messagesByThread := map[string]*ent.Message{}
	for _, message := range messages {
		thread, err := message.Thread(context.Background())
		if err != nil {
			return nil, err
		}
		messagesByThread[string(thread.ID)] = message
	}
	creates := []*ent.BookmarkCreate{}
	for _, threadMessages := range messagesByThread {
		numBookmarks := gofakeit.Number(minBookmarksPerThread, maxBookmarksPerThread)
		for i := 0; i < numBookmarks; i++ {
			create := common.EntClient.Bookmark.Create().SetMessage(threadMessages).SetUser(user)
			creates = append(creates, create)
		}
	}
	return common.EntClient.Bookmark.CreateBulk(creates...).Save(context.Background())
}
