/* plugin:   (StaticAnalysis) for x64dbg <http://www.x64dbg.com>
 * author:   tr4ceflow@gmail.com <http://blog.traceflow.com>
 * license:  GLPv3
 */
#pragma once
#include "meta.h"
#include <list>
class ApiDB
{
private:
	bool mValid;  // "database" ok ?
public:
	ApiDB(void);
	~ApiDB(void);

	const bool ok() const;

	FunctionInfo_t find(std::string name);

	std::list<FunctionInfo_t> mInfo;

	
};

