<h1 align="center">URL Shortener</h1>

<div align="center">

![License Information Badge](https://img.shields.io/github/license/timkral5/url_shortener?logo=github&label=License&color=blue)
![Latest Release Version Badge](https://img.shields.io/github/v/tag/timkral5/url_shortener?logo=github&label=Latest%20Version&color=blue)
[![Master Last Commit Badge](https://img.shields.io/github/last-commit/timkral5/url_shortener/master?label=Last%20Commit%20-%20Master&logo=git&color=blue)](https://github.com/TimKral5/url_shortener)
[![Development Last Commit Badge](https://img.shields.io/github/last-commit/timkral5/url_shortener/development?label=Last%20Commit%20-%20Development&logo=git&color=red)](https://github.com/TimKral5/url_shortener/tree/development)
[![Go Version](https://img.shields.io/badge/Go_Version-1.24.5-deepskyblue?logo=go)](https://go.dev)
[![Docker Version](https://img.shields.io/badge/Docker_Version-18.3.3-deepskyblue?logo=docker)](https://docker.com)
![Operating Systems Badge](https://img.shields.io/badge/OS-linux%20%7C%20windows-blue?style=flat&logo=Linux&logoColor=b0c0c0)
![Architectures Badge](https://img.shields.io/badge/CPU-x86%20%7C%20x86__64%20-blue?style=flat&logo=amd&logoColor=b0c0c0)

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
