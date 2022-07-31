package main

import (
	"reflect"
	"testing"
)

func TestNewEmployee(t *testing.T) {
	type args struct {
		_name     string
		_position string
		_location string
	}
	tests := []struct {
		name string
		args args
		want Employee
	}{
		{name: "Test 1", args: args{_name: "Gain", _position: "Backend Engineer", _location: USA}, want: NewEmployee("Gain", "Backend Engineer", "USA")},
		{name: "Test 2", args: args{_name: "Sam", _position: "Frontend Engineer", _location: CHINA}, want: NewEmployee("Sam", "Frontend Engineer", "China")},
		{name: "Test 3", args: args{_name: "Jane", _position: "DevOps Engineer", _location: CHINA}, want: NewEmployee("Jane", "DevOps Engineer", "China")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmployee(tt.args._name, tt.args._position, tt.args._location); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmployee() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestNewEmployeeFactory(t *testing.T) {
//	type args struct {
//		location string
//		position string
//		salary   int
//	}
//	tests := []struct {
//		name string
//		args args
//		want EmployeeFactory
//	}{
//		{name: "Test 1", args: args{location: USA, position: "DevOps Engineer", salary: 200000}, want: NewEmployeeFactory("USA", "DevOps Engineer", 200000)},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewEmployeeFactory(tt.args.location, tt.args.position, tt.args.salary); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewEmployeeFactory() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
