#!/bin/bash

repo_user="timkral5"
repo_id="url_shortener"
repo_url="http://github.com/timkral5/url_shortener"

go_version=1.24.5
docker_version=28.3.3

repo_branches=(master development)
repo_branch_names=(Master Development)
repo_branch_colors=(blue red)

# Disect readme file
readme_file="$(cat README.md)"
badges_start="$(echo -E "$readme_file" | grep -n '<!-- {{badges:start}} -->' | cut -d ':' -f 1)"
badges_end="$(echo -E "$readme_file" | grep -n '<!-- {{badges:end}} -->' | cut -d ':' -f 1)"

# Generate new readme file
echo "$(echo -E "$readme_file" | sed -n "1,${badges_start}p")"
echo

# Project info
echo "![License Information](https://img.shields.io/github/license/$repo_user/$repo_id?logo=github&label=License)"
echo "![Latest Release](https://img.shields.io/github/v/release/$repo_user/$repo_id?logo=github&label=Latest%20Release&color=blue&include_prereleases)"

# Branch workflow status
for (( i=0; i<${#repo_branches[@]}; i++ )); do
    branch_id="${repo_branches[$i]}"
    branch_name="${repo_branch_names[$i]}"
    branch_color="${repo_branch_colors[$i]}"
    echo "![$branch_name Branch Status](https://img.shields.io/github/check-runs/$repo_user/$repo_id/$branch_id?logo=github&label=$branch_name%20Status)"
done

# Branch last commit date
for (( i=0; i<${#repo_branches[@]}; i++ )); do
    branch_id="${repo_branches[$i]}"
    branch_name="${repo_branch_names[$i]}"
    branch_color="${repo_branch_colors[$i]}"
    echo "![$branch_name Last Commit](https://img.shields.io/github/last-commit/$repo_user/$repo_id/$branch_id?logo=git&color=$branch_color&label=Last%20Commit%20-%20$branch_name)"
done

# Version info
echo "[![Go Version](https://img.shields.io/badge/Go_Version-$go_version-deepskyblue?logo=go)](https://go.dev)"
echo "[![Docker Version](https://img.shields.io/badge/Docker_Version-$go_version-deepskyblue?logo=docker)](https://docker.com)"

# Compatibility information
echo "![Operating Systems Badge](https://img.shields.io/badge/OS-linux%20%7C%20windows-blue?style=flat&logo=Linux&logoColor=b0c0c0)"
echo "![Architectures Badge](https://img.shields.io/badge/CPU-x86%20%7C%20x86__64%20-blue?style=flat&logo=amd&logoColor=b0c0c0)"

echo
echo "$(echo -E "$readme_file" | sed -n "${badges_end},\$p")"
