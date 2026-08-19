package pages

// 本文件仅用于测试，导出 pages 包内部的 unexported 方法/字段供 _test.go 使用。
// Go 编译器只在 `go test` 时编译 *_test.go 文件。

// ===== LogPage 测试辅助 =====

// FlushForTest 暴露 unexported flush，供测试排空 pending 缓冲到 logView。
func (p *LogPage) FlushForTest() {
	p.flush()
}

// TextForTest 返回日志视图导出的文本。
func (p *LogPage) TextForTest() string {
	return p.logView.Export()
}

// EntryCountForTest 返回 logView 中的条目数。
func (p *LogPage) EntryCountForTest() int {
	return p.logView.GetEntryCount()
}

// PendingCountForTest 返回 pending 缓冲中的行数。
func (p *LogPage) PendingCountForTest() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return len(p.pending)
}

// ===== EventsPage 测试辅助 =====

// HooksForTest 返回 hooks 列表的副本。
func (p *EventsPage) HooksForTest() []HookRecord {
	p.mu.Lock()
	defer p.mu.Unlock()
	out := make([]HookRecord, len(p.hooks))
	copy(out, p.hooks)
	return out
}

// AddHookForTest 直接调用 addHook 供测试用。
func (p *EventsPage) AddHookForTest(rec HookRecord) {
	p.addHook(rec)
}

// SetAddrInputForTest 设置地址输入框文本。
func (p *EventsPage) SetAddrInputForTest(s string) {
	p.addrInput.Editor.SetText(s)
}

// SetPIDInputForTest 设置 PID 输入框文本。
func (p *EventsPage) SetPIDInputForTest(s string) {
	p.pidInput.Editor.SetText(s)
}

// SetTagInputForTest 设置 Tag 输入框文本。
func (p *EventsPage) SetTagInputForTest(s string) {
	p.tagInput.Editor.SetText(s)
}

// SetHookEnabledForTest 直接修改 hooks 列表中指定 tag 的 Enabled 状态。
func (p *EventsPage) SetHookEnabledForTest(tag uint64, enabled bool) {
	p.mu.Lock()
	for i := range p.hooks {
		if p.hooks[i].Tag == tag {
			p.hooks[i].Enabled = enabled
			break
		}
	}
	p.mu.Unlock()
}

// RemoveHookForTest 直接从 hooks 列表删除指定 tag。
func (p *EventsPage) RemoveHookForTest(tag uint64) {
	p.mu.Lock()
	for i, h := range p.hooks {
		if h.Tag == tag {
			p.hooks = append(p.hooks[:i], p.hooks[i+1:]...)
			break
		}
	}
	p.mu.Unlock()
	p.refreshList()
}

// RefreshListForTest 暴露 refreshList 供测试调用。
func (p *EventsPage) RefreshListForTest() {
	p.refreshList()
}

// ===== BreaksPage 测试辅助 =====

// SetAddrInputForTest 设置地址输入框文本。
func (p *BreaksPage) SetAddrInputForTest(s string) {
	p.addrInput.Editor.SetText(s)
}

// SetTagInputForTest 设置 Tag 输入框文本。
func (p *BreaksPage) SetTagInputForTest(s string) {
	p.tagInput.Editor.SetText(s)
}

// ===== CpuPage 测试辅助 =====

// SetAddrInputForTest 设置地址输入框文本。
func (c *CpuPage) SetAddrInputForTest(s string) {
	c.addrEditor.Editor.SetText(s)
}
