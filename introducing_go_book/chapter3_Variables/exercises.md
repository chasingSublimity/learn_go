1. What are two ways to create a new variable?
  - With or without type inference
    - With: x := 'foo'
    - Without: var x string = 'foo'

2. What is the value of x after running x := 5; x += 1?
  - 6

3. What is scope? How do you determine the scope of a variable in Go?
  - Scope refers to where a given variable can be accessed or referenced within a program.
  - Variables in Scope are *block scoped*, meaning that they are accessible within the curly braces in which they were defined.

4. What is the difference between var and const?
  - Const === constant. Constant variables cannot be redefined.

5. Using the example program as a starting point, write a program that converts from Fahrenheit into Celsius (C = (F − 32) * 5/9).
```
package main

import "fmt"

func main() {
	fahrenheitTemp := 50

	fmt.Println((fahrenheitTemp - 32) * 5 / 9)

}
```

6. Write another program that converts from feet into meters (1 ft = 0.3048 m).
```
package main

import "fmt"

func main() {
	feetMeasurement := 10

	fmt.Println(feetMeasurement * .3048)

}
```
