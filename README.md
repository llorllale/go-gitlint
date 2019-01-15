# go-gitlint
Go lint your commit messages!

## Requirements

As an architect of other (not necessarily golang) projects hosted on GitHub I need:

* Commit titles and bodies merged to the development branch conform to an arbitrary regex
* Commit titles to include the relevant issue's ID
* Lint commit msgs on varios development platforms (Windows, Linux, Mac)
* (BONUS) Pre-commit hook to validate my commit's msg
* (BONUS) Performance (because a slow pre-commit hook would render the git workflow unmanageable)
