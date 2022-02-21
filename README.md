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

operator ::= + | - | * | / | ^ | % | |> <span style="color:red">”|>” is defined below as expression |> expression</span>

var-declaration ::= (const | var) (type | “”) variablename = expression<span style="color:red"> type | “” means can define variable without telling its type, for example var x = 1 or var x int = 1 should mean same thing</span>

array-declaration ::= (var | const) arrayname = type[ | expression*]  <span style="color:red">expressions are comma-separated</span>


expression ::= int | str | variablename  | expression `|>` expression

<div style="margin-left:92px">
| expression[expression]  <span style="color:red">array at index</span>

| len(expression)  <span style="color:red">length of expression</span>

| type[ | expression*]  <span style="color:red">declaration of an array of type</span>

| print(expression) <span style="color:red"> prints something</span>
</div>

statement ::= var-declaration  variable declaration

<div style="margin-left:92px">
| expression[expression] = expression 

| {statement*} <span style="color:red">block<span>

| function-call <span style="color:red">calls a function</span>

| if (expression) statement else statement <span style="color:red">if statement</span>

| for-loop declaration ::= for ( expression ; expression; expression ) {statement*} <span style="color:red">for loops</span>

| return expression <span style="color:red">returns an expression</span>
</div>

function-definition ::= fn funcname(paramname type){ statement* } <span style="color:red">var-declarations are comma-separated</span>

function-call ::= funcname( | param*)  <span style="color:red">params are comma seperated</span>

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
