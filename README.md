# garden-containerd (spike)

Attempt at a [containerd](https://containerd.io/) backend for [Garden](https://github.com/cloudfoundry/garden).

**Disclaimer**: This is a spike, built during a Cloud Foundry hack day. It is
as far from production ready as code can possibly get. There are no tests,
abstractions or encapsulations. The code is riddled with race conditions.  More
importantly, it doesn't actually work (at least not in bosh). The
garden-containerd code in `src/garden-containerd` and the version of containerd
that is submoduled work together (to a point) in a VM based on Ubuntu Zesty
(17.04), with Linux kernel version 4.12.  Using this setup, you can create
containers and run exactly one process in them using
[gaol](https://github.com/contraband/gaol).

For example:

```
gaol create -r 'docker.io/library/busybox:latest' -n test
gaol run test -a -c '/bin/echo hi'
```

The bosh blobs are checked in here because it's a hack day and life's too short
to get an S3 bucket for that. The version of runc in there is built from
51b501dab1889ca609db9c536ac976f0f53e7021, the latest commit as of today
(7/7/17).
