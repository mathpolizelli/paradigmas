# Linguagem escolhida: **Python**

**Documentação:** https://docs.python.org/3/reference/grammar.html

**Notação utilizada:** PEG

## Regras em PEG

```text
if_statement <- "if" condition ":" block

assignment_statement <- identifier "=" expression

condition <- comparison

comparison <- expression comparison_operator expression

comparison_operator <- "<"

block <- NEWLINE INDENT statement_list DEDENT

statement_list <- statement

statement <- assignment_statement / procedure_statement

procedure_statement <- "print" "(" arguments ")"

arguments <- expression "," expression

expression <- integer_literal / identifier

identifier <- "a" / "b" / "print"

integer_literal <- "10" / "20"
```

## Regras equivalentes em BNF

```text
<if_statement> ::= if <condition> : <block>

<assignment_statement> ::= <identifier> = <expression>

<condition> ::= <comparison>

<comparison> ::= <expression> <comparison_operator> <expression>

<comparison_operator> ::= <

<block> ::= NEWLINE INDENT <statement_list> DEDENT

<statement_list> ::= <statement>

<statement> ::= <assignment_statement> | <procedure_statement>

<procedure_statement> ::= print ( <arguments> )

<arguments> ::= <expression> , <expression>

<expression> ::= <integer_literal> | <identifier>

<identifier> ::= a | b | print

<integer_literal> ::= 10 | 20
```

## Código analisado

```python
a = 10
b = 20

if a < 20:
    print(a, b)
```

## Derivação

```text
⇒ statements

⇒ assignment
   assignment
   if_stmt

⇒ NAME '=' expression
   NAME '=' expression
   'if' named_expression ':' block

⇒ NAME '=' expression
   NAME '=' expression
   'if' expression ':' NEWLINE INDENT statements DEDENT

⇒ NAME '=' expression
   NAME '=' expression
   'if' comparison ':' NEWLINE INDENT simple_stmts DEDENT

⇒ NAME '=' expression
   NAME '=' expression
   'if' NAME '<' NUMBER ':' NEWLINE INDENT
       NAME '(' arguments ')'
   DEDENT
```