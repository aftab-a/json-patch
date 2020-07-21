package skiparrays

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/emirpasic/gods/lists/arraylist"
)

func TestNew(t *testing.T) {
	type args struct {
		rowlen int
		values []int
	}
	tests := []struct {
		name string
		args args
		want *SkipArray
	}{
		{
			name: "Base case with empty values",
			args: args{
				rowlen: 10,
				values: []int{},
			},
			want: &SkipArray{
				rows: []*row{&row{
					pos:      0,
					elements: arraylist.New(),
				}},
				rowLen: 10,
			},
		},
		{
			name: "Base case with 9 values",
			args: args{
				rowlen: 10,
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: &SkipArray{
				rows: []*row{&row{
					pos:      0,
					elements: arraylist.New(1, 2, 3, 4, 5, 6, 7, 8, 9),
				}},
				rowLen: 10,
			},
		},
		{
			name: "Base case with 13 values",
			args: args{
				rowlen: 10,
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			},
			want: &SkipArray{
				rows: []*row{&row{
					pos:      0,
					elements: arraylist.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
				},
					&row{
						pos:      10,
						elements: arraylist.New(11, 12, 13),
					}},
				rowLen: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := make([]interface{}, len(tt.args.values))
			for i, v := range tt.args.values {
				s[i] = v
			}
			if got := New(tt.args.rowlen, s...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v %v %v", *got, *tt.want, len(got.rows), len(tt.want.rows))
			}
		})
	}
}

func cmpArr(s1, s2 *SkipArray) bool {
	if s1.rowLen != s2.rowLen {
		return false
	}

	if len(s1.rows) != len(s2.rows) {
		return false
	}

	/*if !reflect.DeepEqual(s1.rows, s2.rows) {
		return false
	}*/

	for i := range s1.rows {
		if s1.rows[i].pos != s2.rows[i].pos {
			log.Println("Pos mismatch for row", i, s1.rows[i].pos, s2.rows[i].pos)
			return false
		}
		v1 := s1.rows[i].elements.Values()
		v2 := s2.rows[i].elements.Values()

		if len(v1) != len(v2) {
			fmt.Println("Mismatched row len", len(v1), len(v2))
			fmt.Println(v1, v2)
			return false
		}

		for j := range v1 {
			if !reflect.DeepEqual(v1[j], v2[j]) {
				fmt.Println(v1[j], v2[j])
				return false
			} else {
				fmt.Println(v1[j], v2[j])
			}
		}
	}
	return true
}

func TestSkipArray_Insert(t *testing.T) {
	type fields struct {
		rows   []*row
		rowLen int
	}
	type args struct {
		pos   int
		value interface{}
	}
	type want struct {
		rnum   int
		rpos   int
		values []int
		ridx   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{

		{
			name: "Insert at pos -1",
			fields: fields{
				rows: []*row{&row{
					pos:      0,
					elements: arraylist.New(1, 2, 3, 4, 5, 6, 7, 8, 9),
				}},
				rowLen: 10,
			},
			args: args{
				pos:   0,
				value: 1,
			},
			want: want{
				rnum:   0,
				rpos:   -1,
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
		{
			name: "Insert at pos 0",
			fields: fields{
				rows: []*row{&row{
					pos:      0,
					elements: arraylist.New(1, 2, 3, 4, 5, 6, 7, 8, 9),
				}},
				rowLen: 10,
			},
			args: args{
				pos:   0,
				value: 1,
			},
			want: want{
				rnum:   0,
				rpos:   0,
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
		{
			name: "Insert at mid pos",
			fields: fields{
				rows: []*row{&row{
					pos:      0,
					elements: arraylist.New(1, 2, 3, 4, 5, 6, 7, 8, 9),
				}},
				rowLen: 10,
			},
			args: args{
				pos:   3,
				value: 1,
			},
			want: want{
				rnum:   0,
				rpos:   3,
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
		{
			name: "Insert at end pos",
			fields: fields{
				rows: []*row{&row{
					pos:      0,
					elements: arraylist.New(1, 2, 3, 4, 5, 6, 7, 8, 9),
				}},
				rowLen: 10,
			},
			args: args{
				pos:   9,
				value: 1,
			},
			want: want{
				rnum:   0,
				rpos:   9,
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
		{
			name: "Insert at first pos of second array",
			fields: fields{
				rows: []*row{&row{
					pos:      0,
					elements: arraylist.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
				},
					&row{
						pos:      10,
						elements: arraylist.New(11, 12, 13),
					}},
				rowLen: 10,
			},
			args: args{
				pos:   10,
				value: 1,
			},
			want: want{
				rnum:   1,
				rpos:   0,
				values: []int{11, 12, 13},
				ridx:   10,
			},
		},
		{
			name: "Insert at second pos of first array",
			fields: fields{
				rows: []*row{&row{
					pos:      0,
					elements: arraylist.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
				},
					&row{
						pos:      10,
						elements: arraylist.New(11, 12, 13),
					}},
				rowLen: 10,
			},
			args: args{
				pos:   1,
				value: 1,
			},
			want: want{
				rnum:   0,
				rpos:   1,
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				ridx:   0,
			},
		},
		{
			name: "Insert at last pos of second array",
			fields: fields{
				rows: []*row{&row{
					pos:      0,
					elements: arraylist.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
				},
					&row{
						pos:      10,
						elements: arraylist.New(11, 12, 13),
					}},
				rowLen: 10,
			},
			args: args{
				pos:   12,
				value: 1,
			},
			want: want{
				rnum:   1,
				rpos:   2,
				values: []int{11, 12, 13},
				ridx:   10,
			},
		},
		{
			name: "Insert at last pos of second array",
			fields: fields{
				rows: []*row{&row{
					pos:      0,
					elements: arraylist.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
				},
					&row{
						pos:      10,
						elements: arraylist.New(11, 12, 13, 14, 15, 16),
					}},
				rowLen: 10,
			},
			args: args{
				pos:   12,
				value: 1,
			},
			want: want{
				rnum:   1,
				rpos:   2,
				values: []int{11, 12, 13, 14, 15, 16},
				ridx:   10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &SkipArray{
				rows:   tt.fields.rows,
				rowLen: tt.fields.rowLen,
			}
			rows := make([]*row, len(tt.fields.rows))
			for i := range tt.fields.rows {
				rows[i] = tt.fields.rows[i]
			}
			a2 := &SkipArray{
				rows:   rows,
				rowLen: tt.fields.rowLen,
			}
			a.Insert(tt.args.pos, tt.args.value)
			if tt.want.rpos >= 0 {
				s := make([]interface{}, len(tt.want.values))
				for i, v := range tt.want.values {
					s[i] = v
				}
				alist := arraylist.New(s...)
				alist.Insert(tt.want.rpos, tt.args.value)
				r := newRow(tt.want.ridx)
				r.elements = alist
				for i := tt.want.rnum + 1; i < len(a2.rows)-1; i++ {
					a2.rows[i].pos++
				}
				a2.rows[tt.want.rnum] = r
				//if !reflect.DeepEqual(a, a2) {
				if !cmpArr(a, a2) {
					t.Errorf("Inserted() = %v, want %v", *a, *a2)
				}
			}
		})
	}
}

func TestSkipArray_Remove(t *testing.T) {
	type fields struct {
		rows   []*row
		rowLen int
	}
	type args struct {
		pos int
	}
	type want struct {
		rnum   int
		rpos   int
		values []int
		ridx   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "Remove the element at index 0",
			fields: fields{
				rows: []*row{&row{
					pos:      0,
					elements: arraylist.New(1, 2, 3, 4, 5, 6, 7, 8, 9),
				}},
				rowLen: 10,
			},
			args: args{
				pos: 0,
			},
			want: want{
				rnum:   0,
				rpos:   0,
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				ridx:   0,
			},
		},
		{
			name: "Remove the element at last index",
			fields: fields{
				rows: []*row{&row{
					pos:      0,
					elements: arraylist.New(1, 2, 3, 4, 5, 6, 7, 8, 9),
				}},
				rowLen: 10,
			},
			args: args{
				pos: 9,
			},
			want: want{
				rnum:   0,
				rpos:   9,
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				ridx:   0,
			},
		},
		{
			name: "Remove the mid element",
			fields: fields{
				rows: []*row{&row{
					pos:      0,
					elements: arraylist.New(1, 2, 3, 4, 5, 6, 7, 8, 9),
				}},
				rowLen: 10,
			},
			args: args{
				pos: 5,
			},
			want: want{
				rnum:   0,
				rpos:   5,
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				ridx:   0,
			},
		},
		{
			name: "Remove the first pos of second array",
			fields: fields{
				rows: []*row{&row{
					pos:      0,
					elements: arraylist.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
				},
					&row{
						pos:      10,
						elements: arraylist.New(11, 12, 13),
					}},
				rowLen: 10,
			},
			args: args{
				pos: 10,
			},
			want: want{
				rnum:   1,
				rpos:   0,
				values: []int{11, 12, 13},
				ridx:   10,
			},
		},
		{
			name: "Remove the first pos of first array",
			fields: fields{
				rows: []*row{&row{
					pos:      0,
					elements: arraylist.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
				},
					&row{
						pos:      10,
						elements: arraylist.New(11, 12, 13),
					}},
				rowLen: 10,
			},
			args: args{
				pos: 0,
			},
			want: want{
				rnum:   0,
				rpos:   0,
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				ridx:   0,
			},
		},
		{
			name: "Remove the last pos of last array",
			fields: fields{
				rows: []*row{&row{
					pos:      0,
					elements: arraylist.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
				},
					&row{
						pos:      10,
						elements: arraylist.New(11, 12, 13),
					}},
				rowLen: 10,
			},
			args: args{
				pos: 12,
			},
			want: want{
				rnum:   1,
				rpos:   2,
				values: []int{11, 12, 13},
				ridx:   10,
			},
		},
		{
			name: "Remove the last element of second array",
			fields: fields{
				rows: []*row{&row{
					pos:      0,
					elements: arraylist.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
				},
					&row{
						pos:      10,
						elements: arraylist.New(11),
					},
					&row{
						pos:      11,
						elements: arraylist.New(12, 13),
					}},
				rowLen: 10,
			},
			args: args{
				pos: 10,
			},
			want: want{
				rnum:   1,
				rpos:   0,
				values: []int{11},
				ridx:   10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &SkipArray{
				rows:   tt.fields.rows,
				rowLen: tt.fields.rowLen,
			}
			rows := make([]*row, len(tt.fields.rows))
			for i := range tt.fields.rows {
				rows[i] = tt.fields.rows[i]
			}
			a2 := &SkipArray{
				rows:   rows,
				rowLen: tt.fields.rowLen,
			}
			a.Remove(tt.args.pos)
			if tt.want.rpos >= 0 {
				s := make([]interface{}, len(tt.want.values))
				for i, v := range tt.want.values {
					s[i] = v
				}
				alist := arraylist.New(s...)
				//fmt.Println("Size of alist", alist.Size())
				alist.Remove(tt.want.rpos)
				//fmt.Println("Size of alist", alist.Size())
				if alist.Size() > 0 {
					r := newRow(tt.want.ridx)
					r.elements = alist
					for i := tt.want.rnum + 1; i < len(a2.rows)-1; i++ {
						a2.rows[i].pos--
					}
					a2.rows[tt.want.rnum] = r
					//if !reflect.DeepEqual(a, a2) {
				} else {
					copy(a2.rows[tt.want.rnum:], a2.rows[tt.want.rnum+1:])
					a2.rows = a2.rows[:len(a2.rows)-1]
				}
				if !cmpArr(a, a2) {
					t.Errorf("Inserted() = %v, want %v", *a, *a2)
				}
			}
		})
	}
}
