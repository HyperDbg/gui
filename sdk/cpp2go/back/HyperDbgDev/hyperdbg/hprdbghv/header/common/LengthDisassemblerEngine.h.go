package common

//back\HyperDbgDev\hyperdbg\hprdbghv\header\common\LengthDisassemblerEngine.h.back

type (
	LengthDisassemblerEngine interface {
		findByte() (ok bool)   //col:39
		parseModRM() (ok bool) //col:61
		ldisasm() (ok bool)    //col:126
	}
)

func NewLengthDisassemblerEngine() { return &lengthDisassemblerEngine{} }

func (l *lengthDisassemblerEngine) findByte() (ok bool) { //col:39
	/*findByte(const UINT8 * arr, const size_t N, const UINT8 x)
	  {
	      for (size_t i = 0; i < N; i++)
	      {
	          if (arr[i] == x)
	          {
	              return TRUE;
	          }
	      };
	      return FALSE;
	  }*/
	return true
}

func (l *lengthDisassemblerEngine) parseModRM() (ok bool) { //col:61
	/*parseModRM(UINT8 ** b, const BOOLEAN addressPrefix)
	  {
	      UINT8 modrm = *++*b;
	      if (!addressPrefix || **b >= 0x40)
	      {
	          if (**b < 0xC0 && (**b & 0b111) == 0b100 && !addressPrefix)
	              hasSIB = TRUE, (*b)++;
	              (*b)++;
	              *b += (addressPrefix) ? 2 : 4;
	              *b += (modrm & 0b01000000) ? 1 : 4;
	      }
	      else if (addressPrefix && modrm == 0x26)
	          *b += 2;
	  };*/
	return true
}

func (l *lengthDisassemblerEngine) ldisasm() (ok bool) { //col:126
	/*ldisasm(const void * const address, const BOOLEAN x86_64_mode)
	  {
	      size_t  offset        = 0;
	      BOOLEAN operandPrefix = FALSE, addressPrefix = FALSE, rexW = FALSE;
	      UINT8 * b = (UINT8 *)(address);
	      for (int i = 0; i < 14 && findByte(legacyPrefixes, sizeof(legacyPrefixes), *b) || ((x86_64_mode) ? (LDISASM_R == 4) : FALSE); i++, b++)
	      {
	          if (*b == 0x66)
	              operandPrefix = TRUE;
	          else if (*b == 0x67)
	              addressPrefix = TRUE;
	          else if (LDISASM_R == 4 && LDISASM_C >= 8)
	              rexW = TRUE;
	      }
	      {
	          b++;
	          {
	              if (*b++ == 0x3A)
	                  offset++;
	              parseModRM(&b, addressPrefix);
	          }
	          {
	                  offset += 4;
	                  offset++;
	              if (findByte(op2modrm, sizeof(op2modrm), *b) || (LDISASM_R != 3 && LDISASM_R > 0 && LDISASM_R < 7) || *b >= 0xD0 || (LDISASM_R == 7 && LDISASM_C != 7) || LDISASM_R == 9 || LDISASM_R == 0xB || (LDISASM_R == 0xC && LDISASM_C < 8) || (LDISASM_R == 0 && LDISASM_C < 4))
	                  parseModRM(&b, addressPrefix);
	          }
	      }
	      {
	              offset++;
	              offset += 2;
	              offset += 3;
	              offset += (rexW) ? 8 : (operandPrefix ? 2 : 4);
	          else if (LDISASM_R == 0xA && LDISASM_C < 4)
	              offset += (rexW) ? 8 : (addressPrefix ? 2 : 4);
	              offset += operandPrefix ? 4 : 6;
	          if (findByte(op1modrm, sizeof(op1modrm), *b) || (LDISASM_R < 4 && (LDISASM_C < 4 || (LDISASM_C >= 8 && LDISASM_C < 0xC))) || LDISASM_R == 8 || (LDISASM_R == 0xD && LDISASM_C >= 8))
	              parseModRM(&b, addressPrefix);
	      }
	      return (size_t)((ptrdiff_t)(++b + offset) - (ptrdiff_t)(address));
	  }*/
	return true
}
