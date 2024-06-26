# Set 1
## Challenge 1 - hex to base64

### Hexadecimal (Hex)
Definition:
Hexadecimal is a base-16 number system that uses sixteen symbols to represent values. These symbols are:

The numbers 0 to 9 (representing values 0 to 9).
The letters A to F (or a to f) (representing values 10 to 15).

### Base64
Definition:
Base64 is a binary-to-text encoding scheme that represents `binary data in an ASCII string format` by translating it into a base-64 representation. The Base64 alphabet consists of 64 characters:
* The uppercase letters A-Z.
* The lowercase letters a-z.
* The digits 0-9.
* The symbols + and /.
* An equals sign (=) is used for padding to ensure the encoded string length is a multiple of 4.

Usage:
Base64 is often used to encode data that needs to be stored and transferred over media that are designed to deal with textual data. This includes email via MIME, as well as storing complex data in XML or JSON.

Example:
Consider string "Hello":

1. Convert each character to its ASCII binary equivalent:
    * H is 01001000(2) = 72(10) 
    * e is 01100101(2) = 101(10)
    * l is 01101100(2) - 108(10)
    * l is 01101100(2) - 108(10)
    * o is 01101111(2) - 111(10)
2. Combine these binaries: 0100100001100101011011000110110001101111.
3. Split into 6-bit groups: base  010010 000110 010101 101100 011011 000110 1111_00. (if last chunk is shorter than 6, add "0" to make it length = 6)
4. Convert each 6-bit group to its corresponding Base64 character:
    * 010010(2)=18(10) is S
    * 000110(2)=6(10)  is G
    * 010101(2)=21(10) is V
    * 101100(2)=44(10) is s
    * 011011(2)=27(10) is b
    * 000110(2)=6(10)  is G
    * 111100(2)=60(10) is 8
5. The Base64 encoding of "Hello" is SGVsbG8= (with padding = to ensure the output length is a multiple of 4).
