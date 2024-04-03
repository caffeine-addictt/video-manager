# **Contributing**

When contributing to this repository, please first discuss the change you wish to make via issue,
email, or any other method with the owners of this repository before making a change.

Please note we have a [code of conduct](CODE_OF_CONDUCT.md); please follow it in all your interactions with the project.



## Pull Request Process

1. Ensure any install or build dependencies are removed before the end of the layer when doing a
   build.
2. Update the README.md with details of changes to the interface; this includes new environment variables, exposed ports, valid file locations and container parameters.
3. Increase the version numbers in any examples files and the README.md to the new version that this
   Pull Request would represent. The versioning scheme we use is [SemVer](http://semver.org/).
4. You may merge the Pull Request once you have the sign-off of two other developers, or if you
   do not have permission to do that, you may request the second reviewer to merge it for you.



## Issue Report Process

1. Go to the project's issues.
2. Select the template that better fits your issue.
3. Read the instructions carefully and write within the template guidelines.
4. Submit it and wait for support.



## Commit Message Guidelines

When committing, commit messages are prefixed with one of the following depending on the type of change made.

- `feat:` when a new feature is introduced with the changes.
- `fix:` when a bug fix has occurred.
- `chore:` for changes that do not relate to a fix or feature and do not modify _source_ or _tests_. (like updating dependencies)
- `refactor:` for refactoring code that neither fixes a bug nor adds a feature.
- `docs:` when changes are made to documentation.
- `style:` when changes that do not affect the code, but modify formatting.
- `test:` when changes to tests are made.
- `perf:` for changes that improve performance.
- `ci:` for changes that affect CI.
- `build:` for changes that affect the build system or external dependencies.
- `revert:` when reverting changes.

A parenthesis can be placed after the type of change to indicate the scope of the change. Below list some example commit messages.

```sh
git commit -m "docs(client): Updated README.md in client/"
git commit -m "revert(server): Fall back to old typing"
git commit -m "docs: Moved README.md"
```
