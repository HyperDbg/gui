extern crate alloc;

use alloc::string::String;
use alloc::string::ToString;
use alloc::vec::Vec;
use alloc::collections::BTreeMap;
use alloc::format;
use core::fmt;

#[derive(Debug, Clone, PartialEq)]
pub enum ExprError {
    UnexpectedToken(String),
    UnexpectedEnd,
    InvalidNumber(String),
    InvalidIdentifier(String),
    UndefinedVariable(String),
    TypeMismatch,
    DivisionByZero,
    InvalidExpression,
}

impl fmt::Display for ExprError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            ExprError::UnexpectedToken(t) => write!(f, "unexpected token: {}", t),
            ExprError::UnexpectedEnd => write!(f, "unexpected end of expression"),
            ExprError::InvalidNumber(n) => write!(f, "invalid number: {}", n),
            ExprError::InvalidIdentifier(i) => write!(f, "invalid identifier: {}", i),
            ExprError::UndefinedVariable(v) => write!(f, "undefined variable: {}", v),
            ExprError::TypeMismatch => write!(f, "type mismatch"),
            ExprError::DivisionByZero => write!(f, "division by zero"),
            ExprError::InvalidExpression => write!(f, "invalid expression"),
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq)]
enum TokenKind {
    Ident,
    Number,
    HexNumber,
    String,
    Eq,
    Ne,
    Lt,
    Le,
    Gt,
    Ge,
    And,
    Or,
    Not,
    Plus,
    Minus,
    Star,
    Slash,
    Percent,
    LParen,
    RParen,
    LBracket,
    RBracket,
    Dot,
    Comma,
    True,
    False,
    Nil,
    Eof,
}

#[derive(Debug, Clone)]
struct Token {
    kind: TokenKind,
    value: String,
    pos: usize,
}

struct Lexer {
    input: Vec<char>,
    pos: usize,
}

impl Lexer {
    fn new(input: &str) -> Self {
        Self {
            input: input.chars().collect(),
            pos: 0,
        }
    }

    fn peek(&self) -> Option<char> {
        self.input.get(self.pos).copied()
    }

    fn next(&mut self) -> Option<char> {
        let c = self.peek();
        self.pos += 1;
        c
    }

    fn skip_whitespace(&mut self) {
        while let Some(c) = self.peek() {
            if c.is_whitespace() {
                self.pos += 1;
            } else {
                break;
            }
        }
    }

    fn read_number(&mut self) -> Token {
        let start = self.pos;
        let mut value = String::new();
        
        while let Some(c) = self.peek() {
            if c.is_ascii_digit() || c == '_' {
                if c != '_' {
                    value.push(c);
                }
                self.pos += 1;
            } else {
                break;
            }
        }
        
        Token {
            kind: TokenKind::Number,
            value,
            pos: start,
        }
    }

    fn read_hex_number(&mut self) -> Token {
        let start = self.pos;
        let mut value = String::new();
        
        if self.peek() == Some('0') {
            value.push('0');
            self.pos += 1;
        }
        if self.peek() == Some('x') || self.peek() == Some('X') {
            value.push('x');
            self.pos += 1;
        }
        
        while let Some(c) = self.peek() {
            if c.is_ascii_hexdigit() || c == '_' {
                if c != '_' {
                    value.push(c);
                }
                self.pos += 1;
            } else {
                break;
            }
        }
        
        Token {
            kind: TokenKind::HexNumber,
            value,
            pos: start,
        }
    }

    fn read_ident(&mut self) -> Token {
        let start = self.pos;
        let mut value = String::new();
        
        while let Some(c) = self.peek() {
            if c.is_ascii_alphanumeric() || c == '_' {
                value.push(c);
                self.pos += 1;
            } else {
                break;
            }
        }
        
        let kind = match value.as_str() {
            "true" => TokenKind::True,
            "false" => TokenKind::False,
            "nil" => TokenKind::Nil,
            _ => TokenKind::Ident,
        };
        
        Token {
            kind,
            value,
            pos: start,
        }
    }

    fn read_string(&mut self, quote: char) -> Token {
        let start = self.pos;
        let mut value = String::new();
        
        self.next();
        
        while let Some(c) = self.peek() {
            if c == quote {
                self.pos += 1;
                break;
            } else if c == '\\' {
                self.pos += 1;
                if let Some(escaped) = self.peek() {
                    match escaped {
                        'n' => value.push('\n'),
                        't' => value.push('\t'),
                        'r' => value.push('\r'),
                        '\\' => value.push('\\'),
                        '"' => value.push('"'),
                        '\'' => value.push('\''),
                        _ => value.push(escaped),
                    }
                    self.pos += 1;
                }
            } else {
                value.push(c);
                self.pos += 1;
            }
        }
        
        Token {
            kind: TokenKind::String,
            value,
            pos: start,
        }
    }

    fn next_token(&mut self) -> Token {
        self.skip_whitespace();
        
        let start = self.pos;
        
        match self.peek() {
            None => Token {
                kind: TokenKind::Eof,
                value: String::new(),
                pos: start,
            },
            Some(c) if c.is_ascii_digit() => {
                if c == '0' && self.input.get(self.pos + 1).map(|&c| c == 'x' || c == 'X').unwrap_or(false) {
                    self.read_hex_number()
                } else {
                    self.read_number()
                }
            }
            Some(c) if c.is_ascii_alphabetic() || c == '_' => self.read_ident(),
            Some('"') | Some('\'') => {
                let quote = self.next().unwrap();
                self.read_string(quote)
            }
            Some('=') => {
                self.pos += 1;
                if self.peek() == Some('=') {
                    self.pos += 1;
                    Token { kind: TokenKind::Eq, value: "==".to_string(), pos: start }
                } else {
                    Token { kind: TokenKind::Eq, value: "=".to_string(), pos: start }
                }
            }
            Some('!') => {
                self.pos += 1;
                if self.peek() == Some('=') {
                    self.pos += 1;
                    Token { kind: TokenKind::Ne, value: "!=".to_string(), pos: start }
                } else {
                    Token { kind: TokenKind::Not, value: "!".to_string(), pos: start }
                }
            }
            Some('<') => {
                self.pos += 1;
                if self.peek() == Some('=') {
                    self.pos += 1;
                    Token { kind: TokenKind::Le, value: "<=".to_string(), pos: start }
                } else {
                    Token { kind: TokenKind::Lt, value: "<".to_string(), pos: start }
                }
            }
            Some('>') => {
                self.pos += 1;
                if self.peek() == Some('=') {
                    self.pos += 1;
                    Token { kind: TokenKind::Ge, value: ">=".to_string(), pos: start }
                } else {
                    Token { kind: TokenKind::Gt, value: ">".to_string(), pos: start }
                }
            }
            Some('&') => {
                self.pos += 1;
                if self.peek() == Some('&') {
                    self.pos += 1;
                    Token { kind: TokenKind::And, value: "&&".to_string(), pos: start }
                } else {
                    Token { kind: TokenKind::And, value: "&".to_string(), pos: start }
                }
            }
            Some('|') => {
                self.pos += 1;
                if self.peek() == Some('|') {
                    self.pos += 1;
                    Token { kind: TokenKind::Or, value: "||".to_string(), pos: start }
                } else {
                    Token { kind: TokenKind::Or, value: "|".to_string(), pos: start }
                }
            }
            Some('+') => {
                self.pos += 1;
                Token { kind: TokenKind::Plus, value: "+".to_string(), pos: start }
            }
            Some('-') => {
                self.pos += 1;
                Token { kind: TokenKind::Minus, value: "-".to_string(), pos: start }
            }
            Some('*') => {
                self.pos += 1;
                Token { kind: TokenKind::Star, value: "*".to_string(), pos: start }
            }
            Some('/') => {
                self.pos += 1;
                Token { kind: TokenKind::Slash, value: "/".to_string(), pos: start }
            }
            Some('%') => {
                self.pos += 1;
                Token { kind: TokenKind::Percent, value: "%".to_string(), pos: start }
            }
            Some('(') => {
                self.pos += 1;
                Token { kind: TokenKind::LParen, value: "(".to_string(), pos: start }
            }
            Some(')') => {
                self.pos += 1;
                Token { kind: TokenKind::RParen, value: ")".to_string(), pos: start }
            }
            Some('[') => {
                self.pos += 1;
                Token { kind: TokenKind::LBracket, value: "[".to_string(), pos: start }
            }
            Some(']') => {
                self.pos += 1;
                Token { kind: TokenKind::RBracket, value: "]".to_string(), pos: start }
            }
            Some('.') => {
                self.pos += 1;
                Token { kind: TokenKind::Dot, value: ".".to_string(), pos: start }
            }
            Some(',') => {
                self.pos += 1;
                Token { kind: TokenKind::Comma, value: ",".to_string(), pos: start }
            }
            Some(c) => {
                self.pos += 1;
                Token {
                    kind: TokenKind::Ident,
                    value: c.to_string(),
                    pos: start,
                }
            }
        }
    }
}

#[derive(Debug, Clone)]
pub enum Value {
    Bool(bool),
    Int(i64),
    Uint(u64),
    Float(f64),
    String(String),
    Nil,
}

impl Value {
    pub fn as_bool(&self) -> Result<bool, ExprError> {
        match self {
            Value::Bool(b) => Ok(*b),
            _ => Err(ExprError::TypeMismatch),
        }
    }

    pub fn as_int(&self) -> Result<i64, ExprError> {
        match self {
            Value::Int(i) => Ok(*i),
            Value::Uint(u) => Ok(*u as i64),
            Value::Float(f) => Ok(*f as i64),
            _ => Err(ExprError::TypeMismatch),
        }
    }

    pub fn as_uint(&self) -> Result<u64, ExprError> {
        match self {
            Value::Uint(u) => Ok(*u),
            Value::Int(i) => Ok(*i as u64),
            Value::Float(f) => Ok(*f as u64),
            _ => Err(ExprError::TypeMismatch),
        }
    }

    pub fn as_float(&self) -> Result<f64, ExprError> {
        match self {
            Value::Float(f) => Ok(*f),
            Value::Int(i) => Ok(*i as f64),
            Value::Uint(u) => Ok(*u as f64),
            _ => Err(ExprError::TypeMismatch),
        }
    }

    pub fn as_string(&self) -> Result<&str, ExprError> {
        match self {
            Value::String(s) => Ok(s),
            _ => Err(ExprError::TypeMismatch),
        }
    }

    pub fn is_truthy(&self) -> bool {
        match self {
            Value::Bool(b) => *b,
            Value::Int(i) => *i != 0,
            Value::Uint(u) => *u != 0,
            Value::Float(f) => *f != 0.0,
            Value::String(s) => !s.is_empty(),
            Value::Nil => false,
        }
    }

    pub fn type_name(&self) -> &'static str {
        match self {
            Value::Bool(_) => "bool",
            Value::Int(_) => "int",
            Value::Uint(_) => "uint",
            Value::Float(_) => "float",
            Value::String(_) => "string",
            Value::Nil => "nil",
        }
    }
}

impl fmt::Display for Value {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            Value::Bool(b) => write!(f, "{}", b),
            Value::Int(i) => write!(f, "{}", i),
            Value::Uint(u) => write!(f, "{}", u),
            Value::Float(fl) => write!(f, "{}", fl),
            Value::String(s) => write!(f, "\"{}\"", s),
            Value::Nil => write!(f, "nil"),
        }
    }
}

impl PartialEq for Value {
    fn eq(&self, other: &Self) -> bool {
        match (self, other) {
            (Value::Bool(a), Value::Bool(b)) => a == b,
            (Value::Int(a), Value::Int(b)) => a == b,
            (Value::Int(a), Value::Uint(b)) => *a as u64 == *b,
            (Value::Uint(a), Value::Uint(b)) => a == b,
            (Value::Uint(a), Value::Int(b)) => *a == *b as u64,
            (Value::Float(a), Value::Float(b)) => a == b,
            (Value::String(a), Value::String(b)) => a == b,
            (Value::Nil, Value::Nil) => true,
            _ => false,
        }
    }
}

impl PartialOrd for Value {
    fn partial_cmp(&self, other: &Self) -> Option<core::cmp::Ordering> {
        match (self, other) {
            (Value::Int(a), Value::Int(b)) => a.partial_cmp(b),
            (Value::Int(a), Value::Uint(b)) => (*a as u64).partial_cmp(b),
            (Value::Uint(a), Value::Uint(b)) => a.partial_cmp(b),
            (Value::Uint(a), Value::Int(b)) => a.partial_cmp(&(*b as u64)),
            (Value::Float(a), Value::Float(b)) => a.partial_cmp(b),
            (Value::String(a), Value::String(b)) => a.partial_cmp(b),
            _ => None,
        }
    }
}

pub type Variables = BTreeMap<String, Value>;

pub struct ExprEvaluator {
    lexer: Lexer,
    current: Token,
}

impl ExprEvaluator {
    pub fn new(input: &str) -> Self {
        let mut lexer = Lexer::new(input);
        let current = lexer.next_token();
        Self { lexer, current }
    }

    fn advance(&mut self) {
        self.current = self.lexer.next_token();
    }

    fn expect(&mut self, kind: TokenKind) -> Result<(), ExprError> {
        if self.current.kind == kind {
            self.advance();
            Ok(())
        } else {
            Err(ExprError::UnexpectedToken(self.current.value.clone()))
        }
    }

    pub fn eval(&mut self, vars: &Variables) -> Result<Value, ExprError> {
        self.parse_or(vars)
    }

    fn parse_or(&mut self, vars: &Variables) -> Result<Value, ExprError> {
        let mut left = self.parse_and(vars)?;
        
        while self.current.kind == TokenKind::Or {
            self.advance();
            let right = self.parse_and(vars)?;
            left = Value::Bool(left.as_bool()? || right.as_bool()?);
        }
        
        Ok(left)
    }

    fn parse_and(&mut self, vars: &Variables) -> Result<Value, ExprError> {
        let mut left = self.parse_equality(vars)?;
        
        while self.current.kind == TokenKind::And {
            self.advance();
            let right = self.parse_equality(vars)?;
            left = Value::Bool(left.as_bool()? && right.as_bool()?);
        }
        
        Ok(left)
    }

    fn parse_equality(&mut self, vars: &Variables) -> Result<Value, ExprError> {
        let mut left = self.parse_comparison(vars)?;
        
        loop {
            match self.current.kind {
                TokenKind::Eq => {
                    self.advance();
                    let right = self.parse_comparison(vars)?;
                    left = Value::Bool(left == right);
                }
                TokenKind::Ne => {
                    self.advance();
                    let right = self.parse_comparison(vars)?;
                    left = Value::Bool(left != right);
                }
                _ => break,
            }
        }
        
        Ok(left)
    }

    fn parse_comparison(&mut self, vars: &Variables) -> Result<Value, ExprError> {
        let mut left = self.parse_additive(vars)?;
        
        loop {
            match self.current.kind {
                TokenKind::Lt => {
                    self.advance();
                    let right = self.parse_additive(vars)?;
                    left = Value::Bool(left.partial_cmp(&right) == Some(core::cmp::Ordering::Less));
                }
                TokenKind::Le => {
                    self.advance();
                    let right = self.parse_additive(vars)?;
                    left = Value::Bool(matches!(left.partial_cmp(&right), Some(core::cmp::Ordering::Less | core::cmp::Ordering::Equal)));
                }
                TokenKind::Gt => {
                    self.advance();
                    let right = self.parse_additive(vars)?;
                    left = Value::Bool(left.partial_cmp(&right) == Some(core::cmp::Ordering::Greater));
                }
                TokenKind::Ge => {
                    self.advance();
                    let right = self.parse_additive(vars)?;
                    left = Value::Bool(matches!(left.partial_cmp(&right), Some(core::cmp::Ordering::Greater | core::cmp::Ordering::Equal)));
                }
                _ => break,
            }
        }
        
        Ok(left)
    }

    fn parse_additive(&mut self, vars: &Variables) -> Result<Value, ExprError> {
        let mut left = self.parse_multiplicative(vars)?;
        
        loop {
            match self.current.kind {
                TokenKind::Plus => {
                    self.advance();
                    let right = self.parse_multiplicative(vars)?;
                    left = self.add(left, right)?;
                }
                TokenKind::Minus => {
                    self.advance();
                    let right = self.parse_multiplicative(vars)?;
                    left = self.sub(left, right)?;
                }
                _ => break,
            }
        }
        
        Ok(left)
    }

    fn parse_multiplicative(&mut self, vars: &Variables) -> Result<Value, ExprError> {
        let mut left = self.parse_unary(vars)?;
        
        loop {
            match self.current.kind {
                TokenKind::Star => {
                    self.advance();
                    let right = self.parse_unary(vars)?;
                    left = self.mul(left, right)?;
                }
                TokenKind::Slash => {
                    self.advance();
                    let right = self.parse_unary(vars)?;
                    left = self.div(left, right)?;
                }
                TokenKind::Percent => {
                    self.advance();
                    let right = self.parse_unary(vars)?;
                    left = self.rem(left, right)?;
                }
                _ => break,
            }
        }
        
        Ok(left)
    }

    fn parse_unary(&mut self, vars: &Variables) -> Result<Value, ExprError> {
        match self.current.kind {
            TokenKind::Not => {
                self.advance();
                let val = self.parse_unary(vars)?;
                Ok(Value::Bool(!val.as_bool()?))
            }
            TokenKind::Minus => {
                self.advance();
                let val = self.parse_unary(vars)?;
                match val {
                    Value::Int(i) => Ok(Value::Int(-i)),
                    Value::Float(f) => Ok(Value::Float(-f)),
                    _ => Err(ExprError::TypeMismatch),
                }
            }
            _ => self.parse_primary(vars),
        }
    }

    fn parse_primary(&mut self, vars: &Variables) -> Result<Value, ExprError> {
        match self.current.kind {
            TokenKind::True => {
                self.advance();
                Ok(Value::Bool(true))
            }
            TokenKind::False => {
                self.advance();
                Ok(Value::Bool(false))
            }
            TokenKind::Nil => {
                self.advance();
                Ok(Value::Nil)
            }
            TokenKind::Number => {
                let value = self.current.value.clone();
                self.advance();
                let num = value.parse::<i64>().map_err(|_| ExprError::InvalidNumber(value))?;
                Ok(Value::Int(num))
            }
            TokenKind::HexNumber => {
                let value = self.current.value.clone();
                self.advance();
                let hex_str = value.trim_start_matches("0x").trim_start_matches("0X");
                let num = u64::from_str_radix(hex_str, 16).map_err(|_| ExprError::InvalidNumber(value))?;
                Ok(Value::Uint(num))
            }
            TokenKind::String => {
                let value = self.current.value.clone();
                self.advance();
                Ok(Value::String(value))
            }
            TokenKind::Ident => {
                let name = self.current.value.clone();
                self.advance();
                
                if self.current.kind == TokenKind::LParen {
                    return self.parse_function_call(&name, vars);
                }
                
                vars.get(&name)
                    .cloned()
                    .ok_or_else(|| ExprError::UndefinedVariable(name))
            }
            TokenKind::LParen => {
                self.advance();
                let val = self.parse_or(vars)?;
                self.expect(TokenKind::RParen)?;
                Ok(val)
            }
            _ => Err(ExprError::UnexpectedToken(self.current.value.clone())),
        }
    }

    fn parse_function_call(&mut self, _name: &str, _vars: &Variables) -> Result<Value, ExprError> {
        self.expect(TokenKind::LParen)?;
        
        let mut args = Vec::new();
        if self.current.kind != TokenKind::RParen {
            loop {
                args.push(self.parse_or(_vars)?);
                if self.current.kind == TokenKind::Comma {
                    self.advance();
                } else {
                    break;
                }
            }
        }
        
        self.expect(TokenKind::RParen)?;
        
        Ok(Value::Nil)
    }

    fn add(&self, a: Value, b: Value) -> Result<Value, ExprError> {
        match (&a, &b) {
            (Value::Int(x), Value::Int(y)) => Ok(Value::Int(x + y)),
            (Value::Int(x), Value::Uint(y)) => Ok(Value::Int(x + *y as i64)),
            (Value::Uint(x), Value::Uint(y)) => Ok(Value::Uint(x + y)),
            (Value::Uint(x), Value::Int(y)) => Ok(Value::Uint(x + *y as u64)),
            (Value::Float(x), Value::Float(y)) => Ok(Value::Float(x + y)),
            (Value::String(x), Value::String(y)) => Ok(Value::String(format!("{}{}", x, y))),
            _ => Err(ExprError::TypeMismatch),
        }
    }

    fn sub(&self, a: Value, b: Value) -> Result<Value, ExprError> {
        match (&a, &b) {
            (Value::Int(x), Value::Int(y)) => Ok(Value::Int(x - y)),
            (Value::Int(x), Value::Uint(y)) => Ok(Value::Int(x - *y as i64)),
            (Value::Uint(x), Value::Uint(y)) => Ok(Value::Uint(x - y)),
            (Value::Uint(x), Value::Int(y)) => Ok(Value::Uint(x - *y as u64)),
            (Value::Float(x), Value::Float(y)) => Ok(Value::Float(x - y)),
            _ => Err(ExprError::TypeMismatch),
        }
    }

    fn mul(&self, a: Value, b: Value) -> Result<Value, ExprError> {
        match (&a, &b) {
            (Value::Int(x), Value::Int(y)) => Ok(Value::Int(x * y)),
            (Value::Int(x), Value::Uint(y)) => Ok(Value::Int(x * *y as i64)),
            (Value::Uint(x), Value::Uint(y)) => Ok(Value::Uint(x * y)),
            (Value::Uint(x), Value::Int(y)) => Ok(Value::Uint(x * *y as u64)),
            (Value::Float(x), Value::Float(y)) => Ok(Value::Float(x * y)),
            _ => Err(ExprError::TypeMismatch),
        }
    }

    fn div(&self, a: Value, b: Value) -> Result<Value, ExprError> {
        match (&a, &b) {
            (Value::Int(x), Value::Int(y)) => {
                if *y == 0 {
                    return Err(ExprError::DivisionByZero);
                }
                Ok(Value::Int(x / y))
            }
            (Value::Uint(x), Value::Uint(y)) => {
                if *y == 0 {
                    return Err(ExprError::DivisionByZero);
                }
                Ok(Value::Uint(x / y))
            }
            (Value::Float(x), Value::Float(y)) => {
                if *y == 0.0 {
                    return Err(ExprError::DivisionByZero);
                }
                Ok(Value::Float(x / y))
            }
            _ => Err(ExprError::TypeMismatch),
        }
    }

    fn rem(&self, a: Value, b: Value) -> Result<Value, ExprError> {
        match (&a, &b) {
            (Value::Int(x), Value::Int(y)) => {
                if *y == 0 {
                    return Err(ExprError::DivisionByZero);
                }
                Ok(Value::Int(x % y))
            }
            (Value::Uint(x), Value::Uint(y)) => {
                if *y == 0 {
                    return Err(ExprError::DivisionByZero);
                }
                Ok(Value::Uint(x % y))
            }
            _ => Err(ExprError::TypeMismatch),
        }
    }
}

pub fn eval_expr(input: &str, vars: &Variables) -> Result<Value, ExprError> {
    let mut evaluator = ExprEvaluator::new(input);
    evaluator.eval(vars)
}

pub fn eval_bool(input: &str, vars: &Variables) -> Result<bool, ExprError> {
    eval_expr(input, vars)?.as_bool()
}

pub struct HookFilter {
    condition: String,
    compiled: Option<CompiledFilter>,
}

#[derive(Debug, Clone)]
struct CompiledFilter {
    tokens: Vec<Token>,
}

impl HookFilter {
    pub fn new(condition: &str) -> Self {
        Self {
            condition: condition.to_string(),
            compiled: None,
        }
    }

    pub fn eval(&self, vars: &Variables) -> Result<bool, ExprError> {
        eval_bool(&self.condition, vars)
    }

    pub fn condition(&self) -> &str {
        &self.condition
    }
}

#[derive(Debug, Clone)]
pub struct HookArgs {
    pub name: String,
    pub values: BTreeMap<String, Value>,
}

impl HookArgs {
    pub fn new(name: &str) -> Self {
        Self {
            name: name.to_string(),
            values: BTreeMap::new(),
        }
    }

    pub fn with_uint(mut self, name: &str, value: u64) -> Self {
        self.values.insert(name.to_string(), Value::Uint(value));
        self
    }

    pub fn with_int(mut self, name: &str, value: i64) -> Self {
        self.values.insert(name.to_string(), Value::Int(value));
        self
    }

    pub fn with_bool(mut self, name: &str, value: bool) -> Self {
        self.values.insert(name.to_string(), Value::Bool(value));
        self
    }

    pub fn with_string(mut self, name: &str, value: &str) -> Self {
        self.values.insert(name.to_string(), Value::String(value.to_string()));
        self
    }

    pub fn to_variables(&self) -> Variables {
        self.values.clone()
    }
}

pub fn eval_hook_filter(condition: &str, args: &HookArgs) -> Result<bool, ExprError> {
    let vars = args.to_variables();
    eval_bool(condition, &vars)
}

#[macro_export]
macro_rules! hook_args {
    ($name:expr, $($key:expr => $value:expr),* $(,)?) => {{
        let mut args = $crate::expr::HookArgs::new($name);
        $(
            args = args.with_uint($key, $value as u64);
        )*
        args
    }};
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_basic_arithmetic() {
        let vars = Variables::new();
        
        assert_eq!(eval_expr("1 + 2", &vars).unwrap(), Value::Int(3));
        assert_eq!(eval_expr("10 - 3", &vars).unwrap(), Value::Int(7));
        assert_eq!(eval_expr("4 * 5", &vars).unwrap(), Value::Int(20));
        assert_eq!(eval_expr("20 / 4", &vars).unwrap(), Value::Int(5));
        assert_eq!(eval_expr("17 % 5", &vars).unwrap(), Value::Int(2));
    }

    #[test]
    fn test_comparison() {
        let vars = Variables::new();
        
        assert_eq!(eval_expr("1 == 1", &vars).unwrap(), Value::Bool(true));
        assert_eq!(eval_expr("1 != 2", &vars).unwrap(), Value::Bool(true));
        assert_eq!(eval_expr("1 < 2", &vars).unwrap(), Value::Bool(true));
        assert_eq!(eval_expr("2 <= 2", &vars).unwrap(), Value::Bool(true));
        assert_eq!(eval_expr("3 > 2", &vars).unwrap(), Value::Bool(true));
        assert_eq!(eval_expr("3 >= 3", &vars).unwrap(), Value::Bool(true));
    }

    #[test]
    fn test_logical() {
        let vars = Variables::new();
        
        assert_eq!(eval_expr("true && true", &vars).unwrap(), Value::Bool(true));
        assert_eq!(eval_expr("true && false", &vars).unwrap(), Value::Bool(false));
        assert_eq!(eval_expr("false || true", &vars).unwrap(), Value::Bool(true));
        assert_eq!(eval_expr("!false", &vars).unwrap(), Value::Bool(true));
        assert_eq!(eval_expr("!(1 == 2)", &vars).unwrap(), Value::Bool(true));
    }

    #[test]
    fn test_hex_numbers() {
        let vars = Variables::new();
        
        assert_eq!(eval_expr("0x10", &vars).unwrap(), Value::Uint(16));
        assert_eq!(eval_expr("0xFF", &vars).unwrap(), Value::Uint(255));
        assert_eq!(eval_expr("0x1234", &vars).unwrap(), Value::Uint(0x1234));
    }

    #[test]
    fn test_variables() {
        let mut vars = Variables::new();
        vars.insert("IoControlCode".to_string(), Value::Uint(0x12345678));
        vars.insert("InputBufferLength".to_string(), Value::Uint(100));
        
        assert_eq!(eval_expr("IoControlCode == 0x12345678", &vars).unwrap(), Value::Bool(true));
        assert_eq!(eval_expr("InputBufferLength > 50", &vars).unwrap(), Value::Bool(true));
        assert_eq!(eval_expr("IoControlCode == 0x12345678 && InputBufferLength > 0", &vars).unwrap(), Value::Bool(true));
    }

    #[test]
    fn test_complex_expression() {
        let mut vars = Variables::new();
        vars.insert("x".to_string(), Value::Int(10));
        vars.insert("y".to_string(), Value::Int(20));
        
        assert_eq!(eval_expr("(x + y) * 2", &vars).unwrap(), Value::Int(60));
        assert_eq!(eval_expr("x > 5 && y < 30", &vars).unwrap(), Value::Bool(true));
        assert_eq!(eval_expr("!(x == y)", &vars).unwrap(), Value::Bool(true));
    }

    #[test]
    fn test_hook_filter_example() {
        let mut vars = Variables::new();
        vars.insert("IoControlCode".to_string(), Value::Uint(0x12345678));
        vars.insert("InputBufferLength".to_string(), Value::Uint(256));
        vars.insert("OutputBufferLength".to_string(), Value::Uint(512));
        
        let result = eval_bool(
            "IoControlCode == 0x12345678 && InputBufferLength > 0",
            &vars
        ).unwrap();
        assert!(result);
        
        let result = eval_bool(
            "IoControlCode == 0x12345678 && InputBufferLength > 0 && OutputBufferLength > 0",
            &vars
        ).unwrap();
        assert!(result);
    }

    #[test]
    fn test_hook_args() {
        let args = HookArgs::new("NtDeviceIoControlFile")
            .with_uint("IoControlCode", 0x12345678)
            .with_uint("InputBufferLength", 256)
            .with_uint("OutputBufferLength", 512);
        
        let result = eval_hook_filter(
            "IoControlCode == 0x12345678 && InputBufferLength > 0",
            &args
        ).unwrap();
        assert!(result);
    }

    #[test]
    fn test_hook_filter_struct() {
        let filter = HookFilter::new("IoControlCode == 0x12345678 && InputBufferLength > 0");
        
        let args = HookArgs::new("NtDeviceIoControlFile")
            .with_uint("IoControlCode", 0x12345678)
            .with_uint("InputBufferLength", 256);
        
        assert!(filter.eval(&args.to_variables()).unwrap());
    }

    #[test]
    fn test_string_operations() {
        let vars = Variables::new();
        
        assert_eq!(eval_expr("\"hello\" == \"hello\"", &vars).unwrap(), Value::Bool(true));
        assert_eq!(eval_expr("\"hello\" != \"world\"", &vars).unwrap(), Value::Bool(true));
    }

    #[test]
    fn test_parentheses() {
        let vars = Variables::new();
        
        assert_eq!(eval_expr("(1 + 2) * 3", &vars).unwrap(), Value::Int(9));
        assert_eq!(eval_expr("1 + (2 * 3)", &vars).unwrap(), Value::Int(7));
        assert_eq!(eval_expr("((1 + 2) * 3)", &vars).unwrap(), Value::Int(9));
    }

    #[test]
    fn test_unary_minus() {
        let vars = Variables::new();
        
        assert_eq!(eval_expr("-5", &vars).unwrap(), Value::Int(-5));
        assert_eq!(eval_expr("-(-5)", &vars).unwrap(), Value::Int(5));
        assert_eq!(eval_expr("10 + -5", &vars).unwrap(), Value::Int(5));
    }

    #[test]
    fn test_boolean_literals() {
        let vars = Variables::new();
        
        assert_eq!(eval_expr("true", &vars).unwrap(), Value::Bool(true));
        assert_eq!(eval_expr("false", &vars).unwrap(), Value::Bool(false));
        assert_eq!(eval_expr("true == true", &vars).unwrap(), Value::Bool(true));
    }

    #[test]
    fn test_nil() {
        let vars = Variables::new();
        
        assert_eq!(eval_expr("nil", &vars).unwrap(), Value::Nil);
        assert_eq!(eval_expr("nil == nil", &vars).unwrap(), Value::Bool(true));
    }

    #[test]
    fn test_error_undefined_variable() {
        let vars = Variables::new();
        
        match eval_expr("undefined_var", &vars) {
            Err(ExprError::UndefinedVariable(name)) => assert_eq!(name, "undefined_var"),
            _ => panic!("Expected UndefinedVariable error"),
        }
    }

    #[test]
    fn test_error_division_by_zero() {
        let vars = Variables::new();
        
        match eval_expr("1 / 0", &vars) {
            Err(ExprError::DivisionByZero) => (),
            _ => panic!("Expected DivisionByZero error"),
        }
    }

    #[test]
    fn test_operator_precedence() {
        let vars = Variables::new();
        
        assert_eq!(eval_expr("1 + 2 * 3", &vars).unwrap(), Value::Int(7));
        assert_eq!(eval_expr("1 * 2 + 3", &vars).unwrap(), Value::Int(5));
        assert_eq!(eval_expr("1 + 2 < 3 + 1", &vars).unwrap(), Value::Bool(true));
        assert_eq!(eval_expr("true || false && false", &vars).unwrap(), Value::Bool(true));
    }
}
