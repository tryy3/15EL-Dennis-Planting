/*
 * Calculator.cpp - Library for creating calculators.
 * Created by Dennis Planting, 21 January, 2017.
 */

#include "Calculator.h"
#include "Arduino.h"
#include <StandardCplusplus.h>
#include <vector>
#include <math.h>

using namespace std;

Calculator::Calculator() {
    _buffer = "";
}

void Calculator::addChar(char c) {
    _buffer += c;
}

void Calculator::addWord(String str) {
    _buffer += str;
}

void Calculator::delChar() {
    if (_buffer.length() <= 0) { return; }
    _buffer = _buffer.substring(0, _buffer.length()-1);
}
void Calculator::clear() {
    _buffer = "";
}

String Calculator::getBuffer() {
    return _buffer;
}

String Calculator::evaluate() {
    Eval eval(_buffer);
    double x = eval.parse();
    return String(x,14); // LCD can only print 16 char, x.(14 decimals left)
}

Eval::Eval(String str) {
    _str = str;
}

double Eval::parse() {
    nextChar();
    double x = parseExpression();
    if (_pos < _str.length()) {
        return -1;
    }
    return x;
}

void Eval::nextChar() {
    if (++_pos < _str.length()) {
        _ch = _str.charAt(_pos);
    } else {
        _ch = -1;
    }
}

bool Eval::eat(int charToEat) {
    while (_ch == ' ') {
        nextChar();
    }
    if (_ch == charToEat) {
        nextChar();
        return true;
    }
    return false;
}

// Grammar:
// expression = term | expression `+` term | expression `-` term
// term = factor | term `*` factor | term `/` factor
// factor = `+` factor | `-` factor | `(` expression `)`
//        | number | functionName factor | factor `^` factor
double Eval::parseExpression() {
    double x = parseTerm();
    while(true) {
        if (eat('+')) { // Addition
            x += parseTerm();
        } else if (eat('-')) { // subtraction;
            x -= parseTerm();
        } else {
            return x;
        }
    }
}

double Eval::parseTerm() {
    double x = parseFactor();
    while(true) {
        if (eat('*')) { // Multiplication
            x *= parseFactor();
        } else if (eat('/')) { // Subtraction;
            x /= parseFactor();
        } else {
            return x;
        }
    }
}

double Eval::parseFactor() {
    if (eat('+')) { // Unary plus
        return parseFactor();
    }
    if (eat('-')) { // Unary minus
        return -parseFactor();
    }

    double x;
    int startPos = _pos;
    if (eat('(')) { // Parantheses
        x = parseExpression();
        eat(')');
    } else if ((_ch >= '0' && _ch <= '9') || _ch == '.') { // Numbers
        while((_ch >= '0' && _ch <= '9') || _ch == '.')  { nextChar(); }
        x = atof(_str.substring(startPos, _pos).c_str());
    } else if (_ch >= 'a' && _ch <= 'z') {
        while(_ch >= 'a' && _ch <= 'z')  { nextChar(); }
        String func = _str.substring(startPos, _pos);
        x = parseFactor();
        if (func.equals("sqrt")) { x = sqrt(x); }
        else if (func.equals("sin")) { x = sin(x); }
        else if (func.equals("cos")) { x = cos(x); }
        else if (func.equals("tan")) { x = tan(x); }
        else { return -1; } // unknown function
    } else {
        return -1; // unknown char
    }

    if (eat('^')) { x = pow(x, parseFactor()); } // Exponentiation.

    return x;
}