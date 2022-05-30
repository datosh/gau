# gau

Golang ArchUnit (gau) Test Framework


## Credit

* [Java ArchUnit](https://www.archunit.org/)
* [matthewmcnew/archtest](https://github.com/matthewmcnew/archtest)
    * not maintained, slow for bigger projects
* [fdaines/arch-go](https://github.com/fdaines/arch-go)
    * not unit test based: requires extra CLI tool

## Ideas

* Support checking of not using specific functions / classes of packages, e.g.,
"do not use fmt.Prinln", but some logging library instead.

## Open Problems

* Is Method Chaining really the best way for golang? How do we communicate errors? Fail on last call? Return nil?
