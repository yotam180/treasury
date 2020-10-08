# AltFS

AltFS is a simple file system wrapper that knows to take files from different mount points, with priority.

Let's observe how it works:

Assume you have to following folder structure mounted on /mnt1:
    /mnt1
        /a
            a.txt
        /b
            b.txt

And the following folder structure on /mnt2:
    /mnt2
        /a
            a.info
            a.txt
        /c
            c.info

AltFS can be defined with read mount-points and write mount-points, which will be prioritized. Assume an AltFS is initialized with the next parameters:

ReadFS: /mnt1 /mnt2
WriteFS: /mnt2

Then, reading a file will first try to find it inside /mnt1, then in /mnt2. But writing a file will write it directly to /mnt2.

For example, `Read("/a/a.txt")` is going to return /mnt1/a/a.txt, `Read("/a/a.info")` will return /mnt2/a/a.info, and `Write("/a/a.txt")` is going to write into /mnt2/a.txt

> TODO: The Write example might not be the desirable outcome. We might want to prevent "overriding" files in such way that they are written to the only write-FS but the changes are not visible since the read-FS is the one that will determine the file data.
> We might want to allow ordering of files according to date?
