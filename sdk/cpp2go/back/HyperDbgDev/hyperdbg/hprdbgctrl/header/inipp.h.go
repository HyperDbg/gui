package header
//back\HyperDbgDev\hyperdbg\hprdbgctrl\header\inipp.h.back

type (
Inipp interface{
Copyright ()(ok bool)//col:52
rtrim()(ok bool)//col:60
rtrim2()(ok bool)//col:67
replace()(ok bool)//col:84
extract()(ok bool)//col:104
extract()(ok bool)//col:112
get_value()(ok bool)//col:122
get_value()(ok bool)//col:129
    virtual bool is_section_start()(ok bool)//col:168
    Ini()(ok bool)//col:315
}
)

func NewInipp() { return & inipp{} }

func (i *inipp)Copyright ()(ok bool){//col:52
/*Copyright (c) 2017-2020 Matthias C. M. Troffaes
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
#pragma once
#include <algorithm>
#include <cctype>
#include <cstring>
#include <functional>
#include <iostream>
#include <list>
#include <vector>
#include <locale>
#include <map>
#include <memory>
#include <sstream>
#include <string>
namespace inipp {
namespace detail {
template <class CharT>
inline void
ltrim(std::basic_string<CharT> & s, const std::locale & loc)
{
    s.erase(s.begin(),
            std::find_if(s.begin(), s.end(), [&loc](CharT ch) { return !std::isspace(ch, loc); }));
}*/
return true
}

func (i *inipp)rtrim()(ok bool){//col:60
/*rtrim(std::basic_string<CharT> & s, const std::locale & loc)
{
    s.erase(std::find_if(s.rbegin(), s.rend(), [&loc](CharT ch) { return !std::isspace(ch, loc); }).base(),
            s.end());
}*/
return true
}

func (i *inipp)rtrim2()(ok bool){//col:67
/*rtrim2(std::basic_string<CharT> & s, UnaryPredicate pred)
{
    s.erase(std::find_if(s.begin(), s.end(), pred), s.end());
}*/
return true
}

func (i *inipp)replace()(ok bool){//col:84
/*replace(std::basic_string<CharT> & str, const std::basic_string<CharT> & from, const std::basic_string<CharT> & to)
{
    auto   changed   = false;
    size_t start_pos = 0;
    while ((start_pos = str.find(from, start_pos)) != std::basic_string<CharT>::npos)
    {
        str.replace(start_pos, from.length(), to);
        start_pos += to.length();
        changed = true;
    }
    return changed;
}*/
return true
}

func (i *inipp)extract()(ok bool){//col:104
/*extract(const std::basic_string<CharT> & value, T & dst)
{
    CharT                           c;
    std::basic_istringstream<CharT> is {value};
    T                               result;
    if ((is >> std::boolalpha >> result) && !(is >> c))
    {
        dst = result;
        return true;
    }
    else
    {
        return false;
    }
}*/
return true
}

func (i *inipp)extract()(ok bool){//col:112
/*extract(const std::basic_string<CharT> & value, std::basic_string<CharT> & dst)
{
    dst = value;
    return true;
}*/
return true
}

func (i *inipp)get_value()(ok bool){//col:122
/*get_value(const std::map<std::basic_string<CharT>, std::basic_string<CharT>> & sec, const std::basic_string<CharT> & key, T & dst)
{
    const auto it = sec.find(key);
    if (it == sec.end())
        return false;
    return extract(it->second, dst);
}*/
return true
}

func (i *inipp)get_value()(ok bool){//col:129
/*get_value(const std::map<std::basic_string<CharT>, std::basic_string<CharT>> & sec, const CharT * key, T & dst)
{
    return get_value(sec, std::basic_string<CharT>(key), dst);
}*/
return true
}

func (i *inipp)    virtual bool is_section_start()(ok bool){//col:168
/*    virtual bool is_section_start(CharT ch) const { return ch == char_section_start; }
    virtual bool is_section_end(CharT ch) const { return ch == char_section_end; }
    virtual bool is_assign(CharT ch) const { return ch == char_assign; }
    virtual bool is_comment(CharT ch) const { return ch == char_comment; }
    const CharT char_interpol;
    const CharT char_interpol_start;
    const CharT char_interpol_sep;
    const CharT char_interpol_end;
    Format(CharT section_start, CharT section_end, CharT assign, CharT comment, CharT interpol, CharT interpol_start, CharT interpol_sep, CharT interpol_end) :
        char_section_start(section_start), char_section_end(section_end), char_assign(assign), char_comment(comment), char_interpol(interpol), char_interpol_start(interpol_start), char_interpol_sep(interpol_sep), char_interpol_end(interpol_end) { }
    Format() :
        Format('[', ']', '=', ';', '$', '{', ':', '}') { }
    const std::basic_string<CharT> local_symbol(const std::basic_string<CharT> & name) const
    {
        return char_interpol + (char_interpol_start + name + char_interpol_end);
    }
    const std::basic_string<CharT> global_symbol(const std::basic_string<CharT> & sec_name, const std::basic_string<CharT> & name) const
    {
        return local_symbol(sec_name + char_interpol_sep + name);
    }
};*/
return true
}

func (i *inipp)    Ini()(ok bool){//col:315
/*    Ini() :
        format(std::make_shared<Format<CharT>>()) {};
    Ini(std::shared_ptr<Format<CharT>> fmt) :
        format(fmt) {};
    void generate(std::basic_ostream<CharT> & os) const
    {
        for (auto const & sec : sections)
        {
            os << format->char_section_start << sec.first << format->char_section_end << std::endl;
            for (auto const & val : sec.second)
            {
                os << val.first << format->char_assign << val.second << std::endl;
            }
            os << std::endl;
        }
    }
    void parse(std::basic_istream<CharT> & is)
    {
        String            line;
        String            section;
        const std::locale loc {"C"};
        while (std::getline(is, line))
        {
            detail::ltrim(line, loc);
            detail::rtrim(line, loc);
            const auto length = line.length();
            if (length > 0)
            {
                const auto   pos   = std::find_if(line.begin(), line.end(), [this](CharT ch) { return format->is_assign(ch); });
                const auto & front = line.front();
                if (format->is_comment(front))
                {
                }
                else if (format->is_section_start(front))
                {
                    if (format->is_section_end(line.back()))
                        section = line.substr(1, length - 2);
                    else
                        errors.push_back(line);
                }
                else if (pos != line.begin() && pos != line.end())
                {
                    String variable(line.begin(), pos);
                    String value(pos + 1, line.end());
                    detail::rtrim(variable, loc);
                    detail::ltrim(value, loc);
                    auto & sec = sections[section];
                    if (sec.find(variable) == sec.end())
                        sec.emplace(variable, value);
                    else
                        errors.push_back(line);
                }
                else
                {
                    errors.push_back(line);
                }
            }
        }
    }
    void interpolate()
    {
        int  global_iteration = 0;
        auto changed          = false;
        for (auto & sec : sections)
            replace_symbols(local_symbols(sec.first, sec.second), sec.second);
        do
        {
            changed         = false;
            const auto syms = global_symbols();
            for (auto & sec : sections)
                changed |= replace_symbols(syms, sec.second);
        } while (changed && (max_interpolation_depth > global_iteration++));
    }
    void default_section(const Section & sec)
    {
        for (auto & sec2 : sections)
            for (const auto & val : sec)
                sec2.second.insert(val);
    }
    void strip_trailing_comments()
    {
        const std::locale loc {"C"};
        for (auto & sec : sections)
            for (auto & val : sec.second)
            {
                detail::rtrim2(val.second, [this](CharT ch) { return format->is_comment(ch); });
                detail::rtrim(val.second, loc);
            }
    }
    void clear()
    {
        sections.clear();
        errors.clear();
    }
private:
    using Symbols = std::vector<std::pair<String, String>>;
    const Symbols local_symbols(const String & sec_name, const Section & sec) const
    {
        Symbols result;
        for (const auto & val : sec)
            result.emplace_back(format->local_symbol(val.first), format->global_symbol(sec_name, val.first));
        return result;
    }
    const Symbols global_symbols() const
    {
        Symbols result;
        for (const auto & sec : sections)
            for (const auto & val : sec.second)
                result.emplace_back(format->global_symbol(sec.first, val.first), val.second);
        return result;
    }
    bool replace_symbols(const Symbols & syms, Section & sec) const
    {
        auto changed = false;
        for (auto & sym : syms)
            for (auto & val : sec)
                changed |= detail::replace(val.second, sym.first, sym.second);
        return changed;
    }
};*/
return true
}



