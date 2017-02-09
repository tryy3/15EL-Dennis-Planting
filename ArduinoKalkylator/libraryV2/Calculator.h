/*
 * Calculator.h - Library for creating calculators.
 * Created by Dennis Planting, 21 January, 2017.
 */
#ifndef Calculator_h
#define Calculator_h

#include "Arduino.h"

class Calculator {
  public:
    Calculator();
    void addChar(char c);
    void addWord(String str);
    void delChar();
    void clear();
    String getBuffer();
    String evaluate();
  private:
    String _buffer;
};

// Baserad på http://stackoverflow.com/a/26227947/2817262
class Eval {
  public:
    Eval(String str);
    double parse();
  private:
    int _pos;
    int _ch;
    String _str;
    void nextChar();
    bool eat(int charToEat);
    double parseExpression();
    double parseTerm();
    double parseFactor();
};

#endif
