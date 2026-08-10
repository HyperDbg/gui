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
    └── walker/       # PDB parsing demos
    └── x64dbg/      # Debugger UI demos (IMPORTANT: contains treetable, symbol, disassembly examples)
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
├── treetable/       # TreeTable widget (IMPORTANT: for large data display)
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

### 3. Explore App Directory

Check `D:\笔记本\p\ux\app\` for application lifecycle:

```
D:\笔记本\p\ux\app\
├── debug.go         # Main application entry and utilities
```

**Key functions**:
- `app.Run(title, renderFunc)` - Start application
- `app.RunWithConfig(config, renderFunc)` - Start with custom config
- `app.RequestRedraw()` - Request UI redraw from background thread

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

---

## Critical: TreeTable Usage Pattern

**This is the MOST IMPORTANT widget for displaying large datasets (functions, symbols, disassembly, etc.)**

### Correct TreeTable Usage

```go
// 1. Define row struct with table tag
type FuncRow struct {
    Name    string `table:"Name"`
    Address string `table:"Address"`
    Size    string `table:"Size"`
}

// 2. Initialize TreeTable
table := treetable.NewTreeTable[FuncRow]()

// 3. Configure TableContext with SetRootRowsCallBack
table.AirTable.TableContext = treetable.TableContext{
    SetRootRowsCallBack: func() {},  // Empty callback - data is populated in Update()
    JsonName:            "Functions",
}

// 4. In Update() method - check for pending data and update table
func (v *MyViewer) Update(gtx layout.Context) {
    // ... button click handling ...

    // Check and apply pending data in main thread
    v.dataMu.Lock()
    if v.dataReady {
        v.dataReady = false
        rows := v.pendingData
        v.dataMu.Unlock()

        // Clear and repopulate table
        table.Root().SetChildren(nil)
        for _, row := range rows {
            table.Root().AddChild(table.NewNode(row))
        }
        table.AirTable.Refresh()
    } else {
        v.dataMu.Unlock()
    }
}
```

### Background Data Loading Pattern

```go
type Viewer struct {
    table          *treetable.TreeTable[RowType]
    pendingData    []RowType
    dataReady      bool
    dataMu         sync.Mutex
    loading        bool
    loadingMu      sync.Mutex
}

func (v *Viewer) loadDataAsync() {
    v.loadingMu.Lock()
    if v.loading {
        v.loadingMu.Unlock()
        return
    }
    v.loading = true
    v.loadingMu.Unlock()

    defer func() {
        v.loadingMu.Lock()
        v.loading = false
        v.loadingMu.Unlock()
        app.RequestRedraw()  // CRITICAL: Request redraw after background work
    }()

    // Do heavy work (PDB parsing, etc.) in goroutine
    data := doHeavyWork()

    // Store pending data (NOT modifying table directly!)
    v.dataMu.Lock()
    v.pendingData = data
    v.dataReady = true
    v.dataMu.Unlock()
}
```

### DO NOT DO THIS

- ❌ Modifying table data directly in goroutine
- ❌ Calling `table.Root().AddChild()` in background thread
- ❌ Forgetting `app.RequestRedraw()` after background work
- ❌ Using `list.List` for large datasets (causes UI lag)

### DO THIS

- ✅ Store pending data in goroutine, process in Update()
- ✅ Use `sync.Mutex` for thread-safe data sharing
- ✅ Call `app.RequestRedraw()` to trigger UI refresh
- ✅ Use `treetable.TreeTable` for large datasets

---

## Widget Selection Guide

| Need | Widget | Factory Function |
|------|--------|------------------|
| Large data display | treetable | `treetable.NewTreeTable[T]()` |
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
| Tree data | tree | `tree.New()` |
| Progress | progress | `progress.New()` |
| Dialog | dialog | `dialog.New()` |
| Menu | menu | `menu.New()` |
| Date input | datepicker | `datepicker.New()` |
| Time input | timepicker | `timepicker.New()` |
| Floating action | fab | `fab.New()` |

---

## Quick Start Template

```go
package main

import (
    "image"
    "sync"

    "gioui.org/layout"
    "gioui.org/op/clip"
    "gioui.org/op/paint"
    "gioui.org/unit"
    "github.com/ddkwork/ux/app"
    "github.com/ddkwork/ux/wdk"
    "github.com/ddkwork/ux/widget/button"
    "github.com/ddkwork/ux/widget/input"
    "github.com/ddkwork/ux/widget/treetable"
)

type RowData struct {
    Name  string `table:"Name"`
    Value string `table:"Value"`
}

type MyUI struct {
    table       *treetable.TreeTable[RowData]
    loadButton  *button.Button
    input       *input.Input

    // Background loading state
    loading      bool
    loadingMu    sync.Mutex
    pendingData  []RowData
    dataReady    bool
    dataMu       sync.Mutex
}

func NewMyUI() *MyUI {
    ui := &MyUI{
        table:      treetable.NewTreeTable[RowData](),
        loadButton: button.Filled(),
        input:      input.FilledTextInput(),
    }

    ui.table.AirTable.TableContext = treetable.TableContext{
        SetRootRowsCallBack: func() {},
        JsonName:            "Data",
    }

    return ui
}

func (ui *MyUI) Update(gtx layout.Context) {
    // Handle button clicks in main thread
    if ui.loadButton.Clicked(gtx) {
        go ui.loadDataAsync()
    }

    // Process pending data in main thread (CRITICAL!)
    ui.dataMu.Lock()
    if ui.dataReady {
        ui.dataReady = false
        rows := ui.pendingData
        ui.dataMu.Unlock()

        ui.table.Root().SetChildren(nil)
        for _, row := range rows {
            ui.table.Root().AddChild(ui.table.NewNode(row))
        }
        ui.table.AirTable.Refresh()
    } else {
        ui.dataMu.Unlock()
    }
}

func (ui *MyUI) loadDataAsync() {
    ui.loadingMu.Lock()
    if ui.loading {
        ui.loadingMu.Unlock()
        return
    }
    ui.loading = true
    ui.loadingMu.Unlock()

    defer func() {
        ui.loadingMu.Lock()
        ui.loading = false
        ui.loadingMu.Unlock()
        app.RequestRedraw()
    }()

    // Do heavy work here...
    data := []RowData{{Name: "Item1", Value: "Value1"}}

    // Store pending data
    ui.dataMu.Lock()
    ui.pendingData = data
    ui.dataReady = true
    ui.dataMu.Unlock()
}

func (ui *MyUI) Layout(gtx layout.Context) layout.Dimensions {
    materialTheme := wdk.GetMaterialTheme(gtx)
    paint.FillShape(gtx.Ops, materialTheme.Scheme.Background.Color.AsNRGBA(),
        clip.Rect(image.Rect(0, 0, gtx.Constraints.Max.X, gtx.Constraints.Max.Y)).Op())

    return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
        layout.Rigid(func(gtx layout.Context) layout.Dimensions {
            label := "Load"
            ui.loadingMu.Lock()
            if ui.loading {
                label = "Loading..."
            }
            ui.loadingMu.Unlock()
            return ui.loadButton.Layout(gtx, label)
        }),
        layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
            return ui.table.AirTable.Layout(gtx)
        }),
    )
}

func main() {
    ui := NewMyUI()
    app.Run("My App", func(gtx layout.Context) {
        ui.Update(gtx)
        ui.Layout(gtx)
    })
}
```

---

## Widget Creation Reference

### TreeTable (IMPORTANT!)

```go
// Define row type with table tags
type MyRow struct {
    Column1 string `table:"Column 1"`
    Column2 string `table:"Column 2"`
}

// Create table
table := treetable.NewTreeTable[MyRow]()

// Configure context
table.AirTable.TableContext = treetable.TableContext{
    SetRootRowsCallBack: func() {},  // Called during layout
    JsonName:            "MyTable",
}

// Add data (in Update method, not in background thread!)
table.Root().SetChildren(nil)
table.Root().AddChild(table.NewNode(MyRow{Column1: "value1", Column2: "value2"}))
table.AirTable.Refresh()
```

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

---

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

---

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

---

## Text Labels

```go
wdk.TitleL(gtx, "Large Title")
wdk.TitleM(gtx, "Medium Title")
wdk.TitleS(gtx, "Small Title")
wdk.LabelL(gtx, "Large Label")
wdk.LabelM(gtx, "Medium Label")
wdk.LabelS(gtx, "Small Label")
```

---

## Code Search Strategy

### Finding Examples

1. **Search by widget name**:
   ```
   Grep pattern: treetable\.NewTreeTable
   path: d:\笔记本\p\ux\examples
   ```

2. **Search by pattern**:
   ```
   Grep pattern: func.*Update\(\).*error
   path: d:\笔记本\p\ux\examples\x64dbg
   ```

3. **Search for background loading**:
   ```
   Grep pattern: go func|goroutine|async
   path: d:\笔记本\p\ux\examples\x64dbg
   ```

4. **Search for RequestRedraw**:
   ```
   Grep pattern: RequestRedraw
   path: d:\笔记本\p\ux
   ```

5. **Search for mutex patterns**:
   ```
   Grep pattern: sync\.Mutex|loadingMu|dataMu
   path: d:\笔记本\p\ux\examples
   ```

### Key Files to Reference

- `d:\笔记本\p\ux\examples\x64dbg\debugger\symbol\symbol.go` - TreeTable + background loading pattern
- `d:\笔记本\p\ux\examples\x64dbg\debugger\disassembly\disassembly.go` - TreeTable with complex data
- `d:\笔记本\p\ux\examples\x64dbg\debugger\debugger.go` - UI update pattern
- `d:\笔记本\p\ux\widget\treetable\table.go` - TreeTable implementation

---

## Common Pitfalls & Solutions

### 1. TreeTable Data Not Showing

**Problem**: Background thread modifies table but UI doesn't update
**Solution**: 
- Store pending data in goroutine
- Process in Update() method (main thread)
- Call `app.RequestRedraw()` after goroutine completes

### 2. UI Lag with Large Data

**Problem**: List widget causes lag with thousands of items
**Solution**: Use `treetable.TreeTable` instead of `list.List`

### 3. Input Widget Panic

**Wrong**: `&input.Input{Editor: wdk.NewEditor(true, false)}`
**Correct**: `input.FilledTextInput()`

Input requires `contextArea` and `contextMenu` to be initialized. Always use factory functions.

### 4. Editor Arguments

**Wrong**: `wdk.NewEditor()`
**Correct**: `wdk.NewEditor(true, false)  // (singleLine, enableSubmit)`

### 5. IconWidget is a Function

**Wrong**: `icon.Layout(gtx)`
**Correct**: `icon(gtx, color)`

Signature: `type IconWidget func(gtx layout.Context, foreground theme.MatColor) layout.Dimensions`

### 6. Chip Layout Arguments

**Wrong**: `chip.Layout(gtx)`
**Correct**: `chip.Layout(gtx, "Label")`

### 7. Theme Color Access

**Wrong**: `materialTheme.Scheme.OnBackground.Color`
**Correct**: `materialTheme.Scheme.Background.OnColor`

MatColorSet structure: `{Color, OnColor}`

### 8. Missing Imports

Common required imports:
```go
import (
    "image"
    "sync"
    "gioui.org/layout"
    "gioui.org/op/clip"
    "gioui.org/op/paint"
    "gioui.org/unit"
    "github.com/ddkwork/ux/app"
)
```

---

## Debugging Checklist

1. ✅ Use factory functions for widgets
2. ✅ Check nil pointers before Layout calls
3. ✅ Verify method signatures match
4. ✅ Ensure all required imports present
5. ✅ Check theme color access pattern
6. ✅ Use `wdk.GetMaterialTheme(gtx)` for theme access
7. ✅ Use `treetable.TreeTable` for large datasets
8. ✅ Process table updates in Update() method, not background thread
9. ✅ Call `app.RequestRedraw()` after background work
10. ✅ Use `sync.Mutex` for thread-safe data sharing
