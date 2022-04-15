package internal

import (
	"reflect"
	"testing"
)

func TestGetKeyOperator(t *testing.T) {
	tests := []struct {
		name string
		want Operators
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetKeyOperator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKeyOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_keyOp_Degenerate(t *testing.T) {
	type fields struct {
		Template string
	}
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kp := &keyOp{
				Template: tt.fields.Template,
			}
			got, got1, err := kp.Degenerate(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Degenerate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Degenerate() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Degenerate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_keyOp_Generate(t *testing.T) {
	//type fields struct {
	//	Template string
	//}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		//fields fields
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				x: 5,
				y: 50,
			},
			want: "5_50",
		},
		{
			name: "success large integers",
			args: args{
				x: 50000,
				y: 999999,
			},
			want: "50000_999999",
		},
	}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		kp := &keyOp{
	//			Template: tt.fields.Template,
	//		}
	//		if got := kp.Generate(tt.args.x, tt.args.y); got != tt.want {
	//			t.Errorf("Generate() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kp := GetKeyOperator()
			if got := kp.Generate(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("keyOp.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
