# Changelog
All notable changes to this project will be documented in this file.

This project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [unreleased]

- Render A tags in markdown
- MakeTOC adds links for named headers for easier sharing

## [0.13.0] 2020-11-07

- Add toc.MakeTOC combining generate id and generating toc
- Add package files
- Add CSS.Import for eg. font imports

## [0.12.0] 2020-10-25

- Added Page.SaveAs with filename argument
- Remove filename argument from NewPage, use NewFile instead
- Add apidoc example

## [0.11.0] 2020-09-24

- Use json.Indent for nice json formatting
- Increase pre-element indentation to fours paces for markdown output

## [0.10.0] 2020-09-22

- Add package toc for working with table of contents

## [0.9.0] 2020-06-03

- Add CSS struct for generating inline stylesheets
- Element implements io.WriterTo
- Add func Attr for arbitrary attributes
- Add attributes for mouse events and more
- HtmlWriter also renders any implementation of WriterTo

## [0.8.3] 2020-04-10

- Add Em tag
- Add attributes action, autocomplete, method, tabindex, value, fieldset
- HtmlWriter adds newlines after table and tr tags
- Markdown writer indents pre tags even if no newlines are present
- Newline after header and nav elements

## [0.8.2] 2020-03-17

- HtmlWriter no longer includes newlines in A tags
- No newline before closing div
- Text elements are not separated with newlines

## [0.8.1] 2020-01-19

- Add apidoc subpackage for documenting HTTP APIs
- Add Alt attribute

## [0.7.0] 2020-01-16

- Add MarkdownWriter

## [0.6.0] 2020-01-15

- Add Element.String simpler use when outside of other elements
- Type Page for saving pages
- Renamed Tag to Element
- Renamed Attr to Attribute
- Renamed constructors NewElement and NewSimpleElement
- Making some Element fields public

## [0.5.0] - 2019-12-31

- WriterTo signature matches io.WriterTo

## [0.4.0] - 2019-12-31

- Moved site/* and doctype/* into package web

## [0.3.0] - 2019-12-29

- Expose Attr Name and Val for generic cases

## [0.2.1] - 2019-12-29

- Add quote element
- Add attribute constructors Id, Class

## [0.2.0] - 2019-12-29

- Add Tag.With method to add children or attributes post creation
- Add attribute constructors Type and Href

## [0.1.0] - 2019-12-28

- Add package doctype for html rendering
- Add cmd toc for checking links
