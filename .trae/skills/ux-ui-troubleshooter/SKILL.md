---
name: "ux-ui-troubleshooter"
description: "Comprehensive guide for creating gio/ux UIs. Invoke when developing any gio/ux UI, creating widgets, layouts, or troubleshooting panics and compilation errors."
---

# Gio/UX UI Development Guide

Complete guide for creating UIs with the gio/ux widget library. Use this skill when developing any gio/ux interface.

## Development Workflow

When creating a new UI, follow this workflow to explore and select appropriate widgets:

### 1. Explore Examples Directory

First, check `D:\笔记本\p\ux\examples\` for similar UI patterns:

```
D:\笔记本\p\ux\examples\
├── designer/        # Full UI designer app - comprehensive widget usage
├── simple/          # Simple examples
└── hypedbg/         # HyperDbg specific examples
```

**Action**: Read example files to understand:
- How widgets are initialized
- Layout patterns used
- Event handling patterns
- Integration with business logic

### 2. Explore Widget Directory

Check `D:\笔记本\p\ux\widget\` for available widgets and their APIs:

```
D:\笔记本\p\ux\widget\
├── button/          # Button variants
├── input/           # Text input widgets
├── card/            # Card containers
├── chip/            # Chips and tags
├── logview/         # Log viewer
├── menu/            # Menus and dropdowns
├── dialog/          # Dialogs
├── slider/          # Sliders
├── switch/          # Toggle switches
├── checkbox/        # Checkboxes
├── radiobutton/     # Radio buttons
├── progress/        # Progress indicators
├── tab/             # Tab navigation
├── table/           # Table views
├── tree/            # Tree views
├── fab/             # Floating action buttons
├── navigationrail/  # Navigation rail
├── navigationbar/   # Bottom navigation
├── appbar/          # App bars
├── datepicker/      # Date picker
├── timepicker/      # Time picker
└── ...              # More widgets
```

**For each widget, check**:
- `<widget>/` main implementation file
- `<widget>/`_variants.go for factory functions
- `<widget>/`_style.go for styling options

### 3. Explore App Directory

Check `D:\笔记本\p\ux\app\` for application lifecycle:

```
D:\笔记本\p\ux\app\
├── app.go           # Main application entry
└── ...              # Theme, window management
```

**Key functions**:
- `app.Run(title, renderFunc)` - Start application
- `app.RunWithConfig(config, renderFunc)` - Start with custom config

### 4. Explore WDK Directory

Check `D:\笔记本\p\ux\wdk\` for utilities:

```
D:\笔记本\p\ux\wdk\
├── editor.go        # Text editor widget
├── theme.go         # Theme access
├── icons.go         # Icon handling
├── text.go          # Text rendering
└── ...
```

**Key utilities**:
- `wdk.GetMaterialTheme(gtx)` - Get current theme
- `wdk.NewEditor(singleLine, submit)` - Create editor
- `wdk.RequireIconWidget(icon)` - Get icon widget
- `wdk.TitleL/M/S(gtx, text)` - Title labels
- `wdk.LabelL/M/S(gtx, text)` - Body labels

### 5. Widget Selection Guide

| Need | Widget | Factory Function |
|------|--------|------------------|
| Text input | input | `input.FilledTextInput()` |
| Multi-line text | input | `input.FilledTextArea()` |
| Compact input | input | `input.CompactInput()` |
| Primary action | button | `button.Filled()` |
| Secondary action | button | `button.Elevated()` |
| Tertiary action | button | `button.Outlined()` |
| Text only action | button | `button.Text()` |
| Container | card | `&card.Card{Kind: card.Elevated}` |
| Status indicator | chip | `chip.Assist()` |
| Filter option | chip | `chip.Filter()` |
| Log output | logview | `logview.New()` |
| Numeric input | slider | `slider.StandardSlider()` |
| Toggle option | switch | `switch.New()` |
| Multi-select | checkbox | `checkbox.New()` |
| Single-select | radiobutton | `radiobutton.New()` |
| Navigation | navigationrail | `navigationrail.New()` |
| Tabs | tab | `tab.New()` |
| Table data | table | `table.New()` |
| Tree data | tree | `tree.New()` |
| Progress | progress | `progress.New()` |
| Dialog | dialog | `dialog.New()` |
| Menu | menu | `menu.New()` |
| Date input | datepicker | `datepicker.New()` |
| Time input | timepicker | `timepicker.New()` |
| Floating action | fab | `fab.New()` |

## Quick Start Template

```go
package main

import (
    "image"

    "gioui.org/layout"
    "gioui.org/op/clip"
    "gioui.org/op/paint"
    "gioui.org/unit"
    "github.com/ddkwork/ux/app"
    "github.com/ddkwork/ux/wdk"
    "github.com/ddkwork/ux/widget/button"
    "github.com/ddkwork/ux/widget/card"
    "github.com/ddkwork/ux/widget/input"
    "github.com/ddkwork/ux/widget/logview"
    "golang.org/x/exp/shiny/materialdesign/icons"
)

type MyUI struct {
    // Widgets here
}

func NewMyUI() *MyUI {
    d := &MyUI{}
    // Initialize widgets using factory functions
    return d
}

func (d *MyUI) Update(gtx layout.Context) {
    // Handle button clicks and events
}

func (d *MyUI) Layout(gtx layout.Context) layout.Dimensions {
    materialTheme := wdk.GetMaterialTheme(gtx)
    paint.FillShape(gtx.Ops, materialTheme.Scheme.Background.Color.AsNRGBA(), 
        clip.Rect(image.Rect(0, 0, gtx.Constraints.Max.X, gtx.Constraints.Max.Y)).Op())

    return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
        // Layout children
    )
}

func main() {
    d := NewMyUI()
    app.Run("My App", func(gtx layout.Context) {
        d.Update(gtx)
        d.Layout(gtx)
    })
}
```

## Widget Creation Reference

### Input Widgets

```go
// Single line text input
inputWidget := input.FilledTextInput()
inputWidget.LabelText = "Label"
inputWidget.Editor.SetText("default value")
text := inputWidget.Editor.GetText()

// Multi-line text area
textArea := input.FilledTextArea()

// Compact input
compactInput := input.CompactInput()
```

### Buttons

```go
// Button types
filledBtn := button.Filled()
elevatedBtn := button.Elevated()
outlinedBtn := button.Outlined()
textBtn := button.Text()
tonalBtn := button.FilledTonal()

// Layout
filledBtn.Layout(gtx, "Button Text")

// Click detection
if filledBtn.Clicked(gtx) {
    // Handle click
}
```

### Chips

```go
// Chip types
assistChip := chip.Assist()
actionChip := chip.Action()
filterChip := chip.Filter()

// Layout (requires label parameter)
assistChip.Layout(gtx, "Chip Label")
```

### Cards

```go
c := &card.Card{}
c.Kind = card.Elevated  // or card.Filled, card.Outlined
return c.Layout(gtx, card.Content(func(gtx layout.Context) layout.Dimensions {
    return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
        // Card content
    })
}))
```

### LogView

```go
log := logview.New()
log.Info("message")
log.Warn("warning")
log.Error("error")
log.Debugf("format: %s", value)
```

### Icons

```go
// Get icon widget
icon := wdk.RequireIconWidget(icons.HardwareMemory)

// Use icon (IconWidget is a function)
materialTheme := wdk.GetMaterialTheme(gtx)
icon(gtx, materialTheme.Scheme.Background.OnColor)
```

## Layout Patterns

### Vertical Layout

```go
layout.Flex{Axis: layout.Vertical}.Layout(gtx,
    layout.Rigid(func(gtx layout.Context) layout.Dimensions {
        // Fixed height content
    }),
    layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
        // Flexible content (fills remaining space)
    }),
)
```

### Horizontal Layout

```go
layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
    layout.Rigid(func(gtx layout.Context) layout.Dimensions {
        // Fixed width content
    }),
    layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
        // Flexible content
    }),
)
```

### Spacing & Padding

```go
layout.Rigid(layout.Spacer{Width: unit.Dp(16)}.Layout)
layout.Rigid(layout.Spacer{Height: unit.Dp(8)}.Layout)

layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
    // Padded content
})
```

## Theme & Colors

```go
materialTheme := wdk.GetMaterialTheme(gtx)

// Color access pattern
materialTheme.Scheme.Primary.Color        // Primary color
materialTheme.Scheme.Primary.OnColor      // Text on primary
materialTheme.Scheme.Background.Color     // Background color
materialTheme.Scheme.Background.OnColor   // Text on background
materialTheme.Scheme.Error.Color          // Error color
materialTheme.Scheme.Surface.Color        // Surface color
```

## Text Labels

```go
wdk.TitleL(gtx, "Large Title")
wdk.TitleM(gtx, "Medium Title")
wdk.TitleS(gtx, "Small Title")
wdk.LabelL(gtx, "Large Label")
wdk.LabelM(gtx, "Medium Label")
wdk.LabelS(gtx, "Small Label")
```

## Common Pitfalls & Solutions

### 1. Input Widget Panic

**Wrong**: `&input.Input{Editor: wdk.NewEditor(true, false)}`
**Correct**: `input.FilledTextInput()`

Input requires `contextArea` and `contextMenu` to be initialized. Always use factory functions.

### 2. Editor Arguments

**Wrong**: `wdk.NewEditor()`
**Correct**: `wdk.NewEditor(true, false)  // (singleLine, enableSubmit)`

### 3. IconWidget is a Function

**Wrong**: `icon.Layout(gtx)`
**Correct**: `icon(gtx, color)`

Signature: `type IconWidget func(gtx layout.Context, foreground theme.MatColor) layout.Dimensions`

### 4. Chip Layout Arguments

**Wrong**: `chip.Layout(gtx)`
**Correct**: `chip.Layout(gtx, "Label")`

### 5. Theme Color Access

**Wrong**: `materialTheme.Scheme.OnBackground.Color`
**Correct**: `materialTheme.Scheme.Background.OnColor`

MatColorSet structure: `{Color, OnColor}`

### 6. Missing Imports

Common required imports:
```go
import (
    "image"
    "gioui.org/layout"
    "gioui.org/op/clip"
    "gioui.org/op/paint"
    "gioui.org/unit"
)
```

## Debugging Checklist

1. Use factory functions for widgets
2. Check nil pointers before Layout calls
3. Verify method signatures match
4. Ensure all required imports present
5. Check theme color access pattern
6. Use `wdk.GetMaterialTheme(gtx)` for theme access
7. Read widget source in `D:\笔记本\p\ux\widget\` for exact API
8. Check examples in `D:\笔记本\p\ux\examples\` for usage patterns
