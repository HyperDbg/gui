package windef

import (
	os "os"
	unsafe "unsafe"
)

type DWORD = uint64
type BOOL = int32
type BYTE = uint8
type WORD = uint16
type FLOAT = float32
type PFLOAT = *float32
type INT = int32
type UINT = uint32
type PUINT = *uint32
type PBOOL = *int32
type LPBOOL = *int32
type PBYTE = *uint8
type LPBYTE = *uint8
type PINT = *int32
type LPINT = *int32
type PWORD = *uint16
type LPWORD = *uint16
type LPLONG = *int64
type PDWORD = *uint64
type LPDWORD = *uint64
type LPVOID = unsafe.Pointer
type LPCVOID = unsafe.Pointer
type ULONG = uint64
type PULONG = *uint64
type USHORT = uint16
type PUSHORT = *uint16
type UCHAR = uint8
type PUCHAR = *uint8
type CHAR = int8
type SHORT = int16
type LONG = int64

var BuildDateTime [20]uint8 = [20]uint8{uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(7))))), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(8))))), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(9))))), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(10))))), uint8('-'), uint8(func() int32 {
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'O' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'N' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'D' {
		return '1'
	} else {
		return '0'
	}
}()), uint8(func() int32 {
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'J' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'a' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'n' {
		return '1'
	} else {
		return func() int32 {
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'F' {
				return '2'
			} else {
				return func() int32 {
					if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'M' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'a' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'r' {
						return '3'
					} else {
						return func() int32 {
							if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'A' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'p' {
								return '4'
							} else {
								return func() int32 {
									if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'M' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'a' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'y' {
										return '5'
									} else {
										return func() int32 {
											if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'J' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'u' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'n' {
												return '6'
											} else {
												return func() int32 {
													if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'J' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'u' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'l' {
														return '7'
													} else {
														return func() int32 {
															if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'A' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'u' {
																return '8'
															} else {
																return func() int32 {
																	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'S' {
																		return '9'
																	} else {
																		return func() int32 {
																			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'O' {
																				return '0'
																			} else {
																				return func() int32 {
																					if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'N' {
																						return '1'
																					} else {
																						return func() int32 {
																							if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'D' {
																								return '2'
																							} else {
																								return '?'
																							}
																						}()
																					}
																				}()
																			}
																		}()
																	}
																}()
															}
														}()
													}
												}()
											}
										}()
									}
								}()
							}
						}()
					}
				}()
			}
		}()
	}
}()), uint8('-'), uint8(func() int32 {
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(4))))) >= '0' {
		return int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(4)))))
	} else {
		return '0'
	}
}()), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(5))))), uint8(' '), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'0', '6', ':', '3', '2', ':', '5', '0', '\x00'})))) + uintptr(int32(0))))), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'0', '6', ':', '3', '2', ':', '5', '0', '\x00'})))) + uintptr(int32(1))))), uint8(':'), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'0', '6', ':', '3', '2', ':', '5', '0', '\x00'})))) + uintptr(int32(3))))), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'0', '6', ':', '3', '2', ':', '5', '0', '\x00'})))) + uintptr(int32(4))))), uint8(':'), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'0', '6', ':', '3', '2', ':', '5', '0', '\x00'})))) + uintptr(int32(6))))), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'0', '6', ':', '3', '2', ':', '5', '0', '\x00'})))) + uintptr(int32(7))))), uint8('\x00')}
var CompleteVersion [7]uint8 = [7]uint8{uint8('v'), uint8(48), uint8('.'), uint8(50), uint8('.'), uint8(48), uint8('\x00')}
var BuildVersion [9]uint8 = [9]uint8{uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(7))))), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(8))))), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(9))))), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(10))))), uint8(func() int32 {
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'O' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'N' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'D' {
		return '1'
	} else {
		return '0'
	}
}()), uint8(func() int32 {
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'J' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'a' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'n' {
		return '1'
	} else {
		return func() int32 {
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'F' {
				return '2'
			} else {
				return func() int32 {
					if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'M' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'a' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'r' {
						return '3'
					} else {
						return func() int32 {
							if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'A' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'p' {
								return '4'
							} else {
								return func() int32 {
									if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'M' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'a' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'y' {
										return '5'
									} else {
										return func() int32 {
											if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'J' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'u' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'n' {
												return '6'
											} else {
												return func() int32 {
													if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'J' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'u' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'l' {
														return '7'
													} else {
														return func() int32 {
															if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'A' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'u' {
																return '8'
															} else {
																return func() int32 {
																	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'S' {
																		return '9'
																	} else {
																		return func() int32 {
																			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'O' {
																				return '0'
																			} else {
																				return func() int32 {
																					if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'N' {
																						return '1'
																					} else {
																						return func() int32 {
																							if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'D' {
																								return '2'
																							} else {
																								return '?'
																							}
																						}()
																					}
																				}()
																			}
																		}()
																	}
																}()
															}
														}()
													}
												}()
											}
										}()
									}
								}()
							}
						}()
					}
				}()
			}
		}()
	}
}()), uint8(func() int32 {
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(4))))) >= '0' {
		return int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(4)))))
	} else {
		return '0'
	}
}()), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(5))))), uint8('\x00')}

func _cgo_main() int32 {
	return int32(0)
}
func main() {
	os.Exit(int(_cgo_main()))
}
