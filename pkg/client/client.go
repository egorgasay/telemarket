package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
)

// Client is a struct for working with Telegram API.
type Client struct {
	updates    chan Update
	done       chan struct{}
	token      string
	httpClient *http.Client
}

// telegramAPIHost is a constant for Telegram API address.
const (
	telegramAPIHost = "api.telegram.org"
)

// New is a constructor for Client.
func New(token string) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		token:      "bot" + token,
		done:       make(chan struct{}),
	}
}

// GetUpdatesChan is a method for getting updates from TelegramAPI.
func (c *Client) GetUpdatesChan(uc UpdateConfig) chan Update {
	c.updates = make(chan Update, 100)
	go c.receive(uc)

	return c.updates
}

// StopReceivingUpdates is a method for stopping receiving updates from TelegramAPI.
func (c *Client) StopReceivingUpdates() {
	close(c.done)
	close(c.updates)
}

// NewUpdate is a method for creating UpdateConfig.
func NewUpdate(offset int) UpdateConfig {
	return UpdateConfig{
		Offset:  offset,
		Limit:   0,
		Timeout: 0,
	}
}

// SendMessage is a method for sending message to Telegram.
func (c *Client) receive(updateConf UpdateConfig) {
	for {
		select {
		case <-c.done:
			return
		default:
		}

		updates, err := c.getBatchUpdates(updateConf)
		if err != nil {
			log.Println(err.Error())
			time.Sleep(time.Second) // in case something wrong with the network or Telegram API.
			continue
		}

		for _, update := range updates {
			if updateConf.Offset <= update.UpdateID {
				updateConf.Offset = update.UpdateID + 1
				c.updates <- update
			}
		}
	}
}

// getBatchUpdates is a method for getting batch updates.
func (c *Client) getBatchUpdates(updateConf UpdateConfig) ([]Update, error) {
	r, err := c.doRequest("getUpdates", strings.NewReader(fmt.Sprintf("allowed_updates=null&offset=%d&timeout=%d", updateConf.Offset, updateConf.Timeout)))
	if err != nil {
		return nil, fmt.Errorf("can't get batch updates %e", err)
	}

	if !r.Ok {
		return nil, fmt.Errorf("can't get batch updates: ErrorCode: %d", r.ErrorCode)
	}

	var updates = make([]Update, 0)
	if err = json.Unmarshal(r.Result, &updates); err != nil {
		return nil, fmt.Errorf("can't unmarshal updates %e", err)
	}

	return updates, nil
}

// doRequest is a method for sending request to Telegram API.
func (c *Client) doRequest(endpoint string, body io.Reader) (*Response, error) {
	requestPath := path.Join(telegramAPIHost, c.token, endpoint)

	r, err := http.NewRequest("POST", "https://"+requestPath, body)
	if err != nil {
		return nil, fmt.Errorf("can't create request %e", err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("can't do request %e", err)
	}
	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("can't decode response %e", err)
	}

	return &response, nil
}

// ISend is an interface for sending messages.
type ISend interface {
	getParams() map[string]string
	getEndpoint() string
}

// Send is a method for sending messages.
func (c *Client) Send(s ISend) (Message, error) {
	params := s.getParams()
	query := makeQuery(params)
	r, err := c.doRequest(s.getEndpoint(), strings.NewReader(query))
	if err != nil {
		return Message{}, fmt.Errorf("can't send message %e", err)
	}

	var message Message
	if err = json.Unmarshal(r.Result, &message); err != nil {
		return Message{}, fmt.Errorf("can't unmarshal message %e", err)
	}

	return message, nil
}

// NewEditMessageTextAndMarkup is a method for editing message text and markup.
func NewEditMessageTextAndMarkup(chatID int, messageID int, text string, markup InlineKeyboardMarkup) ISend {
	return EditMessageTextConfig{
		BaseEdit: BaseEdit{
			ChatID:      int64(chatID),
			MessageID:   messageID,
			ReplyMarkup: &markup,
		},
		Text: text,
	}
}

// NewMessage is a method for creating new message.
func NewMessage(chatID int64, text string) MessageConfig {
	return MessageConfig{
		BaseChat: BaseChat{
			ChatID:           chatID,
			ReplyToMessageID: 0,
		},
		Text:                  text,
		DisableWebPagePreview: false,
	}
}

// jsonAnything is a method for converting interface to JSON string.
func jsonAnything(i interface{}) string {
	b, err := json.Marshal(i)
	if err != nil {
		log.Println(err)
	}
	return string(b)
}

// makeQuery is a method for creating query string.
func makeQuery(params map[string]string) string {
	var query = make([]string, 0)
	for k, v := range params {
		query = append(query, fmt.Sprintf("%s=%s&", k, v))
	}

	return strings.Trim(strings.Join(query, ""), "&")
}
