# Wasmo

Wasmo is a set of tools for Golang-Wasm projects

## LinkVar
```
myGoVar := &MyStruct{}
wasmo.LinkVar(myVarJS, "myGoVar", myGoVar)
...
myGoVar := wasmo.GetLinkedVar(myVarJS, "myGoVar").(*MyStruct)
```