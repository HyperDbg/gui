package src

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\src\String.c.back

const (
	ZYDIS_MAXCHARS_DEC_32 = 10 //col:1
	ZYDIS_MAXCHARS_DEC_64 = 20 //col:2
	ZYDIS_MAXCHARS_HEX_32 = 8  //col:3
	ZYDIS_MAXCHARS_HEX_64 = 16 //col:4
)

type (
	String interface {
		ZyanStatus_ZydisStringAppendDecU32() (ok bool) //col:36
		ZyanStatus_ZydisStringAppendDecU64() (ok bool) //col:72
		ZyanStatus_ZydisStringAppendHexU32() (ok bool) //col:130
		ZyanStatus_ZydisStringAppendHexU64() (ok bool) //col:189
		ZyanStatus_ZydisStringAppendDecU() (ok bool)   //col:211
		ZyanStatus_ZydisStringAppendHexU() (ok bool)   //col:236
	}
	string struct{}
)

func NewString() String { return &string{} }

func (s *string) ZyanStatus_ZydisStringAppendDecU32() (ok bool) { //col:36
	/*ZyanStatus ZydisStringAppendDecU32(ZyanString* string, ZyanU32 value, ZyanU8 padding_length)
	  {
	      ZYAN_ASSERT(string);
	      ZYAN_ASSERT(!string->vector.allocator);
	      char buffer[ZYDIS_MAXCHARS_DEC_32];
	      char *buffer_end = &buffer[ZYDIS_MAXCHARS_DEC_32];
	      char *buffer_write_pointer = buffer_end;
	      while (value >= 100)
	      {
	          const ZyanU32 value_old = value;
	          buffer_write_pointer -= 2;
	          value /= 100;
	          ZYAN_MEMCPY(buffer_write_pointer, &DECIMAL_LOOKUP[(value_old - (value * 100)) * 2], 2);
	      }
	      buffer_write_pointer -= 2;
	      ZYAN_MEMCPY(buffer_write_pointer, &DECIMAL_LOOKUP[value * 2], 2);
	      const ZyanUSize offset_odd    = (ZyanUSize)(value < 10);
	      const ZyanUSize length_number = buffer_end - buffer_write_pointer - offset_odd;
	      const ZyanUSize length_total  = ZYAN_MAX(length_number, padding_length);
	      const ZyanUSize length_target = string->vector.size;
	      if (string->vector.size + length_total > string->vector.capacity)
	      {
	          return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
	      }
	      ZyanUSize offset_write = 0;
	      if (padding_length > length_number)
	      {
	          offset_write = padding_length - length_number;
	          ZYAN_MEMSET((char*)string->vector.data + length_target - 1, '0', offset_write);
	      }
	      ZYAN_MEMCPY((char*)string->vector.data + length_target + offset_write - 1,
	          buffer_write_pointer + offset_odd, length_number);
	      string->vector.size = length_target + length_total;
	      ZYDIS_STRING_NULLTERMINATE(string);
	      return ZYAN_STATUS_SUCCESS;
	  }*/
	return true
}

func (s *string) ZyanStatus_ZydisStringAppendDecU64() (ok bool) { //col:72
	/*ZyanStatus ZydisStringAppendDecU64(ZyanString* string, ZyanU64 value, ZyanU8 padding_length)
	  {
	      ZYAN_ASSERT(string);
	      ZYAN_ASSERT(!string->vector.allocator);
	      char buffer[ZYDIS_MAXCHARS_DEC_64];
	      char *buffer_end = &buffer[ZYDIS_MAXCHARS_DEC_64];
	      char *buffer_write_pointer = buffer_end;
	      while (value >= 100)
	      {
	          const ZyanU64 value_old = value;
	          buffer_write_pointer -= 2;
	          value /= 100;
	          ZYAN_MEMCPY(buffer_write_pointer, &DECIMAL_LOOKUP[(value_old - (value * 100)) * 2], 2);
	      }
	      buffer_write_pointer -= 2;
	      ZYAN_MEMCPY(buffer_write_pointer, &DECIMAL_LOOKUP[value * 2], 2);
	      const ZyanUSize offset_odd    = (ZyanUSize)(value < 10);
	      const ZyanUSize length_number = buffer_end - buffer_write_pointer - offset_odd;
	      const ZyanUSize length_total  = ZYAN_MAX(length_number, padding_length);
	      const ZyanUSize length_target = string->vector.size;
	      if (string->vector.size + length_total > string->vector.capacity)
	      {
	          return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
	      }
	      ZyanUSize offset_write = 0;
	      if (padding_length > length_number)
	      {
	          offset_write = padding_length - length_number;
	          ZYAN_MEMSET((char*)string->vector.data + length_target - 1, '0', offset_write);
	      }
	      ZYAN_MEMCPY((char*)string->vector.data + length_target + offset_write - 1,
	          buffer_write_pointer + offset_odd, length_number);
	      string->vector.size = length_target + length_total;
	      ZYDIS_STRING_NULLTERMINATE(string);
	      return ZYAN_STATUS_SUCCESS;
	  }*/
	return true
}

func (s *string) ZyanStatus_ZydisStringAppendHexU32() (ok bool) { //col:130
	/*ZyanStatus ZydisStringAppendHexU32(ZyanString* string, ZyanU32 value, ZyanU8 padding_length,
	      ZyanBool uppercase)
	  {
	      ZYAN_ASSERT(string);
	      ZYAN_ASSERT(!string->vector.allocator);
	      const ZyanUSize len = string->vector.size;
	      const ZyanUSize remaining = string->vector.capacity - string->vector.size;
	      if (remaining < (ZyanUSize)padding_length)
	      {
	          return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
	      }
	      if (!value)
	      {
	          const ZyanU8 n = (padding_length ? padding_length : 1);
	          if (remaining < (ZyanUSize)n)
	          {
	              return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
	          }
	          ZYAN_MEMSET((char*)string->vector.data + len - 1, '0', n);
	          string->vector.size = len + n;
	          ZYDIS_STRING_NULLTERMINATE(string);
	          return ZYAN_STATUS_SUCCESS;
	      }
	      ZyanU8 n = 0;
	      char* buffer = ZYAN_NULL;
	      for (ZyanI8 i = ZYDIS_MAXCHARS_HEX_32 - 1; i >= 0; --i)
	      {
	          const ZyanU8 v = (value >> i * 4) & 0x0F;
	          if (!n)
	          {
	              if (!v)
	              {
	                  continue;
	              }
	              if (remaining <= (ZyanU8)i)
	              {
	                  return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
	              }
	              buffer = (char*)string->vector.data + len - 1;
	              if (padding_length > i)
	              {
	                  n = padding_length - i - 1;
	                  ZYAN_MEMSET(buffer, '0', n);
	              }
	          }
	          ZYAN_ASSERT(buffer);
	          if (uppercase)
	          {
	              buffer[n++] = "0123456789ABCDEF"[v];
	          } else
	          {
	              buffer[n++] = "0123456789abcdef"[v];
	          }
	      }
	      string->vector.size = len + n;
	      ZYDIS_STRING_NULLTERMINATE(string);
	      return ZYAN_STATUS_SUCCESS;
	  }*/
	return true
}

func (s *string) ZyanStatus_ZydisStringAppendHexU64() (ok bool) { //col:189
	/*ZyanStatus ZydisStringAppendHexU64(ZyanString* string, ZyanU64 value, ZyanU8 padding_length,
	      ZyanBool uppercase)
	  {
	      ZYAN_ASSERT(string);
	      ZYAN_ASSERT(!string->vector.allocator);
	      const ZyanUSize len = string->vector.size;
	      const ZyanUSize remaining = string->vector.capacity - string->vector.size;
	      if (remaining < (ZyanUSize)padding_length)
	      {
	          return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
	      }
	      if (!value)
	      {
	          const ZyanU8 n = (padding_length ? padding_length : 1);
	          if (remaining < (ZyanUSize)n)
	          {
	              return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
	          }
	          ZYAN_MEMSET((char*)string->vector.data + len - 1, '0', n);
	          string->vector.size = len + n;
	          ZYDIS_STRING_NULLTERMINATE(string);
	          return ZYAN_STATUS_SUCCESS;
	      }
	      ZyanU8 n = 0;
	      char* buffer = ZYAN_NULL;
	      for (ZyanI8 i = ((value & 0xFFFFFFFF00000000) ?
	          ZYDIS_MAXCHARS_HEX_64 : ZYDIS_MAXCHARS_HEX_32) - 1; i >= 0; --i)
	      {
	          const ZyanU8 v = (value >> i * 4) & 0x0F;
	          if (!n)
	          {
	              if (!v)
	              {
	                  continue;
	              }
	              if (remaining <= (ZyanU8)i)
	              {
	                  return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
	              }
	              buffer = (char*)string->vector.data + len - 1;
	              if (padding_length > i)
	              {
	                  n = padding_length - i - 1;
	                  ZYAN_MEMSET(buffer, '0', n);
	              }
	          }
	          ZYAN_ASSERT(buffer);
	          if (uppercase)
	          {
	              buffer[n++] = "0123456789ABCDEF"[v];
	          } else
	          {
	              buffer[n++] = "0123456789abcdef"[v];
	          }
	      }
	      string->vector.size = len + n;
	      ZYDIS_STRING_NULLTERMINATE(string);
	      return ZYAN_STATUS_SUCCESS;
	  }*/
	return true
}

func (s *string) ZyanStatus_ZydisStringAppendDecU() (ok bool) { //col:211
	/*ZyanStatus ZydisStringAppendDecU(ZyanString* string, ZyanU64 value, ZyanU8 padding_length,
	      const ZyanStringView* prefix, const ZyanStringView* suffix)
	  {
	      if (prefix)
	      {
	          ZYAN_CHECK(ZydisStringAppend(string, prefix));
	      }
	  #if defined(ZYAN_X64) || defined(ZYAN_AARCH64)
	      ZYAN_CHECK(ZydisStringAppendDecU64(string, value, padding_length));
	  #else
	      if (value & 0xFFFFFFFF00000000)
	      {
	          ZYAN_CHECK(ZydisStringAppendDecU64(string, value, padding_length));
	      }
	      ZYAN_CHECK(ZydisStringAppendDecU32(string, (ZyanU32)value, padding_length));
	  #endif
	      if (suffix)
	      {
	          return ZydisStringAppend(string, suffix);
	      }
	      return ZYAN_STATUS_SUCCESS;
	  }*/
	return true
}

func (s *string) ZyanStatus_ZydisStringAppendHexU() (ok bool) { //col:236
	/*ZyanStatus ZydisStringAppendHexU(ZyanString* string, ZyanU64 value, ZyanU8 padding_length,
	      ZyanBool uppercase, const ZyanStringView* prefix, const ZyanStringView* suffix)
	  {
	      if (prefix)
	      {
	          ZYAN_CHECK(ZydisStringAppend(string, prefix));
	      }
	  #if defined(ZYAN_X64) || defined(ZYAN_AARCH64)
	      ZYAN_CHECK(ZydisStringAppendHexU64(string, value, padding_length, uppercase));
	  #else
	      if (value & 0xFFFFFFFF00000000)
	      {
	          ZYAN_CHECK(ZydisStringAppendHexU64(string, value, padding_length, uppercase));
	      }
	      else
	      {
	          ZYAN_CHECK(ZydisStringAppendHexU32(string, (ZyanU32)value, padding_length, uppercase));
	      }
	  #endif
	      if (suffix)
	      {
	          return ZydisStringAppend(string, suffix);
	      }
	      return ZYAN_STATUS_SUCCESS;
	  }*/
	return true
}
