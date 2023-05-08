package client

import (
	"testing"
)

func Test_makeQuery(t *testing.T) {
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{
				params: map[string]string{
					"name": "test",
					"age":  "20",
				},
			},
			want: "name=test&age=20",
		},
		{
			name: "empty",
			args: args{
				params: map[string]string{},
			},

			want: "",
		},
		{
			name: "nil",
			args: args{
				params: nil,
			},
		},
		{
			name: "one param",
			args: args{
				params: map[string]string{
					"name": "test",
				},
			},
			want: "name=test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeQuery(tt.args.params); got != tt.want {
				t.Errorf("makeQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonAnything(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{
				i: "test",
			},
			want: "\"test\"",
		},
		{
			name: "int",
			args: args{
				i: 1,
			},
			want: "1",
		},
		{
			name: "float",
			args: args{
				i: 1.1,
			},
			want: "1.1",
		},
		{
			name: "bool",
			args: args{
				i: true,
			},
			want: "true",
		},
		{
			name: "nil",
			args: args{
				i: nil,
			},
			want: "null",
		},
		{
			name: "array",
			args: args{
				i: []string{"test", "test2"},
			},
			want: "[\"test\",\"test2\"]",
		},
		{
			name: "map",
			args: args{
				i: map[string]string{"test": "test"},
			},
			want: "{\"test\":\"test\"}",
		},
		{
			name: "struct",
			args: args{
				i: struct {
					Name string
					Age  int
					Bool bool
				}{
					Name: "test",
					Age:  20,
					Bool: true,
				},
			},
			want: "{\"Name\":\"test\",\"Age\":20,\"Bool\":true}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jsonAnything(tt.args.i); got != tt.want {
				t.Errorf("jsonAnything() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMessage(t *testing.T) {
	m := NewMessage(1234, "test")
	if m.BaseChat.ChatID != 1234 {
		t.Errorf("NewMessage() ID = %v, want 1234", m.BaseChat.ChatID)
	}
	if m.Text != "test" {
		t.Errorf("NewMessage() Text = %v, want test", m.Text)
	}
}

func TestNewEditMessageTextAndMarkup(t *testing.T) {
	m := NewEditMessageTextAndMarkup(1234, 100, "test2", NewKeyboardWithMarkup(
		NewKeyboardRow(NewKeyboardButtonWithData("test", "test")), NewKeyboardRow(NewKeyboardButtonWithData("test2", "test2"))))
	params := m.getParams()
	if params["chat_id"] != "1234" {
		t.Errorf("NewEditMessageTextAndMarkup() ChatID = %v, want 1234", params["chat_id"])
	}

	if params["message_id"] != "100" {
		t.Errorf("NewEditMessageTextAndMarkup() MessageID = %v, want 100", params["message_id"])
	}

	if params["text"] != "test2" {
		t.Errorf("NewEditMessageTextAndMarkup() Text = %v, want test2", params["text"])
	}

	if params["reply_markup"] != "{\"inline_keyboard\":[[{\"text\":\"test\",\"callback_data\":\"test\"}],[{\"text\":\"test2\",\"callback_data\":\"test2\"}]]}" {
		t.Errorf("NewEditMessageTextAndMarkup() ReplyMarkup = %v, want {\"inline_keyboard\":[[{\"text\":\"test\",\"callback_data\":\"test\"}],[{\"text\":\"test2\",\"callback_data\":\"test2\"}]]}", params["reply_markup"])
	}
}

func TestClient_doRequest(t *testing.T) {
	c := New("test")
	request, err := c.doRequest("test", nil)
	if err != nil {
		t.Errorf("doRequest() error = %v, want nil", err)
	}

	if request.ErrorCode != 404 {
		t.Errorf("doRequest() = %v, want 404", request.ErrorCode)
	}
}
