package toolbar

import (
	_ "embed"
)

type (
	Object struct {
		CloseButton    *unison.Button
		OpenButton     *unison.Button
		RestartButton  *unison.Button
		ThreadsButton  *unison.Button
		CpuButton      *unison.Button
		OptionsButton  *unison.Button
		StepinButton   *unison.Button
		AppearButton   *unison.Button
		GotoButton     *unison.Button
		ModulesButton  *unison.Button
		WindowsButton  *unison.Button
		SourceButton   *unison.Button
		HandlesButton  *unison.Button
		MemoryButton   *unison.Button
		RunButton      *unison.Button
		TraceButton    *unison.Button
		LogButton      *unison.Button
		PauseButton    *unison.Button
		StackButton    *unison.Button
		TillretButton  *unison.Button
		TroverButton   *unison.Button
		PatchesButton  *unison.Button
		RefersButton   *unison.Button
		StepoverButton *unison.Button
		TrinButton     *unison.Button
		BreaksButton   *unison.Button
		HelpButton     *unison.Button
	}
)

func New() *Object {
	return &Object{
		CloseButton:    CloseKind.Button(),
		OpenButton:     OpenKind.Button(),
		RestartButton:  RestartKind.Button(),
		ThreadsButton:  ThreadsKind.Button(),
		CpuButton:      CpuKind.Button(),
		OptionsButton:  OptionsKind.Button(),
		StepinButton:   StepinKind.Button(),
		AppearButton:   AppearKind.Button(),
		GotoButton:     GotoKind.Button(),
		ModulesButton:  ModulesKind.Button(),
		WindowsButton:  WindowsKind.Button(),
		SourceButton:   SourceKind.Button(),
		HandlesButton:  HandlesKind.Button(),
		MemoryButton:   MemoryKind.Button(),
		RunButton:      RunKind.Button(),
		TraceButton:    TraceKind.Button(),
		LogButton:      LogKind.Button(),
		PauseButton:    PauseKind.Button(),
		StackButton:    StackKind.Button(),
		TillretButton:  TillretKind.Button(),
		TroverButton:   TroverKind.Button(),
		PatchesButton:  PatchesKind.Button(),
		RefersButton:   RefersKind.Button(),
		StepoverButton: StepoverKind.Button(),
		TrinButton:     TrinKind.Button(),
		BreaksButton:   BreaksKind.Button(),
		HelpButton:     HelpKind.Button(),
	}
}

func (o *Object) CanvasObject() *unison.Panel {
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlowLayout{
		HSpacing: unison.StdHSpacing,
		VSpacing: unison.StdVSpacing,
	})
	panel.AddChild(OpenKind.Button())
	panel.AddChild(PatchesKind.Button())
	panel.AddChild(RefersKind.Button())
	panel.AddChild(StepoverKind.Button())
	panel.AddChild(TrinKind.Button())
	panel.AddChild(BreaksKind.Button())
	panel.AddChild(ThreadsKind.Button())
	panel.AddChild(CloseKind.Button())
	panel.AddChild(RestartKind.Button())
	panel.AddChild(CpuKind.Button())
	panel.AddChild(OptionsKind.Button())
	panel.AddChild(StepinKind.Button())
	panel.AddChild(AppearKind.Button())
	panel.AddChild(GotoKind.Button())
	panel.AddChild(ModulesKind.Button())
	panel.AddChild(WindowsKind.Button())
	panel.AddChild(RunKind.Button())
	panel.AddChild(SourceKind.Button())
	panel.AddChild(HandlesKind.Button())
	panel.AddChild(MemoryKind.Button())
	panel.AddChild(TraceKind.Button())
	panel.AddChild(LogKind.Button())
	panel.AddChild(PauseKind.Button())
	panel.AddChild(StackKind.Button())
	panel.AddChild(TillretKind.Button())
	panel.AddChild(TroverKind.Button())
	panel.AddChild(HelpKind.Button())

	return panel
}
