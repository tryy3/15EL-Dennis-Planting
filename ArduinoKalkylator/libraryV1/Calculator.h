/*
 * Calculator.h - Library for creating calculators.
 * Created by Dennis Planting, 21 January, 2017.
 */
#ifndef Calculator_h
#define Calculator_h

#include "Arduino.h"

// Method är olika typer av metoder, det är enum för att enkelt
// kunna räkna ut olika typer av metoder.
enum Method{NUMBER, PLUS, MINUS, DIVIDE, MULTIPLY, POWER};

// Value är en klass som håller olika värden, siffror, plus, minus osv.
class Value {
  public:
    Value(Method method, String value);
    Method getMethod();
    String getValue();
  private:
    Method _method;
    String _value;
};

// Calculator är huvud klassen för denna library.
class Calculator {
  public:
    Calculator();
    void addChar(char c);
    void delChar();
    void clear();
    String getBuffer();
    String evaluate();
  private:
    String _buffer;
};

#endif
