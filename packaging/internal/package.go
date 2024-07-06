package internal

// Package performs the platform-specific packaging for GCS.
func Package() {
	Configure()
	loadBaseImages()
	platformPackage()
}

func Configure() {
	// AppName = "HyperDbg"
	// cmdline.AppCmdName = "hyperdbg"
	// cmdline.License = "Mozilla Public License, version 2.0"
	// cmdline.CopyrightStartYear = "1998"
	// cmdline.CopyrightHolder = "ddkwork"
	// cmdline.AppIdentifier = "com.hyperdbg"
}
