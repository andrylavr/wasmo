# Wasmo

Wasmo is a set of tools / functions / utils for Golang-Wasm projects

## LinkVar
```
myGoVar := &MyStruct{}
wasmo.LinkVar(myVarJS, "myGoVar", myGoVar)
...
myGoVar := wasmo.GetLinkedVar(myVarJS, "myGoVar").(*MyStruct)
```