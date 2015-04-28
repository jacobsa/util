// Copyright 2015 Aaron Jacobs. All Rights Reserved.
// Author: aaronjjacobs@gmail.com (Aaron Jacobs)
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

package lrucache

import (
	"container/list"
	"errors"
)

// An LRU cache for arbitrary values indexed by string keys. External
// synchronization is required. Gob encoding/decoding is supported as long as
// all values are registered using gob.Register.
//
// May be used directly as a field in a larger struct. Must be created with New
// or initialized using gob decoding.
type Cache struct {
	/////////////////////////
	// Constant data
	/////////////////////////

	capacity int

	/////////////////////////
	// Mutable state
	/////////////////////////

	// List of elements, with least recently used at the tail.
	//
	// INVARIANT: elems.Len() <= capacity
	// INVARIANT: Each element is of type elem
	elems list.List

	// Index of elements by name.
	//
	// INVARIANT: Contains all and only the elements of elem
	// INVARIANT: For each k, v: v.Value.(elem).Key == k
	index map[string]*list.Element
}

type elem struct {
	Key   string
	Value interface{}
}

// Initialize a cache with the supplied capacity.
func New(capacity int) (c Cache) {
	c.capacity = capacity
	c.index = make(map[string]*list.Element)
	return
}

// Panic if any internal invariants have been violated. The careful user can
// arrange to call this at crucial moments.
func (c *Cache) CheckInvariants() {
	panic("TODO")
}

////////////////////////////////////////////////////////////////////////
// Cache interface
////////////////////////////////////////////////////////////////////////

// Insert the supplied value into the cache, overwriting any previous entry for
// the given key. The value must be non-nil. Return the overwritten value, or
// nil if none.
func (c *Cache) Insert(
	key string,
	value interface{}) (prev interface{}) {
	panic("TODO")
}

// Erase any entry for the supplied key.
func (c *Cache) Erase(key string) {
	panic("TODO")
}

// Look up a previously-inserted value for the given key. Return nil if no
// value is present.
func (c *Cache) LookUp(key string) interface{} {
	panic("TODO")
}

////////////////////////////////////////////////////////////////////////
// Gob encoding
////////////////////////////////////////////////////////////////////////

func (c *Cache) GobEncode() (b []byte, err error) {
	err = errors.New("TODO")
	return
}

func (c *Cache) GobDecode(b []byte) (err error) {
	err = errors.New("TODO")
	return
}
