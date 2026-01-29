package main

import (
	"fmt"
	"reflect"
)

// Person defines a basic data structure. 
// Note: 'age' is unexported (lowercase), making it invisible to certain reflection operations.
type Person struct {
	Name string
	age  int
}

type Greeter struct{}

func (g Greeter) Greet(name string) string {
	return "Hello " + name
}

// Future Reference: Method Set Marshalling
// Reflection allows a program to inspect its own structure. 
// When using reflect.Type, you are looking at the "blueprint," 
// while reflect.Value looks at the "actual data" in memory.
func main() {
	// --- Section 1: Dynamic Method Invocation ---
	g := Greeter{}
	t := reflect.TypeOf(g)
	v := reflect.ValueOf(g)

	fmt.Printf("Inspecting Type: %s\n", t)

	// Logic Block: Iterating over methods of a struct
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("Method %d: %s\n", i, method.Name)
	}

	// Logic Block: Dynamically calling a method by name
	// We use reflect.Value to call methods because it holds the function pointer.
	m := v.MethodByName("Greet")
	if m.IsValid() {
		args := []reflect.Value{reflect.ValueOf("Alice")}
		results := m.Call(args)
		fmt.Println("Greet Result:", results[0].String())
	}

	// --- Section 2: Struct Field Inspection & Mutation ---
	p := Person{Name: "Everest", age: 30}
	vPerson := reflect.ValueOf(p)

	// Logic Block: Accessing field values
	for i := 0; i < vPerson.NumField(); i++ {
		field := vPerson.Field(i)
		fmt.Printf("Field %d (%s): %v\n", i, vPerson.Type().Field(i).Name, field)
	}

	// Future Reference: Settability and Pointers
	// To modify a value via reflection, you must pass a pointer (&p).
	// .Elem() "dereferences" the pointer to get the underlying settable value.
	vSettable := reflect.ValueOf(&p).Elem()
	nameField := vSettable.FieldByName("Name")

	// Logic Block: Safe mutation of fields
	if nameField.CanSet() {
		nameField.SetString("Hello Paisa")
	} else {
		fmt.Println("Error: Field 'Name' is not settable")
	}
	fmt.Println("Modified Person:", p)

	// --- Section 3: Deep Type Inspection (Kinds vs. Types) ---
	x := 42
	vInt := reflect.ValueOf(x)
	
	// Logic Block: Distinguishing between Type (int) and Kind (reflect.Int)
	fmt.Printf("\nValue: %v | Type: %s | Kind: %s\n", vInt, vInt.Type(), vInt.Kind())
	fmt.Println("Is it an Int?", vInt.Kind() == reflect.Int)
	fmt.Println("Is it zero value?", vInt.IsZero())

	// --- Section 4: Pointer Dereferencing & Interfaces ---
	y := 10
	vPtr := reflect.ValueOf(&y)      // Type: *int
	vData := reflect.ValueOf(&y).Elem() // Type: int (following the pointer)

	fmt.Printf("\nPointer Type: %s | Elem Type: %s\n", vPtr.Type(), vData.Type())

	// Logic Block: Direct integer manipulation
	vData.SetInt(80)
	fmt.Println("Modified y via reflection:", y)

	// Logic Block: Handling Empty Interfaces
	var itf interface{} = "Reflective Go"
	vItf := reflect.ValueOf(itf)
	if vItf.Kind() == reflect.String {
		fmt.Println("Interface contains string:", vItf.String())
	}
}
//