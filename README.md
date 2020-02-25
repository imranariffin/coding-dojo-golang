# Coding Dojo Golang

## Golang

If you're not familiar with Go or you'd like a good warm up the official [Tour of Go](https://tour.golang.org/welcome/1) is a good place to hangout at before starting this problem.

## Problem
### CSV of Similar Email Addresses

#### Valid email address

A valid email address is made up of a 'local name', alias ('@') & a domain
name. For example:

```
golang.dev@coding.dojo.com
```
local name:'golang.dev'
alias: '@'
domain name: 'coding.dojo.com'

A valid email address must also only contain lowercase letters.

```
Golang.dev@coding.dojo.com # Invalid
golang.dev@Coding.dojo.com # Invalid
golang.dev@coding.dojo.com # Valid
```

A valid email address may contain '+'`'s & `'.'`'s in the **local name**.

E.g. All email addresses are valid:
```
golang.dev+cool@coding.dojo.com
go.lang.dev@coding.dojo.com
go.lang.dev+cool@coding.dojo.com
```

#### Similar email addresses

There are two rules regarding similar addresses:

1. Ignore dots ('.') in local name
2. Ignore all characters after '+' in local name

e.g.

These two are similar (rule 1 - drop the dots)
```
go.lang.dev@coding.dojo.com
golangdev@coding.dojo.com
```
These two are similar (rule 2 - ignore all characters after '+')
```
golangdev+cool@coding.dojo.com
golangdev@coding.dojo.com
```
These two are similar (rule 1 & 2)
```
golangdev+cool@coding.dojo.com
golangdev@coding.dojo.com
```

#### Question

Given a CSV of emails as input, can you complete the following function to return number of rows with emails that are all similar?

You can assume each email provided is valid. An input will have a consistent number of emails in each row but different test cases may have different number of emails in a row.

```go
func getSimilarEmailsCount(emailCSV []string) int {
  // complete the function
}
```

#### Sample test case 1
```go
import "fmt"
input := []string{
  "alice@mail.com,alic.e@mail.com,a.lice@mail.com", // all similar - accepted
  "bob@mail.com,bob+dev@mail.com,b.ob+dev@mail.com", // all similar - accepted
  "bob@mail.com,b.o.b@mail.com,bob.dev@mail.com", // one not similar - rejected
}
got := getSimilarEmailsCount(input);
want := 2;
fmt.Println(got == want);
```

#### Sample test case 2
```go
input := []string{
  "alice@mail.com,alic.e@mail.com", // all similar - accepted
  "bob@mail.com,b.o.b@mail.com", // all similar - accepted
  "foo@bar.com,foo@bar.xyz", // not similar - rejected
  "john@mail.com,j.o.h.n+123@mail.com", // all similar - accepted
}
got := getSimilarEmailsCount(input);
want := 3;
fmt.Println(got == want);
```

#### Running the tests
```bash
go test # Run all tests
go test -run 'TestTwoRowOneBothSimilar' # Run 'TestTwoRowOneBothSimilar' only
go test -run 'TestTwoRowOneBothSimilar|TestTwoRowOneBothSimilar' # Run 'TestTwoRowOneBothSimilar' & 'TestTwoRowOneBothSimilar' only
```

#### Credit

This problem is an extension of the "Unique Email Addresses" posted on LeetCode [here](https://leetcode.com/problems/unique-email-addresses/).
