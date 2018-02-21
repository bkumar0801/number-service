# Number Service Travel Audience Go Challenge

I have not seen any ambiguities in the task description.

But couple of assumptions are made :

Assumption 1: URL `/numbers?u=http://example.com/primes&u=http://foobar.com/fibo HTTP/1.0` has two valid URLs in query parameter "u": http://example.com/primes and http://foobar.com/fibo (not considered space once a syntactically valid URL found for example in  `http://foobar.com/fibo HTTP/1.0` query parameter ` HTTP/1.0` is ignored.

Assumption 2: we can even choose to pass "timeout" as a command line parameter but since 500ms is explicitely mentioned in the requirement, this value has been kept as a constant. We can always change this value in constant package.

Some of the test scenarios have been covered using Fake implementation of ResourceBuilder interface methods.

