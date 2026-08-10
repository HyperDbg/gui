#pragma once
#include "pluginmain.h"
class IObserver
{
public:
	IObserver(void);
	~IObserver(void);
protected:
	duint mBase;
	duint mSize;
	StaticAnalysis *mObservable;

public:
	StaticAnalysis_Interface(StaticAnalysis *parent);
	virtual ~StaticAnalysis_Interface() {}

	// initialization before any analysis
	void initialise(const int_t Base,const int_t Size);
	// clear all extracted informations
	virtual void clear() = 0;
	// each derived class will see each instruction only once (in a analysis step)
	virtual void see(DISASM* disasm) = 0;
	// this methods process all gathered informations
	virtual bool think() = 0;
};

