# VulnDB

[![Build Status](https://travis-ci.org/gerred/vulndb.svg?branch=master)](https://travis-ci.org/gerred/vulndb)

## Building the DB

Build a DB with a list of files. Eg: 
`vulndb build nvdcve-2.0-2015.xml nvdcve-2.0-2014.xml`
`vulndb build ~/Downloads/*.xml`

If the DB already exists, it will throw an error. The `-f` flag will force a removal of the existing database.

## Searching the DB

Searching can either be done with a CPE URI or a raw string. Omitted fields or `*` will be treated as wildcards.

Example:

`vulndb search cpe:/o:microsoft`
`vulndb search cpe:/o:microsoft:windows_8.1:*`

If a raw string is used, it will do a full text search on everything.

Example:

`vulndb search drupal`

# NVD/CVE XML Feed with CVSS and CPE mappings

https://nvd.nist.gov/download.cfm#RSS

https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-2015.xml.gz

## Contributions

Contributions are welcome.

## Creators

**Remco Verhoef**
- <https://twitter.com/remco_verhoef>
- <https://twitter.com/dutchcoders>

## Copyright and license

Code and documentation copyright 2011-2015 Remco Verhoef.

Code released under [the MIT license](LICENSE).
