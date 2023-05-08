package client

import (
	"reflect"
	"testing"
)

func TestNewKeyboardRow(t *testing.T) {
	type args struct {
		buttons []KeyboardButton
	}
	tests := []struct {
		name string
		args args
		want []KeyboardButton
	}{
		{
			name: "empty",
			args: args{
				buttons: []KeyboardButton{},
			},
			want: []KeyboardButton{},
		},
		{
			name: "one",
			args: args{
				buttons: []KeyboardButton{
					{
						Text: "one",
					},
				},
			},
			want: []KeyboardButton{
				{
					Text: "one",
				},
			},
		},
		{
			name: "two",
			args: args{
				buttons: []KeyboardButton{
					{
						Text: "one",
					},
					{
						Text: "two",
					},
				},
			},
			want: []KeyboardButton{
				{
					Text: "one",
				},
				{
					Text: "two",
				},
			},
		},
		{
			name: "nil",
			args: args{
				buttons: nil,
			},
			want: []KeyboardButton{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKeyboardRow(tt.args.buttons...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKeyboardRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKeyboardWithMarkup(t *testing.T) {
	type args struct {
		rows [][]KeyboardButton
	}
	tests := []struct {
		name string
		args args
		want InlineKeyboardMarkup
	}{
		{
			name: "empty",
			args: args{
				rows: [][]KeyboardButton{},
			},
			want: InlineKeyboardMarkup{[][]KeyboardButton{}},
		},
		{
			name: "one",
			args: args{
				rows: [][]KeyboardButton{
					{
						{
							Text: "one",
						},
					},
				},
			},
			want: InlineKeyboardMarkup{
				InlineKeyboard: [][]KeyboardButton{
					{
						{
							Text: "one",
						},
					},
				},
			},
		},
		{
			name: "three",
			args: args{
				rows: [][]KeyboardButton{
					{
						{
							Text: "one",
						},
					},
					{
						{
							Text: "two",
						},
					},
					{
						{
							Text: "three",
						},
					},
				},
			},
			want: InlineKeyboardMarkup{
				InlineKeyboard: [][]KeyboardButton{
					{
						{
							Text: "one",
						},
					},
					{
						{
							Text: "two",
						},
					},
					{
						{
							Text: "three",
						},
					},
				},
			},
		},
		{
			name: "nil",
			args: args{
				rows: nil,
			},
			want: InlineKeyboardMarkup{[][]KeyboardButton{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKeyboardWithMarkup(tt.args.rows...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKeyboardWithMarkup() = %v, want %v", got, tt.want)
			}
		})
	}
}
