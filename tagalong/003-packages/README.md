# Packages

Every Go program is made up of packages.

Programs start running in package `main`.

This program is using the packages with import paths "fmt" and "math/rand".

By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package `rand`.

Note: The environment in which these programs are executed is deterministic, so each time you run the example program rand.Intn will return the same number.

(To see a different number, seed the number generator; see rand.Seed. Time is constant in the playground, so you will need to use something else as the seed.)

## Import statement
This code groups the imports into a parenthesized, "factored" import statement.

## Exported names
In Go, a name is exported if it begins with a capital letter. For example, `Pizza` is an exported name, as is `Pi`, which is exported from the `math` package.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.