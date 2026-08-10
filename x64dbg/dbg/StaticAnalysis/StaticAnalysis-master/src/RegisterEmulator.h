#pragma once
#include "meta.h"
class RegisterEmulator
{
	duint mRCX;
	duint mRDX;
	duint mR8;
	duint mR9;
public:
	RegisterEmulator();
	~RegisterEmulator();

	void emulate(const DISASM* BeaStruct);

	const duint rcx() const;
	const duint rdx() const;
	const duint r8() const;
	const duint r9() const;
};

