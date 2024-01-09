# Testing

Automated tests utilize [go-cmdtest]. Testing is based on the [Vale styles][vale-packages].

## Running Tests

### Requirements

- [Go](https://go.dev/doc/install)
- [Vale](https://vale.sh/)
- Python
- [docutils](https://pypi.org/project/docutils/)

```shell
make test
```

To run a specific test
```
go test -v -run Test/testdata/AltTag/test .
```

To render the test output in html, call rst2html
```
/usr/local/bin/rst2html.py ./testdata/FutureTense/test.rst
```

## Adding Tests

1. Add a folder named the rule you're going to test in the `testdata` folder.
2. Each folder contains a `test.rst` file, which is the test data, and a `test.ct` file, which contains the expected output of the test.
3. Refer to the existing [sample test](./testdata/FutureTense/test.ct) and the [go-cmdtest] README for more details.

[go-cmdtest]: https://github.com/google/go-cmdtest
[vale-packages]: https://github.com/errata-ai/packages/blob/master/pkg/main_test.go
