package src

type (
	List interface {
		static_ZyanStatus_ZyanListAllocateNode() (ok bool) //col:54
		ZyanStatus_ZyanListInitCustomBuffer() (ok bool)    //col:72
		ZyanStatus_ZyanListDestroy() (ok bool)             //col:99
		ZyanStatus_ZyanListGetTailNode() (ok bool)         //col:108
		ZyanStatus_ZyanListGetPrevNode() (ok bool)         //col:117
		ZyanStatus_ZyanListGetNextNode() (ok bool)         //col:126
		const_voidPtr_ZyanListGetNodeData() (ok bool)      //col:186
		ZyanStatus_ZyanListPushFront() (ok bool)           //col:207
		ZyanStatus_ZyanListEmplaceBack() (ok bool)         //col:231
		ZyanStatus_ZyanListEmplaceFront() (ok bool)        //col:255
		ZyanStatus_ZyanListPopBack() (ok bool)             //col:383
	}
	list struct{}
)

func NewList() List { return &list{} }

func (l *list) static_ZyanStatus_ZyanListAllocateNode() (ok bool) { //col:54

	return true
}

func (l *list) ZyanStatus_ZyanListInitCustomBuffer() (ok bool) { //col:72

	return true
}

func (l *list) ZyanStatus_ZyanListDestroy() (ok bool) { //col:99

	return true
}

func (l *list) ZyanStatus_ZyanListGetTailNode() (ok bool) { //col:108

	return true
}

func (l *list) ZyanStatus_ZyanListGetPrevNode() (ok bool) { //col:117

	return true
}

func (l *list) ZyanStatus_ZyanListGetNextNode() (ok bool) { //col:126

	return true
}

func (l *list) const_voidPtr_ZyanListGetNodeData() (ok bool) { //col:186

	return true
}

func (l *list) ZyanStatus_ZyanListPushFront() (ok bool) { //col:207

	return true
}

func (l *list) ZyanStatus_ZyanListEmplaceBack() (ok bool) { //col:231

	return true
}

func (l *list) ZyanStatus_ZyanListEmplaceFront() (ok bool) { //col:255

	return true
}

func (l *list) ZyanStatus_ZyanListPopBack() (ok bool) { //col:383

	return true
}
