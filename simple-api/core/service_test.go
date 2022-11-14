package core

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

type MockRepo struct{}

var mockData = []Todo{
	{
		ID:        1,
		Title:     "What is Lorem Ipsum?",
		Desc:      "is simply dummy text of the printing and typesetting industry.",
		Done:      true,
		CreatedAt: time.Now(),
	},
	{
		ID:        2,
		Title:     "Where does it come from?",
		Desc:      "Contrary to popular belief, Lorem Ipsum is not simply random text.",
		Done:      false,
		CreatedAt: time.Now(),
	},
	{
		ID:        3,
		Title:     "Why do we use it?",
		Desc:      "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout.",
		Done:      false,
		CreatedAt: time.Now(),
	},
}

func (mr *MockRepo) Get() []Todo {
	return mockData
}

func (mr *MockRepo) GetByPK(id uint) (Todo, error) {
	index := -1
	for k, v := range mockData {
		if id == v.ID {
			index = k
			break
		}
	}
	if index == -1 {
		return Todo{}, errors.New("Not Found Todo")
	}
	return mockData[index], nil
}

func (mr *MockRepo) Create(title string, desc string) (Todo, error) {
	if title == "i make some mistake" {
		return Todo{}, errors.New("Make Todo Error")
	}
	lastTodo := len(mockData) - 1
	lastID := uint(0)
	if lastTodo > -1 {
		lastID = mockData[lastTodo].ID + 1
	}
	td := Todo{
		ID:        lastID,
		Title:     title,
		Desc:      desc,
		Done:      false,
		CreatedAt: time.Now(),
	}
	mockData = append(mockData, td)
	return td, nil
}

func (mr *MockRepo) UpdateByPk(id uint, title *string, desc *string, done *bool) error {
	index := -1
	for k, v := range mockData {
		if id == v.ID {
			index = k
			break
		}
	}
	if index == -1 {
		return errors.New("Not Found Todo")
	}

	mockData[index].Title = *title
	mockData[index].Desc = *desc
	mockData[index].Done = *done

	return nil
}

func (mr *MockRepo) RemoveByPK(id uint) error {
	index := -1
	for k, v := range mockData {
		if id == v.ID {
			index = k
			break
		}
	}
	if index == -1 {
		return errors.New("Not Found Todo")
	}
	if mockData[index].Done {
		return errors.New("Not Found Todo")
	}

	mockData = append(mockData[:index], mockData[index+1:]...)
	return nil
}

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
