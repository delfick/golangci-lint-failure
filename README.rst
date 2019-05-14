Golangci-lint Failure
=====================

If I run::

   $ golangci-lint run

Then I get no errors.

If I run::

   $ golangci-lint run ./thing/thing_test.go

I get::

   thing/thing_test.go:9:2: undeclared name: `now` (typecheck)
         now = func() time.Time { return time.Unix(1, 0) }
         ^
   thing/thing_test.go:10:12: undeclared name: `Value` (typecheck)
         if got := Value(); got != 1 {

Because it doesn't see that ``now`` and ``Value`` are defined in thing.go
