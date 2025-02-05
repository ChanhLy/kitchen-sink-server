package users

import (
	"context"
	"database/sql"
	database "go-server/db"
	"reflect"
	"testing"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		data database.CreateUserParams
		ctx  context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    database.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "first user",
			args: args{
				data: database.CreateUserParams{
					Name:  "a",
					Email: sql.NullString{String: "a@a.com", Valid: true},
				},
				ctx: context.Background(),
			},
			want: database.User{
				ID:    1,
				Name:  "a",
				Email: sql.NullString{String: "a@a.com", Valid: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateUser(tt.args.data, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserById(t *testing.T) {
	type args struct {
		id  int64
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    database.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "a",
			args: args{
				id:  1,
				ctx: context.Background(),
			},
			want: database.User{
				ID:    1,
				Name:  "a",
				Email: sql.NullString{String: "a@a.com", Valid: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateUser(database.CreateUserParams{
				Name:  tt.want.Name,
				Email: tt.want.Email,
			}, tt.args.ctx)
			if err != nil {
				t.Errorf("GetUserById() error = %v, unable to create user", err)
				return
			}

			got, err := GetUserById(tt.args.id, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUsers(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []database.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}
