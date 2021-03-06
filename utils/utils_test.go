package utils

import (
	"reflect"
	"testing"

	"github.com/gan-of-culture/go-hentai-scraper/config"
)

func TestGetLastItem(t *testing.T) {
	tests := []struct {
		name string
		list []string
		want string
	}{
		{
			name: "String slice",
			list: []string{"1", "2", "3", "last item"},
			want: "last item",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := GetLastItemString(tt.list)

			if item != tt.want {
				t.Errorf("Got: %v - want: %v", item, tt.want)
			}
		})
	}
}

func TestCalcSizeInByte(t *testing.T) {
	tests := []struct {
		name   string
		number float64
		unit   string
		want   int64
	}{
		{
			name:   "Kilobytes to Bytes",
			number: 752,
			unit:   "KB",
			want:   752000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes := CalcSizeInByte(tt.number, tt.unit)

			if bytes != tt.want {
				t.Errorf("Got: %v - want: %v", bytes, tt.want)
			}
		})
	}
}

func TestNeedDownloadList(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		pages string
	}{
		{
			name: "pages test",
			args: args{
				len: 3,
			},
			pages: "1, 3",
			want:  []int{1, 3},
		},
		{
			name: "from to item selection 1",
			args: args{
				len: 10,
			},
			pages: "1-3, 5, 7-8, 10",
			want:  []int{1, 2, 3, 5, 7, 8, 10},
		},
		{
			name: "from to item selection 2",
			args: args{
				len: 10,
			},
			pages: "1,2, 4 , 5, 7-8  , 10",
			want:  []int{1, 2, 4, 5, 7, 8, 10},
		},
		{
			name: "from to item selection 3",
			args: args{
				len: 10,
			},
			pages: "5-1, 2",
			want:  []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.Pages = tt.pages
			if got := NeedDownloadList(tt.args.len); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NeedDownloadList() = %v, want %v", got, tt.want)
			}
		})
	}
}
