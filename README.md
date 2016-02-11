[![Slack Room][slack-badge]][slack-link]

# [Fisherman][fisherman] Index

The Fisherman Index is a plain text flat database externally managed and independent from Fisherman.

The index lists records, each consisting of the following fields:

* **name**, **url**, **info**, **author** and one or more **tags**.

Fields are separated by a new line **\n**. Tags are separated by one space. For example:

```
fishtape
https://github.com/fishery/fishtape
TAP producer and test harness
test tool harness tap unit-testing
bucaran
```

The easiest way to submit new plugins for registration is using the [`submit`][fisher-submit] plugin:

```
fisher install submit
```

For usage see `fisher help submit` after the installation is complete.

You can also submit a new plugin manually and create a pull request.

```fish
git clone https://github.com/fisherman/fisher-index
cd index
echo "$name\n$url\n$info\n$author\n$tags\n\n" >> index
git push origin master
```

## Disclaimer

All the plugins listed in this index are property of their respective owners. Follow the provided URL and see the bundled LICENSE or COPYING file for copyright information.

<!-- Badges -->

[slack-link]:   https://fisherman-wharf.herokuapp.com/
[slack-badge]:  https://img.shields.io/badge/slack-join%20the%20chat-00B9FF.svg?style=flat-square

<!-- Links -->

[fisherman]:        https://github.com/fisherman/fisherman
[fisher-submit]:    https://github.com/fishery/submit
