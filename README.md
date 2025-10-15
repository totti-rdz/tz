# tz CLI tool

A lightweight CLI tool to streamline development commands across multiple projects.

## The Problem

Working with multiple projects in different languages and frameworks means remembering different commands:

- Some projects use `npm install`, others use `yarn add`
- Dev servers might be `npm run dev`, `npm start`, or something else
- Test commands vary: `npm test`, `go test`, `pytest`, etc.

**tz** solves this by letting you use the same commands everywhere: `tz install`, `tz dev`, `tz test`

## Why tz?

**Zero project pollution** - Unlike task runners that require config files in each project, tz stores all mappings in `~/.tz/config.json`. Your projects stay clean, and your team isn't forced to use your workflow tool.
