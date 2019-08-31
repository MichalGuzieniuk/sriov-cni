package objx

import (
<<<<<<< HEAD
	"reflect"
=======
	"fmt"
>>>>>>> update vendor with go mod vendor
	"regexp"
	"strconv"
	"strings"
)

<<<<<<< HEAD
const (
	// PathSeparator is the character used to separate the elements
	// of the keypath.
	//
	// For example, `location.address.city`
	PathSeparator string = "."

	// arrayAccesRegexString is the regex used to extract the array number
	// from the access path
	arrayAccesRegexString = `^(.+)\[([0-9]+)\]$`

	// mapAccessRegexString is the regex used to extract the map key
	// from the access path
	mapAccessRegexString = `^([^\[]*)\[([^\]]+)\](.*)$`
)
=======
// arrayAccesRegexString is the regex used to extract the array number
// from the access path
const arrayAccesRegexString = `^(.+)\[([0-9]+)\]$`
>>>>>>> update vendor with go mod vendor

// arrayAccesRegex is the compiled arrayAccesRegexString
var arrayAccesRegex = regexp.MustCompile(arrayAccesRegexString)

// mapAccessRegex is the compiled mapAccessRegexString
var mapAccessRegex = regexp.MustCompile(mapAccessRegexString)

// Get gets the value using the specified selector and
// returns it inside a new Obj object.
//
// If it cannot find the value, Get will return a nil
// value inside an instance of Obj.
//
// Get can only operate directly on map[string]interface{} and []interface.
//
// Example
//
// To access the title of the third chapter of the second book, do:
//
//    o.Get("books[1].chapters[2].title")
func (m Map) Get(selector string) *Value {
	rawObj := access(m, selector, nil, false, false)
	return &Value{data: rawObj}
}

// Set sets the value using the specified selector and
// returns the object on which Set was called.
//
// Set can only operate directly on map[string]interface{} and []interface
//
// Example
//
// To set the title of the third chapter of the second book, do:
//
//    o.Set("books[1].chapters[2].title","Time to Go")
func (m Map) Set(selector string, value interface{}) Map {
	access(m, selector, value, true, false)
	return m
}

<<<<<<< HEAD
// getIndex returns the index, which is hold in s by two braches.
// It also returns s withour the index part, e.g. name[1] will return (1, name).
// If no index is found, -1 is returned
func getIndex(s string) (int, string) {
	arrayMatches := arrayAccesRegex.FindStringSubmatch(s)
	if len(arrayMatches) > 0 {
		// Get the key into the map
		selector := arrayMatches[1]
		// Get the index into the array at the key
		// We know this cannt fail because arrayMatches[2] is an int for sure
		index, _ := strconv.Atoi(arrayMatches[2])
		return index, selector
	}
	return -1, s
}

// getKey returns the key which is held in s by two brackets.
// It also returns the next selector.
func getKey(s string) (string, string) {
	selSegs := strings.SplitN(s, PathSeparator, 2)
	thisSel := selSegs[0]
	nextSel := ""

	if len(selSegs) > 1 {
		nextSel = selSegs[1]
	}

	mapMatches := mapAccessRegex.FindStringSubmatch(s)
	if len(mapMatches) > 0 {
		if _, err := strconv.Atoi(mapMatches[2]); err != nil {
			thisSel = mapMatches[1]
			nextSel = "[" + mapMatches[2] + "]" + mapMatches[3]

			if thisSel == "" {
				thisSel = mapMatches[2]
				nextSel = mapMatches[3]
			}

			if nextSel == "" {
				selSegs = []string{"", ""}
			} else if nextSel[0] == '.' {
				nextSel = nextSel[1:]
			}
		}
	}

	return thisSel, nextSel
}

// access accesses the object using the selector and performs the
// appropriate action.
func access(current interface{}, selector string, value interface{}, isSet bool) interface{} {
	thisSel, nextSel := getKey(selector)

	index := -1
	if strings.Contains(thisSel, "[") {
		index, thisSel = getIndex(thisSel)
	}

	if curMap, ok := current.(Map); ok {
		current = map[string]interface{}(curMap)
	}
	// get the object in question
	switch current.(type) {
	case map[string]interface{}:
		curMSI := current.(map[string]interface{})
		if nextSel == "" && isSet {
			curMSI[thisSel] = value
			return nil
=======
// access accesses the object using the selector and performs the
// appropriate action.
func access(current, selector, value interface{}, isSet, panics bool) interface{} {

	switch selector.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:

		if array, ok := current.([]interface{}); ok {
			index := intFromInterface(selector)

			if index >= len(array) {
				if panics {
					panic(fmt.Sprintf("objx: Index %d is out of range. Slice only contains %d items.", index, len(array)))
				}
				return nil
			}

			return array[index]
		}

		return nil

	case string:

		selStr := selector.(string)
		selSegs := strings.SplitN(selStr, PathSeparator, 2)
		thisSel := selSegs[0]
		index := -1
		var err error

		if strings.Contains(thisSel, "[") {
			arrayMatches := arrayAccesRegex.FindStringSubmatch(thisSel)

			if len(arrayMatches) > 0 {
				// Get the key into the map
				thisSel = arrayMatches[1]

				// Get the index into the array at the key
				index, err = strconv.Atoi(arrayMatches[2])

				if err != nil {
					// This should never happen. If it does, something has gone
					// seriously wrong. Panic.
					panic("objx: Array index is not an integer.  Must use array[int].")
				}
			}
>>>>>>> update vendor with go mod vendor
		}

		if curMap, ok := current.(Map); ok {
			current = map[string]interface{}(curMap)
		}

<<<<<<< HEAD
		current = curMSI[thisSel]
	default:
		current = nil
	}

	// do we need to access the item of an array?
	if index > -1 {
		if array, ok := interSlice(current); ok {
			if index < len(array) {
				current = array[index]
			} else {
				current = nil
=======
		// get the object in question
		switch current.(type) {
		case map[string]interface{}:
			curMSI := current.(map[string]interface{})
			if len(selSegs) <= 1 && isSet {
				curMSI[thisSel] = value
				return nil
>>>>>>> update vendor with go mod vendor
			}
			current = curMSI[thisSel]
		default:
			current = nil
		}
<<<<<<< HEAD
	}
	if nextSel != "" {
		current = access(current, nextSel, value, isSet)
=======

		if current == nil && panics {
			panic(fmt.Sprintf("objx: '%v' invalid on object.", selector))
		}

		// do we need to access the item of an array?
		if index > -1 {
			if array, ok := current.([]interface{}); ok {
				if index < len(array) {
					current = array[index]
				} else {
					if panics {
						panic(fmt.Sprintf("objx: Index %d is out of range. Slice only contains %d items.", index, len(array)))
					}
					current = nil
				}
			}
		}

		if len(selSegs) > 1 {
			current = access(current, selSegs[1], value, isSet, panics)
		}

>>>>>>> update vendor with go mod vendor
	}
	return current
}

<<<<<<< HEAD
func interSlice(slice interface{}) ([]interface{}, bool) {
	if array, ok := slice.([]interface{}); ok {
		return array, ok
	}

	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil, false
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret, true
=======
// intFromInterface converts an interface object to the largest
// representation of an unsigned integer using a type switch and
// assertions
func intFromInterface(selector interface{}) int {
	var value int
	switch selector.(type) {
	case int:
		value = selector.(int)
	case int8:
		value = int(selector.(int8))
	case int16:
		value = int(selector.(int16))
	case int32:
		value = int(selector.(int32))
	case int64:
		value = int(selector.(int64))
	case uint:
		value = int(selector.(uint))
	case uint8:
		value = int(selector.(uint8))
	case uint16:
		value = int(selector.(uint16))
	case uint32:
		value = int(selector.(uint32))
	case uint64:
		value = int(selector.(uint64))
	default:
		panic("objx: array access argument is not an integer type (this should never happen)")
	}
	return value
>>>>>>> update vendor with go mod vendor
}
