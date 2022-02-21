# GoScript
The goal of this project is to design a language and subsequently implement a complete compiler for this language.

This language will be defined in Backus-Naur Form(BNF) which will also include tokens as well as concrete and abstract syntax.

Our language will also be Statically Typed, meaning it will reject ill-typed programs.

GoScript will have subprograms to evaluate expressions as well as subprogramming for other features that will be listed below.

It will have control structures which will allow for the code to be conditionally executed as well as the ability for entire computations to be abstracted over.

## Language Name: 
GoScript

## Compiler Implementation Language and Reasoning: 
Go-lang. Our team member Vladimir has experience with golang and the rest of our team members would like to learn it. Has an implemented structure to help with unit testing.

## Target Language: 
Javascript

## Language Description:  
has Go-lang like syntax (https://go.dev/) with some features inspired by Elixir (https://elixir-lang.org/). Like Go, it has static typing system and like Elixir, it has pipe operators identical to Elixir’s.

## Abstract Syntax: 
variablename is a variable

fn is a function

Str is a string

Float is a float

Null is null

type[] is an array of type

type ::= Int | Float | Bool | Str | Null | type[]

operator ::= + | - | * | / | ^ | % | |>  `”|>” is defined below as expression |> expression`

expression ::= int | str | variablename  | expression operator expression

| expression[expression]  `array at index`

| len(expression) `length of expression`

| type[ | expression*] `declaration of an array of type`

| print(expression) ` prints something`

var-declaration ::= (const | var) (type | “”) variablename = expression `type | “” means can define variable without telling its type, for example var x = 1 or var x int = 1 should mean same thing`

array-declaration ::= (var | const) arrayname = type[ | expression*] `expressions are comma-separated`

statement ::= var-declaration `variable declaration`

| expression[expression] = expression 

| {statement*} `block`

| function-call `calls a function`

| if (expression) statement else statement `if statement`

| for-loop declaration ::= for ( expression ; expression; expression ) {statement*} `for loops`

| return expression `returns an expression`

function-definition ::= fn funcname(paramname type){ statement* } `var-declarations are comma-separated`

function-call ::= funcname( | param*) `params are comma seperated`

Program ::= function-definition*

## Computational Abstraction Non-Trivial Feature: 
Higher Order Functions

## Non-Trivial Feature #2: 
 Mutable/Immutable variables

## Non-Trivial Feature #3: 
Pipe Operators
https://elixirschool.com/en/lessons/basics/pipe_operator 

## Work Planned for Custom Component:  
Pipe Operators
