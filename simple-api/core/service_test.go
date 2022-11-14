package core

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		r Repo
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSrv_GetAll(t *testing.T) {
	type fields struct {
		r Repo
	}
	tests := []struct {
		name   string
		fields fields
		want   []Todo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Srv{
				r: tt.fields.r,
			}
			if got := s.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Srv.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSrv_GetByID(t *testing.T) {
	type fields struct {
		r Repo
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Todo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Srv{
				r: tt.fields.r,
			}
			got, err := s.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Srv.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Srv.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSrv_Make(t *testing.T) {
	type fields struct {
		r Repo
	}
	type args struct {
		title string
		desc  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Todo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Srv{
				r: tt.fields.r,
			}
			if got := s.Make(tt.args.title, tt.args.desc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Srv.Make() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSrv_UpdateByID(t *testing.T) {
	type fields struct {
		r Repo
	}
	type args struct {
		id    uint
		title *string
		desc  *string
		done  *bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Srv{
				r: tt.fields.r,
			}
			if err := s.UpdateByID(tt.args.id, tt.args.title, tt.args.desc, tt.args.done); (err != nil) != tt.wantErr {
				t.Errorf("Srv.UpdateByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSrv_RemoveByID(t *testing.T) {
	type fields struct {
		r Repo
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Srv{
				r: tt.fields.r,
			}
			if err := s.RemoveByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Srv.RemoveByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
