package chessboard_cipher

import "testing"

func TestPolybiusEncrypt(t *testing.T) {
	type args struct {
		text       string
		chessboard []Chessboard
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				text: "CC",
			},
			want: "1313",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PolybiusEncrypt(tt.args.text, tt.args.chessboard...)
			if (err != nil) != tt.wantErr {
				t.Errorf("PolybiusEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PolybiusEncrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
