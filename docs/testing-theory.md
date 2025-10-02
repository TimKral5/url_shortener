# Testing Theory

## Test Types

**Functional Tests**

Describes tests that evaluate functions within a solution that are
usually defined during planning.

**Non-functional Tests**

These are tests that evaluate things such as performance, compliance
to clean-code principles and such.

**Abstract Test Cases**

Describes test cases that evaluate functions that aren't limited to
single outcomes but instead allow a number of results.

**Concrete Test Cases**

In contrast to abstract tests, concrete test cases set specific
conditions for input and output and mark it as failed if the output
doesn't fit the specification.

## Test Methods

**White-box Tests**

These types of tests require the source code or at the very least,
individual "units" to be available. These units or entire sets of
units can then be tested individually through unit tests, for
example.

**Black-box Tests**

In black-box test cases, the source of the solution is not known.
Instead, the solution is tested as a whole, possibly even with
connections to external services.

**Automatic Testing**

Testing can be done manually, though, this is not very efficient,
especially given the possibility that a change may affect several
components/units of a solution at once. This is why there are
frameworks for unit testing that allows the setup of tests that can
be executed on command.

Testing can also be automated through CI/CD pipelines.

## Test Levels

### Unit Tests / Component Tests

Unit tests evaluate the functions of individual units. They are, by
design, white-box tests.

Component tests can be, depending on the definition, a different kind
of tests where multiple units are tested together.

### Integration Tests

Integration tests test the interactions between a solution and
external services. They can be white-box but also black-box tests.

### System Tests

System tests don't only test individual aspects of a solution, but
instead do testing on the solution and its environment as a whole.

These ultimately are black-box testing by nature.

### Acceptance Tests

These tests are usually done together with the customer and intended
to confirm that the requirements set by the customer are met and that
the customers are satisfied.

