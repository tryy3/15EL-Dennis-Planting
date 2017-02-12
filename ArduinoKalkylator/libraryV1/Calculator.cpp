/*
 * Calculator.cpp - Library for creating calculators.
 * Created by Dennis Planting, 21 January, 2017.
 */

/** Error table
 *
 * 101 = Första Value är inte ett nummer.
 * 102 = Sista Value är inte ett nummer.
 * 103 = Nästa Value är inte ett nummer.
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

    String nummer = "";
    for (int i = 0; i < _buffer.length(); i++) {
        char c = _buffer.charAt(i);

        if (c == '+' || c == '-' ||
            c == '*' || c == '/' ||
            c == '^' || c == '%') {

            // Skapa en ny Value från tidigare nummer och lägg till den i listan.
            Value value(Method::NUMBER, nummer);
            values.push_back(value);
            nummer = "";

            // Skapa en ny Value.
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
            nummer += c;
        }
    }

    // Avsluta med att skapa en ny Value från tidigare nummer.
    if (nummer != "") {
        Value value(Method::NUMBER, nummer);
        values.push_back(value);
        nummer = "";
    }

    // Kolla om första och sista Value var ett nummer.
    if (values.at(0).getMethod() != Method::NUMBER) {
        return "Error 101"; // Första Value är inte ett nummer.
    }
    if (values.at(values.size()-1).getMethod() != Method::NUMBER){
        return "Error 102"; // Sista Value är inte ett nummer.
    }

    float result = values.at(0).getValue().toFloat(); // Nuvarande resultat
    for (int i = 1; i < values.size()-1; i+=2) {
        Method method = values.at(i).getMethod();

        if (values.at(i+1).getMethod() != Method::NUMBER) {
            return "Error 103"; // Nästa Value är inte ett nummer.
        }

        float nextValue = values.at(i+1).getValue().toFloat(); // Nästa Value
        
        if (method == Method::PLUS) {
            result += nextValue;
        } else if (method == Method::MINUS) {
            result -= nextValue;
        } else if (method == Method::DIVIDE) {
            result /= nextValue;
        } else if (method == Method::MULTIPLY) {
            result *= nextValue;
        } else if (method == Method::POWER) {
            result = pow(v, nextValue);
        }
    }
    return String(result);
}