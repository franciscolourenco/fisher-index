[![Build Status][travis-badge]][travis-link]
[![Slack Room][slack-badge]][slack-link]

# Index

The index is a text file that list plugins, each consisting of the following fields:

```
name
url
plugin description
one or more tags
author
```

For example:

```
fishtape
https://github.com/fisherman/fishtape
TAP producer and test harness
test tap harness runner
bucaran
```

The fisherman index powers the [online] search.

## Adding plugins to this list

Create a pull request in this repository.

```fish
git clone https://github.com/fisherman/fisher-index
cd index
echo "$name\n$url\n$info\n$tags\n$author\n\n" >> index
git push origin master
```

## Test coverage

- [x] Index integrity
- [x] Uniqueness of plugin/prompt name
- [ ] Description
- [x] Integrity of URL
- [x] Uniqueness of URL
- [x] Uniqueness of tags within a plugin/prompt
- [x] Tags are shorter no more than 15 characters
- [x] No more than 4 tags per plugin
- [ ] Author.

## Disclaimer

All the plugins listed in this index are property of their respective owners. Follow the provided URL and see the bundled LICENSE or COPYING file for copyright information.

<!-- Badges -->

[slack-link]: https://fisherman-wharf.herokuapp.com
[slack-badge]: https://fisherman-wharf.herokuapp.com/badge.svg
[travis-link]: https://travis-ci.org/fisherman/index
[travis-badge]: https://img.shields.io/travis/fisherman/index.svg
[fisherman]: https://github.com/fisherman/fisherman
[online]: http://fisherman.sh
