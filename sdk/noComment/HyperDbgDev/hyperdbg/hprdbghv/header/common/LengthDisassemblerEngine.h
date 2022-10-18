#pragma once
#ifndef LDISASM_H
#define LDISASM_H
#define LDISASM_R                                                              \
  (*b >>                                                                       \
   4) 
#define LDISASM_C                                                              \
  (*b & 0xF) 
static const UINT8 legacyPrefixes[] = {0xF0, 0xF2, 0xF3, 0x2E, 0x36, 0x3E,
                                       0x26, 0x64, 0x65, 0x66, 0x67};
static const UINT8 op1modrm[] = {0x62, 0x63, 0x69, 0x6B, 0xC0, 0xC1,
                                 0xC4, 0xC5, 0xC6, 0xC7, 0xD0, 0xD1,
                                 0xD2, 0xD3, 0xF6, 0xF7, 0xFE, 0xFF};
static const UINT8 op1imm8[] = {0x6A, 0x6B, 0x80, 0x82, 0x83, 0xA8, 0xC0,
                                0xC1, 0xC6, 0xCD, 0xD4, 0xD5, 0xEB};
static const UINT8 op1imm32[] = {0x68, 0x69, 0x81, 0xA9, 0xC7, 0xE8, 0xE9};
static const UINT8 op2modrm[] = {0x0D, 0xA3, 0xA4, 0xA5, 0xAB,
                                 0xAC, 0xAD, 0xAE, 0xAF};
inline BOOLEAN findByte(const UINT8 *arr, const size_t N, const UINT8 x) {
  for (size_t i = 0; i < N; i++) {
    if (arr[i] == x) {
      return TRUE;
    }
  };
  return FALSE;
}

inline void parseModRM(UINT8 **b, const BOOLEAN addressPrefix) {
  UINT8 modrm = *++*b;
  if (!addressPrefix || **b >= 0x40) {
    BOOLEAN hasSIB = FALSE; 
    if (**b < 0xC0 && (**b & 0b111) == 0b100 && !addressPrefix)
      hasSIB = TRUE, (*b)++;
    if (modrm >= 0x40 && modrm <= 0x7F) 
      (*b)++;
    else if ((modrm <= 0x3F && (modrm & 0b111) == 0b101) ||
             (modrm >= 0x80 && modrm <= 0xBF)) 
      *b += (addressPrefix) ? 2 : 4;
    else if (hasSIB && (**b & 0b111) == 0b101) 
      *b += (modrm & 0b01000000) ? 1 : 4;
  } else if (addressPrefix && modrm == 0x26)
    *b += 2;
};
inline size_t ldisasm(const void *const address, const BOOLEAN x86_64_mode) {
  size_t offset = 0;
  BOOLEAN operandPrefix = FALSE, addressPrefix = FALSE, rexW = FALSE;
  UINT8 *b = (UINT8 *)(address);
  for (int i = 0;
       i < 14 && findByte(legacyPrefixes, sizeof(legacyPrefixes), *b) ||
       ((x86_64_mode) ? (LDISASM_R == 4) : FALSE);
       i++, b++) {
    if (*b == 0x66)
      operandPrefix = TRUE;
    else if (*b == 0x67)
      addressPrefix = TRUE;
    else if (LDISASM_R == 4 && LDISASM_C >= 8)
      rexW = TRUE;
  }
  if (*b == 0x0F) 
  {
    b++;
    if (*b == 0x38 || *b == 0x3A) 
    {
      if (*b++ == 0x3A)
        offset++;
      parseModRM(&b, addressPrefix);
    } else 
    {
      if (LDISASM_R == 8) 
        offset += 4;
      else if ((LDISASM_R == 7 && LDISASM_C < 4) || *b == 0xA4 || *b == 0xC2 ||
               (*b > 0xC3 && *b <= 0xC6) || *b == 0xBA || *b == 0xAC) 
        offset++;
      if (findByte(op2modrm, sizeof(op2modrm), *b) ||
          (LDISASM_R != 3 && LDISASM_R > 0 && LDISASM_R < 7) || *b >= 0xD0 ||
          (LDISASM_R == 7 && LDISASM_C != 7) || LDISASM_R == 9 ||
          LDISASM_R == 0xB || (LDISASM_R == 0xC && LDISASM_C < 8) ||
          (LDISASM_R == 0 && LDISASM_C < 4))
        parseModRM(&b, addressPrefix);
    }
  } else 
  {
    if ((LDISASM_R == 0xE && LDISASM_C < 8) ||
        (LDISASM_R == 0xB && LDISASM_C < 8) || LDISASM_R == 7 ||
        (LDISASM_R < 4 && (LDISASM_C == 4 || LDISASM_C == 0xC)) ||
        (*b == 0xF6 && !(*(b + 1) & 48)) ||
        findByte(op1imm8, sizeof(op1imm8), *b)) 
      offset++;
    else if (*b == 0xC2 || *b == 0xCA) 
      offset += 2;
    else if (*b == 0xC8) 
      offset += 3;
    else if ((LDISASM_R < 4 && (LDISASM_C == 5 || LDISASM_C == 0xD)) ||
             (LDISASM_R == 0xB && LDISASM_C >= 8) ||
             (*b == 0xF7 && !(*(b + 1) & 48)) ||
             findByte(op1imm32, sizeof(op1imm32), *b)) 
      offset += (rexW) ? 8 : (operandPrefix ? 2 : 4);
    else if (LDISASM_R == 0xA && LDISASM_C < 4)
      offset += (rexW) ? 8 : (addressPrefix ? 2 : 4);
    else if (*b == 0xEA || *b == 0x9A) 
      offset += operandPrefix ? 4 : 6;
    if (findByte(op1modrm, sizeof(op1modrm), *b) ||
        (LDISASM_R < 4 &&
         (LDISASM_C < 4 || (LDISASM_C >= 8 && LDISASM_C < 0xC))) ||
        LDISASM_R == 8 || (LDISASM_R == 0xD && LDISASM_C >= 8))
      parseModRM(&b, addressPrefix);
  }
  return (size_t)((ptrdiff_t)(++b + offset) - (ptrdiff_t)(address));
}

#endif 
