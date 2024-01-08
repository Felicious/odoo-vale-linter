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

To run only one test
```
go test -v -run Test/testdata/AltTag/test .
```

To render the test output in html, call rst2html
```
/usr/local/bin/rst2html.py ./testdata/FutureTense/test.rst
```

## Adding Tests

Refer to the existing [sample test](./testdata/FutureTense/test.ct) and the [go-cmdtest] README for details on writing tests.

[go-cmdtest]: https://github.com/google/go-cmdtest
[vale-packages]: https://github.com/errata-ai/packages/blob/master/pkg/main_test.go
