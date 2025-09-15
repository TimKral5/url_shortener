<h1 align="center">URL Shortener</h1>

<!--
   # run this to regenerate badges.
   echo -E "$(./badge-gen.sh)" > README.md
-->

<div align="center">
<!-- {{badges:start}} -->

![License Information](https://img.shields.io/github/license/timkral5/url_shortener?logo=github&label=License)
![Latest Release](https://img.shields.io/github/v/release/timkral5/url_shortener?logo=github&label=Latest%20Release&color=blue&include_prereleases)
![Master Branch Status](https://img.shields.io/github/check-runs/timkral5/url_shortener/master?logo=github&label=Master%20Status)
![Development Branch Status](https://img.shields.io/github/check-runs/timkral5/url_shortener/development?logo=github&label=Development%20Status)
![Master Last Commit](https://img.shields.io/github/last-commit/timkral5/url_shortener/master?logo=git&color=blue&label=Last%20Commit%20-%20Master)
![Development Last Commit](https://img.shields.io/github/last-commit/timkral5/url_shortener/development?logo=git&color=red&label=Last%20Commit%20-%20Development)
[![Go Version](https://img.shields.io/badge/Go_Version-1.24.5-deepskyblue?logo=go)](https://go.dev)
[![Docker Version](https://img.shields.io/badge/Docker_Version-28.3.3-deepskyblue?logo=docker)](https://docker.com)
![Operating Systems Badge](https://img.shields.io/badge/OS-linux%20%7C%20windows-blue?style=flat&logo=Linux&logoColor=b0c0c0)
![Architectures Badge](https://img.shields.io/badge/CPU-x86%20%7C%20x86__64%20-blue?style=flat&logo=amd&logoColor=b0c0c0)

<!-- {{badges:end}} -->
</div>

<p align="center">
  <i>A simple-to-use URL shortener written in Go</i> 😎
</p>

## Table of Contents

- [About](#about)
- [How to Build](#how-to-build)
- [Documentation](#documentation)
- [License](#license)
- [Contacts](#contacts)

## About

**URL Shortener** is a project built in Go designed to be used to
shorten URLs.

## How to Build

For all available information in respects for building the project,
refer to this excerpt from the `make` command.

```plain
Available targets:

General:
   h,   help         Show this prompt.
   c,   clean        Clean up all generated files.

Deployment:
   b,  build         Build the project and its executables.
   bd, docker-build  Build the docker container.
   cu, compose-up    Launch the compose configuration for production.
   cd, compose-down  Terminate the compose configuration for production.

Documentation:
   g,   godoc        Launch documenation server.
   m,   mkdocs       Launch documenation server.

Tests:
   t,   test         Run all unit tests of the project.
   bm,  benchmarks   Run all benchmarks of the project.
   i,   integration  Run all integration tests of the project.

Code Quality:
   l,   lint         Run the linter on all project files.
   f,   format       Format all .go files in the project.

Development:
   r,   run          Launch the url_shortener executable.
   du, dev-up        Launch the compose configuration for development.
   dd, dev-down      Terminate the compose configuration for development.
   s,   stats        Show repository stats.
```

## Documentation

> **TODO**

## License

This product is distributed under the GPLv2 license. Therefore, it is
provided with the freedom of usage under very lax restrictions. In
turn, it comes without any warranty.

Refer to the [LICENSE](./LICENSE) file for more information.

## Contacts

For more information about the project, refer to the issue section
and the README of the GitHub repository.
