package examples
//back\HyperDbgDev\hyperdbg\dependencies\zydis\examples\ZydisPerfTest.c.back

const(
COLOR_DEFAULT =       ZYAN_VT100SGR_FG_DEFAULT //col:1
COLOR_ERROR =         ZYAN_VT100SGR_FG_BRIGHT_RED //col:2
COLOR_VALUE_R =       ZYAN_VT100SGR_FG_BRIGHT_RED //col:3
COLOR_VALUE_G =       ZYAN_VT100SGR_FG_BRIGHT_GREEN //col:4
COLOR_VALUE_B =       ZYAN_VT100SGR_FG_CYAN //col:5
CVT100_OUT(sequence) = (g_vt100_stdout ? (sequence) : "") //col:6
CVT100_ERR(sequence) = (g_vt100_stderr ? (sequence) : "") //col:7
)

type (
ZydisPerfTest interface{
static_void_StartCounter()(ok bool)//col:13
static_double_GetCounter()(ok bool)//col:19
static_void_StartCounter()(ok bool)//col:23
static_double_GetCounter()(ok bool)//col:32
static_void_StartCounter()(ok bool)//col:36
static_double_GetCounter()(ok bool)//col:43
static_void_AdjustProcessAndThreadPriority()(ok bool)//col:82
static_ZyanU64_ProcessBuffer()(ok bool)//col:137
static_void_TestPerformance()(ok bool)//col:190
static_void_GenerateTestData()(ok bool)//col:276
int_main()(ok bool)//col:385
}
)

func NewZydisPerfTest() { return & zydisPerfTest{} }

func (z *zydisPerfTest)static_void_StartCounter()(ok bool){//col:13
/*static void StartCounter(void)
{
    LARGE_INTEGER li;
    if (!QueryPerformanceFrequency(&li))
    {
        ZYAN_FPRINTF(ZYAN_STDERR, "%sError: QueryPerformanceFrequency failed!%s\n",
            CVT100_ERR(COLOR_ERROR), CVT100_ERR(ZYAN_VT100SGR_RESET));
        exit(EXIT_FAILURE);
    }
    counter_freq = (double)li.QuadPart / 1000.0;
    QueryPerformanceCounter(&li);
    counter_start = li.QuadPart;
}*/
return true
}

func (z *zydisPerfTest)static_double_GetCounter()(ok bool){//col:19
/*static double GetCounter(void)
{
    LARGE_INTEGER li;
    QueryPerformanceCounter(&li);
    return (double)(li.QuadPart - counter_start) / counter_freq;
}*/
return true
}

func (z *zydisPerfTest)static_void_StartCounter()(ok bool){//col:23
/*static void StartCounter(void)
{
    counter_start = mach_absolute_time();
}*/
return true
}

func (z *zydisPerfTest)static_double_GetCounter()(ok bool){//col:32
/*static double GetCounter(void)
{
    ZyanU64 elapsed = mach_absolute_time() - counter_start;
    if (timebase_info.denom == 0)
    {
        mach_timebase_info(&timebase_info);
    }
    return (double)elapsed * timebase_info.numer / timebase_info.denom / 1000000;
}*/
return true
}

func (z *zydisPerfTest)static_void_StartCounter()(ok bool){//col:36
/*static void StartCounter(void)
{
    gettimeofday(&t1, NULL);
}*/
return true
}

func (z *zydisPerfTest)static_double_GetCounter()(ok bool){//col:43
/*static double GetCounter(void)
{
    struct timeval t2;
    gettimeofday(&t2, NULL);
    double t = (t2.tv_sec - t1.tv_sec) * 1000.0;
    return t + (t2.tv_usec - t1.tv_usec) / 1000.0;
}*/
return true
}

func (z *zydisPerfTest)static_void_AdjustProcessAndThreadPriority()(ok bool){//col:82
/*static void AdjustProcessAndThreadPriority(void)
{
#if defined(ZYAN_WINDOWS)
    SYSTEM_INFO info;
    GetSystemInfo(&info);
    if (info.dwNumberOfProcessors > 1)
    {
        if (!SetThreadAffinityMask(GetCurrentThread(), (DWORD_PTR)1))
        {
            ZYAN_FPRINTF(ZYAN_STDERR, "%sWarning: Could not set thread affinity mask%s\n",
                CVT100_ERR(ZYAN_VT100SGR_FG_YELLOW), CVT100_ERR(ZYAN_VT100SGR_RESET));
        }
        if (!SetPriorityClass(GetCurrentProcess(), REALTIME_PRIORITY_CLASS))
        {
            ZYAN_FPRINTF(ZYAN_STDERR, "%sWarning: Could not set process priority class%s\n",
                CVT100_ERR(ZYAN_VT100SGR_FG_YELLOW), CVT100_ERR(ZYAN_VT100SGR_RESET));
        }
        if (!SetThreadPriority(GetCurrentThread(), THREAD_PRIORITY_TIME_CRITICAL))
        {
            ZYAN_FPRINTF(ZYAN_STDERR, "%sWarning: Could not set thread priority class%s\n",
                CVT100_ERR(ZYAN_VT100SGR_FG_YELLOW), CVT100_ERR(ZYAN_VT100SGR_RESET));
        }
    }
#elif defined(ZYAN_LINUX) || defined(ZYAN_FREEBSD)
    pthread_t thread = pthread_self();
#if defined(ZYAN_LINUX)
    cpu_set_t cpus;
#else  
    cpuset_t cpus;
#endif
    CPU_ZERO(&cpus);
    CPU_SET(0, &cpus);
    if (pthread_setaffinity_np(thread, sizeof(cpus), &cpus))
    {
        ZYAN_FPRINTF(ZYAN_STDERR, "%sWarning: Could not set thread affinity mask%s\n",
            CVT100_ERR(ZYAN_VT100SGR_FG_YELLOW), CVT100_ERR(ZYAN_VT100SGR_RESET));
    }
#endif
}*/
return true
}

func (z *zydisPerfTest)static_ZyanU64_ProcessBuffer()(ok bool){//col:137
/*static ZyanU64 ProcessBuffer(const ZydisDecoder* decoder, const ZydisFormatter* formatter,
    const ZyanU8* buffer, ZyanUSize length, ZyanBool format, ZyanBool tokenize, ZyanBool use_cache)
{
    ZyanU64 count = 0;
    ZyanUSize offset = 0;
    ZyanStatus status;
    ZydisDecodedInstruction instruction_data;
    ZydisDecodedInstruction* instruction;
    char format_buffer[256];
    while (length - offset > 0)
    {
        if (use_cache)
        {
            ZYAN_UNREACHABLE;
        } else
        {
            status = ZydisDecoderDecodeBuffer(decoder, buffer + offset, length - offset,
                &instruction_data);
            instruction = &instruction_data;
        }
        if (status == ZYDIS_STATUS_NO_MORE_DATA)
        {
            break;
        }
        if (!ZYAN_SUCCESS(status))
        {
            ZYAN_FPRINTF(ZYAN_STDERR, "%sUnexpected decoding error. Data: ",
                CVT100_ERR(COLOR_ERROR));
            for (ZyanUSize i = 0; i < ZYAN_MIN(ZYDIS_MAX_INSTRUCTION_LENGTH,
                length - offset); ++i)
            {
                ZYAN_FPRINTF(ZYAN_STDERR, "%02X ", (ZyanU8)buffer[offset + i]);
            }
            ZYAN_FPRINTF(ZYAN_STDERR, "%s\n", CVT100_ERR(ZYAN_VT100SGR_RESET));
            ZYAN_ASSERT(ZYAN_FALSE);
            exit(EXIT_FAILURE);
        }
        if (format)
        {
            if (tokenize)
            {
                const ZydisFormatterToken* token;
                ZydisFormatterTokenizeInstruction(formatter, instruction, format_buffer,
                    sizeof(format_buffer), offset, &token);
            } else
            {
                ZydisFormatterFormatInstruction(formatter, instruction, format_buffer,
                    sizeof(format_buffer), offset);
            }
        }
        offset += instruction->length;
        ++count;
    }
    return count;
}*/
return true
}

func (z *zydisPerfTest)static_void_TestPerformance()(ok bool){//col:190
/*static void TestPerformance(const ZyanU8* buffer, ZyanUSize length, ZyanBool minimal_mode,
    ZyanBool format, ZyanBool tokenize, ZyanBool use_cache)
{
    ZydisDecoder decoder;
    if (!ZYAN_SUCCESS(ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64,
        ZYDIS_ADDRESS_WIDTH_64)))
    {
        ZYAN_FPRINTF(ZYAN_STDERR, "%sFailed to initialize decoder%s\n",
            CVT100_ERR(COLOR_ERROR), CVT100_ERR(ZYAN_VT100SGR_RESET));
        exit(EXIT_FAILURE);
    }
    if (!ZYAN_SUCCESS(ZydisDecoderEnableMode(&decoder, ZYDIS_DECODER_MODE_MINIMAL, minimal_mode)))
    {
        ZYAN_FPRINTF(ZYAN_STDERR, "%sFailed to adjust decoder-mode%s\n",
            CVT100_ERR(COLOR_ERROR), CVT100_ERR(ZYAN_VT100SGR_RESET));
        exit(EXIT_FAILURE);
    }
    ZydisFormatter formatter;
    if (format)
    {
        if (!ZYAN_SUCCESS(ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL)) ||
            !ZYAN_SUCCESS(ZydisFormatterSetProperty(&formatter,
                ZYDIS_FORMATTER_PROP_FORCE_SEGMENT, ZYAN_TRUE)) ||
            !ZYAN_SUCCESS(ZydisFormatterSetProperty(&formatter,
                ZYDIS_FORMATTER_PROP_FORCE_SIZE, ZYAN_TRUE)))
        {
            ZYAN_FPRINTF(ZYAN_STDERR, "%sFailed to initialize instruction-formatter%s\n",
                CVT100_ERR(COLOR_ERROR), CVT100_ERR(ZYAN_VT100SGR_RESET));
            exit(EXIT_FAILURE);
        }
    }
    ProcessBuffer(&decoder, &formatter, buffer, length, format, tokenize, use_cache);
    ZyanU64 count = 0;
    StartCounter();
    for (ZyanU8 j = 0; j < 100; ++j)
    {
        count += ProcessBuffer(&decoder, &formatter, buffer, length, format,
            tokenize, use_cache);
    }
    const char* color[4];
    color[0] = minimal_mode ? CVT100_OUT(COLOR_VALUE_G) : CVT100_OUT(COLOR_VALUE_B);
    color[1] = format       ? CVT100_OUT(COLOR_VALUE_G) : CVT100_OUT(COLOR_VALUE_B);
    color[2] = tokenize     ? CVT100_OUT(COLOR_VALUE_G) : CVT100_OUT(COLOR_VALUE_B);
    color[3] = use_cache    ? CVT100_OUT(COLOR_VALUE_G) : CVT100_OUT(COLOR_VALUE_B);
    ZYAN_PRINTF("Minimal-Mode %s%d%s, Format %s%d%s, Tokenize %s%d%s, Caching %s%d%s, " \
        "Instructions: %s%6.2fM%s, Time: %s%8.2f%s msec\n",
        color[0], minimal_mode, CVT100_OUT(COLOR_DEFAULT),
        color[1], format, CVT100_OUT(COLOR_DEFAULT),
        color[2], tokenize, CVT100_OUT(COLOR_DEFAULT),
        color[3], use_cache, CVT100_OUT(COLOR_DEFAULT),
        CVT100_OUT(COLOR_VALUE_B), (double)count / 1000000, CVT100_OUT(COLOR_DEFAULT),
        CVT100_OUT(COLOR_VALUE_G), GetCounter(), CVT100_OUT(COLOR_DEFAULT));
}*/
return true
}

func (z *zydisPerfTest)static_void_GenerateTestData()(ok bool){//col:276
/*static void GenerateTestData(FILE* file, ZyanU8 encoding)
{
    ZydisDecoder decoder;
    if (!ZYAN_SUCCESS(
        ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64, ZYDIS_ADDRESS_WIDTH_64)))
    {
        ZYAN_FPRINTF(ZYAN_STDERR, "%sFailed to initialize decoder%s\n", CVT100_ERR(COLOR_ERROR),
            CVT100_ERR(ZYAN_VT100SGR_RESET));
        exit(EXIT_FAILURE);
    }
    ZyanU8 last = 0;
    ZyanU32 count = 0;
    ZydisDecodedInstruction instruction;
    while (count < 100000)
    {
        ZyanU8 data[ZYDIS_MAX_INSTRUCTION_LENGTH];
        for (int i = 0; i < ZYDIS_MAX_INSTRUCTION_LENGTH; ++i)
        {
            data[i] = rand() % 256;
        }
        const ZyanU8 offset = rand() % (ZYDIS_MAX_INSTRUCTION_LENGTH - 2);
        switch (encoding)
        {
        case 0:
            break;
        case 1:
            data[offset    ] = 0x0F;
            data[offset + 1] = 0x0F;
            break;
        case 2:
            data[offset    ] = 0x8F;
            break;
        case 3:
            data[offset    ] = 0xC4;
            break;
        case 4:
            data[offset    ] = 0xC5;
            break;
        case 5:
        case 6:
            data[offset    ] = 0x62;
            break;
        default:
            ZYAN_UNREACHABLE;
        }
        if (ZYAN_SUCCESS(ZydisDecoderDecodeBuffer(&decoder, data, sizeof(data), &instruction)))
        {
            ZyanBool b = ZYAN_FALSE;
            switch (encoding)
            {
            case 0:
                b = (instruction.encoding == ZYDIS_INSTRUCTION_ENCODING_LEGACY);
                break;
            case 1:
                b = (instruction.encoding == ZYDIS_INSTRUCTION_ENCODING_3DNOW);
                break;
            case 2:
                b = (instruction.encoding == ZYDIS_INSTRUCTION_ENCODING_XOP);
                break;
            case 3:
            case 4:
                b = (instruction.encoding == ZYDIS_INSTRUCTION_ENCODING_VEX);
                break;
            case 5:
                b = (instruction.encoding == ZYDIS_INSTRUCTION_ENCODING_EVEX);
                break;
            case 6:
                b = (instruction.encoding == ZYDIS_INSTRUCTION_ENCODING_MVEX);
                break;
            default:
                ZYAN_UNREACHABLE;
            }
            if (b)
            {
                fwrite(&data[0], sizeof(ZyanU8), instruction.length, file);
                ++count;
                const ZyanU8 p = (ZyanU8)((double)count / 100000 * 100);
                if (last < p)
                {
                    last = p;
                    ZYAN_PRINTF("%3.0d%%\n", p);
                }
            }
        }
    }
}*/
return true
}

func (z *zydisPerfTest)int_main()(ok bool){//col:385
/*int main(int argc, char** argv)
{
    g_vt100_stdout = (ZyanTerminalIsTTY(ZYAN_STDSTREAM_OUT) == ZYAN_STATUS_TRUE) &&
                     ZYAN_SUCCESS(ZyanTerminalEnableVT100(ZYAN_STDSTREAM_OUT));
    g_vt100_stderr = (ZyanTerminalIsTTY(ZYAN_STDSTREAM_ERR) == ZYAN_STATUS_TRUE) &&
                     ZYAN_SUCCESS(ZyanTerminalEnableVT100(ZYAN_STDSTREAM_ERR));
    if (ZydisGetVersion() != ZYDIS_VERSION)
    {
        ZYAN_FPRINTF(ZYAN_STDERR, "%sInvalid zydis version%s\n",
            CVT100_ERR(COLOR_ERROR), CVT100_ERR(ZYAN_VT100SGR_RESET));
        return EXIT_FAILURE;
    }
    if (argc < 3 || (ZYAN_STRCMP(argv[1], "-test") && ZYAN_STRCMP(argv[1], "-generate")))
    {
        ZYAN_FPRINTF(ZYAN_STDERR, "%sUsage: %s -[test|generate] [directory]%s\n",
            CVT100_ERR(COLOR_ERROR), (argc > 0 ? argv[0] : "PerfTest"),
            CVT100_ERR(ZYAN_VT100SGR_RESET));
        return EXIT_FAILURE;
    }
    ZyanBool generate = ZYAN_FALSE;
    if (!ZYAN_STRCMP(argv[1], "-generate"))
    {
        generate = ZYAN_TRUE;
    }
    const char* directory = argv[2];
    static const struct
    {
        const char* encoding;
        const char* filename;
    } tests[7] =
    {
        { "DEFAULT", "enc_default.dat" },
        { "3DNOW"  , "enc_3dnow.dat"   },
        { "XOP"    , "enc_xop.dat"     },
        { "VEX_C4" , "enc_vex_c4.dat"  },
        { "VEX_C5" , "enc_vex_c5.dat"  },
        { "EVEX"   , "enc_evex.dat"    },
        { "MVEX"   , "enc_mvex.dat"    }
    };
    if (generate)
    {
        time_t t;
        srand((unsigned)time(&t));
    } else
    {
        AdjustProcessAndThreadPriority();
    }
    for (ZyanU8 i = 0; i < ZYAN_ARRAY_LENGTH(tests); ++i)
    {
        FILE* file;
        const ZyanUSize len = strlen(directory);
        char buf[1024];
        strncpy(&buf[0], directory, sizeof(buf) - 1);
        if (generate)
        {
            file = fopen(strncat(buf, tests[i].filename, sizeof(buf) - len - 1), "wb");
        } else
        {
            file = fopen(strncat(buf, tests[i].filename, sizeof(buf) - len - 1), "rb");
        }
        if (!file)
        {
            ZYAN_FPRINTF(ZYAN_STDERR, "%sCould not open file \"%s\": %s%s\n",
                CVT100_ERR(COLOR_ERROR), &buf[0], strerror(ZYAN_ERRNO),
                CVT100_ERR(ZYAN_VT100SGR_RESET));
            continue;
        }
        if (generate)
        {
            ZYAN_PRINTF("Generating %s%s%s ...\n", CVT100_OUT(COLOR_VALUE_B), tests[i].encoding,
                CVT100_OUT(ZYAN_VT100SGR_RESET));
            GenerateTestData(file, i);
        } else
        {
            fseek(file, 0L, SEEK_END);
            const long length = ftell(file);
            void* buffer = malloc(length);
            if (!buffer)
            {
                ZYAN_FPRINTF(ZYAN_STDERR,
                    "%sFailed to allocate %" PRIu64 " bytes on the heap%s\n",
                    CVT100_ERR(COLOR_ERROR), (ZyanU64)length, CVT100_ERR(ZYAN_VT100SGR_RESET));
                goto NextFile2;
            }
            rewind(file);
            if (fread(buffer, 1, length, file) != (ZyanUSize)length)
            {
                ZYAN_FPRINTF(ZYAN_STDERR,
                    "%sCould not read %" PRIu64 " bytes from file \"%s\"%s\n",
                    CVT100_ERR(COLOR_ERROR), (ZyanU64)length, &buf[0],
                    CVT100_ERR(ZYAN_VT100SGR_RESET));
                goto NextFile1;
            }
            ZYAN_PRINTF("%sTesting %s%s%s ...\n", CVT100_OUT(ZYAN_VT100SGR_FG_MAGENTA),
                CVT100_OUT(ZYAN_VT100SGR_FG_BRIGHT_MAGENTA), tests[i].encoding,
                CVT100_OUT(COLOR_DEFAULT));
            TestPerformance(buffer, length, ZYAN_TRUE , ZYAN_FALSE, ZYAN_FALSE, ZYAN_FALSE);
            TestPerformance(buffer, length, ZYAN_FALSE, ZYAN_FALSE, ZYAN_FALSE, ZYAN_FALSE);
            TestPerformance(buffer, length, ZYAN_FALSE, ZYAN_TRUE , ZYAN_FALSE, ZYAN_FALSE);
            TestPerformance(buffer, length, ZYAN_FALSE, ZYAN_TRUE , ZYAN_TRUE , ZYAN_FALSE);
            ZYAN_PUTS("");
        NextFile1:
            free(buffer);
        }
    NextFile2:
        fclose(file);
    }
    return 0;
}*/
return true
}



