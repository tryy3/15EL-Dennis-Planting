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
    return String(x,14); // Eftersom en LCD bara kan ha 16 karaktärer så om man har decimaler, så kommer man bara ha 14 decimaler i strängen.
}



Eval::Eval(String str) {
    _str = str;
}

double Eval::parse() {
    nextChar();
    double result = parseExpression();
    if (_pos < _str.length()) {
        return -1;
    }
    return result;
}

void Eval::nextChar() {
    // Leta efter nästa Char.
    if (++_pos < _str.length()) {
        _ch = _str.charAt(_pos);
    } else {
        _ch = -1;
    }
}

bool Eval::eat(int charToEat) {
    // Hitta nästa Char
    while (_ch == ' ') {
        nextChar();
    }
    // Om Char är samma som charToEat, leta efter nästa och return med true.
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
//
// Hur denna klass funkar, beskrivs i dokumentationen i README
double Eval::parseExpression() {
    double result = parseTerm();
    while(true) {
        if (eat('+')) { // Addition
            result += parseTerm();
        } else if (eat('-')) { // Subtraktion;
            result -= parseTerm();
        } else {
            return result;
        }
    }
}

double Eval::parseTerm() {
    double result = parseFactor();
    while(true) {
        if (eat('*')) { // Multiplikation
            result *= parseFactor();
        } else if (eat('/')) { // Division;
            result /= parseFactor();
        } else {
            return result;
        }
    }
}

double Eval::parseFactor() {
    // Mer om unary https://msdn.microsoft.com/en-us/library/ewkkxkwb.aspx
    if (eat('+')) { // Unary plus
        return parseFactor();
    }
    if (eat('-')) { // Unary minus
        return -parseFactor();
    }

    double result;
    int startPos = _pos;

    // Paranteser
    if (eat('(')) {
        result = parseExpression();
        eat(')');
    }
    // Nummer
    else if ((_ch >= '0' && _ch <= '9') || _ch == '.') {
        while((_ch >= '0' && _ch <= '9') || _ch == '.')  { nextChar(); }
        result = atof(_str.substring(startPos, _pos).c_str());
    } 
    // Funkltioner
    else if (_ch >= 'a' && _ch <= 'z') {
        while(_ch >= 'a' && _ch <= 'z')  { nextChar(); }
        String func = _str.substring(startPos, _pos);
       result = parseFactor();
        if (func.equals("sqrt")) { result = sqrt(result); }
        else if (func.equals("sin")) { result = sin(result); }
        else if (func.equals("cos")) { result = cos(result); }
        else if (func.equals("tan")) { result = tan(result); }
        else { return -1; } // Unknown funktion
    } else {
        return -1; // Unknown Char
    }

    if (eat('^')) { result = pow(result, parseFactor()); } // Exponentiering

    return result;
}