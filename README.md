Tee hash
====

Compute the hash of a flow stdin->stdout and write it to a local file for later use.

```
cat some/file | tee-hash -output hash.sha256 | aws s3 cp - s3://bucket/some/file
```

`hash.sha256` container the sha256 of the uploaded files.

This allows to compute the hash of the file without reading the file twice.
