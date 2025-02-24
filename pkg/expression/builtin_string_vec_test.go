// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package expression

import (
	"math/rand"
	"testing"

	"github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/pingcap/tidb/pkg/parser/charset"
	"github.com/pingcap/tidb/pkg/parser/mysql"
	"github.com/pingcap/tidb/pkg/types"
)

type randSpaceStrGener struct {
	lenBegin int
	lenEnd   int
}

func (g *randSpaceStrGener) gen() interface{} {
	n := rand.Intn(g.lenEnd-g.lenBegin) + g.lenBegin
	buf := make([]byte, n)
	for i := range buf {
		x := rand.Intn(150)
		if x < 10 {
			buf[i] = byte('0' + x)
		} else if x-10 < 26 {
			buf[i] = byte('a' + x - 10)
		} else if x < 62 {
			buf[i] = byte('A' + x - 10 - 26)
		} else {
			buf[i] = byte(' ')
		}
	}
	return string(buf)
}

var vecBuiltinStringCases = map[string][]vecExprBenchCase{
	ast.Length: {
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString}, geners: []dataGenerator{newDefaultGener(0.2, types.ETString)}},
	},
	ast.ASCII: {
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString}, geners: []dataGenerator{newDefaultGener(0.2, types.ETString)}},
	},
	ast.Concat: {
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString, types.ETString, types.ETString}},
	},
	ast.ConcatWS: {
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETString, types.ETString, types.ETString},
			geners:        []dataGenerator{&constStrGener{","}},
		},
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETString, types.ETString, types.ETString},
			geners:        []dataGenerator{newDefaultGener(1, types.ETString)},
		},
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETString, types.ETString},
			geners: []dataGenerator{
				&constStrGener{"<------------------>"},
				&constStrGener{"1413006"},
				&constStrGener{"idlfmv"},
			},
		},
	},
	ast.Convert: {
		{
			retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString, types.ETString},
			constants: []*Constant{nil, {Value: types.NewDatum("utf8"), RetType: types.NewFieldType(mysql.TypeString)}},
		},
		{
			retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString, types.ETString},
			constants: []*Constant{nil, {Value: types.NewDatum("binary"), RetType: types.NewFieldType(mysql.TypeString)}},
		},
		{
			retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString, types.ETString},
			constants: []*Constant{nil, {Value: types.NewDatum("utf8mb4"), RetType: types.NewFieldType(mysql.TypeString)}},
		},
		{
			retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString, types.ETString},
			constants: []*Constant{nil, {Value: types.NewDatum("ascii"), RetType: types.NewFieldType(mysql.TypeString)}},
		},
		{
			retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString, types.ETString},
			constants: []*Constant{nil, {Value: types.NewDatum("latin1"), RetType: types.NewFieldType(mysql.TypeString)}},
		},
	},
	ast.Substring: {
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETInt},
			geners:        []dataGenerator{newRandLenStrGener(0, 20), newRangeInt64Gener(-25, 25)},
		},
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETInt, types.ETInt},
			geners:        []dataGenerator{newRandLenStrGener(0, 20), newRangeInt64Gener(-25, 25), newRangeInt64Gener(-25, 25)},
		},
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETInt, types.ETInt},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
				types.NewFieldTypeBuilder().SetType(mysql.TypeLonglong).BuildP(),
				types.NewFieldTypeBuilder().SetType(mysql.TypeLonglong).BuildP()},
			geners: []dataGenerator{newRandLenStrGener(0, 20), newRangeInt64Gener(-25, 25), newRangeInt64Gener(-25, 25)},
		},
	},
	ast.SubstringIndex: {
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETString, types.ETInt},
			geners:        []dataGenerator{newRandLenStrGener(0, 20), newRandLenStrGener(0, 2), newRangeInt64Gener(-4, 4)},
		},
	},
	ast.Locate: {
		{
			retEvalType:   types.ETInt,
			childrenTypes: []types.EvalType{types.ETString, types.ETString},
			geners:        []dataGenerator{newRandLenStrGener(0, 10), newRandLenStrGener(0, 20)},
		},
		{
			retEvalType:   types.ETInt,
			childrenTypes: []types.EvalType{types.ETString, types.ETString},
			geners:        []dataGenerator{newRandLenStrGener(1, 2), newRandLenStrGener(0, 20)},
		},
		{
			retEvalType:   types.ETInt,
			childrenTypes: []types.EvalType{types.ETString, types.ETString},
			geners:        []dataGenerator{newSelectStringGener([]string{"01", "10", "001", "110", "0001", "1110"}), newSelectStringGener([]string{"010010001000010", "101101110111101"})},
		},
		{
			retEvalType:   types.ETInt,
			childrenTypes: []types.EvalType{types.ETString, types.ETString, types.ETInt},
			geners:        []dataGenerator{newRandLenStrGener(0, 10), newRandLenStrGener(0, 20), newRangeInt64Gener(-10, 20)},
		},
		{
			retEvalType:   types.ETInt,
			childrenTypes: []types.EvalType{types.ETString, types.ETString, types.ETInt},
			geners:        []dataGenerator{newRandLenStrGener(1, 2), newRandLenStrGener(0, 10), newRangeInt64Gener(0, 8)},
		},
		{
			retEvalType:   types.ETInt,
			childrenTypes: []types.EvalType{types.ETString, types.ETString},
			geners:        []dataGenerator{newSelectStringGener([]string{"01", "10", "001", "110", "0001", "1110"}), newSelectStringGener([]string{"010010001000010", "101101110111101"})},
		},
		{
			retEvalType:   types.ETInt,
			childrenTypes: []types.EvalType{types.ETString, types.ETString},
			childrenFieldTypes: []*types.FieldType{nil,
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP()},
			geners: []dataGenerator{newRandLenStrGener(0, 10), newRandLenStrGener(0, 20)},
		},
		{
			retEvalType:   types.ETInt,
			childrenTypes: []types.EvalType{types.ETString, types.ETString},
			childrenFieldTypes: []*types.FieldType{nil,
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP()},
			geners: []dataGenerator{newRandLenStrGener(1, 2), newRandLenStrGener(0, 20)},
		},
		{
			retEvalType:   types.ETInt,
			childrenTypes: []types.EvalType{types.ETString, types.ETString},
			childrenFieldTypes: []*types.FieldType{nil,
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP()},
			geners: []dataGenerator{newSelectStringGener([]string{"01", "10", "001", "110", "0001", "1110"}), newSelectStringGener([]string{"010010001000010", "101101110111101"})},
		},
		{
			retEvalType:   types.ETInt,
			childrenTypes: []types.EvalType{types.ETString, types.ETString},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(), nil},
			geners: []dataGenerator{newRandLenStrGener(0, 10), newRandLenStrGener(0, 20)},
		},
		{
			retEvalType:   types.ETInt,
			childrenTypes: []types.EvalType{types.ETString, types.ETString},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(), nil},
			geners: []dataGenerator{newRandLenStrGener(1, 2), newRandLenStrGener(0, 20)},
		},
		{
			retEvalType:   types.ETInt,
			childrenTypes: []types.EvalType{types.ETString, types.ETString},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(), nil},
			geners: []dataGenerator{newSelectStringGener([]string{"01", "10", "001", "110", "0001", "1110"}), newSelectStringGener([]string{"010010001000010", "101101110111101"})},
		},
		{
			retEvalType:   types.ETInt,
			childrenTypes: []types.EvalType{types.ETString, types.ETString},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP()},
			geners: []dataGenerator{newRandLenStrGener(0, 10), newRandLenStrGener(0, 20)},
		},
		{
			retEvalType:   types.ETInt,
			childrenTypes: []types.EvalType{types.ETString, types.ETString},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP()},
			geners: []dataGenerator{newRandLenStrGener(1, 2), newRandLenStrGener(0, 20)},
		},
		{
			retEvalType:   types.ETInt,
			childrenTypes: []types.EvalType{types.ETString, types.ETString},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP()},
			geners: []dataGenerator{newSelectStringGener([]string{"01", "10", "001", "110", "0001", "1110"}), newSelectStringGener([]string{"010010001000010", "101101110111101"})},
		},
		{
			retEvalType:   types.ETInt,
			childrenTypes: []types.EvalType{types.ETString, types.ETString, types.ETInt},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
				types.NewFieldTypeBuilder().SetType(mysql.TypeInt24).BuildP()},
			geners: []dataGenerator{newRandLenStrGener(0, 10), newRandLenStrGener(0, 20), newRangeInt64Gener(-10, 20)},
		},
		{
			retEvalType:   types.ETInt,
			childrenTypes: []types.EvalType{types.ETString, types.ETString, types.ETInt},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
				types.NewFieldTypeBuilder().SetType(mysql.TypeInt24).BuildP()},
			geners: []dataGenerator{newSelectStringGener([]string{"01", "10", "001", "110", "0001", "1110"}), newSelectStringGener([]string{"010010001000010", "101101110111101"}), newRangeInt64Gener(-10, 20)},
		},
	},
	ast.Hex: {
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString}, geners: []dataGenerator{newRandHexStrGener(10, 100)}},
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETInt}},
	},
	ast.Unhex: {
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString}, geners: []dataGenerator{newRandHexStrGener(10, 100)}},
	},
	ast.Trim: {
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString}, geners: []dataGenerator{&randSpaceStrGener{10, 100}}},
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString, types.ETString}, geners: []dataGenerator{newRandLenStrGener(10, 20), newRandLenStrGener(5, 25)}},
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETString, types.ETInt},
			geners:        []dataGenerator{newRandLenStrGener(10, 20), newRandLenStrGener(5, 25), newRangeInt64Gener(0, 4)},
			constants:     []*Constant{nil, nil, {Value: types.NewDatum(ast.TrimBoth), RetType: types.NewFieldType(mysql.TypeLonglong)}},
		},
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETString, types.ETInt},
			geners:        []dataGenerator{newRandLenStrGener(10, 20), newRandLenStrGener(5, 25), newRangeInt64Gener(0, 4)},
			constants:     []*Constant{nil, nil, {Value: types.NewDatum(ast.TrimLeading), RetType: types.NewFieldType(mysql.TypeLonglong)}},
		},
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETString, types.ETInt},
			geners:        []dataGenerator{newRandLenStrGener(10, 20), newRandLenStrGener(5, 25), newRangeInt64Gener(0, 4)},
			constants:     []*Constant{nil, nil, {Value: types.NewDatum(ast.TrimTrailing), RetType: types.NewFieldType(mysql.TypeLonglong)}},
		},
	},
	ast.Translate: {
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETString, types.ETString},
			geners:        []dataGenerator{newRandLenStrGener(10, 20), newRandLenStrGener(5, 25), newRandLenStrGener(5, 25)},
		},
	},
	ast.LTrim: {
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString}, geners: []dataGenerator{&randSpaceStrGener{10, 100}}},
	},
	ast.RTrim: {
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString}, geners: []dataGenerator{&randSpaceStrGener{10, 100}}},
	},
	ast.Lpad: {
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETInt, types.ETString},
			geners:        []dataGenerator{newRandLenStrGener(0, 20), newRangeInt64Gener(168435456, 368435456), newRandLenStrGener(0, 10)},
		},
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETInt, types.ETString},
			geners:        []dataGenerator{newDefaultGener(0.2, types.ETString), newDefaultGener(0.2, types.ETInt), newDefaultGener(0.2, types.ETString)},
		},
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETInt, types.ETString},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP()},
			geners: []dataGenerator{newRandLenStrGener(0, 20), newRangeInt64Gener(168435456, 368435456), newRandLenStrGener(0, 10)},
		},
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETInt, types.ETString},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP()},
			geners: []dataGenerator{newDefaultGener(0.2, types.ETString), newDefaultGener(0.2, types.ETInt), newDefaultGener(0.2, types.ETString)},
		},
	},
	ast.Rpad: {
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETInt, types.ETString},
			geners:        []dataGenerator{newRandLenStrGener(0, 20), newRangeInt64Gener(168435456, 368435456), newRandLenStrGener(0, 10)},
		},
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETInt, types.ETString},
			geners:        []dataGenerator{newDefaultGener(0.2, types.ETString), newDefaultGener(0.2, types.ETInt), newDefaultGener(0.2, types.ETString)},
		},
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETInt, types.ETString},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP()},
			geners: []dataGenerator{newRandLenStrGener(0, 20), newRangeInt64Gener(168435456, 368435456), newRandLenStrGener(0, 10)},
		},
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETInt, types.ETString},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
			},
			geners: []dataGenerator{newDefaultGener(0.2, types.ETString), newDefaultGener(0.2, types.ETInt), newDefaultGener(0.2, types.ETString)},
		},
	},
	ast.CharLength: {
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString}},
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
			},
		},
	},
	ast.BitLength: {
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString}},
	},
	ast.CharFunc: {
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETInt, types.ETInt, types.ETInt, types.ETString},
			geners:        []dataGenerator{&charInt64Gener{}, &charInt64Gener{}, &charInt64Gener{}, nil},
			constants:     []*Constant{nil, nil, nil, {Value: types.NewDatum("ascii"), RetType: types.NewFieldType(mysql.TypeString)}},
		},
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETInt, types.ETInt, types.ETInt, types.ETString},
			geners:        []dataGenerator{&charInt64Gener{}, nil, &charInt64Gener{}, nil},
			constants:     []*Constant{nil, nil, nil, {Value: types.NewDatum("ascii"), RetType: types.NewFieldType(mysql.TypeString)}},
		},
	},
	ast.FindInSet: {
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString, types.ETString}},
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString, types.ETString}, geners: []dataGenerator{&constStrGener{"case"}, &constStrGener{"test,case"}}},
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString, types.ETString}, geners: []dataGenerator{&constStrGener{""}, &constStrGener{"test,case"}}},
	},
}

var vecBuiltinStringCases2 = map[string][]vecExprBenchCase{
	ast.MakeSet: {
		{aesModes: "aes-128-ecb", retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETInt, types.ETString, types.ETString, types.ETString, types.ETString, types.ETString, types.ETString, types.ETString, types.ETString}},
	},
	ast.Oct: {
		{aesModes: "aes-128-ecb", retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETInt}},
		{aesModes: "aes-128-ecb", retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString}, geners: []dataGenerator{&numStrGener{*newRangeInt64Gener(-10, 10)}}},
	},
	ast.Quote: {
		{aesModes: "aes-128-ecb", retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString}},
	},
	ast.Ord: {
		{aesModes: "aes-128-ecb", retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString}},
	},
	ast.Bin: {
		{aesModes: "aes-128-ecb", retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETInt}},
	},
	ast.ToBase64: {
		{aesModes: "aes-128-ecb", retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString}, geners: []dataGenerator{newRandLenStrGener(0, 10)}},
	},
	ast.FromBase64: {
		{aesModes: "aes-128-ecb", retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString}, geners: []dataGenerator{newRandLenStrGener(10, 100)}},
	},
	ast.ExportSet: {
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETInt, types.ETString, types.ETString},
			geners:        []dataGenerator{newRangeInt64Gener(10, 100), &constStrGener{"Y"}, &constStrGener{"N"}},
		},
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETInt, types.ETString, types.ETString, types.ETString},
			geners:        []dataGenerator{newRangeInt64Gener(10, 100), &constStrGener{"Y"}, &constStrGener{"N"}, &constStrGener{","}},
		},
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETInt, types.ETString, types.ETString, types.ETString, types.ETInt},
			geners:        []dataGenerator{newRangeInt64Gener(10, 100), &constStrGener{"Y"}, &constStrGener{"N"}, &constStrGener{","}, newRangeInt64Gener(-10, 70)},
		},
	},
	ast.Repeat: {
		{aesModes: "aes-128-ecb", retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString, types.ETInt}, geners: []dataGenerator{newRandLenStrGener(10, 20), newRangeInt64Gener(-10, 10)}},
	},
	ast.Lower: {
		{aesModes: "aes-128-ecb", retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString}},
		{aesModes: "aes-128-ecb", retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString}, geners: []dataGenerator{newSelectStringGener([]string{"one week’s time TEST", "one week's time TEST", "ABC测试DEF", "ABCテストABC"})}},
	},
	ast.IsNull: {
		{aesModes: "aes-128-ecb", retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString}, geners: []dataGenerator{newRandLenStrGener(10, 20)}},
		{aesModes: "aes-128-ecb", retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString}, geners: []dataGenerator{newDefaultGener(0.2, types.ETString)}},
	},
	ast.Upper: {
		{aesModes: "aes-128-ecb", retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString}},
		{aesModes: "aes-128-ecb", retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString}, geners: []dataGenerator{newSelectStringGener([]string{"one week’s time TEST", "one week's time TEST", "abc测试DeF", "AbCテストAbC"})}},
	},
	ast.Right: {
		{aesModes: "aes-128-ecb", retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString, types.ETInt}},
		// need to add BinaryFlag for the Binary func
		{aesModes: "aes-128-ecb", retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString, types.ETInt},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
				types.NewFieldTypeBuilder().SetType(mysql.TypeLonglong).BuildP()},
			geners: []dataGenerator{
				newRandLenStrGener(10, 20),
				newRangeInt64Gener(-10, 20),
			},
		},
	},
	ast.Left: {
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString, types.ETInt}},
		// need to add BinaryFlag for the Binary func
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString, types.ETInt},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
				types.NewFieldTypeBuilder().SetType(mysql.TypeLonglong).BuildP()},
			geners: []dataGenerator{
				newRandLenStrGener(10, 20),
				newRangeInt64Gener(-10, 20),
			},
		},
	},
	ast.Space: {
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETInt}, geners: []dataGenerator{newRangeInt64Gener(-10, 2000)}},
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETInt}, geners: []dataGenerator{newRangeInt64Gener(5, 10)}},
	},
	ast.Reverse: {
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString}, geners: []dataGenerator{newRandLenStrGener(10, 20)}},
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString}, geners: []dataGenerator{newDefaultGener(0.2, types.ETString)}},
		// need to add BinaryFlag for the Binary func
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP()},
		},
	},
	ast.Instr: {
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString, types.ETString}},
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString, types.ETString}, geners: []dataGenerator{&constStrGener{"test,case"}, &constStrGener{"case"}}},
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString, types.ETString}, geners: []dataGenerator{&constStrGener{"test,case"}, &constStrGener{"testcase"}}},
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString, types.ETString},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
			},
		},
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString, types.ETString},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
			},
			geners: []dataGenerator{&constStrGener{"test,case"}, &constStrGener{"case"}},
		},
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString, types.ETString},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
			},
			geners: []dataGenerator{&constStrGener{"test,case"}, &constStrGener{""}},
		},
	},
	ast.Replace: {
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString, types.ETString, types.ETString}, geners: []dataGenerator{newRandLenStrGener(10, 20), newRandLenStrGener(0, 10), newRandLenStrGener(0, 10)}},
	},
	ast.InsertFunc: {
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString, types.ETInt, types.ETInt, types.ETString}, geners: []dataGenerator{newRandLenStrGener(10, 20), newRangeInt64Gener(-10, 20), newRangeInt64Gener(0, 100), newRandLenStrGener(0, 10)}},
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString, types.ETInt, types.ETInt, types.ETString},
			childrenFieldTypes: []*types.FieldType{
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
				types.NewFieldTypeBuilder().SetType(mysql.TypeLonglong).BuildP(),
				types.NewFieldTypeBuilder().SetType(mysql.TypeLonglong).BuildP(),
				types.NewFieldTypeBuilder().SetType(mysql.TypeString).SetFlag(mysql.BinaryFlag).SetCharset(charset.CharsetBin).SetCollate(charset.CollationBin).BuildP(),
			},
		},
	},
	ast.Elt: {
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETInt, types.ETString, types.ETString, types.ETString}, geners: []dataGenerator{newRangeInt64Gener(-1, 5)}},
	},
	ast.FromUnixTime: {
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETDecimal, types.ETString},
			geners: []dataGenerator{
				gener{*newDefaultGener(0.9, types.ETDecimal)},
				&constStrGener{"%y-%m-%d"},
			},
		},
	},
	ast.Strcmp: {
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString, types.ETString}},
		{retEvalType: types.ETInt, childrenTypes: []types.EvalType{types.ETString, types.ETString}, geners: []dataGenerator{
			newSelectStringGener(
				[]string{
					"test",
				},
			),
			newSelectStringGener(
				[]string{
					"test",
				},
			),
		}},
	},
	ast.Format: {
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETDecimal, types.ETInt}, geners: []dataGenerator{
			newRangeDecimalGener(-10000, 10000, 0),
			newRangeInt64Gener(-10, 40),
		}},
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETReal, types.ETInt}, geners: []dataGenerator{
			newRangeRealGener(-10000, 10000, 0),
			newRangeInt64Gener(-10, 40),
		}},
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETDecimal, types.ETInt}, geners: []dataGenerator{
			newRangeDecimalGener(-10000, 10000, 1),
			newRangeInt64Gener(-10, 40),
		}},
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETReal, types.ETInt}, geners: []dataGenerator{
			newRangeRealGener(-10000, 10000, 1),
			newRangeInt64Gener(-10, 40),
		}},
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETString, types.ETString}, geners: []dataGenerator{
			newRealStringGener(),
			&numStrGener{*newRangeInt64Gener(-10, 40)},
		}},
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETDecimal, types.ETInt, types.ETString}, geners: []dataGenerator{
			newRangeDecimalGener(-10000, 10000, 0.5),
			newRangeInt64Gener(-10, 40),
			newNullWrappedGener(0.1, &constStrGener{"en_US"}),
		}},
		{retEvalType: types.ETString, childrenTypes: []types.EvalType{types.ETReal, types.ETInt, types.ETString}, geners: []dataGenerator{
			newRangeRealGener(-10000, 10000, 0.5),
			newRangeInt64Gener(-10, 40),
			newNullWrappedGener(0.1, &constStrGener{"en_US"}),
		}},
	},
}

func TestVectorizedBuiltinStringEvalOneVec(t *testing.T) {
	testVectorizedEvalOneVec(t, vecBuiltinStringCases)
}

func TestVectorizedBuiltinStringFunc(t *testing.T) {
	testVectorizedBuiltinFunc(t, vecBuiltinStringCases)
}

func BenchmarkVectorizedBuiltinStringEvalOneVec(b *testing.B) {
	benchmarkVectorizedEvalOneVec(b, vecBuiltinStringCases)
}

func BenchmarkVectorizedBuiltinStringFunc(b *testing.B) {
	benchmarkVectorizedBuiltinFunc(b, vecBuiltinStringCases)
}

func TestVectorizedBuiltinStringEvalOneVec2(t *testing.T) {
	testVectorizedEvalOneVec(t, vecBuiltinStringCases2)
}

func TestVectorizedBuiltinStringFunc2(t *testing.T) {
	testVectorizedBuiltinFunc(t, vecBuiltinStringCases2)
}

func BenchmarkVectorizedBuiltinStringEvalOneVec2(b *testing.B) {
	benchmarkVectorizedEvalOneVec(b, vecBuiltinStringCases2)
}

func BenchmarkVectorizedBuiltinStringFunc2(b *testing.B) {
	benchmarkVectorizedBuiltinFunc(b, vecBuiltinStringCases2)
}
