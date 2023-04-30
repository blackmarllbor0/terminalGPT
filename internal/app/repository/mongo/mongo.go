package mongo

import (
	"context"
	configI "terminalGPT/config/interfaces"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	client *mongo.Client

	configReader configI.IConfigReader
}

func NewClient(configReader configI.IConfigReader) *Client {
	return &Client{
		configReader: configReader,
	}
}

func (c *Client) Connection() error {
	clientOpts := options.Client().ApplyURI(c.configReader.GetString("db.uri"))

	var err error
	c.client, err = mongo.NewClient(clientOpts)
	if err != nil {
		return err
	}

	if err := c.client.Connect(context.Background()); err != nil {
		return err
	}

	if err := c.client.Ping(context.Background(), nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) Disconnect() {
	_ = c.client.Disconnect(nil)
}
