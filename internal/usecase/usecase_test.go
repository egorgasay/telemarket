package usecase

import (
	"bot/internal/entity"
	"bot/internal/storage"
	"bot/pkg"
	"reflect"
	"testing"
)

func TestUseCase_AddRate(t *testing.T) {
	uc := UseCase{storage: storage.New()}
	uc.AddRate(100)

	avg := uc.storage.GetAVG()
	if avg != 100 {
		t.Errorf("avg = %v, want %v", avg, 100)
	}

	uc.AddRate(1)

	avg = uc.storage.GetAVG()
	if avg != 50.5 {
		t.Errorf("avg = %v, want %v", avg, 50.5)
	}

	uc.AddRate(44)

	avg = uc.storage.GetAVG()
	if avg != 48.333333333333336 {
		t.Errorf("avg = %v, want %v", avg, 48.333333333333336)
	}
}

func TestUseCase_GetAll(t *testing.T) {
	type fields struct {
		storage *storage.Storage
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			name:   "ok",
			fields: fields{storage.New()},
			want: []string{
				"HATE ⬜️", "HATE ⬛️",
			},
		},
		{
			name:    "empty",
			fields:  fields{&storage.Storage{}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				storage: tt.fields.storage,
			}
			got, err := u.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !pkg.Like(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetItemByName(t *testing.T) {
	type fields struct {
		storage *storage.Storage
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
			name:   "ok",
			fields: fields{storage: storage.New()},
			args: args{
				name: "HATE ⬜️",
			},
			want: entity.Item{
				Name:        "HATE ⬜",
				Price:       1500,
				Quantity:    0,
				Description: "100% хлопок.",
			},
		},
		{
			name:   "not found",
			fields: fields{storage: storage.New()},
			args: args{
				name: "test2",
			},
			want:    entity.Item{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				storage: tt.fields.storage,
			}
			got, err := u.GetItemByName(tt.args.name)
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

func TestUseCase_GetRate(t *testing.T) {
	s := storage.New()
	s.AddRate(100)
	s.AddRate(1)
	s.AddRate(44)

	uc := UseCase{storage: s}

	rate := uc.GetRate()
	if rate != 48.34 {
		t.Errorf("rate = %v, want %v", rate, 48.34)
	}

	uc.AddRate(100)

	rate = uc.GetRate()
	if rate != 61.25 {
		t.Errorf("rate = %v, want %v", rate, 61.25)
	}
}
