package kdserial

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\ioaccess.c.back

type (
	Ioaccess interface {
		ReadPort8() (ok bool)       //col:5
		ReadPort16() (ok bool)      //col:10
		ReadPort32() (ok bool)      //col:15
		WritePort8() (ok bool)      //col:21
		WritePort16() (ok bool)     //col:27
		WritePort32() (ok bool)     //col:33
		ReadRegister8() (ok bool)   //col:38
		ReadRegister16() (ok bool)  //col:43
		ReadRegister32() (ok bool)  //col:48
		WriteRegister8() (ok bool)  //col:54
		WriteRegister16() (ok bool) //col:60
		WriteRegister32() (ok bool) //col:66
	}
	ioaccess struct{}
)

func NewIoaccess() Ioaccess { return &ioaccess{} }

func (i *ioaccess) ReadPort8() (ok bool) { //col:5
	/*ReadPort8(
	      PUCHAR Port)
	  {
	      return READ_PORT_UCHAR(Port);
	  }*/
	return true
}

func (i *ioaccess) ReadPort16() (ok bool) { //col:10
	/*ReadPort16(
	      PUSHORT Port)
	  {
	      return READ_PORT_USHORT(Port);
	  }*/
	return true
}

func (i *ioaccess) ReadPort32() (ok bool) { //col:15
	/*ReadPort32(
	      PULONG Port)
	  {
	      return READ_PORT_ULONG(Port);
	  }*/
	return true
}

func (i *ioaccess) WritePort8() (ok bool) { //col:21
	/*WritePort8(
	      PUCHAR      Port,
	      const UCHAR Value)
	  {
	      WRITE_PORT_UCHAR(Port, Value);
	  }*/
	return true
}

func (i *ioaccess) WritePort16() (ok bool) { //col:27
	/*WritePort16(
	      PUSHORT      Port,
	      const USHORT Value)
	  {
	      WRITE_PORT_USHORT(Port, Value);
	  }*/
	return true
}

func (i *ioaccess) WritePort32() (ok bool) { //col:33
	/*WritePort32(
	      PULONG      Port,
	      const ULONG Value)
	  {
	      WRITE_PORT_ULONG(Port, Value);
	  }*/
	return true
}

func (i *ioaccess) ReadRegister8() (ok bool) { //col:38
	/*ReadRegister8(
	      PUCHAR Register)
	  {
	      return READ_REGISTER_UCHAR(Register);
	  }*/
	return true
}

func (i *ioaccess) ReadRegister16() (ok bool) { //col:43
	/*ReadRegister16(
	      PUSHORT Register)
	  {
	      return READ_REGISTER_USHORT(Register);
	  }*/
	return true
}

func (i *ioaccess) ReadRegister32() (ok bool) { //col:48
	/*ReadRegister32(
	      PULONG Register)
	  {
	      return READ_REGISTER_ULONG(Register);
	  }*/
	return true
}

func (i *ioaccess) WriteRegister8() (ok bool) { //col:54
	/*WriteRegister8(
	      PUCHAR      Register,
	      const UCHAR Value)
	  {
	      WRITE_REGISTER_UCHAR(Register, Value);
	  }*/
	return true
}

func (i *ioaccess) WriteRegister16() (ok bool) { //col:60
	/*WriteRegister16(
	      PUSHORT      Register,
	      const USHORT Value)
	  {
	      WRITE_REGISTER_USHORT(Register, Value);
	  }*/
	return true
}

func (i *ioaccess) WriteRegister32() (ok bool) { //col:66
	/*WriteRegister32(
	      PULONG      Register,
	      const ULONG Value)
	  {
	      WRITE_REGISTER_ULONG(Register, Value);
	  }*/
	return true
}
