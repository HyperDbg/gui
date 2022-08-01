package header
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\header\inipp.h.back

type (
Inipp interface{
ltrim()(ok bool)//col:5
rtrim()(ok bool)//col:10
rtrim2()(ok bool)//col:14
replace()(ok bool)//col:26
extract()(ok bool)//col:41
extract()(ok bool)//col:46
get_value()(ok bool)//col:53
get_value()(ok bool)//col:57
____virtual_bool_is_section_start()(ok bool)//col:78
____Ini()(ok bool)//col:209
____virtual_bool_is_section_end()(ok bool)//col:230
____Ini()(ok bool)//col:361
____virtual_bool_is_assign()(ok bool)//col:381
____Ini()(ok bool)//col:512
____virtual_bool_is_comment()(ok bool)//col:531
____Ini()(ok bool)//col:662
____Format()(ok bool)//col:676
____Ini()(ok bool)//col:807
________char_section_start()(ok bool)//col:820
____Ini()(ok bool)//col:951
____Format()(ok bool)//col:963
____Ini()(ok bool)//col:1094
________Format()(ok bool)//col:1105
____Ini()(ok bool)//col:1236
____const_std::basic_string<CharT>_local_symbol()(ok bool)//col:1246
____Ini()(ok bool)//col:1377
________return_char_interpol_+_()(ok bool)//col:1385
____Ini()(ok bool)//col:1516
____const_std::basic_string<CharT>_global_symbol()(ok bool)//col:1522
____Ini()(ok bool)//col:1653
________return_local_symbol()(ok bool)//col:1657
____Ini()(ok bool)//col:1788
____Ini()(ok bool)//col:1909
________format()(ok bool)//col:2029
____Ini()(ok bool)//col:2148
________format()(ok bool)//col:2266
____void_generate()(ok bool)//col:2383
________for_()(ok bool)//col:2498
____________for_()(ok bool)//col:2610
____void_parse()(ok bool)//col:2715
________while_()(ok bool)//col:2815
____________detail::ltrim()(ok bool)//col:2913
____________detail::rtrim()(ok bool)//col:3010
____________const_auto_length_=_line.length()(ok bool)//col:3106
____________if_()(ok bool)//col:3201
________________const_auto___pos___=_std::find_if()(ok bool)//col:3294
________________const_auto_&_front_=_line.front()(ok bool)//col:3386
________________if_()(ok bool)//col:3477
________________else_if_()(ok bool)//col:3565
____________________if_()(ok bool)//col:3651
________________________section_=_line.substr()(ok bool)//col:3736
________________________errors.push_back()(ok bool)//col:3819
________________else_if_()(ok bool)//col:3900
____________________String_variable()(ok bool)//col:3979
____________________String_value()(ok bool)//col:4057
____________________detail::rtrim()(ok bool)//col:4134
____________________detail::ltrim()(ok bool)//col:4210
____________________if_()(ok bool)//col:4284
________________________sec.emplace()(ok bool)//col:4357
________________________errors.push_back()(ok bool)//col:4428
____________________errors.push_back()(ok bool)//col:4495
____void_interpolate()(ok bool)//col:4557
________for_()(ok bool)//col:4615
____________replace_symbols()(ok bool)//col:4672
____________const_auto_syms_=_global_symbols()(ok bool)//col:4725
____________for_()(ok bool)//col:4777
________________changed_|=_replace_symbols()(ok bool)//col:4828
________}_while_()(ok bool)//col:4878
____void_default_section()(ok bool)//col:4926
________for_()(ok bool)//col:4972
____________for_()(ok bool)//col:5017
________________sec2.second.insert()(ok bool)//col:5061
____void_strip_trailing_comments()(ok bool)//col:5103
________for_()(ok bool)//col:5142
____________for_()(ok bool)//col:5180
________________detail::rtrim2()(ok bool)//col:5216
________________detail::rtrim()(ok bool)//col:5251
____void_clear()(ok bool)//col:5283
________sections.clear()(ok bool)//col:5313
________errors.clear()(ok bool)//col:5342
____const_Symbols_local_symbols()(ok bool)//col:5367
________for_()(ok bool)//col:5389
____________result.emplace_back()(ok bool)//col:5410
____const_Symbols_global_symbols()(ok bool)//col:5428
________for_()(ok bool)//col:5443
____________for_()(ok bool)//col:5457
________________result.emplace_back()(ok bool)//col:5470
____bool_replace_symbols()(ok bool)//col:5480
}
)

func NewInipp() { return & inipp{} }

func (i *inipp)ltrim()(ok bool){//col:5
/*ltrim(std::basic_string<CharT> & s, const std::locale & loc)
{
    s.erase(s.begin(),
            std::find_if(s.begin(), s.end(), [&loc](CharT ch) { return !std::isspace(ch, loc); }));
}*/
return true
}

func (i *inipp)rtrim()(ok bool){//col:10
/*rtrim(std::basic_string<CharT> & s, const std::locale & loc)
{
    s.erase(std::find_if(s.rbegin(), s.rend(), [&loc](CharT ch) { return !std::isspace(ch, loc); }).base(),
            s.end());
}*/
return true
}

func (i *inipp)rtrim2()(ok bool){//col:14
/*rtrim2(std::basic_string<CharT> & s, UnaryPredicate pred)
{
    s.erase(std::find_if(s.begin(), s.end(), pred), s.end());
}*/
return true
}

func (i *inipp)replace()(ok bool){//col:26
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

func (i *inipp)extract()(ok bool){//col:41
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

func (i *inipp)extract()(ok bool){//col:46
/*extract(const std::basic_string<CharT> & value, std::basic_string<CharT> & dst)
{
    dst = value;
    return true;
}*/
return true
}

func (i *inipp)get_value()(ok bool){//col:53
/*get_value(const std::map<std::basic_string<CharT>, std::basic_string<CharT>> & sec, const std::basic_string<CharT> & key, T & dst)
{
    const auto it = sec.find(key);
    if (it == sec.end())
        return false;
    return extract(it->second, dst);
}*/
return true
}

func (i *inipp)get_value()(ok bool){//col:57
/*get_value(const std::map<std::basic_string<CharT>, std::basic_string<CharT>> & sec, const CharT * key, T & dst)
{
    return get_value(sec, std::basic_string<CharT>(key), dst);
}*/
return true
}

func (i *inipp)____virtual_bool_is_section_start()(ok bool){//col:78
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

func (i *inipp)____Ini()(ok bool){//col:209
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

func (i *inipp)____virtual_bool_is_section_end()(ok bool){//col:230
/*    virtual bool is_section_end(CharT ch) const { return ch == char_section_end; }
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

func (i *inipp)____Ini()(ok bool){//col:361
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

func (i *inipp)____virtual_bool_is_assign()(ok bool){//col:381
/*    virtual bool is_assign(CharT ch) const { return ch == char_assign; }
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

func (i *inipp)____Ini()(ok bool){//col:512
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

func (i *inipp)____virtual_bool_is_comment()(ok bool){//col:531
/*    virtual bool is_comment(CharT ch) const { return ch == char_comment; }
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

func (i *inipp)____Ini()(ok bool){//col:662
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

func (i *inipp)____Format()(ok bool){//col:676
/*    Format(CharT section_start, CharT section_end, CharT assign, CharT comment, CharT interpol, CharT interpol_start, CharT interpol_sep, CharT interpol_end) :
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

func (i *inipp)____Ini()(ok bool){//col:807
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

func (i *inipp)________char_section_start()(ok bool){//col:820
/*        char_section_start(section_start), char_section_end(section_end), char_assign(assign), char_comment(comment), char_interpol(interpol), char_interpol_start(interpol_start), char_interpol_sep(interpol_sep), char_interpol_end(interpol_end) { }
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

func (i *inipp)____Ini()(ok bool){//col:951
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

func (i *inipp)____Format()(ok bool){//col:963
/*    Format() :
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

func (i *inipp)____Ini()(ok bool){//col:1094
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

func (i *inipp)________Format()(ok bool){//col:1105
/*        Format('[', ']', '=', ';', '$', '{', ':', '}') { }
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

func (i *inipp)____Ini()(ok bool){//col:1236
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

func (i *inipp)____const_std::basic_string<CharT>_local_symbol()(ok bool){//col:1246
/*    const std::basic_string<CharT> local_symbol(const std::basic_string<CharT> & name) const
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

func (i *inipp)____Ini()(ok bool){//col:1377
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

func (i *inipp)________return_char_interpol_+_()(ok bool){//col:1385
/*        return char_interpol + (char_interpol_start + name + char_interpol_end);
    }
    const std::basic_string<CharT> global_symbol(const std::basic_string<CharT> & sec_name, const std::basic_string<CharT> & name) const
    {
        return local_symbol(sec_name + char_interpol_sep + name);
    }
};*/
return true
}

func (i *inipp)____Ini()(ok bool){//col:1516
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

func (i *inipp)____const_std::basic_string<CharT>_global_symbol()(ok bool){//col:1522
/*    const std::basic_string<CharT> global_symbol(const std::basic_string<CharT> & sec_name, const std::basic_string<CharT> & name) const
    {
        return local_symbol(sec_name + char_interpol_sep + name);
    }
};*/
return true
}

func (i *inipp)____Ini()(ok bool){//col:1653
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

func (i *inipp)________return_local_symbol()(ok bool){//col:1657
/*        return local_symbol(sec_name + char_interpol_sep + name);
    }
};*/
return true
}

func (i *inipp)____Ini()(ok bool){//col:1788
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

func (i *inipp)____Ini()(ok bool){//col:1909
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

func (i *inipp)________format()(ok bool){//col:2029
/*        format(std::make_shared<Format<CharT>>()) {};
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

func (i *inipp)____Ini()(ok bool){//col:2148
/*    Ini(std::shared_ptr<Format<CharT>> fmt) :
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

func (i *inipp)________format()(ok bool){//col:2266
/*        format(fmt) {};
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

func (i *inipp)____void_generate()(ok bool){//col:2383
/*    void generate(std::basic_ostream<CharT> & os) const
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

func (i *inipp)________for_()(ok bool){//col:2498
/*        for (auto const & sec : sections)
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

func (i *inipp)____________for_()(ok bool){//col:2610
/*            for (auto const & val : sec.second)
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

func (i *inipp)____void_parse()(ok bool){//col:2715
/*    void parse(std::basic_istream<CharT> & is)
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

func (i *inipp)________while_()(ok bool){//col:2815
/*        while (std::getline(is, line))
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

func (i *inipp)____________detail::ltrim()(ok bool){//col:2913
/*            detail::ltrim(line, loc);
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

func (i *inipp)____________detail::rtrim()(ok bool){//col:3010
/*            detail::rtrim(line, loc);
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

func (i *inipp)____________const_auto_length_=_line.length()(ok bool){//col:3106
/*            const auto length = line.length();
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

func (i *inipp)____________if_()(ok bool){//col:3201
/*            if (length > 0)
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

func (i *inipp)________________const_auto___pos___=_std::find_if()(ok bool){//col:3294
/*                const auto   pos   = std::find_if(line.begin(), line.end(), [this](CharT ch) { return format->is_assign(ch); });
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

func (i *inipp)________________const_auto_&_front_=_line.front()(ok bool){//col:3386
/*                const auto & front = line.front();
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

func (i *inipp)________________if_()(ok bool){//col:3477
/*                if (format->is_comment(front))
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

func (i *inipp)________________else_if_()(ok bool){//col:3565
/*                else if (format->is_section_start(front))
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

func (i *inipp)____________________if_()(ok bool){//col:3651
/*                    if (format->is_section_end(line.back()))
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

func (i *inipp)________________________section_=_line.substr()(ok bool){//col:3736
/*                        section = line.substr(1, length - 2);
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

func (i *inipp)________________________errors.push_back()(ok bool){//col:3819
/*                        errors.push_back(line);
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

func (i *inipp)________________else_if_()(ok bool){//col:3900
/*                else if (pos != line.begin() && pos != line.end())
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

func (i *inipp)____________________String_variable()(ok bool){//col:3979
/*                    String variable(line.begin(), pos);
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

func (i *inipp)____________________String_value()(ok bool){//col:4057
/*                    String value(pos + 1, line.end());
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

func (i *inipp)____________________detail::rtrim()(ok bool){//col:4134
/*                    detail::rtrim(variable, loc);
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

func (i *inipp)____________________detail::ltrim()(ok bool){//col:4210
/*                    detail::ltrim(value, loc);
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

func (i *inipp)____________________if_()(ok bool){//col:4284
/*                    if (sec.find(variable) == sec.end())
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

func (i *inipp)________________________sec.emplace()(ok bool){//col:4357
/*                        sec.emplace(variable, value);
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

func (i *inipp)________________________errors.push_back()(ok bool){//col:4428
/*                        errors.push_back(line);
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

func (i *inipp)____________________errors.push_back()(ok bool){//col:4495
/*                    errors.push_back(line);
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

func (i *inipp)____void_interpolate()(ok bool){//col:4557
/*    void interpolate()
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

func (i *inipp)________for_()(ok bool){//col:4615
/*        for (auto & sec : sections)
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

func (i *inipp)____________replace_symbols()(ok bool){//col:4672
/*            replace_symbols(local_symbols(sec.first, sec.second), sec.second);
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

func (i *inipp)____________const_auto_syms_=_global_symbols()(ok bool){//col:4725
/*            const auto syms = global_symbols();
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

func (i *inipp)____________for_()(ok bool){//col:4777
/*            for (auto & sec : sections)
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

func (i *inipp)________________changed_|=_replace_symbols()(ok bool){//col:4828
/*                changed |= replace_symbols(syms, sec.second);
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

func (i *inipp)________}_while_()(ok bool){//col:4878
/*        } while (changed && (max_interpolation_depth > global_iteration++));
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

func (i *inipp)____void_default_section()(ok bool){//col:4926
/*    void default_section(const Section & sec)
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

func (i *inipp)________for_()(ok bool){//col:4972
/*        for (auto & sec2 : sections)
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

func (i *inipp)____________for_()(ok bool){//col:5017
/*            for (const auto & val : sec)
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

func (i *inipp)________________sec2.second.insert()(ok bool){//col:5061
/*                sec2.second.insert(val);
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

func (i *inipp)____void_strip_trailing_comments()(ok bool){//col:5103
/*    void strip_trailing_comments()
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

func (i *inipp)________for_()(ok bool){//col:5142
/*        for (auto & sec : sections)
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

func (i *inipp)____________for_()(ok bool){//col:5180
/*            for (auto & val : sec.second)
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

func (i *inipp)________________detail::rtrim2()(ok bool){//col:5216
/*                detail::rtrim2(val.second, [this](CharT ch) { return format->is_comment(ch); });
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

func (i *inipp)________________detail::rtrim()(ok bool){//col:5251
/*                detail::rtrim(val.second, loc);
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

func (i *inipp)____void_clear()(ok bool){//col:5283
/*    void clear()
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

func (i *inipp)________sections.clear()(ok bool){//col:5313
/*        sections.clear();
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

func (i *inipp)________errors.clear()(ok bool){//col:5342
/*        errors.clear();
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

func (i *inipp)____const_Symbols_local_symbols()(ok bool){//col:5367
/*    const Symbols local_symbols(const String & sec_name, const Section & sec) const
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

func (i *inipp)________for_()(ok bool){//col:5389
/*        for (const auto & val : sec)
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

func (i *inipp)____________result.emplace_back()(ok bool){//col:5410
/*            result.emplace_back(format->local_symbol(val.first), format->global_symbol(sec_name, val.first));
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

func (i *inipp)____const_Symbols_global_symbols()(ok bool){//col:5428
/*    const Symbols global_symbols() const
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

func (i *inipp)________for_()(ok bool){//col:5443
/*        for (const auto & sec : sections)
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

func (i *inipp)____________for_()(ok bool){//col:5457
/*            for (const auto & val : sec.second)
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

func (i *inipp)________________result.emplace_back()(ok bool){//col:5470
/*                result.emplace_back(format->global_symbol(sec.first, val.first), val.second);
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

func (i *inipp)____bool_replace_symbols()(ok bool){//col:5480
/*    bool replace_symbols(const Symbols & syms, Section & sec) const
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



