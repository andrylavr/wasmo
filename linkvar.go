package wasmo

import "syscall/js"

var linkTable = map[int]interface{}{}

// LinkVar is used to link an existing Go-variable to JS-variable with some name
func LinkVar(jsValue js.Value, name string, someVar interface{}) {
	found := false
	for _, v := range linkTable {
		if v == someVar {
			found = true
			break
		}
	}
	if !found {
		var index int
		for index = range linkTable {
			break
		}
		for n := range linkTable {
			if n > index {
				index = n
			}
		}
		linkTable[index] = someVar
		jsValue.Set(name, index)
	}
}

// UnlinkVar deletes an existing link
func UnlinkVar(jsValue js.Value, name string) {
	index := jsValue.Get(name).Int()
	delete(linkTable, index)
}

// GetLinkedVar returns the associated variable
func GetLinkedVar(jsValue js.Value, name string) interface{} {
	index := jsValue.Get(name).Int()
	return linkTable[index]
}
