# Testing

Automated tests utilize [go-cmdtest]. Testing is based on the [Vale styles][vale-packages].

ℹ️ **Any styles added to the `styles/Odoo` folder should have a test in the `testdata` folder.**

## GitHub Actions Continuous Integration (CI)

The CI workflow is defined in the [`.github/workflows/main.yml`](../.github/workflows/main.yml) file. It runs on every push and on every pull request, reporting any CI failures as a status check on the pull request.

The CI workflow runs the following steps:

1. Install Go
2. Install Python
3. Install docutils
4. Install Vale
5. Get Vale version
3. Run Vale with `make test`

## Running Tests Locally

Tests can be ran locally before pushing to GitHub, or when developing new tests.

### Requirements

- [Go 1.20+](https://go.dev/doc/install)
- [Vale v3.0+](https://vale.sh/)
- [Python 3+]()
- [docutils 0.16](https://pypi.org/project/docutils/)

To run all tests

```
make test
```

To run a specific test

```
go test -v -run Test/testdata/<STYLE-TEST>/test .

Example:
go test -v -run Test/testdata/WordList/test .
```

To render the test output in html, call rst2html

```
/usr/local/bin/rst2html.py ./testdata/FutureTense/test.rst
```

## Adding Tests

Each [Odoo Style](styles/Odoo) should have a test for it within the `testdata` folder. To add a new test, create a folder within the `testdata` folder with, the same name as the style, containing the following files:

- `.vale.ini`: contains the vale configuration for the test with a [declaration of the style][based-on-styles] to test with.
- `test.ct`: contains the expected output of the test with commands to run vale on the `test.rst` file.
- `test.rst`: contains the RST to run the test on, should be thorough and representative to _real_ RST.

```
styles
└── Odoo
    ├── Wordlist.yml
    └── ...
testdata
├── WordList
│   ├── .vale.ini
│   ├── test.ct
│   └── test.rst
└── ...
```

Once a test has been added, run `make test` to ensure that the test passes. A passing test will output `PASS` along with the expected output, a failing test will output `FAIL` along with the expected output and the actual output. If the test fails, ensure that the expected output matches the actual output and that the `.vale.ini` file is configured correctly.

### Test Files

#### .vale.ini contents

The `.vale.ini` file contains the vale configuration for the test with a [declaration of the style][based-on-styles] to test with. The following outlines the contents of the file:

First, declare the `StylesPath = ../../styles` to point to the styles folder.

Then, declare the `MinAlertLevel = <error-level>` to match the *error-level* of the style. This is required to ensure that the test fails if the error-level is not met.

Next, declare the filetype to lint with `[*.rst]` and the `TokenIgnores` and `BlockIgnores` to allow linting within RST directives in the `test.rst` file. These can be copied directly from the main .vale.ini file at the root of this directory.

Finally, declare the style to test with `[Odoo]` and the rules to test with `<STYLE>.<RULE> = YES`.

For example:

```ini
StylesPath = ../../styles

MinAlertLevel = error

[*.rst]

TokenIgnores = (:[^\s]+:`*`)
TokenIgnores = ((?s) *(\.\. [^\s]+::))
BlockIgnores = (?s) *(\.\. [^\s]+::)

Odoo.Spacing = YES
```

#### test.ct contents

The `test.ct` file contains the expected output of the test with commands to run vale on the `test.rst` file. The following outlines the contents of the file:

First, declare the `cdf` command to change the directory to the root of the repository.

Then, declare the `vale` command to run vale on the `test.rst` file with the following options:

- `--output=line`: output the results in a line format
- `--sort`: sort the results
- `--normalize`: normalize the results
- `--relative`: output the relative path of the file
- `--no-exit`: do not exit on error

Finally, declare the expected output of the test.

> ℹ️ Tip:</br>
> A helpful command to generate the expected output is to run the `vale` command with the options stated above on the `test.rst` file. Doing so outputs the expected output of the test in the terminal, which can be copied into the `test.ct` file.

For example:
```go
$ cdf ${ROOTDIR}
$ vale --output=line --sort --normalize --relative --no-exit test.rst
test.rst:1:15:Odoo.Spacing:Double space in 'has a' detected.
test.rst:3:17:Odoo.Spacing:Double space in 'Create and' detected.
test.rst:5:37:Odoo.Spacing:Double space in 'Configuration -->' detected.
test.rst:7:13:Odoo.Spacing:Double space in 'Abbrevaition (ABBR)' detected.
test.rst:10:12:Odoo.Spacing:Double space in 'a note' detected.
test.rst:10:25:Odoo.Spacing:Double space in 'a double' detected.
test.rst:13:15:Odoo.Spacing:Double space in 'has a' detected.
test.rst:15:17:Odoo.Spacing:Double space in 'Create and' detected.
test.rst:17:37:Odoo.Spacing:Double space in 'Configuration -->' detected.
test.rst:29:44:Odoo.Spacing:Double space in 'has a' detected.
```

Refer to the [go-cmdtest] README for more details.

#### test.rst contents

The `test.rst` file contains the RST to run the test on, should be thorough and representative to _real_ RST that the style applies to.

There does not need to be an example for every rule in the style, but there should be enough examples to ensure that the style is working as expected.

Be sure to add RST directives, like :guilabel:, :menuselection: .. note::, etc. to ensure that the style is working as expected within RST directives.

For example:
```rst
This sentence has  a double space.

This :guilabel:`Create  and Edit...` has a double space.

This :menuselection:`Accounting --> Configuration  --> Settings` has a double space.

This :abbr:`Abbrevaition  (ABBR)` has a double space.

.. note::
   This is a  note with a  double space.
```

[go-cmdtest]: https://github.com/google/go-cmdtest
[vale-packages]: https://github.com/errata-ai/packages/blob/master/pkg/main_test.go
[based-on-styles]: https://vale.sh/docs/topics/config/#basedonstyles