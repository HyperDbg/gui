package pages

import (
	"strings"
	"sync"
	"testing"
)

func TestLogPage_Write(t *testing.T) {
	lp := NewLog()
	lp.Write([]byte("line1\nline2\n"))

	if lp.PendingCountForTest() != 2 {
		t.Errorf("pending count = %d, want 2", lp.PendingCountForTest())
	}
}

func TestLogPage_WriteNoTrailingNewline(t *testing.T) {
	lp := NewLog()
	lp.Write([]byte("no newline"))

	if lp.PendingCountForTest() != 1 {
		t.Errorf("pending count = %d, want 1", lp.PendingCountForTest())
	}
}

func TestLogPage_WriteEmpty(t *testing.T) {
	lp := NewLog()
	lp.Write([]byte(""))

	if lp.PendingCountForTest() != 0 {
		t.Errorf("pending count = %d, want 0", lp.PendingCountForTest())
	}
}

func TestLogPage_WriteOnlyNewline(t *testing.T) {
	lp := NewLog()
	lp.Write([]byte("\n"))

	// 尾部 \n 被去掉后是空字符串，不追加
	if lp.PendingCountForTest() != 0 {
		t.Errorf("pending count = %d, want 0", lp.PendingCountForTest())
	}
}

func TestLogPage_Flush(t *testing.T) {
	lp := NewLog()
	lp.Write([]byte("flush me\n"))

	lp.FlushForTest()
	if lp.PendingCountForTest() != 0 {
		t.Errorf("pending after flush = %d, want 0", lp.PendingCountForTest())
	}
	if lp.EntryCountForTest() != 1 {
		t.Errorf("entry count = %d, want 1", lp.EntryCountForTest())
	}
}

func TestLogPage_Printf(t *testing.T) {
	lp := NewLog()
	lp.Printf("val=%d str=%s", 42, "hello")

	lp.FlushForTest()
	text := lp.TextForTest()
	if !strings.Contains(text, "val=42") {
		t.Errorf("text = %q, want contains 'val=42'", text)
	}
	if !strings.Contains(text, "str=hello") {
		t.Errorf("text = %q, want contains 'str=hello'", text)
	}
}

func TestLogPage_Println(t *testing.T) {
	lp := NewLog()
	lp.Println("a", "b", "c")

	lp.FlushForTest()
	text := lp.TextForTest()
	if !strings.Contains(text, "a b c") {
		t.Errorf("text = %q, want contains 'a b c'", text)
	}
}

func TestLogPage_Clear(t *testing.T) {
	lp := NewLog()
	lp.Write([]byte("to be cleared\n"))
	lp.FlushForTest()

	lp.Clear()
	if lp.PendingCountForTest() != 0 {
		t.Errorf("pending after clear = %d, want 0", lp.PendingCountForTest())
	}
}

func TestLogPage_ConcurrentWrites(t *testing.T) {
	lp := NewLog()

	var wg sync.WaitGroup
	for i := range 50 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			lp.Printf("concurrent %d\n", n)
		}(i)
	}
	wg.Wait()

	if lp.PendingCountForTest() != 50 {
		t.Errorf("pending count = %d, want 50", lp.PendingCountForTest())
	}
}
