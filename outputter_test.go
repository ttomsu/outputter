package outputter_test

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/ttomsu/outputter"
)

func TestWriting(t *testing.T) {
	type args struct {
		opts []outputter.Option
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no opts",
			want: "Hello, World!",
		},
		{
			name: "buffered",
			want: "Hello, World!",
			args: args{
				opts: []outputter.Option{
					outputter.WithBufferedStdOut(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outBuf, errBuf := &bytes.Buffer{}, &bytes.Buffer{}
			opts := append([]outputter.Option{outputter.WithStdOut(outBuf), outputter.WithStdErr(errBuf)}, tt.args.opts...)
			out := outputter.New(opts...)
			fmt.Fprint(out.StdOut, tt.want)
			if err := out.Close(); err != nil {
				t.Fatalf("error closing: %w", err)
			}
			if !reflect.DeepEqual(outBuf.String(), tt.want) {
				t.Errorf("Printed = %q, want %q", outBuf.String(), tt.want)
			}
		})
	}
}
