# vulnd

[![Build Status](https://travis-ci.org/gerred/vulndb.svg?branch=master)](https://travis-ci.org/gerred/vulndb)

## Building the DB

Build a DB with a list of files. Eg: `vulndb build nvdcve-2.0-2015.xml nvdcve-2.0-2014.xml`

If the DB already exists, it will throw an error. The `-f` flag will force a removal of the existing database.

## Searching the DB

Searching can either be done with a CPE URI or a raw string. Omitted fields or `*` will be treated as wildcards.

Example:

`vulndb search cpe:/o:microsoft`
`vulndb search cpe:/o:microsoft:windows_8.1:*`

If a raw string is used, it will do a full text search on everything.

Example:

`vulndb search drupal`
