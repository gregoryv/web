# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [unreleased]
### Changed

- htmlwriter adds newlines after table and tr tags
- Markdown writer indents pre tags even if no newlines are present
- Newline after header and nav elements

## [0.8.2] 2020-03-17
### Changed

- HtmlWriter no longer includes newlines in A tags
- No newline before closing div
- Text elements are not separated with newlines

## [0.8.1] 2020-01-19
### Added

- apidoc subpackage for documenting HTTP APIs
- Alt attribute

## [0.7.0] 2020-01-16
### Added

- MarkdownWriter

## [0.6.0] 2020-01-15
### Added

- Element.String simpler use when outside of other elements
- Type Page for saving pages

### Changed

- Renamed Tag to Element
- Renamed Attr to Attribute
- Renamed constructors NewElement and NewSimpleElement
- Making some Element fields public

## [0.5.0] - 2019-12-31
### Changed

- WriterTo signature matches io.WriterTo

## [0.4.0] - 2019-12-31
### Changed

- Moved site/* and doctype/* into package web

## [0.3.0] - 2019-12-29
### Changed

- Expose Attr Name and Val for generic cases

## [0.2.1] - 2019-12-29
### Added

- Tag constructor Quote
- Attribute constructors Id, Class

## [0.2.0] - 2019-12-29
### Added

- Tag.With method to add children or attributes post creation
- Attribute constructors Type and Href

## [0.1.0] - 2019-12-28
### Added

- package doctype for html rendering
- cmd toc for checking links
