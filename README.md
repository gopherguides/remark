# 🎯 Remark - Advanced Markdown Processing & HTML Generation

> **⚠️ DEPRECATION NOTICE**
> This project is no longer maintained and has been completely replaced by a ground-up rebuild in the [Hype](https://github.com/gopherguides/hype) project. Please use Hype for all new projects.

[![Tests](https://github.com/gopherguides/remark/workflows/Tests/badge.svg)](https://github.com/gopherguides/remark/actions)
[![codecov](https://codecov.io/gh/gopherguides/remark/branch/main/graph/badge.svg)](https://codecov.io/gh/gopherguides/remark)
[![GoDoc](https://godoc.org/github.com/gopherguides/remark?status.svg)](https://godoc.org/github.com/gopherguides/remark)
[![Go Report Card](https://goreportcard.com/badge/github.com/gopherguides/remark)](https://goreportcard.com/report/github.com/gopherguides/remark)

**Remark** is a powerful Go package that transforms markdown into structured,
processable content. Built for educational platforms, documentation systems, and
content management, it provides a sophisticated tag-based architecture for
parsing, processing, and rendering markdown with advanced features like file
inclusion, code syntax highlighting, metadata extraction, and HTML generation.

## 🚀 Features

### Core Capabilities

- **Dual Parser Support**: Process both Markdown (`md`) and HTML (`htm`) content
- **Tag-Based Architecture**: Structured content representation through a
  unified `Tag` interface
- **File Inclusion**: Dynamic content composition with `<include>` tags
- **Advanced Code Blocks**: Syntax highlighting with snippet support and
  language detection
- **Metadata Extraction**: Parse and process document metadata from HTML
  details tags
- **Section Management**: Automatic content sectioning with horizontal rule delimiters
- **HTML Generation**: Convert processed content to clean HTML output
- **Table of Contents**: Generate structured TOCs from heading hierarchies

### Advanced Features

- **Template Processing**: Go template integration for dynamic content generation
- **Image Processing**: Automatic image validation and path resolution
- **Link Processing**: Smart link handling and validation
- **Attribute Management**: Rich attribute system for all content elements
- **Custom Printers**: Extensible rendering system with custom tag processors
- **Snippet Management**: Advanced code snippet extraction and processing

## 📦 Installation

### Package Installation

```bash
go get github.com/gopherguides/remark
```

### CLI Tools

```bash
# Install the remark CLI processor
go install github.com/gopherguides/remark/cmd/remark@latest

# Install the table of contents generator
go install github.com/gopherguides/remark/cmd/retoc@latest
```

## 🎮 Quick Start

### Basic Markdown Processing

```go
package main

import (
    "fmt"
    "github.com/gopherguides/remark/md"
)

func main() {
    markdown := `# Hello World

This is a **markdown** document with:
- Lists
- Code blocks
- And more!

## Code Example

` + "```go" + `
func main() {
    fmt.Println("Hello, World!")
}
` + "```" + `
`

    // Parse the markdown
    tags, err := md.Parse(".", []byte(markdown))
    if err != nil {
        panic(err)
    }

    // Output the processed content
    fmt.Println(tags)
}
```

### Advanced Processing with File Inclusion

```markdown
# My Course Module

Welcome to the course!

---

<include src="setup.md"></include>

---

<include src="assignments/assignment01.md"></include>
```

### HTML Processing

```go
package main

import (
    "fmt"
    "github.com/gopherguides/remark/htm"
)

func main() {
    html := `<div class="content">
        <h1>Course Title</h1>
        <details>
            course: advanced-go
            difficulty: intermediate
        </details>
        <code src="./examples/hello.go"></code>
    </div>`

    // Parse HTML content
    doc, err := htm.Parse([]byte(html))
    if err != nil {
        panic(err)
    }

    // Access metadata
    metadata := doc.Metadata()
    fmt.Printf("Course: %s\n", metadata["course"])
    fmt.Printf("Difficulty: %s\n", metadata["difficulty"])
}
```

## 🛠️ CLI Tools

### remark - Markdown Processor

Process markdown from stdin and output structured content:

```bash
# Process a markdown file
cat document.md | remark

# Set working directory for includes
MARKED_ORIGIN=/path/to/content cat document.md | remark
```

### retoc - Table of Contents Generator

Generate structured table of contents from markdown files:

```bash
# Generate TOC for specific files
retoc /path/to/content/

# Example output:
# Course Introduction
#     Getting Started
#     Requirements
#         System Requirements
#         Software Installation
#     First Steps
```

## 🏗️ Architecture

### Tag System

Remark uses a unified `Tag` interface for all content elements:

```go
type Tag interface {
    Attrs() Attributes
    GetChildren() Tags
    Options() tags.Options
    TagName() string
    fmt.Stringer
}
```

### Core Components

#### 1. **Generic Tags**

Universal container for any HTML-like element:

```go
generic := remark.NewGeneric("div")
generic.Set("class", "content")
generic.Append(remark.String("Hello World"))
```

#### 2. **Headings**

Structured heading elements with level information:

```go
type Heading struct {
    *Generic
    Level int  // 1-6 for h1-h6
}
```

#### 3. **Code Blocks**

Advanced code processing with syntax highlighting:

```go
type CodeBlock struct {
    *Generic
    Language string
    Snippets Snippets
}
```

#### 4. **Sections**

Document sections with metadata support:

```go
type Section struct {
    *Generic
    Title string
}
```

### Parsers

#### Markdown Parser (`md`)

- Full CommonMark compliance
- Extended syntax support (tables, strikethrough, etc.)
- Template processing
- File inclusion
- Custom extension support

#### HTML Parser (`htm`)

- Clean HTML parsing
- Metadata extraction from `<details>` tags
- Image validation
- Custom tag processing

## 🎯 Use Cases

### Educational Content Management

Perfect for course materials, tutorials, and documentation:

```markdown
# Week 1: Introduction to Go

<details>
overview: true
difficulty: beginner
duration: 2 hours
</details>

Welcome to our Go programming course!

<include src="setup-instructions.md"></include>

## Your First Program

<code src="examples/hello.go"></code>
```

### Documentation Systems

Build comprehensive documentation with cross-references:

```markdown
# API Documentation

<include src="authentication.md"></include>

<include src="endpoints/users.md"></include>

<include src="examples/complete-example.md"></include>
```

### Content Publishing

Create rich content with embedded examples:

```markdown
# Tutorial: Building a Web Server

<code src="server.go" snippet="basic-server"></code>

The code above shows...

<code src="server.go" snippet="with-middleware"></code>
```

## 🔧 Advanced Features

### Custom Printers

Create custom rendering logic:

```go
printer := &htm.Printer{}

// Custom code block renderer
printer.Set("code", func(t remark.Tag) (string, error) {
    code := t.(*remark.CodeBlock)
    return fmt.Sprintf(`<pre class="custom"><code>%s</code></pre>`, 
        html.EscapeString(code.Children.String())), nil
})

// Render with custom logic
html, err := printer.Print(tags...)
```

### Metadata Processing

Extract and use document metadata:

```go
// Find sections with overview metadata
for _, tag := range tags {
    if section, ok := tag.(*md.Section); ok {
        if section.Overview() {
            overview := tags.Overview() // Get overview text
            fmt.Printf("Overview: %s\n", overview)
        }
    }
}
```

### Snippet Management

Process code snippets with markers:

```go
// In your Go file:
// snippet:start:basic-server
func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
// snippet:end:basic-server
```

```markdown
<code src="server.go" snippet="basic-server"></code>
```

## 🤝 Integration with Hype

Remark is designed to work seamlessly with the
[Hype](https://github.com/gopherguides/hype) package for advanced content
generation and templating. Together, they form a powerful content processing
pipeline for educational platforms and documentation systems.

## 📚 API Reference

### Core Functions

#### Markdown Processing

```go
// Parse markdown from bytes
md.Parse(root string, src []byte) (remark.Tags, error)

// Parse markdown from file
md.ParseFile(filename string) (remark.Tags, error)

// Create new parser
md.NewParser(root string) *Parser
```

#### HTML API Functions

```go
// Parse HTML content
htm.Parse(src []byte) (*Document, error)

// Create new HTML parser
htm.NewParser(root string) *Parser

// Print tags to HTML
htm.Print(tags ...remark.Tag) (string, error)
```

#### Tag Operations

```go
// Find specific tags
tags.FindFirst(name string) (Tag, bool)
tags.FindAll(name string) Tags

// Get document body
tags.Body() (Tag, bool)

// Extract overview
tags.Overview() string
```

### Interfaces

#### Essential Interfaces

```go
// Core tag interface
type Tag interface {
    Attrs() Attributes
    GetChildren() Tags
    Options() tags.Options
    TagName() string
    fmt.Stringer
}

// Metadata support
type Metadatable interface {
    Tag
    Metadata() Metadata
}

// Appendable content
type Appendable interface {
    Tag
    Append(tag Tag)
}
```

## 🎨 Examples

### Complete Processing Pipeline

```go
func processContentDirectory(dir string) error {
    return filepath.Walk(dir, func(path string, info os.FileInfo, 
        err error) error {
        if filepath.Ext(path) != ".md" {
            return nil
        }

        // Parse the markdown file
        tags, err := md.ParseFile(path)
        if err != nil {
            return err
        }

        // Process each section
        for _, tag := range tags {
            if section, ok := tag.(*md.Section); ok {
                fmt.Printf("Section: %s\n", section.Title)
                
                // Extract metadata
                metadata := section.Metadata()
                if difficulty, ok := metadata["difficulty"]; ok {
                    fmt.Printf("Difficulty: %s\n", difficulty)
                }

                // Find code blocks
                codeBlocks := section.GetChildren().FindAll("code")
                fmt.Printf("Code blocks: %d\n", len(codeBlocks))
            }
        }

        return nil
    })
}
```

### Custom Content Processor

```go
type CourseProcessor struct {
    parser *md.Parser
}

func (cp *CourseProcessor) ProcessCourse(content []byte) (*Course, error) {
    tags, err := cp.parser.Parse(content)
    if err != nil {
        return nil, err
    }

    course := &Course{
        Modules: make([]Module, 0),
    }

    for _, tag := range tags {
        if section, ok := tag.(*md.Section); ok {
            module := Module{
                Title:   section.Title,
                Content: section.GetChildren().String(),
            }

            // Extract difficulty and duration
            metadata := section.Metadata()
            module.Difficulty = metadata["difficulty"]
            module.Duration = metadata["duration"]

            course.Modules = append(course.Modules, module)
        }
    }

    return course, nil
}
```

## 🔄 Requirements

- **Go**: 1.24 or higher
- **Go Modules**: Required for dependency management

## 📝 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE)
file for details.

## 🤝 Contributing

We welcome contributions! Please see our contributing guidelines and feel free
to submit issues or pull requests.

## 🎯 Built For

- **Educational Platforms**: Course content management and delivery
- **Documentation Systems**: Technical documentation with code examples
- **Content Management**: Rich content processing and publishing
- **Static Site Generation**: Advanced markdown processing for websites

---

**Remark** - Transform your markdown into structured, powerful content. Built
with ❤️ by the [Gopher Guides](https://github.com/gopherguides) team.
