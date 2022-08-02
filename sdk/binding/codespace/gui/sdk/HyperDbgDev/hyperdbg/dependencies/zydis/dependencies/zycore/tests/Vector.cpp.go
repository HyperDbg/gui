package tests

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\tests\Vector.cpp.back

type (
	Vector interface {
		____void_SetUp() (ok bool)                //col:28
		static_ZyanStatus_FreeZyanU16() (ok bool) //col:33
	}
	vector struct{}
)

func NewVector() Vector { return &vector{} }

func (v *vector) ____void_SetUp() (ok bool) { //col:28
	/*
	   	void SetUp() override
	   	{
	   	    m_has_fixed_capacity = GetParam();
	   	    if (!m_has_fixed_capacity)
	   	    {
	   	        ASSERT_EQ(ZyanVectorInit(&m_vector, sizeof(ZyanU64), m_test_size,
	   	            reinterpret_cast<ZyanMemberProcedure>(ZYAN_NULL)), ZYAN_STATUS_SUCCESS);
	   	        m_buffer.reserve(m_test_size);
	   	        ASSERT_EQ(ZyanVectorInitCustomBuffer(&m_vector, sizeof(ZyanU64), m_buffer.data(),
	   	            m_test_size, reinterpret_cast<ZyanMemberProcedure>(ZYAN_NULL)),
	   	            ZYAN_STATUS_SUCCESS);
	   	void TearDown() override
	   	{
	   	    EXPECT_EQ(ZyanVectorDestroy(&m_vector), ZYAN_STATUS_SUCCESS);
	   	void SetUp() override
	   	{
	   	    VectorTestBase::SetUp();
	   	    if (m_has_fixed_capacity)
	   	    {
	   	        m_buffer.resize(m_test_size);
	   	    for (ZyanU64 i = 0; i < m_test_size; ++i)
	   	    {
	   	        ASSERT_EQ(ZyanVectorPushBack(&m_vector, &i), ZYAN_STATUS_SUCCESS);

	   static ZyanStatus InitZyanU64(ZyanU64* object)

	   	{
	   	    *object = 1337;
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (v *vector) static_ZyanStatus_FreeZyanU16() (ok bool) { //col:33
	/*
	   static ZyanStatus FreeZyanU16(ZyanU16* object)

	   	{
	   	    *object = 0;
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

