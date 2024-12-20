# Changelog
All notable changes to this project will be documented in this file.

This project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [0.27.0] 2024-12-19

- Add func apidoc.DocumentRouter
- Use type apidoc.Doc to index routes for additional documentation

## [0.26.1] 2024-03-22

- Update dependencies

## [0.26.0] 2024-03-17

- Add funcs For, Pattern and Main

## [0.25.1] 2024-02-15

- Bump license year to 2024

## [0.25.0] 2023-10-21

- Set go 1.21 in go.mod
- Remove link checking code to minimise dependencies

## [0.24.0] 2022-07-17

- Page.SaveAs handles paths/filename 
- Page.SaveTo tries to create directories before saving
- Add query expr #id

## [0.23.0] 2022-01-20

- Add css style query capabilities using func Query and ParseExpr
- files.LoadFunc parses comments aswell
- Remove cmd/*
- Remove gregoryv/workdir dependency
- Update gregoryv/english dependency

## [0.22.0] 2021-10-23

- Add theme GoishColors
- Add method CSS.With to simplify combinations
- Add package theme with GoldenSpace

## [0.21.0] 2021-09-23

- MarkdownEncoder does not trim spaces for Pre tags
- Func GenerateID and TOC related funcs generate uniq ids for similar titles

## [0.20.0] 2021-07-25

- Add type ElementBuilder which both html- and markdown- encoders recognize
- Add input related attributes for, max, min, maxlength, pattern
  placeholder and size

## [0.19.0] 2021-06-01

- Add func Comment() for html comments
- Add NewSafePage and NewSafeHtmlEncoder with auto escaping of string values
- Add attribute formaction

## [0.18.0] 2021-05-27

- Add files.MustLoadFunc for loading func bodies
- Remove sweb package, imposed context switching as the structure was hidden

## [0.17.0] 2021-04-09

- Add package sweb for sequential writing style
- Remove CheckLinks channel argument
- Remove CheckLink
- Fixe func LinkAll and ILinkAll to exclude words in existing A elements

## [0.16.0] 2021-03-04

- Add CSS.SaveAs and related methods for saving to file
- Add CSS.ServeHTTP for using stylesheets as static resources

## [0.15.0] 2021-02-03

- Add Page.ServeHTTP for using pages as static resources
- Add func LinkAll and ILinkAll for simple link injection
- Add func Wrap with special wrapper element that can be used to group elements
- Add type Hn for dynamic headings
- Allow for multiple media in CSS, using Media and SetMedia methods

## [0.14.0] 2020-11-16

- Page SaveAs, SaveTo and WriteTo use markdown for .md suffix in filename
- Rename MarkdownWriter to MarkdownEncoder
- Rename HtmlWriter to HtmlEncoder and WriteTo method to Encode
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
- Rename Tag to Element
- Rename Attr to Attribute
- Rename constructors NewElement and NewSimpleElement
- Make some Element fields public

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
