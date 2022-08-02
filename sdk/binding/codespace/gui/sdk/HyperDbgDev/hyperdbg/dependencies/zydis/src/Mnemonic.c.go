package src



type (
	Mnemonic interface {
		const_charPtr_ZydisMnemonicGetString() (ok bool)                    //col:8
		const_ZydisShortStringPtr_ZydisMnemonicGetStringWrapped() (ok bool) //col:16
	}
	mnemonic struct{}
)

func NewMnemonic() Mnemonic { return &mnemonic{} }

func (m *mnemonic) const_charPtr_ZydisMnemonicGetString() (ok bool) { //col:8











	return true
}


func (m *mnemonic) const_ZydisShortStringPtr_ZydisMnemonicGetStringWrapped() (ok bool) { //col:16











	return true
}


