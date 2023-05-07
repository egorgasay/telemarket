package storage

import (
	"bot/internal/entity"
	"bot/pkg"
	"math"
	"reflect"
	"sync"
	"testing"
)

func TestStorage_GetItemByName(t *testing.T) {
	type fields struct {
		RWMutex    *sync.RWMutex
		items      map[string]entity.Item
		allRates   int
		countRates int
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Item
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				items: map[string]entity.Item{
					"test": {
						Name:  "test",
						Price: 100,
					},
				},
			},
			args: args{
				name: "test",
			},
			want: entity.Item{
				Name:  "test",
				Price: 100,
			},
		},
		{
			name: "not found",
			fields: fields{
				items: map[string]entity.Item{
					"test": {
						Name: "test",
					},
				},
			},
			args: args{
				name: "test2",
			},
			want:    entity.Item{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				items:      tt.fields.items,
				allRates:   tt.fields.allRates,
				countRates: tt.fields.countRates,
			}
			got, err := s.GetItemByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetItemByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItemByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetAll(t *testing.T) {
	type fields struct {
		items      map[string]entity.Item
		allRates   int
		countRates int
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "ok",
			fields: fields{
				items: map[string]entity.Item{
					"test": {
						Name:  "test",
						Price: 100,
					},
					"test2": {
						Name:  "test2",
						Price: 200,
					},
					"test3": {
						Name:  "test3",
						Price: 300,
					},
				},
			},
			want: []string{
				"test", "test2", "test3",
			},
		},
		{
			name: "empty",
			fields: fields{
				items: map[string]entity.Item{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				RWMutex:    sync.RWMutex{},
				items:      tt.fields.items,
				allRates:   tt.fields.allRates,
				countRates: tt.fields.countRates,
			}
			if got := s.GetAll(); !pkg.Like(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetAVG(t *testing.T) {
	s := &Storage{
		items: map[string]entity.Item{
			"test": {
				Name:  "test",
				Price: 100,
			},

			"test2": {
				Name:  "test2",
				Price: 200,
			},
		},
		allRates:   0,
		countRates: 0,
	}

	avg := s.GetAVG()
	if !math.IsNaN(avg) {
		t.Errorf("GetAVG() = %v, want NaN", avg)
	}

	s.countRates = 2
	s.allRates = 100

	avg = s.GetAVG()
	if avg != 50 {
		t.Errorf("GetAVG() = %v, want 50", avg)
	}

	s.countRates = 3
	s.allRates = 47

	avg = s.GetAVG()
	if avg != 15.666666666666666 {
		t.Errorf("GetAVG() = %v, want 15.666666666666666", avg)
	}
}

func TestStorage_AddRate(t *testing.T) {
	s := &Storage{
		items: map[string]entity.Item{
			"test": {
				Name:  "test",
				Price: 100,
			},

			"test2": {
				Name:  "test2",
				Price: 200,
			},
		},
		allRates:   0,
		countRates: 0,
	}

	s.AddRate(100)

	if s.allRates != 100 && s.countRates != 1 {
		t.Errorf("AddRate() = %v, %v, want 100, 1", s.allRates, s.countRates)
	}

	s.AddRate(200)

	if s.allRates != 300 && s.countRates != 2 {
		t.Errorf("AddRate() = %v, %v, want 300, 2", s.allRates, s.countRates)
	}

	s.AddRate(300)

	if s.allRates != 600 && s.countRates != 3 {
		t.Errorf("AddRate() = %v, %v, want 600, 3", s.allRates, s.countRates)
	}
}
