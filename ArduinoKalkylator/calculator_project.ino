// Inkludera Calulator library
#include <Calculator.h>

#include <LiquidCrystal.h>
#include <Key.h>
#include <Keypad.h>

// Rader och kolumer från keypad
const byte ROWS = 5;
const byte COLS = 5;

Calculator calc;

/* Functions list
 *  
 *  d = Delete (remove last char)
 *  c = Clear
 *  s = sqrt
 *  i = sin
 *  o = cos
 *  t = tan
 */

char keys[ROWS][COLS]= {
  {'+', '-', '*', '/', '^'},
  {'1', '2', '3', '(', ')'},
  {'4', '5', '6', 'i', 'o'},
  {'7', '8', '9', 't', 's'},
  {'.', '0', '=', 'd', 'c'}
};

byte rowPins[ROWS] = {A2, A3, A4, A5};
byte colPins[COLS] = {2,3,4,5,6};

Keypad keypad = Keypad(makeKeymap(keys), rowPins, colPins, ROWS, COLS);

LiquidCrystal lcd(8,7,9,10,11,12);

void setup() {
  Serial.begin(9600); // Test
  lcd.begin(16,2);
  lcd.clear();
}

void loop() {
  char key = keypad.getKey();

  if (key) {
    if (key == 'i') {
      calc.addWord("sin");
      clr();
    } else if (key == 'o') {
      calc.addWord("cos");
      clr();
    } else if (key == 't') {
      calc.addWord("tan");
      clr();
    } else if (key == 's') {
      calc.addWord("sqrt");
      clr();
    } else if (key == 'c') {
      calc.clear();
      clr();
    }
    else if (key == 'd') {
      calc.delChar();
      clr();
    }
    else if (key == '=') {
      out(1, calc.evaluate());
    } else {
      calc.addChar(key);
      clr();
    }
  }
}

// Clear screen och kör out funktion.
void clr() {
  lcd.clear();
  out(0, calc.getBuffer());
}

// Enkel funktion för att printa ut meddelande till både LCD och Serial.
void out(int line, String s) {
  lcd.setCursor(0,line);
  lcd.print(s);
  Serial.println(s);
}

