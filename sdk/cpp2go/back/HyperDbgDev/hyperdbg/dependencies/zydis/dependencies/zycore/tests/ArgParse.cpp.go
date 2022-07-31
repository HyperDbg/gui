package tests
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\tests\ArgParse.cpp.back

type (
ArgParse interface{
  Zyan Core Library ()(ok bool)//col:50
static auto UnnamedArgTest()(ok bool)//col:81
TEST()(ok bool)//col:88
TEST()(ok bool)//col:95
TEST()(ok bool)//col:115
TEST()(ok bool)//col:173
TEST()(ok bool)//col:220
TEST()(ok bool)//col:255
TEST()(ok bool)//col:308
int main()(ok bool)//col:318
}
)

func NewArgParse() { return & argParse{} }

func (a *argParse)  Zyan Core Library ()(ok bool){//col:50
/*  Zyan Core Library (Zycore-C)
  Original Author : Joel Hoener
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
#include <string_view>
#include <gtest/gtest.h>
#include <Zycore/ArgParse.h>
#include <Zycore/LibC.h>
auto cvt_string_view(const ZyanStringView *sv)
{
    const char* buf;
    if (ZYAN_FAILED(ZyanStringViewGetData(sv, &buf))) throw std::exception{};
    ZyanUSize len;
    if (ZYAN_FAILED(ZyanStringViewGetSize(sv, &len))) throw std::exception{};
    return std::string_view{buf, len};
}*/
return true
}

func (a *argParse)static auto UnnamedArgTest()(ok bool){//col:81
/*static auto UnnamedArgTest(ZyanU64 min, ZyanU64 max)
{
    const char* argv[]
    {
        "./test", "a", "xxx"
    };
    ZyanArgParseConfig cfg
    {
    };
    ZyanVector parsed;
    const char* err_tok = nullptr;
    ZYAN_MEMSET(&parsed, 0, sizeof(parsed));
    auto status = ZyanArgParse(&cfg, &parsed, &err_tok);
    return std::make_tuple(status, parsed, err_tok);
}*/
return true
}

func (a *argParse)TEST()(ok bool){//col:88
/*TEST(UnnamedArgs, TooFew)
{
    auto [status, parsed, err_tok] = UnnamedArgTest(5, 5);
    ASSERT_EQ(status, ZYAN_STATUS_TOO_FEW_ARGS);
    ASSERT_STREQ(err_tok, nullptr);
}*/
return true
}

func (a *argParse)TEST()(ok bool){//col:95
/*TEST(UnnamedArgs, TooMany)
{
    auto [status, parsed, err_tok] = UnnamedArgTest(1, 1);
    ASSERT_EQ(status, ZYAN_STATUS_TOO_MANY_ARGS);
    ASSERT_STREQ(err_tok, "xxx");
}*/
return true
}

func (a *argParse)TEST()(ok bool){//col:115
/*TEST(UnnamedArgs, PerfectFit)
{
    auto [status, parsed, err_tok] = UnnamedArgTest(2, 2);
    ASSERT_TRUE(ZYAN_SUCCESS(status));
    ZyanUSize size;
    ASSERT_TRUE(ZYAN_SUCCESS(ZyanVectorGetSize(&parsed, &size)));
    ASSERT_EQ(size, 2);
    auto arg = (const ZyanArgParseArg*)ZyanVectorGet(&parsed, 0);
    ASSERT_NE(arg, nullptr);
    ASSERT_TRUE(arg->has_value);
    ASSERT_EQ(cvt_string_view(&arg->value), "a");
    arg = (const ZyanArgParseArg*)ZyanVectorGet(&parsed, 1);
    ASSERT_NE(arg, nullptr);
    ASSERT_TRUE(arg->has_value);
    ASSERT_EQ(cvt_string_view(&arg->value), "xxx");
}*/
return true
}

func (a *argParse)TEST()(ok bool){//col:173
/*TEST(DashArg, MixedBoolAndValueArgs)
{
    const char* argv[]
    {
        "./test", "-aio42", "-n", "xxx"
    };
    ZyanArgParseDefinition args[]
    {
        {"-o", ZYAN_FALSE, ZYAN_FALSE},
        {"-a", ZYAN_TRUE, ZYAN_FALSE},
        {"-n", ZYAN_FALSE, ZYAN_FALSE},
        {"-i", ZYAN_TRUE, ZYAN_FALSE},
        {nullptr, ZYAN_FALSE, ZYAN_FALSE}
    };
    ZyanArgParseConfig cfg
    {
    };
    ZyanVector parsed;
    ZYAN_MEMSET(&parsed, 0, sizeof(parsed));
    auto status = ZyanArgParse(&cfg, &parsed, nullptr);
    ASSERT_TRUE(ZYAN_SUCCESS(status));
    ZyanUSize size;
    ASSERT_TRUE(ZYAN_SUCCESS(ZyanVectorGetSize(&parsed, &size)));
    ASSERT_EQ(size, 4);
    const ZyanArgParseArg* arg;
    ASSERT_TRUE(ZYAN_SUCCESS(ZyanVectorGetPointer(&parsed, 0, (const void**)&arg)));
    ASSERT_STREQ(arg->def->name, "-a");
    ASSERT_FALSE(arg->has_value);
    ASSERT_TRUE(ZYAN_SUCCESS(ZyanVectorGetPointer(&parsed, 1, (const void**)&arg)));
    ASSERT_STREQ(arg->def->name, "-i");
    ASSERT_FALSE(arg->has_value);
    ASSERT_TRUE(ZYAN_SUCCESS(ZyanVectorGetPointer(&parsed, 2, (const void**)&arg)));
    ASSERT_STREQ(arg->def->name, "-o");
    ASSERT_TRUE(arg->has_value);
    ASSERT_EQ(cvt_string_view(&arg->value), "42");
    ASSERT_TRUE(ZYAN_SUCCESS(ZyanVectorGetPointer(&parsed, 3, (const void**)&arg)));
    ASSERT_STREQ(arg->def->name, "-n");
    ASSERT_TRUE(arg->has_value);
    ASSERT_EQ(cvt_string_view(&arg->value), "xxx");
}*/
return true
}

func (a *argParse)TEST()(ok bool){//col:220
/*TEST(DoubleDashArg, PerfectFit)
{
    const char* argv[]
    {
        "./test", "--help", "--stuff", "1337"
    };
    ZyanArgParseDefinition args[]
    {
        {"--help", ZYAN_TRUE, ZYAN_FALSE},
        {"--stuff", ZYAN_FALSE, ZYAN_FALSE},
        {nullptr, ZYAN_FALSE, ZYAN_FALSE}
    };
    ZyanArgParseConfig cfg
    {
    };
    ZyanVector parsed;
    ZYAN_MEMSET(&parsed, 0, sizeof(parsed));
    auto status = ZyanArgParse(&cfg, &parsed, nullptr);
    ASSERT_TRUE(ZYAN_SUCCESS(status));
    ZyanUSize size;
    ASSERT_TRUE(ZYAN_SUCCESS(ZyanVectorGetSize(&parsed, &size)));
    ASSERT_EQ(size, 2);
    const ZyanArgParseArg* arg;
    ASSERT_TRUE(ZYAN_SUCCESS(ZyanVectorGetPointer(&parsed, 0, (const void**)&arg)));
    ASSERT_STREQ(arg->def->name, "--help");
    ASSERT_FALSE(arg->has_value);
    ASSERT_TRUE(ZYAN_SUCCESS(ZyanVectorGetPointer(&parsed, 1, (const void**)&arg)));
    ASSERT_STREQ(arg->def->name, "--stuff");
    ASSERT_TRUE(arg->has_value);
    ASSERT_EQ(cvt_string_view(&arg->value), "1337");
}*/
return true
}

func (a *argParse)TEST()(ok bool){//col:255
/*TEST(MixedArgs, MissingRequiredArg)
{
    const char* argv[]
    {
        "./test", "blah.c", "woof.moo"
    };
    ZyanArgParseDefinition args[]
    {
        {"--feature-xyz", ZYAN_TRUE, ZYAN_FALSE},
        {"-n", ZYAN_FALSE, ZYAN_TRUE},
        {nullptr, ZYAN_FALSE, ZYAN_FALSE}
    };
    ZyanArgParseConfig cfg
    {
    };
    ZyanVector parsed;
    ZYAN_MEMSET(&parsed, 0, sizeof(parsed));
    const char* err_tok = nullptr;
    auto status = ZyanArgParse(&cfg, &parsed, &err_tok);
    ASSERT_EQ(status, ZYAN_STATUS_REQUIRED_ARG_MISSING);
    ASSERT_STREQ(err_tok, "-n");
}*/
return true
}

func (a *argParse)TEST()(ok bool){//col:308
/*TEST(MixedArgs, Stuff)
{
    const char* argv[]
    {
        "./test", "--feature-xyz", "-n5", "blah.c", "woof.moo"
    };
    ZyanArgParseDefinition args[]
    {
        {"--feature-xyz", ZYAN_TRUE, ZYAN_FALSE},
        {"-n", ZYAN_FALSE, ZYAN_FALSE},
        {nullptr, ZYAN_FALSE, ZYAN_FALSE}
    };
    ZyanArgParseConfig cfg
    {
    };
    ZyanVector parsed;
    ZYAN_MEMSET(&parsed, 0, sizeof(parsed));
    auto status = ZyanArgParse(&cfg, &parsed, nullptr);
    ASSERT_TRUE(ZYAN_SUCCESS(status));
    ZyanUSize size;
    ASSERT_TRUE(ZYAN_SUCCESS(ZyanVectorGetSize(&parsed, &size)));
    ASSERT_EQ(size, 4);
    const ZyanArgParseArg* arg;
    ASSERT_TRUE(ZYAN_SUCCESS(ZyanVectorGetPointer(&parsed, 0, (const void**)&arg)));
    ASSERT_STREQ(arg->def->name, "--feature-xyz");
    ASSERT_FALSE(arg->has_value);
    ASSERT_TRUE(ZYAN_SUCCESS(ZyanVectorGetPointer(&parsed, 1, (const void**)&arg)));
    ASSERT_STREQ(arg->def->name, "-n");
    ASSERT_TRUE(arg->has_value);
    ASSERT_EQ(cvt_string_view(&arg->value), "5");
    ASSERT_TRUE(ZYAN_SUCCESS(ZyanVectorGetPointer(&parsed, 2, (const void**)&arg)));
    ASSERT_EQ(arg->def, nullptr);
    ASSERT_TRUE(arg->has_value);
    ASSERT_EQ(cvt_string_view(&arg->value), "blah.c");
    ASSERT_TRUE(ZYAN_SUCCESS(ZyanVectorGetPointer(&parsed, 3, (const void**)&arg)));
    ASSERT_EQ(arg->def, nullptr);
    ASSERT_TRUE(arg->has_value);
    ASSERT_EQ(cvt_string_view(&arg->value), "woof.moo");
}*/
return true
}

func (a *argParse)int main()(ok bool){//col:318
/*int main(int argc, char **argv)
{
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}*/
return true
}



