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
Has Go-lang like syntax (https://go.dev/) with some features inspired by Elixir (https://elixir-lang.org/). Like Go, it has static typing system and like Elixir, it has pipe operators identical to Elixir’s.

## Abstract Syntax: 
variablename is a variable

fn is a function

string is a textual value (“abc”) 

number is a numeric value (1, 1.0, …)

Str is string data type

Int is an integer data type

Float is a floating point number data type

type[] is an array of type

type ::= Int | Float | Bool | Str | type[] | (type*) -> type 

operator ::= + | - | * | / | ^ | % | |>

expression ::= int | str | variablename  | expression operator expression

| expression[expression]  `array at index`

| len(expression) `length of expression`

| type[expression*] `declaration of an array of type // array with no bounds or initialized with values //calling out of bounds is UB`

| funcname(expression*) `calls a function`

| expression(expression*) `higher order function call`

| (param*) -> expression `initializing higher order function`

var-declaration ::= (const | var) type variablename = expression `must declare type when creating variable`

statement ::= var-declaration `variable declaration`

| expression[expression] = expression 

| variablename = expression

| {statement*} `block`

| print(expression)  `prints something`

| if (expression) statement else statement `if statement`

| for-loop declaration ::= for ( statement ; expression; statement) statement `for loops`

| return expression `returns an expression`

param ::= type variablename `parameters are immutable`

function-definition ::= fn type funcname(param*) { statement* } `param(s) are comma-separated`

Program ::= function-definition*

## Computational Abstraction Non-Trivial Feature: 
Higher Order Functions `without using javascript higher order functions`

## Non-Trivial Feature #2: 
 Mutable/Immutable variables

## Non-Trivial Feature #3: 
Pipe Operators
https://elixirschool.com/en/lessons/basics/pipe_operator 

## Work Planned for Custom Component:  
Higher Order Functions `without using javascript higher order functions`
