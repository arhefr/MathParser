# MathParser
This is golang package for parsing math expressions. With ```MathParser``` you can easily working with prefix, infix and postfix notation of math expressions.
## Install
```shell
go get -u "github.com/arhefr/MathParser"
```
## Instruction
Import the package in your project.
```Go
import . "github.com/arhefr/MathParser"
```
- ### Parser(string)[]string
  Breaks a math expression into its individual components.\
  **Example:**
  ```Go
  fmt.Println(Parser("(1488.0 + 1488.0) * 1488"))
  ```
  ```-> ["(", "1488.0", "+", "1488.0", ")", "*", "1488"]```
- ### InfixPostfix(string) []string
  Converts an infix to a postfix notation math expression.\
  **Example:**
  ```Go
  fmt.Println(InfixPostfix("(1488.0 + 1488.0) * 1488"))
  ```
  ```-> ["1488.0", "1488.0", "+", "1488", "*"]```
- ### InfixPrefix(string) []string
  Converts an infix to a prefix notation math expression.\
  **Example:**
  ```Go
  fmt.Println(InfixPrefix("(1488.0 + 1488.0) * 1488"))
  ```
  ```-> ["*", "+", "1488.0", "1488.0", "1488"]```
- ### PostfixInfix([]string) (string, error)
  Converts an postfix to a infix notation math expression. If math expression incorrect return an error.\
  **Example:**
  ```Go
  fmt.Println(PostfixInfix([]string{"1488.0", "1488.0", "+", "1488", "*"}))
  ```
  ```-> ((1488.0+1488.0)*1488) <nil>```
  ### PrefixInfix([]string) (string, error)
  Converts an prefix to a infix notation math expression. If math expression incorrect return an error.\
  **Example:**
  ```Go
  fmt.Println(PrefixInfix([]string{"*", "+", "1488.0", "1488.0", "1488"}))
  ```
  ```-> ((1488.0+1488.0)*1488) <nil>```
