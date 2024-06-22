package hookApiMgr_test

import (
	"github.com/ddkwork/hyperdbgui/ux/plugin/scriptGen"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"github.com/ddkwork/golibrary/stream/clang"
)

func TestName(t *testing.T) {
	h := scriptGen.New()
	h.ProcessNameFilter("ssd.exe", "")
	clang.New().Format(`C:\Users\Admin\Desktop\gui\sdk\scriptGen\cpuid.ds`)
	// h := hookApiMgr.NewGeneratedFile()
	// h.TestIopXxxControlFile()
}

func TestName1(t *testing.T) {
	/*
	    	!cpuid script {
	   			 	printf("$pname\t%s\n", $pname);
	       		 	printf("rax,rbx,rcx,rdx: %llx\t%llx\t%llx\t%llx\t\n", @rax,@rbx,@rcx,@rdx);
	       		    if ($context == 1) {
	       		 	    @rax=0x000306A9;
	       		 		@rcx=0x3DBAE3BF;
	       		 		@rdx=0xBFEBFBFF;
	       		      }
	   		 	}

	*/
	/*
			 	if (
		db($pname  ) == 73 &&
		db($pname+1) == 73 &&
		db($pname+2) == 64 &&
		db($pname+3) == 2e &&
		db($pname+3) == 65 &&
		db($pname+3) == 78 &&
		db($pname+3) == 65
			) {

			    }
	*/
	ABCDE := stream.NewBuffer("ABCDE")
	ssd := stream.NewBuffer("ssd.exe")
	mylog.HexDump("ABCDE", ABCDE.Buffer.Bytes())
	mylog.HexDump("ssd.exe", ssd.Buffer.Bytes())
}
