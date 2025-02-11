/*
Copyright 2020 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package evalengine

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"vitess.io/vitess/go/sqltypes"
)

// more tests in go/sqlparser/expressions_test.go

func TestBinaryOpTypes(t *testing.T) {
	type testcase struct {
		l, r, e sqltypes.Type
	}
	type ops struct {
		op        ArithmeticOp
		testcases []testcase
	}

	tests := []ops{
		{
			op: &OpAddition{},
			testcases: []testcase{
				{sqltypes.Int64, sqltypes.Int64, sqltypes.Int64},
				{sqltypes.Uint64, sqltypes.Int64, sqltypes.Uint64},
				{sqltypes.Float64, sqltypes.Int64, sqltypes.Float64},
				{sqltypes.Int64, sqltypes.Uint64, sqltypes.Int64},
				{sqltypes.Uint64, sqltypes.Uint64, sqltypes.Uint64},
				{sqltypes.Float64, sqltypes.Uint64, sqltypes.Float64},
				{sqltypes.Int64, sqltypes.Float64, sqltypes.Int64},
				{sqltypes.Uint64, sqltypes.Float64, sqltypes.Uint64},
				{sqltypes.Float64, sqltypes.Float64, sqltypes.Float64},
			},
		}, {
			op: &OpSubstraction{},
			testcases: []testcase{
				{sqltypes.Int64, sqltypes.Int64, sqltypes.Int64},
				{sqltypes.Uint64, sqltypes.Int64, sqltypes.Uint64},
				{sqltypes.Float64, sqltypes.Int64, sqltypes.Float64},
				{sqltypes.Int64, sqltypes.Uint64, sqltypes.Int64},
				{sqltypes.Uint64, sqltypes.Uint64, sqltypes.Uint64},
				{sqltypes.Float64, sqltypes.Uint64, sqltypes.Float64},
				{sqltypes.Int64, sqltypes.Float64, sqltypes.Int64},
				{sqltypes.Uint64, sqltypes.Float64, sqltypes.Uint64},
				{sqltypes.Float64, sqltypes.Float64, sqltypes.Float64},
			},
		}, {
			op: &OpMultiplication{},
			testcases: []testcase{
				{sqltypes.Int64, sqltypes.Int64, sqltypes.Int64},
				{sqltypes.Uint64, sqltypes.Int64, sqltypes.Uint64},
				{sqltypes.Float64, sqltypes.Int64, sqltypes.Float64},
				{sqltypes.Int64, sqltypes.Uint64, sqltypes.Int64},
				{sqltypes.Uint64, sqltypes.Uint64, sqltypes.Uint64},
				{sqltypes.Float64, sqltypes.Uint64, sqltypes.Float64},
				{sqltypes.Int64, sqltypes.Float64, sqltypes.Int64},
				{sqltypes.Uint64, sqltypes.Float64, sqltypes.Uint64},
				{sqltypes.Float64, sqltypes.Float64, sqltypes.Float64},
			},
		}, {
			op: &OpDivision{},
			testcases: []testcase{
				{sqltypes.Int64, sqltypes.Int64, sqltypes.Float64},
				{sqltypes.Uint64, sqltypes.Int64, sqltypes.Float64},
				{sqltypes.Float64, sqltypes.Int64, sqltypes.Float64},
				{sqltypes.Int64, sqltypes.Uint64, sqltypes.Float64},
				{sqltypes.Uint64, sqltypes.Uint64, sqltypes.Float64},
				{sqltypes.Float64, sqltypes.Uint64, sqltypes.Float64},
				{sqltypes.Int64, sqltypes.Float64, sqltypes.Float64},
				{sqltypes.Uint64, sqltypes.Float64, sqltypes.Float64},
				{sqltypes.Float64, sqltypes.Float64, sqltypes.Float64},
			},
		},
	}

	for _, op := range tests {
		for _, tc := range op.testcases {
			name := fmt.Sprintf("%s %s %s", tc.l.String(), reflect.TypeOf(op.op).String(), tc.r.String())
			t.Run(name, func(t *testing.T) {
				result := op.op.typeof(tc.l)
				assert.Equal(t, tc.e, result)
			})
		}
	}
}
