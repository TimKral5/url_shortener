# Test Concept

> **Note:** Terminology and concepts used are defined
[here](./testing-theory.md).

This document is based on the **IEEE 829** standard. Read more about
it [here](https://en.wikipedia.org/wiki/Test_plan).

## Introduction

The *URL Shortener* project is a solution made to shorten URLs
through an API. The generated URLs then point to the API and upon
access, redirect the user to the desired destination.

## Test Items

1. API and Endpoints
2. Authentication Interfaces
3. Cache Interfaces
4. Database Interfaces
5. Client Libraries

## Features to be tested

- [ ] 1. API Endpoints:
    - [ ] 1. Fetching Documentation (JSON/YAML/HTML1/HTML2)
    - [ ] 2. Shortening URL (JSON)
    - [ ] 3. Fetching URL (JSON)
    - [ ] 4. Performing Redirect (HTTP status)
- [ ] 2. Authentication Interfaces
    - [ ] 1. Auth-less implementation
    - [ ] 2. Single-token implementation
    - [ ] 3. LDAP implementation (JWT/Cookies)
- [ ] 3. Cache Interfaces
    - [ ] 1. Memcached implementation
    - [ ] 2. Cache-less implementation
- [ ] 4. Database Interfaces
    - [ ] 1. MongoDB implementation
    - [ ] 2. MariaDB implementation

## Features not to be tested

The following list of systems and components will not be tested
individually, due to them being thoroughly integrated into other
tests or them being validated through visual indicators.

- Hashing Utilities
- Logging Utilities
- Client Libraries (generated)

## Approach

**Unit tests**, **component tests** and **integration tests** are
done through the internal system of **Go**.

Additionally, benchmarks and fuzzy-tests can also be done through
that system.

For now, there will be no system tests, though, they could be done
manually.

## Item pass/fail criteria

- [ ] 1.1. The documentation should be accessible and show the right
  content.
- [ ] 1.2. The endpoint should be accessible and the response should
  contain the hash generated to access the full URL.
- [ ] 1.3. The endpoint should be accessible and the response should
  contain the full URL or the hash.
- [ ] 1.4. The endpoint should be accessible and the response should
  perform a redirect through a HTTP 3xx redirect.

- [ ] 2.1. The endpoints should be accessible without restrictions.
- [ ] 2.2. The endpoints should only be accessible through providing
  the configured token.
- [ ] 2.3. The endpoints should be accessible through JWT or cookies
  with the user credentials and permissions that were configured.

- [ ] 3.1.
- [ ] 3.2.

- [ ] 4.1.
- [ ] 4.2.

## Test Deliverables

Testing is fundamentally done using **Go**'s `go test`. However,
system could, in the future, be done through shell scripts.

## Testing Tasks

Following types of tests will be applied to this project:

1. Unit/Component Tests
2. Integration Tests
3. System Tests (eventually)

## Environmental Needs

In order to run any tests at all, the **Go** cli needs to be present.
Also, in order to run the integration tests, **Docker** needs to be
installed as well.

## Schedule

> TODO

