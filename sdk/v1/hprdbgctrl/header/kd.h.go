package header

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\header\kd.h.back

type (
	Kd interface {
		____HKeyHolder() (ok bool) //col:11
	}
	kd struct{}
)

func NewKd() Kd { return &kd{} }

func (k *kd) ____HKeyHolder() (ok bool) { //col:11
	/*
		    HKeyHolder() :
		        m_Key(nullptr) { }
		    HKeyHolder(const HKeyHolder &) = delete;
		    HKeyHolder & operator=(const HKeyHolder &) = delete;
		    ~HKeyHolder()
		    {
		        if (m_Key != nullptr)
		            RegCloseKey(m_Key);
		    operator HKEY() const { return m_Key; }
		    HKEY * operator&() { return &m_Key; }
		};
	*/
	return true
}

