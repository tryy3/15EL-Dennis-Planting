/*
 * Calculator.cpp - Library for creating calculators.
 * Created by Dennis Planting, 21 January, 2017.
 */

/** Error table
 *
 * 101 = First value is not a number.
 * 102 = Last value is not a number.
 * 103 = Next value is not a number.
 */

#include "Calculator.h"
#include "Arduino.h"
#include <StandardCplusplus.h>
#include <vector>

using namespace std;

Value::Value(Method method, String value) {
    _method = method;
    _value = value;
}
Method Value::getMethod() {
    return _method;
}
String Value::getValue() {
    return _value;
}


Calculator::Calculator() {
    _buffer = "";
}

void Calculator::addChar(char c) {
    _buffer += c;
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
    vector<Value> values;

    String val = "";
    bool push_value = false;
    for (int i = 0; i < _buffer.length(); i++) {
        char c = _buffer.charAt(i);

        if (c == '+' || c == '-' ||
            c == '*' || c == '/' ||
            c == '^' || c == '%') {

            Value value(Method::NUMBER, val);
            values.push_back(value);
            val = "";

            if (c == '+') {
                Value method(Method::PLUS, "+");
                values.push_back(method);
            } else if (c == '-') {
                Value method(Method::MINUS, "-");
                values.push_back(method);
            } else if (c == '*') {
                Value method(Method::MULTIPLY, "*");
                values.push_back(method);
            } else if (c == '/') {
                Value method(Method::DIVIDE, "/");
                values.push_back(method);
            } else if (c == '^') {
                Value method(Method::POWER, "^");
                values.push_back(method);
            }
        } else {
            val += c;
        }
    }

    if (val != "") {
        Value value(Method::NUMBER, val);
        values.push_back(value);
        val = "";
    }

    // Check if first and last value is numbers.
    if (values.at(0).getMethod() != Method::NUMBER) {
        return "Error 101"; // First value is not a number
    }
    if (values.at(values.size()-1).getMethod() != Method::NUMBER){
        return "Error 102"; // Second value is not a number
    }

    float v = values.at(0).getValue().toFloat(); // The current value.
    for (int i = 1; i < values.size()-1; i+=2) {
        Method method = values.at(i).getMethod();

        if (values.at(i+1).getMethod() != Method::NUMBER) {
            return "Error 103"; // Next value is not a number.
        }

        float n = values.at(i+1).getValue().toFloat(); // Next value.
        
        if (method == Method::PLUS) {
            v += n;
        } else if (method == Method::MINUS) {
            v -= n;
        } else if (method == Method::DIVIDE) {
            v /= n;
        } else if (method == Method::MULTIPLY) {
            v *= n;
        } else if (method == Method::POWER) {
            v = pow(v, n);
        }
    }
    return String(v);
}