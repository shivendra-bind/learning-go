### new data type
Go lets you create new types from existing ones.

The syntax is `type MyName OriginalType`

[Stringer](https://pkg.go.dev/fmt#Stringer)

## nil
- Pointers can be nil
- When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.
- Useful for when you want to describe a value that could be missing