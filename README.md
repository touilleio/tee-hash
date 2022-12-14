Tee hash
====

`tee-hash` computes the sha256 hash on the fly of a flow stdin -> stdout and write the resulting hash to a local file for later use.

Typical use case is for tar'ing a big folder and sending it to S3 without the need to store it locally first and reading it twice:

```
tar cf - $folder | tee-hash -output hash.sha256 | aws s3 cp - s3://bucket/$folder.tar
```

`hash.sha256` container the sha256 of the uploaded files.

```
aws s3 cp hash.sha256 s3://bucket/$folder.tar.sha256sum
```

This allows to compute the hash of the file without reading the file twice.

# Usage

```
go install github.com/touilleio/tee-hash@latest
echo 1234 | tee-hash
cat hash.sha256
```

# Without `tee-hash`

If you would have done it without `tee-hash`, you would probably have done it the following way:

```
tar cf $folder.tar $folder <- !! requires almost the same amount of local storage
aws s3 cp $folder.tar s3://bucket/$folder.tar <- reads $folder.tar
sha256sum $folder.tar | awk '{print $a}' > $folder.tar.sha256sum <- !! read $folder.tar a second times
aws s3 cp $folder.tar.sha256sum s3://bucket/$folder.tar.sha256sum
```
