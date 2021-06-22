Issue: https://github.com/bazelbuild/rules_go/issues/2878

Not working version (0.27.0):
```
$ bazel run :caller
caller.cc: In function 'int main()':
caller.cc:4:14: error: 'SomeFunc' was not declared in this scope
     SomeFunc();
              ^

$ ls -la bazel-out/k8-fastbuild-*/bin/somefunc.h
bazel-out/k8-fastbuild-ST-b8cfed987b33/bin/somefunc.h -> (...)/execroot/rules_go_empty_header/bazel-out/k8-fastbuild-ST-b8cfed987b33/bin/external/org_golang_x_sys/unix/unix_/_cgo_install.h

```

Patched working version (uncomment working version in WORKSPACE):
```
$ bazel run :caller
$ ls -la bazel-out/k8-fastbuild-*/bin/somefunc.h

bazel-out/k8-fastbuild-ST-b8cfed987b33/bin/somefunc.h -> (...)/execroot/rules_go_empty_header/bazel-out/k8-fastbuild-ST-b8cfed987b33/bin/somefunc_/_cgo_install.h
```

