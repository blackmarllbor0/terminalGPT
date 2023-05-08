package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	configI "terminalGPT/config/interfaces"
)

type Client struct {
	client      *mongo.Client
	collections map[string]*mongo.Collection
	ctx         context.Context

	configReader configI.ConfigReader
}

func NewClient(configReader configI.ConfigReader, ctx context.Context) *Client {
	return &Client{
		configReader: configReader,
		ctx:          ctx,
	}
}

func (c *Client) Connection() error {
	clientOpts := options.Client().ApplyURI(c.configReader.GetString("db.uri"))

	var err error
	c.client, err = mongo.NewClient(clientOpts)
	if err != nil {
		return err
	}

	if err := c.client.Connect(c.ctx); err != nil {
		return err
	}

	if err := c.client.Ping(c.ctx, nil); err != nil {
		return err
	}

	c.collections = make(map[string]*mongo.Collection)

	c.collections[CHATS] = c.client.
		Database(c.configReader.GetString("db.db-name")).
		Collection(c.configReader.GetString("db.collection"))

	return nil
}

func (c *Client) Disconnect() {
	_ = c.client.Disconnect(nil)
}

func (c *Client) GetCollectionByName(name string) *mongo.Collection {
	return c.collections[name]
}
