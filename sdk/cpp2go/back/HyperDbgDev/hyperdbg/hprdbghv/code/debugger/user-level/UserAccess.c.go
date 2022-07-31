package user-level
type (
	UserAccess interface {
		UserAccessAllocateAndGetImagePathFromProcessId() (ok bool)              //col:161
		UserAccessGetPebFromProcessId() (ok bool)                               //col:231
		UserAccessGetBaseAndEntrypointOfMainModuleIfLoadedInVmxRoot() (ok bool) //col:324
		UserAccessPrintLoadedModulesX64() (ok bool)                             //col:474
		UserAccessPrintLoadedModulesX86() (ok bool)                             //col:616
		UserAccessPrintLoadedModulesX86_2() (ok bool)                           //col:699
		UserAccessIsWow64Process() (ok bool)                                    //col:754
		UserAccessGetLoadedModules() (ok bool)                                  //col:826
		UserAccessCheckForLoadedModuleDetails() (ok bool)                       //col:882
	}
)
func NewUserAccess() { return &userAccess{} }
func (u *userAccess) UserAccessAllocateAndGetImagePathFromProcessId() (ok bool) { //col:161
	return true
}

func (u *userAccess) UserAccessGetPebFromProcessId() (ok bool) { //col:231
	return true
}

func (u *userAccess) UserAccessGetBaseAndEntrypointOfMainModuleIfLoadedInVmxRoot() (ok bool) { //col:324
	return true
}

func (u *userAccess) UserAccessPrintLoadedModulesX64() (ok bool) { //col:474
	return true
}

func (u *userAccess) UserAccessPrintLoadedModulesX86() (ok bool) { //col:616
	return true
}

func (u *userAccess) UserAccessPrintLoadedModulesX86_2() (ok bool) { //col:699
	return true
}

func (u *userAccess) UserAccessIsWow64Process() (ok bool) { //col:754
	return true
}

func (u *userAccess) UserAccessGetLoadedModules() (ok bool) { //col:826
	return true
}

func (u *userAccess) UserAccessCheckForLoadedModuleDetails() (ok bool) { //col:882
	return true
}

