package base62

import "testing"

func TestEncode(t *testing.T) {
	type args struct {
		seq uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestEncode",
			args: args{seq: 12345},
			want: "DNH",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.seq); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		seq string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "TestDecode",
			args: args{seq: "DNH"},
			want: 12345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.args.seq); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
