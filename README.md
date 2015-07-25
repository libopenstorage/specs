# Lib Open Storage Specifications

This project is about defining an interface to orchestrate the provisioning of storage volumes for Linux containers and creating a specification by which a volume's properties can be defined.  It also provides a UX specification defining how an end user would provision storage for stateful applications deployed in Linux containers.

We hope that this work will align with the Open Containers Project.

We also encourage anyone that has a vested interest in supporting stateful Linux containers to join this effort. 

# The Six Factors of Storage for Linux Containers

This specification defines the following six factors of provisioning storage to Linux Containers:

1. A vendor neutral software spec that can be included as part of the Container spec.  This spec defines the properties of the storage volume, such as a specific class of service and snapshot requirements.  It also allows for an initial format of the volume with the required filesystem and its properties.

2. A RESTful API for orchestrating the provisioning of a volume, such that it can be driven by the container orchestration software.

3. An interface for binding a volume into a container's namespace.

4. An interface for container granular storage operations.  These include Snapshots, Class of Service (CoS) and Clones.

5. An interface for extracting metrics and logs from the storage provider on a per-volume basis.

6. The User Experience (UX) for an end user to interact with Open Storage.

# Contributing

The specification and code is licensed under the Apache 2.0 license found in 
the `LICENSE` file of this repository.  

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
