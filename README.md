# dr-gc-total-size

[The garbage collection of Docker Registry](https://docs.docker.com/registry/garbage-collection/) accepts a `--dry-run` parameter.

This result is below.

```
hello-world
hello-world: marking manifest sha256:fea8895f450959fa676bcc1df0611ea93823a735a01205fd8622846041d0c7cf
hello-world: marking blob sha256:03f4658f8b782e12230c1783426bd3bacce651ce582a4ffb6fbbfa2079428ecb
hello-world: marking blob sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4
hello-world: marking configuration sha256:690ed74de00f99a7d00a98a5ad855ac4febd66412be132438f9b8dbd300a937d
ubuntu

4 blobs marked, 5 blobs eligible for deletion
blob eligible for deletion: sha256:28e09fddaacbfc8a13f82871d9d66141a6ed9ca526cb9ed295ef545ab4559b81
blob eligible for deletion: sha256:7e15ce58ccb2181a8fced7709e9893206f0937cc9543bc0c8178ea1cf4d7e7b5
blob eligible for deletion: sha256:87192bdbe00f8f2a62527f36bb4c7c7f4eaf9307e4b87e8334fb6abec1765bcb
blob eligible for deletion: sha256:b549a9959a664038fc35c155a95742cf12297672ca0ae35735ec027d55bf4e97
blob eligible for deletion: sha256:f251d679a7c61455f06d793e43c06786d7766c88b8c24edf242b2c08e3c3f599
```

This result is useful for keeping track of blobs that are deleted.
But do you want to know how much size can be reduced anyway?
