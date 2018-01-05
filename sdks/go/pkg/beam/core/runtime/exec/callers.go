// File generated by specialize. Do not edit.

// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exec

import (
	"reflect"

	"github.com/apache/beam/sdks/go/pkg/beam/core/typex"
)

// TODO(herohde) 1/4/2018: Potential targets for type-specialization include simple predicate,
// dofn and combiner forms.
//
//   (1) func(X) bool
//   (2) func(X) error
//   (3) func(X, X) X
//   (4) func(ctx?, X) (X, error?)
//   (5) func(ctx?, X, func(X)) (error?)
//
// For now, we just do #3.

func init() {
	RegisterCaller(reflect.TypeOf((*func([]byte, []byte) []byte)(nil)).Elem(), callMakerByteSliceM)
	RegisterCaller(reflect.TypeOf((*func(bool, bool) bool)(nil)).Elem(), callMakerBoolM)
	RegisterCaller(reflect.TypeOf((*func(string, string) string)(nil)).Elem(), callMakerStringM)
	RegisterCaller(reflect.TypeOf((*func(int, int) int)(nil)).Elem(), callMakerIntM)
	RegisterCaller(reflect.TypeOf((*func(int8, int8) int8)(nil)).Elem(), callMakerInt8M)
	RegisterCaller(reflect.TypeOf((*func(int16, int16) int16)(nil)).Elem(), callMakerInt16M)
	RegisterCaller(reflect.TypeOf((*func(int32, int32) int32)(nil)).Elem(), callMakerInt32M)
	RegisterCaller(reflect.TypeOf((*func(int64, int64) int64)(nil)).Elem(), callMakerInt64M)
	RegisterCaller(reflect.TypeOf((*func(uint, uint) uint)(nil)).Elem(), callMakerUintM)
	RegisterCaller(reflect.TypeOf((*func(uint8, uint8) uint8)(nil)).Elem(), callMakerUint8M)
	RegisterCaller(reflect.TypeOf((*func(uint16, uint16) uint16)(nil)).Elem(), callMakerUint16M)
	RegisterCaller(reflect.TypeOf((*func(uint32, uint32) uint32)(nil)).Elem(), callMakerUint32M)
	RegisterCaller(reflect.TypeOf((*func(uint64, uint64) uint64)(nil)).Elem(), callMakerUint64M)
	RegisterCaller(reflect.TypeOf((*func(float32, float32) float32)(nil)).Elem(), callMakerFloat32M)
	RegisterCaller(reflect.TypeOf((*func(float64, float64) float64)(nil)).Elem(), callMakerFloat64M)
	RegisterCaller(reflect.TypeOf((*func(typex.T, typex.T) typex.T)(nil)).Elem(), callMakerTypex_TM)
	RegisterCaller(reflect.TypeOf((*func(typex.U, typex.U) typex.U)(nil)).Elem(), callMakerTypex_UM)
	RegisterCaller(reflect.TypeOf((*func(typex.V, typex.V) typex.V)(nil)).Elem(), callMakerTypex_VM)
	RegisterCaller(reflect.TypeOf((*func(typex.W, typex.W) typex.W)(nil)).Elem(), callMakerTypex_WM)
	RegisterCaller(reflect.TypeOf((*func(typex.X, typex.X) typex.X)(nil)).Elem(), callMakerTypex_XM)
	RegisterCaller(reflect.TypeOf((*func(typex.Y, typex.Y) typex.Y)(nil)).Elem(), callMakerTypex_YM)
	RegisterCaller(reflect.TypeOf((*func(typex.Z, typex.Z) typex.Z)(nil)).Elem(), callMakerTypex_ZM)
}

type nativeByteSliceMCaller struct {
	fn func([]byte, []byte) []byte
}

func callMakerByteSliceM(fn reflect.Value) Caller {
	f := fn.Interface().(func([]byte, []byte) []byte)
	return &nativeByteSliceMCaller{fn: f}
}

func (c *nativeByteSliceMCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().([]byte), args[1].Interface().([]byte))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeBoolMCaller struct {
	fn func(bool, bool) bool
}

func callMakerBoolM(fn reflect.Value) Caller {
	f := fn.Interface().(func(bool, bool) bool)
	return &nativeBoolMCaller{fn: f}
}

func (c *nativeBoolMCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(bool), args[1].Interface().(bool))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeStringMCaller struct {
	fn func(string, string) string
}

func callMakerStringM(fn reflect.Value) Caller {
	f := fn.Interface().(func(string, string) string)
	return &nativeStringMCaller{fn: f}
}

func (c *nativeStringMCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(string), args[1].Interface().(string))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeIntMCaller struct {
	fn func(int, int) int
}

func callMakerIntM(fn reflect.Value) Caller {
	f := fn.Interface().(func(int, int) int)
	return &nativeIntMCaller{fn: f}
}

func (c *nativeIntMCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(int), args[1].Interface().(int))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeInt8MCaller struct {
	fn func(int8, int8) int8
}

func callMakerInt8M(fn reflect.Value) Caller {
	f := fn.Interface().(func(int8, int8) int8)
	return &nativeInt8MCaller{fn: f}
}

func (c *nativeInt8MCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(int8), args[1].Interface().(int8))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeInt16MCaller struct {
	fn func(int16, int16) int16
}

func callMakerInt16M(fn reflect.Value) Caller {
	f := fn.Interface().(func(int16, int16) int16)
	return &nativeInt16MCaller{fn: f}
}

func (c *nativeInt16MCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(int16), args[1].Interface().(int16))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeInt32MCaller struct {
	fn func(int32, int32) int32
}

func callMakerInt32M(fn reflect.Value) Caller {
	f := fn.Interface().(func(int32, int32) int32)
	return &nativeInt32MCaller{fn: f}
}

func (c *nativeInt32MCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(int32), args[1].Interface().(int32))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeInt64MCaller struct {
	fn func(int64, int64) int64
}

func callMakerInt64M(fn reflect.Value) Caller {
	f := fn.Interface().(func(int64, int64) int64)
	return &nativeInt64MCaller{fn: f}
}

func (c *nativeInt64MCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(int64), args[1].Interface().(int64))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeUintMCaller struct {
	fn func(uint, uint) uint
}

func callMakerUintM(fn reflect.Value) Caller {
	f := fn.Interface().(func(uint, uint) uint)
	return &nativeUintMCaller{fn: f}
}

func (c *nativeUintMCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(uint), args[1].Interface().(uint))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeUint8MCaller struct {
	fn func(uint8, uint8) uint8
}

func callMakerUint8M(fn reflect.Value) Caller {
	f := fn.Interface().(func(uint8, uint8) uint8)
	return &nativeUint8MCaller{fn: f}
}

func (c *nativeUint8MCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(uint8), args[1].Interface().(uint8))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeUint16MCaller struct {
	fn func(uint16, uint16) uint16
}

func callMakerUint16M(fn reflect.Value) Caller {
	f := fn.Interface().(func(uint16, uint16) uint16)
	return &nativeUint16MCaller{fn: f}
}

func (c *nativeUint16MCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(uint16), args[1].Interface().(uint16))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeUint32MCaller struct {
	fn func(uint32, uint32) uint32
}

func callMakerUint32M(fn reflect.Value) Caller {
	f := fn.Interface().(func(uint32, uint32) uint32)
	return &nativeUint32MCaller{fn: f}
}

func (c *nativeUint32MCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(uint32), args[1].Interface().(uint32))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeUint64MCaller struct {
	fn func(uint64, uint64) uint64
}

func callMakerUint64M(fn reflect.Value) Caller {
	f := fn.Interface().(func(uint64, uint64) uint64)
	return &nativeUint64MCaller{fn: f}
}

func (c *nativeUint64MCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(uint64), args[1].Interface().(uint64))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeFloat32MCaller struct {
	fn func(float32, float32) float32
}

func callMakerFloat32M(fn reflect.Value) Caller {
	f := fn.Interface().(func(float32, float32) float32)
	return &nativeFloat32MCaller{fn: f}
}

func (c *nativeFloat32MCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(float32), args[1].Interface().(float32))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeFloat64MCaller struct {
	fn func(float64, float64) float64
}

func callMakerFloat64M(fn reflect.Value) Caller {
	f := fn.Interface().(func(float64, float64) float64)
	return &nativeFloat64MCaller{fn: f}
}

func (c *nativeFloat64MCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(float64), args[1].Interface().(float64))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeTypex_TMCaller struct {
	fn func(typex.T, typex.T) typex.T
}

func callMakerTypex_TM(fn reflect.Value) Caller {
	f := fn.Interface().(func(typex.T, typex.T) typex.T)
	return &nativeTypex_TMCaller{fn: f}
}

func (c *nativeTypex_TMCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(typex.T), args[1].Interface().(typex.T))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeTypex_UMCaller struct {
	fn func(typex.U, typex.U) typex.U
}

func callMakerTypex_UM(fn reflect.Value) Caller {
	f := fn.Interface().(func(typex.U, typex.U) typex.U)
	return &nativeTypex_UMCaller{fn: f}
}

func (c *nativeTypex_UMCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(typex.U), args[1].Interface().(typex.U))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeTypex_VMCaller struct {
	fn func(typex.V, typex.V) typex.V
}

func callMakerTypex_VM(fn reflect.Value) Caller {
	f := fn.Interface().(func(typex.V, typex.V) typex.V)
	return &nativeTypex_VMCaller{fn: f}
}

func (c *nativeTypex_VMCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(typex.V), args[1].Interface().(typex.V))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeTypex_WMCaller struct {
	fn func(typex.W, typex.W) typex.W
}

func callMakerTypex_WM(fn reflect.Value) Caller {
	f := fn.Interface().(func(typex.W, typex.W) typex.W)
	return &nativeTypex_WMCaller{fn: f}
}

func (c *nativeTypex_WMCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(typex.W), args[1].Interface().(typex.W))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeTypex_XMCaller struct {
	fn func(typex.X, typex.X) typex.X
}

func callMakerTypex_XM(fn reflect.Value) Caller {
	f := fn.Interface().(func(typex.X, typex.X) typex.X)
	return &nativeTypex_XMCaller{fn: f}
}

func (c *nativeTypex_XMCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(typex.X), args[1].Interface().(typex.X))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeTypex_YMCaller struct {
	fn func(typex.Y, typex.Y) typex.Y
}

func callMakerTypex_YM(fn reflect.Value) Caller {
	f := fn.Interface().(func(typex.Y, typex.Y) typex.Y)
	return &nativeTypex_YMCaller{fn: f}
}

func (c *nativeTypex_YMCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(typex.Y), args[1].Interface().(typex.Y))
	return []reflect.Value{reflect.ValueOf(out)}
}

type nativeTypex_ZMCaller struct {
	fn func(typex.Z, typex.Z) typex.Z
}

func callMakerTypex_ZM(fn reflect.Value) Caller {
	f := fn.Interface().(func(typex.Z, typex.Z) typex.Z)
	return &nativeTypex_ZMCaller{fn: f}
}

func (c *nativeTypex_ZMCaller) Call(args []reflect.Value) []reflect.Value {
	out := c.fn(args[0].Interface().(typex.Z), args[1].Interface().(typex.Z))
	return []reflect.Value{reflect.ValueOf(out)}
}
