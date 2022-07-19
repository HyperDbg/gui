#include "pch.h"
BOOLEAN
TestCase(std::vector<std::string> & TestCases) {
    string        Line;
    std::ifstream File(TEST_CASE_FILE_NAME);
    if (File.is_open()) {
        while (std::getline(File, Line)) {
            TestCases.push_back(Line);
        }
        File.close();
    } else {
        printf("err, test case file not found (%s) \npress enter to continue", TEST_CASE_FILE_NAME);
        _getch();
        return FALSE;
    }
    return TRUE;
}
