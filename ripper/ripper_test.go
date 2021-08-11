package ripper

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func mustReadFile(t *testing.T, path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	return string(content)
}

func TestRipJSON(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "sample1",
			args: args{text: mustReadFile(t, "testdata/sample1.txt")},
			want: []byte("[{\"changed\":false,\"message\":\"hello world!\"}]"),
			wantErr: false,
		},
		{
			name: "sample2",
			args: args{text: mustReadFile(t, "testdata/sample2.txt")},
			want: []byte("[{\"random\":\"JSON\"},{\"why\":\"is this here?\"}]"),
			wantErr: false,
		},
		{
			name: "sample3",
			args: args{text: mustReadFile(t, "testdata/sample3.txt")},
			want: []byte("[{\"changed\":false,\"some-json\":{\"hello\":\"world\"}}]"),
			wantErr: false,
		},
		{
			name: "sample4",
			args: args{text: mustReadFile(t, "testdata/sample4.txt")},
			want: []byte("[]"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RipJSON(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("RipJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RipJSON() got = %s, want %s", got, tt.want)
			}
		})
	}
}
