//  Copyright (c) 2015 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package firestorm

import (
	"reflect"
	"testing"

	"github.com/acgshare/bleve/index"
)

func TestStoredRows(t *testing.T) {
	tests := []struct {
		input  index.IndexRow
		outKey []byte
		outVal []byte
	}{
		{
			NewStoredRow([]byte("doca"), 5, 7, nil, []byte("tcat")),
			[]byte{StoredKeyPrefix[0], 'd', 'o', 'c', 'a', ByteSeparator, 5, 7, 0},
			[]byte{10, 4, 't', 'c', 'a', 't'},
		},
		{
			NewStoredRow([]byte("doca"), 5, 7, []uint64{1, 1}, []byte("tcat")),
			[]byte{StoredKeyPrefix[0], 'd', 'o', 'c', 'a', ByteSeparator, 5, 7, 0, 1, 1},
			[]byte{10, 4, 't', 'c', 'a', 't'},
		},
	}

	// test going from struct to k/v bytes
	for i, test := range tests {
		rk := test.input.Key()
		if !reflect.DeepEqual(rk, test.outKey) {
			t.Errorf("Expected key to be %v got: %v", test.outKey, rk)
		}
		rv := test.input.Value()
		if !reflect.DeepEqual(rv, test.outVal) {
			t.Errorf("Expected value to be %v got: %v for %d", test.outVal, rv, i)
		}
	}

	// now test going back from k/v bytes to struct
	for i, test := range tests {
		row, err := NewStoredRowKV(test.outKey, test.outVal)
		if err != nil {
			t.Errorf("error parsking key/value: %v", err)
		}
		if !reflect.DeepEqual(row, test.input) {
			t.Errorf("Expected: %#v got: %#v for %d", test.input, row, i)
		}
	}
}
