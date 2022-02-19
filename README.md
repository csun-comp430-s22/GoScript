# GoScript

## Language Name: 
GoScript

## Compiler Implementation Language and Reasoning: 
Go-lang. Our team member Vladimir has experience with golang and the rest of our team members would like to learn it. Has an implemented structure to help with unit testing.

## Target Language: 
Javascript

## Language Description:  
has Go-lang like syntax (https://go.dev/) with some features inspired by Swift (https://developer.apple.com/swift/) and Elixir (https://elixir-lang.org/). Like Go, it has structs and a static typing system. Language includes operator overloading similar to one in Swift and pipe operator identical to Elixir’s.

## Abstract Syntax: 
varname is name of variable
funcname is  name of function
structname is name of struct
type ::= Int | Float | Bool | Str | Null | type[] | Map[type]type | structname
operator ::= + | - | * | / | ^ | % | |> ”|>” is defined below as expression operator expression
expression ::= number | variable | expression operator expression 
| type[] | expression[expression]  array at index
| len(expression)  length
| array-declaration 
| map-declaration 
| print(expression)
statement ::= var-declaration 
        | expression[expression] = expression 
        | if (expression) stmt else stmt if statement
        | for-loop declaration ::= for ( <expression> ; <expression>; <expression> ) <statement> for loops

var-declaration ::= (const | var) (type | “”) varname = expression type | “” means can define variable without telling its type, for example var x = 1 or var x int = 1 should mean same thing
array-declaration ::= [ expression* ]
map-declaration ::= { (expression: expression)* }
struct-definition ::= struct structname {
variable-declaration
// need to add function declaration too
	}
array-declaration ::= (var | const) arrayname = [ | expression] 
       |  (var | const) []arrayname = [expression] expressions are comma-separated
function-def ::= fn funcname(varname type){ statement } var-declarations are comma-separated


## Computational Abstraction Non-Trivial Feature: 
Operator Overloading

## Non-Trivial Feature #2: 
Static Typing System(As seen in Java, C, C++) https://stackoverflow.com/a/1517670

## Non-Trivial Feature #3: 
Pipe Operators
https://elixirschool.com/en/lessons/basics/pipe_operator 

## Work Planned for Custom Component:  
Operator Overloading
