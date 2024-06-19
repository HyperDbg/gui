package ARImpRec

import (
	"github.com/can1357/gengo/clang"
	"github.com/can1357/gengo/gengo"
	"github.com/ddkwork/golibrary/mylog"
	"testing"
)

func TestGetProcName(t *testing.T) {
	GetProcName()
}

func TestARImpRec(t *testing.T) {
	pkg := gengo.NewPackage("ARImpRec",
		gengo.WithRemovePrefix(
			"ARImpRec_", "Zyan_", "Zycore_",
			"ARImpRec", "Zyan", "Zycore",
		),
		gengo.WithInferredMethods([]gengo.MethodInferenceRule{
			{Name: "ARImpRecDecoder", Receiver: "Decoder"},
			{Name: "ARImpRecEncoder", Receiver: "EncoderRequest"},
			{Name: "ARImpRecFormatterBuffer", Receiver: "FormatterBuffer"},
			{Name: "ARImpRecFormatter", Receiver: "ARImpRecFormatter *"},
			{Name: "ZyanVector", Receiver: "Vector"},
			{Name: "ZyanStringView", Receiver: "StringView"},
			{Name: "ZyanString", Receiver: "String"},
			{Name: "ARImpRecRegister", Receiver: "Register"},
			{Name: "ARImpRecMnemonic", Receiver: "Mnemonic"},
			{Name: "ARImpRecISASet", Receiver: "ISASet"},
			{Name: "ARImpRecISAExt", Receiver: "ISAExt"},
			{Name: "ARImpRecCategory", Receiver: "Category"},
		}),
		gengo.WithForcedSynthetic(
			"ARImpRecShortString_",
			"struct ARImpRecShortString_",
		),
	)
	mylog.Check(pkg.Transform("ARImpRec", &clang.Options{
		Sources: []string{"ARImpRec/ARImpRec.h"},
		// Sources: []string{"./ARImpRec.h"},
		AdditionalParams: []string{
			//"-DZYAN_NO_LIBC",
			//"-DZYAN_STATIC_ASSERT",
			//"-IC:\\Users\\Admin\\Desktop\\ARImpRec\\dependencies\\zycore\\include",
			//"-IC:\\Users\\Admin\\Desktop\\ARImpRec\\include",
		},
	}))

	mylog.Check(pkg.WriteToDir("ARImpRec"))
}
